# V2KvLedgerErasureShredRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The shredId. | 
**SubjectRef** | **string** | The subject reference (never the subject id). | 
**ShreddedAt** | **time.Time** |  | 
**FullyErasedAfter** | **time.Time** |  | 
**Entries** | **int32** | The number of shred entries recorded. | 
**Reason** | **string** |  | 
**Authority** | **string** |  | 
**RequestedBy** | **string** |  | 
**LedgerForked** | Pointer to **bool** |  | [optional] 
**Supersedes** | Pointer to **string** |  | [optional] 
**RefKeyVersion** | Pointer to **int32** | The version of the reference key that made subjectRef. | [optional] 

## Methods

### NewV2KvLedgerErasureShredRecord

`func NewV2KvLedgerErasureShredRecord(id string, subjectRef string, shreddedAt time.Time, fullyErasedAfter time.Time, entries int32, reason string, authority string, requestedBy string, ) *V2KvLedgerErasureShredRecord`

NewV2KvLedgerErasureShredRecord instantiates a new V2KvLedgerErasureShredRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredRecordWithDefaults

`func NewV2KvLedgerErasureShredRecordWithDefaults() *V2KvLedgerErasureShredRecord`

NewV2KvLedgerErasureShredRecordWithDefaults instantiates a new V2KvLedgerErasureShredRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2KvLedgerErasureShredRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2KvLedgerErasureShredRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2KvLedgerErasureShredRecord) SetId(v string)`

SetId sets Id field to given value.


### GetSubjectRef

`func (o *V2KvLedgerErasureShredRecord) GetSubjectRef() string`

GetSubjectRef returns the SubjectRef field if non-nil, zero value otherwise.

### GetSubjectRefOk

`func (o *V2KvLedgerErasureShredRecord) GetSubjectRefOk() (*string, bool)`

GetSubjectRefOk returns a tuple with the SubjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRef

`func (o *V2KvLedgerErasureShredRecord) SetSubjectRef(v string)`

SetSubjectRef sets SubjectRef field to given value.


### GetShreddedAt

`func (o *V2KvLedgerErasureShredRecord) GetShreddedAt() time.Time`

GetShreddedAt returns the ShreddedAt field if non-nil, zero value otherwise.

### GetShreddedAtOk

`func (o *V2KvLedgerErasureShredRecord) GetShreddedAtOk() (*time.Time, bool)`

GetShreddedAtOk returns a tuple with the ShreddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShreddedAt

`func (o *V2KvLedgerErasureShredRecord) SetShreddedAt(v time.Time)`

SetShreddedAt sets ShreddedAt field to given value.


### GetFullyErasedAfter

`func (o *V2KvLedgerErasureShredRecord) GetFullyErasedAfter() time.Time`

GetFullyErasedAfter returns the FullyErasedAfter field if non-nil, zero value otherwise.

### GetFullyErasedAfterOk

`func (o *V2KvLedgerErasureShredRecord) GetFullyErasedAfterOk() (*time.Time, bool)`

GetFullyErasedAfterOk returns a tuple with the FullyErasedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyErasedAfter

`func (o *V2KvLedgerErasureShredRecord) SetFullyErasedAfter(v time.Time)`

SetFullyErasedAfter sets FullyErasedAfter field to given value.


### GetEntries

`func (o *V2KvLedgerErasureShredRecord) GetEntries() int32`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerErasureShredRecord) GetEntriesOk() (*int32, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerErasureShredRecord) SetEntries(v int32)`

SetEntries sets Entries field to given value.


### GetReason

`func (o *V2KvLedgerErasureShredRecord) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvLedgerErasureShredRecord) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvLedgerErasureShredRecord) SetReason(v string)`

SetReason sets Reason field to given value.


### GetAuthority

`func (o *V2KvLedgerErasureShredRecord) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *V2KvLedgerErasureShredRecord) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *V2KvLedgerErasureShredRecord) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetRequestedBy

`func (o *V2KvLedgerErasureShredRecord) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *V2KvLedgerErasureShredRecord) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *V2KvLedgerErasureShredRecord) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.


### GetLedgerForked

`func (o *V2KvLedgerErasureShredRecord) GetLedgerForked() bool`

GetLedgerForked returns the LedgerForked field if non-nil, zero value otherwise.

### GetLedgerForkedOk

`func (o *V2KvLedgerErasureShredRecord) GetLedgerForkedOk() (*bool, bool)`

GetLedgerForkedOk returns a tuple with the LedgerForked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerForked

`func (o *V2KvLedgerErasureShredRecord) SetLedgerForked(v bool)`

SetLedgerForked sets LedgerForked field to given value.

### HasLedgerForked

`func (o *V2KvLedgerErasureShredRecord) HasLedgerForked() bool`

HasLedgerForked returns a boolean if a field has been set.

### GetSupersedes

`func (o *V2KvLedgerErasureShredRecord) GetSupersedes() string`

GetSupersedes returns the Supersedes field if non-nil, zero value otherwise.

### GetSupersedesOk

`func (o *V2KvLedgerErasureShredRecord) GetSupersedesOk() (*string, bool)`

GetSupersedesOk returns a tuple with the Supersedes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersedes

`func (o *V2KvLedgerErasureShredRecord) SetSupersedes(v string)`

SetSupersedes sets Supersedes field to given value.

### HasSupersedes

`func (o *V2KvLedgerErasureShredRecord) HasSupersedes() bool`

HasSupersedes returns a boolean if a field has been set.

### GetRefKeyVersion

`func (o *V2KvLedgerErasureShredRecord) GetRefKeyVersion() int32`

GetRefKeyVersion returns the RefKeyVersion field if non-nil, zero value otherwise.

### GetRefKeyVersionOk

`func (o *V2KvLedgerErasureShredRecord) GetRefKeyVersionOk() (*int32, bool)`

GetRefKeyVersionOk returns a tuple with the RefKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefKeyVersion

`func (o *V2KvLedgerErasureShredRecord) SetRefKeyVersion(v int32)`

SetRefKeyVersion sets RefKeyVersion field to given value.

### HasRefKeyVersion

`func (o *V2KvLedgerErasureShredRecord) HasRefKeyVersion() bool`

HasRefKeyVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


