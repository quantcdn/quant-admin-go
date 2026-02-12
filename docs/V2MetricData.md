# V2MetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | [**[]V2MetricDataPoint**](V2MetricDataPoint.md) | Time series data points | 
**PeriodTotal** | **float32** | Total value for the period | 
**AllTimeTotal** | **float32** | All-time total value | 
**PeriodAverage** | **float32** | Average value per time unit in the period | 

## Methods

### NewV2MetricData

`func NewV2MetricData(series []V2MetricDataPoint, periodTotal float32, allTimeTotal float32, periodAverage float32, ) *V2MetricData`

NewV2MetricData instantiates a new V2MetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MetricDataWithDefaults

`func NewV2MetricDataWithDefaults() *V2MetricData`

NewV2MetricDataWithDefaults instantiates a new V2MetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *V2MetricData) GetSeries() []V2MetricDataPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *V2MetricData) GetSeriesOk() (*[]V2MetricDataPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *V2MetricData) SetSeries(v []V2MetricDataPoint)`

SetSeries sets Series field to given value.


### GetPeriodTotal

`func (o *V2MetricData) GetPeriodTotal() float32`

GetPeriodTotal returns the PeriodTotal field if non-nil, zero value otherwise.

### GetPeriodTotalOk

`func (o *V2MetricData) GetPeriodTotalOk() (*float32, bool)`

GetPeriodTotalOk returns a tuple with the PeriodTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodTotal

`func (o *V2MetricData) SetPeriodTotal(v float32)`

SetPeriodTotal sets PeriodTotal field to given value.


### GetAllTimeTotal

`func (o *V2MetricData) GetAllTimeTotal() float32`

GetAllTimeTotal returns the AllTimeTotal field if non-nil, zero value otherwise.

### GetAllTimeTotalOk

`func (o *V2MetricData) GetAllTimeTotalOk() (*float32, bool)`

GetAllTimeTotalOk returns a tuple with the AllTimeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTimeTotal

`func (o *V2MetricData) SetAllTimeTotal(v float32)`

SetAllTimeTotal sets AllTimeTotal field to given value.


### GetPeriodAverage

`func (o *V2MetricData) GetPeriodAverage() float32`

GetPeriodAverage returns the PeriodAverage field if non-nil, zero value otherwise.

### GetPeriodAverageOk

`func (o *V2MetricData) GetPeriodAverageOk() (*float32, bool)`

GetPeriodAverageOk returns a tuple with the PeriodAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodAverage

`func (o *V2MetricData) SetPeriodAverage(v float32)`

SetPeriodAverage sets PeriodAverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


