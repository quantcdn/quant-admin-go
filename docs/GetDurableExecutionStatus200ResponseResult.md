# GetDurableExecutionStatus200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**GetDurableExecutionStatus200ResponseResultResponse**](GetDurableExecutionStatus200ResponseResultResponse.md) |  | [optional] 
**Usage** | Pointer to [**GetDurableExecutionStatus200ResponseResultUsage**](GetDurableExecutionStatus200ResponseResultUsage.md) |  | [optional] 
**ToolExecutions** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetDurableExecutionStatus200ResponseResult

`func NewGetDurableExecutionStatus200ResponseResult() *GetDurableExecutionStatus200ResponseResult`

NewGetDurableExecutionStatus200ResponseResult instantiates a new GetDurableExecutionStatus200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDurableExecutionStatus200ResponseResultWithDefaults

`func NewGetDurableExecutionStatus200ResponseResultWithDefaults() *GetDurableExecutionStatus200ResponseResult`

NewGetDurableExecutionStatus200ResponseResultWithDefaults instantiates a new GetDurableExecutionStatus200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *GetDurableExecutionStatus200ResponseResult) GetResponse() GetDurableExecutionStatus200ResponseResultResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GetDurableExecutionStatus200ResponseResult) GetResponseOk() (*GetDurableExecutionStatus200ResponseResultResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GetDurableExecutionStatus200ResponseResult) SetResponse(v GetDurableExecutionStatus200ResponseResultResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GetDurableExecutionStatus200ResponseResult) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetUsage

`func (o *GetDurableExecutionStatus200ResponseResult) GetUsage() GetDurableExecutionStatus200ResponseResultUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetDurableExecutionStatus200ResponseResult) GetUsageOk() (*GetDurableExecutionStatus200ResponseResultUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetDurableExecutionStatus200ResponseResult) SetUsage(v GetDurableExecutionStatus200ResponseResultUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetDurableExecutionStatus200ResponseResult) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetToolExecutions

`func (o *GetDurableExecutionStatus200ResponseResult) GetToolExecutions() []map[string]interface{}`

GetToolExecutions returns the ToolExecutions field if non-nil, zero value otherwise.

### GetToolExecutionsOk

`func (o *GetDurableExecutionStatus200ResponseResult) GetToolExecutionsOk() (*[]map[string]interface{}, bool)`

GetToolExecutionsOk returns a tuple with the ToolExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolExecutions

`func (o *GetDurableExecutionStatus200ResponseResult) SetToolExecutions(v []map[string]interface{})`

SetToolExecutions sets ToolExecutions field to given value.

### HasToolExecutions

`func (o *GetDurableExecutionStatus200ResponseResult) HasToolExecutions() bool`

HasToolExecutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


