(function() {
    'use strict';
    var enums = Browse.enums;
    var render = Browse.render;
    var col = window.Layer8ColumnFactory;

    Browse.columns = {
        SpringListing: [
            ...col.id('listingId'),
            ...col.col('title', 'Title'),
            ...col.money('maxPrice', 'Max Price'),
            ...col.status('status', 'Status', enums.LISTING_STATUS_VALUES, render.listingStatus),
            ...col.enum('desiredCondition', 'Condition', enums.ITEM_CONDITION_VALUES, render.itemCondition),
            ...col.col('location', 'Location'),
            ...col.col('buyerDisplayName', 'Buyer'),
            ...col.number('bidCount', 'Bids'),
            ...col.date('auditInfo.createdDate', 'Posted')
        ],
        SpringCategory: [
            ...col.id('categoryId'),
            ...col.col('name', 'Name'),
            ...col.col('description', 'Description'),
            ...col.boolean('isActive', 'Active'),
            ...col.number('sortOrder', 'Sort Order'),
            ...col.number('listingCount', 'Listings')
        ]
    };

    Browse.primaryKeys = {
        SpringListing: 'listingId',
        SpringCategory: 'categoryId'
    };
})();
