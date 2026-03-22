# TokensCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name for the token | 
**Preset** | Pointer to **NullableString** | Preset scope bundle | [optional] 
**Scopes** | Pointer to **[]string** | Individual scopes (mutually exclusive with preset) | [optional] 
**Projects** | Pointer to **[]int32** | Project IDs to restrict this token to | [optional] 
**ExpiresIn** | Pointer to **NullableString** | Token expiration period | [optional] 

## Methods

### NewTokensCreateRequest

`func NewTokensCreateRequest(name string, ) *TokensCreateRequest`

NewTokensCreateRequest instantiates a new TokensCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensCreateRequestWithDefaults

`func NewTokensCreateRequestWithDefaults() *TokensCreateRequest`

NewTokensCreateRequestWithDefaults instantiates a new TokensCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TokensCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokensCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokensCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPreset

`func (o *TokensCreateRequest) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *TokensCreateRequest) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *TokensCreateRequest) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *TokensCreateRequest) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### SetPresetNil

`func (o *TokensCreateRequest) SetPresetNil(b bool)`

 SetPresetNil sets the value for Preset to be an explicit nil

### UnsetPreset
`func (o *TokensCreateRequest) UnsetPreset()`

UnsetPreset ensures that no value is present for Preset, not even an explicit nil
### GetScopes

`func (o *TokensCreateRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokensCreateRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokensCreateRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokensCreateRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *TokensCreateRequest) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *TokensCreateRequest) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetProjects

`func (o *TokensCreateRequest) GetProjects() []int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *TokensCreateRequest) GetProjectsOk() (*[]int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *TokensCreateRequest) SetProjects(v []int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *TokensCreateRequest) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *TokensCreateRequest) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *TokensCreateRequest) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetExpiresIn

`func (o *TokensCreateRequest) GetExpiresIn() string`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TokensCreateRequest) GetExpiresInOk() (*string, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TokensCreateRequest) SetExpiresIn(v string)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *TokensCreateRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### SetExpiresInNil

`func (o *TokensCreateRequest) SetExpiresInNil(b bool)`

 SetExpiresInNil sets the value for ExpiresIn to be an explicit nil

### UnsetExpiresIn
`func (o *TokensCreateRequest) UnsetExpiresIn()`

UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


