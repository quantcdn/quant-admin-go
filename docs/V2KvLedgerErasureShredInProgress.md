# V2KvLedgerErasureShredInProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Ledger** | **string** |  | 
**SubjectRef** | **string** | The pseudonym of the subject. | 
**Entries** | [**[]V2KvLedgerReceipt**](V2KvLedgerReceipt.md) | The shred entries recorded so far. The final 200 lists them again. | 

## Methods

### NewV2KvLedgerErasureShredInProgress

`func NewV2KvLedgerErasureShredInProgress(status string, ledger string, subjectRef string, entries []V2KvLedgerReceipt, ) *V2KvLedgerErasureShredInProgress`

NewV2KvLedgerErasureShredInProgress instantiates a new V2KvLedgerErasureShredInProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredInProgressWithDefaults

`func NewV2KvLedgerErasureShredInProgressWithDefaults() *V2KvLedgerErasureShredInProgress`

NewV2KvLedgerErasureShredInProgressWithDefaults instantiates a new V2KvLedgerErasureShredInProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *V2KvLedgerErasureShredInProgress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V2KvLedgerErasureShredInProgress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V2KvLedgerErasureShredInProgress) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLedger

`func (o *V2KvLedgerErasureShredInProgress) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerErasureShredInProgress) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerErasureShredInProgress) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### GetSubjectRef

`func (o *V2KvLedgerErasureShredInProgress) GetSubjectRef() string`

GetSubjectRef returns the SubjectRef field if non-nil, zero value otherwise.

### GetSubjectRefOk

`func (o *V2KvLedgerErasureShredInProgress) GetSubjectRefOk() (*string, bool)`

GetSubjectRefOk returns a tuple with the SubjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRef

`func (o *V2KvLedgerErasureShredInProgress) SetSubjectRef(v string)`

SetSubjectRef sets SubjectRef field to given value.


### GetEntries

`func (o *V2KvLedgerErasureShredInProgress) GetEntries() []V2KvLedgerReceipt`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *V2KvLedgerErasureShredInProgress) GetEntriesOk() (*[]V2KvLedgerReceipt, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *V2KvLedgerErasureShredInProgress) SetEntries(v []V2KvLedgerReceipt)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


