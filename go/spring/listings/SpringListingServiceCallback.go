package listings

import (
	l8common "github.com/saichler/l8common/go/types/l8common"
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

func newSpringListingServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&spring.SpringListing{}, vnic).
		Require(func(v interface{}) string { return v.(*spring.SpringListing).Title }, "Title").
		Require(func(v interface{}) string { return v.(*spring.SpringListing).BuyerId }, "BuyerId").
		Require(func(v interface{}) string { return v.(*spring.SpringListing).Description }, "Description").
		Require(func(v interface{}) string { return v.(*spring.SpringListing).CategoryId }, "CategoryId").
		Enum(func(v interface{}) int32 { return int32(v.(*spring.SpringListing).DesiredCondition) }, spring.SpringItemCondition_name, "DesiredCondition").
		Money(func(v interface{}) *l8common.Money { return v.(*spring.SpringListing).MaxPrice }, "MaxPrice").
		Enum(func(v interface{}) int32 { return int32(v.(*spring.SpringListing).Status) }, spring.SpringListingStatus_name, "Status").
		StatusTransition(&common.StatusTransitionConfig{
			StatusGetter:  func(v interface{}) int32 { return int32(v.(*spring.SpringListing).Status) },
			StatusSetter:  func(v interface{}, s int32) { v.(*spring.SpringListing).Status = spring.SpringListingStatus(s) },
			FilterBuilder: func(v interface{}) interface{} { return &spring.SpringListing{ListingId: v.(*spring.SpringListing).ListingId} },
			ServiceName:   ServiceName,
			ServiceArea:   ServiceArea,
			InitialStatus: int32(spring.SpringListingStatus_SPRING_LISTING_STATUS_DRAFT),
			Transitions: map[int32][]int32{
				int32(spring.SpringListingStatus_SPRING_LISTING_STATUS_DRAFT): {
					int32(spring.SpringListingStatus_SPRING_LISTING_STATUS_ACTIVE),
					int32(spring.SpringListingStatus_SPRING_LISTING_STATUS_CANCELLED),
				},
				int32(spring.SpringListingStatus_SPRING_LISTING_STATUS_ACTIVE): {
					int32(spring.SpringListingStatus_SPRING_LISTING_STATUS_FULFILLED),
					int32(spring.SpringListingStatus_SPRING_LISTING_STATUS_EXPIRED),
					int32(spring.SpringListingStatus_SPRING_LISTING_STATUS_CANCELLED),
				},
			},
			StatusNames: map[int32]string{
				0: "Unspecified",
				1: "Draft",
				2: "Active",
				3: "Fulfilled",
				4: "Expired",
				5: "Cancelled",
			},
		}).
		Build()
}
