# V2KvLedgerMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The ledger name. | [optional] 
**Id** | **string** | A ULID, a hyphen and the two-digit shard. | 
**Shard** | **string** |  | 
**Seq** | **int32** | The sequence number in the shard. | 
**At** | **int64** | Server time, epoch milliseconds. | 
**Hash** | **string** |  | 
**PrevHash** | **string** | The hash of the previous entry in the same shard. | 
**Source** | Pointer to [**V2KvLedgerProvenance**](V2KvLedgerProvenance.md) |  | [optional] 
**Subject** | Pointer to **string** | Present when the ledger has a subjectField: what the subject field holds. | [optional] 

## Methods

### NewV2KvLedgerMeta

`func NewV2KvLedgerMeta(id string, shard string, seq int32, at int64, hash string, prevHash string, ) *V2KvLedgerMeta`

NewV2KvLedgerMeta instantiates a new V2KvLedgerMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerMetaWithDefaults

`func NewV2KvLedgerMetaWithDefaults() *V2KvLedgerMeta`

NewV2KvLedgerMetaWithDefaults instantiates a new V2KvLedgerMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2KvLedgerMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvLedgerMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvLedgerMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2KvLedgerMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *V2KvLedgerMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2KvLedgerMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2KvLedgerMeta) SetId(v string)`

SetId sets Id field to given value.


### GetShard

`func (o *V2KvLedgerMeta) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvLedgerMeta) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvLedgerMeta) SetShard(v string)`

SetShard sets Shard field to given value.


### GetSeq

`func (o *V2KvLedgerMeta) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V2KvLedgerMeta) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V2KvLedgerMeta) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetAt

`func (o *V2KvLedgerMeta) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *V2KvLedgerMeta) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *V2KvLedgerMeta) SetAt(v int64)`

SetAt sets At field to given value.


### GetHash

`func (o *V2KvLedgerMeta) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *V2KvLedgerMeta) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *V2KvLedgerMeta) SetHash(v string)`

SetHash sets Hash field to given value.


### GetPrevHash

`func (o *V2KvLedgerMeta) GetPrevHash() string`

GetPrevHash returns the PrevHash field if non-nil, zero value otherwise.

### GetPrevHashOk

`func (o *V2KvLedgerMeta) GetPrevHashOk() (*string, bool)`

GetPrevHashOk returns a tuple with the PrevHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevHash

`func (o *V2KvLedgerMeta) SetPrevHash(v string)`

SetPrevHash sets PrevHash field to given value.


### GetSource

`func (o *V2KvLedgerMeta) GetSource() V2KvLedgerProvenance`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V2KvLedgerMeta) GetSourceOk() (*V2KvLedgerProvenance, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V2KvLedgerMeta) SetSource(v V2KvLedgerProvenance)`

SetSource sets Source field to given value.

### HasSource

`func (o *V2KvLedgerMeta) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubject

`func (o *V2KvLedgerMeta) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *V2KvLedgerMeta) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *V2KvLedgerMeta) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *V2KvLedgerMeta) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


