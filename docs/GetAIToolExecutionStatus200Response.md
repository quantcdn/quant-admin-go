# GetAIToolExecutionStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** |  | 
**ToolName** | **string** |  | 
**Status** | **string** | Execution status: pending, running, complete, or failed | 
**Result** | Pointer to [**GetAIToolExecutionStatus200ResponseResult**](GetAIToolExecutionStatus200ResponseResult.md) |  | [optional] 
**Error** | Pointer to **string** | Error message (only present when status&#x3D;&#39;failed&#39;) | [optional] 
**CreatedAt** | **int32** | Unix timestamp when execution was created | 
**StartedAt** | Pointer to **int32** | Unix timestamp when execution started (if status &gt;&#x3D; &#39;running&#39;) | [optional] 
**CompletedAt** | Pointer to **int32** | Unix timestamp when execution completed (if status in [&#39;complete&#39;, &#39;failed&#39;]) | [optional] 
**Duration** | Pointer to **float32** | Execution duration in seconds (if completed) | [optional] 

## Methods

### NewGetAIToolExecutionStatus200Response

`func NewGetAIToolExecutionStatus200Response(executionId string, toolName string, status string, createdAt int32, ) *GetAIToolExecutionStatus200Response`

NewGetAIToolExecutionStatus200Response instantiates a new GetAIToolExecutionStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAIToolExecutionStatus200ResponseWithDefaults

`func NewGetAIToolExecutionStatus200ResponseWithDefaults() *GetAIToolExecutionStatus200Response`

NewGetAIToolExecutionStatus200ResponseWithDefaults instantiates a new GetAIToolExecutionStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *GetAIToolExecutionStatus200Response) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *GetAIToolExecutionStatus200Response) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *GetAIToolExecutionStatus200Response) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetToolName

`func (o *GetAIToolExecutionStatus200Response) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *GetAIToolExecutionStatus200Response) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *GetAIToolExecutionStatus200Response) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### GetStatus

`func (o *GetAIToolExecutionStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAIToolExecutionStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAIToolExecutionStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResult

`func (o *GetAIToolExecutionStatus200Response) GetResult() GetAIToolExecutionStatus200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAIToolExecutionStatus200Response) GetResultOk() (*GetAIToolExecutionStatus200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAIToolExecutionStatus200Response) SetResult(v GetAIToolExecutionStatus200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAIToolExecutionStatus200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *GetAIToolExecutionStatus200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAIToolExecutionStatus200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAIToolExecutionStatus200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetAIToolExecutionStatus200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetAIToolExecutionStatus200Response) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAIToolExecutionStatus200Response) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAIToolExecutionStatus200Response) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *GetAIToolExecutionStatus200Response) GetStartedAt() int32`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetAIToolExecutionStatus200Response) GetStartedAtOk() (*int32, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetAIToolExecutionStatus200Response) SetStartedAt(v int32)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetAIToolExecutionStatus200Response) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetAIToolExecutionStatus200Response) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetAIToolExecutionStatus200Response) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetAIToolExecutionStatus200Response) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetAIToolExecutionStatus200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetDuration

`func (o *GetAIToolExecutionStatus200Response) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetAIToolExecutionStatus200Response) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetAIToolExecutionStatus200Response) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetAIToolExecutionStatus200Response) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


