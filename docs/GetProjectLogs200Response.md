# GetProjectLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]GetProjectLogs200ResponseLogsInner**](GetProjectLogs200ResponseLogsInner.md) | Structured CloudFront access log entries. Each entry carries request, response, timing and cache fields as emitted by the edge. | [optional] 
**Count** | Pointer to **int32** | Number of entries in this response | [optional] 
**NextToken** | Pointer to **NullableString** | Token for the next page, or null when there are no more entries | [optional] 

## Methods

### NewGetProjectLogs200Response

`func NewGetProjectLogs200Response() *GetProjectLogs200Response`

NewGetProjectLogs200Response instantiates a new GetProjectLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectLogs200ResponseWithDefaults

`func NewGetProjectLogs200ResponseWithDefaults() *GetProjectLogs200Response`

NewGetProjectLogs200ResponseWithDefaults instantiates a new GetProjectLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *GetProjectLogs200Response) GetLogs() []GetProjectLogs200ResponseLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *GetProjectLogs200Response) GetLogsOk() (*[]GetProjectLogs200ResponseLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *GetProjectLogs200Response) SetLogs(v []GetProjectLogs200ResponseLogsInner)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *GetProjectLogs200Response) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetCount

`func (o *GetProjectLogs200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetProjectLogs200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetProjectLogs200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetProjectLogs200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNextToken

`func (o *GetProjectLogs200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetProjectLogs200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetProjectLogs200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetProjectLogs200Response) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### SetNextTokenNil

`func (o *GetProjectLogs200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *GetProjectLogs200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


