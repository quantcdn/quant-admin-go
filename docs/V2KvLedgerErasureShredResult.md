# V2KvLedgerErasureShredResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledger** | **string** |  | 
**Entries** | [**[]V2KvLedgerReceipt**](V2KvLedgerReceipt.md) | The shred entry of each collection that holds the subject (empty for alreadyShredded, a DROPPED ledger and a forked store). | 
**ShreddedAt** | **time.Time** | When the key was destroyed (the first shred for alreadyShredded). | 
**FullyErasedAfter** | **time.Time** | shreddedAt + 35 days: after this time no backup of the key remains. | 
**AlreadyShredded** | Pointer to **bool** | Present only when the subject was shredded before this call and this call wrote nothing. | [optional] 
**LedgerForked** | Pointer to **bool** | Present only when the store is forked: the key is destroyed and no shred entry is recorded. | [optional] 
**ShredId** | Pointer to **string** | The id of the erasure record (the same on every repeat). | [optional] 
**SubjectRef** | **string** | The pseudonym of the subject in the erasure records and receipts. | 
**RefKeyVersion** | **int32** | The version of the reference key that made subjectRef. | 
**Receipt** | Pointer to [**V2KvLedgerErasureSignedShredReceipt**](V2KvLedgerErasureSignedShredReceipt.md) |  | [optional] 

## Methods

### NewV2KvLedgerErasureShredResult

`func NewV2KvLedgerErasureShredResult(ledger string, entries []V2KvLedgerReceipt, shreddedAt time.Time, fullyErasedAfter time.Time, subjectRef string, refKeyVersion int32, ) *V2KvLedgerErasureShredResult`

NewV2KvLedgerErasureShredResult instantiates a new V2KvLedgerErasureShredResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredResultWithDefaults

`func NewV2KvLedgerErasureShredResultWithDefaults() *V2KvLedgerErasureShredResult`

NewV2KvLedgerErasureShredResultWithDefaults instantiates a new V2KvLedgerErasureShredResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedger

`func (o *V2KvLedgerErasureShredResult) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerErasureShredResult) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerErasureShredResult) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### GetEntries

`func (o *V2KvLedgerErasureShredResult) GetEntries() []V2KvLedgerReceipt`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerErasureShredResult) GetEntriesOk() (*[]V2KvLedgerReceipt, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerErasureShredResult) SetEntries(v []V2KvLedgerReceipt)`

SetEntries sets Entries field to given value.


### GetShreddedAt

`func (o *V2KvLedgerErasureShredResult) GetShreddedAt() time.Time`

GetShreddedAt returns the ShreddedAt field if non-nil, zero value otherwise.

### GetShreddedAtOk

`func (o *V2KvLedgerErasureShredResult) GetShreddedAtOk() (*time.Time, bool)`

GetShreddedAtOk returns a tuple with the ShreddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShreddedAt

`func (o *V2KvLedgerErasureShredResult) SetShreddedAt(v time.Time)`

SetShreddedAt sets ShreddedAt field to given value.


### GetFullyErasedAfter

`func (o *V2KvLedgerErasureShredResult) GetFullyErasedAfter() time.Time`

GetFullyErasedAfter returns the FullyErasedAfter field if non-nil, zero value otherwise.

### GetFullyErasedAfterOk

`func (o *V2KvLedgerErasureShredResult) GetFullyErasedAfterOk() (*time.Time, bool)`

GetFullyErasedAfterOk returns a tuple with the FullyErasedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyErasedAfter

`func (o *V2KvLedgerErasureShredResult) SetFullyErasedAfter(v time.Time)`

SetFullyErasedAfter sets FullyErasedAfter field to given value.


### GetAlreadyShredded

`func (o *V2KvLedgerErasureShredResult) GetAlreadyShredded() bool`

GetAlreadyShredded returns the AlreadyShredded field if non-nil, zero value otherwise.

### GetAlreadyShreddedOk

`func (o *V2KvLedgerErasureShredResult) GetAlreadyShreddedOk() (*bool, bool)`

GetAlreadyShreddedOk returns a tuple with the AlreadyShredded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyShredded

`func (o *V2KvLedgerErasureShredResult) SetAlreadyShredded(v bool)`

SetAlreadyShredded sets AlreadyShredded field to given value.

### HasAlreadyShredded

`func (o *V2KvLedgerErasureShredResult) HasAlreadyShredded() bool`

HasAlreadyShredded returns a boolean if a field has been set.

### GetLedgerForked

`func (o *V2KvLedgerErasureShredResult) GetLedgerForked() bool`

GetLedgerForked returns the LedgerForked field if non-nil, zero value otherwise.

### GetLedgerForkedOk

`func (o *V2KvLedgerErasureShredResult) GetLedgerForkedOk() (*bool, bool)`

GetLedgerForkedOk returns a tuple with the LedgerForked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerForked

`func (o *V2KvLedgerErasureShredResult) SetLedgerForked(v bool)`

SetLedgerForked sets LedgerForked field to given value.

### HasLedgerForked

`func (o *V2KvLedgerErasureShredResult) HasLedgerForked() bool`

HasLedgerForked returns a boolean if a field has been set.

### GetShredId

`func (o *V2KvLedgerErasureShredResult) GetShredId() string`

GetShredId returns the ShredId field if non-nil, zero value otherwise.

### GetShredIdOk

`func (o *V2KvLedgerErasureShredResult) GetShredIdOk() (*string, bool)`

GetShredIdOk returns a tuple with the ShredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShredId

`func (o *V2KvLedgerErasureShredResult) SetShredId(v string)`

SetShredId sets ShredId field to given value.

### HasShredId

`func (o *V2KvLedgerErasureShredResult) HasShredId() bool`

HasShredId returns a boolean if a field has been set.

### GetSubjectRef

`func (o *V2KvLedgerErasureShredResult) GetSubjectRef() string`

GetSubjectRef returns the SubjectRef field if non-nil, zero value otherwise.

### GetSubjectRefOk

`func (o *V2KvLedgerErasureShredResult) GetSubjectRefOk() (*string, bool)`

GetSubjectRefOk returns a tuple with the SubjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRef

`func (o *V2KvLedgerErasureShredResult) SetSubjectRef(v string)`

SetSubjectRef sets SubjectRef field to given value.


### GetRefKeyVersion

`func (o *V2KvLedgerErasureShredResult) GetRefKeyVersion() int32`

GetRefKeyVersion returns the RefKeyVersion field if non-nil, zero value otherwise.

### GetRefKeyVersionOk

`func (o *V2KvLedgerErasureShredResult) GetRefKeyVersionOk() (*int32, bool)`

GetRefKeyVersionOk returns a tuple with the RefKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefKeyVersion

`func (o *V2KvLedgerErasureShredResult) SetRefKeyVersion(v int32)`

SetRefKeyVersion sets RefKeyVersion field to given value.


### GetReceipt

`func (o *V2KvLedgerErasureShredResult) GetReceipt() V2KvLedgerErasureSignedShredReceipt`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *V2KvLedgerErasureShredResult) GetReceiptOk() (*V2KvLedgerErasureSignedShredReceipt, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *V2KvLedgerErasureShredResult) SetReceipt(v V2KvLedgerErasureSignedShredReceipt)`

SetReceipt sets Receipt field to given value.

### HasReceipt

`func (o *V2KvLedgerErasureShredResult) HasReceipt() bool`

HasReceipt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


