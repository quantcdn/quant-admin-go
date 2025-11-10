# ContainerImageReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Specifies whether the image is internal (ECR) or external (e.g., Docker Hub) | 
**Identifier** | **string** | The image identifier. For &#39;internal&#39; type, this is the image tag. For &#39;external&#39; type, this is the full image name. | 

## Methods

### NewContainerImageReference

`func NewContainerImageReference(type_ string, identifier string, ) *ContainerImageReference`

NewContainerImageReference instantiates a new ContainerImageReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageReferenceWithDefaults

`func NewContainerImageReferenceWithDefaults() *ContainerImageReference`

NewContainerImageReferenceWithDefaults instantiates a new ContainerImageReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContainerImageReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerImageReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerImageReference) SetType(v string)`

SetType sets Type field to given value.


### GetIdentifier

`func (o *ContainerImageReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ContainerImageReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ContainerImageReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


