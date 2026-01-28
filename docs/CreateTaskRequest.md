# CreateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Task title | 
**Description** | Pointer to **string** | Detailed task description | [optional] 
**TaskListId** | Pointer to **string** | Task list ID for grouping related tasks (implicit - lists are created automatically) | [optional] 
**Status** | Pointer to **string** | Initial task status | [optional] [default to "pending"]
**AssignedAgentId** | Pointer to **string** | Pre-assign task to specific agent | [optional] 
**CreatedByAgentId** | Pointer to **string** | Agent ID that created this task | [optional] 
**DependsOn** | Pointer to **[]string** | Task IDs that must complete before this task can start | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Flexible JSON metadata for task-specific data | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum retry attempts on failure | [optional] [default to 3]
**BlockedReason** | Pointer to **string** | Reason task is blocked (when status is blocked) | [optional] 
**BlockedByTaskIds** | Pointer to **[]string** | Task IDs that are blocking this task | [optional] 

## Methods

### NewCreateTaskRequest

`func NewCreateTaskRequest(title string, ) *CreateTaskRequest`

NewCreateTaskRequest instantiates a new CreateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskRequestWithDefaults

`func NewCreateTaskRequestWithDefaults() *CreateTaskRequest`

NewCreateTaskRequestWithDefaults instantiates a new CreateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateTaskRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateTaskRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateTaskRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CreateTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaskListId

`func (o *CreateTaskRequest) GetTaskListId() string`

GetTaskListId returns the TaskListId field if non-nil, zero value otherwise.

### GetTaskListIdOk

`func (o *CreateTaskRequest) GetTaskListIdOk() (*string, bool)`

GetTaskListIdOk returns a tuple with the TaskListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskListId

`func (o *CreateTaskRequest) SetTaskListId(v string)`

SetTaskListId sets TaskListId field to given value.

### HasTaskListId

`func (o *CreateTaskRequest) HasTaskListId() bool`

HasTaskListId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateTaskRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTaskRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTaskRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTaskRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssignedAgentId

`func (o *CreateTaskRequest) GetAssignedAgentId() string`

GetAssignedAgentId returns the AssignedAgentId field if non-nil, zero value otherwise.

### GetAssignedAgentIdOk

`func (o *CreateTaskRequest) GetAssignedAgentIdOk() (*string, bool)`

GetAssignedAgentIdOk returns a tuple with the AssignedAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAgentId

`func (o *CreateTaskRequest) SetAssignedAgentId(v string)`

SetAssignedAgentId sets AssignedAgentId field to given value.

### HasAssignedAgentId

`func (o *CreateTaskRequest) HasAssignedAgentId() bool`

HasAssignedAgentId returns a boolean if a field has been set.

### GetCreatedByAgentId

`func (o *CreateTaskRequest) GetCreatedByAgentId() string`

GetCreatedByAgentId returns the CreatedByAgentId field if non-nil, zero value otherwise.

### GetCreatedByAgentIdOk

`func (o *CreateTaskRequest) GetCreatedByAgentIdOk() (*string, bool)`

GetCreatedByAgentIdOk returns a tuple with the CreatedByAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgentId

`func (o *CreateTaskRequest) SetCreatedByAgentId(v string)`

SetCreatedByAgentId sets CreatedByAgentId field to given value.

### HasCreatedByAgentId

`func (o *CreateTaskRequest) HasCreatedByAgentId() bool`

HasCreatedByAgentId returns a boolean if a field has been set.

### GetDependsOn

`func (o *CreateTaskRequest) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *CreateTaskRequest) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *CreateTaskRequest) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *CreateTaskRequest) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateTaskRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateTaskRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateTaskRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateTaskRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMaxRetries

`func (o *CreateTaskRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *CreateTaskRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *CreateTaskRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *CreateTaskRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetBlockedReason

`func (o *CreateTaskRequest) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *CreateTaskRequest) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *CreateTaskRequest) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *CreateTaskRequest) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetBlockedByTaskIds

`func (o *CreateTaskRequest) GetBlockedByTaskIds() []string`

GetBlockedByTaskIds returns the BlockedByTaskIds field if non-nil, zero value otherwise.

### GetBlockedByTaskIdsOk

`func (o *CreateTaskRequest) GetBlockedByTaskIdsOk() (*[]string, bool)`

GetBlockedByTaskIdsOk returns a tuple with the BlockedByTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedByTaskIds

`func (o *CreateTaskRequest) SetBlockedByTaskIds(v []string)`

SetBlockedByTaskIds sets BlockedByTaskIds field to given value.

### HasBlockedByTaskIds

`func (o *CreateTaskRequest) HasBlockedByTaskIds() bool`

HasBlockedByTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


