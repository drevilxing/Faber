package main

import (
	"faberGo/pkg/config"
	"fmt"
)

var Structure *config.GenerateConfig

func main() {
	environmentPath := "./environmentConfig.json"
	errEnv := GenerateDeploymentConfigExample(environmentPath)
	if nil != errEnv {
		fmt.Println(errEnv.Error())
	}

	structure := LoadGenerateConfigFromJsonFile(environmentPath)

	sdkPath := "./sdkConfig.json"
	errSdk := GenerateGoSdkConfigExample(sdkPath, structure)
	if nil != errSdk {
		fmt.Println(errSdk.Error())
	}
}
