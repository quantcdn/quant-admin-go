# CreateCronJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ScheduleExpression** | **string** |  | 
**Command** | **[]string** |  | 
**TargetContainerName** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] [default to true]

## Methods

### NewCreateCronJobRequest

`func NewCreateCronJobRequest(name string, scheduleExpression string, command []string, ) *CreateCronJobRequest`

NewCreateCronJobRequest instantiates a new CreateCronJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCronJobRequestWithDefaults

`func NewCreateCronJobRequestWithDefaults() *CreateCronJobRequest`

NewCreateCronJobRequestWithDefaults instantiates a new CreateCronJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCronJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCronJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCronJobRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCronJobRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCronJobRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCronJobRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCronJobRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateCronJobRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateCronJobRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduleExpression

`func (o *CreateCronJobRequest) GetScheduleExpression() string`

GetScheduleExpression returns the ScheduleExpression field if non-nil, zero value otherwise.

### GetScheduleExpressionOk

`func (o *CreateCronJobRequest) GetScheduleExpressionOk() (*string, bool)`

GetScheduleExpressionOk returns a tuple with the ScheduleExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleExpression

`func (o *CreateCronJobRequest) SetScheduleExpression(v string)`

SetScheduleExpression sets ScheduleExpression field to given value.


### GetCommand

`func (o *CreateCronJobRequest) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateCronJobRequest) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateCronJobRequest) SetCommand(v []string)`

SetCommand sets Command field to given value.


### GetTargetContainerName

`func (o *CreateCronJobRequest) GetTargetContainerName() string`

GetTargetContainerName returns the TargetContainerName field if non-nil, zero value otherwise.

### GetTargetContainerNameOk

`func (o *CreateCronJobRequest) GetTargetContainerNameOk() (*string, bool)`

GetTargetContainerNameOk returns a tuple with the TargetContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContainerName

`func (o *CreateCronJobRequest) SetTargetContainerName(v string)`

SetTargetContainerName sets TargetContainerName field to given value.

### HasTargetContainerName

`func (o *CreateCronJobRequest) HasTargetContainerName() bool`

HasTargetContainerName returns a boolean if a field has been set.

### SetTargetContainerNameNil

`func (o *CreateCronJobRequest) SetTargetContainerNameNil(b bool)`

 SetTargetContainerNameNil sets the value for TargetContainerName to be an explicit nil

### UnsetTargetContainerName
`func (o *CreateCronJobRequest) UnsetTargetContainerName()`

UnsetTargetContainerName ensures that no value is present for TargetContainerName, not even an explicit nil
### GetIsEnabled

`func (o *CreateCronJobRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateCronJobRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateCronJobRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CreateCronJobRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *CreateCronJobRequest) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *CreateCronJobRequest) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


