# QueryVectorCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | Natural language search query (mutually exclusive with vector) | [optional] 
**Vector** | Pointer to **[]float32** | Pre-computed embedding vector (mutually exclusive with query). Array length must match collection dimension. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of results to return | [optional] [default to 5]
**Threshold** | Pointer to **float32** | Minimum similarity score (0-1, higher &#x3D; more relevant) | [optional] [default to 0.7]
**IncludeEmbeddings** | Pointer to **bool** | Include embedding vectors in response (for debugging) | [optional] [default to false]
**Filter** | Pointer to [**QueryVectorCollectionRequestFilter**](QueryVectorCollectionRequestFilter.md) |  | [optional] 
**ListByMetadata** | Pointer to **bool** | If true, skip semantic search and return all documents matching the filter. Requires filter. Supports cursor pagination. | [optional] [default to false]
**Cursor** | Pointer to **string** | Pagination cursor for listByMetadata mode. Use nextCursor from previous response. Opaque format - do not construct manually. | [optional] 
**SortBy** | Pointer to **string** | Field to sort by in listByMetadata mode | [optional] [default to "created_at"]
**SortOrder** | Pointer to **string** | Sort direction in listByMetadata mode | [optional] [default to "desc"]

## Methods

### NewQueryVectorCollectionRequest

`func NewQueryVectorCollectionRequest() *QueryVectorCollectionRequest`

NewQueryVectorCollectionRequest instantiates a new QueryVectorCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryVectorCollectionRequestWithDefaults

`func NewQueryVectorCollectionRequestWithDefaults() *QueryVectorCollectionRequest`

NewQueryVectorCollectionRequestWithDefaults instantiates a new QueryVectorCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *QueryVectorCollectionRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryVectorCollectionRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryVectorCollectionRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryVectorCollectionRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetVector

`func (o *QueryVectorCollectionRequest) GetVector() []float32`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *QueryVectorCollectionRequest) GetVectorOk() (*[]float32, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *QueryVectorCollectionRequest) SetVector(v []float32)`

SetVector sets Vector field to given value.

### HasVector

`func (o *QueryVectorCollectionRequest) HasVector() bool`

HasVector returns a boolean if a field has been set.

### GetLimit

`func (o *QueryVectorCollectionRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryVectorCollectionRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryVectorCollectionRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryVectorCollectionRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetThreshold

`func (o *QueryVectorCollectionRequest) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *QueryVectorCollectionRequest) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *QueryVectorCollectionRequest) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *QueryVectorCollectionRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetIncludeEmbeddings

`func (o *QueryVectorCollectionRequest) GetIncludeEmbeddings() bool`

GetIncludeEmbeddings returns the IncludeEmbeddings field if non-nil, zero value otherwise.

### GetIncludeEmbeddingsOk

`func (o *QueryVectorCollectionRequest) GetIncludeEmbeddingsOk() (*bool, bool)`

GetIncludeEmbeddingsOk returns a tuple with the IncludeEmbeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEmbeddings

`func (o *QueryVectorCollectionRequest) SetIncludeEmbeddings(v bool)`

SetIncludeEmbeddings sets IncludeEmbeddings field to given value.

### HasIncludeEmbeddings

`func (o *QueryVectorCollectionRequest) HasIncludeEmbeddings() bool`

HasIncludeEmbeddings returns a boolean if a field has been set.

### GetFilter

`func (o *QueryVectorCollectionRequest) GetFilter() QueryVectorCollectionRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *QueryVectorCollectionRequest) GetFilterOk() (*QueryVectorCollectionRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *QueryVectorCollectionRequest) SetFilter(v QueryVectorCollectionRequestFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *QueryVectorCollectionRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetListByMetadata

`func (o *QueryVectorCollectionRequest) GetListByMetadata() bool`

GetListByMetadata returns the ListByMetadata field if non-nil, zero value otherwise.

### GetListByMetadataOk

`func (o *QueryVectorCollectionRequest) GetListByMetadataOk() (*bool, bool)`

GetListByMetadataOk returns a tuple with the ListByMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListByMetadata

`func (o *QueryVectorCollectionRequest) SetListByMetadata(v bool)`

SetListByMetadata sets ListByMetadata field to given value.

### HasListByMetadata

`func (o *QueryVectorCollectionRequest) HasListByMetadata() bool`

HasListByMetadata returns a boolean if a field has been set.

### GetCursor

`func (o *QueryVectorCollectionRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *QueryVectorCollectionRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *QueryVectorCollectionRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *QueryVectorCollectionRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetSortBy

`func (o *QueryVectorCollectionRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *QueryVectorCollectionRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *QueryVectorCollectionRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *QueryVectorCollectionRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSortOrder

`func (o *QueryVectorCollectionRequest) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *QueryVectorCollectionRequest) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *QueryVectorCollectionRequest) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *QueryVectorCollectionRequest) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


