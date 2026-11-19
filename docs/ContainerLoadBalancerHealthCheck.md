# ContainerLoadBalancerHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **NullableString** | Path the load balancer requests. Starts with /. | [optional] [default to "/"]
**IntervalSeconds** | Pointer to **NullableInt32** | Seconds between checks. | [optional] [default to 45]

## Methods

### NewContainerLoadBalancerHealthCheck

`func NewContainerLoadBalancerHealthCheck() *ContainerLoadBalancerHealthCheck`

NewContainerLoadBalancerHealthCheck instantiates a new ContainerLoadBalancerHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLoadBalancerHealthCheckWithDefaults

`func NewContainerLoadBalancerHealthCheckWithDefaults() *ContainerLoadBalancerHealthCheck`

NewContainerLoadBalancerHealthCheckWithDefaults instantiates a new ContainerLoadBalancerHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ContainerLoadBalancerHealthCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContainerLoadBalancerHealthCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContainerLoadBalancerHealthCheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ContainerLoadBalancerHealthCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ContainerLoadBalancerHealthCheck) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ContainerLoadBalancerHealthCheck) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetIntervalSeconds

`func (o *ContainerLoadBalancerHealthCheck) GetIntervalSeconds() int32`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *ContainerLoadBalancerHealthCheck) GetIntervalSecondsOk() (*int32, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *ContainerLoadBalancerHealthCheck) SetIntervalSeconds(v int32)`

SetIntervalSeconds sets IntervalSeconds field to given value.

### HasIntervalSeconds

`func (o *ContainerLoadBalancerHealthCheck) HasIntervalSeconds() bool`

HasIntervalSeconds returns a boolean if a field has been set.

### SetIntervalSecondsNil

`func (o *ContainerLoadBalancerHealthCheck) SetIntervalSecondsNil(b bool)`

 SetIntervalSecondsNil sets the value for IntervalSeconds to be an explicit nil

### UnsetIntervalSeconds
`func (o *ContainerLoadBalancerHealthCheck) UnsetIntervalSeconds()`

UnsetIntervalSeconds ensures that no value is present for IntervalSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


