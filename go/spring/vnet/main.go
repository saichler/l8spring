package main

import (
	"github.com/saichler/l8bus/go/overlay/vnet"
	"github.com/saichler/l8spring/go/spring/common"
)

func main() {
	resources := common.CreateResources("vnet", false)
	net := vnet.NewVNet(resources)
	net.Start()
	resources.Logger().Info("vnet started!")
	common.WaitForSignal(resources)
}
