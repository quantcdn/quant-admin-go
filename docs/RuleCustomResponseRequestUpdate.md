# RuleCustomResponseRequestUpdate

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
**CustomResponseBody** | Pointer to **string** |  | [optional] 
**CustomResponseStatusCode** | Pointer to **int32** |  | [optional] [default to 200]

## Methods

### NewRuleCustomResponseRequestUpdate

`func NewRuleCustomResponseRequestUpdate() *RuleCustomResponseRequestUpdate`

NewRuleCustomResponseRequestUpdate instantiates a new RuleCustomResponseRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleCustomResponseRequestUpdateWithDefaults

`func NewRuleCustomResponseRequestUpdateWithDefaults() *RuleCustomResponseRequestUpdate`

NewRuleCustomResponseRequestUpdateWithDefaults instantiates a new RuleCustomResponseRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleCustomResponseRequestUpdate) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleCustomResponseRequestUpdate) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleCustomResponseRequestUpdate) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RuleCustomResponseRequestUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *RuleCustomResponseRequestUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleCustomResponseRequestUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleCustomResponseRequestUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleCustomResponseRequestUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleCustomResponseRequestUpdate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleCustomResponseRequestUpdate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleCustomResponseRequestUpdate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleCustomResponseRequestUpdate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *RuleCustomResponseRequestUpdate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RuleCustomResponseRequestUpdate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RuleCustomResponseRequestUpdate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RuleCustomResponseRequestUpdate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleCustomResponseRequestUpdate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleCustomResponseRequestUpdate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleCustomResponseRequestUpdate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RuleCustomResponseRequestUpdate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *RuleCustomResponseRequestUpdate) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleCustomResponseRequestUpdate) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleCustomResponseRequestUpdate) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RuleCustomResponseRequestUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCountry

`func (o *RuleCustomResponseRequestUpdate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleCustomResponseRequestUpdate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleCustomResponseRequestUpdate) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleCustomResponseRequestUpdate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleCustomResponseRequestUpdate) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleCustomResponseRequestUpdate) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleCustomResponseRequestUpdate) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleCustomResponseRequestUpdate) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleCustomResponseRequestUpdate) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleCustomResponseRequestUpdate) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleCustomResponseRequestUpdate) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleCustomResponseRequestUpdate) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleCustomResponseRequestUpdate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleCustomResponseRequestUpdate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleCustomResponseRequestUpdate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleCustomResponseRequestUpdate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleCustomResponseRequestUpdate) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleCustomResponseRequestUpdate) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleCustomResponseRequestUpdate) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleCustomResponseRequestUpdate) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleCustomResponseRequestUpdate) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleCustomResponseRequestUpdate) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleCustomResponseRequestUpdate) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleCustomResponseRequestUpdate) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleCustomResponseRequestUpdate) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleCustomResponseRequestUpdate) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleCustomResponseRequestUpdate) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleCustomResponseRequestUpdate) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleCustomResponseRequestUpdate) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleCustomResponseRequestUpdate) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleCustomResponseRequestUpdate) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleCustomResponseRequestUpdate) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleCustomResponseRequestUpdate) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleCustomResponseRequestUpdate) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleCustomResponseRequestUpdate) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleCustomResponseRequestUpdate) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetCustomResponseBody

`func (o *RuleCustomResponseRequestUpdate) GetCustomResponseBody() string`

GetCustomResponseBody returns the CustomResponseBody field if non-nil, zero value otherwise.

### GetCustomResponseBodyOk

`func (o *RuleCustomResponseRequestUpdate) GetCustomResponseBodyOk() (*string, bool)`

GetCustomResponseBodyOk returns a tuple with the CustomResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseBody

`func (o *RuleCustomResponseRequestUpdate) SetCustomResponseBody(v string)`

SetCustomResponseBody sets CustomResponseBody field to given value.

### HasCustomResponseBody

`func (o *RuleCustomResponseRequestUpdate) HasCustomResponseBody() bool`

HasCustomResponseBody returns a boolean if a field has been set.

### GetCustomResponseStatusCode

`func (o *RuleCustomResponseRequestUpdate) GetCustomResponseStatusCode() int32`

GetCustomResponseStatusCode returns the CustomResponseStatusCode field if non-nil, zero value otherwise.

### GetCustomResponseStatusCodeOk

`func (o *RuleCustomResponseRequestUpdate) GetCustomResponseStatusCodeOk() (*int32, bool)`

GetCustomResponseStatusCodeOk returns a tuple with the CustomResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseStatusCode

`func (o *RuleCustomResponseRequestUpdate) SetCustomResponseStatusCode(v int32)`

SetCustomResponseStatusCode sets CustomResponseStatusCode field to given value.

### HasCustomResponseStatusCode

`func (o *RuleCustomResponseRequestUpdate) HasCustomResponseStatusCode() bool`

HasCustomResponseStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


