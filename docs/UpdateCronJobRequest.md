# UpdateCronJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**ScheduleExpression** | Pointer to **NullableString** |  | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**TargetContainerName** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateCronJobRequest

`func NewUpdateCronJobRequest() *UpdateCronJobRequest`

NewUpdateCronJobRequest instantiates a new UpdateCronJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCronJobRequestWithDefaults

`func NewUpdateCronJobRequestWithDefaults() *UpdateCronJobRequest`

NewUpdateCronJobRequestWithDefaults instantiates a new UpdateCronJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateCronJobRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCronJobRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCronJobRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCronJobRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateCronJobRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateCronJobRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleExpression

`func (o *UpdateCronJobRequest) GetScheduleExpression() string`

GetScheduleExpression returns the ScheduleExpression field if non-nil, zero value otherwise.

### GetScheduleExpressionOk

`func (o *UpdateCronJobRequest) GetScheduleExpressionOk() (*string, bool)`

GetScheduleExpressionOk returns a tuple with the ScheduleExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleExpression

`func (o *UpdateCronJobRequest) SetScheduleExpression(v string)`

SetScheduleExpression sets ScheduleExpression field to given value.

### HasScheduleExpression

`func (o *UpdateCronJobRequest) HasScheduleExpression() bool`

HasScheduleExpression returns a boolean if a field has been set.

### SetScheduleExpressionNil

`func (o *UpdateCronJobRequest) SetScheduleExpressionNil(b bool)`

 SetScheduleExpressionNil sets the value for ScheduleExpression to be an explicit nil

### UnsetScheduleExpression
`func (o *UpdateCronJobRequest) UnsetScheduleExpression()`

UnsetScheduleExpression ensures that no value is present for ScheduleExpression, not even an explicit nil
### GetCommand

`func (o *UpdateCronJobRequest) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *UpdateCronJobRequest) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *UpdateCronJobRequest) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *UpdateCronJobRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *UpdateCronJobRequest) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *UpdateCronJobRequest) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetTargetContainerName

`func (o *UpdateCronJobRequest) GetTargetContainerName() string`

GetTargetContainerName returns the TargetContainerName field if non-nil, zero value otherwise.

### GetTargetContainerNameOk

`func (o *UpdateCronJobRequest) GetTargetContainerNameOk() (*string, bool)`

GetTargetContainerNameOk returns a tuple with the TargetContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContainerName

`func (o *UpdateCronJobRequest) SetTargetContainerName(v string)`

SetTargetContainerName sets TargetContainerName field to given value.

### HasTargetContainerName

`func (o *UpdateCronJobRequest) HasTargetContainerName() bool`

HasTargetContainerName returns a boolean if a field has been set.

### SetTargetContainerNameNil

`func (o *UpdateCronJobRequest) SetTargetContainerNameNil(b bool)`

 SetTargetContainerNameNil sets the value for TargetContainerName to be an explicit nil

### UnsetTargetContainerName
`func (o *UpdateCronJobRequest) UnsetTargetContainerName()`

UnsetTargetContainerName ensures that no value is present for TargetContainerName, not even an explicit nil
### GetIsEnabled

`func (o *UpdateCronJobRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateCronJobRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateCronJobRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UpdateCronJobRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *UpdateCronJobRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *UpdateCronJobRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


