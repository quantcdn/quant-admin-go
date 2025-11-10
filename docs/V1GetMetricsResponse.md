# V1GetMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Data** | [**V1GetMetricsResponseData**](V1GetMetricsResponseData.md) |  | 

## Methods

### NewV1GetMetricsResponse

`func NewV1GetMetricsResponse(success bool, data V1GetMetricsResponseData, ) *V1GetMetricsResponse`

NewV1GetMetricsResponse instantiates a new V1GetMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetMetricsResponseWithDefaults

`func NewV1GetMetricsResponseWithDefaults() *V1GetMetricsResponse`

NewV1GetMetricsResponseWithDefaults instantiates a new V1GetMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *V1GetMetricsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V1GetMetricsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V1GetMetricsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *V1GetMetricsResponse) GetData() V1GetMetricsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1GetMetricsResponse) GetDataOk() (*V1GetMetricsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1GetMetricsResponse) SetData(v V1GetMetricsResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


