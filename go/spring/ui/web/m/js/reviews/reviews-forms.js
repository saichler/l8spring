(function() {
    'use strict';
    var enums = MobileReviews.enums;
    var f = window.Layer8FormFactory;

    MobileReviews.forms = {
        SpringReview: f.form('Review', [
            f.section('Review Details', [
                ...f.reference('dealId', 'Deal', 'SpringDeal', true),
                ...f.text('title', 'Title'),
                ...f.number('rating', 'Rating (1-5)', true),
                ...f.textarea('content', 'Review Content', true),
                ...f.select('reviewerRole', 'Role', enums.REVIEWER_ROLE)
            ]),
            f.section('Status', [
                ...f.select('status', 'Status', enums.REVIEW_STATUS),
                ...f.text('reviewerDisplayName', 'Reviewer', false),
                ...f.text('revieweeDisplayName', 'Reviewee', false)
            ]),
            f.section('Audit', [
                ...f.audit()
            ])
        ])
    };
})();
