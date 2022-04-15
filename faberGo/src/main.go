package main

import (
	"faberGo/pkg/config"
	"fmt"
)

var Structure *config.GenerateConfig

func main() {
	environmentPath := "./environmentConfig.json"
	err := GenerateDeploymentConfigExample(environmentPath)
	if nil != err {
		fmt.Println(err.Error())
	}
	sdkPath := "./sdkConfig.json"
	err = GenerateGoSdkConfigExample(sdkPath)
	if nil != err {
		fmt.Println(err.Error())
	}
}
