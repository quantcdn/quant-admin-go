# ListFiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]ListFiles200ResponseFilesInner**](ListFiles200ResponseFilesInner.md) |  | [optional] 
**NextCursor** | Pointer to **NullableString** | Cursor for next page | [optional] 
**HasMore** | Pointer to **bool** | True if more results available | [optional] 

## Methods

### NewListFiles200Response

`func NewListFiles200Response() *ListFiles200Response`

NewListFiles200Response instantiates a new ListFiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFiles200ResponseWithDefaults

`func NewListFiles200ResponseWithDefaults() *ListFiles200Response`

NewListFiles200ResponseWithDefaults instantiates a new ListFiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ListFiles200Response) GetFiles() []ListFiles200ResponseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ListFiles200Response) GetFilesOk() (*[]ListFiles200ResponseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ListFiles200Response) SetFiles(v []ListFiles200ResponseFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ListFiles200Response) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetNextCursor

`func (o *ListFiles200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListFiles200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListFiles200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListFiles200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ListFiles200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListFiles200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *ListFiles200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListFiles200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListFiles200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListFiles200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


