# QueryVectorCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]QueryVectorCollection200ResponseResultsInner**](QueryVectorCollection200ResponseResultsInner.md) |  | [optional] 
**Query** | Pointer to **NullableString** | Original query text (null if vector or metadata search was used) | [optional] 
**SearchMode** | Pointer to **string** | Search mode used: text (query provided), vector (pre-computed), metadata (listByMetadata) | [optional] 
**Filter** | Pointer to **map[string]interface{}** | Filter that was applied (if any) | [optional] 
**Count** | Pointer to **int32** | Number of results returned | [optional] 
**ExecutionTimeMs** | Pointer to **int32** | Query execution time in milliseconds | [optional] 
**CollectionId** | Pointer to **string** |  | [optional] 
**HasMore** | Pointer to **bool** | True if more results available (listByMetadata mode only) | [optional] 
**NextCursor** | Pointer to **NullableString** | Cursor for next page. Pass as cursor param to continue. Null when no more results. Only in listByMetadata mode. | [optional] 
**Pagination** | Pointer to [**NullableQueryVectorCollection200ResponsePagination**](QueryVectorCollection200ResponsePagination.md) |  | [optional] 

## Methods

### NewQueryVectorCollection200Response

`func NewQueryVectorCollection200Response() *QueryVectorCollection200Response`

NewQueryVectorCollection200Response instantiates a new QueryVectorCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryVectorCollection200ResponseWithDefaults

`func NewQueryVectorCollection200ResponseWithDefaults() *QueryVectorCollection200Response`

NewQueryVectorCollection200ResponseWithDefaults instantiates a new QueryVectorCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *QueryVectorCollection200Response) GetResults() []QueryVectorCollection200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryVectorCollection200Response) GetResultsOk() (*[]QueryVectorCollection200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryVectorCollection200Response) SetResults(v []QueryVectorCollection200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *QueryVectorCollection200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetQuery

`func (o *QueryVectorCollection200Response) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryVectorCollection200Response) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryVectorCollection200Response) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryVectorCollection200Response) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *QueryVectorCollection200Response) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *QueryVectorCollection200Response) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetSearchMode

`func (o *QueryVectorCollection200Response) GetSearchMode() string`

GetSearchMode returns the SearchMode field if non-nil, zero value otherwise.

### GetSearchModeOk

`func (o *QueryVectorCollection200Response) GetSearchModeOk() (*string, bool)`

GetSearchModeOk returns a tuple with the SearchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMode

`func (o *QueryVectorCollection200Response) SetSearchMode(v string)`

SetSearchMode sets SearchMode field to given value.

### HasSearchMode

`func (o *QueryVectorCollection200Response) HasSearchMode() bool`

HasSearchMode returns a boolean if a field has been set.

### GetFilter

`func (o *QueryVectorCollection200Response) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *QueryVectorCollection200Response) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *QueryVectorCollection200Response) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *QueryVectorCollection200Response) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *QueryVectorCollection200Response) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *QueryVectorCollection200Response) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetCount

`func (o *QueryVectorCollection200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *QueryVectorCollection200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *QueryVectorCollection200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *QueryVectorCollection200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetExecutionTimeMs

`func (o *QueryVectorCollection200Response) GetExecutionTimeMs() int32`

GetExecutionTimeMs returns the ExecutionTimeMs field if non-nil, zero value otherwise.

### GetExecutionTimeMsOk

`func (o *QueryVectorCollection200Response) GetExecutionTimeMsOk() (*int32, bool)`

GetExecutionTimeMsOk returns a tuple with the ExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeMs

`func (o *QueryVectorCollection200Response) SetExecutionTimeMs(v int32)`

SetExecutionTimeMs sets ExecutionTimeMs field to given value.

### HasExecutionTimeMs

`func (o *QueryVectorCollection200Response) HasExecutionTimeMs() bool`

HasExecutionTimeMs returns a boolean if a field has been set.

### GetCollectionId

`func (o *QueryVectorCollection200Response) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *QueryVectorCollection200Response) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *QueryVectorCollection200Response) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *QueryVectorCollection200Response) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetHasMore

`func (o *QueryVectorCollection200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *QueryVectorCollection200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *QueryVectorCollection200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *QueryVectorCollection200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNextCursor

`func (o *QueryVectorCollection200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *QueryVectorCollection200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *QueryVectorCollection200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *QueryVectorCollection200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *QueryVectorCollection200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *QueryVectorCollection200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetPagination

`func (o *QueryVectorCollection200Response) GetPagination() QueryVectorCollection200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *QueryVectorCollection200Response) GetPaginationOk() (*QueryVectorCollection200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *QueryVectorCollection200Response) SetPagination(v QueryVectorCollection200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *QueryVectorCollection200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### SetPaginationNil

`func (o *QueryVectorCollection200Response) SetPaginationNil(b bool)`

 SetPaginationNil sets the value for Pagination to be an explicit nil

### UnsetPagination
`func (o *QueryVectorCollection200Response) UnsetPagination()`

UnsetPagination ensures that no value is present for Pagination, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


