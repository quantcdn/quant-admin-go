# V2KvLedgerReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | collection:id | 
**Id** | **string** | ULID, a hyphen and the two-digit shard. Increases within a shard. | 
**Shard** | **string** |  | 
**Seq** | **int32** | Sequence number of the entry on its shard. | 
**At** | **int64** | Epoch milliseconds (the ULID time of id). | 
**Hash** | **string** | SHA-256 hex of this entry. | 
**PrevHash** | **string** | SHA-256 hex of the previous entry on the shard (the genesis hash for seq 1). | 
**Ledger** | **string** | The ledger name. | 
**Source** | Pointer to [**V2KvLedgerProvenance**](V2KvLedgerProvenance.md) |  | [optional] 
**Replayed** | Pointer to **bool** | Present only on a retry with the same idempotencyKey: this is the first receipt. | [optional] 

## Methods

### NewV2KvLedgerReceipt

`func NewV2KvLedgerReceipt(key string, id string, shard string, seq int32, at int64, hash string, prevHash string, ledger string, ) *V2KvLedgerReceipt`

NewV2KvLedgerReceipt instantiates a new V2KvLedgerReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerReceiptWithDefaults

`func NewV2KvLedgerReceiptWithDefaults() *V2KvLedgerReceipt`

NewV2KvLedgerReceiptWithDefaults instantiates a new V2KvLedgerReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvLedgerReceipt) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvLedgerReceipt) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvLedgerReceipt) SetKey(v string)`

SetKey sets Key field to given value.


### GetId

`func (o *V2KvLedgerReceipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2KvLedgerReceipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2KvLedgerReceipt) SetId(v string)`

SetId sets Id field to given value.


### GetShard

`func (o *V2KvLedgerReceipt) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvLedgerReceipt) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvLedgerReceipt) SetShard(v string)`

SetShard sets Shard field to given value.


### GetSeq

`func (o *V2KvLedgerReceipt) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V2KvLedgerReceipt) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V2KvLedgerReceipt) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetAt

`func (o *V2KvLedgerReceipt) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *V2KvLedgerReceipt) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *V2KvLedgerReceipt) SetAt(v int64)`

SetAt sets At field to given value.


### GetHash

`func (o *V2KvLedgerReceipt) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *V2KvLedgerReceipt) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *V2KvLedgerReceipt) SetHash(v string)`

SetHash sets Hash field to given value.


### GetPrevHash

`func (o *V2KvLedgerReceipt) GetPrevHash() string`

GetPrevHash returns the PrevHash field if non-nil, zero value otherwise.

### GetPrevHashOk

`func (o *V2KvLedgerReceipt) GetPrevHashOk() (*string, bool)`

GetPrevHashOk returns a tuple with the PrevHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevHash

`func (o *V2KvLedgerReceipt) SetPrevHash(v string)`

SetPrevHash sets PrevHash field to given value.


### GetLedger

`func (o *V2KvLedgerReceipt) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerReceipt) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerReceipt) SetLedger(v string)`

SetLedger sets Ledger field to given value.


### GetSource

`func (o *V2KvLedgerReceipt) GetSource() V2KvLedgerProvenance`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V2KvLedgerReceipt) GetSourceOk() (*V2KvLedgerProvenance, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V2KvLedgerReceipt) SetSource(v V2KvLedgerProvenance)`

SetSource sets Source field to given value.

### HasSource

`func (o *V2KvLedgerReceipt) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetReplayed

`func (o *V2KvLedgerReceipt) GetReplayed() bool`

GetReplayed returns the Replayed field if non-nil, zero value otherwise.

### GetReplayedOk

`func (o *V2KvLedgerReceipt) GetReplayedOk() (*bool, bool)`

GetReplayedOk returns a tuple with the Replayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayed

`func (o *V2KvLedgerReceipt) SetReplayed(v bool)`

SetReplayed sets Replayed field to given value.

### HasReplayed

`func (o *V2KvLedgerReceipt) HasReplayed() bool`

HasReplayed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


