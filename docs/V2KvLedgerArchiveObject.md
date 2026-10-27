# V2KvLedgerArchiveObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | ledgers/&lt;org&gt;/&lt;store&gt;/&lt;ledger&gt;/&lt;collection base64url&gt;/&lt;shard&gt;/&lt;period&gt;-&lt;fromSeq 12&gt;-&lt;toSeq 12&gt;.jsonl.gz. Save the file under this path. | 
**Url** | **string** | Presigned GET URL, valid for up to 3600 seconds. | 
**Collection** | **string** |  | 
**Shard** | **string** |  | 
**FromSeq** | **int32** |  | 
**ToSeq** | **int32** |  | 
**Period** | **string** |  | 

## Methods

### NewV2KvLedgerArchiveObject

`func NewV2KvLedgerArchiveObject(key string, url string, collection string, shard string, fromSeq int32, toSeq int32, period string, ) *V2KvLedgerArchiveObject`

NewV2KvLedgerArchiveObject instantiates a new V2KvLedgerArchiveObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerArchiveObjectWithDefaults

`func NewV2KvLedgerArchiveObjectWithDefaults() *V2KvLedgerArchiveObject`

NewV2KvLedgerArchiveObjectWithDefaults instantiates a new V2KvLedgerArchiveObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V2KvLedgerArchiveObject) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2KvLedgerArchiveObject) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2KvLedgerArchiveObject) SetKey(v string)`

SetKey sets Key field to given value.


### GetUrl

`func (o *V2KvLedgerArchiveObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2KvLedgerArchiveObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2KvLedgerArchiveObject) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCollection

`func (o *V2KvLedgerArchiveObject) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvLedgerArchiveObject) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvLedgerArchiveObject) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetShard

`func (o *V2KvLedgerArchiveObject) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *V2KvLedgerArchiveObject) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *V2KvLedgerArchiveObject) SetShard(v string)`

SetShard sets Shard field to given value.


### GetFromSeq

`func (o *V2KvLedgerArchiveObject) GetFromSeq() int32`

GetFromSeq returns the FromSeq field if non-nil, zero value otherwise.

### GetFromSeqOk

`func (o *V2KvLedgerArchiveObject) GetFromSeqOk() (*int32, bool)`

GetFromSeqOk returns a tuple with the FromSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSeq

`func (o *V2KvLedgerArchiveObject) SetFromSeq(v int32)`

SetFromSeq sets FromSeq field to given value.


### GetToSeq

`func (o *V2KvLedgerArchiveObject) GetToSeq() int32`

GetToSeq returns the ToSeq field if non-nil, zero value otherwise.

### GetToSeqOk

`func (o *V2KvLedgerArchiveObject) GetToSeqOk() (*int32, bool)`

GetToSeqOk returns a tuple with the ToSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSeq

`func (o *V2KvLedgerArchiveObject) SetToSeq(v int32)`

SetToSeq sets ToSeq field to given value.


### GetPeriod

`func (o *V2KvLedgerArchiveObject) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *V2KvLedgerArchiveObject) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *V2KvLedgerArchiveObject) SetPeriod(v string)`

SetPeriod sets Period field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


