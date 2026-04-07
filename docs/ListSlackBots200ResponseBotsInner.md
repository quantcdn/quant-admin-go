# ListSlackBots200ResponseBotsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SetupType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**SystemPrompt** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**AssignedSkills** | Pointer to **[]string** |  | [optional] 
**AllowedCollections** | Pointer to **[]string** |  | [optional] 
**AllowedSubAgents** | Pointer to **[]string** |  | [optional] 
**GuardrailPreset** | Pointer to **string** |  | [optional] 
**FilterPolicies** | Pointer to **[]string** |  | [optional] 
**LongContext** | Pointer to **bool** |  | [optional] 
**SessionTtlDays** | Pointer to **int32** |  | [optional] 
**KeywordsEnabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListSlackBots200ResponseBotsInner

`func NewListSlackBots200ResponseBotsInner() *ListSlackBots200ResponseBotsInner`

NewListSlackBots200ResponseBotsInner instantiates a new ListSlackBots200ResponseBotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSlackBots200ResponseBotsInnerWithDefaults

`func NewListSlackBots200ResponseBotsInnerWithDefaults() *ListSlackBots200ResponseBotsInner`

NewListSlackBots200ResponseBotsInnerWithDefaults instantiates a new ListSlackBots200ResponseBotsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotId

`func (o *ListSlackBots200ResponseBotsInner) GetBotId() string`

GetBotId returns the BotId field if non-nil, zero value otherwise.

### GetBotIdOk

`func (o *ListSlackBots200ResponseBotsInner) GetBotIdOk() (*string, bool)`

GetBotIdOk returns a tuple with the BotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotId

`func (o *ListSlackBots200ResponseBotsInner) SetBotId(v string)`

SetBotId sets BotId field to given value.

### HasBotId

`func (o *ListSlackBots200ResponseBotsInner) HasBotId() bool`

HasBotId returns a boolean if a field has been set.

### GetName

`func (o *ListSlackBots200ResponseBotsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSlackBots200ResponseBotsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSlackBots200ResponseBotsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSlackBots200ResponseBotsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSetupType

`func (o *ListSlackBots200ResponseBotsInner) GetSetupType() string`

GetSetupType returns the SetupType field if non-nil, zero value otherwise.

### GetSetupTypeOk

`func (o *ListSlackBots200ResponseBotsInner) GetSetupTypeOk() (*string, bool)`

GetSetupTypeOk returns a tuple with the SetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupType

`func (o *ListSlackBots200ResponseBotsInner) SetSetupType(v string)`

SetSetupType sets SetupType field to given value.

### HasSetupType

`func (o *ListSlackBots200ResponseBotsInner) HasSetupType() bool`

HasSetupType returns a boolean if a field has been set.

### GetStatus

`func (o *ListSlackBots200ResponseBotsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSlackBots200ResponseBotsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSlackBots200ResponseBotsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSlackBots200ResponseBotsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnected

`func (o *ListSlackBots200ResponseBotsInner) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ListSlackBots200ResponseBotsInner) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ListSlackBots200ResponseBotsInner) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ListSlackBots200ResponseBotsInner) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *ListSlackBots200ResponseBotsInner) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *ListSlackBots200ResponseBotsInner) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *ListSlackBots200ResponseBotsInner) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *ListSlackBots200ResponseBotsInner) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetModelId

`func (o *ListSlackBots200ResponseBotsInner) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ListSlackBots200ResponseBotsInner) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ListSlackBots200ResponseBotsInner) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ListSlackBots200ResponseBotsInner) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetTemperature

`func (o *ListSlackBots200ResponseBotsInner) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ListSlackBots200ResponseBotsInner) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ListSlackBots200ResponseBotsInner) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ListSlackBots200ResponseBotsInner) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *ListSlackBots200ResponseBotsInner) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ListSlackBots200ResponseBotsInner) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ListSlackBots200ResponseBotsInner) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ListSlackBots200ResponseBotsInner) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetAllowedTools

`func (o *ListSlackBots200ResponseBotsInner) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *ListSlackBots200ResponseBotsInner) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *ListSlackBots200ResponseBotsInner) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *ListSlackBots200ResponseBotsInner) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetAssignedSkills

`func (o *ListSlackBots200ResponseBotsInner) GetAssignedSkills() []string`

GetAssignedSkills returns the AssignedSkills field if non-nil, zero value otherwise.

### GetAssignedSkillsOk

`func (o *ListSlackBots200ResponseBotsInner) GetAssignedSkillsOk() (*[]string, bool)`

GetAssignedSkillsOk returns a tuple with the AssignedSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedSkills

`func (o *ListSlackBots200ResponseBotsInner) SetAssignedSkills(v []string)`

SetAssignedSkills sets AssignedSkills field to given value.

### HasAssignedSkills

`func (o *ListSlackBots200ResponseBotsInner) HasAssignedSkills() bool`

HasAssignedSkills returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *ListSlackBots200ResponseBotsInner) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *ListSlackBots200ResponseBotsInner) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *ListSlackBots200ResponseBotsInner) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *ListSlackBots200ResponseBotsInner) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetAllowedSubAgents

`func (o *ListSlackBots200ResponseBotsInner) GetAllowedSubAgents() []string`

GetAllowedSubAgents returns the AllowedSubAgents field if non-nil, zero value otherwise.

### GetAllowedSubAgentsOk

`func (o *ListSlackBots200ResponseBotsInner) GetAllowedSubAgentsOk() (*[]string, bool)`

GetAllowedSubAgentsOk returns a tuple with the AllowedSubAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSubAgents

`func (o *ListSlackBots200ResponseBotsInner) SetAllowedSubAgents(v []string)`

SetAllowedSubAgents sets AllowedSubAgents field to given value.

### HasAllowedSubAgents

`func (o *ListSlackBots200ResponseBotsInner) HasAllowedSubAgents() bool`

HasAllowedSubAgents returns a boolean if a field has been set.

### GetGuardrailPreset

`func (o *ListSlackBots200ResponseBotsInner) GetGuardrailPreset() string`

GetGuardrailPreset returns the GuardrailPreset field if non-nil, zero value otherwise.

### GetGuardrailPresetOk

`func (o *ListSlackBots200ResponseBotsInner) GetGuardrailPresetOk() (*string, bool)`

GetGuardrailPresetOk returns a tuple with the GuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailPreset

`func (o *ListSlackBots200ResponseBotsInner) SetGuardrailPreset(v string)`

SetGuardrailPreset sets GuardrailPreset field to given value.

### HasGuardrailPreset

`func (o *ListSlackBots200ResponseBotsInner) HasGuardrailPreset() bool`

HasGuardrailPreset returns a boolean if a field has been set.

### GetFilterPolicies

`func (o *ListSlackBots200ResponseBotsInner) GetFilterPolicies() []string`

GetFilterPolicies returns the FilterPolicies field if non-nil, zero value otherwise.

### GetFilterPoliciesOk

`func (o *ListSlackBots200ResponseBotsInner) GetFilterPoliciesOk() (*[]string, bool)`

GetFilterPoliciesOk returns a tuple with the FilterPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPolicies

`func (o *ListSlackBots200ResponseBotsInner) SetFilterPolicies(v []string)`

SetFilterPolicies sets FilterPolicies field to given value.

### HasFilterPolicies

`func (o *ListSlackBots200ResponseBotsInner) HasFilterPolicies() bool`

HasFilterPolicies returns a boolean if a field has been set.

### GetLongContext

`func (o *ListSlackBots200ResponseBotsInner) GetLongContext() bool`

GetLongContext returns the LongContext field if non-nil, zero value otherwise.

### GetLongContextOk

`func (o *ListSlackBots200ResponseBotsInner) GetLongContextOk() (*bool, bool)`

GetLongContextOk returns a tuple with the LongContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongContext

`func (o *ListSlackBots200ResponseBotsInner) SetLongContext(v bool)`

SetLongContext sets LongContext field to given value.

### HasLongContext

`func (o *ListSlackBots200ResponseBotsInner) HasLongContext() bool`

HasLongContext returns a boolean if a field has been set.

### GetSessionTtlDays

`func (o *ListSlackBots200ResponseBotsInner) GetSessionTtlDays() int32`

GetSessionTtlDays returns the SessionTtlDays field if non-nil, zero value otherwise.

### GetSessionTtlDaysOk

`func (o *ListSlackBots200ResponseBotsInner) GetSessionTtlDaysOk() (*int32, bool)`

GetSessionTtlDaysOk returns a tuple with the SessionTtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTtlDays

`func (o *ListSlackBots200ResponseBotsInner) SetSessionTtlDays(v int32)`

SetSessionTtlDays sets SessionTtlDays field to given value.

### HasSessionTtlDays

`func (o *ListSlackBots200ResponseBotsInner) HasSessionTtlDays() bool`

HasSessionTtlDays returns a boolean if a field has been set.

### GetKeywordsEnabled

`func (o *ListSlackBots200ResponseBotsInner) GetKeywordsEnabled() bool`

GetKeywordsEnabled returns the KeywordsEnabled field if non-nil, zero value otherwise.

### GetKeywordsEnabledOk

`func (o *ListSlackBots200ResponseBotsInner) GetKeywordsEnabledOk() (*bool, bool)`

GetKeywordsEnabledOk returns a tuple with the KeywordsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywordsEnabled

`func (o *ListSlackBots200ResponseBotsInner) SetKeywordsEnabled(v bool)`

SetKeywordsEnabled sets KeywordsEnabled field to given value.

### HasKeywordsEnabled

`func (o *ListSlackBots200ResponseBotsInner) HasKeywordsEnabled() bool`

HasKeywordsEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListSlackBots200ResponseBotsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListSlackBots200ResponseBotsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListSlackBots200ResponseBotsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListSlackBots200ResponseBotsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


