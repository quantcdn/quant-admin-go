# UploadFile201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** |  | [optional] 
**S3Uri** | Pointer to **string** | S3 URI (direct upload only) | [optional] 
**Url** | Pointer to **string** | Presigned download URL (direct upload only) | [optional] 
**UploadUrl** | Pointer to **string** | Presigned PUT URL (presigned upload only) | [optional] 
**S3Key** | Pointer to **string** | S3 object key (presigned upload only) | [optional] 
**ExpiresIn** | Pointer to **int32** | URL expiry in seconds (presigned upload only) | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUploadFile201Response

`func NewUploadFile201Response() *UploadFile201Response`

NewUploadFile201Response instantiates a new UploadFile201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFile201ResponseWithDefaults

`func NewUploadFile201ResponseWithDefaults() *UploadFile201Response`

NewUploadFile201ResponseWithDefaults instantiates a new UploadFile201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *UploadFile201Response) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *UploadFile201Response) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *UploadFile201Response) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *UploadFile201Response) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetS3Uri

`func (o *UploadFile201Response) GetS3Uri() string`

GetS3Uri returns the S3Uri field if non-nil, zero value otherwise.

### GetS3UriOk

`func (o *UploadFile201Response) GetS3UriOk() (*string, bool)`

GetS3UriOk returns a tuple with the S3Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Uri

`func (o *UploadFile201Response) SetS3Uri(v string)`

SetS3Uri sets S3Uri field to given value.

### HasS3Uri

`func (o *UploadFile201Response) HasS3Uri() bool`

HasS3Uri returns a boolean if a field has been set.

### GetUrl

`func (o *UploadFile201Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadFile201Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadFile201Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UploadFile201Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUploadUrl

`func (o *UploadFile201Response) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *UploadFile201Response) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *UploadFile201Response) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *UploadFile201Response) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetS3Key

`func (o *UploadFile201Response) GetS3Key() string`

GetS3Key returns the S3Key field if non-nil, zero value otherwise.

### GetS3KeyOk

`func (o *UploadFile201Response) GetS3KeyOk() (*string, bool)`

GetS3KeyOk returns a tuple with the S3Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Key

`func (o *UploadFile201Response) SetS3Key(v string)`

SetS3Key sets S3Key field to given value.

### HasS3Key

`func (o *UploadFile201Response) HasS3Key() bool`

HasS3Key returns a boolean if a field has been set.

### GetExpiresIn

`func (o *UploadFile201Response) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *UploadFile201Response) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *UploadFile201Response) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *UploadFile201Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetFilename

`func (o *UploadFile201Response) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadFile201Response) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadFile201Response) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *UploadFile201Response) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetContentType

`func (o *UploadFile201Response) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UploadFile201Response) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UploadFile201Response) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *UploadFile201Response) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *UploadFile201Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UploadFile201Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UploadFile201Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *UploadFile201Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMetadata

`func (o *UploadFile201Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UploadFile201Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UploadFile201Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UploadFile201Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UploadFile201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UploadFile201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UploadFile201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UploadFile201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


