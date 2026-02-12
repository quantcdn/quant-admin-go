# V2MetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**V2MetricsMeta**](V2MetricsMeta.md) |  | 
**Data** | [**map[string]V2MetricData**](V2MetricData.md) | Metrics data keyed by metric name | 

## Methods

### NewV2MetricsResponse

`func NewV2MetricsResponse(meta V2MetricsMeta, data map[string]V2MetricData, ) *V2MetricsResponse`

NewV2MetricsResponse instantiates a new V2MetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MetricsResponseWithDefaults

`func NewV2MetricsResponseWithDefaults() *V2MetricsResponse`

NewV2MetricsResponseWithDefaults instantiates a new V2MetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *V2MetricsResponse) GetMeta() V2MetricsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V2MetricsResponse) GetMetaOk() (*V2MetricsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V2MetricsResponse) SetMeta(v V2MetricsMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *V2MetricsResponse) GetData() map[string]V2MetricData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2MetricsResponse) GetDataOk() (*map[string]V2MetricData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2MetricsResponse) SetData(v map[string]V2MetricData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


