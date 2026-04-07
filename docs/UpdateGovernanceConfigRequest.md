# UpdateGovernanceConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiEnabled** | **bool** |  | 
**ModelPolicy** | **string** |  | 
**ModelList** | Pointer to **[]string** |  | [optional] 
**MandatoryGuardrailPreset** | Pointer to **NullableString** |  | [optional] 
**MandatoryFilterPolicies** | Pointer to **[]string** |  | [optional] 
**SpendLimits** | Pointer to [**GetGovernanceConfig200ResponseSpendLimits**](GetGovernanceConfig200ResponseSpendLimits.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateGovernanceConfigRequest

`func NewUpdateGovernanceConfigRequest(aiEnabled bool, modelPolicy string, ) *UpdateGovernanceConfigRequest`

NewUpdateGovernanceConfigRequest instantiates a new UpdateGovernanceConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGovernanceConfigRequestWithDefaults

`func NewUpdateGovernanceConfigRequestWithDefaults() *UpdateGovernanceConfigRequest`

NewUpdateGovernanceConfigRequestWithDefaults instantiates a new UpdateGovernanceConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiEnabled

`func (o *UpdateGovernanceConfigRequest) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *UpdateGovernanceConfigRequest) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *UpdateGovernanceConfigRequest) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.


### GetModelPolicy

`func (o *UpdateGovernanceConfigRequest) GetModelPolicy() string`

GetModelPolicy returns the ModelPolicy field if non-nil, zero value otherwise.

### GetModelPolicyOk

`func (o *UpdateGovernanceConfigRequest) GetModelPolicyOk() (*string, bool)`

GetModelPolicyOk returns a tuple with the ModelPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPolicy

`func (o *UpdateGovernanceConfigRequest) SetModelPolicy(v string)`

SetModelPolicy sets ModelPolicy field to given value.


### GetModelList

`func (o *UpdateGovernanceConfigRequest) GetModelList() []string`

GetModelList returns the ModelList field if non-nil, zero value otherwise.

### GetModelListOk

`func (o *UpdateGovernanceConfigRequest) GetModelListOk() (*[]string, bool)`

GetModelListOk returns a tuple with the ModelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelList

`func (o *UpdateGovernanceConfigRequest) SetModelList(v []string)`

SetModelList sets ModelList field to given value.

### HasModelList

`func (o *UpdateGovernanceConfigRequest) HasModelList() bool`

HasModelList returns a boolean if a field has been set.

### GetMandatoryGuardrailPreset

`func (o *UpdateGovernanceConfigRequest) GetMandatoryGuardrailPreset() string`

GetMandatoryGuardrailPreset returns the MandatoryGuardrailPreset field if non-nil, zero value otherwise.

### GetMandatoryGuardrailPresetOk

`func (o *UpdateGovernanceConfigRequest) GetMandatoryGuardrailPresetOk() (*string, bool)`

GetMandatoryGuardrailPresetOk returns a tuple with the MandatoryGuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryGuardrailPreset

`func (o *UpdateGovernanceConfigRequest) SetMandatoryGuardrailPreset(v string)`

SetMandatoryGuardrailPreset sets MandatoryGuardrailPreset field to given value.

### HasMandatoryGuardrailPreset

`func (o *UpdateGovernanceConfigRequest) HasMandatoryGuardrailPreset() bool`

HasMandatoryGuardrailPreset returns a boolean if a field has been set.

### SetMandatoryGuardrailPresetNil

`func (o *UpdateGovernanceConfigRequest) SetMandatoryGuardrailPresetNil(b bool)`

 SetMandatoryGuardrailPresetNil sets the value for MandatoryGuardrailPreset to be an explicit nil

### UnsetMandatoryGuardrailPreset
`func (o *UpdateGovernanceConfigRequest) UnsetMandatoryGuardrailPreset()`

UnsetMandatoryGuardrailPreset ensures that no value is present for MandatoryGuardrailPreset, not even an explicit nil
### GetMandatoryFilterPolicies

`func (o *UpdateGovernanceConfigRequest) GetMandatoryFilterPolicies() []string`

GetMandatoryFilterPolicies returns the MandatoryFilterPolicies field if non-nil, zero value otherwise.

### GetMandatoryFilterPoliciesOk

`func (o *UpdateGovernanceConfigRequest) GetMandatoryFilterPoliciesOk() (*[]string, bool)`

GetMandatoryFilterPoliciesOk returns a tuple with the MandatoryFilterPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryFilterPolicies

`func (o *UpdateGovernanceConfigRequest) SetMandatoryFilterPolicies(v []string)`

SetMandatoryFilterPolicies sets MandatoryFilterPolicies field to given value.

### HasMandatoryFilterPolicies

`func (o *UpdateGovernanceConfigRequest) HasMandatoryFilterPolicies() bool`

HasMandatoryFilterPolicies returns a boolean if a field has been set.

### GetSpendLimits

`func (o *UpdateGovernanceConfigRequest) GetSpendLimits() GetGovernanceConfig200ResponseSpendLimits`

GetSpendLimits returns the SpendLimits field if non-nil, zero value otherwise.

### GetSpendLimitsOk

`func (o *UpdateGovernanceConfigRequest) GetSpendLimitsOk() (*GetGovernanceConfig200ResponseSpendLimits, bool)`

GetSpendLimitsOk returns a tuple with the SpendLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimits

`func (o *UpdateGovernanceConfigRequest) SetSpendLimits(v GetGovernanceConfig200ResponseSpendLimits)`

SetSpendLimits sets SpendLimits field to given value.

### HasSpendLimits

`func (o *UpdateGovernanceConfigRequest) HasSpendLimits() bool`

HasSpendLimits returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateGovernanceConfigRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateGovernanceConfigRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateGovernanceConfigRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateGovernanceConfigRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


