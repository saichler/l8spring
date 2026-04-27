package mocks

import (
	"fmt"

	l8m "github.com/saichler/l8common/go/mocks"
	spring "github.com/saichler/l8spring/go/types/spring"
)

const springPrefix = "/spring/10"

func RunAllPhases(client *SpringClient, store *MockDataStore) error {
	initUserNames(store)

	if err := runPhase1Categories(client, store); err != nil {
		return err
	}
	if err := runPhase2Listings(client, store); err != nil {
		return err
	}
	if err := runPhase3Bids(client, store); err != nil {
		return err
	}
	if err := runPhase4Deals(client, store); err != nil {
		return err
	}
	if err := runPhase5Reviews(client, store); err != nil {
		return err
	}
	return nil
}

func initUserNames(store *MockDataStore) {
	store.UserNames = make([]string, 20)
	for i := 0; i < 20; i++ {
		store.UserNames[i] = l8m.RandomName()
	}
}

func runPhase1Categories(client *SpringClient, store *MockDataStore) error {
	fmt.Println("  Phase 1: Categories")
	categories := generateCategories()
	_, err := client.Post(springPrefix+"/Category", &spring.SpringCategoryList{List: categories})
	if err != nil {
		return fmt.Errorf("failed to post categories: %w", err)
	}
	for _, c := range categories {
		store.CategoryIDs = append(store.CategoryIDs, c.CategoryId)
	}
	fmt.Printf("    Created %d categories\n", len(categories))
	return nil
}

func runPhase2Listings(client *SpringClient, store *MockDataStore) error {
	fmt.Println("  Phase 2: Listings")
	listings := generateListings(store)
	_, err := client.Post(springPrefix+"/Listing", &spring.SpringListingList{List: listings})
	if err != nil {
		return fmt.Errorf("failed to post listings: %w", err)
	}
	for _, l := range listings {
		store.ListingIDs = append(store.ListingIDs, l.ListingId)
		if l.Status == spring.SpringListingStatus_SPRING_LISTING_STATUS_ACTIVE {
			store.ActiveListingIDs = append(store.ActiveListingIDs, l.ListingId)
		}
	}
	fmt.Printf("    Created %d listings (%d active)\n", len(listings), len(store.ActiveListingIDs))
	return nil
}

func runPhase3Bids(client *SpringClient, store *MockDataStore) error {
	fmt.Println("  Phase 3: Bids")
	bids := generateBids(store)
	_, err := client.Post(springPrefix+"/Bid", &spring.SpringBidList{List: bids})
	if err != nil {
		return fmt.Errorf("failed to post bids: %w", err)
	}
	for _, b := range bids {
		store.BidIDs = append(store.BidIDs, b.BidId)
	}
	fmt.Printf("    Created %d bids\n", len(bids))
	return nil
}

func runPhase4Deals(client *SpringClient, store *MockDataStore) error {
	fmt.Println("  Phase 4: Deals")
	deals := generateDeals(store)
	_, err := client.Post(springPrefix+"/Deal", &spring.SpringDealList{List: deals})
	if err != nil {
		return fmt.Errorf("failed to post deals: %w", err)
	}
	for _, d := range deals {
		store.DealIDs = append(store.DealIDs, d.DealId)
		if d.Status == spring.SpringDealStatus_SPRING_DEAL_STATUS_COMPLETED {
			store.CompletedDealIDs = append(store.CompletedDealIDs, d.DealId)
		}
	}
	fmt.Printf("    Created %d deals (%d completed)\n", len(deals), len(store.CompletedDealIDs))
	return nil
}

func runPhase5Reviews(client *SpringClient, store *MockDataStore) error {
	fmt.Println("  Phase 5: Reviews")
	reviews := generateReviews(store)
	_, err := client.Post(springPrefix+"/Review", &spring.SpringReviewList{List: reviews})
	if err != nil {
		return fmt.Errorf("failed to post reviews: %w", err)
	}
	for _, r := range reviews {
		store.ReviewIDs = append(store.ReviewIDs, r.ReviewId)
	}
	fmt.Printf("    Created %d reviews\n", len(reviews))
	return nil
}
