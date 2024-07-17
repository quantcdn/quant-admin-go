# ProjectRequestUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] [default to "au"]
**AllowQueryParams** | Pointer to **bool** |  | [optional] 
**BasicAuthUsername** | Pointer to **string** |  | [optional] 
**BasicAuthPassword** | Pointer to **string** |  | [optional] 
**BasicAuthPreviewOnly** | Pointer to **string** |  | [optional] 
**CustomS3SyncBucket** | Pointer to **string** |  | [optional] 
**CustomS3SyncRegion** | Pointer to **string** |  | [optional] 
**CustomS3SyncAccessKey** | Pointer to **string** |  | [optional] 
**CustomS3SyncSecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectRequestUpdate

`func NewProjectRequestUpdate() *ProjectRequestUpdate`

NewProjectRequestUpdate instantiates a new ProjectRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRequestUpdateWithDefaults

`func NewProjectRequestUpdateWithDefaults() *ProjectRequestUpdate`

NewProjectRequestUpdateWithDefaults instantiates a new ProjectRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectRequestUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectRequestUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectRequestUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectRequestUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *ProjectRequestUpdate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProjectRequestUpdate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProjectRequestUpdate) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProjectRequestUpdate) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAllowQueryParams

`func (o *ProjectRequestUpdate) GetAllowQueryParams() bool`

GetAllowQueryParams returns the AllowQueryParams field if non-nil, zero value otherwise.

### GetAllowQueryParamsOk

`func (o *ProjectRequestUpdate) GetAllowQueryParamsOk() (*bool, bool)`

GetAllowQueryParamsOk returns a tuple with the AllowQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryParams

`func (o *ProjectRequestUpdate) SetAllowQueryParams(v bool)`

SetAllowQueryParams sets AllowQueryParams field to given value.

### HasAllowQueryParams

`func (o *ProjectRequestUpdate) HasAllowQueryParams() bool`

HasAllowQueryParams returns a boolean if a field has been set.

### GetBasicAuthUsername

`func (o *ProjectRequestUpdate) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *ProjectRequestUpdate) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *ProjectRequestUpdate) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.

### HasBasicAuthUsername

`func (o *ProjectRequestUpdate) HasBasicAuthUsername() bool`

HasBasicAuthUsername returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *ProjectRequestUpdate) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *ProjectRequestUpdate) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *ProjectRequestUpdate) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.

### HasBasicAuthPassword

`func (o *ProjectRequestUpdate) HasBasicAuthPassword() bool`

HasBasicAuthPassword returns a boolean if a field has been set.

### GetBasicAuthPreviewOnly

`func (o *ProjectRequestUpdate) GetBasicAuthPreviewOnly() string`

GetBasicAuthPreviewOnly returns the BasicAuthPreviewOnly field if non-nil, zero value otherwise.

### GetBasicAuthPreviewOnlyOk

`func (o *ProjectRequestUpdate) GetBasicAuthPreviewOnlyOk() (*string, bool)`

GetBasicAuthPreviewOnlyOk returns a tuple with the BasicAuthPreviewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPreviewOnly

`func (o *ProjectRequestUpdate) SetBasicAuthPreviewOnly(v string)`

SetBasicAuthPreviewOnly sets BasicAuthPreviewOnly field to given value.

### HasBasicAuthPreviewOnly

`func (o *ProjectRequestUpdate) HasBasicAuthPreviewOnly() bool`

HasBasicAuthPreviewOnly returns a boolean if a field has been set.

### GetCustomS3SyncBucket

`func (o *ProjectRequestUpdate) GetCustomS3SyncBucket() string`

GetCustomS3SyncBucket returns the CustomS3SyncBucket field if non-nil, zero value otherwise.

### GetCustomS3SyncBucketOk

`func (o *ProjectRequestUpdate) GetCustomS3SyncBucketOk() (*string, bool)`

GetCustomS3SyncBucketOk returns a tuple with the CustomS3SyncBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncBucket

`func (o *ProjectRequestUpdate) SetCustomS3SyncBucket(v string)`

SetCustomS3SyncBucket sets CustomS3SyncBucket field to given value.

### HasCustomS3SyncBucket

`func (o *ProjectRequestUpdate) HasCustomS3SyncBucket() bool`

HasCustomS3SyncBucket returns a boolean if a field has been set.

### GetCustomS3SyncRegion

`func (o *ProjectRequestUpdate) GetCustomS3SyncRegion() string`

GetCustomS3SyncRegion returns the CustomS3SyncRegion field if non-nil, zero value otherwise.

### GetCustomS3SyncRegionOk

`func (o *ProjectRequestUpdate) GetCustomS3SyncRegionOk() (*string, bool)`

GetCustomS3SyncRegionOk returns a tuple with the CustomS3SyncRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncRegion

`func (o *ProjectRequestUpdate) SetCustomS3SyncRegion(v string)`

SetCustomS3SyncRegion sets CustomS3SyncRegion field to given value.

### HasCustomS3SyncRegion

`func (o *ProjectRequestUpdate) HasCustomS3SyncRegion() bool`

HasCustomS3SyncRegion returns a boolean if a field has been set.

### GetCustomS3SyncAccessKey

`func (o *ProjectRequestUpdate) GetCustomS3SyncAccessKey() string`

GetCustomS3SyncAccessKey returns the CustomS3SyncAccessKey field if non-nil, zero value otherwise.

### GetCustomS3SyncAccessKeyOk

`func (o *ProjectRequestUpdate) GetCustomS3SyncAccessKeyOk() (*string, bool)`

GetCustomS3SyncAccessKeyOk returns a tuple with the CustomS3SyncAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncAccessKey

`func (o *ProjectRequestUpdate) SetCustomS3SyncAccessKey(v string)`

SetCustomS3SyncAccessKey sets CustomS3SyncAccessKey field to given value.

### HasCustomS3SyncAccessKey

`func (o *ProjectRequestUpdate) HasCustomS3SyncAccessKey() bool`

HasCustomS3SyncAccessKey returns a boolean if a field has been set.

### GetCustomS3SyncSecretKey

`func (o *ProjectRequestUpdate) GetCustomS3SyncSecretKey() string`

GetCustomS3SyncSecretKey returns the CustomS3SyncSecretKey field if non-nil, zero value otherwise.

### GetCustomS3SyncSecretKeyOk

`func (o *ProjectRequestUpdate) GetCustomS3SyncSecretKeyOk() (*string, bool)`

GetCustomS3SyncSecretKeyOk returns a tuple with the CustomS3SyncSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncSecretKey

`func (o *ProjectRequestUpdate) SetCustomS3SyncSecretKey(v string)`

SetCustomS3SyncSecretKey sets CustomS3SyncSecretKey field to given value.

### HasCustomS3SyncSecretKey

`func (o *ProjectRequestUpdate) HasCustomS3SyncSecretKey() bool`

HasCustomS3SyncSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


