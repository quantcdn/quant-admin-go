# CrawlerScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ScheduleCronString** | **string** |  | 

## Methods

### NewCrawlerScheduleRequest

`func NewCrawlerScheduleRequest(scheduleCronString string, ) *CrawlerScheduleRequest`

NewCrawlerScheduleRequest instantiates a new CrawlerScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlerScheduleRequestWithDefaults

`func NewCrawlerScheduleRequestWithDefaults() *CrawlerScheduleRequest`

NewCrawlerScheduleRequestWithDefaults instantiates a new CrawlerScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CrawlerScheduleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrawlerScheduleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrawlerScheduleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrawlerScheduleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleCronString

`func (o *CrawlerScheduleRequest) GetScheduleCronString() string`

GetScheduleCronString returns the ScheduleCronString field if non-nil, zero value otherwise.

### GetScheduleCronStringOk

`func (o *CrawlerScheduleRequest) GetScheduleCronStringOk() (*string, bool)`

GetScheduleCronStringOk returns a tuple with the ScheduleCronString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCronString

`func (o *CrawlerScheduleRequest) SetScheduleCronString(v string)`

SetScheduleCronString sets ScheduleCronString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


