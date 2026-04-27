(function() {
    'use strict';
    var enums = MobileBrowse.enums;
    var f = window.Layer8FormFactory;

    MobileBrowse.forms = {
        SpringListing: f.form('Listing', [
            f.section('Listing Details', [
                ...f.text('title', 'Title', true),
                ...f.textarea('description', 'Description', true),
                ...f.reference('categoryId', 'Category', 'SpringCategory', true),
                ...f.select('desiredCondition', 'Desired Condition', enums.ITEM_CONDITION),
                ...f.money('maxPrice', 'Max Price'),
                ...f.number('quantity', 'Quantity')
            ]),
            f.section('Shipping & Location', [
                ...f.text('location', 'Location'),
                ...f.select('shippingPreference', 'Shipping Preference', enums.SHIPPING_PREFERENCE),
                ...f.date('expiresAt', 'Expires At')
            ]),
            f.section('Status', [
                ...f.select('status', 'Status', enums.LISTING_STATUS),
                ...f.text('buyerDisplayName', 'Buyer', false),
                ...f.number('viewCount', 'Views'),
                ...f.number('bidCount', 'Bids')
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
        ]),
        SpringCategory: f.form('Category', [
            f.section('Category Details', [
                ...f.text('name', 'Name', true),
                ...f.textarea('description', 'Description'),
                ...f.reference('parentCategoryId', 'Parent Category', 'SpringCategory'),
                ...f.checkbox('isActive', 'Active'),
                ...f.number('sortOrder', 'Sort Order')
            ]),
            f.section('Audit', [
                ...f.audit()
            ])
        ])
    };
})();
