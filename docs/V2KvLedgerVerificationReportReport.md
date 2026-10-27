# V2KvLedgerVerificationReportReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V** | **int32** |  | 
**Type** | **string** |  | 
**Org** | **string** |  | 
**Store** | **string** |  | 
**Ledger** | **string** |  | 
**Id** | **string** |  | 
**At** | **time.Time** |  | 
**StartedAt** | **time.Time** |  | 
**Trigger** | **string** |  | 
**Scope** | **map[string]interface{}** |  | 
**Complete** | **bool** |  | 
**Ok** | **bool** |  | 
**Entries** | **int32** |  | 
**ShardsFailed** | **int32** |  | 
**Failures** | [**[]V2KvLedgerFailure**](V2KvLedgerFailure.md) |  | 

## Methods

### NewV2KvLedgerVerificationReportReport

`func NewV2KvLedgerVerificationReportReport(v int32, type_ string, org string, store string, ledger string, id string, at time.Time, startedAt time.Time, trigger string, scope map[string]interface{}, complete bool, ok bool, entries int32, shardsFailed int32, failures []V2KvLedgerFailure, ) *V2KvLedgerVerificationReportReport`

NewV2KvLedgerVerificationReportReport instantiates a new V2KvLedgerVerificationReportReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerVerificationReportReportWithDefaults

`func NewV2KvLedgerVerificationReportReportWithDefaults() *V2KvLedgerVerificationReportReport`

NewV2KvLedgerVerificationReportReportWithDefaults instantiates a new V2KvLedgerVerificationReportReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV

`func (o *V2KvLedgerVerificationReportReport) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *V2KvLedgerVerificationReportReport) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *V2KvLedgerVerificationReportReport) SetV(v int32)`

SetV sets V field to given value.


### GetType

`func (o *V2KvLedgerVerificationReportReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2KvLedgerVerificationReportReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2KvLedgerVerificationReportReport) SetType(v string)`

SetType sets Type field to given value.


### GetOrg

`func (o *V2KvLedgerVerificationReportReport) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *V2KvLedgerVerificationReportReport) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *V2KvLedgerVerificationReportReport) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetStore

`func (o *V2KvLedgerVerificationReportReport) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *V2KvLedgerVerificationReportReport) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *V2KvLedgerVerificationReportReport) SetStore(v string)`

SetStore sets Store field to given value.


### GetLedger

`func (o *V2KvLedgerVerificationReportReport) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerVerificationReportReport) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerVerificationReportReport) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### GetId

`func (o *V2KvLedgerVerificationReportReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2KvLedgerVerificationReportReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2KvLedgerVerificationReportReport) SetId(v string)`

SetId sets Id field to given value.


### GetAt

`func (o *V2KvLedgerVerificationReportReport) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *V2KvLedgerVerificationReportReport) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *V2KvLedgerVerificationReportReport) SetAt(v time.Time)`

SetAt sets At field to given value.


### GetStartedAt

`func (o *V2KvLedgerVerificationReportReport) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V2KvLedgerVerificationReportReport) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V2KvLedgerVerificationReportReport) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetTrigger

`func (o *V2KvLedgerVerificationReportReport) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *V2KvLedgerVerificationReportReport) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *V2KvLedgerVerificationReportReport) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetScope

`func (o *V2KvLedgerVerificationReportReport) GetScope() map[string]interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *V2KvLedgerVerificationReportReport) GetScopeOk() (*map[string]interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *V2KvLedgerVerificationReportReport) SetScope(v map[string]interface{})`

SetScope sets Scope field to given value.


### GetComplete

`func (o *V2KvLedgerVerificationReportReport) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *V2KvLedgerVerificationReportReport) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *V2KvLedgerVerificationReportReport) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetOk

`func (o *V2KvLedgerVerificationReportReport) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *V2KvLedgerVerificationReportReport) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *V2KvLedgerVerificationReportReport) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetEntries

`func (o *V2KvLedgerVerificationReportReport) GetEntries() int32`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerVerificationReportReport) GetEntriesOk() (*int32, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerVerificationReportReport) SetEntries(v int32)`

SetEntries sets Entries field to given value.


### GetShardsFailed

`func (o *V2KvLedgerVerificationReportReport) GetShardsFailed() int32`

GetShardsFailed returns the ShardsFailed field if non-nil, zero value otherwise.

### GetShardsFailedOk

`func (o *V2KvLedgerVerificationReportReport) GetShardsFailedOk() (*int32, bool)`

GetShardsFailedOk returns a tuple with the ShardsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardsFailed

`func (o *V2KvLedgerVerificationReportReport) SetShardsFailed(v int32)`

SetShardsFailed sets ShardsFailed field to given value.


### GetFailures

`func (o *V2KvLedgerVerificationReportReport) GetFailures() []V2KvLedgerFailure`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *V2KvLedgerVerificationReportReport) GetFailuresOk() (*[]V2KvLedgerFailure, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *V2KvLedgerVerificationReportReport) SetFailures(v []V2KvLedgerFailure)`

SetFailures sets Failures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


