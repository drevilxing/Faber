package sdk

import (
	"faberGo/pkg/sdk/client"
	"faberGo/pkg/sdk/entryMatcher"
	"faberGo/pkg/sdk/target"
)

type GoSDK struct {
	Name          string                     `json:"name"`
	Description   string                     `json:"description"`
	Version       string                     `json:"version"`
	Client        *client.Config             `json:"client"`
	Channel       *[]*target.ChannelConfig   `json:"channel"`
	Organizations *[]*target.OrgConfig       `json:"organizations"`
	Orderers      *[]*target.OrdererConfig   `json:"orderers"`
	Peers         *[]*target.PeerConfig      `json:"peers"`
	CA            *[]*target.CAConfig        `json:"certificateAuthorities"`
	EntryMatchers *entryMatcher.EntryMatcher `json:"entryMatchers"`
}

func GenerateGoSDK(name string, desc string, version string, org string) *GoSDK {
	return &GoSDK{
		Name:          name,
		Description:   desc,
		Version:       version,
		Client:        client.GenerateDefaultClientConfig(org),
		Channel:       nil,
		Organizations: nil,
		Orderers:      nil,
		Peers:         nil,
		CA:            nil,
		EntryMatchers: nil,
	}
}
