# QueryVectorCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]QueryVectorCollection200ResponseResultsInner**](QueryVectorCollection200ResponseResultsInner.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**ExecutionTimeMs** | Pointer to **int32** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


