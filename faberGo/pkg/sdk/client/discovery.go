package client

type DiscoveryTimeout struct {
	Connection string `json:"connection"`
	Response   string `json:"response"`
}

type Discovery struct {
	Timeout *DiscoveryTimeout `json:"timeout"`
}

func GenerateDefaultDiscovery() *Discovery {
	return &Discovery{Timeout: &DiscoveryTimeout{
		Connection: "15s",
		Response:   "15s",
	}}
}
