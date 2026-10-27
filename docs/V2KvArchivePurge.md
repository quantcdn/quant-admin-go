# V2KvArchivePurge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurgeId** | Pointer to **string** | OpenAPI schemas and path parameters for the retained ledger archive routes of deleted stores (A5). Shapes follow the backend contract C-1 (ArchiveStatus, DryRun, Purge, PurgeReceipt, manifest). The list uses the KV page envelope (data, next_cursor, has_more, page_end). | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | COMPLETE, RETAINED, BLOCKED and CANCELLED are final. RETAINED: a lock or the quiet period was found at the gate, nothing was deleted. BLOCKED: needs an operator. | [optional] 
**RequestedBy** | Pointer to **string** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**RetainedUntil** | Pointer to **NullableTime** | On a RETAINED purge, the latest lock found. | [optional] 
**Counts** | Pointer to [**V2KvArchivePurgeCounts**](V2KvArchivePurgeCounts.md) |  | [optional] 
**LastError** | Pointer to [**NullableV2KvArchivePurgeLastError**](V2KvArchivePurgeLastError.md) |  | [optional] 

## Methods

### NewV2KvArchivePurge

`func NewV2KvArchivePurge() *V2KvArchivePurge`

NewV2KvArchivePurge instantiates a new V2KvArchivePurge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvArchivePurgeWithDefaults

`func NewV2KvArchivePurgeWithDefaults() *V2KvArchivePurge`

NewV2KvArchivePurgeWithDefaults instantiates a new V2KvArchivePurge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurgeId

`func (o *V2KvArchivePurge) GetPurgeId() string`

GetPurgeId returns the PurgeId field if non-nil, zero value otherwise.

### GetPurgeIdOk

`func (o *V2KvArchivePurge) GetPurgeIdOk() (*string, bool)`

GetPurgeIdOk returns a tuple with the PurgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeId

`func (o *V2KvArchivePurge) SetPurgeId(v string)`

SetPurgeId sets PurgeId field to given value.

### HasPurgeId

`func (o *V2KvArchivePurge) HasPurgeId() bool`

HasPurgeId returns a boolean if a field has been set.

### GetStore

`func (o *V2KvArchivePurge) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *V2KvArchivePurge) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *V2KvArchivePurge) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *V2KvArchivePurge) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetState

`func (o *V2KvArchivePurge) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvArchivePurge) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvArchivePurge) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V2KvArchivePurge) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRequestedBy

`func (o *V2KvArchivePurge) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *V2KvArchivePurge) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *V2KvArchivePurge) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *V2KvArchivePurge) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetRequestedAt

`func (o *V2KvArchivePurge) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *V2KvArchivePurge) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *V2KvArchivePurge) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *V2KvArchivePurge) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetReason

`func (o *V2KvArchivePurge) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvArchivePurge) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvArchivePurge) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V2KvArchivePurge) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *V2KvArchivePurge) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *V2KvArchivePurge) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetUpdatedAt

`func (o *V2KvArchivePurge) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2KvArchivePurge) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2KvArchivePurge) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2KvArchivePurge) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *V2KvArchivePurge) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *V2KvArchivePurge) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *V2KvArchivePurge) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *V2KvArchivePurge) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *V2KvArchivePurge) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *V2KvArchivePurge) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetRetainedUntil

`func (o *V2KvArchivePurge) GetRetainedUntil() time.Time`

GetRetainedUntil returns the RetainedUntil field if non-nil, zero value otherwise.

### GetRetainedUntilOk

`func (o *V2KvArchivePurge) GetRetainedUntilOk() (*time.Time, bool)`

GetRetainedUntilOk returns a tuple with the RetainedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedUntil

`func (o *V2KvArchivePurge) SetRetainedUntil(v time.Time)`

SetRetainedUntil sets RetainedUntil field to given value.

### HasRetainedUntil

`func (o *V2KvArchivePurge) HasRetainedUntil() bool`

HasRetainedUntil returns a boolean if a field has been set.

### SetRetainedUntilNil

`func (o *V2KvArchivePurge) SetRetainedUntilNil(b bool)`

 SetRetainedUntilNil sets the value for RetainedUntil to be an explicit nil

### UnsetRetainedUntil
`func (o *V2KvArchivePurge) UnsetRetainedUntil()`

UnsetRetainedUntil ensures that no value is present for RetainedUntil, not even an explicit nil
### GetCounts

`func (o *V2KvArchivePurge) GetCounts() V2KvArchivePurgeCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *V2KvArchivePurge) GetCountsOk() (*V2KvArchivePurgeCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *V2KvArchivePurge) SetCounts(v V2KvArchivePurgeCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *V2KvArchivePurge) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetLastError

`func (o *V2KvArchivePurge) GetLastError() V2KvArchivePurgeLastError`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *V2KvArchivePurge) GetLastErrorOk() (*V2KvArchivePurgeLastError, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *V2KvArchivePurge) SetLastError(v V2KvArchivePurgeLastError)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *V2KvArchivePurge) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *V2KvArchivePurge) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *V2KvArchivePurge) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


