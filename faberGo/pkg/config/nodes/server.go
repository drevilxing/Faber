package nodes

type ServerConfig struct {
	Host       string `json:"host"`
	SSHPort    string `json:"ssh_port"`
	FabricPort string `json:"fabric_port"`
	user       string
	pwd        string
	key        string
	password   bool
}

func (that *ServerConfig) SetConnection(user string, pwd string, key string) {
	if pwd == "" {
		that.password = false
		that.key = key
	} else {
		that.password = true
	}
	that.user = user
}
