# V2KvLedgerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledger** | [**V2KvLedger**](V2KvLedger.md) |  | 
**Instances** | [**[]V2KvLedgerInstance**](V2KvLedgerInstance.md) |  | 
**LastDigestAt** | Pointer to **time.Time** |  | [optional] 
**RetainUntil** | Pointer to **time.Time** |  | [optional] 
**ModeRefused** | Pointer to **int32** |  | [optional] 
**DigestModeRefused** | Pointer to **int32** |  | [optional] 
**ModeRetryAfter** | Pointer to **time.Time** |  | [optional] 
**Anchor** | [**V2KvLedgerStatusAnchor**](V2KvLedgerStatusAnchor.md) |  | 
**LastVerify** | Pointer to [**V2KvLedgerLastVerify**](V2KvLedgerLastVerify.md) |  | [optional] 
**VerifyProgress** | Pointer to **map[string]interface{}** | Present while a scheduled verify runs or is paused between daily runs. | [optional] 
**Forked** | Pointer to **map[string]interface{}** | Present when the store table no longer matches its archive; appends are refused with LEDGER_FORKED. | [optional] 
**NextCursor** | **NullableString** | Pass as cursor for the next page of instances; null on the last page. | 

## Methods

### NewV2KvLedgerStatus

`func NewV2KvLedgerStatus(ledger V2KvLedger, instances []V2KvLedgerInstance, anchor V2KvLedgerStatusAnchor, nextCursor NullableString, ) *V2KvLedgerStatus`

NewV2KvLedgerStatus instantiates a new V2KvLedgerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerStatusWithDefaults

`func NewV2KvLedgerStatusWithDefaults() *V2KvLedgerStatus`

NewV2KvLedgerStatusWithDefaults instantiates a new V2KvLedgerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedger

`func (o *V2KvLedgerStatus) GetLedger() V2KvLedger`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerStatus) GetLedgerOk() (*V2KvLedger, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerStatus) SetLedger(v V2KvLedger)`

SetLedger sets Ledger field to given value.


### GetInstances

`func (o *V2KvLedgerStatus) GetInstances() []V2KvLedgerInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *V2KvLedgerStatus) GetInstancesOk() (*[]V2KvLedgerInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *V2KvLedgerStatus) SetInstances(v []V2KvLedgerInstance)`

SetInstances sets Instances field to given value.


### GetLastDigestAt

`func (o *V2KvLedgerStatus) GetLastDigestAt() time.Time`

GetLastDigestAt returns the LastDigestAt field if non-nil, zero value otherwise.

### GetLastDigestAtOk

`func (o *V2KvLedgerStatus) GetLastDigestAtOk() (*time.Time, bool)`

GetLastDigestAtOk returns a tuple with the LastDigestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDigestAt

`func (o *V2KvLedgerStatus) SetLastDigestAt(v time.Time)`

SetLastDigestAt sets LastDigestAt field to given value.

### HasLastDigestAt

`func (o *V2KvLedgerStatus) HasLastDigestAt() bool`

HasLastDigestAt returns a boolean if a field has been set.

### GetRetainUntil

`func (o *V2KvLedgerStatus) GetRetainUntil() time.Time`

GetRetainUntil returns the RetainUntil field if non-nil, zero value otherwise.

### GetRetainUntilOk

`func (o *V2KvLedgerStatus) GetRetainUntilOk() (*time.Time, bool)`

GetRetainUntilOk returns a tuple with the RetainUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntil

`func (o *V2KvLedgerStatus) SetRetainUntil(v time.Time)`

SetRetainUntil sets RetainUntil field to given value.

### HasRetainUntil

`func (o *V2KvLedgerStatus) HasRetainUntil() bool`

HasRetainUntil returns a boolean if a field has been set.

### GetModeRefused

`func (o *V2KvLedgerStatus) GetModeRefused() int32`

GetModeRefused returns the ModeRefused field if non-nil, zero value otherwise.

### GetModeRefusedOk

`func (o *V2KvLedgerStatus) GetModeRefusedOk() (*int32, bool)`

GetModeRefusedOk returns a tuple with the ModeRefused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeRefused

`func (o *V2KvLedgerStatus) SetModeRefused(v int32)`

SetModeRefused sets ModeRefused field to given value.

### HasModeRefused

`func (o *V2KvLedgerStatus) HasModeRefused() bool`

HasModeRefused returns a boolean if a field has been set.

### GetDigestModeRefused

`func (o *V2KvLedgerStatus) GetDigestModeRefused() int32`

GetDigestModeRefused returns the DigestModeRefused field if non-nil, zero value otherwise.

### GetDigestModeRefusedOk

`func (o *V2KvLedgerStatus) GetDigestModeRefusedOk() (*int32, bool)`

GetDigestModeRefusedOk returns a tuple with the DigestModeRefused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestModeRefused

`func (o *V2KvLedgerStatus) SetDigestModeRefused(v int32)`

SetDigestModeRefused sets DigestModeRefused field to given value.

### HasDigestModeRefused

`func (o *V2KvLedgerStatus) HasDigestModeRefused() bool`

HasDigestModeRefused returns a boolean if a field has been set.

### GetModeRetryAfter

`func (o *V2KvLedgerStatus) GetModeRetryAfter() time.Time`

GetModeRetryAfter returns the ModeRetryAfter field if non-nil, zero value otherwise.

### GetModeRetryAfterOk

`func (o *V2KvLedgerStatus) GetModeRetryAfterOk() (*time.Time, bool)`

GetModeRetryAfterOk returns a tuple with the ModeRetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeRetryAfter

`func (o *V2KvLedgerStatus) SetModeRetryAfter(v time.Time)`

SetModeRetryAfter sets ModeRetryAfter field to given value.

### HasModeRetryAfter

`func (o *V2KvLedgerStatus) HasModeRetryAfter() bool`

HasModeRetryAfter returns a boolean if a field has been set.

### GetAnchor

`func (o *V2KvLedgerStatus) GetAnchor() V2KvLedgerStatusAnchor`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *V2KvLedgerStatus) GetAnchorOk() (*V2KvLedgerStatusAnchor, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *V2KvLedgerStatus) SetAnchor(v V2KvLedgerStatusAnchor)`

SetAnchor sets Anchor field to given value.


### GetLastVerify

`func (o *V2KvLedgerStatus) GetLastVerify() V2KvLedgerLastVerify`

GetLastVerify returns the LastVerify field if non-nil, zero value otherwise.

### GetLastVerifyOk

`func (o *V2KvLedgerStatus) GetLastVerifyOk() (*V2KvLedgerLastVerify, bool)`

GetLastVerifyOk returns a tuple with the LastVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerify

`func (o *V2KvLedgerStatus) SetLastVerify(v V2KvLedgerLastVerify)`

SetLastVerify sets LastVerify field to given value.

### HasLastVerify

`func (o *V2KvLedgerStatus) HasLastVerify() bool`

HasLastVerify returns a boolean if a field has been set.

### GetVerifyProgress

`func (o *V2KvLedgerStatus) GetVerifyProgress() map[string]interface{}`

GetVerifyProgress returns the VerifyProgress field if non-nil, zero value otherwise.

### GetVerifyProgressOk

`func (o *V2KvLedgerStatus) GetVerifyProgressOk() (*map[string]interface{}, bool)`

GetVerifyProgressOk returns a tuple with the VerifyProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyProgress

`func (o *V2KvLedgerStatus) SetVerifyProgress(v map[string]interface{})`

SetVerifyProgress sets VerifyProgress field to given value.

### HasVerifyProgress

`func (o *V2KvLedgerStatus) HasVerifyProgress() bool`

HasVerifyProgress returns a boolean if a field has been set.

### GetForked

`func (o *V2KvLedgerStatus) GetForked() map[string]interface{}`

GetForked returns the Forked field if non-nil, zero value otherwise.

### GetForkedOk

`func (o *V2KvLedgerStatus) GetForkedOk() (*map[string]interface{}, bool)`

GetForkedOk returns a tuple with the Forked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForked

`func (o *V2KvLedgerStatus) SetForked(v map[string]interface{})`

SetForked sets Forked field to given value.

### HasForked

`func (o *V2KvLedgerStatus) HasForked() bool`

HasForked returns a boolean if a field has been set.

### GetNextCursor

`func (o *V2KvLedgerStatus) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V2KvLedgerStatus) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V2KvLedgerStatus) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *V2KvLedgerStatus) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *V2KvLedgerStatus) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


