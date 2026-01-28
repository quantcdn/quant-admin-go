# GetAIOrchestrationStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrchestrationId** | **string** | Unique orchestration identifier | 
**Status** | **string** | Current orchestration status | 
**ToolCount** | **int32** | Total number of async tools in this orchestration | 
**CompletedTools** | Pointer to **int32** | Number of tools that have completed | [optional] 
**SynthesizedResponse** | Pointer to **string** | AI-synthesized response combining all tool results (only present when status&#x3D;complete) | [optional] 
**Tools** | Pointer to [**[]GetAIOrchestrationStatus200ResponseToolsInner**](GetAIOrchestrationStatus200ResponseToolsInner.md) | Status of individual tool executions | [optional] 
**Error** | Pointer to **string** | Error message (only present when status&#x3D;failed) | [optional] 
**CreatedAt** | **time.Time** | When orchestration was created | 
**CompletedAt** | Pointer to **time.Time** | When orchestration completed (if status in complete or failed) | [optional] 

## Methods

### NewGetAIOrchestrationStatus200Response

`func NewGetAIOrchestrationStatus200Response(orchestrationId string, status string, toolCount int32, createdAt time.Time, ) *GetAIOrchestrationStatus200Response`

NewGetAIOrchestrationStatus200Response instantiates a new GetAIOrchestrationStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAIOrchestrationStatus200ResponseWithDefaults

`func NewGetAIOrchestrationStatus200ResponseWithDefaults() *GetAIOrchestrationStatus200Response`

NewGetAIOrchestrationStatus200ResponseWithDefaults instantiates a new GetAIOrchestrationStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestrationId

`func (o *GetAIOrchestrationStatus200Response) GetOrchestrationId() string`

GetOrchestrationId returns the OrchestrationId field if non-nil, zero value otherwise.

### GetOrchestrationIdOk

`func (o *GetAIOrchestrationStatus200Response) GetOrchestrationIdOk() (*string, bool)`

GetOrchestrationIdOk returns a tuple with the OrchestrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationId

`func (o *GetAIOrchestrationStatus200Response) SetOrchestrationId(v string)`

SetOrchestrationId sets OrchestrationId field to given value.


### GetStatus

`func (o *GetAIOrchestrationStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAIOrchestrationStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAIOrchestrationStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetToolCount

`func (o *GetAIOrchestrationStatus200Response) GetToolCount() int32`

GetToolCount returns the ToolCount field if non-nil, zero value otherwise.

### GetToolCountOk

`func (o *GetAIOrchestrationStatus200Response) GetToolCountOk() (*int32, bool)`

GetToolCountOk returns a tuple with the ToolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCount

`func (o *GetAIOrchestrationStatus200Response) SetToolCount(v int32)`

SetToolCount sets ToolCount field to given value.


### GetCompletedTools

`func (o *GetAIOrchestrationStatus200Response) GetCompletedTools() int32`

GetCompletedTools returns the CompletedTools field if non-nil, zero value otherwise.

### GetCompletedToolsOk

`func (o *GetAIOrchestrationStatus200Response) GetCompletedToolsOk() (*int32, bool)`

GetCompletedToolsOk returns a tuple with the CompletedTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTools

`func (o *GetAIOrchestrationStatus200Response) SetCompletedTools(v int32)`

SetCompletedTools sets CompletedTools field to given value.

### HasCompletedTools

`func (o *GetAIOrchestrationStatus200Response) HasCompletedTools() bool`

HasCompletedTools returns a boolean if a field has been set.

### GetSynthesizedResponse

`func (o *GetAIOrchestrationStatus200Response) GetSynthesizedResponse() string`

GetSynthesizedResponse returns the SynthesizedResponse field if non-nil, zero value otherwise.

### GetSynthesizedResponseOk

`func (o *GetAIOrchestrationStatus200Response) GetSynthesizedResponseOk() (*string, bool)`

GetSynthesizedResponseOk returns a tuple with the SynthesizedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthesizedResponse

`func (o *GetAIOrchestrationStatus200Response) SetSynthesizedResponse(v string)`

SetSynthesizedResponse sets SynthesizedResponse field to given value.

### HasSynthesizedResponse

`func (o *GetAIOrchestrationStatus200Response) HasSynthesizedResponse() bool`

HasSynthesizedResponse returns a boolean if a field has been set.

### GetTools

`func (o *GetAIOrchestrationStatus200Response) GetTools() []GetAIOrchestrationStatus200ResponseToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *GetAIOrchestrationStatus200Response) GetToolsOk() (*[]GetAIOrchestrationStatus200ResponseToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *GetAIOrchestrationStatus200Response) SetTools(v []GetAIOrchestrationStatus200ResponseToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *GetAIOrchestrationStatus200Response) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetError

`func (o *GetAIOrchestrationStatus200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAIOrchestrationStatus200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAIOrchestrationStatus200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetAIOrchestrationStatus200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetAIOrchestrationStatus200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAIOrchestrationStatus200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAIOrchestrationStatus200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *GetAIOrchestrationStatus200Response) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetAIOrchestrationStatus200Response) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetAIOrchestrationStatus200Response) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetAIOrchestrationStatus200Response) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


