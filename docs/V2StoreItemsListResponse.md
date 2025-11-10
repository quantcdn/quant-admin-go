# V2StoreItemsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]string** | List of item keys | [optional] 
**NextCursor** | Pointer to **NullableString** | Cursor for next page of results | [optional] 

## Methods

### NewV2StoreItemsListResponse

`func NewV2StoreItemsListResponse() *V2StoreItemsListResponse`

NewV2StoreItemsListResponse instantiates a new V2StoreItemsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2StoreItemsListResponseWithDefaults

`func NewV2StoreItemsListResponseWithDefaults() *V2StoreItemsListResponse`

NewV2StoreItemsListResponseWithDefaults instantiates a new V2StoreItemsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2StoreItemsListResponse) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2StoreItemsListResponse) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2StoreItemsListResponse) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *V2StoreItemsListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNextCursor

`func (o *V2StoreItemsListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *V2StoreItemsListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *V2StoreItemsListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *V2StoreItemsListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *V2StoreItemsListResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *V2StoreItemsListResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


