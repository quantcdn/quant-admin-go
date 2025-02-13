# RuleProxyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** |  | 
**Host** | Pointer to **string** |  | [optional] 
**AuthUser** | Pointer to **string** |  | [optional] 
**AuthPass** | Pointer to **string** |  | [optional] 
**DisableSslVerify** | Pointer to **bool** |  | [optional] 
**CacheLifetime** | Pointer to **int32** |  | [optional] 
**OnlyProxy404** | Pointer to **bool** |  | [optional] 
**InjectHeaders** | Pointer to **map[string]string** |  | [optional] 
**ProxyStripHeaders** | Pointer to **[]string** |  | [optional] 
**ProxyStripRequestHeaders** | Pointer to **[]string** |  | [optional] 
**Failover** | Pointer to [**FailoverConfig**](FailoverConfig.md) |  | [optional] 
**Notify** | Pointer to **string** |  | [optional] 
**NotifyConfig** | Pointer to [**NotifyConfig**](NotifyConfig.md) |  | [optional] 
**WafEnabled** | **bool** |  | [default to false]
**WafConfig** | Pointer to [**WAFConfig**](WAFConfig.md) |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** |  | [optional] 
**ProxyInlineFnEnabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRuleProxyAction

`func NewRuleProxyAction(to string, wafEnabled bool, ) *RuleProxyAction`

NewRuleProxyAction instantiates a new RuleProxyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleProxyActionWithDefaults

`func NewRuleProxyActionWithDefaults() *RuleProxyAction`

NewRuleProxyActionWithDefaults instantiates a new RuleProxyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetAuthUser

`func (o *RuleProxyAction) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *RuleProxyAction) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *RuleProxyAction) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *RuleProxyAction) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

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

### GetProxyStripRequestHeaders

`func (o *RuleProxyAction) GetProxyStripRequestHeaders() []string`

GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field if non-nil, zero value otherwise.

### GetProxyStripRequestHeadersOk

`func (o *RuleProxyAction) GetProxyStripRequestHeadersOk() (*[]string, bool)`

GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripRequestHeaders

`func (o *RuleProxyAction) SetProxyStripRequestHeaders(v []string)`

SetProxyStripRequestHeaders sets ProxyStripRequestHeaders field to given value.

### HasProxyStripRequestHeaders

`func (o *RuleProxyAction) HasProxyStripRequestHeaders() bool`

HasProxyStripRequestHeaders returns a boolean if a field has been set.

### GetFailover

`func (o *RuleProxyAction) GetFailover() FailoverConfig`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *RuleProxyAction) GetFailoverOk() (*FailoverConfig, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *RuleProxyAction) SetFailover(v FailoverConfig)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *RuleProxyAction) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

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

### GetNotifyConfig

`func (o *RuleProxyAction) GetNotifyConfig() NotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *RuleProxyAction) GetNotifyConfigOk() (*NotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *RuleProxyAction) SetNotifyConfig(v NotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *RuleProxyAction) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

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


### GetWafConfig

`func (o *RuleProxyAction) GetWafConfig() WAFConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *RuleProxyAction) GetWafConfigOk() (*WAFConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *RuleProxyAction) SetWafConfig(v WAFConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *RuleProxyAction) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.

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

### GetProxyInlineFnEnabled

`func (o *RuleProxyAction) GetProxyInlineFnEnabled() bool`

GetProxyInlineFnEnabled returns the ProxyInlineFnEnabled field if non-nil, zero value otherwise.

### GetProxyInlineFnEnabledOk

`func (o *RuleProxyAction) GetProxyInlineFnEnabledOk() (*bool, bool)`

GetProxyInlineFnEnabledOk returns a tuple with the ProxyInlineFnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyInlineFnEnabled

`func (o *RuleProxyAction) SetProxyInlineFnEnabled(v bool)`

SetProxyInlineFnEnabled sets ProxyInlineFnEnabled field to given value.

### HasProxyInlineFnEnabled

`func (o *RuleProxyAction) HasProxyInlineFnEnabled() bool`

HasProxyInlineFnEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


