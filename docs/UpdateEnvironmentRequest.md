# UpdateEnvironmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComposeDefinition** | [**Compose**](Compose.md) |  | 
**MinCapacity** | Pointer to **NullableInt32** | Optional. Minimum number of tasks for auto-scaling. If provided at root level, will be merged into composeDefinition. | [optional] 
**MaxCapacity** | Pointer to **NullableInt32** | Optional. Maximum number of tasks for auto-scaling. If provided at root level, will be merged into composeDefinition. | [optional] 

## Methods

### NewUpdateEnvironmentRequest

`func NewUpdateEnvironmentRequest(composeDefinition Compose, ) *UpdateEnvironmentRequest`

NewUpdateEnvironmentRequest instantiates a new UpdateEnvironmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEnvironmentRequestWithDefaults

`func NewUpdateEnvironmentRequestWithDefaults() *UpdateEnvironmentRequest`

NewUpdateEnvironmentRequestWithDefaults instantiates a new UpdateEnvironmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComposeDefinition

`func (o *UpdateEnvironmentRequest) GetComposeDefinition() Compose`

GetComposeDefinition returns the ComposeDefinition field if non-nil, zero value otherwise.

### GetComposeDefinitionOk

`func (o *UpdateEnvironmentRequest) GetComposeDefinitionOk() (*Compose, bool)`

GetComposeDefinitionOk returns a tuple with the ComposeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeDefinition

`func (o *UpdateEnvironmentRequest) SetComposeDefinition(v Compose)`

SetComposeDefinition sets ComposeDefinition field to given value.


### GetMinCapacity

`func (o *UpdateEnvironmentRequest) GetMinCapacity() int32`

GetMinCapacity returns the MinCapacity field if non-nil, zero value otherwise.

### GetMinCapacityOk

`func (o *UpdateEnvironmentRequest) GetMinCapacityOk() (*int32, bool)`

GetMinCapacityOk returns a tuple with the MinCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacity

`func (o *UpdateEnvironmentRequest) SetMinCapacity(v int32)`

SetMinCapacity sets MinCapacity field to given value.

### HasMinCapacity

`func (o *UpdateEnvironmentRequest) HasMinCapacity() bool`

HasMinCapacity returns a boolean if a field has been set.

### SetMinCapacityNil

`func (o *UpdateEnvironmentRequest) SetMinCapacityNil(b bool)`

 SetMinCapacityNil sets the value for MinCapacity to be an explicit nil

### UnsetMinCapacity
`func (o *UpdateEnvironmentRequest) UnsetMinCapacity()`

UnsetMinCapacity ensures that no value is present for MinCapacity, not even an explicit nil
### GetMaxCapacity

`func (o *UpdateEnvironmentRequest) GetMaxCapacity() int32`

GetMaxCapacity returns the MaxCapacity field if non-nil, zero value otherwise.

### GetMaxCapacityOk

`func (o *UpdateEnvironmentRequest) GetMaxCapacityOk() (*int32, bool)`

GetMaxCapacityOk returns a tuple with the MaxCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacity

`func (o *UpdateEnvironmentRequest) SetMaxCapacity(v int32)`

SetMaxCapacity sets MaxCapacity field to given value.

### HasMaxCapacity

`func (o *UpdateEnvironmentRequest) HasMaxCapacity() bool`

HasMaxCapacity returns a boolean if a field has been set.

### SetMaxCapacityNil

`func (o *UpdateEnvironmentRequest) SetMaxCapacityNil(b bool)`

 SetMaxCapacityNil sets the value for MaxCapacity to be an explicit nil

### UnsetMaxCapacity
`func (o *UpdateEnvironmentRequest) UnsetMaxCapacity()`

UnsetMaxCapacity ensures that no value is present for MaxCapacity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


