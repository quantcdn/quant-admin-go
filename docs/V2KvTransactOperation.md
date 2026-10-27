# V2KvTransactOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Key** | **string** |  | 
**Version** | Pointer to **string** | check only: the key must have this version and not be expired. | [optional] 
**Value** | Pointer to **interface{}** | set only. Any JSON value. | [optional] 
**Secret** | Pointer to **bool** | set only. Encrypt the value. | [optional] 
**Ttl** | Pointer to **int32** | set only. Seconds (60 to 31,536,000). | [optional] 
**IfAbsent** | Pointer to **bool** | set only. | [optional] 
**IfVersion** | Pointer to **string** | set and delete. | [optional] 
**Amount** | Pointer to **int64** | increment only. Nonzero. | [optional] [default to 1]

## Methods

### NewV2KvTransactOperation

`func NewV2KvTransactOperation(type_ string, key string, ) *V2KvTransactOperation`

NewV2KvTransactOperation instantiates a new V2KvTransactOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvTransactOperationWithDefaults

`func NewV2KvTransactOperationWithDefaults() *V2KvTransactOperation`

NewV2KvTransactOperationWithDefaults instantiates a new V2KvTransactOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V2KvTransactOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2KvTransactOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2KvTransactOperation) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *V2KvTransactOperation) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvTransactOperation) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvTransactOperation) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *V2KvTransactOperation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V2KvTransactOperation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V2KvTransactOperation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V2KvTransactOperation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetValue

`func (o *V2KvTransactOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2KvTransactOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2KvTransactOperation) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *V2KvTransactOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *V2KvTransactOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *V2KvTransactOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *V2KvTransactOperation) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V2KvTransactOperation) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V2KvTransactOperation) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *V2KvTransactOperation) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTtl

`func (o *V2KvTransactOperation) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *V2KvTransactOperation) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *V2KvTransactOperation) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *V2KvTransactOperation) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetIfAbsent

`func (o *V2KvTransactOperation) GetIfAbsent() bool`

GetIfAbsent returns the IfAbsent field if non-nil, zero value otherwise.

### GetIfAbsentOk

`func (o *V2KvTransactOperation) GetIfAbsentOk() (*bool, bool)`

GetIfAbsentOk returns a tuple with the IfAbsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfAbsent

`func (o *V2KvTransactOperation) SetIfAbsent(v bool)`

SetIfAbsent sets IfAbsent field to given value.

### HasIfAbsent

`func (o *V2KvTransactOperation) HasIfAbsent() bool`

HasIfAbsent returns a boolean if a field has been set.

### GetIfVersion

`func (o *V2KvTransactOperation) GetIfVersion() string`

GetIfVersion returns the IfVersion field if non-nil, zero value otherwise.

### GetIfVersionOk

`func (o *V2KvTransactOperation) GetIfVersionOk() (*string, bool)`

GetIfVersionOk returns a tuple with the IfVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfVersion

`func (o *V2KvTransactOperation) SetIfVersion(v string)`

SetIfVersion sets IfVersion field to given value.

### HasIfVersion

`func (o *V2KvTransactOperation) HasIfVersion() bool`

HasIfVersion returns a boolean if a field has been set.

### GetAmount

`func (o *V2KvTransactOperation) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *V2KvTransactOperation) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *V2KvTransactOperation) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *V2KvTransactOperation) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


