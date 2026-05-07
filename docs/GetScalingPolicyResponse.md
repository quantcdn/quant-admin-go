# GetScalingPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** |  | [optional] 
**TargetValue** | Pointer to **float64** |  | [optional] 
**ScaleInCooldownSeconds** | Pointer to **int32** |  | [optional] 
**ScaleOutCooldownSeconds** | Pointer to **int32** |  | [optional] 
**PolicyName** | Pointer to **string** | Name of the underlying Application Auto Scaling policy. | [optional] 
**ResourceLabel** | Pointer to **NullableString** | ALB ResourceLabel for RPS policies (target group identifier). | [optional] 

## Methods

### NewGetScalingPolicyResponse

`func NewGetScalingPolicyResponse() *GetScalingPolicyResponse`

NewGetScalingPolicyResponse instantiates a new GetScalingPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetScalingPolicyResponseWithDefaults

`func NewGetScalingPolicyResponseWithDefaults() *GetScalingPolicyResponse`

NewGetScalingPolicyResponseWithDefaults instantiates a new GetScalingPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *GetScalingPolicyResponse) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GetScalingPolicyResponse) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GetScalingPolicyResponse) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *GetScalingPolicyResponse) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetTargetValue

`func (o *GetScalingPolicyResponse) GetTargetValue() float64`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *GetScalingPolicyResponse) GetTargetValueOk() (*float64, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *GetScalingPolicyResponse) SetTargetValue(v float64)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *GetScalingPolicyResponse) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### GetScaleInCooldownSeconds

`func (o *GetScalingPolicyResponse) GetScaleInCooldownSeconds() int32`

GetScaleInCooldownSeconds returns the ScaleInCooldownSeconds field if non-nil, zero value otherwise.

### GetScaleInCooldownSecondsOk

`func (o *GetScalingPolicyResponse) GetScaleInCooldownSecondsOk() (*int32, bool)`

GetScaleInCooldownSecondsOk returns a tuple with the ScaleInCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleInCooldownSeconds

`func (o *GetScalingPolicyResponse) SetScaleInCooldownSeconds(v int32)`

SetScaleInCooldownSeconds sets ScaleInCooldownSeconds field to given value.

### HasScaleInCooldownSeconds

`func (o *GetScalingPolicyResponse) HasScaleInCooldownSeconds() bool`

HasScaleInCooldownSeconds returns a boolean if a field has been set.

### GetScaleOutCooldownSeconds

`func (o *GetScalingPolicyResponse) GetScaleOutCooldownSeconds() int32`

GetScaleOutCooldownSeconds returns the ScaleOutCooldownSeconds field if non-nil, zero value otherwise.

### GetScaleOutCooldownSecondsOk

`func (o *GetScalingPolicyResponse) GetScaleOutCooldownSecondsOk() (*int32, bool)`

GetScaleOutCooldownSecondsOk returns a tuple with the ScaleOutCooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleOutCooldownSeconds

`func (o *GetScalingPolicyResponse) SetScaleOutCooldownSeconds(v int32)`

SetScaleOutCooldownSeconds sets ScaleOutCooldownSeconds field to given value.

### HasScaleOutCooldownSeconds

`func (o *GetScalingPolicyResponse) HasScaleOutCooldownSeconds() bool`

HasScaleOutCooldownSeconds returns a boolean if a field has been set.

### GetPolicyName

`func (o *GetScalingPolicyResponse) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *GetScalingPolicyResponse) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *GetScalingPolicyResponse) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *GetScalingPolicyResponse) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetResourceLabel

`func (o *GetScalingPolicyResponse) GetResourceLabel() string`

GetResourceLabel returns the ResourceLabel field if non-nil, zero value otherwise.

### GetResourceLabelOk

`func (o *GetScalingPolicyResponse) GetResourceLabelOk() (*string, bool)`

GetResourceLabelOk returns a tuple with the ResourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLabel

`func (o *GetScalingPolicyResponse) SetResourceLabel(v string)`

SetResourceLabel sets ResourceLabel field to given value.

### HasResourceLabel

`func (o *GetScalingPolicyResponse) HasResourceLabel() bool`

HasResourceLabel returns a boolean if a field has been set.

### SetResourceLabelNil

`func (o *GetScalingPolicyResponse) SetResourceLabelNil(b bool)`

 SetResourceLabelNil sets the value for ResourceLabel to be an explicit nil

### UnsetResourceLabel
`func (o *GetScalingPolicyResponse) UnsetResourceLabel()`

UnsetResourceLabel ensures that no value is present for ResourceLabel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


