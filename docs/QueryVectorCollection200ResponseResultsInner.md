# QueryVectorCollection200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** | Document text content | [optional] 
**Similarity** | Pointer to **float32** | Cosine similarity score (1.0 for metadata-only queries) | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Embedding** | Pointer to **[]float32** | Vector embedding (only if includeEmbeddings&#x3D;true) | [optional] 

## Methods

### NewQueryVectorCollection200ResponseResultsInner

`func NewQueryVectorCollection200ResponseResultsInner() *QueryVectorCollection200ResponseResultsInner`

NewQueryVectorCollection200ResponseResultsInner instantiates a new QueryVectorCollection200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryVectorCollection200ResponseResultsInnerWithDefaults

`func NewQueryVectorCollection200ResponseResultsInnerWithDefaults() *QueryVectorCollection200ResponseResultsInner`

NewQueryVectorCollection200ResponseResultsInnerWithDefaults instantiates a new QueryVectorCollection200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *QueryVectorCollection200ResponseResultsInner) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *QueryVectorCollection200ResponseResultsInner) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *QueryVectorCollection200ResponseResultsInner) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *QueryVectorCollection200ResponseResultsInner) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetContent

`func (o *QueryVectorCollection200ResponseResultsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *QueryVectorCollection200ResponseResultsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *QueryVectorCollection200ResponseResultsInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *QueryVectorCollection200ResponseResultsInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSimilarity

`func (o *QueryVectorCollection200ResponseResultsInner) GetSimilarity() float32`

GetSimilarity returns the Similarity field if non-nil, zero value otherwise.

### GetSimilarityOk

`func (o *QueryVectorCollection200ResponseResultsInner) GetSimilarityOk() (*float32, bool)`

GetSimilarityOk returns a tuple with the Similarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarity

`func (o *QueryVectorCollection200ResponseResultsInner) SetSimilarity(v float32)`

SetSimilarity sets Similarity field to given value.

### HasSimilarity

`func (o *QueryVectorCollection200ResponseResultsInner) HasSimilarity() bool`

HasSimilarity returns a boolean if a field has been set.

### GetMetadata

`func (o *QueryVectorCollection200ResponseResultsInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QueryVectorCollection200ResponseResultsInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QueryVectorCollection200ResponseResultsInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QueryVectorCollection200ResponseResultsInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEmbedding

`func (o *QueryVectorCollection200ResponseResultsInner) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *QueryVectorCollection200ResponseResultsInner) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *QueryVectorCollection200ResponseResultsInner) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.

### HasEmbedding

`func (o *QueryVectorCollection200ResponseResultsInner) HasEmbedding() bool`

HasEmbedding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


