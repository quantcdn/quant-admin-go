# ChatInference200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** | Text response content | [optional] 
**ToolUse** | Pointer to [**ChatInference200ResponseResponseToolUse**](ChatInference200ResponseResponseToolUse.md) |  | [optional] 

## Methods

### NewChatInference200ResponseResponse

`func NewChatInference200ResponseResponse() *ChatInference200ResponseResponse`

NewChatInference200ResponseResponse instantiates a new ChatInference200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInference200ResponseResponseWithDefaults

`func NewChatInference200ResponseResponseWithDefaults() *ChatInference200ResponseResponse`

NewChatInference200ResponseResponseWithDefaults instantiates a new ChatInference200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ChatInference200ResponseResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatInference200ResponseResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatInference200ResponseResponse) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ChatInference200ResponseResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *ChatInference200ResponseResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatInference200ResponseResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatInference200ResponseResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ChatInference200ResponseResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetToolUse

`func (o *ChatInference200ResponseResponse) GetToolUse() ChatInference200ResponseResponseToolUse`

GetToolUse returns the ToolUse field if non-nil, zero value otherwise.

### GetToolUseOk

`func (o *ChatInference200ResponseResponse) GetToolUseOk() (*ChatInference200ResponseResponseToolUse, bool)`

GetToolUseOk returns a tuple with the ToolUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUse

`func (o *ChatInference200ResponseResponse) SetToolUse(v ChatInference200ResponseResponseToolUse)`

SetToolUse sets ToolUse field to given value.

### HasToolUse

`func (o *ChatInference200ResponseResponse) HasToolUse() bool`

HasToolUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


