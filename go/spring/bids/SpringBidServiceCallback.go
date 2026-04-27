package bids

import (
	l8common "github.com/saichler/l8common/go/types/l8common"
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

func newSpringBidServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&spring.SpringBid{}, vnic).
		Require(func(v interface{}) string { return v.(*spring.SpringBid).ListingId }, "ListingId").
		Require(func(v interface{}) string { return v.(*spring.SpringBid).SellerId }, "SellerId").
		Require(func(v interface{}) string { return v.(*spring.SpringBid).Description }, "Description").
		Money(func(v interface{}) *l8common.Money { return v.(*spring.SpringBid).Price }, "Price").
		Enum(func(v interface{}) int32 { return int32(v.(*spring.SpringBid).ItemCondition) }, spring.SpringItemCondition_name, "ItemCondition").
		Enum(func(v interface{}) int32 { return int32(v.(*spring.SpringBid).Status) }, spring.SpringBidStatus_name, "Status").
		StatusTransition(&common.StatusTransitionConfig{
			StatusGetter:  func(v interface{}) int32 { return int32(v.(*spring.SpringBid).Status) },
			StatusSetter:  func(v interface{}, s int32) { v.(*spring.SpringBid).Status = spring.SpringBidStatus(s) },
			FilterBuilder: func(v interface{}) interface{} { return &spring.SpringBid{BidId: v.(*spring.SpringBid).BidId} },
			ServiceName:   ServiceName,
			ServiceArea:   ServiceArea,
			InitialStatus: int32(spring.SpringBidStatus_SPRING_BID_STATUS_PENDING),
			Transitions: map[int32][]int32{
				int32(spring.SpringBidStatus_SPRING_BID_STATUS_PENDING): {
					int32(spring.SpringBidStatus_SPRING_BID_STATUS_ACCEPTED),
					int32(spring.SpringBidStatus_SPRING_BID_STATUS_REJECTED),
					int32(spring.SpringBidStatus_SPRING_BID_STATUS_WITHDRAWN),
					int32(spring.SpringBidStatus_SPRING_BID_STATUS_EXPIRED),
				},
			},
			StatusNames: map[int32]string{
				0: "Unspecified",
				1: "Pending",
				2: "Accepted",
				3: "Rejected",
				4: "Withdrawn",
				5: "Expired",
			},
		}).
		After(afterBidUpdate).
		Build()
}
