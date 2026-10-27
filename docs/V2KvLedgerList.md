# V2KvLedgerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledgers** | [**[]V2KvLedgerListItem**](V2KvLedgerListItem.md) | Every ledger of the store (all states); empty for a format 1 store. A store has at most 20 ledgers. | 

## Methods

### NewV2KvLedgerList

`func NewV2KvLedgerList(ledgers []V2KvLedgerListItem, ) *V2KvLedgerList`

NewV2KvLedgerList instantiates a new V2KvLedgerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerListWithDefaults

`func NewV2KvLedgerListWithDefaults() *V2KvLedgerList`

NewV2KvLedgerListWithDefaults instantiates a new V2KvLedgerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedgers

`func (o *V2KvLedgerList) GetLedgers() []V2KvLedgerListItem`

GetLedgers returns the Ledgers field if non-nil, zero value otherwise.

### GetLedgersOk

`func (o *V2KvLedgerList) GetLedgersOk() (*[]V2KvLedgerListItem, bool)`

GetLedgersOk returns a tuple with the Ledgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgers

`func (o *V2KvLedgerList) SetLedgers(v []V2KvLedgerListItem)`

SetLedgers sets Ledgers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


