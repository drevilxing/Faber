package client

const LogInfo = "info"

type Log struct {
	Level string `json:"level" yaml:"level"`
}

func GenerateDefaultLog() *Log {
	return &Log{
		Level: LogInfo,
	}
}
