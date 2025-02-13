# RuleFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionConfig** | [**RuleFunctionAction**](RuleFunctionAction.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | **string** |  | 
**RuleId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **[]string** |  | [optional] 
**Domain** | Pointer to **[]string** |  | [optional] 
**Disabled** | **bool** |  | [default to false]
**OnlyWithCookie** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**MethodIs** | Pointer to **[]string** |  | [optional] 
**MethodIsNot** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IpIs** | Pointer to **[]string** |  | [optional] 
**IpIsNot** | Pointer to **[]string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryIs** | Pointer to **[]string** |  | [optional] 
**CountryIsNot** | Pointer to **[]string** |  | [optional] 
**Action** | **string** |  | 

## Methods

### NewRuleFunction

`func NewRuleFunction(actionConfig RuleFunctionAction, uuid string, disabled bool, action string, ) *RuleFunction`

NewRuleFunction instantiates a new RuleFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleFunctionWithDefaults

`func NewRuleFunctionWithDefaults() *RuleFunction`

NewRuleFunctionWithDefaults instantiates a new RuleFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionConfig

`func (o *RuleFunction) GetActionConfig() RuleFunctionAction`

GetActionConfig returns the ActionConfig field if non-nil, zero value otherwise.

### GetActionConfigOk

`func (o *RuleFunction) GetActionConfigOk() (*RuleFunctionAction, bool)`

GetActionConfigOk returns a tuple with the ActionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConfig

`func (o *RuleFunction) SetActionConfig(v RuleFunctionAction)`

SetActionConfig sets ActionConfig field to given value.


### GetName

`func (o *RuleFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleFunction) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleFunction) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleFunction) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetRuleId

`func (o *RuleFunction) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleFunction) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleFunction) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *RuleFunction) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetUrl

`func (o *RuleFunction) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleFunction) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleFunction) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RuleFunction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *RuleFunction) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleFunction) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleFunction) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RuleFunction) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleFunction) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleFunction) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleFunction) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetOnlyWithCookie

`func (o *RuleFunction) GetOnlyWithCookie() string`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *RuleFunction) GetOnlyWithCookieOk() (*string, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *RuleFunction) SetOnlyWithCookie(v string)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *RuleFunction) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetMethod

`func (o *RuleFunction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleFunction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleFunction) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleFunction) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleFunction) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleFunction) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleFunction) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleFunction) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleFunction) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleFunction) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleFunction) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleFunction) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleFunction) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleFunction) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleFunction) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleFunction) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleFunction) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleFunction) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleFunction) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleFunction) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleFunction) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleFunction) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleFunction) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleFunction) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetCountry

`func (o *RuleFunction) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleFunction) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleFunction) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleFunction) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleFunction) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleFunction) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleFunction) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleFunction) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleFunction) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleFunction) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleFunction) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleFunction) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetAction

`func (o *RuleFunction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RuleFunction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RuleFunction) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


