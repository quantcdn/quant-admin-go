# GetAIUsageStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRequests** | Pointer to **int32** | Total number of API requests | [optional] 
**TotalTokens** | Pointer to **int32** | Total tokens consumed across all requests | [optional] 
**ByModel** | Pointer to [**map[string]GetAIUsageStats200ResponseByModelValue**](GetAIUsageStats200ResponseByModelValue.md) | Usage breakdown by model ID | [optional] 

## Methods

### NewGetAIUsageStats200Response

`func NewGetAIUsageStats200Response() *GetAIUsageStats200Response`

NewGetAIUsageStats200Response instantiates a new GetAIUsageStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAIUsageStats200ResponseWithDefaults

`func NewGetAIUsageStats200ResponseWithDefaults() *GetAIUsageStats200Response`

NewGetAIUsageStats200ResponseWithDefaults instantiates a new GetAIUsageStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRequests

`func (o *GetAIUsageStats200Response) GetTotalRequests() int32`

GetTotalRequests returns the TotalRequests field if non-nil, zero value otherwise.

### GetTotalRequestsOk

`func (o *GetAIUsageStats200Response) GetTotalRequestsOk() (*int32, bool)`

GetTotalRequestsOk returns a tuple with the TotalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequests

`func (o *GetAIUsageStats200Response) SetTotalRequests(v int32)`

SetTotalRequests sets TotalRequests field to given value.

### HasTotalRequests

`func (o *GetAIUsageStats200Response) HasTotalRequests() bool`

HasTotalRequests returns a boolean if a field has been set.

### GetTotalTokens

`func (o *GetAIUsageStats200Response) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *GetAIUsageStats200Response) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *GetAIUsageStats200Response) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *GetAIUsageStats200Response) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetByModel

`func (o *GetAIUsageStats200Response) GetByModel() map[string]GetAIUsageStats200ResponseByModelValue`

GetByModel returns the ByModel field if non-nil, zero value otherwise.

### GetByModelOk

`func (o *GetAIUsageStats200Response) GetByModelOk() (*map[string]GetAIUsageStats200ResponseByModelValue, bool)`

GetByModelOk returns a tuple with the ByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByModel

`func (o *GetAIUsageStats200Response) SetByModel(v map[string]GetAIUsageStats200ResponseByModelValue)`

SetByModel sets ByModel field to given value.

### HasByModel

`func (o *GetAIUsageStats200Response) HasByModel() bool`

HasByModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


