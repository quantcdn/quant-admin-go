# GetProjectEdgeEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]GetProjectEdgeEvents200ResponseEventsInner**](GetProjectEdgeEvents200ResponseEventsInner.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**NextToken** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetProjectEdgeEvents200Response

`func NewGetProjectEdgeEvents200Response() *GetProjectEdgeEvents200Response`

NewGetProjectEdgeEvents200Response instantiates a new GetProjectEdgeEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectEdgeEvents200ResponseWithDefaults

`func NewGetProjectEdgeEvents200ResponseWithDefaults() *GetProjectEdgeEvents200Response`

NewGetProjectEdgeEvents200ResponseWithDefaults instantiates a new GetProjectEdgeEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *GetProjectEdgeEvents200Response) GetEvents() []GetProjectEdgeEvents200ResponseEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetProjectEdgeEvents200Response) GetEventsOk() (*[]GetProjectEdgeEvents200ResponseEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetProjectEdgeEvents200Response) SetEvents(v []GetProjectEdgeEvents200ResponseEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetProjectEdgeEvents200Response) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetCount

`func (o *GetProjectEdgeEvents200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetProjectEdgeEvents200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetProjectEdgeEvents200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetProjectEdgeEvents200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNextToken

`func (o *GetProjectEdgeEvents200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetProjectEdgeEvents200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetProjectEdgeEvents200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetProjectEdgeEvents200Response) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### SetNextTokenNil

`func (o *GetProjectEdgeEvents200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *GetProjectEdgeEvents200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


