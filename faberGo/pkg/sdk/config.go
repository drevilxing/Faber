package sdk

import (
	"encoding/json"
	"errors"
	"faberGo/pkg/config"
	"faberGo/pkg/sdk/client"
	"faberGo/pkg/sdk/target"
	"fmt"
	"os"
)

type GoSDK struct {
	Name           string                   `json:"name" yaml:"name"`
	Description    string                   `json:"description" yaml:"description"`
	Version        string                   `json:"version" yaml:"version"`
	Client         *client.Config           `json:"client" yaml:"client"`
	Channels       *[]*target.ChannelConfig `json:"channels" yaml:"channels"`
	Organizations  *[]*target.OrgConfig     `json:"organizations" yaml:"organizations"`
	Orderers       *[]*target.OrdererConfig `json:"orderers" yaml:"orderers"`
	Peers          *[]*target.PeerConfig    `json:"peers" yaml:"peers"`
	CA             *[]*target.CAConfig      `json:"certificateAuthorities" yaml:"certificateAuthorities"`
	EntityMatchers *target.EntityMatcher    `json:"entityMatchers" yaml:"entityMatchers"`
	host           *[]string
}

func GenerateGoSDK(name string, desc string, version string, org string, generate *config.GenerateConfig) *GoSDK {
	// 生成基础配置文件
	sdkConfig := &GoSDK{
		Name:        name,
		Description: desc,
		Version:     version,
		Client:      client.GenerateDefaultClientConfig(org),
		Channels:    &[]*target.ChannelConfig{
			//target.GenerateSimpleChannel(channel),
		},
		Organizations: &[]*target.OrgConfig{
			//target.GenerateDefaultOrgConfig(org, mspId),
		},
		Orderers:       &[]*target.OrdererConfig{},
		Peers:          &[]*target.PeerConfig{},
		CA:             &[]*target.CAConfig{},
		EntityMatchers: target.GenerateDefaultEntityMatcher(),
		host:           &[]string{},
	}

	// 通道部分
	// 遍历区块链网络
	for _, element := range *generate.Blockchains {
		// 遍历通道并加入配置文件
		for _, channelElement := range *element.Channels {
			*sdkConfig.Channels = append(*sdkConfig.Channels, target.GenerateDefaultChannel(channelElement))
			*sdkConfig.EntityMatchers.Channel = append(*sdkConfig.EntityMatchers.Channel, &target.Matcher{
				Pattern:    channelElement + "$",
				MappedHost: channelElement,
			})
		}
	}
	// 组织部分
	// 遍历组织
	for _, element := range *generate.Groups {
		*sdkConfig.Organizations = append(*sdkConfig.Organizations, target.GenerateDefaultOrgConfig(element.Key, target.MspId(element.Key)))
		orgPointer, err := sdkConfig.FindOrg(element.Key)
		if nil != err {
			return nil
		}
		// 添加CA以及其他节点到组织信息内
		// 添加CA节点
		orgPointer.AddPeer(element.CA)
		*sdkConfig.CA = append(*sdkConfig.CA, target.GenerateDefaultCAConfig(element.CA, "localhost:7054"))
		sdkConfig.EntityMatchers.AddCA(element.CA)
		// 添加Orderer节点
		for _, peer := range *element.Node.Orderer {
			orgPointer.AddPeer(peer)
			// 添加到对应channel的order信息
			for _, channelElement := range *element.Channel {
				channelPointer, errChannelPointer := sdkConfig.FindChannel(channelElement)
				if nil != errChannelPointer {
					return nil
				}
				channelPointer.AddOrderer(peer)
			}
			// 添加到对应orderers内
			*sdkConfig.Orderers = append(*sdkConfig.Orderers, target.GenerateDefaultOrdererConfig(peer, "localhost:7050"))
			sdkConfig.EntityMatchers.AddOrderer(peer)
		}
		for _, peer := range *element.Node.LeaderPeers {
			orgPointer.AddPeer(peer)
			// 添加到对应channel的peer信息
			for _, channelElement := range *element.Channel {
				channelPointer, errChannelPointer := sdkConfig.FindChannel(channelElement)
				if nil != errChannelPointer {
					return nil
				}
				channelPointer.AddPeer(target.GenerateDefaultPeer(peer))
			}
			// 添加节点信息
			sdkConfig.AddPeer(target.GenerateDefaultPeerConfig(peer, "localhost:7051", "localhost:7053"))
			sdkConfig.EntityMatchers.AddPeer(peer)
		}
		for _, peer := range *element.Node.AnchorPeers {
			orgPointer.AddPeer(peer)
			// 添加到对应channel的peer信息
			for _, channelElement := range *element.Channel {
				channelPointer, errChannelPointer := sdkConfig.FindChannel(channelElement)
				if nil != errChannelPointer {
					return nil
				}
				channelPointer.AddPeer(target.GenerateDefaultPeer(peer))
			}
			// 添加节点信息
			sdkConfig.AddPeer(target.GenerateDefaultPeerConfig(peer, "localhost:7051", "localhost:7053"))
			sdkConfig.EntityMatchers.AddPeer(peer)
		}
		for _, peer := range *element.Node.CommittingPeers {
			orgPointer.AddPeer(peer)
			// 添加到对应channel的peer信息
			for _, channelElement := range *element.Channel {
				channelPointer, errChannelPointer := sdkConfig.FindChannel(channelElement)
				if nil != errChannelPointer {
					return nil
				}
				channelPointer.AddPeer(target.GenerateDefaultPeer(peer))
			}
			// 添加节点信息
			sdkConfig.AddPeer(target.GenerateDefaultPeerConfig(peer, "localhost:7051", "localhost:7053"))
			sdkConfig.EntityMatchers.AddPeer(peer)
		}
		for _, peer := range *element.Node.EndorsingPeers {
			orgPointer.AddPeer(peer)
			// 添加到对应channel的peer信息
			for _, channelElement := range *element.Channel {
				fmt.Println(channelElement)
				channelPointer, errChannelPointer := sdkConfig.FindChannel(channelElement)
				if nil != errChannelPointer {
					return nil
				}
				channelPointer.AddPeer(target.GenerateEndorsingPeer(peer))
			}
			// 添加节点信息
			sdkConfig.AddPeer(target.GenerateDefaultPeerConfig(peer, "localhost:7051", "localhost:7053"))
			sdkConfig.EntityMatchers.AddPeer(peer)
		}
	}
	for _, element := range *generate.Nodes {
		//fmt.Println(element.Key, element.Bootstrap, element.Address)
		*sdkConfig.host = append(*sdkConfig.host, element.Address.Host+":"+element.Address.FabricPort+" "+element.Key)
	}
	return sdkConfig
}

func (that *GoSDK) FindChannel(key string) (*target.ChannelConfig, error) {
	for _, element := range *that.Channels {
		if key == element.Key {
			return element, nil
		}
	}
	return nil, errors.New(config.NotFoundErr)
}

func (that *GoSDK) FindOrg(key string) (*target.OrgConfig, error) {
	for _, element := range *that.Organizations {
		if key == element.Key {
			return element, nil
		}
	}
	return nil, errors.New(config.NotFoundErr)
}

func (that *GoSDK) FindOrderer(key string) (*target.OrdererConfig, error) {
	for _, element := range *that.Orderers {
		if key == element.Key {
			return element, nil
		}
	}
	return nil, errors.New(config.NotFoundErr)
}

func (that *GoSDK) AddPeer(peer *target.PeerConfig) {
	for _, element := range *that.Peers {
		if peer.Key == element.Key {
			return
		}
	}
	*that.Peers = append(*that.Peers, peer)
}

func (that *GoSDK) FindPeer(key string) (*target.PeerConfig, error) {
	for _, element := range *that.Peers {
		if key == element.Key {
			return element, nil
		}
	}
	return nil, errors.New(config.NotFoundErr)
}

func (that *GoSDK) FindCA(key string) (*target.CAConfig, error) {
	for _, element := range *that.CA {
		if key == element.Key {
			return element, nil
		}
	}
	return nil, errors.New(config.NotFoundErr)
}

func (that *GoSDK) SaveHost(path string) error {
	data, err := json.Marshal(*that.host)
	if nil != err {
		return err
	}
	file, err := os.OpenFile(path+"host", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
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
