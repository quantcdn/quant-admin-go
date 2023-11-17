# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**MachineName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**OrganisationId** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ProjectType** | Pointer to **string** |  | [optional] 
**ProjectParentId** | Pointer to **int32** |  | [optional] 
**SecurityScore** | Pointer to **int32** |  | [optional] 
**GitUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMachineName

`func (o *Project) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *Project) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *Project) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *Project) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Project) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Project) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Project) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Project) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Project) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Project) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Project) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Project) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Project) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Project) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Project) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Project) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetOrganisationId

`func (o *Project) GetOrganisationId() int64`

GetOrganisationId returns the OrganisationId field if non-nil, zero value otherwise.

### GetOrganisationIdOk

`func (o *Project) GetOrganisationIdOk() (*int64, bool)`

GetOrganisationIdOk returns a tuple with the OrganisationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationId

`func (o *Project) SetOrganisationId(v int64)`

SetOrganisationId sets OrganisationId field to given value.

### HasOrganisationId

`func (o *Project) HasOrganisationId() bool`

HasOrganisationId returns a boolean if a field has been set.

### GetUuid

`func (o *Project) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Project) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Project) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Project) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetRegion

`func (o *Project) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Project) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Project) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Project) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetProjectType

`func (o *Project) GetProjectType() string`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *Project) GetProjectTypeOk() (*string, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *Project) SetProjectType(v string)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *Project) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### GetProjectParentId

`func (o *Project) GetProjectParentId() int32`

GetProjectParentId returns the ProjectParentId field if non-nil, zero value otherwise.

### GetProjectParentIdOk

`func (o *Project) GetProjectParentIdOk() (*int32, bool)`

GetProjectParentIdOk returns a tuple with the ProjectParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectParentId

`func (o *Project) SetProjectParentId(v int32)`

SetProjectParentId sets ProjectParentId field to given value.

### HasProjectParentId

`func (o *Project) HasProjectParentId() bool`

HasProjectParentId returns a boolean if a field has been set.

### GetSecurityScore

`func (o *Project) GetSecurityScore() int32`

GetSecurityScore returns the SecurityScore field if non-nil, zero value otherwise.

### GetSecurityScoreOk

`func (o *Project) GetSecurityScoreOk() (*int32, bool)`

GetSecurityScoreOk returns a tuple with the SecurityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScore

`func (o *Project) SetSecurityScore(v int32)`

SetSecurityScore sets SecurityScore field to given value.

### HasSecurityScore

`func (o *Project) HasSecurityScore() bool`

HasSecurityScore returns a boolean if a field has been set.

### GetGitUrl

`func (o *Project) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *Project) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *Project) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *Project) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

