# PatchEnvironmentCompose202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** |  | [optional] 
**TaskCpu** | Pointer to **string** |  | [optional] 
**TaskMemory** | Pointer to **string** |  | [optional] 
**MinCapacity** | Pointer to **int32** |  | [optional] 
**MaxCapacity** | Pointer to **int32** |  | [optional] 
**Containers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpotConfiguration** | Pointer to [**PatchEnvironmentCompose202ResponseSpotConfiguration**](PatchEnvironmentCompose202ResponseSpotConfiguration.md) |  | [optional] 
**EnableCrossEnvNetworking** | Pointer to **bool** |  | [optional] 
**EnableCrossAppNetworking** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchEnvironmentCompose202Response

`func NewPatchEnvironmentCompose202Response() *PatchEnvironmentCompose202Response`

NewPatchEnvironmentCompose202Response instantiates a new PatchEnvironmentCompose202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchEnvironmentCompose202ResponseWithDefaults

`func NewPatchEnvironmentCompose202ResponseWithDefaults() *PatchEnvironmentCompose202Response`

NewPatchEnvironmentCompose202ResponseWithDefaults instantiates a new PatchEnvironmentCompose202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *PatchEnvironmentCompose202Response) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *PatchEnvironmentCompose202Response) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *PatchEnvironmentCompose202Response) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *PatchEnvironmentCompose202Response) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetTaskCpu

`func (o *PatchEnvironmentCompose202Response) GetTaskCpu() string`

GetTaskCpu returns the TaskCpu field if non-nil, zero value otherwise.

### GetTaskCpuOk

`func (o *PatchEnvironmentCompose202Response) GetTaskCpuOk() (*string, bool)`

GetTaskCpuOk returns a tuple with the TaskCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCpu

`func (o *PatchEnvironmentCompose202Response) SetTaskCpu(v string)`

SetTaskCpu sets TaskCpu field to given value.

### HasTaskCpu

`func (o *PatchEnvironmentCompose202Response) HasTaskCpu() bool`

HasTaskCpu returns a boolean if a field has been set.

### GetTaskMemory

`func (o *PatchEnvironmentCompose202Response) GetTaskMemory() string`

GetTaskMemory returns the TaskMemory field if non-nil, zero value otherwise.

### GetTaskMemoryOk

`func (o *PatchEnvironmentCompose202Response) GetTaskMemoryOk() (*string, bool)`

GetTaskMemoryOk returns a tuple with the TaskMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMemory

`func (o *PatchEnvironmentCompose202Response) SetTaskMemory(v string)`

SetTaskMemory sets TaskMemory field to given value.

### HasTaskMemory

`func (o *PatchEnvironmentCompose202Response) HasTaskMemory() bool`

HasTaskMemory returns a boolean if a field has been set.

### GetMinCapacity

`func (o *PatchEnvironmentCompose202Response) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *PatchEnvironmentCompose202Response) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *PatchEnvironmentCompose202Response) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *PatchEnvironmentCompose202Response) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *PatchEnvironmentCompose202Response) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *PatchEnvironmentCompose202Response) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *PatchEnvironmentCompose202Response) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *PatchEnvironmentCompose202Response) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetContainers

`func (o *PatchEnvironmentCompose202Response) GetContainers() []map[string]interface{}`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *PatchEnvironmentCompose202Response) GetContainersOk() (*[]map[string]interface{}, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *PatchEnvironmentCompose202Response) SetContainers(v []map[string]interface{})`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *PatchEnvironmentCompose202Response) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetSpotConfiguration

`func (o *PatchEnvironmentCompose202Response) GetSpotConfiguration() PatchEnvironmentCompose202ResponseSpotConfiguration`

GetSpotConfiguration returns the SpotConfiguration field if non-nil, zero value otherwise.

### GetSpotConfigurationOk

`func (o *PatchEnvironmentCompose202Response) GetSpotConfigurationOk() (*PatchEnvironmentCompose202ResponseSpotConfiguration, bool)`

GetSpotConfigurationOk returns a tuple with the SpotConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotConfiguration

`func (o *PatchEnvironmentCompose202Response) SetSpotConfiguration(v PatchEnvironmentCompose202ResponseSpotConfiguration)`

SetSpotConfiguration sets SpotConfiguration field to given value.

### HasSpotConfiguration

`func (o *PatchEnvironmentCompose202Response) HasSpotConfiguration() bool`

HasSpotConfiguration returns a boolean if a field has been set.

### GetEnableCrossEnvNetworking

`func (o *PatchEnvironmentCompose202Response) GetEnableCrossEnvNetworking() bool`

GetEnableCrossEnvNetworking returns the EnableCrossEnvNetworking field if non-nil, zero value otherwise.

### GetEnableCrossEnvNetworkingOk

`func (o *PatchEnvironmentCompose202Response) GetEnableCrossEnvNetworkingOk() (*bool, bool)`

GetEnableCrossEnvNetworkingOk returns a tuple with the EnableCrossEnvNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCrossEnvNetworking

`func (o *PatchEnvironmentCompose202Response) SetEnableCrossEnvNetworking(v bool)`

SetEnableCrossEnvNetworking sets EnableCrossEnvNetworking field to given value.

### HasEnableCrossEnvNetworking

`func (o *PatchEnvironmentCompose202Response) HasEnableCrossEnvNetworking() bool`

HasEnableCrossEnvNetworking returns a boolean if a field has been set.

### GetEnableCrossAppNetworking

`func (o *PatchEnvironmentCompose202Response) GetEnableCrossAppNetworking() bool`

GetEnableCrossAppNetworking returns the EnableCrossAppNetworking field if non-nil, zero value otherwise.

### GetEnableCrossAppNetworkingOk

`func (o *PatchEnvironmentCompose202Response) GetEnableCrossAppNetworkingOk() (*bool, bool)`

GetEnableCrossAppNetworkingOk returns a tuple with the EnableCrossAppNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCrossAppNetworking

`func (o *PatchEnvironmentCompose202Response) SetEnableCrossAppNetworking(v bool)`

SetEnableCrossAppNetworking sets EnableCrossAppNetworking field to given value.

### HasEnableCrossAppNetworking

`func (o *PatchEnvironmentCompose202Response) HasEnableCrossAppNetworking() bool`

HasEnableCrossAppNetworking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


