# UpdateAISessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewMessages** | Pointer to [**[]UpdateAISessionRequestNewMessagesInner**](UpdateAISessionRequestNewMessagesInner.md) | New messages to append to conversation | [optional] 
**TokensUsed** | Pointer to **int32** | Tokens consumed in this turn | [optional] 
**Status** | Pointer to **string** | Update session status | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Update custom metadata | [optional] 

## Methods

### NewUpdateAISessionRequest

`func NewUpdateAISessionRequest() *UpdateAISessionRequest`

NewUpdateAISessionRequest instantiates a new UpdateAISessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAISessionRequestWithDefaults

`func NewUpdateAISessionRequestWithDefaults() *UpdateAISessionRequest`

NewUpdateAISessionRequestWithDefaults instantiates a new UpdateAISessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewMessages

`func (o *UpdateAISessionRequest) GetNewMessages() []UpdateAISessionRequestNewMessagesInner`

GetNewMessages returns the NewMessages field if non-nil, zero value otherwise.

### GetNewMessagesOk

`func (o *UpdateAISessionRequest) GetNewMessagesOk() (*[]UpdateAISessionRequestNewMessagesInner, bool)`

GetNewMessagesOk returns a tuple with the NewMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMessages

`func (o *UpdateAISessionRequest) SetNewMessages(v []UpdateAISessionRequestNewMessagesInner)`

SetNewMessages sets NewMessages field to given value.

### HasNewMessages

`func (o *UpdateAISessionRequest) HasNewMessages() bool`

HasNewMessages returns a boolean if a field has been set.

### GetTokensUsed

`func (o *UpdateAISessionRequest) GetTokensUsed() int32`

GetTokensUsed returns the TokensUsed field if non-nil, zero value otherwise.

### GetTokensUsedOk

`func (o *UpdateAISessionRequest) GetTokensUsedOk() (*int32, bool)`

GetTokensUsedOk returns a tuple with the TokensUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensUsed

`func (o *UpdateAISessionRequest) SetTokensUsed(v int32)`

SetTokensUsed sets TokensUsed field to given value.

### HasTokensUsed

`func (o *UpdateAISessionRequest) HasTokensUsed() bool`

HasTokensUsed returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAISessionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAISessionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAISessionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAISessionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateAISessionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateAISessionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateAISessionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateAISessionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


