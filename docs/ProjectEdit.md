# ProjectEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Region** | **string** |  | 
**AllowQueryParams** | **bool** |  | 
**BasicAuthUsername** | Pointer to **string** |  | [optional] 
**BasicAuthPassword** | Pointer to **string** |  | [optional] 
**BasicAuthPreviewOnly** | Pointer to **bool** |  | [optional] 
**CustomS3SyncBucket** | Pointer to **string** | Optional bucket name for S3 sync function | [optional] 
**CustomS3SyncRegion** | Pointer to **string** | Bucket region for S3 sync. Required if s3 sync bucket is provided | [optional] 
**CustomS3SyncAccessKey** | Pointer to **string** | Access key for S3 sync. Required if s3 sync bucket is provided | [optional] 
**CustomS3SyncSecretKey** | Pointer to **string** | Secret key for S3 sync. Required if s3 sync bucket is provided | [optional] 

## Methods

### NewProjectEdit

`func NewProjectEdit(name string, region string, allowQueryParams bool, ) *ProjectEdit`

NewProjectEdit instantiates a new ProjectEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectEditWithDefaults

`func NewProjectEditWithDefaults() *ProjectEdit`

NewProjectEditWithDefaults instantiates a new ProjectEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectEdit) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *ProjectEdit) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProjectEdit) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProjectEdit) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAllowQueryParams

`func (o *ProjectEdit) GetAllowQueryParams() bool`

GetAllowQueryParams returns the AllowQueryParams field if non-nil, zero value otherwise.

### GetAllowQueryParamsOk

`func (o *ProjectEdit) GetAllowQueryParamsOk() (*bool, bool)`

GetAllowQueryParamsOk returns a tuple with the AllowQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryParams

`func (o *ProjectEdit) SetAllowQueryParams(v bool)`

SetAllowQueryParams sets AllowQueryParams field to given value.


### GetBasicAuthUsername

`func (o *ProjectEdit) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *ProjectEdit) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *ProjectEdit) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.

### HasBasicAuthUsername

`func (o *ProjectEdit) HasBasicAuthUsername() bool`

HasBasicAuthUsername returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *ProjectEdit) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *ProjectEdit) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *ProjectEdit) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.

### HasBasicAuthPassword

`func (o *ProjectEdit) HasBasicAuthPassword() bool`

HasBasicAuthPassword returns a boolean if a field has been set.

### GetBasicAuthPreviewOnly

`func (o *ProjectEdit) GetBasicAuthPreviewOnly() bool`

GetBasicAuthPreviewOnly returns the BasicAuthPreviewOnly field if non-nil, zero value otherwise.

### GetBasicAuthPreviewOnlyOk

`func (o *ProjectEdit) GetBasicAuthPreviewOnlyOk() (*bool, bool)`

GetBasicAuthPreviewOnlyOk returns a tuple with the BasicAuthPreviewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPreviewOnly

`func (o *ProjectEdit) SetBasicAuthPreviewOnly(v bool)`

SetBasicAuthPreviewOnly sets BasicAuthPreviewOnly field to given value.

### HasBasicAuthPreviewOnly

`func (o *ProjectEdit) HasBasicAuthPreviewOnly() bool`

HasBasicAuthPreviewOnly returns a boolean if a field has been set.

### GetCustomS3SyncBucket

`func (o *ProjectEdit) GetCustomS3SyncBucket() string`

GetCustomS3SyncBucket returns the CustomS3SyncBucket field if non-nil, zero value otherwise.

### GetCustomS3SyncBucketOk

`func (o *ProjectEdit) GetCustomS3SyncBucketOk() (*string, bool)`

GetCustomS3SyncBucketOk returns a tuple with the CustomS3SyncBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncBucket

`func (o *ProjectEdit) SetCustomS3SyncBucket(v string)`

SetCustomS3SyncBucket sets CustomS3SyncBucket field to given value.

### HasCustomS3SyncBucket

`func (o *ProjectEdit) HasCustomS3SyncBucket() bool`

HasCustomS3SyncBucket returns a boolean if a field has been set.

### GetCustomS3SyncRegion

`func (o *ProjectEdit) GetCustomS3SyncRegion() string`

GetCustomS3SyncRegion returns the CustomS3SyncRegion field if non-nil, zero value otherwise.

### GetCustomS3SyncRegionOk

`func (o *ProjectEdit) GetCustomS3SyncRegionOk() (*string, bool)`

GetCustomS3SyncRegionOk returns a tuple with the CustomS3SyncRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncRegion

`func (o *ProjectEdit) SetCustomS3SyncRegion(v string)`

SetCustomS3SyncRegion sets CustomS3SyncRegion field to given value.

### HasCustomS3SyncRegion

`func (o *ProjectEdit) HasCustomS3SyncRegion() bool`

HasCustomS3SyncRegion returns a boolean if a field has been set.

### GetCustomS3SyncAccessKey

`func (o *ProjectEdit) GetCustomS3SyncAccessKey() string`

GetCustomS3SyncAccessKey returns the CustomS3SyncAccessKey field if non-nil, zero value otherwise.

### GetCustomS3SyncAccessKeyOk

`func (o *ProjectEdit) GetCustomS3SyncAccessKeyOk() (*string, bool)`

GetCustomS3SyncAccessKeyOk returns a tuple with the CustomS3SyncAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncAccessKey

`func (o *ProjectEdit) SetCustomS3SyncAccessKey(v string)`

SetCustomS3SyncAccessKey sets CustomS3SyncAccessKey field to given value.

### HasCustomS3SyncAccessKey

`func (o *ProjectEdit) HasCustomS3SyncAccessKey() bool`

HasCustomS3SyncAccessKey returns a boolean if a field has been set.

### GetCustomS3SyncSecretKey

`func (o *ProjectEdit) GetCustomS3SyncSecretKey() string`

GetCustomS3SyncSecretKey returns the CustomS3SyncSecretKey field if non-nil, zero value otherwise.

### GetCustomS3SyncSecretKeyOk

`func (o *ProjectEdit) GetCustomS3SyncSecretKeyOk() (*string, bool)`

GetCustomS3SyncSecretKeyOk returns a tuple with the CustomS3SyncSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncSecretKey

`func (o *ProjectEdit) SetCustomS3SyncSecretKey(v string)`

SetCustomS3SyncSecretKey sets CustomS3SyncSecretKey field to given value.

### HasCustomS3SyncSecretKey

`func (o *ProjectEdit) HasCustomS3SyncSecretKey() bool`

HasCustomS3SyncSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


