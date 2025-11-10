# ListBackups200ResponseBackupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Engine** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**TaskArn** | Pointer to **string** |  | [optional] 
**S3Key** | Pointer to **string** |  | [optional] 
**BucketName** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**SizeFormatted** | Pointer to **string** |  | [optional] 
**FileExists** | Pointer to **bool** |  | [optional] 

## Methods

### NewListBackups200ResponseBackupsInner

`func NewListBackups200ResponseBackupsInner() *ListBackups200ResponseBackupsInner`

NewListBackups200ResponseBackupsInner instantiates a new ListBackups200ResponseBackupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackups200ResponseBackupsInnerWithDefaults

`func NewListBackups200ResponseBackupsInnerWithDefaults() *ListBackups200ResponseBackupsInner`

NewListBackups200ResponseBackupsInnerWithDefaults instantiates a new ListBackups200ResponseBackupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *ListBackups200ResponseBackupsInner) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *ListBackups200ResponseBackupsInner) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *ListBackups200ResponseBackupsInner) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *ListBackups200ResponseBackupsInner) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetStatus

`func (o *ListBackups200ResponseBackupsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBackups200ResponseBackupsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBackups200ResponseBackupsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBackups200ResponseBackupsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ListBackups200ResponseBackupsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListBackups200ResponseBackupsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListBackups200ResponseBackupsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListBackups200ResponseBackupsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEngine

`func (o *ListBackups200ResponseBackupsInner) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ListBackups200ResponseBackupsInner) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ListBackups200ResponseBackupsInner) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *ListBackups200ResponseBackupsInner) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetDescription

`func (o *ListBackups200ResponseBackupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListBackups200ResponseBackupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListBackups200ResponseBackupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListBackups200ResponseBackupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListBackups200ResponseBackupsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListBackups200ResponseBackupsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListBackups200ResponseBackupsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListBackups200ResponseBackupsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListBackups200ResponseBackupsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListBackups200ResponseBackupsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListBackups200ResponseBackupsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListBackups200ResponseBackupsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTaskArn

`func (o *ListBackups200ResponseBackupsInner) GetTaskArn() string`

GetTaskArn returns the TaskArn field if non-nil, zero value otherwise.

### GetTaskArnOk

`func (o *ListBackups200ResponseBackupsInner) GetTaskArnOk() (*string, bool)`

GetTaskArnOk returns a tuple with the TaskArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArn

`func (o *ListBackups200ResponseBackupsInner) SetTaskArn(v string)`

SetTaskArn sets TaskArn field to given value.

### HasTaskArn

`func (o *ListBackups200ResponseBackupsInner) HasTaskArn() bool`

HasTaskArn returns a boolean if a field has been set.

### GetS3Key

`func (o *ListBackups200ResponseBackupsInner) GetS3Key() string`

GetS3Key returns the S3Key field if non-nil, zero value otherwise.

### GetS3KeyOk

`func (o *ListBackups200ResponseBackupsInner) GetS3KeyOk() (*string, bool)`

GetS3KeyOk returns a tuple with the S3Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Key

`func (o *ListBackups200ResponseBackupsInner) SetS3Key(v string)`

SetS3Key sets S3Key field to given value.

### HasS3Key

`func (o *ListBackups200ResponseBackupsInner) HasS3Key() bool`

HasS3Key returns a boolean if a field has been set.

### GetBucketName

`func (o *ListBackups200ResponseBackupsInner) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ListBackups200ResponseBackupsInner) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ListBackups200ResponseBackupsInner) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ListBackups200ResponseBackupsInner) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetSize

`func (o *ListBackups200ResponseBackupsInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListBackups200ResponseBackupsInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListBackups200ResponseBackupsInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListBackups200ResponseBackupsInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeFormatted

`func (o *ListBackups200ResponseBackupsInner) GetSizeFormatted() string`

GetSizeFormatted returns the SizeFormatted field if non-nil, zero value otherwise.

### GetSizeFormattedOk

`func (o *ListBackups200ResponseBackupsInner) GetSizeFormattedOk() (*string, bool)`

GetSizeFormattedOk returns a tuple with the SizeFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeFormatted

`func (o *ListBackups200ResponseBackupsInner) SetSizeFormatted(v string)`

SetSizeFormatted sets SizeFormatted field to given value.

### HasSizeFormatted

`func (o *ListBackups200ResponseBackupsInner) HasSizeFormatted() bool`

HasSizeFormatted returns a boolean if a field has been set.

### GetFileExists

`func (o *ListBackups200ResponseBackupsInner) GetFileExists() bool`

GetFileExists returns the FileExists field if non-nil, zero value otherwise.

### GetFileExistsOk

`func (o *ListBackups200ResponseBackupsInner) GetFileExistsOk() (*bool, bool)`

GetFileExistsOk returns a tuple with the FileExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExists

`func (o *ListBackups200ResponseBackupsInner) SetFileExists(v bool)`

SetFileExists sets FileExists field to given value.

### HasFileExists

`func (o *ListBackups200ResponseBackupsInner) HasFileExists() bool`

HasFileExists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


