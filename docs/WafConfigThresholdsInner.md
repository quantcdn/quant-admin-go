# WafConfigThresholdsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Threshold type | [optional] 
**Rps** | Pointer to **int32** | Requests per second limit (for ip/header) | [optional] 
**Hits** | Pointer to **int32** | Hit count limit (for waf_hit_by_ip) | [optional] 
**Minutes** | Pointer to **int32** | Time window in minutes (for waf_hit_by_ip) | [optional] 
**Cooldown** | Pointer to **int32** | Cooldown period in seconds | [optional] 
**Mode** | Pointer to **string** | Threshold enforcement mode | [optional] [default to "disabled"]
**Value** | Pointer to **NullableString** | Header name (for header type) | [optional] 
**NotifySlack** | Pointer to **NullableString** | Slack webhook for this threshold | [optional] 

## Methods

### NewWafConfigThresholdsInner

`func NewWafConfigThresholdsInner() *WafConfigThresholdsInner`

NewWafConfigThresholdsInner instantiates a new WafConfigThresholdsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafConfigThresholdsInnerWithDefaults

`func NewWafConfigThresholdsInnerWithDefaults() *WafConfigThresholdsInner`

NewWafConfigThresholdsInnerWithDefaults instantiates a new WafConfigThresholdsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafConfigThresholdsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafConfigThresholdsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafConfigThresholdsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WafConfigThresholdsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRps

`func (o *WafConfigThresholdsInner) GetRps() int32`

GetRps returns the Rps field if non-nil, zero value otherwise.

### GetRpsOk

`func (o *WafConfigThresholdsInner) GetRpsOk() (*int32, bool)`

GetRpsOk returns a tuple with the Rps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRps

`func (o *WafConfigThresholdsInner) SetRps(v int32)`

SetRps sets Rps field to given value.

### HasRps

`func (o *WafConfigThresholdsInner) HasRps() bool`

HasRps returns a boolean if a field has been set.

### GetHits

`func (o *WafConfigThresholdsInner) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *WafConfigThresholdsInner) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *WafConfigThresholdsInner) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *WafConfigThresholdsInner) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMinutes

`func (o *WafConfigThresholdsInner) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *WafConfigThresholdsInner) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *WafConfigThresholdsInner) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *WafConfigThresholdsInner) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetCooldown

`func (o *WafConfigThresholdsInner) GetCooldown() int32`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *WafConfigThresholdsInner) GetCooldownOk() (*int32, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *WafConfigThresholdsInner) SetCooldown(v int32)`

SetCooldown sets Cooldown field to given value.

### HasCooldown

`func (o *WafConfigThresholdsInner) HasCooldown() bool`

HasCooldown returns a boolean if a field has been set.

### GetMode

`func (o *WafConfigThresholdsInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WafConfigThresholdsInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WafConfigThresholdsInner) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WafConfigThresholdsInner) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetValue

`func (o *WafConfigThresholdsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WafConfigThresholdsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WafConfigThresholdsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WafConfigThresholdsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WafConfigThresholdsInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WafConfigThresholdsInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetNotifySlack

`func (o *WafConfigThresholdsInner) GetNotifySlack() string`

GetNotifySlack returns the NotifySlack field if non-nil, zero value otherwise.

### GetNotifySlackOk

`func (o *WafConfigThresholdsInner) GetNotifySlackOk() (*string, bool)`

GetNotifySlackOk returns a tuple with the NotifySlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlack

`func (o *WafConfigThresholdsInner) SetNotifySlack(v string)`

SetNotifySlack sets NotifySlack field to given value.

### HasNotifySlack

`func (o *WafConfigThresholdsInner) HasNotifySlack() bool`

HasNotifySlack returns a boolean if a field has been set.

### SetNotifySlackNil

`func (o *WafConfigThresholdsInner) SetNotifySlackNil(b bool)`

 SetNotifySlackNil sets the value for NotifySlack to be an explicit nil

### UnsetNotifySlack
`func (o *WafConfigThresholdsInner) UnsetNotifySlack()`

UnsetNotifySlack ensures that no value is present for NotifySlack, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


