package nodes

type ServerConfig struct {
	Host       string `json:"host"`
	SSHPort    string `json:"ssh_port"`
	FabricPort string `json:"fabric_port"`
}
