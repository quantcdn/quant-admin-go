# GetEnvironmentLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogEvents** | Pointer to [**[]GetEnvironmentLogs200ResponseLogEventsInner**](GetEnvironmentLogs200ResponseLogEventsInner.md) | Array of log events | [optional] 
**LogGroupName** | Pointer to **NullableString** | CloudWatch log group the events were read from | [optional] 
**Pagination** | Pointer to [**GetEnvironmentLogs200ResponsePagination**](GetEnvironmentLogs200ResponsePagination.md) |  | [optional] 
**NextToken** | Pointer to **NullableString** | Same as pagination.nextToken; kept for backward compatibility | [optional] 

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

### GetLogGroupName

`func (o *GetEnvironmentLogs200Response) GetLogGroupName() string`

GetLogGroupName returns the LogGroupName field if non-nil, zero value otherwise.

### GetLogGroupNameOk

`func (o *GetEnvironmentLogs200Response) GetLogGroupNameOk() (*string, bool)`

GetLogGroupNameOk returns a tuple with the LogGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGroupName

`func (o *GetEnvironmentLogs200Response) SetLogGroupName(v string)`

SetLogGroupName sets LogGroupName field to given value.

### HasLogGroupName

`func (o *GetEnvironmentLogs200Response) HasLogGroupName() bool`

HasLogGroupName returns a boolean if a field has been set.

### SetLogGroupNameNil

`func (o *GetEnvironmentLogs200Response) SetLogGroupNameNil(b bool)`

 SetLogGroupNameNil sets the value for LogGroupName to be an explicit nil

### UnsetLogGroupName
`func (o *GetEnvironmentLogs200Response) UnsetLogGroupName()`

UnsetLogGroupName ensures that no value is present for LogGroupName, not even an explicit nil
### GetPagination

`func (o *GetEnvironmentLogs200Response) GetPagination() GetEnvironmentLogs200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetEnvironmentLogs200Response) GetPaginationOk() (*GetEnvironmentLogs200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetEnvironmentLogs200Response) SetPagination(v GetEnvironmentLogs200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetEnvironmentLogs200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

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


