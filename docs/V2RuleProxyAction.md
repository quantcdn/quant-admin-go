# V2RuleProxyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | Target URL to proxy to | 
**Host** | Pointer to **string** | Host header override | [optional] 
**AuthUser** | Pointer to **string** | Basic auth username | [optional] 
**AuthPass** | Pointer to **string** | Basic auth password | [optional] 
**DisableSslVerify** | Pointer to **bool** | Disable SSL verification | [optional] [default to false]
**CacheLifetime** | Pointer to **NullableString** | Cache lifetime | [optional] 
**OnlyProxy404** | Pointer to **bool** | Only proxy 404 responses | [optional] [default to false]
**InjectHeaders** | Pointer to **map[string]string** | Headers to inject | [optional] 
**ProxyStripHeaders** | Pointer to **[]string** | Headers to strip from response | [optional] 
**ProxyStripRequestHeaders** | Pointer to **[]string** | Headers to strip from request | [optional] 
**OriginTimeout** | Pointer to **string** | Origin timeout | [optional] 
**FailoverMode** | Pointer to **bool** | Enable failover mode | [optional] [default to false]
**FailoverOriginTtfb** | Pointer to **string** | Failover TTFB threshold | [optional] [default to "2000"]
**FailoverOriginStatusCodes** | Pointer to **[]string** | Status codes for failover (default: 200,404,301,302,304) | [optional] 
**FailoverLifetime** | Pointer to **string** | Failover cache lifetime | [optional] [default to "300"]
**Notify** | Pointer to **string** | Notification type (none, slack) | [optional] [default to "none"]
**NotifyConfig** | Pointer to [**NullableV2RuleProxyActionNotifyConfig**](V2RuleProxyActionNotifyConfig.md) |  | [optional] 
**WafEnabled** | Pointer to **bool** | WAF enabled | [optional] [default to false]
**WafConfig** | Pointer to [**NullableWafConfig**](WafConfig.md) |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** | Proxy alert enabled | [optional] 
**ProxyInlineFnEnabled** | Pointer to **bool** | Proxy inline function enabled | [optional] [default to false]
**ApplicationProxy** | Pointer to **bool** | Enable Quant Cloud application proxy mode | [optional] [default to false]
**ApplicationName** | Pointer to **string** | Quant Cloud application name (required when application_proxy is true) | [optional] 
**ApplicationEnvironment** | Pointer to **string** | Quant Cloud application environment (required when application_proxy is true) | [optional] 
**ApplicationContainer** | Pointer to **string** | Quant Cloud application container (required when application_proxy is true) | [optional] 
**ApplicationPort** | Pointer to **int32** | Quant Cloud application port (required when application_proxy is true) | [optional] 
**QuantCloudSelection** | Pointer to [**NullableV2RuleProxyActionQuantCloudSelection**](V2RuleProxyActionQuantCloudSelection.md) |  | [optional] 

## Methods

### NewV2RuleProxyAction

`func NewV2RuleProxyAction(to string, ) *V2RuleProxyAction`

NewV2RuleProxyAction instantiates a new V2RuleProxyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleProxyActionWithDefaults

`func NewV2RuleProxyActionWithDefaults() *V2RuleProxyAction`

NewV2RuleProxyActionWithDefaults instantiates a new V2RuleProxyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *V2RuleProxyAction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V2RuleProxyAction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V2RuleProxyAction) SetTo(v string)`

SetTo sets To field to given value.


### GetHost

`func (o *V2RuleProxyAction) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2RuleProxyAction) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2RuleProxyAction) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2RuleProxyAction) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAuthUser

`func (o *V2RuleProxyAction) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *V2RuleProxyAction) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *V2RuleProxyAction) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *V2RuleProxyAction) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetAuthPass

`func (o *V2RuleProxyAction) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *V2RuleProxyAction) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *V2RuleProxyAction) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *V2RuleProxyAction) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *V2RuleProxyAction) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *V2RuleProxyAction) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *V2RuleProxyAction) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *V2RuleProxyAction) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *V2RuleProxyAction) GetCacheLifetime() string`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *V2RuleProxyAction) GetCacheLifetimeOk() (*string, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *V2RuleProxyAction) SetCacheLifetime(v string)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *V2RuleProxyAction) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### SetCacheLifetimeNil

`func (o *V2RuleProxyAction) SetCacheLifetimeNil(b bool)`

 SetCacheLifetimeNil sets the value for CacheLifetime to be an explicit nil

### UnsetCacheLifetime
`func (o *V2RuleProxyAction) UnsetCacheLifetime()`

UnsetCacheLifetime ensures that no value is present for CacheLifetime, not even an explicit nil
### GetOnlyProxy404

`func (o *V2RuleProxyAction) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *V2RuleProxyAction) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *V2RuleProxyAction) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *V2RuleProxyAction) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *V2RuleProxyAction) GetInjectHeaders() map[string]string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *V2RuleProxyAction) GetInjectHeadersOk() (*map[string]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *V2RuleProxyAction) SetInjectHeaders(v map[string]string)`

SetInjectHeaders sets InjectHeaders field to given value.

### HasInjectHeaders

`func (o *V2RuleProxyAction) HasInjectHeaders() bool`

HasInjectHeaders returns a boolean if a field has been set.

### SetInjectHeadersNil

`func (o *V2RuleProxyAction) SetInjectHeadersNil(b bool)`

 SetInjectHeadersNil sets the value for InjectHeaders to be an explicit nil

### UnsetInjectHeaders
`func (o *V2RuleProxyAction) UnsetInjectHeaders()`

UnsetInjectHeaders ensures that no value is present for InjectHeaders, not even an explicit nil
### GetProxyStripHeaders

`func (o *V2RuleProxyAction) GetProxyStripHeaders() []string`

GetProxyStripHeaders returns the ProxyStripHeaders field if non-nil, zero value otherwise.

### GetProxyStripHeadersOk

`func (o *V2RuleProxyAction) GetProxyStripHeadersOk() (*[]string, bool)`

GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripHeaders

`func (o *V2RuleProxyAction) SetProxyStripHeaders(v []string)`

SetProxyStripHeaders sets ProxyStripHeaders field to given value.

### HasProxyStripHeaders

`func (o *V2RuleProxyAction) HasProxyStripHeaders() bool`

HasProxyStripHeaders returns a boolean if a field has been set.

### GetProxyStripRequestHeaders

`func (o *V2RuleProxyAction) GetProxyStripRequestHeaders() []string`

GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field if non-nil, zero value otherwise.

### GetProxyStripRequestHeadersOk

`func (o *V2RuleProxyAction) GetProxyStripRequestHeadersOk() (*[]string, bool)`

GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripRequestHeaders

`func (o *V2RuleProxyAction) SetProxyStripRequestHeaders(v []string)`

SetProxyStripRequestHeaders sets ProxyStripRequestHeaders field to given value.

### HasProxyStripRequestHeaders

`func (o *V2RuleProxyAction) HasProxyStripRequestHeaders() bool`

HasProxyStripRequestHeaders returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *V2RuleProxyAction) GetOriginTimeout() string`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *V2RuleProxyAction) GetOriginTimeoutOk() (*string, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *V2RuleProxyAction) SetOriginTimeout(v string)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *V2RuleProxyAction) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetFailoverMode

`func (o *V2RuleProxyAction) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *V2RuleProxyAction) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *V2RuleProxyAction) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *V2RuleProxyAction) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *V2RuleProxyAction) GetFailoverOriginTtfb() string`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *V2RuleProxyAction) GetFailoverOriginTtfbOk() (*string, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *V2RuleProxyAction) SetFailoverOriginTtfb(v string)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *V2RuleProxyAction) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverOriginStatusCodes

`func (o *V2RuleProxyAction) GetFailoverOriginStatusCodes() []string`

GetFailoverOriginStatusCodes returns the FailoverOriginStatusCodes field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodesOk

`func (o *V2RuleProxyAction) GetFailoverOriginStatusCodesOk() (*[]string, bool)`

GetFailoverOriginStatusCodesOk returns a tuple with the FailoverOriginStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCodes

`func (o *V2RuleProxyAction) SetFailoverOriginStatusCodes(v []string)`

SetFailoverOriginStatusCodes sets FailoverOriginStatusCodes field to given value.

### HasFailoverOriginStatusCodes

`func (o *V2RuleProxyAction) HasFailoverOriginStatusCodes() bool`

HasFailoverOriginStatusCodes returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *V2RuleProxyAction) GetFailoverLifetime() string`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *V2RuleProxyAction) GetFailoverLifetimeOk() (*string, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *V2RuleProxyAction) SetFailoverLifetime(v string)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *V2RuleProxyAction) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetNotify

`func (o *V2RuleProxyAction) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *V2RuleProxyAction) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *V2RuleProxyAction) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *V2RuleProxyAction) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *V2RuleProxyAction) GetNotifyConfig() V2RuleProxyActionNotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *V2RuleProxyAction) GetNotifyConfigOk() (*V2RuleProxyActionNotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *V2RuleProxyAction) SetNotifyConfig(v V2RuleProxyActionNotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *V2RuleProxyAction) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

### SetNotifyConfigNil

`func (o *V2RuleProxyAction) SetNotifyConfigNil(b bool)`

 SetNotifyConfigNil sets the value for NotifyConfig to be an explicit nil

### UnsetNotifyConfig
`func (o *V2RuleProxyAction) UnsetNotifyConfig()`

UnsetNotifyConfig ensures that no value is present for NotifyConfig, not even an explicit nil
### GetWafEnabled

`func (o *V2RuleProxyAction) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *V2RuleProxyAction) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *V2RuleProxyAction) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *V2RuleProxyAction) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.

### GetWafConfig

`func (o *V2RuleProxyAction) GetWafConfig() WafConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *V2RuleProxyAction) GetWafConfigOk() (*WafConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *V2RuleProxyAction) SetWafConfig(v WafConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *V2RuleProxyAction) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.

### SetWafConfigNil

`func (o *V2RuleProxyAction) SetWafConfigNil(b bool)`

 SetWafConfigNil sets the value for WafConfig to be an explicit nil

### UnsetWafConfig
`func (o *V2RuleProxyAction) UnsetWafConfig()`

UnsetWafConfig ensures that no value is present for WafConfig, not even an explicit nil
### GetProxyAlertEnabled

`func (o *V2RuleProxyAction) GetProxyAlertEnabled() bool`

GetProxyAlertEnabled returns the ProxyAlertEnabled field if non-nil, zero value otherwise.

### GetProxyAlertEnabledOk

`func (o *V2RuleProxyAction) GetProxyAlertEnabledOk() (*bool, bool)`

GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAlertEnabled

`func (o *V2RuleProxyAction) SetProxyAlertEnabled(v bool)`

SetProxyAlertEnabled sets ProxyAlertEnabled field to given value.

### HasProxyAlertEnabled

`func (o *V2RuleProxyAction) HasProxyAlertEnabled() bool`

HasProxyAlertEnabled returns a boolean if a field has been set.

### GetProxyInlineFnEnabled

`func (o *V2RuleProxyAction) GetProxyInlineFnEnabled() bool`

GetProxyInlineFnEnabled returns the ProxyInlineFnEnabled field if non-nil, zero value otherwise.

### GetProxyInlineFnEnabledOk

`func (o *V2RuleProxyAction) GetProxyInlineFnEnabledOk() (*bool, bool)`

GetProxyInlineFnEnabledOk returns a tuple with the ProxyInlineFnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyInlineFnEnabled

`func (o *V2RuleProxyAction) SetProxyInlineFnEnabled(v bool)`

SetProxyInlineFnEnabled sets ProxyInlineFnEnabled field to given value.

### HasProxyInlineFnEnabled

`func (o *V2RuleProxyAction) HasProxyInlineFnEnabled() bool`

HasProxyInlineFnEnabled returns a boolean if a field has been set.

### GetApplicationProxy

`func (o *V2RuleProxyAction) GetApplicationProxy() bool`

GetApplicationProxy returns the ApplicationProxy field if non-nil, zero value otherwise.

### GetApplicationProxyOk

`func (o *V2RuleProxyAction) GetApplicationProxyOk() (*bool, bool)`

GetApplicationProxyOk returns a tuple with the ApplicationProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationProxy

`func (o *V2RuleProxyAction) SetApplicationProxy(v bool)`

SetApplicationProxy sets ApplicationProxy field to given value.

### HasApplicationProxy

`func (o *V2RuleProxyAction) HasApplicationProxy() bool`

HasApplicationProxy returns a boolean if a field has been set.

### GetApplicationName

`func (o *V2RuleProxyAction) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *V2RuleProxyAction) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *V2RuleProxyAction) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *V2RuleProxyAction) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationEnvironment

`func (o *V2RuleProxyAction) GetApplicationEnvironment() string`

GetApplicationEnvironment returns the ApplicationEnvironment field if non-nil, zero value otherwise.

### GetApplicationEnvironmentOk

`func (o *V2RuleProxyAction) GetApplicationEnvironmentOk() (*string, bool)`

GetApplicationEnvironmentOk returns a tuple with the ApplicationEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEnvironment

`func (o *V2RuleProxyAction) SetApplicationEnvironment(v string)`

SetApplicationEnvironment sets ApplicationEnvironment field to given value.

### HasApplicationEnvironment

`func (o *V2RuleProxyAction) HasApplicationEnvironment() bool`

HasApplicationEnvironment returns a boolean if a field has been set.

### GetApplicationContainer

`func (o *V2RuleProxyAction) GetApplicationContainer() string`

GetApplicationContainer returns the ApplicationContainer field if non-nil, zero value otherwise.

### GetApplicationContainerOk

`func (o *V2RuleProxyAction) GetApplicationContainerOk() (*string, bool)`

GetApplicationContainerOk returns a tuple with the ApplicationContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationContainer

`func (o *V2RuleProxyAction) SetApplicationContainer(v string)`

SetApplicationContainer sets ApplicationContainer field to given value.

### HasApplicationContainer

`func (o *V2RuleProxyAction) HasApplicationContainer() bool`

HasApplicationContainer returns a boolean if a field has been set.

### GetApplicationPort

`func (o *V2RuleProxyAction) GetApplicationPort() int32`

GetApplicationPort returns the ApplicationPort field if non-nil, zero value otherwise.

### GetApplicationPortOk

`func (o *V2RuleProxyAction) GetApplicationPortOk() (*int32, bool)`

GetApplicationPortOk returns a tuple with the ApplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPort

`func (o *V2RuleProxyAction) SetApplicationPort(v int32)`

SetApplicationPort sets ApplicationPort field to given value.

### HasApplicationPort

`func (o *V2RuleProxyAction) HasApplicationPort() bool`

HasApplicationPort returns a boolean if a field has been set.

### GetQuantCloudSelection

`func (o *V2RuleProxyAction) GetQuantCloudSelection() V2RuleProxyActionQuantCloudSelection`

GetQuantCloudSelection returns the QuantCloudSelection field if non-nil, zero value otherwise.

### GetQuantCloudSelectionOk

`func (o *V2RuleProxyAction) GetQuantCloudSelectionOk() (*V2RuleProxyActionQuantCloudSelection, bool)`

GetQuantCloudSelectionOk returns a tuple with the QuantCloudSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantCloudSelection

`func (o *V2RuleProxyAction) SetQuantCloudSelection(v V2RuleProxyActionQuantCloudSelection)`

SetQuantCloudSelection sets QuantCloudSelection field to given value.

### HasQuantCloudSelection

`func (o *V2RuleProxyAction) HasQuantCloudSelection() bool`

HasQuantCloudSelection returns a boolean if a field has been set.

### SetQuantCloudSelectionNil

`func (o *V2RuleProxyAction) SetQuantCloudSelectionNil(b bool)`

 SetQuantCloudSelectionNil sets the value for QuantCloudSelection to be an explicit nil

### UnsetQuantCloudSelection
`func (o *V2RuleProxyAction) UnsetQuantCloudSelection()`

UnsetQuantCloudSelection ensures that no value is present for QuantCloudSelection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


