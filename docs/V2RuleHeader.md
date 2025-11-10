# V2RuleHeader

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
**ActionConfig** | [**V2RuleHeaderAction**](V2RuleHeaderAction.md) |  | 

## Methods

### NewV2RuleHeader

`func NewV2RuleHeader(uuid string, disabled bool, action string, message string, error_ bool, actionConfig V2RuleHeaderAction, ) *V2RuleHeader`

NewV2RuleHeader instantiates a new V2RuleHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleHeaderWithDefaults

`func NewV2RuleHeaderWithDefaults() *V2RuleHeader`

NewV2RuleHeaderWithDefaults instantiates a new V2RuleHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2RuleHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleHeader) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleHeader) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleHeader) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleHeader) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleHeader) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetRuleId

`func (o *V2RuleHeader) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *V2RuleHeader) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *V2RuleHeader) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *V2RuleHeader) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleHeader) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleHeader) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleHeader) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleHeader) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleHeader) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleHeader) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleHeader) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V2RuleHeader) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *V2RuleHeader) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleHeader) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleHeader) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V2RuleHeader) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleHeader) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleHeader) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleHeader) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetOnlyWithCookie

`func (o *V2RuleHeader) GetOnlyWithCookie() string`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *V2RuleHeader) GetOnlyWithCookieOk() (*string, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *V2RuleHeader) SetOnlyWithCookie(v string)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *V2RuleHeader) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleHeader) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleHeader) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleHeader) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleHeader) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleHeader) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleHeader) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleHeader) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleHeader) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleHeader) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleHeader) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleHeader) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleHeader) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleHeader) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleHeader) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleHeader) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleHeader) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleHeader) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleHeader) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleHeader) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleHeader) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleHeader) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleHeader) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleHeader) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleHeader) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetCountry

`func (o *V2RuleHeader) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleHeader) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleHeader) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleHeader) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleHeader) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleHeader) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleHeader) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleHeader) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleHeader) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleHeader) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleHeader) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleHeader) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetAction

`func (o *V2RuleHeader) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V2RuleHeader) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V2RuleHeader) SetAction(v string)`

SetAction sets Action field to given value.


### GetMessage

`func (o *V2RuleHeader) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleHeader) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleHeader) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleHeader) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleHeader) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleHeader) SetError(v bool)`

SetError sets Error field to given value.


### GetActionConfig

`func (o *V2RuleHeader) GetActionConfig() V2RuleHeaderAction`

GetActionConfig returns the ActionConfig field if non-nil, zero value otherwise.

### GetActionConfigOk

`func (o *V2RuleHeader) GetActionConfigOk() (*V2RuleHeaderAction, bool)`

GetActionConfigOk returns a tuple with the ActionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConfig

`func (o *V2RuleHeader) SetActionConfig(v V2RuleHeaderAction)`

SetActionConfig sets ActionConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


