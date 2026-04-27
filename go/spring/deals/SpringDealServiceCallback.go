package deals

import (
	l8common "github.com/saichler/l8common/go/types/l8common"
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

func newSpringDealServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&spring.SpringDeal{}, vnic).
		Require(func(v interface{}) string { return v.(*spring.SpringDeal).ListingId }, "ListingId").
		Require(func(v interface{}) string { return v.(*spring.SpringDeal).BidId }, "BidId").
		Require(func(v interface{}) string { return v.(*spring.SpringDeal).BuyerId }, "BuyerId").
		Require(func(v interface{}) string { return v.(*spring.SpringDeal).SellerId }, "SellerId").
		Require(func(v interface{}) string { return v.(*spring.SpringDeal).ItemTitle }, "ItemTitle").
		Money(func(v interface{}) *l8common.Money { return v.(*spring.SpringDeal).AgreedPrice }, "AgreedPrice").
		Money(func(v interface{}) *l8common.Money { return v.(*spring.SpringDeal).TotalAmount }, "TotalAmount").
		Enum(func(v interface{}) int32 { return int32(v.(*spring.SpringDeal).Status) }, spring.SpringDealStatus_name, "Status").
		StatusTransition(&common.StatusTransitionConfig{
			StatusGetter:  func(v interface{}) int32 { return int32(v.(*spring.SpringDeal).Status) },
			StatusSetter:  func(v interface{}, s int32) { v.(*spring.SpringDeal).Status = spring.SpringDealStatus(s) },
			FilterBuilder: func(v interface{}) interface{} { return &spring.SpringDeal{DealId: v.(*spring.SpringDeal).DealId} },
			ServiceName:   ServiceName,
			ServiceArea:   ServiceArea,
			InitialStatus: int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_PENDING_PAYMENT),
			Transitions: map[int32][]int32{
				int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_PENDING_PAYMENT): {
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_PAID),
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_CANCELLED),
				},
				int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_PAID): {
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_SHIPPED),
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_CANCELLED),
				},
				int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_SHIPPED): {
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_DELIVERED),
				},
				int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_DELIVERED): {
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_COMPLETED),
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_DISPUTED),
				},
				int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_DISPUTED): {
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_COMPLETED),
					int32(spring.SpringDealStatus_SPRING_DEAL_STATUS_REFUNDED),
				},
			},
			StatusNames: map[int32]string{
				0: "Unspecified",
				1: "Pending Payment",
				2: "Paid",
				3: "Shipped",
				4: "Delivered",
				5: "Completed",
				6: "Disputed",
				7: "Cancelled",
				8: "Refunded",
			},
		}).
		Build()
}
