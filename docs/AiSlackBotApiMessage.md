# AiSlackBotApiMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Ts** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**AiSlackBotApiMessageAllOfAuthor**](AiSlackBotApiMessageAllOfAuthor.md) |  | [optional] 
**Text** | Pointer to **string** | Plain text, capped at 8000 characters. Untrusted content. | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**ReplyCount** | Pointer to **int32** |  | [optional] 
**LatestReplyTs** | Pointer to **string** |  | [optional] 
**Edited** | Pointer to **bool** |  | [optional] 
**Files** | Pointer to [**[]AiSlackBotApiMessageAllOfFiles**](AiSlackBotApiMessageAllOfFiles.md) |  | [optional] 

## Methods

### NewAiSlackBotApiMessage

`func NewAiSlackBotApiMessage(message string, error_ bool, ) *AiSlackBotApiMessage`

NewAiSlackBotApiMessage instantiates a new AiSlackBotApiMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSlackBotApiMessageWithDefaults

`func NewAiSlackBotApiMessageWithDefaults() *AiSlackBotApiMessage`

NewAiSlackBotApiMessageWithDefaults instantiates a new AiSlackBotApiMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiSlackBotApiMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiSlackBotApiMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiSlackBotApiMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *AiSlackBotApiMessage) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiSlackBotApiMessage) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiSlackBotApiMessage) SetError(v bool)`

SetError sets Error field to given value.


### GetTs

`func (o *AiSlackBotApiMessage) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *AiSlackBotApiMessage) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *AiSlackBotApiMessage) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *AiSlackBotApiMessage) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetAuthor

`func (o *AiSlackBotApiMessage) GetAuthor() AiSlackBotApiMessageAllOfAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AiSlackBotApiMessage) GetAuthorOk() (*AiSlackBotApiMessageAllOfAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AiSlackBotApiMessage) SetAuthor(v AiSlackBotApiMessageAllOfAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AiSlackBotApiMessage) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetText

`func (o *AiSlackBotApiMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiSlackBotApiMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiSlackBotApiMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AiSlackBotApiMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTruncated

`func (o *AiSlackBotApiMessage) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *AiSlackBotApiMessage) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *AiSlackBotApiMessage) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *AiSlackBotApiMessage) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetReplyCount

`func (o *AiSlackBotApiMessage) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *AiSlackBotApiMessage) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *AiSlackBotApiMessage) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *AiSlackBotApiMessage) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetLatestReplyTs

`func (o *AiSlackBotApiMessage) GetLatestReplyTs() string`

GetLatestReplyTs returns the LatestReplyTs field if non-nil, zero value otherwise.

### GetLatestReplyTsOk

`func (o *AiSlackBotApiMessage) GetLatestReplyTsOk() (*string, bool)`

GetLatestReplyTsOk returns a tuple with the LatestReplyTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReplyTs

`func (o *AiSlackBotApiMessage) SetLatestReplyTs(v string)`

SetLatestReplyTs sets LatestReplyTs field to given value.

### HasLatestReplyTs

`func (o *AiSlackBotApiMessage) HasLatestReplyTs() bool`

HasLatestReplyTs returns a boolean if a field has been set.

### GetEdited

`func (o *AiSlackBotApiMessage) GetEdited() bool`

GetEdited returns the Edited field if non-nil, zero value otherwise.

### GetEditedOk

`func (o *AiSlackBotApiMessage) GetEditedOk() (*bool, bool)`

GetEditedOk returns a tuple with the Edited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdited

`func (o *AiSlackBotApiMessage) SetEdited(v bool)`

SetEdited sets Edited field to given value.

### HasEdited

`func (o *AiSlackBotApiMessage) HasEdited() bool`

HasEdited returns a boolean if a field has been set.

### GetFiles

`func (o *AiSlackBotApiMessage) GetFiles() []AiSlackBotApiMessageAllOfFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *AiSlackBotApiMessage) GetFilesOk() (*[]AiSlackBotApiMessageAllOfFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *AiSlackBotApiMessage) SetFiles(v []AiSlackBotApiMessageAllOfFiles)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *AiSlackBotApiMessage) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


