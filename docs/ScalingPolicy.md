# ScalingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** |  | [optional] 
**TargetValue** | Pointer to **float32** |  | [optional] 
**ScaleInCooldownSeconds** | Pointer to **int32** |  | [optional] 
**ScaleOutCooldownSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewScalingPolicy

`func NewScalingPolicy() *ScalingPolicy`

NewScalingPolicy instantiates a new ScalingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalingPolicyWithDefaults

`func NewScalingPolicyWithDefaults() *ScalingPolicy`

NewScalingPolicyWithDefaults instantiates a new ScalingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *ScalingPolicy) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ScalingPolicy) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ScalingPolicy) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ScalingPolicy) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetTargetValue

`func (o *ScalingPolicy) GetTargetValue() float32`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *ScalingPolicy) GetTargetValueOk() (*float32, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *ScalingPolicy) SetTargetValue(v float32)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *ScalingPolicy) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### GetScaleInCooldownSeconds

`func (o *ScalingPolicy) GetScaleInCooldownSeconds() int32`

GetScaleInCooldownSeconds returns the ScaleInCooldownSeconds field if non-nil, zero value otherwise.

### GetScaleInCooldownSecondsOk

`func (o *ScalingPolicy) GetScaleInCooldownSecondsOk() (*int32, bool)`

GetScaleInCooldownSecondsOk returns a tuple with the ScaleInCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleInCooldownSeconds

`func (o *ScalingPolicy) SetScaleInCooldownSeconds(v int32)`

SetScaleInCooldownSeconds sets ScaleInCooldownSeconds field to given value.

### HasScaleInCooldownSeconds

`func (o *ScalingPolicy) HasScaleInCooldownSeconds() bool`

HasScaleInCooldownSeconds returns a boolean if a field has been set.

### GetScaleOutCooldownSeconds

`func (o *ScalingPolicy) GetScaleOutCooldownSeconds() int32`

GetScaleOutCooldownSeconds returns the ScaleOutCooldownSeconds field if non-nil, zero value otherwise.

### GetScaleOutCooldownSecondsOk

`func (o *ScalingPolicy) GetScaleOutCooldownSecondsOk() (*int32, bool)`

GetScaleOutCooldownSecondsOk returns a tuple with the ScaleOutCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleOutCooldownSeconds

`func (o *ScalingPolicy) SetScaleOutCooldownSeconds(v int32)`

SetScaleOutCooldownSeconds sets ScaleOutCooldownSeconds field to given value.

### HasScaleOutCooldownSeconds

`func (o *ScalingPolicy) HasScaleOutCooldownSeconds() bool`

HasScaleOutCooldownSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


