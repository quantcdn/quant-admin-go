# ChatInferenceRequestToolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to [**[]ChatInferenceRequestToolConfigToolsInner**](ChatInferenceRequestToolConfigToolsInner.md) |  | [optional] 
**AutoExecute** | Pointer to **bool** | When true, backend automatically executes tools and feeds results back to AI. For async tools (e.g., image generation), returns executionId for polling. Security: Use allowedTools to whitelist which tools can auto-execute. | [optional] [default to false]
**AllowedTools** | Pointer to **[]string** | Whitelist of tool names that can be auto-executed. Required when autoExecute is true for security. Example: [&#39;get_weather&#39;, &#39;generate_image&#39;] | [optional] 

## Methods

### NewChatInferenceRequestToolConfig

`func NewChatInferenceRequestToolConfig() *ChatInferenceRequestToolConfig`

NewChatInferenceRequestToolConfig instantiates a new ChatInferenceRequestToolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInferenceRequestToolConfigWithDefaults

`func NewChatInferenceRequestToolConfigWithDefaults() *ChatInferenceRequestToolConfig`

NewChatInferenceRequestToolConfigWithDefaults instantiates a new ChatInferenceRequestToolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *ChatInferenceRequestToolConfig) GetTools() []ChatInferenceRequestToolConfigToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ChatInferenceRequestToolConfig) GetToolsOk() (*[]ChatInferenceRequestToolConfigToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ChatInferenceRequestToolConfig) SetTools(v []ChatInferenceRequestToolConfigToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ChatInferenceRequestToolConfig) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetAutoExecute

`func (o *ChatInferenceRequestToolConfig) GetAutoExecute() bool`

GetAutoExecute returns the AutoExecute field if non-nil, zero value otherwise.

### GetAutoExecuteOk

`func (o *ChatInferenceRequestToolConfig) GetAutoExecuteOk() (*bool, bool)`

GetAutoExecuteOk returns a tuple with the AutoExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExecute

`func (o *ChatInferenceRequestToolConfig) SetAutoExecute(v bool)`

SetAutoExecute sets AutoExecute field to given value.

### HasAutoExecute

`func (o *ChatInferenceRequestToolConfig) HasAutoExecute() bool`

HasAutoExecute returns a boolean if a field has been set.

### GetAllowedTools

`func (o *ChatInferenceRequestToolConfig) GetAllowedTools() []string`

GetAllowedTools returns the AllowedTools field if non-nil, zero value otherwise.

### GetAllowedToolsOk

`func (o *ChatInferenceRequestToolConfig) GetAllowedToolsOk() (*[]string, bool)`

GetAllowedToolsOk returns a tuple with the AllowedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTools

`func (o *ChatInferenceRequestToolConfig) SetAllowedTools(v []string)`

SetAllowedTools sets AllowedTools field to given value.

### HasAllowedTools

`func (o *ChatInferenceRequestToolConfig) HasAllowedTools() bool`

HasAllowedTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


