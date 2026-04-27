(function() {
    'use strict';
    var ref = window.Layer8RefFactory;

    Layer8DReferenceRegistry.register({
        ...ref.simple('SpringCategory', 'categoryId', 'name', 'Category'),
        ...ref.simple('SpringListing', 'listingId', 'title', 'Listing'),
        ...ref.simple('SpringBid', 'bidId', 'bidId', 'Bid'),
        ...ref.simple('SpringDeal', 'dealId', 'itemTitle', 'Deal'),
        ...ref.simple('SpringReview', 'reviewId', 'title', 'Review')
    });
})();
