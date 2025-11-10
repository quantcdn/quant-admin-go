# ListBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backups** | Pointer to [**[]ListBackups200ResponseBackupsInner**](ListBackups200ResponseBackupsInner.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**NextToken** | Pointer to **NullableString** | Token for retrieving the next page of results, if more data is available | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewListBackups200Response

`func NewListBackups200Response() *ListBackups200Response`

NewListBackups200Response instantiates a new ListBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackups200ResponseWithDefaults

`func NewListBackups200ResponseWithDefaults() *ListBackups200Response`

NewListBackups200ResponseWithDefaults instantiates a new ListBackups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackups

`func (o *ListBackups200Response) GetBackups() []ListBackups200ResponseBackupsInner`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *ListBackups200Response) GetBackupsOk() (*[]ListBackups200ResponseBackupsInner, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *ListBackups200Response) SetBackups(v []ListBackups200ResponseBackupsInner)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *ListBackups200Response) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetCount

`func (o *ListBackups200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListBackups200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListBackups200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListBackups200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNextToken

`func (o *ListBackups200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListBackups200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListBackups200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListBackups200Response) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### SetNextTokenNil

`func (o *ListBackups200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *ListBackups200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil
### GetMessage

`func (o *ListBackups200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListBackups200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListBackups200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListBackups200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


