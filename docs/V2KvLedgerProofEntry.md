# V2KvLedgerProofEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Collection** | Pointer to **string** |  | [optional] 
**Shard** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int32** |  | [optional] 
**At** | Pointer to **int64** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**PrevHash** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewV2KvLedgerProofEntry

`func NewV2KvLedgerProofEntry() *V2KvLedgerProofEntry`

NewV2KvLedgerProofEntry instantiates a new V2KvLedgerProofEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerProofEntryWithDefaults

`func NewV2KvLedgerProofEntryWithDefaults() *V2KvLedgerProofEntry`

NewV2KvLedgerProofEntryWithDefaults instantiates a new V2KvLedgerProofEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvLedgerProofEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvLedgerProofEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvLedgerProofEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *V2KvLedgerProofEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCollection

`func (o *V2KvLedgerProofEntry) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerProofEntry) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerProofEntry) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *V2KvLedgerProofEntry) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetShard

`func (o *V2KvLedgerProofEntry) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvLedgerProofEntry) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvLedgerProofEntry) SetShard(v string)`

SetShard sets Shard field to given value.

### HasShard

`func (o *V2KvLedgerProofEntry) HasShard() bool`

HasShard returns a boolean if a field has been set.

### GetId

`func (o *V2KvLedgerProofEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2KvLedgerProofEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2KvLedgerProofEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V2KvLedgerProofEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSeq

`func (o *V2KvLedgerProofEntry) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V2KvLedgerProofEntry) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V2KvLedgerProofEntry) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *V2KvLedgerProofEntry) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetAt

`func (o *V2KvLedgerProofEntry) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *V2KvLedgerProofEntry) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *V2KvLedgerProofEntry) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *V2KvLedgerProofEntry) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetValue

`func (o *V2KvLedgerProofEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2KvLedgerProofEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2KvLedgerProofEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V2KvLedgerProofEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPrevHash

`func (o *V2KvLedgerProofEntry) GetPrevHash() string`

GetPrevHash returns the PrevHash field if non-nil, zero value otherwise.

### GetPrevHashOk

`func (o *V2KvLedgerProofEntry) GetPrevHashOk() (*string, bool)`

GetPrevHashOk returns a tuple with the PrevHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevHash

`func (o *V2KvLedgerProofEntry) SetPrevHash(v string)`

SetPrevHash sets PrevHash field to given value.

### HasPrevHash

`func (o *V2KvLedgerProofEntry) HasPrevHash() bool`

HasPrevHash returns a boolean if a field has been set.

### GetHash

`func (o *V2KvLedgerProofEntry) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *V2KvLedgerProofEntry) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *V2KvLedgerProofEntry) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *V2KvLedgerProofEntry) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


