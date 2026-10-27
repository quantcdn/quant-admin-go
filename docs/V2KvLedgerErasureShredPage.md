# V2KvLedgerErasureShredPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]V2KvLedgerErasureShredRecord**](V2KvLedgerErasureShredRecord.md) | Newest first, at most 50. | 
**NextCursor** | **NullableString** | Pass as cursor for the next page. Null on the last page. | 
**HasMore** | **bool** |  | 
**PageEnd** | [**NullableV2KvPageEnd**](V2KvPageEnd.md) |  | 
**Ledger** | **string** |  | 

## Methods

### NewV2KvLedgerErasureShredPage

`func NewV2KvLedgerErasureShredPage(data []V2KvLedgerErasureShredRecord, nextCursor NullableString, hasMore bool, pageEnd NullableV2KvPageEnd, ledger string, ) *V2KvLedgerErasureShredPage`

NewV2KvLedgerErasureShredPage instantiates a new V2KvLedgerErasureShredPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerErasureShredPageWithDefaults

`func NewV2KvLedgerErasureShredPageWithDefaults() *V2KvLedgerErasureShredPage`

NewV2KvLedgerErasureShredPageWithDefaults instantiates a new V2KvLedgerErasureShredPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2KvLedgerErasureShredPage) GetData() []V2KvLedgerErasureShredRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2KvLedgerErasureShredPage) GetDataOk() (*[]V2KvLedgerErasureShredRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2KvLedgerErasureShredPage) SetData(v []V2KvLedgerErasureShredRecord)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *V2KvLedgerErasureShredPage) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V2KvLedgerErasureShredPage) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V2KvLedgerErasureShredPage) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *V2KvLedgerErasureShredPage) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *V2KvLedgerErasureShredPage) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *V2KvLedgerErasureShredPage) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *V2KvLedgerErasureShredPage) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *V2KvLedgerErasureShredPage) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageEnd

`func (o *V2KvLedgerErasureShredPage) GetPageEnd() V2KvPageEnd`

GetPageEnd returns the PageEnd field if non-nil, zero value otherwise.

### GetPageEndOk

`func (o *V2KvLedgerErasureShredPage) GetPageEndOk() (*V2KvPageEnd, bool)`

GetPageEndOk returns a tuple with the PageEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageEnd

`func (o *V2KvLedgerErasureShredPage) SetPageEnd(v V2KvPageEnd)`

SetPageEnd sets PageEnd field to given value.


### SetPageEndNil

`func (o *V2KvLedgerErasureShredPage) SetPageEndNil(b bool)`

 SetPageEndNil sets the value for PageEnd to be an explicit nil

### UnsetPageEnd
`func (o *V2KvLedgerErasureShredPage) UnsetPageEnd()`

UnsetPageEnd ensures that no value is present for PageEnd, not even an explicit nil
### GetLedger

`func (o *V2KvLedgerErasureShredPage) GetLedger() string`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *V2KvLedgerErasureShredPage) GetLedgerOk() (*string, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *V2KvLedgerErasureShredPage) SetLedger(v string)`

SetLedger sets Ledger field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


