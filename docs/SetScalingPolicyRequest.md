# SetScalingPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | Metric to track for scaling. | 
**TargetValue** | **float64** | Target value. Percentage for CPU/Memory; req/sec per task for RPS. | 
**ScaleInCooldownSeconds** | Pointer to **int32** | Cooldown (seconds) before another scale-in can start. | [optional] [default to 300]
**ScaleOutCooldownSeconds** | Pointer to **int32** | Cooldown (seconds) before another scale-out can start. | [optional] [default to 60]

## Methods

### NewSetScalingPolicyRequest

`func NewSetScalingPolicyRequest(metric string, targetValue float64, ) *SetScalingPolicyRequest`

NewSetScalingPolicyRequest instantiates a new SetScalingPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetScalingPolicyRequestWithDefaults

`func NewSetScalingPolicyRequestWithDefaults() *SetScalingPolicyRequest`

NewSetScalingPolicyRequestWithDefaults instantiates a new SetScalingPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *SetScalingPolicyRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *SetScalingPolicyRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *SetScalingPolicyRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetTargetValue

`func (o *SetScalingPolicyRequest) GetTargetValue() float64`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *SetScalingPolicyRequest) GetTargetValueOk() (*float64, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *SetScalingPolicyRequest) SetTargetValue(v float64)`

SetTargetValue sets TargetValue field to given value.


### GetScaleInCooldownSeconds

`func (o *SetScalingPolicyRequest) GetScaleInCooldownSeconds() int32`

GetScaleInCooldownSeconds returns the ScaleInCooldownSeconds field if non-nil, zero value otherwise.

### GetScaleInCooldownSecondsOk

`func (o *SetScalingPolicyRequest) GetScaleInCooldownSecondsOk() (*int32, bool)`

GetScaleInCooldownSecondsOk returns a tuple with the ScaleInCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleInCooldownSeconds

`func (o *SetScalingPolicyRequest) SetScaleInCooldownSeconds(v int32)`

SetScaleInCooldownSeconds sets ScaleInCooldownSeconds field to given value.

### HasScaleInCooldownSeconds

`func (o *SetScalingPolicyRequest) HasScaleInCooldownSeconds() bool`

HasScaleInCooldownSeconds returns a boolean if a field has been set.

### GetScaleOutCooldownSeconds

`func (o *SetScalingPolicyRequest) GetScaleOutCooldownSeconds() int32`

GetScaleOutCooldownSeconds returns the ScaleOutCooldownSeconds field if non-nil, zero value otherwise.

### GetScaleOutCooldownSecondsOk

`func (o *SetScalingPolicyRequest) GetScaleOutCooldownSecondsOk() (*int32, bool)`

GetScaleOutCooldownSecondsOk returns a tuple with the ScaleOutCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleOutCooldownSeconds

`func (o *SetScalingPolicyRequest) SetScaleOutCooldownSeconds(v int32)`

SetScaleOutCooldownSeconds sets ScaleOutCooldownSeconds field to given value.

### HasScaleOutCooldownSeconds

`func (o *SetScalingPolicyRequest) HasScaleOutCooldownSeconds() bool`

HasScaleOutCooldownSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


