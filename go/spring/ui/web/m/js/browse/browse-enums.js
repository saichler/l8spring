(function() {
    'use strict';
    window.MobileBrowse = window.MobileBrowse || {};
    var factory = window.Layer8EnumFactory;
    var createStatusRenderer = Layer8DRenderers.createStatusRenderer;

    var LISTING_STATUS = factory.create([
        ['Unspecified', null, ''],
        ['Draft', 'draft', 'layer8d-status-pending'],
        ['Active', 'active', 'layer8d-status-active'],
        ['Fulfilled', 'fulfilled', 'layer8d-status-completed'],
        ['Expired', 'expired', 'layer8d-status-inactive'],
        ['Cancelled', 'cancelled', 'layer8d-status-inactive']
    ]);

    var ITEM_CONDITION = factory.create([
        ['Unspecified', null, ''],
        ['New', 'new', 'layer8d-status-active'],
        ['Like New', 'like-new', 'layer8d-status-active'],
        ['Good', 'good', 'layer8d-status-info'],
        ['Fair', 'fair', 'layer8d-status-pending'],
        ['Any', 'any', '']
    ]);

    var SHIPPING_PREFERENCE = factory.simple([
        'Unspecified', 'Local Pickup', 'Shipping Only', 'Either'
    ]);

    MobileBrowse.enums = {
        LISTING_STATUS: LISTING_STATUS.enum,
        LISTING_STATUS_VALUES: LISTING_STATUS.values,
        LISTING_STATUS_CLASSES: LISTING_STATUS.classes,
        ITEM_CONDITION: ITEM_CONDITION.enum,
        ITEM_CONDITION_VALUES: ITEM_CONDITION.values,
        ITEM_CONDITION_CLASSES: ITEM_CONDITION.classes,
        SHIPPING_PREFERENCE: SHIPPING_PREFERENCE.enum
    };

    MobileBrowse.render = {};
    MobileBrowse.render.listingStatus = createStatusRenderer(LISTING_STATUS.enum, LISTING_STATUS.classes);
    MobileBrowse.render.itemCondition = createStatusRenderer(ITEM_CONDITION.enum, ITEM_CONDITION.classes);
    MobileBrowse.render.shippingPreference = function(value) {
        return Layer8DRenderers.renderEnum(value, SHIPPING_PREFERENCE.enum);
    };
})();
