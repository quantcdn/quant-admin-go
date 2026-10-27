# V2KvLedgerErasureSubjectPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledger** | **string** |  | 
**SubjectRef** | **string** | The reference that the erasure records and receipts use for this subject. The answer never carries the subject id. | 
**RefKeyVersion** | **int32** | The version of the reference key that made subjectRef. | 
**Found** | **bool** |  | 
**KeyState** | **string** | live: a shred would destroy the key. shredded: the subject is already shredded. | 
**ShreddedAt** | Pointer to **time.Time** | Present only when keyState is shredded. | [optional] 
**Complete** | **bool** | True when every collection was counted to its end. Even then each count means at least N (eventual reads). | 
**Entries** | Pointer to **int32** | The total of the collection counts. Present only when keyState is live. | [optional] 
**Collections** | [**[]V2KvLedgerErasureSubjectPreviewCollectionsInner**](V2KvLedgerErasureSubjectPreviewCollectionsInner.md) | Every collection that holds the subject, in name order. | 

## Methods

### NewV2KvLedgerErasureSubjectPreview

`func NewV2KvLedgerErasureSubjectPreview(ledger string, subjectRef string, refKeyVersion int32, found bool, keyState string, complete bool, collections []V2KvLedgerErasureSubjectPreviewCollectionsInner, ) *V2KvLedgerErasureSubjectPreview`

NewV2KvLedgerErasureSubjectPreview instantiates a new V2KvLedgerErasureSubjectPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureSubjectPreviewWithDefaults

`func NewV2KvLedgerErasureSubjectPreviewWithDefaults() *V2KvLedgerErasureSubjectPreview`

NewV2KvLedgerErasureSubjectPreviewWithDefaults instantiates a new V2KvLedgerErasureSubjectPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedger

`func (o *V2KvLedgerErasureSubjectPreview) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerErasureSubjectPreview) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerErasureSubjectPreview) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### GetSubjectRef

`func (o *V2KvLedgerErasureSubjectPreview) GetSubjectRef() string`

GetSubjectRef returns the SubjectRef field if non-nil, zero value otherwise.

### GetSubjectRefOk

`func (o *V2KvLedgerErasureSubjectPreview) GetSubjectRefOk() (*string, bool)`

GetSubjectRefOk returns a tuple with the SubjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRef

`func (o *V2KvLedgerErasureSubjectPreview) SetSubjectRef(v string)`

SetSubjectRef sets SubjectRef field to given value.


### GetRefKeyVersion

`func (o *V2KvLedgerErasureSubjectPreview) GetRefKeyVersion() int32`

GetRefKeyVersion returns the RefKeyVersion field if non-nil, zero value otherwise.

### GetRefKeyVersionOk

`func (o *V2KvLedgerErasureSubjectPreview) GetRefKeyVersionOk() (*int32, bool)`

GetRefKeyVersionOk returns a tuple with the RefKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefKeyVersion

`func (o *V2KvLedgerErasureSubjectPreview) SetRefKeyVersion(v int32)`

SetRefKeyVersion sets RefKeyVersion field to given value.


### GetFound

`func (o *V2KvLedgerErasureSubjectPreview) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *V2KvLedgerErasureSubjectPreview) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *V2KvLedgerErasureSubjectPreview) SetFound(v bool)`

SetFound sets Found field to given value.


### GetKeyState

`func (o *V2KvLedgerErasureSubjectPreview) GetKeyState() string`

GetKeyState returns the KeyState field if non-nil, zero value otherwise.

### GetKeyStateOk

`func (o *V2KvLedgerErasureSubjectPreview) GetKeyStateOk() (*string, bool)`

GetKeyStateOk returns a tuple with the KeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyState

`func (o *V2KvLedgerErasureSubjectPreview) SetKeyState(v string)`

SetKeyState sets KeyState field to given value.


### GetShreddedAt

`func (o *V2KvLedgerErasureSubjectPreview) GetShreddedAt() time.Time`

GetShreddedAt returns the ShreddedAt field if non-nil, zero value otherwise.

### GetShreddedAtOk

`func (o *V2KvLedgerErasureSubjectPreview) GetShreddedAtOk() (*time.Time, bool)`

GetShreddedAtOk returns a tuple with the ShreddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShreddedAt

`func (o *V2KvLedgerErasureSubjectPreview) SetShreddedAt(v time.Time)`

SetShreddedAt sets ShreddedAt field to given value.

### HasShreddedAt

`func (o *V2KvLedgerErasureSubjectPreview) HasShreddedAt() bool`

HasShreddedAt returns a boolean if a field has been set.

### GetComplete

`func (o *V2KvLedgerErasureSubjectPreview) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *V2KvLedgerErasureSubjectPreview) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *V2KvLedgerErasureSubjectPreview) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetEntries

`func (o *V2KvLedgerErasureSubjectPreview) GetEntries() int32`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerErasureSubjectPreview) GetEntriesOk() (*int32, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerErasureSubjectPreview) SetEntries(v int32)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *V2KvLedgerErasureSubjectPreview) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetCollections

`func (o *V2KvLedgerErasureSubjectPreview) GetCollections() []V2KvLedgerErasureSubjectPreviewCollectionsInner`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *V2KvLedgerErasureSubjectPreview) GetCollectionsOk() (*[]V2KvLedgerErasureSubjectPreviewCollectionsInner, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *V2KvLedgerErasureSubjectPreview) SetCollections(v []V2KvLedgerErasureSubjectPreviewCollectionsInner)`

SetCollections sets Collections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


