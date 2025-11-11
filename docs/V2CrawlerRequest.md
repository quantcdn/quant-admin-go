# V2CrawlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Crawler name | [optional] 
**Domain** | **string** | Domain to crawl | 
**BrowserMode** | Pointer to **bool** | Enable browser mode | [optional] [default to false]
**Urls** | Pointer to **[]string** | URLs to crawl | [optional] 
**StartUrls** | Pointer to **[]string** | Starting URLs for crawl | [optional] 
**Headers** | Pointer to **map[string]string** | Custom headers | [optional] 
**Exclude** | Pointer to **[]string** | URL patterns to exclude (regex) | [optional] 
**Include** | Pointer to **[]string** | URL patterns to include (regex) | [optional] 
**WebhookUrl** | Pointer to **string** | Webhook URL for notifications | [optional] 
**WebhookAuthHeader** | Pointer to **string** | Authorization header for webhook | [optional] 
**WebhookExtraVars** | Pointer to **string** | Extra variables for webhook | [optional] 
**Workers** | Pointer to **int32** | Number of concurrent workers (default: 2, non-default requires verification) | [optional] 
**Delay** | Pointer to **float32** | Delay between requests in seconds (default: 4, non-default requires verification) | [optional] 
**Depth** | Pointer to **int32** | Maximum crawl depth, -1 for unlimited | [optional] 
**MaxHits** | Pointer to **int32** | Maximum total requests, 0 for unlimited (default: 0, non-default requires verification) | [optional] 
**MaxHtml** | Pointer to **int32** | Maximum HTML pages, 0 for unlimited (default: org limit, non-default requires verification) | [optional] 
**StatusOk** | Pointer to **[]int32** | HTTP status codes that will result in content being captured and pushed to Quant | [optional] 
**Sitemap** | Pointer to [**[]V2CrawlerSitemapInner**](V2CrawlerSitemapInner.md) | Sitemap configuration | [optional] 
**AllowedDomains** | Pointer to **[]string** | Allowed domains for multi-domain crawling, automatically enables merge_domains | [optional] 
**UserAgent** | Pointer to **string** | Custom user agent, only when browser_mode is false | [optional] 
**Assets** | Pointer to [**V2CrawlerAssets**](V2CrawlerAssets.md) |  | [optional] 
**MaxErrors** | Pointer to **int32** | Maximum errors before stopping crawl | [optional] 

## Methods

### NewV2CrawlerRequest

`func NewV2CrawlerRequest(domain string, ) *V2CrawlerRequest`

NewV2CrawlerRequest instantiates a new V2CrawlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CrawlerRequestWithDefaults

`func NewV2CrawlerRequestWithDefaults() *V2CrawlerRequest`

NewV2CrawlerRequestWithDefaults instantiates a new V2CrawlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2CrawlerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2CrawlerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2CrawlerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2CrawlerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *V2CrawlerRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2CrawlerRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2CrawlerRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetBrowserMode

`func (o *V2CrawlerRequest) GetBrowserMode() bool`

GetBrowserMode returns the BrowserMode field if non-nil, zero value otherwise.

### GetBrowserModeOk

`func (o *V2CrawlerRequest) GetBrowserModeOk() (*bool, bool)`

GetBrowserModeOk returns a tuple with the BrowserMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMode

`func (o *V2CrawlerRequest) SetBrowserMode(v bool)`

SetBrowserMode sets BrowserMode field to given value.

### HasBrowserMode

`func (o *V2CrawlerRequest) HasBrowserMode() bool`

HasBrowserMode returns a boolean if a field has been set.

### GetUrls

`func (o *V2CrawlerRequest) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *V2CrawlerRequest) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *V2CrawlerRequest) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *V2CrawlerRequest) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetStartUrls

`func (o *V2CrawlerRequest) GetStartUrls() []string`

GetStartUrls returns the StartUrls field if non-nil, zero value otherwise.

### GetStartUrlsOk

`func (o *V2CrawlerRequest) GetStartUrlsOk() (*[]string, bool)`

GetStartUrlsOk returns a tuple with the StartUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrls

`func (o *V2CrawlerRequest) SetStartUrls(v []string)`

SetStartUrls sets StartUrls field to given value.

### HasStartUrls

`func (o *V2CrawlerRequest) HasStartUrls() bool`

HasStartUrls returns a boolean if a field has been set.

### GetHeaders

`func (o *V2CrawlerRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *V2CrawlerRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *V2CrawlerRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *V2CrawlerRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetExclude

`func (o *V2CrawlerRequest) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *V2CrawlerRequest) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *V2CrawlerRequest) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *V2CrawlerRequest) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *V2CrawlerRequest) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *V2CrawlerRequest) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *V2CrawlerRequest) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *V2CrawlerRequest) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *V2CrawlerRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *V2CrawlerRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *V2CrawlerRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *V2CrawlerRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookAuthHeader

`func (o *V2CrawlerRequest) GetWebhookAuthHeader() string`

GetWebhookAuthHeader returns the WebhookAuthHeader field if non-nil, zero value otherwise.

### GetWebhookAuthHeaderOk

`func (o *V2CrawlerRequest) GetWebhookAuthHeaderOk() (*string, bool)`

GetWebhookAuthHeaderOk returns a tuple with the WebhookAuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookAuthHeader

`func (o *V2CrawlerRequest) SetWebhookAuthHeader(v string)`

SetWebhookAuthHeader sets WebhookAuthHeader field to given value.

### HasWebhookAuthHeader

`func (o *V2CrawlerRequest) HasWebhookAuthHeader() bool`

HasWebhookAuthHeader returns a boolean if a field has been set.

### GetWebhookExtraVars

`func (o *V2CrawlerRequest) GetWebhookExtraVars() string`

GetWebhookExtraVars returns the WebhookExtraVars field if non-nil, zero value otherwise.

### GetWebhookExtraVarsOk

`func (o *V2CrawlerRequest) GetWebhookExtraVarsOk() (*string, bool)`

GetWebhookExtraVarsOk returns a tuple with the WebhookExtraVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookExtraVars

`func (o *V2CrawlerRequest) SetWebhookExtraVars(v string)`

SetWebhookExtraVars sets WebhookExtraVars field to given value.

### HasWebhookExtraVars

`func (o *V2CrawlerRequest) HasWebhookExtraVars() bool`

HasWebhookExtraVars returns a boolean if a field has been set.

### GetWorkers

`func (o *V2CrawlerRequest) GetWorkers() int32`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *V2CrawlerRequest) GetWorkersOk() (*int32, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *V2CrawlerRequest) SetWorkers(v int32)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *V2CrawlerRequest) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetDelay

`func (o *V2CrawlerRequest) GetDelay() float32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *V2CrawlerRequest) GetDelayOk() (*float32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *V2CrawlerRequest) SetDelay(v float32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *V2CrawlerRequest) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDepth

`func (o *V2CrawlerRequest) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *V2CrawlerRequest) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *V2CrawlerRequest) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *V2CrawlerRequest) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetMaxHits

`func (o *V2CrawlerRequest) GetMaxHits() int32`

GetMaxHits returns the MaxHits field if non-nil, zero value otherwise.

### GetMaxHitsOk

`func (o *V2CrawlerRequest) GetMaxHitsOk() (*int32, bool)`

GetMaxHitsOk returns a tuple with the MaxHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHits

`func (o *V2CrawlerRequest) SetMaxHits(v int32)`

SetMaxHits sets MaxHits field to given value.

### HasMaxHits

`func (o *V2CrawlerRequest) HasMaxHits() bool`

HasMaxHits returns a boolean if a field has been set.

### GetMaxHtml

`func (o *V2CrawlerRequest) GetMaxHtml() int32`

GetMaxHtml returns the MaxHtml field if non-nil, zero value otherwise.

### GetMaxHtmlOk

`func (o *V2CrawlerRequest) GetMaxHtmlOk() (*int32, bool)`

GetMaxHtmlOk returns a tuple with the MaxHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHtml

`func (o *V2CrawlerRequest) SetMaxHtml(v int32)`

SetMaxHtml sets MaxHtml field to given value.

### HasMaxHtml

`func (o *V2CrawlerRequest) HasMaxHtml() bool`

HasMaxHtml returns a boolean if a field has been set.

### GetStatusOk

`func (o *V2CrawlerRequest) GetStatusOk() []int32`

GetStatusOk returns the StatusOk field if non-nil, zero value otherwise.

### GetStatusOkOk

`func (o *V2CrawlerRequest) GetStatusOkOk() (*[]int32, bool)`

GetStatusOkOk returns a tuple with the StatusOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOk

`func (o *V2CrawlerRequest) SetStatusOk(v []int32)`

SetStatusOk sets StatusOk field to given value.

### HasStatusOk

`func (o *V2CrawlerRequest) HasStatusOk() bool`

HasStatusOk returns a boolean if a field has been set.

### GetSitemap

`func (o *V2CrawlerRequest) GetSitemap() []V2CrawlerSitemapInner`

GetSitemap returns the Sitemap field if non-nil, zero value otherwise.

### GetSitemapOk

`func (o *V2CrawlerRequest) GetSitemapOk() (*[]V2CrawlerSitemapInner, bool)`

GetSitemapOk returns a tuple with the Sitemap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitemap

`func (o *V2CrawlerRequest) SetSitemap(v []V2CrawlerSitemapInner)`

SetSitemap sets Sitemap field to given value.

### HasSitemap

`func (o *V2CrawlerRequest) HasSitemap() bool`

HasSitemap returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *V2CrawlerRequest) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *V2CrawlerRequest) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *V2CrawlerRequest) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *V2CrawlerRequest) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetUserAgent

`func (o *V2CrawlerRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *V2CrawlerRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *V2CrawlerRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *V2CrawlerRequest) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetAssets

`func (o *V2CrawlerRequest) GetAssets() V2CrawlerAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *V2CrawlerRequest) GetAssetsOk() (*V2CrawlerAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *V2CrawlerRequest) SetAssets(v V2CrawlerAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *V2CrawlerRequest) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetMaxErrors

`func (o *V2CrawlerRequest) GetMaxErrors() int32`

GetMaxErrors returns the MaxErrors field if non-nil, zero value otherwise.

### GetMaxErrorsOk

`func (o *V2CrawlerRequest) GetMaxErrorsOk() (*int32, bool)`

GetMaxErrorsOk returns a tuple with the MaxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxErrors

`func (o *V2CrawlerRequest) SetMaxErrors(v int32)`

SetMaxErrors sets MaxErrors field to given value.

### HasMaxErrors

`func (o *V2CrawlerRequest) HasMaxErrors() bool`

HasMaxErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


