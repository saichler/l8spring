package main

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8spring/go/spring/ui"
)

func main() {
	svr := common.CreateWebServer("web", ui.RegisterTypes)
	svr.Start()
}
