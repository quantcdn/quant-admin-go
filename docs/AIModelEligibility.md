# AIModelEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | **string** |  | 
**ResolvedId** | **string** |  | 
**Eligible** | **bool** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewAIModelEligibility

`func NewAIModelEligibility(modelId string, resolvedId string, eligible bool, ) *AIModelEligibility`

NewAIModelEligibility instantiates a new AIModelEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIModelEligibilityWithDefaults

`func NewAIModelEligibilityWithDefaults() *AIModelEligibility`

NewAIModelEligibilityWithDefaults instantiates a new AIModelEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *AIModelEligibility) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *AIModelEligibility) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *AIModelEligibility) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetResolvedId

`func (o *AIModelEligibility) GetResolvedId() string`

GetResolvedId returns the ResolvedId field if non-nil, zero value otherwise.

### GetResolvedIdOk

`func (o *AIModelEligibility) GetResolvedIdOk() (*string, bool)`

GetResolvedIdOk returns a tuple with the ResolvedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedId

`func (o *AIModelEligibility) SetResolvedId(v string)`

SetResolvedId sets ResolvedId field to given value.


### GetEligible

`func (o *AIModelEligibility) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *AIModelEligibility) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *AIModelEligibility) SetEligible(v bool)`

SetEligible sets Eligible field to given value.


### GetReason

`func (o *AIModelEligibility) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AIModelEligibility) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AIModelEligibility) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AIModelEligibility) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


