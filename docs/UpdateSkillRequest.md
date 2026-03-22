# UpdateSkillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**TriggerCondition** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**RequiredTools** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**Files** | Pointer to **map[string]interface{}** |  | [optional] 
**DisableModelInvocation** | Pointer to **bool** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateSkillRequest

`func NewUpdateSkillRequest() *UpdateSkillRequest`

NewUpdateSkillRequest instantiates a new UpdateSkillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSkillRequestWithDefaults

`func NewUpdateSkillRequestWithDefaults() *UpdateSkillRequest`

NewUpdateSkillRequestWithDefaults instantiates a new UpdateSkillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSkillRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSkillRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSkillRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSkillRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSkillRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSkillRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSkillRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSkillRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContent

`func (o *UpdateSkillRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateSkillRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateSkillRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateSkillRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *UpdateSkillRequest) GetTriggerCondition() string`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *UpdateSkillRequest) GetTriggerConditionOk() (*string, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *UpdateSkillRequest) SetTriggerCondition(v string)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *UpdateSkillRequest) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetTags

`func (o *UpdateSkillRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateSkillRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateSkillRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateSkillRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRequiredTools

`func (o *UpdateSkillRequest) GetRequiredTools() []string`

GetRequiredTools returns the RequiredTools field if non-nil, zero value otherwise.

### GetRequiredToolsOk

`func (o *UpdateSkillRequest) GetRequiredToolsOk() (*[]string, bool)`

GetRequiredToolsOk returns a tuple with the RequiredTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredTools

`func (o *UpdateSkillRequest) SetRequiredTools(v []string)`

SetRequiredTools sets RequiredTools field to given value.

### HasRequiredTools

`func (o *UpdateSkillRequest) HasRequiredTools() bool`

HasRequiredTools returns a boolean if a field has been set.

### GetSource

`func (o *UpdateSkillRequest) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpdateSkillRequest) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpdateSkillRequest) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *UpdateSkillRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetFiles

`func (o *UpdateSkillRequest) GetFiles() map[string]interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *UpdateSkillRequest) GetFilesOk() (*map[string]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *UpdateSkillRequest) SetFiles(v map[string]interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *UpdateSkillRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetDisableModelInvocation

`func (o *UpdateSkillRequest) GetDisableModelInvocation() bool`

GetDisableModelInvocation returns the DisableModelInvocation field if non-nil, zero value otherwise.

### GetDisableModelInvocationOk

`func (o *UpdateSkillRequest) GetDisableModelInvocationOk() (*bool, bool)`

GetDisableModelInvocationOk returns a tuple with the DisableModelInvocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModelInvocation

`func (o *UpdateSkillRequest) SetDisableModelInvocation(v bool)`

SetDisableModelInvocation sets DisableModelInvocation field to given value.

### HasDisableModelInvocation

`func (o *UpdateSkillRequest) HasDisableModelInvocation() bool`

HasDisableModelInvocation returns a boolean if a field has been set.

### GetAllowedTools

`func (o *UpdateSkillRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *UpdateSkillRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *UpdateSkillRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *UpdateSkillRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetNamespace

`func (o *UpdateSkillRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UpdateSkillRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UpdateSkillRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UpdateSkillRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


