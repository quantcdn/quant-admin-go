# AiSearchIngestPagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** |  | [optional] 
**Pages** | [**[]AiSearchIngestPagesRequestPagesInner**](AiSearchIngestPagesRequestPagesInner.md) |  | 

## Methods

### NewAiSearchIngestPagesRequest

`func NewAiSearchIngestPagesRequest(pages []AiSearchIngestPagesRequestPagesInner, ) *AiSearchIngestPagesRequest`

NewAiSearchIngestPagesRequest instantiates a new AiSearchIngestPagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSearchIngestPagesRequestWithDefaults

`func NewAiSearchIngestPagesRequestWithDefaults() *AiSearchIngestPagesRequest`

NewAiSearchIngestPagesRequestWithDefaults instantiates a new AiSearchIngestPagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *AiSearchIngestPagesRequest) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *AiSearchIngestPagesRequest) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *AiSearchIngestPagesRequest) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *AiSearchIngestPagesRequest) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetPages

`func (o *AiSearchIngestPagesRequest) GetPages() []AiSearchIngestPagesRequestPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *AiSearchIngestPagesRequest) GetPagesOk() (*[]AiSearchIngestPagesRequestPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *AiSearchIngestPagesRequest) SetPages(v []AiSearchIngestPagesRequestPagesInner)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


