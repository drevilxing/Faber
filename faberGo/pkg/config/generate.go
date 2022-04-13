package config

import (
	"errors"
	"faberGo/pkg/config/blockchain"
	"faberGo/pkg/config/nodes"
	"faberGo/pkg/config/org"
	"fmt"
)

const NotFoundErr = "Not Found. "

const DefaultOrder = "127.0.0.1:7050"
const DefaultPeer = "127.0.0.1:7051"
const DefaultCA = "127.0.0.1:7054"

type GenerateConfig struct {
	Groups      *[]*org.Org               `json:"groups"`
	Nodes       *[]*nodes.Node            `json:"nodes"`
	Blockchains *[]*blockchain.Blockchain `json:"blockchains"`
}

func GenGenerateConfig() GenerateConfig {
	return GenerateConfig{
		Groups:      &[]*org.Org{},
		Nodes:       &[]*nodes.Node{},
		Blockchains: &[]*blockchain.Blockchain{},
	}
}

func (that GenerateConfig) AddGroup(org *org.Org) {
	*that.Groups = append(*that.Groups, org)
}

func (that GenerateConfig) GetGroupByKey(org string) (*org.Org, error) {
	for _, element := range *that.Groups {
		if org == element.Key {
			return element, nil
		}
	}
	return nil, errors.New(NotFoundErr)
}

//func (that GenerateConfig) addNode(node nodes.Node) (*org.Org, error) {
//	orgPointer, err := that.GetGroupByKey(node.Org)
//	if nil != err {
//		return nil, err
//	}
//	node.Key = node.Key + "." + orgPointer.Key
//	*that.Nodes = append(*that.Nodes, node)
//	return orgPointer, nil
//}

func (that GenerateConfig) AddNode(node *nodes.Node) {
	orgPointer, err := that.GetGroupByKey(node.Org)
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	node.Key += "." + orgPointer.Key
	*that.Nodes = append(*that.Nodes, node)
	for _, element := range *node.Type {
		orgPointer.AddNode(node.Key, element)
	}
}

func (that GenerateConfig) AddBlockchain(blockchain *blockchain.Blockchain) {
	*that.Blockchains = append(*that.Blockchains, blockchain)
}
