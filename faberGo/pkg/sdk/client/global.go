package client

type GlobalTimeout struct {
	Query   string `json:"query" yaml:"query"`
	Execute string `json:"execute" yaml:"execute"`
	Resmgmt string `json:"resmgmt" yaml:"resmgmt"`
}

type GlobalCache struct {
	ConnectionIdle    string `json:"connectionIdle" yaml:"connectionIdle"`
	EventServiceIdle  string `json:"eventServiceIdle" yaml:"eventServiceIdle"`
	ChannelConfig     string `json:"channelConfig" yaml:"channelConfig"`
	ChannelMembership string `json:"channelMembership" yaml:"channelMembership"`
	Discovery         string `json:"discovery" yaml:"discovery"`
	Selection         string `json:"selection" yaml:"selection"`
}

type Global struct {
	Timeout *GlobalTimeout `json:"timeout" yaml:"timeout"`
	Cache   *GlobalCache   `json:"cache" yaml:"cache"`
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
