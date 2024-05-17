# RuleCustomResponserequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomResponseStatusCode** | **int32** |  | [default to 200]
**CustomResponseBody** | **string** |  | 
**Domain** | **string** |  | [default to "any"]
**Name** | Pointer to **string** |  | [optional] 
**Disabled** | **bool** |  | [default to false]
**Country** | Pointer to **string** |  | [optional] 
**CountryIs** | Pointer to **[]string** |  | [optional] 
**CountryIsNot** | Pointer to **[]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**MethodIs** | Pointer to **[]string** |  | [optional] 
**MethodIsNot** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IpIs** | Pointer to **[]string** |  | [optional] 
**IpIsNot** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRuleCustomResponserequest

`func NewRuleCustomResponserequest(customResponseStatusCode int32, customResponseBody string, domain string, disabled bool, ) *RuleCustomResponserequest`

NewRuleCustomResponserequest instantiates a new RuleCustomResponserequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleCustomResponserequestWithDefaults

`func NewRuleCustomResponserequestWithDefaults() *RuleCustomResponserequest`

NewRuleCustomResponserequestWithDefaults instantiates a new RuleCustomResponserequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomResponseStatusCode

`func (o *RuleCustomResponserequest) GetCustomResponseStatusCode() int32`

GetCustomResponseStatusCode returns the CustomResponseStatusCode field if non-nil, zero value otherwise.

### GetCustomResponseStatusCodeOk

`func (o *RuleCustomResponserequest) GetCustomResponseStatusCodeOk() (*int32, bool)`

GetCustomResponseStatusCodeOk returns a tuple with the CustomResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseStatusCode

`func (o *RuleCustomResponserequest) SetCustomResponseStatusCode(v int32)`

SetCustomResponseStatusCode sets CustomResponseStatusCode field to given value.


### GetCustomResponseBody

`func (o *RuleCustomResponserequest) GetCustomResponseBody() string`

GetCustomResponseBody returns the CustomResponseBody field if non-nil, zero value otherwise.

### GetCustomResponseBodyOk

`func (o *RuleCustomResponserequest) GetCustomResponseBodyOk() (*string, bool)`

GetCustomResponseBodyOk returns a tuple with the CustomResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseBody

`func (o *RuleCustomResponserequest) SetCustomResponseBody(v string)`

SetCustomResponseBody sets CustomResponseBody field to given value.


### GetDomain

`func (o *RuleCustomResponserequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleCustomResponserequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleCustomResponserequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *RuleCustomResponserequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleCustomResponserequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleCustomResponserequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleCustomResponserequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleCustomResponserequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleCustomResponserequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleCustomResponserequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetCountry

`func (o *RuleCustomResponserequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleCustomResponserequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleCustomResponserequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleCustomResponserequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleCustomResponserequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleCustomResponserequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleCustomResponserequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleCustomResponserequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleCustomResponserequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleCustomResponserequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleCustomResponserequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleCustomResponserequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleCustomResponserequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleCustomResponserequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleCustomResponserequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleCustomResponserequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleCustomResponserequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleCustomResponserequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleCustomResponserequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleCustomResponserequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleCustomResponserequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleCustomResponserequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleCustomResponserequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleCustomResponserequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleCustomResponserequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleCustomResponserequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleCustomResponserequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleCustomResponserequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleCustomResponserequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleCustomResponserequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleCustomResponserequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleCustomResponserequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleCustomResponserequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleCustomResponserequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleCustomResponserequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleCustomResponserequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


