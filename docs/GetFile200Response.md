# GetFile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** |  | [optional] 
**S3Uri** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | Presigned download URL (1 hour) | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetFile200Response

`func NewGetFile200Response() *GetFile200Response`

NewGetFile200Response instantiates a new GetFile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFile200ResponseWithDefaults

`func NewGetFile200ResponseWithDefaults() *GetFile200Response`

NewGetFile200ResponseWithDefaults instantiates a new GetFile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *GetFile200Response) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *GetFile200Response) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *GetFile200Response) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *GetFile200Response) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetS3Uri

`func (o *GetFile200Response) GetS3Uri() string`

GetS3Uri returns the S3Uri field if non-nil, zero value otherwise.

### GetS3UriOk

`func (o *GetFile200Response) GetS3UriOk() (*string, bool)`

GetS3UriOk returns a tuple with the S3Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Uri

`func (o *GetFile200Response) SetS3Uri(v string)`

SetS3Uri sets S3Uri field to given value.

### HasS3Uri

`func (o *GetFile200Response) HasS3Uri() bool`

HasS3Uri returns a boolean if a field has been set.

### GetUrl

`func (o *GetFile200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetFile200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetFile200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetFile200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFilename

`func (o *GetFile200Response) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *GetFile200Response) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *GetFile200Response) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *GetFile200Response) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetContentType

`func (o *GetFile200Response) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetFile200Response) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetFile200Response) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GetFile200Response) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *GetFile200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetFile200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetFile200Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetFile200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMetadata

`func (o *GetFile200Response) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetFile200Response) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetFile200Response) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetFile200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetFile200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetFile200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetFile200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetFile200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


