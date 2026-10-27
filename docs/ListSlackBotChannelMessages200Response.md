# ListSlackBotChannelMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTrust** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to [**[]AiSlackBotApiMessage**](AiSlackBotApiMessage.md) |  | [optional] 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListSlackBotChannelMessages200Response

`func NewListSlackBotChannelMessages200Response() *ListSlackBotChannelMessages200Response`

NewListSlackBotChannelMessages200Response instantiates a new ListSlackBotChannelMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSlackBotChannelMessages200ResponseWithDefaults

`func NewListSlackBotChannelMessages200ResponseWithDefaults() *ListSlackBotChannelMessages200Response`

NewListSlackBotChannelMessages200ResponseWithDefaults instantiates a new ListSlackBotChannelMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTrust

`func (o *ListSlackBotChannelMessages200Response) GetContentTrust() string`

GetContentTrust returns the ContentTrust field if non-nil, zero value otherwise.

### GetContentTrustOk

`func (o *ListSlackBotChannelMessages200Response) GetContentTrustOk() (*string, bool)`

GetContentTrustOk returns a tuple with the ContentTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTrust

`func (o *ListSlackBotChannelMessages200Response) SetContentTrust(v string)`

SetContentTrust sets ContentTrust field to given value.

### HasContentTrust

`func (o *ListSlackBotChannelMessages200Response) HasContentTrust() bool`

HasContentTrust returns a boolean if a field has been set.

### GetChannel

`func (o *ListSlackBotChannelMessages200Response) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ListSlackBotChannelMessages200Response) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ListSlackBotChannelMessages200Response) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ListSlackBotChannelMessages200Response) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetMessages

`func (o *ListSlackBotChannelMessages200Response) GetMessages() []AiSlackBotApiMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListSlackBotChannelMessages200Response) GetMessagesOk() (*[]AiSlackBotApiMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListSlackBotChannelMessages200Response) SetMessages(v []AiSlackBotApiMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ListSlackBotChannelMessages200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNextCursor

`func (o *ListSlackBotChannelMessages200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListSlackBotChannelMessages200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListSlackBotChannelMessages200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListSlackBotChannelMessages200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ListSlackBotChannelMessages200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListSlackBotChannelMessages200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


