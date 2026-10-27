# V2KvLedgerVerifyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ok** | **bool** |  | 
**Entries** | **int32** | Live entries checked in this call. | 
**Shards** | [**[]V2KvLedgerVerifyShard**](V2KvLedgerVerifyShard.md) |  | 
**VerificationId** | Pointer to **string** | The verification record; KV_ledgers_verifications_report returns it signed. | [optional] 
**NextCursor** | **NullableString** |  | 

## Methods

### NewV2KvLedgerVerifyResult

`func NewV2KvLedgerVerifyResult(ok bool, entries int32, shards []V2KvLedgerVerifyShard, nextCursor NullableString, ) *V2KvLedgerVerifyResult`

NewV2KvLedgerVerifyResult instantiates a new V2KvLedgerVerifyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerVerifyResultWithDefaults

`func NewV2KvLedgerVerifyResultWithDefaults() *V2KvLedgerVerifyResult`

NewV2KvLedgerVerifyResultWithDefaults instantiates a new V2KvLedgerVerifyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOk

`func (o *V2KvLedgerVerifyResult) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *V2KvLedgerVerifyResult) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *V2KvLedgerVerifyResult) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetEntries

`func (o *V2KvLedgerVerifyResult) GetEntries() int32`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerVerifyResult) GetEntriesOk() (*int32, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerVerifyResult) SetEntries(v int32)`

SetEntries sets Entries field to given value.


### GetShards

`func (o *V2KvLedgerVerifyResult) GetShards() []V2KvLedgerVerifyShard`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *V2KvLedgerVerifyResult) GetShardsOk() (*[]V2KvLedgerVerifyShard, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *V2KvLedgerVerifyResult) SetShards(v []V2KvLedgerVerifyShard)`

SetShards sets Shards field to given value.


### GetVerificationId

`func (o *V2KvLedgerVerifyResult) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *V2KvLedgerVerifyResult) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *V2KvLedgerVerifyResult) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *V2KvLedgerVerifyResult) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.

### GetNextCursor

`func (o *V2KvLedgerVerifyResult) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V2KvLedgerVerifyResult) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V2KvLedgerVerifyResult) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *V2KvLedgerVerifyResult) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *V2KvLedgerVerifyResult) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


