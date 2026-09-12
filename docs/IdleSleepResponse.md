# IdleSleepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**IdleMinutes** | **NullableInt32** |  | 
**State** | **string** |  | 
**StateChangedAt** | **NullableTime** |  | 

## Methods

### NewIdleSleepResponse

`func NewIdleSleepResponse(enabled bool, idleMinutes NullableInt32, state string, stateChangedAt NullableTime, ) *IdleSleepResponse`

NewIdleSleepResponse instantiates a new IdleSleepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdleSleepResponseWithDefaults

`func NewIdleSleepResponseWithDefaults() *IdleSleepResponse`

NewIdleSleepResponseWithDefaults instantiates a new IdleSleepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *IdleSleepResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdleSleepResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdleSleepResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdleMinutes

`func (o *IdleSleepResponse) GetIdleMinutes() int32`

GetIdleMinutes returns the IdleMinutes field if non-nil, zero value otherwise.

### GetIdleMinutesOk

`func (o *IdleSleepResponse) GetIdleMinutesOk() (*int32, bool)`

GetIdleMinutesOk returns a tuple with the IdleMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleMinutes

`func (o *IdleSleepResponse) SetIdleMinutes(v int32)`

SetIdleMinutes sets IdleMinutes field to given value.


### SetIdleMinutesNil

`func (o *IdleSleepResponse) SetIdleMinutesNil(b bool)`

 SetIdleMinutesNil sets the value for IdleMinutes to be an explicit nil

### UnsetIdleMinutes
`func (o *IdleSleepResponse) UnsetIdleMinutes()`

UnsetIdleMinutes ensures that no value is present for IdleMinutes, not even an explicit nil
### GetState

`func (o *IdleSleepResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IdleSleepResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IdleSleepResponse) SetState(v string)`

SetState sets State field to given value.


### GetStateChangedAt

`func (o *IdleSleepResponse) GetStateChangedAt() time.Time`

GetStateChangedAt returns the StateChangedAt field if non-nil, zero value otherwise.

### GetStateChangedAtOk

`func (o *IdleSleepResponse) GetStateChangedAtOk() (*time.Time, bool)`

GetStateChangedAtOk returns a tuple with the StateChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedAt

`func (o *IdleSleepResponse) SetStateChangedAt(v time.Time)`

SetStateChangedAt sets StateChangedAt field to given value.


### SetStateChangedAtNil

`func (o *IdleSleepResponse) SetStateChangedAtNil(b bool)`

 SetStateChangedAtNil sets the value for StateChangedAt to be an explicit nil

### UnsetStateChangedAt
`func (o *IdleSleepResponse) UnsetStateChangedAt()`

UnsetStateChangedAt ensures that no value is present for StateChangedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


