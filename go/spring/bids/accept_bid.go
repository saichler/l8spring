package bids

import (
	"fmt"
	l8common "github.com/saichler/l8common/go/types/l8common"
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

func afterBidUpdate(entity interface{}, action ifs.Action, vnic ifs.IVNic) error {
	if action != ifs.PUT {
		return nil
	}
	bid := entity.(*spring.SpringBid)
	if bid.Status != spring.SpringBidStatus_SPRING_BID_STATUS_ACCEPTED {
		return nil
	}

	listing, err := fetchListing(bid.ListingId, vnic)
	if err != nil {
		return fmt.Errorf("failed to fetch listing for deal creation: %w", err)
	}

	deal := buildDeal(bid, listing)
	_, err = common.PostEntity("Deal", 10, deal, vnic)
	if err != nil {
		return fmt.Errorf("failed to create deal: %w", err)
	}

	listing.Status = spring.SpringListingStatus_SPRING_LISTING_STATUS_FULFILLED
	err = common.PutEntity("Listing", 10, listing, vnic)
	if err != nil {
		return fmt.Errorf("failed to update listing to fulfilled: %w", err)
	}

	return nil
}

func fetchListing(listingId string, vnic ifs.IVNic) (*spring.SpringListing, error) {
	filter := &spring.SpringListing{ListingId: listingId}
	result, err := common.GetEntity("Listing", 10, filter, vnic)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("listing %s not found", listingId)
	}
	return result.(*spring.SpringListing), nil
}

func buildDeal(bid *spring.SpringBid, listing *spring.SpringListing) *spring.SpringDeal {
	var totalAmount int64
	if bid.Price != nil {
		totalAmount = bid.Price.Amount
	}
	if bid.ShippingCost != nil {
		totalAmount += bid.ShippingCost.Amount
	}

	currencyCode := "USD"
	if bid.Price != nil && bid.Price.CurrencyId != "" {
		currencyCode = bid.Price.CurrencyId
	}

	return &spring.SpringDeal{
		ListingId:         listing.ListingId,
		BidId:             bid.BidId,
		BuyerId:           listing.BuyerId,
		SellerId:          bid.SellerId,
		BuyerDisplayName:  listing.BuyerDisplayName,
		SellerDisplayName: bid.SellerDisplayName,
		ItemTitle:         listing.Title,
		AgreedPrice:       bid.Price,
		ShippingCost:      bid.ShippingCost,
		TotalAmount: &l8common.Money{
			Amount:       totalAmount,
			CurrencyId: currencyCode,
		},
		Status: spring.SpringDealStatus_SPRING_DEAL_STATUS_PENDING_PAYMENT,
	}
}
