# V2CrawlerScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Name** | **string** | Schedule name | 
**ScheduleCronString** | **string** | Cron schedule string | 

## Methods

### NewV2CrawlerScheduleRequest

`func NewV2CrawlerScheduleRequest(message string, error_ bool, name string, scheduleCronString string, ) *V2CrawlerScheduleRequest`

NewV2CrawlerScheduleRequest instantiates a new V2CrawlerScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2CrawlerScheduleRequestWithDefaults

`func NewV2CrawlerScheduleRequestWithDefaults() *V2CrawlerScheduleRequest`

NewV2CrawlerScheduleRequestWithDefaults instantiates a new V2CrawlerScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2CrawlerScheduleRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2CrawlerScheduleRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2CrawlerScheduleRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2CrawlerScheduleRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2CrawlerScheduleRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2CrawlerScheduleRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetName

`func (o *V2CrawlerScheduleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2CrawlerScheduleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2CrawlerScheduleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScheduleCronString

`func (o *V2CrawlerScheduleRequest) GetScheduleCronString() string`

GetScheduleCronString returns the ScheduleCronString field if non-nil, zero value otherwise.

### GetScheduleCronStringOk

`func (o *V2CrawlerScheduleRequest) GetScheduleCronStringOk() (*string, bool)`

GetScheduleCronStringOk returns a tuple with the ScheduleCronString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleCronString

`func (o *V2CrawlerScheduleRequest) SetScheduleCronString(v string)`

SetScheduleCronString sets ScheduleCronString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


