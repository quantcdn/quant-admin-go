# V2KvLedgerErasureShredEntriesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]V2KvLedgerErasureShredEntry**](V2KvLedgerErasureShredEntry.md) | Sorted by key. | 
**NextCursor** | **NullableString** | Pass as cursor for the next page. Null on the last page. | 
**HasMore** | **bool** |  | 
**PageEnd** | [**NullableV2KvPageEnd**](V2KvPageEnd.md) |  | 
**ShredId** | **string** |  | 
**EntryCount** | **int32** |  | 

## Methods

### NewV2KvLedgerErasureShredEntriesPage

`func NewV2KvLedgerErasureShredEntriesPage(data []V2KvLedgerErasureShredEntry, nextCursor NullableString, hasMore bool, pageEnd NullableV2KvPageEnd, shredId string, entryCount int32, ) *V2KvLedgerErasureShredEntriesPage`

NewV2KvLedgerErasureShredEntriesPage instantiates a new V2KvLedgerErasureShredEntriesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredEntriesPageWithDefaults

`func NewV2KvLedgerErasureShredEntriesPageWithDefaults() *V2KvLedgerErasureShredEntriesPage`

NewV2KvLedgerErasureShredEntriesPageWithDefaults instantiates a new V2KvLedgerErasureShredEntriesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2KvLedgerErasureShredEntriesPage) GetData() []V2KvLedgerErasureShredEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2KvLedgerErasureShredEntriesPage) GetDataOk() (*[]V2KvLedgerErasureShredEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2KvLedgerErasureShredEntriesPage) SetData(v []V2KvLedgerErasureShredEntry)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *V2KvLedgerErasureShredEntriesPage) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V2KvLedgerErasureShredEntriesPage) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V2KvLedgerErasureShredEntriesPage) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *V2KvLedgerErasureShredEntriesPage) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *V2KvLedgerErasureShredEntriesPage) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *V2KvLedgerErasureShredEntriesPage) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *V2KvLedgerErasureShredEntriesPage) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *V2KvLedgerErasureShredEntriesPage) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageEnd

`func (o *V2KvLedgerErasureShredEntriesPage) GetPageEnd() V2KvPageEnd`

GetPageEnd returns the PageEnd field if non-nil, zero value otherwise.

### GetPageEndOk

`func (o *V2KvLedgerErasureShredEntriesPage) GetPageEndOk() (*V2KvPageEnd, bool)`

GetPageEndOk returns a tuple with the PageEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageEnd

`func (o *V2KvLedgerErasureShredEntriesPage) SetPageEnd(v V2KvPageEnd)`

SetPageEnd sets PageEnd field to given value.


### SetPageEndNil

`func (o *V2KvLedgerErasureShredEntriesPage) SetPageEndNil(b bool)`

 SetPageEndNil sets the value for PageEnd to be an explicit nil

### UnsetPageEnd
`func (o *V2KvLedgerErasureShredEntriesPage) UnsetPageEnd()`

UnsetPageEnd ensures that no value is present for PageEnd, not even an explicit nil
### GetShredId

`func (o *V2KvLedgerErasureShredEntriesPage) GetShredId() string`

GetShredId returns the ShredId field if non-nil, zero value otherwise.

### GetShredIdOk

`func (o *V2KvLedgerErasureShredEntriesPage) GetShredIdOk() (*string, bool)`

GetShredIdOk returns a tuple with the ShredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShredId

`func (o *V2KvLedgerErasureShredEntriesPage) SetShredId(v string)`

SetShredId sets ShredId field to given value.


### GetEntryCount

`func (o *V2KvLedgerErasureShredEntriesPage) GetEntryCount() int32`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *V2KvLedgerErasureShredEntriesPage) GetEntryCountOk() (*int32, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *V2KvLedgerErasureShredEntriesPage) SetEntryCount(v int32)`

SetEntryCount sets EntryCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


