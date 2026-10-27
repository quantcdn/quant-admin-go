# V2KvTransactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | Pointer to **string** | A repeat of the same operations within 24 hours returns the first result with replayed true; other operations with the same key fail VALIDATION. Send it to make a lost-connection 503 retryable. | [optional] 
**Operations** | [**[]V2KvTransactOperation**](V2KvTransactOperation.md) | 1 to 25 operations on different keys. More than 25 is 400 LIMIT_EXCEEDED. | 

## Methods

### NewV2KvTransactRequest

`func NewV2KvTransactRequest(operations []V2KvTransactOperation, ) *V2KvTransactRequest`

NewV2KvTransactRequest instantiates a new V2KvTransactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvTransactRequestWithDefaults

`func NewV2KvTransactRequestWithDefaults() *V2KvTransactRequest`

NewV2KvTransactRequestWithDefaults instantiates a new V2KvTransactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdempotencyKey

`func (o *V2KvTransactRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *V2KvTransactRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *V2KvTransactRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *V2KvTransactRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetOperations

`func (o *V2KvTransactRequest) GetOperations() []V2KvTransactOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *V2KvTransactRequest) GetOperationsOk() (*[]V2KvTransactOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *V2KvTransactRequest) SetOperations(v []V2KvTransactOperation)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


