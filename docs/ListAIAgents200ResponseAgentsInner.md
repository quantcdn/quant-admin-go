# ListAIAgents200ResponseAgentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**AssignedSkills** | Pointer to **[]string** |  | [optional] 
**LongContext** | Pointer to **bool** |  | [optional] 
**GuardrailPreset** | Pointer to **string** |  | [optional] 
**IsGlobal** | Pointer to **bool** |  | [optional] 
**HasOverlay** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListAIAgents200ResponseAgentsInner

`func NewListAIAgents200ResponseAgentsInner() *ListAIAgents200ResponseAgentsInner`

NewListAIAgents200ResponseAgentsInner instantiates a new ListAIAgents200ResponseAgentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAIAgents200ResponseAgentsInnerWithDefaults

`func NewListAIAgents200ResponseAgentsInnerWithDefaults() *ListAIAgents200ResponseAgentsInner`

NewListAIAgents200ResponseAgentsInnerWithDefaults instantiates a new ListAIAgents200ResponseAgentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ListAIAgents200ResponseAgentsInner) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ListAIAgents200ResponseAgentsInner) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ListAIAgents200ResponseAgentsInner) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ListAIAgents200ResponseAgentsInner) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetName

`func (o *ListAIAgents200ResponseAgentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAIAgents200ResponseAgentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAIAgents200ResponseAgentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListAIAgents200ResponseAgentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListAIAgents200ResponseAgentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListAIAgents200ResponseAgentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListAIAgents200ResponseAgentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListAIAgents200ResponseAgentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *ListAIAgents200ResponseAgentsInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ListAIAgents200ResponseAgentsInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ListAIAgents200ResponseAgentsInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ListAIAgents200ResponseAgentsInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetModelId

`func (o *ListAIAgents200ResponseAgentsInner) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ListAIAgents200ResponseAgentsInner) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ListAIAgents200ResponseAgentsInner) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ListAIAgents200ResponseAgentsInner) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetTemperature

`func (o *ListAIAgents200ResponseAgentsInner) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ListAIAgents200ResponseAgentsInner) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ListAIAgents200ResponseAgentsInner) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ListAIAgents200ResponseAgentsInner) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *ListAIAgents200ResponseAgentsInner) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ListAIAgents200ResponseAgentsInner) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ListAIAgents200ResponseAgentsInner) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ListAIAgents200ResponseAgentsInner) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetAllowedTools

`func (o *ListAIAgents200ResponseAgentsInner) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *ListAIAgents200ResponseAgentsInner) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *ListAIAgents200ResponseAgentsInner) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *ListAIAgents200ResponseAgentsInner) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetAssignedSkills

`func (o *ListAIAgents200ResponseAgentsInner) GetAssignedSkills() []string`

GetAssignedSkills returns the AssignedSkills field if non-nil, zero value otherwise.

### GetAssignedSkillsOk

`func (o *ListAIAgents200ResponseAgentsInner) GetAssignedSkillsOk() (*[]string, bool)`

GetAssignedSkillsOk returns a tuple with the AssignedSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedSkills

`func (o *ListAIAgents200ResponseAgentsInner) SetAssignedSkills(v []string)`

SetAssignedSkills sets AssignedSkills field to given value.

### HasAssignedSkills

`func (o *ListAIAgents200ResponseAgentsInner) HasAssignedSkills() bool`

HasAssignedSkills returns a boolean if a field has been set.

### GetLongContext

`func (o *ListAIAgents200ResponseAgentsInner) GetLongContext() bool`

GetLongContext returns the LongContext field if non-nil, zero value otherwise.

### GetLongContextOk

`func (o *ListAIAgents200ResponseAgentsInner) GetLongContextOk() (*bool, bool)`

GetLongContextOk returns a tuple with the LongContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongContext

`func (o *ListAIAgents200ResponseAgentsInner) SetLongContext(v bool)`

SetLongContext sets LongContext field to given value.

### HasLongContext

`func (o *ListAIAgents200ResponseAgentsInner) HasLongContext() bool`

HasLongContext returns a boolean if a field has been set.

### GetGuardrailPreset

`func (o *ListAIAgents200ResponseAgentsInner) GetGuardrailPreset() string`

GetGuardrailPreset returns the GuardrailPreset field if non-nil, zero value otherwise.

### GetGuardrailPresetOk

`func (o *ListAIAgents200ResponseAgentsInner) GetGuardrailPresetOk() (*string, bool)`

GetGuardrailPresetOk returns a tuple with the GuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailPreset

`func (o *ListAIAgents200ResponseAgentsInner) SetGuardrailPreset(v string)`

SetGuardrailPreset sets GuardrailPreset field to given value.

### HasGuardrailPreset

`func (o *ListAIAgents200ResponseAgentsInner) HasGuardrailPreset() bool`

HasGuardrailPreset returns a boolean if a field has been set.

### GetIsGlobal

`func (o *ListAIAgents200ResponseAgentsInner) GetIsGlobal() bool`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *ListAIAgents200ResponseAgentsInner) GetIsGlobalOk() (*bool, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *ListAIAgents200ResponseAgentsInner) SetIsGlobal(v bool)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *ListAIAgents200ResponseAgentsInner) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### GetHasOverlay

`func (o *ListAIAgents200ResponseAgentsInner) GetHasOverlay() bool`

GetHasOverlay returns the HasOverlay field if non-nil, zero value otherwise.

### GetHasOverlayOk

`func (o *ListAIAgents200ResponseAgentsInner) GetHasOverlayOk() (*bool, bool)`

GetHasOverlayOk returns a tuple with the HasOverlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOverlay

`func (o *ListAIAgents200ResponseAgentsInner) SetHasOverlay(v bool)`

SetHasOverlay sets HasOverlay field to given value.

### HasHasOverlay

`func (o *ListAIAgents200ResponseAgentsInner) HasHasOverlay() bool`

HasHasOverlay returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListAIAgents200ResponseAgentsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListAIAgents200ResponseAgentsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListAIAgents200ResponseAgentsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListAIAgents200ResponseAgentsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListAIAgents200ResponseAgentsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListAIAgents200ResponseAgentsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListAIAgents200ResponseAgentsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListAIAgents200ResponseAgentsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


