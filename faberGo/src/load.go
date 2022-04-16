package main

import (
	"encoding/json"
	"faberGo/pkg/config"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadGenerateConfigFromJsonFile(path string) *config.GenerateConfig {
	jsonFile, err := os.Open(path)

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(string(byteValue))
	return LoadGenerateConfigFromJsonByte(byteValue)
}

func LoadGenerateConfigFromJsonByte(data []byte) *config.GenerateConfig {
	config := &config.GenerateConfig{}
	err := json.Unmarshal(data, config)
	if nil != err {
		fmt.Println(err.Error())
		return nil
	}
	return config
}
