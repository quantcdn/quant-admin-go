# V2KvExportDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Part** | **int32** |  | 
**Url** | **string** | Presigned URL, valid for 1 hour, new on every call. | 
**SizeBytes** | **int32** |  | 

## Methods

### NewV2KvExportDownload

`func NewV2KvExportDownload(part int32, url string, sizeBytes int32, ) *V2KvExportDownload`

NewV2KvExportDownload instantiates a new V2KvExportDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvExportDownloadWithDefaults

`func NewV2KvExportDownloadWithDefaults() *V2KvExportDownload`

NewV2KvExportDownloadWithDefaults instantiates a new V2KvExportDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPart

`func (o *V2KvExportDownload) GetPart() int32`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *V2KvExportDownload) GetPartOk() (*int32, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *V2KvExportDownload) SetPart(v int32)`

SetPart sets Part field to given value.


### GetUrl

`func (o *V2KvExportDownload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2KvExportDownload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2KvExportDownload) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSizeBytes

`func (o *V2KvExportDownload) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *V2KvExportDownload) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *V2KvExportDownload) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


