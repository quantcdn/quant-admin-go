# ImportSkillCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**Source** | [**ImportSkillCollectionRequestSource**](ImportSkillCollectionRequestSource.md) |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**InstalledBy** | Pointer to **string** |  | [optional] 

## Methods

### NewImportSkillCollectionRequest

`func NewImportSkillCollectionRequest(namespace string, source ImportSkillCollectionRequestSource, ) *ImportSkillCollectionRequest`

NewImportSkillCollectionRequest instantiates a new ImportSkillCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportSkillCollectionRequestWithDefaults

`func NewImportSkillCollectionRequestWithDefaults() *ImportSkillCollectionRequest`

NewImportSkillCollectionRequestWithDefaults instantiates a new ImportSkillCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ImportSkillCollectionRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ImportSkillCollectionRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ImportSkillCollectionRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSource

`func (o *ImportSkillCollectionRequest) GetSource() ImportSkillCollectionRequestSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ImportSkillCollectionRequest) GetSourceOk() (*ImportSkillCollectionRequestSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ImportSkillCollectionRequest) SetSource(v ImportSkillCollectionRequestSource)`

SetSource sets Source field to given value.


### GetTags

`func (o *ImportSkillCollectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportSkillCollectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportSkillCollectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportSkillCollectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetInstalledBy

`func (o *ImportSkillCollectionRequest) GetInstalledBy() string`

GetInstalledBy returns the InstalledBy field if non-nil, zero value otherwise.

### GetInstalledByOk

`func (o *ImportSkillCollectionRequest) GetInstalledByOk() (*string, bool)`

GetInstalledByOk returns a tuple with the InstalledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledBy

`func (o *ImportSkillCollectionRequest) SetInstalledBy(v string)`

SetInstalledBy sets InstalledBy field to given value.

### HasInstalledBy

`func (o *ImportSkillCollectionRequest) HasInstalledBy() bool`

HasInstalledBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


