# ListAIToolExecutions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executions** | [**[]ListAIToolExecutions200ResponseExecutionsInner**](ListAIToolExecutions200ResponseExecutionsInner.md) |  | 
**Count** | **int32** | Number of executions returned | 
**OrgId** | **string** | Organization identifier | 
**Status** | Pointer to **string** | Filter applied (or &#39;all&#39;) | [optional] 

## Methods

### NewListAIToolExecutions200Response

`func NewListAIToolExecutions200Response(executions []ListAIToolExecutions200ResponseExecutionsInner, count int32, orgId string, ) *ListAIToolExecutions200Response`

NewListAIToolExecutions200Response instantiates a new ListAIToolExecutions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAIToolExecutions200ResponseWithDefaults

`func NewListAIToolExecutions200ResponseWithDefaults() *ListAIToolExecutions200Response`

NewListAIToolExecutions200ResponseWithDefaults instantiates a new ListAIToolExecutions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutions

`func (o *ListAIToolExecutions200Response) GetExecutions() []ListAIToolExecutions200ResponseExecutionsInner`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *ListAIToolExecutions200Response) GetExecutionsOk() (*[]ListAIToolExecutions200ResponseExecutionsInner, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *ListAIToolExecutions200Response) SetExecutions(v []ListAIToolExecutions200ResponseExecutionsInner)`

SetExecutions sets Executions field to given value.


### GetCount

`func (o *ListAIToolExecutions200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListAIToolExecutions200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListAIToolExecutions200Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetOrgId

`func (o *ListAIToolExecutions200Response) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ListAIToolExecutions200Response) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ListAIToolExecutions200Response) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetStatus

`func (o *ListAIToolExecutions200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAIToolExecutions200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAIToolExecutions200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListAIToolExecutions200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


