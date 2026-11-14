# AIAutoModeEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orchestrator** | [**NullableAIModelEligibility**](AIModelEligibility.md) |  | 
**Roles** | [**AIAutoModeEligibilityRoles**](AIAutoModeEligibilityRoles.md) |  | 
**Usable** | **bool** |  | 

## Methods

### NewAIAutoModeEligibility

`func NewAIAutoModeEligibility(orchestrator NullableAIModelEligibility, roles AIAutoModeEligibilityRoles, usable bool, ) *AIAutoModeEligibility`

NewAIAutoModeEligibility instantiates a new AIAutoModeEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAutoModeEligibilityWithDefaults

`func NewAIAutoModeEligibilityWithDefaults() *AIAutoModeEligibility`

NewAIAutoModeEligibilityWithDefaults instantiates a new AIAutoModeEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestrator

`func (o *AIAutoModeEligibility) GetOrchestrator() AIModelEligibility`

GetOrchestrator returns the Orchestrator field if non-nil, zero value otherwise.

### GetOrchestratorOk

`func (o *AIAutoModeEligibility) GetOrchestratorOk() (*AIModelEligibility, bool)`

GetOrchestratorOk returns a tuple with the Orchestrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrator

`func (o *AIAutoModeEligibility) SetOrchestrator(v AIModelEligibility)`

SetOrchestrator sets Orchestrator field to given value.


### SetOrchestratorNil

`func (o *AIAutoModeEligibility) SetOrchestratorNil(b bool)`

 SetOrchestratorNil sets the value for Orchestrator to be an explicit nil

### UnsetOrchestrator
`func (o *AIAutoModeEligibility) UnsetOrchestrator()`

UnsetOrchestrator ensures that no value is present for Orchestrator, not even an explicit nil
### GetRoles

`func (o *AIAutoModeEligibility) GetRoles() AIAutoModeEligibilityRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AIAutoModeEligibility) GetRolesOk() (*AIAutoModeEligibilityRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AIAutoModeEligibility) SetRoles(v AIAutoModeEligibilityRoles)`

SetRoles sets Roles field to given value.


### GetUsable

`func (o *AIAutoModeEligibility) GetUsable() bool`

GetUsable returns the Usable field if non-nil, zero value otherwise.

### GetUsableOk

`func (o *AIAutoModeEligibility) GetUsableOk() (*bool, bool)`

GetUsableOk returns a tuple with the Usable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsable

`func (o *AIAutoModeEligibility) SetUsable(v bool)`

SetUsable sets Usable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


