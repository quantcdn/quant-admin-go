# AiSearchIngestPagesRequestPagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Title** | **string** |  | 
**Content** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**FetchedAt** | Pointer to **string** |  | [optional] 
**PreProcessed** | Pointer to **bool** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAiSearchIngestPagesRequestPagesInner

`func NewAiSearchIngestPagesRequestPagesInner(url string, title string, content string, ) *AiSearchIngestPagesRequestPagesInner`

NewAiSearchIngestPagesRequestPagesInner instantiates a new AiSearchIngestPagesRequestPagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSearchIngestPagesRequestPagesInnerWithDefaults

`func NewAiSearchIngestPagesRequestPagesInnerWithDefaults() *AiSearchIngestPagesRequestPagesInner`

NewAiSearchIngestPagesRequestPagesInnerWithDefaults instantiates a new AiSearchIngestPagesRequestPagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AiSearchIngestPagesRequestPagesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AiSearchIngestPagesRequestPagesInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *AiSearchIngestPagesRequestPagesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiSearchIngestPagesRequestPagesInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *AiSearchIngestPagesRequestPagesInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiSearchIngestPagesRequestPagesInner) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentType

`func (o *AiSearchIngestPagesRequestPagesInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AiSearchIngestPagesRequestPagesInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AiSearchIngestPagesRequestPagesInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetFetchedAt

`func (o *AiSearchIngestPagesRequestPagesInner) GetFetchedAt() string`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetFetchedAtOk() (*string, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *AiSearchIngestPagesRequestPagesInner) SetFetchedAt(v string)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *AiSearchIngestPagesRequestPagesInner) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### GetPreProcessed

`func (o *AiSearchIngestPagesRequestPagesInner) GetPreProcessed() bool`

GetPreProcessed returns the PreProcessed field if non-nil, zero value otherwise.

### GetPreProcessedOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetPreProcessedOk() (*bool, bool)`

GetPreProcessedOk returns a tuple with the PreProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProcessed

`func (o *AiSearchIngestPagesRequestPagesInner) SetPreProcessed(v bool)`

SetPreProcessed sets PreProcessed field to given value.

### HasPreProcessed

`func (o *AiSearchIngestPagesRequestPagesInner) HasPreProcessed() bool`

HasPreProcessed returns a boolean if a field has been set.

### GetSummary

`func (o *AiSearchIngestPagesRequestPagesInner) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AiSearchIngestPagesRequestPagesInner) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AiSearchIngestPagesRequestPagesInner) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTags

`func (o *AiSearchIngestPagesRequestPagesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AiSearchIngestPagesRequestPagesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AiSearchIngestPagesRequestPagesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTopics

`func (o *AiSearchIngestPagesRequestPagesInner) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *AiSearchIngestPagesRequestPagesInner) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *AiSearchIngestPagesRequestPagesInner) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *AiSearchIngestPagesRequestPagesInner) HasTopics() bool`

HasTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


