# V1MetricDayStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**DayTotal** | **int32** |  | 
**DaySeries** | **map[string]float32** |  | 
**DayAverage** | **float32** |  | 

## Methods

### NewV1MetricDayStats

`func NewV1MetricDayStats(total int32, dayTotal int32, daySeries map[string]float32, dayAverage float32, ) *V1MetricDayStats`

NewV1MetricDayStats instantiates a new V1MetricDayStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetricDayStatsWithDefaults

`func NewV1MetricDayStatsWithDefaults() *V1MetricDayStats`

NewV1MetricDayStatsWithDefaults instantiates a new V1MetricDayStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V1MetricDayStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1MetricDayStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1MetricDayStats) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetDayTotal

`func (o *V1MetricDayStats) GetDayTotal() int32`

GetDayTotal returns the DayTotal field if non-nil, zero value otherwise.

### GetDayTotalOk

`func (o *V1MetricDayStats) GetDayTotalOk() (*int32, bool)`

GetDayTotalOk returns a tuple with the DayTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTotal

`func (o *V1MetricDayStats) SetDayTotal(v int32)`

SetDayTotal sets DayTotal field to given value.


### GetDaySeries

`func (o *V1MetricDayStats) GetDaySeries() map[string]float32`

GetDaySeries returns the DaySeries field if non-nil, zero value otherwise.

### GetDaySeriesOk

`func (o *V1MetricDayStats) GetDaySeriesOk() (*map[string]float32, bool)`

GetDaySeriesOk returns a tuple with the DaySeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaySeries

`func (o *V1MetricDayStats) SetDaySeries(v map[string]float32)`

SetDaySeries sets DaySeries field to given value.


### GetDayAverage

`func (o *V1MetricDayStats) GetDayAverage() float32`

GetDayAverage returns the DayAverage field if non-nil, zero value otherwise.

### GetDayAverageOk

`func (o *V1MetricDayStats) GetDayAverageOk() (*float32, bool)`

GetDayAverageOk returns a tuple with the DayAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayAverage

`func (o *V1MetricDayStats) SetDayAverage(v float32)`

SetDayAverage sets DayAverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


