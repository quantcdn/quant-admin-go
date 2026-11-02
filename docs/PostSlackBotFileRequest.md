# PostSlackBotFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**ThreadTs** | Pointer to **string** | Absent: a root post | [optional] 
**Content** | **string** | The image, base64 (standard alphabet, no whitespace, no data: prefix). 1 byte to 4 MiB decoded. | 
**Filename** | Pointer to **string** | Default screenshot.&lt;ext&gt;. The extension must match the image type (png, jpg, jpeg, gif, webp). | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**AltText** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Caption. Same mention rules as a message. | [optional] 

## Methods

### NewPostSlackBotFileRequest

`func NewPostSlackBotFileRequest(channel string, content string, ) *PostSlackBotFileRequest`

NewPostSlackBotFileRequest instantiates a new PostSlackBotFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSlackBotFileRequestWithDefaults

`func NewPostSlackBotFileRequestWithDefaults() *PostSlackBotFileRequest`

NewPostSlackBotFileRequestWithDefaults instantiates a new PostSlackBotFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *PostSlackBotFileRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostSlackBotFileRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostSlackBotFileRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetThreadTs

`func (o *PostSlackBotFileRequest) GetThreadTs() string`

GetThreadTs returns the ThreadTs field if non-nil, zero value otherwise.

### GetThreadTsOk

`func (o *PostSlackBotFileRequest) GetThreadTsOk() (*string, bool)`

GetThreadTsOk returns a tuple with the ThreadTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadTs

`func (o *PostSlackBotFileRequest) SetThreadTs(v string)`

SetThreadTs sets ThreadTs field to given value.

### HasThreadTs

`func (o *PostSlackBotFileRequest) HasThreadTs() bool`

HasThreadTs returns a boolean if a field has been set.

### GetContent

`func (o *PostSlackBotFileRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PostSlackBotFileRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PostSlackBotFileRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetFilename

`func (o *PostSlackBotFileRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PostSlackBotFileRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PostSlackBotFileRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PostSlackBotFileRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetTitle

`func (o *PostSlackBotFileRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostSlackBotFileRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostSlackBotFileRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PostSlackBotFileRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAltText

`func (o *PostSlackBotFileRequest) GetAltText() string`

GetAltText returns the AltText field if non-nil, zero value otherwise.

### GetAltTextOk

`func (o *PostSlackBotFileRequest) GetAltTextOk() (*string, bool)`

GetAltTextOk returns a tuple with the AltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltText

`func (o *PostSlackBotFileRequest) SetAltText(v string)`

SetAltText sets AltText field to given value.

### HasAltText

`func (o *PostSlackBotFileRequest) HasAltText() bool`

HasAltText returns a boolean if a field has been set.

### GetText

`func (o *PostSlackBotFileRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostSlackBotFileRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostSlackBotFileRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostSlackBotFileRequest) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


