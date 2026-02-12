# V2MetricDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | [**V2MetricDataPointTimestamp**](V2MetricDataPointTimestamp.md) |  | 
**Value** | **float32** | Metric value at this timestamp | 

## Methods

### NewV2MetricDataPoint

`func NewV2MetricDataPoint(timestamp V2MetricDataPointTimestamp, value float32, ) *V2MetricDataPoint`

NewV2MetricDataPoint instantiates a new V2MetricDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MetricDataPointWithDefaults

`func NewV2MetricDataPointWithDefaults() *V2MetricDataPoint`

NewV2MetricDataPointWithDefaults instantiates a new V2MetricDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *V2MetricDataPoint) GetTimestamp() V2MetricDataPointTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *V2MetricDataPoint) GetTimestampOk() (*V2MetricDataPointTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *V2MetricDataPoint) SetTimestamp(v V2MetricDataPointTimestamp)`

SetTimestamp sets Timestamp field to given value.


### GetValue

`func (o *V2MetricDataPoint) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2MetricDataPoint) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2MetricDataPoint) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


