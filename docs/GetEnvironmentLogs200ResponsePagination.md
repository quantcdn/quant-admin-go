# GetEnvironmentLogs200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Page size that was applied | [optional] 
**HasMore** | Pointer to **bool** | True when another page is available | [optional] 
**NextToken** | Pointer to **NullableString** | Token for the next page. Present only when hasMore is true. | [optional] 
**Total** | Pointer to **NullableInt32** | Total events in the time range. Present only when includeTotal&#x3D;true. | [optional] 
**TotalPages** | Pointer to **NullableInt32** | ceil(total / limit). Present only when includeTotal&#x3D;true. | [optional] 

## Methods

### NewGetEnvironmentLogs200ResponsePagination

`func NewGetEnvironmentLogs200ResponsePagination() *GetEnvironmentLogs200ResponsePagination`

NewGetEnvironmentLogs200ResponsePagination instantiates a new GetEnvironmentLogs200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnvironmentLogs200ResponsePaginationWithDefaults

`func NewGetEnvironmentLogs200ResponsePaginationWithDefaults() *GetEnvironmentLogs200ResponsePagination`

NewGetEnvironmentLogs200ResponsePaginationWithDefaults instantiates a new GetEnvironmentLogs200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetEnvironmentLogs200ResponsePagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetEnvironmentLogs200ResponsePagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetEnvironmentLogs200ResponsePagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetEnvironmentLogs200ResponsePagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetHasMore

`func (o *GetEnvironmentLogs200ResponsePagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetEnvironmentLogs200ResponsePagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetEnvironmentLogs200ResponsePagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetEnvironmentLogs200ResponsePagination) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNextToken

`func (o *GetEnvironmentLogs200ResponsePagination) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetEnvironmentLogs200ResponsePagination) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetEnvironmentLogs200ResponsePagination) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetEnvironmentLogs200ResponsePagination) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### SetNextTokenNil

`func (o *GetEnvironmentLogs200ResponsePagination) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *GetEnvironmentLogs200ResponsePagination) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil
### GetTotal

`func (o *GetEnvironmentLogs200ResponsePagination) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetEnvironmentLogs200ResponsePagination) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetEnvironmentLogs200ResponsePagination) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetEnvironmentLogs200ResponsePagination) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GetEnvironmentLogs200ResponsePagination) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GetEnvironmentLogs200ResponsePagination) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetTotalPages

`func (o *GetEnvironmentLogs200ResponsePagination) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetEnvironmentLogs200ResponsePagination) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetEnvironmentLogs200ResponsePagination) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetEnvironmentLogs200ResponsePagination) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### SetTotalPagesNil

`func (o *GetEnvironmentLogs200ResponsePagination) SetTotalPagesNil(b bool)`

 SetTotalPagesNil sets the value for TotalPages to be an explicit nil

### UnsetTotalPages
`func (o *GetEnvironmentLogs200ResponsePagination) UnsetTotalPages()`

UnsetTotalPages ensures that no value is present for TotalPages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


