# V2KvLedgerLastVerify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The verification record; KV_ledgers_verifications_report returns it signed. | [optional] 
**At** | **time.Time** |  | 
**StartedAt** | **time.Time** |  | 
**DurationMs** | Pointer to **int32** |  | [optional] 
**Ok** | **bool** |  | 
**Entries** | **int32** |  | 
**ShardsFailed** | **int32** |  | 
**Complete** | **bool** |  | 
**Pages** | Pointer to **int32** |  | [optional] 
**Trigger** | **string** |  | 
**Scope** | [**V2KvLedgerLastVerifyScope**](V2KvLedgerLastVerifyScope.md) |  | 
**FirstFailure** | Pointer to [**V2KvLedgerFailure**](V2KvLedgerFailure.md) |  | [optional] 

## Methods

### NewV2KvLedgerLastVerify

`func NewV2KvLedgerLastVerify(at time.Time, startedAt time.Time, ok bool, entries int32, shardsFailed int32, complete bool, trigger string, scope V2KvLedgerLastVerifyScope, ) *V2KvLedgerLastVerify`

NewV2KvLedgerLastVerify instantiates a new V2KvLedgerLastVerify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerLastVerifyWithDefaults

`func NewV2KvLedgerLastVerifyWithDefaults() *V2KvLedgerLastVerify`

NewV2KvLedgerLastVerifyWithDefaults instantiates a new V2KvLedgerLastVerify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2KvLedgerLastVerify) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2KvLedgerLastVerify) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2KvLedgerLastVerify) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V2KvLedgerLastVerify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAt

`func (o *V2KvLedgerLastVerify) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *V2KvLedgerLastVerify) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *V2KvLedgerLastVerify) SetAt(v time.Time)`

SetAt sets At field to given value.


### GetStartedAt

`func (o *V2KvLedgerLastVerify) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V2KvLedgerLastVerify) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V2KvLedgerLastVerify) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetDurationMs

`func (o *V2KvLedgerLastVerify) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *V2KvLedgerLastVerify) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *V2KvLedgerLastVerify) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *V2KvLedgerLastVerify) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetOk

`func (o *V2KvLedgerLastVerify) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *V2KvLedgerLastVerify) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *V2KvLedgerLastVerify) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetEntries

`func (o *V2KvLedgerLastVerify) GetEntries() int32`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerLastVerify) GetEntriesOk() (*int32, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerLastVerify) SetEntries(v int32)`

SetEntries sets Entries field to given value.


### GetShardsFailed

`func (o *V2KvLedgerLastVerify) GetShardsFailed() int32`

GetShardsFailed returns the ShardsFailed field if non-nil, zero value otherwise.

### GetShardsFailedOk

`func (o *V2KvLedgerLastVerify) GetShardsFailedOk() (*int32, bool)`

GetShardsFailedOk returns a tuple with the ShardsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardsFailed

`func (o *V2KvLedgerLastVerify) SetShardsFailed(v int32)`

SetShardsFailed sets ShardsFailed field to given value.


### GetComplete

`func (o *V2KvLedgerLastVerify) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *V2KvLedgerLastVerify) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *V2KvLedgerLastVerify) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetPages

`func (o *V2KvLedgerLastVerify) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *V2KvLedgerLastVerify) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *V2KvLedgerLastVerify) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *V2KvLedgerLastVerify) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetTrigger

`func (o *V2KvLedgerLastVerify) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *V2KvLedgerLastVerify) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *V2KvLedgerLastVerify) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetScope

`func (o *V2KvLedgerLastVerify) GetScope() V2KvLedgerLastVerifyScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *V2KvLedgerLastVerify) GetScopeOk() (*V2KvLedgerLastVerifyScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *V2KvLedgerLastVerify) SetScope(v V2KvLedgerLastVerifyScope)`

SetScope sets Scope field to given value.


### GetFirstFailure

`func (o *V2KvLedgerLastVerify) GetFirstFailure() V2KvLedgerFailure`

GetFirstFailure returns the FirstFailure field if non-nil, zero value otherwise.

### GetFirstFailureOk

`func (o *V2KvLedgerLastVerify) GetFirstFailureOk() (*V2KvLedgerFailure, bool)`

GetFirstFailureOk returns a tuple with the FirstFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFailure

`func (o *V2KvLedgerLastVerify) SetFirstFailure(v V2KvLedgerFailure)`

SetFirstFailure sets FirstFailure field to given value.

### HasFirstFailure

`func (o *V2KvLedgerLastVerify) HasFirstFailure() bool`

HasFirstFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


