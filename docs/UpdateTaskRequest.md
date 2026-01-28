# UpdateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TaskListId** | Pointer to **NullableString** | Move task to different list or remove from list (set null) | [optional] 
**Status** | Pointer to **string** | Task status (triggers automatic timestamp updates) | [optional] 
**AssignedAgentId** | Pointer to **NullableString** | Reassign task to different agent | [optional] 
**DependsOn** | Pointer to **[]string** | Update task dependencies | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Update task metadata (replaces entire metadata object) | [optional] 
**Progress** | Pointer to **float32** | Progress from 0.0 to 1.0 | [optional] 
**ProgressMessage** | Pointer to **string** | Human-readable progress message | [optional] 
**Result** | Pointer to **map[string]interface{}** | Task result data (set when completing task) | [optional] 
**Error** | Pointer to **NullableString** | Error message (set when task fails) | [optional] 
**RetryCount** | Pointer to **int32** | Update retry count | [optional] 
**MaxRetries** | Pointer to **int32** | Update maximum retry attempts | [optional] 
**BlockedReason** | Pointer to **NullableString** | Reason task is blocked (set when status is blocked) | [optional] 
**BlockedByTaskIds** | Pointer to **[]string** | Task IDs that are blocking this task | [optional] 

## Methods

### NewUpdateTaskRequest

`func NewUpdateTaskRequest() *UpdateTaskRequest`

NewUpdateTaskRequest instantiates a new UpdateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskRequestWithDefaults

`func NewUpdateTaskRequestWithDefaults() *UpdateTaskRequest`

NewUpdateTaskRequestWithDefaults instantiates a new UpdateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateTaskRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateTaskRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateTaskRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateTaskRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaskListId

`func (o *UpdateTaskRequest) GetTaskListId() string`

GetTaskListId returns the TaskListId field if non-nil, zero value otherwise.

### GetTaskListIdOk

`func (o *UpdateTaskRequest) GetTaskListIdOk() (*string, bool)`

GetTaskListIdOk returns a tuple with the TaskListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskListId

`func (o *UpdateTaskRequest) SetTaskListId(v string)`

SetTaskListId sets TaskListId field to given value.

### HasTaskListId

`func (o *UpdateTaskRequest) HasTaskListId() bool`

HasTaskListId returns a boolean if a field has been set.

### SetTaskListIdNil

`func (o *UpdateTaskRequest) SetTaskListIdNil(b bool)`

 SetTaskListIdNil sets the value for TaskListId to be an explicit nil

### UnsetTaskListId
`func (o *UpdateTaskRequest) UnsetTaskListId()`

UnsetTaskListId ensures that no value is present for TaskListId, not even an explicit nil
### GetStatus

`func (o *UpdateTaskRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateTaskRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateTaskRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateTaskRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssignedAgentId

`func (o *UpdateTaskRequest) GetAssignedAgentId() string`

GetAssignedAgentId returns the AssignedAgentId field if non-nil, zero value otherwise.

### GetAssignedAgentIdOk

`func (o *UpdateTaskRequest) GetAssignedAgentIdOk() (*string, bool)`

GetAssignedAgentIdOk returns a tuple with the AssignedAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAgentId

`func (o *UpdateTaskRequest) SetAssignedAgentId(v string)`

SetAssignedAgentId sets AssignedAgentId field to given value.

### HasAssignedAgentId

`func (o *UpdateTaskRequest) HasAssignedAgentId() bool`

HasAssignedAgentId returns a boolean if a field has been set.

### SetAssignedAgentIdNil

`func (o *UpdateTaskRequest) SetAssignedAgentIdNil(b bool)`

 SetAssignedAgentIdNil sets the value for AssignedAgentId to be an explicit nil

### UnsetAssignedAgentId
`func (o *UpdateTaskRequest) UnsetAssignedAgentId()`

UnsetAssignedAgentId ensures that no value is present for AssignedAgentId, not even an explicit nil
### GetDependsOn

`func (o *UpdateTaskRequest) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *UpdateTaskRequest) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *UpdateTaskRequest) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *UpdateTaskRequest) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateTaskRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateTaskRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateTaskRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateTaskRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProgress

`func (o *UpdateTaskRequest) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *UpdateTaskRequest) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *UpdateTaskRequest) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *UpdateTaskRequest) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressMessage

`func (o *UpdateTaskRequest) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *UpdateTaskRequest) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *UpdateTaskRequest) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *UpdateTaskRequest) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### GetResult

`func (o *UpdateTaskRequest) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateTaskRequest) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateTaskRequest) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *UpdateTaskRequest) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *UpdateTaskRequest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpdateTaskRequest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpdateTaskRequest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UpdateTaskRequest) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *UpdateTaskRequest) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *UpdateTaskRequest) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetRetryCount

`func (o *UpdateTaskRequest) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *UpdateTaskRequest) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *UpdateTaskRequest) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *UpdateTaskRequest) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetMaxRetries

`func (o *UpdateTaskRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *UpdateTaskRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *UpdateTaskRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *UpdateTaskRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetBlockedReason

`func (o *UpdateTaskRequest) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *UpdateTaskRequest) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *UpdateTaskRequest) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *UpdateTaskRequest) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### SetBlockedReasonNil

`func (o *UpdateTaskRequest) SetBlockedReasonNil(b bool)`

 SetBlockedReasonNil sets the value for BlockedReason to be an explicit nil

### UnsetBlockedReason
`func (o *UpdateTaskRequest) UnsetBlockedReason()`

UnsetBlockedReason ensures that no value is present for BlockedReason, not even an explicit nil
### GetBlockedByTaskIds

`func (o *UpdateTaskRequest) GetBlockedByTaskIds() []string`

GetBlockedByTaskIds returns the BlockedByTaskIds field if non-nil, zero value otherwise.

### GetBlockedByTaskIdsOk

`func (o *UpdateTaskRequest) GetBlockedByTaskIdsOk() (*[]string, bool)`

GetBlockedByTaskIdsOk returns a tuple with the BlockedByTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedByTaskIds

`func (o *UpdateTaskRequest) SetBlockedByTaskIds(v []string)`

SetBlockedByTaskIds sets BlockedByTaskIds field to given value.

### HasBlockedByTaskIds

`func (o *UpdateTaskRequest) HasBlockedByTaskIds() bool`

HasBlockedByTaskIds returns a boolean if a field has been set.

### SetBlockedByTaskIdsNil

`func (o *UpdateTaskRequest) SetBlockedByTaskIdsNil(b bool)`

 SetBlockedByTaskIdsNil sets the value for BlockedByTaskIds to be an explicit nil

### UnsetBlockedByTaskIds
`func (o *UpdateTaskRequest) UnsetBlockedByTaskIds()`

UnsetBlockedByTaskIds ensures that no value is present for BlockedByTaskIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


