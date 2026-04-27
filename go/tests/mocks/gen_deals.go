package mocks

import (
	"fmt"
	"math/rand"

	l8common "github.com/saichler/l8common/go/types/l8common"
	l8m "github.com/saichler/l8common/go/mocks"
	spring "github.com/saichler/l8spring/go/types/spring"
)

func generateDeals(store *MockDataStore) []*spring.SpringDeal {
	count := 20
	deals := make([]*spring.SpringDeal, 0, count)

	for i := 0; i < count; i++ {
		buyerName := l8m.PickRef(store.UserNames, i)
		sellerName := l8m.PickRef(store.UserNames, i+5)
		agreedPrice := int64(rand.Intn(180000) + 2000)
		shippingCost := int64(rand.Intn(3000) + 500)

		var status spring.SpringDealStatus
		switch {
		case i < count*10/100:
			status = spring.SpringDealStatus_SPRING_DEAL_STATUS_PENDING_PAYMENT
		case i < count*30/100:
			status = spring.SpringDealStatus_SPRING_DEAL_STATUS_PAID
		case i < count*45/100:
			status = spring.SpringDealStatus_SPRING_DEAL_STATUS_SHIPPED
		case i < count*55/100:
			status = spring.SpringDealStatus_SPRING_DEAL_STATUS_DELIVERED
		case i < count*80/100:
			status = spring.SpringDealStatus_SPRING_DEAL_STATUS_COMPLETED
		case i < count*90/100:
			status = spring.SpringDealStatus_SPRING_DEAL_STATUS_CANCELLED
		default:
			status = spring.SpringDealStatus_SPRING_DEAL_STATUS_DISPUTED
		}

		deal := &spring.SpringDeal{
			DealId:            l8m.GenID("DEAL", i),
			ListingId:         l8m.PickRef(store.ListingIDs, i),
			BidId:             l8m.PickRef(store.BidIDs, i),
			BuyerId:           fmt.Sprintf("user-%03d", (i%10)+1),
			SellerId:          fmt.Sprintf("user-%03d", ((i+5)%10)+1),
			BuyerDisplayName:  buyerName,
			SellerDisplayName: sellerName,
			ItemTitle:         listingTitles[i%len(listingTitles)],
			AgreedPrice:       &l8common.Money{Amount: agreedPrice, CurrencyId: "USD"},
			ShippingCost:      &l8common.Money{Amount: shippingCost, CurrencyId: "USD"},
			TotalAmount:       &l8common.Money{Amount: agreedPrice + shippingCost, CurrencyId: "USD"},
			Status:            status,
			AuditInfo:         l8m.CreateAuditInfo(),
		}

		// Add shipping info for shipped/delivered/completed deals
		if status >= spring.SpringDealStatus_SPRING_DEAL_STATUS_SHIPPED {
			deal.TrackingNumber = fmt.Sprintf("1Z%09d", rand.Intn(999999999))
			deal.ShippingCarrier = shippingCarriers[i%len(shippingCarriers)]
			deal.ShippedAt = l8m.RandomPastDate(1, 15)
		}
		if status >= spring.SpringDealStatus_SPRING_DEAL_STATUS_DELIVERED {
			deal.DeliveredAt = l8m.RandomPastDate(0, 10)
		}
		if status == spring.SpringDealStatus_SPRING_DEAL_STATUS_COMPLETED {
			deal.CompletedAt = l8m.RandomPastDate(0, 5)
		}
		if status == spring.SpringDealStatus_SPRING_DEAL_STATUS_PAID {
			deal.PaymentReference = fmt.Sprintf("PAY-%06d", rand.Intn(999999))
		}

		// Add messages
		msgCount := rand.Intn(3) + 1
		deal.Messages = make([]*spring.SpringDealMessage, 0, msgCount)
		for j := 0; j < msgCount; j++ {
			msg := &spring.SpringDealMessage{
				MessageId:         fmt.Sprintf("MSG-%03d-%02d", i+1, j+1),
				SenderId:          deal.BuyerId,
				SenderDisplayName: buyerName,
				Content:           dealMessages[(i+j)%len(dealMessages)],
				SentAt:            l8m.RandomPastDate(0, 20),
			}
			if j%2 == 1 {
				msg.SenderId = deal.SellerId
				msg.SenderDisplayName = sellerName
			}
			deal.Messages = append(deal.Messages, msg)
		}

		deals = append(deals, deal)
	}
	return deals
}
