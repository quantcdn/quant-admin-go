# ApplicationFilesystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesystemId** | Pointer to **string** | EFS filesystem ID | [optional] 
**MountPath** | Pointer to **string** | Default mount path in containers | [optional] 

## Methods

### NewApplicationFilesystem

`func NewApplicationFilesystem() *ApplicationFilesystem`

NewApplicationFilesystem instantiates a new ApplicationFilesystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFilesystemWithDefaults

`func NewApplicationFilesystemWithDefaults() *ApplicationFilesystem`

NewApplicationFilesystemWithDefaults instantiates a new ApplicationFilesystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystemId

`func (o *ApplicationFilesystem) GetFilesystemId() string`

GetFilesystemId returns the FilesystemId field if non-nil, zero value otherwise.

### GetFilesystemIdOk

`func (o *ApplicationFilesystem) GetFilesystemIdOk() (*string, bool)`

GetFilesystemIdOk returns a tuple with the FilesystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemId

`func (o *ApplicationFilesystem) SetFilesystemId(v string)`

SetFilesystemId sets FilesystemId field to given value.

### HasFilesystemId

`func (o *ApplicationFilesystem) HasFilesystemId() bool`

HasFilesystemId returns a boolean if a field has been set.

### GetMountPath

`func (o *ApplicationFilesystem) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *ApplicationFilesystem) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *ApplicationFilesystem) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *ApplicationFilesystem) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


