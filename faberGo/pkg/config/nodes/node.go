package nodes

type Node struct {
	Key       string        `json:"key"`
	Org       string        `json:"org"`
	Address   *ServerConfig `json:"address"`
	Bootstrap *[]string     `json:"bootstrap"`
	Type      *[]string     `json:"type"`
}

func (that Node) AddType(tag string) {
	*that.Type = append(*that.Type, tag)
}
