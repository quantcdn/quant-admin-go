# Threshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Cooldown** | Pointer to **int32** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] [default to "disabled"]
**Rps** | Pointer to **int32** |  | [optional] 
**NotifySlack** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Hits** | Pointer to **int32** |  | [optional] 
**Minutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewThreshold

`func NewThreshold() *Threshold`

NewThreshold instantiates a new Threshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdWithDefaults

`func NewThresholdWithDefaults() *Threshold`

NewThresholdWithDefaults instantiates a new Threshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Threshold) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Threshold) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Threshold) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Threshold) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCooldown

`func (o *Threshold) GetCooldown() int32`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *Threshold) GetCooldownOk() (*int32, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *Threshold) SetCooldown(v int32)`

SetCooldown sets Cooldown field to given value.

### HasCooldown

`func (o *Threshold) HasCooldown() bool`

HasCooldown returns a boolean if a field has been set.

### GetMode

`func (o *Threshold) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Threshold) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Threshold) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Threshold) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRps

`func (o *Threshold) GetRps() int32`

GetRps returns the Rps field if non-nil, zero value otherwise.

### GetRpsOk

`func (o *Threshold) GetRpsOk() (*int32, bool)`

GetRpsOk returns a tuple with the Rps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRps

`func (o *Threshold) SetRps(v int32)`

SetRps sets Rps field to given value.

### HasRps

`func (o *Threshold) HasRps() bool`

HasRps returns a boolean if a field has been set.

### GetNotifySlack

`func (o *Threshold) GetNotifySlack() string`

GetNotifySlack returns the NotifySlack field if non-nil, zero value otherwise.

### GetNotifySlackOk

`func (o *Threshold) GetNotifySlackOk() (*string, bool)`

GetNotifySlackOk returns a tuple with the NotifySlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlack

`func (o *Threshold) SetNotifySlack(v string)`

SetNotifySlack sets NotifySlack field to given value.

### HasNotifySlack

`func (o *Threshold) HasNotifySlack() bool`

HasNotifySlack returns a boolean if a field has been set.

### GetType

`func (o *Threshold) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Threshold) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Threshold) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Threshold) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHits

`func (o *Threshold) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *Threshold) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *Threshold) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *Threshold) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMinutes

`func (o *Threshold) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *Threshold) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *Threshold) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *Threshold) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


