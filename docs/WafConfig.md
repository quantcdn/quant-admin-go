# WafConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | WAF operation mode | [optional] [default to "report"]
**ParanoiaLevel** | Pointer to **int32** | OWASP paranoia level | [optional] [default to 1]
**AllowRules** | Pointer to **[]string** | WAF rule IDs to allow/whitelist | [optional] 
**AllowIp** | Pointer to **[]string** | IP addresses to allow | [optional] 
**BlockIp** | Pointer to **[]string** | IP addresses to block | [optional] 
**BlockAsn** | Pointer to **[]string** | ASN numbers to block | [optional] 
**BlockUa** | Pointer to **[]string** | User agent patterns to block | [optional] 
**BlockReferer** | Pointer to **[]string** | Referer patterns to block | [optional] 
**NotifySlack** | Pointer to **string** | Slack webhook URL for notifications | [optional] 
**NotifySlackHitsRpm** | Pointer to **int32** | Minimum hits per minute to trigger Slack notification | [optional] 
**NotifyEmail** | Pointer to **[]string** | Email addresses for notifications | [optional] 
**Httpbl** | Pointer to [**WafConfigHttpbl**](WafConfigHttpbl.md) |  | [optional] 
**BlockLists** | Pointer to [**WafConfigBlockLists**](WafConfigBlockLists.md) |  | [optional] 
**Thresholds** | Pointer to [**[]WafConfigThresholdsInner**](WafConfigThresholdsInner.md) | Rate limiting thresholds | [optional] 

## Methods

### NewWafConfig

`func NewWafConfig() *WafConfig`

NewWafConfig instantiates a new WafConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafConfigWithDefaults

`func NewWafConfigWithDefaults() *WafConfig`

NewWafConfigWithDefaults instantiates a new WafConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *WafConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WafConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WafConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WafConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetParanoiaLevel

`func (o *WafConfig) GetParanoiaLevel() int32`

GetParanoiaLevel returns the ParanoiaLevel field if non-nil, zero value otherwise.

### GetParanoiaLevelOk

`func (o *WafConfig) GetParanoiaLevelOk() (*int32, bool)`

GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParanoiaLevel

`func (o *WafConfig) SetParanoiaLevel(v int32)`

SetParanoiaLevel sets ParanoiaLevel field to given value.

### HasParanoiaLevel

`func (o *WafConfig) HasParanoiaLevel() bool`

HasParanoiaLevel returns a boolean if a field has been set.

### GetAllowRules

`func (o *WafConfig) GetAllowRules() []string`

GetAllowRules returns the AllowRules field if non-nil, zero value otherwise.

### GetAllowRulesOk

`func (o *WafConfig) GetAllowRulesOk() (*[]string, bool)`

GetAllowRulesOk returns a tuple with the AllowRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRules

`func (o *WafConfig) SetAllowRules(v []string)`

SetAllowRules sets AllowRules field to given value.

### HasAllowRules

`func (o *WafConfig) HasAllowRules() bool`

HasAllowRules returns a boolean if a field has been set.

### GetAllowIp

`func (o *WafConfig) GetAllowIp() []string`

GetAllowIp returns the AllowIp field if non-nil, zero value otherwise.

### GetAllowIpOk

`func (o *WafConfig) GetAllowIpOk() (*[]string, bool)`

GetAllowIpOk returns a tuple with the AllowIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIp

`func (o *WafConfig) SetAllowIp(v []string)`

SetAllowIp sets AllowIp field to given value.

### HasAllowIp

`func (o *WafConfig) HasAllowIp() bool`

HasAllowIp returns a boolean if a field has been set.

### GetBlockIp

`func (o *WafConfig) GetBlockIp() []string`

GetBlockIp returns the BlockIp field if non-nil, zero value otherwise.

### GetBlockIpOk

`func (o *WafConfig) GetBlockIpOk() (*[]string, bool)`

GetBlockIpOk returns a tuple with the BlockIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIp

`func (o *WafConfig) SetBlockIp(v []string)`

SetBlockIp sets BlockIp field to given value.

### HasBlockIp

`func (o *WafConfig) HasBlockIp() bool`

HasBlockIp returns a boolean if a field has been set.

### GetBlockAsn

`func (o *WafConfig) GetBlockAsn() []string`

GetBlockAsn returns the BlockAsn field if non-nil, zero value otherwise.

### GetBlockAsnOk

`func (o *WafConfig) GetBlockAsnOk() (*[]string, bool)`

GetBlockAsnOk returns a tuple with the BlockAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAsn

`func (o *WafConfig) SetBlockAsn(v []string)`

SetBlockAsn sets BlockAsn field to given value.

### HasBlockAsn

`func (o *WafConfig) HasBlockAsn() bool`

HasBlockAsn returns a boolean if a field has been set.

### GetBlockUa

`func (o *WafConfig) GetBlockUa() []string`

GetBlockUa returns the BlockUa field if non-nil, zero value otherwise.

### GetBlockUaOk

`func (o *WafConfig) GetBlockUaOk() (*[]string, bool)`

GetBlockUaOk returns a tuple with the BlockUa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUa

`func (o *WafConfig) SetBlockUa(v []string)`

SetBlockUa sets BlockUa field to given value.

### HasBlockUa

`func (o *WafConfig) HasBlockUa() bool`

HasBlockUa returns a boolean if a field has been set.

### GetBlockReferer

`func (o *WafConfig) GetBlockReferer() []string`

GetBlockReferer returns the BlockReferer field if non-nil, zero value otherwise.

### GetBlockRefererOk

`func (o *WafConfig) GetBlockRefererOk() (*[]string, bool)`

GetBlockRefererOk returns a tuple with the BlockReferer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReferer

`func (o *WafConfig) SetBlockReferer(v []string)`

SetBlockReferer sets BlockReferer field to given value.

### HasBlockReferer

`func (o *WafConfig) HasBlockReferer() bool`

HasBlockReferer returns a boolean if a field has been set.

### GetNotifySlack

`func (o *WafConfig) GetNotifySlack() string`

GetNotifySlack returns the NotifySlack field if non-nil, zero value otherwise.

### GetNotifySlackOk

`func (o *WafConfig) GetNotifySlackOk() (*string, bool)`

GetNotifySlackOk returns a tuple with the NotifySlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlack

`func (o *WafConfig) SetNotifySlack(v string)`

SetNotifySlack sets NotifySlack field to given value.

### HasNotifySlack

`func (o *WafConfig) HasNotifySlack() bool`

HasNotifySlack returns a boolean if a field has been set.

### GetNotifySlackHitsRpm

`func (o *WafConfig) GetNotifySlackHitsRpm() int32`

GetNotifySlackHitsRpm returns the NotifySlackHitsRpm field if non-nil, zero value otherwise.

### GetNotifySlackHitsRpmOk

`func (o *WafConfig) GetNotifySlackHitsRpmOk() (*int32, bool)`

GetNotifySlackHitsRpmOk returns a tuple with the NotifySlackHitsRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySlackHitsRpm

`func (o *WafConfig) SetNotifySlackHitsRpm(v int32)`

SetNotifySlackHitsRpm sets NotifySlackHitsRpm field to given value.

### HasNotifySlackHitsRpm

`func (o *WafConfig) HasNotifySlackHitsRpm() bool`

HasNotifySlackHitsRpm returns a boolean if a field has been set.

### GetNotifyEmail

`func (o *WafConfig) GetNotifyEmail() []string`

GetNotifyEmail returns the NotifyEmail field if non-nil, zero value otherwise.

### GetNotifyEmailOk

`func (o *WafConfig) GetNotifyEmailOk() (*[]string, bool)`

GetNotifyEmailOk returns a tuple with the NotifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmail

`func (o *WafConfig) SetNotifyEmail(v []string)`

SetNotifyEmail sets NotifyEmail field to given value.

### HasNotifyEmail

`func (o *WafConfig) HasNotifyEmail() bool`

HasNotifyEmail returns a boolean if a field has been set.

### GetHttpbl

`func (o *WafConfig) GetHttpbl() WafConfigHttpbl`

GetHttpbl returns the Httpbl field if non-nil, zero value otherwise.

### GetHttpblOk

`func (o *WafConfig) GetHttpblOk() (*WafConfigHttpbl, bool)`

GetHttpblOk returns a tuple with the Httpbl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpbl

`func (o *WafConfig) SetHttpbl(v WafConfigHttpbl)`

SetHttpbl sets Httpbl field to given value.

### HasHttpbl

`func (o *WafConfig) HasHttpbl() bool`

HasHttpbl returns a boolean if a field has been set.

### GetBlockLists

`func (o *WafConfig) GetBlockLists() WafConfigBlockLists`

GetBlockLists returns the BlockLists field if non-nil, zero value otherwise.

### GetBlockListsOk

`func (o *WafConfig) GetBlockListsOk() (*WafConfigBlockLists, bool)`

GetBlockListsOk returns a tuple with the BlockLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockLists

`func (o *WafConfig) SetBlockLists(v WafConfigBlockLists)`

SetBlockLists sets BlockLists field to given value.

### HasBlockLists

`func (o *WafConfig) HasBlockLists() bool`

HasBlockLists returns a boolean if a field has been set.

### GetThresholds

`func (o *WafConfig) GetThresholds() []WafConfigThresholdsInner`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *WafConfig) GetThresholdsOk() (*[]WafConfigThresholdsInner, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *WafConfig) SetThresholds(v []WafConfigThresholdsInner)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *WafConfig) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


