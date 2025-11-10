# DownloadBackup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadUrl** | Pointer to **string** | Pre-signed S3 URL for download | [optional] 
**ExpiresAt** | Pointer to **time.Time** | URL expiration time | [optional] 
**Filename** | Pointer to **string** | Suggested filename for download | [optional] 

## Methods

### NewDownloadBackup200Response

`func NewDownloadBackup200Response() *DownloadBackup200Response`

NewDownloadBackup200Response instantiates a new DownloadBackup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadBackup200ResponseWithDefaults

`func NewDownloadBackup200ResponseWithDefaults() *DownloadBackup200Response`

NewDownloadBackup200ResponseWithDefaults instantiates a new DownloadBackup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadUrl

`func (o *DownloadBackup200Response) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *DownloadBackup200Response) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *DownloadBackup200Response) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *DownloadBackup200Response) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DownloadBackup200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DownloadBackup200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DownloadBackup200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DownloadBackup200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFilename

`func (o *DownloadBackup200Response) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *DownloadBackup200Response) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *DownloadBackup200Response) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *DownloadBackup200Response) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


