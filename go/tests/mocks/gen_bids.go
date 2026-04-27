package mocks

import (
	"fmt"
	"math/rand"

	l8common "github.com/saichler/l8common/go/types/l8common"
	l8m "github.com/saichler/l8common/go/mocks"
	spring "github.com/saichler/l8spring/go/types/spring"
)

func generateBids(store *MockDataStore) []*spring.SpringBid {
	count := len(store.ActiveListingIDs) * 2
	if count > 50 {
		count = 50
	}
	bids := make([]*spring.SpringBid, 0, count)

	conditions := []spring.SpringItemCondition{
		spring.SpringItemCondition_SPRING_ITEM_CONDITION_NEW,
		spring.SpringItemCondition_SPRING_ITEM_CONDITION_LIKE_NEW,
		spring.SpringItemCondition_SPRING_ITEM_CONDITION_GOOD,
	}

	for i := 0; i < count; i++ {
		userName := l8m.PickRef(store.UserNames, i+5)

		var status spring.SpringBidStatus
		switch {
		case i < count*50/100:
			status = spring.SpringBidStatus_SPRING_BID_STATUS_PENDING
		case i < count*70/100:
			status = spring.SpringBidStatus_SPRING_BID_STATUS_ACCEPTED
		case i < count*85/100:
			status = spring.SpringBidStatus_SPRING_BID_STATUS_REJECTED
		default:
			status = spring.SpringBidStatus_SPRING_BID_STATUS_WITHDRAWN
		}

		bid := &spring.SpringBid{
			BidId:             l8m.GenID("BID", i),
			ListingId:         l8m.PickRef(store.ActiveListingIDs, i),
			SellerId:          fmt.Sprintf("user-%03d", (i%10)+1),
			SellerDisplayName: userName,
			Price:             &l8common.Money{Amount: int64(rand.Intn(150000) + 1000), CurrencyId: "USD"},
			ItemCondition:     conditions[i%len(conditions)],
			Description:       bidMessages[i%len(bidMessages)],
			EstimatedShipDays: int32(rand.Intn(7) + 1),
			ShippingCost:      &l8common.Money{Amount: int64(rand.Intn(3000) + 500), CurrencyId: "USD"},
			Status:            status,
			SellerRating:      float32(rand.Intn(20)+30) / 10.0,
			AuditInfo:         l8m.CreateAuditInfo(),
		}

		bids = append(bids, bid)
	}
	return bids
}
