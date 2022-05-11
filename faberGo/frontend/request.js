import axios from "axios";

const server = "localhost"
const port = 9000
const base = "/faber"

function backend() {
    return server + ":" + port + base
}

function header() {
    return {
        "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
    };
}

function get(path, params) {
    return new Promise(((resolve, reject) => {
        axios({
            headers: header(),
            method: "get",
            url: path,
            params: params,
        }).then(res => {
            resolve(res)
        }).catch(err => {
            reject(err)
        })
    }));
}

function post(path, data) {
    const inside = new URLSearchParams();
    Object.keys(data).forEach((key) => {
        inside.append(key, data[key]);
    });
    return new Promise(((resolve, reject) => {
        axios({
            headers: header(),
            method: "post",
            url: path,
            data: inside,
        }).then(res => {
            resolve(res)
        }).catch(err => {
            reject(err)
        })
    }));
}

// 完成区块链网络的创建和删除工作。

function networkCreate(name) {
    return post(backend() + "/network/create", { name: name });
}

function networkList() {
    return get(backend() + "/network/list", {});
}

function networkDelete(name) {
    return post(backend() + "/network/delete", { name: name });
}

function networkOpen(name) {
    return post(backend() + "/network/open", { name: name })
}

// 区块链网络上组织、节点、通道管理

function blockchainCreate(name, key) {
    return post(backend() + "/blockchain/create", { name: name, key: key })
}

function blockchainChannelCreate(name, blockchain) {
    return post(backend() + "/blockchain/channel/create", { name: name, blockchain: blockchain })
}

function blockchainOrganizationCreate(name, blockchain) {
    return post(backend() + "/blockchain/org/create", { name: name, blockchain: blockchain })
}

function blockchainOrganizationJoinChannel(org, channel) {
    return post(backend() + "/blockchain/org/join/channel", { org: org, channel: channel })
}

function blockchainNodeCreate(name, org, blockchain, host, ssh, fabricPort, user, pwd, key, bootstrap, type) {
    return post(backend() + "/blockchain/node/create", {
        name, org, blockchain, host, ssh, fabricPort, user, pwd, key, bootstrap, type
    })
}

// 区块链配置文件保存、环境安装、区块链网络部署

function configSave() {
    return post(backend() + "/environment/config/save", {})
}

function configFetch() {
    return post(backend() + "/environment/config/fetch", {})
}

function configGenerate() {
    return post(backend() + "/environment/config/generate", {})
}

function install() {
    return post(backend() + "/environment/install", {})
}

function deploy() {
    return post(backend() + "/environment/deploy", {})
}
