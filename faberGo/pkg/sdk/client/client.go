package client

type Config struct {
	Organization    string           `json:"organization" yaml:"organization"`
	Logging         *Log             `json:"logging" yaml:"logging"`
	Peer            *Peer            `json:"peer" yaml:"peer"`
	EventService    *EventService    `json:"eventService" yaml:"eventService"`
	Orderer         *Orderer         `json:"orderer" yaml:"orderer"`
	Discovery       *Discovery       `json:"discovery" yaml:"discovery"`
	Global          *Global          `json:"global" yaml:"global"`
	CryptoConfig    *CryptoConfig    `json:"cryptoconfig" yaml:"cryptoConfig"`
	CredentialStore *CredentialStore `json:"credentialStore" yaml:"credentialStore"`
	BCCSP           *BCCSP           `json:"BCCSP" yaml:"BCCSP"`
	TlsCerts        *TlsCerts        `json:"tlsCerts" yaml:"tlsCerts"`
}

func GenerateDefaultClientConfig(organization string) *Config {
	return &Config{
		Organization:    organization,
		Logging:         GenerateDefaultLog(),
		Peer:            GenerateDefaultPeer(),
		EventService:    GenerateDefaultEventService(EventServiceTypeDeliver),
		Orderer:         GenerateDefaultOrderer(),
		Discovery:       GenerateDefaultDiscovery(),
		Global:          GenerateDefaultGlobal(),
		CryptoConfig:    GenerateDefaultCryptoConfig(),
		CredentialStore: GenerateDefaultCredentialStore(),
		BCCSP:           GenerateDefaultBCCSP(),
		TlsCerts:        GenerateDefaultTlsCerts(),
	}
}
