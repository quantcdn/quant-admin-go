# DeleteVectorDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurgeAll** | Pointer to **bool** | Delete ALL documents in collection | [optional] 
**DocumentIds** | Pointer to **[]string** | Delete specific documents by UUID | [optional] 
**Metadata** | Pointer to [**DeleteVectorDocumentsRequestMetadata**](DeleteVectorDocumentsRequestMetadata.md) |  | [optional] 

## Methods

### NewDeleteVectorDocumentsRequest

`func NewDeleteVectorDocumentsRequest() *DeleteVectorDocumentsRequest`

NewDeleteVectorDocumentsRequest instantiates a new DeleteVectorDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVectorDocumentsRequestWithDefaults

`func NewDeleteVectorDocumentsRequestWithDefaults() *DeleteVectorDocumentsRequest`

NewDeleteVectorDocumentsRequestWithDefaults instantiates a new DeleteVectorDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurgeAll

`func (o *DeleteVectorDocumentsRequest) GetPurgeAll() bool`

GetPurgeAll returns the PurgeAll field if non-nil, zero value otherwise.

### GetPurgeAllOk

`func (o *DeleteVectorDocumentsRequest) GetPurgeAllOk() (*bool, bool)`

GetPurgeAllOk returns a tuple with the PurgeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeAll

`func (o *DeleteVectorDocumentsRequest) SetPurgeAll(v bool)`

SetPurgeAll sets PurgeAll field to given value.

### HasPurgeAll

`func (o *DeleteVectorDocumentsRequest) HasPurgeAll() bool`

HasPurgeAll returns a boolean if a field has been set.

### GetDocumentIds

`func (o *DeleteVectorDocumentsRequest) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *DeleteVectorDocumentsRequest) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *DeleteVectorDocumentsRequest) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *DeleteVectorDocumentsRequest) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetMetadata

`func (o *DeleteVectorDocumentsRequest) GetMetadata() DeleteVectorDocumentsRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeleteVectorDocumentsRequest) GetMetadataOk() (*DeleteVectorDocumentsRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeleteVectorDocumentsRequest) SetMetadata(v DeleteVectorDocumentsRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeleteVectorDocumentsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


