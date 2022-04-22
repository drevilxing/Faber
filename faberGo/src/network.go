package main

import (
	"encoding/json"
	"errors"
	"faberGo/pkg/config"
	"faberGo/pkg/https"
	"net/http"
)

func defaultHeaderNetwork(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Token")
}

func handlerNetworkCreate(writer http.ResponseWriter, request *http.Request) {
	w, r, err := https.Dealer{
		Header:   defaultHeaderNetwork,
		Handlers: nil,
	}.Deal(writer, request)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	postRequest := https.DealPostRequest(w, r, []string{
		"name",
	}...)
	if nil != postRequest.Err {
		return
	}

	// 检查
	for _, element := range *configs {
		if element.Key == r.PostFormValue("name") {
			https.SendResponseInternalError(w, r, errors.New("Duplicate Key. "), https.ResponseFalse)
			return
		}
	}
	// 创建
	*currentConfig = config.GenGenerateConfig()
	// 添加到配置列表
	*configs = append(*configs, currentConfig)
	// 设置名称
	currentConfig.Key = r.PostFormValue("name")
	// 保存配置文件
	err = currentConfig.SaveToPath(ConfigFiles)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	https.SendResponseOK(writer, request)
}

func handlerNetworkList(writer http.ResponseWriter, request *http.Request) {
	w, r, err := https.Dealer{
		Header:   defaultHeaderNetwork,
		Handlers: nil,
	}.Deal(writer, request)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	getRequest := https.DealGetRequest(w, r, []string{}...)
	if nil != getRequest.Err {
		return
	}
	dataTemp := &[]json.RawMessage{}
	for _, element := range *configs {
		temp, errIn := json.Marshal(*element)

		if nil != errIn {
			https.SendResponseInternalError(w, r, err, https.ResponseFalse)
			return
		}
		*dataTemp = append(*dataTemp, temp)
	}
	data, err := json.Marshal(*dataTemp)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	https.SendJsonResponse(w, r, https.JsonResponse{
		Status:  https.ResponseTrue,
		Message: data,
		Code:    http.StatusOK,
	})
}

func handlerNetworkDelete(writer http.ResponseWriter, request *http.Request) {
	w, r, err := https.Dealer{
		Header:   defaultHeaderNetwork,
		Handlers: nil,
	}.Deal(writer, request)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	postRequest := https.DealPostRequest(w, r, []string{
		"name",
	}...)
	if nil != postRequest.Err {
		return
	}

	// 检查并删除
	for index, element := range *configs {
		if element.Key == r.PostFormValue("name") {

			errIn := config.Remove(ConfigFiles, (*configs)[index].Key)
			if nil != errIn {
				https.SendResponseInternalError(w, r, err, https.ResponseFalse)
				return
			}
			*configs = append((*configs)[:index], (*configs)[index+1:]...)
			break
		}
	}
	https.SendResponseOK(writer, request)
}
