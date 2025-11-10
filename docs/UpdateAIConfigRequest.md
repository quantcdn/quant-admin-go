# UpdateAIConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledModels** | Pointer to **[]string** |  | [optional] 
**DefaultModel** | Pointer to **string** |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 

## Methods

### NewUpdateAIConfigRequest

`func NewUpdateAIConfigRequest() *UpdateAIConfigRequest`

NewUpdateAIConfigRequest instantiates a new UpdateAIConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAIConfigRequestWithDefaults

`func NewUpdateAIConfigRequestWithDefaults() *UpdateAIConfigRequest`

NewUpdateAIConfigRequestWithDefaults instantiates a new UpdateAIConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledModels

`func (o *UpdateAIConfigRequest) GetEnabledModels() []string`

GetEnabledModels returns the EnabledModels field if non-nil, zero value otherwise.

### GetEnabledModelsOk

`func (o *UpdateAIConfigRequest) GetEnabledModelsOk() (*[]string, bool)`

GetEnabledModelsOk returns a tuple with the EnabledModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledModels

`func (o *UpdateAIConfigRequest) SetEnabledModels(v []string)`

SetEnabledModels sets EnabledModels field to given value.

### HasEnabledModels

`func (o *UpdateAIConfigRequest) HasEnabledModels() bool`

HasEnabledModels returns a boolean if a field has been set.

### GetDefaultModel

`func (o *UpdateAIConfigRequest) GetDefaultModel() string`

GetDefaultModel returns the DefaultModel field if non-nil, zero value otherwise.

### GetDefaultModelOk

`func (o *UpdateAIConfigRequest) GetDefaultModelOk() (*string, bool)`

GetDefaultModelOk returns a tuple with the DefaultModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModel

`func (o *UpdateAIConfigRequest) SetDefaultModel(v string)`

SetDefaultModel sets DefaultModel field to given value.

### HasDefaultModel

`func (o *UpdateAIConfigRequest) HasDefaultModel() bool`

HasDefaultModel returns a boolean if a field has been set.

### GetMaxTokens

`func (o *UpdateAIConfigRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *UpdateAIConfigRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *UpdateAIConfigRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *UpdateAIConfigRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTemperature

`func (o *UpdateAIConfigRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *UpdateAIConfigRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *UpdateAIConfigRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *UpdateAIConfigRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


