# AIAutoModeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**OrchestratorModel** | Pointer to **string** | Required when enabled is true. May be left out when enabled is false. | [optional] 
**Group** | [**AIAutoModeInputGroup**](AIAutoModeInputGroup.md) |  | 
**WatcherIntervalMinutes** | **int32** |  | 
**MaxParallelTasks** | Pointer to **int32** |  | [optional] 
**Enforcement** | **string** |  | 
**AutonomousDefault** | **bool** |  | 

## Methods

### NewAIAutoModeInput

`func NewAIAutoModeInput(enabled bool, group AIAutoModeInputGroup, watcherIntervalMinutes int32, enforcement string, autonomousDefault bool, ) *AIAutoModeInput`

NewAIAutoModeInput instantiates a new AIAutoModeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAutoModeInputWithDefaults

`func NewAIAutoModeInputWithDefaults() *AIAutoModeInput`

NewAIAutoModeInputWithDefaults instantiates a new AIAutoModeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AIAutoModeInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AIAutoModeInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AIAutoModeInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOrchestratorModel

`func (o *AIAutoModeInput) GetOrchestratorModel() string`

GetOrchestratorModel returns the OrchestratorModel field if non-nil, zero value otherwise.

### GetOrchestratorModelOk

`func (o *AIAutoModeInput) GetOrchestratorModelOk() (*string, bool)`

GetOrchestratorModelOk returns a tuple with the OrchestratorModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorModel

`func (o *AIAutoModeInput) SetOrchestratorModel(v string)`

SetOrchestratorModel sets OrchestratorModel field to given value.

### HasOrchestratorModel

`func (o *AIAutoModeInput) HasOrchestratorModel() bool`

HasOrchestratorModel returns a boolean if a field has been set.

### GetGroup

`func (o *AIAutoModeInput) GetGroup() AIAutoModeInputGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AIAutoModeInput) GetGroupOk() (*AIAutoModeInputGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AIAutoModeInput) SetGroup(v AIAutoModeInputGroup)`

SetGroup sets Group field to given value.


### GetWatcherIntervalMinutes

`func (o *AIAutoModeInput) GetWatcherIntervalMinutes() int32`

GetWatcherIntervalMinutes returns the WatcherIntervalMinutes field if non-nil, zero value otherwise.

### GetWatcherIntervalMinutesOk

`func (o *AIAutoModeInput) GetWatcherIntervalMinutesOk() (*int32, bool)`

GetWatcherIntervalMinutesOk returns a tuple with the WatcherIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatcherIntervalMinutes

`func (o *AIAutoModeInput) SetWatcherIntervalMinutes(v int32)`

SetWatcherIntervalMinutes sets WatcherIntervalMinutes field to given value.


### GetMaxParallelTasks

`func (o *AIAutoModeInput) GetMaxParallelTasks() int32`

GetMaxParallelTasks returns the MaxParallelTasks field if non-nil, zero value otherwise.

### GetMaxParallelTasksOk

`func (o *AIAutoModeInput) GetMaxParallelTasksOk() (*int32, bool)`

GetMaxParallelTasksOk returns a tuple with the MaxParallelTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelTasks

`func (o *AIAutoModeInput) SetMaxParallelTasks(v int32)`

SetMaxParallelTasks sets MaxParallelTasks field to given value.

### HasMaxParallelTasks

`func (o *AIAutoModeInput) HasMaxParallelTasks() bool`

HasMaxParallelTasks returns a boolean if a field has been set.

### GetEnforcement

`func (o *AIAutoModeInput) GetEnforcement() string`

GetEnforcement returns the Enforcement field if non-nil, zero value otherwise.

### GetEnforcementOk

`func (o *AIAutoModeInput) GetEnforcementOk() (*string, bool)`

GetEnforcementOk returns a tuple with the Enforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcement

`func (o *AIAutoModeInput) SetEnforcement(v string)`

SetEnforcement sets Enforcement field to given value.


### GetAutonomousDefault

`func (o *AIAutoModeInput) GetAutonomousDefault() bool`

GetAutonomousDefault returns the AutonomousDefault field if non-nil, zero value otherwise.

### GetAutonomousDefaultOk

`func (o *AIAutoModeInput) GetAutonomousDefaultOk() (*bool, bool)`

GetAutonomousDefaultOk returns a tuple with the AutonomousDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomousDefault

`func (o *AIAutoModeInput) SetAutonomousDefault(v bool)`

SetAutonomousDefault sets AutonomousDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


