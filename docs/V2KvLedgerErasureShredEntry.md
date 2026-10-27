# V2KvLedgerErasureShredEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The entry key, &lt;collection&gt;:&lt;id&gt;. | 
**Shard** | **string** |  | 
**Seq** | **int32** |  | 
**Hash** | **string** |  | 

## Methods

### NewV2KvLedgerErasureShredEntry

`func NewV2KvLedgerErasureShredEntry(key string, shard string, seq int32, hash string, ) *V2KvLedgerErasureShredEntry`

NewV2KvLedgerErasureShredEntry instantiates a new V2KvLedgerErasureShredEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredEntryWithDefaults

`func NewV2KvLedgerErasureShredEntryWithDefaults() *V2KvLedgerErasureShredEntry`

NewV2KvLedgerErasureShredEntryWithDefaults instantiates a new V2KvLedgerErasureShredEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvLedgerErasureShredEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvLedgerErasureShredEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvLedgerErasureShredEntry) SetKey(v string)`

SetKey sets Key field to given value.


### GetShard

`func (o *V2KvLedgerErasureShredEntry) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvLedgerErasureShredEntry) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvLedgerErasureShredEntry) SetShard(v string)`

SetShard sets Shard field to given value.


### GetSeq

`func (o *V2KvLedgerErasureShredEntry) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V2KvLedgerErasureShredEntry) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V2KvLedgerErasureShredEntry) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetHash

`func (o *V2KvLedgerErasureShredEntry) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *V2KvLedgerErasureShredEntry) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *V2KvLedgerErasureShredEntry) SetHash(v string)`

SetHash sets Hash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


