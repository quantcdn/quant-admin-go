# V2RuleScheduleWeeklyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **[]string** |  | 
**Start** | **string** | Local time HH:MM, 00:00 to 23:59. | 
**End** | **string** | Local time HH:MM, 00:00 to 24:00. | 

## Methods

### NewV2RuleScheduleWeeklyInner

`func NewV2RuleScheduleWeeklyInner(days []string, start string, end string, ) *V2RuleScheduleWeeklyInner`

NewV2RuleScheduleWeeklyInner instantiates a new V2RuleScheduleWeeklyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleScheduleWeeklyInnerWithDefaults

`func NewV2RuleScheduleWeeklyInnerWithDefaults() *V2RuleScheduleWeeklyInner`

NewV2RuleScheduleWeeklyInnerWithDefaults instantiates a new V2RuleScheduleWeeklyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *V2RuleScheduleWeeklyInner) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *V2RuleScheduleWeeklyInner) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *V2RuleScheduleWeeklyInner) SetDays(v []string)`

SetDays sets Days field to given value.


### GetStart

`func (o *V2RuleScheduleWeeklyInner) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V2RuleScheduleWeeklyInner) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V2RuleScheduleWeeklyInner) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *V2RuleScheduleWeeklyInner) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V2RuleScheduleWeeklyInner) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V2RuleScheduleWeeklyInner) SetEnd(v string)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


