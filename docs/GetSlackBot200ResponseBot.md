# GetSlackBot200ResponseBot

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
**AllowedChannels** | Pointer to **[]string** |  | [optional] 
**AllowedUsers** | Pointer to **[]string** |  | [optional] 
**DeniedUsers** | Pointer to **[]string** |  | [optional] 
**AllowGuests** | Pointer to **bool** |  | [optional] 
**HomeTabContent** | Pointer to **string** |  | [optional] 
**AgentAccessControl** | Pointer to **map[string]interface{}** |  | [optional] 
**KeywordsEnabled** | Pointer to **bool** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetSlackBot200ResponseBot

`func NewGetSlackBot200ResponseBot() *GetSlackBot200ResponseBot`

NewGetSlackBot200ResponseBot instantiates a new GetSlackBot200ResponseBot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSlackBot200ResponseBotWithDefaults

`func NewGetSlackBot200ResponseBotWithDefaults() *GetSlackBot200ResponseBot`

NewGetSlackBot200ResponseBotWithDefaults instantiates a new GetSlackBot200ResponseBot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotId

`func (o *GetSlackBot200ResponseBot) GetBotId() string`

GetBotId returns the BotId field if non-nil, zero value otherwise.

### GetBotIdOk

`func (o *GetSlackBot200ResponseBot) GetBotIdOk() (*string, bool)`

GetBotIdOk returns a tuple with the BotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotId

`func (o *GetSlackBot200ResponseBot) SetBotId(v string)`

SetBotId sets BotId field to given value.

### HasBotId

`func (o *GetSlackBot200ResponseBot) HasBotId() bool`

HasBotId returns a boolean if a field has been set.

### GetName

`func (o *GetSlackBot200ResponseBot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSlackBot200ResponseBot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSlackBot200ResponseBot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSlackBot200ResponseBot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSetupType

`func (o *GetSlackBot200ResponseBot) GetSetupType() string`

GetSetupType returns the SetupType field if non-nil, zero value otherwise.

### GetSetupTypeOk

`func (o *GetSlackBot200ResponseBot) GetSetupTypeOk() (*string, bool)`

GetSetupTypeOk returns a tuple with the SetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupType

`func (o *GetSlackBot200ResponseBot) SetSetupType(v string)`

SetSetupType sets SetupType field to given value.

### HasSetupType

`func (o *GetSlackBot200ResponseBot) HasSetupType() bool`

HasSetupType returns a boolean if a field has been set.

### GetStatus

`func (o *GetSlackBot200ResponseBot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSlackBot200ResponseBot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSlackBot200ResponseBot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSlackBot200ResponseBot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnected

`func (o *GetSlackBot200ResponseBot) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *GetSlackBot200ResponseBot) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *GetSlackBot200ResponseBot) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *GetSlackBot200ResponseBot) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *GetSlackBot200ResponseBot) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *GetSlackBot200ResponseBot) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *GetSlackBot200ResponseBot) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *GetSlackBot200ResponseBot) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetModelId

`func (o *GetSlackBot200ResponseBot) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *GetSlackBot200ResponseBot) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *GetSlackBot200ResponseBot) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *GetSlackBot200ResponseBot) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetTemperature

`func (o *GetSlackBot200ResponseBot) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetSlackBot200ResponseBot) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetSlackBot200ResponseBot) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetSlackBot200ResponseBot) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *GetSlackBot200ResponseBot) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *GetSlackBot200ResponseBot) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *GetSlackBot200ResponseBot) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *GetSlackBot200ResponseBot) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetAllowedTools

`func (o *GetSlackBot200ResponseBot) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *GetSlackBot200ResponseBot) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *GetSlackBot200ResponseBot) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *GetSlackBot200ResponseBot) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetAssignedSkills

`func (o *GetSlackBot200ResponseBot) GetAssignedSkills() []string`

GetAssignedSkills returns the AssignedSkills field if non-nil, zero value otherwise.

### GetAssignedSkillsOk

`func (o *GetSlackBot200ResponseBot) GetAssignedSkillsOk() (*[]string, bool)`

GetAssignedSkillsOk returns a tuple with the AssignedSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedSkills

`func (o *GetSlackBot200ResponseBot) SetAssignedSkills(v []string)`

SetAssignedSkills sets AssignedSkills field to given value.

### HasAssignedSkills

`func (o *GetSlackBot200ResponseBot) HasAssignedSkills() bool`

HasAssignedSkills returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *GetSlackBot200ResponseBot) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *GetSlackBot200ResponseBot) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *GetSlackBot200ResponseBot) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *GetSlackBot200ResponseBot) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetAllowedSubAgents

`func (o *GetSlackBot200ResponseBot) GetAllowedSubAgents() []string`

GetAllowedSubAgents returns the AllowedSubAgents field if non-nil, zero value otherwise.

### GetAllowedSubAgentsOk

`func (o *GetSlackBot200ResponseBot) GetAllowedSubAgentsOk() (*[]string, bool)`

GetAllowedSubAgentsOk returns a tuple with the AllowedSubAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSubAgents

`func (o *GetSlackBot200ResponseBot) SetAllowedSubAgents(v []string)`

SetAllowedSubAgents sets AllowedSubAgents field to given value.

### HasAllowedSubAgents

`func (o *GetSlackBot200ResponseBot) HasAllowedSubAgents() bool`

HasAllowedSubAgents returns a boolean if a field has been set.

### GetGuardrailPreset

`func (o *GetSlackBot200ResponseBot) GetGuardrailPreset() string`

GetGuardrailPreset returns the GuardrailPreset field if non-nil, zero value otherwise.

### GetGuardrailPresetOk

`func (o *GetSlackBot200ResponseBot) GetGuardrailPresetOk() (*string, bool)`

GetGuardrailPresetOk returns a tuple with the GuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailPreset

`func (o *GetSlackBot200ResponseBot) SetGuardrailPreset(v string)`

SetGuardrailPreset sets GuardrailPreset field to given value.

### HasGuardrailPreset

`func (o *GetSlackBot200ResponseBot) HasGuardrailPreset() bool`

HasGuardrailPreset returns a boolean if a field has been set.

### GetFilterPolicies

`func (o *GetSlackBot200ResponseBot) GetFilterPolicies() []string`

GetFilterPolicies returns the FilterPolicies field if non-nil, zero value otherwise.

### GetFilterPoliciesOk

`func (o *GetSlackBot200ResponseBot) GetFilterPoliciesOk() (*[]string, bool)`

GetFilterPoliciesOk returns a tuple with the FilterPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPolicies

`func (o *GetSlackBot200ResponseBot) SetFilterPolicies(v []string)`

SetFilterPolicies sets FilterPolicies field to given value.

### HasFilterPolicies

`func (o *GetSlackBot200ResponseBot) HasFilterPolicies() bool`

HasFilterPolicies returns a boolean if a field has been set.

### GetLongContext

`func (o *GetSlackBot200ResponseBot) GetLongContext() bool`

GetLongContext returns the LongContext field if non-nil, zero value otherwise.

### GetLongContextOk

`func (o *GetSlackBot200ResponseBot) GetLongContextOk() (*bool, bool)`

GetLongContextOk returns a tuple with the LongContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongContext

`func (o *GetSlackBot200ResponseBot) SetLongContext(v bool)`

SetLongContext sets LongContext field to given value.

### HasLongContext

`func (o *GetSlackBot200ResponseBot) HasLongContext() bool`

HasLongContext returns a boolean if a field has been set.

### GetSessionTtlDays

`func (o *GetSlackBot200ResponseBot) GetSessionTtlDays() int32`

GetSessionTtlDays returns the SessionTtlDays field if non-nil, zero value otherwise.

### GetSessionTtlDaysOk

`func (o *GetSlackBot200ResponseBot) GetSessionTtlDaysOk() (*int32, bool)`

GetSessionTtlDaysOk returns a tuple with the SessionTtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTtlDays

`func (o *GetSlackBot200ResponseBot) SetSessionTtlDays(v int32)`

SetSessionTtlDays sets SessionTtlDays field to given value.

### HasSessionTtlDays

`func (o *GetSlackBot200ResponseBot) HasSessionTtlDays() bool`

HasSessionTtlDays returns a boolean if a field has been set.

### GetAllowedChannels

`func (o *GetSlackBot200ResponseBot) GetAllowedChannels() []string`

GetAllowedChannels returns the AllowedChannels field if non-nil, zero value otherwise.

### GetAllowedChannelsOk

`func (o *GetSlackBot200ResponseBot) GetAllowedChannelsOk() (*[]string, bool)`

GetAllowedChannelsOk returns a tuple with the AllowedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedChannels

`func (o *GetSlackBot200ResponseBot) SetAllowedChannels(v []string)`

SetAllowedChannels sets AllowedChannels field to given value.

### HasAllowedChannels

`func (o *GetSlackBot200ResponseBot) HasAllowedChannels() bool`

HasAllowedChannels returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *GetSlackBot200ResponseBot) GetAllowedUsers() []string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *GetSlackBot200ResponseBot) GetAllowedUsersOk() (*[]string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *GetSlackBot200ResponseBot) SetAllowedUsers(v []string)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *GetSlackBot200ResponseBot) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetDeniedUsers

`func (o *GetSlackBot200ResponseBot) GetDeniedUsers() []string`

GetDeniedUsers returns the DeniedUsers field if non-nil, zero value otherwise.

### GetDeniedUsersOk

`func (o *GetSlackBot200ResponseBot) GetDeniedUsersOk() (*[]string, bool)`

GetDeniedUsersOk returns a tuple with the DeniedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedUsers

`func (o *GetSlackBot200ResponseBot) SetDeniedUsers(v []string)`

SetDeniedUsers sets DeniedUsers field to given value.

### HasDeniedUsers

`func (o *GetSlackBot200ResponseBot) HasDeniedUsers() bool`

HasDeniedUsers returns a boolean if a field has been set.

### GetAllowGuests

`func (o *GetSlackBot200ResponseBot) GetAllowGuests() bool`

GetAllowGuests returns the AllowGuests field if non-nil, zero value otherwise.

### GetAllowGuestsOk

`func (o *GetSlackBot200ResponseBot) GetAllowGuestsOk() (*bool, bool)`

GetAllowGuestsOk returns a tuple with the AllowGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuests

`func (o *GetSlackBot200ResponseBot) SetAllowGuests(v bool)`

SetAllowGuests sets AllowGuests field to given value.

### HasAllowGuests

`func (o *GetSlackBot200ResponseBot) HasAllowGuests() bool`

HasAllowGuests returns a boolean if a field has been set.

### GetHomeTabContent

`func (o *GetSlackBot200ResponseBot) GetHomeTabContent() string`

GetHomeTabContent returns the HomeTabContent field if non-nil, zero value otherwise.

### GetHomeTabContentOk

`func (o *GetSlackBot200ResponseBot) GetHomeTabContentOk() (*string, bool)`

GetHomeTabContentOk returns a tuple with the HomeTabContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeTabContent

`func (o *GetSlackBot200ResponseBot) SetHomeTabContent(v string)`

SetHomeTabContent sets HomeTabContent field to given value.

### HasHomeTabContent

`func (o *GetSlackBot200ResponseBot) HasHomeTabContent() bool`

HasHomeTabContent returns a boolean if a field has been set.

### GetAgentAccessControl

`func (o *GetSlackBot200ResponseBot) GetAgentAccessControl() map[string]interface{}`

GetAgentAccessControl returns the AgentAccessControl field if non-nil, zero value otherwise.

### GetAgentAccessControlOk

`func (o *GetSlackBot200ResponseBot) GetAgentAccessControlOk() (*map[string]interface{}, bool)`

GetAgentAccessControlOk returns a tuple with the AgentAccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAccessControl

`func (o *GetSlackBot200ResponseBot) SetAgentAccessControl(v map[string]interface{})`

SetAgentAccessControl sets AgentAccessControl field to given value.

### HasAgentAccessControl

`func (o *GetSlackBot200ResponseBot) HasAgentAccessControl() bool`

HasAgentAccessControl returns a boolean if a field has been set.

### GetKeywordsEnabled

`func (o *GetSlackBot200ResponseBot) GetKeywordsEnabled() bool`

GetKeywordsEnabled returns the KeywordsEnabled field if non-nil, zero value otherwise.

### GetKeywordsEnabledOk

`func (o *GetSlackBot200ResponseBot) GetKeywordsEnabledOk() (*bool, bool)`

GetKeywordsEnabledOk returns a tuple with the KeywordsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywordsEnabled

`func (o *GetSlackBot200ResponseBot) SetKeywordsEnabled(v bool)`

SetKeywordsEnabled sets KeywordsEnabled field to given value.

### HasKeywordsEnabled

`func (o *GetSlackBot200ResponseBot) HasKeywordsEnabled() bool`

HasKeywordsEnabled returns a boolean if a field has been set.

### GetKeywords

`func (o *GetSlackBot200ResponseBot) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *GetSlackBot200ResponseBot) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *GetSlackBot200ResponseBot) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *GetSlackBot200ResponseBot) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetSlackBot200ResponseBot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetSlackBot200ResponseBot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetSlackBot200ResponseBot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetSlackBot200ResponseBot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetSlackBot200ResponseBot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetSlackBot200ResponseBot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetSlackBot200ResponseBot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetSlackBot200ResponseBot) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


