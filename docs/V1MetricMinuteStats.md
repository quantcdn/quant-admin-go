# V1MetricMinuteStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**MinuteTotal** | **int32** |  | 
**MinuteSeries** | **map[string]float32** |  | 
**MinuteAverage** | **float32** |  | 

## Methods

### NewV1MetricMinuteStats

`func NewV1MetricMinuteStats(total int32, minuteTotal int32, minuteSeries map[string]float32, minuteAverage float32, ) *V1MetricMinuteStats`

NewV1MetricMinuteStats instantiates a new V1MetricMinuteStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetricMinuteStatsWithDefaults

`func NewV1MetricMinuteStatsWithDefaults() *V1MetricMinuteStats`

NewV1MetricMinuteStatsWithDefaults instantiates a new V1MetricMinuteStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V1MetricMinuteStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1MetricMinuteStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1MetricMinuteStats) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetMinuteTotal

`func (o *V1MetricMinuteStats) GetMinuteTotal() int32`

GetMinuteTotal returns the MinuteTotal field if non-nil, zero value otherwise.

### GetMinuteTotalOk

`func (o *V1MetricMinuteStats) GetMinuteTotalOk() (*int32, bool)`

GetMinuteTotalOk returns a tuple with the MinuteTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinuteTotal

`func (o *V1MetricMinuteStats) SetMinuteTotal(v int32)`

SetMinuteTotal sets MinuteTotal field to given value.


### GetMinuteSeries

`func (o *V1MetricMinuteStats) GetMinuteSeries() map[string]float32`

GetMinuteSeries returns the MinuteSeries field if non-nil, zero value otherwise.

### GetMinuteSeriesOk

`func (o *V1MetricMinuteStats) GetMinuteSeriesOk() (*map[string]float32, bool)`

GetMinuteSeriesOk returns a tuple with the MinuteSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinuteSeries

`func (o *V1MetricMinuteStats) SetMinuteSeries(v map[string]float32)`

SetMinuteSeries sets MinuteSeries field to given value.


### GetMinuteAverage

`func (o *V1MetricMinuteStats) GetMinuteAverage() float32`

GetMinuteAverage returns the MinuteAverage field if non-nil, zero value otherwise.

### GetMinuteAverageOk

`func (o *V1MetricMinuteStats) GetMinuteAverageOk() (*float32, bool)`

GetMinuteAverageOk returns a tuple with the MinuteAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinuteAverage

`func (o *V1MetricMinuteStats) SetMinuteAverage(v float32)`

SetMinuteAverage sets MinuteAverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


