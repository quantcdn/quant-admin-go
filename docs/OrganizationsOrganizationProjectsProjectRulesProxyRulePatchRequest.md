# OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryIs** | Pointer to **[]string** |  | [optional] 
**CountryIsNot** | Pointer to **[]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**MethodIs** | Pointer to **[]string** |  | [optional] 
**MethodIsNot** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IpIs** | Pointer to **[]string** |  | [optional] 
**IpIsNot** | Pointer to **[]string** |  | [optional] 
**OnlyWithCookie** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**AuthUser** | Pointer to **string** |  | [optional] 
**AuthPass** | Pointer to **string** |  | [optional] 
**DisableSslVerify** | Pointer to **bool** |  | [optional] 
**CacheLifetime** | Pointer to **int32** |  | [optional] 
**OnlyProxy404** | Pointer to **bool** |  | [optional] 
**StripHeaders** | Pointer to **[]string** |  | [optional] 
**WafEnabled** | Pointer to **bool** |  | [optional] 
**WafConfig** | Pointer to [**OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig**](OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest

`func NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest() *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest`

NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest instantiates a new OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWithDefaults

`func NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWithDefaults() *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest`

NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWithDefaults instantiates a new OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetCountry

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetOnlyWithCookie

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetOnlyWithCookie() string`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetOnlyWithCookieOk() (*string, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetOnlyWithCookie(v string)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetUrl

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisabled

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTo

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetHost

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAuthUser

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetAuthPass

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetCacheLifetime() int32`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetCacheLifetimeOk() (*int32, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetCacheLifetime(v int32)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### GetOnlyProxy404

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetStripHeaders

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetStripHeaders() []string`

GetStripHeaders returns the StripHeaders field if non-nil, zero value otherwise.

### GetStripHeadersOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetStripHeadersOk() (*[]string, bool)`

GetStripHeadersOk returns a tuple with the StripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripHeaders

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetStripHeaders(v []string)`

SetStripHeaders sets StripHeaders field to given value.

### HasStripHeaders

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasStripHeaders() bool`

HasStripHeaders returns a boolean if a field has been set.

### GetWafEnabled

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.

### GetWafConfig

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetWafConfig() OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) GetWafConfigOk() (*OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) SetWafConfig(v OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


