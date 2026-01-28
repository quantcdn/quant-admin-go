# SubmitToolCallbackRequestToolResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolUseId** | **string** | The toolUseId from pendingTools | 
**Result** | **map[string]interface{}** | The result of executing the tool | 

## Methods

### NewSubmitToolCallbackRequestToolResultsInner

`func NewSubmitToolCallbackRequestToolResultsInner(toolUseId string, result map[string]interface{}, ) *SubmitToolCallbackRequestToolResultsInner`

NewSubmitToolCallbackRequestToolResultsInner instantiates a new SubmitToolCallbackRequestToolResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitToolCallbackRequestToolResultsInnerWithDefaults

`func NewSubmitToolCallbackRequestToolResultsInnerWithDefaults() *SubmitToolCallbackRequestToolResultsInner`

NewSubmitToolCallbackRequestToolResultsInnerWithDefaults instantiates a new SubmitToolCallbackRequestToolResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolUseId

`func (o *SubmitToolCallbackRequestToolResultsInner) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *SubmitToolCallbackRequestToolResultsInner) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *SubmitToolCallbackRequestToolResultsInner) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetResult

`func (o *SubmitToolCallbackRequestToolResultsInner) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SubmitToolCallbackRequestToolResultsInner) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SubmitToolCallbackRequestToolResultsInner) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


