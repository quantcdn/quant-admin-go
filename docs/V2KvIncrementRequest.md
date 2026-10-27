# V2KvIncrementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | A nonzero integer. The result must stay a safe integer. | [optional] [default to 1]
**Ttl** | Pointer to **int32** | Seconds (60 to 31,536,000). Applied only when the key is created or reset after expiry. | [optional] 

## Methods

### NewV2KvIncrementRequest

`func NewV2KvIncrementRequest() *V2KvIncrementRequest`

NewV2KvIncrementRequest instantiates a new V2KvIncrementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvIncrementRequestWithDefaults

`func NewV2KvIncrementRequestWithDefaults() *V2KvIncrementRequest`

NewV2KvIncrementRequestWithDefaults instantiates a new V2KvIncrementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *V2KvIncrementRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *V2KvIncrementRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *V2KvIncrementRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *V2KvIncrementRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTtl

`func (o *V2KvIncrementRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *V2KvIncrementRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *V2KvIncrementRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *V2KvIncrementRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


