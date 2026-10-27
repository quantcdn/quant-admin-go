# V2KvBatchGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | **[]string** | 1 to 100 keys. Duplicates are read once. More than 100 is 400 LIMIT_EXCEEDED. | 
**Consistent** | Pointer to **bool** | false for an eventually consistent read. | [optional] [default to true]

## Methods

### NewV2KvBatchGetRequest

`func NewV2KvBatchGetRequest(keys []string, ) *V2KvBatchGetRequest`

NewV2KvBatchGetRequest instantiates a new V2KvBatchGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvBatchGetRequestWithDefaults

`func NewV2KvBatchGetRequestWithDefaults() *V2KvBatchGetRequest`

NewV2KvBatchGetRequestWithDefaults instantiates a new V2KvBatchGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *V2KvBatchGetRequest) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *V2KvBatchGetRequest) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *V2KvBatchGetRequest) SetKeys(v []string)`

SetKeys sets Keys field to given value.


### GetConsistent

`func (o *V2KvBatchGetRequest) GetConsistent() bool`

GetConsistent returns the Consistent field if non-nil, zero value otherwise.

### GetConsistentOk

`func (o *V2KvBatchGetRequest) GetConsistentOk() (*bool, bool)`

GetConsistentOk returns a tuple with the Consistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistent

`func (o *V2KvBatchGetRequest) SetConsistent(v bool)`

SetConsistent sets Consistent field to given value.

### HasConsistent

`func (o *V2KvBatchGetRequest) HasConsistent() bool`

HasConsistent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


