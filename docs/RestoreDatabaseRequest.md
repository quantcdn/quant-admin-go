# RestoreDatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **string** | The backup ID to restore (must match path param) | 
**AcknowledgeDataloss** | **bool** | Must be true to confirm existing data will be overwritten | 

## Methods

### NewRestoreDatabaseRequest

`func NewRestoreDatabaseRequest(backupId string, acknowledgeDataloss bool, ) *RestoreDatabaseRequest`

NewRestoreDatabaseRequest instantiates a new RestoreDatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreDatabaseRequestWithDefaults

`func NewRestoreDatabaseRequestWithDefaults() *RestoreDatabaseRequest`

NewRestoreDatabaseRequestWithDefaults instantiates a new RestoreDatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *RestoreDatabaseRequest) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *RestoreDatabaseRequest) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *RestoreDatabaseRequest) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetAcknowledgeDataloss

`func (o *RestoreDatabaseRequest) GetAcknowledgeDataloss() bool`

GetAcknowledgeDataloss returns the AcknowledgeDataloss field if non-nil, zero value otherwise.

### GetAcknowledgeDatalossOk

`func (o *RestoreDatabaseRequest) GetAcknowledgeDatalossOk() (*bool, bool)`

GetAcknowledgeDatalossOk returns a tuple with the AcknowledgeDataloss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeDataloss

`func (o *RestoreDatabaseRequest) SetAcknowledgeDataloss(v bool)`

SetAcknowledgeDataloss sets AcknowledgeDataloss field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


