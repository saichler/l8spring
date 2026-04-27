(function() {
    'use strict';
    var enums = MobileMyBids.enums;
    var render = MobileMyBids.render;
    var col = window.Layer8ColumnFactory;

    MobileMyBids.columns = {
        SpringBid: [
            ...col.id('bidId'),
            ...col.col('listingId', 'Listing'),
            ...col.money('price', 'Price'),
            ...col.status('status', 'Status', enums.BID_STATUS_VALUES, render.bidStatus),
            ...col.enum('itemCondition', 'Condition', enums.ITEM_CONDITION_VALUES, render.itemCondition),
            ...col.col('description', 'Description'),
            ...col.money('shippingCost', 'Shipping'),
            ...col.number('estimatedShipDays', 'Ship Days'),
            ...col.col('sellerDisplayName', 'Seller'),
            ...col.date('auditInfo.createdDate', 'Submitted')
        ]
    };

    MobileMyBids.columns.SpringBid[1].primary = true;
    MobileMyBids.columns.SpringBid[2].secondary = true;

    MobileMyBids.primaryKeys = {
        SpringBid: 'bidId'
    };
})();
