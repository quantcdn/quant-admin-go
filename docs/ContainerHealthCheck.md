# ContainerHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | The command to run to determine if the container is healthy | [optional] 
**Interval** | Pointer to **int32** | Time period (seconds) between health checks | [optional] [default to 30]
**Timeout** | Pointer to **int32** | Time period (seconds) to wait for a health check to return | [optional] [default to 5]
**Retries** | Pointer to **int32** | Number of times to retry a failed health check | [optional] [default to 3]
**StartPeriod** | Pointer to **NullableInt32** | Grace period (seconds) to ignore unhealthy checks after container starts | [optional] 

## Methods

### NewContainerHealthCheck

`func NewContainerHealthCheck() *ContainerHealthCheck`

NewContainerHealthCheck instantiates a new ContainerHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerHealthCheckWithDefaults

`func NewContainerHealthCheckWithDefaults() *ContainerHealthCheck`

NewContainerHealthCheckWithDefaults instantiates a new ContainerHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ContainerHealthCheck) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ContainerHealthCheck) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ContainerHealthCheck) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ContainerHealthCheck) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetInterval

`func (o *ContainerHealthCheck) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ContainerHealthCheck) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ContainerHealthCheck) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ContainerHealthCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTimeout

`func (o *ContainerHealthCheck) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ContainerHealthCheck) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ContainerHealthCheck) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ContainerHealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRetries

`func (o *ContainerHealthCheck) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *ContainerHealthCheck) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *ContainerHealthCheck) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *ContainerHealthCheck) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetStartPeriod

`func (o *ContainerHealthCheck) GetStartPeriod() int32`

GetStartPeriod returns the StartPeriod field if non-nil, zero value otherwise.

### GetStartPeriodOk

`func (o *ContainerHealthCheck) GetStartPeriodOk() (*int32, bool)`

GetStartPeriodOk returns a tuple with the StartPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPeriod

`func (o *ContainerHealthCheck) SetStartPeriod(v int32)`

SetStartPeriod sets StartPeriod field to given value.

### HasStartPeriod

`func (o *ContainerHealthCheck) HasStartPeriod() bool`

HasStartPeriod returns a boolean if a field has been set.

### SetStartPeriodNil

`func (o *ContainerHealthCheck) SetStartPeriodNil(b bool)`

 SetStartPeriodNil sets the value for StartPeriod to be an explicit nil

### UnsetStartPeriod
`func (o *ContainerHealthCheck) UnsetStartPeriod()`

UnsetStartPeriod ensures that no value is present for StartPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


