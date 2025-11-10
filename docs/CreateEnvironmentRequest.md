# CreateEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvName** | **string** | Environment name (e.g., &#39;staging&#39;, &#39;development&#39;) | 
**MinCapacity** | Pointer to **int32** | Minimum number of instances | [optional] 
**MaxCapacity** | Pointer to **int32** | Maximum number of instances | [optional] 
**CloneConfigurationFrom** | Pointer to **string** | Clone configuration from an existing environment | [optional] 
**ComposeDefinition** | Pointer to [**Compose**](Compose.md) |  | [optional] 
**ImageSuffix** | Pointer to **string** | Optional image tag suffix for cloning | [optional] 
**SpotConfiguration** | Pointer to [**SpotConfiguration**](SpotConfiguration.md) |  | [optional] 

## Methods

### NewCreateEnvironmentRequest

`func NewCreateEnvironmentRequest(envName string, ) *CreateEnvironmentRequest`

NewCreateEnvironmentRequest instantiates a new CreateEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentRequestWithDefaults

`func NewCreateEnvironmentRequestWithDefaults() *CreateEnvironmentRequest`

NewCreateEnvironmentRequestWithDefaults instantiates a new CreateEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvName

`func (o *CreateEnvironmentRequest) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *CreateEnvironmentRequest) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *CreateEnvironmentRequest) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.


### GetMinCapacity

`func (o *CreateEnvironmentRequest) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *CreateEnvironmentRequest) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *CreateEnvironmentRequest) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *CreateEnvironmentRequest) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### GetMaxCapacity

`func (o *CreateEnvironmentRequest) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *CreateEnvironmentRequest) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *CreateEnvironmentRequest) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *CreateEnvironmentRequest) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### GetCloneConfigurationFrom

`func (o *CreateEnvironmentRequest) GetCloneConfigurationFrom() string`

GetCloneConfigurationFrom returns the CloneConfigurationFrom field if non-nil, zero value otherwise.

### GetCloneConfigurationFromOk

`func (o *CreateEnvironmentRequest) GetCloneConfigurationFromOk() (*string, bool)`

GetCloneConfigurationFromOk returns a tuple with the CloneConfigurationFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneConfigurationFrom

`func (o *CreateEnvironmentRequest) SetCloneConfigurationFrom(v string)`

SetCloneConfigurationFrom sets CloneConfigurationFrom field to given value.

### HasCloneConfigurationFrom

`func (o *CreateEnvironmentRequest) HasCloneConfigurationFrom() bool`

HasCloneConfigurationFrom returns a boolean if a field has been set.

### GetComposeDefinition

`func (o *CreateEnvironmentRequest) GetComposeDefinition() Compose`

GetComposeDefinition returns the ComposeDefinition field if non-nil, zero value otherwise.

### GetComposeDefinitionOk

`func (o *CreateEnvironmentRequest) GetComposeDefinitionOk() (*Compose, bool)`

GetComposeDefinitionOk returns a tuple with the ComposeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeDefinition

`func (o *CreateEnvironmentRequest) SetComposeDefinition(v Compose)`

SetComposeDefinition sets ComposeDefinition field to given value.

### HasComposeDefinition

`func (o *CreateEnvironmentRequest) HasComposeDefinition() bool`

HasComposeDefinition returns a boolean if a field has been set.

### GetImageSuffix

`func (o *CreateEnvironmentRequest) GetImageSuffix() string`

GetImageSuffix returns the ImageSuffix field if non-nil, zero value otherwise.

### GetImageSuffixOk

`func (o *CreateEnvironmentRequest) GetImageSuffixOk() (*string, bool)`

GetImageSuffixOk returns a tuple with the ImageSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSuffix

`func (o *CreateEnvironmentRequest) SetImageSuffix(v string)`

SetImageSuffix sets ImageSuffix field to given value.

### HasImageSuffix

`func (o *CreateEnvironmentRequest) HasImageSuffix() bool`

HasImageSuffix returns a boolean if a field has been set.

### GetSpotConfiguration

`func (o *CreateEnvironmentRequest) GetSpotConfiguration() SpotConfiguration`

GetSpotConfiguration returns the SpotConfiguration field if non-nil, zero value otherwise.

### GetSpotConfigurationOk

`func (o *CreateEnvironmentRequest) GetSpotConfigurationOk() (*SpotConfiguration, bool)`

GetSpotConfigurationOk returns a tuple with the SpotConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotConfiguration

`func (o *CreateEnvironmentRequest) SetSpotConfiguration(v SpotConfiguration)`

SetSpotConfiguration sets SpotConfiguration field to given value.

### HasSpotConfiguration

`func (o *CreateEnvironmentRequest) HasSpotConfiguration() bool`

HasSpotConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


