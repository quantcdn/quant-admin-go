# RuleCustomResponseRequest

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
**CustomResponseStatusCode** | Pointer to **int32** |  | [optional] 
**CustomResponseBody** | Pointer to **string** |  | [optional] 

## Methods

### NewRuleCustomResponseRequest

`func NewRuleCustomResponseRequest() *RuleCustomResponseRequest`

NewRuleCustomResponseRequest instantiates a new RuleCustomResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleCustomResponseRequestWithDefaults

`func NewRuleCustomResponseRequestWithDefaults() *RuleCustomResponseRequest`

NewRuleCustomResponseRequestWithDefaults instantiates a new RuleCustomResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleCustomResponseRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleCustomResponseRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleCustomResponseRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RuleCustomResponseRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetCountry

`func (o *RuleCustomResponseRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleCustomResponseRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleCustomResponseRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleCustomResponseRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleCustomResponseRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleCustomResponseRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleCustomResponseRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleCustomResponseRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleCustomResponseRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleCustomResponseRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleCustomResponseRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleCustomResponseRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleCustomResponseRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleCustomResponseRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleCustomResponseRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleCustomResponseRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleCustomResponseRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleCustomResponseRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleCustomResponseRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleCustomResponseRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleCustomResponseRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleCustomResponseRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleCustomResponseRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleCustomResponseRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleCustomResponseRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleCustomResponseRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleCustomResponseRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleCustomResponseRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleCustomResponseRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleCustomResponseRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleCustomResponseRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleCustomResponseRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleCustomResponseRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleCustomResponseRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleCustomResponseRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleCustomResponseRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetOnlyWithCookie

`func (o *RuleCustomResponseRequest) GetOnlyWithCookie() string`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *RuleCustomResponseRequest) GetOnlyWithCookieOk() (*string, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *RuleCustomResponseRequest) SetOnlyWithCookie(v string)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *RuleCustomResponseRequest) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetUrl

`func (o *RuleCustomResponseRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleCustomResponseRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleCustomResponseRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RuleCustomResponseRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *RuleCustomResponseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleCustomResponseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleCustomResponseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleCustomResponseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleCustomResponseRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleCustomResponseRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleCustomResponseRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RuleCustomResponseRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTo

`func (o *RuleCustomResponseRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RuleCustomResponseRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RuleCustomResponseRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RuleCustomResponseRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetHost

`func (o *RuleCustomResponseRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RuleCustomResponseRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RuleCustomResponseRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RuleCustomResponseRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAuthUser

`func (o *RuleCustomResponseRequest) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *RuleCustomResponseRequest) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *RuleCustomResponseRequest) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *RuleCustomResponseRequest) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetAuthPass

`func (o *RuleCustomResponseRequest) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *RuleCustomResponseRequest) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *RuleCustomResponseRequest) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *RuleCustomResponseRequest) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *RuleCustomResponseRequest) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *RuleCustomResponseRequest) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *RuleCustomResponseRequest) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *RuleCustomResponseRequest) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *RuleCustomResponseRequest) GetCacheLifetime() int32`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *RuleCustomResponseRequest) GetCacheLifetimeOk() (*int32, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *RuleCustomResponseRequest) SetCacheLifetime(v int32)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *RuleCustomResponseRequest) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### GetOnlyProxy404

`func (o *RuleCustomResponseRequest) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *RuleCustomResponseRequest) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *RuleCustomResponseRequest) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *RuleCustomResponseRequest) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetStripHeaders

`func (o *RuleCustomResponseRequest) GetStripHeaders() []string`

GetStripHeaders returns the StripHeaders field if non-nil, zero value otherwise.

### GetStripHeadersOk

`func (o *RuleCustomResponseRequest) GetStripHeadersOk() (*[]string, bool)`

GetStripHeadersOk returns a tuple with the StripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripHeaders

`func (o *RuleCustomResponseRequest) SetStripHeaders(v []string)`

SetStripHeaders sets StripHeaders field to given value.

### HasStripHeaders

`func (o *RuleCustomResponseRequest) HasStripHeaders() bool`

HasStripHeaders returns a boolean if a field has been set.

### GetCustomResponseStatusCode

`func (o *RuleCustomResponseRequest) GetCustomResponseStatusCode() int32`

GetCustomResponseStatusCode returns the CustomResponseStatusCode field if non-nil, zero value otherwise.

### GetCustomResponseStatusCodeOk

`func (o *RuleCustomResponseRequest) GetCustomResponseStatusCodeOk() (*int32, bool)`

GetCustomResponseStatusCodeOk returns a tuple with the CustomResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseStatusCode

`func (o *RuleCustomResponseRequest) SetCustomResponseStatusCode(v int32)`

SetCustomResponseStatusCode sets CustomResponseStatusCode field to given value.

### HasCustomResponseStatusCode

`func (o *RuleCustomResponseRequest) HasCustomResponseStatusCode() bool`

HasCustomResponseStatusCode returns a boolean if a field has been set.

### GetCustomResponseBody

`func (o *RuleCustomResponseRequest) GetCustomResponseBody() string`

GetCustomResponseBody returns the CustomResponseBody field if non-nil, zero value otherwise.

### GetCustomResponseBodyOk

`func (o *RuleCustomResponseRequest) GetCustomResponseBodyOk() (*string, bool)`

GetCustomResponseBodyOk returns a tuple with the CustomResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseBody

`func (o *RuleCustomResponseRequest) SetCustomResponseBody(v string)`

SetCustomResponseBody sets CustomResponseBody field to given value.

### HasCustomResponseBody

`func (o *RuleCustomResponseRequest) HasCustomResponseBody() bool`

HasCustomResponseBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


