# ListSkills200ResponseSkillsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**TriggerCondition** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**InstalledAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListSkills200ResponseSkillsInner

`func NewListSkills200ResponseSkillsInner() *ListSkills200ResponseSkillsInner`

NewListSkills200ResponseSkillsInner instantiates a new ListSkills200ResponseSkillsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSkills200ResponseSkillsInnerWithDefaults

`func NewListSkills200ResponseSkillsInnerWithDefaults() *ListSkills200ResponseSkillsInner`

NewListSkills200ResponseSkillsInnerWithDefaults instantiates a new ListSkills200ResponseSkillsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillId

`func (o *ListSkills200ResponseSkillsInner) GetSkillId() string`

GetSkillId returns the SkillId field if non-nil, zero value otherwise.

### GetSkillIdOk

`func (o *ListSkills200ResponseSkillsInner) GetSkillIdOk() (*string, bool)`

GetSkillIdOk returns a tuple with the SkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillId

`func (o *ListSkills200ResponseSkillsInner) SetSkillId(v string)`

SetSkillId sets SkillId field to given value.

### HasSkillId

`func (o *ListSkills200ResponseSkillsInner) HasSkillId() bool`

HasSkillId returns a boolean if a field has been set.

### GetName

`func (o *ListSkills200ResponseSkillsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSkills200ResponseSkillsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSkills200ResponseSkillsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSkills200ResponseSkillsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListSkills200ResponseSkillsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListSkills200ResponseSkillsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListSkills200ResponseSkillsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListSkills200ResponseSkillsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *ListSkills200ResponseSkillsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListSkills200ResponseSkillsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListSkills200ResponseSkillsInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListSkills200ResponseSkillsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSource

`func (o *ListSkills200ResponseSkillsInner) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListSkills200ResponseSkillsInner) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListSkills200ResponseSkillsInner) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *ListSkills200ResponseSkillsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTriggerCondition

`func (o *ListSkills200ResponseSkillsInner) GetTriggerCondition() string`

GetTriggerCondition returns the TriggerCondition field if non-nil, zero value otherwise.

### GetTriggerConditionOk

`func (o *ListSkills200ResponseSkillsInner) GetTriggerConditionOk() (*string, bool)`

GetTriggerConditionOk returns a tuple with the TriggerCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCondition

`func (o *ListSkills200ResponseSkillsInner) SetTriggerCondition(v string)`

SetTriggerCondition sets TriggerCondition field to given value.

### HasTriggerCondition

`func (o *ListSkills200ResponseSkillsInner) HasTriggerCondition() bool`

HasTriggerCondition returns a boolean if a field has been set.

### GetNamespace

`func (o *ListSkills200ResponseSkillsInner) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListSkills200ResponseSkillsInner) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListSkills200ResponseSkillsInner) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ListSkills200ResponseSkillsInner) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetInstalledAt

`func (o *ListSkills200ResponseSkillsInner) GetInstalledAt() time.Time`

GetInstalledAt returns the InstalledAt field if non-nil, zero value otherwise.

### GetInstalledAtOk

`func (o *ListSkills200ResponseSkillsInner) GetInstalledAtOk() (*time.Time, bool)`

GetInstalledAtOk returns a tuple with the InstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledAt

`func (o *ListSkills200ResponseSkillsInner) SetInstalledAt(v time.Time)`

SetInstalledAt sets InstalledAt field to given value.

### HasInstalledAt

`func (o *ListSkills200ResponseSkillsInner) HasInstalledAt() bool`

HasInstalledAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListSkills200ResponseSkillsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListSkills200ResponseSkillsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListSkills200ResponseSkillsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListSkills200ResponseSkillsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


