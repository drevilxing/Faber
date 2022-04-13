package blockchain

type Blockchain struct {
	Key      string    `json:"key"`
	Name     string    `json:"name"`
	Channels *[]string `json:"channels"`
}
