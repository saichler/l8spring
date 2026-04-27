(function() {
    'use strict';

    window.sections = {
        dashboard: 'sections/dashboard.html',
        marketplace: 'sections/marketplace.html',
        system: 'sections/system.html'
    };

    window.sectionInitializers = {
        dashboard: function() {
            if (typeof initSpringDashboard === 'function') initSpringDashboard();
        },
        marketplace: function() {
            if (typeof initializeMarketplace === 'function') initializeMarketplace();
        },
        system: function() {
            if (typeof initializeL8Sys === 'function') initializeL8Sys();
        }
    };
})();
