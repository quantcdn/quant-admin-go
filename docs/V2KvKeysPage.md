# V2KvKeysPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]V2KvEntry**](V2KvEntry.md) |  | 
**NextCursor** | **NullableString** | Signed; valid 24 hours; bound to the store and the query. Null on the last page. | 
**HasMore** | **bool** |  | 
**PageEnd** | [**NullableV2KvPageEnd**](V2KvPageEnd.md) |  | 
**Indexed** | Pointer to **bool** | Present when collection or index was sent. True for an indexed query on a format 2 store. | [optional] 
**Ledger** | Pointer to **string** | Present on a ledger list (the collection of a ledger): the ledger name. | [optional] 
**HiddenLedgerCollections** | Pointer to **int32** | Present when ledger rows were left out because the role does not have ledger_read: the number of ledger collections hidden from this page. | [optional] 
**Notice** | Pointer to **string** | Present with hidden_ledger_collections. | [optional] 

## Methods

### NewV2KvKeysPage

`func NewV2KvKeysPage(data []V2KvEntry, nextCursor NullableString, hasMore bool, pageEnd NullableV2KvPageEnd, ) *V2KvKeysPage`

NewV2KvKeysPage instantiates a new V2KvKeysPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvKeysPageWithDefaults

`func NewV2KvKeysPageWithDefaults() *V2KvKeysPage`

NewV2KvKeysPageWithDefaults instantiates a new V2KvKeysPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2KvKeysPage) GetData() []V2KvEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2KvKeysPage) GetDataOk() (*[]V2KvEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2KvKeysPage) SetData(v []V2KvEntry)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *V2KvKeysPage) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V2KvKeysPage) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V2KvKeysPage) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *V2KvKeysPage) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *V2KvKeysPage) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *V2KvKeysPage) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *V2KvKeysPage) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *V2KvKeysPage) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageEnd

`func (o *V2KvKeysPage) GetPageEnd() V2KvPageEnd`

GetPageEnd returns the PageEnd field if non-nil, zero value otherwise.

### GetPageEndOk

`func (o *V2KvKeysPage) GetPageEndOk() (*V2KvPageEnd, bool)`

GetPageEndOk returns a tuple with the PageEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageEnd

`func (o *V2KvKeysPage) SetPageEnd(v V2KvPageEnd)`

SetPageEnd sets PageEnd field to given value.


### SetPageEndNil

`func (o *V2KvKeysPage) SetPageEndNil(b bool)`

 SetPageEndNil sets the value for PageEnd to be an explicit nil

### UnsetPageEnd
`func (o *V2KvKeysPage) UnsetPageEnd()`

UnsetPageEnd ensures that no value is present for PageEnd, not even an explicit nil
### GetIndexed

`func (o *V2KvKeysPage) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *V2KvKeysPage) GetIndexedOk() (*bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexed

`func (o *V2KvKeysPage) SetIndexed(v bool)`

SetIndexed sets Indexed field to given value.

### HasIndexed

`func (o *V2KvKeysPage) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### GetLedger

`func (o *V2KvKeysPage) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvKeysPage) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvKeysPage) SetLedger(v string)`

SetLedger sets Ledger field to given value.

### HasLedger

`func (o *V2KvKeysPage) HasLedger() bool`

HasLedger returns a boolean if a field has been set.

### GetHiddenLedgerCollections

`func (o *V2KvKeysPage) GetHiddenLedgerCollections() int32`

GetHiddenLedgerCollections returns the HiddenLedgerCollections field if non-nil, zero value otherwise.

### GetHiddenLedgerCollectionsOk

`func (o *V2KvKeysPage) GetHiddenLedgerCollectionsOk() (*int32, bool)`

GetHiddenLedgerCollectionsOk returns a tuple with the HiddenLedgerCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenLedgerCollections

`func (o *V2KvKeysPage) SetHiddenLedgerCollections(v int32)`

SetHiddenLedgerCollections sets HiddenLedgerCollections field to given value.

### HasHiddenLedgerCollections

`func (o *V2KvKeysPage) HasHiddenLedgerCollections() bool`

HasHiddenLedgerCollections returns a boolean if a field has been set.

### GetNotice

`func (o *V2KvKeysPage) GetNotice() string`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *V2KvKeysPage) GetNoticeOk() (*string, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *V2KvKeysPage) SetNotice(v string)`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *V2KvKeysPage) HasNotice() bool`

HasNotice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


