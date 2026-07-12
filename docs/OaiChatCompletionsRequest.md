# OaiChatCompletionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | A model id from GET /oai/v1/models | 
**Messages** | [**[]OaiChatCompletionsRequestMessagesInner**](OaiChatCompletionsRequestMessagesInner.md) |  | 
**Stream** | Pointer to **bool** | Stream the response as SSE chat.completion.chunk events | [optional] [default to false]
**MaxTokens** | Pointer to **int32** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**TopP** | Pointer to **float32** |  | [optional] 
**Tools** | Pointer to **[]map[string]interface{}** | OpenAI function tool definitions | [optional] 
**ToolChoice** | Pointer to **interface{}** | auto | none | required | {type:function, function:{name}} | [optional] 
**StreamOptions** | Pointer to **map[string]interface{}** | {include_usage: true} to emit a final usage chunk when streaming | [optional] 

## Methods

### NewOaiChatCompletionsRequest

`func NewOaiChatCompletionsRequest(model string, messages []OaiChatCompletionsRequestMessagesInner, ) *OaiChatCompletionsRequest`

NewOaiChatCompletionsRequest instantiates a new OaiChatCompletionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOaiChatCompletionsRequestWithDefaults

`func NewOaiChatCompletionsRequestWithDefaults() *OaiChatCompletionsRequest`

NewOaiChatCompletionsRequestWithDefaults instantiates a new OaiChatCompletionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *OaiChatCompletionsRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OaiChatCompletionsRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OaiChatCompletionsRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMessages

`func (o *OaiChatCompletionsRequest) GetMessages() []OaiChatCompletionsRequestMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *OaiChatCompletionsRequest) GetMessagesOk() (*[]OaiChatCompletionsRequestMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *OaiChatCompletionsRequest) SetMessages(v []OaiChatCompletionsRequestMessagesInner)`

SetMessages sets Messages field to given value.


### GetStream

`func (o *OaiChatCompletionsRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *OaiChatCompletionsRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *OaiChatCompletionsRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *OaiChatCompletionsRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetMaxTokens

`func (o *OaiChatCompletionsRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *OaiChatCompletionsRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *OaiChatCompletionsRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *OaiChatCompletionsRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTemperature

`func (o *OaiChatCompletionsRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *OaiChatCompletionsRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *OaiChatCompletionsRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *OaiChatCompletionsRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *OaiChatCompletionsRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *OaiChatCompletionsRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *OaiChatCompletionsRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *OaiChatCompletionsRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetTools

`func (o *OaiChatCompletionsRequest) GetTools() []map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *OaiChatCompletionsRequest) GetToolsOk() (*[]map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *OaiChatCompletionsRequest) SetTools(v []map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *OaiChatCompletionsRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *OaiChatCompletionsRequest) GetToolChoice() interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *OaiChatCompletionsRequest) GetToolChoiceOk() (*interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *OaiChatCompletionsRequest) SetToolChoice(v interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *OaiChatCompletionsRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *OaiChatCompletionsRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *OaiChatCompletionsRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetStreamOptions

`func (o *OaiChatCompletionsRequest) GetStreamOptions() map[string]interface{}`

GetStreamOptions returns the StreamOptions field if non-nil, zero value otherwise.

### GetStreamOptionsOk

`func (o *OaiChatCompletionsRequest) GetStreamOptionsOk() (*map[string]interface{}, bool)`

GetStreamOptionsOk returns a tuple with the StreamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamOptions

`func (o *OaiChatCompletionsRequest) SetStreamOptions(v map[string]interface{})`

SetStreamOptions sets StreamOptions field to given value.

### HasStreamOptions

`func (o *OaiChatCompletionsRequest) HasStreamOptions() bool`

HasStreamOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


