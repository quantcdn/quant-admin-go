# V2CrawlerBrowserConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptureApiResponses** | Pointer to **bool** | Store XHR/fetch responses as files, so a static copy can serve a site whose navigation or content is rendered client-side from a JSON endpoint | [optional] 
**WaitForNetworkIdle** | Pointer to **int32** | Wait for the network to settle before capture, in milliseconds. Useful for API-driven sites | [optional] 
**UseRenderedHtml** | Pointer to **bool** | Store the JavaScript-modified DOM instead of the original HTML response | [optional] 

## Methods

### NewV2CrawlerBrowserConfig

`func NewV2CrawlerBrowserConfig() *V2CrawlerBrowserConfig`

NewV2CrawlerBrowserConfig instantiates a new V2CrawlerBrowserConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CrawlerBrowserConfigWithDefaults

`func NewV2CrawlerBrowserConfigWithDefaults() *V2CrawlerBrowserConfig`

NewV2CrawlerBrowserConfigWithDefaults instantiates a new V2CrawlerBrowserConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptureApiResponses

`func (o *V2CrawlerBrowserConfig) GetCaptureApiResponses() bool`

GetCaptureApiResponses returns the CaptureApiResponses field if non-nil, zero value otherwise.

### GetCaptureApiResponsesOk

`func (o *V2CrawlerBrowserConfig) GetCaptureApiResponsesOk() (*bool, bool)`

GetCaptureApiResponsesOk returns a tuple with the CaptureApiResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureApiResponses

`func (o *V2CrawlerBrowserConfig) SetCaptureApiResponses(v bool)`

SetCaptureApiResponses sets CaptureApiResponses field to given value.

### HasCaptureApiResponses

`func (o *V2CrawlerBrowserConfig) HasCaptureApiResponses() bool`

HasCaptureApiResponses returns a boolean if a field has been set.

### GetWaitForNetworkIdle

`func (o *V2CrawlerBrowserConfig) GetWaitForNetworkIdle() int32`

GetWaitForNetworkIdle returns the WaitForNetworkIdle field if non-nil, zero value otherwise.

### GetWaitForNetworkIdleOk

`func (o *V2CrawlerBrowserConfig) GetWaitForNetworkIdleOk() (*int32, bool)`

GetWaitForNetworkIdleOk returns a tuple with the WaitForNetworkIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForNetworkIdle

`func (o *V2CrawlerBrowserConfig) SetWaitForNetworkIdle(v int32)`

SetWaitForNetworkIdle sets WaitForNetworkIdle field to given value.

### HasWaitForNetworkIdle

`func (o *V2CrawlerBrowserConfig) HasWaitForNetworkIdle() bool`

HasWaitForNetworkIdle returns a boolean if a field has been set.

### GetUseRenderedHtml

`func (o *V2CrawlerBrowserConfig) GetUseRenderedHtml() bool`

GetUseRenderedHtml returns the UseRenderedHtml field if non-nil, zero value otherwise.

### GetUseRenderedHtmlOk

`func (o *V2CrawlerBrowserConfig) GetUseRenderedHtmlOk() (*bool, bool)`

GetUseRenderedHtmlOk returns a tuple with the UseRenderedHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRenderedHtml

`func (o *V2CrawlerBrowserConfig) SetUseRenderedHtml(v bool)`

SetUseRenderedHtml sets UseRenderedHtml field to given value.

### HasUseRenderedHtml

`func (o *V2CrawlerBrowserConfig) HasUseRenderedHtml() bool`

HasUseRenderedHtml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


