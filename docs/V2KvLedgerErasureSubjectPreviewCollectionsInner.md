# V2KvLedgerErasureSubjectPreviewCollectionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** |  | 
**Entries** | **int32** | At least this many entries of the subject (shred entries not counted). | 
**Complete** | **bool** | false: the time budget ended before this collection was counted to its end. | 
**FirstAt** | Pointer to **time.Time** |  | [optional] 
**LastAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewV2KvLedgerErasureSubjectPreviewCollectionsInner

`func NewV2KvLedgerErasureSubjectPreviewCollectionsInner(collection string, entries int32, complete bool, ) *V2KvLedgerErasureSubjectPreviewCollectionsInner`

NewV2KvLedgerErasureSubjectPreviewCollectionsInner instantiates a new V2KvLedgerErasureSubjectPreviewCollectionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureSubjectPreviewCollectionsInnerWithDefaults

`func NewV2KvLedgerErasureSubjectPreviewCollectionsInnerWithDefaults() *V2KvLedgerErasureSubjectPreviewCollectionsInner`

NewV2KvLedgerErasureSubjectPreviewCollectionsInnerWithDefaults instantiates a new V2KvLedgerErasureSubjectPreviewCollectionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetEntries

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetEntries() int32`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetEntriesOk() (*int32, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) SetEntries(v int32)`

SetEntries sets Entries field to given value.


### GetComplete

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetFirstAt

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetFirstAt() time.Time`

GetFirstAt returns the FirstAt field if non-nil, zero value otherwise.

### GetFirstAtOk

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetFirstAtOk() (*time.Time, bool)`

GetFirstAtOk returns a tuple with the FirstAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAt

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) SetFirstAt(v time.Time)`

SetFirstAt sets FirstAt field to given value.

### HasFirstAt

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) HasFirstAt() bool`

HasFirstAt returns a boolean if a field has been set.

### GetLastAt

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetLastAt() time.Time`

GetLastAt returns the LastAt field if non-nil, zero value otherwise.

### GetLastAtOk

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) GetLastAtOk() (*time.Time, bool)`

GetLastAtOk returns a tuple with the LastAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAt

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) SetLastAt(v time.Time)`

SetLastAt sets LastAt field to given value.

### HasLastAt

`func (o *V2KvLedgerErasureSubjectPreviewCollectionsInner) HasLastAt() bool`

HasLastAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


