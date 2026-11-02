# CreateSlackInteractionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**ThreadTs** | Pointer to **string** |  | [optional] 
**Text** | **string** |  | 
**Actions** | [**[]CreateSlackInteractionRequestActionsInner**](CreateSlackInteractionRequestActionsInner.md) |  | 
**ResponseMode** | Pointer to **string** |  | [optional] 
**ExpiresInSeconds** | Pointer to **int32** |  | [optional] [default to 3600]
**AllowedUsers** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Opaque caller data, at most 4096 bytes of canonical JSON | [optional] 

## Methods

### NewCreateSlackInteractionRequest

`func NewCreateSlackInteractionRequest(channel string, text string, actions []CreateSlackInteractionRequestActionsInner, ) *CreateSlackInteractionRequest`

NewCreateSlackInteractionRequest instantiates a new CreateSlackInteractionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSlackInteractionRequestWithDefaults

`func NewCreateSlackInteractionRequestWithDefaults() *CreateSlackInteractionRequest`

NewCreateSlackInteractionRequestWithDefaults instantiates a new CreateSlackInteractionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CreateSlackInteractionRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateSlackInteractionRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateSlackInteractionRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetThreadTs

`func (o *CreateSlackInteractionRequest) GetThreadTs() string`

GetThreadTs returns the ThreadTs field if non-nil, zero value otherwise.

### GetThreadTsOk

`func (o *CreateSlackInteractionRequest) GetThreadTsOk() (*string, bool)`

GetThreadTsOk returns a tuple with the ThreadTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadTs

`func (o *CreateSlackInteractionRequest) SetThreadTs(v string)`

SetThreadTs sets ThreadTs field to given value.

### HasThreadTs

`func (o *CreateSlackInteractionRequest) HasThreadTs() bool`

HasThreadTs returns a boolean if a field has been set.

### GetText

`func (o *CreateSlackInteractionRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateSlackInteractionRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateSlackInteractionRequest) SetText(v string)`

SetText sets Text field to given value.


### GetActions

`func (o *CreateSlackInteractionRequest) GetActions() []CreateSlackInteractionRequestActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateSlackInteractionRequest) GetActionsOk() (*[]CreateSlackInteractionRequestActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateSlackInteractionRequest) SetActions(v []CreateSlackInteractionRequestActionsInner)`

SetActions sets Actions field to given value.


### GetResponseMode

`func (o *CreateSlackInteractionRequest) GetResponseMode() string`

GetResponseMode returns the ResponseMode field if non-nil, zero value otherwise.

### GetResponseModeOk

`func (o *CreateSlackInteractionRequest) GetResponseModeOk() (*string, bool)`

GetResponseModeOk returns a tuple with the ResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMode

`func (o *CreateSlackInteractionRequest) SetResponseMode(v string)`

SetResponseMode sets ResponseMode field to given value.

### HasResponseMode

`func (o *CreateSlackInteractionRequest) HasResponseMode() bool`

HasResponseMode returns a boolean if a field has been set.

### GetExpiresInSeconds

`func (o *CreateSlackInteractionRequest) GetExpiresInSeconds() int32`

GetExpiresInSeconds returns the ExpiresInSeconds field if non-nil, zero value otherwise.

### GetExpiresInSecondsOk

`func (o *CreateSlackInteractionRequest) GetExpiresInSecondsOk() (*int32, bool)`

GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSeconds

`func (o *CreateSlackInteractionRequest) SetExpiresInSeconds(v int32)`

SetExpiresInSeconds sets ExpiresInSeconds field to given value.

### HasExpiresInSeconds

`func (o *CreateSlackInteractionRequest) HasExpiresInSeconds() bool`

HasExpiresInSeconds returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *CreateSlackInteractionRequest) GetAllowedUsers() []string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *CreateSlackInteractionRequest) GetAllowedUsersOk() (*[]string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *CreateSlackInteractionRequest) SetAllowedUsers(v []string)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *CreateSlackInteractionRequest) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSlackInteractionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSlackInteractionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSlackInteractionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSlackInteractionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


