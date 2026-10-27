# V2KvEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | Pointer to **interface{}** | The stored JSON value (any JSON type), decrypted when it was stored as a secret. In a key list a secret value reads [ENCRYPTED]. | [optional] 
**Encrypted** | Pointer to **bool** | True when the value was stored as a secret. | [optional] 
**Status** | Pointer to **string** | Single-key reads only. | [optional] 
**Message** | Pointer to **string** | Single-key reads only. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **int64** | Epoch seconds when the key expires. Absent when the key does not expire. | [optional] 
**Version** | Pointer to **string** | Format 2 stores only. A ULID, new on every write. Absent on ledger entries. | [optional] 
**Ledger** | Pointer to [**V2KvLedgerMeta**](V2KvLedgerMeta.md) |  | [optional] 
**Personal** | Pointer to **[]string** | Ledger entries only: the personal fields that read \&quot;[PERSONAL]\&quot;. [\&quot;*\&quot;] when the whole value is hidden (the ledger definition could not be read). | [optional] 

## Methods

### NewV2KvEntry

`func NewV2KvEntry(key string, ) *V2KvEntry`

NewV2KvEntry instantiates a new V2KvEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvEntryWithDefaults

`func NewV2KvEntryWithDefaults() *V2KvEntry`

NewV2KvEntryWithDefaults instantiates a new V2KvEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvEntry) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *V2KvEntry) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2KvEntry) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2KvEntry) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *V2KvEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *V2KvEntry) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *V2KvEntry) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEncrypted

`func (o *V2KvEntry) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *V2KvEntry) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *V2KvEntry) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *V2KvEntry) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetStatus

`func (o *V2KvEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2KvEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2KvEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V2KvEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *V2KvEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2KvEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2KvEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V2KvEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *V2KvEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2KvEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2KvEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2KvEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2KvEntry) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2KvEntry) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2KvEntry) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2KvEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *V2KvEntry) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *V2KvEntry) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *V2KvEntry) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *V2KvEntry) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetVersion

`func (o *V2KvEntry) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V2KvEntry) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V2KvEntry) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V2KvEntry) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLedger

`func (o *V2KvEntry) GetLedger() V2KvLedgerMeta`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvEntry) GetLedgerOk() (*V2KvLedgerMeta, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvEntry) SetLedger(v V2KvLedgerMeta)`

SetLedger sets Ledger field to given value.

### HasLedger

`func (o *V2KvEntry) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### GetPersonal

`func (o *V2KvEntry) GetPersonal() []string`

GetPersonal returns the Personal field if non-nil, zero value otherwise.

### GetPersonalOk

`func (o *V2KvEntry) GetPersonalOk() (*[]string, bool)`

GetPersonalOk returns a tuple with the Personal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonal

`func (o *V2KvEntry) SetPersonal(v []string)`

SetPersonal sets Personal field to given value.

### HasPersonal

`func (o *V2KvEntry) HasPersonal() bool`

HasPersonal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


