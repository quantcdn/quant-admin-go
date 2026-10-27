# V2KvLedgerVerifyShard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** |  | 
**Shard** | **string** |  | 
**FromSeq** | **int32** | The first seq checked in this call (0 when no entry was in range). | 
**ToSeq** | **int32** |  | 
**Ok** | **bool** |  | 
**FirstBrokenSeq** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ObjectFromSeq** | Pointer to **int32** |  | [optional] 
**ObjectToSeq** | Pointer to **int32** |  | [optional] 

## Methods

### NewV2KvLedgerVerifyShard

`func NewV2KvLedgerVerifyShard(collection string, shard string, fromSeq int32, toSeq int32, ok bool, ) *V2KvLedgerVerifyShard`

NewV2KvLedgerVerifyShard instantiates a new V2KvLedgerVerifyShard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerVerifyShardWithDefaults

`func NewV2KvLedgerVerifyShardWithDefaults() *V2KvLedgerVerifyShard`

NewV2KvLedgerVerifyShardWithDefaults instantiates a new V2KvLedgerVerifyShard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *V2KvLedgerVerifyShard) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerVerifyShard) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerVerifyShard) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetShard

`func (o *V2KvLedgerVerifyShard) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvLedgerVerifyShard) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvLedgerVerifyShard) SetShard(v string)`

SetShard sets Shard field to given value.


### GetFromSeq

`func (o *V2KvLedgerVerifyShard) GetFromSeq() int32`

GetFromSeq returns the FromSeq field if non-nil, zero value otherwise.

### GetFromSeqOk

`func (o *V2KvLedgerVerifyShard) GetFromSeqOk() (*int32, bool)`

GetFromSeqOk returns a tuple with the FromSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSeq

`func (o *V2KvLedgerVerifyShard) SetFromSeq(v int32)`

SetFromSeq sets FromSeq field to given value.


### GetToSeq

`func (o *V2KvLedgerVerifyShard) GetToSeq() int32`

GetToSeq returns the ToSeq field if non-nil, zero value otherwise.

### GetToSeqOk

`func (o *V2KvLedgerVerifyShard) GetToSeqOk() (*int32, bool)`

GetToSeqOk returns a tuple with the ToSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSeq

`func (o *V2KvLedgerVerifyShard) SetToSeq(v int32)`

SetToSeq sets ToSeq field to given value.


### GetOk

`func (o *V2KvLedgerVerifyShard) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *V2KvLedgerVerifyShard) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *V2KvLedgerVerifyShard) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetFirstBrokenSeq

`func (o *V2KvLedgerVerifyShard) GetFirstBrokenSeq() int32`

GetFirstBrokenSeq returns the FirstBrokenSeq field if non-nil, zero value otherwise.

### GetFirstBrokenSeqOk

`func (o *V2KvLedgerVerifyShard) GetFirstBrokenSeqOk() (*int32, bool)`

GetFirstBrokenSeqOk returns a tuple with the FirstBrokenSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstBrokenSeq

`func (o *V2KvLedgerVerifyShard) SetFirstBrokenSeq(v int32)`

SetFirstBrokenSeq sets FirstBrokenSeq field to given value.

### HasFirstBrokenSeq

`func (o *V2KvLedgerVerifyShard) HasFirstBrokenSeq() bool`

HasFirstBrokenSeq returns a boolean if a field has been set.

### GetReason

`func (o *V2KvLedgerVerifyShard) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V2KvLedgerVerifyShard) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V2KvLedgerVerifyShard) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V2KvLedgerVerifyShard) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetObjectFromSeq

`func (o *V2KvLedgerVerifyShard) GetObjectFromSeq() int32`

GetObjectFromSeq returns the ObjectFromSeq field if non-nil, zero value otherwise.

### GetObjectFromSeqOk

`func (o *V2KvLedgerVerifyShard) GetObjectFromSeqOk() (*int32, bool)`

GetObjectFromSeqOk returns a tuple with the ObjectFromSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFromSeq

`func (o *V2KvLedgerVerifyShard) SetObjectFromSeq(v int32)`

SetObjectFromSeq sets ObjectFromSeq field to given value.

### HasObjectFromSeq

`func (o *V2KvLedgerVerifyShard) HasObjectFromSeq() bool`

HasObjectFromSeq returns a boolean if a field has been set.

### GetObjectToSeq

`func (o *V2KvLedgerVerifyShard) GetObjectToSeq() int32`

GetObjectToSeq returns the ObjectToSeq field if non-nil, zero value otherwise.

### GetObjectToSeqOk

`func (o *V2KvLedgerVerifyShard) GetObjectToSeqOk() (*int32, bool)`

GetObjectToSeqOk returns a tuple with the ObjectToSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectToSeq

`func (o *V2KvLedgerVerifyShard) SetObjectToSeq(v int32)`

SetObjectToSeq sets ObjectToSeq field to given value.

### HasObjectToSeq

`func (o *V2KvLedgerVerifyShard) HasObjectToSeq() bool`

HasObjectToSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


