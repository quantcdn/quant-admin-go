# V2KvLedgerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** |  | 
**Shard** | **string** |  | 
**HeadSeq** | **int32** | Sequence number of the last entry. | 
**HeadAt** | Pointer to **int64** | Epoch ms of the last entry. | [optional] 
**ArchivedSeq** | **int32** | Last sequence number in the archive. | 
**ArchiveLagSeconds** | **int32** | Age of the oldest unarchived entry; 0 when up to date. | 
**Broken** | Pointer to **map[string]interface{}** | Set when the archive worker found a chain fault (seq, reason, at). | [optional] 

## Methods

### NewV2KvLedgerInstance

`func NewV2KvLedgerInstance(collection string, shard string, headSeq int32, archivedSeq int32, archiveLagSeconds int32, ) *V2KvLedgerInstance`

NewV2KvLedgerInstance instantiates a new V2KvLedgerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerInstanceWithDefaults

`func NewV2KvLedgerInstanceWithDefaults() *V2KvLedgerInstance`

NewV2KvLedgerInstanceWithDefaults instantiates a new V2KvLedgerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *V2KvLedgerInstance) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerInstance) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerInstance) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetShard

`func (o *V2KvLedgerInstance) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvLedgerInstance) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvLedgerInstance) SetShard(v string)`

SetShard sets Shard field to given value.


### GetHeadSeq

`func (o *V2KvLedgerInstance) GetHeadSeq() int32`

GetHeadSeq returns the HeadSeq field if non-nil, zero value otherwise.

### GetHeadSeqOk

`func (o *V2KvLedgerInstance) GetHeadSeqOk() (*int32, bool)`

GetHeadSeqOk returns a tuple with the HeadSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSeq

`func (o *V2KvLedgerInstance) SetHeadSeq(v int32)`

SetHeadSeq sets HeadSeq field to given value.


### GetHeadAt

`func (o *V2KvLedgerInstance) GetHeadAt() int64`

GetHeadAt returns the HeadAt field if non-nil, zero value otherwise.

### GetHeadAtOk

`func (o *V2KvLedgerInstance) GetHeadAtOk() (*int64, bool)`

GetHeadAtOk returns a tuple with the HeadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadAt

`func (o *V2KvLedgerInstance) SetHeadAt(v int64)`

SetHeadAt sets HeadAt field to given value.

### HasHeadAt

`func (o *V2KvLedgerInstance) HasHeadAt() bool`

HasHeadAt returns a boolean if a field has been set.

### GetArchivedSeq

`func (o *V2KvLedgerInstance) GetArchivedSeq() int32`

GetArchivedSeq returns the ArchivedSeq field if non-nil, zero value otherwise.

### GetArchivedSeqOk

`func (o *V2KvLedgerInstance) GetArchivedSeqOk() (*int32, bool)`

GetArchivedSeqOk returns a tuple with the ArchivedSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedSeq

`func (o *V2KvLedgerInstance) SetArchivedSeq(v int32)`

SetArchivedSeq sets ArchivedSeq field to given value.


### GetArchiveLagSeconds

`func (o *V2KvLedgerInstance) GetArchiveLagSeconds() int32`

GetArchiveLagSeconds returns the ArchiveLagSeconds field if non-nil, zero value otherwise.

### GetArchiveLagSecondsOk

`func (o *V2KvLedgerInstance) GetArchiveLagSecondsOk() (*int32, bool)`

GetArchiveLagSecondsOk returns a tuple with the ArchiveLagSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLagSeconds

`func (o *V2KvLedgerInstance) SetArchiveLagSeconds(v int32)`

SetArchiveLagSeconds sets ArchiveLagSeconds field to given value.


### GetBroken

`func (o *V2KvLedgerInstance) GetBroken() map[string]interface{}`

GetBroken returns the Broken field if non-nil, zero value otherwise.

### GetBrokenOk

`func (o *V2KvLedgerInstance) GetBrokenOk() (*map[string]interface{}, bool)`

GetBrokenOk returns a tuple with the Broken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroken

`func (o *V2KvLedgerInstance) SetBroken(v map[string]interface{})`

SetBroken sets Broken field to given value.

### HasBroken

`func (o *V2KvLedgerInstance) HasBroken() bool`

HasBroken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


