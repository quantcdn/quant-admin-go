# CreateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | Application name | 
**ComposeDefinition** | [**Compose**](Compose.md) |  | 
**MinCapacity** | Pointer to **NullableInt32** | Minimum task count for auto-scaling | [optional] [default to 1]
**MaxCapacity** | Pointer to **NullableInt32** | Maximum task count for auto-scaling | [optional] [default to 1]
**Database** | Pointer to [**NullableCreateApplicationRequestDatabase**](CreateApplicationRequestDatabase.md) |  | [optional] 
**Filesystem** | Pointer to [**NullableCreateApplicationRequestFilesystem**](CreateApplicationRequestFilesystem.md) |  | [optional] 
**Environment** | Pointer to [**[]CreateApplicationRequestEnvironmentInner**](CreateApplicationRequestEnvironmentInner.md) | Optional. Initial environment variables for the production environment. | [optional] 

## Methods

### NewCreateApplicationRequest

`func NewCreateApplicationRequest(appName string, composeDefinition Compose, ) *CreateApplicationRequest`

NewCreateApplicationRequest instantiates a new CreateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestWithDefaults

`func NewCreateApplicationRequestWithDefaults() *CreateApplicationRequest`

NewCreateApplicationRequestWithDefaults instantiates a new CreateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *CreateApplicationRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *CreateApplicationRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *CreateApplicationRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetComposeDefinition

`func (o *CreateApplicationRequest) GetComposeDefinition() Compose`

GetComposeDefinition returns the ComposeDefinition field if non-nil, zero value otherwise.

### GetComposeDefinitionOk

`func (o *CreateApplicationRequest) GetComposeDefinitionOk() (*Compose, bool)`

GetComposeDefinitionOk returns a tuple with the ComposeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeDefinition

`func (o *CreateApplicationRequest) SetComposeDefinition(v Compose)`

SetComposeDefinition sets ComposeDefinition field to given value.


### GetMinCapacity

`func (o *CreateApplicationRequest) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *CreateApplicationRequest) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *CreateApplicationRequest) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *CreateApplicationRequest) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### SetMinCapacityNil

`func (o *CreateApplicationRequest) SetMinCapacityNil(b bool)`

 SetMinCapacityNil sets the value for MinCapacity to be an explicit nil

### UnsetMinCapacity
`func (o *CreateApplicationRequest) UnsetMinCapacity()`

UnsetMinCapacity ensures that no value is present for MinCapacity, not even an explicit nil
### GetMaxCapacity

`func (o *CreateApplicationRequest) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *CreateApplicationRequest) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *CreateApplicationRequest) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *CreateApplicationRequest) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### SetMaxCapacityNil

`func (o *CreateApplicationRequest) SetMaxCapacityNil(b bool)`

 SetMaxCapacityNil sets the value for MaxCapacity to be an explicit nil

### UnsetMaxCapacity
`func (o *CreateApplicationRequest) UnsetMaxCapacity()`

UnsetMaxCapacity ensures that no value is present for MaxCapacity, not even an explicit nil
### GetDatabase

`func (o *CreateApplicationRequest) GetDatabase() CreateApplicationRequestDatabase`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *CreateApplicationRequest) GetDatabaseOk() (*CreateApplicationRequestDatabase, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *CreateApplicationRequest) SetDatabase(v CreateApplicationRequestDatabase)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *CreateApplicationRequest) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### SetDatabaseNil

`func (o *CreateApplicationRequest) SetDatabaseNil(b bool)`

 SetDatabaseNil sets the value for Database to be an explicit nil

### UnsetDatabase
`func (o *CreateApplicationRequest) UnsetDatabase()`

UnsetDatabase ensures that no value is present for Database, not even an explicit nil
### GetFilesystem

`func (o *CreateApplicationRequest) GetFilesystem() CreateApplicationRequestFilesystem`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *CreateApplicationRequest) GetFilesystemOk() (*CreateApplicationRequestFilesystem, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *CreateApplicationRequest) SetFilesystem(v CreateApplicationRequestFilesystem)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *CreateApplicationRequest) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### SetFilesystemNil

`func (o *CreateApplicationRequest) SetFilesystemNil(b bool)`

 SetFilesystemNil sets the value for Filesystem to be an explicit nil

### UnsetFilesystem
`func (o *CreateApplicationRequest) UnsetFilesystem()`

UnsetFilesystem ensures that no value is present for Filesystem, not even an explicit nil
### GetEnvironment

`func (o *CreateApplicationRequest) GetEnvironment() []CreateApplicationRequestEnvironmentInner`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateApplicationRequest) GetEnvironmentOk() (*[]CreateApplicationRequestEnvironmentInner, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateApplicationRequest) SetEnvironment(v []CreateApplicationRequestEnvironmentInner)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateApplicationRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *CreateApplicationRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CreateApplicationRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


