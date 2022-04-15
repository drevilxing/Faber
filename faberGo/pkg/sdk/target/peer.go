package target

type PeerConfig struct {
	Key         string       `json:"key"`
	Url         string       `json:"url"`
	EventUrl    string       `json:"eventUrl"`
	GrpcOptions *GrpcOptions `json:"grpcOptions"`
	TlsCACerts  *TlsCACerts  `json:"tlsCaCerts"`
}

func GenerateDefaultPeerConfig(key string, url string, eventUrl string) *PeerConfig {
	return &PeerConfig{
		Key:         key,
		Url:         "localhost:7051",
		EventUrl:    "localhost:7053",
		GrpcOptions: GenerateDefaultGrpcOptions(key),
		TlsCACerts:  &TlsCACerts{Path: "/etc/hyperledger/fabric/tls/tlscacerts"},
	}
}
