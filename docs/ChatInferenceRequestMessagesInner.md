# ChatInferenceRequestMessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**Content** | [**ChatInferenceRequestMessagesInnerContent**](ChatInferenceRequestMessagesInnerContent.md) |  | 

## Methods

### NewChatInferenceRequestMessagesInner

`func NewChatInferenceRequestMessagesInner(role string, content ChatInferenceRequestMessagesInnerContent, ) *ChatInferenceRequestMessagesInner`

NewChatInferenceRequestMessagesInner instantiates a new ChatInferenceRequestMessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInferenceRequestMessagesInnerWithDefaults

`func NewChatInferenceRequestMessagesInnerWithDefaults() *ChatInferenceRequestMessagesInner`

NewChatInferenceRequestMessagesInnerWithDefaults instantiates a new ChatInferenceRequestMessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ChatInferenceRequestMessagesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatInferenceRequestMessagesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatInferenceRequestMessagesInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *ChatInferenceRequestMessagesInner) GetContent() ChatInferenceRequestMessagesInnerContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatInferenceRequestMessagesInner) GetContentOk() (*ChatInferenceRequestMessagesInnerContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatInferenceRequestMessagesInner) SetContent(v ChatInferenceRequestMessagesInnerContent)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


