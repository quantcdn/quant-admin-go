# CiValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** | Lower-case owner/name. Present when scope is repo. | [optional] 
**Branches** | Pointer to **[]string** | The branch rule. Absent when the value applies to every branch. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Variables only. A secret value is never returned. | [optional] 

## Methods

### NewCiValue

`func NewCiValue() *CiValue`

NewCiValue instantiates a new CiValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiValueWithDefaults

`func NewCiValueWithDefaults() *CiValue`

NewCiValueWithDefaults instantiates a new CiValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CiValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CiValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CiValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CiValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *CiValue) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CiValue) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CiValue) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CiValue) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetScope

`func (o *CiValue) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CiValue) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CiValue) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CiValue) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRepo

`func (o *CiValue) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CiValue) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CiValue) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CiValue) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetBranches

`func (o *CiValue) GetBranches() []string`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *CiValue) GetBranchesOk() (*[]string, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *CiValue) SetBranches(v []string)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *CiValue) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CiValue) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CiValue) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CiValue) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CiValue) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CiValue) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CiValue) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CiValue) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CiValue) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetValue

`func (o *CiValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CiValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CiValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CiValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


