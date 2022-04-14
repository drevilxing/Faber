package client

type OrdererTimeout struct {
	Connection string `json:"connection"`
	Response   string `json:"response"`
}

type Orderer struct {
	Timeout *OrdererTimeout `json:"timeout"`
}

func GenerateDefaultOrderer() *Orderer {
	return &Orderer{Timeout: &OrdererTimeout{
		Connection: "15s",
		Response:   "15s",
	}}
}
