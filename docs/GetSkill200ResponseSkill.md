# GetSkill200ResponseSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**TriggerCondition** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**RequiredTools** | Pointer to **[]string** |  | [optional] 
**Files** | Pointer to **map[string]interface{}** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**DisableModelInvocation** | Pointer to **bool** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**InstalledAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetSkill200ResponseSkill

`func NewGetSkill200ResponseSkill() *GetSkill200ResponseSkill`

NewGetSkill200ResponseSkill instantiates a new GetSkill200ResponseSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSkill200ResponseSkillWithDefaults

`func NewGetSkill200ResponseSkillWithDefaults() *GetSkill200ResponseSkill`

NewGetSkill200ResponseSkillWithDefaults instantiates a new GetSkill200ResponseSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillId

`func (o *GetSkill200ResponseSkill) GetSkillId() string`

GetSkillId returns the SkillId field if non-nil, zero value otherwise.

### GetSkillIdOk

`func (o *GetSkill200ResponseSkill) GetSkillIdOk() (*string, bool)`

GetSkillIdOk returns a tuple with the SkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillId

`func (o *GetSkill200ResponseSkill) SetSkillId(v string)`

SetSkillId sets SkillId field to given value.

### HasSkillId

`func (o *GetSkill200ResponseSkill) HasSkillId() bool`

HasSkillId returns a boolean if a field has been set.

### GetName

`func (o *GetSkill200ResponseSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSkill200ResponseSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSkill200ResponseSkill) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSkill200ResponseSkill) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetSkill200ResponseSkill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSkill200ResponseSkill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSkill200ResponseSkill) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSkill200ResponseSkill) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContent

`func (o *GetSkill200ResponseSkill) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetSkill200ResponseSkill) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetSkill200ResponseSkill) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetSkill200ResponseSkill) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *GetSkill200ResponseSkill) GetTriggerCondition() string`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *GetSkill200ResponseSkill) GetTriggerConditionOk() (*string, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *GetSkill200ResponseSkill) SetTriggerCondition(v string)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *GetSkill200ResponseSkill) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetTags

`func (o *GetSkill200ResponseSkill) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetSkill200ResponseSkill) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetSkill200ResponseSkill) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetSkill200ResponseSkill) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSource

`func (o *GetSkill200ResponseSkill) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetSkill200ResponseSkill) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetSkill200ResponseSkill) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *GetSkill200ResponseSkill) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRequiredTools

`func (o *GetSkill200ResponseSkill) GetRequiredTools() []string`

GetRequiredTools returns the RequiredTools field if non-nil, zero value otherwise.

### GetRequiredToolsOk

`func (o *GetSkill200ResponseSkill) GetRequiredToolsOk() (*[]string, bool)`

GetRequiredToolsOk returns a tuple with the RequiredTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredTools

`func (o *GetSkill200ResponseSkill) SetRequiredTools(v []string)`

SetRequiredTools sets RequiredTools field to given value.

### HasRequiredTools

`func (o *GetSkill200ResponseSkill) HasRequiredTools() bool`

HasRequiredTools returns a boolean if a field has been set.

### GetFiles

`func (o *GetSkill200ResponseSkill) GetFiles() map[string]interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetSkill200ResponseSkill) GetFilesOk() (*map[string]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetSkill200ResponseSkill) SetFiles(v map[string]interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GetSkill200ResponseSkill) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetNamespace

`func (o *GetSkill200ResponseSkill) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GetSkill200ResponseSkill) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GetSkill200ResponseSkill) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GetSkill200ResponseSkill) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetDisableModelInvocation

`func (o *GetSkill200ResponseSkill) GetDisableModelInvocation() bool`

GetDisableModelInvocation returns the DisableModelInvocation field if non-nil, zero value otherwise.

### GetDisableModelInvocationOk

`func (o *GetSkill200ResponseSkill) GetDisableModelInvocationOk() (*bool, bool)`

GetDisableModelInvocationOk returns a tuple with the DisableModelInvocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModelInvocation

`func (o *GetSkill200ResponseSkill) SetDisableModelInvocation(v bool)`

SetDisableModelInvocation sets DisableModelInvocation field to given value.

### HasDisableModelInvocation

`func (o *GetSkill200ResponseSkill) HasDisableModelInvocation() bool`

HasDisableModelInvocation returns a boolean if a field has been set.

### GetAllowedTools

`func (o *GetSkill200ResponseSkill) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *GetSkill200ResponseSkill) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *GetSkill200ResponseSkill) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *GetSkill200ResponseSkill) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetInstalledAt

`func (o *GetSkill200ResponseSkill) GetInstalledAt() time.Time`

GetInstalledAt returns the InstalledAt field if non-nil, zero value otherwise.

### GetInstalledAtOk

`func (o *GetSkill200ResponseSkill) GetInstalledAtOk() (*time.Time, bool)`

GetInstalledAtOk returns a tuple with the InstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAt

`func (o *GetSkill200ResponseSkill) SetInstalledAt(v time.Time)`

SetInstalledAt sets InstalledAt field to given value.

### HasInstalledAt

`func (o *GetSkill200ResponseSkill) HasInstalledAt() bool`

HasInstalledAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetSkill200ResponseSkill) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetSkill200ResponseSkill) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetSkill200ResponseSkill) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetSkill200ResponseSkill) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


