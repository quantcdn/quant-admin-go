# ListAIModels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Models** | Pointer to [**[]ListAIModels200ResponseModelsInner**](ListAIModels200ResponseModelsInner.md) |  | [optional] 

## Methods

### NewListAIModels200Response

`func NewListAIModels200Response() *ListAIModels200Response`

NewListAIModels200Response instantiates a new ListAIModels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAIModels200ResponseWithDefaults

`func NewListAIModels200ResponseWithDefaults() *ListAIModels200Response`

NewListAIModels200ResponseWithDefaults instantiates a new ListAIModels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListAIModels200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListAIModels200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListAIModels200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListAIModels200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetModels

`func (o *ListAIModels200Response) GetModels() []ListAIModels200ResponseModelsInner`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ListAIModels200Response) GetModelsOk() (*[]ListAIModels200ResponseModelsInner, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ListAIModels200Response) SetModels(v []ListAIModels200ResponseModelsInner)`

SetModels sets Models field to given value.

### HasModels

`func (o *ListAIModels200Response) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


