(function() {
    'use strict';
    window.MobileDeals = window.MobileDeals || {};
    var factory = window.Layer8EnumFactory;
    var createStatusRenderer = Layer8DRenderers.createStatusRenderer;

    var DEAL_STATUS = factory.create([
        ['Unspecified', null, ''],
        ['Pending Payment', 'pending-payment', 'layer8d-status-pending'],
        ['Paid', 'paid', 'layer8d-status-info'],
        ['Shipped', 'shipped', 'layer8d-status-info'],
        ['Delivered', 'delivered', 'layer8d-status-active'],
        ['Completed', 'completed', 'layer8d-status-completed'],
        ['Disputed', 'disputed', 'layer8d-status-warning'],
        ['Cancelled', 'cancelled', 'layer8d-status-inactive'],
        ['Refunded', 'refunded', 'layer8d-status-inactive']
    ]);

    var DISPUTE_STATUS = factory.create([
        ['Unspecified', null, ''],
        ['Open', 'open', 'layer8d-status-warning'],
        ['Under Review', 'under-review', 'layer8d-status-pending'],
        ['Resolved (Buyer)', 'resolved-buyer', 'layer8d-status-completed'],
        ['Resolved (Seller)', 'resolved-seller', 'layer8d-status-completed'],
        ['Escalated', 'escalated', 'layer8d-status-inactive']
    ]);

    MobileDeals.enums = {
        DEAL_STATUS: DEAL_STATUS.enum,
        DEAL_STATUS_VALUES: DEAL_STATUS.values,
        DEAL_STATUS_CLASSES: DEAL_STATUS.classes,
        DISPUTE_STATUS: DISPUTE_STATUS.enum,
        DISPUTE_STATUS_VALUES: DISPUTE_STATUS.values,
        DISPUTE_STATUS_CLASSES: DISPUTE_STATUS.classes
    };

    MobileDeals.render = {};
    MobileDeals.render.dealStatus = createStatusRenderer(DEAL_STATUS.enum, DEAL_STATUS.classes);
    MobileDeals.render.disputeStatus = createStatusRenderer(DISPUTE_STATUS.enum, DISPUTE_STATUS.classes);
})();
