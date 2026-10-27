# V2KvEntryPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **interface{}** | Any JSON value. An empty object {} is stored as {}. | 
**Secret** | Pointer to **bool** | Encrypt the value with KMS. Lists show [ENCRYPTED]. Not in a collection with an index on the value (INDEX_FIELD_ENCRYPTED). | [optional] [default to false]
**Ttl** | Pointer to **int32** | Seconds until the key expires (60 to 31,536,000). A write without ttl removes any expiry. | [optional] 
**IfAbsent** | Pointer to **bool** | Format 2 only. Write only if the key is missing or expired (201). Not with ifVersion. | [optional] 
**IfVersion** | Pointer to **string** | Format 2 only. Write only if the current version is this one. | [optional] 

## Methods

### NewV2KvEntryPutRequest

`func NewV2KvEntryPutRequest(value interface{}, ) *V2KvEntryPutRequest`

NewV2KvEntryPutRequest instantiates a new V2KvEntryPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvEntryPutRequestWithDefaults

`func NewV2KvEntryPutRequestWithDefaults() *V2KvEntryPutRequest`

NewV2KvEntryPutRequestWithDefaults instantiates a new V2KvEntryPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *V2KvEntryPutRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2KvEntryPutRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2KvEntryPutRequest) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *V2KvEntryPutRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *V2KvEntryPutRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSecret

`func (o *V2KvEntryPutRequest) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V2KvEntryPutRequest) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V2KvEntryPutRequest) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *V2KvEntryPutRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTtl

`func (o *V2KvEntryPutRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *V2KvEntryPutRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *V2KvEntryPutRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *V2KvEntryPutRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetIfAbsent

`func (o *V2KvEntryPutRequest) GetIfAbsent() bool`

GetIfAbsent returns the IfAbsent field if non-nil, zero value otherwise.

### GetIfAbsentOk

`func (o *V2KvEntryPutRequest) GetIfAbsentOk() (*bool, bool)`

GetIfAbsentOk returns a tuple with the IfAbsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfAbsent

`func (o *V2KvEntryPutRequest) SetIfAbsent(v bool)`

SetIfAbsent sets IfAbsent field to given value.

### HasIfAbsent

`func (o *V2KvEntryPutRequest) HasIfAbsent() bool`

HasIfAbsent returns a boolean if a field has been set.

### GetIfVersion

`func (o *V2KvEntryPutRequest) GetIfVersion() string`

GetIfVersion returns the IfVersion field if non-nil, zero value otherwise.

### GetIfVersionOk

`func (o *V2KvEntryPutRequest) GetIfVersionOk() (*string, bool)`

GetIfVersionOk returns a tuple with the IfVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfVersion

`func (o *V2KvEntryPutRequest) SetIfVersion(v string)`

SetIfVersion sets IfVersion field to given value.

### HasIfVersion

`func (o *V2KvEntryPutRequest) HasIfVersion() bool`

HasIfVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


