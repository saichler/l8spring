(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'Marketplace',
        defaultModule: 'browse',
        defaultService: 'listings',
        sectionSelector: 'browse',
        initializerName: 'initializeMarketplace',
        requiredNamespaces: ['Browse', 'MyBids', 'Deals', 'Reviews']
    });
})();
