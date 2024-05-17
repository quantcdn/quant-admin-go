# CrawlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Domain** | **string** |  | 
**BrowserMode** | Pointer to **bool** |  | [optional] 
**UrlList** | **[]string** |  | 
**Headers** | **string** |  | 

## Methods

### NewCrawlerRequest

`func NewCrawlerRequest(domain string, urlList []string, headers string, ) *CrawlerRequest`

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

### GetUrlList

`func (o *CrawlerRequest) GetUrlList() []string`

GetUrlList returns the UrlList field if non-nil, zero value otherwise.

### GetUrlListOk

`func (o *CrawlerRequest) GetUrlListOk() (*[]string, bool)`

GetUrlListOk returns a tuple with the UrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlList

`func (o *CrawlerRequest) SetUrlList(v []string)`

SetUrlList sets UrlList field to given value.


### GetHeaders

`func (o *CrawlerRequest) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CrawlerRequest) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CrawlerRequest) SetHeaders(v string)`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


