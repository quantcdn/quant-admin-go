# V2KvLedgerVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to **string** | One concrete collection of the ledger (no *). Absent: every collection. | [optional] 
**Since** | Pointer to [**V2KvLedgerVerifyRequestSince**](V2KvLedgerVerifyRequestSince.md) |  | [optional] 
**Until** | Pointer to [**V2KvLedgerVerifyRequestUntil**](V2KvLedgerVerifyRequestUntil.md) |  | [optional] 
**Cursor** | Pointer to **string** | The next_cursor of the previous call (signed, bound to the range, lasts 24 hours). | [optional] 
**MaxEntries** | Pointer to **int32** |  | [optional] [default to 10000]

## Methods

### NewV2KvLedgerVerifyRequest

`func NewV2KvLedgerVerifyRequest() *V2KvLedgerVerifyRequest`

NewV2KvLedgerVerifyRequest instantiates a new V2KvLedgerVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerVerifyRequestWithDefaults

`func NewV2KvLedgerVerifyRequestWithDefaults() *V2KvLedgerVerifyRequest`

NewV2KvLedgerVerifyRequestWithDefaults instantiates a new V2KvLedgerVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *V2KvLedgerVerifyRequest) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerVerifyRequest) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerVerifyRequest) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *V2KvLedgerVerifyRequest) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetSince

`func (o *V2KvLedgerVerifyRequest) GetSince() V2KvLedgerVerifyRequestSince`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *V2KvLedgerVerifyRequest) GetSinceOk() (*V2KvLedgerVerifyRequestSince, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *V2KvLedgerVerifyRequest) SetSince(v V2KvLedgerVerifyRequestSince)`

SetSince sets Since field to given value.

### HasSince

`func (o *V2KvLedgerVerifyRequest) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetUntil

`func (o *V2KvLedgerVerifyRequest) GetUntil() V2KvLedgerVerifyRequestUntil`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *V2KvLedgerVerifyRequest) GetUntilOk() (*V2KvLedgerVerifyRequestUntil, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *V2KvLedgerVerifyRequest) SetUntil(v V2KvLedgerVerifyRequestUntil)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *V2KvLedgerVerifyRequest) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetCursor

`func (o *V2KvLedgerVerifyRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *V2KvLedgerVerifyRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *V2KvLedgerVerifyRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *V2KvLedgerVerifyRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetMaxEntries

`func (o *V2KvLedgerVerifyRequest) GetMaxEntries() int32`

GetMaxEntries returns the MaxEntries field if non-nil, zero value otherwise.

### GetMaxEntriesOk

`func (o *V2KvLedgerVerifyRequest) GetMaxEntriesOk() (*int32, bool)`

GetMaxEntriesOk returns a tuple with the MaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntries

`func (o *V2KvLedgerVerifyRequest) SetMaxEntries(v int32)`

SetMaxEntries sets MaxEntries field to given value.

### HasMaxEntries

`func (o *V2KvLedgerVerifyRequest) HasMaxEntries() bool`

HasMaxEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


