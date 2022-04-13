package sdk

type ClientConfig struct {
	Organization    string        `json:"organization"`
	Logging         string        `json:"logging"`
	Peer            PeerConfig    `json:"peer"`
	Orderer         OrdererConfig `json:"orderer"`
	CryptoPath      string        `json:"crypto_path"`
	CredentialStore string        `json:"credential_store"`
	CryptoStore     string        `json:"crypto_store"`
}
