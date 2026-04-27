package main

import (
	"fmt"
	"github.com/saichler/l8bus/go/overlay/vnic"
	"github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/spring/services"
	"github.com/saichler/l8types/go/ifs"
	"os"
	"os/exec"
	"time"
)

func main() {
	res := common.CreateResources("SpringServices", false)
	ifs.SetNetworkMode(ifs.NETWORK_K8s)
	nic := vnic.NewVirtualNetworkInterface(res, nil)
	nic.Start()
	nic.WaitForConnection()

	if len(os.Args) == 1 {
		startDb(nic)
	}

	dbcred := nic.Resources().SysConfig().DataStoreConfig.Type
	dbname := nic.Resources().SysConfig().DataStoreConfig.Name

	services.ActivateAllServices(dbcred, dbname, nic)

	common.WaitForSignal(res)
}

func startDb(nic ifs.IVNic) {
	dbcred := nic.Resources().SysConfig().DataStoreConfig.Type
	dbname := nic.Resources().SysConfig().DataStoreConfig.Name

	_, user, pass, _, err := nic.Resources().Security().Credential(dbcred, dbname, nic.Resources())
	if err != nil {
		panic(dbcred + " " + dbname + " " + err.Error())
	}
	if user == "admin" && pass == "admin" {
		dbname = "admin"
	}

	var cmd *exec.Cmd
	cmd = exec.Command("nohup", "/start-postgres.sh", dbname, user, pass)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	time.Sleep(time.Second * 5)
}
