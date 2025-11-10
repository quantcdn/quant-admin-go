# Embeddings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embeddings** | [**Embeddings200ResponseEmbeddings**](Embeddings200ResponseEmbeddings.md) |  | 
**Model** | **string** | Model used to generate embeddings | 
**Dimension** | **int32** | Dimensionality of each embedding vector | 
**Usage** | [**Embeddings200ResponseUsage**](Embeddings200ResponseUsage.md) |  | 

## Methods

### NewEmbeddings200Response

`func NewEmbeddings200Response(embeddings Embeddings200ResponseEmbeddings, model string, dimension int32, usage Embeddings200ResponseUsage, ) *Embeddings200Response`

NewEmbeddings200Response instantiates a new Embeddings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddings200ResponseWithDefaults

`func NewEmbeddings200ResponseWithDefaults() *Embeddings200Response`

NewEmbeddings200ResponseWithDefaults instantiates a new Embeddings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbeddings

`func (o *Embeddings200Response) GetEmbeddings() Embeddings200ResponseEmbeddings`

GetEmbeddings returns the Embeddings field if non-nil, zero value otherwise.

### GetEmbeddingsOk

`func (o *Embeddings200Response) GetEmbeddingsOk() (*Embeddings200ResponseEmbeddings, bool)`

GetEmbeddingsOk returns a tuple with the Embeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddings

`func (o *Embeddings200Response) SetEmbeddings(v Embeddings200ResponseEmbeddings)`

SetEmbeddings sets Embeddings field to given value.


### GetModel

`func (o *Embeddings200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Embeddings200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Embeddings200Response) SetModel(v string)`

SetModel sets Model field to given value.


### GetDimension

`func (o *Embeddings200Response) GetDimension() int32`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *Embeddings200Response) GetDimensionOk() (*int32, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *Embeddings200Response) SetDimension(v int32)`

SetDimension sets Dimension field to given value.


### GetUsage

`func (o *Embeddings200Response) GetUsage() Embeddings200ResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Embeddings200Response) GetUsageOk() (*Embeddings200ResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Embeddings200Response) SetUsage(v Embeddings200ResponseUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


