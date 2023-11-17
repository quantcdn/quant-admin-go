# ProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**AllowQueryParams** | Pointer to **bool** |  | [optional] 
**BasicAuthUsername** | Pointer to **string** |  | [optional] 
**BasicAuthPassword** | Pointer to **string** |  | [optional] 
**BasicAuthPreviewOnly** | Pointer to **string** |  | [optional] 
**CustomS3SyncBucket** | Pointer to **string** |  | [optional] 
**CustomS3SyncRegion** | Pointer to **string** |  | [optional] 
**CustomS3SyncAccessKey** | Pointer to **string** |  | [optional] 
**CustomS3SyncSecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectRequest

`func NewProjectRequest() *ProjectRequest`

NewProjectRequest instantiates a new ProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRequestWithDefaults

`func NewProjectRequestWithDefaults() *ProjectRequest`

NewProjectRequestWithDefaults instantiates a new ProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *ProjectRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProjectRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProjectRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProjectRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAllowQueryParams

`func (o *ProjectRequest) GetAllowQueryParams() bool`

GetAllowQueryParams returns the AllowQueryParams field if non-nil, zero value otherwise.

### GetAllowQueryParamsOk

`func (o *ProjectRequest) GetAllowQueryParamsOk() (*bool, bool)`

GetAllowQueryParamsOk returns a tuple with the AllowQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryParams

`func (o *ProjectRequest) SetAllowQueryParams(v bool)`

SetAllowQueryParams sets AllowQueryParams field to given value.

### HasAllowQueryParams

`func (o *ProjectRequest) HasAllowQueryParams() bool`

HasAllowQueryParams returns a boolean if a field has been set.

### GetBasicAuthUsername

`func (o *ProjectRequest) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *ProjectRequest) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *ProjectRequest) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.

### HasBasicAuthUsername

`func (o *ProjectRequest) HasBasicAuthUsername() bool`

HasBasicAuthUsername returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *ProjectRequest) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *ProjectRequest) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *ProjectRequest) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.

### HasBasicAuthPassword

`func (o *ProjectRequest) HasBasicAuthPassword() bool`

HasBasicAuthPassword returns a boolean if a field has been set.

### GetBasicAuthPreviewOnly

`func (o *ProjectRequest) GetBasicAuthPreviewOnly() string`

GetBasicAuthPreviewOnly returns the BasicAuthPreviewOnly field if non-nil, zero value otherwise.

### GetBasicAuthPreviewOnlyOk

`func (o *ProjectRequest) GetBasicAuthPreviewOnlyOk() (*string, bool)`

GetBasicAuthPreviewOnlyOk returns a tuple with the BasicAuthPreviewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPreviewOnly

`func (o *ProjectRequest) SetBasicAuthPreviewOnly(v string)`

SetBasicAuthPreviewOnly sets BasicAuthPreviewOnly field to given value.

### HasBasicAuthPreviewOnly

`func (o *ProjectRequest) HasBasicAuthPreviewOnly() bool`

HasBasicAuthPreviewOnly returns a boolean if a field has been set.

### GetCustomS3SyncBucket

`func (o *ProjectRequest) GetCustomS3SyncBucket() string`

GetCustomS3SyncBucket returns the CustomS3SyncBucket field if non-nil, zero value otherwise.

### GetCustomS3SyncBucketOk

`func (o *ProjectRequest) GetCustomS3SyncBucketOk() (*string, bool)`

GetCustomS3SyncBucketOk returns a tuple with the CustomS3SyncBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncBucket

`func (o *ProjectRequest) SetCustomS3SyncBucket(v string)`

SetCustomS3SyncBucket sets CustomS3SyncBucket field to given value.

### HasCustomS3SyncBucket

`func (o *ProjectRequest) HasCustomS3SyncBucket() bool`

HasCustomS3SyncBucket returns a boolean if a field has been set.

### GetCustomS3SyncRegion

`func (o *ProjectRequest) GetCustomS3SyncRegion() string`

GetCustomS3SyncRegion returns the CustomS3SyncRegion field if non-nil, zero value otherwise.

### GetCustomS3SyncRegionOk

`func (o *ProjectRequest) GetCustomS3SyncRegionOk() (*string, bool)`

GetCustomS3SyncRegionOk returns a tuple with the CustomS3SyncRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncRegion

`func (o *ProjectRequest) SetCustomS3SyncRegion(v string)`

SetCustomS3SyncRegion sets CustomS3SyncRegion field to given value.

### HasCustomS3SyncRegion

`func (o *ProjectRequest) HasCustomS3SyncRegion() bool`

HasCustomS3SyncRegion returns a boolean if a field has been set.

### GetCustomS3SyncAccessKey

`func (o *ProjectRequest) GetCustomS3SyncAccessKey() string`

GetCustomS3SyncAccessKey returns the CustomS3SyncAccessKey field if non-nil, zero value otherwise.

### GetCustomS3SyncAccessKeyOk

`func (o *ProjectRequest) GetCustomS3SyncAccessKeyOk() (*string, bool)`

GetCustomS3SyncAccessKeyOk returns a tuple with the CustomS3SyncAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncAccessKey

`func (o *ProjectRequest) SetCustomS3SyncAccessKey(v string)`

SetCustomS3SyncAccessKey sets CustomS3SyncAccessKey field to given value.

### HasCustomS3SyncAccessKey

`func (o *ProjectRequest) HasCustomS3SyncAccessKey() bool`

HasCustomS3SyncAccessKey returns a boolean if a field has been set.

### GetCustomS3SyncSecretKey

`func (o *ProjectRequest) GetCustomS3SyncSecretKey() string`

GetCustomS3SyncSecretKey returns the CustomS3SyncSecretKey field if non-nil, zero value otherwise.

### GetCustomS3SyncSecretKeyOk

`func (o *ProjectRequest) GetCustomS3SyncSecretKeyOk() (*string, bool)`

GetCustomS3SyncSecretKeyOk returns a tuple with the CustomS3SyncSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomS3SyncSecretKey

`func (o *ProjectRequest) SetCustomS3SyncSecretKey(v string)`

SetCustomS3SyncSecretKey sets CustomS3SyncSecretKey field to given value.

### HasCustomS3SyncSecretKey

`func (o *ProjectRequest) HasCustomS3SyncSecretKey() bool`

HasCustomS3SyncSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


