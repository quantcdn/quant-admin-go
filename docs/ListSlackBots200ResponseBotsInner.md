# ListSlackBots200ResponseBotsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotId** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**SetupType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
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

### GetAgentId

`func (o *ListSlackBots200ResponseBotsInner) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ListSlackBots200ResponseBotsInner) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ListSlackBots200ResponseBotsInner) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ListSlackBots200ResponseBotsInner) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

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


