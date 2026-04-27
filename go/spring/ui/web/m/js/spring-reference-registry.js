(function() {
    'use strict';
    var ref = window.Layer8RefFactory;

    window.Layer8MReferenceRegistrySpring = {
        ...ref.simple('SpringCategory', 'categoryId', 'name', 'Category'),
        ...ref.simple('SpringListing', 'listingId', 'title', 'Listing'),
        ...ref.simple('SpringBid', 'bidId', 'bidId', 'Bid'),
        ...ref.simple('SpringDeal', 'dealId', 'itemTitle', 'Deal'),
        ...ref.simple('SpringReview', 'reviewId', 'title', 'Review')
    };

    Layer8MReferenceRegistry.register(window.Layer8MReferenceRegistrySpring);
})();
