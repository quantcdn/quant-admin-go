# V2RuleFunctionRequest

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
**FnUuid** | **string** | UUID of a published function of type edge_function in this project. A create, or a change of fn_uuid, that names any other function returns 422. | 

## Methods

### NewV2RuleFunctionRequest

`func NewV2RuleFunctionRequest(domain []string, url []string, fnUuid string, ) *V2RuleFunctionRequest`

NewV2RuleFunctionRequest instantiates a new V2RuleFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleFunctionRequestWithDefaults

`func NewV2RuleFunctionRequestWithDefaults() *V2RuleFunctionRequest`

NewV2RuleFunctionRequestWithDefaults instantiates a new V2RuleFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *V2RuleFunctionRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleFunctionRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleFunctionRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *V2RuleFunctionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleFunctionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleFunctionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleFunctionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleFunctionRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleFunctionRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleFunctionRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2RuleFunctionRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleFunctionRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleFunctionRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleFunctionRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleFunctionRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleFunctionRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleFunctionRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleFunctionRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V2RuleFunctionRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleFunctionRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleFunctionRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleFunctionRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *V2RuleFunctionRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleFunctionRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleFunctionRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleFunctionRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleFunctionRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleFunctionRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleFunctionRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleFunctionRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleFunctionRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleFunctionRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleFunctionRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleFunctionRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleFunctionRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleFunctionRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleFunctionRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleFunctionRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleFunctionRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleFunctionRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleFunctionRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleFunctionRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleFunctionRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleFunctionRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleFunctionRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleFunctionRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleFunctionRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleFunctionRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleFunctionRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleFunctionRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleFunctionRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleFunctionRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleFunctionRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleFunctionRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleFunctionRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleFunctionRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleFunctionRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleFunctionRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetAsn

`func (o *V2RuleFunctionRequest) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *V2RuleFunctionRequest) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *V2RuleFunctionRequest) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *V2RuleFunctionRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetAsnIs

`func (o *V2RuleFunctionRequest) GetAsnIs() []string`

GetAsnIs returns the AsnIs field if non-nil, zero value otherwise.

### GetAsnIsOk

`func (o *V2RuleFunctionRequest) GetAsnIsOk() (*[]string, bool)`

GetAsnIsOk returns a tuple with the AsnIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnIs

`func (o *V2RuleFunctionRequest) SetAsnIs(v []string)`

SetAsnIs sets AsnIs field to given value.

### HasAsnIs

`func (o *V2RuleFunctionRequest) HasAsnIs() bool`

HasAsnIs returns a boolean if a field has been set.

### GetAsnIsNot

`func (o *V2RuleFunctionRequest) GetAsnIsNot() []string`

GetAsnIsNot returns the AsnIsNot field if non-nil, zero value otherwise.

### GetAsnIsNotOk

`func (o *V2RuleFunctionRequest) GetAsnIsNotOk() (*[]string, bool)`

GetAsnIsNotOk returns a tuple with the AsnIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnIsNot

`func (o *V2RuleFunctionRequest) SetAsnIsNot(v []string)`

SetAsnIsNot sets AsnIsNot field to given value.

### HasAsnIsNot

`func (o *V2RuleFunctionRequest) HasAsnIsNot() bool`

HasAsnIsNot returns a boolean if a field has been set.

### GetUa

`func (o *V2RuleFunctionRequest) GetUa() string`

GetUa returns the Ua field if non-nil, zero value otherwise.

### GetUaOk

`func (o *V2RuleFunctionRequest) GetUaOk() (*string, bool)`

GetUaOk returns a tuple with the Ua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUa

`func (o *V2RuleFunctionRequest) SetUa(v string)`

SetUa sets Ua field to given value.

### HasUa

`func (o *V2RuleFunctionRequest) HasUa() bool`

HasUa returns a boolean if a field has been set.

### GetUaIs

`func (o *V2RuleFunctionRequest) GetUaIs() []string`

GetUaIs returns the UaIs field if non-nil, zero value otherwise.

### GetUaIsOk

`func (o *V2RuleFunctionRequest) GetUaIsOk() (*[]string, bool)`

GetUaIsOk returns a tuple with the UaIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaIs

`func (o *V2RuleFunctionRequest) SetUaIs(v []string)`

SetUaIs sets UaIs field to given value.

### HasUaIs

`func (o *V2RuleFunctionRequest) HasUaIs() bool`

HasUaIs returns a boolean if a field has been set.

### GetUaIsNot

`func (o *V2RuleFunctionRequest) GetUaIsNot() []string`

GetUaIsNot returns the UaIsNot field if non-nil, zero value otherwise.

### GetUaIsNotOk

`func (o *V2RuleFunctionRequest) GetUaIsNotOk() (*[]string, bool)`

GetUaIsNotOk returns a tuple with the UaIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaIsNot

`func (o *V2RuleFunctionRequest) SetUaIsNot(v []string)`

SetUaIsNot sets UaIsNot field to given value.

### HasUaIsNot

`func (o *V2RuleFunctionRequest) HasUaIsNot() bool`

HasUaIsNot returns a boolean if a field has been set.

### GetTrafficSource

`func (o *V2RuleFunctionRequest) GetTrafficSource() string`

GetTrafficSource returns the TrafficSource field if non-nil, zero value otherwise.

### GetTrafficSourceOk

`func (o *V2RuleFunctionRequest) GetTrafficSourceOk() (*string, bool)`

GetTrafficSourceOk returns a tuple with the TrafficSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSource

`func (o *V2RuleFunctionRequest) SetTrafficSource(v string)`

SetTrafficSource sets TrafficSource field to given value.

### HasTrafficSource

`func (o *V2RuleFunctionRequest) HasTrafficSource() bool`

HasTrafficSource returns a boolean if a field has been set.

### GetTrafficSourceIs

`func (o *V2RuleFunctionRequest) GetTrafficSourceIs() []string`

GetTrafficSourceIs returns the TrafficSourceIs field if non-nil, zero value otherwise.

### GetTrafficSourceIsOk

`func (o *V2RuleFunctionRequest) GetTrafficSourceIsOk() (*[]string, bool)`

GetTrafficSourceIsOk returns a tuple with the TrafficSourceIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSourceIs

`func (o *V2RuleFunctionRequest) SetTrafficSourceIs(v []string)`

SetTrafficSourceIs sets TrafficSourceIs field to given value.

### HasTrafficSourceIs

`func (o *V2RuleFunctionRequest) HasTrafficSourceIs() bool`

HasTrafficSourceIs returns a boolean if a field has been set.

### GetTrafficSourceIsNot

`func (o *V2RuleFunctionRequest) GetTrafficSourceIsNot() []string`

GetTrafficSourceIsNot returns the TrafficSourceIsNot field if non-nil, zero value otherwise.

### GetTrafficSourceIsNotOk

`func (o *V2RuleFunctionRequest) GetTrafficSourceIsNotOk() (*[]string, bool)`

GetTrafficSourceIsNotOk returns a tuple with the TrafficSourceIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSourceIsNot

`func (o *V2RuleFunctionRequest) SetTrafficSourceIsNot(v []string)`

SetTrafficSourceIsNot sets TrafficSourceIsNot field to given value.

### HasTrafficSourceIsNot

`func (o *V2RuleFunctionRequest) HasTrafficSourceIsNot() bool`

HasTrafficSourceIsNot returns a boolean if a field has been set.

### GetCookieName

`func (o *V2RuleFunctionRequest) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *V2RuleFunctionRequest) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *V2RuleFunctionRequest) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *V2RuleFunctionRequest) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### GetFnUuid

`func (o *V2RuleFunctionRequest) GetFnUuid() string`

GetFnUuid returns the FnUuid field if non-nil, zero value otherwise.

### GetFnUuidOk

`func (o *V2RuleFunctionRequest) GetFnUuidOk() (*string, bool)`

GetFnUuidOk returns a tuple with the FnUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFnUuid

`func (o *V2RuleFunctionRequest) SetFnUuid(v string)`

SetFnUuid sets FnUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


