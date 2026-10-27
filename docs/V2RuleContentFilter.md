# V2RuleContentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Rule name | [optional] 
**Uuid** | **string** | Rule UUID | 
**RuleId** | Pointer to **string** | Rule ID | [optional] 
**Weight** | Pointer to **int32** | Rule order. Must be 0 or greater; lower weights are processed first. If omitted on create, the rule is placed last. | [optional] 
**Url** | Pointer to **[]string** | URL patterns | [optional] 
**Domain** | Pointer to **[]string** | Exact hostnames the rule matches, or &#39;any&#39;. Wildcards are not supported. | [optional] 
**Disabled** | **bool** | Whether rule is disabled | [default to false]
**OnlyWithCookie** | Pointer to **string** | Only apply with cookie | [optional] 
**Method** | Pointer to **string** | HTTP method | [optional] 
**MethodIs** | Pointer to **[]string** | Allowed HTTP methods | [optional] 
**MethodIsNot** | Pointer to **[]string** | Excluded HTTP methods | [optional] 
**Ip** | Pointer to **string** | IP address | [optional] 
**IpIs** | Pointer to **[]string** | Allowed IP addresses | [optional] 
**IpIsNot** | Pointer to **[]string** | Excluded IP addresses | [optional] 
**Asn** | Pointer to **string** | ASN condition type (asn_is, asn_is_not, any) | [optional] 
**AsnIs** | Pointer to **[]string** | Allowed AS numbers | [optional] 
**AsnIsNot** | Pointer to **[]string** | Excluded AS numbers | [optional] 
**Ua** | Pointer to **string** | User agent condition type (ua_is, ua_is_not, any) | [optional] 
**UaIs** | Pointer to **[]string** | User agent strings to match | [optional] 
**UaIsNot** | Pointer to **[]string** | User agent strings to exclude | [optional] 
**TrafficSource** | Pointer to **string** | Traffic source condition type (traffic_source_is, traffic_source_is_not, any) | [optional] 
**TrafficSourceIs** | Pointer to **[]string** | Traffic sources to match (cloud, vpn, tor-exit, abuser, crawler) | [optional] 
**TrafficSourceIsNot** | Pointer to **[]string** | Traffic sources to exclude (cloud, vpn, tor-exit, abuser, crawler) | [optional] 
**Country** | Pointer to **string** | Country code | [optional] 
**CountryIs** | Pointer to **[]string** | Allowed countries | [optional] 
**CountryIsNot** | Pointer to **[]string** | Excluded countries | [optional] 
**Action** | **string** | Rule action | 
**ActionConfig** | [**V2RuleContentFilterAction**](V2RuleContentFilterAction.md) |  | 

## Methods

### NewV2RuleContentFilter

`func NewV2RuleContentFilter(uuid string, disabled bool, action string, actionConfig V2RuleContentFilterAction, ) *V2RuleContentFilter`

NewV2RuleContentFilter instantiates a new V2RuleContentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleContentFilterWithDefaults

`func NewV2RuleContentFilterWithDefaults() *V2RuleContentFilter`

NewV2RuleContentFilterWithDefaults instantiates a new V2RuleContentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2RuleContentFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleContentFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleContentFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleContentFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleContentFilter) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleContentFilter) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleContentFilter) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetRuleId

`func (o *V2RuleContentFilter) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *V2RuleContentFilter) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *V2RuleContentFilter) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *V2RuleContentFilter) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleContentFilter) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleContentFilter) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleContentFilter) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleContentFilter) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleContentFilter) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleContentFilter) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleContentFilter) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V2RuleContentFilter) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *V2RuleContentFilter) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleContentFilter) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleContentFilter) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V2RuleContentFilter) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleContentFilter) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleContentFilter) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleContentFilter) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetOnlyWithCookie

`func (o *V2RuleContentFilter) GetOnlyWithCookie() string`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *V2RuleContentFilter) GetOnlyWithCookieOk() (*string, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *V2RuleContentFilter) SetOnlyWithCookie(v string)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *V2RuleContentFilter) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleContentFilter) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleContentFilter) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleContentFilter) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleContentFilter) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleContentFilter) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleContentFilter) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleContentFilter) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleContentFilter) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleContentFilter) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleContentFilter) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleContentFilter) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleContentFilter) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleContentFilter) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleContentFilter) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleContentFilter) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleContentFilter) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleContentFilter) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleContentFilter) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleContentFilter) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleContentFilter) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleContentFilter) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleContentFilter) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleContentFilter) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleContentFilter) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetAsn

`func (o *V2RuleContentFilter) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *V2RuleContentFilter) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *V2RuleContentFilter) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *V2RuleContentFilter) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetAsnIs

`func (o *V2RuleContentFilter) GetAsnIs() []string`

GetAsnIs returns the AsnIs field if non-nil, zero value otherwise.

### GetAsnIsOk

`func (o *V2RuleContentFilter) GetAsnIsOk() (*[]string, bool)`

GetAsnIsOk returns a tuple with the AsnIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnIs

`func (o *V2RuleContentFilter) SetAsnIs(v []string)`

SetAsnIs sets AsnIs field to given value.

### HasAsnIs

`func (o *V2RuleContentFilter) HasAsnIs() bool`

HasAsnIs returns a boolean if a field has been set.

### GetAsnIsNot

`func (o *V2RuleContentFilter) GetAsnIsNot() []string`

GetAsnIsNot returns the AsnIsNot field if non-nil, zero value otherwise.

### GetAsnIsNotOk

`func (o *V2RuleContentFilter) GetAsnIsNotOk() (*[]string, bool)`

GetAsnIsNotOk returns a tuple with the AsnIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnIsNot

`func (o *V2RuleContentFilter) SetAsnIsNot(v []string)`

SetAsnIsNot sets AsnIsNot field to given value.

### HasAsnIsNot

`func (o *V2RuleContentFilter) HasAsnIsNot() bool`

HasAsnIsNot returns a boolean if a field has been set.

### GetUa

`func (o *V2RuleContentFilter) GetUa() string`

GetUa returns the Ua field if non-nil, zero value otherwise.

### GetUaOk

`func (o *V2RuleContentFilter) GetUaOk() (*string, bool)`

GetUaOk returns a tuple with the Ua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUa

`func (o *V2RuleContentFilter) SetUa(v string)`

SetUa sets Ua field to given value.

### HasUa

`func (o *V2RuleContentFilter) HasUa() bool`

HasUa returns a boolean if a field has been set.

### GetUaIs

`func (o *V2RuleContentFilter) GetUaIs() []string`

GetUaIs returns the UaIs field if non-nil, zero value otherwise.

### GetUaIsOk

`func (o *V2RuleContentFilter) GetUaIsOk() (*[]string, bool)`

GetUaIsOk returns a tuple with the UaIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaIs

`func (o *V2RuleContentFilter) SetUaIs(v []string)`

SetUaIs sets UaIs field to given value.

### HasUaIs

`func (o *V2RuleContentFilter) HasUaIs() bool`

HasUaIs returns a boolean if a field has been set.

### GetUaIsNot

`func (o *V2RuleContentFilter) GetUaIsNot() []string`

GetUaIsNot returns the UaIsNot field if non-nil, zero value otherwise.

### GetUaIsNotOk

`func (o *V2RuleContentFilter) GetUaIsNotOk() (*[]string, bool)`

GetUaIsNotOk returns a tuple with the UaIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaIsNot

`func (o *V2RuleContentFilter) SetUaIsNot(v []string)`

SetUaIsNot sets UaIsNot field to given value.

### HasUaIsNot

`func (o *V2RuleContentFilter) HasUaIsNot() bool`

HasUaIsNot returns a boolean if a field has been set.

### GetTrafficSource

`func (o *V2RuleContentFilter) GetTrafficSource() string`

GetTrafficSource returns the TrafficSource field if non-nil, zero value otherwise.

### GetTrafficSourceOk

`func (o *V2RuleContentFilter) GetTrafficSourceOk() (*string, bool)`

GetTrafficSourceOk returns a tuple with the TrafficSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSource

`func (o *V2RuleContentFilter) SetTrafficSource(v string)`

SetTrafficSource sets TrafficSource field to given value.

### HasTrafficSource

`func (o *V2RuleContentFilter) HasTrafficSource() bool`

HasTrafficSource returns a boolean if a field has been set.

### GetTrafficSourceIs

`func (o *V2RuleContentFilter) GetTrafficSourceIs() []string`

GetTrafficSourceIs returns the TrafficSourceIs field if non-nil, zero value otherwise.

### GetTrafficSourceIsOk

`func (o *V2RuleContentFilter) GetTrafficSourceIsOk() (*[]string, bool)`

GetTrafficSourceIsOk returns a tuple with the TrafficSourceIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSourceIs

`func (o *V2RuleContentFilter) SetTrafficSourceIs(v []string)`

SetTrafficSourceIs sets TrafficSourceIs field to given value.

### HasTrafficSourceIs

`func (o *V2RuleContentFilter) HasTrafficSourceIs() bool`

HasTrafficSourceIs returns a boolean if a field has been set.

### GetTrafficSourceIsNot

`func (o *V2RuleContentFilter) GetTrafficSourceIsNot() []string`

GetTrafficSourceIsNot returns the TrafficSourceIsNot field if non-nil, zero value otherwise.

### GetTrafficSourceIsNotOk

`func (o *V2RuleContentFilter) GetTrafficSourceIsNotOk() (*[]string, bool)`

GetTrafficSourceIsNotOk returns a tuple with the TrafficSourceIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSourceIsNot

`func (o *V2RuleContentFilter) SetTrafficSourceIsNot(v []string)`

SetTrafficSourceIsNot sets TrafficSourceIsNot field to given value.

### HasTrafficSourceIsNot

`func (o *V2RuleContentFilter) HasTrafficSourceIsNot() bool`

HasTrafficSourceIsNot returns a boolean if a field has been set.

### GetCountry

`func (o *V2RuleContentFilter) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleContentFilter) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleContentFilter) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleContentFilter) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleContentFilter) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleContentFilter) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleContentFilter) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleContentFilter) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleContentFilter) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleContentFilter) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleContentFilter) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleContentFilter) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetAction

`func (o *V2RuleContentFilter) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V2RuleContentFilter) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V2RuleContentFilter) SetAction(v string)`

SetAction sets Action field to given value.


### GetActionConfig

`func (o *V2RuleContentFilter) GetActionConfig() V2RuleContentFilterAction`

GetActionConfig returns the ActionConfig field if non-nil, zero value otherwise.

### GetActionConfigOk

`func (o *V2RuleContentFilter) GetActionConfigOk() (*V2RuleContentFilterAction, bool)`

GetActionConfigOk returns a tuple with the ActionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConfig

`func (o *V2RuleContentFilter) SetActionConfig(v V2RuleContentFilterAction)`

SetActionConfig sets ActionConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


