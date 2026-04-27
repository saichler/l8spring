package tests

import (
	"fmt"
	"github.com/saichler/l8spring/go/tests/mocks"
	"github.com/saichler/l8types/go/ifs"
	"strings"
	"testing"
)

func testCRUD(t *testing.T, client *mocks.SpringClient) {
	testCRUDCategory(t, client)
	testCRUDListing(t, client)
	testCRUDBid(t, client)
	testCRUDReview(t, client)
}

func testCRUDCategory(t *testing.T, client *mocks.SpringClient) {
	catId := ifs.NewUuid()
	cat := map[string]interface{}{
		"categoryId":  catId,
		"name":        "CRUD Test Category",
		"description": "Created by CRUD test",
		"isActive":    true,
		"sortOrder":   99,
	}

	_, err := client.Post("/spring/10/Category", cat)
	if err != nil {
		t.Fatalf("POST Category failed: %v", err)
	}

	q := mocks.L8QueryText(fmt.Sprintf("select * from SpringCategory where categoryId=%s", catId))
	getResp, err := client.Get("/spring/10/Category", q)
	if err != nil {
		t.Fatalf("GET Category failed: %v", err)
	}
	if !strings.Contains(getResp, "CRUD Test Category") {
		t.Fatalf("GET Category did not return expected name")
	}

	cat["description"] = "Updated by CRUD test"
	_, err = client.Put("/spring/10/Category", cat)
	if err != nil {
		t.Fatalf("PUT Category failed: %v", err)
	}

	delQ := mocks.L8QueryText(fmt.Sprintf("select * from SpringCategory where categoryId=%s", catId))
	_, err = client.Delete("/spring/10/Category", delQ)
	if err != nil {
		t.Fatalf("DELETE Category failed: %v", err)
	}
}

func testCRUDListing(t *testing.T, client *mocks.SpringClient) {
	listingId := ifs.NewUuid()
	listing := map[string]interface{}{
		"listingId":        listingId,
		"buyerId":          "user-crud",
		"buyerDisplayName": "CRUD Buyer",
		"title":            "CRUD Test Listing",
		"description":      "Looking for test item",
		"status":           2,
		"quantity":         1,
	}

	_, err := client.Post("/spring/10/Listing", listing)
	if err != nil {
		t.Fatalf("POST Listing failed: %v", err)
	}

	q := mocks.L8QueryText(fmt.Sprintf("select * from SpringListing where listingId=%s", listingId))
	getResp, err := client.Get("/spring/10/Listing", q)
	if err != nil {
		t.Fatalf("GET Listing failed: %v", err)
	}
	if !strings.Contains(getResp, "CRUD Test Listing") {
		t.Fatalf("GET Listing did not return expected title")
	}

	listing["description"] = "Updated CRUD test listing"
	_, err = client.Put("/spring/10/Listing", listing)
	if err != nil {
		t.Fatalf("PUT Listing failed: %v", err)
	}

	delQ := mocks.L8QueryText(fmt.Sprintf("select * from SpringListing where listingId=%s", listingId))
	_, err = client.Delete("/spring/10/Listing", delQ)
	if err != nil {
		t.Fatalf("DELETE Listing failed: %v", err)
	}
}

func testCRUDBid(t *testing.T, client *mocks.SpringClient) {
	bidId := ifs.NewUuid()
	bid := map[string]interface{}{
		"bidId":             bidId,
		"listingId":         "LST-001",
		"sellerId":          "user-crud",
		"sellerDisplayName": "CRUD Seller",
		"description":       "CRUD Test Bid",
		"status":            1,
		"estimatedShipDays": 3,
	}

	_, err := client.Post("/spring/10/Bid", bid)
	if err != nil {
		t.Fatalf("POST Bid failed: %v", err)
	}

	q := mocks.L8QueryText(fmt.Sprintf("select * from SpringBid where bidId=%s", bidId))
	getResp, err := client.Get("/spring/10/Bid", q)
	if err != nil {
		t.Fatalf("GET Bid failed: %v", err)
	}
	if !strings.Contains(getResp, "CRUD Test Bid") {
		t.Fatalf("GET Bid did not return expected description")
	}

	bid["description"] = "Updated CRUD test bid"
	_, err = client.Put("/spring/10/Bid", bid)
	if err != nil {
		t.Fatalf("PUT Bid failed: %v", err)
	}

	delQ := mocks.L8QueryText(fmt.Sprintf("select * from SpringBid where bidId=%s", bidId))
	_, err = client.Delete("/spring/10/Bid", delQ)
	if err != nil {
		t.Fatalf("DELETE Bid failed: %v", err)
	}
}

func testCRUDReview(t *testing.T, client *mocks.SpringClient) {
	reviewId := ifs.NewUuid()
	review := map[string]interface{}{
		"reviewId":            reviewId,
		"dealId":              "DEAL-001",
		"reviewerId":          "user-crud",
		"revieweeId":          "user-crud-2",
		"reviewerDisplayName": "CRUD Reviewer",
		"rating":              5,
		"title":               "CRUD Test Review",
		"content":             "Test review content",
		"status":              2,
	}

	_, err := client.Post("/spring/10/Review", review)
	if err != nil {
		t.Fatalf("POST Review failed: %v", err)
	}

	q := mocks.L8QueryText(fmt.Sprintf("select * from SpringReview where reviewId=%s", reviewId))
	getResp, err := client.Get("/spring/10/Review", q)
	if err != nil {
		t.Fatalf("GET Review failed: %v", err)
	}
	if !strings.Contains(getResp, "CRUD Test Review") {
		t.Fatalf("GET Review did not return expected title")
	}

	review["content"] = "Updated test review content"
	_, err = client.Put("/spring/10/Review", review)
	if err != nil {
		t.Fatalf("PUT Review failed: %v", err)
	}

	delQ := mocks.L8QueryText(fmt.Sprintf("select * from SpringReview where reviewId=%s", reviewId))
	_, err = client.Delete("/spring/10/Review", delQ)
	if err != nil {
		t.Fatalf("DELETE Review failed: %v", err)
	}
}
