# V2KvLedgerErasureSubjectPreviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | [**V2KvLedgerSubjectId**](V2KvLedgerSubjectId.md) |  | 
**RefKeyVersion** | Pointer to **int32** | Optional: the version of the store&#39;s reference key. Sent to the server unchanged. | [optional] 

## Methods

### NewV2KvLedgerErasureSubjectPreviewRequest

`func NewV2KvLedgerErasureSubjectPreviewRequest(subjectId V2KvLedgerSubjectId, ) *V2KvLedgerErasureSubjectPreviewRequest`

NewV2KvLedgerErasureSubjectPreviewRequest instantiates a new V2KvLedgerErasureSubjectPreviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureSubjectPreviewRequestWithDefaults

`func NewV2KvLedgerErasureSubjectPreviewRequestWithDefaults() *V2KvLedgerErasureSubjectPreviewRequest`

NewV2KvLedgerErasureSubjectPreviewRequestWithDefaults instantiates a new V2KvLedgerErasureSubjectPreviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *V2KvLedgerErasureSubjectPreviewRequest) GetSubjectId() V2KvLedgerSubjectId`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *V2KvLedgerErasureSubjectPreviewRequest) GetSubjectIdOk() (*V2KvLedgerSubjectId, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *V2KvLedgerErasureSubjectPreviewRequest) SetSubjectId(v V2KvLedgerSubjectId)`

SetSubjectId sets SubjectId field to given value.


### GetRefKeyVersion

`func (o *V2KvLedgerErasureSubjectPreviewRequest) GetRefKeyVersion() int32`

GetRefKeyVersion returns the RefKeyVersion field if non-nil, zero value otherwise.

### GetRefKeyVersionOk

`func (o *V2KvLedgerErasureSubjectPreviewRequest) GetRefKeyVersionOk() (*int32, bool)`

GetRefKeyVersionOk returns a tuple with the RefKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefKeyVersion

`func (o *V2KvLedgerErasureSubjectPreviewRequest) SetRefKeyVersion(v int32)`

SetRefKeyVersion sets RefKeyVersion field to given value.

### HasRefKeyVersion

`func (o *V2KvLedgerErasureSubjectPreviewRequest) HasRefKeyVersion() bool`

HasRefKeyVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


