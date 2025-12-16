# ListAIToolExecutions200ResponseExecutionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **string** |  | [optional] 
**ToolName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Execution status: pending, running, complete, or failed | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListAIToolExecutions200ResponseExecutionsInner

`func NewListAIToolExecutions200ResponseExecutionsInner() *ListAIToolExecutions200ResponseExecutionsInner`

NewListAIToolExecutions200ResponseExecutionsInner instantiates a new ListAIToolExecutions200ResponseExecutionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAIToolExecutions200ResponseExecutionsInnerWithDefaults

`func NewListAIToolExecutions200ResponseExecutionsInnerWithDefaults() *ListAIToolExecutions200ResponseExecutionsInner`

NewListAIToolExecutions200ResponseExecutionsInnerWithDefaults instantiates a new ListAIToolExecutions200ResponseExecutionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ListAIToolExecutions200ResponseExecutionsInner) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ListAIToolExecutions200ResponseExecutionsInner) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetToolName

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *ListAIToolExecutions200ResponseExecutionsInner) SetToolName(v string)`

SetToolName sets ToolName field to given value.

### HasToolName

`func (o *ListAIToolExecutions200ResponseExecutionsInner) HasToolName() bool`

HasToolName returns a boolean if a field has been set.

### GetStatus

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAIToolExecutions200ResponseExecutionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListAIToolExecutions200ResponseExecutionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListAIToolExecutions200ResponseExecutionsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListAIToolExecutions200ResponseExecutionsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ListAIToolExecutions200ResponseExecutionsInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ListAIToolExecutions200ResponseExecutionsInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ListAIToolExecutions200ResponseExecutionsInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


