# ContainerDependsOnInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** | The name of the container this container depends on | 
**Condition** | Pointer to **string** | The condition to wait for on the dependency | [optional] 

## Methods

### NewContainerDependsOnInner

`func NewContainerDependsOnInner(containerName string, ) *ContainerDependsOnInner`

NewContainerDependsOnInner instantiates a new ContainerDependsOnInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerDependsOnInnerWithDefaults

`func NewContainerDependsOnInnerWithDefaults() *ContainerDependsOnInner`

NewContainerDependsOnInnerWithDefaults instantiates a new ContainerDependsOnInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *ContainerDependsOnInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ContainerDependsOnInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ContainerDependsOnInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetCondition

`func (o *ContainerDependsOnInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ContainerDependsOnInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ContainerDependsOnInner) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ContainerDependsOnInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


