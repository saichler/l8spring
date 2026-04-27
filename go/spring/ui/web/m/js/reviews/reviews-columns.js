(function() {
    'use strict';
    var enums = MobileReviews.enums;
    var render = MobileReviews.render;
    var col = window.Layer8ColumnFactory;

    MobileReviews.columns = {
        SpringReview: [
            ...col.id('reviewId'),
            ...col.col('title', 'Title'),
            ...col.number('rating', 'Rating'),
            ...col.col('content', 'Content'),
            ...col.status('status', 'Status', enums.REVIEW_STATUS_VALUES, render.reviewStatus),
            ...col.enum('reviewerRole', 'Role', null, render.reviewerRole),
            ...col.col('reviewerDisplayName', 'Reviewer'),
            ...col.col('revieweeDisplayName', 'Reviewee'),
            ...col.date('auditInfo.createdDate', 'Date')
        ]
    };

    MobileReviews.columns.SpringReview[1].primary = true;
    MobileReviews.columns.SpringReview[2].secondary = true;

    MobileReviews.primaryKeys = {
        SpringReview: 'reviewId'
    };
})();
