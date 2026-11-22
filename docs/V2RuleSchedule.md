# V2RuleSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | **string** | IANA time zone name, or UTC. Invalid values return 422. | 
**Mode** | **string** | active_during: the rule applies only inside the windows. inactive_during: the rule applies only outside them. | 
**Weekly** | Pointer to [**[]V2RuleScheduleWeeklyInner**](V2RuleScheduleWeeklyInner.md) | Weekly windows. Stored as an empty list when omitted. | [optional] 
**Once** | Pointer to [**[]V2RuleScheduleOnceInner**](V2RuleScheduleOnceInner.md) | One-off ranges. Stored as an empty list when omitted. | [optional] 

## Methods

### NewV2RuleSchedule

`func NewV2RuleSchedule(timezone string, mode string, ) *V2RuleSchedule`

NewV2RuleSchedule instantiates a new V2RuleSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleScheduleWithDefaults

`func NewV2RuleScheduleWithDefaults() *V2RuleSchedule`

NewV2RuleScheduleWithDefaults instantiates a new V2RuleSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *V2RuleSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V2RuleSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V2RuleSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetMode

`func (o *V2RuleSchedule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V2RuleSchedule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V2RuleSchedule) SetMode(v string)`

SetMode sets Mode field to given value.


### GetWeekly

`func (o *V2RuleSchedule) GetWeekly() []V2RuleScheduleWeeklyInner`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *V2RuleSchedule) GetWeeklyOk() (*[]V2RuleScheduleWeeklyInner, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *V2RuleSchedule) SetWeekly(v []V2RuleScheduleWeeklyInner)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *V2RuleSchedule) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetOnce

`func (o *V2RuleSchedule) GetOnce() []V2RuleScheduleOnceInner`

GetOnce returns the Once field if non-nil, zero value otherwise.

### GetOnceOk

`func (o *V2RuleSchedule) GetOnceOk() (*[]V2RuleScheduleOnceInner, bool)`

GetOnceOk returns a tuple with the Once field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnce

`func (o *V2RuleSchedule) SetOnce(v []V2RuleScheduleOnceInner)`

SetOnce sets Once field to given value.

### HasOnce

`func (o *V2RuleSchedule) HasOnce() bool`

HasOnce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


