# WAFConfigUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |  | [optional] 
**ParanoiaLevel** | Pointer to **int32** |  | [optional] [default to 1]
**AllowRules** | Pointer to **[]string** |  | [optional] 
**AllowIp** | Pointer to **[]string** |  | [optional] 
**BlockIp** | Pointer to **[]string** |  | [optional] 
**BlockUa** | Pointer to **[]string** |  | [optional] 
**BlockReferer** | Pointer to **[]string** |  | [optional] 
**Thresholds** | Pointer to [**[]Threshold**](Threshold.md) |  | [optional] 
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

### NewWAFConfigUpdate

`func NewWAFConfigUpdate() *WAFConfigUpdate`

NewWAFConfigUpdate instantiates a new WAFConfigUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFConfigUpdateWithDefaults

`func NewWAFConfigUpdateWithDefaults() *WAFConfigUpdate`

NewWAFConfigUpdateWithDefaults instantiates a new WAFConfigUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *WAFConfigUpdate) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WAFConfigUpdate) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WAFConfigUpdate) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WAFConfigUpdate) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetParanoiaLevel

`func (o *WAFConfigUpdate) GetParanoiaLevel() int32`

GetParanoiaLevel returns the ParanoiaLevel field if non-nil, zero value otherwise.

### GetParanoiaLevelOk

`func (o *WAFConfigUpdate) GetParanoiaLevelOk() (*int32, bool)`

GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParanoiaLevel

`func (o *WAFConfigUpdate) SetParanoiaLevel(v int32)`

SetParanoiaLevel sets ParanoiaLevel field to given value.

### HasParanoiaLevel

`func (o *WAFConfigUpdate) HasParanoiaLevel() bool`

HasParanoiaLevel returns a boolean if a field has been set.

### GetAllowRules

`func (o *WAFConfigUpdate) GetAllowRules() []string`

GetAllowRules returns the AllowRules field if non-nil, zero value otherwise.

### GetAllowRulesOk

`func (o *WAFConfigUpdate) GetAllowRulesOk() (*[]string, bool)`

GetAllowRulesOk returns a tuple with the AllowRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRules

`func (o *WAFConfigUpdate) SetAllowRules(v []string)`

SetAllowRules sets AllowRules field to given value.

### HasAllowRules

`func (o *WAFConfigUpdate) HasAllowRules() bool`

HasAllowRules returns a boolean if a field has been set.

### GetAllowIp

`func (o *WAFConfigUpdate) GetAllowIp() []string`

GetAllowIp returns the AllowIp field if non-nil, zero value otherwise.

### GetAllowIpOk

`func (o *WAFConfigUpdate) GetAllowIpOk() (*[]string, bool)`

GetAllowIpOk returns a tuple with the AllowIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIp

`func (o *WAFConfigUpdate) SetAllowIp(v []string)`

SetAllowIp sets AllowIp field to given value.

### HasAllowIp

`func (o *WAFConfigUpdate) HasAllowIp() bool`

HasAllowIp returns a boolean if a field has been set.

### GetBlockIp

`func (o *WAFConfigUpdate) GetBlockIp() []string`

GetBlockIp returns the BlockIp field if non-nil, zero value otherwise.

### GetBlockIpOk

`func (o *WAFConfigUpdate) GetBlockIpOk() (*[]string, bool)`

GetBlockIpOk returns a tuple with the BlockIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIp

`func (o *WAFConfigUpdate) SetBlockIp(v []string)`

SetBlockIp sets BlockIp field to given value.

### HasBlockIp

`func (o *WAFConfigUpdate) HasBlockIp() bool`

HasBlockIp returns a boolean if a field has been set.

### GetBlockUa

`func (o *WAFConfigUpdate) GetBlockUa() []string`

GetBlockUa returns the BlockUa field if non-nil, zero value otherwise.

### GetBlockUaOk

`func (o *WAFConfigUpdate) GetBlockUaOk() (*[]string, bool)`

GetBlockUaOk returns a tuple with the BlockUa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUa

`func (o *WAFConfigUpdate) SetBlockUa(v []string)`

SetBlockUa sets BlockUa field to given value.

### HasBlockUa

`func (o *WAFConfigUpdate) HasBlockUa() bool`

HasBlockUa returns a boolean if a field has been set.

### GetBlockReferer

`func (o *WAFConfigUpdate) GetBlockReferer() []string`

GetBlockReferer returns the BlockReferer field if non-nil, zero value otherwise.

### GetBlockRefererOk

`func (o *WAFConfigUpdate) GetBlockRefererOk() (*[]string, bool)`

GetBlockRefererOk returns a tuple with the BlockReferer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReferer

`func (o *WAFConfigUpdate) SetBlockReferer(v []string)`

SetBlockReferer sets BlockReferer field to given value.

### HasBlockReferer

`func (o *WAFConfigUpdate) HasBlockReferer() bool`

HasBlockReferer returns a boolean if a field has been set.

### GetThresholds

`func (o *WAFConfigUpdate) GetThresholds() []Threshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *WAFConfigUpdate) GetThresholdsOk() (*[]Threshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *WAFConfigUpdate) SetThresholds(v []Threshold)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *WAFConfigUpdate) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetHttpblEnabled

`func (o *WAFConfigUpdate) GetHttpblEnabled() map[string]bool`

GetHttpblEnabled returns the HttpblEnabled field if non-nil, zero value otherwise.

### GetHttpblEnabledOk

`func (o *WAFConfigUpdate) GetHttpblEnabledOk() (*map[string]bool, bool)`

GetHttpblEnabledOk returns a tuple with the HttpblEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpblEnabled

`func (o *WAFConfigUpdate) SetHttpblEnabled(v map[string]bool)`

SetHttpblEnabled sets HttpblEnabled field to given value.

### HasHttpblEnabled

`func (o *WAFConfigUpdate) HasHttpblEnabled() bool`

HasHttpblEnabled returns a boolean if a field has been set.

### GetNotifySlackHitsRpm

`func (o *WAFConfigUpdate) GetNotifySlackHitsRpm() int32`

GetNotifySlackHitsRpm returns the NotifySlackHitsRpm field if non-nil, zero value otherwise.

### GetNotifySlackHitsRpmOk

`func (o *WAFConfigUpdate) GetNotifySlackHitsRpmOk() (*int32, bool)`

GetNotifySlackHitsRpmOk returns a tuple with the NotifySlackHitsRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlackHitsRpm

`func (o *WAFConfigUpdate) SetNotifySlackHitsRpm(v int32)`

SetNotifySlackHitsRpm sets NotifySlackHitsRpm field to given value.

### HasNotifySlackHitsRpm

`func (o *WAFConfigUpdate) HasNotifySlackHitsRpm() bool`

HasNotifySlackHitsRpm returns a boolean if a field has been set.

### GetIpRatelimitMode

`func (o *WAFConfigUpdate) GetIpRatelimitMode() string`

GetIpRatelimitMode returns the IpRatelimitMode field if non-nil, zero value otherwise.

### GetIpRatelimitModeOk

`func (o *WAFConfigUpdate) GetIpRatelimitModeOk() (*string, bool)`

GetIpRatelimitModeOk returns a tuple with the IpRatelimitMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRatelimitMode

`func (o *WAFConfigUpdate) SetIpRatelimitMode(v string)`

SetIpRatelimitMode sets IpRatelimitMode field to given value.

### HasIpRatelimitMode

`func (o *WAFConfigUpdate) HasIpRatelimitMode() bool`

HasIpRatelimitMode returns a boolean if a field has been set.

### GetIpRatelimitRps

`func (o *WAFConfigUpdate) GetIpRatelimitRps() int32`

GetIpRatelimitRps returns the IpRatelimitRps field if non-nil, zero value otherwise.

### GetIpRatelimitRpsOk

`func (o *WAFConfigUpdate) GetIpRatelimitRpsOk() (*int32, bool)`

GetIpRatelimitRpsOk returns a tuple with the IpRatelimitRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRatelimitRps

`func (o *WAFConfigUpdate) SetIpRatelimitRps(v int32)`

SetIpRatelimitRps sets IpRatelimitRps field to given value.

### HasIpRatelimitRps

`func (o *WAFConfigUpdate) HasIpRatelimitRps() bool`

HasIpRatelimitRps returns a boolean if a field has been set.

### GetIpRatelimitCooldown

`func (o *WAFConfigUpdate) GetIpRatelimitCooldown() int32`

GetIpRatelimitCooldown returns the IpRatelimitCooldown field if non-nil, zero value otherwise.

### GetIpRatelimitCooldownOk

`func (o *WAFConfigUpdate) GetIpRatelimitCooldownOk() (*int32, bool)`

GetIpRatelimitCooldownOk returns a tuple with the IpRatelimitCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRatelimitCooldown

`func (o *WAFConfigUpdate) SetIpRatelimitCooldown(v int32)`

SetIpRatelimitCooldown sets IpRatelimitCooldown field to given value.

### HasIpRatelimitCooldown

`func (o *WAFConfigUpdate) HasIpRatelimitCooldown() bool`

HasIpRatelimitCooldown returns a boolean if a field has been set.

### GetRequestHeaderRatelimitMode

`func (o *WAFConfigUpdate) GetRequestHeaderRatelimitMode() string`

GetRequestHeaderRatelimitMode returns the RequestHeaderRatelimitMode field if non-nil, zero value otherwise.

### GetRequestHeaderRatelimitModeOk

`func (o *WAFConfigUpdate) GetRequestHeaderRatelimitModeOk() (*string, bool)`

GetRequestHeaderRatelimitModeOk returns a tuple with the RequestHeaderRatelimitMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderRatelimitMode

`func (o *WAFConfigUpdate) SetRequestHeaderRatelimitMode(v string)`

SetRequestHeaderRatelimitMode sets RequestHeaderRatelimitMode field to given value.

### HasRequestHeaderRatelimitMode

`func (o *WAFConfigUpdate) HasRequestHeaderRatelimitMode() bool`

HasRequestHeaderRatelimitMode returns a boolean if a field has been set.

### GetRequestHeaderName

`func (o *WAFConfigUpdate) GetRequestHeaderName() string`

GetRequestHeaderName returns the RequestHeaderName field if non-nil, zero value otherwise.

### GetRequestHeaderNameOk

`func (o *WAFConfigUpdate) GetRequestHeaderNameOk() (*string, bool)`

GetRequestHeaderNameOk returns a tuple with the RequestHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderName

`func (o *WAFConfigUpdate) SetRequestHeaderName(v string)`

SetRequestHeaderName sets RequestHeaderName field to given value.

### HasRequestHeaderName

`func (o *WAFConfigUpdate) HasRequestHeaderName() bool`

HasRequestHeaderName returns a boolean if a field has been set.

### GetRequestHeaderRatelimitRps

`func (o *WAFConfigUpdate) GetRequestHeaderRatelimitRps() int32`

GetRequestHeaderRatelimitRps returns the RequestHeaderRatelimitRps field if non-nil, zero value otherwise.

### GetRequestHeaderRatelimitRpsOk

`func (o *WAFConfigUpdate) GetRequestHeaderRatelimitRpsOk() (*int32, bool)`

GetRequestHeaderRatelimitRpsOk returns a tuple with the RequestHeaderRatelimitRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderRatelimitRps

`func (o *WAFConfigUpdate) SetRequestHeaderRatelimitRps(v int32)`

SetRequestHeaderRatelimitRps sets RequestHeaderRatelimitRps field to given value.

### HasRequestHeaderRatelimitRps

`func (o *WAFConfigUpdate) HasRequestHeaderRatelimitRps() bool`

HasRequestHeaderRatelimitRps returns a boolean if a field has been set.

### GetRequestHeaderRatelimitCooldown

`func (o *WAFConfigUpdate) GetRequestHeaderRatelimitCooldown() int32`

GetRequestHeaderRatelimitCooldown returns the RequestHeaderRatelimitCooldown field if non-nil, zero value otherwise.

### GetRequestHeaderRatelimitCooldownOk

`func (o *WAFConfigUpdate) GetRequestHeaderRatelimitCooldownOk() (*int32, bool)`

GetRequestHeaderRatelimitCooldownOk returns a tuple with the RequestHeaderRatelimitCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderRatelimitCooldown

`func (o *WAFConfigUpdate) SetRequestHeaderRatelimitCooldown(v int32)`

SetRequestHeaderRatelimitCooldown sets RequestHeaderRatelimitCooldown field to given value.

### HasRequestHeaderRatelimitCooldown

`func (o *WAFConfigUpdate) HasRequestHeaderRatelimitCooldown() bool`

HasRequestHeaderRatelimitCooldown returns a boolean if a field has been set.

### GetWafRatelimitMode

`func (o *WAFConfigUpdate) GetWafRatelimitMode() string`

GetWafRatelimitMode returns the WafRatelimitMode field if non-nil, zero value otherwise.

### GetWafRatelimitModeOk

`func (o *WAFConfigUpdate) GetWafRatelimitModeOk() (*string, bool)`

GetWafRatelimitModeOk returns a tuple with the WafRatelimitMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRatelimitMode

`func (o *WAFConfigUpdate) SetWafRatelimitMode(v string)`

SetWafRatelimitMode sets WafRatelimitMode field to given value.

### HasWafRatelimitMode

`func (o *WAFConfigUpdate) HasWafRatelimitMode() bool`

HasWafRatelimitMode returns a boolean if a field has been set.

### GetWafRatelimitHits

`func (o *WAFConfigUpdate) GetWafRatelimitHits() int32`

GetWafRatelimitHits returns the WafRatelimitHits field if non-nil, zero value otherwise.

### GetWafRatelimitHitsOk

`func (o *WAFConfigUpdate) GetWafRatelimitHitsOk() (*int32, bool)`

GetWafRatelimitHitsOk returns a tuple with the WafRatelimitHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRatelimitHits

`func (o *WAFConfigUpdate) SetWafRatelimitHits(v int32)`

SetWafRatelimitHits sets WafRatelimitHits field to given value.

### HasWafRatelimitHits

`func (o *WAFConfigUpdate) HasWafRatelimitHits() bool`

HasWafRatelimitHits returns a boolean if a field has been set.

### GetWafRatelimitRps

`func (o *WAFConfigUpdate) GetWafRatelimitRps() int32`

GetWafRatelimitRps returns the WafRatelimitRps field if non-nil, zero value otherwise.

### GetWafRatelimitRpsOk

`func (o *WAFConfigUpdate) GetWafRatelimitRpsOk() (*int32, bool)`

GetWafRatelimitRpsOk returns a tuple with the WafRatelimitRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRatelimitRps

`func (o *WAFConfigUpdate) SetWafRatelimitRps(v int32)`

SetWafRatelimitRps sets WafRatelimitRps field to given value.

### HasWafRatelimitRps

`func (o *WAFConfigUpdate) HasWafRatelimitRps() bool`

HasWafRatelimitRps returns a boolean if a field has been set.

### GetWafRatelimitCooldown

`func (o *WAFConfigUpdate) GetWafRatelimitCooldown() int32`

GetWafRatelimitCooldown returns the WafRatelimitCooldown field if non-nil, zero value otherwise.

### GetWafRatelimitCooldownOk

`func (o *WAFConfigUpdate) GetWafRatelimitCooldownOk() (*int32, bool)`

GetWafRatelimitCooldownOk returns a tuple with the WafRatelimitCooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafRatelimitCooldown

`func (o *WAFConfigUpdate) SetWafRatelimitCooldown(v int32)`

SetWafRatelimitCooldown sets WafRatelimitCooldown field to given value.

### HasWafRatelimitCooldown

`func (o *WAFConfigUpdate) HasWafRatelimitCooldown() bool`

HasWafRatelimitCooldown returns a boolean if a field has been set.

### GetNotifyEmail

`func (o *WAFConfigUpdate) GetNotifyEmail() []string`

GetNotifyEmail returns the NotifyEmail field if non-nil, zero value otherwise.

### GetNotifyEmailOk

`func (o *WAFConfigUpdate) GetNotifyEmailOk() (*[]string, bool)`

GetNotifyEmailOk returns a tuple with the NotifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmail

`func (o *WAFConfigUpdate) SetNotifyEmail(v []string)`

SetNotifyEmail sets NotifyEmail field to given value.

### HasNotifyEmail

`func (o *WAFConfigUpdate) HasNotifyEmail() bool`

HasNotifyEmail returns a boolean if a field has been set.

### GetNotifySlack

`func (o *WAFConfigUpdate) GetNotifySlack() string`

GetNotifySlack returns the NotifySlack field if non-nil, zero value otherwise.

### GetNotifySlackOk

`func (o *WAFConfigUpdate) GetNotifySlackOk() (*string, bool)`

GetNotifySlackOk returns a tuple with the NotifySlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlack

`func (o *WAFConfigUpdate) SetNotifySlack(v string)`

SetNotifySlack sets NotifySlack field to given value.

### HasNotifySlack

`func (o *WAFConfigUpdate) HasNotifySlack() bool`

HasNotifySlack returns a boolean if a field has been set.

### GetNotifySlackRpm

`func (o *WAFConfigUpdate) GetNotifySlackRpm() int32`

GetNotifySlackRpm returns the NotifySlackRpm field if non-nil, zero value otherwise.

### GetNotifySlackRpmOk

`func (o *WAFConfigUpdate) GetNotifySlackRpmOk() (*int32, bool)`

GetNotifySlackRpmOk returns a tuple with the NotifySlackRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlackRpm

`func (o *WAFConfigUpdate) SetNotifySlackRpm(v int32)`

SetNotifySlackRpm sets NotifySlackRpm field to given value.

### HasNotifySlackRpm

`func (o *WAFConfigUpdate) HasNotifySlackRpm() bool`

HasNotifySlackRpm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


