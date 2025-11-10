# EmbeddingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | [**EmbeddingsRequestInput**](EmbeddingsRequestInput.md) |  | 
**ModelId** | Pointer to **string** | Embedding model to use | [optional] [default to "amazon.titan-embed-text-v2:0"]
**Dimensions** | Pointer to **int32** | Output embedding dimensions. Titan v2 supports: 256, 512, 1024, 8192 | [optional] [default to 1024]
**Normalize** | Pointer to **bool** | Normalize embeddings to unit length (magnitude &#x3D; 1.0) | [optional] [default to true]

## Methods

### NewEmbeddingsRequest

`func NewEmbeddingsRequest(input EmbeddingsRequestInput, ) *EmbeddingsRequest`

NewEmbeddingsRequest instantiates a new EmbeddingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingsRequestWithDefaults

`func NewEmbeddingsRequestWithDefaults() *EmbeddingsRequest`

NewEmbeddingsRequestWithDefaults instantiates a new EmbeddingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *EmbeddingsRequest) GetInput() EmbeddingsRequestInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EmbeddingsRequest) GetInputOk() (*EmbeddingsRequestInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EmbeddingsRequest) SetInput(v EmbeddingsRequestInput)`

SetInput sets Input field to given value.


### GetModelId

`func (o *EmbeddingsRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *EmbeddingsRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *EmbeddingsRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *EmbeddingsRequest) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetDimensions

`func (o *EmbeddingsRequest) GetDimensions() int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *EmbeddingsRequest) GetDimensionsOk() (*int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *EmbeddingsRequest) SetDimensions(v int32)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *EmbeddingsRequest) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetNormalize

`func (o *EmbeddingsRequest) GetNormalize() bool`

GetNormalize returns the Normalize field if non-nil, zero value otherwise.

### GetNormalizeOk

`func (o *EmbeddingsRequest) GetNormalizeOk() (*bool, bool)`

GetNormalizeOk returns a tuple with the Normalize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalize

`func (o *EmbeddingsRequest) SetNormalize(v bool)`

SetNormalize sets Normalize field to given value.

### HasNormalize

`func (o *EmbeddingsRequest) HasNormalize() bool`

HasNormalize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


