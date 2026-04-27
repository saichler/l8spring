package mocks

type MockDataStore struct {
	// Phase 1: Categories
	CategoryIDs []string

	// Phase 2: Listings
	ListingIDs       []string
	ActiveListingIDs []string

	// Phase 3: Bids
	BidIDs []string

	// Phase 4: Deals
	DealIDs          []string
	CompletedDealIDs []string

	// Phase 5: Reviews
	ReviewIDs []string

	// User display names for reference
	UserNames []string
}
