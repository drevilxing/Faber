import yaml
import re
import os
import json
from ruamel.yaml import YAML


class YamlGenerator:
    def __init__(self):
        pass

    def generate(self, **kwargs):
        pass


class ConfigTXYamlGenerator:
    def __init__(self, net_name: str, crypto_base: str):
        self.yml = YAML()
        self.yml.indent(mapping=4, sequence=6, offset=4)
        self.crypto_base = crypto_base
        self.net_name = net_name
        self.configtx = None
        self.filename = ""

    def update_organizations(self, groups, nodes):
        Organizations = []
        for group in groups:
            elems = group['key'].split(".")
            name = elems[0].capitalize()
            Organization = {
                "Name": name,
                "SkipAsForeign": False,
                "ID": name + "MSP",
                "MSPDir": self.crypto_base + "/organizations/" + group['key'] + "/msp"
            }
            if "order" in group['key']:
                Organization["Policies"] = self.configtx["Organizations"][0]["Policies"]
                Organization["OrdererEndpoints"] = group["nodes"]["orderer"]
            else:
                Organization["Policies"] = self.configtx["Organizations"][1]["Policies"]
            Organization["AnchorPeers"] = []
            for url in group["nodes"]["anchor_peers"]:
                for node in nodes:
                    if node['key'] == url:
                        Organization["AnchorPeers"].append({
                            "Host": url,
                            "Port": int(node["address"]["fabric_port"])
                        })
            for Policie in Organization["Policies"]:
                Rule = Organization["Policies"][Policie]["Rule"]
                Organization["Policies"][Policie]["Rule"] = re.sub("Or\S+MSP", name, Rule)
            Organizations.append(Organization)
        self.configtx["Organizations"] = Organizations

    def update_orderer(self, orderers: dict):
        Orderer = self.configtx["Orderer"]
        Addresses = []
        Consenters = []
        for orderer in orderers:
            name = orderer.split(".")[0]
            Addresses.append(orderers[orderer]["address"]["host"] + ":" + orderers[orderer]["address"]["fabric_port"])
            Consenters.append({
                "Host": orderer,
                "Port": int(orderers[orderer]["address"]["fabric_port"]),
                "ClientTLSCert": self.crypto_base + "/organizations/" + ".".join(
                    orderer.split(".")[-3:]) + "/peers/" + orderer + "/tls/server.crt",
                "ServerTLSCert": self.crypto_base + "/organizations/" + ".".join(
                    orderer.split(".")[-3:]) + "/peers/" + orderer + "/tls/server.crt"
            })
        Orderer["Addresses"] = Addresses
        Orderer["EtcdRaft"]["Consenters"] = Consenters

    def update_profiles(self):
        OrdererGenesis = self.net_name + "OrdererGenesis"
        Channel = self.net_name + "Channel"
        Profiles = {
            OrdererGenesis: self.configtx["Profiles"]["TwoOrgsOrdererGenesis"],
            Channel: self.configtx["Profiles"]["TwoOrgsChannel"]
        }
        Profiles[OrdererGenesis]["Consortiums"]["SampleConsortium"]["Organizations"] = []
        Profiles[OrdererGenesis]["Orderer"]["Organizations"] = []
        Profiles[Channel]["Application"]["Organizations"] = []
        for org in self.configtx["Organizations"]:
            if "rder" not in org["Name"]:
                Profiles[OrdererGenesis]["Consortiums"]["SampleConsortium"]["Organizations"].append(org)
                Profiles[Channel]["Application"]["Organizations"].append(org)
            else:
                Profiles[OrdererGenesis]["Orderer"]["Organizations"].append(org)
        self.configtx["Profiles"] = Profiles

    def get_filename(self):
        return self.filename

    def input_from(self, filename: str):
        with open(filename) as file:
            self.configtx = self.yml.load(file)
        return self

    def output_to(self, filename: str):
        if not self.configtx:
            return None
        self.filename = filename
        with open(filename, 'w', encoding="utf-8") as file:
            self.yml.dump(self.configtx, file)
        return self

    def generate(self, groups, nodes, orderers):
        if not self.configtx:
            return None
        self.update_organizations(groups, nodes)
        self.update_orderer(orderers)
        self.update_profiles()
        return self


class CAYamlGenerator(YamlGenerator):
    def __init__(self):
        super().__init__()

    def generate(self, node_id, org_name, fabric_name, fabric_port, crypto_path):
        with open('template/docker-compose-ca-template.yaml') as file:
            docker_yaml = yaml.load(file, Loader=yaml.Loader)
        docker_yaml['networks']['net']['external']['name'] = fabric_name
        ca_information = docker_yaml['services']['ca.org1.test.com']
        ca_information['environment'][1] = f'FABRIC_CA_SERVER_CA_NAME=ca-{org_name}'
        ca_information['environment'][3] = f'FABRIC_CA_SERVER_PORT={fabric_port}'
        ca_information['environment'][4] = f'FABRIC_CA_SERVER_CSR_HOSTS=localhost, {node_id}'
        ca_information['ports'][0] = f'{fabric_port}:{fabric_port}'
        ca_information['volumes'][0] = f'/root/opt/organizations/fabric-ca/{org_name}:/etc/hyperledger/fabric-ca-server'
        ca_information['container_name'] = node_id
        del docker_yaml['services']['ca.org1.test.com']
        docker_yaml['services'][node_id] = ca_information
        file_name = f'docker-compose-ca-{org_name}.yaml'
        with open(file_name, 'w', encoding="utf-8") as file:
            yaml.dump(docker_yaml, file, Dumper=yaml.Dumper)
        return file_name


class OrderYamlGenerator(YamlGenerator):
    def __init__(self):
        super().__init__()

    def generate(self, node_id, org_name, node_name, fabric_name, fabric_port, crypto_path):
        with open('template/docker-compose-orderer-template.yaml') as file:
            docker_yaml = yaml.load(file, Loader=yaml.Loader)
        docker_yaml['networks']['net']['external']['name'] = fabric_name
        order_information = docker_yaml['services']['orderer0.orderer.test.com']
        order_information['container_name'] = node_id
        order_information['environment'][2] = f'ORDERER_GENERAL_LISTENPORT={fabric_port}'
        order_information['environment'][5] = f'ORDERER_GENERAL_LOCALMSPID={org_name.capitalize()}MSP'
        order_information['volumes'][
            1] = f'{crypto_path}/organizations/orderer.test.com/peers/{node_id}/msp:/var/hyperledger/orderer/msp'
        order_information['volumes'][
            2] = f'{crypto_path}/organizations/orderer.test.com/peers/{node_id}/tls/:/var/hyperledger/orderer/tls'
        order_information['volumes'][3] = f'{node_id}:/var/hyperledger/production/orderer'
        order_information['ports'][0] = f'{fabric_port}:{fabric_port}'
        del docker_yaml['services']['orderer0.orderer.test.com']
        docker_yaml['services'][node_id] = order_information
        file_name = f'docker-compose-{org_name}-{node_name}.yaml'
        with open(file_name, 'w', encoding="utf-8") as file:
            yaml.dump(docker_yaml, file, Dumper=yaml.Dumper)
        with open(file_name, 'a') as file:
            file.write(f'volumes:\n  {node_id}:\n')
        return file_name


class PeerYamlGenerator(YamlGenerator):
    def __init__(self):
        super().__init__()

    def generate(self, node_id, fabric_name, fabric_port, crypto_path):
        node_name, org_name, domain = node_id.split('.', 2)
        with open('template/docker-compose-peer-template.yaml') as file:
            docker_yaml = yaml.load(file, Loader=yaml.Loader)
        docker_yaml['networks']['net']['external']['name'] = fabric_name
        peer_information = docker_yaml['services']['peer0.org1.test.com']
        peer_information['container_name'] = node_id
        peer_information['environment'][8] = f'CORE_PEER_ID={node_id}'
        peer_information['environment'][9] = f'CORE_PEER_ADDRESS={node_id}:{fabric_port}'
        peer_information['environment'][11] = f'CORE_PEER_CHAINCODEADDRESS={node_id}:7052'
        peer_information['environment'][14] = f'CORE_PEER_GOSSIP_EXTERNALENDPOINT={node_id}:{fabric_port}'
        peer_information['environment'][15] = f'CORE_PEER_LOCALMSPID={org_name.capitalize()}MSP'
        peer_information['volumes'][
            1] = f'{crypto_path}/organizations/{org_name}.{domain}/peers/{node_id}/msp:/etc/hyperledger/fabric/msp'
        peer_information['volumes'][
            2] = f'{crypto_path}/organizations/{org_name}.{domain}/peers/{node_id}/tls:/etc/hyperledger/fabric/tls'
        peer_information['volumes'][3] = f'{node_id}:/var/hyperledger/production'
        peer_information['ports'][0] = f'{fabric_port}:{fabric_port}'
        del docker_yaml['services']['peer0.org1.test.com']
        docker_yaml['services'][node_id] = peer_information
        file_name = f'docker-compose-{org_name}-{node_name}.yaml'
        with open(file_name, 'w', encoding="utf-8") as file:
            yaml.dump(docker_yaml, file, Dumper=yaml.Dumper)
        with open(file_name, 'a') as file:
            file.write(f'volumes:\n  {node_id}:\n')
        return file_name


def generate_ca(ca_id, ca_information, fabric_name, target_host, crypto_base):
    node_name, group_name, domain = ca_id.split('.', 2)
    address = ca_information['address']
    ca_yaml_generator = CAYamlGenerator()
    file_name = ca_yaml_generator.generate(ca_id, group_name, fabric_name, address['fabric_port'], crypto_base)
    return file_name


def generate_peer(peer_id, peer_information, order_group_id, fabric_name, target_host, ca_port, crypto_base):
    node_name, group_name, domain = peer_id.split('.', 2)
    address = peer_information['address']
    peer_yaml_generator = PeerYamlGenerator()
    file_name = peer_yaml_generator.generate(peer_id, fabric_name, address['fabric_port'], crypto_base)
    return file_name


def generate_configtx(groups: dict, nodes: dict, orderers: dict, fabric_name: str, crypto_base: str):
    configtx = ConfigTXYamlGenerator(fabric_name, crypto_base)
    # 读取yaml文件
    # 生成groups、nodes、orderers
    # 输出至configtx.yaml文件
    # 返回文件名称
    return configtx.input_from("./template/configtx.yaml") \
        .generate(groups, nodes, orderers) \
        .output_to(f"configtx.yaml") \
        .get_filename()


def generate_order(order_id, order_information, fabric_name, channel_id, peer_group_ids, configtx_filename: str, crypto_base='/root/opt'):
    node_name, group_name, domain = order_id.split('.', 2)
    address = order_information['address']
    orderer_yaml_generator = OrderYamlGenerator()
    filename = orderer_yaml_generator.generate(order_id, group_name, node_name, fabric_name, address["fabric_port"], crypto_base)


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
                #peer_group_ids.append(group['key'])
            peer_group_ids.append(group['key'])
            # 生成ca证书

            for node in network_topology_json['nodes']:
                if node['key'] == group['nodes']['ca']:
                    generate_ca(group['nodes']['ca'], node, blockchain['name'], target_host, crypto_path)
        print("成功生成ca配置文件")
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
                        generate_peer(peer_id, node, order_group_id, blockchain['name'], target_host, peer_ca_port, crypto_path)
        orderers = dict()
        print("成功生成peer节点配置文件")

        for node in network_topology_json["nodes"]:
            if "orderer" in node["type"]:
                orderers[node['key']] = node

        # 生成configtx文件
        configtx_filename = generate_configtx(network_topology_json["groups"], network_topology_json["nodes"], orderers, blockchain["name"], crypto_path)
        print("成功生成configtx通道配置文件")

        for group in network_topology_json['groups']:
            if group['key'] == order_group_id:
                for order_id in group['nodes']['orderer']:
                    for node in network_topology_json['nodes']:
                        if node['key'] == order_id:
                            generate_order(order_id, node, blockchain['name'], blockchain['channels'][0], peer_group_ids, configtx_filename, crypto_path)
        print("成功生成orderer节点配置文件")


if __name__ == '__main__':
    json_file = 'config.json'
    with open(json_file) as js:
        network_json = json.load(js)
    parse_json(network_json)
    print('该模块运行完毕')
