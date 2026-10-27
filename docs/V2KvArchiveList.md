# V2KvArchiveList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]V2KvArchiveStatus**](V2KvArchiveStatus.md) |  | 
**NextCursor** | **NullableString** | Pass as cursor for the next page; null on the last page. | 
**HasMore** | **bool** |  | 
**PageEnd** | [**NullableV2KvPageEnd**](V2KvPageEnd.md) |  | 

## Methods

### NewV2KvArchiveList

`func NewV2KvArchiveList(data []V2KvArchiveStatus, nextCursor NullableString, hasMore bool, pageEnd NullableV2KvPageEnd, ) *V2KvArchiveList`

NewV2KvArchiveList instantiates a new V2KvArchiveList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvArchiveListWithDefaults

`func NewV2KvArchiveListWithDefaults() *V2KvArchiveList`

NewV2KvArchiveListWithDefaults instantiates a new V2KvArchiveList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2KvArchiveList) GetData() []V2KvArchiveStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2KvArchiveList) GetDataOk() (*[]V2KvArchiveStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2KvArchiveList) SetData(v []V2KvArchiveStatus)`

SetData sets Data field to given value.


### GetNextCursor

`func (o *V2KvArchiveList) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V2KvArchiveList) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V2KvArchiveList) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *V2KvArchiveList) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *V2KvArchiveList) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *V2KvArchiveList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *V2KvArchiveList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *V2KvArchiveList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetPageEnd

`func (o *V2KvArchiveList) GetPageEnd() V2KvPageEnd`

GetPageEnd returns the PageEnd field if non-nil, zero value otherwise.

### GetPageEndOk

`func (o *V2KvArchiveList) GetPageEndOk() (*V2KvPageEnd, bool)`

GetPageEndOk returns a tuple with the PageEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageEnd

`func (o *V2KvArchiveList) SetPageEnd(v V2KvPageEnd)`

SetPageEnd sets PageEnd field to given value.


### SetPageEndNil

`func (o *V2KvArchiveList) SetPageEndNil(b bool)`

 SetPageEndNil sets the value for PageEnd to be an explicit nil

### UnsetPageEnd
`func (o *V2KvArchiveList) UnsetPageEnd()`

UnsetPageEnd ensures that no value is present for PageEnd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


