package sdk

type GoSDK struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Version     string        `json:"version"`
	Client      ClientConfig  `json:"client"`
	Channel     ChannelConfig `json:"channel"`
}
