# V2KvUpgradeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The current upgrade phase. | 
**DryRun** | **bool** |  | 
**V1Table** | **string** |  | 
**V2Table** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**FrozenAt** | Pointer to **time.Time** | When writes were last frozen. | [optional] 
**SwitchedAt** | Pointer to **time.Time** |  | [optional] 
**RetainUntil** | Pointer to **time.Time** | The format 1 table is kept until this time (14 days after the switch); finalize is allowed after it. | [optional] 
**Total** | Pointer to **int32** | Items counted by preflight. | [optional] 
**Copied** | Pointer to **int32** | Items copied so far. | [optional] 
**InvalidKeys** | Pointer to **[]string** | Up to 20 keys that are not valid in format 2 (preflight). | [optional] 
**OversizeKeys** | Pointer to **[]string** | Up to 20 keys whose items would exceed the 400 KB item limit in format 2 (preflight). | [optional] 
**Error** | Pointer to **string** | Why the upgrade failed. | [optional] 

## Methods

### NewV2KvUpgradeStatus

`func NewV2KvUpgradeStatus(phase string, dryRun bool, v1Table string, v2Table string, startedAt time.Time, ) *V2KvUpgradeStatus`

NewV2KvUpgradeStatus instantiates a new V2KvUpgradeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvUpgradeStatusWithDefaults

`func NewV2KvUpgradeStatusWithDefaults() *V2KvUpgradeStatus`

NewV2KvUpgradeStatusWithDefaults instantiates a new V2KvUpgradeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *V2KvUpgradeStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V2KvUpgradeStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V2KvUpgradeStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetDryRun

`func (o *V2KvUpgradeStatus) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *V2KvUpgradeStatus) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *V2KvUpgradeStatus) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetV1Table

`func (o *V2KvUpgradeStatus) GetV1Table() string`

GetV1Table returns the V1Table field if non-nil, zero value otherwise.

### GetV1TableOk

`func (o *V2KvUpgradeStatus) GetV1TableOk() (*string, bool)`

GetV1TableOk returns a tuple with the V1Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV1Table

`func (o *V2KvUpgradeStatus) SetV1Table(v string)`

SetV1Table sets V1Table field to given value.


### GetV2Table

`func (o *V2KvUpgradeStatus) GetV2Table() string`

GetV2Table returns the V2Table field if non-nil, zero value otherwise.

### GetV2TableOk

`func (o *V2KvUpgradeStatus) GetV2TableOk() (*string, bool)`

GetV2TableOk returns a tuple with the V2Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2Table

`func (o *V2KvUpgradeStatus) SetV2Table(v string)`

SetV2Table sets V2Table field to given value.


### GetStartedAt

`func (o *V2KvUpgradeStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V2KvUpgradeStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V2KvUpgradeStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFrozenAt

`func (o *V2KvUpgradeStatus) GetFrozenAt() time.Time`

GetFrozenAt returns the FrozenAt field if non-nil, zero value otherwise.

### GetFrozenAtOk

`func (o *V2KvUpgradeStatus) GetFrozenAtOk() (*time.Time, bool)`

GetFrozenAtOk returns a tuple with the FrozenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenAt

`func (o *V2KvUpgradeStatus) SetFrozenAt(v time.Time)`

SetFrozenAt sets FrozenAt field to given value.

### HasFrozenAt

`func (o *V2KvUpgradeStatus) HasFrozenAt() bool`

HasFrozenAt returns a boolean if a field has been set.

### GetSwitchedAt

`func (o *V2KvUpgradeStatus) GetSwitchedAt() time.Time`

GetSwitchedAt returns the SwitchedAt field if non-nil, zero value otherwise.

### GetSwitchedAtOk

`func (o *V2KvUpgradeStatus) GetSwitchedAtOk() (*time.Time, bool)`

GetSwitchedAtOk returns a tuple with the SwitchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchedAt

`func (o *V2KvUpgradeStatus) SetSwitchedAt(v time.Time)`

SetSwitchedAt sets SwitchedAt field to given value.

### HasSwitchedAt

`func (o *V2KvUpgradeStatus) HasSwitchedAt() bool`

HasSwitchedAt returns a boolean if a field has been set.

### GetRetainUntil

`func (o *V2KvUpgradeStatus) GetRetainUntil() time.Time`

GetRetainUntil returns the RetainUntil field if non-nil, zero value otherwise.

### GetRetainUntilOk

`func (o *V2KvUpgradeStatus) GetRetainUntilOk() (*time.Time, bool)`

GetRetainUntilOk returns a tuple with the RetainUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntil

`func (o *V2KvUpgradeStatus) SetRetainUntil(v time.Time)`

SetRetainUntil sets RetainUntil field to given value.

### HasRetainUntil

`func (o *V2KvUpgradeStatus) HasRetainUntil() bool`

HasRetainUntil returns a boolean if a field has been set.

### GetTotal

`func (o *V2KvUpgradeStatus) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V2KvUpgradeStatus) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V2KvUpgradeStatus) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V2KvUpgradeStatus) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCopied

`func (o *V2KvUpgradeStatus) GetCopied() int32`

GetCopied returns the Copied field if non-nil, zero value otherwise.

### GetCopiedOk

`func (o *V2KvUpgradeStatus) GetCopiedOk() (*int32, bool)`

GetCopiedOk returns a tuple with the Copied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopied

`func (o *V2KvUpgradeStatus) SetCopied(v int32)`

SetCopied sets Copied field to given value.

### HasCopied

`func (o *V2KvUpgradeStatus) HasCopied() bool`

HasCopied returns a boolean if a field has been set.

### GetInvalidKeys

`func (o *V2KvUpgradeStatus) GetInvalidKeys() []string`

GetInvalidKeys returns the InvalidKeys field if non-nil, zero value otherwise.

### GetInvalidKeysOk

`func (o *V2KvUpgradeStatus) GetInvalidKeysOk() (*[]string, bool)`

GetInvalidKeysOk returns a tuple with the InvalidKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidKeys

`func (o *V2KvUpgradeStatus) SetInvalidKeys(v []string)`

SetInvalidKeys sets InvalidKeys field to given value.

### HasInvalidKeys

`func (o *V2KvUpgradeStatus) HasInvalidKeys() bool`

HasInvalidKeys returns a boolean if a field has been set.

### GetOversizeKeys

`func (o *V2KvUpgradeStatus) GetOversizeKeys() []string`

GetOversizeKeys returns the OversizeKeys field if non-nil, zero value otherwise.

### GetOversizeKeysOk

`func (o *V2KvUpgradeStatus) GetOversizeKeysOk() (*[]string, bool)`

GetOversizeKeysOk returns a tuple with the OversizeKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversizeKeys

`func (o *V2KvUpgradeStatus) SetOversizeKeys(v []string)`

SetOversizeKeys sets OversizeKeys field to given value.

### HasOversizeKeys

`func (o *V2KvUpgradeStatus) HasOversizeKeys() bool`

HasOversizeKeys returns a boolean if a field has been set.

### GetError

`func (o *V2KvUpgradeStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2KvUpgradeStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2KvUpgradeStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V2KvUpgradeStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


