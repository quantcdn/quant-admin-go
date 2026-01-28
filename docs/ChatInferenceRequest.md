# ChatInferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ChatInferenceRequestMessagesInner**](ChatInferenceRequestMessagesInner.md) | Array of chat messages. Content can be a simple string or an array of content blocks for multimodal input. | 
**ModelId** | **string** | Model ID. Use Nova models for multimodal support. | 
**Temperature** | Pointer to **float32** |  | [optional] [default to 0.7]
**MaxTokens** | Pointer to **int32** | Max tokens. Claude 4.5 supports up to 64k. | [optional] [default to 4096]
**TopP** | Pointer to **float32** |  | [optional] 
**Stream** | Pointer to **bool** | Ignored in buffered mode, always returns complete response | [optional] 
**SystemPrompt** | Pointer to **string** | Optional custom system prompt. When tools are enabled, this is prepended with tool usage guidance. | [optional] 
**StopSequences** | Pointer to **[]string** | Custom stop sequences | [optional] 
**ResponseFormat** | Pointer to [**ChatInferenceRequestResponseFormat**](ChatInferenceRequestResponseFormat.md) |  | [optional] 
**ToolConfig** | Pointer to [**ChatInferenceRequestToolConfig**](ChatInferenceRequestToolConfig.md) |  | [optional] 
**SessionId** | Pointer to **string** | Optional session ID for conversation continuity. Omit to use stateless mode, include to continue an existing session. | [optional] 
**Async** | Pointer to **bool** | Enable async/durable execution mode. When true, returns 202 with pollUrl instead of waiting for completion. Use for long-running inference, client-executed tools, or operations &gt;30 seconds. | [optional] [default to false]

## Methods

### NewChatInferenceRequest

`func NewChatInferenceRequest(messages []ChatInferenceRequestMessagesInner, modelId string, ) *ChatInferenceRequest`

NewChatInferenceRequest instantiates a new ChatInferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInferenceRequestWithDefaults

`func NewChatInferenceRequestWithDefaults() *ChatInferenceRequest`

NewChatInferenceRequestWithDefaults instantiates a new ChatInferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatInferenceRequest) GetMessages() []ChatInferenceRequestMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatInferenceRequest) GetMessagesOk() (*[]ChatInferenceRequestMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatInferenceRequest) SetMessages(v []ChatInferenceRequestMessagesInner)`

SetMessages sets Messages field to given value.


### GetModelId

`func (o *ChatInferenceRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ChatInferenceRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ChatInferenceRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetTemperature

`func (o *ChatInferenceRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ChatInferenceRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ChatInferenceRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ChatInferenceRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *ChatInferenceRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ChatInferenceRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ChatInferenceRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ChatInferenceRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTopP

`func (o *ChatInferenceRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ChatInferenceRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ChatInferenceRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ChatInferenceRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetStream

`func (o *ChatInferenceRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatInferenceRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatInferenceRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatInferenceRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *ChatInferenceRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *ChatInferenceRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *ChatInferenceRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *ChatInferenceRequest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetStopSequences

`func (o *ChatInferenceRequest) GetStopSequences() []string`

GetStopSequences returns the StopSequences field if non-nil, zero value otherwise.

### GetStopSequencesOk

`func (o *ChatInferenceRequest) GetStopSequencesOk() (*[]string, bool)`

GetStopSequencesOk returns a tuple with the StopSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSequences

`func (o *ChatInferenceRequest) SetStopSequences(v []string)`

SetStopSequences sets StopSequences field to given value.

### HasStopSequences

`func (o *ChatInferenceRequest) HasStopSequences() bool`

HasStopSequences returns a boolean if a field has been set.

### GetResponseFormat

`func (o *ChatInferenceRequest) GetResponseFormat() ChatInferenceRequestResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ChatInferenceRequest) GetResponseFormatOk() (*ChatInferenceRequestResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ChatInferenceRequest) SetResponseFormat(v ChatInferenceRequestResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ChatInferenceRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetToolConfig

`func (o *ChatInferenceRequest) GetToolConfig() ChatInferenceRequestToolConfig`

GetToolConfig returns the ToolConfig field if non-nil, zero value otherwise.

### GetToolConfigOk

`func (o *ChatInferenceRequest) GetToolConfigOk() (*ChatInferenceRequestToolConfig, bool)`

GetToolConfigOk returns a tuple with the ToolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfig

`func (o *ChatInferenceRequest) SetToolConfig(v ChatInferenceRequestToolConfig)`

SetToolConfig sets ToolConfig field to given value.

### HasToolConfig

`func (o *ChatInferenceRequest) HasToolConfig() bool`

HasToolConfig returns a boolean if a field has been set.

### GetSessionId

`func (o *ChatInferenceRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatInferenceRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatInferenceRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatInferenceRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetAsync

`func (o *ChatInferenceRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *ChatInferenceRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *ChatInferenceRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *ChatInferenceRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


