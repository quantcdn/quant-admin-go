# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvName** | **string** |  | 
**DesiredCount** | Pointer to **int32** |  | [optional] 
**MinCapacity** | Pointer to **int32** |  | [optional] 
**MaxCapacity** | Pointer to **int32** |  | [optional] 
**CloneConfigurationFrom** | Pointer to **string** |  | [optional] 
**ImageSuffix** | Pointer to **string** | Image tag suffix for cloning | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment(envName string, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvName

`func (o *Environment) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *Environment) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *Environment) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.


### GetDesiredCount

`func (o *Environment) GetDesiredCount() int32`

GetDesiredCount returns the DesiredCount field if non-nil, zero value otherwise.

### GetDesiredCountOk

`func (o *Environment) GetDesiredCountOk() (*int32, bool)`

GetDesiredCountOk returns a tuple with the DesiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCount

`func (o *Environment) SetDesiredCount(v int32)`

SetDesiredCount sets DesiredCount field to given value.

### HasDesiredCount

`func (o *Environment) HasDesiredCount() bool`

HasDesiredCount returns a boolean if a field has been set.

### GetMinCapacity

`func (o *Environment) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *Environment) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *Environment) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *Environment) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *Environment) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *Environment) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *Environment) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *Environment) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetCloneConfigurationFrom

`func (o *Environment) GetCloneConfigurationFrom() string`

GetCloneConfigurationFrom returns the CloneConfigurationFrom field if non-nil, zero value otherwise.

### GetCloneConfigurationFromOk

`func (o *Environment) GetCloneConfigurationFromOk() (*string, bool)`

GetCloneConfigurationFromOk returns a tuple with the CloneConfigurationFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneConfigurationFrom

`func (o *Environment) SetCloneConfigurationFrom(v string)`

SetCloneConfigurationFrom sets CloneConfigurationFrom field to given value.

### HasCloneConfigurationFrom

`func (o *Environment) HasCloneConfigurationFrom() bool`

HasCloneConfigurationFrom returns a boolean if a field has been set.

### GetImageSuffix

`func (o *Environment) GetImageSuffix() string`

GetImageSuffix returns the ImageSuffix field if non-nil, zero value otherwise.

### GetImageSuffixOk

`func (o *Environment) GetImageSuffixOk() (*string, bool)`

GetImageSuffixOk returns a tuple with the ImageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSuffix

`func (o *Environment) SetImageSuffix(v string)`

SetImageSuffix sets ImageSuffix field to given value.

### HasImageSuffix

`func (o *Environment) HasImageSuffix() bool`

HasImageSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


