# CreateSkillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Content** | **string** |  | 
**TriggerCondition** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**RequiredTools** | Pointer to **[]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Files** | Pointer to **map[string]interface{}** |  | [optional] 
**DisableModelInvocation** | Pointer to **bool** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**InstalledBy** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSkillRequest

`func NewCreateSkillRequest(name string, content string, triggerCondition string, ) *CreateSkillRequest`

NewCreateSkillRequest instantiates a new CreateSkillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSkillRequestWithDefaults

`func NewCreateSkillRequestWithDefaults() *CreateSkillRequest`

NewCreateSkillRequestWithDefaults instantiates a new CreateSkillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSkillRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSkillRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSkillRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSkillRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSkillRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSkillRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSkillRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContent

`func (o *CreateSkillRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateSkillRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateSkillRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetTriggerCondition

`func (o *CreateSkillRequest) GetTriggerCondition() string`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *CreateSkillRequest) GetTriggerConditionOk() (*string, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *CreateSkillRequest) SetTriggerCondition(v string)`

SetTriggerCondition sets TriggerCondition field to given value.


### GetTags

`func (o *CreateSkillRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSkillRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSkillRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSkillRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRequiredTools

`func (o *CreateSkillRequest) GetRequiredTools() []string`

GetRequiredTools returns the RequiredTools field if non-nil, zero value otherwise.

### GetRequiredToolsOk

`func (o *CreateSkillRequest) GetRequiredToolsOk() (*[]string, bool)`

GetRequiredToolsOk returns a tuple with the RequiredTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredTools

`func (o *CreateSkillRequest) SetRequiredTools(v []string)`

SetRequiredTools sets RequiredTools field to given value.

### HasRequiredTools

`func (o *CreateSkillRequest) HasRequiredTools() bool`

HasRequiredTools returns a boolean if a field has been set.

### GetNamespace

`func (o *CreateSkillRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateSkillRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateSkillRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateSkillRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetFiles

`func (o *CreateSkillRequest) GetFiles() map[string]interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CreateSkillRequest) GetFilesOk() (*map[string]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CreateSkillRequest) SetFiles(v map[string]interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CreateSkillRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetDisableModelInvocation

`func (o *CreateSkillRequest) GetDisableModelInvocation() bool`

GetDisableModelInvocation returns the DisableModelInvocation field if non-nil, zero value otherwise.

### GetDisableModelInvocationOk

`func (o *CreateSkillRequest) GetDisableModelInvocationOk() (*bool, bool)`

GetDisableModelInvocationOk returns a tuple with the DisableModelInvocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModelInvocation

`func (o *CreateSkillRequest) SetDisableModelInvocation(v bool)`

SetDisableModelInvocation sets DisableModelInvocation field to given value.

### HasDisableModelInvocation

`func (o *CreateSkillRequest) HasDisableModelInvocation() bool`

HasDisableModelInvocation returns a boolean if a field has been set.

### GetAllowedTools

`func (o *CreateSkillRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *CreateSkillRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *CreateSkillRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *CreateSkillRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetInstalledBy

`func (o *CreateSkillRequest) GetInstalledBy() string`

GetInstalledBy returns the InstalledBy field if non-nil, zero value otherwise.

### GetInstalledByOk

`func (o *CreateSkillRequest) GetInstalledByOk() (*string, bool)`

GetInstalledByOk returns a tuple with the InstalledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledBy

`func (o *CreateSkillRequest) SetInstalledBy(v string)`

SetInstalledBy sets InstalledBy field to given value.

### HasInstalledBy

`func (o *CreateSkillRequest) HasInstalledBy() bool`

HasInstalledBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


