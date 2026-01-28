# ChatWithAIAgentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The user&#39;s message to the agent | 
**SessionId** | Pointer to **string** | Optional session ID to continue a conversation | [optional] 
**UserId** | Pointer to **string** | Optional user identifier for session isolation | [optional] 
**Stream** | Pointer to **bool** | Whether to stream the response (SSE) | [optional] [default to false]
**Async** | Pointer to **bool** | Enable async/durable execution mode. When true, returns 202 with pollUrl. Use for long-running agent tasks. | [optional] [default to false]
**System** | Pointer to **string** | Optional additional system prompt (appended to agent&#39;s configured prompt) | [optional] 

## Methods

### NewChatWithAIAgentRequest

`func NewChatWithAIAgentRequest(message string, ) *ChatWithAIAgentRequest`

NewChatWithAIAgentRequest instantiates a new ChatWithAIAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithAIAgentRequestWithDefaults

`func NewChatWithAIAgentRequestWithDefaults() *ChatWithAIAgentRequest`

NewChatWithAIAgentRequestWithDefaults instantiates a new ChatWithAIAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ChatWithAIAgentRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatWithAIAgentRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatWithAIAgentRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSessionId

`func (o *ChatWithAIAgentRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatWithAIAgentRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatWithAIAgentRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatWithAIAgentRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetUserId

`func (o *ChatWithAIAgentRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChatWithAIAgentRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChatWithAIAgentRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ChatWithAIAgentRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStream

`func (o *ChatWithAIAgentRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatWithAIAgentRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatWithAIAgentRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatWithAIAgentRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetAsync

`func (o *ChatWithAIAgentRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *ChatWithAIAgentRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *ChatWithAIAgentRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *ChatWithAIAgentRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.

### GetSystem

`func (o *ChatWithAIAgentRequest) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ChatWithAIAgentRequest) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ChatWithAIAgentRequest) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ChatWithAIAgentRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


