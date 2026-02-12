# CreateOrchestrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Orchestration name | 
**Description** | Pointer to **NullableString** | Optional description | [optional] 
**AgentId** | Pointer to **NullableString** | Agent to process items | [optional] 
**ToolId** | Pointer to **NullableString** | Tool to execute for items | [optional] 
**WorkflowId** | Pointer to **NullableString** | Workflow to run for items | [optional] 
**InputSource** | [**CreateOrchestrationRequestInputSource**](CreateOrchestrationRequestInputSource.md) |  | 
**BatchSize** | Pointer to **int32** | Items per batch | [optional] [default to 10]
**Concurrency** | Pointer to **int32** | Concurrent items within a batch | [optional] [default to 1]
**StopCondition** | Pointer to [**CreateOrchestrationRequestStopCondition**](CreateOrchestrationRequestStopCondition.md) |  | [optional] 
**AssignedSkills** | Pointer to **[]string** | Skill IDs to assign | [optional] 
**AutoStart** | Pointer to **bool** | Whether to start immediately | [optional] [default to true]

## Methods

### NewCreateOrchestrationRequest

`func NewCreateOrchestrationRequest(name string, inputSource CreateOrchestrationRequestInputSource, ) *CreateOrchestrationRequest`

NewCreateOrchestrationRequest instantiates a new CreateOrchestrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrchestrationRequestWithDefaults

`func NewCreateOrchestrationRequestWithDefaults() *CreateOrchestrationRequest`

NewCreateOrchestrationRequestWithDefaults instantiates a new CreateOrchestrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrchestrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrchestrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrchestrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateOrchestrationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrchestrationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrchestrationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrchestrationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateOrchestrationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateOrchestrationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAgentId

`func (o *CreateOrchestrationRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CreateOrchestrationRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CreateOrchestrationRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CreateOrchestrationRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *CreateOrchestrationRequest) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *CreateOrchestrationRequest) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetToolId

`func (o *CreateOrchestrationRequest) GetToolId() string`

GetToolId returns the ToolId field if non-nil, zero value otherwise.

### GetToolIdOk

`func (o *CreateOrchestrationRequest) GetToolIdOk() (*string, bool)`

GetToolIdOk returns a tuple with the ToolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolId

`func (o *CreateOrchestrationRequest) SetToolId(v string)`

SetToolId sets ToolId field to given value.

### HasToolId

`func (o *CreateOrchestrationRequest) HasToolId() bool`

HasToolId returns a boolean if a field has been set.

### SetToolIdNil

`func (o *CreateOrchestrationRequest) SetToolIdNil(b bool)`

 SetToolIdNil sets the value for ToolId to be an explicit nil

### UnsetToolId
`func (o *CreateOrchestrationRequest) UnsetToolId()`

UnsetToolId ensures that no value is present for ToolId, not even an explicit nil
### GetWorkflowId

`func (o *CreateOrchestrationRequest) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CreateOrchestrationRequest) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CreateOrchestrationRequest) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *CreateOrchestrationRequest) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *CreateOrchestrationRequest) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *CreateOrchestrationRequest) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetInputSource

`func (o *CreateOrchestrationRequest) GetInputSource() CreateOrchestrationRequestInputSource`

GetInputSource returns the InputSource field if non-nil, zero value otherwise.

### GetInputSourceOk

`func (o *CreateOrchestrationRequest) GetInputSourceOk() (*CreateOrchestrationRequestInputSource, bool)`

GetInputSourceOk returns a tuple with the InputSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSource

`func (o *CreateOrchestrationRequest) SetInputSource(v CreateOrchestrationRequestInputSource)`

SetInputSource sets InputSource field to given value.


### GetBatchSize

`func (o *CreateOrchestrationRequest) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *CreateOrchestrationRequest) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *CreateOrchestrationRequest) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *CreateOrchestrationRequest) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetConcurrency

`func (o *CreateOrchestrationRequest) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *CreateOrchestrationRequest) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *CreateOrchestrationRequest) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *CreateOrchestrationRequest) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetStopCondition

`func (o *CreateOrchestrationRequest) GetStopCondition() CreateOrchestrationRequestStopCondition`

GetStopCondition returns the StopCondition field if non-nil, zero value otherwise.

### GetStopConditionOk

`func (o *CreateOrchestrationRequest) GetStopConditionOk() (*CreateOrchestrationRequestStopCondition, bool)`

GetStopConditionOk returns a tuple with the StopCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopCondition

`func (o *CreateOrchestrationRequest) SetStopCondition(v CreateOrchestrationRequestStopCondition)`

SetStopCondition sets StopCondition field to given value.

### HasStopCondition

`func (o *CreateOrchestrationRequest) HasStopCondition() bool`

HasStopCondition returns a boolean if a field has been set.

### GetAssignedSkills

`func (o *CreateOrchestrationRequest) GetAssignedSkills() []string`

GetAssignedSkills returns the AssignedSkills field if non-nil, zero value otherwise.

### GetAssignedSkillsOk

`func (o *CreateOrchestrationRequest) GetAssignedSkillsOk() (*[]string, bool)`

GetAssignedSkillsOk returns a tuple with the AssignedSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedSkills

`func (o *CreateOrchestrationRequest) SetAssignedSkills(v []string)`

SetAssignedSkills sets AssignedSkills field to given value.

### HasAssignedSkills

`func (o *CreateOrchestrationRequest) HasAssignedSkills() bool`

HasAssignedSkills returns a boolean if a field has been set.

### GetAutoStart

`func (o *CreateOrchestrationRequest) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *CreateOrchestrationRequest) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *CreateOrchestrationRequest) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *CreateOrchestrationRequest) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


