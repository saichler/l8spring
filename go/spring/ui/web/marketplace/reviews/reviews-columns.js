(function() {
    'use strict';
    var enums = Reviews.enums;
    var render = Reviews.render;
    var col = window.Layer8ColumnFactory;

    Reviews.columns = {
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

    Reviews.primaryKeys = {
        SpringReview: 'reviewId'
    };
})();
