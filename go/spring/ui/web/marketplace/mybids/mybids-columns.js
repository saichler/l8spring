(function() {
    'use strict';
    var enums = MyBids.enums;
    var render = MyBids.render;
    var col = window.Layer8ColumnFactory;

    MyBids.columns = {
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

    MyBids.primaryKeys = {
        SpringBid: 'bidId'
    };
})();
