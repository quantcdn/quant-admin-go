# CreateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeName** | **string** | Volume name | 
**Description** | Pointer to **NullableString** | Volume description | [optional] 
**RootDirectory** | Pointer to **NullableString** | Root directory path | [optional] 

## Methods

### NewCreateVolumeRequest

`func NewCreateVolumeRequest(volumeName string, ) *CreateVolumeRequest`

NewCreateVolumeRequest instantiates a new CreateVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeRequestWithDefaults

`func NewCreateVolumeRequestWithDefaults() *CreateVolumeRequest`

NewCreateVolumeRequestWithDefaults instantiates a new CreateVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeName

`func (o *CreateVolumeRequest) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *CreateVolumeRequest) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *CreateVolumeRequest) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.


### GetDescription

`func (o *CreateVolumeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVolumeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVolumeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVolumeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateVolumeRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateVolumeRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRootDirectory

`func (o *CreateVolumeRequest) GetRootDirectory() string`

GetRootDirectory returns the RootDirectory field if non-nil, zero value otherwise.

### GetRootDirectoryOk

`func (o *CreateVolumeRequest) GetRootDirectoryOk() (*string, bool)`

GetRootDirectoryOk returns a tuple with the RootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDirectory

`func (o *CreateVolumeRequest) SetRootDirectory(v string)`

SetRootDirectory sets RootDirectory field to given value.

### HasRootDirectory

`func (o *CreateVolumeRequest) HasRootDirectory() bool`

HasRootDirectory returns a boolean if a field has been set.

### SetRootDirectoryNil

`func (o *CreateVolumeRequest) SetRootDirectoryNil(b bool)`

 SetRootDirectoryNil sets the value for RootDirectory to be an explicit nil

### UnsetRootDirectory
`func (o *CreateVolumeRequest) UnsetRootDirectory()`

UnsetRootDirectory ensures that no value is present for RootDirectory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


