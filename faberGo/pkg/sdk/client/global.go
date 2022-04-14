package client

type GlobalTimeout struct {
	Query   string `json:"query"`
	Execute string `json:"execute"`
	Resmgmt string `json:"resmgmt"`
}

type GlobalCache struct {
	ConnectionIdle    string `json:"connectionIdle"`
	EventServiceIdle  string `json:"eventServiceIdle"`
	ChannelConfig     string `json:"channelConfig"`
	ChannelMembership string `json:"channelMembership"`
	Discovery         string `json:"discovery"`
	Selection         string `json:"selection"`
}

type Global struct {
	Timeout *GlobalTimeout `json:"timeout"`
	Cache   *GlobalCache   `json:"cache"`
}

func GenerateDefaultGlobal() *Global {
	return &Global{
		Timeout: &GlobalTimeout{
			Query:   "180s",
			Execute: "180s",
			Resmgmt: "180s",
		},
		Cache: &GlobalCache{
			ConnectionIdle:    "30s",
			EventServiceIdle:  "2m",
			ChannelConfig:     "30m",
			ChannelMembership: "30s",
			Discovery:         "10s",
			Selection:         "10m",
		},
	}
}
