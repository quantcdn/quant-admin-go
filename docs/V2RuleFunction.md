# V2RuleFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Rule name | [optional] 
**Uuid** | **string** | Rule UUID | 
**RuleId** | Pointer to **string** | Rule ID | [optional] 
**Weight** | Pointer to **int32** | Rule weight | [optional] [default to 0]
**Url** | Pointer to **[]string** | URL patterns | [optional] 
**Domain** | Pointer to **[]string** | Domain patterns | [optional] 
**Disabled** | **bool** | Whether rule is disabled | [default to false]
**OnlyWithCookie** | Pointer to **string** | Only apply with cookie | [optional] 
**Method** | Pointer to **string** | HTTP method | [optional] 
**MethodIs** | Pointer to **[]string** | Allowed HTTP methods | [optional] 
**MethodIsNot** | Pointer to **[]string** | Excluded HTTP methods | [optional] 
**Ip** | Pointer to **string** | IP address | [optional] 
**IpIs** | Pointer to **[]string** | Allowed IP addresses | [optional] 
**IpIsNot** | Pointer to **[]string** | Excluded IP addresses | [optional] 
**Country** | Pointer to **string** | Country code | [optional] 
**CountryIs** | Pointer to **[]string** | Allowed countries | [optional] 
**CountryIsNot** | Pointer to **[]string** | Excluded countries | [optional] 
**Action** | **string** | Rule action | 
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**ActionConfig** | [**V2RuleFunctionAction**](V2RuleFunctionAction.md) |  | 

## Methods

### NewV2RuleFunction

`func NewV2RuleFunction(uuid string, disabled bool, action string, message string, error_ bool, actionConfig V2RuleFunctionAction, ) *V2RuleFunction`

NewV2RuleFunction instantiates a new V2RuleFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleFunctionWithDefaults

`func NewV2RuleFunctionWithDefaults() *V2RuleFunction`

NewV2RuleFunctionWithDefaults instantiates a new V2RuleFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2RuleFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleFunction) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleFunction) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleFunction) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetRuleId

`func (o *V2RuleFunction) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *V2RuleFunction) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *V2RuleFunction) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *V2RuleFunction) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleFunction) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleFunction) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleFunction) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleFunction) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleFunction) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleFunction) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleFunction) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V2RuleFunction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *V2RuleFunction) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleFunction) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleFunction) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V2RuleFunction) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleFunction) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleFunction) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleFunction) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetOnlyWithCookie

`func (o *V2RuleFunction) GetOnlyWithCookie() string`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *V2RuleFunction) GetOnlyWithCookieOk() (*string, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *V2RuleFunction) SetOnlyWithCookie(v string)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *V2RuleFunction) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleFunction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleFunction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleFunction) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleFunction) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleFunction) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleFunction) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleFunction) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleFunction) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleFunction) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleFunction) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleFunction) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleFunction) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleFunction) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleFunction) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleFunction) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleFunction) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleFunction) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleFunction) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleFunction) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleFunction) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleFunction) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleFunction) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleFunction) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleFunction) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetCountry

`func (o *V2RuleFunction) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleFunction) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleFunction) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleFunction) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleFunction) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleFunction) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleFunction) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleFunction) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleFunction) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleFunction) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleFunction) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleFunction) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetAction

`func (o *V2RuleFunction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V2RuleFunction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V2RuleFunction) SetAction(v string)`

SetAction sets Action field to given value.


### GetMessage

`func (o *V2RuleFunction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleFunction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleFunction) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleFunction) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleFunction) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleFunction) SetError(v bool)`

SetError sets Error field to given value.


### GetActionConfig

`func (o *V2RuleFunction) GetActionConfig() V2RuleFunctionAction`

GetActionConfig returns the ActionConfig field if non-nil, zero value otherwise.

### GetActionConfigOk

`func (o *V2RuleFunction) GetActionConfigOk() (*V2RuleFunctionAction, bool)`

GetActionConfigOk returns a tuple with the ActionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConfig

`func (o *V2RuleFunction) SetActionConfig(v V2RuleFunctionAction)`

SetActionConfig sets ActionConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


