# OaiChatCompletions200ResponseChoicesInnerMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**ToolCalls** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewOaiChatCompletions200ResponseChoicesInnerMessage

`func NewOaiChatCompletions200ResponseChoicesInnerMessage() *OaiChatCompletions200ResponseChoicesInnerMessage`

NewOaiChatCompletions200ResponseChoicesInnerMessage instantiates a new OaiChatCompletions200ResponseChoicesInnerMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOaiChatCompletions200ResponseChoicesInnerMessageWithDefaults

`func NewOaiChatCompletions200ResponseChoicesInnerMessageWithDefaults() *OaiChatCompletions200ResponseChoicesInnerMessage`

NewOaiChatCompletions200ResponseChoicesInnerMessageWithDefaults instantiates a new OaiChatCompletions200ResponseChoicesInnerMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetToolCalls

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) GetToolCalls() []map[string]interface{}`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) GetToolCallsOk() (*[]map[string]interface{}, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) SetToolCalls(v []map[string]interface{})`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *OaiChatCompletions200ResponseChoicesInnerMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


