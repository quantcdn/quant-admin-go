# UpdateSlackBotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Display name for the bot | [optional] 
**Status** | Pointer to **string** | Enable or disable the bot | [optional] 
**SystemPrompt** | Pointer to **string** | System prompt for the backing AI agent | [optional] 
**ModelId** | Pointer to **string** | AI model identifier | [optional] 
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

### NewUpdateSlackBotRequest

`func NewUpdateSlackBotRequest() *UpdateSlackBotRequest`

NewUpdateSlackBotRequest instantiates a new UpdateSlackBotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSlackBotRequestWithDefaults

`func NewUpdateSlackBotRequestWithDefaults() *UpdateSlackBotRequest`

NewUpdateSlackBotRequestWithDefaults instantiates a new UpdateSlackBotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSlackBotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSlackBotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSlackBotRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSlackBotRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateSlackBotRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateSlackBotRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateSlackBotRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateSlackBotRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *UpdateSlackBotRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *UpdateSlackBotRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *UpdateSlackBotRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *UpdateSlackBotRequest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetModelId

`func (o *UpdateSlackBotRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *UpdateSlackBotRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *UpdateSlackBotRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *UpdateSlackBotRequest) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetTemperature

`func (o *UpdateSlackBotRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *UpdateSlackBotRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *UpdateSlackBotRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *UpdateSlackBotRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *UpdateSlackBotRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *UpdateSlackBotRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *UpdateSlackBotRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *UpdateSlackBotRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetAllowedTools

`func (o *UpdateSlackBotRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *UpdateSlackBotRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *UpdateSlackBotRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *UpdateSlackBotRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetAssignedSkills

`func (o *UpdateSlackBotRequest) GetAssignedSkills() []string`

GetAssignedSkills returns the AssignedSkills field if non-nil, zero value otherwise.

### GetAssignedSkillsOk

`func (o *UpdateSlackBotRequest) GetAssignedSkillsOk() (*[]string, bool)`

GetAssignedSkillsOk returns a tuple with the AssignedSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedSkills

`func (o *UpdateSlackBotRequest) SetAssignedSkills(v []string)`

SetAssignedSkills sets AssignedSkills field to given value.

### HasAssignedSkills

`func (o *UpdateSlackBotRequest) HasAssignedSkills() bool`

HasAssignedSkills returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *UpdateSlackBotRequest) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *UpdateSlackBotRequest) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *UpdateSlackBotRequest) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *UpdateSlackBotRequest) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetAllowedSubAgents

`func (o *UpdateSlackBotRequest) GetAllowedSubAgents() []string`

GetAllowedSubAgents returns the AllowedSubAgents field if non-nil, zero value otherwise.

### GetAllowedSubAgentsOk

`func (o *UpdateSlackBotRequest) GetAllowedSubAgentsOk() (*[]string, bool)`

GetAllowedSubAgentsOk returns a tuple with the AllowedSubAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSubAgents

`func (o *UpdateSlackBotRequest) SetAllowedSubAgents(v []string)`

SetAllowedSubAgents sets AllowedSubAgents field to given value.

### HasAllowedSubAgents

`func (o *UpdateSlackBotRequest) HasAllowedSubAgents() bool`

HasAllowedSubAgents returns a boolean if a field has been set.

### GetGuardrailPreset

`func (o *UpdateSlackBotRequest) GetGuardrailPreset() string`

GetGuardrailPreset returns the GuardrailPreset field if non-nil, zero value otherwise.

### GetGuardrailPresetOk

`func (o *UpdateSlackBotRequest) GetGuardrailPresetOk() (*string, bool)`

GetGuardrailPresetOk returns a tuple with the GuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailPreset

`func (o *UpdateSlackBotRequest) SetGuardrailPreset(v string)`

SetGuardrailPreset sets GuardrailPreset field to given value.

### HasGuardrailPreset

`func (o *UpdateSlackBotRequest) HasGuardrailPreset() bool`

HasGuardrailPreset returns a boolean if a field has been set.

### GetFilterPolicies

`func (o *UpdateSlackBotRequest) GetFilterPolicies() []string`

GetFilterPolicies returns the FilterPolicies field if non-nil, zero value otherwise.

### GetFilterPoliciesOk

`func (o *UpdateSlackBotRequest) GetFilterPoliciesOk() (*[]string, bool)`

GetFilterPoliciesOk returns a tuple with the FilterPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPolicies

`func (o *UpdateSlackBotRequest) SetFilterPolicies(v []string)`

SetFilterPolicies sets FilterPolicies field to given value.

### HasFilterPolicies

`func (o *UpdateSlackBotRequest) HasFilterPolicies() bool`

HasFilterPolicies returns a boolean if a field has been set.

### GetLongContext

`func (o *UpdateSlackBotRequest) GetLongContext() bool`

GetLongContext returns the LongContext field if non-nil, zero value otherwise.

### GetLongContextOk

`func (o *UpdateSlackBotRequest) GetLongContextOk() (*bool, bool)`

GetLongContextOk returns a tuple with the LongContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongContext

`func (o *UpdateSlackBotRequest) SetLongContext(v bool)`

SetLongContext sets LongContext field to given value.

### HasLongContext

`func (o *UpdateSlackBotRequest) HasLongContext() bool`

HasLongContext returns a boolean if a field has been set.

### GetSessionTtlDays

`func (o *UpdateSlackBotRequest) GetSessionTtlDays() int32`

GetSessionTtlDays returns the SessionTtlDays field if non-nil, zero value otherwise.

### GetSessionTtlDaysOk

`func (o *UpdateSlackBotRequest) GetSessionTtlDaysOk() (*int32, bool)`

GetSessionTtlDaysOk returns a tuple with the SessionTtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTtlDays

`func (o *UpdateSlackBotRequest) SetSessionTtlDays(v int32)`

SetSessionTtlDays sets SessionTtlDays field to given value.

### HasSessionTtlDays

`func (o *UpdateSlackBotRequest) HasSessionTtlDays() bool`

HasSessionTtlDays returns a boolean if a field has been set.

### GetAllowedChannels

`func (o *UpdateSlackBotRequest) GetAllowedChannels() []string`

GetAllowedChannels returns the AllowedChannels field if non-nil, zero value otherwise.

### GetAllowedChannelsOk

`func (o *UpdateSlackBotRequest) GetAllowedChannelsOk() (*[]string, bool)`

GetAllowedChannelsOk returns a tuple with the AllowedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChannels

`func (o *UpdateSlackBotRequest) SetAllowedChannels(v []string)`

SetAllowedChannels sets AllowedChannels field to given value.

### HasAllowedChannels

`func (o *UpdateSlackBotRequest) HasAllowedChannels() bool`

HasAllowedChannels returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *UpdateSlackBotRequest) GetAllowedUsers() []string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *UpdateSlackBotRequest) GetAllowedUsersOk() (*[]string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *UpdateSlackBotRequest) SetAllowedUsers(v []string)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *UpdateSlackBotRequest) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetDeniedUsers

`func (o *UpdateSlackBotRequest) GetDeniedUsers() []string`

GetDeniedUsers returns the DeniedUsers field if non-nil, zero value otherwise.

### GetDeniedUsersOk

`func (o *UpdateSlackBotRequest) GetDeniedUsersOk() (*[]string, bool)`

GetDeniedUsersOk returns a tuple with the DeniedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedUsers

`func (o *UpdateSlackBotRequest) SetDeniedUsers(v []string)`

SetDeniedUsers sets DeniedUsers field to given value.

### HasDeniedUsers

`func (o *UpdateSlackBotRequest) HasDeniedUsers() bool`

HasDeniedUsers returns a boolean if a field has been set.

### GetAllowGuests

`func (o *UpdateSlackBotRequest) GetAllowGuests() bool`

GetAllowGuests returns the AllowGuests field if non-nil, zero value otherwise.

### GetAllowGuestsOk

`func (o *UpdateSlackBotRequest) GetAllowGuestsOk() (*bool, bool)`

GetAllowGuestsOk returns a tuple with the AllowGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuests

`func (o *UpdateSlackBotRequest) SetAllowGuests(v bool)`

SetAllowGuests sets AllowGuests field to given value.

### HasAllowGuests

`func (o *UpdateSlackBotRequest) HasAllowGuests() bool`

HasAllowGuests returns a boolean if a field has been set.

### GetHomeTabContent

`func (o *UpdateSlackBotRequest) GetHomeTabContent() string`

GetHomeTabContent returns the HomeTabContent field if non-nil, zero value otherwise.

### GetHomeTabContentOk

`func (o *UpdateSlackBotRequest) GetHomeTabContentOk() (*string, bool)`

GetHomeTabContentOk returns a tuple with the HomeTabContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeTabContent

`func (o *UpdateSlackBotRequest) SetHomeTabContent(v string)`

SetHomeTabContent sets HomeTabContent field to given value.

### HasHomeTabContent

`func (o *UpdateSlackBotRequest) HasHomeTabContent() bool`

HasHomeTabContent returns a boolean if a field has been set.

### GetAgentAccessControl

`func (o *UpdateSlackBotRequest) GetAgentAccessControl() map[string]interface{}`

GetAgentAccessControl returns the AgentAccessControl field if non-nil, zero value otherwise.

### GetAgentAccessControlOk

`func (o *UpdateSlackBotRequest) GetAgentAccessControlOk() (*map[string]interface{}, bool)`

GetAgentAccessControlOk returns a tuple with the AgentAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAccessControl

`func (o *UpdateSlackBotRequest) SetAgentAccessControl(v map[string]interface{})`

SetAgentAccessControl sets AgentAccessControl field to given value.

### HasAgentAccessControl

`func (o *UpdateSlackBotRequest) HasAgentAccessControl() bool`

HasAgentAccessControl returns a boolean if a field has been set.

### GetKeywordsEnabled

`func (o *UpdateSlackBotRequest) GetKeywordsEnabled() bool`

GetKeywordsEnabled returns the KeywordsEnabled field if non-nil, zero value otherwise.

### GetKeywordsEnabledOk

`func (o *UpdateSlackBotRequest) GetKeywordsEnabledOk() (*bool, bool)`

GetKeywordsEnabledOk returns a tuple with the KeywordsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywordsEnabled

`func (o *UpdateSlackBotRequest) SetKeywordsEnabled(v bool)`

SetKeywordsEnabled sets KeywordsEnabled field to given value.

### HasKeywordsEnabled

`func (o *UpdateSlackBotRequest) HasKeywordsEnabled() bool`

HasKeywordsEnabled returns a boolean if a field has been set.

### GetKeywords

`func (o *UpdateSlackBotRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *UpdateSlackBotRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *UpdateSlackBotRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *UpdateSlackBotRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


