package client

type OrdererTimeout struct {
	Connection string `json:"connection" yaml:"connection"`
	Response   string `json:"response" yaml:"response"`
}

type Orderer struct {
	Timeout *OrdererTimeout `json:"timeout" yaml:"timeout"`
}

func GenerateDefaultOrderer() *Orderer {
	return &Orderer{Timeout: &OrdererTimeout{
		Connection: "15s",
		Response:   "15s",
	}}
}
