# RuleProxyRequestCreate

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
**CacheLifetime** | Pointer to **NullableString** |  | [optional] 
**OnlyProxy404** | Pointer to **bool** |  | [optional] [default to false]
**InjectHeaders** | Pointer to **map[string]string** |  | [optional] 
**ProxyStripHeaders** | Pointer to **[]string** |  | [optional] 
**ProxyStripRequestHeaders** | Pointer to **[]string** |  | [optional] 
**OriginTimeout** | Pointer to **int32** |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** |  | [optional] 
**StaticErrorPage** | Pointer to **string** |  | [optional] 
**StaticErrorPageStatusCodes** | Pointer to **[]string** |  | [optional] 
**FailoverMode** | Pointer to **bool** |  | [optional] [default to false]
**FailoverOriginTtfb** | Pointer to **string** |  | [optional] [default to "2000"]
**FailoverOriginStatusCodes** | Pointer to **[]string** |  | [optional] [default to ["200","404","301","302","304"]]
**FailoverLifetime** | Pointer to **string** |  | [optional] [default to "300"]
**Notify** | Pointer to **string** |  | [optional] [default to "none"]
**NotifyConfig** | Pointer to [**NotifyConfig**](NotifyConfig.md) |  | [optional] 
**WafEnabled** | Pointer to **bool** |  | [optional] [default to false]
**WafConfig** | Pointer to [**WAFConfig**](WAFConfig.md) |  | [optional] 
**ApplicationProxy** | Pointer to **bool** |  | [optional] [default to false]
**ApplicationName** | Pointer to **string** |  | [optional] 
**ApplicationEnvironment** | Pointer to **string** |  | [optional] 
**ApplicationContainer** | Pointer to **string** |  | [optional] 
**ApplicationPort** | Pointer to **int32** |  | [optional] 

## Methods

### NewRuleProxyRequestCreate

`func NewRuleProxyRequestCreate(domain []string, disabled bool, url []string, to string, ) *RuleProxyRequestCreate`

NewRuleProxyRequestCreate instantiates a new RuleProxyRequestCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleProxyRequestCreateWithDefaults

`func NewRuleProxyRequestCreateWithDefaults() *RuleProxyRequestCreate`

NewRuleProxyRequestCreateWithDefaults instantiates a new RuleProxyRequestCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleProxyRequestCreate) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleProxyRequestCreate) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleProxyRequestCreate) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *RuleProxyRequestCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleProxyRequestCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleProxyRequestCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleProxyRequestCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleProxyRequestCreate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleProxyRequestCreate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleProxyRequestCreate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleProxyRequestCreate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *RuleProxyRequestCreate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RuleProxyRequestCreate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RuleProxyRequestCreate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RuleProxyRequestCreate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleProxyRequestCreate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleProxyRequestCreate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleProxyRequestCreate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetUrl

`func (o *RuleProxyRequestCreate) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleProxyRequestCreate) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleProxyRequestCreate) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *RuleProxyRequestCreate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleProxyRequestCreate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleProxyRequestCreate) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleProxyRequestCreate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleProxyRequestCreate) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleProxyRequestCreate) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleProxyRequestCreate) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleProxyRequestCreate) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleProxyRequestCreate) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleProxyRequestCreate) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleProxyRequestCreate) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleProxyRequestCreate) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleProxyRequestCreate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleProxyRequestCreate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleProxyRequestCreate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleProxyRequestCreate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleProxyRequestCreate) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleProxyRequestCreate) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleProxyRequestCreate) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleProxyRequestCreate) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleProxyRequestCreate) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleProxyRequestCreate) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleProxyRequestCreate) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleProxyRequestCreate) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleProxyRequestCreate) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleProxyRequestCreate) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleProxyRequestCreate) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleProxyRequestCreate) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleProxyRequestCreate) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleProxyRequestCreate) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleProxyRequestCreate) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleProxyRequestCreate) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleProxyRequestCreate) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleProxyRequestCreate) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleProxyRequestCreate) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleProxyRequestCreate) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetTo

`func (o *RuleProxyRequestCreate) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RuleProxyRequestCreate) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RuleProxyRequestCreate) SetTo(v string)`

SetTo sets To field to given value.


### GetHost

`func (o *RuleProxyRequestCreate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RuleProxyRequestCreate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RuleProxyRequestCreate) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RuleProxyRequestCreate) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAuthUser

`func (o *RuleProxyRequestCreate) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *RuleProxyRequestCreate) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *RuleProxyRequestCreate) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *RuleProxyRequestCreate) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetAuthPass

`func (o *RuleProxyRequestCreate) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *RuleProxyRequestCreate) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *RuleProxyRequestCreate) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *RuleProxyRequestCreate) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *RuleProxyRequestCreate) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *RuleProxyRequestCreate) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *RuleProxyRequestCreate) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *RuleProxyRequestCreate) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *RuleProxyRequestCreate) GetCacheLifetime() string`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *RuleProxyRequestCreate) GetCacheLifetimeOk() (*string, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *RuleProxyRequestCreate) SetCacheLifetime(v string)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *RuleProxyRequestCreate) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### SetCacheLifetimeNil

`func (o *RuleProxyRequestCreate) SetCacheLifetimeNil(b bool)`

 SetCacheLifetimeNil sets the value for CacheLifetime to be an explicit nil

### UnsetCacheLifetime
`func (o *RuleProxyRequestCreate) UnsetCacheLifetime()`

UnsetCacheLifetime ensures that no value is present for CacheLifetime, not even an explicit nil
### GetOnlyProxy404

`func (o *RuleProxyRequestCreate) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *RuleProxyRequestCreate) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *RuleProxyRequestCreate) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *RuleProxyRequestCreate) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *RuleProxyRequestCreate) GetInjectHeaders() map[string]string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *RuleProxyRequestCreate) GetInjectHeadersOk() (*map[string]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *RuleProxyRequestCreate) SetInjectHeaders(v map[string]string)`

SetInjectHeaders sets InjectHeaders field to given value.

### HasInjectHeaders

`func (o *RuleProxyRequestCreate) HasInjectHeaders() bool`

HasInjectHeaders returns a boolean if a field has been set.

### SetInjectHeadersNil

`func (o *RuleProxyRequestCreate) SetInjectHeadersNil(b bool)`

 SetInjectHeadersNil sets the value for InjectHeaders to be an explicit nil

### UnsetInjectHeaders
`func (o *RuleProxyRequestCreate) UnsetInjectHeaders()`

UnsetInjectHeaders ensures that no value is present for InjectHeaders, not even an explicit nil
### GetProxyStripHeaders

`func (o *RuleProxyRequestCreate) GetProxyStripHeaders() []string`

GetProxyStripHeaders returns the ProxyStripHeaders field if non-nil, zero value otherwise.

### GetProxyStripHeadersOk

`func (o *RuleProxyRequestCreate) GetProxyStripHeadersOk() (*[]string, bool)`

GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripHeaders

`func (o *RuleProxyRequestCreate) SetProxyStripHeaders(v []string)`

SetProxyStripHeaders sets ProxyStripHeaders field to given value.

### HasProxyStripHeaders

`func (o *RuleProxyRequestCreate) HasProxyStripHeaders() bool`

HasProxyStripHeaders returns a boolean if a field has been set.

### GetProxyStripRequestHeaders

`func (o *RuleProxyRequestCreate) GetProxyStripRequestHeaders() []string`

GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field if non-nil, zero value otherwise.

### GetProxyStripRequestHeadersOk

`func (o *RuleProxyRequestCreate) GetProxyStripRequestHeadersOk() (*[]string, bool)`

GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripRequestHeaders

`func (o *RuleProxyRequestCreate) SetProxyStripRequestHeaders(v []string)`

SetProxyStripRequestHeaders sets ProxyStripRequestHeaders field to given value.

### HasProxyStripRequestHeaders

`func (o *RuleProxyRequestCreate) HasProxyStripRequestHeaders() bool`

HasProxyStripRequestHeaders returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *RuleProxyRequestCreate) GetOriginTimeout() int32`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *RuleProxyRequestCreate) GetOriginTimeoutOk() (*int32, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *RuleProxyRequestCreate) SetOriginTimeout(v int32)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *RuleProxyRequestCreate) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetProxyAlertEnabled

`func (o *RuleProxyRequestCreate) GetProxyAlertEnabled() bool`

GetProxyAlertEnabled returns the ProxyAlertEnabled field if non-nil, zero value otherwise.

### GetProxyAlertEnabledOk

`func (o *RuleProxyRequestCreate) GetProxyAlertEnabledOk() (*bool, bool)`

GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAlertEnabled

`func (o *RuleProxyRequestCreate) SetProxyAlertEnabled(v bool)`

SetProxyAlertEnabled sets ProxyAlertEnabled field to given value.

### HasProxyAlertEnabled

`func (o *RuleProxyRequestCreate) HasProxyAlertEnabled() bool`

HasProxyAlertEnabled returns a boolean if a field has been set.

### GetStaticErrorPage

`func (o *RuleProxyRequestCreate) GetStaticErrorPage() string`

GetStaticErrorPage returns the StaticErrorPage field if non-nil, zero value otherwise.

### GetStaticErrorPageOk

`func (o *RuleProxyRequestCreate) GetStaticErrorPageOk() (*string, bool)`

GetStaticErrorPageOk returns a tuple with the StaticErrorPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticErrorPage

`func (o *RuleProxyRequestCreate) SetStaticErrorPage(v string)`

SetStaticErrorPage sets StaticErrorPage field to given value.

### HasStaticErrorPage

`func (o *RuleProxyRequestCreate) HasStaticErrorPage() bool`

HasStaticErrorPage returns a boolean if a field has been set.

### GetStaticErrorPageStatusCodes

`func (o *RuleProxyRequestCreate) GetStaticErrorPageStatusCodes() []string`

GetStaticErrorPageStatusCodes returns the StaticErrorPageStatusCodes field if non-nil, zero value otherwise.

### GetStaticErrorPageStatusCodesOk

`func (o *RuleProxyRequestCreate) GetStaticErrorPageStatusCodesOk() (*[]string, bool)`

GetStaticErrorPageStatusCodesOk returns a tuple with the StaticErrorPageStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticErrorPageStatusCodes

`func (o *RuleProxyRequestCreate) SetStaticErrorPageStatusCodes(v []string)`

SetStaticErrorPageStatusCodes sets StaticErrorPageStatusCodes field to given value.

### HasStaticErrorPageStatusCodes

`func (o *RuleProxyRequestCreate) HasStaticErrorPageStatusCodes() bool`

HasStaticErrorPageStatusCodes returns a boolean if a field has been set.

### GetFailoverMode

`func (o *RuleProxyRequestCreate) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *RuleProxyRequestCreate) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *RuleProxyRequestCreate) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *RuleProxyRequestCreate) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *RuleProxyRequestCreate) GetFailoverOriginTtfb() string`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *RuleProxyRequestCreate) GetFailoverOriginTtfbOk() (*string, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *RuleProxyRequestCreate) SetFailoverOriginTtfb(v string)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *RuleProxyRequestCreate) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverOriginStatusCodes

`func (o *RuleProxyRequestCreate) GetFailoverOriginStatusCodes() []string`

GetFailoverOriginStatusCodes returns the FailoverOriginStatusCodes field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodesOk

`func (o *RuleProxyRequestCreate) GetFailoverOriginStatusCodesOk() (*[]string, bool)`

GetFailoverOriginStatusCodesOk returns a tuple with the FailoverOriginStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCodes

`func (o *RuleProxyRequestCreate) SetFailoverOriginStatusCodes(v []string)`

SetFailoverOriginStatusCodes sets FailoverOriginStatusCodes field to given value.

### HasFailoverOriginStatusCodes

`func (o *RuleProxyRequestCreate) HasFailoverOriginStatusCodes() bool`

HasFailoverOriginStatusCodes returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *RuleProxyRequestCreate) GetFailoverLifetime() string`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *RuleProxyRequestCreate) GetFailoverLifetimeOk() (*string, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *RuleProxyRequestCreate) SetFailoverLifetime(v string)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *RuleProxyRequestCreate) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetNotify

`func (o *RuleProxyRequestCreate) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *RuleProxyRequestCreate) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *RuleProxyRequestCreate) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *RuleProxyRequestCreate) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *RuleProxyRequestCreate) GetNotifyConfig() NotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *RuleProxyRequestCreate) GetNotifyConfigOk() (*NotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *RuleProxyRequestCreate) SetNotifyConfig(v NotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *RuleProxyRequestCreate) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

### GetWafEnabled

`func (o *RuleProxyRequestCreate) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *RuleProxyRequestCreate) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *RuleProxyRequestCreate) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *RuleProxyRequestCreate) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.

### GetWafConfig

`func (o *RuleProxyRequestCreate) GetWafConfig() WAFConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *RuleProxyRequestCreate) GetWafConfigOk() (*WAFConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *RuleProxyRequestCreate) SetWafConfig(v WAFConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *RuleProxyRequestCreate) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.

### GetApplicationProxy

`func (o *RuleProxyRequestCreate) GetApplicationProxy() bool`

GetApplicationProxy returns the ApplicationProxy field if non-nil, zero value otherwise.

### GetApplicationProxyOk

`func (o *RuleProxyRequestCreate) GetApplicationProxyOk() (*bool, bool)`

GetApplicationProxyOk returns a tuple with the ApplicationProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationProxy

`func (o *RuleProxyRequestCreate) SetApplicationProxy(v bool)`

SetApplicationProxy sets ApplicationProxy field to given value.

### HasApplicationProxy

`func (o *RuleProxyRequestCreate) HasApplicationProxy() bool`

HasApplicationProxy returns a boolean if a field has been set.

### GetApplicationName

`func (o *RuleProxyRequestCreate) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *RuleProxyRequestCreate) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *RuleProxyRequestCreate) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *RuleProxyRequestCreate) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationEnvironment

`func (o *RuleProxyRequestCreate) GetApplicationEnvironment() string`

GetApplicationEnvironment returns the ApplicationEnvironment field if non-nil, zero value otherwise.

### GetApplicationEnvironmentOk

`func (o *RuleProxyRequestCreate) GetApplicationEnvironmentOk() (*string, bool)`

GetApplicationEnvironmentOk returns a tuple with the ApplicationEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEnvironment

`func (o *RuleProxyRequestCreate) SetApplicationEnvironment(v string)`

SetApplicationEnvironment sets ApplicationEnvironment field to given value.

### HasApplicationEnvironment

`func (o *RuleProxyRequestCreate) HasApplicationEnvironment() bool`

HasApplicationEnvironment returns a boolean if a field has been set.

### GetApplicationContainer

`func (o *RuleProxyRequestCreate) GetApplicationContainer() string`

GetApplicationContainer returns the ApplicationContainer field if non-nil, zero value otherwise.

### GetApplicationContainerOk

`func (o *RuleProxyRequestCreate) GetApplicationContainerOk() (*string, bool)`

GetApplicationContainerOk returns a tuple with the ApplicationContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationContainer

`func (o *RuleProxyRequestCreate) SetApplicationContainer(v string)`

SetApplicationContainer sets ApplicationContainer field to given value.

### HasApplicationContainer

`func (o *RuleProxyRequestCreate) HasApplicationContainer() bool`

HasApplicationContainer returns a boolean if a field has been set.

### GetApplicationPort

`func (o *RuleProxyRequestCreate) GetApplicationPort() int32`

GetApplicationPort returns the ApplicationPort field if non-nil, zero value otherwise.

### GetApplicationPortOk

`func (o *RuleProxyRequestCreate) GetApplicationPortOk() (*int32, bool)`

GetApplicationPortOk returns a tuple with the ApplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPort

`func (o *RuleProxyRequestCreate) SetApplicationPort(v int32)`

SetApplicationPort sets ApplicationPort field to given value.

### HasApplicationPort

`func (o *RuleProxyRequestCreate) HasApplicationPort() bool`

HasApplicationPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


