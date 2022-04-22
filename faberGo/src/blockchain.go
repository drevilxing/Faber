package main

import (
	"faberGo/pkg/config/blockchain"
	"faberGo/pkg/config/nodes"
	"faberGo/pkg/config/org"
	"faberGo/pkg/https"
	"net/http"
	"strings"
)

func defaultHeaderBlockchain(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Token")
}

type RequestNetworkCreate struct {
	// 网络名称
	Name string `json:"name"`
}
type RequestBlockchainCreate struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func handlerBlockchainCreate(writer http.ResponseWriter, request *http.Request) {
	w, r, err := https.Dealer{
		Header:   defaultHeaderBlockchain,
		Handlers: nil,
	}.Deal(writer, request)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	postRequest := https.DealPostRequest(w, r, []string{
		"name", "key",
	}...)
	if nil != postRequest.Err {
		return
	}
	currentConfig.AddBlockchain(&blockchain.Blockchain{
		Key:      r.PostFormValue("key"),
		Name:     r.PostFormValue("name"),
		Channels: &[]string{},
	})
	https.SendResponseOK(writer, request)
}

func handlerBlockchainOrgCreate(writer http.ResponseWriter, request *http.Request) {
	w, r, err := https.Dealer{
		Header:   defaultHeaderBlockchain,
		Handlers: nil,
	}.Deal(writer, request)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	postRequest := https.DealPostRequest(w, r, []string{
		"name", "blockchain",
	}...)
	if nil != postRequest.Err {
		return
	}
	currentConfig.AddGroup(org.GenOrg(r.PostFormValue("blockchain"), r.PostFormValue("name")))
	https.SendResponseOK(writer, request)
}

func handlerBlockchainOrgNodeCreate(writer http.ResponseWriter, request *http.Request) {
	w, r, err := https.Dealer{
		Header:   defaultHeaderBlockchain,
		Handlers: nil,
	}.Deal(writer, request)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	postRequest := https.DealPostRequest(w, r, []string{
		"name", "org", "blockchain", "host", "ssh", "fabricPort", "user", "pwd", "key", "bootstrap", "type",
	}...)
	if nil != postRequest.Err {
		return
	}
	node := &nodes.Node{
		Key: r.PostFormValue("name"),
		Org: r.PostFormValue("org"),
		Address: &nodes.ServerConfig{
			Host:       r.PostFormValue("host"),
			SSHPort:    r.PostFormValue("ssh"),
			FabricPort: r.PostFormValue("fabricPort"),
		},
		Bootstrap: &[]string{},
		Type:      &[]string{},
	}
	currentConfig.AddNode(node)
	node.SetConnection(r.PostFormValue("user"), r.PostFormValue("pwd"), r.PostFormValue("key"))
	*node.Bootstrap = strings.Split(r.PostFormValue("bootstrap"), ";")
	*node.Type = strings.Split(r.PostFormValue("type"), ";")
	https.SendResponseOK(writer, request)
}

func handlerBlockchainChannelCreate(writer http.ResponseWriter, request *http.Request) {
	w, r, err := https.Dealer{
		Header:   defaultHeaderBlockchain,
		Handlers: nil,
	}.Deal(writer, request)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	postRequest := https.DealPostRequest(w, r, []string{
		"name", "blockchain",
	}...)
	if nil != postRequest.Err {
		return
	}
	currentConfig.AddBlockchainChannel(r.PostFormValue("blockchain"), r.PostFormValue("name"))
	https.SendResponseOK(writer, request)
}

func handlerBlockchainOrgJoinChannel(writer http.ResponseWriter, request *http.Request) {
	w, r, err := https.Dealer{
		Header:   defaultHeaderBlockchain,
		Handlers: nil,
	}.Deal(writer, request)
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	postRequest := https.DealPostRequest(w, r, []string{
		"org", "channel",
	}...)
	if nil != postRequest.Err {
		return
	}
	orgPointer, err := currentConfig.GetGroupByKey(r.PostFormValue("org"))
	if nil != err {
		https.SendResponseInternalError(w, r, err, https.ResponseFalse)
		return
	}
	orgPointer.JoinChannel(r.PostFormValue("channel"))
	https.SendResponseOK(writer, request)
}
