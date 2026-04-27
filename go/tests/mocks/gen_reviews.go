package mocks

import (
	"fmt"
	"math/rand"

	l8m "github.com/saichler/l8common/go/mocks"
	spring "github.com/saichler/l8spring/go/types/spring"
)

func generateReviews(store *MockDataStore) []*spring.SpringReview {
	count := len(store.CompletedDealIDs)
	if count > 30 {
		count = 30
	}
	reviews := make([]*spring.SpringReview, 0, count*2)

	for i := 0; i < count; i++ {
		buyerName := l8m.PickRef(store.UserNames, i)
		sellerName := l8m.PickRef(store.UserNames, i+5)

		var status spring.SpringReviewStatus
		if i < count*80/100 {
			status = spring.SpringReviewStatus_SPRING_REVIEW_STATUS_PUBLISHED
		} else {
			status = spring.SpringReviewStatus_SPRING_REVIEW_STATUS_PENDING
		}

		// Buyer reviews seller
		buyerReview := &spring.SpringReview{
			ReviewId:            l8m.GenID("REV", i*2),
			DealId:              l8m.PickRef(store.CompletedDealIDs, i),
			ReviewerId:          fmt.Sprintf("user-%03d", (i%10)+1),
			RevieweeId:          fmt.Sprintf("user-%03d", ((i+5)%10)+1),
			ReviewerDisplayName: buyerName,
			ReviewerRole:        spring.SpringReviewerRole_SPRING_REVIEWER_ROLE_BUYER,
			Rating:              int32(rand.Intn(3) + 3),
			Title:               reviewTitles[i%len(reviewTitles)],
			Content:             reviewContents[i%len(reviewContents)],
			Status:              status,
			AuditInfo:           l8m.CreateAuditInfo(),
		}
		reviews = append(reviews, buyerReview)

		// Seller reviews buyer
		sellerReview := &spring.SpringReview{
			ReviewId:            l8m.GenID("REV", i*2+1),
			DealId:              l8m.PickRef(store.CompletedDealIDs, i),
			ReviewerId:          fmt.Sprintf("user-%03d", ((i+5)%10)+1),
			RevieweeId:          fmt.Sprintf("user-%03d", (i%10)+1),
			ReviewerDisplayName: sellerName,
			ReviewerRole:        spring.SpringReviewerRole_SPRING_REVIEWER_ROLE_SELLER,
			Rating:              int32(rand.Intn(3) + 3),
			Title:               reviewTitles[(i+3)%len(reviewTitles)],
			Content:             reviewContents[(i+3)%len(reviewContents)],
			Status:              status,
			AuditInfo:           l8m.CreateAuditInfo(),
		}
		reviews = append(reviews, sellerReview)
	}
	return reviews
}
