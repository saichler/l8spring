package listings

import (
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Listing"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName, ServiceArea: ServiceArea,
		PrimaryKey: "ListingId", Callback: newSpringListingServiceCallback(vnic),
	}, &spring.SpringListing{}, &spring.SpringListingList{}, creds, dbname, vnic)
}
