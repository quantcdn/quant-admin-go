# UpdateSlackBotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** | Change the backing AI agent | [optional] 
**Status** | Pointer to **string** | Enable or disable the bot | [optional] 
**SessionTtlDays** | Pointer to **int32** | Session TTL in days | [optional] 
**AllowedChannels** | Pointer to **[]string** | Slack channel IDs the bot may respond in | [optional] 
**KeywordsEnabled** | Pointer to **bool** | Whether keyword triggers are enabled | [optional] 
**Keywords** | Pointer to **[]string** | Keywords that trigger the bot | [optional] 
**SlashCommands** | Pointer to **[]string** | Slash commands the bot responds to | [optional] 

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

### GetAgentId

`func (o *UpdateSlackBotRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *UpdateSlackBotRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *UpdateSlackBotRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *UpdateSlackBotRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

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

### GetSlashCommands

`func (o *UpdateSlackBotRequest) GetSlashCommands() []string`

GetSlashCommands returns the SlashCommands field if non-nil, zero value otherwise.

### GetSlashCommandsOk

`func (o *UpdateSlackBotRequest) GetSlashCommandsOk() (*[]string, bool)`

GetSlashCommandsOk returns a tuple with the SlashCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashCommands

`func (o *UpdateSlackBotRequest) SetSlashCommands(v []string)`

SetSlashCommands sets SlashCommands field to given value.

### HasSlashCommands

`func (o *UpdateSlackBotRequest) HasSlashCommands() bool`

HasSlashCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


