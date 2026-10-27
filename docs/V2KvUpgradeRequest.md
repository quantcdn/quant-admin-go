# V2KvUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | start: preflight, copy and switch to format 2. rollback: copy back to format 1 after the switch. finalize: delete the format 1 table after retainUntil; no rollback is possible after it. resume: queue the current step again when the phase has not changed for more than 960 seconds. | 
**DryRun** | Pointer to **bool** | start only: run preflight and stop at phase checked. Must be a JSON boolean. | [optional] [default to false]

## Methods

### NewV2KvUpgradeRequest

`func NewV2KvUpgradeRequest(action string, ) *V2KvUpgradeRequest`

NewV2KvUpgradeRequest instantiates a new V2KvUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvUpgradeRequestWithDefaults

`func NewV2KvUpgradeRequestWithDefaults() *V2KvUpgradeRequest`

NewV2KvUpgradeRequestWithDefaults instantiates a new V2KvUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *V2KvUpgradeRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V2KvUpgradeRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V2KvUpgradeRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetDryRun

`func (o *V2KvUpgradeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *V2KvUpgradeRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *V2KvUpgradeRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *V2KvUpgradeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


