# PutCiValueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Value** | **string** |  | 
**Repo** | Pointer to **string** | Owner/name. Lower-cased by the server. Leave out for an organisation value. | [optional] 
**Branches** | Pointer to **[]string** | Only jobs on these branches get the value. Leave out for every branch. | [optional] 

## Methods

### NewPutCiValueRequest

`func NewPutCiValueRequest(kind string, value string, ) *PutCiValueRequest`

NewPutCiValueRequest instantiates a new PutCiValueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutCiValueRequestWithDefaults

`func NewPutCiValueRequestWithDefaults() *PutCiValueRequest`

NewPutCiValueRequestWithDefaults instantiates a new PutCiValueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PutCiValueRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PutCiValueRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PutCiValueRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetValue

`func (o *PutCiValueRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PutCiValueRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PutCiValueRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetRepo

`func (o *PutCiValueRequest) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PutCiValueRequest) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PutCiValueRequest) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PutCiValueRequest) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetBranches

`func (o *PutCiValueRequest) GetBranches() []string`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *PutCiValueRequest) GetBranchesOk() (*[]string, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *PutCiValueRequest) SetBranches(v []string)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *PutCiValueRequest) HasBranches() bool`

HasBranches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


