# ListTasks200ResponseTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**TaskListId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AssignedAgentId** | Pointer to **NullableString** |  | [optional] 
**Progress** | Pointer to **float32** |  | [optional] 
**BlockedReason** | Pointer to **NullableString** |  | [optional] 
**BlockedByTaskIds** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewListTasks200ResponseTasksInner

`func NewListTasks200ResponseTasksInner() *ListTasks200ResponseTasksInner`

NewListTasks200ResponseTasksInner instantiates a new ListTasks200ResponseTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasks200ResponseTasksInnerWithDefaults

`func NewListTasks200ResponseTasksInnerWithDefaults() *ListTasks200ResponseTasksInner`

NewListTasks200ResponseTasksInnerWithDefaults instantiates a new ListTasks200ResponseTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *ListTasks200ResponseTasksInner) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ListTasks200ResponseTasksInner) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ListTasks200ResponseTasksInner) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ListTasks200ResponseTasksInner) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetOrgId

`func (o *ListTasks200ResponseTasksInner) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ListTasks200ResponseTasksInner) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ListTasks200ResponseTasksInner) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ListTasks200ResponseTasksInner) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTaskListId

`func (o *ListTasks200ResponseTasksInner) GetTaskListId() string`

GetTaskListId returns the TaskListId field if non-nil, zero value otherwise.

### GetTaskListIdOk

`func (o *ListTasks200ResponseTasksInner) GetTaskListIdOk() (*string, bool)`

GetTaskListIdOk returns a tuple with the TaskListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskListId

`func (o *ListTasks200ResponseTasksInner) SetTaskListId(v string)`

SetTaskListId sets TaskListId field to given value.

### HasTaskListId

`func (o *ListTasks200ResponseTasksInner) HasTaskListId() bool`

HasTaskListId returns a boolean if a field has been set.

### SetTaskListIdNil

`func (o *ListTasks200ResponseTasksInner) SetTaskListIdNil(b bool)`

 SetTaskListIdNil sets the value for TaskListId to be an explicit nil

### UnsetTaskListId
`func (o *ListTasks200ResponseTasksInner) UnsetTaskListId()`

UnsetTaskListId ensures that no value is present for TaskListId, not even an explicit nil
### GetTitle

`func (o *ListTasks200ResponseTasksInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListTasks200ResponseTasksInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListTasks200ResponseTasksInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListTasks200ResponseTasksInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ListTasks200ResponseTasksInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListTasks200ResponseTasksInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListTasks200ResponseTasksInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListTasks200ResponseTasksInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ListTasks200ResponseTasksInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListTasks200ResponseTasksInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListTasks200ResponseTasksInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListTasks200ResponseTasksInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssignedAgentId

`func (o *ListTasks200ResponseTasksInner) GetAssignedAgentId() string`

GetAssignedAgentId returns the AssignedAgentId field if non-nil, zero value otherwise.

### GetAssignedAgentIdOk

`func (o *ListTasks200ResponseTasksInner) GetAssignedAgentIdOk() (*string, bool)`

GetAssignedAgentIdOk returns a tuple with the AssignedAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAgentId

`func (o *ListTasks200ResponseTasksInner) SetAssignedAgentId(v string)`

SetAssignedAgentId sets AssignedAgentId field to given value.

### HasAssignedAgentId

`func (o *ListTasks200ResponseTasksInner) HasAssignedAgentId() bool`

HasAssignedAgentId returns a boolean if a field has been set.

### SetAssignedAgentIdNil

`func (o *ListTasks200ResponseTasksInner) SetAssignedAgentIdNil(b bool)`

 SetAssignedAgentIdNil sets the value for AssignedAgentId to be an explicit nil

### UnsetAssignedAgentId
`func (o *ListTasks200ResponseTasksInner) UnsetAssignedAgentId()`

UnsetAssignedAgentId ensures that no value is present for AssignedAgentId, not even an explicit nil
### GetProgress

`func (o *ListTasks200ResponseTasksInner) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ListTasks200ResponseTasksInner) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ListTasks200ResponseTasksInner) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ListTasks200ResponseTasksInner) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetBlockedReason

`func (o *ListTasks200ResponseTasksInner) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *ListTasks200ResponseTasksInner) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *ListTasks200ResponseTasksInner) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *ListTasks200ResponseTasksInner) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### SetBlockedReasonNil

`func (o *ListTasks200ResponseTasksInner) SetBlockedReasonNil(b bool)`

 SetBlockedReasonNil sets the value for BlockedReason to be an explicit nil

### UnsetBlockedReason
`func (o *ListTasks200ResponseTasksInner) UnsetBlockedReason()`

UnsetBlockedReason ensures that no value is present for BlockedReason, not even an explicit nil
### GetBlockedByTaskIds

`func (o *ListTasks200ResponseTasksInner) GetBlockedByTaskIds() []string`

GetBlockedByTaskIds returns the BlockedByTaskIds field if non-nil, zero value otherwise.

### GetBlockedByTaskIdsOk

`func (o *ListTasks200ResponseTasksInner) GetBlockedByTaskIdsOk() (*[]string, bool)`

GetBlockedByTaskIdsOk returns a tuple with the BlockedByTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedByTaskIds

`func (o *ListTasks200ResponseTasksInner) SetBlockedByTaskIds(v []string)`

SetBlockedByTaskIds sets BlockedByTaskIds field to given value.

### HasBlockedByTaskIds

`func (o *ListTasks200ResponseTasksInner) HasBlockedByTaskIds() bool`

HasBlockedByTaskIds returns a boolean if a field has been set.

### SetBlockedByTaskIdsNil

`func (o *ListTasks200ResponseTasksInner) SetBlockedByTaskIdsNil(b bool)`

 SetBlockedByTaskIdsNil sets the value for BlockedByTaskIds to be an explicit nil

### UnsetBlockedByTaskIds
`func (o *ListTasks200ResponseTasksInner) UnsetBlockedByTaskIds()`

UnsetBlockedByTaskIds ensures that no value is present for BlockedByTaskIds, not even an explicit nil
### GetCreatedAt

`func (o *ListTasks200ResponseTasksInner) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListTasks200ResponseTasksInner) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListTasks200ResponseTasksInner) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListTasks200ResponseTasksInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListTasks200ResponseTasksInner) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListTasks200ResponseTasksInner) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListTasks200ResponseTasksInner) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListTasks200ResponseTasksInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


