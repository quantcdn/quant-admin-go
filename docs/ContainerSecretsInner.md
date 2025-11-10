# ContainerSecretsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The environment variable name to be set in the container | 
**ValueFrom** | **string** | The key of the secret in the environment&#39;s &#39;app-secrets&#39; store | 

## Methods

### NewContainerSecretsInner

`func NewContainerSecretsInner(name string, valueFrom string, ) *ContainerSecretsInner`

NewContainerSecretsInner instantiates a new ContainerSecretsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerSecretsInnerWithDefaults

`func NewContainerSecretsInnerWithDefaults() *ContainerSecretsInner`

NewContainerSecretsInnerWithDefaults instantiates a new ContainerSecretsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerSecretsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerSecretsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerSecretsInner) SetName(v string)`

SetName sets Name field to given value.


### GetValueFrom

`func (o *ContainerSecretsInner) GetValueFrom() string`

GetValueFrom returns the ValueFrom field if non-nil, zero value otherwise.

### GetValueFromOk

`func (o *ContainerSecretsInner) GetValueFromOk() (*string, bool)`

GetValueFromOk returns a tuple with the ValueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFrom

`func (o *ContainerSecretsInner) SetValueFrom(v string)`

SetValueFrom sets ValueFrom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


