package environment

import "faberGo/pkg/connect"

func OrganizationCheck(commands *connect.Commands) {
	commands.Append("mkdir /root/opt/organizations/peerOrganizations")
	commands.Append("mkdir /root/opt/organizations/ordererOrganizations")
}

func OrganizationCryptogenConfig() {

}

func OrganizationCreateOrderer() {

}

func OrganizationCreate(commands *connect.Commands) {
	commands.Append("cryptogen generate --config=/root/opt/config/crypto-config.yaml --output=\"crypto-config\"")
}
