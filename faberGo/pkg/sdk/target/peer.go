package target

type PeerConfig struct {
	Key         string       `json:"key" yaml:"key"`
	Url         string       `json:"url" yaml:"url"`
	EventUrl    string       `json:"eventUrl" yaml:"eventUrl"`
	GrpcOptions *GrpcOptions `json:"grpcOptions" yaml:"grpcOptions"`
	TlsCACerts  *TlsCACerts  `json:"tlsCaCerts" yaml:"tlsCACerts"`
}

func GenerateDefaultPeerConfig(key string, url string, eventUrl string) *PeerConfig {
	// TODO 需要在生成网络时修改对应证书文件名为tls.pem
	return &PeerConfig{
		Key:         key,
		Url:         "localhost:7051",
		EventUrl:    "localhost:7053",
		GrpcOptions: GenerateDefaultGrpcOptions(key),
		TlsCACerts:  &TlsCACerts{Path: "/etc/hyperledger/fabric/tls/tlscacerts/tls.pem"},
	}
}
