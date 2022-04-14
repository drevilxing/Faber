package client

type Config struct {
	Organization    string           `json:"organization"`
	Logging         *Log             `json:"logging"`
	Peer            *Peer            `json:"peer"`
	EventService    *EventService    `json:"eventService"`
	Orderer         *Orderer         `json:"orderer"`
	Discovery       *Discovery       `json:"discovery"`
	Global          *Global          `json:"global"`
	CryptoConfig    *CryptoConfig    `json:"cryptoconfig"`
	CredentialStore *CredentialStore `json:"credentialStore"`
	BCCSP           *BCCSP           `json:"BCCSP"`
	TlsCerts        *TlsCerts        `json:"tlsCerts"`
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
