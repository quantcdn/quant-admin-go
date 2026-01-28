# GetTask200Response

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
**CreatedByAgentId** | Pointer to **NullableString** |  | [optional] 
**DependsOn** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Progress** | Pointer to **float32** |  | [optional] 
**ProgressMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **map[string]interface{}** | Task result data when completed | [optional] 
**Error** | Pointer to **NullableString** | Error message if status is failed | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**MaxRetries** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**UpdatedAt** | Pointer to **int64** | Unix timestamp in milliseconds | [optional] 
**StartedAt** | Pointer to **NullableInt64** | When status changed to in_progress | [optional] 
**CompletedAt** | Pointer to **NullableInt64** | When task completed/failed/cancelled | [optional] 
**ExpiresAt** | Pointer to **NullableInt64** | TTL timestamp for completed tasks | [optional] 
**BlockedReason** | Pointer to **NullableString** | Reason task is blocked | [optional] 
**BlockedByTaskIds** | Pointer to **[]string** | Task IDs that are blocking this task | [optional] 
**BlockedAt** | Pointer to **NullableInt64** | When status changed to blocked | [optional] 

## Methods

### NewGetTask200Response

`func NewGetTask200Response() *GetTask200Response`

NewGetTask200Response instantiates a new GetTask200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTask200ResponseWithDefaults

`func NewGetTask200ResponseWithDefaults() *GetTask200Response`

NewGetTask200ResponseWithDefaults instantiates a new GetTask200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *GetTask200Response) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GetTask200Response) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GetTask200Response) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *GetTask200Response) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetOrgId

`func (o *GetTask200Response) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GetTask200Response) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GetTask200Response) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *GetTask200Response) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTaskListId

`func (o *GetTask200Response) GetTaskListId() string`

GetTaskListId returns the TaskListId field if non-nil, zero value otherwise.

### GetTaskListIdOk

`func (o *GetTask200Response) GetTaskListIdOk() (*string, bool)`

GetTaskListIdOk returns a tuple with the TaskListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskListId

`func (o *GetTask200Response) SetTaskListId(v string)`

SetTaskListId sets TaskListId field to given value.

### HasTaskListId

`func (o *GetTask200Response) HasTaskListId() bool`

HasTaskListId returns a boolean if a field has been set.

### SetTaskListIdNil

`func (o *GetTask200Response) SetTaskListIdNil(b bool)`

 SetTaskListIdNil sets the value for TaskListId to be an explicit nil

### UnsetTaskListId
`func (o *GetTask200Response) UnsetTaskListId()`

UnsetTaskListId ensures that no value is present for TaskListId, not even an explicit nil
### GetTitle

`func (o *GetTask200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetTask200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetTask200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetTask200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *GetTask200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTask200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTask200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTask200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *GetTask200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTask200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTask200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTask200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssignedAgentId

`func (o *GetTask200Response) GetAssignedAgentId() string`

GetAssignedAgentId returns the AssignedAgentId field if non-nil, zero value otherwise.

### GetAssignedAgentIdOk

`func (o *GetTask200Response) GetAssignedAgentIdOk() (*string, bool)`

GetAssignedAgentIdOk returns a tuple with the AssignedAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAgentId

`func (o *GetTask200Response) SetAssignedAgentId(v string)`

SetAssignedAgentId sets AssignedAgentId field to given value.

### HasAssignedAgentId

`func (o *GetTask200Response) HasAssignedAgentId() bool`

HasAssignedAgentId returns a boolean if a field has been set.

### SetAssignedAgentIdNil

`func (o *GetTask200Response) SetAssignedAgentIdNil(b bool)`

 SetAssignedAgentIdNil sets the value for AssignedAgentId to be an explicit nil

### UnsetAssignedAgentId
`func (o *GetTask200Response) UnsetAssignedAgentId()`

UnsetAssignedAgentId ensures that no value is present for AssignedAgentId, not even an explicit nil
### GetCreatedByAgentId

`func (o *GetTask200Response) GetCreatedByAgentId() string`

GetCreatedByAgentId returns the CreatedByAgentId field if non-nil, zero value otherwise.

### GetCreatedByAgentIdOk

`func (o *GetTask200Response) GetCreatedByAgentIdOk() (*string, bool)`

GetCreatedByAgentIdOk returns a tuple with the CreatedByAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAgentId

`func (o *GetTask200Response) SetCreatedByAgentId(v string)`

SetCreatedByAgentId sets CreatedByAgentId field to given value.

### HasCreatedByAgentId

`func (o *GetTask200Response) HasCreatedByAgentId() bool`

HasCreatedByAgentId returns a boolean if a field has been set.

### SetCreatedByAgentIdNil

`func (o *GetTask200Response) SetCreatedByAgentIdNil(b bool)`

 SetCreatedByAgentIdNil sets the value for CreatedByAgentId to be an explicit nil

### UnsetCreatedByAgentId
`func (o *GetTask200Response) UnsetCreatedByAgentId()`

UnsetCreatedByAgentId ensures that no value is present for CreatedByAgentId, not even an explicit nil
### GetDependsOn

`func (o *GetTask200Response) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *GetTask200Response) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *GetTask200Response) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *GetTask200Response) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetMetadata

`func (o *GetTask200Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetTask200Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetTask200Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetTask200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProgress

`func (o *GetTask200Response) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *GetTask200Response) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *GetTask200Response) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *GetTask200Response) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressMessage

`func (o *GetTask200Response) GetProgressMessage() string`

GetProgressMessage returns the ProgressMessage field if non-nil, zero value otherwise.

### GetProgressMessageOk

`func (o *GetTask200Response) GetProgressMessageOk() (*string, bool)`

GetProgressMessageOk returns a tuple with the ProgressMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMessage

`func (o *GetTask200Response) SetProgressMessage(v string)`

SetProgressMessage sets ProgressMessage field to given value.

### HasProgressMessage

`func (o *GetTask200Response) HasProgressMessage() bool`

HasProgressMessage returns a boolean if a field has been set.

### GetResult

`func (o *GetTask200Response) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetTask200Response) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetTask200Response) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *GetTask200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *GetTask200Response) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *GetTask200Response) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetError

`func (o *GetTask200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetTask200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetTask200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetTask200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *GetTask200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *GetTask200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetRetryCount

`func (o *GetTask200Response) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *GetTask200Response) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *GetTask200Response) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *GetTask200Response) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetMaxRetries

`func (o *GetTask200Response) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *GetTask200Response) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *GetTask200Response) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *GetTask200Response) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetTask200Response) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTask200Response) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTask200Response) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetTask200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetTask200Response) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetTask200Response) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetTask200Response) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetTask200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *GetTask200Response) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetTask200Response) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetTask200Response) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetTask200Response) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GetTask200Response) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GetTask200Response) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GetTask200Response) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetTask200Response) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetTask200Response) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetTask200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GetTask200Response) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GetTask200Response) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetExpiresAt

`func (o *GetTask200Response) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetTask200Response) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetTask200Response) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetTask200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GetTask200Response) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GetTask200Response) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetBlockedReason

`func (o *GetTask200Response) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *GetTask200Response) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *GetTask200Response) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *GetTask200Response) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### SetBlockedReasonNil

`func (o *GetTask200Response) SetBlockedReasonNil(b bool)`

 SetBlockedReasonNil sets the value for BlockedReason to be an explicit nil

### UnsetBlockedReason
`func (o *GetTask200Response) UnsetBlockedReason()`

UnsetBlockedReason ensures that no value is present for BlockedReason, not even an explicit nil
### GetBlockedByTaskIds

`func (o *GetTask200Response) GetBlockedByTaskIds() []string`

GetBlockedByTaskIds returns the BlockedByTaskIds field if non-nil, zero value otherwise.

### GetBlockedByTaskIdsOk

`func (o *GetTask200Response) GetBlockedByTaskIdsOk() (*[]string, bool)`

GetBlockedByTaskIdsOk returns a tuple with the BlockedByTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedByTaskIds

`func (o *GetTask200Response) SetBlockedByTaskIds(v []string)`

SetBlockedByTaskIds sets BlockedByTaskIds field to given value.

### HasBlockedByTaskIds

`func (o *GetTask200Response) HasBlockedByTaskIds() bool`

HasBlockedByTaskIds returns a boolean if a field has been set.

### SetBlockedByTaskIdsNil

`func (o *GetTask200Response) SetBlockedByTaskIdsNil(b bool)`

 SetBlockedByTaskIdsNil sets the value for BlockedByTaskIds to be an explicit nil

### UnsetBlockedByTaskIds
`func (o *GetTask200Response) UnsetBlockedByTaskIds()`

UnsetBlockedByTaskIds ensures that no value is present for BlockedByTaskIds, not even an explicit nil
### GetBlockedAt

`func (o *GetTask200Response) GetBlockedAt() int64`

GetBlockedAt returns the BlockedAt field if non-nil, zero value otherwise.

### GetBlockedAtOk

`func (o *GetTask200Response) GetBlockedAtOk() (*int64, bool)`

GetBlockedAtOk returns a tuple with the BlockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedAt

`func (o *GetTask200Response) SetBlockedAt(v int64)`

SetBlockedAt sets BlockedAt field to given value.

### HasBlockedAt

`func (o *GetTask200Response) HasBlockedAt() bool`

HasBlockedAt returns a boolean if a field has been set.

### SetBlockedAtNil

`func (o *GetTask200Response) SetBlockedAtNil(b bool)`

 SetBlockedAtNil sets the value for BlockedAt to be an explicit nil

### UnsetBlockedAt
`func (o *GetTask200Response) UnsetBlockedAt()`

UnsetBlockedAt ensures that no value is present for BlockedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


