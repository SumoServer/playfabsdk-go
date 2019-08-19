package economy

import "time"
                    
// AddVirtualCurrenciesRequest given an entity type, entity identifier and container details, will increase the entity's currencies by a specific
// amount.
type AddVirtualCurrenciesRequestModel struct {
    // Currencies a list of currencies to modify
    Currencies []CurrencyDetailsModel `json:"Currencies,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// AddVirtualCurrenciesResult 
type AddVirtualCurrenciesResultModel struct {
    // Currencies 
    Currencies []CurrencyResponseDtoModel `json:"Currencies,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// CatalogAlternateId 
type CatalogAlternateIdModel struct {
    // Type type of the alternate Id
    Type string `json:"Type,omitempty"`
    // Value value of the alternate Id
    Value string `json:"Value,omitempty"`
}

// CatalogConfig 
type CatalogConfigModel struct {
    // Admins a list of title player accounts that will have admin permissions.
    Admins []string `json:"Admins,omitempty"`
    // Catalog the set of configuration that only applies to catalog items.
    Catalog *CatalogSpecificConfigModel `json:"Catalog,omitempty"`
    // DeepLinkFormats a list of deep link formats.
    DeepLinkFormats []DeepLinkFormatModel `json:"DeepLinkFormats,omitempty"`
    // DisplayPropertyIndexInfos a list of display properties to index.
    DisplayPropertyIndexInfos []DisplayPropertyIndexInfoModel `json:"DisplayPropertyIndexInfos,omitempty"`
    // Reviewers a set of player entity keys that are allowed to review content.
    Reviewers []string `json:"Reviewers,omitempty"`
    // UserGeneratedContent the set of configuration that only applies to user generated contents.
    UserGeneratedContent *UserGeneratedContentSpecificConfigModel `json:"UserGeneratedContent,omitempty"`
}

// CatalogItemMetadata 
type CatalogItemMetadataModel struct {
    // AllowMultipleStacks indicates whether there can be multiple stacks of this item.
    AllowMultipleStacks bool `json:"AllowMultipleStacks,omitempty"`
    // AlternateIds the alternate IDs associated with this item.
    AlternateIds []CatalogAlternateIdModel `json:"AlternateIds,omitempty"`
    // Contents the set of contents associated with this item.
    Contents []ContentModel `json:"Contents,omitempty"`
    // ContentType the client-defined type of the item.
    ContentType string `json:"ContentType,omitempty"`
    // CreationDate the date and time when this item was created.
    CreationDate time.Time `json:"CreationDate,omitempty"`
    // CreatorEntityKey the ID of the creator of this catalog item.
    CreatorEntityKey *EntityKeyModel `json:"CreatorEntityKey,omitempty"`
    // DeepLinks the set of platform specific deep links for this item.
    DeepLinks []DeepLinkModel `json:"DeepLinks,omitempty"`
    // Description a dictionary of localized descriptions. Key is language code and localized string is the value. The neutral locale is
// required.
    Description map[string]string `json:"Description,omitempty"`
    // DisplayProperties game specific properties for display purposes. This is an arbitrary JSON blob.
    DisplayProperties interface{} `json:"DisplayProperties,omitempty"`
    // DisplayVersion the user provided version of the item for display purposes.
    DisplayVersion string `json:"DisplayVersion,omitempty"`
    // EndDate the date of when the item will cease to be available. If not provided then the product will be available indefinitely.
    EndDate time.Time `json:"EndDate,omitempty"`
    // ETag the current ETag value that can be used for optimistic concurrency in the If-None-Match header.
    ETag string `json:"ETag,omitempty"`
    // Id the unique ID of the item.
    Id string `json:"Id,omitempty"`
    // Images the images associated with this item. Images can be thumbnails or screenshots.
    Images []ImageModel `json:"Images,omitempty"`
    // IsConsumable indicates whether the item is a consumable or durable.
    IsConsumable bool `json:"IsConsumable,omitempty"`
    // IsHidden indicates if the item is hidden.
    IsHidden bool `json:"IsHidden,omitempty"`
    // IsStackable indicates if this item can be stacked or if they should be handled individually.
    IsStackable bool `json:"IsStackable,omitempty"`
    // ItemReferences the item references associated with this item.
    ItemReferences []CatalogItemReferenceModel `json:"ItemReferences,omitempty"`
    // LastModifiedDate the date and time this item was last updated.
    LastModifiedDate time.Time `json:"LastModifiedDate,omitempty"`
    // Moderation the moderation state for this item.
    Moderation *ModerationStateModel `json:"Moderation,omitempty"`
    // PayoutInfo the payout information of the payee.
    PayoutInfo *PayoutInfoModel `json:"PayoutInfo,omitempty"`
    // Platforms the platforms supported by this item.
    Platforms []string `json:"Platforms,omitempty"`
    // Price the base price of this item.
    Price *CatalogPriceModel `json:"Price,omitempty"`
    // Rating rating summary for this UGC item.
    Rating *RatingModel `json:"Rating,omitempty"`
    // SourceEntityKey the title or namespace that this item belongs to.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
    // StartDate the date of when the item will be available. If not provided then the product will appear immediately.
    StartDate time.Time `json:"StartDate,omitempty"`
    // Subscription additional metadata for subscription items.
    Subscription *SubscriptionDetailsModel `json:"Subscription,omitempty"`
    // Tags the list of tags that are associated with this item.
    Tags []string `json:"Tags,omitempty"`
    // Title a dictionary of localized titles. Key is language code and localized string is the value. The neutral locale is
// required.
    Title map[string]string `json:"Title,omitempty"`
    // Type the high-level type of the item.
    Type string `json:"Type,omitempty"`
}

// CatalogItemReference 
type CatalogItemReferenceModel struct {
    // Amount the amount of the catalog item.
    Amount int32 `json:"Amount,omitempty"`
    // Id the unique ID of the catalog item.
    Id string `json:"Id,omitempty"`
    // Price the price of the catalog item.
    Price *CatalogPriceModel `json:"Price,omitempty"`
}

// CatalogPrice 
type CatalogPriceModel struct {
    // PayoutAmount the amount the payee will receive.
    PayoutAmount float64 `json:"PayoutAmount,omitempty"`
    // Prices prices of the catalog item.
    Prices []CatalogPriceInstanceModel `json:"Prices,omitempty"`
    // Sort sort setting of the catalog item.
    Sort int32 `json:"Sort,omitempty"`
}

// CatalogPriceAmount 
type CatalogPriceAmountModel struct {
    // Amount the amount of the catalog price.
    Amount int32 `json:"Amount,omitempty"`
    // CurrencyId the currency ID of the catalog price.
    CurrencyId string `json:"CurrencyId,omitempty"`
}

// CatalogPriceInstance 
type CatalogPriceInstanceModel struct {
    // Amounts the amounts of the catalog item price.
    Amounts []CatalogPriceAmountModel `json:"Amounts,omitempty"`
}

// CatalogSearchRequest 
type CatalogSearchRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Filter an OData filter used to refine the search query.
    Filter string `json:"Filter,omitempty"`
    // OrderBy an OData orderBy used to order the results of the search query.
    OrderBy string `json:"OrderBy,omitempty"`
    // Search the text to search for.
    Search string `json:"Search,omitempty"`
    // Select an OData select query option used to augment the search results. If not defined, the default search result metadata will
// be returned.
    Select string `json:"Select,omitempty"`
    // Skip the number of results to skip.
    Skip int32 `json:"Skip,omitempty"`
    // SourceEntityKey the title or namespace to search under.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
    // TitleId the ID of the title that this catalog item is associated with.
    TitleId string `json:"TitleId,omitempty"`
    // Top the number of results to retrieve.
    Top int32 `json:"Top,omitempty"`
}

// CatalogSearchResult 
type CatalogSearchResultModel struct {
    // Count the total count of hits for the search query.
    Count int32 `json:"Count,omitempty"`
    // Items the paginated set of results for the search query.
    Items []CatalogItemMetadataModel `json:"Items,omitempty"`
}

// CatalogSpecificConfig 
type CatalogSpecificConfigModel struct {
    // ContentTypes the set of content types that will be used for validation and if no values are provided then anything is allowed.
    ContentTypes []string `json:"ContentTypes,omitempty"`
    // Tags the set of tags that will be used for validation and if no values are provided then anything is allowed.
    Tags []string `json:"Tags,omitempty"`
}

// ConcernCategory 
type ConcernCategory string
  
const (
     ConcernCategoryNone ConcernCategory = "None"
     ConcernCategoryOffensiveContent ConcernCategory = "OffensiveContent"
     ConcernCategoryChildExploitation ConcernCategory = "ChildExploitation"
     ConcernCategoryMalwareOrVirus ConcernCategory = "MalwareOrVirus"
     ConcernCategoryPrivacyConcerns ConcernCategory = "PrivacyConcerns"
     ConcernCategoryMisleadingApp ConcernCategory = "MisleadingApp"
     ConcernCategoryPoorPerformance ConcernCategory = "PoorPerformance"
     ConcernCategoryReviewResponse ConcernCategory = "ReviewResponse"
     ConcernCategorySpamAdvertising ConcernCategory = "SpamAdvertising"
     ConcernCategoryProfanity ConcernCategory = "Profanity"
      )
// ConsumeInventoryItemsRequest given an entity type, entity identifier and container details, will consume the specified inventory items.
type ConsumeInventoryItemsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items a list of Items to modify
    Items []InventoryItemDetailsModel `json:"Items,omitempty"`
}

// ConsumeInventoryItemsResult 
type ConsumeInventoryItemsResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items 
    Items []InventoryItemResponseDtoModel `json:"Items,omitempty"`
}

// ContainerType 
type ContainerType string
  
const (
     ContainerTypeNone ContainerType = "None"
     ContainerTypeBundles ContainerType = "Bundles"
     ContainerTypeStores ContainerType = "Stores"
     ContainerTypeSubscriptions ContainerType = "Subscriptions"
      )
// Content 
type ContentModel struct {
    // Id 
    Id string `json:"Id,omitempty"`
    // MaxClientVersion the maximum client version that this content is compatible with.
    MaxClientVersion string `json:"MaxClientVersion,omitempty"`
    // MinClientVersion the minimum client version that this content is compatible with.
    MinClientVersion string `json:"MinClientVersion,omitempty"`
    // Tags the list of tags that are associated with this content.
    Tags []string `json:"Tags,omitempty"`
    // Type the client-defined type of the content.
    Type string `json:"Type,omitempty"`
    // Url the Azure CDN URL for retrieval of the catalog item binary content.
    Url string `json:"Url,omitempty"`
}

// CreateBundleRequest create bundle request
type CreateBundleRequestModel struct {
    // AllowOverwrite allow overwrite
    AllowOverwrite bool `json:"AllowOverwrite,omitempty"`
    // Bundle bundle details
    Bundle CatalogItemMetadataModel `json:"Bundle,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// CreateBundleResult create bundle result
type CreateBundleResultModel struct {
    // Bundle 
    Bundle *CatalogItemMetadataModel `json:"Bundle,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// CreateCurrencyRequest given an entity type, and currency details, create a currency.
type CreateCurrencyRequestModel struct {
    // Currency metadata describing the new currency item to be created.
    Currency CatalogItemMetadataModel `json:"Currency,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// CreateCurrencyResult 
type CreateCurrencyResultModel struct {
    // CurrencyMetadata updated metadata describing the catalog item just created.
    CurrencyMetadata *CatalogItemMetadataModel `json:"CurrencyMetadata,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// CreateDraftItemRequest the item will not be published to the public catalog until the user makes a call to the PublishItem API.
type CreateDraftItemRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Item metadata describing the new catalog item to be created.
    Item CatalogItemMetadataModel `json:"Item,omitempty"`
}

// CreateDraftItemResult 
type CreateDraftItemResultModel struct {
    // Item updated metadata describing the catalog item just created.
    Item *CatalogItemMetadataModel `json:"Item,omitempty"`
}

// CreateOrUpdateReviewRequest 
type CreateOrUpdateReviewRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId the ID of the item to submit a review for.
    ItemId string `json:"ItemId,omitempty"`
    // Review the review to submit.
    Review *ReviewSubmissionModel `json:"Review,omitempty"`
}

// CreateOrUpdateReviewResult 
type CreateOrUpdateReviewResultModel struct {
}

// CreateStoreRequest create store request
type CreateStoreRequestModel struct {
    // AllowOverwrite allow overwrite
    AllowOverwrite bool `json:"AllowOverwrite,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Store store details
    Store CatalogItemMetadataModel `json:"Store,omitempty"`
}

// CreateStoreResult create store result
type CreateStoreResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Store 
    Store *CatalogItemMetadataModel `json:"Store,omitempty"`
}

// CreateSubscriptionRequest create subscription request
type CreateSubscriptionRequestModel struct {
    // AllowOverwrite allow overwrite
    AllowOverwrite bool `json:"AllowOverwrite,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Subscription subscription details
    Subscription CatalogItemMetadataModel `json:"Subscription,omitempty"`
}

// CreateSubscriptionResult create subscription result
type CreateSubscriptionResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Subscription 
    Subscription *CatalogItemMetadataModel `json:"Subscription,omitempty"`
}

// CreateUploadUrlsRequest upload URLs point to Azure Blobs; clients must follow the Microsoft Azure Storage Blob Service REST API pattern for
// uploading content. The response contains upload URLs and IDs for each file. The IDs and URLs returned must be added to
// the item metadata and commited using the CreateDraftItem or UpdateDraftItem Item APIs.
type CreateUploadUrlsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Files description of the files to be uploaded by the client.
    Files []UploadInfoModel `json:"Files,omitempty"`
    // SourceEntityKey the title or namespace that this content belongs to.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// CreateUploadUrlsResult 
type CreateUploadUrlsResultModel struct {
    // UploadUrls urls and metadata for the files to be uploaded by the client.
    UploadUrls []UploadUrlMetadataModel `json:"UploadUrls,omitempty"`
}

// CurrencyDetails 
type CurrencyDetailsModel struct {
    // AlternateId 
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // Amount 
    Amount int32 `json:"Amount,omitempty"`
    // CurrencyId 
    CurrencyId string `json:"CurrencyId,omitempty"`
    // InstanceId 
    InstanceId string `json:"InstanceId,omitempty"`
}

// CurrencyResponseDto 
type CurrencyResponseDtoModel struct {
    // Amount 
    Amount int32 `json:"Amount,omitempty"`
    // ChangedAmount 
    ChangedAmount int32 `json:"ChangedAmount,omitempty"`
    // CurrencyId 
    CurrencyId string `json:"CurrencyId,omitempty"`
}

// DeepLink 
type DeepLinkModel struct {
    // Platform target platform for this deep link.
    Platform string `json:"Platform,omitempty"`
    // Url the deep link for this platform.
    Url string `json:"Url,omitempty"`
}

// DeepLinkFormat 
type DeepLinkFormatModel struct {
    // Format the format of the deep link to return. The format should contain '{id}' to represent where the item ID should be placed.
    Format string `json:"Format,omitempty"`
    // Platform the target platform for the deep link.
    Platform string `json:"Platform,omitempty"`
}

// DeleteBundleByFriendlyIdRequest delete bundle by friendly Id request
type DeleteBundleByFriendlyIdRequestModel struct {
    // FriendlyId the friendly Id of target bundle.
    FriendlyId string `json:"FriendlyId,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // SourceEntityKey the title or namespace the bundle is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// DeleteBundleByIdRequest delete bundle by Id request
type DeleteBundleByIdRequestModel struct {
    // Id id of target bundle
    Id string `json:"Id,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // SourceEntityKey the title or namespace the bundle is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// DeleteBundleResult delete bundle result
type DeleteBundleResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// DeleteItemRequest 
type DeleteItemRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId iD of the catalog item to delete.
    ItemId string `json:"ItemId,omitempty"`
}

// DeleteItemResult 
type DeleteItemResultModel struct {
}

// DeleteStoreByFriendlyIdRequest delete store by Friendly Id request
type DeleteStoreByFriendlyIdRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // FriendlyId the friendly Id of target store.
    FriendlyId string `json:"FriendlyId,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // SourceEntityKey the title or namespace the store is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// DeleteStoreByIdRequest delete store by Id request
type DeleteStoreByIdRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Id id of target store
    Id string `json:"Id,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // SourceEntityKey the title or namespace the store is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// DeleteStoreResult delete store result
type DeleteStoreResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// DeleteSubscriptionByFriendlyIdRequest delete Subscription by friendly Id request
type DeleteSubscriptionByFriendlyIdRequestModel struct {
    // FriendlyId the friendly Id of target subscription.
    FriendlyId string `json:"FriendlyId,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // SourceEntityKey the title or namespace the subscription is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// DeleteSubscriptionByIdRequest delete subscription by Id request
type DeleteSubscriptionByIdRequestModel struct {
    // Id the Id of target subscription
    Id string `json:"Id,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // SourceEntityKey the title or namespace the subscription is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// DeleteSubscriptionResult delete subscription result
type DeleteSubscriptionResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// DisplayPropertyIndexInfo 
type DisplayPropertyIndexInfoModel struct {
    // Name the property name in the 'DisplayProperties' property to be indexed.
    Name string `json:"Name,omitempty"`
    // Type the type of the property to be indexed.
    Type DisplayPropertyType `json:"Type,omitempty"`
}

// DisplayPropertyType 
type DisplayPropertyType string
  
const (
     DisplayPropertyTypeQueryDateTime DisplayPropertyType = "QueryDateTime"
     DisplayPropertyTypeQueryDouble DisplayPropertyType = "QueryDouble"
     DisplayPropertyTypeQueryString DisplayPropertyType = "QueryString"
     DisplayPropertyTypeSearchString DisplayPropertyType = "SearchString"
      )
// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://api.playfab.com/docs/tutorials/entities/entitytypes
    Type string `json:"Type,omitempty"`
}

// GetBundleByFriendlyIdRequest get bundle by friendly Id request
type GetBundleByFriendlyIdRequestModel struct {
    // ExpandReferencedItems whether to expand the referenced items
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // FriendlyId the friendly Id of target bundle.
    FriendlyId string `json:"FriendlyId,omitempty"`
    // SourceEntityKey the title or namespace the item is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// GetBundleByIdRequest get bundle by Id request
type GetBundleByIdRequestModel struct {
    // ExpandReferencedItems whether to expand the referenced items
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // Id id of target bundle
    Id string `json:"Id,omitempty"`
    // SourceEntityKey the title or namespace the item is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// GetBundleByMarketplaceOfferIdRequest get bundle by marketplace offer Id request
type GetBundleByMarketplaceOfferIdRequestModel struct {
    // ExpandReferencedItems whether to expand the referenced items
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // MarketplaceOfferId the Marketplace offer Id of target bundle.
    MarketplaceOfferId *CatalogAlternateIdModel `json:"MarketplaceOfferId,omitempty"`
    // SourceEntityKey the title or namespace the item is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// GetBundleResult get bundle result
type GetBundleResultModel struct {
    // Bundle 
    Bundle *CatalogItemMetadataModel `json:"Bundle,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // ReferencedItems 
    ReferencedItems []CatalogItemMetadataModel `json:"ReferencedItems,omitempty"`
}

// GetCatalogConfigRequest 
type GetCatalogConfigRequestModel struct {
}

// GetCatalogConfigResult 
type GetCatalogConfigResultModel struct {
    // Config 
    Config *CatalogConfigModel `json:"Config,omitempty"`
}

// GetCurrencyByIdRequest 
type GetCurrencyByIdRequestModel struct {
    // CurrencyId iD of the currency to retrieve.
    CurrencyId string `json:"CurrencyId,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// GetCurrencyByIdResult 
type GetCurrencyByIdResultModel struct {
    // Currency full metadata of the currency requested.
    Currency *CatalogItemMetadataModel `json:"Currency,omitempty"`
}

// GetDraftItemRequest 
type GetDraftItemRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId iD of the catalog item to retrieve from the working catalog.
    ItemId string `json:"ItemId,omitempty"`
}

// GetDraftItemResult 
type GetDraftItemResultModel struct {
    // Item full metadata of the catalog item requested.
    Item *CatalogItemMetadataModel `json:"Item,omitempty"`
}

// GetDraftItemsRequest 
type GetDraftItemsRequestModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items created by the caller, if any are available. Should be null on
// initial request.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count number of items to retrieve. Maximum page size is 10.
    Count int32 `json:"Count,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// GetDraftItemsResult 
type GetDraftItemsResultModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of items, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count the total number of items created by the caller.
    Count int32 `json:"Count,omitempty"`
    // Items a set of items created by the caller.
    Items []CatalogItemMetadataModel `json:"Items,omitempty"`
}

// GetInventoryItemsRequest given an entity type, entity identifier and container details, will get the entity's inventory items.
type GetInventoryItemsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // ReceiptData property bag for receipt generation
    ReceiptData map[string]string `json:"ReceiptData,omitempty"`
    // ShouldGenerateReceipt whether or not to generate a receipt
    ShouldGenerateReceipt bool `json:"ShouldGenerateReceipt,omitempty"`
}

// GetInventoryItemsResult 
type GetInventoryItemsResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items 
    Items []InventoryItemResponseDtoModel `json:"Items,omitempty"`
    // Receipt 
    Receipt string `json:"Receipt,omitempty"`
    // Subscriptions 
    Subscriptions []SubscriptionItemResponseDtoModel `json:"Subscriptions,omitempty"`
}

// GetMyReviewRequest 
type GetMyReviewRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId the ID of the item to retrieve the user's review for.
    ItemId string `json:"ItemId,omitempty"`
}

// GetMyReviewResult 
type GetMyReviewResultModel struct {
    // Review the review the user submitted for the requested item.
    Review *ReviewModel `json:"Review,omitempty"`
}

// GetPublishedItemRequest 
type GetPublishedItemRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId iD of the catalog item to retrieve from the working catalog.
    ItemId string `json:"ItemId,omitempty"`
}

// GetPublishedItemResult 
type GetPublishedItemResultModel struct {
    // Item full metadata of the catalog item requested.
    Item *CatalogItemMetadataModel `json:"Item,omitempty"`
}

// GetReviewsRequest 
type GetReviewsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId the ID of the item to retrieve reviews for.
    ItemId string `json:"ItemId,omitempty"`
    // OrderBy an OData orderBy used to order the results of the query.
    OrderBy string `json:"OrderBy,omitempty"`
    // Skip the number of review results to skip. If not specified, defaults to 0.
    Skip int32 `json:"Skip,omitempty"`
    // Top the number of review results to retrieve. If not specified, defaults to 10.
    Top int32 `json:"Top,omitempty"`
}

// GetReviewsResult 
type GetReviewsResultModel struct {
    // ContinuationToken an opaque token used to retrieve the next page of reviews, if any are available.
    ContinuationToken string `json:"ContinuationToken,omitempty"`
    // Count the total number of reviews associated with this item.
    Count int32 `json:"Count,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Reviews the paginated set of results.
    Reviews []ReviewModel `json:"Reviews,omitempty"`
}

// GetStoreByFriendlyIdRequest get store by friendly Id request
type GetStoreByFriendlyIdRequestModel struct {
    // ExpandReferencedItems whether to expand the referenced items
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // FriendlyId the friendly Id of target store.
    FriendlyId string `json:"FriendlyId,omitempty"`
    // SourceEntityKey the title or namespace the item is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// GetStoreByIdRequest get store by Id request
type GetStoreByIdRequestModel struct {
    // ExpandReferencedItems whether to expand the referenced items
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // Id id of target store
    Id string `json:"Id,omitempty"`
    // SourceEntityKey the title or namespace the item is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// GetStoreResult get store result
type GetStoreResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // ReferencedItems 
    ReferencedItems []CatalogItemMetadataModel `json:"ReferencedItems,omitempty"`
    // Store 
    Store *CatalogItemMetadataModel `json:"Store,omitempty"`
}

// GetSubscriptionByFriendlyIdRequest get subscription by friendly Id request
type GetSubscriptionByFriendlyIdRequestModel struct {
    // ExpandReferencedItems whether to expand the referenced items
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // FriendlyId the friendly Id of target subscription.
    FriendlyId string `json:"FriendlyId,omitempty"`
    // SourceEntityKey the title or namespace the item is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// GetSubscriptionByIdRequest get subscription by Id request
type GetSubscriptionByIdRequestModel struct {
    // ExpandReferencedItems whether to expand the referenced items
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // Id id of target subscription
    Id string `json:"Id,omitempty"`
    // SourceEntityKey the title or namespace the item is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// GetSubscriptionByMarketplaceOfferIdRequest get subscription by marketplace offer Id request
type GetSubscriptionByMarketplaceOfferIdRequestModel struct {
    // ExpandReferencedItems whether to expand the referenced items
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // MarketplaceOfferId the Marketplace offer Id of target bundle.
    MarketplaceOfferId *CatalogAlternateIdModel `json:"MarketplaceOfferId,omitempty"`
    // SourceEntityKey the title or namespace the item is in.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// GetSubscriptionResult get subscription result
type GetSubscriptionResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // ReferencedItems 
    ReferencedItems []CatalogItemMetadataModel `json:"ReferencedItems,omitempty"`
    // Subscription 
    Subscription *CatalogItemMetadataModel `json:"Subscription,omitempty"`
}

// GetUgcItemModerationStateRequest 
type GetUgcItemModerationStateRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId iD of the UGC item to get the moderation state for.
    ItemId string `json:"ItemId,omitempty"`
}

// GetUgcItemModerationStateResult 
type GetUgcItemModerationStateResultModel struct {
    // State the current moderation state for the requested item.
    State *ModerationStateModel `json:"State,omitempty"`
}

// GetVirtualCurrenciesRequest given an entity type, entity identifier and container details, will get the entity's virtual currencies.
type GetVirtualCurrenciesRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// GetVirtualCurrenciesResult 
type GetVirtualCurrenciesResultModel struct {
    // Currencies 
    Currencies []CurrencyResponseDtoModel `json:"Currencies,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// GrantInventoryItemsRequest given an entity type, entity identifier and container details, will grant the specified inventory items.
type GrantInventoryItemsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items a list of Items to modify
    Items []InventoryItemDetailsModel `json:"Items,omitempty"`
}

// GrantInventoryItemsResult 
type GrantInventoryItemsResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items 
    Items []InventoryItemResponseDtoModel `json:"Items,omitempty"`
}

// HelpfulnessVoteRequest 
type HelpfulnessVoteRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IsHelpful 
    IsHelpful bool `json:"IsHelpful,omitempty"`
    // ReviewId the ID of the review to submit a helpfulness vote for.
    ReviewId string `json:"ReviewId,omitempty"`
}

// HelpfulnessVoteResult 
type HelpfulnessVoteResultModel struct {
}

// Image 
type ImageModel struct {
    // Id 
    Id string `json:"Id,omitempty"`
    // Locales the optional list of locales that are supported by this image.
    Locales []string `json:"Locales,omitempty"`
    // Tag the client-defined tag associated with this image.
    Tag string `json:"Tag,omitempty"`
    // Type the client-defined type of this image.
    Type string `json:"Type,omitempty"`
    // Url the URL for retrieval of the image.
    Url string `json:"Url,omitempty"`
}

// InventoryItemDetails 
type InventoryItemDetailsModel struct {
    // AlternateId 
    AlternateId *CatalogAlternateIdModel `json:"AlternateId,omitempty"`
    // Duration 
    Duration string `json:"Duration,omitempty"`
    // ExpirationDate 
    ExpirationDate time.Time `json:"ExpirationDate,omitempty"`
    // InstanceId 
    InstanceId string `json:"InstanceId,omitempty"`
    // IsSubscription 
    IsSubscription bool `json:"IsSubscription,omitempty"`
    // ItemId 
    ItemId string `json:"ItemId,omitempty"`
    // Marketplace 
    Marketplace string `json:"Marketplace,omitempty"`
    // MergeProperties 
    MergeProperties bool `json:"MergeProperties,omitempty"`
    // Origin 
    Origin string `json:"Origin,omitempty"`
    // OriginId 
    OriginId string `json:"OriginId,omitempty"`
    // Properties 
    Properties map[string]string `json:"Properties,omitempty"`
    // Quantity 
    Quantity int32 `json:"Quantity,omitempty"`
}

// InventoryItemResponseDto 
type InventoryItemResponseDtoModel struct {
    // Amount 
    Amount int32 `json:"Amount,omitempty"`
    // ChangedAmount 
    ChangedAmount int32 `json:"ChangedAmount,omitempty"`
    // InstanceId 
    InstanceId string `json:"InstanceId,omitempty"`
    // ItemId 
    ItemId string `json:"ItemId,omitempty"`
    // Properties 
    Properties map[string]string `json:"Properties,omitempty"`
    // Receipt 
    Receipt string `json:"Receipt,omitempty"`
}

// ModerationState 
type ModerationStateModel struct {
    // LastModifiedDate the date and time this moderation state was last updated.
    LastModifiedDate time.Time `json:"LastModifiedDate,omitempty"`
    // Reason the current stated reason for the associated item being moderated.
    Reason string `json:"Reason,omitempty"`
    // Status the current moderation status for the associated item.
    Status ModerationStatus `json:"Status,omitempty"`
}

// ModerationStatus 
type ModerationStatus string
  
const (
     ModerationStatusUnknown ModerationStatus = "Unknown"
     ModerationStatusAwaitingModeration ModerationStatus = "AwaitingModeration"
     ModerationStatusApproved ModerationStatus = "Approved"
     ModerationStatusRejected ModerationStatus = "Rejected"
      )
// PayoutInfo 
type PayoutInfoModel struct {
    // AccountSellerId the Dev Center account ID of the payee.
    AccountSellerId string `json:"AccountSellerId,omitempty"`
    // TaxCode the tax code for payout calculations.
    TaxCode string `json:"TaxCode,omitempty"`
    // Uaid the Universal account ID of the payee.
    Uaid string `json:"Uaid,omitempty"`
}

// PublishItemRequest the call kicks off a workflow to publish the item to the public catalog. The Publish Status API should be used to
// monitor the publish job.
type PublishItemRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ETag eTag of the catalog item to published from the working catalog to the public catalog. Used for optimistic concurrency.
// If the provided ETag does not match the ETag in the current working catalog, the request will be reject. If not
// provided, the current version of the document in the working catalog will be published.
    ETag string `json:"ETag,omitempty"`
    // ItemId iD of the catalog item to publish from the working catalog to the public catalog.
    ItemId string `json:"ItemId,omitempty"`
}

// PublishItemResult 
type PublishItemResultModel struct {
}

// PublishResult 
type PublishResult string
  
const (
     PublishResultUnknown PublishResult = "Unknown"
     PublishResultPending PublishResult = "Pending"
     PublishResultSucceeded PublishResult = "Succeeded"
     PublishResultFailed PublishResult = "Failed"
     PublishResultCanceled PublishResult = "Canceled"
      )
// PublishStatusRequest 
type PublishStatusRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId iD of the catalog item to publish from the working catalog to the public catalog.
    ItemId string `json:"ItemId,omitempty"`
}

// PublishStatusResult 
type PublishStatusResultModel struct {
    // Result high level state of the publish.
    Result PublishResult `json:"Result,omitempty"`
    // StatusMessage descriptive message about the current status of the publish.
    StatusMessage string `json:"StatusMessage,omitempty"`
}

// PurchaseItemByFriendlyIdRequest purchase item by friendly Id request
type PurchaseItemByFriendlyIdRequestModel struct {
    // AutoConsume sets a value indicating whether the purchased item should be automatically consumed. The specified item quantity will be
// both increased and decreased.
    AutoConsume bool `json:"AutoConsume,omitempty"`
    // Currencies the currencies used to pay for item.
    Currencies []PurchaseItemCurrencyModel `json:"Currencies,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // ItemFriendlyId friendly identifier of the item to purchase.
    ItemFriendlyId string `json:"ItemFriendlyId,omitempty"`
    // Properties additional properties to be associated with purchased item.
    Properties map[string]string `json:"Properties,omitempty"`
    // Quantity the quantity of the item to purchase.
    Quantity int32 `json:"Quantity,omitempty"`
    // ReturnInventory indicates if the full inventory should be returned.
    ReturnInventory bool `json:"ReturnInventory,omitempty"`
    // Store the store to buy the item through.
    Store *PurchaseStoreInfoModel `json:"Store,omitempty"`
    // Uaid the Universal account ID for the creator of the item.
    Uaid string `json:"Uaid,omitempty"`
}

// PurchaseItemByIdRequest purchase item by Id request
type PurchaseItemByIdRequestModel struct {
    // AutoConsume sets a value indicating whether the purchased item should be automatically consumed. The specified item quantity will be
// both increased and decreased.
    AutoConsume bool `json:"AutoConsume,omitempty"`
    // Currencies the currencies used to pay for item.
    Currencies []PurchaseItemCurrencyModel `json:"Currencies,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // ItemId unique identifier of the item to purchase.
    ItemId string `json:"ItemId,omitempty"`
    // Properties additional properties to be associated with purchased item.
    Properties map[string]string `json:"Properties,omitempty"`
    // Quantity the quantity of the item to purchase.
    Quantity int32 `json:"Quantity,omitempty"`
    // ReturnInventory indicates if the full inventory should be returned.
    ReturnInventory bool `json:"ReturnInventory,omitempty"`
    // Store the store to buy the item through.
    Store *PurchaseStoreInfoModel `json:"Store,omitempty"`
    // Uaid the Universal account ID for the creator of the item.
    Uaid string `json:"Uaid,omitempty"`
}

// PurchaseItemCurrency 
type PurchaseItemCurrencyModel struct {
    // CurrencyId the unique identifier of the currency.
    CurrencyId string `json:"CurrencyId,omitempty"`
    // ExpectedPrice price the client expects to pay for the item.
    ExpectedPrice int32 `json:"ExpectedPrice,omitempty"`
}

// PurchaseItemResult purchase item result
type PurchaseItemResultModel struct {
    // Currencies details of the currencies used to purchase the items.
    Currencies []PurchaseItemResultCurrencyModel `json:"Currencies,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // InventoryItems details of the current inventory items.
    InventoryItems []PurchaseItemResultItemModel `json:"InventoryItems,omitempty"`
    // InventoryTooLarge indicates whether the number of inventory items is too large to be returned.
    InventoryTooLarge bool `json:"InventoryTooLarge,omitempty"`
    // PurchasedItems details for the items purchased.
    PurchasedItems []PurchaseItemResultItemModel `json:"PurchasedItems,omitempty"`
}

// PurchaseItemResultCurrency 
type PurchaseItemResultCurrencyModel struct {
    // Amount the total amount of the currency on inventory.
    Amount int32 `json:"Amount,omitempty"`
    // ChangedAmount the amount of the currency used in this call.
    ChangedAmount int32 `json:"ChangedAmount,omitempty"`
    // CurrencyId the ID of the currency used in this purchase.
    CurrencyId string `json:"CurrencyId,omitempty"`
}

// PurchaseItemResultItem 
type PurchaseItemResultItemModel struct {
    // Amount the total amount of the item on inventory.
    Amount int32 `json:"Amount,omitempty"`
    // ChangedAmount the amount of the item purchased in this call.
    ChangedAmount int32 `json:"ChangedAmount,omitempty"`
    // FriendlyId friendly id of the item.
    FriendlyId string `json:"FriendlyId,omitempty"`
    // InstanceId unique item identifier for this specific instance of the item.
    InstanceId string `json:"InstanceId,omitempty"`
    // ItemId unique identifier of the item.
    ItemId string `json:"ItemId,omitempty"`
    // ItemType the type of item.
    ItemType string `json:"ItemType,omitempty"`
    // Properties the properties associated to the currency.
    Properties map[string]string `json:"Properties,omitempty"`
}

// PurchaseStoreInfo 
type PurchaseStoreInfoModel struct {
    // FriendlyId the friendly identifier of the store.
    FriendlyId string `json:"FriendlyId,omitempty"`
    // Id the unique identifier of the store.
    Id string `json:"Id,omitempty"`
}

// Rating 
type RatingModel struct {
    // Average the average rating for this item
    Average float32 `json:"Average,omitempty"`
    // Count1Star the total count of 1 star ratings for this item
    Count1Star int32 `json:"Count1Star,omitempty"`
    // Count2Star the total count of 2 star ratings for this item
    Count2Star int32 `json:"Count2Star,omitempty"`
    // Count3Star the total count of 3 star ratings for this item
    Count3Star int32 `json:"Count3Star,omitempty"`
    // Count4Star the total count of 4 star ratings for this item
    Count4Star int32 `json:"Count4Star,omitempty"`
    // Count5Star the total count of 5 star ratings for this item
    Count5Star int32 `json:"Count5Star,omitempty"`
    // TotalCount the total count of ratings for this item
    TotalCount int32 `json:"TotalCount,omitempty"`
}

// ReportItemRequest 
type ReportItemRequestModel struct {
    // Concern category of concern for this report.
    Concern ConcernCategory `json:"Concern,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId iD of the item to report.
    ItemId string `json:"ItemId,omitempty"`
    // Reason the string reason for this report.
    Reason string `json:"Reason,omitempty"`
}

// ReportItemResult 
type ReportItemResultModel struct {
}

// ReportReviewRequest submit a report for an inappropriate review, allowing the submitting user to specify their concern.
type ReportReviewRequestModel struct {
    // ConcernCategory the reason this review is being reported.
    ConcernCategory ConcernCategory `json:"ConcernCategory,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ReviewId the ID of the review to submit a report for.
    ReviewId string `json:"ReviewId,omitempty"`
}

// ReportReviewResult 
type ReportReviewResultModel struct {
}

// Review 
type ReviewModel struct {
    // HelpfulNegative the number of negative helpfulness votes for this review.
    HelpfulNegative int32 `json:"HelpfulNegative,omitempty"`
    // HelpfulnessVotes total number of helpfulness votes for this review.
    HelpfulnessVotes int32 `json:"HelpfulnessVotes,omitempty"`
    // HelpfulPositive the number of positive helpfulness votes for this review.
    HelpfulPositive int32 `json:"HelpfulPositive,omitempty"`
    // IsInstalled indicates whether the review author has the item installed.
    IsInstalled bool `json:"IsInstalled,omitempty"`
    // ItemId the ID of the item being reviewed.
    ItemId string `json:"ItemId,omitempty"`
    // ItemVersion the version of the item being reviewed.
    ItemVersion string `json:"ItemVersion,omitempty"`
    // Locale the locale for which this review was submitted in.
    Locale string `json:"Locale,omitempty"`
    // Rating star rating associated with this review.
    Rating int32 `json:"Rating,omitempty"`
    // ReviewerId the ID of the author of the review.
    ReviewerId string `json:"ReviewerId,omitempty"`
    // ReviewId the ID of the review.
    ReviewId string `json:"ReviewId,omitempty"`
    // ReviewText the full text of this review.
    ReviewText string `json:"ReviewText,omitempty"`
    // Submitted the date and time this review was last submitted.
    Submitted time.Time `json:"Submitted,omitempty"`
    // Title the title of this review.
    Title string `json:"Title,omitempty"`
}

// ReviewSubmission 
type ReviewSubmissionModel struct {
    // IsInstalled indicates whether the review author has the item installed.
    IsInstalled bool `json:"IsInstalled,omitempty"`
    // ItemVersion the version of the item being reviewed.
    ItemVersion string `json:"ItemVersion,omitempty"`
    // Rating star rating associated with this review.
    Rating int32 `json:"Rating,omitempty"`
    // ReviewText the full text of this review.
    ReviewText string `json:"ReviewText,omitempty"`
    // Title the title of this review.
    Title string `json:"Title,omitempty"`
}

// ReviewSummaryRequest 
type ReviewSummaryRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId the ID of the item to retrieve the reviews summary for.
    ItemId string `json:"ItemId,omitempty"`
}

// ReviewSummaryResult 
type ReviewSummaryResultModel struct {
    // LeastFavorableReview 
    LeastFavorableReview *ReviewModel `json:"LeastFavorableReview,omitempty"`
    // MostFavorableReview 
    MostFavorableReview *ReviewModel `json:"MostFavorableReview,omitempty"`
    // Rating the summary of ratings associated with this item.
    Rating *RatingModel `json:"Rating,omitempty"`
    // ReviewsCount the total number of reviews associated with this item.
    ReviewsCount int32 `json:"ReviewsCount,omitempty"`
}

// ReviewTakedown 
type ReviewTakedownModel struct {
    // ItemId the ID of the item associated with the review to take down.
    ItemId string `json:"ItemId,omitempty"`
    // ReviewId the ID of the review to take down.
    ReviewId string `json:"ReviewId,omitempty"`
}

// SearchBundlesRequest search for bundles request
type SearchBundlesRequestModel struct {
    // ExpandReferencedItems whether the referenced items should be expanded.
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // Filter an OData filter used to refine the search query.
    Filter string `json:"Filter,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // OrderBy an OData orderBy used to order the results of the search query.
    OrderBy string `json:"OrderBy,omitempty"`
    // Search the text to search for
    Search string `json:"Search,omitempty"`
    // Skip the number of search results to skip.
    Skip int32 `json:"Skip,omitempty"`
    // SourceEntityKey the title or namespace to search under.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
    // Top the number of search results to retrieve.
    Top int32 `json:"Top,omitempty"`
}

// SearchBundlesResult search for bundles result
type SearchBundlesResultModel struct {
    // Bundles 
    Bundles []GetBundleResultModel `json:"Bundles,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// SearchForItemsContainingItemByFriendlyIdRequest search for bundles/subscriptions/stores containing any item from a list of items. The items are referenced by Friendly
// Ids.
type SearchForItemsContainingItemByFriendlyIdRequestModel struct {
    // ContainerType do we want to restrict the Search to Bundles/Subscriptions/Stores?
    ContainerType ContainerType `json:"ContainerType,omitempty"`
    // FriendlyIds friendlyIds of the items
    FriendlyIds []string `json:"FriendlyIds,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // SourceEntityKey the title or namespace to search under.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// SearchForItemsContainingItemByIdRequest search for bundles/subscriptions/stores containing any item from a list of items. The items are referenced by Ids.
type SearchForItemsContainingItemByIdRequestModel struct {
    // ContainerType do we want to restrict the Search to Bundles/Subscriptions/Stores?
    ContainerType ContainerType `json:"ContainerType,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Ids ids of the items
    Ids []string `json:"Ids,omitempty"`
    // SourceEntityKey the title or namespace to search under.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
}

// SearchForItemsContainingItemResult search for items containing an item result
type SearchForItemsContainingItemResultModel struct {
    // Bundles list of Bundles containing the item.
    Bundles []CatalogItemMetadataModel `json:"Bundles,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Stores list of Stores containing the item.
    Stores []CatalogItemMetadataModel `json:"Stores,omitempty"`
    // Subscriptions list of Subscriptions containing the item.
    Subscriptions []CatalogItemMetadataModel `json:"Subscriptions,omitempty"`
}

// SearchInStoreByFriendlyIdRequest search for items within a store request. The store is referenced by Friendly Id.
type SearchInStoreByFriendlyIdRequestModel struct {
    // Filter an OData filter used to refine the search query.
    Filter string `json:"Filter,omitempty"`
    // FriendlyId friendlyId of target store
    FriendlyId string `json:"FriendlyId,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // OrderBy an OData orderBy used to order the results of the search query.
    OrderBy string `json:"OrderBy,omitempty"`
    // Search the text to search for
    Search string `json:"Search,omitempty"`
    // Skip the number of search results to skip.
    Skip int32 `json:"Skip,omitempty"`
    // SourceEntityKey the title or namespace to search under.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
    // Top the number of search results to retrieve.
    Top int32 `json:"Top,omitempty"`
}

// SearchInStoreByIdRequest search for items within a store request. The store is referenced by Id.
type SearchInStoreByIdRequestModel struct {
    // Filter an OData filter used to refine the search query.
    Filter string `json:"Filter,omitempty"`
    // Id id of target store
    Id string `json:"Id,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // OrderBy an OData orderBy used to order the results of the search query.
    OrderBy string `json:"OrderBy,omitempty"`
    // Search the text to search for
    Search string `json:"Search,omitempty"`
    // Skip the number of search results to skip.
    Skip int32 `json:"Skip,omitempty"`
    // SourceEntityKey the title or namespace to search under.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
    // Top the number of search results to retrieve.
    Top int32 `json:"Top,omitempty"`
}

// SearchInStoreResult search within a store result
type SearchInStoreResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items list of Items within the Store that satisfy the search condition.
    Items []CatalogItemMetadataModel `json:"Items,omitempty"`
}

// SearchStoresRequest search for stores request
type SearchStoresRequestModel struct {
    // ExpandReferencedItems whether the referenced items should be expanded.
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // Filter an OData filter used to refine the search query.
    Filter string `json:"Filter,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // OrderBy an OData orderBy used to order the results of the search query.
    OrderBy string `json:"OrderBy,omitempty"`
    // Search the text to search for
    Search string `json:"Search,omitempty"`
    // Skip the number of search results to skip.
    Skip int32 `json:"Skip,omitempty"`
    // SourceEntityKey the title or namespace to search under.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
    // Top the number of search results to retrieve.
    Top int32 `json:"Top,omitempty"`
}

// SearchStoresResult search for stores result
type SearchStoresResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Stores 
    Stores []GetStoreResultModel `json:"Stores,omitempty"`
}

// SearchSubscriptionsRequest search for subscriptions request
type SearchSubscriptionsRequestModel struct {
    // ExpandReferencedItems whether the referenced items should be expanded.
    ExpandReferencedItems bool `json:"ExpandReferencedItems,omitempty"`
    // Filter an OData filter used to refine the search query.
    Filter string `json:"Filter,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // OrderBy an OData orderBy used to order the results of the search query.
    OrderBy string `json:"OrderBy,omitempty"`
    // Search the text to search for
    Search string `json:"Search,omitempty"`
    // Skip the number of search results to skip.
    Skip int32 `json:"Skip,omitempty"`
    // SourceEntityKey the title or namespace to search under.
    SourceEntityKey *EntityKeyModel `json:"SourceEntityKey,omitempty"`
    // Top the number of search results to retrieve.
    Top int32 `json:"Top,omitempty"`
}

// SearchSubscriptionsResult search for subscriptions result
type SearchSubscriptionsResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Subscriptions 
    Subscriptions []GetSubscriptionResultModel `json:"Subscriptions,omitempty"`
}

// SetInventoryItemsRequest given an entity type, entity identifier and container details, will set the entity's inventory items
type SetInventoryItemsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items a list of Items to modify
    Items []InventoryItemDetailsModel `json:"Items,omitempty"`
}

// SetInventoryItemsResult 
type SetInventoryItemsResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items 
    Items []InventoryItemResponseDtoModel `json:"Items,omitempty"`
}

// SetUgcItemModerationStateRequest 
type SetUgcItemModerationStateRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ItemId iD of the UGC item to set the moderation state for.
    ItemId string `json:"ItemId,omitempty"`
    // Reason the reason for the moderation state change for the associated UGC item.
    Reason string `json:"Reason,omitempty"`
    // Status the status to set for the associated UGC item.
    Status ModerationStatus `json:"Status,omitempty"`
}

// SetUgcItemModerationStateResult 
type SetUgcItemModerationStateResultModel struct {
}

// SetVirtualCurrenciesRequest given an entity type, entity identifier and container details, will set the entity's currencies to a specific amount.
type SetVirtualCurrenciesRequestModel struct {
    // Currencies a list of currencies to modify
    Currencies []CurrencyDetailsModel `json:"Currencies,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// SetVirtualCurrenciesResult 
type SetVirtualCurrenciesResultModel struct {
    // Currencies 
    Currencies []CurrencyResponseDtoModel `json:"Currencies,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// SubscriptionDetails 
type SubscriptionDetailsModel struct {
    // DurationInSeconds the length of time that the subscription will last.
    DurationInSeconds float64 `json:"DurationInSeconds,omitempty"`
}

// SubscriptionItemResponseDto 
type SubscriptionItemResponseDtoModel struct {
    // ExpirationDate 
    ExpirationDate time.Time `json:"ExpirationDate,omitempty"`
    // ItemIds 
    ItemIds []string `json:"ItemIds,omitempty"`
    // Marketplace 
    Marketplace string `json:"Marketplace,omitempty"`
    // OfferId 
    OfferId string `json:"OfferId,omitempty"`
    // Receipt 
    Receipt string `json:"Receipt,omitempty"`
}

// SubtractVirtualCurrenciesRequest given an entity type, entity identifier and container details, will decrease the entity's currencies by a specific
// amount.
type SubtractVirtualCurrenciesRequestModel struct {
    // Currencies a list of currencies to modify
    Currencies []CurrencyDetailsModel `json:"Currencies,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// SubtractVirtualCurrenciesResult 
type SubtractVirtualCurrenciesResultModel struct {
    // Currencies 
    Currencies []CurrencyResponseDtoModel `json:"Currencies,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// TakedownReviewsRequest submit a request to takedown one or more reviews, removing them from public view. Authors will still be able to see
// their reviews after being taken down.
type TakedownReviewsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Reviews the set of reviews to take down.
    Reviews []ReviewTakedownModel `json:"Reviews,omitempty"`
}

// TakedownReviewsResult 
type TakedownReviewsResultModel struct {
}

// UpdateBundleRequest update bundle request
type UpdateBundleRequestModel struct {
    // Bundle bundle details
    Bundle CatalogItemMetadataModel `json:"Bundle,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// UpdateBundleResult update bundle result
type UpdateBundleResultModel struct {
    // Bundle 
    Bundle *CatalogItemMetadataModel `json:"Bundle,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
}

// UpdateCatalogConfigRequest 
type UpdateCatalogConfigRequestModel struct {
    // Config 
    Config *CatalogConfigModel `json:"Config,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// UpdateCatalogConfigResult 
type UpdateCatalogConfigResultModel struct {
}

// UpdateCurrencyRequest 
type UpdateCurrencyRequestModel struct {
    // Currency updated metadata describing the currency item to be updated.
    Currency CatalogItemMetadataModel `json:"Currency,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// UpdateCurrencyResult 
type UpdateCurrencyResultModel struct {
    // Currency updated metadata describing the currency item just updated.
    Currency *CatalogItemMetadataModel `json:"Currency,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// UpdateDraftItemRequest 
type UpdateDraftItemRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Item updated metadata describing the catalog item to be updated.
    Item CatalogItemMetadataModel `json:"Item,omitempty"`
}

// UpdateDraftItemResult 
type UpdateDraftItemResultModel struct {
    // Item updated metadata describing the catalog item just updated.
    Item *CatalogItemMetadataModel `json:"Item,omitempty"`
}

// UpdateInventoryItemsPropertiesRequest update the specified inventory items.
type UpdateInventoryItemsPropertiesRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items a list of Items to modify
    Items []InventoryItemDetailsModel `json:"Items,omitempty"`
}

// UpdateInventoryPropertiesItemsResult 
type UpdateInventoryPropertiesItemsResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Items 
    Items []InventoryItemResponseDtoModel `json:"Items,omitempty"`
    // Subscriptions 
    Subscriptions []SubscriptionItemResponseDtoModel `json:"Subscriptions,omitempty"`
}

// UpdateStoreRequest update store request
type UpdateStoreRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Store bundle details
    Store CatalogItemMetadataModel `json:"Store,omitempty"`
}

// UpdateStoreResult update store result
type UpdateStoreResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Store 
    Store *CatalogItemMetadataModel `json:"Store,omitempty"`
}

// UpdateSubscriptionRequest update subscription request
type UpdateSubscriptionRequestModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Subscription subscription details
    Subscription CatalogItemMetadataModel `json:"Subscription,omitempty"`
}

// UpdateSubscriptionResult update subscription result
type UpdateSubscriptionResultModel struct {
    // IdempotencyId the Idempotency Id for this request
    IdempotencyId string `json:"IdempotencyId,omitempty"`
    // Subscription 
    Subscription *CatalogItemMetadataModel `json:"Subscription,omitempty"`
}

// UploadInfo 
type UploadInfoModel struct {
    // FileName name of the file to be uploaded.
    FileName string `json:"FileName,omitempty"`
    // FileSize size of the file to be uploaded, in bytes.
    FileSize int32 `json:"FileSize,omitempty"`
}

// UploadUrlMetadata 
type UploadUrlMetadataModel struct {
    // FileName name of the file for which this upload url was requested.
    FileName string `json:"FileName,omitempty"`
    // Id unique identifier for the binary content to be uploaded to the target url.
    Id string `json:"Id,omitempty"`
    // Url url for the binary content to be uploaded to.
    Url string `json:"Url,omitempty"`
}

// UserGeneratedContentSpecificConfig 
type UserGeneratedContentSpecificConfigModel struct {
    // ContentTypes the set of content types that will be used for validation and if no values are provided then anything is allowed.
    ContentTypes []string `json:"ContentTypes,omitempty"`
    // Enabled flag defining whether user generated content is enabled.
    Enabled bool `json:"Enabled,omitempty"`
    // Tags the set of tags that will be used for validation and if no values are provided then anything is allowed.
    Tags []string `json:"Tags,omitempty"`
}
