# QueryVectorCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Natural language search query | 
**Count** | Pointer to **int32** | Number of results to return | [optional] [default to 5]

## Methods

### NewQueryVectorCollectionRequest

`func NewQueryVectorCollectionRequest(query string, ) *QueryVectorCollectionRequest`

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


### GetCount

`func (o *QueryVectorCollectionRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *QueryVectorCollectionRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *QueryVectorCollectionRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *QueryVectorCollectionRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


