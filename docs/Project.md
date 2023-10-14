# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**MachineName** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ProjectConfig**](ProjectConfig.md) |  | [optional] 
**Domains** | Pointer to [**[]ProjectDomains**](ProjectDomains.md) |  | [optional] 

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

### GetToken

`func (o *Project) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Project) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Project) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Project) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetConfig

`func (o *Project) GetConfig() ProjectConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Project) GetConfigOk() (*ProjectConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Project) SetConfig(v ProjectConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Project) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDomains

`func (o *Project) GetDomains() []ProjectDomains`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Project) GetDomainsOk() (*[]ProjectDomains, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Project) SetDomains(v []ProjectDomains)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Project) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


