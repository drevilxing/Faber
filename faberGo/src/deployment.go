package main

import (
	"encoding/json"
	"faberGo/pkg/config"
	"faberGo/pkg/config/blockchain"
	"faberGo/pkg/config/nodes"
	"faberGo/pkg/config/org"
	"os"
)

const fabric = "fabric-1"
const fabricDraw = "FabricDraw"
const channel1 = "channel-1"

// org

const orgOrderer = "orderer.test.com"
const org0 = "org0.test.com"
const org1 = "org1.test.com"
const org2 = "org2.test.com"

func GenerateDeploymentConfigExample(path string) error {
	generateConfig := config.GenGenerateConfig()

	// Blockchain配置
	generateConfig.AddBlockchain(&blockchain.Blockchain{
		Key:      fabric,
		Name:     fabricDraw,
		Channels: &[]string{channel1},
	})

	// 组织配置
	generateConfig.AddGroup(org.GenOrg(fabric, orgOrderer))
	generateConfig.AddGroup(org.GenOrg(fabric, org0))
	generateConfig.AddGroup(org.GenOrg(fabric, org1))
	generateConfig.AddGroup(org.GenOrg(fabric, org2))

	// 组织节点添加
	generateConfig.AddNode(&nodes.Node{
		Key: "ca",
		Org: orgOrderer,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "7054",
		},
		Bootstrap: &[]string{},
		Type:      &[]string{org.NodeCA},
	})
	generateConfig.AddNode(&nodes.Node{
		Key: "orderer0",
		Org: orgOrderer,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "7050",
		},
		Bootstrap: &[]string{},
		Type:      &[]string{org.NodeOrderer},
	})
	generateConfig.AddNode(&nodes.Node{
		Key: "orderer1",
		Org: orgOrderer,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "8054",
		},
		Bootstrap: &[]string{},
		Type:      &[]string{org.NodeOrderer},
	})
	generateConfig.AddNode(&nodes.Node{
		Key: "orderer2",
		Org: orgOrderer,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "9050",
		},
		Bootstrap: &[]string{},
		Type:      &[]string{org.NodeOrderer},
	})

	generateConfig.AddNode(&nodes.Node{
		Key: "ca",
		Org: org0,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "9054",
		},
		Bootstrap: &[]string{},
		Type:      &[]string{org.NodeCA},
	})
	generateConfig.AddNode(&nodes.Node{
		Key: "peer0",
		Org: org0,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "8051",
		},
		Bootstrap: &[]string{"127.0.0.1:7051"},
		Type:      &[]string{org.NodeLeader, org.NodeAnchor, org.NodeCommit, org.NodeEndorse},
	})

	generateConfig.AddNode(&nodes.Node{
		Key: "ca",
		Org: org1,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "10054",
		},
		Bootstrap: &[]string{},
		Type:      &[]string{org.NodeCA},
	})
	generateConfig.AddNode(&nodes.Node{
		Key: "peer0",
		Org: org1,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "11051",
		},
		Bootstrap: &[]string{"127.0.0.1:7051"},
		Type:      &[]string{org.NodeLeader, org.NodeAnchor, org.NodeCommit, org.NodeEndorse},
	})

	generateConfig.AddNode(&nodes.Node{
		Key: "ca",
		Org: org2,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "12054",
		},
		Bootstrap: &[]string{},
		Type:      &[]string{org.NodeCA},
	})
	generateConfig.AddNode(&nodes.Node{
		Key: "peer0",
		Org: org2,
		Address: &nodes.ServerConfig{
			Host:       "172.20.10.3",
			SSHPort:    "22",
			FabricPort: "13051",
		},
		Bootstrap: &[]string{"127.0.0.1:7051"},
		Type:      &[]string{org.NodeLeader, org.NodeAnchor, org.NodeCommit, org.NodeEndorse},
	})

	// 加入通道
	orgPointer, err := generateConfig.GetGroupByKey(orgOrderer)
	if nil != err {
		return err
	}
	orgPointer.JoinChannel(channel1)

	orgPointer, err = generateConfig.GetGroupByKey(org0)
	if nil != err {
		return err
	}
	orgPointer.JoinChannel(channel1)

	orgPointer, err = generateConfig.GetGroupByKey(org1)
	if nil != err {
		return err
	}
	orgPointer.JoinChannel(channel1)

	orgPointer, err = generateConfig.GetGroupByKey(org2)
	if nil != err {
		return err
	}
	orgPointer.JoinChannel(channel1)

	// 全局保存配置信息
	Structure = &generateConfig
	// 处理结构体数据JSON
	data, err := json.Marshal(generateConfig)
	if nil != err {
		return err
	}
	//fmt.Println(string(data))
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if nil != err {
		return err
	}
	_, err = file.Write(data)
	if nil != err {
		return err
	}
	return file.Sync()
}
