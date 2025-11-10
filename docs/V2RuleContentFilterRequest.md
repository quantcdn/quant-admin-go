# V2RuleContentFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **[]string** | Domain patterns (default: any) | 
**Name** | Pointer to **string** | Rule name | [optional] 
**Uuid** | Pointer to **string** | Rule UUID | [optional] 
**Weight** | Pointer to **int32** | Rule weight | [optional] [default to 0]
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
**FnUuid** | **string** | Function UUID | 

## Methods

### NewV2RuleContentFilterRequest

`func NewV2RuleContentFilterRequest(domain []string, url []string, fnUuid string, ) *V2RuleContentFilterRequest`

NewV2RuleContentFilterRequest instantiates a new V2RuleContentFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleContentFilterRequestWithDefaults

`func NewV2RuleContentFilterRequestWithDefaults() *V2RuleContentFilterRequest`

NewV2RuleContentFilterRequestWithDefaults instantiates a new V2RuleContentFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *V2RuleContentFilterRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleContentFilterRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleContentFilterRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *V2RuleContentFilterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleContentFilterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleContentFilterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleContentFilterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleContentFilterRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleContentFilterRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleContentFilterRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2RuleContentFilterRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleContentFilterRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleContentFilterRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleContentFilterRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleContentFilterRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleContentFilterRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleContentFilterRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleContentFilterRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V2RuleContentFilterRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleContentFilterRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleContentFilterRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleContentFilterRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *V2RuleContentFilterRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleContentFilterRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleContentFilterRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleContentFilterRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleContentFilterRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleContentFilterRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleContentFilterRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleContentFilterRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleContentFilterRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleContentFilterRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleContentFilterRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleContentFilterRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleContentFilterRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleContentFilterRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleContentFilterRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleContentFilterRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleContentFilterRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleContentFilterRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleContentFilterRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleContentFilterRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleContentFilterRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleContentFilterRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleContentFilterRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleContentFilterRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleContentFilterRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleContentFilterRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleContentFilterRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleContentFilterRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleContentFilterRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleContentFilterRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleContentFilterRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleContentFilterRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleContentFilterRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleContentFilterRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleContentFilterRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleContentFilterRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetFnUuid

`func (o *V2RuleContentFilterRequest) GetFnUuid() string`

GetFnUuid returns the FnUuid field if non-nil, zero value otherwise.

### GetFnUuidOk

`func (o *V2RuleContentFilterRequest) GetFnUuidOk() (*string, bool)`

GetFnUuidOk returns a tuple with the FnUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFnUuid

`func (o *V2RuleContentFilterRequest) SetFnUuid(v string)`

SetFnUuid sets FnUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


