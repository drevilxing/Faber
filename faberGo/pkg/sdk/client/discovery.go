package client

type DiscoveryTimeout struct {
	Connection string `json:"connection" yaml:"connection"`
	Response   string `json:"response" yaml:"response"`
}

type Discovery struct {
	Timeout *DiscoveryTimeout `json:"timeout" yaml:"timeout"`
}

func GenerateDefaultDiscovery() *Discovery {
	return &Discovery{Timeout: &DiscoveryTimeout{
		Connection: "15s",
		Response:   "15s",
	}}
}
