# ProxyNotifyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginStatusCodes** | Pointer to **[]string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] [default to "60"]
**SlackWebhook** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewProxyNotifyConfig

`func NewProxyNotifyConfig() *ProxyNotifyConfig`

NewProxyNotifyConfig instantiates a new ProxyNotifyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyNotifyConfigWithDefaults

`func NewProxyNotifyConfigWithDefaults() *ProxyNotifyConfig`

NewProxyNotifyConfigWithDefaults instantiates a new ProxyNotifyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginStatusCodes

`func (o *ProxyNotifyConfig) GetOriginStatusCodes() []string`

GetOriginStatusCodes returns the OriginStatusCodes field if non-nil, zero value otherwise.

### GetOriginStatusCodesOk

`func (o *ProxyNotifyConfig) GetOriginStatusCodesOk() (*[]string, bool)`

GetOriginStatusCodesOk returns a tuple with the OriginStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatusCodes

`func (o *ProxyNotifyConfig) SetOriginStatusCodes(v []string)`

SetOriginStatusCodes sets OriginStatusCodes field to given value.

### HasOriginStatusCodes

`func (o *ProxyNotifyConfig) HasOriginStatusCodes() bool`

HasOriginStatusCodes returns a boolean if a field has been set.

### GetPeriod

`func (o *ProxyNotifyConfig) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ProxyNotifyConfig) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ProxyNotifyConfig) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ProxyNotifyConfig) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

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


