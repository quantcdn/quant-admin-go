# V2SearchExtractors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | XPath for the record title | [optional] [default to "//h1[1]"]
**Content** | Pointer to **string** | XPath for the record body | [optional] [default to "//meta[@name='description'][1]/@content"]
**Summary** | Pointer to **string** | XPath for the record summary | [optional] [default to "//meta[@name='description'][1]/@content"]
**Image** | Pointer to **string** | XPath for the record image | [optional] [default to "//meta[@property='og:image'][1]/@content"]

## Methods

### NewV2SearchExtractors

`func NewV2SearchExtractors() *V2SearchExtractors`

NewV2SearchExtractors instantiates a new V2SearchExtractors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SearchExtractorsWithDefaults

`func NewV2SearchExtractorsWithDefaults() *V2SearchExtractors`

NewV2SearchExtractorsWithDefaults instantiates a new V2SearchExtractors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *V2SearchExtractors) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V2SearchExtractors) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V2SearchExtractors) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V2SearchExtractors) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *V2SearchExtractors) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V2SearchExtractors) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V2SearchExtractors) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *V2SearchExtractors) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSummary

`func (o *V2SearchExtractors) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *V2SearchExtractors) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *V2SearchExtractors) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *V2SearchExtractors) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetImage

`func (o *V2SearchExtractors) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V2SearchExtractors) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V2SearchExtractors) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V2SearchExtractors) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


