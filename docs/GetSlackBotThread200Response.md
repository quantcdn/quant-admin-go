# GetSlackBotThread200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTrust** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**ThreadTs** | Pointer to **string** |  | [optional] 
**Permalink** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to [**[]AiSlackBotApiMessage**](AiSlackBotApiMessage.md) |  | [optional] 
**CodeJob** | Pointer to [**NullableGetSlackBotThread200ResponseCodeJob**](GetSlackBotThread200ResponseCodeJob.md) |  | [optional] 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetSlackBotThread200Response

`func NewGetSlackBotThread200Response() *GetSlackBotThread200Response`

NewGetSlackBotThread200Response instantiates a new GetSlackBotThread200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSlackBotThread200ResponseWithDefaults

`func NewGetSlackBotThread200ResponseWithDefaults() *GetSlackBotThread200Response`

NewGetSlackBotThread200ResponseWithDefaults instantiates a new GetSlackBotThread200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTrust

`func (o *GetSlackBotThread200Response) GetContentTrust() string`

GetContentTrust returns the ContentTrust field if non-nil, zero value otherwise.

### GetContentTrustOk

`func (o *GetSlackBotThread200Response) GetContentTrustOk() (*string, bool)`

GetContentTrustOk returns a tuple with the ContentTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTrust

`func (o *GetSlackBotThread200Response) SetContentTrust(v string)`

SetContentTrust sets ContentTrust field to given value.

### HasContentTrust

`func (o *GetSlackBotThread200Response) HasContentTrust() bool`

HasContentTrust returns a boolean if a field has been set.

### GetChannel

`func (o *GetSlackBotThread200Response) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *GetSlackBotThread200Response) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *GetSlackBotThread200Response) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *GetSlackBotThread200Response) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetThreadTs

`func (o *GetSlackBotThread200Response) GetThreadTs() string`

GetThreadTs returns the ThreadTs field if non-nil, zero value otherwise.

### GetThreadTsOk

`func (o *GetSlackBotThread200Response) GetThreadTsOk() (*string, bool)`

GetThreadTsOk returns a tuple with the ThreadTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadTs

`func (o *GetSlackBotThread200Response) SetThreadTs(v string)`

SetThreadTs sets ThreadTs field to given value.

### HasThreadTs

`func (o *GetSlackBotThread200Response) HasThreadTs() bool`

HasThreadTs returns a boolean if a field has been set.

### GetPermalink

`func (o *GetSlackBotThread200Response) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *GetSlackBotThread200Response) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *GetSlackBotThread200Response) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *GetSlackBotThread200Response) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetMessages

`func (o *GetSlackBotThread200Response) GetMessages() []AiSlackBotApiMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetSlackBotThread200Response) GetMessagesOk() (*[]AiSlackBotApiMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetSlackBotThread200Response) SetMessages(v []AiSlackBotApiMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *GetSlackBotThread200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetCodeJob

`func (o *GetSlackBotThread200Response) GetCodeJob() GetSlackBotThread200ResponseCodeJob`

GetCodeJob returns the CodeJob field if non-nil, zero value otherwise.

### GetCodeJobOk

`func (o *GetSlackBotThread200Response) GetCodeJobOk() (*GetSlackBotThread200ResponseCodeJob, bool)`

GetCodeJobOk returns a tuple with the CodeJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeJob

`func (o *GetSlackBotThread200Response) SetCodeJob(v GetSlackBotThread200ResponseCodeJob)`

SetCodeJob sets CodeJob field to given value.

### HasCodeJob

`func (o *GetSlackBotThread200Response) HasCodeJob() bool`

HasCodeJob returns a boolean if a field has been set.

### SetCodeJobNil

`func (o *GetSlackBotThread200Response) SetCodeJobNil(b bool)`

 SetCodeJobNil sets the value for CodeJob to be an explicit nil

### UnsetCodeJob
`func (o *GetSlackBotThread200Response) UnsetCodeJob()`

UnsetCodeJob ensures that no value is present for CodeJob, not even an explicit nil
### GetNextCursor

`func (o *GetSlackBotThread200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetSlackBotThread200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetSlackBotThread200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *GetSlackBotThread200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *GetSlackBotThread200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetSlackBotThread200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


