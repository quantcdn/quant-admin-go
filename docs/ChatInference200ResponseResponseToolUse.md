# ChatInference200ResponseResponseToolUse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolUseId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]interface{}** |  | [optional] 
**ExecutionId** | Pointer to **string** | Present for async tools with autoExecute | [optional] 
**Status** | Pointer to **string** | Execution status (pending/running/complete/failed) - present for async tools with autoExecute | [optional] 
**Result** | Pointer to [**ChatInference200ResponseResponseToolUseOneOfResult**](ChatInference200ResponseResponseToolUseOneOfResult.md) |  | [optional] 

## Methods

### NewChatInference200ResponseResponseToolUse

`func NewChatInference200ResponseResponseToolUse() *ChatInference200ResponseResponseToolUse`

NewChatInference200ResponseResponseToolUse instantiates a new ChatInference200ResponseResponseToolUse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatInference200ResponseResponseToolUseWithDefaults

`func NewChatInference200ResponseResponseToolUseWithDefaults() *ChatInference200ResponseResponseToolUse`

NewChatInference200ResponseResponseToolUseWithDefaults instantiates a new ChatInference200ResponseResponseToolUse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolUseId

`func (o *ChatInference200ResponseResponseToolUse) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *ChatInference200ResponseResponseToolUse) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *ChatInference200ResponseResponseToolUse) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.

### HasToolUseId

`func (o *ChatInference200ResponseResponseToolUse) HasToolUseId() bool`

HasToolUseId returns a boolean if a field has been set.

### GetName

`func (o *ChatInference200ResponseResponseToolUse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatInference200ResponseResponseToolUse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatInference200ResponseResponseToolUse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChatInference200ResponseResponseToolUse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInput

`func (o *ChatInference200ResponseResponseToolUse) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ChatInference200ResponseResponseToolUse) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ChatInference200ResponseResponseToolUse) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *ChatInference200ResponseResponseToolUse) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetExecutionId

`func (o *ChatInference200ResponseResponseToolUse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ChatInference200ResponseResponseToolUse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ChatInference200ResponseResponseToolUse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *ChatInference200ResponseResponseToolUse) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetStatus

`func (o *ChatInference200ResponseResponseToolUse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatInference200ResponseResponseToolUse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatInference200ResponseResponseToolUse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChatInference200ResponseResponseToolUse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ChatInference200ResponseResponseToolUse) GetResult() ChatInference200ResponseResponseToolUseOneOfResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ChatInference200ResponseResponseToolUse) GetResultOk() (*ChatInference200ResponseResponseToolUseOneOfResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ChatInference200ResponseResponseToolUse) SetResult(v ChatInference200ResponseResponseToolUseOneOfResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ChatInference200ResponseResponseToolUse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


