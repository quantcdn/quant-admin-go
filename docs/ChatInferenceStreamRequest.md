# ChatInferenceStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ChatInferenceStreamRequestMessagesInner**](ChatInferenceStreamRequestMessagesInner.md) | Array of chat messages. Content can be a simple string or an array of content blocks for multimodal input. | 
**ModelId** | **string** | Model ID. Use Nova models for multimodal support. | 
**Temperature** | Pointer to **float32** |  | [optional] [default to 0.7]
**MaxTokens** | Pointer to **int32** | Max tokens. Claude 4.5 supports up to 64k. | [optional] [default to 4096]
**TopP** | Pointer to **float32** |  | [optional] 
**SystemPrompt** | Pointer to **string** | Optional custom system prompt. When tools are enabled, this is prepended with tool usage guidance. | [optional] 
**StopSequences** | Pointer to **[]string** | Custom stop sequences | [optional] 
**ResponseFormat** | Pointer to [**ChatInferenceRequestResponseFormat**](ChatInferenceRequestResponseFormat.md) |  | [optional] 
**ToolConfig** | Pointer to [**ChatInferenceRequestToolConfig**](ChatInferenceRequestToolConfig.md) |  | [optional] 
**SessionId** | Pointer to **string** | Optional session ID for conversation continuity. Omit to use stateless mode, include to continue an existing session. | [optional] 
**Async** | Pointer to **bool** | Enable async/durable execution mode. When true, returns 202 with pollUrl instead of streaming. Use for long-running inference, client-executed tools, or operations &gt;30 seconds. | [optional] [default to false]

## Methods

### NewChatInferenceStreamRequest

`func NewChatInferenceStreamRequest(messages []ChatInferenceStreamRequestMessagesInner, modelId string, ) *ChatInferenceStreamRequest`

NewChatInferenceStreamRequest instantiates a new ChatInferenceStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInferenceStreamRequestWithDefaults

`func NewChatInferenceStreamRequestWithDefaults() *ChatInferenceStreamRequest`

NewChatInferenceStreamRequestWithDefaults instantiates a new ChatInferenceStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatInferenceStreamRequest) GetMessages() []ChatInferenceStreamRequestMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatInferenceStreamRequest) GetMessagesOk() (*[]ChatInferenceStreamRequestMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatInferenceStreamRequest) SetMessages(v []ChatInferenceStreamRequestMessagesInner)`

SetMessages sets Messages field to given value.


### GetModelId

`func (o *ChatInferenceStreamRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ChatInferenceStreamRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ChatInferenceStreamRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetTemperature

`func (o *ChatInferenceStreamRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ChatInferenceStreamRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ChatInferenceStreamRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ChatInferenceStreamRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *ChatInferenceStreamRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ChatInferenceStreamRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ChatInferenceStreamRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ChatInferenceStreamRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTopP

`func (o *ChatInferenceStreamRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ChatInferenceStreamRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ChatInferenceStreamRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ChatInferenceStreamRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *ChatInferenceStreamRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *ChatInferenceStreamRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *ChatInferenceStreamRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *ChatInferenceStreamRequest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetStopSequences

`func (o *ChatInferenceStreamRequest) GetStopSequences() []string`

GetStopSequences returns the StopSequences field if non-nil, zero value otherwise.

### GetStopSequencesOk

`func (o *ChatInferenceStreamRequest) GetStopSequencesOk() (*[]string, bool)`

GetStopSequencesOk returns a tuple with the StopSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSequences

`func (o *ChatInferenceStreamRequest) SetStopSequences(v []string)`

SetStopSequences sets StopSequences field to given value.

### HasStopSequences

`func (o *ChatInferenceStreamRequest) HasStopSequences() bool`

HasStopSequences returns a boolean if a field has been set.

### GetResponseFormat

`func (o *ChatInferenceStreamRequest) GetResponseFormat() ChatInferenceRequestResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ChatInferenceStreamRequest) GetResponseFormatOk() (*ChatInferenceRequestResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ChatInferenceStreamRequest) SetResponseFormat(v ChatInferenceRequestResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ChatInferenceStreamRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetToolConfig

`func (o *ChatInferenceStreamRequest) GetToolConfig() ChatInferenceRequestToolConfig`

GetToolConfig returns the ToolConfig field if non-nil, zero value otherwise.

### GetToolConfigOk

`func (o *ChatInferenceStreamRequest) GetToolConfigOk() (*ChatInferenceRequestToolConfig, bool)`

GetToolConfigOk returns a tuple with the ToolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfig

`func (o *ChatInferenceStreamRequest) SetToolConfig(v ChatInferenceRequestToolConfig)`

SetToolConfig sets ToolConfig field to given value.

### HasToolConfig

`func (o *ChatInferenceStreamRequest) HasToolConfig() bool`

HasToolConfig returns a boolean if a field has been set.

### GetSessionId

`func (o *ChatInferenceStreamRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatInferenceStreamRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatInferenceStreamRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatInferenceStreamRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetAsync

`func (o *ChatInferenceStreamRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *ChatInferenceStreamRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *ChatInferenceStreamRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *ChatInferenceStreamRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


