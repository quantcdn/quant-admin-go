# GetCrawlRun200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current status of the run (provisioning, running, stopped, canceled). | [optional] 
**StartedAt** | Pointer to **int32** | Unix timestamp denoting the start date and time of the run. | [optional] 
**CompletedAt** | Pointer to **int32** | Unix timestamp denoting the finish date and time of the run. Value is null if the run is incomplete. | [optional] 

## Methods

### NewGetCrawlRun200Response

`func NewGetCrawlRun200Response() *GetCrawlRun200Response`

NewGetCrawlRun200Response instantiates a new GetCrawlRun200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCrawlRun200ResponseWithDefaults

`func NewGetCrawlRun200ResponseWithDefaults() *GetCrawlRun200Response`

NewGetCrawlRun200ResponseWithDefaults instantiates a new GetCrawlRun200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetCrawlRun200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCrawlRun200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCrawlRun200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCrawlRun200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetCrawlRun200Response) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetCrawlRun200Response) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetCrawlRun200Response) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetCrawlRun200Response) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetCrawlRun200Response) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetCrawlRun200Response) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetCrawlRun200Response) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetCrawlRun200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


