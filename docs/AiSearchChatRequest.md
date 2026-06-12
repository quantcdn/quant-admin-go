# AiSearchChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**ContextUrl** | Pointer to **string** |  | [optional] 
**MaxContextChunks** | Pointer to **int32** |  | [optional] 

## Methods

### NewAiSearchChatRequest

`func NewAiSearchChatRequest(message string, ) *AiSearchChatRequest`

NewAiSearchChatRequest instantiates a new AiSearchChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSearchChatRequestWithDefaults

`func NewAiSearchChatRequestWithDefaults() *AiSearchChatRequest`

NewAiSearchChatRequestWithDefaults instantiates a new AiSearchChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiSearchChatRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiSearchChatRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiSearchChatRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSessionId

`func (o *AiSearchChatRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AiSearchChatRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AiSearchChatRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AiSearchChatRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetContextUrl

`func (o *AiSearchChatRequest) GetContextUrl() string`

GetContextUrl returns the ContextUrl field if non-nil, zero value otherwise.

### GetContextUrlOk

`func (o *AiSearchChatRequest) GetContextUrlOk() (*string, bool)`

GetContextUrlOk returns a tuple with the ContextUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextUrl

`func (o *AiSearchChatRequest) SetContextUrl(v string)`

SetContextUrl sets ContextUrl field to given value.

### HasContextUrl

`func (o *AiSearchChatRequest) HasContextUrl() bool`

HasContextUrl returns a boolean if a field has been set.

### GetMaxContextChunks

`func (o *AiSearchChatRequest) GetMaxContextChunks() int32`

GetMaxContextChunks returns the MaxContextChunks field if non-nil, zero value otherwise.

### GetMaxContextChunksOk

`func (o *AiSearchChatRequest) GetMaxContextChunksOk() (*int32, bool)`

GetMaxContextChunksOk returns a tuple with the MaxContextChunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContextChunks

`func (o *AiSearchChatRequest) SetMaxContextChunks(v int32)`

SetMaxContextChunks sets MaxContextChunks field to given value.

### HasMaxContextChunks

`func (o *AiSearchChatRequest) HasMaxContextChunks() bool`

HasMaxContextChunks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


