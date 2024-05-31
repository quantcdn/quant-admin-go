# RuleProxyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**WafEnabled** | **bool** |  | [default to false]
**ProxyStripHeaders** | Pointer to **[]string** |  | [optional] 
**AuthPass** | Pointer to **string** |  | [optional] 
**OriginTimeout** | Pointer to **int32** |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** |  | [optional] 
**Notify** | Pointer to **string** |  | [optional] 
**FailoverOriginStatusCode** | Pointer to **[]string** |  | [optional] 
**FailoverOriginTtfb** | Pointer to **int32** |  | [optional] 
**FailoverLifetime** | Pointer to **int32** |  | [optional] 
**OnlyProxy404** | Pointer to **bool** |  | [optional] 
**InjectHeaders** | Pointer to **map[string]string** |  | [optional] 
**WafConfig** | Pointer to [**RuleWAFConfig**](RuleWAFConfig.md) |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**CacheLifetime** | Pointer to **int32** |  | [optional] 
**FailoverMode** | Pointer to **bool** |  | [optional] 
**DisableSslVerify** | Pointer to **bool** |  | [optional] 
**NotifyConfig** | [**ProxyNotifyConfig**](ProxyNotifyConfig.md) |  | 

## Methods

### NewRuleProxyAction

`func NewRuleProxyAction(wafEnabled bool, notifyConfig ProxyNotifyConfig, ) *RuleProxyAction`

NewRuleProxyAction instantiates a new RuleProxyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleProxyActionWithDefaults

`func NewRuleProxyActionWithDefaults() *RuleProxyAction`

NewRuleProxyActionWithDefaults instantiates a new RuleProxyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RuleProxyAction) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RuleProxyAction) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RuleProxyAction) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RuleProxyAction) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetWafEnabled

`func (o *RuleProxyAction) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *RuleProxyAction) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *RuleProxyAction) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.


### GetProxyStripHeaders

`func (o *RuleProxyAction) GetProxyStripHeaders() []string`

GetProxyStripHeaders returns the ProxyStripHeaders field if non-nil, zero value otherwise.

### GetProxyStripHeadersOk

`func (o *RuleProxyAction) GetProxyStripHeadersOk() (*[]string, bool)`

GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripHeaders

`func (o *RuleProxyAction) SetProxyStripHeaders(v []string)`

SetProxyStripHeaders sets ProxyStripHeaders field to given value.

### HasProxyStripHeaders

`func (o *RuleProxyAction) HasProxyStripHeaders() bool`

HasProxyStripHeaders returns a boolean if a field has been set.

### GetAuthPass

`func (o *RuleProxyAction) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *RuleProxyAction) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *RuleProxyAction) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *RuleProxyAction) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *RuleProxyAction) GetOriginTimeout() int32`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *RuleProxyAction) GetOriginTimeoutOk() (*int32, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *RuleProxyAction) SetOriginTimeout(v int32)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *RuleProxyAction) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetProxyAlertEnabled

`func (o *RuleProxyAction) GetProxyAlertEnabled() bool`

GetProxyAlertEnabled returns the ProxyAlertEnabled field if non-nil, zero value otherwise.

### GetProxyAlertEnabledOk

`func (o *RuleProxyAction) GetProxyAlertEnabledOk() (*bool, bool)`

GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAlertEnabled

`func (o *RuleProxyAction) SetProxyAlertEnabled(v bool)`

SetProxyAlertEnabled sets ProxyAlertEnabled field to given value.

### HasProxyAlertEnabled

`func (o *RuleProxyAction) HasProxyAlertEnabled() bool`

HasProxyAlertEnabled returns a boolean if a field has been set.

### GetNotify

`func (o *RuleProxyAction) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *RuleProxyAction) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *RuleProxyAction) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *RuleProxyAction) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetFailoverOriginStatusCode

`func (o *RuleProxyAction) GetFailoverOriginStatusCode() []string`

GetFailoverOriginStatusCode returns the FailoverOriginStatusCode field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodeOk

`func (o *RuleProxyAction) GetFailoverOriginStatusCodeOk() (*[]string, bool)`

GetFailoverOriginStatusCodeOk returns a tuple with the FailoverOriginStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCode

`func (o *RuleProxyAction) SetFailoverOriginStatusCode(v []string)`

SetFailoverOriginStatusCode sets FailoverOriginStatusCode field to given value.

### HasFailoverOriginStatusCode

`func (o *RuleProxyAction) HasFailoverOriginStatusCode() bool`

HasFailoverOriginStatusCode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *RuleProxyAction) GetFailoverOriginTtfb() int32`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *RuleProxyAction) GetFailoverOriginTtfbOk() (*int32, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *RuleProxyAction) SetFailoverOriginTtfb(v int32)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *RuleProxyAction) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *RuleProxyAction) GetFailoverLifetime() int32`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *RuleProxyAction) GetFailoverLifetimeOk() (*int32, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *RuleProxyAction) SetFailoverLifetime(v int32)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *RuleProxyAction) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetOnlyProxy404

`func (o *RuleProxyAction) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *RuleProxyAction) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *RuleProxyAction) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *RuleProxyAction) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *RuleProxyAction) GetInjectHeaders() map[string]string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *RuleProxyAction) GetInjectHeadersOk() (*map[string]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *RuleProxyAction) SetInjectHeaders(v map[string]string)`

SetInjectHeaders sets InjectHeaders field to given value.

### HasInjectHeaders

`func (o *RuleProxyAction) HasInjectHeaders() bool`

HasInjectHeaders returns a boolean if a field has been set.

### GetWafConfig

`func (o *RuleProxyAction) GetWafConfig() RuleWAFConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *RuleProxyAction) GetWafConfigOk() (*RuleWAFConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *RuleProxyAction) SetWafConfig(v RuleWAFConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *RuleProxyAction) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.

### GetTo

`func (o *RuleProxyAction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RuleProxyAction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RuleProxyAction) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RuleProxyAction) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *RuleProxyAction) GetCacheLifetime() int32`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *RuleProxyAction) GetCacheLifetimeOk() (*int32, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *RuleProxyAction) SetCacheLifetime(v int32)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *RuleProxyAction) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### GetFailoverMode

`func (o *RuleProxyAction) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *RuleProxyAction) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *RuleProxyAction) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *RuleProxyAction) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *RuleProxyAction) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *RuleProxyAction) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *RuleProxyAction) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *RuleProxyAction) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *RuleProxyAction) GetNotifyConfig() ProxyNotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *RuleProxyAction) GetNotifyConfigOk() (*ProxyNotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *RuleProxyAction) SetNotifyConfig(v ProxyNotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


