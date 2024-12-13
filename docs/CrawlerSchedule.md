# CrawlerSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CrawlerConfigId** | **int32** |  | 
**ProjectId** | **int32** |  | 
**CrawlerLastRunId** | **int32** |  | 
**ScheduleCronString** | **string** |  | 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCrawlerSchedule

`func NewCrawlerSchedule(id int32, crawlerConfigId int32, projectId int32, crawlerLastRunId int32, scheduleCronString string, ) *CrawlerSchedule`

NewCrawlerSchedule instantiates a new CrawlerSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlerScheduleWithDefaults

`func NewCrawlerScheduleWithDefaults() *CrawlerSchedule`

NewCrawlerScheduleWithDefaults instantiates a new CrawlerSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CrawlerSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrawlerSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrawlerSchedule) SetId(v int32)`

SetId sets Id field to given value.


### GetCrawlerConfigId

`func (o *CrawlerSchedule) GetCrawlerConfigId() int32`

GetCrawlerConfigId returns the CrawlerConfigId field if non-nil, zero value otherwise.

### GetCrawlerConfigIdOk

`func (o *CrawlerSchedule) GetCrawlerConfigIdOk() (*int32, bool)`

GetCrawlerConfigIdOk returns a tuple with the CrawlerConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerConfigId

`func (o *CrawlerSchedule) SetCrawlerConfigId(v int32)`

SetCrawlerConfigId sets CrawlerConfigId field to given value.


### GetProjectId

`func (o *CrawlerSchedule) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CrawlerSchedule) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CrawlerSchedule) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetCrawlerLastRunId

`func (o *CrawlerSchedule) GetCrawlerLastRunId() int32`

GetCrawlerLastRunId returns the CrawlerLastRunId field if non-nil, zero value otherwise.

### GetCrawlerLastRunIdOk

`func (o *CrawlerSchedule) GetCrawlerLastRunIdOk() (*int32, bool)`

GetCrawlerLastRunIdOk returns a tuple with the CrawlerLastRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerLastRunId

`func (o *CrawlerSchedule) SetCrawlerLastRunId(v int32)`

SetCrawlerLastRunId sets CrawlerLastRunId field to given value.


### GetScheduleCronString

`func (o *CrawlerSchedule) GetScheduleCronString() string`

GetScheduleCronString returns the ScheduleCronString field if non-nil, zero value otherwise.

### GetScheduleCronStringOk

`func (o *CrawlerSchedule) GetScheduleCronStringOk() (*string, bool)`

GetScheduleCronStringOk returns a tuple with the ScheduleCronString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCronString

`func (o *CrawlerSchedule) SetScheduleCronString(v string)`

SetScheduleCronString sets ScheduleCronString field to given value.


### GetCreatedAt

`func (o *CrawlerSchedule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrawlerSchedule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrawlerSchedule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CrawlerSchedule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CrawlerSchedule) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CrawlerSchedule) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CrawlerSchedule) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CrawlerSchedule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CrawlerSchedule) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CrawlerSchedule) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CrawlerSchedule) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CrawlerSchedule) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


