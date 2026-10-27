# V2RuleAuthFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **[]string** | Exact hostnames the rule matches, or &#39;any&#39; (the default). Wildcards such as *.example.com are not supported. | 
**Name** | Pointer to **string** | Rule name | [optional] 
**Uuid** | Pointer to **string** | Rule UUID | [optional] 
**Weight** | Pointer to **int32** | Rule order. Must be 0 or greater; lower weights are processed first. If omitted on create, the rule is placed last. | [optional] 
**Disabled** | Pointer to **bool** | Whether rule is disabled | [optional] [default to false]
**Url** | **[]string** | URL patterns | 
**Country** | Pointer to **string** | Country filter type (country_is, country_is_not, any) | [optional] 
**CountryIs** | Pointer to **[]string** | Allowed countries | [optional] 
**CountryIsNot** | Pointer to **[]string** | Excluded countries | [optional] 
**Method** | Pointer to **string** | Method filter type (method_is, method_is_not, any) | [optional] 
**MethodIs** | Pointer to **[]string** | Allowed HTTP methods | [optional] 
**MethodIsNot** | Pointer to **[]string** | Excluded HTTP methods | [optional] 
**Ip** | Pointer to **string** | IP filter type (ip_is, ip_is_not, any) | [optional] 
**IpIs** | Pointer to **[]string** | Allowed IP addresses | [optional] 
**IpIsNot** | Pointer to **[]string** | Excluded IP addresses | [optional] 
**Asn** | Pointer to **string** | ASN filter type (asn_is, asn_is_not, any). A PATCH that omits it keeps the current value. Invalid values return 422. | [optional] 
**AsnIs** | Pointer to **[]string** | Allowed AS numbers. Positive integers up to 4294967295, stored as strings. A PATCH replaces the list. | [optional] 
**AsnIsNot** | Pointer to **[]string** | Excluded AS numbers. Positive integers up to 4294967295, stored as strings. A PATCH replaces the list. | [optional] 
**Ua** | Pointer to **string** | User agent filter type (ua_is, ua_is_not, any). Invalid values return 422. | [optional] 
**UaIs** | Pointer to **[]string** | User agent strings to match. Each item has at most 255 characters. A PATCH replaces the list. | [optional] 
**UaIsNot** | Pointer to **[]string** | User agent strings to exclude. Each item has at most 255 characters. A PATCH replaces the list. | [optional] 
**TrafficSource** | Pointer to **string** | Traffic source filter type (traffic_source_is, traffic_source_is_not, any). Invalid values return 422. | [optional] 
**TrafficSourceIs** | Pointer to **[]string** | Traffic sources to match. A PATCH replaces the list. | [optional] 
**TrafficSourceIsNot** | Pointer to **[]string** | Traffic sources to exclude. A PATCH replaces the list. | [optional] 
**CookieName** | Pointer to **string** | Write name of only_with_cookie: the rule applies only when this cookie is present. Empty string clears it. | [optional] 
**FnUuid** | **string** | UUID of a published function of type edge_auth in this project. A create, or a change of fn_uuid, that names any other function returns 422. | 

## Methods

### NewV2RuleAuthFunctionRequest

`func NewV2RuleAuthFunctionRequest(domain []string, url []string, fnUuid string, ) *V2RuleAuthFunctionRequest`

NewV2RuleAuthFunctionRequest instantiates a new V2RuleAuthFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleAuthFunctionRequestWithDefaults

`func NewV2RuleAuthFunctionRequestWithDefaults() *V2RuleAuthFunctionRequest`

NewV2RuleAuthFunctionRequestWithDefaults instantiates a new V2RuleAuthFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *V2RuleAuthFunctionRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleAuthFunctionRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleAuthFunctionRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *V2RuleAuthFunctionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleAuthFunctionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleAuthFunctionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleAuthFunctionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleAuthFunctionRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleAuthFunctionRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleAuthFunctionRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2RuleAuthFunctionRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleAuthFunctionRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleAuthFunctionRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleAuthFunctionRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleAuthFunctionRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleAuthFunctionRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleAuthFunctionRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleAuthFunctionRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V2RuleAuthFunctionRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleAuthFunctionRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleAuthFunctionRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleAuthFunctionRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *V2RuleAuthFunctionRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleAuthFunctionRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleAuthFunctionRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleAuthFunctionRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleAuthFunctionRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleAuthFunctionRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleAuthFunctionRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleAuthFunctionRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleAuthFunctionRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleAuthFunctionRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleAuthFunctionRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleAuthFunctionRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleAuthFunctionRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleAuthFunctionRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleAuthFunctionRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleAuthFunctionRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleAuthFunctionRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleAuthFunctionRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleAuthFunctionRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleAuthFunctionRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleAuthFunctionRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleAuthFunctionRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleAuthFunctionRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleAuthFunctionRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleAuthFunctionRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleAuthFunctionRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleAuthFunctionRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleAuthFunctionRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleAuthFunctionRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleAuthFunctionRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleAuthFunctionRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleAuthFunctionRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleAuthFunctionRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleAuthFunctionRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleAuthFunctionRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleAuthFunctionRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetAsn

`func (o *V2RuleAuthFunctionRequest) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *V2RuleAuthFunctionRequest) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *V2RuleAuthFunctionRequest) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *V2RuleAuthFunctionRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetAsnIs

`func (o *V2RuleAuthFunctionRequest) GetAsnIs() []string`

GetAsnIs returns the AsnIs field if non-nil, zero value otherwise.

### GetAsnIsOk

`func (o *V2RuleAuthFunctionRequest) GetAsnIsOk() (*[]string, bool)`

GetAsnIsOk returns a tuple with the AsnIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnIs

`func (o *V2RuleAuthFunctionRequest) SetAsnIs(v []string)`

SetAsnIs sets AsnIs field to given value.

### HasAsnIs

`func (o *V2RuleAuthFunctionRequest) HasAsnIs() bool`

HasAsnIs returns a boolean if a field has been set.

### GetAsnIsNot

`func (o *V2RuleAuthFunctionRequest) GetAsnIsNot() []string`

GetAsnIsNot returns the AsnIsNot field if non-nil, zero value otherwise.

### GetAsnIsNotOk

`func (o *V2RuleAuthFunctionRequest) GetAsnIsNotOk() (*[]string, bool)`

GetAsnIsNotOk returns a tuple with the AsnIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnIsNot

`func (o *V2RuleAuthFunctionRequest) SetAsnIsNot(v []string)`

SetAsnIsNot sets AsnIsNot field to given value.

### HasAsnIsNot

`func (o *V2RuleAuthFunctionRequest) HasAsnIsNot() bool`

HasAsnIsNot returns a boolean if a field has been set.

### GetUa

`func (o *V2RuleAuthFunctionRequest) GetUa() string`

GetUa returns the Ua field if non-nil, zero value otherwise.

### GetUaOk

`func (o *V2RuleAuthFunctionRequest) GetUaOk() (*string, bool)`

GetUaOk returns a tuple with the Ua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUa

`func (o *V2RuleAuthFunctionRequest) SetUa(v string)`

SetUa sets Ua field to given value.

### HasUa

`func (o *V2RuleAuthFunctionRequest) HasUa() bool`

HasUa returns a boolean if a field has been set.

### GetUaIs

`func (o *V2RuleAuthFunctionRequest) GetUaIs() []string`

GetUaIs returns the UaIs field if non-nil, zero value otherwise.

### GetUaIsOk

`func (o *V2RuleAuthFunctionRequest) GetUaIsOk() (*[]string, bool)`

GetUaIsOk returns a tuple with the UaIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaIs

`func (o *V2RuleAuthFunctionRequest) SetUaIs(v []string)`

SetUaIs sets UaIs field to given value.

### HasUaIs

`func (o *V2RuleAuthFunctionRequest) HasUaIs() bool`

HasUaIs returns a boolean if a field has been set.

### GetUaIsNot

`func (o *V2RuleAuthFunctionRequest) GetUaIsNot() []string`

GetUaIsNot returns the UaIsNot field if non-nil, zero value otherwise.

### GetUaIsNotOk

`func (o *V2RuleAuthFunctionRequest) GetUaIsNotOk() (*[]string, bool)`

GetUaIsNotOk returns a tuple with the UaIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaIsNot

`func (o *V2RuleAuthFunctionRequest) SetUaIsNot(v []string)`

SetUaIsNot sets UaIsNot field to given value.

### HasUaIsNot

`func (o *V2RuleAuthFunctionRequest) HasUaIsNot() bool`

HasUaIsNot returns a boolean if a field has been set.

### GetTrafficSource

`func (o *V2RuleAuthFunctionRequest) GetTrafficSource() string`

GetTrafficSource returns the TrafficSource field if non-nil, zero value otherwise.

### GetTrafficSourceOk

`func (o *V2RuleAuthFunctionRequest) GetTrafficSourceOk() (*string, bool)`

GetTrafficSourceOk returns a tuple with the TrafficSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSource

`func (o *V2RuleAuthFunctionRequest) SetTrafficSource(v string)`

SetTrafficSource sets TrafficSource field to given value.

### HasTrafficSource

`func (o *V2RuleAuthFunctionRequest) HasTrafficSource() bool`

HasTrafficSource returns a boolean if a field has been set.

### GetTrafficSourceIs

`func (o *V2RuleAuthFunctionRequest) GetTrafficSourceIs() []string`

GetTrafficSourceIs returns the TrafficSourceIs field if non-nil, zero value otherwise.

### GetTrafficSourceIsOk

`func (o *V2RuleAuthFunctionRequest) GetTrafficSourceIsOk() (*[]string, bool)`

GetTrafficSourceIsOk returns a tuple with the TrafficSourceIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSourceIs

`func (o *V2RuleAuthFunctionRequest) SetTrafficSourceIs(v []string)`

SetTrafficSourceIs sets TrafficSourceIs field to given value.

### HasTrafficSourceIs

`func (o *V2RuleAuthFunctionRequest) HasTrafficSourceIs() bool`

HasTrafficSourceIs returns a boolean if a field has been set.

### GetTrafficSourceIsNot

`func (o *V2RuleAuthFunctionRequest) GetTrafficSourceIsNot() []string`

GetTrafficSourceIsNot returns the TrafficSourceIsNot field if non-nil, zero value otherwise.

### GetTrafficSourceIsNotOk

`func (o *V2RuleAuthFunctionRequest) GetTrafficSourceIsNotOk() (*[]string, bool)`

GetTrafficSourceIsNotOk returns a tuple with the TrafficSourceIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSourceIsNot

`func (o *V2RuleAuthFunctionRequest) SetTrafficSourceIsNot(v []string)`

SetTrafficSourceIsNot sets TrafficSourceIsNot field to given value.

### HasTrafficSourceIsNot

`func (o *V2RuleAuthFunctionRequest) HasTrafficSourceIsNot() bool`

HasTrafficSourceIsNot returns a boolean if a field has been set.

### GetCookieName

`func (o *V2RuleAuthFunctionRequest) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *V2RuleAuthFunctionRequest) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *V2RuleAuthFunctionRequest) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *V2RuleAuthFunctionRequest) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetFnUuid

`func (o *V2RuleAuthFunctionRequest) GetFnUuid() string`

GetFnUuid returns the FnUuid field if non-nil, zero value otherwise.

### GetFnUuidOk

`func (o *V2RuleAuthFunctionRequest) GetFnUuidOk() (*string, bool)`

GetFnUuidOk returns a tuple with the FnUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFnUuid

`func (o *V2RuleAuthFunctionRequest) SetFnUuid(v string)`

SetFnUuid sets FnUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


