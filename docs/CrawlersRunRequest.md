# CrawlersRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | Pointer to **[]string** | Optional URLs to crawl (overrides crawler&#39;s default URL configuration). If not provided, the crawler will use its configured URLs or perform a full crawl. | [optional] 

## Methods

### NewCrawlersRunRequest

`func NewCrawlersRunRequest() *CrawlersRunRequest`

NewCrawlersRunRequest instantiates a new CrawlersRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlersRunRequestWithDefaults

`func NewCrawlersRunRequestWithDefaults() *CrawlersRunRequest`

NewCrawlersRunRequestWithDefaults instantiates a new CrawlersRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *CrawlersRunRequest) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CrawlersRunRequest) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CrawlersRunRequest) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *CrawlersRunRequest) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


