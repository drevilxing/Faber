package main

import "faberGo/pkg/https"

var server *https.Server

func StartingBasicServer() {
	server = &https.Server{
		Path: "",
		Port: 9000,
		Mask: "/faber",
		Routers: []https.Router{
			{
				Name: "network",
				Mask: "/network",
				Methods: []https.Method{
					{
						Pattern: "/create",
						Handler: handlerNetworkCreate,
					},
					{
						Pattern: "/list",
						Handler: handlerNetworkList,
					},
					{
						Pattern: "/delete",
						Handler: handlerNetworkDelete,
					},
					{
						Pattern: "/open",
						Handler: handlerNetworkOpen,
					},
				},
			},
			{
				Name: "blockchain",
				Mask: "/blockchain",
				Methods: []https.Method{
					{
						Pattern: "/create",
						Handler: handlerBlockchainCreate,
					},
					//https.Method{
					//	Pattern: "/update",
					//	Handler: nil,
					//},
					//https.Method{
					//	Pattern: "/read",
					//	Handler: nil,
					//},
					//https.Method{
					//	Pattern: "/delete",
					//	Handler: nil,
					//},
					{
						Pattern: "/channel/create",
						Handler: handlerBlockchainChannelCreate,
					},
					{
						Pattern: "/org/create",
						Handler: handlerBlockchainOrgCreate,
					},
					{
						Pattern: "/org/join/channel",
						Handler: handlerBlockchainOrgJoinChannel,
					},
					{
						Pattern: "/node/create",
						Handler: handlerBlockchainOrgNodeCreate,
					},
				},
			},
			{
				Name: "environment",
				Mask: "/environment",
				Methods: []https.Method{
					{
						Pattern: "/config/save",
						Handler: handlerConfigSave,
					},
					{
						Pattern: "/config/fetch",
						Handler: handlerConfigFetch,
					},
					{
						Pattern: "/config/generate",
						Handler: handlerConfigGenerate,
					},
					{
						Pattern: "/install",
						Handler: handlerInstall,
					},
					{
						Pattern: "/deploy",
						Handler: handlerDeploy,
					},
				},
			},
		},
	}
}
