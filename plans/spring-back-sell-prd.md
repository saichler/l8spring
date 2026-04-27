# Spring Back Sell - Product Requirements Document

**Project:** l8spring
**Platform Name:** Spring Back Sell
**Version:** 1.0
**Status:** Draft - Pending Approval
**Date:** 2026-04-25

---

## 1. Vision & Purpose

Spring Back Sell is a **reverse commerce platform** where **buyers drive the market**. Instead of sellers listing products and buyers searching for them, buyers post listings describing what they want to buy and the maximum price they are willing to pay. Sellers browse these buyer requests and submit bids offering to fulfill them at or below the buyer's price. When a buyer accepts a bid, a deal is created and the transaction proceeds through payment, shipping, and completion.

### Target Users
- **Buyers** looking for specific products at their desired price point
- **Sellers** with inventory who want demand-driven sales instead of competing on traditional marketplaces
- **Casual users** who want to sell items they no longer need to people actively looking for them

### Core Value Proposition
- Eliminates the seller's guessing game about pricing — buyers state what they will pay
- Reduces unsold inventory — sellers respond only to confirmed demand
- Creates competitive pricing among sellers, benefiting buyers
- Enables price discovery from the buyer's perspective

---

## 2. System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                     Spring Back Sell                             │
│                                                                 │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────┐   │
│  │ Category │  │ Listing  │  │   Bid    │  │    Deal      │   │
│  │ Service  │  │ Service  │  │ Service  │  │   Service    │   │
│  │ (SA=10)  │  │ (SA=10)  │  │ (SA=10)  │  │   (SA=10)    │   │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └──────┬───────┘   │
│       │              │              │               │           │
│  ┌────┴──────────────┴──────────────┴───────────────┴───────┐  │
│  │                    L8 ORM / VNet                          │  │
│  └──────────────────────┬───────────────────────────────────┘  │
│                         │                                       │
│  ┌──────────────────────┴───────────────────────────────────┐  │
│  │                   PostgreSQL                              │  │
│  └───────────────────────────────────────────────────────────┘  │
│                                                                 │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │              UI (Desktop + Mobile)                        │  │
│  │  ┌──────────┐ ┌──────────┐ ┌─────────┐ ┌─────────────┐  │  │
│  │  │ Browse   │ │ My       │ │ My Bids │ │ My Deals    │  │  │
│  │  │ Listings │ │ Listings │ │         │ │ & Reviews   │  │  │
│  │  └──────────┘ └──────────┘ └─────────┘ └─────────────┘  │  │
│  └───────────────────────────────────────────────────────────┘  │
│                                                                 │
│  ┌───────────────┐  ┌───────────────┐                          │
│  │  Review       │  │    SYS        │                          │
│  │  Service      │  │   Module      │                          │
│  │  (SA=10)      │  │  (built-in)   │                          │
│  └───────────────┘  └───────────────┘                          │
└─────────────────────────────────────────────────────────────────┘
```

### Technology Stack
- **Backend:** Go services on the Layer 8 framework (l8types, l8orm, l8bus, l8secure)
- **Frontend:** l8ui shared component library (desktop + mobile)
- **Database:** PostgreSQL via l8orm
- **Networking:** l8bus virtual network (VNet)
- **Security:** ISecurityProvider implementation for authentication and authorization
- **Deployment:** Docker images on Kubernetes

---

## 3. Data Model

### 3.1 Prime Objects

The following entities pass the Prime Object test (independent existence, own lifecycle, directly queried, no parent dependency):

#### SpringCategory

Product categories for organizing listings. Supports hierarchical nesting via self-referencing parentCategoryId.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| category_id | string | PK | Auto-generated |
| name | string | Yes | Category display name |
| description | string | No | Category description |
| parent_category_id | string | No | Parent category for hierarchy |
| icon_url | string | No | Category icon URL |
| is_active | bool | Yes | Whether category is visible |
| sort_order | int32 | No | Display ordering |
| listing_count | int64 | No | Number of active listings (computed) |
| audit_info | AuditInfo | Auto | Created/modified tracking |
| custom_fields | map | No | Extensibility |

#### SpringListing

A buyer's request to purchase a specific product at a stated price. This is the core entity of the platform.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| listing_id | string | PK | Auto-generated |
| buyer_id | string | Yes | Security user ID of the buyer |
| buyer_display_name | string | Yes | Buyer's display name |
| title | string | Yes | What the buyer wants (e.g., "iPhone 15 Pro 256GB") |
| description | string | Yes | Detailed requirements, acceptable conditions, etc. |
| category_id | string | Yes | Reference to SpringCategory |
| desired_condition | int32 | Yes | Enum: condition the buyer will accept |
| max_price | Money | Yes | Maximum price the buyer is willing to pay |
| quantity | int32 | Yes | Number of items wanted (default 1) |
| status | int32 | Yes | Enum: listing lifecycle status |
| expires_at | int64 | No | Expiration timestamp (0 = no expiration) |
| location | string | No | Buyer's general location for shipping |
| shipping_preference | int32 | Yes | Enum: how the buyer wants to receive the item |
| tags | repeated string | No | Searchable tags |
| view_count | int64 | No | Number of times viewed (computed) |
| bid_count | int32 | No | Number of bids received (computed) |
| images | repeated SpringListingImage | No | Reference images of desired product |
| audit_info | AuditInfo | Auto | Created/modified tracking |
| custom_fields | map | No | Extensibility |

**SpringListingImage** (embedded child — meaningless without parent listing):

| Field | Type | Description |
|-------|------|-------------|
| image_id | string | Image identifier |
| storage_path | string | File path from FileStore (NOT binary, NOT URL) |
| file_name | string | Original filename for display |
| file_size | int64 | Size in bytes for display |
| caption | string | Image description |
| sort_order | int32 | Display ordering |

#### SpringBid

A seller's offer to fulfill a buyer's listing. Bids have their own lifecycle (submitted, accepted, rejected, withdrawn) and sellers query "all my bids" across listings, so this passes the Prime Object test.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| bid_id | string | PK | Auto-generated |
| listing_id | string | Yes | Reference to SpringListing being fulfilled |
| seller_id | string | Yes | Security user ID of the seller |
| seller_display_name | string | Yes | Seller's display name |
| price | Money | Yes | Seller's offered price (typically <= listing max_price) |
| item_condition | int32 | Yes | Enum: actual condition of seller's item |
| description | string | Yes | Seller describes their specific item |
| estimated_ship_days | int32 | No | Estimated business days to ship |
| shipping_cost | Money | No | Shipping cost (0 = free shipping) |
| status | int32 | Yes | Enum: bid lifecycle status |
| images | repeated SpringBidImage | No | Photos of the actual item |
| seller_rating | float | No | Seller's average rating (computed) |
| audit_info | AuditInfo | Auto | Created/modified tracking |
| custom_fields | map | No | Extensibility |

**SpringBidImage** (embedded child — meaningless without parent bid):

| Field | Type | Description |
|-------|------|-------------|
| image_id | string | Image identifier |
| storage_path | string | File path from FileStore (NOT binary, NOT URL) |
| file_name | string | Original filename for display |
| file_size | int64 | Size in bytes for display |
| caption | string | Image description |
| sort_order | int32 | Display ordering |

#### SpringDeal

Created when a buyer accepts a bid. Tracks the full transaction lifecycle from payment through delivery and completion.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| deal_id | string | PK | Auto-generated |
| listing_id | string | Yes | Reference to SpringListing |
| bid_id | string | Yes | Reference to accepted SpringBid |
| buyer_id | string | Yes | Buyer's security user ID |
| seller_id | string | Yes | Seller's security user ID |
| buyer_display_name | string | Yes | Buyer's display name |
| seller_display_name | string | Yes | Seller's display name |
| item_title | string | Yes | Copied from listing title for quick reference |
| agreed_price | Money | Yes | The accepted bid price |
| shipping_cost | Money | No | Shipping cost from the bid |
| total_amount | Money | Yes | agreed_price + shipping_cost |
| status | int32 | Yes | Enum: deal lifecycle status |
| payment_reference | string | No | External payment system reference |
| tracking_number | string | No | Shipping tracking number |
| shipping_carrier | string | No | Carrier name (UPS, FedEx, USPS, etc.) |
| shipped_at | int64 | No | Timestamp when shipped |
| delivered_at | int64 | No | Timestamp when delivered |
| completed_at | int64 | No | Timestamp when deal completed |
| messages | repeated SpringDealMessage | No | Buyer-seller communication |
| dispute | SpringDealDispute | No | Dispute details if raised |
| audit_info | AuditInfo | Auto | Created/modified tracking |
| custom_fields | map | No | Extensibility |

**SpringDealMessage** (embedded child — communication within a deal context):

| Field | Type | Description |
|-------|------|-------------|
| message_id | string | Message identifier |
| sender_id | string | Who sent it |
| sender_display_name | string | Display name |
| content | string | Message text |
| sent_at | int64 | Timestamp |

**SpringDealDispute** (embedded child — dispute is part of deal lifecycle):

| Field | Type | Description |
|-------|------|-------------|
| reason | string | Dispute reason code |
| description | string | Detailed explanation |
| status | int32 | Enum: dispute status |
| resolution | string | How it was resolved |
| opened_at | int64 | When dispute was opened |
| resolved_at | int64 | When dispute was resolved |
| resolved_by | string | Admin who resolved it |

#### SpringReview

Post-deal reviews. Both buyer and seller can review each other after a deal completes. Reviews are queried independently ("show all reviews for seller X") and have their own moderation lifecycle, passing the Prime Object test.

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| review_id | string | PK | Auto-generated |
| deal_id | string | Yes | Reference to SpringDeal |
| reviewer_id | string | Yes | Who wrote the review |
| reviewee_id | string | Yes | Who is being reviewed |
| reviewer_display_name | string | Yes | Reviewer's display name |
| reviewer_role | int32 | Yes | Enum: buyer or seller |
| rating | int32 | Yes | 1-5 star rating |
| title | string | No | Review headline |
| content | string | Yes | Review text |
| status | int32 | Yes | Enum: review moderation status |
| audit_info | AuditInfo | Auto | Created/modified tracking |
| custom_fields | map | No | Extensibility |

### 3.2 Prime Object Justification

| Entity | Independence | Own Lifecycle | Direct Query | No Parent Dependency | Result |
|--------|:---:|:---:|:---:|:---:|:---:|
| SpringCategory | Yes — exists as taxonomy | Yes — created/deactivated independently | Yes — browsed directly | Yes | Prime |
| SpringListing | Yes — buyer's request | Yes — draft → active → fulfilled/expired | Yes — browsed/searched | Yes | Prime |
| SpringBid | Yes — seller's offer | Yes — pending → accepted/rejected/withdrawn | Yes — "my bids" view | Yes | Prime |
| SpringDeal | Yes — completed transaction | Yes — payment → shipping → completion | Yes — "my deals" view | Yes | Prime |
| SpringReview | Yes — feedback record | Yes — pending → published/flagged/removed | Yes — "reviews for user X" | Yes | Prime |

### 3.3 Embedded Children Justification

| Type | Parent | Why NOT Prime |
|------|--------|---------------|
| SpringListingImage | SpringListing | Meaningless without listing, no independent lifecycle |
| SpringBidImage | SpringBid | Meaningless without bid, no independent lifecycle |
| SpringDealMessage | SpringDeal | Communication within deal context only |
| SpringDealDispute | SpringDeal | Dispute lifecycle tied to deal, always viewed within deal |

---

## 4. Enums

### SpringListingStatus
| Value | Code | Description |
|-------|------|-------------|
| SPRING_LISTING_STATUS_UNSPECIFIED | 0 | Invalid/unset |
| SPRING_LISTING_STATUS_DRAFT | 1 | Not yet published |
| SPRING_LISTING_STATUS_ACTIVE | 2 | Accepting bids |
| SPRING_LISTING_STATUS_FULFILLED | 3 | Deal made, listing complete |
| SPRING_LISTING_STATUS_EXPIRED | 4 | Past expiration date with no deal |
| SPRING_LISTING_STATUS_CANCELLED | 5 | Buyer cancelled |

### SpringBidStatus
| Value | Code | Description |
|-------|------|-------------|
| SPRING_BID_STATUS_UNSPECIFIED | 0 | Invalid/unset |
| SPRING_BID_STATUS_PENDING | 1 | Awaiting buyer review |
| SPRING_BID_STATUS_ACCEPTED | 2 | Buyer accepted this bid |
| SPRING_BID_STATUS_REJECTED | 3 | Buyer rejected this bid |
| SPRING_BID_STATUS_WITHDRAWN | 4 | Seller withdrew the bid |
| SPRING_BID_STATUS_EXPIRED | 5 | Listing expired before action |

### SpringDealStatus
| Value | Code | Description |
|-------|------|-------------|
| SPRING_DEAL_STATUS_UNSPECIFIED | 0 | Invalid/unset |
| SPRING_DEAL_STATUS_PENDING_PAYMENT | 1 | Awaiting buyer payment |
| SPRING_DEAL_STATUS_PAID | 2 | Payment received |
| SPRING_DEAL_STATUS_SHIPPED | 3 | Seller shipped the item |
| SPRING_DEAL_STATUS_DELIVERED | 4 | Item delivered to buyer |
| SPRING_DEAL_STATUS_COMPLETED | 5 | Both parties confirmed, deal closed |
| SPRING_DEAL_STATUS_DISPUTED | 6 | Dispute raised |
| SPRING_DEAL_STATUS_CANCELLED | 7 | Deal cancelled before shipment |
| SPRING_DEAL_STATUS_REFUNDED | 8 | Payment refunded |

### SpringDisputeStatus
| Value | Code | Description |
|-------|------|-------------|
| SPRING_DISPUTE_STATUS_UNSPECIFIED | 0 | Invalid/unset |
| SPRING_DISPUTE_STATUS_OPEN | 1 | Dispute filed |
| SPRING_DISPUTE_STATUS_UNDER_REVIEW | 2 | Admin reviewing |
| SPRING_DISPUTE_STATUS_RESOLVED_BUYER | 3 | Resolved in buyer's favor |
| SPRING_DISPUTE_STATUS_RESOLVED_SELLER | 4 | Resolved in seller's favor |
| SPRING_DISPUTE_STATUS_ESCALATED | 5 | Escalated to higher authority |

### SpringReviewStatus
| Value | Code | Description |
|-------|------|-------------|
| SPRING_REVIEW_STATUS_UNSPECIFIED | 0 | Invalid/unset |
| SPRING_REVIEW_STATUS_PENDING | 1 | Awaiting moderation |
| SPRING_REVIEW_STATUS_PUBLISHED | 2 | Visible to all users |
| SPRING_REVIEW_STATUS_FLAGGED | 3 | Flagged for review |
| SPRING_REVIEW_STATUS_REMOVED | 4 | Removed by admin |

### SpringItemCondition
| Value | Code | Description |
|-------|------|-------------|
| SPRING_ITEM_CONDITION_UNSPECIFIED | 0 | Invalid/unset |
| SPRING_ITEM_CONDITION_NEW | 1 | Brand new, sealed |
| SPRING_ITEM_CONDITION_LIKE_NEW | 2 | Opened but unused or minimal use |
| SPRING_ITEM_CONDITION_GOOD | 3 | Used, normal wear |
| SPRING_ITEM_CONDITION_FAIR | 4 | Visible wear but functional |
| SPRING_ITEM_CONDITION_ANY | 5 | Buyer accepts any condition |

### SpringShippingPreference
| Value | Code | Description |
|-------|------|-------------|
| SPRING_SHIPPING_PREFERENCE_UNSPECIFIED | 0 | Invalid/unset |
| SPRING_SHIPPING_PREFERENCE_SHIPPING | 1 | Ship to buyer |
| SPRING_SHIPPING_PREFERENCE_LOCAL_PICKUP | 2 | Local pickup only |
| SPRING_SHIPPING_PREFERENCE_EITHER | 3 | Either option acceptable |

### SpringReviewerRole
| Value | Code | Description |
|-------|------|-------------|
| SPRING_REVIEWER_ROLE_UNSPECIFIED | 0 | Invalid/unset |
| SPRING_REVIEWER_ROLE_BUYER | 1 | Review written by buyer |
| SPRING_REVIEWER_ROLE_SELLER | 2 | Review written by seller |

---

## 5. Workflows

### 5.1 Listing Lifecycle

```
  ┌───────┐
  │ DRAFT │ ← Buyer creates listing
  └───┬───┘
      │ publish
      v
  ┌────────┐
  │ ACTIVE │ ← Accepting bids from sellers
  └───┬────┘
      │
      ├─── accept bid ──> FULFILLED (deal created)
      │
      ├─── expires_at reached ──> EXPIRED
      │
      └─── buyer cancels ──> CANCELLED
```

### 5.2 Bid Lifecycle

```
  ┌─────────┐
  │ PENDING │ ← Seller submits bid
  └────┬────┘
       │
       ├─── buyer accepts ──> ACCEPTED (deal created)
       │
       ├─── buyer rejects ──> REJECTED
       │
       ├─── seller withdraws ──> WITHDRAWN
       │
       └─── listing expires ──> EXPIRED
```

### 5.3 Deal Lifecycle

```
  ┌─────────────────┐
  │ PENDING_PAYMENT  │ ← Created when bid accepted
  └────────┬────────┘
           │ payment confirmed
           v
       ┌──────┐
       │ PAID │
       └──┬───┘
          │ seller ships
          v
     ┌─────────┐
     │ SHIPPED │
     └────┬────┘
          │ delivery confirmed
          v
    ┌───────────┐
    │ DELIVERED │
    └─────┬─────┘
          │
          ├─── both parties confirm ──> COMPLETED
          │
          └─── dispute raised ──> DISPUTED
                                    │
                                    ├── resolved buyer ──> REFUNDED
                                    └── resolved seller ──> COMPLETED

  At any point before SHIPPED:
       ──> CANCELLED (mutual or admin)
```

### 5.4 Review Flow

```
  Deal reaches COMPLETED status
       │
       ├── Buyer can review seller
       │      └── PENDING → PUBLISHED (auto or moderated)
       │
       └── Seller can review buyer
              └── PENDING → PUBLISHED (auto or moderated)

  Admin can FLAG or REMOVE any review
```

---

## 6. UI Requirements

### 6.1 Desktop Modules

The UI is organized into a single module ("Marketplace") with sub-modules:

| Sub-Module | Services | Description |
|------------|----------|-------------|
| Browse | Listings, Categories | Browse active listings, search, filter by category. Also powers "My Listings" tab via `baseWhereClause` filter (`buyerId=<currentUser>`) — same model, same enums/columns/forms. |
| My Bids | My Bids | Seller's submitted bids across all listings |
| Deals | My Deals | Active and completed transactions |
| Reviews | Reviews | Reviews given and received |

### 6.2 View Types

| Service | Supported Views | Notes |
|---------|----------------|-------|
| Listings (browse) | table, kanban (by status) | Kanban lanes: Active, Fulfilled, Expired |
| My Listings | table, kanban | Same as Listings but filtered by `buyerId=<currentUser>` |
| My Bids | table, timeline | Timeline by bid date |
| My Deals | table, kanban, timeline | Kanban by deal status |
| Categories | table, tree | Tree view for category hierarchy |
| Reviews | table | Standard table view |

### 6.3 Dashboard

The dashboard section displays KPI widgets:

| Widget | Metric | Description |
|--------|--------|-------------|
| Active Listings | Count | Total active listings on the platform |
| Pending Bids | Count | Bids awaiting buyer response |
| Deals In Progress | Count | Deals not yet completed |
| Completed Deals | Count | Total completed transactions |
| Average Deal Value | Money | Average total_amount of completed deals |
| Average Rating | Number | Platform-wide average review rating |

### 6.4 Mobile Parity

Per `mobile-rules.md` Rule 2, every desktop feature must have functional parity on mobile. The following table maps each desktop sub-module to its mobile equivalent:

| Desktop Sub-Module | Desktop Files | Mobile Sub-Module | Mobile Files |
|-------------------|--------------|-------------------|-------------|
| browse/ (enums, columns, forms) | 3 files | m/js/browse/ (enums, columns, forms) | 3 files |
| *(mylistings reuses browse/ — no separate files)* | — | *(mylistings reuses m/js/browse/ — no separate files)* | — |
| mybids/ (enums, columns, forms) | 3 files | m/js/mybids/ (enums, columns, forms) | 3 files |
| deals/ (enums, columns, forms) | 3 files | m/js/deals/ (enums, columns, forms) | 3 files |
| reviews/ (enums, columns, forms) | 3 files | m/js/reviews/ (enums, columns, forms) | 3 files |
| marketplace-config.js | 1 file | mobile-config-spring.js | 1 file |
| marketplace-init.js | 1 file | spring-index.js (Layer8MModuleRegistry) | 1 file |
| reference-registry-spring.js | 1 file | spring-reference-registry.js | 1 file |
| app.html (script tags) | wired | m/app.html (script tags) | wired |
| sections.js + app.js | 2 files | app-core.js | 1 file |

#### Mobile-Specific Requirements

- **Mobile columns** must include `primary: true` and `secondary: true` markers for card display (per `checklist.md`)
- **Mobile nav config** must define `hasSubModules: true` + service configs with `idField`, `endpoint`, `model` (per `adding-module-mobile.md`)
- **m/app.html sidebar** must include `<a>` with `data-section="dashboard" data-module="marketplace"` (per `checklist.md`)
- **Layer8MNav data lookups** must register `window.MobileSpring` in `_getServiceColumns`, `_getServiceFormDef`, and `_getServiceTransformData` arrays (per `adding-module-mobile.md`)
- **Mobile reference registry** must register with `Layer8MReferenceRegistry.register()` (per `layer8m-nav.md`)

#### Feature Parity Verification Checklist

For each sub-module, verify on BOTH desktop and mobile:
- [ ] Table/card data loads with correct columns
- [ ] Row/card click opens detail popup with all fields populated
- [ ] Add form works (creates record)
- [ ] Edit form works (updates record)
- [ ] Delete works (with confirmation)
- [ ] Status columns show labels, not raw numbers
- [ ] Reference fields show display names, not raw IDs
- [ ] Money fields show formatted currency
- [ ] Date fields show formatted dates
- [ ] Inline tables (bids on listing, messages on deal) render correctly
- [ ] View types (kanban, tree, timeline) render on both platforms

### 6.5 Listing Detail Popup

When viewing a listing, the detail popup shows:
- **Overview tab:** Title, description, category, condition, max price, shipping preference, status
- **Bids tab:** Inline table of all bids on this listing (with accept/reject actions for the buyer)
- **Images tab:** Inline table of images using FileStore pattern — each row has `storagePath` (type `file`), `caption` (text), `sortOrder` (number). Upload via `Layer8FileUpload`, download via stored `storagePath`.

#### Image Upload Pattern (per `file-upload-pattern.md`)

Listing and bid images use the framework's FileStore service (global, service area 0). Entities store `storage_path` strings — no binary data in proto fields, no custom upload endpoints.

**Listing form images section:**
```javascript
f.section('Images', [
    ...f.inlineTable('images', 'Images', [
        { key: 'imageId', label: 'ID', hidden: true },
        { key: 'storagePath', label: 'Image', type: 'file' },
        { key: 'caption', label: 'Caption', type: 'text' },
        { key: 'sortOrder', label: 'Order', type: 'number' }
    ]),
])
```

**Bid form images section:** Same pattern as listing images.

### 6.6 Deal Detail Popup

When viewing a deal, the detail popup shows:
- **Overview tab:** Item title, buyer/seller info, agreed price, shipping cost, total, status
- **Shipping tab:** Tracking number, carrier, shipped/delivered timestamps
- **Messages tab:** Inline table of deal messages with ability to add new messages
- **Dispute tab:** Dispute details if applicable (read-only unless admin)

---

## 7. Access Control

### 7.1 Roles

| Role | Permissions |
|------|-------------|
| Buyer | Create/edit/cancel own listings; view bids on own listings; accept/reject bids; manage own deals; write reviews |
| Seller | Browse listings; submit/withdraw bids; manage shipping on own deals; write reviews |
| User | Both Buyer and Seller permissions (default role — all users can buy and sell) |
| Admin | All user permissions + manage categories, moderate reviews, resolve disputes, view all data |
| System | Built-in SYS module access (health, security, modules, logs) |

### 7.2 Security Implementation

Security is implemented via `ifs.ISecurityProvider` as required by the Layer 8 framework. Users, roles, and credentials are provisioned through the security config JSON or the Security API (SYS module).

### 7.3 Security Config Design

The security config JSON (`go/secure/plugin/spring/spring.json`) defines roles with allow/deny rules following the l8secure framework pattern (reference: `l8secure/go/secure/plugin/phy/phy.json`).

#### Allow Rules (action-level authorization)

| Role | Entity | Actions | Attributes |
|------|--------|---------|------------|
| user | SpringCategory | GET | `*: *` |
| user | SpringListing | POST, PUT, GET, DELETE | `*: *` |
| user | SpringBid | POST, PUT, GET, DELETE | `*: *` |
| user | SpringDeal | PUT, GET | `*: *` |
| user | SpringReview | POST, GET | `*: *` |
| admin | *(all types)* | all (`-999`) | `*: *` |

#### Deny Rules with Row-Level Scoping

Row-level scoping uses L8Query deny rules with `${userId}` placeholder. The framework's `ScopeView()` runs after every GET, removing rows that match the deny query.

| Rule ID | Role | Entity | Deny Query | Effect |
|---------|------|--------|------------|--------|
| `user-deny-other-drafts` | user | SpringListing | `select * from SpringListing where buyerId!=${userId} and status=1` | Users cannot see other buyers' DRAFT listings. Active/fulfilled/expired listings remain visible to all (public marketplace). |
| `user-deny-other-deals` | user | SpringDeal | `select * from SpringDeal where buyerId!=${userId} and sellerId!=${userId}` | Users can only see deals where they are the buyer OR the seller. |

#### Field-Level Denials

| Rule ID | Role | Field | Effect |
|---------|------|-------|--------|
| `user-deny-deal-payment` | user | `springdeal.paymentreference` | Payment reference is admin-only |

#### Public vs. User-Scoped Data

| Entity | Public Data | User-Scoped Data |
|--------|-------------|------------------|
| SpringCategory | All categories (read-only for non-admin) | None — categories are fully public |
| SpringListing | Active, Fulfilled, Expired, Cancelled listings | Draft listings (only visible to the creating buyer) |
| SpringBid | All bids are public — transparency creates price competition among sellers, benefiting buyers | None |
| SpringDeal | None — deals are private | Scoped to buyer or seller of the deal |
| SpringReview | All reviews are public — enables informed buying decisions and marketplace trust | None |

#### Pre-Defined Users

```json
"users": {
  "admin": {
    "userName": "admin",
    "password": "admin",
    "roles": { "admin": true }
  }
}
```

Additional users are created via the Security API (SYS module) at runtime. All marketplace users receive the `user` role by default, which grants both buyer and seller capabilities.

#### Security Config File Location

```
go/secure/plugin/spring/spring.json
```

---

## 8. Service Area & Service Names

**Module:** spring
**Service Area:** 10 (all services share this area)
**PREFIX:** `/spring`

| Protobuf Type | ServiceName | Area | Endpoint | Primary Key |
|---------------|-------------|------|----------|-------------|
| SpringCategory | Category | 10 | /spring/10/Category | categoryId |
| SpringListing | Listing | 10 | /spring/10/Listing | listingId |
| SpringBid | Bid | 10 | /spring/10/Bid | bidId |
| SpringDeal | Deal | 10 | /spring/10/Deal | dealId |
| SpringReview | Review | 10 | /spring/10/Review | reviewId |

All ServiceName values are within the 10-character limit.

### L8Query Examples
```
select * from SpringListing where status=2 sort-by auditInfo.createdDate descending
select * from SpringBid where listingId=LST-001
select * from SpringDeal where buyerId=USR-001 sort-by auditInfo.createdDate descending
select * from SpringReview where revieweeId=USR-002
select * from SpringCategory where isActive=true sort-by sortOrder
```

---

## 9. Implementation Reference

### 9.1 Project Directory Structure

```
l8spring/
├── proto/
│   ├── make-bindings.sh
│   ├── spring-categories.proto
│   ├── spring-listings.proto
│   ├── spring-bids.proto
│   ├── spring-deals.proto
│   └── spring-reviews.proto
├── go/
│   ├── go.mod
│   ├── go.sum
│   ├── vendor/
│   ├── build-all-images.sh
│   ├── run-local.sh
│   ├── secure/
│   │   └── plugin/
│   │       └── spring/
│   │           └── spring.json            # Security config (roles, deny rules, users)
│   ├── spring/
│   │   ├── common/
│   │   │   └── defaults.go               # PREFIX="/spring", shared constants
│   │   ├── categories/
│   │   │   ├── SpringCategoryService.go
│   │   │   └── SpringCategoryServiceCallback.go
│   │   ├── listings/
│   │   │   ├── SpringListingService.go
│   │   │   └── SpringListingServiceCallback.go
│   │   ├── bids/
│   │   │   ├── SpringBidService.go
│   │   │   └── SpringBidServiceCallback.go
│   │   ├── deals/
│   │   │   ├── SpringDealService.go
│   │   │   └── SpringDealServiceCallback.go
│   │   ├── reviews/
│   │   │   ├── SpringReviewService.go
│   │   │   └── SpringReviewServiceCallback.go
│   │   ├── ui/
│   │   │   ├── main.go                    # UI server + type registration
│   │   │   ├── web/
│   │   │   │   ├── app.html               # Desktop app shell
│   │   │   │   ├── login.html
│   │   │   │   ├── login.json
│   │   │   │   ├── register/
│   │   │   │   │   └── index.html
│   │   │   │   ├── l8ui/                  # Shared UI library (submodule)
│   │   │   │   ├── js/
│   │   │   │   │   ├── sections.js
│   │   │   │   │   ├── app.js
│   │   │   │   │   └── reference-registry-spring.js
│   │   │   │   ├── sections/
│   │   │   │   │   └── marketplace.html
│   │   │   │   ├── marketplace/
│   │   │   │   │   ├── marketplace-config.js
│   │   │   │   │   ├── browse/              # Shared by browse + mylistings (same model)
│   │   │   │   │   │   ├── browse-enums.js
│   │   │   │   │   │   ├── browse-columns.js
│   │   │   │   │   │   └── browse-forms.js
│   │   │   │   │   ├── mybids/
│   │   │   │   │   │   ├── mybids-enums.js
│   │   │   │   │   │   ├── mybids-columns.js
│   │   │   │   │   │   └── mybids-forms.js
│   │   │   │   │   ├── deals/
│   │   │   │   │   │   ├── deals-enums.js
│   │   │   │   │   │   ├── deals-columns.js
│   │   │   │   │   │   └── deals-forms.js
│   │   │   │   │   ├── reviews/
│   │   │   │   │   │   ├── reviews-enums.js
│   │   │   │   │   │   ├── reviews-columns.js
│   │   │   │   │   │   └── reviews-forms.js
│   │   │   │   │   └── marketplace-init.js
│   │   │   │   └── m/                     # Mobile
│   │   │   │       ├── app.html
│   │   │   │       └── js/
│   │   │   │           ├── app-core.js
│   │   │   │           ├── mobile-config-spring.js
│   │   │   │           ├── browse/              # Shared by browse + mylistings (same model)
│   │   │   │           │   ├── browse-enums.js
│   │   │   │           │   ├── browse-columns.js
│   │   │   │           │   └── browse-forms.js
│   │   │   │           ├── mybids/
│   │   │   │           │   ├── mybids-enums.js
│   │   │   │           │   ├── mybids-columns.js
│   │   │   │           │   └── mybids-forms.js
│   │   │   │           ├── deals/
│   │   │   │           │   ├── deals-enums.js
│   │   │   │           │   ├── deals-columns.js
│   │   │   │           │   └── deals-forms.js
│   │   │   │           ├── reviews/
│   │   │   │           │   ├── reviews-enums.js
│   │   │   │           │   ├── reviews-columns.js
│   │   │   │           │   └── reviews-forms.js
│   │   │   │           ├── spring-index.js
│   │   │   │           └── spring-reference-registry.js
│   │   │   ├── build.sh
│   │   │   └── Dockerfile
│   │   ├── main/
│   │   │   ├── main.go                    # Backend server entry point
│   │   │   ├── build.sh
│   │   │   └── Dockerfile
│   │   └── vnet/
│   │       ├── main.go                    # Virtual network entry point
│   │       ├── build.sh
│   │       └── Dockerfile
│   ├── types/
│   │   └── spring/                        # Generated .pb.go files
│   ├── tests/
│   │   └── mocks/
│   │       ├── cmd/
│   │       │   └── main.go               # Mock data CLI
│   │       ├── data.go                    # Curated name/data arrays
│   │       ├── store.go                   # ID slices
│   │       ├── main_phases.go             # Phase orchestration
│   │       ├── gen_spring_categories.go
│   │       ├── gen_spring_listings.go
│   │       ├── gen_spring_bids.go
│   │       ├── gen_spring_deals.go
│   │       └── gen_spring_reviews.go
│   └── k8s/
│       ├── deploy.sh
│       ├── undeploy.sh
│       ├── spring.yaml
│       ├── spring-web.yaml
│       └── spring-vnet.yaml
└── plans/
    └── spring-back-sell-prd.md            # This document
```

### 9.2 Protobuf Pattern

```protobuf
syntax = "proto3";
package spring;
option go_package = "./types/spring";

import "l8common.proto";
import "api.proto";

// @PrimeObject
message SpringListing {
    string listing_id = 1;
    string buyer_id = 2;
    string buyer_display_name = 3;
    string title = 4;
    string description = 5;
    string category_id = 6;
    SpringItemCondition desired_condition = 7;
    l8common.Money max_price = 8;
    int32 quantity = 9;
    SpringListingStatus status = 10;
    int64 expires_at = 11;
    string location = 12;
    SpringShippingPreference shipping_preference = 13;
    repeated string tags = 14;
    int64 view_count = 15;
    int32 bid_count = 16;
    repeated SpringListingImage images = 17;
    l8common.AuditInfo audit_info = 20;
    map<string, string> custom_fields = 21;
}

message SpringListingImage {
    string image_id = 1;
    string storage_path = 2;
    string file_name = 3;
    int64 file_size = 4;
    string caption = 5;
    int32 sort_order = 6;
}

// SpringBidImage follows the same pattern (storage_path, file_name, file_size, caption, sort_order)

message SpringListingList {
    repeated SpringListing list = 1;
    l8api.L8MetaData metadata = 2;
}
```

### 9.3 Service Pattern

```go
package listings

import (
    common "github.com/saichler/l8spring/go/spring/common"
    "github.com/saichler/l8spring/go/types/spring"
    "github.com/saichler/l8types/go/ifs"
)

const (
    ServiceName = "Listing"
    ServiceArea = byte(10)
)

func Activate(creds, dbname string, vnic ifs.IVNic) {
    common.ActivateService(common.ServiceConfig{
        ServiceName: ServiceName,
        ServiceArea: ServiceArea,
        PrimaryKey:  "ListingId",
        Callback:    newSpringListingServiceCallback(vnic),
    }, &spring.SpringListing{}, &spring.SpringListingList{}, creds, dbname, vnic)
}
```

### 9.4 ServiceCallback Pattern

```go
package listings

import (
    common "github.com/saichler/l8spring/go/spring/common"
    "github.com/saichler/l8spring/go/types/spring"
    "github.com/saichler/l8types/go/ifs"
)

type SpringListingServiceCallback struct {
    vnic ifs.IVNic
}

func newSpringListingServiceCallback(vnic ifs.IVNic) *SpringListingServiceCallback {
    return &SpringListingServiceCallback{vnic: vnic}
}

func (cb *SpringListingServiceCallback) Before(action ifs.Action, item interface{}) error {
    entity := item.(*spring.SpringListing)

    if action == ifs.POST {
        common.GenerateID(&entity.ListingId)
    }

    return cb.validate(entity, action)
}

func (cb *SpringListingServiceCallback) After(action ifs.Action, item interface{}) {}

func (cb *SpringListingServiceCallback) validate(entity *spring.SpringListing, action ifs.Action) error {
    common.ValidateRequired(entity.Title, "Title")
    common.ValidateRequired(entity.BuyerId, "BuyerId")
    common.ValidateRequired(entity.Description, "Description")
    common.ValidateRequired(entity.CategoryId, "CategoryId")
    common.ValidateEnum(int32(entity.DesiredCondition), spring.SpringItemCondition_name, "DesiredCondition")
    common.ValidateMoney(entity.MaxPrice, "MaxPrice")
    common.ValidateEnum(int32(entity.Status), spring.SpringListingStatus_name, "Status")
    return nil
}
```

### 9.5 UI Module Config Pattern

```javascript
// marketplace/marketplace-config.js
(function() {
    'use strict';
    const svc = Layer8ModuleConfigFactory.service;
    const mod = Layer8ModuleConfigFactory.module;

    // Note: 'mylistings' reuses the Browse namespace (same SpringListing model).
    // The only difference is a baseWhereClause filter (buyerId=<currentUser>).
    // No separate MyListings data files (enums/columns/forms) are needed.
    Layer8ModuleConfigFactory.create({
        namespace: 'Marketplace',
        modules: {
            'browse': mod('Browse', '🔍', [
                svc('listings', 'Listings', '📋', '/10/Listing', 'SpringListing'),
                svc('categories', 'Categories', '📂', '/10/Category', 'SpringCategory')
            ]),
            'mylistings': mod('My Listings', '📝', [
                svc('mylistings', 'My Listings', '📝', '/10/Listing', 'SpringListing')
            ]),
            'mybids': mod('My Bids', '🏷️', [
                svc('mybids', 'My Bids', '🏷️', '/10/Bid', 'SpringBid')
            ]),
            'deals': mod('Deals', '🤝', [
                svc('deals', 'My Deals', '🤝', '/10/Deal', 'SpringDeal')
            ]),
            'reviews': mod('Reviews', '⭐', [
                svc('reviews', 'Reviews', '⭐', '/10/Review', 'SpringReview')
            ])
        },
        submodules: ['Browse', 'MyBids', 'Deals', 'Reviews']
    });
})();
```

### 9.6 Module Init Pattern

```javascript
// marketplace/marketplace-init.js
(function() {
    'use strict';
    Layer8DModuleFactory.create({
        namespace: 'Marketplace',
        defaultModule: 'browse',
        defaultService: 'listings',
        sectionSelector: 'browse',           // MUST match defaultModule
        initializerName: 'initializeMarketplace',
        requiredNamespaces: ['Browse', 'MyBids', 'Deals', 'Reviews']
    });
})();
```

Note: `sectionSelector` matches `defaultModule` (`'browse'`) per `module-init-section-selector.md`. The section HTML must have `<div class="l8-module-content active" data-module="browse">`.

### 9.7 login.json

```json
{
    "login": {
        "appTitle": "Spring Back Sell",
        "authEndpoint": "/auth",
        "redirectUrl": "/app.html",
        "sessionTimeout": 30,
        "tfaEnabled": true
    },
    "app": {
        "dateFormat": "mm/dd/yyyy",
        "apiPrefix": "/spring",
        "healthPath": "/0/Health"
    }
}
```

### 9.8 common/defaults.go

```go
package common

import (
    l8common "github.com/saichler/l8common/go/common"
    "github.com/saichler/l8types/go/ifs"
)

const PREFIX = "/spring"

func CreateResources(alias string, logVnet bool) ifs.IResources {
    return l8common.CreateResources(alias, logVnet)
}

var WaitForSignal = l8common.WaitForSignal
var OpenDBConection = l8common.OpenDBConection
var ActivateService = l8common.ActivateService
var GenerateID = l8common.GenerateID
var ValidateRequired = l8common.ValidateRequired
var ValidateEnum = l8common.ValidateEnum
var ValidateMoney = l8common.ValidateMoney
```

### 9.9 Mobile Nav Config Pattern

```javascript
// m/js/mobile-config-spring.js
(function() {
    'use strict';

    LAYER8M_NAV_CONFIG.modules.push(
        { key: 'marketplace', label: 'Marketplace', icon: 'marketplace', hasSubModules: true }
    );

    LAYER8M_NAV_CONFIG.marketplace = {
        subModules: [
            { key: 'browse', label: 'Browse', icon: 'search' },
            { key: 'mylistings', label: 'My Listings', icon: 'listing' },
            { key: 'mybids', label: 'My Bids', icon: 'bid' },
            { key: 'deals', label: 'Deals', icon: 'deal' },
            { key: 'reviews', label: 'Reviews', icon: 'review' }
        ],
        services: {
            'browse': [
                { key: 'listings', label: 'Listings', icon: 'listing',
                  endpoint: '/10/Listing', model: 'SpringListing', idField: 'listingId',
                  supportedViews: ['table', 'kanban'] },
                { key: 'categories', label: 'Categories', icon: 'category',
                  endpoint: '/10/Category', model: 'SpringCategory', idField: 'categoryId',
                  supportedViews: ['table', 'tree'] }
            ],
            // mylistings reuses Browse columns/forms (same SpringListing model).
            // The only difference is baseWhereClause filtering to the current user's listings.
            'mylistings': [
                { key: 'mylistings', label: 'My Listings', icon: 'listing',
                  endpoint: '/10/Listing', model: 'SpringListing', idField: 'listingId',
                  supportedViews: ['table', 'kanban'] }
            ],
            'mybids': [
                { key: 'mybids', label: 'My Bids', icon: 'bid',
                  endpoint: '/10/Bid', model: 'SpringBid', idField: 'bidId',
                  supportedViews: ['table', 'timeline'] }
            ],
            'deals': [
                { key: 'deals', label: 'My Deals', icon: 'deal',
                  endpoint: '/10/Deal', model: 'SpringDeal', idField: 'dealId',
                  supportedViews: ['table', 'kanban', 'timeline'] }
            ],
            'reviews': [
                { key: 'reviews', label: 'Reviews', icon: 'review',
                  endpoint: '/10/Review', model: 'SpringReview', idField: 'reviewId' }
            ]
        }
    };
})();
```

Note: `idField` values use the **JSON tag name** (camelCase), not the Go struct field name, per `js-protobuf-field-names.md`.

### 9.10 Mobile Column Pattern (with primary/secondary)

```javascript
// m/js/browse/browse-columns.js
(function() {
    'use strict';
    var enums = MobileBrowse.enums;
    var render = MobileBrowse.render;
    var col = window.Layer8ColumnFactory;

    MobileBrowse.columns = {
        SpringListing: [
            ...col.id('listingId'),
            ...col.col('title', 'Title'),              // primary: true added below
            ...col.money('maxPrice', 'Max Price'),     // secondary: true added below
            ...col.status('status', 'Status', enums.LISTING_STATUS_VALUES, render.listingStatus),
            ...col.enum('desiredCondition', 'Condition', null, render.itemCondition),
            ...col.col('location', 'Location'),
            ...col.col('buyerDisplayName', 'Buyer'),
            ...col.date('auditInfo.createdDate', 'Posted')
        ],
        SpringCategory: [
            ...col.id('categoryId'),
            ...col.col('name', 'Name'),
            ...col.col('description', 'Description'),
            ...col.boolean('isActive', 'Active')
        ]
    };

    // Card display markers
    MobileBrowse.columns.SpringListing[1].primary = true;    // title
    MobileBrowse.columns.SpringListing[2].secondary = true;  // maxPrice
    MobileBrowse.columns.SpringCategory[1].primary = true;   // name
    MobileBrowse.columns.SpringCategory[2].secondary = true; // description

    MobileBrowse.primaryKeys = {
        SpringListing: 'listingId',
        SpringCategory: 'categoryId'
    };
})();
```

### 9.11 Mobile Registry Index Pattern

```javascript
// m/js/spring-index.js
// Note: 'My Listings' maps to MobileBrowse (same SpringListing model, filtered by buyer).
// No separate MobileMyListings namespace — reuses MobileBrowse's columns/forms/enums.
(function() {
    'use strict';
    window.MobileSpring = Layer8MModuleRegistry.create('MobileSpring', {
        'Browse': MobileBrowse,
        'My Listings': MobileBrowse,
        'My Bids': MobileMyBids,
        'Deals': MobileDeals,
        'Reviews': MobileReviews
    });
})();
```

### 9.12 Mobile Reference Registry Pattern

```javascript
// m/js/spring-reference-registry.js
(function() {
    'use strict';
    var ref = window.Layer8RefFactory;

    window.Layer8MReferenceRegistrySpring = {
        ...ref.simple('SpringCategory', 'categoryId', 'name', 'Category'),
        ...ref.simple('SpringListing', 'listingId', 'title', 'Listing'),
        ...ref.simple('SpringBid', 'bidId', 'bidId', 'Bid'),
        ...ref.simple('SpringDeal', 'dealId', 'itemTitle', 'Deal'),
        ...ref.simple('SpringReview', 'reviewId', 'title', 'Review')
    };

    Layer8MReferenceRegistry.register(window.Layer8MReferenceRegistrySpring);
})();
```

---

## 10. Deployment

### 10.1 Docker Images

| Image | Directory | Base Image (Runtime) | K8s Kind |
|-------|-----------|---------------------|----------|
| saichler/spring | go/spring/main/ | saichler/erp-postgres | StatefulSet |
| saichler/spring-web | go/spring/ui/ | saichler/erp-security | DaemonSet (hostNetwork) |
| saichler/spring-vnet | go/spring/vnet/ | saichler/erp-security | DaemonSet (hostNetwork) |

### 10.2 build.sh Pattern

Each service directory gets a `build.sh`:
```bash
#!/usr/bin/env bash
set -e
docker build --no-cache --platform=linux/amd64 -t saichler/spring:latest .
docker push saichler/spring:latest
```

### 10.3 Dockerfile Pattern

```dockerfile
FROM saichler/builder:latest AS build
COPY main.go /home/src/github.com/saichler/build/
RUN go mod init
RUN GOPROXY=direct GOPRIVATE=github.com go mod tidy
RUN go build -o spring

FROM saichler/erp-postgres:latest AS final
COPY --from=build /home/src/github.com/saichler/build/spring /home/run/spring
ENTRYPOINT ["/home/run/spring"]
```

### 10.4 K8s Manifests

Each manifest follows l8erp patterns with required entries:

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: spring
  labels:
    name: spring
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  namespace: spring
  name: spring
  labels:
    app: spring
spec:
  selector:
    matchLabels:
      app: spring
  template:
    spec:
      containers:
        - name: spring
          image: saichler/spring:latest
          imagePullPolicy: Always
          env:
            - name: NODE_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.hostIP
          volumeMounts:
            - name: hdata
              mountPath: /data
      volumes:
        - name: hdata
          hostPath:
            path: /data
            type: DirectoryOrCreate
```

### 10.5 build-all-images.sh

```bash
#!/usr/bin/env bash
set -e
cd spring/vnet && ./build.sh && cd ../..
cd spring/main && ./build.sh && cd ../..
cd spring/ui && ./build.sh && cd ../..
```

### 10.6 k8s/deploy.sh

```bash
#!/usr/bin/env bash
kubectl apply -f ./spring-vnet.yaml
kubectl apply -f ./spring.yaml
kubectl apply -f ./spring-web.yaml
```

### 10.7 k8s/undeploy.sh

```bash
#!/usr/bin/env bash
kubectl delete -f ./spring-web.yaml
kubectl delete -f ./spring.yaml
kubectl delete -f ./spring-vnet.yaml
```

---

## 11. Local Development Setup

A `run-local.sh` script will be created at `go/run-local.sh`, copied and adapted from `l8erp/go/run-local.sh`. It will:

1. Clean and fetch dependencies (`go mod init`, `go mod tidy`, `go mod vendor`)
2. Start PostgreSQL container (`saichler/unsecure-postgres`)
3. Build all binaries into `demo/` directory:
   - `mocks_demo` (mock data generator)
   - `vnet_demo` (virtual network)
   - `spring_demo` (backend server)
   - `ui_demo` (web UI server)
4. Copy web assets to `demo/`
5. Generate `kill_demo.sh` cleanup script
6. Start services in order: vnet → spring → ui
7. Upload mock data via `mocks_demo`
8. Wait for user, then clean up

---

## 12. Mock Data

### 12.1 Data Arrays (data.go)

Curated arrays for realistic mock data:
- Product names (electronics, furniture, clothing, sporting goods, books, etc.)
- Category names (Electronics, Home & Garden, Fashion, Sports, Books, Automotive, etc.)
- Listing titles combining product type + specific model
- Review text templates (positive, neutral, negative)
- Location names (cities across US)
- Shipping carrier names

### 12.2 Store (store.go)

```go
type MockDataStore struct {
    // Phase 1: Foundation
    CategoryIDs []string

    // Phase 2: Listings
    ListingIDs []string

    // Phase 3: Bids
    BidIDs []string

    // Phase 4: Deals
    DealIDs []string

    // Phase 5: Reviews
    ReviewIDs []string
}
```

### 12.3 Phase Ordering

| Phase | Entities | Dependencies |
|-------|----------|-------------|
| 1 | SpringCategory (30 categories) | None |
| 2 | SpringListing (200 listings) | CategoryIDs |
| 3 | SpringBid (500 bids) | ListingIDs |
| 4 | SpringDeal (100 deals) | ListingIDs, BidIDs |
| 5 | SpringReview (150 reviews) | DealIDs |

### 12.4 Mock Data Distributions

- **Listings:** 10% Draft, 50% Active, 25% Fulfilled, 10% Expired, 5% Cancelled
- **Bids:** 40% Pending, 20% Accepted, 20% Rejected, 10% Withdrawn, 10% Expired
- **Deals:** 10% Pending Payment, 15% Paid, 20% Shipped, 15% Delivered, 30% Completed, 5% Disputed, 3% Cancelled, 2% Refunded
- **Reviews:** 5% Pending, 90% Published, 3% Flagged, 2% Removed
- **Ratings:** Normal distribution centered at 4 stars (1=5%, 2=10%, 3=15%, 4=35%, 5=35%)
- **Prices:** Range $5 - $5,000, weighted toward $20-$500 range

---

## 13. Duplication Audit

Per `plan-duplication-audit.md`, this section audits all planned files for behavioral vs. configuration code to confirm no duplication.

### Behavioral Code (shared — lives in l8ui, NOT duplicated)

All of the following is handled by existing l8ui shared components. No sub-module reimplements any of it:

| Behavior | Shared Component |
|----------|-----------------|
| Table rendering, pagination, sorting, filtering | `Layer8DTable` / `Layer8MEditTable` |
| Form rendering, data collection, validation | `Layer8DForms` / `Layer8MForms` |
| Detail popup open/close, view/edit/add modes | `Layer8DFormsModal` / `Layer8MPopup` |
| Tab navigation, sub-nav switching | `Layer8DModuleNavigation` |
| CRUD operations (POST/PUT/DELETE) | `Layer8DModuleCRUD` / `Layer8MNavCrud` |
| Service registry, endpoint resolution | `Layer8DServiceRegistry` |
| Module bootstrap (wiring all of the above) | `Layer8DModuleFactory` / `Layer8MModuleRegistry` |
| View type switching (kanban, tree, timeline) | `Layer8DViewFactory` / `Layer8MViewFactory` |
| Reference picker search/select | `Layer8DReferencePicker` / `Layer8MReferencePicker` |
| Dashboard KPI widgets | `Layer8DWidget` |

### Configuration Code (per sub-module — unique data only)

Each sub-module produces only data files:

| File | Content | Lines (est.) |
|------|---------|:---:|
| `*-enums.js` | Enum maps, status classes, renderers | ~30-50 |
| `*-columns.js` | Column definitions (keys, labels, renderers) | ~30-60 |
| `*-forms.js` | Form definitions (sections, fields, types) | ~40-80 |

### Shared files (created once, not per sub-module)

| File | Content | Lines (est.) |
|------|---------|:---:|
| `marketplace-config.js` | Module/service registry (all sub-modules) | ~30 |
| `marketplace-init.js` | Single `Layer8DModuleFactory.create()` call | ~12 |
| `reference-registry-spring.js` | All model reference configs | ~20 |
| `spring-index.js` (mobile) | `Layer8MModuleRegistry.create()` call | ~10 |

### Duplication Calculation

- Behavioral lines per sub-module: **0** (all in shared l8ui components)
- Configuration lines per sub-module: ~100-190 (enums + columns + forms)
- Total structural/behavioral code for new module: ~72 lines (config + init + registry + mobile index)
- This is well under the 50-line threshold in `maintainability.md` for structural boilerplate

### Cross-Sub-Module Dedup: Browse and My Listings

Browse and My Listings both operate on the same model (`SpringListing`) with the same endpoint (`/10/Listing`). The only difference is a `baseWhereClause` filter (`buyerId=<currentUser>` for My Listings). Therefore:

- **No separate mylistings/ data files** (enums, columns, forms) are created — desktop or mobile
- Both sub-modules use the `Browse` namespace's definitions
- The mobile registry maps `'My Listings'` to `MobileBrowse` (not a separate `MobileMyListings`)
- The desktop `submodules` array omits `'MyListings'` — it shares `Browse`

This eliminates 6 unnecessary files (3 desktop + 3 mobile) that would have been exact duplicates.

### Conclusion

No extraction phase is needed. The shared abstraction layer (l8ui) already exists and handles 100% of behavioral logic. Each sub-module is config-only. The browse/mylistings deduplication ensures no data files are duplicated across sub-modules that share the same model.

---

## 14. Traceability Matrix

Every gap and action item from the analysis sections mapped to the implementation phase that addresses it.

| # | Source Section | Gap / Action Item | Phase |
|---|---------------|-------------------|-------|
| 1 | 3.1 Data Model | Define SpringCategory proto message + List type | Phase 1 |
| 2 | 3.1 Data Model | Define SpringListing proto message + embedded images + List type | Phase 1 |
| 3 | 3.1 Data Model | Define SpringBid proto message + embedded images + List type | Phase 1 |
| 4 | 3.1 Data Model | Define SpringDeal proto message + embedded messages/dispute + List type | Phase 1 |
| 5 | 3.1 Data Model | Define SpringReview proto message + List type | Phase 1 |
| 6 | 4 Enums | Define all 8 enums in proto files with UNSPECIFIED zero values | Phase 1 |
| 7 | 5.1 Workflow | Implement listing status transitions (Draft→Active→Fulfilled/Expired/Cancelled) | Phase 2 |
| 8 | 5.2 Workflow | Implement bid status transitions (Pending→Accepted/Rejected/Withdrawn/Expired) | Phase 2 |
| 9 | 5.3 Workflow | Implement deal status transitions (PendingPayment→Paid→Shipped→Delivered→Completed) | Phase 2 |
| 10 | 5.3 Workflow | Implement deal dispute flow (Delivered→Disputed→Resolved/Escalated) | Phase 2 |
| 11 | 5.2+5.3 Workflow | Business logic: accept bid → create deal + update listing/bid statuses | Phase 2 |
| 12 | 6.1 UI | Create marketplace module config with 5 sub-modules | Phase 3 |
| 13 | 6.1 UI | Create browse sub-module (enums, columns, forms) for Listing + Category | Phase 3 |
| 14 | 6.1 UI | *(Removed — mylistings reuses Browse data files, no separate sub-module)* | N/A |
| 15 | 6.1 UI | Create section HTML with correct container IDs | Phase 3 |
| 16 | 6.1 UI | Create marketplace init file (sectionSelector = defaultModule) | Phase 3 |
| 17 | 6.1 UI | Wire into app.html and sections.js | Phase 3 |
| 18 | 6.1 UI | Create mybids sub-module (enums, columns, forms) | Phase 4 |
| 19 | 6.1 UI | Create deals sub-module with inline message table | Phase 4 |
| 20 | 6.1 UI | Create reviews sub-module (enums, columns, forms) | Phase 4 |
| 21 | 6.2 Views | Add kanban view config for listings (by status) | Phase 6 |
| 22 | 6.2 Views | Add kanban view config for deals (by status) | Phase 6 |
| 23 | 6.2 Views | Add tree view config for categories (by parentCategoryId) | Phase 6 |
| 24 | 6.2 Views | Add timeline view config for bids and deals | Phase 6 |
| 25 | 6.3 Dashboard | Create dashboard with 6 KPI widgets | Phase 6 |
| 26 | 6.4 Mobile | Create m/js/browse/ enums, columns (primary/secondary), forms — parity with desktop browse/ | Phase 5 |
| 27 | 6.4 Mobile | *(Removed — mylistings reuses m/js/browse/ data files, no separate mobile sub-module)* | N/A |
| 28 | 6.4 Mobile | Create m/js/mybids/ enums, columns (primary/secondary), forms — parity with desktop mybids/ | Phase 5 |
| 29 | 6.4 Mobile | Create m/js/deals/ enums, columns (primary/secondary), forms — parity with desktop deals/ | Phase 5 |
| 30 | 6.4 Mobile | Create m/js/reviews/ enums, columns (primary/secondary), forms — parity with desktop reviews/ | Phase 5 |
| 31 | 6.4 Mobile | Create mobile-config-spring.js nav config (hasSubModules, services with idField) | Phase 5 |
| 32 | 6.4 Mobile | Create spring-index.js (Layer8MModuleRegistry.create with all 5 sub-modules) | Phase 5 |
| 33 | 6.4 Mobile | Create spring-reference-registry.js + register with Layer8MReferenceRegistry | Phase 5 |
| 34 | 6.4 Mobile | Create m/app.html with script tags in correct loading order | Phase 5 |
| 35 | 6.4 Mobile | Create m/app-core.js adapted from l8erp mobile pattern | Phase 5 |
| 36 | 6.4 Mobile | Add sidebar link: data-section="dashboard" data-module="marketplace" | Phase 5 |
| 37 | 6.4 Mobile | Register MobileSpring in Layer8MNav data lookup arrays (_getServiceColumns, _getServiceFormDef, _getServiceTransformData) | Phase 5 |
| 38 | 6.4 Mobile | Verify desktop/mobile parity for all 5 sub-modules (feature parity checklist) | Phase 5 |
| 39 | 6.5 Detail | Listing detail popup: overview + bids inline table + images (FileStore) | Phase 4 |
| 39b | 6.5 Detail | Image upload via FileStore pattern: inline table with `type: 'file'` columns for listing + bid images | Phase 3+4 |
| 40 | 6.6 Detail | Deal detail popup: overview + shipping + messages inline + dispute | Phase 4 |
| 41 | 7.2 Security | Implement ISecurityProvider with User/Admin roles | Phase 1 |
| 41b | 7.3 Security Config | Create spring.json with allow rules, deny/scope rules, field denials, pre-defined admin user | Phase 1 |
| 42 | 8 Services | Create SpringCategory service + callback | Phase 1 |
| 43 | 8 Services | Create SpringListing service + callback | Phase 1 |
| 44 | 8 Services | Create SpringBid service + callback | Phase 2 |
| 45 | 8 Services | Create SpringDeal service + callback | Phase 2 |
| 46 | 8 Services | Create SpringReview service + callback | Phase 2 |
| 47 | 8 Services | Create reference registry entries for all 5 models (desktop + mobile) | Phase 3+4+5 |
| 48 | 9.7 login.json | Create login.json with apiPrefix=/spring | Phase 1 |
| 49 | 9.8 defaults | Create common/defaults.go with PREFIX and re-exports | Phase 1 |
| 50 | 10 Deployment | Create build.sh + Dockerfile for spring, spring-web, spring-vnet | Phase 8 |
| 51 | 10 Deployment | Create K8s manifests with required entries | Phase 8 |
| 52 | 10 Deployment | Create build-all-images.sh, deploy.sh, undeploy.sh | Phase 8 |
| 53 | 11 Local Dev | Create run-local.sh adapted from l8erp | Phase 8 |
| 54 | 12 Mock Data | Create data.go with curated arrays | Phase 7 |
| 55 | 12 Mock Data | Create store.go with ID slices | Phase 7 |
| 56 | 12 Mock Data | Create gen_spring_*.go generator files (5 files) | Phase 7 |
| 57 | 12 Mock Data | Create main_phases.go phase orchestration | Phase 7 |
| 58 | 12 Mock Data | Create mock CLI entry point | Phase 7 |
| 59 | Testing | CRUD tests for all 5 services | Phase 7 |
| 60 | Testing | Validation tests (required fields, enums, status transitions) | Phase 7 |
| 61 | Testing | Business logic tests (accept bid → create deal) | Phase 7 |
| 62 | Verification | End-to-end smoke test of all sections on desktop + mobile | Phase 9 |

61 items assigned to phases, 2 removed (mylistings dedup — items 14 and 27). No orphaned gaps.

---

## 15. Implementation Phases

### Phase 1: Foundation & Infrastructure
- [ ] Create proto files for all 5 Prime Objects with enums
- [ ] Run `make-bindings.sh` to generate `.pb.go` files
- [ ] Create `go/spring/common/defaults.go` with PREFIX and re-exports
- [ ] Create SpringCategory service + callback (with validation)
- [ ] Create SpringListing service + callback (with validation)
- [ ] Create security config (`go/secure/plugin/spring/spring.json`) with roles, allow/deny rules, row-level scoping, and pre-defined admin user
- [ ] Set up l8ui submodule: copy `../l8ui/setup-l8ui-submodule.sh` to `go/spring/ui/web/` and run it
- [ ] Create `login.json` with correct apiPrefix
- [ ] Create backend `main.go` (register services, start listening)
- [ ] Create vnet `main.go`
- [ ] Create UI `main.go` (register types, serve web assets)
- [ ] Verify `go build ./...` passes

### Phase 2: Core Services
- [ ] Create SpringBid service + callback (with validation)
- [ ] Create SpringDeal service + callback (with validation)
- [ ] Create SpringReview service + callback (with validation)
- [ ] Implement status transition validation for Listing, Bid, Deal
- [ ] Implement business logic: accepting a bid creates a deal and updates listing/bid statuses
- [ ] Verify all services compile and activate

### Phase 3: Desktop UI - Browse & Listings
- [ ] Create marketplace module config (`marketplace-config.js`) — mylistings tab reuses Browse namespace
- [ ] Create browse sub-module: enums, columns, forms for SpringListing (with image inline table using `type: 'file'`) and SpringCategory (shared by mylistings)
- [ ] Create section HTML (`sections/marketplace.html`) — includes browse and mylistings tabs
- [ ] Create marketplace init file (`marketplace-init.js`)
- [ ] Wire into `app.html` (script tags) and `sections.js`
- [ ] Create reference registry for SpringCategory
- [ ] Create `app.js` (adapted from l8erp, remove ModConfig/currency calls)
- [ ] Verify listings display in table and detail popup (both browse and mylistings tabs)

### Phase 4: Desktop UI - Bids, Deals, Reviews
- [ ] Create mybids sub-module: enums, columns, forms (with image inline table using `type: 'file'`)
- [ ] Create deals sub-module: enums, columns, forms (with inline message table)
- [ ] Create reviews sub-module: enums, columns, forms
- [ ] Add reference registry entries for all models
- [ ] Add view type support (kanban for listings/deals, tree for categories)
- [ ] Verify all sections render correctly

### Phase 5: Mobile UI
- [ ] Create `m/js/browse/` enums, columns (with `primary`/`secondary` markers), forms — parity with desktop browse/ (shared by mylistings)
- [ ] Create `m/js/mybids/` enums, columns (with `primary`/`secondary`), forms — parity with desktop mybids/
- [ ] Create `m/js/deals/` enums, columns (with `primary`/`secondary`), forms — parity with desktop deals/
- [ ] Create `m/js/reviews/` enums, columns (with `primary`/`secondary`), forms — parity with desktop reviews/
- [ ] Create `mobile-config-spring.js` nav config (`hasSubModules: true`, services with `idField`, `endpoint`, `model`)
- [ ] Create `spring-index.js` (`Layer8MModuleRegistry.create` — map 'My Listings' to MobileBrowse, no separate namespace)
- [ ] Create `spring-reference-registry.js` and register with `Layer8MReferenceRegistry.register()`
- [ ] Create `m/app.html` with all script tags in correct loading order (per `mobile-script-loading-order.md`)
- [ ] Create `m/app-core.js` adapted from l8erp mobile pattern (no ModConfig call)
- [ ] Add sidebar link: `<a data-section="dashboard" data-module="marketplace">Marketplace</a>`
- [ ] Register `window.MobileSpring` in Layer8MNav data lookup arrays (`_getServiceColumns`, `_getServiceFormDef`, `_getServiceTransformData`)
- [ ] Run feature parity verification checklist (section 6.4) for all sub-modules on both platforms

### Phase 6: Dashboard & Advanced Views
- [ ] Create dashboard with KPI widgets (active listings, deals, ratings)
- [ ] Add kanban view for listings (by status)
- [ ] Add kanban view for deals (by status)
- [ ] Add tree view for categories
- [ ] Add timeline view for bids and deals

### Phase 7: Mock Data & Testing
- [ ] Create `data.go` with curated name arrays
- [ ] Create `store.go` with ID slices
- [ ] Create generator files: categories, listings, bids, deals, reviews
- [ ] Create phase orchestration (`main_phases.go`)
- [ ] Create mock CLI entry point (`cmd/main.go`)
- [ ] Verify `go build ./tests/mocks/` passes
- [ ] Create integration tests in `go/tests/`
- [ ] CRUD tests for all 5 services
- [ ] Validation tests (required fields, enum validation, status transitions)
- [ ] Business logic tests (accept bid → create deal flow)

### Phase 8: Deployment & Local Dev
- [ ] Create `build.sh` and `Dockerfile` for all 3 images (spring, spring-web, spring-vnet)
- [ ] Create `build-all-images.sh`
- [ ] Create K8s manifests (spring.yaml, spring-web.yaml, spring-vnet.yaml)
- [ ] Create `k8s/deploy.sh` and `k8s/undeploy.sh`
- [ ] Create `run-local.sh` (adapted from l8erp)
- [ ] End-to-end verification: start system, upload mocks, browse in browser

### Phase 9: End-to-End Verification

For every section affected by this implementation:
1. Navigate to each section in the desktop UI
2. Verify table data loads (not blank)
3. Verify row click opens detail popup with populated data
4. Verify CRUD operations (add, edit, delete) work
5. Verify status transitions work (accept bid → deal created)
6. Verify kanban/tree/timeline views render correctly
7. Repeat all checks on mobile
8. Verify dashboard KPI widgets show correct counts

Sections to verify:
- [ ] Browse > Listings
- [ ] Browse > Categories
- [ ] My Listings
- [ ] My Bids
- [ ] Deals
- [ ] Reviews
- [ ] Dashboard
- [ ] SYS > Health
- [ ] SYS > Security
- [ ] Mobile: all of the above

---

## 16. Non-Functional Requirements

| Requirement | Target |
|-------------|--------|
| Page load time | < 2 seconds |
| API response time | < 500ms (p95) |
| Concurrent users | 1,000+ |
| Listing search | < 1 second for filtered queries |
| Database | PostgreSQL via l8orm |
| Uptime | 99.9% |
| Browser support | Chrome, Firefox, Safari, Edge (latest 2 versions) |
| Mobile support | iOS Safari, Android Chrome |

---

## 17. Compliance Checklist

### Project Structure & Architecture
- [x] Project structure follows l8erp layout (directory names, file naming, organization)
- [x] Go module root at `go/` with standard subdirectories
- [x] Proto files in `proto/` with `make-bindings.sh`
- [x] Types generated into `go/types/spring/`
- [x] Tests in `go/tests/` (not alongside source code)

### Protobuf Design
- [x] Enum zero values are `_UNSPECIFIED` (all 8 enums verified)
- [x] List types use `repeated X list = 1` + `l8api.L8MetaData metadata = 2` convention
- [x] No direct struct references between Prime Objects — only ID fields (string)
- [x] Child entities (images, messages, disputes) are embedded `repeated` fields, not separate services
- [x] Proto package is `spring`, go_package is `./types/spring`

### Service Design
- [x] All ServiceName values are 10 characters or less (Category=8, Listing=7, Bid=3, Deal=4, Review=6)
- [x] ServiceArea is consistent within module (all use byte(10))
- [x] ServiceCallback auto-generates primary key on POST via `common.GenerateID()`
- [x] Types will be registered in UI main.go via `introspect.AddPrimaryKeyDecorator` + `registry.Register`

### File Upload
- [x] Image fields use `storage_path` string (not binary data, not external URLs) per `file-upload-pattern.md`
- [x] Image upload/download via global FileStore service (service area 0) — no custom upload endpoints
- [x] Form definitions use `f.inlineTable()` with `type: 'file'` columns for image galleries
- [x] Proto fields follow `storage_path`/`file_name`/`file_size` convention

### Prime Object Rules
- [x] All 5 entities pass the Prime Object test (documented in section 3.2)
- [x] All 4 embedded children justified as non-Prime (documented in section 3.3)
- [x] No standalone UI (config, columns, forms, nav) for embedded children
- [x] Children rendered as inline tables within parent forms

### UI Design
- [x] All UI module integration steps planned (config, enums, columns, forms, init, section HTML, app.html wiring, sections.js)
- [x] Desktop and mobile parity addressed (Phase 5 dedicated to mobile)
- [x] CSS classes use `l8-` prefix (shared from `layer8-section-layout.css`)
- [x] Table container IDs follow `{moduleKey}-{serviceKey}-table-container` pattern
- [x] View types use registered view types via Layer8DViewFactory (kanban, tree, timeline)
- [x] `--layer8d-*` CSS custom properties for theming
- [x] No ModConfig/currency fetches in app.js (not applicable to this project)
- [x] login.json adapted with correct apiPrefix (`/spring`) and appTitle

### Mock Data
- [x] All 5 services have mock data generators planned
- [x] Phase ordering accounts for dependencies (categories → listings → bids → deals → reviews)
- [x] Store has ID slices for all entities
- [x] Realistic distributions documented

### Deployment
- [x] New services include build.sh, Dockerfile, K8s YAML
- [x] K8s YAMLs include namespace labels, resource labels, NODE_IP env, hdata volume
- [x] `build-all-images.sh` planned
- [x] `k8s/deploy.sh` and `k8s/undeploy.sh` planned
- [x] `run-local.sh` section included (adapted from l8erp)

### Configuration
- [x] login.json adapted from l8erp (appTitle, apiPrefix changed)
- [x] ModConfig handling addressed (removed — not applicable)
- [x] Currency/exchange rate caches not loaded (not needed)

### Security
- [x] Authentication via `ifs.ISecurityProvider` (no custom auth)
- [x] Users/roles provisioned via security config JSON or Security API
- [x] No programmatic user creation in mock generators
- [x] Security config JSON designed with allow rules, deny/scope rules, and field-level denials
- [x] Row-level data scoping uses `${userId}` deny rules (no custom ServiceCallback filtering)
- [x] Public vs. user-scoped data explicitly documented per entity type

### Testing
- [x] Tests planned in `go/tests/` directory (not alongside source)
- [x] CRUD tests for all services
- [x] Validation tests for required fields and enums
- [x] Business logic tests for bid acceptance flow
- [x] End-to-end verification phase included

### Code Quality
- [x] All files planned to stay under 500 lines
- [x] No Go generics used
- [x] Shared abstractions for common patterns (Layer8DModuleFactory, Layer8MModuleRegistry)
- [x] No duplicate behavioral code across sub-modules
