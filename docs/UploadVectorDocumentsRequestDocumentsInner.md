# UploadVectorDocumentsRequestDocumentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Document text content | 
**Metadata** | Pointer to [**UploadVectorDocumentsRequestDocumentsInnerMetadata**](UploadVectorDocumentsRequestDocumentsInnerMetadata.md) |  | [optional] 

## Methods

### NewUploadVectorDocumentsRequestDocumentsInner

`func NewUploadVectorDocumentsRequestDocumentsInner(content string, ) *UploadVectorDocumentsRequestDocumentsInner`

NewUploadVectorDocumentsRequestDocumentsInner instantiates a new UploadVectorDocumentsRequestDocumentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadVectorDocumentsRequestDocumentsInnerWithDefaults

`func NewUploadVectorDocumentsRequestDocumentsInnerWithDefaults() *UploadVectorDocumentsRequestDocumentsInner`

NewUploadVectorDocumentsRequestDocumentsInnerWithDefaults instantiates a new UploadVectorDocumentsRequestDocumentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *UploadVectorDocumentsRequestDocumentsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UploadVectorDocumentsRequestDocumentsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UploadVectorDocumentsRequestDocumentsInner) SetContent(v string)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *UploadVectorDocumentsRequestDocumentsInner) GetMetadata() UploadVectorDocumentsRequestDocumentsInnerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UploadVectorDocumentsRequestDocumentsInner) GetMetadataOk() (*UploadVectorDocumentsRequestDocumentsInnerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UploadVectorDocumentsRequestDocumentsInner) SetMetadata(v UploadVectorDocumentsRequestDocumentsInnerMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UploadVectorDocumentsRequestDocumentsInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


