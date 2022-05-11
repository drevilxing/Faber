package target

type ChannelPeer struct {
	Key            string `json:"key" yaml:"key"`
	EndorsingPeer  bool   `json:"endorsingPeer" yaml:"endorsingPeer"`
	ChaincodeQuery bool   `json:"chaincodeQuery" yaml:"chaincodeQuery"`
	LedgerQuery    bool   `json:"ledgerQuery" yaml:"ledgerQuery"`
	EventSource    bool   `json:"eventSource" yaml:"eventSource"`
}

type RetryOpts struct {
	Attempts       int64  `json:"attempts" yaml:"attempts"`
	InitialBackoff string `json:"initialBackoff" yaml:"initialBackoff"`
	MaxBackoff     string `json:"maxBackoff" yaml:"maxBackoff"`
	BackoffFactor  string `json:"backoffFactor" yaml:"backoffFactor"`
}

type QueryChannelConfig struct {
	MinResponses int64      `json:"minResponses" yaml:"minResponses"`
	MaxTargets   int64      `json:"maxTargets" yaml:"maxTargets"`
	RetryOpts    *RetryOpts `json:"retryOpts" yaml:"retryOpts"`
}

type ChannelPolices struct {
	QueryChannelConfig *QueryChannelConfig `json:"queryChannelConfig" yaml:"queryChannelConfig"`
}

type ChannelSelection struct {
	SortingStrategy         string `json:"SortingStrategy" yaml:"sortingStrategy"`
	Balancer                string `json:"Balancer" yaml:"balancer"`
	BlockHeightLagThreshold int64  `json:"BlockHeightLagThreshold" yaml:"blockHeightLagThreshold"`
}

type ChannelEventService struct {
	ResolverStrategy                 string `json:"resolverStrategy" yaml:"resolverStrategy"`
	MinBlockHeightResolverMode       string `json:"minBlockHeightResolverMode" yaml:"minBlockHeightResolverMode"`
	BlockHeightLagThreshold          int64  `json:"blockHeightLagThreshold" yaml:"blockHeightLagThreshold"`
	ReconnectBlockHeightLagThreshold int64  `json:"reconnectBlockHeightLagThreshold" yaml:"reconnectBlockHeightLagThreshold"`
	PeerMonitor                      string `json:"peerMonitor" yaml:"peerMonitor"`
	PeerMonitorPeriod                string `json:"peerMonitorPeriod" yaml:"peerMonitorPeriod"`
}

type ChannelConfig struct {
	Key          string               `json:"key" yaml:"key" yaml:"key"`
	Orderers     *[]string            `json:"orderers" yaml:"orderers" yaml:"orderers"`
	Peers        *[]*ChannelPeer      `json:"peers" yaml:"peers" yaml:"peers"`
	Polices      *ChannelPolices      `json:"polices" yaml:"polices" yaml:"polices"`
	Selection    *ChannelSelection    `json:"selection" yaml:"selection" yaml:"selection"`
	EventService *ChannelEventService `json:"eventService" yaml:"eventService" yaml:"eventService"`
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
	//fmt.Println(*that.Peers)
	for _, element := range *that.Peers {
		if peer.Key == element.Key {
			element.EndorsingPeer = element.EndorsingPeer || peer.EndorsingPeer
			element.ChaincodeQuery = element.ChaincodeQuery || peer.ChaincodeQuery
			element.LedgerQuery = element.LedgerQuery || peer.LedgerQuery
			element.EventSource = element.EventSource || peer.EventSource
			return
		}
	}
	*that.Peers = append(*that.Peers, peer)
}

func (that *ChannelConfig) AddOrderer(orderer string) {
	*that.Orderers = append(*that.Orderers, orderer)
}
