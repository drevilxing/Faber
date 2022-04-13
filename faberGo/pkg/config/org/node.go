package org

type Node struct {
	CA              string    `json:"ca"`
	Orderer         *[]string `json:"orderer"`
	LeaderPeers     *[]string `json:"leader_peers"`
	AnchorPeers     *[]string `json:"anchor_peers"`
	CommittingPeers *[]string `json:"committing_peers"`
	EndorsingPeers  *[]string `json:"endorsing_peers"`
}

func GenNodes() *Node {
	return &Node{
		CA:              "",
		Orderer:         &[]string{},
		LeaderPeers:     &[]string{},
		AnchorPeers:     &[]string{},
		CommittingPeers: &[]string{},
		EndorsingPeers:  &[]string{},
	}
}

func (that *Node) SetCA(ca string) {
	that.CA = ca
}

func (that *Node) AddOrderer(orderer string) {
	*that.Orderer = append(*that.Orderer, orderer)
}

func (that *Node) AddLeader(leader string) {
	*that.LeaderPeers = append(*that.LeaderPeers, leader)
}

func (that *Node) AddAnchor(anchor string) {
	*that.AnchorPeers = append(*that.AnchorPeers, anchor)
}

func (that *Node) AddCommit(commit string) {
	*that.CommittingPeers = append(*that.CommittingPeers, commit)
}

func (that *Node) AddEndorse(endorse string) {
	*that.EndorsingPeers = append(*that.EndorsingPeers, endorse)
}
