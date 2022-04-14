package client

type PeerTimeoutDiscovery struct {
	GreylistExpiry string `json:"greylistExpiry"`
}

type PeerTimeout struct {
	Connection string                `json:"connection"`
	Response   string                `json:"response"`
	Discovery  *PeerTimeoutDiscovery `json:"discovery"`
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
