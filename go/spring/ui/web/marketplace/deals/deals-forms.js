(function() {
    'use strict';
    var enums = Deals.enums;
    var f = window.Layer8FormFactory;

    Deals.forms = {
        SpringDeal: f.form('Deal', [
            f.section('Deal Details', [
                ...f.text('itemTitle', 'Item Title', true),
                ...f.reference('listingId', 'Listing', 'SpringListing', true),
                ...f.reference('bidId', 'Bid', 'SpringBid', true),
                ...f.text('buyerDisplayName', 'Buyer', false),
                ...f.text('sellerDisplayName', 'Seller', false)
            ]),
            f.section('Financials', [
                ...f.money('agreedPrice', 'Agreed Price'),
                ...f.money('shippingCost', 'Shipping Cost'),
                ...f.money('totalAmount', 'Total Amount'),
                ...f.text('paymentReference', 'Payment Reference')
            ]),
            f.section('Shipping', [
                ...f.select('status', 'Status', enums.DEAL_STATUS),
                ...f.text('trackingNumber', 'Tracking Number'),
                ...f.text('shippingCarrier', 'Shipping Carrier'),
                ...f.date('shippedAt', 'Shipped At'),
                ...f.date('deliveredAt', 'Delivered At'),
                ...f.date('completedAt', 'Completed At')
            ]),
            f.section('Messages', [
                ...f.inlineTable('messages', 'Messages', [
                    { key: 'messageId', label: 'ID', hidden: true },
                    { key: 'senderId', label: 'Sender', type: 'text' },
                    { key: 'senderDisplayName', label: 'Name', type: 'text' },
                    { key: 'content', label: 'Message', type: 'text' },
                    { key: 'sentAt', label: 'Sent', type: 'date' }
                ])
            ]),
            f.section('Audit', [
                ...f.audit()
            ])
        ])
    };
})();
