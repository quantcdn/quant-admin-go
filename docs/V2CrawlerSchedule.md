# V2CrawlerSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Id** | **int32** | Schedule ID | 
**Name** | Pointer to **string** | Schedule name | [optional] 
**CrawlerConfigId** | **int32** | Crawler config ID | 
**ProjectId** | **int32** | Project ID | 
**CrawlerLastRunId** | **int32** | Last run ID | 
**ScheduleCronString** | **string** | Cron schedule string | 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update timestamp | [optional] 

## Methods

### NewV2CrawlerSchedule

`func NewV2CrawlerSchedule(message string, error_ bool, id int32, crawlerConfigId int32, projectId int32, crawlerLastRunId int32, scheduleCronString string, ) *V2CrawlerSchedule`

NewV2CrawlerSchedule instantiates a new V2CrawlerSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CrawlerScheduleWithDefaults

`func NewV2CrawlerScheduleWithDefaults() *V2CrawlerSchedule`

NewV2CrawlerScheduleWithDefaults instantiates a new V2CrawlerSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2CrawlerSchedule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2CrawlerSchedule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2CrawlerSchedule) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2CrawlerSchedule) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2CrawlerSchedule) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2CrawlerSchedule) SetError(v bool)`

SetError sets Error field to given value.


### GetId

`func (o *V2CrawlerSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2CrawlerSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2CrawlerSchedule) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *V2CrawlerSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2CrawlerSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2CrawlerSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2CrawlerSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCrawlerConfigId

`func (o *V2CrawlerSchedule) GetCrawlerConfigId() int32`

GetCrawlerConfigId returns the CrawlerConfigId field if non-nil, zero value otherwise.

### GetCrawlerConfigIdOk

`func (o *V2CrawlerSchedule) GetCrawlerConfigIdOk() (*int32, bool)`

GetCrawlerConfigIdOk returns a tuple with the CrawlerConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerConfigId

`func (o *V2CrawlerSchedule) SetCrawlerConfigId(v int32)`

SetCrawlerConfigId sets CrawlerConfigId field to given value.


### GetProjectId

`func (o *V2CrawlerSchedule) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *V2CrawlerSchedule) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *V2CrawlerSchedule) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetCrawlerLastRunId

`func (o *V2CrawlerSchedule) GetCrawlerLastRunId() int32`

GetCrawlerLastRunId returns the CrawlerLastRunId field if non-nil, zero value otherwise.

### GetCrawlerLastRunIdOk

`func (o *V2CrawlerSchedule) GetCrawlerLastRunIdOk() (*int32, bool)`

GetCrawlerLastRunIdOk returns a tuple with the CrawlerLastRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlerLastRunId

`func (o *V2CrawlerSchedule) SetCrawlerLastRunId(v int32)`

SetCrawlerLastRunId sets CrawlerLastRunId field to given value.


### GetScheduleCronString

`func (o *V2CrawlerSchedule) GetScheduleCronString() string`

GetScheduleCronString returns the ScheduleCronString field if non-nil, zero value otherwise.

### GetScheduleCronStringOk

`func (o *V2CrawlerSchedule) GetScheduleCronStringOk() (*string, bool)`

GetScheduleCronStringOk returns a tuple with the ScheduleCronString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCronString

`func (o *V2CrawlerSchedule) SetScheduleCronString(v string)`

SetScheduleCronString sets ScheduleCronString field to given value.


### GetCreatedAt

`func (o *V2CrawlerSchedule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *V2CrawlerSchedule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *V2CrawlerSchedule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *V2CrawlerSchedule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *V2CrawlerSchedule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *V2CrawlerSchedule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *V2CrawlerSchedule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *V2CrawlerSchedule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


