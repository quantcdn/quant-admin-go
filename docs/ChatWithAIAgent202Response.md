# ChatWithAIAgent202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique request identifier for polling | 
**AgentId** | **string** | The agent processing the request | 
**AgentName** | Pointer to **string** | Human-readable agent name | [optional] 
**SessionId** | Pointer to **string** | Session ID (if provided) | [optional] 
**Status** | **string** | Initial status | 
**Message** | Pointer to **string** |  | [optional] 
**PollUrl** | **string** | URL to poll for execution status | 

## Methods

### NewChatWithAIAgent202Response

`func NewChatWithAIAgent202Response(requestId string, agentId string, status string, pollUrl string, ) *ChatWithAIAgent202Response`

NewChatWithAIAgent202Response instantiates a new ChatWithAIAgent202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithAIAgent202ResponseWithDefaults

`func NewChatWithAIAgent202ResponseWithDefaults() *ChatWithAIAgent202Response`

NewChatWithAIAgent202ResponseWithDefaults instantiates a new ChatWithAIAgent202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ChatWithAIAgent202Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ChatWithAIAgent202Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ChatWithAIAgent202Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetAgentId

`func (o *ChatWithAIAgent202Response) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ChatWithAIAgent202Response) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ChatWithAIAgent202Response) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetAgentName

`func (o *ChatWithAIAgent202Response) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *ChatWithAIAgent202Response) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *ChatWithAIAgent202Response) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *ChatWithAIAgent202Response) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetSessionId

`func (o *ChatWithAIAgent202Response) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatWithAIAgent202Response) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatWithAIAgent202Response) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatWithAIAgent202Response) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStatus

`func (o *ChatWithAIAgent202Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatWithAIAgent202Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatWithAIAgent202Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ChatWithAIAgent202Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatWithAIAgent202Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatWithAIAgent202Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatWithAIAgent202Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPollUrl

`func (o *ChatWithAIAgent202Response) GetPollUrl() string`

GetPollUrl returns the PollUrl field if non-nil, zero value otherwise.

### GetPollUrlOk

`func (o *ChatWithAIAgent202Response) GetPollUrlOk() (*string, bool)`

GetPollUrlOk returns a tuple with the PollUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollUrl

`func (o *ChatWithAIAgent202Response) SetPollUrl(v string)`

SetPollUrl sets PollUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


