# PostSlackBotMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**ThreadTs** | Pointer to **string** | Absent: a root post | [optional] 
**Text** | **string** | 1-8000 characters after trim | 
**Format** | Pointer to **string** | mrkdwn: Slack mrkdwn sections (default). markdown: standard Markdown in one Slack markdown block, including tables. | [optional] [default to "mrkdwn"]

## Methods

### NewPostSlackBotMessageRequest

`func NewPostSlackBotMessageRequest(channel string, text string, ) *PostSlackBotMessageRequest`

NewPostSlackBotMessageRequest instantiates a new PostSlackBotMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSlackBotMessageRequestWithDefaults

`func NewPostSlackBotMessageRequestWithDefaults() *PostSlackBotMessageRequest`

NewPostSlackBotMessageRequestWithDefaults instantiates a new PostSlackBotMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *PostSlackBotMessageRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostSlackBotMessageRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostSlackBotMessageRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetThreadTs

`func (o *PostSlackBotMessageRequest) GetThreadTs() string`

GetThreadTs returns the ThreadTs field if non-nil, zero value otherwise.

### GetThreadTsOk

`func (o *PostSlackBotMessageRequest) GetThreadTsOk() (*string, bool)`

GetThreadTsOk returns a tuple with the ThreadTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadTs

`func (o *PostSlackBotMessageRequest) SetThreadTs(v string)`

SetThreadTs sets ThreadTs field to given value.

### HasThreadTs

`func (o *PostSlackBotMessageRequest) HasThreadTs() bool`

HasThreadTs returns a boolean if a field has been set.

### GetText

`func (o *PostSlackBotMessageRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostSlackBotMessageRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostSlackBotMessageRequest) SetText(v string)`

SetText sets Text field to given value.


### GetFormat

`func (o *PostSlackBotMessageRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PostSlackBotMessageRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PostSlackBotMessageRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PostSlackBotMessageRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


