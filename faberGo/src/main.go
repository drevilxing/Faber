package main

import "faberGo/src/yaml"

const ConfigFiles = "./config"
const YamlFiles = "./assert"

func main() {
	yaml.GenerateCryptoConfigExample()

	// 从服务器对应位置加载配置文件信息
	//ListGenerateConfigFile(ConfigFiles)
	//fmt.Println(configs)
	//environmentPath := "./environmentConfig.json"
	//errEnv := GenerateDeploymentConfigExample(environmentPath)
	//if nil != errEnv {
	//	fmt.Println(errEnv.Error())
	//}

	//structure := LoadGenerateConfigFromJsonFile(environmentPath)

	//sdkPath := "./sdkConfig.yaml"
	//errSdk := GenerateGoSdkConfigExample(sdkPath, structure)
	//if nil != errSdk {
	//	fmt.Println(errSdk.Error())
	//}
	StartingBasicServer()
	//server.Listen()
}
