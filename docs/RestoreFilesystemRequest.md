# RestoreFilesystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **string** | The backup ID to restore (must match path param) | 
**AcknowledgeDataloss** | **bool** | Must be true. tar extraction overwrites same-named files in the target EFS in place; pre-existing files not in the archive are preserved. | 

## Methods

### NewRestoreFilesystemRequest

`func NewRestoreFilesystemRequest(backupId string, acknowledgeDataloss bool, ) *RestoreFilesystemRequest`

NewRestoreFilesystemRequest instantiates a new RestoreFilesystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreFilesystemRequestWithDefaults

`func NewRestoreFilesystemRequestWithDefaults() *RestoreFilesystemRequest`

NewRestoreFilesystemRequestWithDefaults instantiates a new RestoreFilesystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *RestoreFilesystemRequest) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *RestoreFilesystemRequest) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *RestoreFilesystemRequest) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetAcknowledgeDataloss

`func (o *RestoreFilesystemRequest) GetAcknowledgeDataloss() bool`

GetAcknowledgeDataloss returns the AcknowledgeDataloss field if non-nil, zero value otherwise.

### GetAcknowledgeDatalossOk

`func (o *RestoreFilesystemRequest) GetAcknowledgeDatalossOk() (*bool, bool)`

GetAcknowledgeDatalossOk returns a tuple with the AcknowledgeDataloss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeDataloss

`func (o *RestoreFilesystemRequest) SetAcknowledgeDataloss(v bool)`

SetAcknowledgeDataloss sets AcknowledgeDataloss field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


