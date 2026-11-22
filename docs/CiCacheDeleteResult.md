# CiCacheDeleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedEntries** | **int64** |  | 
**DeletedBytes** | **int64** |  | 
**Complete** | **bool** | false when the delete stopped at its time limit (about 12 seconds), a storage call failed or timed out, or storage refused some objects. false means: send the same request again until it is true. If a call returns false with deletedEntries 0, stop and try again later, because a storage call failed, timed out or refused objects. | 

## Methods

### NewCiCacheDeleteResult

`func NewCiCacheDeleteResult(deletedEntries int64, deletedBytes int64, complete bool, ) *CiCacheDeleteResult`

NewCiCacheDeleteResult instantiates a new CiCacheDeleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiCacheDeleteResultWithDefaults

`func NewCiCacheDeleteResultWithDefaults() *CiCacheDeleteResult`

NewCiCacheDeleteResultWithDefaults instantiates a new CiCacheDeleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedEntries

`func (o *CiCacheDeleteResult) GetDeletedEntries() int64`

GetDeletedEntries returns the DeletedEntries field if non-nil, zero value otherwise.

### GetDeletedEntriesOk

`func (o *CiCacheDeleteResult) GetDeletedEntriesOk() (*int64, bool)`

GetDeletedEntriesOk returns a tuple with the DeletedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedEntries

`func (o *CiCacheDeleteResult) SetDeletedEntries(v int64)`

SetDeletedEntries sets DeletedEntries field to given value.


### GetDeletedBytes

`func (o *CiCacheDeleteResult) GetDeletedBytes() int64`

GetDeletedBytes returns the DeletedBytes field if non-nil, zero value otherwise.

### GetDeletedBytesOk

`func (o *CiCacheDeleteResult) GetDeletedBytesOk() (*int64, bool)`

GetDeletedBytesOk returns a tuple with the DeletedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBytes

`func (o *CiCacheDeleteResult) SetDeletedBytes(v int64)`

SetDeletedBytes sets DeletedBytes field to given value.


### GetComplete

`func (o *CiCacheDeleteResult) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *CiCacheDeleteResult) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *CiCacheDeleteResult) SetComplete(v bool)`

SetComplete sets Complete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


