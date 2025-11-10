# ContainerMountPointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVolume** | **string** | The name of the logical volume | 
**ContainerPath** | **string** | The path inside the container where the volume is mounted | 
**ReadOnly** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewContainerMountPointsInner

`func NewContainerMountPointsInner(sourceVolume string, containerPath string, ) *ContainerMountPointsInner`

NewContainerMountPointsInner instantiates a new ContainerMountPointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerMountPointsInnerWithDefaults

`func NewContainerMountPointsInnerWithDefaults() *ContainerMountPointsInner`

NewContainerMountPointsInnerWithDefaults instantiates a new ContainerMountPointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceVolume

`func (o *ContainerMountPointsInner) GetSourceVolume() string`

GetSourceVolume returns the SourceVolume field if non-nil, zero value otherwise.

### GetSourceVolumeOk

`func (o *ContainerMountPointsInner) GetSourceVolumeOk() (*string, bool)`

GetSourceVolumeOk returns a tuple with the SourceVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolume

`func (o *ContainerMountPointsInner) SetSourceVolume(v string)`

SetSourceVolume sets SourceVolume field to given value.


### GetContainerPath

`func (o *ContainerMountPointsInner) GetContainerPath() string`

GetContainerPath returns the ContainerPath field if non-nil, zero value otherwise.

### GetContainerPathOk

`func (o *ContainerMountPointsInner) GetContainerPathOk() (*string, bool)`

GetContainerPathOk returns a tuple with the ContainerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPath

`func (o *ContainerMountPointsInner) SetContainerPath(v string)`

SetContainerPath sets ContainerPath field to given value.


### GetReadOnly

`func (o *ContainerMountPointsInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ContainerMountPointsInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ContainerMountPointsInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ContainerMountPointsInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


