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
**BlockLists** | Pointer to [**WafConfigBlockLists**](WafConfigBlockLists.md) |  | [optional] 
**Httpbl** | Pointer to [**ProxyConfigHttpbl**](ProxyConfigHttpbl.md) |  | [optional] 
**Thresholds** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HttpblEnabled** | Pointer to **map[string]bool** |  | [optional] 
**NotifySlackHitsRpm** | Pointer to **int32** |  | [optional] 
**IpRatelimitMode** | Pointer to **string** |  | [optional] [default to "disabled"]
**IpRatelimitRps** | Pointer to **int32** |  | [optional] [default to 5]
**IpRatelimitCooldown** | Pointer to **int32** |  | [optional] [default to 30]
**RequestHeaderRatelimitMode** | Pointer to **string** |  | [optional] [default to "disabled"]
**RequestHeaderName** | Pointer to **string** |  | [optional] 
**RequestHeaderRatelimitRps** | Pointer to **int32** |  | [optional] [default to 5]
**RequestHeaderRatelimitCooldown** | Pointer to **int32** |  | [optional] [default to 30]
**WafRatelimitMode** | Pointer to **string** |  | [optional] [default to "disabled"]
**WafRatelimitHits** | Pointer to **int32** |  | [optional] [default to 10]
**WafRatelimitRps** | Pointer to **int32** |  | [optional] [default to 5]
**WafRatelimitCooldown** | Pointer to **int32** |  | [optional] [default to 300]
**NotifyEmail** | Pointer to **[]string** |  | [optional] 
**NotifySlack** | Pointer to **string** |  | [optional] 
**NotifySlackRpm** | Pointer to **int32** |  | [optional] 

## Methods

### NewRuleWAFConfig

`func NewRuleWAFConfig(mode string, ) *RuleWAFConfig`

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

### GetBlockLists

`func (o *RuleWAFConfig) GetBlockLists() WafConfigBlockLists`

GetBlockLists returns the BlockLists field if non-nil, zero value otherwise.

### GetBlockListsOk

`func (o *RuleWAFConfig) GetBlockListsOk() (*WafConfigBlockLists, bool)`

GetBlockListsOk returns a tuple with the BlockLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLists

`func (o *RuleWAFConfig) SetBlockLists(v WafConfigBlockLists)`

SetBlockLists sets BlockLists field to given value.

### HasBlockLists

`func (o *RuleWAFConfig) HasBlockLists() bool`

HasBlockLists returns a boolean if a field has been set.

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

### GetThresholds

`func (o *RuleWAFConfig) GetThresholds() []map[string]interface{}`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *RuleWAFConfig) GetThresholdsOk() (*[]map[string]interface{}, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *RuleWAFConfig) SetThresholds(v []map[string]interface{})`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *RuleWAFConfig) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetHttpblEnabled

`func (o *RuleWAFConfig) GetHttpblEnabled() map[string]bool`

GetHttpblEnabled returns the HttpblEnabled field if non-nil, zero value otherwise.

### GetHttpblEnabledOk

`func (o *RuleWAFConfig) GetHttpblEnabledOk() (*map[string]bool, bool)`

GetHttpblEnabledOk returns a tuple with the HttpblEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpblEnabled

`func (o *RuleWAFConfig) SetHttpblEnabled(v map[string]bool)`

SetHttpblEnabled sets HttpblEnabled field to given value.

### HasHttpblEnabled

`func (o *RuleWAFConfig) HasHttpblEnabled() bool`

HasHttpblEnabled returns a boolean if a field has been set.

### GetNotifySlackHitsRpm

`func (o *RuleWAFConfig) GetNotifySlackHitsRpm() int32`

GetNotifySlackHitsRpm returns the NotifySlackHitsRpm field if non-nil, zero value otherwise.

### GetNotifySlackHitsRpmOk

`func (o *RuleWAFConfig) GetNotifySlackHitsRpmOk() (*int32, bool)`

GetNotifySlackHitsRpmOk returns a tuple with the NotifySlackHitsRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlackHitsRpm

`func (o *RuleWAFConfig) SetNotifySlackHitsRpm(v int32)`

SetNotifySlackHitsRpm sets NotifySlackHitsRpm field to given value.

### HasNotifySlackHitsRpm

`func (o *RuleWAFConfig) HasNotifySlackHitsRpm() bool`

HasNotifySlackHitsRpm returns a boolean if a field has been set.

### GetIpRatelimitMode

`func (o *RuleWAFConfig) GetIpRatelimitMode() string`

GetIpRatelimitMode returns the IpRatelimitMode field if non-nil, zero value otherwise.

### GetIpRatelimitModeOk

`func (o *RuleWAFConfig) GetIpRatelimitModeOk() (*string, bool)`

GetIpRatelimitModeOk returns a tuple with the IpRatelimitMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRatelimitMode

`func (o *RuleWAFConfig) SetIpRatelimitMode(v string)`

SetIpRatelimitMode sets IpRatelimitMode field to given value.

### HasIpRatelimitMode

`func (o *RuleWAFConfig) HasIpRatelimitMode() bool`

HasIpRatelimitMode returns a boolean if a field has been set.

### GetIpRatelimitRps

`func (o *RuleWAFConfig) GetIpRatelimitRps() int32`

GetIpRatelimitRps returns the IpRatelimitRps field if non-nil, zero value otherwise.

### GetIpRatelimitRpsOk

`func (o *RuleWAFConfig) GetIpRatelimitRpsOk() (*int32, bool)`

GetIpRatelimitRpsOk returns a tuple with the IpRatelimitRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRatelimitRps

`func (o *RuleWAFConfig) SetIpRatelimitRps(v int32)`

SetIpRatelimitRps sets IpRatelimitRps field to given value.

### HasIpRatelimitRps

`func (o *RuleWAFConfig) HasIpRatelimitRps() bool`

HasIpRatelimitRps returns a boolean if a field has been set.

### GetIpRatelimitCooldown

`func (o *RuleWAFConfig) GetIpRatelimitCooldown() int32`

GetIpRatelimitCooldown returns the IpRatelimitCooldown field if non-nil, zero value otherwise.

### GetIpRatelimitCooldownOk

`func (o *RuleWAFConfig) GetIpRatelimitCooldownOk() (*int32, bool)`

GetIpRatelimitCooldownOk returns a tuple with the IpRatelimitCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRatelimitCooldown

`func (o *RuleWAFConfig) SetIpRatelimitCooldown(v int32)`

SetIpRatelimitCooldown sets IpRatelimitCooldown field to given value.

### HasIpRatelimitCooldown

`func (o *RuleWAFConfig) HasIpRatelimitCooldown() bool`

HasIpRatelimitCooldown returns a boolean if a field has been set.

### GetRequestHeaderRatelimitMode

`func (o *RuleWAFConfig) GetRequestHeaderRatelimitMode() string`

GetRequestHeaderRatelimitMode returns the RequestHeaderRatelimitMode field if non-nil, zero value otherwise.

### GetRequestHeaderRatelimitModeOk

`func (o *RuleWAFConfig) GetRequestHeaderRatelimitModeOk() (*string, bool)`

GetRequestHeaderRatelimitModeOk returns a tuple with the RequestHeaderRatelimitMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderRatelimitMode

`func (o *RuleWAFConfig) SetRequestHeaderRatelimitMode(v string)`

SetRequestHeaderRatelimitMode sets RequestHeaderRatelimitMode field to given value.

### HasRequestHeaderRatelimitMode

`func (o *RuleWAFConfig) HasRequestHeaderRatelimitMode() bool`

HasRequestHeaderRatelimitMode returns a boolean if a field has been set.

### GetRequestHeaderName

`func (o *RuleWAFConfig) GetRequestHeaderName() string`

GetRequestHeaderName returns the RequestHeaderName field if non-nil, zero value otherwise.

### GetRequestHeaderNameOk

`func (o *RuleWAFConfig) GetRequestHeaderNameOk() (*string, bool)`

GetRequestHeaderNameOk returns a tuple with the RequestHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderName

`func (o *RuleWAFConfig) SetRequestHeaderName(v string)`

SetRequestHeaderName sets RequestHeaderName field to given value.

### HasRequestHeaderName

`func (o *RuleWAFConfig) HasRequestHeaderName() bool`

HasRequestHeaderName returns a boolean if a field has been set.

### GetRequestHeaderRatelimitRps

`func (o *RuleWAFConfig) GetRequestHeaderRatelimitRps() int32`

GetRequestHeaderRatelimitRps returns the RequestHeaderRatelimitRps field if non-nil, zero value otherwise.

### GetRequestHeaderRatelimitRpsOk

`func (o *RuleWAFConfig) GetRequestHeaderRatelimitRpsOk() (*int32, bool)`

GetRequestHeaderRatelimitRpsOk returns a tuple with the RequestHeaderRatelimitRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderRatelimitRps

`func (o *RuleWAFConfig) SetRequestHeaderRatelimitRps(v int32)`

SetRequestHeaderRatelimitRps sets RequestHeaderRatelimitRps field to given value.

### HasRequestHeaderRatelimitRps

`func (o *RuleWAFConfig) HasRequestHeaderRatelimitRps() bool`

HasRequestHeaderRatelimitRps returns a boolean if a field has been set.

### GetRequestHeaderRatelimitCooldown

`func (o *RuleWAFConfig) GetRequestHeaderRatelimitCooldown() int32`

GetRequestHeaderRatelimitCooldown returns the RequestHeaderRatelimitCooldown field if non-nil, zero value otherwise.

### GetRequestHeaderRatelimitCooldownOk

`func (o *RuleWAFConfig) GetRequestHeaderRatelimitCooldownOk() (*int32, bool)`

GetRequestHeaderRatelimitCooldownOk returns a tuple with the RequestHeaderRatelimitCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderRatelimitCooldown

`func (o *RuleWAFConfig) SetRequestHeaderRatelimitCooldown(v int32)`

SetRequestHeaderRatelimitCooldown sets RequestHeaderRatelimitCooldown field to given value.

### HasRequestHeaderRatelimitCooldown

`func (o *RuleWAFConfig) HasRequestHeaderRatelimitCooldown() bool`

HasRequestHeaderRatelimitCooldown returns a boolean if a field has been set.

### GetWafRatelimitMode

`func (o *RuleWAFConfig) GetWafRatelimitMode() string`

GetWafRatelimitMode returns the WafRatelimitMode field if non-nil, zero value otherwise.

### GetWafRatelimitModeOk

`func (o *RuleWAFConfig) GetWafRatelimitModeOk() (*string, bool)`

GetWafRatelimitModeOk returns a tuple with the WafRatelimitMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRatelimitMode

`func (o *RuleWAFConfig) SetWafRatelimitMode(v string)`

SetWafRatelimitMode sets WafRatelimitMode field to given value.

### HasWafRatelimitMode

`func (o *RuleWAFConfig) HasWafRatelimitMode() bool`

HasWafRatelimitMode returns a boolean if a field has been set.

### GetWafRatelimitHits

`func (o *RuleWAFConfig) GetWafRatelimitHits() int32`

GetWafRatelimitHits returns the WafRatelimitHits field if non-nil, zero value otherwise.

### GetWafRatelimitHitsOk

`func (o *RuleWAFConfig) GetWafRatelimitHitsOk() (*int32, bool)`

GetWafRatelimitHitsOk returns a tuple with the WafRatelimitHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRatelimitHits

`func (o *RuleWAFConfig) SetWafRatelimitHits(v int32)`

SetWafRatelimitHits sets WafRatelimitHits field to given value.

### HasWafRatelimitHits

`func (o *RuleWAFConfig) HasWafRatelimitHits() bool`

HasWafRatelimitHits returns a boolean if a field has been set.

### GetWafRatelimitRps

`func (o *RuleWAFConfig) GetWafRatelimitRps() int32`

GetWafRatelimitRps returns the WafRatelimitRps field if non-nil, zero value otherwise.

### GetWafRatelimitRpsOk

`func (o *RuleWAFConfig) GetWafRatelimitRpsOk() (*int32, bool)`

GetWafRatelimitRpsOk returns a tuple with the WafRatelimitRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRatelimitRps

`func (o *RuleWAFConfig) SetWafRatelimitRps(v int32)`

SetWafRatelimitRps sets WafRatelimitRps field to given value.

### HasWafRatelimitRps

`func (o *RuleWAFConfig) HasWafRatelimitRps() bool`

HasWafRatelimitRps returns a boolean if a field has been set.

### GetWafRatelimitCooldown

`func (o *RuleWAFConfig) GetWafRatelimitCooldown() int32`

GetWafRatelimitCooldown returns the WafRatelimitCooldown field if non-nil, zero value otherwise.

### GetWafRatelimitCooldownOk

`func (o *RuleWAFConfig) GetWafRatelimitCooldownOk() (*int32, bool)`

GetWafRatelimitCooldownOk returns a tuple with the WafRatelimitCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRatelimitCooldown

`func (o *RuleWAFConfig) SetWafRatelimitCooldown(v int32)`

SetWafRatelimitCooldown sets WafRatelimitCooldown field to given value.

### HasWafRatelimitCooldown

`func (o *RuleWAFConfig) HasWafRatelimitCooldown() bool`

HasWafRatelimitCooldown returns a boolean if a field has been set.

### GetNotifyEmail

`func (o *RuleWAFConfig) GetNotifyEmail() []string`

GetNotifyEmail returns the NotifyEmail field if non-nil, zero value otherwise.

### GetNotifyEmailOk

`func (o *RuleWAFConfig) GetNotifyEmailOk() (*[]string, bool)`

GetNotifyEmailOk returns a tuple with the NotifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmail

`func (o *RuleWAFConfig) SetNotifyEmail(v []string)`

SetNotifyEmail sets NotifyEmail field to given value.

### HasNotifyEmail

`func (o *RuleWAFConfig) HasNotifyEmail() bool`

HasNotifyEmail returns a boolean if a field has been set.

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

### GetNotifySlackRpm

`func (o *RuleWAFConfig) GetNotifySlackRpm() int32`

GetNotifySlackRpm returns the NotifySlackRpm field if non-nil, zero value otherwise.

### GetNotifySlackRpmOk

`func (o *RuleWAFConfig) GetNotifySlackRpmOk() (*int32, bool)`

GetNotifySlackRpmOk returns a tuple with the NotifySlackRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlackRpm

`func (o *RuleWAFConfig) SetNotifySlackRpm(v int32)`

SetNotifySlackRpm sets NotifySlackRpm field to given value.

### HasNotifySlackRpm

`func (o *RuleWAFConfig) HasNotifySlackRpm() bool`

HasNotifySlackRpm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


