# AiSlackBotApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Error** | **bool** | Thin proxy for the Slack bot API: list API-enabled bots, read channel messages and threads, and post as the bot. The portal validates the request, requires a verified email, calls the platform and passes its status and body through. The platform owns the bot, channel and rate-limit rules and their error codes.  Only the values that SlackBotApiRequest returns go to the platform, never raw input. The activity entry (ai_slack_api) never holds message text. | 
**Code** | **string** |  | 
**Retryable** | **bool** |  | 
**SlackError** | Pointer to **string** | The Slack error string. Only with code SLACK_ERROR. | [optional] 

## Methods

### NewAiSlackBotApiError

`func NewAiSlackBotApiError(message string, error_ bool, code string, retryable bool, ) *AiSlackBotApiError`

NewAiSlackBotApiError instantiates a new AiSlackBotApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSlackBotApiErrorWithDefaults

`func NewAiSlackBotApiErrorWithDefaults() *AiSlackBotApiError`

NewAiSlackBotApiErrorWithDefaults instantiates a new AiSlackBotApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiSlackBotApiError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiSlackBotApiError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiSlackBotApiError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *AiSlackBotApiError) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiSlackBotApiError) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiSlackBotApiError) SetError(v bool)`

SetError sets Error field to given value.


### GetCode

`func (o *AiSlackBotApiError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiSlackBotApiError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiSlackBotApiError) SetCode(v string)`

SetCode sets Code field to given value.


### GetRetryable

`func (o *AiSlackBotApiError) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *AiSlackBotApiError) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *AiSlackBotApiError) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.


### GetSlackError

`func (o *AiSlackBotApiError) GetSlackError() string`

GetSlackError returns the SlackError field if non-nil, zero value otherwise.

### GetSlackErrorOk

`func (o *AiSlackBotApiError) GetSlackErrorOk() (*string, bool)`

GetSlackErrorOk returns a tuple with the SlackError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackError

`func (o *AiSlackBotApiError) SetSlackError(v string)`

SetSlackError sets SlackError field to given value.

### HasSlackError

`func (o *AiSlackBotApiError) HasSlackError() bool`

HasSlackError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


