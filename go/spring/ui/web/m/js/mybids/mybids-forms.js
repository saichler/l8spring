(function() {
    'use strict';
    var enums = MobileMyBids.enums;
    var f = window.Layer8FormFactory;

    MobileMyBids.forms = {
        SpringBid: f.form('Bid', [
            f.section('Bid Details', [
                ...f.reference('listingId', 'Listing', 'SpringListing', true),
                ...f.money('price', 'Price'),
                ...f.textarea('description', 'Description', true),
                ...f.select('itemCondition', 'Item Condition', enums.ITEM_CONDITION)
            ]),
            f.section('Shipping', [
                ...f.money('shippingCost', 'Shipping Cost'),
                ...f.number('estimatedShipDays', 'Estimated Ship Days')
            ]),
            f.section('Status', [
                ...f.select('status', 'Status', enums.BID_STATUS),
                ...f.text('sellerDisplayName', 'Seller', false)
            ]),
            f.section('Images', [
                ...f.inlineTable('images', 'Images', [
                    { key: 'imageId', label: 'ID', hidden: true },
                    { key: 'storagePath', label: 'Image', type: 'file' },
                    { key: 'caption', label: 'Caption', type: 'text' },
                    { key: 'sortOrder', label: 'Order', type: 'number' }
                ])
            ]),
            f.section('Audit', [
                ...f.audit()
            ])
        ])
    };
})();
