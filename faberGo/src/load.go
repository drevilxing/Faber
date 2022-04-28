package main

import (
	"encoding/json"
	"faberGo/pkg/config"
	"fmt"
	"io/ioutil"
	"os"
)

var configs *[]*config.GenerateConfig
var currentConfig *config.GenerateConfig

func ListGenerateConfigFile(path string) {
	conf := &[]*config.GenerateConfig{}
	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// 获取文件，并读入配置文件
	for _, file := range files {
		*conf = append(*conf, LoadGenerateConfigFromJsonFile(path+"/"+file.Name()))
	}
	configs = conf
}

func LoadGenerateConfigFromJsonFile(path string) *config.GenerateConfig {
	jsonFile, err := os.Open(path)

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer func(jsonFile *os.File) {
		errIn := jsonFile.Close()
		if errIn != nil {
			return
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(string(byteValue))
	return LoadGenerateConfigFromJsonByte(byteValue)
}

func LoadGenerateConfigFromJsonByte(data []byte) *config.GenerateConfig {
	generateConfig := &config.GenerateConfig{}
	err := json.Unmarshal(data, generateConfig)
	if nil != err {
		fmt.Println(err.Error())
		return nil
	}
	return generateConfig
}
