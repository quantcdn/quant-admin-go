# V1ContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The content URL. Must be relative and start with a leading &#39;/&#39; | 
**Content** | **string** | The content (e.g. html) | 
**Published** | **bool** | If the asset should be published | 
**ContentTimestamp** | Pointer to **int32** | User defined timestamp for this content item | [optional] 
**Info** | Pointer to [**V1Info**](V1Info.md) |  | [optional] 
**Transitions** | Pointer to [**[]V1Transition**](V1Transition.md) |  | [optional] 

## Methods

### NewV1ContentRequest

`func NewV1ContentRequest(url string, content string, published bool, ) *V1ContentRequest`

NewV1ContentRequest instantiates a new V1ContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContentRequestWithDefaults

`func NewV1ContentRequestWithDefaults() *V1ContentRequest`

NewV1ContentRequestWithDefaults instantiates a new V1ContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1ContentRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1ContentRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1ContentRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetContent

`func (o *V1ContentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V1ContentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V1ContentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetPublished

`func (o *V1ContentRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *V1ContentRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *V1ContentRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.


### GetContentTimestamp

`func (o *V1ContentRequest) GetContentTimestamp() int32`

GetContentTimestamp returns the ContentTimestamp field if non-nil, zero value otherwise.

### GetContentTimestampOk

`func (o *V1ContentRequest) GetContentTimestampOk() (*int32, bool)`

GetContentTimestampOk returns a tuple with the ContentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTimestamp

`func (o *V1ContentRequest) SetContentTimestamp(v int32)`

SetContentTimestamp sets ContentTimestamp field to given value.

### HasContentTimestamp

`func (o *V1ContentRequest) HasContentTimestamp() bool`

HasContentTimestamp returns a boolean if a field has been set.

### GetInfo

`func (o *V1ContentRequest) GetInfo() V1Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *V1ContentRequest) GetInfoOk() (*V1Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *V1ContentRequest) SetInfo(v V1Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *V1ContentRequest) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetTransitions

`func (o *V1ContentRequest) GetTransitions() []V1Transition`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *V1ContentRequest) GetTransitionsOk() (*[]V1Transition, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *V1ContentRequest) SetTransitions(v []V1Transition)`

SetTransitions sets Transitions field to given value.

### HasTransitions

`func (o *V1ContentRequest) HasTransitions() bool`

HasTransitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


