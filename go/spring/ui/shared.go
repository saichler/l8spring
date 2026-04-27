package ui

import (
	"github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

func RegisterTypes(resources ifs.IResources) {
	common.RegisterType(resources, &spring.SpringCategory{}, &spring.SpringCategoryList{}, "CategoryId")
	common.RegisterType(resources, &spring.SpringListing{}, &spring.SpringListingList{}, "ListingId")
	common.RegisterType(resources, &spring.SpringBid{}, &spring.SpringBidList{}, "BidId")
	common.RegisterType(resources, &spring.SpringDeal{}, &spring.SpringDealList{}, "DealId")
	common.RegisterType(resources, &spring.SpringReview{}, &spring.SpringReviewList{}, "ReviewId")
}
