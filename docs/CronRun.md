# CronRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** |  | [optional] 
**RunType** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**ExitCode** | Pointer to **NullableInt32** |  | [optional] 
**Output** | Pointer to **[]string** |  | [optional] 
**ScheduleName** | Pointer to **NullableString** |  | [optional] 
**TargetContainerName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCronRun

`func NewCronRun() *CronRun`

NewCronRun instantiates a new CronRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronRunWithDefaults

`func NewCronRunWithDefaults() *CronRun`

NewCronRunWithDefaults instantiates a new CronRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *CronRun) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CronRun) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CronRun) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *CronRun) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetRunType

`func (o *CronRun) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *CronRun) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *CronRun) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *CronRun) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetCommand

`func (o *CronRun) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CronRun) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CronRun) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CronRun) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *CronRun) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *CronRun) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetStatus

`func (o *CronRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CronRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CronRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CronRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *CronRun) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CronRun) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CronRun) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CronRun) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *CronRun) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *CronRun) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *CronRun) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CronRun) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CronRun) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CronRun) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *CronRun) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *CronRun) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetExitCode

`func (o *CronRun) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *CronRun) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *CronRun) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *CronRun) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### SetExitCodeNil

`func (o *CronRun) SetExitCodeNil(b bool)`

 SetExitCodeNil sets the value for ExitCode to be an explicit nil

### UnsetExitCode
`func (o *CronRun) UnsetExitCode()`

UnsetExitCode ensures that no value is present for ExitCode, not even an explicit nil
### GetOutput

`func (o *CronRun) GetOutput() []string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CronRun) GetOutputOk() (*[]string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CronRun) SetOutput(v []string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CronRun) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetScheduleName

`func (o *CronRun) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *CronRun) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *CronRun) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *CronRun) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### SetScheduleNameNil

`func (o *CronRun) SetScheduleNameNil(b bool)`

 SetScheduleNameNil sets the value for ScheduleName to be an explicit nil

### UnsetScheduleName
`func (o *CronRun) UnsetScheduleName()`

UnsetScheduleName ensures that no value is present for ScheduleName, not even an explicit nil
### GetTargetContainerName

`func (o *CronRun) GetTargetContainerName() string`

GetTargetContainerName returns the TargetContainerName field if non-nil, zero value otherwise.

### GetTargetContainerNameOk

`func (o *CronRun) GetTargetContainerNameOk() (*string, bool)`

GetTargetContainerNameOk returns a tuple with the TargetContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContainerName

`func (o *CronRun) SetTargetContainerName(v string)`

SetTargetContainerName sets TargetContainerName field to given value.

### HasTargetContainerName

`func (o *CronRun) HasTargetContainerName() bool`

HasTargetContainerName returns a boolean if a field has been set.

### SetTargetContainerNameNil

`func (o *CronRun) SetTargetContainerNameNil(b bool)`

 SetTargetContainerNameNil sets the value for TargetContainerName to be an explicit nil

### UnsetTargetContainerName
`func (o *CronRun) UnsetTargetContainerName()`

UnsetTargetContainerName ensures that no value is present for TargetContainerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


