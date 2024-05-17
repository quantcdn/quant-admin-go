# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MachineName** | **string** |  | 
**Name** | **string** |  | 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ProjectType** | Pointer to **string** |  | [optional] [default to "normal"]
**GitUrl** | Pointer to **string** |  | [optional] 
**SecurityScore** | Pointer to **string** |  | [optional] 
**ParentProjectId** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**FastlyMigrated** | Pointer to **int32** |  | [optional] [default to 1]
**AllowQueryParams** | Pointer to **bool** |  | [optional] 

## Methods

### NewProject

`func NewProject(machineName string, name string, ) *Project`

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

`func (o *Project) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int32)`

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


### GetOrganizationId

`func (o *Project) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Project) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Project) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Project) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

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

### GetSecurityScore

`func (o *Project) GetSecurityScore() string`

GetSecurityScore returns the SecurityScore field if non-nil, zero value otherwise.

### GetSecurityScoreOk

`func (o *Project) GetSecurityScoreOk() (*string, bool)`

GetSecurityScoreOk returns a tuple with the SecurityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScore

`func (o *Project) SetSecurityScore(v string)`

SetSecurityScore sets SecurityScore field to given value.

### HasSecurityScore

`func (o *Project) HasSecurityScore() bool`

HasSecurityScore returns a boolean if a field has been set.

### GetParentProjectId

`func (o *Project) GetParentProjectId() int32`

GetParentProjectId returns the ParentProjectId field if non-nil, zero value otherwise.

### GetParentProjectIdOk

`func (o *Project) GetParentProjectIdOk() (*int32, bool)`

GetParentProjectIdOk returns a tuple with the ParentProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProjectId

`func (o *Project) SetParentProjectId(v int32)`

SetParentProjectId sets ParentProjectId field to given value.

### HasParentProjectId

`func (o *Project) HasParentProjectId() bool`

HasParentProjectId returns a boolean if a field has been set.

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

### GetFastlyMigrated

`func (o *Project) GetFastlyMigrated() int32`

GetFastlyMigrated returns the FastlyMigrated field if non-nil, zero value otherwise.

### GetFastlyMigratedOk

`func (o *Project) GetFastlyMigratedOk() (*int32, bool)`

GetFastlyMigratedOk returns a tuple with the FastlyMigrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastlyMigrated

`func (o *Project) SetFastlyMigrated(v int32)`

SetFastlyMigrated sets FastlyMigrated field to given value.

### HasFastlyMigrated

`func (o *Project) HasFastlyMigrated() bool`

HasFastlyMigrated returns a boolean if a field has been set.

### GetAllowQueryParams

`func (o *Project) GetAllowQueryParams() bool`

GetAllowQueryParams returns the AllowQueryParams field if non-nil, zero value otherwise.

### GetAllowQueryParamsOk

`func (o *Project) GetAllowQueryParamsOk() (*bool, bool)`

GetAllowQueryParamsOk returns a tuple with the AllowQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryParams

`func (o *Project) SetAllowQueryParams(v bool)`

SetAllowQueryParams sets AllowQueryParams field to given value.

### HasAllowQueryParams

`func (o *Project) HasAllowQueryParams() bool`

HasAllowQueryParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


