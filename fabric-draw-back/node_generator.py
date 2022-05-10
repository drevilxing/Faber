import json
import paramiko
import time
import os
import stat


# 递归获取文件
def sftp_get_r(sftp_client, remote_path, local_path):
    try:
        sftp_client.stat(remote_path)
    except IOError:
        return
    if not os.path.exists(local_path):
        os.mkdir(local_path)
    for item in sftp_client.listdir(remote_path):
        if stat.S_ISDIR(sftp_client.stat(f'{remote_path}/{item}').st_mode):
            sftp_get_r(sftp_client, f'{remote_path}/{item}', os.path.join(local_path, item))
        else:
            sftp_client.get(f'{remote_path}/{item}', os.path.join(local_path, item))


# 递归传输文件
def sftp_put_r(sftp_client, local_path, remote_path):
    if not os.path.exists(local_path):
        return
    path = ""
    for dir in remote_path.split("/"):
        if dir == "":
            continue
        path += f"/{dir}"
        try:
            sftp_client.listdir(path)
        except IOError:
            sftp_client.mkdir(path)
    try:
        sftp_client.stat(remote_path)
    except IOError:
        sftp_client.mkdir(remote_path)
    for item in os.listdir(local_path):
        if os.path.isfile(os.path.join(local_path, item)):
            sftp_client.put(os.path.join(local_path, item), f'{remote_path}/{item}')
        else:
            sftp_put_r(sftp_client, os.path.join(local_path, item), f'{remote_path}/{item}')


def start_ca(ca_id, ca_information, fabric_name, target_host, crypto_base):
    node_name, group_name, domain = ca_id.split('.', 2)
    address = ca_information['address']
    # 连接服务器
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(hostname=address['host'], port=address['ssh_port'], username='root', password='linux')
    stdin, stdout, stderr = ssh.exec_command(f'if [ ! -d {crypto_base} ]; then mkdir -p {crypto_base}; fi')
    stdout.channel.recv_exit_status()
    ftp_client = ssh.open_sftp()
    # 搭建节点
    file_name = 'node_build.py'
    ftp_client.put(file_name, f'{crypto_base}/{file_name}')
    if group_name == 'orderer':
        # 初始化docker swarm
        stdin, stdout, stderr = ssh.exec_command(f'python3 {crypto_base}/node_build.py --func_name init_docker_swarm {target_host} {fabric_name} {crypto_base}')
        stdout.channel.recv_exit_status()
        ftp_client.get(f'{crypto_base}/token', 'token')
    else:
        try:
            ftp_client.stat(f'{crypto_base}/token')
        except IOError:
            node_host = address['host']
            ftp_client.put('token', f'{crypto_base}/token')
            # 加入docker swarm
            stdin, stdout, stderr = ssh.exec_command(f'python3 {crypto_base}/node_build.py --func_name join_docker_swarm {node_host} {target_host} {crypto_base}')
            stdout.channel.recv_exit_status()
    # 获取ca配置文件
    file_name = f'docker-compose-ca-{group_name}.yaml'
    # 将配置文件传输至服务器
    ftp_client.put(file_name, f'{crypto_base}/{file_name}')
    # 启动ca容器
    stdin, stdout, stderr = ssh.exec_command(f'docker-compose -f {crypto_base}/{file_name} up -d')
    stdout.channel.recv_exit_status()
    time.sleep(4)
    # 获取服务器的证书文件
    tls_cert_path = f'organizations/fabric-ca/{group_name}'
    if not os.path.exists(tls_cert_path):
        os.makedirs(tls_cert_path)
    ftp_client.get(f'{crypto_base}/{tls_cert_path}/tls-cert.pem', f'{tls_cert_path}/tls-cert.pem')
    ftp_client.close()
    ssh.close()


def start_order_msp(order_id, order_information, ca_port, crypto_base):
    node_name, group_name, domain = order_id.split('.', 2)
    # 连接服务器
    address = order_information['address']
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(hostname=address['host'], port=address['ssh_port'], username='root', password='linux')
    tls_cert_path = f'organizations/fabric-ca/{group_name}'
    # 创建证书文件夹
    stdin, stdout, stderr = ssh.exec_command(f'if [ ! -d {crypto_base}/{tls_cert_path} ]; then mkdir -p {crypto_base}/{tls_cert_path}; fi')
    stdout.channel.recv_exit_status()
    ftp_client = ssh.open_sftp()
    # 上传证书文件
    ftp_client.put(f'{tls_cert_path}/tls-cert.pem', f'{crypto_base}/{tls_cert_path}/tls-cert.pem')
    file_name = 'node_build.py'
    ftp_client.put(file_name, f'{crypto_base}/{file_name}')
    # 生成组织msp
    stdin, stdout, stderr = ssh.exec_command(f'python3 {crypto_base}/node_build.py --func_name org_msp_generate {group_name} {domain} {ca_port} {crypto_base}')
    stdout.channel.recv_exit_status()
    # print(stdout.readlines(), stderr.readlines())
    stdin, stdout, stderr = ssh.exec_command(f'python3 {crypto_base}/node_build.py --func_name peer_msp_generate {node_name} {group_name} {domain} {ca_port} {crypto_base}')
    stdout.channel.recv_exit_status()
    # print(stdout.readlines(), stderr.readlines())
    tls_ca_path = f'organizations/{group_name}.{domain}/tlsca'
    if not os.path.exists(tls_ca_path):
        os.makedirs(tls_ca_path)
    # 获取msp文件
    ftp_client.get(f'{crypto_base}/{tls_ca_path}/tlsca.{group_name}.{domain}-cert.pem', f'{tls_ca_path}/tlsca.{group_name}.{domain}-cert.pem')
    server_path = f'organizations/{group_name}.{domain}/peers/{order_id}/tls'
    if not os.path.exists(server_path):
        os.makedirs(server_path)
    ftp_client.get(f'{crypto_base}/{server_path}/server.crt', f'{server_path}/server.crt')
    ftp_client.close()


def start_peer(peer_id, peer_information, order_group_id, fabric_name, target_host, ca_port, crypto_base):
    node_name, group_name, domain = peer_id.split('.', 2)
    address = peer_information['address']
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(hostname=address['host'], port=address['ssh_port'], username='root', password='linux')
    tls_cert_path = f'organizations/fabric-ca/{group_name}'
    stdin, stdout, stderr = ssh.exec_command(f'if [ ! -d {crypto_base}/{tls_cert_path} ]; then mkdir -p {crypto_base}/{tls_cert_path}; fi')
    stdout.channel.recv_exit_status()
    ftp_client = ssh.open_sftp()
    # 上传peer节点证书文件
    ftp_client.put(f'{tls_cert_path}/tls-cert.pem', f'{crypto_base}/{tls_cert_path}/tls-cert.pem')
    tls_ca_path = f'organizations/{order_group_id}/tlsca'
    stdin, stdout, stderr = ssh.exec_command(f'if [ ! -d {crypto_base}/{tls_ca_path} ]; then mkdir -p {crypto_base}/{tls_ca_path}; fi')
    stdout.channel.recv_exit_status()
    ftp_client.put(f'{tls_ca_path}/tlsca.{order_group_id}-cert.pem', f'{crypto_base}/{tls_ca_path}/tlsca.{order_group_id}-cert.pem')
    # 构建节点
    file_name = 'node_build.py'
    ftp_client.put(file_name, f'{crypto_base}/{file_name}')
    try:
        ftp_client.stat(f'{crypto_base}/token')
    except IOError:
        node_host = address['host']
        ftp_client.put('token', f'{crypto_base}/token')
        stdin, stdout, stderr = ssh.exec_command(f'python {crypto_base}/node_build.py --func_name join_docker_swarm {node_host} {target_host} {crypto_base}')
        stdout.channel.recv_exit_status()

    # 获取配置文件并上传
    node_name, org_name, domain = peer_id.split('.', 2)
    file_name = f'docker-compose-{org_name}-{node_name}.yaml'
    ftp_client.put(file_name, f'{crypto_base}/{file_name}')
    # 构建节点
    stdin, stdout, stderr = ssh.exec_command(f'python3 {crypto_base}/node_build.py --func_name org_msp_generate {group_name} {domain} {ca_port} {crypto_base}')
    stdout.channel.recv_exit_status()
    print(stderr.readlines())
    stdin, stdout, stderr = ssh.exec_command(f'python3 {crypto_base}/node_build.py --func_name peer_msp_generate {node_name} {group_name} {domain} {ca_port} {crypto_base}')
    stdout.channel.recv_exit_status()
    print(stderr.readlines())
    # 启动容器
    stdin, stdout, stderr = ssh.exec_command(f'docker-compose -f {crypto_base}/{file_name} up -d')
    stdout.channel.recv_exit_status()
    print(stderr.readlines())
    time.sleep(3)
    peer_path = f'organizations/{group_name}.{domain}'
    if not os.path.exists(peer_path):
        os.makedirs(peer_path)
    sftp_get_r(ftp_client, f'{crypto_base}/{peer_path}', peer_path)
    ftp_client.close()


def start_order(order_id, order_information, fabric_name, channel_id, peer_group_ids, configtx_filename: str, crypto_base='/root/opt'):
    node_name, group_name, domain = order_id.split('.', 2)
    address = order_information['address']
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(hostname=address['host'], port=address['ssh_port'], username='root', password='linux')
    ssh.exec_command(f'if [ ! -d {crypto_base}/channel-artifacts ]; then mkdir -p {crypto_base}/channel-artifacts; fi')
    ftp_client = ssh.open_sftp()
    sftp_put_r(ftp_client, f"organizations/{group_name}.{domain}/peers", f"{crypto_base}/organizations/{group_name}.{domain}/peers")
    for peer in peer_group_ids:
        sftp_put_r(ftp_client, f"organizations/{peer}/msp/cacerts", f"{crypto_base}/organizations/{peer}/msp/cacerts")
        ftp_client.put(f"organizations/{peer}/msp/config.yaml", f"{crypto_base}/organizations/{peer}/msp/config.yaml")

    file_name = f'docker-compose-{group_name}-{node_name}.yaml'
    ftp_client.put(file_name, f"{crypto_base}/{file_name}")
    ftp_client.put(configtx_filename, f'{crypto_base}/{configtx_filename}')
    while True:
        try:
            ftp_client.stat(f'{crypto_base}/{configtx_filename}')
            ftp_client.stat(f'{crypto_base}/{file_name}')
            print("File exists.")
            break
        except IOError:
            print("File not exists.")
            time.sleep(2)
    ftp_client.close()
    stdin, stdout, stderr = ssh.exec_command(f'python3 {crypto_base}/node_build.py --func_name init_channel_artifacts {fabric_name} {channel_id} "{crypto_base}" {peer_group_ids} ')
    stdout.channel.recv_exit_status()
    print(stderr.readlines())
    stdin, stdout, stderr = ssh.exec_command(f'docker-compose -f {crypto_base}/{file_name} up -d')
    stdout.channel.recv_exit_status()
    print(stderr.readlines())
    time.sleep(4)
    ssh.close()


def parse_json(network_topology_json):
    order_group_id = ''
    order_ca_port = ''
    target_host = ''
    peer_group_ids = []
    crypto_path = "/root/opt"
    # 读取group信息
    for blockchain in network_topology_json['blockchains']:
        for group in network_topology_json['groups']:
            # orderer节点
            if group['key'].split('.', 1)[0] == 'orderer':
                # 获取orderer节点groupid
                order_group_id = group['key']
                # 填入ca端口号
                for node in network_topology_json['nodes']:
                    if node['key'] == group['nodes']['ca']:
                        order_ca_port = node['address']['fabric_port']
                # 填入ca的ip地址
                for node in network_topology_json['nodes']:
                    if node['key'] == group['nodes']['ca']:
                        target_host = node['address']['host']
            #else:
                # 添加peer节点
            peer_group_ids.append(group['key'])
            # 生成ca证书

            for node in network_topology_json['nodes']:
                if node['key'] == group['nodes']['ca']:
                    start_ca(group['nodes']['ca'], node, blockchain['name'], target_host, crypto_path)
        print("成功启动ca节点")

        # 生成每个组织的msp
        for group in network_topology_json['groups']:
            if group['key'] == order_group_id:
                for order_id in group ['nodes']['orderer']:
                    for node in  network_topology_json['nodes']:
                        if node['key'] == order_id:
                            start_order_msp(order_id, node, order_ca_port, crypto_path)

        # 对每个peer节点
        for org_id in peer_group_ids:
            # 获取peer结点的信息
            for node in network_topology_json['nodes']:
                for group in network_topology_json['groups']:
                    if group['key'] == org_id:
                        if node['key'] == group['nodes']['ca']:
                            peer_ca_port = node['address']['fabric_port']

            for group in network_topology_json['groups']:
                if group['key'] == org_id:
                    leader_peers_ids = group['nodes']['leader_peers']
                    anchor_peers_ids = group['nodes']['anchor_peers']
                    committing_peers_ids = group['nodes']['committing_peers']
                    endorsing_peers_ids = group['nodes']['endorsing_peers']

            peer_ids = list(set(leader_peers_ids).union(
                set(anchor_peers_ids).union(set(committing_peers_ids)).union(set(endorsing_peers_ids))))
            # 生成peer节点

            for peer_id in peer_ids:
                for node in network_topology_json['nodes']:
                    if peer_id == node['key']:
                        start_peer(peer_id, node, order_group_id, blockchain['name'], target_host, peer_ca_port, crypto_path)
        orderers = dict()
        print("成功启动peer节点")

        for node in network_topology_json["nodes"]:
            if "orderer" in node["type"]:
                orderers[node['key']] = node

        # 获取configtx文件
        configtx_filename = 'configtx.yaml'

        for group in network_topology_json['groups']:
            if group['key'] == order_group_id:
                for order_id in group['nodes']['orderer']:
                    for node in network_topology_json['nodes']:
                        if node['key'] == order_id:
                            start_order(order_id, node, blockchain['name'], blockchain['channels'][0], peer_group_ids, configtx_filename, crypto_path)
        print("成功启动orderer节点")


if __name__ == '__main__':
    json_file = 'config.json'
    with open(json_file) as js:
        network_json = json.load(js)
    parse_json(network_json)
    print('该模块运行完毕')
