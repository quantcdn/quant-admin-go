# UpdateAIAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**SystemPrompt** | Pointer to **string** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**AllowedCollections** | Pointer to **[]string** |  | [optional] 
**AssignedSkills** | Pointer to **[]string** | Skill IDs to assign to this agent | [optional] 
**LongContext** | Pointer to **bool** | Enable 1M context window support | [optional] 
**GuardrailPreset** | Pointer to **string** | Guardrail preset name | [optional] 
**FilterPolicies** | Pointer to **[]string** | Filter policy IDs to apply to this agent&#39;s inference requests | [optional] 

## Methods

### NewUpdateAIAgentRequest

`func NewUpdateAIAgentRequest() *UpdateAIAgentRequest`

NewUpdateAIAgentRequest instantiates a new UpdateAIAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAIAgentRequestWithDefaults

`func NewUpdateAIAgentRequestWithDefaults() *UpdateAIAgentRequest`

NewUpdateAIAgentRequestWithDefaults instantiates a new UpdateAIAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAIAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAIAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAIAgentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAIAgentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAIAgentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAIAgentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAIAgentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAIAgentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *UpdateAIAgentRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateAIAgentRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateAIAgentRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateAIAgentRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *UpdateAIAgentRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *UpdateAIAgentRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *UpdateAIAgentRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *UpdateAIAgentRequest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetTemperature

`func (o *UpdateAIAgentRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *UpdateAIAgentRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *UpdateAIAgentRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *UpdateAIAgentRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetModelId

`func (o *UpdateAIAgentRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *UpdateAIAgentRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *UpdateAIAgentRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *UpdateAIAgentRequest) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetMaxTokens

`func (o *UpdateAIAgentRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *UpdateAIAgentRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *UpdateAIAgentRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *UpdateAIAgentRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetAllowedTools

`func (o *UpdateAIAgentRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *UpdateAIAgentRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *UpdateAIAgentRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *UpdateAIAgentRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *UpdateAIAgentRequest) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *UpdateAIAgentRequest) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *UpdateAIAgentRequest) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *UpdateAIAgentRequest) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetAssignedSkills

`func (o *UpdateAIAgentRequest) GetAssignedSkills() []string`

GetAssignedSkills returns the AssignedSkills field if non-nil, zero value otherwise.

### GetAssignedSkillsOk

`func (o *UpdateAIAgentRequest) GetAssignedSkillsOk() (*[]string, bool)`

GetAssignedSkillsOk returns a tuple with the AssignedSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedSkills

`func (o *UpdateAIAgentRequest) SetAssignedSkills(v []string)`

SetAssignedSkills sets AssignedSkills field to given value.

### HasAssignedSkills

`func (o *UpdateAIAgentRequest) HasAssignedSkills() bool`

HasAssignedSkills returns a boolean if a field has been set.

### GetLongContext

`func (o *UpdateAIAgentRequest) GetLongContext() bool`

GetLongContext returns the LongContext field if non-nil, zero value otherwise.

### GetLongContextOk

`func (o *UpdateAIAgentRequest) GetLongContextOk() (*bool, bool)`

GetLongContextOk returns a tuple with the LongContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongContext

`func (o *UpdateAIAgentRequest) SetLongContext(v bool)`

SetLongContext sets LongContext field to given value.

### HasLongContext

`func (o *UpdateAIAgentRequest) HasLongContext() bool`

HasLongContext returns a boolean if a field has been set.

### GetGuardrailPreset

`func (o *UpdateAIAgentRequest) GetGuardrailPreset() string`

GetGuardrailPreset returns the GuardrailPreset field if non-nil, zero value otherwise.

### GetGuardrailPresetOk

`func (o *UpdateAIAgentRequest) GetGuardrailPresetOk() (*string, bool)`

GetGuardrailPresetOk returns a tuple with the GuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailPreset

`func (o *UpdateAIAgentRequest) SetGuardrailPreset(v string)`

SetGuardrailPreset sets GuardrailPreset field to given value.

### HasGuardrailPreset

`func (o *UpdateAIAgentRequest) HasGuardrailPreset() bool`

HasGuardrailPreset returns a boolean if a field has been set.

### GetFilterPolicies

`func (o *UpdateAIAgentRequest) GetFilterPolicies() []string`

GetFilterPolicies returns the FilterPolicies field if non-nil, zero value otherwise.

### GetFilterPoliciesOk

`func (o *UpdateAIAgentRequest) GetFilterPoliciesOk() (*[]string, bool)`

GetFilterPoliciesOk returns a tuple with the FilterPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPolicies

`func (o *UpdateAIAgentRequest) SetFilterPolicies(v []string)`

SetFilterPolicies sets FilterPolicies field to given value.

### HasFilterPolicies

`func (o *UpdateAIAgentRequest) HasFilterPolicies() bool`

HasFilterPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


