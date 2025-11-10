# CreateApplicationRequestFilesystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to **bool** | Whether to create a shared filesystem | [optional] [default to true]
**MountPath** | Pointer to **string** | Mount path inside containers | [optional] [default to "/mnt/data"]

## Methods

### NewCreateApplicationRequestFilesystem

`func NewCreateApplicationRequestFilesystem() *CreateApplicationRequestFilesystem`

NewCreateApplicationRequestFilesystem instantiates a new CreateApplicationRequestFilesystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestFilesystemWithDefaults

`func NewCreateApplicationRequestFilesystemWithDefaults() *CreateApplicationRequestFilesystem`

NewCreateApplicationRequestFilesystemWithDefaults instantiates a new CreateApplicationRequestFilesystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *CreateApplicationRequestFilesystem) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CreateApplicationRequestFilesystem) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CreateApplicationRequestFilesystem) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CreateApplicationRequestFilesystem) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetMountPath

`func (o *CreateApplicationRequestFilesystem) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *CreateApplicationRequestFilesystem) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *CreateApplicationRequestFilesystem) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *CreateApplicationRequestFilesystem) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


