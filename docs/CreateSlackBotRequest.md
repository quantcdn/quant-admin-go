# CreateSlackBotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the bot | 
**SetupType** | **string** | Whether to use Quant-managed or customer-provided Slack app | 
**SystemPrompt** | **string** | System prompt for the backing AI agent | 
**ModelId** | **string** | AI model identifier | 
**Temperature** | Pointer to **float32** | Sampling temperature | [optional] 
**MaxTokens** | Pointer to **int32** | Maximum response tokens | [optional] 
**AllowedTools** | Pointer to **[]string** | Tools the agent may use | [optional] 
**AssignedSkills** | Pointer to **[]string** | Skills assigned to the agent | [optional] 
**AllowedCollections** | Pointer to **[]string** | Vector DB collections the agent may query | [optional] 
**AllowedSubAgents** | Pointer to **[]string** | Sub-agents the agent may call | [optional] 
**GuardrailPreset** | Pointer to **string** | Guardrail preset name | [optional] 
**FilterPolicies** | Pointer to **[]string** | Content filter policies | [optional] 
**LongContext** | Pointer to **bool** | Enable long context mode | [optional] 
**SessionTtlDays** | Pointer to **int32** | Session TTL in days | [optional] 
**AllowedChannels** | Pointer to **[]string** | Slack channel IDs the bot may respond in | [optional] 
**AllowedUsers** | Pointer to **[]string** | Slack user IDs allowed to interact with the bot | [optional] 
**DeniedUsers** | Pointer to **[]string** | Slack user IDs denied from interacting with the bot | [optional] 
**AllowGuests** | Pointer to **bool** | Whether guest users may interact with the bot | [optional] 
**HomeTabContent** | Pointer to **string** | Content shown on the bot&#39;s Home tab in Slack | [optional] 
**AgentAccessControl** | Pointer to **map[string]interface{}** | Agent-level access control settings | [optional] 
**KeywordsEnabled** | Pointer to **bool** | Whether keyword triggers are enabled | [optional] 
**Keywords** | Pointer to **[]string** | Keywords that trigger the bot | [optional] 

## Methods

### NewCreateSlackBotRequest

`func NewCreateSlackBotRequest(name string, setupType string, systemPrompt string, modelId string, ) *CreateSlackBotRequest`

NewCreateSlackBotRequest instantiates a new CreateSlackBotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSlackBotRequestWithDefaults

`func NewCreateSlackBotRequestWithDefaults() *CreateSlackBotRequest`

NewCreateSlackBotRequestWithDefaults instantiates a new CreateSlackBotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSlackBotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSlackBotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSlackBotRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSetupType

`func (o *CreateSlackBotRequest) GetSetupType() string`

GetSetupType returns the SetupType field if non-nil, zero value otherwise.

### GetSetupTypeOk

`func (o *CreateSlackBotRequest) GetSetupTypeOk() (*string, bool)`

GetSetupTypeOk returns a tuple with the SetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupType

`func (o *CreateSlackBotRequest) SetSetupType(v string)`

SetSetupType sets SetupType field to given value.


### GetSystemPrompt

`func (o *CreateSlackBotRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *CreateSlackBotRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *CreateSlackBotRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.


### GetModelId

`func (o *CreateSlackBotRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateSlackBotRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateSlackBotRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetTemperature

`func (o *CreateSlackBotRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *CreateSlackBotRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *CreateSlackBotRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *CreateSlackBotRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *CreateSlackBotRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *CreateSlackBotRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *CreateSlackBotRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *CreateSlackBotRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetAllowedTools

`func (o *CreateSlackBotRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *CreateSlackBotRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *CreateSlackBotRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *CreateSlackBotRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetAssignedSkills

`func (o *CreateSlackBotRequest) GetAssignedSkills() []string`

GetAssignedSkills returns the AssignedSkills field if non-nil, zero value otherwise.

### GetAssignedSkillsOk

`func (o *CreateSlackBotRequest) GetAssignedSkillsOk() (*[]string, bool)`

GetAssignedSkillsOk returns a tuple with the AssignedSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedSkills

`func (o *CreateSlackBotRequest) SetAssignedSkills(v []string)`

SetAssignedSkills sets AssignedSkills field to given value.

### HasAssignedSkills

`func (o *CreateSlackBotRequest) HasAssignedSkills() bool`

HasAssignedSkills returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *CreateSlackBotRequest) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *CreateSlackBotRequest) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *CreateSlackBotRequest) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *CreateSlackBotRequest) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetAllowedSubAgents

`func (o *CreateSlackBotRequest) GetAllowedSubAgents() []string`

GetAllowedSubAgents returns the AllowedSubAgents field if non-nil, zero value otherwise.

### GetAllowedSubAgentsOk

`func (o *CreateSlackBotRequest) GetAllowedSubAgentsOk() (*[]string, bool)`

GetAllowedSubAgentsOk returns a tuple with the AllowedSubAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSubAgents

`func (o *CreateSlackBotRequest) SetAllowedSubAgents(v []string)`

SetAllowedSubAgents sets AllowedSubAgents field to given value.

### HasAllowedSubAgents

`func (o *CreateSlackBotRequest) HasAllowedSubAgents() bool`

HasAllowedSubAgents returns a boolean if a field has been set.

### GetGuardrailPreset

`func (o *CreateSlackBotRequest) GetGuardrailPreset() string`

GetGuardrailPreset returns the GuardrailPreset field if non-nil, zero value otherwise.

### GetGuardrailPresetOk

`func (o *CreateSlackBotRequest) GetGuardrailPresetOk() (*string, bool)`

GetGuardrailPresetOk returns a tuple with the GuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailPreset

`func (o *CreateSlackBotRequest) SetGuardrailPreset(v string)`

SetGuardrailPreset sets GuardrailPreset field to given value.

### HasGuardrailPreset

`func (o *CreateSlackBotRequest) HasGuardrailPreset() bool`

HasGuardrailPreset returns a boolean if a field has been set.

### GetFilterPolicies

`func (o *CreateSlackBotRequest) GetFilterPolicies() []string`

GetFilterPolicies returns the FilterPolicies field if non-nil, zero value otherwise.

### GetFilterPoliciesOk

`func (o *CreateSlackBotRequest) GetFilterPoliciesOk() (*[]string, bool)`

GetFilterPoliciesOk returns a tuple with the FilterPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPolicies

`func (o *CreateSlackBotRequest) SetFilterPolicies(v []string)`

SetFilterPolicies sets FilterPolicies field to given value.

### HasFilterPolicies

`func (o *CreateSlackBotRequest) HasFilterPolicies() bool`

HasFilterPolicies returns a boolean if a field has been set.

### GetLongContext

`func (o *CreateSlackBotRequest) GetLongContext() bool`

GetLongContext returns the LongContext field if non-nil, zero value otherwise.

### GetLongContextOk

`func (o *CreateSlackBotRequest) GetLongContextOk() (*bool, bool)`

GetLongContextOk returns a tuple with the LongContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongContext

`func (o *CreateSlackBotRequest) SetLongContext(v bool)`

SetLongContext sets LongContext field to given value.

### HasLongContext

`func (o *CreateSlackBotRequest) HasLongContext() bool`

HasLongContext returns a boolean if a field has been set.

### GetSessionTtlDays

`func (o *CreateSlackBotRequest) GetSessionTtlDays() int32`

GetSessionTtlDays returns the SessionTtlDays field if non-nil, zero value otherwise.

### GetSessionTtlDaysOk

`func (o *CreateSlackBotRequest) GetSessionTtlDaysOk() (*int32, bool)`

GetSessionTtlDaysOk returns a tuple with the SessionTtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTtlDays

`func (o *CreateSlackBotRequest) SetSessionTtlDays(v int32)`

SetSessionTtlDays sets SessionTtlDays field to given value.

### HasSessionTtlDays

`func (o *CreateSlackBotRequest) HasSessionTtlDays() bool`

HasSessionTtlDays returns a boolean if a field has been set.

### GetAllowedChannels

`func (o *CreateSlackBotRequest) GetAllowedChannels() []string`

GetAllowedChannels returns the AllowedChannels field if non-nil, zero value otherwise.

### GetAllowedChannelsOk

`func (o *CreateSlackBotRequest) GetAllowedChannelsOk() (*[]string, bool)`

GetAllowedChannelsOk returns a tuple with the AllowedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChannels

`func (o *CreateSlackBotRequest) SetAllowedChannels(v []string)`

SetAllowedChannels sets AllowedChannels field to given value.

### HasAllowedChannels

`func (o *CreateSlackBotRequest) HasAllowedChannels() bool`

HasAllowedChannels returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *CreateSlackBotRequest) GetAllowedUsers() []string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *CreateSlackBotRequest) GetAllowedUsersOk() (*[]string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *CreateSlackBotRequest) SetAllowedUsers(v []string)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *CreateSlackBotRequest) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetDeniedUsers

`func (o *CreateSlackBotRequest) GetDeniedUsers() []string`

GetDeniedUsers returns the DeniedUsers field if non-nil, zero value otherwise.

### GetDeniedUsersOk

`func (o *CreateSlackBotRequest) GetDeniedUsersOk() (*[]string, bool)`

GetDeniedUsersOk returns a tuple with the DeniedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedUsers

`func (o *CreateSlackBotRequest) SetDeniedUsers(v []string)`

SetDeniedUsers sets DeniedUsers field to given value.

### HasDeniedUsers

`func (o *CreateSlackBotRequest) HasDeniedUsers() bool`

HasDeniedUsers returns a boolean if a field has been set.

### GetAllowGuests

`func (o *CreateSlackBotRequest) GetAllowGuests() bool`

GetAllowGuests returns the AllowGuests field if non-nil, zero value otherwise.

### GetAllowGuestsOk

`func (o *CreateSlackBotRequest) GetAllowGuestsOk() (*bool, bool)`

GetAllowGuestsOk returns a tuple with the AllowGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuests

`func (o *CreateSlackBotRequest) SetAllowGuests(v bool)`

SetAllowGuests sets AllowGuests field to given value.

### HasAllowGuests

`func (o *CreateSlackBotRequest) HasAllowGuests() bool`

HasAllowGuests returns a boolean if a field has been set.

### GetHomeTabContent

`func (o *CreateSlackBotRequest) GetHomeTabContent() string`

GetHomeTabContent returns the HomeTabContent field if non-nil, zero value otherwise.

### GetHomeTabContentOk

`func (o *CreateSlackBotRequest) GetHomeTabContentOk() (*string, bool)`

GetHomeTabContentOk returns a tuple with the HomeTabContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeTabContent

`func (o *CreateSlackBotRequest) SetHomeTabContent(v string)`

SetHomeTabContent sets HomeTabContent field to given value.

### HasHomeTabContent

`func (o *CreateSlackBotRequest) HasHomeTabContent() bool`

HasHomeTabContent returns a boolean if a field has been set.

### GetAgentAccessControl

`func (o *CreateSlackBotRequest) GetAgentAccessControl() map[string]interface{}`

GetAgentAccessControl returns the AgentAccessControl field if non-nil, zero value otherwise.

### GetAgentAccessControlOk

`func (o *CreateSlackBotRequest) GetAgentAccessControlOk() (*map[string]interface{}, bool)`

GetAgentAccessControlOk returns a tuple with the AgentAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAccessControl

`func (o *CreateSlackBotRequest) SetAgentAccessControl(v map[string]interface{})`

SetAgentAccessControl sets AgentAccessControl field to given value.

### HasAgentAccessControl

`func (o *CreateSlackBotRequest) HasAgentAccessControl() bool`

HasAgentAccessControl returns a boolean if a field has been set.

### GetKeywordsEnabled

`func (o *CreateSlackBotRequest) GetKeywordsEnabled() bool`

GetKeywordsEnabled returns the KeywordsEnabled field if non-nil, zero value otherwise.

### GetKeywordsEnabledOk

`func (o *CreateSlackBotRequest) GetKeywordsEnabledOk() (*bool, bool)`

GetKeywordsEnabledOk returns a tuple with the KeywordsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywordsEnabled

`func (o *CreateSlackBotRequest) SetKeywordsEnabled(v bool)`

SetKeywordsEnabled sets KeywordsEnabled field to given value.

### HasKeywordsEnabled

`func (o *CreateSlackBotRequest) HasKeywordsEnabled() bool`

HasKeywordsEnabled returns a boolean if a field has been set.

### GetKeywords

`func (o *CreateSlackBotRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CreateSlackBotRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CreateSlackBotRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *CreateSlackBotRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


