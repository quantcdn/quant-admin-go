# GetSlackBot200ResponseBot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotId** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**SetupType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**SessionTtlDays** | Pointer to **int32** |  | [optional] 
**AllowedChannels** | Pointer to **[]string** |  | [optional] 
**KeywordsEnabled** | Pointer to **bool** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**SlashCommands** | Pointer to **[]string** |  | [optional] 
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

### GetAgentId

`func (o *GetSlackBot200ResponseBot) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *GetSlackBot200ResponseBot) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *GetSlackBot200ResponseBot) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *GetSlackBot200ResponseBot) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

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

### GetSlashCommands

`func (o *GetSlackBot200ResponseBot) GetSlashCommands() []string`

GetSlashCommands returns the SlashCommands field if non-nil, zero value otherwise.

### GetSlashCommandsOk

`func (o *GetSlackBot200ResponseBot) GetSlashCommandsOk() (*[]string, bool)`

GetSlashCommandsOk returns a tuple with the SlashCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlashCommands

`func (o *GetSlackBot200ResponseBot) SetSlashCommands(v []string)`

SetSlashCommands sets SlashCommands field to given value.

### HasSlashCommands

`func (o *GetSlackBot200ResponseBot) HasSlashCommands() bool`

HasSlashCommands returns a boolean if a field has been set.

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


