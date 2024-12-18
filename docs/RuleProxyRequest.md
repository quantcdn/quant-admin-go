# RuleProxyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **[]string** |  | [default to ["any"]]
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
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
**OnlyWithCookie** | Pointer to **bool** |  | [optional] [default to false]
**CookieName** | Pointer to **string** |  | [optional] 
**Proxy** | [**ProxyConfig**](ProxyConfig.md) |  | 
**Failover** | Pointer to [**FailoverConfig**](FailoverConfig.md) |  | [optional] 
**Notify** | Pointer to **string** |  | [optional] [default to "none"]
**NotifyConfig** | Pointer to [**NotifyConfig**](NotifyConfig.md) |  | [optional] 
**WafEnabled** | Pointer to **bool** |  | [optional] [default to false]
**WafConfig** | Pointer to [**WAFConfig**](WAFConfig.md) |  | [optional] 

## Methods

### NewRuleProxyRequest

`func NewRuleProxyRequest(domain []string, disabled bool, url []string, proxy ProxyConfig, ) *RuleProxyRequest`

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

### GetOnlyWithCookie

`func (o *RuleProxyRequest) GetOnlyWithCookie() bool`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *RuleProxyRequest) GetOnlyWithCookieOk() (*bool, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *RuleProxyRequest) SetOnlyWithCookie(v bool)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *RuleProxyRequest) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetCookieName

`func (o *RuleProxyRequest) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *RuleProxyRequest) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *RuleProxyRequest) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *RuleProxyRequest) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetProxy

`func (o *RuleProxyRequest) GetProxy() ProxyConfig`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *RuleProxyRequest) GetProxyOk() (*ProxyConfig, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *RuleProxyRequest) SetProxy(v ProxyConfig)`

SetProxy sets Proxy field to given value.


### GetFailover

`func (o *RuleProxyRequest) GetFailover() FailoverConfig`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *RuleProxyRequest) GetFailoverOk() (*FailoverConfig, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *RuleProxyRequest) SetFailover(v FailoverConfig)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *RuleProxyRequest) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

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


