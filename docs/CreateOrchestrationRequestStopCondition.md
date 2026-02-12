# CreateOrchestrationRequestStopCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "all_complete"]
**MaxIterations** | Pointer to **int32** | Max iterations (for type&#x3D;max_iterations) | [optional] 
**ConditionPrompt** | Pointer to **string** | AI prompt to evaluate stop (for type&#x3D;condition) | [optional] 

## Methods

### NewCreateOrchestrationRequestStopCondition

`func NewCreateOrchestrationRequestStopCondition() *CreateOrchestrationRequestStopCondition`

NewCreateOrchestrationRequestStopCondition instantiates a new CreateOrchestrationRequestStopCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrchestrationRequestStopConditionWithDefaults

`func NewCreateOrchestrationRequestStopConditionWithDefaults() *CreateOrchestrationRequestStopCondition`

NewCreateOrchestrationRequestStopConditionWithDefaults instantiates a new CreateOrchestrationRequestStopCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOrchestrationRequestStopCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrchestrationRequestStopCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrchestrationRequestStopCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOrchestrationRequestStopCondition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMaxIterations

`func (o *CreateOrchestrationRequestStopCondition) GetMaxIterations() int32`

GetMaxIterations returns the MaxIterations field if non-nil, zero value otherwise.

### GetMaxIterationsOk

`func (o *CreateOrchestrationRequestStopCondition) GetMaxIterationsOk() (*int32, bool)`

GetMaxIterationsOk returns a tuple with the MaxIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIterations

`func (o *CreateOrchestrationRequestStopCondition) SetMaxIterations(v int32)`

SetMaxIterations sets MaxIterations field to given value.

### HasMaxIterations

`func (o *CreateOrchestrationRequestStopCondition) HasMaxIterations() bool`

HasMaxIterations returns a boolean if a field has been set.

### GetConditionPrompt

`func (o *CreateOrchestrationRequestStopCondition) GetConditionPrompt() string`

GetConditionPrompt returns the ConditionPrompt field if non-nil, zero value otherwise.

### GetConditionPromptOk

`func (o *CreateOrchestrationRequestStopCondition) GetConditionPromptOk() (*string, bool)`

GetConditionPromptOk returns a tuple with the ConditionPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionPrompt

`func (o *CreateOrchestrationRequestStopCondition) SetConditionPrompt(v string)`

SetConditionPrompt sets ConditionPrompt field to given value.

### HasConditionPrompt

`func (o *CreateOrchestrationRequestStopCondition) HasConditionPrompt() bool`

HasConditionPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


