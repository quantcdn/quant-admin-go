# V1MetricMonthStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**MonthTotal** | **int32** |  | 
**MonthSeries** | **map[string]float32** |  | 
**MonthAverage** | **float32** |  | 

## Methods

### NewV1MetricMonthStats

`func NewV1MetricMonthStats(total int32, monthTotal int32, monthSeries map[string]float32, monthAverage float32, ) *V1MetricMonthStats`

NewV1MetricMonthStats instantiates a new V1MetricMonthStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MetricMonthStatsWithDefaults

`func NewV1MetricMonthStatsWithDefaults() *V1MetricMonthStats`

NewV1MetricMonthStatsWithDefaults instantiates a new V1MetricMonthStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *V1MetricMonthStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1MetricMonthStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1MetricMonthStats) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetMonthTotal

`func (o *V1MetricMonthStats) GetMonthTotal() int32`

GetMonthTotal returns the MonthTotal field if non-nil, zero value otherwise.

### GetMonthTotalOk

`func (o *V1MetricMonthStats) GetMonthTotalOk() (*int32, bool)`

GetMonthTotalOk returns a tuple with the MonthTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthTotal

`func (o *V1MetricMonthStats) SetMonthTotal(v int32)`

SetMonthTotal sets MonthTotal field to given value.


### GetMonthSeries

`func (o *V1MetricMonthStats) GetMonthSeries() map[string]float32`

GetMonthSeries returns the MonthSeries field if non-nil, zero value otherwise.

### GetMonthSeriesOk

`func (o *V1MetricMonthStats) GetMonthSeriesOk() (*map[string]float32, bool)`

GetMonthSeriesOk returns a tuple with the MonthSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthSeries

`func (o *V1MetricMonthStats) SetMonthSeries(v map[string]float32)`

SetMonthSeries sets MonthSeries field to given value.


### GetMonthAverage

`func (o *V1MetricMonthStats) GetMonthAverage() float32`

GetMonthAverage returns the MonthAverage field if non-nil, zero value otherwise.

### GetMonthAverageOk

`func (o *V1MetricMonthStats) GetMonthAverageOk() (*float32, bool)`

GetMonthAverageOk returns a tuple with the MonthAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthAverage

`func (o *V1MetricMonthStats) SetMonthAverage(v float32)`

SetMonthAverage sets MonthAverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


