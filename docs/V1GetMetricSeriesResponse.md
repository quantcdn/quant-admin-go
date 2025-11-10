# V1GetMetricSeriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**V1GetMetricSeriesResponseData**](V1GetMetricSeriesResponseData.md) |  | 

## Methods

### NewV1GetMetricSeriesResponse

`func NewV1GetMetricSeriesResponse(success bool, data V1GetMetricSeriesResponseData, ) *V1GetMetricSeriesResponse`

NewV1GetMetricSeriesResponse instantiates a new V1GetMetricSeriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetMetricSeriesResponseWithDefaults

`func NewV1GetMetricSeriesResponseWithDefaults() *V1GetMetricSeriesResponse`

NewV1GetMetricSeriesResponseWithDefaults instantiates a new V1GetMetricSeriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V1GetMetricSeriesResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V1GetMetricSeriesResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V1GetMetricSeriesResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *V1GetMetricSeriesResponse) GetData() V1GetMetricSeriesResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1GetMetricSeriesResponse) GetDataOk() (*V1GetMetricSeriesResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1GetMetricSeriesResponse) SetData(v V1GetMetricSeriesResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


