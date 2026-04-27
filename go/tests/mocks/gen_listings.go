package mocks

import (
	"fmt"
	"math/rand"

	l8common "github.com/saichler/l8common/go/types/l8common"
	l8m "github.com/saichler/l8common/go/mocks"
	spring "github.com/saichler/l8spring/go/types/spring"
)

func generateListings(store *MockDataStore) []*spring.SpringListing {
	count := len(listingTitles)
	listings := make([]*spring.SpringListing, 0, count)

	for i := 0; i < count; i++ {
		userName := l8m.PickRef(store.UserNames, i)
		desc := listingDescriptions[i%len(listingDescriptions)]

		// Status distribution: 60% active, 15% draft, 10% fulfilled, 10% expired, 5% cancelled
		var status spring.SpringListingStatus
		switch {
		case i < count*60/100:
			status = spring.SpringListingStatus_SPRING_LISTING_STATUS_ACTIVE
		case i < count*75/100:
			status = spring.SpringListingStatus_SPRING_LISTING_STATUS_DRAFT
		case i < count*85/100:
			status = spring.SpringListingStatus_SPRING_LISTING_STATUS_FULFILLED
		case i < count*95/100:
			status = spring.SpringListingStatus_SPRING_LISTING_STATUS_EXPIRED
		default:
			status = spring.SpringListingStatus_SPRING_LISTING_STATUS_CANCELLED
		}

		conditions := []spring.SpringItemCondition{
			spring.SpringItemCondition_SPRING_ITEM_CONDITION_NEW,
			spring.SpringItemCondition_SPRING_ITEM_CONDITION_LIKE_NEW,
			spring.SpringItemCondition_SPRING_ITEM_CONDITION_GOOD,
			spring.SpringItemCondition_SPRING_ITEM_CONDITION_ANY,
		}

		shippingPrefs := []spring.SpringShippingPreference{
			spring.SpringShippingPreference_SPRING_SHIPPING_PREFERENCE_SHIPPING,
			spring.SpringShippingPreference_SPRING_SHIPPING_PREFERENCE_LOCAL_PICKUP,
			spring.SpringShippingPreference_SPRING_SHIPPING_PREFERENCE_EITHER,
		}

		listing := &spring.SpringListing{
			ListingId:          l8m.GenID("LST", i),
			BuyerId:            fmt.Sprintf("user-%03d", (i%10)+1),
			BuyerDisplayName:   userName,
			Title:              listingTitles[i],
			Description:        desc,
			CategoryId:         l8m.PickRef(store.CategoryIDs, i),
			DesiredCondition:   conditions[i%len(conditions)],
			MaxPrice:           &l8common.Money{Amount: int64(rand.Intn(200000) + 2000), CurrencyId: "USD"},
			Quantity:           int32(rand.Intn(3) + 1),
			Status:             status,
			ExpiresAt:          l8m.RandomFutureDate(3, 30),
			Location:           cities[i%len(cities)],
			ShippingPreference: shippingPrefs[i%len(shippingPrefs)],
			Tags:               []string{"popular", categoryNames[i%len(categoryNames)]},
			ViewCount:          int64(rand.Intn(500)),
			BidCount:           int32(rand.Intn(10)),
			AuditInfo:          l8m.CreateAuditInfo(),
		}

		listings = append(listings, listing)
	}
	return listings
}
