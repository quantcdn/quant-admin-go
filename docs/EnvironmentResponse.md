# EnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvName** | **string** | Environment name | 
**Status** | Pointer to **string** | Environment status | [optional] 
**RunningCount** | Pointer to **int32** | Number of running tasks | [optional] 
**DesiredCount** | Pointer to **int32** | Desired number of tasks | [optional] 
**MinCapacity** | Pointer to **int32** | Minimum capacity for autoscaling | [optional] 
**MaxCapacity** | Pointer to **int32** | Maximum capacity for autoscaling | [optional] 
**PublicIpAddress** | Pointer to **NullableString** | Public IP address for SSH access | [optional] 
**DeploymentStatus** | Pointer to **string** | Current deployment status. FAILED indicates the most recent deployment did not complete successfully. | [optional] 
**DeploymentFailureType** | Pointer to **NullableString** | Type of deployment failure when deploymentStatus is FAILED (e.g., &#39;ECS_DEPLOYMENT_CIRCUIT_BREAKER&#39;, &#39;IMAGE_PULL_ERROR&#39;) | [optional] 
**DeploymentFailureReason** | Pointer to **NullableString** | Human-readable explanation of why the deployment failed. Contains details such as wrong image architecture, missing image, or container startup errors. | [optional] 
**TaskDefinition** | Pointer to **map[string]interface{}** | ECS task definition details | [optional] 
**Service** | Pointer to **map[string]interface{}** | ECS service details | [optional] 
**LoadBalancer** | Pointer to **map[string]interface{}** | Load balancer configuration | [optional] 
**SecurityGroup** | Pointer to **map[string]interface{}** | Security group configuration | [optional] 
**Subnet** | Pointer to **map[string]interface{}** | Subnet configuration | [optional] 
**Vpc** | Pointer to **map[string]interface{}** | VPC configuration | [optional] 
**Containers** | Pointer to **[]map[string]interface{}** | Container configurations | [optional] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) | Persistent storage volumes | [optional] 
**Cron** | Pointer to [**[]Cron**](Cron.md) | Scheduled cron jobs | [optional] 
**AlbRouting** | Pointer to **map[string]interface{}** | ALB routing configuration | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update timestamp | [optional] 

## Methods

### NewEnvironmentResponse

`func NewEnvironmentResponse(envName string, ) *EnvironmentResponse`

NewEnvironmentResponse instantiates a new EnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResponseWithDefaults

`func NewEnvironmentResponseWithDefaults() *EnvironmentResponse`

NewEnvironmentResponseWithDefaults instantiates a new EnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvName

`func (o *EnvironmentResponse) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *EnvironmentResponse) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *EnvironmentResponse) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.


### GetStatus

`func (o *EnvironmentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRunningCount

`func (o *EnvironmentResponse) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *EnvironmentResponse) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *EnvironmentResponse) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *EnvironmentResponse) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetDesiredCount

`func (o *EnvironmentResponse) GetDesiredCount() int32`

GetDesiredCount returns the DesiredCount field if non-nil, zero value otherwise.

### GetDesiredCountOk

`func (o *EnvironmentResponse) GetDesiredCountOk() (*int32, bool)`

GetDesiredCountOk returns a tuple with the DesiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCount

`func (o *EnvironmentResponse) SetDesiredCount(v int32)`

SetDesiredCount sets DesiredCount field to given value.

### HasDesiredCount

`func (o *EnvironmentResponse) HasDesiredCount() bool`

HasDesiredCount returns a boolean if a field has been set.

### GetMinCapacity

`func (o *EnvironmentResponse) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *EnvironmentResponse) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *EnvironmentResponse) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *EnvironmentResponse) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *EnvironmentResponse) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *EnvironmentResponse) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *EnvironmentResponse) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *EnvironmentResponse) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *EnvironmentResponse) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *EnvironmentResponse) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *EnvironmentResponse) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *EnvironmentResponse) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddressNil

`func (o *EnvironmentResponse) SetPublicIpAddressNil(b bool)`

 SetPublicIpAddressNil sets the value for PublicIpAddress to be an explicit nil

### UnsetPublicIpAddress
`func (o *EnvironmentResponse) UnsetPublicIpAddress()`

UnsetPublicIpAddress ensures that no value is present for PublicIpAddress, not even an explicit nil
### GetDeploymentStatus

`func (o *EnvironmentResponse) GetDeploymentStatus() string`

GetDeploymentStatus returns the DeploymentStatus field if non-nil, zero value otherwise.

### GetDeploymentStatusOk

`func (o *EnvironmentResponse) GetDeploymentStatusOk() (*string, bool)`

GetDeploymentStatusOk returns a tuple with the DeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentStatus

`func (o *EnvironmentResponse) SetDeploymentStatus(v string)`

SetDeploymentStatus sets DeploymentStatus field to given value.

### HasDeploymentStatus

`func (o *EnvironmentResponse) HasDeploymentStatus() bool`

HasDeploymentStatus returns a boolean if a field has been set.

### GetDeploymentFailureType

`func (o *EnvironmentResponse) GetDeploymentFailureType() string`

GetDeploymentFailureType returns the DeploymentFailureType field if non-nil, zero value otherwise.

### GetDeploymentFailureTypeOk

`func (o *EnvironmentResponse) GetDeploymentFailureTypeOk() (*string, bool)`

GetDeploymentFailureTypeOk returns a tuple with the DeploymentFailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentFailureType

`func (o *EnvironmentResponse) SetDeploymentFailureType(v string)`

SetDeploymentFailureType sets DeploymentFailureType field to given value.

### HasDeploymentFailureType

`func (o *EnvironmentResponse) HasDeploymentFailureType() bool`

HasDeploymentFailureType returns a boolean if a field has been set.

### SetDeploymentFailureTypeNil

`func (o *EnvironmentResponse) SetDeploymentFailureTypeNil(b bool)`

 SetDeploymentFailureTypeNil sets the value for DeploymentFailureType to be an explicit nil

### UnsetDeploymentFailureType
`func (o *EnvironmentResponse) UnsetDeploymentFailureType()`

UnsetDeploymentFailureType ensures that no value is present for DeploymentFailureType, not even an explicit nil
### GetDeploymentFailureReason

`func (o *EnvironmentResponse) GetDeploymentFailureReason() string`

GetDeploymentFailureReason returns the DeploymentFailureReason field if non-nil, zero value otherwise.

### GetDeploymentFailureReasonOk

`func (o *EnvironmentResponse) GetDeploymentFailureReasonOk() (*string, bool)`

GetDeploymentFailureReasonOk returns a tuple with the DeploymentFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentFailureReason

`func (o *EnvironmentResponse) SetDeploymentFailureReason(v string)`

SetDeploymentFailureReason sets DeploymentFailureReason field to given value.

### HasDeploymentFailureReason

`func (o *EnvironmentResponse) HasDeploymentFailureReason() bool`

HasDeploymentFailureReason returns a boolean if a field has been set.

### SetDeploymentFailureReasonNil

`func (o *EnvironmentResponse) SetDeploymentFailureReasonNil(b bool)`

 SetDeploymentFailureReasonNil sets the value for DeploymentFailureReason to be an explicit nil

### UnsetDeploymentFailureReason
`func (o *EnvironmentResponse) UnsetDeploymentFailureReason()`

UnsetDeploymentFailureReason ensures that no value is present for DeploymentFailureReason, not even an explicit nil
### GetTaskDefinition

`func (o *EnvironmentResponse) GetTaskDefinition() map[string]interface{}`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *EnvironmentResponse) GetTaskDefinitionOk() (*map[string]interface{}, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *EnvironmentResponse) SetTaskDefinition(v map[string]interface{})`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *EnvironmentResponse) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.

### GetService

`func (o *EnvironmentResponse) GetService() map[string]interface{}`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *EnvironmentResponse) GetServiceOk() (*map[string]interface{}, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *EnvironmentResponse) SetService(v map[string]interface{})`

SetService sets Service field to given value.

### HasService

`func (o *EnvironmentResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *EnvironmentResponse) GetLoadBalancer() map[string]interface{}`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *EnvironmentResponse) GetLoadBalancerOk() (*map[string]interface{}, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *EnvironmentResponse) SetLoadBalancer(v map[string]interface{})`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *EnvironmentResponse) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *EnvironmentResponse) GetSecurityGroup() map[string]interface{}`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *EnvironmentResponse) GetSecurityGroupOk() (*map[string]interface{}, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *EnvironmentResponse) SetSecurityGroup(v map[string]interface{})`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *EnvironmentResponse) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetSubnet

`func (o *EnvironmentResponse) GetSubnet() map[string]interface{}`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *EnvironmentResponse) GetSubnetOk() (*map[string]interface{}, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *EnvironmentResponse) SetSubnet(v map[string]interface{})`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *EnvironmentResponse) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVpc

`func (o *EnvironmentResponse) GetVpc() map[string]interface{}`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *EnvironmentResponse) GetVpcOk() (*map[string]interface{}, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *EnvironmentResponse) SetVpc(v map[string]interface{})`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *EnvironmentResponse) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetContainers

`func (o *EnvironmentResponse) GetContainers() []map[string]interface{}`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *EnvironmentResponse) GetContainersOk() (*[]map[string]interface{}, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *EnvironmentResponse) SetContainers(v []map[string]interface{})`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *EnvironmentResponse) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetVolumes

`func (o *EnvironmentResponse) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *EnvironmentResponse) GetVolumesOk() (*[]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *EnvironmentResponse) SetVolumes(v []Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *EnvironmentResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetCron

`func (o *EnvironmentResponse) GetCron() []Cron`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *EnvironmentResponse) GetCronOk() (*[]Cron, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *EnvironmentResponse) SetCron(v []Cron)`

SetCron sets Cron field to given value.

### HasCron

`func (o *EnvironmentResponse) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetAlbRouting

`func (o *EnvironmentResponse) GetAlbRouting() map[string]interface{}`

GetAlbRouting returns the AlbRouting field if non-nil, zero value otherwise.

### GetAlbRoutingOk

`func (o *EnvironmentResponse) GetAlbRoutingOk() (*map[string]interface{}, bool)`

GetAlbRoutingOk returns a tuple with the AlbRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbRouting

`func (o *EnvironmentResponse) SetAlbRouting(v map[string]interface{})`

SetAlbRouting sets AlbRouting field to given value.

### HasAlbRouting

`func (o *EnvironmentResponse) HasAlbRouting() bool`

HasAlbRouting returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EnvironmentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


