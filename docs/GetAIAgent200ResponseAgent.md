# GetAIAgent200ResponseAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**SystemPrompt** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**AllowedCollections** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetAIAgent200ResponseAgent

`func NewGetAIAgent200ResponseAgent() *GetAIAgent200ResponseAgent`

NewGetAIAgent200ResponseAgent instantiates a new GetAIAgent200ResponseAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAIAgent200ResponseAgentWithDefaults

`func NewGetAIAgent200ResponseAgentWithDefaults() *GetAIAgent200ResponseAgent`

NewGetAIAgent200ResponseAgentWithDefaults instantiates a new GetAIAgent200ResponseAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *GetAIAgent200ResponseAgent) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *GetAIAgent200ResponseAgent) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *GetAIAgent200ResponseAgent) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *GetAIAgent200ResponseAgent) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetName

`func (o *GetAIAgent200ResponseAgent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAIAgent200ResponseAgent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAIAgent200ResponseAgent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAIAgent200ResponseAgent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetAIAgent200ResponseAgent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAIAgent200ResponseAgent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAIAgent200ResponseAgent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAIAgent200ResponseAgent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *GetAIAgent200ResponseAgent) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetAIAgent200ResponseAgent) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetAIAgent200ResponseAgent) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetAIAgent200ResponseAgent) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *GetAIAgent200ResponseAgent) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *GetAIAgent200ResponseAgent) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *GetAIAgent200ResponseAgent) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *GetAIAgent200ResponseAgent) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetModelId

`func (o *GetAIAgent200ResponseAgent) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *GetAIAgent200ResponseAgent) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *GetAIAgent200ResponseAgent) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *GetAIAgent200ResponseAgent) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetTemperature

`func (o *GetAIAgent200ResponseAgent) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GetAIAgent200ResponseAgent) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GetAIAgent200ResponseAgent) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GetAIAgent200ResponseAgent) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *GetAIAgent200ResponseAgent) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *GetAIAgent200ResponseAgent) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *GetAIAgent200ResponseAgent) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *GetAIAgent200ResponseAgent) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetAllowedTools

`func (o *GetAIAgent200ResponseAgent) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *GetAIAgent200ResponseAgent) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *GetAIAgent200ResponseAgent) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *GetAIAgent200ResponseAgent) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *GetAIAgent200ResponseAgent) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *GetAIAgent200ResponseAgent) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *GetAIAgent200ResponseAgent) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *GetAIAgent200ResponseAgent) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetAIAgent200ResponseAgent) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetAIAgent200ResponseAgent) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetAIAgent200ResponseAgent) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetAIAgent200ResponseAgent) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetAIAgent200ResponseAgent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAIAgent200ResponseAgent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAIAgent200ResponseAgent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetAIAgent200ResponseAgent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetAIAgent200ResponseAgent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetAIAgent200ResponseAgent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetAIAgent200ResponseAgent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetAIAgent200ResponseAgent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


