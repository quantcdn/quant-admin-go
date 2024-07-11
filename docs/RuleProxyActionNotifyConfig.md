# RuleProxyActionNotifyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginStatusCodes** | Pointer to **[]string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] [default to "60"]
**SlackWebhook** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewRuleProxyActionNotifyConfig

`func NewRuleProxyActionNotifyConfig() *RuleProxyActionNotifyConfig`

NewRuleProxyActionNotifyConfig instantiates a new RuleProxyActionNotifyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleProxyActionNotifyConfigWithDefaults

`func NewRuleProxyActionNotifyConfigWithDefaults() *RuleProxyActionNotifyConfig`

NewRuleProxyActionNotifyConfigWithDefaults instantiates a new RuleProxyActionNotifyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginStatusCodes

`func (o *RuleProxyActionNotifyConfig) GetOriginStatusCodes() []string`

GetOriginStatusCodes returns the OriginStatusCodes field if non-nil, zero value otherwise.

### GetOriginStatusCodesOk

`func (o *RuleProxyActionNotifyConfig) GetOriginStatusCodesOk() (*[]string, bool)`

GetOriginStatusCodesOk returns a tuple with the OriginStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginStatusCodes

`func (o *RuleProxyActionNotifyConfig) SetOriginStatusCodes(v []string)`

SetOriginStatusCodes sets OriginStatusCodes field to given value.

### HasOriginStatusCodes

`func (o *RuleProxyActionNotifyConfig) HasOriginStatusCodes() bool`

HasOriginStatusCodes returns a boolean if a field has been set.

### GetPeriod

`func (o *RuleProxyActionNotifyConfig) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RuleProxyActionNotifyConfig) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RuleProxyActionNotifyConfig) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RuleProxyActionNotifyConfig) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetSlackWebhook

`func (o *RuleProxyActionNotifyConfig) GetSlackWebhook() string`

GetSlackWebhook returns the SlackWebhook field if non-nil, zero value otherwise.

### GetSlackWebhookOk

`func (o *RuleProxyActionNotifyConfig) GetSlackWebhookOk() (*string, bool)`

GetSlackWebhookOk returns a tuple with the SlackWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhook

`func (o *RuleProxyActionNotifyConfig) SetSlackWebhook(v string)`

SetSlackWebhook sets SlackWebhook field to given value.

### HasSlackWebhook

`func (o *RuleProxyActionNotifyConfig) HasSlackWebhook() bool`

HasSlackWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


