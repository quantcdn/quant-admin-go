# CreateCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**ContainerName** | Pointer to **NullableString** | Target container (defaults to the first container in the task definition) | [optional] 

## Methods

### NewCreateCommandRequest

`func NewCreateCommandRequest(command string, ) *CreateCommandRequest`

NewCreateCommandRequest instantiates a new CreateCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommandRequestWithDefaults

`func NewCreateCommandRequestWithDefaults() *CreateCommandRequest`

NewCreateCommandRequestWithDefaults instantiates a new CreateCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *CreateCommandRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateCommandRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateCommandRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetContainerName

`func (o *CreateCommandRequest) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *CreateCommandRequest) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *CreateCommandRequest) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *CreateCommandRequest) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### SetContainerNameNil

`func (o *CreateCommandRequest) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *CreateCommandRequest) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


