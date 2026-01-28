# QueryVectorCollectionRequestFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exact** | Pointer to **map[string]interface{}** | Exact match on metadata fields. Keys are metadata field names, values are expected values. | [optional] 
**Contains** | Pointer to **map[string][]string** | Array contains filter for array metadata fields (like tags). Returns documents where the metadata array contains ANY of the specified values. | [optional] 

## Methods

### NewQueryVectorCollectionRequestFilter

`func NewQueryVectorCollectionRequestFilter() *QueryVectorCollectionRequestFilter`

NewQueryVectorCollectionRequestFilter instantiates a new QueryVectorCollectionRequestFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryVectorCollectionRequestFilterWithDefaults

`func NewQueryVectorCollectionRequestFilterWithDefaults() *QueryVectorCollectionRequestFilter`

NewQueryVectorCollectionRequestFilterWithDefaults instantiates a new QueryVectorCollectionRequestFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExact

`func (o *QueryVectorCollectionRequestFilter) GetExact() map[string]interface{}`

GetExact returns the Exact field if non-nil, zero value otherwise.

### GetExactOk

`func (o *QueryVectorCollectionRequestFilter) GetExactOk() (*map[string]interface{}, bool)`

GetExactOk returns a tuple with the Exact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExact

`func (o *QueryVectorCollectionRequestFilter) SetExact(v map[string]interface{})`

SetExact sets Exact field to given value.

### HasExact

`func (o *QueryVectorCollectionRequestFilter) HasExact() bool`

HasExact returns a boolean if a field has been set.

### GetContains

`func (o *QueryVectorCollectionRequestFilter) GetContains() map[string][]string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *QueryVectorCollectionRequestFilter) GetContainsOk() (*map[string][]string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *QueryVectorCollectionRequestFilter) SetContains(v map[string][]string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *QueryVectorCollectionRequestFilter) HasContains() bool`

HasContains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


