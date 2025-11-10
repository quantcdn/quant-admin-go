# GetEnvironmentLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogEvents** | Pointer to [**[]GetEnvironmentLogs200ResponseLogEventsInner**](GetEnvironmentLogs200ResponseLogEventsInner.md) | Array of log events | [optional] 
**NextToken** | Pointer to **NullableString** | Token for fetching next page of results (null if no more pages) | [optional] 

## Methods

### NewGetEnvironmentLogs200Response

`func NewGetEnvironmentLogs200Response() *GetEnvironmentLogs200Response`

NewGetEnvironmentLogs200Response instantiates a new GetEnvironmentLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnvironmentLogs200ResponseWithDefaults

`func NewGetEnvironmentLogs200ResponseWithDefaults() *GetEnvironmentLogs200Response`

NewGetEnvironmentLogs200ResponseWithDefaults instantiates a new GetEnvironmentLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogEvents

`func (o *GetEnvironmentLogs200Response) GetLogEvents() []GetEnvironmentLogs200ResponseLogEventsInner`

GetLogEvents returns the LogEvents field if non-nil, zero value otherwise.

### GetLogEventsOk

`func (o *GetEnvironmentLogs200Response) GetLogEventsOk() (*[]GetEnvironmentLogs200ResponseLogEventsInner, bool)`

GetLogEventsOk returns a tuple with the LogEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEvents

`func (o *GetEnvironmentLogs200Response) SetLogEvents(v []GetEnvironmentLogs200ResponseLogEventsInner)`

SetLogEvents sets LogEvents field to given value.

### HasLogEvents

`func (o *GetEnvironmentLogs200Response) HasLogEvents() bool`

HasLogEvents returns a boolean if a field has been set.

### GetNextToken

`func (o *GetEnvironmentLogs200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetEnvironmentLogs200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetEnvironmentLogs200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetEnvironmentLogs200Response) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### SetNextTokenNil

`func (o *GetEnvironmentLogs200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *GetEnvironmentLogs200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


