(function() {
    'use strict';
    window.LAYER8M_NAV_CONFIG = window.LAYER8M_NAV_CONFIG || { modules: [] };

    LAYER8M_NAV_CONFIG.modules = [
        { key: 'marketplace', label: 'Marketplace', icon: 'marketplace', hasSubModules: true }
    ];

    LAYER8M_NAV_CONFIG.marketplace = {
        subModules: [
            { key: 'browse', label: 'Browse', icon: 'browse' },
            { key: 'mybids', label: 'My Bids', icon: 'mybids' },
            { key: 'deals', label: 'Deals', icon: 'deals' },
            { key: 'reviews', label: 'Reviews', icon: 'reviews' }
        ],
        services: {
            'browse': [
                { key: 'listings', label: 'Listings', icon: 'listings',
                  endpoint: '/10/Listing', model: 'SpringListing', idField: 'listingId',
                  supportedViews: ['table', 'kanban'] },
                { key: 'categories', label: 'Categories', icon: 'categories',
                  endpoint: '/10/Category', model: 'SpringCategory', idField: 'categoryId',
                  supportedViews: ['table', 'tree'] }
            ],
            'mybids': [
                { key: 'mybids', label: 'My Bids', icon: 'mybids',
                  endpoint: '/10/Bid', model: 'SpringBid', idField: 'bidId',
                  supportedViews: ['table', 'timeline'] }
            ],
            'deals': [
                { key: 'deals', label: 'My Deals', icon: 'deals',
                  endpoint: '/10/Deal', model: 'SpringDeal', idField: 'dealId',
                  supportedViews: ['table', 'kanban', 'timeline'] }
            ],
            'reviews': [
                { key: 'reviews', label: 'Reviews', icon: 'reviews',
                  endpoint: '/10/Review', model: 'SpringReview', idField: 'reviewId' }
            ]
        }
    };

    LAYER8M_NAV_CONFIG.icons = {
        'marketplace': '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"/></svg>',
        'browse': '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/></svg>',
        'mybids': '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="12" y1="1" x2="12" y2="23"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/></svg>',
        'deals': '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg>',
        'reviews': '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/></svg>',
        'listings': '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"/><line x1="3" y1="9" x2="21" y2="9"/><line x1="9" y1="21" x2="9" y2="9"/></svg>',
        'categories': '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"/></svg>'
    };

    LAYER8M_NAV_CONFIG.getIcon = function(key) {
        return (this.icons && this.icons[key]) || '';
    };
})();
