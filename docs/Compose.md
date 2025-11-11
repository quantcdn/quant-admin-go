# Compose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]Container**](Container.md) |  | [optional] 
**Architecture** | Pointer to **string** | CPU architecture (X86_64 or ARM64) | [optional] 
**TaskCpu** | Pointer to **int32** | Task-level CPU units (e.g., 256, 512, 1024) | [optional] 
**TaskMemory** | Pointer to **int32** | Task-level memory in MB | [optional] 
**MinCapacity** | Pointer to **int32** | Minimum number of instances | [optional] 
**MaxCapacity** | Pointer to **int32** | Maximum number of instances | [optional] 
**SpotConfiguration** | Pointer to [**SpotConfiguration**](SpotConfiguration.md) |  | [optional] 
**EnableCrossEnvNetworking** | Pointer to **NullableBool** | Optional. Enable cross-environment networking within the same application. When false (default): Uses shared security group for complete isolation (most secure). When true: Uses app-specific security group to enable communication between environments of the same application (e.g., staging can connect to production database). Note: If enableCrossAppNetworking is true, this setting is overridden. | [optional] [default to false]
**EnableCrossAppNetworking** | Pointer to **NullableBool** | Optional. Enable cross-application networking within the same organization. When false (default): Uses shared/app-specific security group based on enableCrossEnvNetworking. When true: Uses org-specific security group to enable container-to-container communication with ALL applications in the same organization via service discovery (microservices architecture). This setting takes priority over enableCrossEnvNetworking. | [optional] [default to false]

## Methods

### NewCompose

`func NewCompose() *Compose`

NewCompose instantiates a new Compose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeWithDefaults

`func NewComposeWithDefaults() *Compose`

NewComposeWithDefaults instantiates a new Compose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *Compose) GetContainers() []Container`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Compose) GetContainersOk() (*[]Container, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Compose) SetContainers(v []Container)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *Compose) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetArchitecture

`func (o *Compose) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Compose) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Compose) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *Compose) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetTaskCpu

`func (o *Compose) GetTaskCpu() int32`

GetTaskCpu returns the TaskCpu field if non-nil, zero value otherwise.

### GetTaskCpuOk

`func (o *Compose) GetTaskCpuOk() (*int32, bool)`

GetTaskCpuOk returns a tuple with the TaskCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCpu

`func (o *Compose) SetTaskCpu(v int32)`

SetTaskCpu sets TaskCpu field to given value.

### HasTaskCpu

`func (o *Compose) HasTaskCpu() bool`

HasTaskCpu returns a boolean if a field has been set.

### GetTaskMemory

`func (o *Compose) GetTaskMemory() int32`

GetTaskMemory returns the TaskMemory field if non-nil, zero value otherwise.

### GetTaskMemoryOk

`func (o *Compose) GetTaskMemoryOk() (*int32, bool)`

GetTaskMemoryOk returns a tuple with the TaskMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMemory

`func (o *Compose) SetTaskMemory(v int32)`

SetTaskMemory sets TaskMemory field to given value.

### HasTaskMemory

`func (o *Compose) HasTaskMemory() bool`

HasTaskMemory returns a boolean if a field has been set.

### GetMinCapacity

`func (o *Compose) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *Compose) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *Compose) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *Compose) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *Compose) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *Compose) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *Compose) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *Compose) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetSpotConfiguration

`func (o *Compose) GetSpotConfiguration() SpotConfiguration`

GetSpotConfiguration returns the SpotConfiguration field if non-nil, zero value otherwise.

### GetSpotConfigurationOk

`func (o *Compose) GetSpotConfigurationOk() (*SpotConfiguration, bool)`

GetSpotConfigurationOk returns a tuple with the SpotConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotConfiguration

`func (o *Compose) SetSpotConfiguration(v SpotConfiguration)`

SetSpotConfiguration sets SpotConfiguration field to given value.

### HasSpotConfiguration

`func (o *Compose) HasSpotConfiguration() bool`

HasSpotConfiguration returns a boolean if a field has been set.

### GetEnableCrossEnvNetworking

`func (o *Compose) GetEnableCrossEnvNetworking() bool`

GetEnableCrossEnvNetworking returns the EnableCrossEnvNetworking field if non-nil, zero value otherwise.

### GetEnableCrossEnvNetworkingOk

`func (o *Compose) GetEnableCrossEnvNetworkingOk() (*bool, bool)`

GetEnableCrossEnvNetworkingOk returns a tuple with the EnableCrossEnvNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCrossEnvNetworking

`func (o *Compose) SetEnableCrossEnvNetworking(v bool)`

SetEnableCrossEnvNetworking sets EnableCrossEnvNetworking field to given value.

### HasEnableCrossEnvNetworking

`func (o *Compose) HasEnableCrossEnvNetworking() bool`

HasEnableCrossEnvNetworking returns a boolean if a field has been set.

### SetEnableCrossEnvNetworkingNil

`func (o *Compose) SetEnableCrossEnvNetworkingNil(b bool)`

 SetEnableCrossEnvNetworkingNil sets the value for EnableCrossEnvNetworking to be an explicit nil

### UnsetEnableCrossEnvNetworking
`func (o *Compose) UnsetEnableCrossEnvNetworking()`

UnsetEnableCrossEnvNetworking ensures that no value is present for EnableCrossEnvNetworking, not even an explicit nil
### GetEnableCrossAppNetworking

`func (o *Compose) GetEnableCrossAppNetworking() bool`

GetEnableCrossAppNetworking returns the EnableCrossAppNetworking field if non-nil, zero value otherwise.

### GetEnableCrossAppNetworkingOk

`func (o *Compose) GetEnableCrossAppNetworkingOk() (*bool, bool)`

GetEnableCrossAppNetworkingOk returns a tuple with the EnableCrossAppNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCrossAppNetworking

`func (o *Compose) SetEnableCrossAppNetworking(v bool)`

SetEnableCrossAppNetworking sets EnableCrossAppNetworking field to given value.

### HasEnableCrossAppNetworking

`func (o *Compose) HasEnableCrossAppNetworking() bool`

HasEnableCrossAppNetworking returns a boolean if a field has been set.

### SetEnableCrossAppNetworkingNil

`func (o *Compose) SetEnableCrossAppNetworkingNil(b bool)`

 SetEnableCrossAppNetworkingNil sets the value for EnableCrossAppNetworking to be an explicit nil

### UnsetEnableCrossAppNetworking
`func (o *Compose) UnsetEnableCrossAppNetworking()`

UnsetEnableCrossAppNetworking ensures that no value is present for EnableCrossAppNetworking, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


