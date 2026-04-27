package reviews

import (
	common "github.com/saichler/l8spring/go/spring/common"
	"github.com/saichler/l8spring/go/types/spring"
	"github.com/saichler/l8types/go/ifs"
)

func newSpringReviewServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&spring.SpringReview{}, vnic).
		Require(func(v interface{}) string { return v.(*spring.SpringReview).DealId }, "DealId").
		Require(func(v interface{}) string { return v.(*spring.SpringReview).ReviewerId }, "ReviewerId").
		Require(func(v interface{}) string { return v.(*spring.SpringReview).RevieweeId }, "RevieweeId").
		Require(func(v interface{}) string { return v.(*spring.SpringReview).Content }, "Content").
		Enum(func(v interface{}) int32 { return int32(v.(*spring.SpringReview).ReviewerRole) }, spring.SpringReviewerRole_name, "ReviewerRole").
		Enum(func(v interface{}) int32 { return int32(v.(*spring.SpringReview).Status) }, spring.SpringReviewStatus_name, "Status").
		StatusTransition(&common.StatusTransitionConfig{
			StatusGetter:  func(v interface{}) int32 { return int32(v.(*spring.SpringReview).Status) },
			StatusSetter:  func(v interface{}, s int32) { v.(*spring.SpringReview).Status = spring.SpringReviewStatus(s) },
			FilterBuilder: func(v interface{}) interface{} { return &spring.SpringReview{ReviewId: v.(*spring.SpringReview).ReviewId} },
			ServiceName:   ServiceName,
			ServiceArea:   ServiceArea,
			InitialStatus: int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_PENDING),
			Transitions: map[int32][]int32{
				int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_PENDING): {
					int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_PUBLISHED),
					int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_REMOVED),
				},
				int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_PUBLISHED): {
					int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_FLAGGED),
					int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_REMOVED),
				},
				int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_FLAGGED): {
					int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_PUBLISHED),
					int32(spring.SpringReviewStatus_SPRING_REVIEW_STATUS_REMOVED),
				},
			},
			StatusNames: map[int32]string{
				0: "Unspecified",
				1: "Pending",
				2: "Published",
				3: "Flagged",
				4: "Removed",
			},
		}).
		Build()
}
