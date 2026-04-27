(function() {
    'use strict';
    window.MyBids = window.MyBids || {};
    var factory = window.Layer8EnumFactory;
    var createStatusRenderer = Layer8DRenderers.createStatusRenderer;
    var renderEnum = Layer8DRenderers.renderEnum;

    var BID_STATUS = factory.create([
        ['Unspecified', null, ''],
        ['Pending', 'pending', 'layer8d-status-pending'],
        ['Accepted', 'accepted', 'layer8d-status-active'],
        ['Rejected', 'rejected', 'layer8d-status-inactive'],
        ['Withdrawn', 'withdrawn', 'layer8d-status-inactive'],
        ['Expired', 'expired', 'layer8d-status-inactive']
    ]);

    var ITEM_CONDITION = Browse.enums.ITEM_CONDITION;
    var ITEM_CONDITION_VALUES = Browse.enums.ITEM_CONDITION_VALUES;
    var ITEM_CONDITION_CLASSES = Browse.enums.ITEM_CONDITION_CLASSES;

    MyBids.enums = {
        BID_STATUS: BID_STATUS.enum,
        BID_STATUS_VALUES: BID_STATUS.values,
        BID_STATUS_CLASSES: BID_STATUS.classes,
        ITEM_CONDITION: ITEM_CONDITION,
        ITEM_CONDITION_VALUES: ITEM_CONDITION_VALUES,
        ITEM_CONDITION_CLASSES: ITEM_CONDITION_CLASSES
    };

    MyBids.render = {};
    MyBids.render.bidStatus = createStatusRenderer(BID_STATUS.enum, BID_STATUS.classes);
    MyBids.render.itemCondition = function(value) {
        return renderEnum(value, ITEM_CONDITION);
    };
})();
