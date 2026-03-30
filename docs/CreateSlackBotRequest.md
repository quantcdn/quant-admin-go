# CreateSlackBotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** | The AI agent that powers this bot | 
**SetupType** | **string** | Whether to use Quant-managed or customer-provided Slack app | 
**SessionTtlDays** | Pointer to **int32** | Session TTL in days | [optional] 
**AllowedChannels** | Pointer to **[]string** | Slack channel IDs the bot may respond in | [optional] 
**KeywordsEnabled** | Pointer to **bool** | Whether keyword triggers are enabled | [optional] 
**Keywords** | Pointer to **[]string** | Keywords that trigger the bot | [optional] 
**SlashCommands** | Pointer to **[]string** | Slash commands the bot responds to | [optional] 

## Methods

### NewCreateSlackBotRequest

`func NewCreateSlackBotRequest(agentId string, setupType string, ) *CreateSlackBotRequest`

NewCreateSlackBotRequest instantiates a new CreateSlackBotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSlackBotRequestWithDefaults

`func NewCreateSlackBotRequestWithDefaults() *CreateSlackBotRequest`

NewCreateSlackBotRequestWithDefaults instantiates a new CreateSlackBotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *CreateSlackBotRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CreateSlackBotRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CreateSlackBotRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


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

### GetSlashCommands

`func (o *CreateSlackBotRequest) GetSlashCommands() []string`

GetSlashCommands returns the SlashCommands field if non-nil, zero value otherwise.

### GetSlashCommandsOk

`func (o *CreateSlackBotRequest) GetSlashCommandsOk() (*[]string, bool)`

GetSlashCommandsOk returns a tuple with the SlashCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashCommands

`func (o *CreateSlackBotRequest) SetSlashCommands(v []string)`

SetSlashCommands sets SlashCommands field to given value.

### HasSlashCommands

`func (o *CreateSlackBotRequest) HasSlashCommands() bool`

HasSlashCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


