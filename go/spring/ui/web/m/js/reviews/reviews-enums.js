(function() {
    'use strict';
    window.MobileReviews = window.MobileReviews || {};
    var factory = window.Layer8EnumFactory;
    var createStatusRenderer = Layer8DRenderers.createStatusRenderer;
    var renderEnum = Layer8DRenderers.renderEnum;

    var REVIEW_STATUS = factory.create([
        ['Unspecified', null, ''],
        ['Pending', 'pending', 'layer8d-status-pending'],
        ['Published', 'published', 'layer8d-status-active'],
        ['Flagged', 'flagged', 'layer8d-status-warning'],
        ['Removed', 'removed', 'layer8d-status-inactive']
    ]);

    var REVIEWER_ROLE = factory.simple([
        'Unspecified', 'Buyer', 'Seller'
    ]);

    MobileReviews.enums = {
        REVIEW_STATUS: REVIEW_STATUS.enum,
        REVIEW_STATUS_VALUES: REVIEW_STATUS.values,
        REVIEW_STATUS_CLASSES: REVIEW_STATUS.classes,
        REVIEWER_ROLE: REVIEWER_ROLE.enum
    };

    MobileReviews.render = {};
    MobileReviews.render.reviewStatus = createStatusRenderer(REVIEW_STATUS.enum, REVIEW_STATUS.classes);
    MobileReviews.render.reviewerRole = function(value) {
        return renderEnum(value, REVIEWER_ROLE.enum);
    };
})();
