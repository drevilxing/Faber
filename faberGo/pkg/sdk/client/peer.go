package client

type PeerTimeoutDiscovery struct {
	GreylistExpiry string `json:"greylistExpiry" yaml:"greylistExpiry"`
}

type PeerTimeout struct {
	Connection string                `json:"connection" yaml:"connection"`
	Response   string                `json:"response" yaml:"response"`
	Discovery  *PeerTimeoutDiscovery `json:"discovery" yaml:"discovery"`
}

type Peer struct {
	Timeout *PeerTimeout `json:"timeout"`
}

func GenerateDefaultPeer() *Peer {
	return &Peer{
		Timeout: &PeerTimeout{
			Connection: "10s",
			Response:   "180s",
			Discovery:  &PeerTimeoutDiscovery{GreylistExpiry: "10s"},
		},
	}
}
