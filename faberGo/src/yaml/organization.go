package yaml

import "faberGo/pkg/yaml/organization"

func GenerateCryptoConfigExample() {
	config := organization.GenerateEmptyCryptogenConfig()
	config.AddOrdererOrg("Orderer", "orderer.test.com", "orderer", []string{"localhost"})
	config.AddPeerOrg("org0", "org0.test.com", []string{"localhost", "localhost"}, 1)
	config.AddPeerOrg("org1", "org1.test.com", []string{"localhost", "localhost"}, 1)
	config.AddPeerOrg("org2", "org2.test.com", []string{"localhost", "localhost"}, 1)
	config.GenerateYaml("./crypto-config")
}
