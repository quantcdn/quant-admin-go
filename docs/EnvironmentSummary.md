# EnvironmentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvName** | Pointer to **string** | Environment name | [optional] 
**Status** | Pointer to **string** | Environment status | [optional] 
**DeploymentStatus** | Pointer to **string** | Current deployment status | [optional] 
**RunningCount** | Pointer to **int32** | Number of running tasks | [optional] 
**DesiredCount** | Pointer to **int32** | Desired number of tasks | [optional] 
**MinCapacity** | Pointer to **int32** | Minimum capacity for autoscaling | [optional] 
**MaxCapacity** | Pointer to **int32** | Maximum capacity for autoscaling | [optional] 

## Methods

### NewEnvironmentSummary

`func NewEnvironmentSummary() *EnvironmentSummary`

NewEnvironmentSummary instantiates a new EnvironmentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSummaryWithDefaults

`func NewEnvironmentSummaryWithDefaults() *EnvironmentSummary`

NewEnvironmentSummaryWithDefaults instantiates a new EnvironmentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvName

`func (o *EnvironmentSummary) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *EnvironmentSummary) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *EnvironmentSummary) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.

### HasEnvName

`func (o *EnvironmentSummary) HasEnvName() bool`

HasEnvName returns a boolean if a field has been set.

### GetStatus

`func (o *EnvironmentSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeploymentStatus

`func (o *EnvironmentSummary) GetDeploymentStatus() string`

GetDeploymentStatus returns the DeploymentStatus field if non-nil, zero value otherwise.

### GetDeploymentStatusOk

`func (o *EnvironmentSummary) GetDeploymentStatusOk() (*string, bool)`

GetDeploymentStatusOk returns a tuple with the DeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentStatus

`func (o *EnvironmentSummary) SetDeploymentStatus(v string)`

SetDeploymentStatus sets DeploymentStatus field to given value.

### HasDeploymentStatus

`func (o *EnvironmentSummary) HasDeploymentStatus() bool`

HasDeploymentStatus returns a boolean if a field has been set.

### GetRunningCount

`func (o *EnvironmentSummary) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *EnvironmentSummary) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *EnvironmentSummary) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *EnvironmentSummary) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetDesiredCount

`func (o *EnvironmentSummary) GetDesiredCount() int32`

GetDesiredCount returns the DesiredCount field if non-nil, zero value otherwise.

### GetDesiredCountOk

`func (o *EnvironmentSummary) GetDesiredCountOk() (*int32, bool)`

GetDesiredCountOk returns a tuple with the DesiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCount

`func (o *EnvironmentSummary) SetDesiredCount(v int32)`

SetDesiredCount sets DesiredCount field to given value.

### HasDesiredCount

`func (o *EnvironmentSummary) HasDesiredCount() bool`

HasDesiredCount returns a boolean if a field has been set.

### GetMinCapacity

`func (o *EnvironmentSummary) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *EnvironmentSummary) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *EnvironmentSummary) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *EnvironmentSummary) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *EnvironmentSummary) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *EnvironmentSummary) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *EnvironmentSummary) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *EnvironmentSummary) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


