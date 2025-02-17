# RuleProxyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **[]string** |  | [default to ["any"]]
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] [default to 0]
**Disabled** | **bool** |  | [default to false]
**Url** | **[]string** |  | 
**Country** | Pointer to **string** |  | [optional] 
**CountryIs** | Pointer to **[]string** |  | [optional] 
**CountryIsNot** | Pointer to **[]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**MethodIs** | Pointer to **[]string** |  | [optional] 
**MethodIsNot** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IpIs** | Pointer to **[]string** |  | [optional] 
**IpIsNot** | Pointer to **[]string** |  | [optional] 
**To** | **string** |  | 
**Host** | Pointer to **string** |  | [optional] 
**AuthUser** | Pointer to **string** |  | [optional] [default to ""]
**AuthPass** | Pointer to **string** |  | [optional] [default to ""]
**DisableSslVerify** | Pointer to **bool** |  | [optional] [default to false]
**CacheLifetime** | Pointer to **NullableInt32** |  | [optional] 
**OnlyProxy404** | Pointer to **bool** |  | [optional] [default to false]
**InjectHeaders** | Pointer to **map[string]string** |  | [optional] 
**ProxyStripHeaders** | Pointer to **[]string** |  | [optional] 
**ProxyStripRequestHeaders** | Pointer to **[]string** |  | [optional] 
**FailoverMode** | Pointer to **bool** |  | [optional] [default to false]
**FailoverOriginTtfb** | Pointer to **string** |  | [optional] [default to "2000"]
**FailoverOriginStatusCodes** | Pointer to **[]string** |  | [optional] [default to ["200","404","301","302","304"]]
**FailoverLifetime** | Pointer to **string** |  | [optional] [default to "300"]
**Notify** | Pointer to **string** |  | [optional] [default to "none"]
**NotifyConfig** | Pointer to [**NotifyConfig**](NotifyConfig.md) |  | [optional] 
**WafEnabled** | Pointer to **bool** |  | [optional] [default to false]
**WafConfig** | Pointer to [**WAFConfig**](WAFConfig.md) |  | [optional] 

## Methods

### NewRuleProxyRequest

`func NewRuleProxyRequest(domain []string, disabled bool, url []string, to string, ) *RuleProxyRequest`

NewRuleProxyRequest instantiates a new RuleProxyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleProxyRequestWithDefaults

`func NewRuleProxyRequestWithDefaults() *RuleProxyRequest`

NewRuleProxyRequestWithDefaults instantiates a new RuleProxyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleProxyRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleProxyRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleProxyRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *RuleProxyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleProxyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleProxyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleProxyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleProxyRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleProxyRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleProxyRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleProxyRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *RuleProxyRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RuleProxyRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RuleProxyRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RuleProxyRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleProxyRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleProxyRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleProxyRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetUrl

`func (o *RuleProxyRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleProxyRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleProxyRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *RuleProxyRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleProxyRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleProxyRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleProxyRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleProxyRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleProxyRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleProxyRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleProxyRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleProxyRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleProxyRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleProxyRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleProxyRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleProxyRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleProxyRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleProxyRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleProxyRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleProxyRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleProxyRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleProxyRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleProxyRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleProxyRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleProxyRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleProxyRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleProxyRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleProxyRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleProxyRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleProxyRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleProxyRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleProxyRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleProxyRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleProxyRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleProxyRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleProxyRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleProxyRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleProxyRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleProxyRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetTo

`func (o *RuleProxyRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RuleProxyRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RuleProxyRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetHost

`func (o *RuleProxyRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RuleProxyRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RuleProxyRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RuleProxyRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAuthUser

`func (o *RuleProxyRequest) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *RuleProxyRequest) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *RuleProxyRequest) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *RuleProxyRequest) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetAuthPass

`func (o *RuleProxyRequest) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *RuleProxyRequest) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *RuleProxyRequest) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *RuleProxyRequest) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *RuleProxyRequest) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *RuleProxyRequest) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *RuleProxyRequest) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *RuleProxyRequest) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *RuleProxyRequest) GetCacheLifetime() int32`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *RuleProxyRequest) GetCacheLifetimeOk() (*int32, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *RuleProxyRequest) SetCacheLifetime(v int32)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *RuleProxyRequest) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### SetCacheLifetimeNil

`func (o *RuleProxyRequest) SetCacheLifetimeNil(b bool)`

 SetCacheLifetimeNil sets the value for CacheLifetime to be an explicit nil

### UnsetCacheLifetime
`func (o *RuleProxyRequest) UnsetCacheLifetime()`

UnsetCacheLifetime ensures that no value is present for CacheLifetime, not even an explicit nil
### GetOnlyProxy404

`func (o *RuleProxyRequest) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *RuleProxyRequest) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *RuleProxyRequest) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *RuleProxyRequest) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *RuleProxyRequest) GetInjectHeaders() map[string]string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *RuleProxyRequest) GetInjectHeadersOk() (*map[string]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *RuleProxyRequest) SetInjectHeaders(v map[string]string)`

SetInjectHeaders sets InjectHeaders field to given value.

### HasInjectHeaders

`func (o *RuleProxyRequest) HasInjectHeaders() bool`

HasInjectHeaders returns a boolean if a field has been set.

### SetInjectHeadersNil

`func (o *RuleProxyRequest) SetInjectHeadersNil(b bool)`

 SetInjectHeadersNil sets the value for InjectHeaders to be an explicit nil

### UnsetInjectHeaders
`func (o *RuleProxyRequest) UnsetInjectHeaders()`

UnsetInjectHeaders ensures that no value is present for InjectHeaders, not even an explicit nil
### GetProxyStripHeaders

`func (o *RuleProxyRequest) GetProxyStripHeaders() []string`

GetProxyStripHeaders returns the ProxyStripHeaders field if non-nil, zero value otherwise.

### GetProxyStripHeadersOk

`func (o *RuleProxyRequest) GetProxyStripHeadersOk() (*[]string, bool)`

GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripHeaders

`func (o *RuleProxyRequest) SetProxyStripHeaders(v []string)`

SetProxyStripHeaders sets ProxyStripHeaders field to given value.

### HasProxyStripHeaders

`func (o *RuleProxyRequest) HasProxyStripHeaders() bool`

HasProxyStripHeaders returns a boolean if a field has been set.

### GetProxyStripRequestHeaders

`func (o *RuleProxyRequest) GetProxyStripRequestHeaders() []string`

GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field if non-nil, zero value otherwise.

### GetProxyStripRequestHeadersOk

`func (o *RuleProxyRequest) GetProxyStripRequestHeadersOk() (*[]string, bool)`

GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripRequestHeaders

`func (o *RuleProxyRequest) SetProxyStripRequestHeaders(v []string)`

SetProxyStripRequestHeaders sets ProxyStripRequestHeaders field to given value.

### HasProxyStripRequestHeaders

`func (o *RuleProxyRequest) HasProxyStripRequestHeaders() bool`

HasProxyStripRequestHeaders returns a boolean if a field has been set.

### GetFailoverMode

`func (o *RuleProxyRequest) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *RuleProxyRequest) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *RuleProxyRequest) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *RuleProxyRequest) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *RuleProxyRequest) GetFailoverOriginTtfb() string`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *RuleProxyRequest) GetFailoverOriginTtfbOk() (*string, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *RuleProxyRequest) SetFailoverOriginTtfb(v string)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *RuleProxyRequest) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverOriginStatusCodes

`func (o *RuleProxyRequest) GetFailoverOriginStatusCodes() []string`

GetFailoverOriginStatusCodes returns the FailoverOriginStatusCodes field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodesOk

`func (o *RuleProxyRequest) GetFailoverOriginStatusCodesOk() (*[]string, bool)`

GetFailoverOriginStatusCodesOk returns a tuple with the FailoverOriginStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCodes

`func (o *RuleProxyRequest) SetFailoverOriginStatusCodes(v []string)`

SetFailoverOriginStatusCodes sets FailoverOriginStatusCodes field to given value.

### HasFailoverOriginStatusCodes

`func (o *RuleProxyRequest) HasFailoverOriginStatusCodes() bool`

HasFailoverOriginStatusCodes returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *RuleProxyRequest) GetFailoverLifetime() string`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *RuleProxyRequest) GetFailoverLifetimeOk() (*string, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *RuleProxyRequest) SetFailoverLifetime(v string)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *RuleProxyRequest) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetNotify

`func (o *RuleProxyRequest) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *RuleProxyRequest) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *RuleProxyRequest) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *RuleProxyRequest) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *RuleProxyRequest) GetNotifyConfig() NotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *RuleProxyRequest) GetNotifyConfigOk() (*NotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *RuleProxyRequest) SetNotifyConfig(v NotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *RuleProxyRequest) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

### GetWafEnabled

`func (o *RuleProxyRequest) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *RuleProxyRequest) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *RuleProxyRequest) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *RuleProxyRequest) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.

### GetWafConfig

`func (o *RuleProxyRequest) GetWafConfig() WAFConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *RuleProxyRequest) GetWafConfigOk() (*WAFConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *RuleProxyRequest) SetWafConfig(v WAFConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *RuleProxyRequest) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


