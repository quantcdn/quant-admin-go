# V2KvLedgerFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** |  | 
**Shard** | **string** |  | 
**Seq** | Pointer to **int32** |  | [optional] 
**Reason** | **string** |  | 

## Methods

### NewV2KvLedgerFailure

`func NewV2KvLedgerFailure(collection string, shard string, reason string, ) *V2KvLedgerFailure`

NewV2KvLedgerFailure instantiates a new V2KvLedgerFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerFailureWithDefaults

`func NewV2KvLedgerFailureWithDefaults() *V2KvLedgerFailure`

NewV2KvLedgerFailureWithDefaults instantiates a new V2KvLedgerFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *V2KvLedgerFailure) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerFailure) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerFailure) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetShard

`func (o *V2KvLedgerFailure) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvLedgerFailure) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvLedgerFailure) SetShard(v string)`

SetShard sets Shard field to given value.


### GetSeq

`func (o *V2KvLedgerFailure) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *V2KvLedgerFailure) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *V2KvLedgerFailure) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *V2KvLedgerFailure) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetReason

`func (o *V2KvLedgerFailure) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvLedgerFailure) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvLedgerFailure) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


