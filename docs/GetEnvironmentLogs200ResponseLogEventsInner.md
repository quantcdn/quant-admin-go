# GetEnvironmentLogs200ResponseLogEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**Message** | Pointer to **string** | Log message content | [optional] 
**IngestionTime** | Pointer to **NullableInt64** | Unix timestamp in milliseconds when CloudWatch ingested the event | [optional] 
**LogStreamName** | Pointer to **NullableString** | CloudWatch log stream, named container/container/taskId | [optional] 

## Methods

### NewGetEnvironmentLogs200ResponseLogEventsInner

`func NewGetEnvironmentLogs200ResponseLogEventsInner() *GetEnvironmentLogs200ResponseLogEventsInner`

NewGetEnvironmentLogs200ResponseLogEventsInner instantiates a new GetEnvironmentLogs200ResponseLogEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnvironmentLogs200ResponseLogEventsInnerWithDefaults

`func NewGetEnvironmentLogs200ResponseLogEventsInnerWithDefaults() *GetEnvironmentLogs200ResponseLogEventsInner`

NewGetEnvironmentLogs200ResponseLogEventsInnerWithDefaults instantiates a new GetEnvironmentLogs200ResponseLogEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMessage

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetIngestionTime

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) GetIngestionTime() int64`

GetIngestionTime returns the IngestionTime field if non-nil, zero value otherwise.

### GetIngestionTimeOk

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) GetIngestionTimeOk() (*int64, bool)`

GetIngestionTimeOk returns a tuple with the IngestionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionTime

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) SetIngestionTime(v int64)`

SetIngestionTime sets IngestionTime field to given value.

### HasIngestionTime

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) HasIngestionTime() bool`

HasIngestionTime returns a boolean if a field has been set.

### SetIngestionTimeNil

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) SetIngestionTimeNil(b bool)`

 SetIngestionTimeNil sets the value for IngestionTime to be an explicit nil

### UnsetIngestionTime
`func (o *GetEnvironmentLogs200ResponseLogEventsInner) UnsetIngestionTime()`

UnsetIngestionTime ensures that no value is present for IngestionTime, not even an explicit nil
### GetLogStreamName

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) GetLogStreamName() string`

GetLogStreamName returns the LogStreamName field if non-nil, zero value otherwise.

### GetLogStreamNameOk

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) GetLogStreamNameOk() (*string, bool)`

GetLogStreamNameOk returns a tuple with the LogStreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStreamName

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) SetLogStreamName(v string)`

SetLogStreamName sets LogStreamName field to given value.

### HasLogStreamName

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) HasLogStreamName() bool`

HasLogStreamName returns a boolean if a field has been set.

### SetLogStreamNameNil

`func (o *GetEnvironmentLogs200ResponseLogEventsInner) SetLogStreamNameNil(b bool)`

 SetLogStreamNameNil sets the value for LogStreamName to be an explicit nil

### UnsetLogStreamName
`func (o *GetEnvironmentLogs200ResponseLogEventsInner) UnsetLogStreamName()`

UnsetLogStreamName ensures that no value is present for LogStreamName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


