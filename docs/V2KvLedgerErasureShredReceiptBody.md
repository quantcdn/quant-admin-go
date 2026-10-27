# V2KvLedgerErasureShredReceiptBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V** | **int32** |  | 
**Type** | **string** |  | 
**Id** | **string** | The shredId. | 
**Org** | **string** |  | 
**Store** | **string** |  | 
**Ledger** | **string** |  | 
**SubjectRef** | **string** |  | 
**RefKeyVersion** | **int32** | The version of the reference key that made subjectRef. | 
**ShreddedAt** | **time.Time** |  | 
**FullyErasedAfter** | **time.Time** |  | 
**Reason** | **string** |  | 
**Authority** | **string** |  | 
**RequestedBy** | **string** |  | 
**EntryCount** | **int32** | The number of shred entries (0 for a forked store or a DROPPED ledger). | 
**EntriesSha256** | **string** | SHA-256 (hex) of the JCS of the shred entries, sorted by key. | 
**Entries** | Pointer to [**[]V2KvLedgerErasureShredEntry**](V2KvLedgerErasureShredEntry.md) | Left out when the receipt would pass 300 KB: then KV_ledgers_shreds_entries serves them. | [optional] 
**LedgerForked** | Pointer to **bool** |  | [optional] 
**Supersedes** | Pointer to **string** | The id of the forked record that this re-shred supersedes. | [optional] 

## Methods

### NewV2KvLedgerErasureShredReceiptBody

`func NewV2KvLedgerErasureShredReceiptBody(v int32, type_ string, id string, org string, store string, ledger string, subjectRef string, refKeyVersion int32, shreddedAt time.Time, fullyErasedAfter time.Time, reason string, authority string, requestedBy string, entryCount int32, entriesSha256 string, ) *V2KvLedgerErasureShredReceiptBody`

NewV2KvLedgerErasureShredReceiptBody instantiates a new V2KvLedgerErasureShredReceiptBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredReceiptBodyWithDefaults

`func NewV2KvLedgerErasureShredReceiptBodyWithDefaults() *V2KvLedgerErasureShredReceiptBody`

NewV2KvLedgerErasureShredReceiptBodyWithDefaults instantiates a new V2KvLedgerErasureShredReceiptBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV

`func (o *V2KvLedgerErasureShredReceiptBody) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *V2KvLedgerErasureShredReceiptBody) SetV(v int32)`

SetV sets V field to given value.


### GetType

`func (o *V2KvLedgerErasureShredReceiptBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2KvLedgerErasureShredReceiptBody) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *V2KvLedgerErasureShredReceiptBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2KvLedgerErasureShredReceiptBody) SetId(v string)`

SetId sets Id field to given value.


### GetOrg

`func (o *V2KvLedgerErasureShredReceiptBody) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *V2KvLedgerErasureShredReceiptBody) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetStore

`func (o *V2KvLedgerErasureShredReceiptBody) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *V2KvLedgerErasureShredReceiptBody) SetStore(v string)`

SetStore sets Store field to given value.


### GetLedger

`func (o *V2KvLedgerErasureShredReceiptBody) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerErasureShredReceiptBody) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### GetSubjectRef

`func (o *V2KvLedgerErasureShredReceiptBody) GetSubjectRef() string`

GetSubjectRef returns the SubjectRef field if non-nil, zero value otherwise.

### GetSubjectRefOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetSubjectRefOk() (*string, bool)`

GetSubjectRefOk returns a tuple with the SubjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRef

`func (o *V2KvLedgerErasureShredReceiptBody) SetSubjectRef(v string)`

SetSubjectRef sets SubjectRef field to given value.


### GetRefKeyVersion

`func (o *V2KvLedgerErasureShredReceiptBody) GetRefKeyVersion() int32`

GetRefKeyVersion returns the RefKeyVersion field if non-nil, zero value otherwise.

### GetRefKeyVersionOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetRefKeyVersionOk() (*int32, bool)`

GetRefKeyVersionOk returns a tuple with the RefKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefKeyVersion

`func (o *V2KvLedgerErasureShredReceiptBody) SetRefKeyVersion(v int32)`

SetRefKeyVersion sets RefKeyVersion field to given value.


### GetShreddedAt

`func (o *V2KvLedgerErasureShredReceiptBody) GetShreddedAt() time.Time`

GetShreddedAt returns the ShreddedAt field if non-nil, zero value otherwise.

### GetShreddedAtOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetShreddedAtOk() (*time.Time, bool)`

GetShreddedAtOk returns a tuple with the ShreddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShreddedAt

`func (o *V2KvLedgerErasureShredReceiptBody) SetShreddedAt(v time.Time)`

SetShreddedAt sets ShreddedAt field to given value.


### GetFullyErasedAfter

`func (o *V2KvLedgerErasureShredReceiptBody) GetFullyErasedAfter() time.Time`

GetFullyErasedAfter returns the FullyErasedAfter field if non-nil, zero value otherwise.

### GetFullyErasedAfterOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetFullyErasedAfterOk() (*time.Time, bool)`

GetFullyErasedAfterOk returns a tuple with the FullyErasedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyErasedAfter

`func (o *V2KvLedgerErasureShredReceiptBody) SetFullyErasedAfter(v time.Time)`

SetFullyErasedAfter sets FullyErasedAfter field to given value.


### GetReason

`func (o *V2KvLedgerErasureShredReceiptBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvLedgerErasureShredReceiptBody) SetReason(v string)`

SetReason sets Reason field to given value.


### GetAuthority

`func (o *V2KvLedgerErasureShredReceiptBody) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *V2KvLedgerErasureShredReceiptBody) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetRequestedBy

`func (o *V2KvLedgerErasureShredReceiptBody) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *V2KvLedgerErasureShredReceiptBody) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.


### GetEntryCount

`func (o *V2KvLedgerErasureShredReceiptBody) GetEntryCount() int32`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetEntryCountOk() (*int32, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *V2KvLedgerErasureShredReceiptBody) SetEntryCount(v int32)`

SetEntryCount sets EntryCount field to given value.


### GetEntriesSha256

`func (o *V2KvLedgerErasureShredReceiptBody) GetEntriesSha256() string`

GetEntriesSha256 returns the EntriesSha256 field if non-nil, zero value otherwise.

### GetEntriesSha256Ok

`func (o *V2KvLedgerErasureShredReceiptBody) GetEntriesSha256Ok() (*string, bool)`

GetEntriesSha256Ok returns a tuple with the EntriesSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesSha256

`func (o *V2KvLedgerErasureShredReceiptBody) SetEntriesSha256(v string)`

SetEntriesSha256 sets EntriesSha256 field to given value.


### GetEntries

`func (o *V2KvLedgerErasureShredReceiptBody) GetEntries() []V2KvLedgerErasureShredEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetEntriesOk() (*[]V2KvLedgerErasureShredEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerErasureShredReceiptBody) SetEntries(v []V2KvLedgerErasureShredEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *V2KvLedgerErasureShredReceiptBody) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetLedgerForked

`func (o *V2KvLedgerErasureShredReceiptBody) GetLedgerForked() bool`

GetLedgerForked returns the LedgerForked field if non-nil, zero value otherwise.

### GetLedgerForkedOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetLedgerForkedOk() (*bool, bool)`

GetLedgerForkedOk returns a tuple with the LedgerForked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerForked

`func (o *V2KvLedgerErasureShredReceiptBody) SetLedgerForked(v bool)`

SetLedgerForked sets LedgerForked field to given value.

### HasLedgerForked

`func (o *V2KvLedgerErasureShredReceiptBody) HasLedgerForked() bool`

HasLedgerForked returns a boolean if a field has been set.

### GetSupersedes

`func (o *V2KvLedgerErasureShredReceiptBody) GetSupersedes() string`

GetSupersedes returns the Supersedes field if non-nil, zero value otherwise.

### GetSupersedesOk

`func (o *V2KvLedgerErasureShredReceiptBody) GetSupersedesOk() (*string, bool)`

GetSupersedesOk returns a tuple with the Supersedes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersedes

`func (o *V2KvLedgerErasureShredReceiptBody) SetSupersedes(v string)`

SetSupersedes sets Supersedes field to given value.

### HasSupersedes

`func (o *V2KvLedgerErasureShredReceiptBody) HasSupersedes() bool`

HasSupersedes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


