# ProxyNotifyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginStatusCode** | **[]int32** |  | 
**Period** | **int32** |  | 
**SlackWebhook** | Pointer to **string** |  | [optional] 

## Methods

### NewProxyNotifyConfig

`func NewProxyNotifyConfig(originStatusCode []int32, period int32, ) *ProxyNotifyConfig`

NewProxyNotifyConfig instantiates a new ProxyNotifyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyNotifyConfigWithDefaults

`func NewProxyNotifyConfigWithDefaults() *ProxyNotifyConfig`

NewProxyNotifyConfigWithDefaults instantiates a new ProxyNotifyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginStatusCode

`func (o *ProxyNotifyConfig) GetOriginStatusCode() []int32`

GetOriginStatusCode returns the OriginStatusCode field if non-nil, zero value otherwise.

### GetOriginStatusCodeOk

`func (o *ProxyNotifyConfig) GetOriginStatusCodeOk() (*[]int32, bool)`

GetOriginStatusCodeOk returns a tuple with the OriginStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatusCode

`func (o *ProxyNotifyConfig) SetOriginStatusCode(v []int32)`

SetOriginStatusCode sets OriginStatusCode field to given value.


### GetPeriod

`func (o *ProxyNotifyConfig) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ProxyNotifyConfig) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ProxyNotifyConfig) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetSlackWebhook

`func (o *ProxyNotifyConfig) GetSlackWebhook() string`

GetSlackWebhook returns the SlackWebhook field if non-nil, zero value otherwise.

### GetSlackWebhookOk

`func (o *ProxyNotifyConfig) GetSlackWebhookOk() (*string, bool)`

GetSlackWebhookOk returns a tuple with the SlackWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhook

`func (o *ProxyNotifyConfig) SetSlackWebhook(v string)`

SetSlackWebhook sets SlackWebhook field to given value.

### HasSlackWebhook

`func (o *ProxyNotifyConfig) HasSlackWebhook() bool`

HasSlackWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


