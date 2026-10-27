# DecideToolProposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decision** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewDecideToolProposalRequest

`func NewDecideToolProposalRequest(decision string, ) *DecideToolProposalRequest`

NewDecideToolProposalRequest instantiates a new DecideToolProposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecideToolProposalRequestWithDefaults

`func NewDecideToolProposalRequestWithDefaults() *DecideToolProposalRequest`

NewDecideToolProposalRequestWithDefaults instantiates a new DecideToolProposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecision

`func (o *DecideToolProposalRequest) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *DecideToolProposalRequest) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *DecideToolProposalRequest) SetDecision(v string)`

SetDecision sets Decision field to given value.


### GetReason

`func (o *DecideToolProposalRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DecideToolProposalRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DecideToolProposalRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DecideToolProposalRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetScope

`func (o *DecideToolProposalRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DecideToolProposalRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DecideToolProposalRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DecideToolProposalRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


