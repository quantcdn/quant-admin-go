# OaiChatCompletions200ResponseChoicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to [**OaiChatCompletions200ResponseChoicesInnerMessage**](OaiChatCompletions200ResponseChoicesInnerMessage.md) |  | [optional] 
**FinishReason** | Pointer to **string** |  | [optional] 

## Methods

### NewOaiChatCompletions200ResponseChoicesInner

`func NewOaiChatCompletions200ResponseChoicesInner() *OaiChatCompletions200ResponseChoicesInner`

NewOaiChatCompletions200ResponseChoicesInner instantiates a new OaiChatCompletions200ResponseChoicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOaiChatCompletions200ResponseChoicesInnerWithDefaults

`func NewOaiChatCompletions200ResponseChoicesInnerWithDefaults() *OaiChatCompletions200ResponseChoicesInner`

NewOaiChatCompletions200ResponseChoicesInnerWithDefaults instantiates a new OaiChatCompletions200ResponseChoicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *OaiChatCompletions200ResponseChoicesInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OaiChatCompletions200ResponseChoicesInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OaiChatCompletions200ResponseChoicesInner) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *OaiChatCompletions200ResponseChoicesInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMessage

`func (o *OaiChatCompletions200ResponseChoicesInner) GetMessage() OaiChatCompletions200ResponseChoicesInnerMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OaiChatCompletions200ResponseChoicesInner) GetMessageOk() (*OaiChatCompletions200ResponseChoicesInnerMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OaiChatCompletions200ResponseChoicesInner) SetMessage(v OaiChatCompletions200ResponseChoicesInnerMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OaiChatCompletions200ResponseChoicesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFinishReason

`func (o *OaiChatCompletions200ResponseChoicesInner) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *OaiChatCompletions200ResponseChoicesInner) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *OaiChatCompletions200ResponseChoicesInner) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *OaiChatCompletions200ResponseChoicesInner) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


