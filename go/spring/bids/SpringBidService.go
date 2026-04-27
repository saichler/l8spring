package bids

import (
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

const (
	ServiceName = "Bid"
	ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
	common.ActivateService(common.ServiceConfig{
		ServiceName: ServiceName, ServiceArea: ServiceArea,
		PrimaryKey: "BidId", Callback: newSpringBidServiceCallback(vnic),
	}, &spring.SpringBid{}, &spring.SpringBidList{}, creds, dbname, vnic)
}
