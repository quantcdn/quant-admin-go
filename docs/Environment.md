# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvName** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**RunningCount** | Pointer to **int32** |  | [optional] 
**DesiredCount** | Pointer to **int32** |  | [optional] 
**MinCapacity** | Pointer to **int32** |  | [optional] 
**MaxCapacity** | Pointer to **int32** |  | [optional] 
**CloneConfigurationFrom** | Pointer to **string** |  | [optional] 
**ImageSuffix** | Pointer to **string** | Image tag suffix for cloning | [optional] 
**TaskDefinition** | Pointer to **map[string]interface{}** |  | [optional] 
**Service** | Pointer to **map[string]interface{}** |  | [optional] 
**LoadBalancer** | Pointer to **map[string]interface{}** |  | [optional] 
**SecurityGroup** | Pointer to **map[string]interface{}** |  | [optional] 
**Subnet** | Pointer to **map[string]interface{}** |  | [optional] 
**Vpc** | Pointer to **map[string]interface{}** |  | [optional] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) |  | [optional] 
**Cron** | Pointer to [**[]Cron**](Cron.md) |  | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment(envName string, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvName

`func (o *Environment) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *Environment) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *Environment) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.


### GetStatus

`func (o *Environment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Environment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Environment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Environment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRunningCount

`func (o *Environment) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *Environment) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *Environment) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *Environment) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetDesiredCount

`func (o *Environment) GetDesiredCount() int32`

GetDesiredCount returns the DesiredCount field if non-nil, zero value otherwise.

### GetDesiredCountOk

`func (o *Environment) GetDesiredCountOk() (*int32, bool)`

GetDesiredCountOk returns a tuple with the DesiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCount

`func (o *Environment) SetDesiredCount(v int32)`

SetDesiredCount sets DesiredCount field to given value.

### HasDesiredCount

`func (o *Environment) HasDesiredCount() bool`

HasDesiredCount returns a boolean if a field has been set.

### GetMinCapacity

`func (o *Environment) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *Environment) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *Environment) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *Environment) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *Environment) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *Environment) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *Environment) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *Environment) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetCloneConfigurationFrom

`func (o *Environment) GetCloneConfigurationFrom() string`

GetCloneConfigurationFrom returns the CloneConfigurationFrom field if non-nil, zero value otherwise.

### GetCloneConfigurationFromOk

`func (o *Environment) GetCloneConfigurationFromOk() (*string, bool)`

GetCloneConfigurationFromOk returns a tuple with the CloneConfigurationFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneConfigurationFrom

`func (o *Environment) SetCloneConfigurationFrom(v string)`

SetCloneConfigurationFrom sets CloneConfigurationFrom field to given value.

### HasCloneConfigurationFrom

`func (o *Environment) HasCloneConfigurationFrom() bool`

HasCloneConfigurationFrom returns a boolean if a field has been set.

### GetImageSuffix

`func (o *Environment) GetImageSuffix() string`

GetImageSuffix returns the ImageSuffix field if non-nil, zero value otherwise.

### GetImageSuffixOk

`func (o *Environment) GetImageSuffixOk() (*string, bool)`

GetImageSuffixOk returns a tuple with the ImageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSuffix

`func (o *Environment) SetImageSuffix(v string)`

SetImageSuffix sets ImageSuffix field to given value.

### HasImageSuffix

`func (o *Environment) HasImageSuffix() bool`

HasImageSuffix returns a boolean if a field has been set.

### GetTaskDefinition

`func (o *Environment) GetTaskDefinition() map[string]interface{}`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *Environment) GetTaskDefinitionOk() (*map[string]interface{}, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *Environment) SetTaskDefinition(v map[string]interface{})`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *Environment) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.

### GetService

`func (o *Environment) GetService() map[string]interface{}`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Environment) GetServiceOk() (*map[string]interface{}, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Environment) SetService(v map[string]interface{})`

SetService sets Service field to given value.

### HasService

`func (o *Environment) HasService() bool`

HasService returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *Environment) GetLoadBalancer() map[string]interface{}`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *Environment) GetLoadBalancerOk() (*map[string]interface{}, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *Environment) SetLoadBalancer(v map[string]interface{})`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *Environment) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *Environment) GetSecurityGroup() map[string]interface{}`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *Environment) GetSecurityGroupOk() (*map[string]interface{}, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *Environment) SetSecurityGroup(v map[string]interface{})`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *Environment) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetSubnet

`func (o *Environment) GetSubnet() map[string]interface{}`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *Environment) GetSubnetOk() (*map[string]interface{}, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *Environment) SetSubnet(v map[string]interface{})`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *Environment) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVpc

`func (o *Environment) GetVpc() map[string]interface{}`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *Environment) GetVpcOk() (*map[string]interface{}, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *Environment) SetVpc(v map[string]interface{})`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *Environment) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetVolumes

`func (o *Environment) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Environment) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Environment) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Environment) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetCron

`func (o *Environment) GetCron() []Cron`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *Environment) GetCronOk() (*[]Cron, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *Environment) SetCron(v []Cron)`

SetCron sets Cron field to given value.

### HasCron

`func (o *Environment) HasCron() bool`

HasCron returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


