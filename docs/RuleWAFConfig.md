# RuleWAFConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** |  | 
**ParanoiaLevel** | Pointer to **int32** |  | [optional] [default to 1]
**AllowRules** | Pointer to **[]string** |  | [optional] 
**AllowIp** | Pointer to **[]string** |  | [optional] 
**BlockIp** | Pointer to **[]string** |  | [optional] 
**BlockUa** | Pointer to **[]string** |  | [optional] 
**BlockReferer** | Pointer to **[]string** |  | [optional] 
**NotifySlack** | Pointer to **string** |  | [optional] 
**NotifySlackHistRpm** | Pointer to **int32** |  | [optional] 
**NotifyEmail** | **string** |  | 
**Httpbl** | Pointer to [**ProxyConfigHttpbl**](ProxyConfigHttpbl.md) |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** |  | [optional] 
**OriginTimeout** | **int32** |  | [default to 30]
**FailoverMode** | Pointer to **bool** |  | [optional] 
**FailoverOriginTtfb** | Pointer to **int32** |  | [optional] 
**FailoverOriginStatusCode** | Pointer to **int32** |  | [optional] 
**FailoverLifetime** | Pointer to **int32** |  | [optional] 
**Notify** | Pointer to **string** |  | [optional] 
**NotifyConfig** | Pointer to [**ProxyNotifyConfig**](ProxyNotifyConfig.md) |  | [optional] 
**InjectHeaders** | **[]string** |  | 

## Methods

### NewRuleWAFConfig

`func NewRuleWAFConfig(mode string, notifyEmail string, originTimeout int32, injectHeaders []string, ) *RuleWAFConfig`

NewRuleWAFConfig instantiates a new RuleWAFConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWAFConfigWithDefaults

`func NewRuleWAFConfigWithDefaults() *RuleWAFConfig`

NewRuleWAFConfigWithDefaults instantiates a new RuleWAFConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RuleWAFConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RuleWAFConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RuleWAFConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### GetParanoiaLevel

`func (o *RuleWAFConfig) GetParanoiaLevel() int32`

GetParanoiaLevel returns the ParanoiaLevel field if non-nil, zero value otherwise.

### GetParanoiaLevelOk

`func (o *RuleWAFConfig) GetParanoiaLevelOk() (*int32, bool)`

GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParanoiaLevel

`func (o *RuleWAFConfig) SetParanoiaLevel(v int32)`

SetParanoiaLevel sets ParanoiaLevel field to given value.

### HasParanoiaLevel

`func (o *RuleWAFConfig) HasParanoiaLevel() bool`

HasParanoiaLevel returns a boolean if a field has been set.

### GetAllowRules

`func (o *RuleWAFConfig) GetAllowRules() []string`

GetAllowRules returns the AllowRules field if non-nil, zero value otherwise.

### GetAllowRulesOk

`func (o *RuleWAFConfig) GetAllowRulesOk() (*[]string, bool)`

GetAllowRulesOk returns a tuple with the AllowRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRules

`func (o *RuleWAFConfig) SetAllowRules(v []string)`

SetAllowRules sets AllowRules field to given value.

### HasAllowRules

`func (o *RuleWAFConfig) HasAllowRules() bool`

HasAllowRules returns a boolean if a field has been set.

### GetAllowIp

`func (o *RuleWAFConfig) GetAllowIp() []string`

GetAllowIp returns the AllowIp field if non-nil, zero value otherwise.

### GetAllowIpOk

`func (o *RuleWAFConfig) GetAllowIpOk() (*[]string, bool)`

GetAllowIpOk returns a tuple with the AllowIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIp

`func (o *RuleWAFConfig) SetAllowIp(v []string)`

SetAllowIp sets AllowIp field to given value.

### HasAllowIp

`func (o *RuleWAFConfig) HasAllowIp() bool`

HasAllowIp returns a boolean if a field has been set.

### GetBlockIp

`func (o *RuleWAFConfig) GetBlockIp() []string`

GetBlockIp returns the BlockIp field if non-nil, zero value otherwise.

### GetBlockIpOk

`func (o *RuleWAFConfig) GetBlockIpOk() (*[]string, bool)`

GetBlockIpOk returns a tuple with the BlockIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIp

`func (o *RuleWAFConfig) SetBlockIp(v []string)`

SetBlockIp sets BlockIp field to given value.

### HasBlockIp

`func (o *RuleWAFConfig) HasBlockIp() bool`

HasBlockIp returns a boolean if a field has been set.

### GetBlockUa

`func (o *RuleWAFConfig) GetBlockUa() []string`

GetBlockUa returns the BlockUa field if non-nil, zero value otherwise.

### GetBlockUaOk

`func (o *RuleWAFConfig) GetBlockUaOk() (*[]string, bool)`

GetBlockUaOk returns a tuple with the BlockUa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUa

`func (o *RuleWAFConfig) SetBlockUa(v []string)`

SetBlockUa sets BlockUa field to given value.

### HasBlockUa

`func (o *RuleWAFConfig) HasBlockUa() bool`

HasBlockUa returns a boolean if a field has been set.

### GetBlockReferer

`func (o *RuleWAFConfig) GetBlockReferer() []string`

GetBlockReferer returns the BlockReferer field if non-nil, zero value otherwise.

### GetBlockRefererOk

`func (o *RuleWAFConfig) GetBlockRefererOk() (*[]string, bool)`

GetBlockRefererOk returns a tuple with the BlockReferer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReferer

`func (o *RuleWAFConfig) SetBlockReferer(v []string)`

SetBlockReferer sets BlockReferer field to given value.

### HasBlockReferer

`func (o *RuleWAFConfig) HasBlockReferer() bool`

HasBlockReferer returns a boolean if a field has been set.

### GetNotifySlack

`func (o *RuleWAFConfig) GetNotifySlack() string`

GetNotifySlack returns the NotifySlack field if non-nil, zero value otherwise.

### GetNotifySlackOk

`func (o *RuleWAFConfig) GetNotifySlackOk() (*string, bool)`

GetNotifySlackOk returns a tuple with the NotifySlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlack

`func (o *RuleWAFConfig) SetNotifySlack(v string)`

SetNotifySlack sets NotifySlack field to given value.

### HasNotifySlack

`func (o *RuleWAFConfig) HasNotifySlack() bool`

HasNotifySlack returns a boolean if a field has been set.

### GetNotifySlackHistRpm

`func (o *RuleWAFConfig) GetNotifySlackHistRpm() int32`

GetNotifySlackHistRpm returns the NotifySlackHistRpm field if non-nil, zero value otherwise.

### GetNotifySlackHistRpmOk

`func (o *RuleWAFConfig) GetNotifySlackHistRpmOk() (*int32, bool)`

GetNotifySlackHistRpmOk returns a tuple with the NotifySlackHistRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlackHistRpm

`func (o *RuleWAFConfig) SetNotifySlackHistRpm(v int32)`

SetNotifySlackHistRpm sets NotifySlackHistRpm field to given value.

### HasNotifySlackHistRpm

`func (o *RuleWAFConfig) HasNotifySlackHistRpm() bool`

HasNotifySlackHistRpm returns a boolean if a field has been set.

### GetNotifyEmail

`func (o *RuleWAFConfig) GetNotifyEmail() string`

GetNotifyEmail returns the NotifyEmail field if non-nil, zero value otherwise.

### GetNotifyEmailOk

`func (o *RuleWAFConfig) GetNotifyEmailOk() (*string, bool)`

GetNotifyEmailOk returns a tuple with the NotifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmail

`func (o *RuleWAFConfig) SetNotifyEmail(v string)`

SetNotifyEmail sets NotifyEmail field to given value.


### GetHttpbl

`func (o *RuleWAFConfig) GetHttpbl() ProxyConfigHttpbl`

GetHttpbl returns the Httpbl field if non-nil, zero value otherwise.

### GetHttpblOk

`func (o *RuleWAFConfig) GetHttpblOk() (*ProxyConfigHttpbl, bool)`

GetHttpblOk returns a tuple with the Httpbl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpbl

`func (o *RuleWAFConfig) SetHttpbl(v ProxyConfigHttpbl)`

SetHttpbl sets Httpbl field to given value.

### HasHttpbl

`func (o *RuleWAFConfig) HasHttpbl() bool`

HasHttpbl returns a boolean if a field has been set.

### GetProxyAlertEnabled

`func (o *RuleWAFConfig) GetProxyAlertEnabled() bool`

GetProxyAlertEnabled returns the ProxyAlertEnabled field if non-nil, zero value otherwise.

### GetProxyAlertEnabledOk

`func (o *RuleWAFConfig) GetProxyAlertEnabledOk() (*bool, bool)`

GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAlertEnabled

`func (o *RuleWAFConfig) SetProxyAlertEnabled(v bool)`

SetProxyAlertEnabled sets ProxyAlertEnabled field to given value.

### HasProxyAlertEnabled

`func (o *RuleWAFConfig) HasProxyAlertEnabled() bool`

HasProxyAlertEnabled returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *RuleWAFConfig) GetOriginTimeout() int32`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *RuleWAFConfig) GetOriginTimeoutOk() (*int32, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *RuleWAFConfig) SetOriginTimeout(v int32)`

SetOriginTimeout sets OriginTimeout field to given value.


### GetFailoverMode

`func (o *RuleWAFConfig) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *RuleWAFConfig) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *RuleWAFConfig) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *RuleWAFConfig) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *RuleWAFConfig) GetFailoverOriginTtfb() int32`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *RuleWAFConfig) GetFailoverOriginTtfbOk() (*int32, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *RuleWAFConfig) SetFailoverOriginTtfb(v int32)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *RuleWAFConfig) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverOriginStatusCode

`func (o *RuleWAFConfig) GetFailoverOriginStatusCode() int32`

GetFailoverOriginStatusCode returns the FailoverOriginStatusCode field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodeOk

`func (o *RuleWAFConfig) GetFailoverOriginStatusCodeOk() (*int32, bool)`

GetFailoverOriginStatusCodeOk returns a tuple with the FailoverOriginStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCode

`func (o *RuleWAFConfig) SetFailoverOriginStatusCode(v int32)`

SetFailoverOriginStatusCode sets FailoverOriginStatusCode field to given value.

### HasFailoverOriginStatusCode

`func (o *RuleWAFConfig) HasFailoverOriginStatusCode() bool`

HasFailoverOriginStatusCode returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *RuleWAFConfig) GetFailoverLifetime() int32`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *RuleWAFConfig) GetFailoverLifetimeOk() (*int32, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *RuleWAFConfig) SetFailoverLifetime(v int32)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *RuleWAFConfig) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetNotify

`func (o *RuleWAFConfig) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *RuleWAFConfig) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *RuleWAFConfig) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *RuleWAFConfig) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *RuleWAFConfig) GetNotifyConfig() ProxyNotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *RuleWAFConfig) GetNotifyConfigOk() (*ProxyNotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *RuleWAFConfig) SetNotifyConfig(v ProxyNotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *RuleWAFConfig) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *RuleWAFConfig) GetInjectHeaders() []string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *RuleWAFConfig) GetInjectHeadersOk() (*[]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *RuleWAFConfig) SetInjectHeaders(v []string)`

SetInjectHeaders sets InjectHeaders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


