# AIAutoModeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**OrchestratorModel** | **string** | Empty when the setting was saved off without an orchestrator. | 
**Group** | [**AIAutoModeConfigGroup**](AIAutoModeConfigGroup.md) |  | 
**WatcherIntervalMinutes** | **int32** |  | 
**MaxParallelTasks** | Pointer to **int32** |  | [optional] 
**Enforcement** | **string** |  | 
**AutonomousDefault** | **bool** |  | 
**UpdatedAt** | **string** |  | 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewAIAutoModeConfig

`func NewAIAutoModeConfig(enabled bool, orchestratorModel string, group AIAutoModeConfigGroup, watcherIntervalMinutes int32, enforcement string, autonomousDefault bool, updatedAt string, ) *AIAutoModeConfig`

NewAIAutoModeConfig instantiates a new AIAutoModeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAutoModeConfigWithDefaults

`func NewAIAutoModeConfigWithDefaults() *AIAutoModeConfig`

NewAIAutoModeConfigWithDefaults instantiates a new AIAutoModeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AIAutoModeConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AIAutoModeConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AIAutoModeConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOrchestratorModel

`func (o *AIAutoModeConfig) GetOrchestratorModel() string`

GetOrchestratorModel returns the OrchestratorModel field if non-nil, zero value otherwise.

### GetOrchestratorModelOk

`func (o *AIAutoModeConfig) GetOrchestratorModelOk() (*string, bool)`

GetOrchestratorModelOk returns a tuple with the OrchestratorModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorModel

`func (o *AIAutoModeConfig) SetOrchestratorModel(v string)`

SetOrchestratorModel sets OrchestratorModel field to given value.


### GetGroup

`func (o *AIAutoModeConfig) GetGroup() AIAutoModeConfigGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AIAutoModeConfig) GetGroupOk() (*AIAutoModeConfigGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AIAutoModeConfig) SetGroup(v AIAutoModeConfigGroup)`

SetGroup sets Group field to given value.


### GetWatcherIntervalMinutes

`func (o *AIAutoModeConfig) GetWatcherIntervalMinutes() int32`

GetWatcherIntervalMinutes returns the WatcherIntervalMinutes field if non-nil, zero value otherwise.

### GetWatcherIntervalMinutesOk

`func (o *AIAutoModeConfig) GetWatcherIntervalMinutesOk() (*int32, bool)`

GetWatcherIntervalMinutesOk returns a tuple with the WatcherIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatcherIntervalMinutes

`func (o *AIAutoModeConfig) SetWatcherIntervalMinutes(v int32)`

SetWatcherIntervalMinutes sets WatcherIntervalMinutes field to given value.


### GetMaxParallelTasks

`func (o *AIAutoModeConfig) GetMaxParallelTasks() int32`

GetMaxParallelTasks returns the MaxParallelTasks field if non-nil, zero value otherwise.

### GetMaxParallelTasksOk

`func (o *AIAutoModeConfig) GetMaxParallelTasksOk() (*int32, bool)`

GetMaxParallelTasksOk returns a tuple with the MaxParallelTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelTasks

`func (o *AIAutoModeConfig) SetMaxParallelTasks(v int32)`

SetMaxParallelTasks sets MaxParallelTasks field to given value.

### HasMaxParallelTasks

`func (o *AIAutoModeConfig) HasMaxParallelTasks() bool`

HasMaxParallelTasks returns a boolean if a field has been set.

### GetEnforcement

`func (o *AIAutoModeConfig) GetEnforcement() string`

GetEnforcement returns the Enforcement field if non-nil, zero value otherwise.

### GetEnforcementOk

`func (o *AIAutoModeConfig) GetEnforcementOk() (*string, bool)`

GetEnforcementOk returns a tuple with the Enforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcement

`func (o *AIAutoModeConfig) SetEnforcement(v string)`

SetEnforcement sets Enforcement field to given value.


### GetAutonomousDefault

`func (o *AIAutoModeConfig) GetAutonomousDefault() bool`

GetAutonomousDefault returns the AutonomousDefault field if non-nil, zero value otherwise.

### GetAutonomousDefaultOk

`func (o *AIAutoModeConfig) GetAutonomousDefaultOk() (*bool, bool)`

GetAutonomousDefaultOk returns a tuple with the AutonomousDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutonomousDefault

`func (o *AIAutoModeConfig) SetAutonomousDefault(v bool)`

SetAutonomousDefault sets AutonomousDefault field to given value.


### GetUpdatedAt

`func (o *AIAutoModeConfig) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AIAutoModeConfig) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AIAutoModeConfig) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedBy

`func (o *AIAutoModeConfig) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AIAutoModeConfig) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AIAutoModeConfig) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AIAutoModeConfig) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


