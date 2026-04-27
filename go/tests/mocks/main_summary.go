package mocks

import "fmt"

func PrintSummary(store *MockDataStore) {
	fmt.Println("\n========================================")
	fmt.Println("  Spring Mock Data Summary")
	fmt.Println("========================================")
	fmt.Printf("  Categories:       %d\n", len(store.CategoryIDs))
	fmt.Printf("  Listings:         %d (%d active)\n", len(store.ListingIDs), len(store.ActiveListingIDs))
	fmt.Printf("  Bids:             %d\n", len(store.BidIDs))
	fmt.Printf("  Deals:            %d (%d completed)\n", len(store.DealIDs), len(store.CompletedDealIDs))
	fmt.Printf("  Reviews:          %d\n", len(store.ReviewIDs))
	fmt.Println("========================================")
}
