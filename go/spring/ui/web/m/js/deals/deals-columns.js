(function() {
    'use strict';
    var enums = MobileDeals.enums;
    var render = MobileDeals.render;
    var col = window.Layer8ColumnFactory;

    MobileDeals.columns = {
        SpringDeal: [
            ...col.id('dealId'),
            ...col.col('itemTitle', 'Item'),
            ...col.money('agreedPrice', 'Price'),
            ...col.money('shippingCost', 'Shipping'),
            ...col.money('totalAmount', 'Total'),
            ...col.status('status', 'Status', enums.DEAL_STATUS_VALUES, render.dealStatus),
            ...col.col('buyerDisplayName', 'Buyer'),
            ...col.col('sellerDisplayName', 'Seller'),
            ...col.col('trackingNumber', 'Tracking'),
            ...col.date('auditInfo.createdDate', 'Created')
        ]
    };

    MobileDeals.columns.SpringDeal[1].primary = true;
    MobileDeals.columns.SpringDeal[2].secondary = true;

    MobileDeals.primaryKeys = {
        SpringDeal: 'dealId'
    };
})();
