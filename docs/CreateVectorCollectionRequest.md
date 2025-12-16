# CreateVectorCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Collection name (used for reference) | 
**Description** | Pointer to **string** |  | [optional] 
**EmbeddingModel** | Pointer to **string** | Embedding model to use (default: amazon.titan-embed-text-v2:0) | [optional] 
**Dimensions** | Pointer to **int32** | Embedding dimensions (default: 1024) | [optional] 

## Methods

### NewCreateVectorCollectionRequest

`func NewCreateVectorCollectionRequest(name string, ) *CreateVectorCollectionRequest`

NewCreateVectorCollectionRequest instantiates a new CreateVectorCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVectorCollectionRequestWithDefaults

`func NewCreateVectorCollectionRequestWithDefaults() *CreateVectorCollectionRequest`

NewCreateVectorCollectionRequestWithDefaults instantiates a new CreateVectorCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVectorCollectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVectorCollectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVectorCollectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateVectorCollectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVectorCollectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVectorCollectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVectorCollectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmbeddingModel

`func (o *CreateVectorCollectionRequest) GetEmbeddingModel() string`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *CreateVectorCollectionRequest) GetEmbeddingModelOk() (*string, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *CreateVectorCollectionRequest) SetEmbeddingModel(v string)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *CreateVectorCollectionRequest) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetDimensions

`func (o *CreateVectorCollectionRequest) GetDimensions() int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CreateVectorCollectionRequest) GetDimensionsOk() (*int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CreateVectorCollectionRequest) SetDimensions(v int32)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CreateVectorCollectionRequest) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


