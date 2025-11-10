# CreateAISessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User identifier for this session | 
**SessionGroup** | Pointer to **string** | Optional user-defined grouping identifier (e.g., app name, environment, tenant). Use any format that makes sense for your application. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional custom metadata for additional context | [optional] 
**ExpirationMinutes** | Pointer to **int32** | Session expiration in minutes | [optional] [default to 60]
**InitialMessages** | Pointer to [**[]CreateAISessionRequestInitialMessagesInner**](CreateAISessionRequestInitialMessagesInner.md) | Initial conversation messages (e.g., system prompt) | [optional] 

## Methods

### NewCreateAISessionRequest

`func NewCreateAISessionRequest(userId string, ) *CreateAISessionRequest`

NewCreateAISessionRequest instantiates a new CreateAISessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAISessionRequestWithDefaults

`func NewCreateAISessionRequestWithDefaults() *CreateAISessionRequest`

NewCreateAISessionRequestWithDefaults instantiates a new CreateAISessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateAISessionRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateAISessionRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateAISessionRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSessionGroup

`func (o *CreateAISessionRequest) GetSessionGroup() string`

GetSessionGroup returns the SessionGroup field if non-nil, zero value otherwise.

### GetSessionGroupOk

`func (o *CreateAISessionRequest) GetSessionGroupOk() (*string, bool)`

GetSessionGroupOk returns a tuple with the SessionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionGroup

`func (o *CreateAISessionRequest) SetSessionGroup(v string)`

SetSessionGroup sets SessionGroup field to given value.

### HasSessionGroup

`func (o *CreateAISessionRequest) HasSessionGroup() bool`

HasSessionGroup returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateAISessionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateAISessionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateAISessionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateAISessionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpirationMinutes

`func (o *CreateAISessionRequest) GetExpirationMinutes() int32`

GetExpirationMinutes returns the ExpirationMinutes field if non-nil, zero value otherwise.

### GetExpirationMinutesOk

`func (o *CreateAISessionRequest) GetExpirationMinutesOk() (*int32, bool)`

GetExpirationMinutesOk returns a tuple with the ExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMinutes

`func (o *CreateAISessionRequest) SetExpirationMinutes(v int32)`

SetExpirationMinutes sets ExpirationMinutes field to given value.

### HasExpirationMinutes

`func (o *CreateAISessionRequest) HasExpirationMinutes() bool`

HasExpirationMinutes returns a boolean if a field has been set.

### GetInitialMessages

`func (o *CreateAISessionRequest) GetInitialMessages() []CreateAISessionRequestInitialMessagesInner`

GetInitialMessages returns the InitialMessages field if non-nil, zero value otherwise.

### GetInitialMessagesOk

`func (o *CreateAISessionRequest) GetInitialMessagesOk() (*[]CreateAISessionRequestInitialMessagesInner, bool)`

GetInitialMessagesOk returns a tuple with the InitialMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMessages

`func (o *CreateAISessionRequest) SetInitialMessages(v []CreateAISessionRequestInitialMessagesInner)`

SetInitialMessages sets InitialMessages field to given value.

### HasInitialMessages

`func (o *CreateAISessionRequest) HasInitialMessages() bool`

HasInitialMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


