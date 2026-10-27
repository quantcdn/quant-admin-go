# V2KvIndexLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | Pointer to **interface{}** | The stored JSON value (any JSON type). | [optional] 
**Status** | **string** |  | 
**Message** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **int32** | Epoch seconds. Absent when the key does not expire. | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** | ULID, new on every write. Absent on ledger entries. | [optional] 
**Ledger** | Pointer to [**V2KvLedgerMeta**](V2KvLedgerMeta.md) |  | [optional] 
**Personal** | Pointer to **[]string** | Ledger entries only: the personal fields the portal replaced with [PERSONAL] (* &#x3D; the whole value). | [optional] 

## Methods

### NewV2KvIndexLookup

`func NewV2KvIndexLookup(key string, status string, message string, ) *V2KvIndexLookup`

NewV2KvIndexLookup instantiates a new V2KvIndexLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvIndexLookupWithDefaults

`func NewV2KvIndexLookupWithDefaults() *V2KvIndexLookup`

NewV2KvIndexLookupWithDefaults instantiates a new V2KvIndexLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvIndexLookup) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvIndexLookup) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvIndexLookup) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *V2KvIndexLookup) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2KvIndexLookup) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2KvIndexLookup) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *V2KvIndexLookup) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *V2KvIndexLookup) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *V2KvIndexLookup) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetStatus

`func (o *V2KvIndexLookup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2KvIndexLookup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2KvIndexLookup) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *V2KvIndexLookup) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2KvIndexLookup) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2KvIndexLookup) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCreatedAt

`func (o *V2KvIndexLookup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2KvIndexLookup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2KvIndexLookup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2KvIndexLookup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2KvIndexLookup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2KvIndexLookup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2KvIndexLookup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2KvIndexLookup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *V2KvIndexLookup) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *V2KvIndexLookup) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *V2KvIndexLookup) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *V2KvIndexLookup) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetEncrypted

`func (o *V2KvIndexLookup) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *V2KvIndexLookup) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *V2KvIndexLookup) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *V2KvIndexLookup) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetVersion

`func (o *V2KvIndexLookup) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V2KvIndexLookup) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V2KvIndexLookup) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V2KvIndexLookup) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLedger

`func (o *V2KvIndexLookup) GetLedger() V2KvLedgerMeta`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvIndexLookup) GetLedgerOk() (*V2KvLedgerMeta, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvIndexLookup) SetLedger(v V2KvLedgerMeta)`

SetLedger sets Ledger field to given value.

### HasLedger

`func (o *V2KvIndexLookup) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### GetPersonal

`func (o *V2KvIndexLookup) GetPersonal() []string`

GetPersonal returns the Personal field if non-nil, zero value otherwise.

### GetPersonalOk

`func (o *V2KvIndexLookup) GetPersonalOk() (*[]string, bool)`

GetPersonalOk returns a tuple with the Personal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonal

`func (o *V2KvIndexLookup) SetPersonal(v []string)`

SetPersonal sets Personal field to given value.

### HasPersonal

`func (o *V2KvIndexLookup) HasPersonal() bool`

HasPersonal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


