# Cron

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ScheduleExpression** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**TargetContainerName** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCron

`func NewCron() *Cron`

NewCron instantiates a new Cron object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronWithDefaults

`func NewCronWithDefaults() *Cron`

NewCronWithDefaults instantiates a new Cron object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cron) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cron) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cron) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cron) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Cron) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cron) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cron) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cron) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Cron) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Cron) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleExpression

`func (o *Cron) GetScheduleExpression() string`

GetScheduleExpression returns the ScheduleExpression field if non-nil, zero value otherwise.

### GetScheduleExpressionOk

`func (o *Cron) GetScheduleExpressionOk() (*string, bool)`

GetScheduleExpressionOk returns a tuple with the ScheduleExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleExpression

`func (o *Cron) SetScheduleExpression(v string)`

SetScheduleExpression sets ScheduleExpression field to given value.

### HasScheduleExpression

`func (o *Cron) HasScheduleExpression() bool`

HasScheduleExpression returns a boolean if a field has been set.

### GetCommand

`func (o *Cron) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Cron) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Cron) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Cron) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetTargetContainerName

`func (o *Cron) GetTargetContainerName() string`

GetTargetContainerName returns the TargetContainerName field if non-nil, zero value otherwise.

### GetTargetContainerNameOk

`func (o *Cron) GetTargetContainerNameOk() (*string, bool)`

GetTargetContainerNameOk returns a tuple with the TargetContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContainerName

`func (o *Cron) SetTargetContainerName(v string)`

SetTargetContainerName sets TargetContainerName field to given value.

### HasTargetContainerName

`func (o *Cron) HasTargetContainerName() bool`

HasTargetContainerName returns a boolean if a field has been set.

### SetTargetContainerNameNil

`func (o *Cron) SetTargetContainerNameNil(b bool)`

 SetTargetContainerNameNil sets the value for TargetContainerName to be an explicit nil

### UnsetTargetContainerName
`func (o *Cron) UnsetTargetContainerName()`

UnsetTargetContainerName ensures that no value is present for TargetContainerName, not even an explicit nil
### GetIsEnabled

`func (o *Cron) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Cron) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Cron) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *Cron) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


