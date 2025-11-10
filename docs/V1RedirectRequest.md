# V1RedirectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The redirect from URL | 
**RedirectUrl** | **string** | The destination URL | 
**RedirectHttpCode** | **int32** | The HTTP code to send with the redirect | 
**Published** | **bool** | If the redirect is published | 
**ContentTimestamp** | Pointer to **int32** | User defined timestamp for this content item | [optional] 
**Info** | Pointer to [**V1Info**](V1Info.md) |  | [optional] 
**Transitions** | Pointer to [**[]V1Transition**](V1Transition.md) |  | [optional] 

## Methods

### NewV1RedirectRequest

`func NewV1RedirectRequest(url string, redirectUrl string, redirectHttpCode int32, published bool, ) *V1RedirectRequest`

NewV1RedirectRequest instantiates a new V1RedirectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RedirectRequestWithDefaults

`func NewV1RedirectRequestWithDefaults() *V1RedirectRequest`

NewV1RedirectRequestWithDefaults instantiates a new V1RedirectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1RedirectRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1RedirectRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1RedirectRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRedirectUrl

`func (o *V1RedirectRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *V1RedirectRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *V1RedirectRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetRedirectHttpCode

`func (o *V1RedirectRequest) GetRedirectHttpCode() int32`

GetRedirectHttpCode returns the RedirectHttpCode field if non-nil, zero value otherwise.

### GetRedirectHttpCodeOk

`func (o *V1RedirectRequest) GetRedirectHttpCodeOk() (*int32, bool)`

GetRedirectHttpCodeOk returns a tuple with the RedirectHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHttpCode

`func (o *V1RedirectRequest) SetRedirectHttpCode(v int32)`

SetRedirectHttpCode sets RedirectHttpCode field to given value.


### GetPublished

`func (o *V1RedirectRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *V1RedirectRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *V1RedirectRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetContentTimestamp

`func (o *V1RedirectRequest) GetContentTimestamp() int32`

GetContentTimestamp returns the ContentTimestamp field if non-nil, zero value otherwise.

### GetContentTimestampOk

`func (o *V1RedirectRequest) GetContentTimestampOk() (*int32, bool)`

GetContentTimestampOk returns a tuple with the ContentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTimestamp

`func (o *V1RedirectRequest) SetContentTimestamp(v int32)`

SetContentTimestamp sets ContentTimestamp field to given value.

### HasContentTimestamp

`func (o *V1RedirectRequest) HasContentTimestamp() bool`

HasContentTimestamp returns a boolean if a field has been set.

### GetInfo

`func (o *V1RedirectRequest) GetInfo() V1Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1RedirectRequest) GetInfoOk() (*V1Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1RedirectRequest) SetInfo(v V1Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1RedirectRequest) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetTransitions

`func (o *V1RedirectRequest) GetTransitions() []V1Transition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *V1RedirectRequest) GetTransitionsOk() (*[]V1Transition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *V1RedirectRequest) SetTransitions(v []V1Transition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *V1RedirectRequest) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


