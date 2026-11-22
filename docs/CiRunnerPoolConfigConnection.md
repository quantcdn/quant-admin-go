# CiRunnerPoolConfigConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**InstallationId** | Pointer to **int64** |  | [optional] 
**GithubOrg** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **int64** | The GitHub App that owns the installation. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCiRunnerPoolConfigConnection

`func NewCiRunnerPoolConfigConnection() *CiRunnerPoolConfigConnection`

NewCiRunnerPoolConfigConnection instantiates a new CiRunnerPoolConfigConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiRunnerPoolConfigConnectionWithDefaults

`func NewCiRunnerPoolConfigConnectionWithDefaults() *CiRunnerPoolConfigConnection`

NewCiRunnerPoolConfigConnectionWithDefaults instantiates a new CiRunnerPoolConfigConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CiRunnerPoolConfigConnection) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CiRunnerPoolConfigConnection) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CiRunnerPoolConfigConnection) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CiRunnerPoolConfigConnection) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetInstallationId

`func (o *CiRunnerPoolConfigConnection) GetInstallationId() int64`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *CiRunnerPoolConfigConnection) GetInstallationIdOk() (*int64, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *CiRunnerPoolConfigConnection) SetInstallationId(v int64)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *CiRunnerPoolConfigConnection) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetGithubOrg

`func (o *CiRunnerPoolConfigConnection) GetGithubOrg() string`

GetGithubOrg returns the GithubOrg field if non-nil, zero value otherwise.

### GetGithubOrgOk

`func (o *CiRunnerPoolConfigConnection) GetGithubOrgOk() (*string, bool)`

GetGithubOrgOk returns a tuple with the GithubOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubOrg

`func (o *CiRunnerPoolConfigConnection) SetGithubOrg(v string)`

SetGithubOrg sets GithubOrg field to given value.

### HasGithubOrg

`func (o *CiRunnerPoolConfigConnection) HasGithubOrg() bool`

HasGithubOrg returns a boolean if a field has been set.

### GetAppId

`func (o *CiRunnerPoolConfigConnection) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CiRunnerPoolConfigConnection) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CiRunnerPoolConfigConnection) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CiRunnerPoolConfigConnection) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetStatus

`func (o *CiRunnerPoolConfigConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CiRunnerPoolConfigConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CiRunnerPoolConfigConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CiRunnerPoolConfigConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


