# V2Crawler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Crawler ID | 
**Name** | Pointer to **string** | Crawler name | [optional] 
**ProjectId** | **int32** | Project ID | 
**Uuid** | **string** | Crawler UUID | 
**Config** | **string** | Crawler configuration (YAML) | 
**Domain** | **string** | Crawler domain | 
**DomainVerified** | Pointer to **int32** | Domain verification status | [optional] [default to 0]
**UrlsList** | Pointer to **string** | URLs list (YAML) | [optional] 
**WebhookUrl** | Pointer to **string** | Webhook URL for notifications | [optional] 
**WebhookAuthHeader** | Pointer to **string** | Authorization header for webhook | [optional] 
**WebhookExtraVars** | Pointer to **string** | Extra variables for webhook | [optional] 
**BrowserMode** | Pointer to **bool** | Browser mode enabled | [optional] 
**Workers** | Pointer to **int32** | Number of concurrent workers | [optional] 
**Delay** | Pointer to **float32** | Delay between requests in seconds | [optional] 
**Depth** | Pointer to **int32** | Maximum crawl depth | [optional] 
**MaxHits** | Pointer to **int32** | Maximum total requests | [optional] 
**MaxHtml** | Pointer to **int32** | Maximum HTML pages | [optional] 
**StatusOk** | Pointer to **[]int32** | HTTP status codes for content capture | [optional] 
**UserAgent** | Pointer to **string** | Custom user agent | [optional] 
**MaxErrors** | Pointer to **int32** | Maximum errors before stopping | [optional] 
**StartUrls** | Pointer to **[]string** | Starting URLs | [optional] 
**Urls** | Pointer to **[]string** | URLs list | [optional] 
**Headers** | Pointer to **map[string]string** | Custom headers | [optional] 
**Exclude** | Pointer to **[]string** | URL patterns to exclude | [optional] 
**Include** | Pointer to **[]string** | URL patterns to include | [optional] 
**Sitemap** | Pointer to [**[]V2CrawlerSitemapInner**](V2CrawlerSitemapInner.md) | Sitemap configuration | [optional] 
**AllowedDomains** | Pointer to **[]string** | Allowed domains | [optional] 
**Assets** | Pointer to [**V2CrawlerAssets**](V2CrawlerAssets.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update timestamp | [optional] 
**DeletedAt** | Pointer to **NullableTime** | Deletion timestamp | [optional] 

## Methods

### NewV2Crawler

`func NewV2Crawler(id int32, projectId int32, uuid string, config string, domain string, ) *V2Crawler`

NewV2Crawler instantiates a new V2Crawler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CrawlerWithDefaults

`func NewV2CrawlerWithDefaults() *V2Crawler`

NewV2CrawlerWithDefaults instantiates a new V2Crawler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2Crawler) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2Crawler) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2Crawler) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *V2Crawler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2Crawler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2Crawler) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2Crawler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *V2Crawler) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *V2Crawler) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *V2Crawler) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetUuid

`func (o *V2Crawler) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2Crawler) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2Crawler) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetConfig

`func (o *V2Crawler) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V2Crawler) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V2Crawler) SetConfig(v string)`

SetConfig sets Config field to given value.


### GetDomain

`func (o *V2Crawler) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2Crawler) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2Crawler) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDomainVerified

`func (o *V2Crawler) GetDomainVerified() int32`

GetDomainVerified returns the DomainVerified field if non-nil, zero value otherwise.

### GetDomainVerifiedOk

`func (o *V2Crawler) GetDomainVerifiedOk() (*int32, bool)`

GetDomainVerifiedOk returns a tuple with the DomainVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainVerified

`func (o *V2Crawler) SetDomainVerified(v int32)`

SetDomainVerified sets DomainVerified field to given value.

### HasDomainVerified

`func (o *V2Crawler) HasDomainVerified() bool`

HasDomainVerified returns a boolean if a field has been set.

### GetUrlsList

`func (o *V2Crawler) GetUrlsList() string`

GetUrlsList returns the UrlsList field if non-nil, zero value otherwise.

### GetUrlsListOk

`func (o *V2Crawler) GetUrlsListOk() (*string, bool)`

GetUrlsListOk returns a tuple with the UrlsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlsList

`func (o *V2Crawler) SetUrlsList(v string)`

SetUrlsList sets UrlsList field to given value.

### HasUrlsList

`func (o *V2Crawler) HasUrlsList() bool`

HasUrlsList returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *V2Crawler) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *V2Crawler) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *V2Crawler) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *V2Crawler) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookAuthHeader

`func (o *V2Crawler) GetWebhookAuthHeader() string`

GetWebhookAuthHeader returns the WebhookAuthHeader field if non-nil, zero value otherwise.

### GetWebhookAuthHeaderOk

`func (o *V2Crawler) GetWebhookAuthHeaderOk() (*string, bool)`

GetWebhookAuthHeaderOk returns a tuple with the WebhookAuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookAuthHeader

`func (o *V2Crawler) SetWebhookAuthHeader(v string)`

SetWebhookAuthHeader sets WebhookAuthHeader field to given value.

### HasWebhookAuthHeader

`func (o *V2Crawler) HasWebhookAuthHeader() bool`

HasWebhookAuthHeader returns a boolean if a field has been set.

### GetWebhookExtraVars

`func (o *V2Crawler) GetWebhookExtraVars() string`

GetWebhookExtraVars returns the WebhookExtraVars field if non-nil, zero value otherwise.

### GetWebhookExtraVarsOk

`func (o *V2Crawler) GetWebhookExtraVarsOk() (*string, bool)`

GetWebhookExtraVarsOk returns a tuple with the WebhookExtraVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookExtraVars

`func (o *V2Crawler) SetWebhookExtraVars(v string)`

SetWebhookExtraVars sets WebhookExtraVars field to given value.

### HasWebhookExtraVars

`func (o *V2Crawler) HasWebhookExtraVars() bool`

HasWebhookExtraVars returns a boolean if a field has been set.

### GetBrowserMode

`func (o *V2Crawler) GetBrowserMode() bool`

GetBrowserMode returns the BrowserMode field if non-nil, zero value otherwise.

### GetBrowserModeOk

`func (o *V2Crawler) GetBrowserModeOk() (*bool, bool)`

GetBrowserModeOk returns a tuple with the BrowserMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMode

`func (o *V2Crawler) SetBrowserMode(v bool)`

SetBrowserMode sets BrowserMode field to given value.

### HasBrowserMode

`func (o *V2Crawler) HasBrowserMode() bool`

HasBrowserMode returns a boolean if a field has been set.

### GetWorkers

`func (o *V2Crawler) GetWorkers() int32`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *V2Crawler) GetWorkersOk() (*int32, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *V2Crawler) SetWorkers(v int32)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *V2Crawler) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetDelay

`func (o *V2Crawler) GetDelay() float32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *V2Crawler) GetDelayOk() (*float32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *V2Crawler) SetDelay(v float32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *V2Crawler) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDepth

`func (o *V2Crawler) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *V2Crawler) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *V2Crawler) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *V2Crawler) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetMaxHits

`func (o *V2Crawler) GetMaxHits() int32`

GetMaxHits returns the MaxHits field if non-nil, zero value otherwise.

### GetMaxHitsOk

`func (o *V2Crawler) GetMaxHitsOk() (*int32, bool)`

GetMaxHitsOk returns a tuple with the MaxHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHits

`func (o *V2Crawler) SetMaxHits(v int32)`

SetMaxHits sets MaxHits field to given value.

### HasMaxHits

`func (o *V2Crawler) HasMaxHits() bool`

HasMaxHits returns a boolean if a field has been set.

### GetMaxHtml

`func (o *V2Crawler) GetMaxHtml() int32`

GetMaxHtml returns the MaxHtml field if non-nil, zero value otherwise.

### GetMaxHtmlOk

`func (o *V2Crawler) GetMaxHtmlOk() (*int32, bool)`

GetMaxHtmlOk returns a tuple with the MaxHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHtml

`func (o *V2Crawler) SetMaxHtml(v int32)`

SetMaxHtml sets MaxHtml field to given value.

### HasMaxHtml

`func (o *V2Crawler) HasMaxHtml() bool`

HasMaxHtml returns a boolean if a field has been set.

### GetStatusOk

`func (o *V2Crawler) GetStatusOk() []int32`

GetStatusOk returns the StatusOk field if non-nil, zero value otherwise.

### GetStatusOkOk

`func (o *V2Crawler) GetStatusOkOk() (*[]int32, bool)`

GetStatusOkOk returns a tuple with the StatusOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOk

`func (o *V2Crawler) SetStatusOk(v []int32)`

SetStatusOk sets StatusOk field to given value.

### HasStatusOk

`func (o *V2Crawler) HasStatusOk() bool`

HasStatusOk returns a boolean if a field has been set.

### GetUserAgent

`func (o *V2Crawler) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *V2Crawler) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *V2Crawler) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *V2Crawler) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetMaxErrors

`func (o *V2Crawler) GetMaxErrors() int32`

GetMaxErrors returns the MaxErrors field if non-nil, zero value otherwise.

### GetMaxErrorsOk

`func (o *V2Crawler) GetMaxErrorsOk() (*int32, bool)`

GetMaxErrorsOk returns a tuple with the MaxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxErrors

`func (o *V2Crawler) SetMaxErrors(v int32)`

SetMaxErrors sets MaxErrors field to given value.

### HasMaxErrors

`func (o *V2Crawler) HasMaxErrors() bool`

HasMaxErrors returns a boolean if a field has been set.

### GetStartUrls

`func (o *V2Crawler) GetStartUrls() []string`

GetStartUrls returns the StartUrls field if non-nil, zero value otherwise.

### GetStartUrlsOk

`func (o *V2Crawler) GetStartUrlsOk() (*[]string, bool)`

GetStartUrlsOk returns a tuple with the StartUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrls

`func (o *V2Crawler) SetStartUrls(v []string)`

SetStartUrls sets StartUrls field to given value.

### HasStartUrls

`func (o *V2Crawler) HasStartUrls() bool`

HasStartUrls returns a boolean if a field has been set.

### GetUrls

`func (o *V2Crawler) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *V2Crawler) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *V2Crawler) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *V2Crawler) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetHeaders

`func (o *V2Crawler) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *V2Crawler) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *V2Crawler) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *V2Crawler) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetExclude

`func (o *V2Crawler) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *V2Crawler) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *V2Crawler) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *V2Crawler) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *V2Crawler) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *V2Crawler) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *V2Crawler) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *V2Crawler) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetSitemap

`func (o *V2Crawler) GetSitemap() []V2CrawlerSitemapInner`

GetSitemap returns the Sitemap field if non-nil, zero value otherwise.

### GetSitemapOk

`func (o *V2Crawler) GetSitemapOk() (*[]V2CrawlerSitemapInner, bool)`

GetSitemapOk returns a tuple with the Sitemap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitemap

`func (o *V2Crawler) SetSitemap(v []V2CrawlerSitemapInner)`

SetSitemap sets Sitemap field to given value.

### HasSitemap

`func (o *V2Crawler) HasSitemap() bool`

HasSitemap returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *V2Crawler) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *V2Crawler) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *V2Crawler) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *V2Crawler) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAssets

`func (o *V2Crawler) GetAssets() V2CrawlerAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *V2Crawler) GetAssetsOk() (*V2CrawlerAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *V2Crawler) SetAssets(v V2CrawlerAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *V2Crawler) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2Crawler) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2Crawler) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2Crawler) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2Crawler) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2Crawler) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2Crawler) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2Crawler) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2Crawler) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *V2Crawler) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *V2Crawler) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *V2Crawler) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *V2Crawler) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *V2Crawler) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *V2Crawler) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


