package tests

import (
	"crypto/tls"
	"fmt"
	"github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/spring/services"
	"github.com/saichler/l8spring/go/tests/mocks"
	l8common "github.com/saichler/l8common/go/common"
	"github.com/saichler/l8types/go/ifs"
	"net/http"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
	tear()
}

func dropAllTables(t *testing.T, vnic ifs.IVNic) {
	creds := common.DB_CREDS
	dbname := common.DB_NAME
	_, user, pass, port, err := vnic.Resources().Security().Credential(creds, dbname, vnic.Resources())
	if err != nil {
		t.Fatalf("Failed to get credentials: %v", err)
	}
	db := l8common.OpenDBConection(dbname, user, pass, port)
	_, err = db.Exec("DROP SCHEMA public CASCADE")
	if err != nil {
		t.Fatalf("Failed to drop schema: %v", err)
	}
	_, err = db.Exec("CREATE SCHEMA public")
	if err != nil {
		t.Fatalf("Failed to recreate schema: %v", err)
	}
	fmt.Println("Cleaned database (dropped and recreated public schema)")
}

func TestAllServices(t *testing.T) {
	erpServicesVnic := topo.VnicByVnetNum(1, 1)
	webServiceVnic := topo.VnicByVnetNum(3, 3)
	log := webServiceVnic.Resources().Logger()

	dropAllTables(t, erpServicesVnic)

	services.ActivateAllServices(common.DB_CREDS, common.DB_NAME, erpServicesVnic)

	port := 9443
	startWebServer(port, webServiceVnic)
	time.Sleep(10 * time.Second)

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client := mocks.NewSpringClient(fmt.Sprintf("https://localhost:%d", port), httpClient)
	err := client.Authenticate("operator", "operator")
	if err != nil {
		log.Fail(t, "Authentication failed: ", err.Error())
		return
	}

	testStore = &mocks.MockDataStore{}
	mocks.RunAllPhases(client, testStore)

	if len(testStore.CategoryIDs) == 0 {
		log.Fail(t, "No categories generated")
	}
	if len(testStore.ListingIDs) == 0 {
		log.Fail(t, "No listings generated")
	}
	if len(testStore.BidIDs) == 0 {
		log.Fail(t, "No bids generated")
	}
	if len(testStore.DealIDs) == 0 {
		log.Fail(t, "No deals generated")
	}
	if len(testStore.ReviewIDs) == 0 {
		log.Fail(t, "No reviews generated")
	}

	mocks.PrintSummary(testStore)

	testCRUD(t, client)
}
