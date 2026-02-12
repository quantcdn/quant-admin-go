# V2MetricsMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | **string** | The period type for this data | 
**Granularity** | **string** | The granularity of data points | 
**StartTime** | **string** | Start time of the data range (ISO8601 or Unix timestamp based on timestamp_format parameter) | 
**EndTime** | **string** | End time of the data range (ISO8601 or Unix timestamp based on timestamp_format parameter) | 
**Metrics** | **[]string** | List of metrics included in the response | 
**Domain** | Pointer to **NullableString** | Domain filter applied (if any) | [optional] 

## Methods

### NewV2MetricsMeta

`func NewV2MetricsMeta(period string, granularity string, startTime string, endTime string, metrics []string, ) *V2MetricsMeta`

NewV2MetricsMeta instantiates a new V2MetricsMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MetricsMetaWithDefaults

`func NewV2MetricsMetaWithDefaults() *V2MetricsMeta`

NewV2MetricsMetaWithDefaults instantiates a new V2MetricsMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *V2MetricsMeta) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *V2MetricsMeta) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *V2MetricsMeta) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetGranularity

`func (o *V2MetricsMeta) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *V2MetricsMeta) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *V2MetricsMeta) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.


### GetStartTime

`func (o *V2MetricsMeta) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V2MetricsMeta) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V2MetricsMeta) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *V2MetricsMeta) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V2MetricsMeta) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V2MetricsMeta) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetMetrics

`func (o *V2MetricsMeta) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *V2MetricsMeta) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *V2MetricsMeta) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.


### GetDomain

`func (o *V2MetricsMeta) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2MetricsMeta) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2MetricsMeta) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V2MetricsMeta) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *V2MetricsMeta) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *V2MetricsMeta) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


