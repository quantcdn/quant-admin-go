# CreateAIAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**Group** | Pointer to **string** |  | [optional] 
**SystemPrompt** | **string** |  | 
**Temperature** | Pointer to **float32** |  | [optional] 
**ModelId** | **string** |  | 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**AllowedTools** | Pointer to **[]string** |  | [optional] 
**AllowedCollections** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **string** | User identifier who created the agent | [optional] 

## Methods

### NewCreateAIAgentRequest

`func NewCreateAIAgentRequest(name string, description string, systemPrompt string, modelId string, ) *CreateAIAgentRequest`

NewCreateAIAgentRequest instantiates a new CreateAIAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAIAgentRequestWithDefaults

`func NewCreateAIAgentRequestWithDefaults() *CreateAIAgentRequest`

NewCreateAIAgentRequestWithDefaults instantiates a new CreateAIAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAIAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAIAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAIAgentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAIAgentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAIAgentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAIAgentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGroup

`func (o *CreateAIAgentRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CreateAIAgentRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CreateAIAgentRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CreateAIAgentRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *CreateAIAgentRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *CreateAIAgentRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *CreateAIAgentRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.


### GetTemperature

`func (o *CreateAIAgentRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *CreateAIAgentRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *CreateAIAgentRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *CreateAIAgentRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetModelId

`func (o *CreateAIAgentRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CreateAIAgentRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CreateAIAgentRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetMaxTokens

`func (o *CreateAIAgentRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *CreateAIAgentRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *CreateAIAgentRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *CreateAIAgentRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetAllowedTools

`func (o *CreateAIAgentRequest) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *CreateAIAgentRequest) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *CreateAIAgentRequest) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *CreateAIAgentRequest) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.

### GetAllowedCollections

`func (o *CreateAIAgentRequest) GetAllowedCollections() []string`

GetAllowedCollections returns the AllowedCollections field if non-nil, zero value otherwise.

### GetAllowedCollectionsOk

`func (o *CreateAIAgentRequest) GetAllowedCollectionsOk() (*[]string, bool)`

GetAllowedCollectionsOk returns a tuple with the AllowedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCollections

`func (o *CreateAIAgentRequest) SetAllowedCollections(v []string)`

SetAllowedCollections sets AllowedCollections field to given value.

### HasAllowedCollections

`func (o *CreateAIAgentRequest) HasAllowedCollections() bool`

HasAllowedCollections returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CreateAIAgentRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateAIAgentRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateAIAgentRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreateAIAgentRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


