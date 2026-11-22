# CiCacheList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalBytes** | **int64** | Size of every entry that matches the filter (partial when totalsComplete is false). | 
**EntryCount** | **int64** | Number of entries that match the filter (partial when totalsComplete is false). | 
**TotalsComplete** | **bool** | false when the count stopped at its time limit (about 10 seconds) on a very large cache; totalBytes and entryCount then cover only the part counted. | 
**Entries** | [**[]CiCacheEntry**](CiCacheEntry.md) |  | 
**NextCursor** | **NullableString** | Opaque. Pass as cursor, with the same filter, for the next page; null on the last page. | 

## Methods

### NewCiCacheList

`func NewCiCacheList(totalBytes int64, entryCount int64, totalsComplete bool, entries []CiCacheEntry, nextCursor NullableString, ) *CiCacheList`

NewCiCacheList instantiates a new CiCacheList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiCacheListWithDefaults

`func NewCiCacheListWithDefaults() *CiCacheList`

NewCiCacheListWithDefaults instantiates a new CiCacheList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalBytes

`func (o *CiCacheList) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *CiCacheList) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *CiCacheList) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.


### GetEntryCount

`func (o *CiCacheList) GetEntryCount() int64`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *CiCacheList) GetEntryCountOk() (*int64, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *CiCacheList) SetEntryCount(v int64)`

SetEntryCount sets EntryCount field to given value.


### GetTotalsComplete

`func (o *CiCacheList) GetTotalsComplete() bool`

GetTotalsComplete returns the TotalsComplete field if non-nil, zero value otherwise.

### GetTotalsCompleteOk

`func (o *CiCacheList) GetTotalsCompleteOk() (*bool, bool)`

GetTotalsCompleteOk returns a tuple with the TotalsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalsComplete

`func (o *CiCacheList) SetTotalsComplete(v bool)`

SetTotalsComplete sets TotalsComplete field to given value.


### GetEntries

`func (o *CiCacheList) GetEntries() []CiCacheEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CiCacheList) GetEntriesOk() (*[]CiCacheEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CiCacheList) SetEntries(v []CiCacheEntry)`

SetEntries sets Entries field to given value.


### GetNextCursor

`func (o *CiCacheList) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CiCacheList) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CiCacheList) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *CiCacheList) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *CiCacheList) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


