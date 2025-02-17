# RuleRedirectRequestUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **[]string** |  | [optional] [default to ["any"]]
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] [default to 0]
**Disabled** | Pointer to **bool** |  | [optional] [default to false]
**Url** | Pointer to **[]string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryIs** | Pointer to **[]string** |  | [optional] 
**CountryIsNot** | Pointer to **[]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**MethodIs** | Pointer to **[]string** |  | [optional] 
**MethodIsNot** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IpIs** | Pointer to **[]string** |  | [optional] 
**IpIsNot** | Pointer to **[]string** |  | [optional] 
**RedirectTo** | Pointer to **string** |  | [optional] 
**RedirectCode** | Pointer to **string** |  | [optional] [default to "301"]

## Methods

### NewRuleRedirectRequestUpdate

`func NewRuleRedirectRequestUpdate() *RuleRedirectRequestUpdate`

NewRuleRedirectRequestUpdate instantiates a new RuleRedirectRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleRedirectRequestUpdateWithDefaults

`func NewRuleRedirectRequestUpdateWithDefaults() *RuleRedirectRequestUpdate`

NewRuleRedirectRequestUpdateWithDefaults instantiates a new RuleRedirectRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleRedirectRequestUpdate) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleRedirectRequestUpdate) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleRedirectRequestUpdate) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RuleRedirectRequestUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *RuleRedirectRequestUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleRedirectRequestUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleRedirectRequestUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleRedirectRequestUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleRedirectRequestUpdate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleRedirectRequestUpdate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleRedirectRequestUpdate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleRedirectRequestUpdate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *RuleRedirectRequestUpdate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RuleRedirectRequestUpdate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RuleRedirectRequestUpdate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RuleRedirectRequestUpdate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleRedirectRequestUpdate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleRedirectRequestUpdate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleRedirectRequestUpdate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RuleRedirectRequestUpdate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *RuleRedirectRequestUpdate) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleRedirectRequestUpdate) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleRedirectRequestUpdate) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RuleRedirectRequestUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCountry

`func (o *RuleRedirectRequestUpdate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleRedirectRequestUpdate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleRedirectRequestUpdate) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleRedirectRequestUpdate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleRedirectRequestUpdate) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleRedirectRequestUpdate) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleRedirectRequestUpdate) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleRedirectRequestUpdate) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleRedirectRequestUpdate) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleRedirectRequestUpdate) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleRedirectRequestUpdate) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleRedirectRequestUpdate) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleRedirectRequestUpdate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleRedirectRequestUpdate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleRedirectRequestUpdate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleRedirectRequestUpdate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleRedirectRequestUpdate) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleRedirectRequestUpdate) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleRedirectRequestUpdate) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleRedirectRequestUpdate) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleRedirectRequestUpdate) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleRedirectRequestUpdate) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleRedirectRequestUpdate) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleRedirectRequestUpdate) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleRedirectRequestUpdate) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleRedirectRequestUpdate) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleRedirectRequestUpdate) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleRedirectRequestUpdate) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleRedirectRequestUpdate) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleRedirectRequestUpdate) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleRedirectRequestUpdate) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleRedirectRequestUpdate) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleRedirectRequestUpdate) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleRedirectRequestUpdate) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleRedirectRequestUpdate) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleRedirectRequestUpdate) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetRedirectTo

`func (o *RuleRedirectRequestUpdate) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *RuleRedirectRequestUpdate) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *RuleRedirectRequestUpdate) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *RuleRedirectRequestUpdate) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### GetRedirectCode

`func (o *RuleRedirectRequestUpdate) GetRedirectCode() string`

GetRedirectCode returns the RedirectCode field if non-nil, zero value otherwise.

### GetRedirectCodeOk

`func (o *RuleRedirectRequestUpdate) GetRedirectCodeOk() (*string, bool)`

GetRedirectCodeOk returns a tuple with the RedirectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectCode

`func (o *RuleRedirectRequestUpdate) SetRedirectCode(v string)`

SetRedirectCode sets RedirectCode field to given value.

### HasRedirectCode

`func (o *RuleRedirectRequestUpdate) HasRedirectCode() bool`

HasRedirectCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


