# GetGovernanceConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** |  | [optional] 
**AiEnabled** | Pointer to **bool** |  | [optional] 
**ModelPolicy** | Pointer to **string** |  | [optional] 
**ModelList** | Pointer to **[]string** |  | [optional] 
**MandatoryGuardrailPreset** | Pointer to **NullableString** |  | [optional] 
**MandatoryFilterPolicies** | Pointer to **[]string** |  | [optional] 
**SpendLimits** | Pointer to [**GetGovernanceConfig200ResponseSpendLimits**](GetGovernanceConfig200ResponseSpendLimits.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetGovernanceConfig200Response

`func NewGetGovernanceConfig200Response() *GetGovernanceConfig200Response`

NewGetGovernanceConfig200Response instantiates a new GetGovernanceConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGovernanceConfig200ResponseWithDefaults

`func NewGetGovernanceConfig200ResponseWithDefaults() *GetGovernanceConfig200Response`

NewGetGovernanceConfig200ResponseWithDefaults instantiates a new GetGovernanceConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *GetGovernanceConfig200Response) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GetGovernanceConfig200Response) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GetGovernanceConfig200Response) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *GetGovernanceConfig200Response) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetAiEnabled

`func (o *GetGovernanceConfig200Response) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *GetGovernanceConfig200Response) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *GetGovernanceConfig200Response) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.

### HasAiEnabled

`func (o *GetGovernanceConfig200Response) HasAiEnabled() bool`

HasAiEnabled returns a boolean if a field has been set.

### GetModelPolicy

`func (o *GetGovernanceConfig200Response) GetModelPolicy() string`

GetModelPolicy returns the ModelPolicy field if non-nil, zero value otherwise.

### GetModelPolicyOk

`func (o *GetGovernanceConfig200Response) GetModelPolicyOk() (*string, bool)`

GetModelPolicyOk returns a tuple with the ModelPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPolicy

`func (o *GetGovernanceConfig200Response) SetModelPolicy(v string)`

SetModelPolicy sets ModelPolicy field to given value.

### HasModelPolicy

`func (o *GetGovernanceConfig200Response) HasModelPolicy() bool`

HasModelPolicy returns a boolean if a field has been set.

### GetModelList

`func (o *GetGovernanceConfig200Response) GetModelList() []string`

GetModelList returns the ModelList field if non-nil, zero value otherwise.

### GetModelListOk

`func (o *GetGovernanceConfig200Response) GetModelListOk() (*[]string, bool)`

GetModelListOk returns a tuple with the ModelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelList

`func (o *GetGovernanceConfig200Response) SetModelList(v []string)`

SetModelList sets ModelList field to given value.

### HasModelList

`func (o *GetGovernanceConfig200Response) HasModelList() bool`

HasModelList returns a boolean if a field has been set.

### GetMandatoryGuardrailPreset

`func (o *GetGovernanceConfig200Response) GetMandatoryGuardrailPreset() string`

GetMandatoryGuardrailPreset returns the MandatoryGuardrailPreset field if non-nil, zero value otherwise.

### GetMandatoryGuardrailPresetOk

`func (o *GetGovernanceConfig200Response) GetMandatoryGuardrailPresetOk() (*string, bool)`

GetMandatoryGuardrailPresetOk returns a tuple with the MandatoryGuardrailPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryGuardrailPreset

`func (o *GetGovernanceConfig200Response) SetMandatoryGuardrailPreset(v string)`

SetMandatoryGuardrailPreset sets MandatoryGuardrailPreset field to given value.

### HasMandatoryGuardrailPreset

`func (o *GetGovernanceConfig200Response) HasMandatoryGuardrailPreset() bool`

HasMandatoryGuardrailPreset returns a boolean if a field has been set.

### SetMandatoryGuardrailPresetNil

`func (o *GetGovernanceConfig200Response) SetMandatoryGuardrailPresetNil(b bool)`

 SetMandatoryGuardrailPresetNil sets the value for MandatoryGuardrailPreset to be an explicit nil

### UnsetMandatoryGuardrailPreset
`func (o *GetGovernanceConfig200Response) UnsetMandatoryGuardrailPreset()`

UnsetMandatoryGuardrailPreset ensures that no value is present for MandatoryGuardrailPreset, not even an explicit nil
### GetMandatoryFilterPolicies

`func (o *GetGovernanceConfig200Response) GetMandatoryFilterPolicies() []string`

GetMandatoryFilterPolicies returns the MandatoryFilterPolicies field if non-nil, zero value otherwise.

### GetMandatoryFilterPoliciesOk

`func (o *GetGovernanceConfig200Response) GetMandatoryFilterPoliciesOk() (*[]string, bool)`

GetMandatoryFilterPoliciesOk returns a tuple with the MandatoryFilterPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryFilterPolicies

`func (o *GetGovernanceConfig200Response) SetMandatoryFilterPolicies(v []string)`

SetMandatoryFilterPolicies sets MandatoryFilterPolicies field to given value.

### HasMandatoryFilterPolicies

`func (o *GetGovernanceConfig200Response) HasMandatoryFilterPolicies() bool`

HasMandatoryFilterPolicies returns a boolean if a field has been set.

### GetSpendLimits

`func (o *GetGovernanceConfig200Response) GetSpendLimits() GetGovernanceConfig200ResponseSpendLimits`

GetSpendLimits returns the SpendLimits field if non-nil, zero value otherwise.

### GetSpendLimitsOk

`func (o *GetGovernanceConfig200Response) GetSpendLimitsOk() (*GetGovernanceConfig200ResponseSpendLimits, bool)`

GetSpendLimitsOk returns a tuple with the SpendLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimits

`func (o *GetGovernanceConfig200Response) SetSpendLimits(v GetGovernanceConfig200ResponseSpendLimits)`

SetSpendLimits sets SpendLimits field to given value.

### HasSpendLimits

`func (o *GetGovernanceConfig200Response) HasSpendLimits() bool`

HasSpendLimits returns a boolean if a field has been set.

### GetVersion

`func (o *GetGovernanceConfig200Response) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetGovernanceConfig200Response) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetGovernanceConfig200Response) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetGovernanceConfig200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


