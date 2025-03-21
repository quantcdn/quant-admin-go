# RuleServeStaticRequestUpdate

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
**StaticFilePath** | Pointer to **string** |  | [optional] 

## Methods

### NewRuleServeStaticRequestUpdate

`func NewRuleServeStaticRequestUpdate() *RuleServeStaticRequestUpdate`

NewRuleServeStaticRequestUpdate instantiates a new RuleServeStaticRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleServeStaticRequestUpdateWithDefaults

`func NewRuleServeStaticRequestUpdateWithDefaults() *RuleServeStaticRequestUpdate`

NewRuleServeStaticRequestUpdateWithDefaults instantiates a new RuleServeStaticRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleServeStaticRequestUpdate) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleServeStaticRequestUpdate) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleServeStaticRequestUpdate) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RuleServeStaticRequestUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *RuleServeStaticRequestUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleServeStaticRequestUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleServeStaticRequestUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleServeStaticRequestUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleServeStaticRequestUpdate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleServeStaticRequestUpdate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleServeStaticRequestUpdate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleServeStaticRequestUpdate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *RuleServeStaticRequestUpdate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RuleServeStaticRequestUpdate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RuleServeStaticRequestUpdate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RuleServeStaticRequestUpdate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleServeStaticRequestUpdate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleServeStaticRequestUpdate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleServeStaticRequestUpdate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RuleServeStaticRequestUpdate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *RuleServeStaticRequestUpdate) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleServeStaticRequestUpdate) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleServeStaticRequestUpdate) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RuleServeStaticRequestUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCountry

`func (o *RuleServeStaticRequestUpdate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleServeStaticRequestUpdate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleServeStaticRequestUpdate) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleServeStaticRequestUpdate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleServeStaticRequestUpdate) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleServeStaticRequestUpdate) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleServeStaticRequestUpdate) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleServeStaticRequestUpdate) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleServeStaticRequestUpdate) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleServeStaticRequestUpdate) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleServeStaticRequestUpdate) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleServeStaticRequestUpdate) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleServeStaticRequestUpdate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleServeStaticRequestUpdate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleServeStaticRequestUpdate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleServeStaticRequestUpdate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleServeStaticRequestUpdate) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleServeStaticRequestUpdate) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleServeStaticRequestUpdate) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleServeStaticRequestUpdate) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleServeStaticRequestUpdate) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleServeStaticRequestUpdate) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleServeStaticRequestUpdate) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleServeStaticRequestUpdate) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleServeStaticRequestUpdate) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleServeStaticRequestUpdate) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleServeStaticRequestUpdate) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleServeStaticRequestUpdate) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleServeStaticRequestUpdate) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleServeStaticRequestUpdate) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleServeStaticRequestUpdate) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleServeStaticRequestUpdate) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleServeStaticRequestUpdate) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleServeStaticRequestUpdate) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleServeStaticRequestUpdate) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleServeStaticRequestUpdate) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetStaticFilePath

`func (o *RuleServeStaticRequestUpdate) GetStaticFilePath() string`

GetStaticFilePath returns the StaticFilePath field if non-nil, zero value otherwise.

### GetStaticFilePathOk

`func (o *RuleServeStaticRequestUpdate) GetStaticFilePathOk() (*string, bool)`

GetStaticFilePathOk returns a tuple with the StaticFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticFilePath

`func (o *RuleServeStaticRequestUpdate) SetStaticFilePath(v string)`

SetStaticFilePath sets StaticFilePath field to given value.

### HasStaticFilePath

`func (o *RuleServeStaticRequestUpdate) HasStaticFilePath() bool`

HasStaticFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


