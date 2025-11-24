# PatchEnvironmentComposeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **NullableString** |  | [optional] 
**TaskCpu** | Pointer to **NullableString** |  | [optional] 
**TaskMemory** | Pointer to **NullableString** |  | [optional] 
**MinCapacity** | Pointer to **NullableInt32** |  | [optional] 
**MaxCapacity** | Pointer to **NullableInt32** |  | [optional] 
**Containers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpotConfiguration** | Pointer to [**NullablePatchEnvironmentComposeRequestSpotConfiguration**](PatchEnvironmentComposeRequestSpotConfiguration.md) |  | [optional] 
**EnableCrossEnvNetworking** | Pointer to **NullableBool** |  | [optional] 
**EnableCrossAppNetworking** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPatchEnvironmentComposeRequest

`func NewPatchEnvironmentComposeRequest() *PatchEnvironmentComposeRequest`

NewPatchEnvironmentComposeRequest instantiates a new PatchEnvironmentComposeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchEnvironmentComposeRequestWithDefaults

`func NewPatchEnvironmentComposeRequestWithDefaults() *PatchEnvironmentComposeRequest`

NewPatchEnvironmentComposeRequestWithDefaults instantiates a new PatchEnvironmentComposeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *PatchEnvironmentComposeRequest) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *PatchEnvironmentComposeRequest) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *PatchEnvironmentComposeRequest) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *PatchEnvironmentComposeRequest) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitectureNil

`func (o *PatchEnvironmentComposeRequest) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *PatchEnvironmentComposeRequest) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetTaskCpu

`func (o *PatchEnvironmentComposeRequest) GetTaskCpu() string`

GetTaskCpu returns the TaskCpu field if non-nil, zero value otherwise.

### GetTaskCpuOk

`func (o *PatchEnvironmentComposeRequest) GetTaskCpuOk() (*string, bool)`

GetTaskCpuOk returns a tuple with the TaskCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCpu

`func (o *PatchEnvironmentComposeRequest) SetTaskCpu(v string)`

SetTaskCpu sets TaskCpu field to given value.

### HasTaskCpu

`func (o *PatchEnvironmentComposeRequest) HasTaskCpu() bool`

HasTaskCpu returns a boolean if a field has been set.

### SetTaskCpuNil

`func (o *PatchEnvironmentComposeRequest) SetTaskCpuNil(b bool)`

 SetTaskCpuNil sets the value for TaskCpu to be an explicit nil

### UnsetTaskCpu
`func (o *PatchEnvironmentComposeRequest) UnsetTaskCpu()`

UnsetTaskCpu ensures that no value is present for TaskCpu, not even an explicit nil
### GetTaskMemory

`func (o *PatchEnvironmentComposeRequest) GetTaskMemory() string`

GetTaskMemory returns the TaskMemory field if non-nil, zero value otherwise.

### GetTaskMemoryOk

`func (o *PatchEnvironmentComposeRequest) GetTaskMemoryOk() (*string, bool)`

GetTaskMemoryOk returns a tuple with the TaskMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMemory

`func (o *PatchEnvironmentComposeRequest) SetTaskMemory(v string)`

SetTaskMemory sets TaskMemory field to given value.

### HasTaskMemory

`func (o *PatchEnvironmentComposeRequest) HasTaskMemory() bool`

HasTaskMemory returns a boolean if a field has been set.

### SetTaskMemoryNil

`func (o *PatchEnvironmentComposeRequest) SetTaskMemoryNil(b bool)`

 SetTaskMemoryNil sets the value for TaskMemory to be an explicit nil

### UnsetTaskMemory
`func (o *PatchEnvironmentComposeRequest) UnsetTaskMemory()`

UnsetTaskMemory ensures that no value is present for TaskMemory, not even an explicit nil
### GetMinCapacity

`func (o *PatchEnvironmentComposeRequest) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *PatchEnvironmentComposeRequest) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *PatchEnvironmentComposeRequest) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *PatchEnvironmentComposeRequest) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### SetMinCapacityNil

`func (o *PatchEnvironmentComposeRequest) SetMinCapacityNil(b bool)`

 SetMinCapacityNil sets the value for MinCapacity to be an explicit nil

### UnsetMinCapacity
`func (o *PatchEnvironmentComposeRequest) UnsetMinCapacity()`

UnsetMinCapacity ensures that no value is present for MinCapacity, not even an explicit nil
### GetMaxCapacity

`func (o *PatchEnvironmentComposeRequest) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *PatchEnvironmentComposeRequest) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *PatchEnvironmentComposeRequest) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *PatchEnvironmentComposeRequest) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### SetMaxCapacityNil

`func (o *PatchEnvironmentComposeRequest) SetMaxCapacityNil(b bool)`

 SetMaxCapacityNil sets the value for MaxCapacity to be an explicit nil

### UnsetMaxCapacity
`func (o *PatchEnvironmentComposeRequest) UnsetMaxCapacity()`

UnsetMaxCapacity ensures that no value is present for MaxCapacity, not even an explicit nil
### GetContainers

`func (o *PatchEnvironmentComposeRequest) GetContainers() []map[string]interface{}`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *PatchEnvironmentComposeRequest) GetContainersOk() (*[]map[string]interface{}, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *PatchEnvironmentComposeRequest) SetContainers(v []map[string]interface{})`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *PatchEnvironmentComposeRequest) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### SetContainersNil

`func (o *PatchEnvironmentComposeRequest) SetContainersNil(b bool)`

 SetContainersNil sets the value for Containers to be an explicit nil

### UnsetContainers
`func (o *PatchEnvironmentComposeRequest) UnsetContainers()`

UnsetContainers ensures that no value is present for Containers, not even an explicit nil
### GetSpotConfiguration

`func (o *PatchEnvironmentComposeRequest) GetSpotConfiguration() PatchEnvironmentComposeRequestSpotConfiguration`

GetSpotConfiguration returns the SpotConfiguration field if non-nil, zero value otherwise.

### GetSpotConfigurationOk

`func (o *PatchEnvironmentComposeRequest) GetSpotConfigurationOk() (*PatchEnvironmentComposeRequestSpotConfiguration, bool)`

GetSpotConfigurationOk returns a tuple with the SpotConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotConfiguration

`func (o *PatchEnvironmentComposeRequest) SetSpotConfiguration(v PatchEnvironmentComposeRequestSpotConfiguration)`

SetSpotConfiguration sets SpotConfiguration field to given value.

### HasSpotConfiguration

`func (o *PatchEnvironmentComposeRequest) HasSpotConfiguration() bool`

HasSpotConfiguration returns a boolean if a field has been set.

### SetSpotConfigurationNil

`func (o *PatchEnvironmentComposeRequest) SetSpotConfigurationNil(b bool)`

 SetSpotConfigurationNil sets the value for SpotConfiguration to be an explicit nil

### UnsetSpotConfiguration
`func (o *PatchEnvironmentComposeRequest) UnsetSpotConfiguration()`

UnsetSpotConfiguration ensures that no value is present for SpotConfiguration, not even an explicit nil
### GetEnableCrossEnvNetworking

`func (o *PatchEnvironmentComposeRequest) GetEnableCrossEnvNetworking() bool`

GetEnableCrossEnvNetworking returns the EnableCrossEnvNetworking field if non-nil, zero value otherwise.

### GetEnableCrossEnvNetworkingOk

`func (o *PatchEnvironmentComposeRequest) GetEnableCrossEnvNetworkingOk() (*bool, bool)`

GetEnableCrossEnvNetworkingOk returns a tuple with the EnableCrossEnvNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCrossEnvNetworking

`func (o *PatchEnvironmentComposeRequest) SetEnableCrossEnvNetworking(v bool)`

SetEnableCrossEnvNetworking sets EnableCrossEnvNetworking field to given value.

### HasEnableCrossEnvNetworking

`func (o *PatchEnvironmentComposeRequest) HasEnableCrossEnvNetworking() bool`

HasEnableCrossEnvNetworking returns a boolean if a field has been set.

### SetEnableCrossEnvNetworkingNil

`func (o *PatchEnvironmentComposeRequest) SetEnableCrossEnvNetworkingNil(b bool)`

 SetEnableCrossEnvNetworkingNil sets the value for EnableCrossEnvNetworking to be an explicit nil

### UnsetEnableCrossEnvNetworking
`func (o *PatchEnvironmentComposeRequest) UnsetEnableCrossEnvNetworking()`

UnsetEnableCrossEnvNetworking ensures that no value is present for EnableCrossEnvNetworking, not even an explicit nil
### GetEnableCrossAppNetworking

`func (o *PatchEnvironmentComposeRequest) GetEnableCrossAppNetworking() bool`

GetEnableCrossAppNetworking returns the EnableCrossAppNetworking field if non-nil, zero value otherwise.

### GetEnableCrossAppNetworkingOk

`func (o *PatchEnvironmentComposeRequest) GetEnableCrossAppNetworkingOk() (*bool, bool)`

GetEnableCrossAppNetworkingOk returns a tuple with the EnableCrossAppNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCrossAppNetworking

`func (o *PatchEnvironmentComposeRequest) SetEnableCrossAppNetworking(v bool)`

SetEnableCrossAppNetworking sets EnableCrossAppNetworking field to given value.

### HasEnableCrossAppNetworking

`func (o *PatchEnvironmentComposeRequest) HasEnableCrossAppNetworking() bool`

HasEnableCrossAppNetworking returns a boolean if a field has been set.

### SetEnableCrossAppNetworkingNil

`func (o *PatchEnvironmentComposeRequest) SetEnableCrossAppNetworkingNil(b bool)`

 SetEnableCrossAppNetworkingNil sets the value for EnableCrossAppNetworking to be an explicit nil

### UnsetEnableCrossAppNetworking
`func (o *PatchEnvironmentComposeRequest) UnsetEnableCrossAppNetworking()`

UnsetEnableCrossAppNetworking ensures that no value is present for EnableCrossAppNetworking, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


