# RunCrawlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | Pointer to **[]string** | Optionally provide manual URLs to crawl. Full crawl will run otherwise. | [optional] 

## Methods

### NewRunCrawlerRequest

`func NewRunCrawlerRequest() *RunCrawlerRequest`

NewRunCrawlerRequest instantiates a new RunCrawlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCrawlerRequestWithDefaults

`func NewRunCrawlerRequestWithDefaults() *RunCrawlerRequest`

NewRunCrawlerRequestWithDefaults instantiates a new RunCrawlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *RunCrawlerRequest) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *RunCrawlerRequest) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *RunCrawlerRequest) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *RunCrawlerRequest) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


