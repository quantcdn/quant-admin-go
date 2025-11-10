# Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** |  | [optional] 
**RunType** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 
**Output** | Pointer to **[]string** |  | [optional] 
**ScheduleName** | Pointer to **string** |  | [optional] 
**TargetContainerName** | Pointer to **string** |  | [optional] 

## Methods

### NewCommand

`func NewCommand() *Command`

NewCommand instantiates a new Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandWithDefaults

`func NewCommandWithDefaults() *Command`

NewCommandWithDefaults instantiates a new Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *Command) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *Command) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *Command) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *Command) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetRunType

`func (o *Command) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *Command) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *Command) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *Command) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetCommand

`func (o *Command) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Command) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Command) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Command) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetStatus

`func (o *Command) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Command) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Command) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Command) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *Command) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Command) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Command) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Command) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Command) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Command) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Command) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Command) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExitCode

`func (o *Command) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *Command) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *Command) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *Command) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetOutput

`func (o *Command) GetOutput() []string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Command) GetOutputOk() (*[]string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Command) SetOutput(v []string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Command) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetScheduleName

`func (o *Command) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *Command) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *Command) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *Command) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetTargetContainerName

`func (o *Command) GetTargetContainerName() string`

GetTargetContainerName returns the TargetContainerName field if non-nil, zero value otherwise.

### GetTargetContainerNameOk

`func (o *Command) GetTargetContainerNameOk() (*string, bool)`

GetTargetContainerNameOk returns a tuple with the TargetContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContainerName

`func (o *Command) SetTargetContainerName(v string)`

SetTargetContainerName sets TargetContainerName field to given value.

### HasTargetContainerName

`func (o *Command) HasTargetContainerName() bool`

HasTargetContainerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


