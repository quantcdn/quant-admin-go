# V2CrawlerRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Run ID | 
**CrawlerConfigId** | **int32** | Crawler config ID | 
**ProjectId** | **int32** | Project ID | 
**LastStatus** | **string** | Run status | 
**TaskId** | **string** | Task ID | 
**StartedAt** | Pointer to **NullableInt32** | Start time (Unix timestamp) | [optional] 
**CompletedAt** | Pointer to **NullableInt32** | Completion time (Unix timestamp) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update timestamp | [optional] 

## Methods

### NewV2CrawlerRun

`func NewV2CrawlerRun(id int32, crawlerConfigId int32, projectId int32, lastStatus string, taskId string, ) *V2CrawlerRun`

NewV2CrawlerRun instantiates a new V2CrawlerRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CrawlerRunWithDefaults

`func NewV2CrawlerRunWithDefaults() *V2CrawlerRun`

NewV2CrawlerRunWithDefaults instantiates a new V2CrawlerRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2CrawlerRun) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2CrawlerRun) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2CrawlerRun) SetId(v int32)`

SetId sets Id field to given value.


### GetCrawlerConfigId

`func (o *V2CrawlerRun) GetCrawlerConfigId() int32`

GetCrawlerConfigId returns the CrawlerConfigId field if non-nil, zero value otherwise.

### GetCrawlerConfigIdOk

`func (o *V2CrawlerRun) GetCrawlerConfigIdOk() (*int32, bool)`

GetCrawlerConfigIdOk returns a tuple with the CrawlerConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerConfigId

`func (o *V2CrawlerRun) SetCrawlerConfigId(v int32)`

SetCrawlerConfigId sets CrawlerConfigId field to given value.


### GetProjectId

`func (o *V2CrawlerRun) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *V2CrawlerRun) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *V2CrawlerRun) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetLastStatus

`func (o *V2CrawlerRun) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *V2CrawlerRun) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *V2CrawlerRun) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.


### GetTaskId

`func (o *V2CrawlerRun) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *V2CrawlerRun) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *V2CrawlerRun) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStartedAt

`func (o *V2CrawlerRun) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V2CrawlerRun) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V2CrawlerRun) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *V2CrawlerRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *V2CrawlerRun) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *V2CrawlerRun) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *V2CrawlerRun) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *V2CrawlerRun) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *V2CrawlerRun) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *V2CrawlerRun) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *V2CrawlerRun) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *V2CrawlerRun) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCreatedAt

`func (o *V2CrawlerRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2CrawlerRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2CrawlerRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2CrawlerRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2CrawlerRun) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2CrawlerRun) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2CrawlerRun) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2CrawlerRun) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


