# SetIdleSleepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the environment sleeps when idle. | 
**IdleMinutes** | Pointer to **int32** | Minutes with no requests before compute sleeps. | [optional] [default to 30]

## Methods

### NewSetIdleSleepRequest

`func NewSetIdleSleepRequest(enabled bool, ) *SetIdleSleepRequest`

NewSetIdleSleepRequest instantiates a new SetIdleSleepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetIdleSleepRequestWithDefaults

`func NewSetIdleSleepRequestWithDefaults() *SetIdleSleepRequest`

NewSetIdleSleepRequestWithDefaults instantiates a new SetIdleSleepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SetIdleSleepRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SetIdleSleepRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SetIdleSleepRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdleMinutes

`func (o *SetIdleSleepRequest) GetIdleMinutes() int32`

GetIdleMinutes returns the IdleMinutes field if non-nil, zero value otherwise.

### GetIdleMinutesOk

`func (o *SetIdleSleepRequest) GetIdleMinutesOk() (*int32, bool)`

GetIdleMinutesOk returns a tuple with the IdleMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleMinutes

`func (o *SetIdleSleepRequest) SetIdleMinutes(v int32)`

SetIdleMinutes sets IdleMinutes field to given value.

### HasIdleMinutes

`func (o *SetIdleSleepRequest) HasIdleMinutes() bool`

HasIdleMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


