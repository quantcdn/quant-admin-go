# RuleProxyRequestWafConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**ParanoiaLevel** | Pointer to **int32** |  | [optional] 
**AllowRules** | Pointer to **[]int32** |  | [optional] 
**AllowIp** | Pointer to **[]string** |  | [optional] 
**BlockIp** | Pointer to **[]string** |  | [optional] 
**BlockUa** | Pointer to **[]string** |  | [optional] 
**BlockReferer** | Pointer to **[]string** |  | [optional] 
**NotifySlack** | Pointer to **string** |  | [optional] 
**NotifySlackHitsRpm** | Pointer to **int32** |  | [optional] 
**NotifyEmail** | Pointer to **string** |  | [optional] 
**Httpbl** | Pointer to [**RuleProxyRequestWafConfigHttpbl**](RuleProxyRequestWafConfigHttpbl.md) |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** |  | [optional] 
**OriginTimeout** | Pointer to **int32** |  | [optional] 
**FailoverMode** | Pointer to **bool** |  | [optional] 
**FailoverOriginTtfb** | Pointer to **int32** |  | [optional] 
**FailoverOriginStatusCode** | Pointer to **int32** |  | [optional] 
**FailoverLifetime** | Pointer to **int32** |  | [optional] 
**Notify** | Pointer to **string** |  | [optional] 
**NotifyConfig** | Pointer to [**RuleProxyRequestWafConfigNotifyConfig**](RuleProxyRequestWafConfigNotifyConfig.md) |  | [optional] 
**InjectHeaders** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRuleProxyRequestWafConfig

`func NewRuleProxyRequestWafConfig() *RuleProxyRequestWafConfig`

NewRuleProxyRequestWafConfig instantiates a new RuleProxyRequestWafConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleProxyRequestWafConfigWithDefaults

`func NewRuleProxyRequestWafConfigWithDefaults() *RuleProxyRequestWafConfig`

NewRuleProxyRequestWafConfigWithDefaults instantiates a new RuleProxyRequestWafConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RuleProxyRequestWafConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RuleProxyRequestWafConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RuleProxyRequestWafConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RuleProxyRequestWafConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetParanoiaLevel

`func (o *RuleProxyRequestWafConfig) GetParanoiaLevel() int32`

GetParanoiaLevel returns the ParanoiaLevel field if non-nil, zero value otherwise.

### GetParanoiaLevelOk

`func (o *RuleProxyRequestWafConfig) GetParanoiaLevelOk() (*int32, bool)`

GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParanoiaLevel

`func (o *RuleProxyRequestWafConfig) SetParanoiaLevel(v int32)`

SetParanoiaLevel sets ParanoiaLevel field to given value.

### HasParanoiaLevel

`func (o *RuleProxyRequestWafConfig) HasParanoiaLevel() bool`

HasParanoiaLevel returns a boolean if a field has been set.

### GetAllowRules

`func (o *RuleProxyRequestWafConfig) GetAllowRules() []int32`

GetAllowRules returns the AllowRules field if non-nil, zero value otherwise.

### GetAllowRulesOk

`func (o *RuleProxyRequestWafConfig) GetAllowRulesOk() (*[]int32, bool)`

GetAllowRulesOk returns a tuple with the AllowRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRules

`func (o *RuleProxyRequestWafConfig) SetAllowRules(v []int32)`

SetAllowRules sets AllowRules field to given value.

### HasAllowRules

`func (o *RuleProxyRequestWafConfig) HasAllowRules() bool`

HasAllowRules returns a boolean if a field has been set.

### GetAllowIp

`func (o *RuleProxyRequestWafConfig) GetAllowIp() []string`

GetAllowIp returns the AllowIp field if non-nil, zero value otherwise.

### GetAllowIpOk

`func (o *RuleProxyRequestWafConfig) GetAllowIpOk() (*[]string, bool)`

GetAllowIpOk returns a tuple with the AllowIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIp

`func (o *RuleProxyRequestWafConfig) SetAllowIp(v []string)`

SetAllowIp sets AllowIp field to given value.

### HasAllowIp

`func (o *RuleProxyRequestWafConfig) HasAllowIp() bool`

HasAllowIp returns a boolean if a field has been set.

### GetBlockIp

`func (o *RuleProxyRequestWafConfig) GetBlockIp() []string`

GetBlockIp returns the BlockIp field if non-nil, zero value otherwise.

### GetBlockIpOk

`func (o *RuleProxyRequestWafConfig) GetBlockIpOk() (*[]string, bool)`

GetBlockIpOk returns a tuple with the BlockIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIp

`func (o *RuleProxyRequestWafConfig) SetBlockIp(v []string)`

SetBlockIp sets BlockIp field to given value.

### HasBlockIp

`func (o *RuleProxyRequestWafConfig) HasBlockIp() bool`

HasBlockIp returns a boolean if a field has been set.

### GetBlockUa

`func (o *RuleProxyRequestWafConfig) GetBlockUa() []string`

GetBlockUa returns the BlockUa field if non-nil, zero value otherwise.

### GetBlockUaOk

`func (o *RuleProxyRequestWafConfig) GetBlockUaOk() (*[]string, bool)`

GetBlockUaOk returns a tuple with the BlockUa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUa

`func (o *RuleProxyRequestWafConfig) SetBlockUa(v []string)`

SetBlockUa sets BlockUa field to given value.

### HasBlockUa

`func (o *RuleProxyRequestWafConfig) HasBlockUa() bool`

HasBlockUa returns a boolean if a field has been set.

### GetBlockReferer

`func (o *RuleProxyRequestWafConfig) GetBlockReferer() []string`

GetBlockReferer returns the BlockReferer field if non-nil, zero value otherwise.

### GetBlockRefererOk

`func (o *RuleProxyRequestWafConfig) GetBlockRefererOk() (*[]string, bool)`

GetBlockRefererOk returns a tuple with the BlockReferer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReferer

`func (o *RuleProxyRequestWafConfig) SetBlockReferer(v []string)`

SetBlockReferer sets BlockReferer field to given value.

### HasBlockReferer

`func (o *RuleProxyRequestWafConfig) HasBlockReferer() bool`

HasBlockReferer returns a boolean if a field has been set.

### GetNotifySlack

`func (o *RuleProxyRequestWafConfig) GetNotifySlack() string`

GetNotifySlack returns the NotifySlack field if non-nil, zero value otherwise.

### GetNotifySlackOk

`func (o *RuleProxyRequestWafConfig) GetNotifySlackOk() (*string, bool)`

GetNotifySlackOk returns a tuple with the NotifySlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlack

`func (o *RuleProxyRequestWafConfig) SetNotifySlack(v string)`

SetNotifySlack sets NotifySlack field to given value.

### HasNotifySlack

`func (o *RuleProxyRequestWafConfig) HasNotifySlack() bool`

HasNotifySlack returns a boolean if a field has been set.

### GetNotifySlackHitsRpm

`func (o *RuleProxyRequestWafConfig) GetNotifySlackHitsRpm() int32`

GetNotifySlackHitsRpm returns the NotifySlackHitsRpm field if non-nil, zero value otherwise.

### GetNotifySlackHitsRpmOk

`func (o *RuleProxyRequestWafConfig) GetNotifySlackHitsRpmOk() (*int32, bool)`

GetNotifySlackHitsRpmOk returns a tuple with the NotifySlackHitsRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlackHitsRpm

`func (o *RuleProxyRequestWafConfig) SetNotifySlackHitsRpm(v int32)`

SetNotifySlackHitsRpm sets NotifySlackHitsRpm field to given value.

### HasNotifySlackHitsRpm

`func (o *RuleProxyRequestWafConfig) HasNotifySlackHitsRpm() bool`

HasNotifySlackHitsRpm returns a boolean if a field has been set.

### GetNotifyEmail

`func (o *RuleProxyRequestWafConfig) GetNotifyEmail() string`

GetNotifyEmail returns the NotifyEmail field if non-nil, zero value otherwise.

### GetNotifyEmailOk

`func (o *RuleProxyRequestWafConfig) GetNotifyEmailOk() (*string, bool)`

GetNotifyEmailOk returns a tuple with the NotifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmail

`func (o *RuleProxyRequestWafConfig) SetNotifyEmail(v string)`

SetNotifyEmail sets NotifyEmail field to given value.

### HasNotifyEmail

`func (o *RuleProxyRequestWafConfig) HasNotifyEmail() bool`

HasNotifyEmail returns a boolean if a field has been set.

### GetHttpbl

`func (o *RuleProxyRequestWafConfig) GetHttpbl() RuleProxyRequestWafConfigHttpbl`

GetHttpbl returns the Httpbl field if non-nil, zero value otherwise.

### GetHttpblOk

`func (o *RuleProxyRequestWafConfig) GetHttpblOk() (*RuleProxyRequestWafConfigHttpbl, bool)`

GetHttpblOk returns a tuple with the Httpbl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpbl

`func (o *RuleProxyRequestWafConfig) SetHttpbl(v RuleProxyRequestWafConfigHttpbl)`

SetHttpbl sets Httpbl field to given value.

### HasHttpbl

`func (o *RuleProxyRequestWafConfig) HasHttpbl() bool`

HasHttpbl returns a boolean if a field has been set.

### GetProxyAlertEnabled

`func (o *RuleProxyRequestWafConfig) GetProxyAlertEnabled() bool`

GetProxyAlertEnabled returns the ProxyAlertEnabled field if non-nil, zero value otherwise.

### GetProxyAlertEnabledOk

`func (o *RuleProxyRequestWafConfig) GetProxyAlertEnabledOk() (*bool, bool)`

GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAlertEnabled

`func (o *RuleProxyRequestWafConfig) SetProxyAlertEnabled(v bool)`

SetProxyAlertEnabled sets ProxyAlertEnabled field to given value.

### HasProxyAlertEnabled

`func (o *RuleProxyRequestWafConfig) HasProxyAlertEnabled() bool`

HasProxyAlertEnabled returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *RuleProxyRequestWafConfig) GetOriginTimeout() int32`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *RuleProxyRequestWafConfig) GetOriginTimeoutOk() (*int32, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *RuleProxyRequestWafConfig) SetOriginTimeout(v int32)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *RuleProxyRequestWafConfig) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetFailoverMode

`func (o *RuleProxyRequestWafConfig) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *RuleProxyRequestWafConfig) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *RuleProxyRequestWafConfig) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *RuleProxyRequestWafConfig) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *RuleProxyRequestWafConfig) GetFailoverOriginTtfb() int32`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *RuleProxyRequestWafConfig) GetFailoverOriginTtfbOk() (*int32, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *RuleProxyRequestWafConfig) SetFailoverOriginTtfb(v int32)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *RuleProxyRequestWafConfig) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverOriginStatusCode

`func (o *RuleProxyRequestWafConfig) GetFailoverOriginStatusCode() int32`

GetFailoverOriginStatusCode returns the FailoverOriginStatusCode field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodeOk

`func (o *RuleProxyRequestWafConfig) GetFailoverOriginStatusCodeOk() (*int32, bool)`

GetFailoverOriginStatusCodeOk returns a tuple with the FailoverOriginStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCode

`func (o *RuleProxyRequestWafConfig) SetFailoverOriginStatusCode(v int32)`

SetFailoverOriginStatusCode sets FailoverOriginStatusCode field to given value.

### HasFailoverOriginStatusCode

`func (o *RuleProxyRequestWafConfig) HasFailoverOriginStatusCode() bool`

HasFailoverOriginStatusCode returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *RuleProxyRequestWafConfig) GetFailoverLifetime() int32`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *RuleProxyRequestWafConfig) GetFailoverLifetimeOk() (*int32, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *RuleProxyRequestWafConfig) SetFailoverLifetime(v int32)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *RuleProxyRequestWafConfig) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetNotify

`func (o *RuleProxyRequestWafConfig) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *RuleProxyRequestWafConfig) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *RuleProxyRequestWafConfig) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *RuleProxyRequestWafConfig) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *RuleProxyRequestWafConfig) GetNotifyConfig() RuleProxyRequestWafConfigNotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *RuleProxyRequestWafConfig) GetNotifyConfigOk() (*RuleProxyRequestWafConfigNotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *RuleProxyRequestWafConfig) SetNotifyConfig(v RuleProxyRequestWafConfigNotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *RuleProxyRequestWafConfig) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *RuleProxyRequestWafConfig) GetInjectHeaders() []string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *RuleProxyRequestWafConfig) GetInjectHeadersOk() (*[]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *RuleProxyRequestWafConfig) SetInjectHeaders(v []string)`

SetInjectHeaders sets InjectHeaders field to given value.

### HasInjectHeaders

`func (o *RuleProxyRequestWafConfig) HasInjectHeaders() bool`

HasInjectHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


