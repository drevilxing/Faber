import paramiko
import time
import io
import os
import stat
import json

from yaml_generator import CAYamlGenerator, OrderYamlGenerator, PeerYamlGenerator, ConfigTXYamlGenerator


def generate_ca(ca_id, ca_information, fabric_name, target_host, crypto_base):
    # 分割成ca org0 test.com
    node_name, group_name, domain = ca_id.split('.', 2)
    # 读取ca地址
    address = ca_information['address']
    ca_yaml_generator = CAYamlGenerator()
    # 生成文件
    ca_yaml_generator.generate(ca_id, group_name, fabric_name, address['fabric_port'], crypto_base)


def generate_peer(peer_id, peer_information, order_group_id, fabric_name, target_host, ca_port, crypto_base):
    address = peer_information['address']
    peer_yaml_generator = PeerYamlGenerator()
    peer_yaml_generator.generate(peer_id, fabric_name, address['fabric_port'], crypto_base)


def generate_order(order_id, order_information, fabric_name, channel_id, peer_group_ids, configtx_filename: str,
                   crypto_base='/root/opt'):
    node_name, group_name, domain = order_id.split('.', 2)
    address = order_information['address']
    orderer_yaml_generator = OrderYamlGenerator()
    orderer_yaml_generator.generate(order_id, group_name, node_name, fabric_name, address["fabric_port"],
                                    crypto_base)


def generate_configtx(groups: dict, nodes: dict, orderers: dict, net_name: str, crypto_base: str):
    configtx = ConfigTXYamlGenerator(net_name, crypto_base)
    # 读取yaml文件
    # 生成groups、nodes、orderers
    # 输出至configtx.yaml文件
    # 返回文件名称
    return configtx.input_from("./template/configtx.yaml") \
        .generate(groups, nodes, orderers) \
        .output_to("./configtx.yaml") \
        .get_filename()


def parse_json(network_topology_json):
    order_group_id = ''
    order_ca_port = ''
    target_host = ''
    peer_group_ids = []

    # 读取group信息
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
        else:
            # 添加peer节点
            peer_group_ids.append(group['key'])
        # 生成ca证书
        for node in network_topology_json['nodes']:
            if node['key'] == group['nodes']['ca']:
                generate_ca(group['nodes']['ca'], node,
                    network_topology_json['blockchains'][0]['name'], target_host, '/root/opt')
    print("成功生成ca证书")

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
                    generate_peer(peer_id, node, order_group_id,
                                  network_topology_json['blockchains'][0]['name'], target_host, peer_ca_port,
                                  '/root/opt')
    orderers = dict()
    print("成功生成peer")

    for node in network_topology_json["nodes"]:
        if "orderer" in node["type"]:
            orderers[node['key']] = node
    # 生成configtx文件
    configtx_filename = generate_configtx(network_topology_json["groups"], network_topology_json["nodes"], orderers,
                                          network_topology_json["blockchains"][0]["name"], "/root/opt")
    print("成功生成configtx")


    for group in network_topology_json['groups']:
        if group['key'] == order_group_id:
            for order_id in group['nodes']['orderer']:
                for node in network_topology_json['nodes']:
                    if node['key'] == order_id:
                        generate_order(order_id, node,
                                       network_topology_json['blockchains'][0]['name'],
                                       network_topology_json['blockchains'][0]['channels'][0], peer_group_ids,
                                       configtx_filename)
    print("成功生成order")



if __name__ == '__main__':
    network_json = {
        "groups": {
            "orderer.test.com": {
                "nodes": {
                    "ca": "ca.orderer.test.com",
                    "orderer": ["orderer0.orderer.test.com", "orderer1.orderer.test.com", "orderer2.orderer.test.com"]
                },
                "blockchains": "fabric-1"
            },
            "org0.test.com": {
                "nodes": {
                    "ca": "ca.org0.test.com",
                    "leader_peers": ["peer0.org0.test.com"],
                    "anchor_peers": ["peer0.org0.test.com"],
                    "committing_peers": ["peer0.org0.test.com"],
                    "endorsing_peers": ["peer0.org0.test.com"]
                },
                "blockchains": "fabric-1",
                "channel": ["channel-1"]
            },
            "org1.test.com": {
                "nodes": {
                    "ca": "ca.org1.test.com",
                    "leader_peers": ["peer0.org1.test.com"],
                    "anchor_peers": ["peer0.org1.test.com"],
                    "committing_peers": ["peer0.org1.test.com"],
                    "endorsing_peers": ["peer0.org1.test.com"]
                },
                "blockchains": "fabric-1",
                "channel": ["channel-1"]
            },
            "org2.test.com": {
                "nodes": {
                    "ca": "ca.org2.test.com",
                    "leader_peers": ["peer0.org2.test.com"],
                    "anchor_peers": ["peer0.org2.test.com"],
                    "committing_peers": ["peer0.org2.test.com"],
                    "endorsing_peers": ["peer0.org2.test.com"]
                },
                "blockchains": "fabric-1",
                "channel": ["channel-1"]
            }
        },
        "nodes": {
            "ca.orderer.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "7054", },
                "type": ["ca"]
            },
            "orderer0.orderer.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "7050", },
                "type": ["orderer"]
            },
            "orderer1.orderer.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "8050", },
                "type": ["orderer"]
            },
            "orderer2.orderer.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "9050"},
                "type": ["orderer"]
            },
            "ca.org0.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "9054"},
                "type": ["ca"]
            },
            "peer0.org0.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "8051"},
                "bootstrap": ["127.0.0.1:7051"],
                "type": ["leader_peer", "anchor_peer", "committing_peer", "endorsing_peers"]
            },
            "ca.org1.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "10054"},
                "type": ["ca"]
            },
            "peer0.org1.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "11051"},
                "bootstrap": ["127.0.0.1:7051"],
                "type": ["leader_peer", "anchor_peer", "committing_peer", "endorsing_peers"]
            },
            "ca.org2.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "12054"},
                "type": ["ca"]
            },
            "peer0.org2.test.com": {
                "address": {"host": "172.20.10.3", "ssh_port": "22", "fabric_port": "13051"},
                "bootstrap": ["127.0.0.1:7051"],
                "type": ["leader_peer", "anchor_peer", "committing_peer", "endorsing_peers"]
            },
        },
        "blockchains": {
            "fabric-1": {
                "name": "FabricDraw",
                "channels": ["channel-1"]
            }
        }
    }
    json_file = 'config.json'
    with open(json_file) as js:
        network_json = json.load(js)
    parse_json(network_json)
    print('ok')
