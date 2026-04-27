(function() {
    'use strict';
    const svc = Layer8ModuleConfigFactory.service;
    const mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'Marketplace',
        modules: {
            'browse': mod('Browse', '', [
                { key: 'listings', label: 'Listings', icon: '', endpoint: '/10/Listing', model: 'SpringListing',
                  supportedViews: ['table', 'kanban'],
                  viewConfig: { laneField: 'status', lanes: {
                      2: { label: 'Active', color: '#22c55e' },
                      4: { label: 'Fulfilled', color: '#0ea5e9' },
                      5: { label: 'Expired', color: '#f59e0b' },
                      6: { label: 'Cancelled', color: '#ef4444' }
                  }, cardTitle: 'title', cardSubtitle: 'maxPrice', cardFields: ['condition', 'shippingPreference'] }
                },
                { key: 'categories', label: 'Categories', icon: '', endpoint: '/10/Category', model: 'SpringCategory',
                  supportedViews: ['table', 'tree'],
                  viewConfig: { parentIdField: 'parentCategoryId', idField: 'categoryId', labelField: 'name' }
                }
            ]),
            'mylistings': mod('My Listings', '', [
                { key: 'mylistings', label: 'My Listings', icon: '', endpoint: '/10/Listing', model: 'SpringListing',
                  supportedViews: ['table', 'kanban'],
                  viewConfig: { laneField: 'status', lanes: {
                      1: { label: 'Draft', color: '#94a3b8' },
                      2: { label: 'Active', color: '#22c55e' },
                      4: { label: 'Fulfilled', color: '#0ea5e9' },
                      5: { label: 'Expired', color: '#f59e0b' },
                      6: { label: 'Cancelled', color: '#ef4444' }
                  }, cardTitle: 'title', cardSubtitle: 'maxPrice' }
                }
            ]),
            'mybids': mod('My Bids', '', [
                { key: 'mybids', label: 'My Bids', icon: '', endpoint: '/10/Bid', model: 'SpringBid',
                  supportedViews: ['table', 'timeline'],
                  viewConfig: { dateField: 'auditInfo.createdDate', titleField: 'listingId', descriptionField: 'message' }
                }
            ]),
            'deals': mod('Deals', '', [
                { key: 'deals', label: 'My Deals', icon: '', endpoint: '/10/Deal', model: 'SpringDeal',
                  supportedViews: ['table', 'kanban', 'timeline'],
                  viewConfig: { laneField: 'status', lanes: {
                      1: { label: 'Pending', color: '#94a3b8' },
                      2: { label: 'In Progress', color: '#0ea5e9' },
                      3: { label: 'Shipped', color: '#f59e0b' },
                      4: { label: 'Delivered', color: '#8b5cf6' },
                      5: { label: 'Completed', color: '#22c55e' },
                      6: { label: 'Cancelled', color: '#ef4444' },
                      7: { label: 'Disputed', color: '#dc2626' }
                  }, cardTitle: 'itemTitle', cardSubtitle: 'totalAmount' }
                }
            ]),
            'reviews': mod('Reviews', '', [
                svc('reviews', 'Reviews', '', '/10/Review', 'SpringReview')
            ])
        },
        submodules: ['Browse', 'MyBids', 'Deals', 'Reviews']
    });
})();
