package target

type ChannelPeer struct {
	Key            string `json:"key"`
	EndorsingPeer  bool   `json:"endorsingPeer"`
	ChaincodeQuery bool   `json:"chaincodeQuery"`
	LedgerQuery    bool   `json:"ledgerQuery"`
	EventSource    bool   `json:"eventSource"`
}

type RetryOpts struct {
	Attempts       int64  `json:"attempts"`
	InitialBackoff string `json:"initialBackoff"`
	MaxBackoff     string `json:"maxBackoff"`
	BackoffFactor  string `json:"backoffFactor"`
}

type QueryChannelConfig struct {
	MinResponses int64      `json:"minResponses"`
	MaxTargets   int64      `json:"maxTargets"`
	RetryOpts    *RetryOpts `json:"retryOpts"`
}

type ChannelPolices struct {
	QueryChannelConfig *QueryChannelConfig `json:"queryChannelConfig"`
}

type ChannelSelection struct {
	SortingStrategy         string `json:"SortingStrategy"`
	Balancer                string `json:"Balancer"`
	BlockHeightLagThreshold int64  `json:"BlockHeightLagThreshold"`
}

type ChannelEventService struct {
	ResolverStrategy                 string `json:"resolverStrategy"`
	MinBlockHeightResolverMode       string `json:"minBlockHeightResolverMode"`
	BlockHeightLagThreshold          int64  `json:"blockHeightLagThreshold"`
	ReconnectBlockHeightLagThreshold int64  `json:"reconnectBlockHeightLagThreshold"`
	PeerMonitor                      string `json:"peerMonitor"`
	PeerMonitorPeriod                string `json:"peerMonitorPeriod"`
}

type ChannelConfig struct {
	Key          string               `json:"name"`
	Orderers     *[]string            `json:"orderers"`
	Peers        *[]*ChannelPeer      `json:"peers"`
	Polices      *ChannelPolices      `json:"polices"`
	Selection    *ChannelSelection    `json:"selection"`
	EventService *ChannelEventService `json:"eventService"`
}

func GenerateDefaultChannel(name string) *ChannelConfig {
	return &ChannelConfig{
		Key:      name,
		Orderers: &[]string{},
		Peers:    &[]*ChannelPeer{},
		Polices: &ChannelPolices{QueryChannelConfig: &QueryChannelConfig{
			MinResponses: 1,
			MaxTargets:   1,
			RetryOpts: &RetryOpts{
				Attempts:       5,
				InitialBackoff: "500ms",
				MaxBackoff:     "5s",
				BackoffFactor:  "2.0",
			},
		}},
		Selection: &ChannelSelection{
			SortingStrategy:         "BlockHeightPriority",
			Balancer:                "RoundRobin",
			BlockHeightLagThreshold: 5,
		},
		EventService: &ChannelEventService{
			ResolverStrategy:                 "PreferOrg",
			MinBlockHeightResolverMode:       "ResolveByThreshold",
			BlockHeightLagThreshold:          5,
			ReconnectBlockHeightLagThreshold: 10,
			PeerMonitor:                      "Enabled",
			PeerMonitorPeriod:                "5s",
		},
	}
}

func GenerateSimpleChannel(name string) *ChannelConfig {
	return &ChannelConfig{
		Key:      name,
		Orderers: &[]string{},
		Peers:    &[]*ChannelPeer{},
	}
}

func GenerateDefaultPeer(key string) *ChannelPeer {
	return &ChannelPeer{
		Key:            key,
		EndorsingPeer:  false,
		ChaincodeQuery: true,
		LedgerQuery:    true,
		EventSource:    true,
	}
}

func GenerateEndorsingPeer(key string) *ChannelPeer {
	return &ChannelPeer{
		Key:            key,
		EndorsingPeer:  true,
		ChaincodeQuery: true,
		LedgerQuery:    true,
		EventSource:    true,
	}
}

//func GenerateChannelPeerConfigFromGenerateConfig() *ChannelPeer {
//	return nil
//}

func (that *ChannelConfig) AddPeer(peer *ChannelPeer) {
	for _, element := range *that.Peers {
		if peer.Key == element.Key {
			element.EndorsingPeer = element.EndorsingPeer || peer.EndorsingPeer
			element.ChaincodeQuery = element.ChaincodeQuery || peer.ChaincodeQuery
			element.LedgerQuery = element.LedgerQuery || peer.LedgerQuery
			element.EventSource = element.EventSource || peer.EventSource
		} else {
			*that.Peers = append(*that.Peers, peer)
		}
	}
}

func (that *ChannelConfig) AddOrderer(orderer string) {
	*that.Orderers = append(*that.Orderers, orderer)
}
