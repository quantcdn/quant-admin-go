# RuleAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | [default to "any"]
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
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
**AuthUser** | **string** |  | 
**AuthPass** | **string** |  | 

## Methods

### NewRuleAuthRequest

`func NewRuleAuthRequest(domain string, disabled bool, url []string, authUser string, authPass string, ) *RuleAuthRequest`

NewRuleAuthRequest instantiates a new RuleAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleAuthRequestWithDefaults

`func NewRuleAuthRequestWithDefaults() *RuleAuthRequest`

NewRuleAuthRequestWithDefaults instantiates a new RuleAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleAuthRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleAuthRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleAuthRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *RuleAuthRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleAuthRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleAuthRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleAuthRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleAuthRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleAuthRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleAuthRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleAuthRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleAuthRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleAuthRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleAuthRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetUrl

`func (o *RuleAuthRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleAuthRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleAuthRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *RuleAuthRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleAuthRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleAuthRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleAuthRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleAuthRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleAuthRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleAuthRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleAuthRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleAuthRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleAuthRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleAuthRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleAuthRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleAuthRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleAuthRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleAuthRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleAuthRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleAuthRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleAuthRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleAuthRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleAuthRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleAuthRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleAuthRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleAuthRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleAuthRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleAuthRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleAuthRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleAuthRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleAuthRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleAuthRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleAuthRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleAuthRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleAuthRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleAuthRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleAuthRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleAuthRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleAuthRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetOnlyWithCookie

`func (o *RuleAuthRequest) GetOnlyWithCookie() bool`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *RuleAuthRequest) GetOnlyWithCookieOk() (*bool, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *RuleAuthRequest) SetOnlyWithCookie(v bool)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *RuleAuthRequest) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetCookieName

`func (o *RuleAuthRequest) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *RuleAuthRequest) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *RuleAuthRequest) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *RuleAuthRequest) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetAuthUser

`func (o *RuleAuthRequest) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *RuleAuthRequest) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *RuleAuthRequest) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.


### GetAuthPass

`func (o *RuleAuthRequest) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *RuleAuthRequest) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *RuleAuthRequest) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


