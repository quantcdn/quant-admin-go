# ChatInference202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique request identifier for polling | 
**SessionId** | Pointer to **string** | Session ID for conversation continuity | [optional] 
**Status** | **string** | Initial execution status | 
**Message** | Pointer to **string** | Human-readable status message | [optional] 
**PollUrl** | **string** | URL to poll for execution status | 

## Methods

### NewChatInference202Response

`func NewChatInference202Response(requestId string, status string, pollUrl string, ) *ChatInference202Response`

NewChatInference202Response instantiates a new ChatInference202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInference202ResponseWithDefaults

`func NewChatInference202ResponseWithDefaults() *ChatInference202Response`

NewChatInference202ResponseWithDefaults instantiates a new ChatInference202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ChatInference202Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ChatInference202Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ChatInference202Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetSessionId

`func (o *ChatInference202Response) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ChatInference202Response) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ChatInference202Response) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ChatInference202Response) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStatus

`func (o *ChatInference202Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatInference202Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatInference202Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ChatInference202Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatInference202Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatInference202Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatInference202Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPollUrl

`func (o *ChatInference202Response) GetPollUrl() string`

GetPollUrl returns the PollUrl field if non-nil, zero value otherwise.

### GetPollUrlOk

`func (o *ChatInference202Response) GetPollUrlOk() (*string, bool)`

GetPollUrlOk returns a tuple with the PollUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollUrl

`func (o *ChatInference202Response) SetPollUrl(v string)`

SetPollUrl sets PollUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


