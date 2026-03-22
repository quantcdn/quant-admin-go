# TokensCreate201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The plain text token (shown once) | [optional] 
**Id** | Pointer to **int32** | Token ID | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Projects** | Pointer to **[]int32** |  | [optional] 
**Preset** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTokensCreate201Response

`func NewTokensCreate201Response() *TokensCreate201Response`

NewTokensCreate201Response instantiates a new TokensCreate201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensCreate201ResponseWithDefaults

`func NewTokensCreate201ResponseWithDefaults() *TokensCreate201Response`

NewTokensCreate201ResponseWithDefaults instantiates a new TokensCreate201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokensCreate201Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokensCreate201Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokensCreate201Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokensCreate201Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetId

`func (o *TokensCreate201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokensCreate201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokensCreate201Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TokensCreate201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TokensCreate201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokensCreate201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokensCreate201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokensCreate201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *TokensCreate201Response) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokensCreate201Response) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokensCreate201Response) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokensCreate201Response) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *TokensCreate201Response) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *TokensCreate201Response) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetProjects

`func (o *TokensCreate201Response) GetProjects() []int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *TokensCreate201Response) GetProjectsOk() (*[]int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *TokensCreate201Response) SetProjects(v []int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *TokensCreate201Response) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *TokensCreate201Response) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *TokensCreate201Response) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetPreset

`func (o *TokensCreate201Response) GetPreset() string`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *TokensCreate201Response) GetPresetOk() (*string, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *TokensCreate201Response) SetPreset(v string)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *TokensCreate201Response) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### SetPresetNil

`func (o *TokensCreate201Response) SetPresetNil(b bool)`

 SetPresetNil sets the value for Preset to be an explicit nil

### UnsetPreset
`func (o *TokensCreate201Response) UnsetPreset()`

UnsetPreset ensures that no value is present for Preset, not even an explicit nil
### GetExpiresAt

`func (o *TokensCreate201Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokensCreate201Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokensCreate201Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokensCreate201Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *TokensCreate201Response) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *TokensCreate201Response) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetCreatedAt

`func (o *TokensCreate201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TokensCreate201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TokensCreate201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TokensCreate201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


