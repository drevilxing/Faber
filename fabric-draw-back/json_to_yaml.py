import yaml
import json


# 删除配置文件中不需要的空字段和false字段
def delete_null_and_false_r(res):
    for k, v in res.items():
        print(k, v)
        if isinstance(v, dict):
            delete_null_and_false_r(v)
        elif v == '':
            del res[k]
        elif v is False:
            del res[k]
    return res


def delete_double_slash(res):
    for peer in res["entityMatchers"]["peers"]:
        peer["pattern"] = peer["pattern"].replace('\\\\', '\\')
    for order in res["entityMatchers"]["orderer"]:
        order["pattern"] = order["pattern"].replace('\\\\', '\\')
    return res


# 将key字段中的域名改写为字典的键值
def manage_key(res):
    # channels部分
    channels = dict()
    for channel in res["channels"]:
        peers = dict()
        for peer in channel['peers']:
            peers[peer["key"]] = peer
            del peers[peer["key"]]["key"]
        channel["peers"] = peers

        channels[channel["name"]] = channel
        del channels[channel["name"]]["name"]
    res["channels"] = channels

    # organizations部分
    organizations = dict()
    for organization in res["organizations"]:
        organizations[organization["key"]] = organization
        del organizations[organization["key"]]["key"]
    res["organizations"] = organizations

    # orderers部分
    orderers = dict()
    for orderer in res["orderers"]:
        orderers[orderer["key"]] = orderer
        del orderers[orderer["key"]]["key"]
    res["orderers"] = orderers

    # peers部分
    peers = dict()
    for peer in res["peers"]:
        peers[peer["key"]] = peer
        del peers[peer["key"]]["key"]
    res["peers"] = peers

    # certificateAuthorities部分
    cas = dict()
    for ca in res["certificateAuthorities"]:
        cas[ca["key"]] = ca
        del cas[ca["key"]]["key"]
    res["certificateAuthorities"] = cas
    return res


# 将go语言生成的json配置文件转为yaml配置文件
def json_to_yaml(path):
    res = {}
    with open(path) as js:
        res = json.load(js)
    # 删除双斜线
    res = delete_double_slash(res)
    # 修改列表为字典
    res = manage_key(res)
    # 删除无用的空字段和false字段
    # res = delete_null_and_false_r(res)
    # 写入文件
    with open('sdkConfig.yaml', 'w', encoding="utf-8") as file:
        yaml.dump(res, file, Dumper=yaml.Dumper)
    with open('../faberGoSDK/src/sdkConfig.yaml', 'w', encoding="utf-8") as file:
        yaml.dump(res, file, Dumper=yaml.Dumper)


json_to_yaml("../sdkConfig.json")
