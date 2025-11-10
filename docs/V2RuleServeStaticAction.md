# V2RuleServeStaticAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
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
**NotifyConfig** | Pointer to [**NullableV2RuleProxyActionAllOfNotifyConfig**](V2RuleProxyActionAllOfNotifyConfig.md) |  | [optional] 
**WafEnabled** | Pointer to **bool** | WAF enabled | [optional] [default to false]
**WafConfig** | Pointer to [**WafConfig**](WafConfig.md) |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** | Proxy alert enabled | [optional] 
**ProxyInlineFnEnabled** | Pointer to **bool** | Proxy inline function enabled | [optional] [default to false]
**QuantCloudSelection** | Pointer to [**NullableV2RuleProxyActionAllOfQuantCloudSelection**](V2RuleProxyActionAllOfQuantCloudSelection.md) |  | [optional] 
**StaticFilePath** | **string** | Path to the static file to serve | 

## Methods

### NewV2RuleServeStaticAction

`func NewV2RuleServeStaticAction(message string, error_ bool, to string, staticFilePath string, ) *V2RuleServeStaticAction`

NewV2RuleServeStaticAction instantiates a new V2RuleServeStaticAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleServeStaticActionWithDefaults

`func NewV2RuleServeStaticActionWithDefaults() *V2RuleServeStaticAction`

NewV2RuleServeStaticActionWithDefaults instantiates a new V2RuleServeStaticAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleServeStaticAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleServeStaticAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleServeStaticAction) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleServeStaticAction) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleServeStaticAction) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleServeStaticAction) SetError(v bool)`

SetError sets Error field to given value.


### GetTo

`func (o *V2RuleServeStaticAction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V2RuleServeStaticAction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V2RuleServeStaticAction) SetTo(v string)`

SetTo sets To field to given value.


### GetHost

`func (o *V2RuleServeStaticAction) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2RuleServeStaticAction) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2RuleServeStaticAction) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2RuleServeStaticAction) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAuthUser

`func (o *V2RuleServeStaticAction) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *V2RuleServeStaticAction) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *V2RuleServeStaticAction) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *V2RuleServeStaticAction) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetAuthPass

`func (o *V2RuleServeStaticAction) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *V2RuleServeStaticAction) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *V2RuleServeStaticAction) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *V2RuleServeStaticAction) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *V2RuleServeStaticAction) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *V2RuleServeStaticAction) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *V2RuleServeStaticAction) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *V2RuleServeStaticAction) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *V2RuleServeStaticAction) GetCacheLifetime() string`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *V2RuleServeStaticAction) GetCacheLifetimeOk() (*string, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *V2RuleServeStaticAction) SetCacheLifetime(v string)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *V2RuleServeStaticAction) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### SetCacheLifetimeNil

`func (o *V2RuleServeStaticAction) SetCacheLifetimeNil(b bool)`

 SetCacheLifetimeNil sets the value for CacheLifetime to be an explicit nil

### UnsetCacheLifetime
`func (o *V2RuleServeStaticAction) UnsetCacheLifetime()`

UnsetCacheLifetime ensures that no value is present for CacheLifetime, not even an explicit nil
### GetOnlyProxy404

`func (o *V2RuleServeStaticAction) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *V2RuleServeStaticAction) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *V2RuleServeStaticAction) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *V2RuleServeStaticAction) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *V2RuleServeStaticAction) GetInjectHeaders() map[string]string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *V2RuleServeStaticAction) GetInjectHeadersOk() (*map[string]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *V2RuleServeStaticAction) SetInjectHeaders(v map[string]string)`

SetInjectHeaders sets InjectHeaders field to given value.

### HasInjectHeaders

`func (o *V2RuleServeStaticAction) HasInjectHeaders() bool`

HasInjectHeaders returns a boolean if a field has been set.

### SetInjectHeadersNil

`func (o *V2RuleServeStaticAction) SetInjectHeadersNil(b bool)`

 SetInjectHeadersNil sets the value for InjectHeaders to be an explicit nil

### UnsetInjectHeaders
`func (o *V2RuleServeStaticAction) UnsetInjectHeaders()`

UnsetInjectHeaders ensures that no value is present for InjectHeaders, not even an explicit nil
### GetProxyStripHeaders

`func (o *V2RuleServeStaticAction) GetProxyStripHeaders() []string`

GetProxyStripHeaders returns the ProxyStripHeaders field if non-nil, zero value otherwise.

### GetProxyStripHeadersOk

`func (o *V2RuleServeStaticAction) GetProxyStripHeadersOk() (*[]string, bool)`

GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripHeaders

`func (o *V2RuleServeStaticAction) SetProxyStripHeaders(v []string)`

SetProxyStripHeaders sets ProxyStripHeaders field to given value.

### HasProxyStripHeaders

`func (o *V2RuleServeStaticAction) HasProxyStripHeaders() bool`

HasProxyStripHeaders returns a boolean if a field has been set.

### GetProxyStripRequestHeaders

`func (o *V2RuleServeStaticAction) GetProxyStripRequestHeaders() []string`

GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field if non-nil, zero value otherwise.

### GetProxyStripRequestHeadersOk

`func (o *V2RuleServeStaticAction) GetProxyStripRequestHeadersOk() (*[]string, bool)`

GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripRequestHeaders

`func (o *V2RuleServeStaticAction) SetProxyStripRequestHeaders(v []string)`

SetProxyStripRequestHeaders sets ProxyStripRequestHeaders field to given value.

### HasProxyStripRequestHeaders

`func (o *V2RuleServeStaticAction) HasProxyStripRequestHeaders() bool`

HasProxyStripRequestHeaders returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *V2RuleServeStaticAction) GetOriginTimeout() string`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *V2RuleServeStaticAction) GetOriginTimeoutOk() (*string, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *V2RuleServeStaticAction) SetOriginTimeout(v string)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *V2RuleServeStaticAction) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetFailoverMode

`func (o *V2RuleServeStaticAction) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *V2RuleServeStaticAction) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *V2RuleServeStaticAction) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *V2RuleServeStaticAction) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *V2RuleServeStaticAction) GetFailoverOriginTtfb() string`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *V2RuleServeStaticAction) GetFailoverOriginTtfbOk() (*string, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *V2RuleServeStaticAction) SetFailoverOriginTtfb(v string)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *V2RuleServeStaticAction) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverOriginStatusCodes

`func (o *V2RuleServeStaticAction) GetFailoverOriginStatusCodes() []string`

GetFailoverOriginStatusCodes returns the FailoverOriginStatusCodes field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodesOk

`func (o *V2RuleServeStaticAction) GetFailoverOriginStatusCodesOk() (*[]string, bool)`

GetFailoverOriginStatusCodesOk returns a tuple with the FailoverOriginStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCodes

`func (o *V2RuleServeStaticAction) SetFailoverOriginStatusCodes(v []string)`

SetFailoverOriginStatusCodes sets FailoverOriginStatusCodes field to given value.

### HasFailoverOriginStatusCodes

`func (o *V2RuleServeStaticAction) HasFailoverOriginStatusCodes() bool`

HasFailoverOriginStatusCodes returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *V2RuleServeStaticAction) GetFailoverLifetime() string`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *V2RuleServeStaticAction) GetFailoverLifetimeOk() (*string, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *V2RuleServeStaticAction) SetFailoverLifetime(v string)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *V2RuleServeStaticAction) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetNotify

`func (o *V2RuleServeStaticAction) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *V2RuleServeStaticAction) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *V2RuleServeStaticAction) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *V2RuleServeStaticAction) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *V2RuleServeStaticAction) GetNotifyConfig() V2RuleProxyActionAllOfNotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *V2RuleServeStaticAction) GetNotifyConfigOk() (*V2RuleProxyActionAllOfNotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *V2RuleServeStaticAction) SetNotifyConfig(v V2RuleProxyActionAllOfNotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *V2RuleServeStaticAction) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

### SetNotifyConfigNil

`func (o *V2RuleServeStaticAction) SetNotifyConfigNil(b bool)`

 SetNotifyConfigNil sets the value for NotifyConfig to be an explicit nil

### UnsetNotifyConfig
`func (o *V2RuleServeStaticAction) UnsetNotifyConfig()`

UnsetNotifyConfig ensures that no value is present for NotifyConfig, not even an explicit nil
### GetWafEnabled

`func (o *V2RuleServeStaticAction) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *V2RuleServeStaticAction) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *V2RuleServeStaticAction) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *V2RuleServeStaticAction) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.

### GetWafConfig

`func (o *V2RuleServeStaticAction) GetWafConfig() WafConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *V2RuleServeStaticAction) GetWafConfigOk() (*WafConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *V2RuleServeStaticAction) SetWafConfig(v WafConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *V2RuleServeStaticAction) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.

### GetProxyAlertEnabled

`func (o *V2RuleServeStaticAction) GetProxyAlertEnabled() bool`

GetProxyAlertEnabled returns the ProxyAlertEnabled field if non-nil, zero value otherwise.

### GetProxyAlertEnabledOk

`func (o *V2RuleServeStaticAction) GetProxyAlertEnabledOk() (*bool, bool)`

GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAlertEnabled

`func (o *V2RuleServeStaticAction) SetProxyAlertEnabled(v bool)`

SetProxyAlertEnabled sets ProxyAlertEnabled field to given value.

### HasProxyAlertEnabled

`func (o *V2RuleServeStaticAction) HasProxyAlertEnabled() bool`

HasProxyAlertEnabled returns a boolean if a field has been set.

### GetProxyInlineFnEnabled

`func (o *V2RuleServeStaticAction) GetProxyInlineFnEnabled() bool`

GetProxyInlineFnEnabled returns the ProxyInlineFnEnabled field if non-nil, zero value otherwise.

### GetProxyInlineFnEnabledOk

`func (o *V2RuleServeStaticAction) GetProxyInlineFnEnabledOk() (*bool, bool)`

GetProxyInlineFnEnabledOk returns a tuple with the ProxyInlineFnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyInlineFnEnabled

`func (o *V2RuleServeStaticAction) SetProxyInlineFnEnabled(v bool)`

SetProxyInlineFnEnabled sets ProxyInlineFnEnabled field to given value.

### HasProxyInlineFnEnabled

`func (o *V2RuleServeStaticAction) HasProxyInlineFnEnabled() bool`

HasProxyInlineFnEnabled returns a boolean if a field has been set.

### GetQuantCloudSelection

`func (o *V2RuleServeStaticAction) GetQuantCloudSelection() V2RuleProxyActionAllOfQuantCloudSelection`

GetQuantCloudSelection returns the QuantCloudSelection field if non-nil, zero value otherwise.

### GetQuantCloudSelectionOk

`func (o *V2RuleServeStaticAction) GetQuantCloudSelectionOk() (*V2RuleProxyActionAllOfQuantCloudSelection, bool)`

GetQuantCloudSelectionOk returns a tuple with the QuantCloudSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantCloudSelection

`func (o *V2RuleServeStaticAction) SetQuantCloudSelection(v V2RuleProxyActionAllOfQuantCloudSelection)`

SetQuantCloudSelection sets QuantCloudSelection field to given value.

### HasQuantCloudSelection

`func (o *V2RuleServeStaticAction) HasQuantCloudSelection() bool`

HasQuantCloudSelection returns a boolean if a field has been set.

### SetQuantCloudSelectionNil

`func (o *V2RuleServeStaticAction) SetQuantCloudSelectionNil(b bool)`

 SetQuantCloudSelectionNil sets the value for QuantCloudSelection to be an explicit nil

### UnsetQuantCloudSelection
`func (o *V2RuleServeStaticAction) UnsetQuantCloudSelection()`

UnsetQuantCloudSelection ensures that no value is present for QuantCloudSelection, not even an explicit nil
### GetStaticFilePath

`func (o *V2RuleServeStaticAction) GetStaticFilePath() string`

GetStaticFilePath returns the StaticFilePath field if non-nil, zero value otherwise.

### GetStaticFilePathOk

`func (o *V2RuleServeStaticAction) GetStaticFilePathOk() (*string, bool)`

GetStaticFilePathOk returns a tuple with the StaticFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticFilePath

`func (o *V2RuleServeStaticAction) SetStaticFilePath(v string)`

SetStaticFilePath sets StaticFilePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


