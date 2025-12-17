# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the container | 
**ImageReference** | [**ContainerImageReference**](ContainerImageReference.md) |  | 
**Cpu** | Pointer to **NullableInt32** | Container-level CPU units | [optional] 
**Memory** | Pointer to **NullableInt32** | Container-level memory hard limit (MiB) | [optional] 
**MemoryReservation** | Pointer to **NullableInt32** | Container-level memory soft limit (MiB) | [optional] 
**ExposedPorts** | Pointer to **[]int32** | List of container ports to expose | [optional] 
**MountPoints** | Pointer to [**[]ContainerMountPointsInner**](ContainerMountPointsInner.md) |  | [optional] 
**Environment** | Pointer to [**[]ContainerEnvironmentInner**](ContainerEnvironmentInner.md) | Environment variables specific to this container | [optional] 
**Secrets** | Pointer to [**[]ContainerSecretsInner**](ContainerSecretsInner.md) | Secrets mapped to environment variables | [optional] 
**HealthCheck** | Pointer to [**NullableContainerHealthCheck**](ContainerHealthCheck.md) |  | [optional] 
**DependsOn** | Pointer to [**[]ContainerDependsOnInner**](ContainerDependsOnInner.md) | Container startup dependencies | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**EntryPoint** | Pointer to **[]string** |  | [optional] 
**WorkingDirectory** | Pointer to **NullableString** |  | [optional] 
**Essential** | Pointer to **NullableBool** |  | [optional] [default to true]
**ReadonlyRootFilesystem** | Pointer to **NullableBool** |  | [optional] [default to false]
**User** | Pointer to **NullableString** |  | [optional] 
**OriginProtection** | Pointer to **NullableBool** | Enable origin protection for all exposed ports on this container. Use originProtectionConfig for advanced options like IP allow lists. | [optional] [default to false]
**OriginProtectionConfig** | Pointer to [**NullableContainerOriginProtectionConfig**](ContainerOriginProtectionConfig.md) |  | [optional] 

## Methods

### NewContainer

`func NewContainer(name string, imageReference ContainerImageReference, ) *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Container) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Container) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Container) SetName(v string)`

SetName sets Name field to given value.


### GetImageReference

`func (o *Container) GetImageReference() ContainerImageReference`

GetImageReference returns the ImageReference field if non-nil, zero value otherwise.

### GetImageReferenceOk

`func (o *Container) GetImageReferenceOk() (*ContainerImageReference, bool)`

GetImageReferenceOk returns a tuple with the ImageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageReference

`func (o *Container) SetImageReference(v ContainerImageReference)`

SetImageReference sets ImageReference field to given value.


### GetCpu

`func (o *Container) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Container) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Container) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Container) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *Container) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *Container) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *Container) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Container) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Container) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Container) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *Container) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *Container) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetMemoryReservation

`func (o *Container) GetMemoryReservation() int32`

GetMemoryReservation returns the MemoryReservation field if non-nil, zero value otherwise.

### GetMemoryReservationOk

`func (o *Container) GetMemoryReservationOk() (*int32, bool)`

GetMemoryReservationOk returns a tuple with the MemoryReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryReservation

`func (o *Container) SetMemoryReservation(v int32)`

SetMemoryReservation sets MemoryReservation field to given value.

### HasMemoryReservation

`func (o *Container) HasMemoryReservation() bool`

HasMemoryReservation returns a boolean if a field has been set.

### SetMemoryReservationNil

`func (o *Container) SetMemoryReservationNil(b bool)`

 SetMemoryReservationNil sets the value for MemoryReservation to be an explicit nil

### UnsetMemoryReservation
`func (o *Container) UnsetMemoryReservation()`

UnsetMemoryReservation ensures that no value is present for MemoryReservation, not even an explicit nil
### GetExposedPorts

`func (o *Container) GetExposedPorts() []int32`

GetExposedPorts returns the ExposedPorts field if non-nil, zero value otherwise.

### GetExposedPortsOk

`func (o *Container) GetExposedPortsOk() (*[]int32, bool)`

GetExposedPortsOk returns a tuple with the ExposedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedPorts

`func (o *Container) SetExposedPorts(v []int32)`

SetExposedPorts sets ExposedPorts field to given value.

### HasExposedPorts

`func (o *Container) HasExposedPorts() bool`

HasExposedPorts returns a boolean if a field has been set.

### SetExposedPortsNil

`func (o *Container) SetExposedPortsNil(b bool)`

 SetExposedPortsNil sets the value for ExposedPorts to be an explicit nil

### UnsetExposedPorts
`func (o *Container) UnsetExposedPorts()`

UnsetExposedPorts ensures that no value is present for ExposedPorts, not even an explicit nil
### GetMountPoints

`func (o *Container) GetMountPoints() []ContainerMountPointsInner`

GetMountPoints returns the MountPoints field if non-nil, zero value otherwise.

### GetMountPointsOk

`func (o *Container) GetMountPointsOk() (*[]ContainerMountPointsInner, bool)`

GetMountPointsOk returns a tuple with the MountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoints

`func (o *Container) SetMountPoints(v []ContainerMountPointsInner)`

SetMountPoints sets MountPoints field to given value.

### HasMountPoints

`func (o *Container) HasMountPoints() bool`

HasMountPoints returns a boolean if a field has been set.

### SetMountPointsNil

`func (o *Container) SetMountPointsNil(b bool)`

 SetMountPointsNil sets the value for MountPoints to be an explicit nil

### UnsetMountPoints
`func (o *Container) UnsetMountPoints()`

UnsetMountPoints ensures that no value is present for MountPoints, not even an explicit nil
### GetEnvironment

`func (o *Container) GetEnvironment() []ContainerEnvironmentInner`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Container) GetEnvironmentOk() (*[]ContainerEnvironmentInner, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Container) SetEnvironment(v []ContainerEnvironmentInner)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Container) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *Container) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *Container) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetSecrets

`func (o *Container) GetSecrets() []ContainerSecretsInner`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *Container) GetSecretsOk() (*[]ContainerSecretsInner, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *Container) SetSecrets(v []ContainerSecretsInner)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *Container) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### SetSecretsNil

`func (o *Container) SetSecretsNil(b bool)`

 SetSecretsNil sets the value for Secrets to be an explicit nil

### UnsetSecrets
`func (o *Container) UnsetSecrets()`

UnsetSecrets ensures that no value is present for Secrets, not even an explicit nil
### GetHealthCheck

`func (o *Container) GetHealthCheck() ContainerHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *Container) GetHealthCheckOk() (*ContainerHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *Container) SetHealthCheck(v ContainerHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *Container) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### SetHealthCheckNil

`func (o *Container) SetHealthCheckNil(b bool)`

 SetHealthCheckNil sets the value for HealthCheck to be an explicit nil

### UnsetHealthCheck
`func (o *Container) UnsetHealthCheck()`

UnsetHealthCheck ensures that no value is present for HealthCheck, not even an explicit nil
### GetDependsOn

`func (o *Container) GetDependsOn() []ContainerDependsOnInner`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *Container) GetDependsOnOk() (*[]ContainerDependsOnInner, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *Container) SetDependsOn(v []ContainerDependsOnInner)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *Container) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *Container) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *Container) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetCommand

`func (o *Container) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Container) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Container) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Container) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *Container) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *Container) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetEntryPoint

`func (o *Container) GetEntryPoint() []string`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *Container) GetEntryPointOk() (*[]string, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *Container) SetEntryPoint(v []string)`

SetEntryPoint sets EntryPoint field to given value.

### HasEntryPoint

`func (o *Container) HasEntryPoint() bool`

HasEntryPoint returns a boolean if a field has been set.

### SetEntryPointNil

`func (o *Container) SetEntryPointNil(b bool)`

 SetEntryPointNil sets the value for EntryPoint to be an explicit nil

### UnsetEntryPoint
`func (o *Container) UnsetEntryPoint()`

UnsetEntryPoint ensures that no value is present for EntryPoint, not even an explicit nil
### GetWorkingDirectory

`func (o *Container) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *Container) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *Container) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *Container) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### SetWorkingDirectoryNil

`func (o *Container) SetWorkingDirectoryNil(b bool)`

 SetWorkingDirectoryNil sets the value for WorkingDirectory to be an explicit nil

### UnsetWorkingDirectory
`func (o *Container) UnsetWorkingDirectory()`

UnsetWorkingDirectory ensures that no value is present for WorkingDirectory, not even an explicit nil
### GetEssential

`func (o *Container) GetEssential() bool`

GetEssential returns the Essential field if non-nil, zero value otherwise.

### GetEssentialOk

`func (o *Container) GetEssentialOk() (*bool, bool)`

GetEssentialOk returns a tuple with the Essential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssential

`func (o *Container) SetEssential(v bool)`

SetEssential sets Essential field to given value.

### HasEssential

`func (o *Container) HasEssential() bool`

HasEssential returns a boolean if a field has been set.

### SetEssentialNil

`func (o *Container) SetEssentialNil(b bool)`

 SetEssentialNil sets the value for Essential to be an explicit nil

### UnsetEssential
`func (o *Container) UnsetEssential()`

UnsetEssential ensures that no value is present for Essential, not even an explicit nil
### GetReadonlyRootFilesystem

`func (o *Container) GetReadonlyRootFilesystem() bool`

GetReadonlyRootFilesystem returns the ReadonlyRootFilesystem field if non-nil, zero value otherwise.

### GetReadonlyRootFilesystemOk

`func (o *Container) GetReadonlyRootFilesystemOk() (*bool, bool)`

GetReadonlyRootFilesystemOk returns a tuple with the ReadonlyRootFilesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyRootFilesystem

`func (o *Container) SetReadonlyRootFilesystem(v bool)`

SetReadonlyRootFilesystem sets ReadonlyRootFilesystem field to given value.

### HasReadonlyRootFilesystem

`func (o *Container) HasReadonlyRootFilesystem() bool`

HasReadonlyRootFilesystem returns a boolean if a field has been set.

### SetReadonlyRootFilesystemNil

`func (o *Container) SetReadonlyRootFilesystemNil(b bool)`

 SetReadonlyRootFilesystemNil sets the value for ReadonlyRootFilesystem to be an explicit nil

### UnsetReadonlyRootFilesystem
`func (o *Container) UnsetReadonlyRootFilesystem()`

UnsetReadonlyRootFilesystem ensures that no value is present for ReadonlyRootFilesystem, not even an explicit nil
### GetUser

`func (o *Container) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Container) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Container) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Container) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Container) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Container) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetOriginProtection

`func (o *Container) GetOriginProtection() bool`

GetOriginProtection returns the OriginProtection field if non-nil, zero value otherwise.

### GetOriginProtectionOk

`func (o *Container) GetOriginProtectionOk() (*bool, bool)`

GetOriginProtectionOk returns a tuple with the OriginProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtection

`func (o *Container) SetOriginProtection(v bool)`

SetOriginProtection sets OriginProtection field to given value.

### HasOriginProtection

`func (o *Container) HasOriginProtection() bool`

HasOriginProtection returns a boolean if a field has been set.

### SetOriginProtectionNil

`func (o *Container) SetOriginProtectionNil(b bool)`

 SetOriginProtectionNil sets the value for OriginProtection to be an explicit nil

### UnsetOriginProtection
`func (o *Container) UnsetOriginProtection()`

UnsetOriginProtection ensures that no value is present for OriginProtection, not even an explicit nil
### GetOriginProtectionConfig

`func (o *Container) GetOriginProtectionConfig() ContainerOriginProtectionConfig`

GetOriginProtectionConfig returns the OriginProtectionConfig field if non-nil, zero value otherwise.

### GetOriginProtectionConfigOk

`func (o *Container) GetOriginProtectionConfigOk() (*ContainerOriginProtectionConfig, bool)`

GetOriginProtectionConfigOk returns a tuple with the OriginProtectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtectionConfig

`func (o *Container) SetOriginProtectionConfig(v ContainerOriginProtectionConfig)`

SetOriginProtectionConfig sets OriginProtectionConfig field to given value.

### HasOriginProtectionConfig

`func (o *Container) HasOriginProtectionConfig() bool`

HasOriginProtectionConfig returns a boolean if a field has been set.

### SetOriginProtectionConfigNil

`func (o *Container) SetOriginProtectionConfigNil(b bool)`

 SetOriginProtectionConfigNil sets the value for OriginProtectionConfig to be an explicit nil

### UnsetOriginProtectionConfig
`func (o *Container) UnsetOriginProtectionConfig()`

UnsetOriginProtectionConfig ensures that no value is present for OriginProtectionConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


