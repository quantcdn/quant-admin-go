# V2KvError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **bool** | Always true. | 
**Message** | **string** |  | 
**Code** | **string** |  | 
**Retryable** | **bool** | True when the same request can be sent again safely. | 
**Key** | Pointer to **string** | CONFLICT: the key the condition named. | [optional] 
**CurrentVersion** | Pointer to **string** | CONFLICT: the current version, when the key exists and has not expired. | [optional] 
**Reason** | Pointer to **string** | CONFLICT: why the condition failed. LEDGER_COLLECTION_NOT_EMPTY, LEDGER_RETAINED, LEDGER_FORKED: the cause. | [optional] 
**RetainUntil** | Pointer to **time.Time** | LEDGER_RETAINED: when the archive retention ends. | [optional] 
**ActiveAfter** | Pointer to **time.Time** | LEDGER_NOT_READY: when the PENDING ledger can become ACTIVE. | [optional] 
**Collection** | Pointer to **string** | LEDGER_RETAINED (unarchived): the collection with unarchived entries. | [optional] 
**Shard** | Pointer to **string** | LEDGER_RETAINED (unarchived): the shard with unarchived entries. | [optional] 
**Operation** | Pointer to **int32** | transact: the index of the failing operation. | [optional] 
**Index** | Pointer to **string** | UNIQUE_VIOLATION: the unique index. | [optional] 
**ExistingKey** | Pointer to **string** | UNIQUE_VIOLATION: the key that holds the value. | [optional] 
**Ledgers** | Pointer to [**[]V2KvErrorLedgersInner**](V2KvErrorLedgersInner.md) | LEDGER_PRESENT (store delete): the ledgers that block the delete. | [optional] 
**Hint** | Pointer to **string** | What to do next, when the KV service gives a hint. | [optional] 
**State** | Pointer to **string** | The current state of the store, index, ledger or export, when the error names it. | [optional] 

## Methods

### NewV2KvError

`func NewV2KvError(error_ bool, message string, code string, retryable bool, ) *V2KvError`

NewV2KvError instantiates a new V2KvError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvErrorWithDefaults

`func NewV2KvErrorWithDefaults() *V2KvError`

NewV2KvErrorWithDefaults instantiates a new V2KvError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V2KvError) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2KvError) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2KvError) SetError(v bool)`

SetError sets Error field to given value.


### GetMessage

`func (o *V2KvError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2KvError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2KvError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *V2KvError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *V2KvError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *V2KvError) SetCode(v string)`

SetCode sets Code field to given value.


### GetRetryable

`func (o *V2KvError) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *V2KvError) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *V2KvError) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.


### GetKey

`func (o *V2KvError) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvError) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvError) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *V2KvError) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *V2KvError) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *V2KvError) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *V2KvError) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *V2KvError) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetReason

`func (o *V2KvError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V2KvError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRetainUntil

`func (o *V2KvError) GetRetainUntil() time.Time`

GetRetainUntil returns the RetainUntil field if non-nil, zero value otherwise.

### GetRetainUntilOk

`func (o *V2KvError) GetRetainUntilOk() (*time.Time, bool)`

GetRetainUntilOk returns a tuple with the RetainUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntil

`func (o *V2KvError) SetRetainUntil(v time.Time)`

SetRetainUntil sets RetainUntil field to given value.

### HasRetainUntil

`func (o *V2KvError) HasRetainUntil() bool`

HasRetainUntil returns a boolean if a field has been set.

### GetActiveAfter

`func (o *V2KvError) GetActiveAfter() time.Time`

GetActiveAfter returns the ActiveAfter field if non-nil, zero value otherwise.

### GetActiveAfterOk

`func (o *V2KvError) GetActiveAfterOk() (*time.Time, bool)`

GetActiveAfterOk returns a tuple with the ActiveAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAfter

`func (o *V2KvError) SetActiveAfter(v time.Time)`

SetActiveAfter sets ActiveAfter field to given value.

### HasActiveAfter

`func (o *V2KvError) HasActiveAfter() bool`

HasActiveAfter returns a boolean if a field has been set.

### GetCollection

`func (o *V2KvError) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvError) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvError) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *V2KvError) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetShard

`func (o *V2KvError) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvError) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvError) SetShard(v string)`

SetShard sets Shard field to given value.

### HasShard

`func (o *V2KvError) HasShard() bool`

HasShard returns a boolean if a field has been set.

### GetOperation

`func (o *V2KvError) GetOperation() int32`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *V2KvError) GetOperationOk() (*int32, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *V2KvError) SetOperation(v int32)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *V2KvError) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetIndex

`func (o *V2KvError) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *V2KvError) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *V2KvError) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *V2KvError) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetExistingKey

`func (o *V2KvError) GetExistingKey() string`

GetExistingKey returns the ExistingKey field if non-nil, zero value otherwise.

### GetExistingKeyOk

`func (o *V2KvError) GetExistingKeyOk() (*string, bool)`

GetExistingKeyOk returns a tuple with the ExistingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingKey

`func (o *V2KvError) SetExistingKey(v string)`

SetExistingKey sets ExistingKey field to given value.

### HasExistingKey

`func (o *V2KvError) HasExistingKey() bool`

HasExistingKey returns a boolean if a field has been set.

### GetLedgers

`func (o *V2KvError) GetLedgers() []V2KvErrorLedgersInner`

GetLedgers returns the Ledgers field if non-nil, zero value otherwise.

### GetLedgersOk

`func (o *V2KvError) GetLedgersOk() (*[]V2KvErrorLedgersInner, bool)`

GetLedgersOk returns a tuple with the Ledgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgers

`func (o *V2KvError) SetLedgers(v []V2KvErrorLedgersInner)`

SetLedgers sets Ledgers field to given value.

### HasLedgers

`func (o *V2KvError) HasLedgers() bool`

HasLedgers returns a boolean if a field has been set.

### GetHint

`func (o *V2KvError) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *V2KvError) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *V2KvError) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *V2KvError) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetState

`func (o *V2KvError) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvError) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvError) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V2KvError) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


