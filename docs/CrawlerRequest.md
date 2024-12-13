# CrawlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Domain** | **string** |  | 
**BrowserMode** | Pointer to **bool** |  | [optional] [default to false]
**Urls** | Pointer to **[]string** |  | [optional] [default to []]
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Exclude** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewCrawlerRequest

`func NewCrawlerRequest(domain string, ) *CrawlerRequest`

NewCrawlerRequest instantiates a new CrawlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlerRequestWithDefaults

`func NewCrawlerRequestWithDefaults() *CrawlerRequest`

NewCrawlerRequestWithDefaults instantiates a new CrawlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CrawlerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrawlerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrawlerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrawlerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *CrawlerRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CrawlerRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CrawlerRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetBrowserMode

`func (o *CrawlerRequest) GetBrowserMode() bool`

GetBrowserMode returns the BrowserMode field if non-nil, zero value otherwise.

### GetBrowserModeOk

`func (o *CrawlerRequest) GetBrowserModeOk() (*bool, bool)`

GetBrowserModeOk returns a tuple with the BrowserMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMode

`func (o *CrawlerRequest) SetBrowserMode(v bool)`

SetBrowserMode sets BrowserMode field to given value.

### HasBrowserMode

`func (o *CrawlerRequest) HasBrowserMode() bool`

HasBrowserMode returns a boolean if a field has been set.

### GetUrls

`func (o *CrawlerRequest) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CrawlerRequest) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CrawlerRequest) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *CrawlerRequest) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetHeaders

`func (o *CrawlerRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CrawlerRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CrawlerRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CrawlerRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetExclude

`func (o *CrawlerRequest) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *CrawlerRequest) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *CrawlerRequest) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *CrawlerRequest) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


