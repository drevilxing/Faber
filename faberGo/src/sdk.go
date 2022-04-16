package main

import (
	"encoding/json"
	"faberGo/pkg/config"
	"faberGo/pkg/sdk"
	"os"
)

func GenerateGoSdkConfigExample(path string, structure *config.GenerateConfig) error {
	sdkConfig := sdk.GenerateGoSDK("faber", "Faber", "1.0.0", orgOrderer, structure)
	err := sdkConfig.SaveHost("")
	if nil != err {
		return err
	}
	data, err := json.Marshal(sdkConfig)
	if nil != err {
		return err
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if nil != err {
		return err
	}
	_, err = file.Write(data)
	if nil != err {
		return err
	}
	return file.Sync()
}
