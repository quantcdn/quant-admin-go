# CrawlerRequestUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**BrowserMode** | Pointer to **bool** |  | [optional] [default to false]
**Urls** | Pointer to **[]string** |  | [optional] [default to []]
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Exclude** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewCrawlerRequestUpdate

`func NewCrawlerRequestUpdate() *CrawlerRequestUpdate`

NewCrawlerRequestUpdate instantiates a new CrawlerRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlerRequestUpdateWithDefaults

`func NewCrawlerRequestUpdateWithDefaults() *CrawlerRequestUpdate`

NewCrawlerRequestUpdateWithDefaults instantiates a new CrawlerRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CrawlerRequestUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrawlerRequestUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrawlerRequestUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrawlerRequestUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *CrawlerRequestUpdate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CrawlerRequestUpdate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CrawlerRequestUpdate) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CrawlerRequestUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetBrowserMode

`func (o *CrawlerRequestUpdate) GetBrowserMode() bool`

GetBrowserMode returns the BrowserMode field if non-nil, zero value otherwise.

### GetBrowserModeOk

`func (o *CrawlerRequestUpdate) GetBrowserModeOk() (*bool, bool)`

GetBrowserModeOk returns a tuple with the BrowserMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserMode

`func (o *CrawlerRequestUpdate) SetBrowserMode(v bool)`

SetBrowserMode sets BrowserMode field to given value.

### HasBrowserMode

`func (o *CrawlerRequestUpdate) HasBrowserMode() bool`

HasBrowserMode returns a boolean if a field has been set.

### GetUrls

`func (o *CrawlerRequestUpdate) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CrawlerRequestUpdate) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CrawlerRequestUpdate) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *CrawlerRequestUpdate) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetHeaders

`func (o *CrawlerRequestUpdate) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CrawlerRequestUpdate) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CrawlerRequestUpdate) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CrawlerRequestUpdate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetExclude

`func (o *CrawlerRequestUpdate) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *CrawlerRequestUpdate) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *CrawlerRequestUpdate) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *CrawlerRequestUpdate) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


