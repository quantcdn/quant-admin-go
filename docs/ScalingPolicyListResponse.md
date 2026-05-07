# ScalingPolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]GetScalingPolicyResponse**](GetScalingPolicyResponse.md) |  | [optional] 

## Methods

### NewScalingPolicyListResponse

`func NewScalingPolicyListResponse() *ScalingPolicyListResponse`

NewScalingPolicyListResponse instantiates a new ScalingPolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalingPolicyListResponseWithDefaults

`func NewScalingPolicyListResponseWithDefaults() *ScalingPolicyListResponse`

NewScalingPolicyListResponseWithDefaults instantiates a new ScalingPolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *ScalingPolicyListResponse) GetPolicies() []GetScalingPolicyResponse`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ScalingPolicyListResponse) GetPoliciesOk() (*[]GetScalingPolicyResponse, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ScalingPolicyListResponse) SetPolicies(v []GetScalingPolicyResponse)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ScalingPolicyListResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


