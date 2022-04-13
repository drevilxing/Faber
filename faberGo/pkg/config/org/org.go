package org

type Org struct {
	Key         string    `json:"key"`
	CA          string    `json:"ca"`
	Node        *Node     `json:"node"`
	Blockchains string    `json:"blockchains"`
	Channel     *[]string `json:"channel"`
}

func GenOrg(blockchain string, org string) *Org {
	return &Org{
		Key:         org,
		CA:          "",
		Node:        GenNodes(),
		Blockchains: blockchain,
		Channel:     &[]string{},
	}
}

func (that *Org) AddChannel(channel string) {
	*that.Channel = append(*that.Channel, channel)
}

func (that *Org) setCA(ca string) {
	that.CA = ca
	that.Node.SetCA(ca)
}

func (that *Org) AddNode(node string, tag string) {
	switch tag {
	case NodeCA:
		that.setCA(node)
		break
	case NodeOrderer:
		that.Node.AddOrderer(node)
		break
	case NodeLeader:
		that.Node.AddLeader(node)
		break
	case NodeAnchor:
		that.Node.AddAnchor(node)
		break
	case NodeCommit:
		that.Node.AddCommit(node)
		break
	case NodeEndorse:
		that.Node.AddEndorse(node)
		break
	default:
		return
	}
}

func (that *Org) UpdateBlockchain(blockchain string) {
	that.Blockchains = blockchain
}

func (that *Org) JoinChannel(channel string) {
	*that.Channel = append(*that.Channel, channel)
}
