(function() {
    'use strict';
    var enums = Deals.enums;
    var render = Deals.render;
    var col = window.Layer8ColumnFactory;

    Deals.columns = {
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

    Deals.primaryKeys = {
        SpringDeal: 'dealId'
    };
})();
