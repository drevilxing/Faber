package config

type GenerateConfig struct {
	Org     []OrgConfig     `json:"org"`
	Node    []NodeConfig    `json:"node"`
	Channel []ChannelConfig `json:"channel"`
}
