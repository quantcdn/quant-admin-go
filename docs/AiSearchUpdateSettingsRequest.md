# AiSearchUpdateSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicAccess** | Pointer to **map[string]interface{}** |  | [optional] 
**RateLimits** | Pointer to **map[string]interface{}** |  | [optional] 
**QueryTextRetention** | Pointer to **string** | What analytics keeps of a visitor&#39;s query. default: the organisation default (store, or hash for strict organisations). | [optional] 

## Methods

### NewAiSearchUpdateSettingsRequest

`func NewAiSearchUpdateSettingsRequest() *AiSearchUpdateSettingsRequest`

NewAiSearchUpdateSettingsRequest instantiates a new AiSearchUpdateSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSearchUpdateSettingsRequestWithDefaults

`func NewAiSearchUpdateSettingsRequestWithDefaults() *AiSearchUpdateSettingsRequest`

NewAiSearchUpdateSettingsRequestWithDefaults instantiates a new AiSearchUpdateSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicAccess

`func (o *AiSearchUpdateSettingsRequest) GetPublicAccess() map[string]interface{}`

GetPublicAccess returns the PublicAccess field if non-nil, zero value otherwise.

### GetPublicAccessOk

`func (o *AiSearchUpdateSettingsRequest) GetPublicAccessOk() (*map[string]interface{}, bool)`

GetPublicAccessOk returns a tuple with the PublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccess

`func (o *AiSearchUpdateSettingsRequest) SetPublicAccess(v map[string]interface{})`

SetPublicAccess sets PublicAccess field to given value.

### HasPublicAccess

`func (o *AiSearchUpdateSettingsRequest) HasPublicAccess() bool`

HasPublicAccess returns a boolean if a field has been set.

### GetRateLimits

`func (o *AiSearchUpdateSettingsRequest) GetRateLimits() map[string]interface{}`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *AiSearchUpdateSettingsRequest) GetRateLimitsOk() (*map[string]interface{}, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *AiSearchUpdateSettingsRequest) SetRateLimits(v map[string]interface{})`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *AiSearchUpdateSettingsRequest) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetQueryTextRetention

`func (o *AiSearchUpdateSettingsRequest) GetQueryTextRetention() string`

GetQueryTextRetention returns the QueryTextRetention field if non-nil, zero value otherwise.

### GetQueryTextRetentionOk

`func (o *AiSearchUpdateSettingsRequest) GetQueryTextRetentionOk() (*string, bool)`

GetQueryTextRetentionOk returns a tuple with the QueryTextRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTextRetention

`func (o *AiSearchUpdateSettingsRequest) SetQueryTextRetention(v string)`

SetQueryTextRetention sets QueryTextRetention field to given value.

### HasQueryTextRetention

`func (o *AiSearchUpdateSettingsRequest) HasQueryTextRetention() bool`

HasQueryTextRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


