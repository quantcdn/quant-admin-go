# UploadFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Base64-encoded file content (for direct upload). Required unless using requestUploadUrl. | [optional] 
**RequestUploadUrl** | Pointer to **bool** | Set to true to get a presigned S3 upload URL instead of uploading directly. | [optional] [default to false]
**Size** | Pointer to **int32** | File size in bytes. Optional but recommended for presigned uploads. | [optional] 
**Filename** | Pointer to **string** | Original filename | [optional] 
**ContentType** | **string** | MIME type of the file | 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata for filtering. Any fields allowed. | [optional] 

## Methods

### NewUploadFileRequest

`func NewUploadFileRequest(contentType string, ) *UploadFileRequest`

NewUploadFileRequest instantiates a new UploadFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadFileRequestWithDefaults

`func NewUploadFileRequestWithDefaults() *UploadFileRequest`

NewUploadFileRequestWithDefaults instantiates a new UploadFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *UploadFileRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UploadFileRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UploadFileRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UploadFileRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRequestUploadUrl

`func (o *UploadFileRequest) GetRequestUploadUrl() bool`

GetRequestUploadUrl returns the RequestUploadUrl field if non-nil, zero value otherwise.

### GetRequestUploadUrlOk

`func (o *UploadFileRequest) GetRequestUploadUrlOk() (*bool, bool)`

GetRequestUploadUrlOk returns a tuple with the RequestUploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUploadUrl

`func (o *UploadFileRequest) SetRequestUploadUrl(v bool)`

SetRequestUploadUrl sets RequestUploadUrl field to given value.

### HasRequestUploadUrl

`func (o *UploadFileRequest) HasRequestUploadUrl() bool`

HasRequestUploadUrl returns a boolean if a field has been set.

### GetSize

`func (o *UploadFileRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UploadFileRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UploadFileRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *UploadFileRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetFilename

`func (o *UploadFileRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadFileRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadFileRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *UploadFileRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetContentType

`func (o *UploadFileRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UploadFileRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UploadFileRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetMetadata

`func (o *UploadFileRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UploadFileRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UploadFileRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UploadFileRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


