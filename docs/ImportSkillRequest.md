# ImportSkillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**ImportSkillRequestSource**](ImportSkillRequestSource.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TriggerCondition** | Pointer to **string** |  | [optional] 
**RequiredTools** | Pointer to **[]string** |  | [optional] 
**DisableModelInvocation** | Pointer to **bool** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**InstalledBy** | Pointer to **string** |  | [optional] 

## Methods

### NewImportSkillRequest

`func NewImportSkillRequest(source ImportSkillRequestSource, ) *ImportSkillRequest`

NewImportSkillRequest instantiates a new ImportSkillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportSkillRequestWithDefaults

`func NewImportSkillRequestWithDefaults() *ImportSkillRequest`

NewImportSkillRequestWithDefaults instantiates a new ImportSkillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ImportSkillRequest) GetSource() ImportSkillRequestSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ImportSkillRequest) GetSourceOk() (*ImportSkillRequestSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ImportSkillRequest) SetSource(v ImportSkillRequestSource)`

SetSource sets Source field to given value.


### GetName

`func (o *ImportSkillRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportSkillRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportSkillRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportSkillRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ImportSkillRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImportSkillRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImportSkillRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImportSkillRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *ImportSkillRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportSkillRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportSkillRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportSkillRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *ImportSkillRequest) GetTriggerCondition() string`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *ImportSkillRequest) GetTriggerConditionOk() (*string, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *ImportSkillRequest) SetTriggerCondition(v string)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *ImportSkillRequest) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetRequiredTools

`func (o *ImportSkillRequest) GetRequiredTools() []string`

GetRequiredTools returns the RequiredTools field if non-nil, zero value otherwise.

### GetRequiredToolsOk

`func (o *ImportSkillRequest) GetRequiredToolsOk() (*[]string, bool)`

GetRequiredToolsOk returns a tuple with the RequiredTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredTools

`func (o *ImportSkillRequest) SetRequiredTools(v []string)`

SetRequiredTools sets RequiredTools field to given value.

### HasRequiredTools

`func (o *ImportSkillRequest) HasRequiredTools() bool`

HasRequiredTools returns a boolean if a field has been set.

### GetDisableModelInvocation

`func (o *ImportSkillRequest) GetDisableModelInvocation() bool`

GetDisableModelInvocation returns the DisableModelInvocation field if non-nil, zero value otherwise.

### GetDisableModelInvocationOk

`func (o *ImportSkillRequest) GetDisableModelInvocationOk() (*bool, bool)`

GetDisableModelInvocationOk returns a tuple with the DisableModelInvocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModelInvocation

`func (o *ImportSkillRequest) SetDisableModelInvocation(v bool)`

SetDisableModelInvocation sets DisableModelInvocation field to given value.

### HasDisableModelInvocation

`func (o *ImportSkillRequest) HasDisableModelInvocation() bool`

HasDisableModelInvocation returns a boolean if a field has been set.

### GetAllowedTools

`func (o *ImportSkillRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *ImportSkillRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *ImportSkillRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *ImportSkillRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetInstalledBy

`func (o *ImportSkillRequest) GetInstalledBy() string`

GetInstalledBy returns the InstalledBy field if non-nil, zero value otherwise.

### GetInstalledByOk

`func (o *ImportSkillRequest) GetInstalledByOk() (*string, bool)`

GetInstalledByOk returns a tuple with the InstalledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledBy

`func (o *ImportSkillRequest) SetInstalledBy(v string)`

SetInstalledBy sets InstalledBy field to given value.

### HasInstalledBy

`func (o *ImportSkillRequest) HasInstalledBy() bool`

HasInstalledBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


