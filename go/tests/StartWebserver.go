package tests

import (
	"github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/spring/ui"
	"github.com/saichler/l8bus/go/overlay/health"
	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8web/go/web/server"
)

func startWebServer(port int, nic ifs.IVNic) ifs.IWebServer {
	ui.RegisterTypes(nic.Resources())

	serverConfig := &server.RestServerConfig{
		Host:           "localhost",
		Port:           port,
		Authentication: true,
		Prefix:         common.PREFIX,
		CertName:       "/data/l8spring",
	}
	svr, err := server.NewRestServer(serverConfig)
	if err != nil {
		panic(err)
	}

	hs, ok := nic.Resources().Services().ServiceHandler(health.ServiceName, 0)
	if ok {
		ws := hs.WebService()
		svr.RegisterWebService(ws, nic)
	}

	sla := ifs.NewServiceLevelAgreement(&server.WebService{}, ifs.WebService, 0, false, nil)
	sla.SetArgs(svr)
	nic.Resources().Services().Activate(sla, nic)

	nic.Resources().Logger().Info("Spring Test Web Server Started!")
	go svr.Start()

	return svr
}
