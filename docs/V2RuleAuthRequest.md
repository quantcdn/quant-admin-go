# V2RuleAuthRequest

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
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**AuthUser** | **string** | Authentication username | 
**AuthPass** | **string** | Authentication password | 

## Methods

### NewV2RuleAuthRequest

`func NewV2RuleAuthRequest(domain []string, url []string, message string, error_ bool, authUser string, authPass string, ) *V2RuleAuthRequest`

NewV2RuleAuthRequest instantiates a new V2RuleAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleAuthRequestWithDefaults

`func NewV2RuleAuthRequestWithDefaults() *V2RuleAuthRequest`

NewV2RuleAuthRequestWithDefaults instantiates a new V2RuleAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *V2RuleAuthRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleAuthRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleAuthRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *V2RuleAuthRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleAuthRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleAuthRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleAuthRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleAuthRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleAuthRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleAuthRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2RuleAuthRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleAuthRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleAuthRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleAuthRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleAuthRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleAuthRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleAuthRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleAuthRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V2RuleAuthRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleAuthRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleAuthRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleAuthRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *V2RuleAuthRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleAuthRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleAuthRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleAuthRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleAuthRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleAuthRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleAuthRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleAuthRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleAuthRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleAuthRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleAuthRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleAuthRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleAuthRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleAuthRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleAuthRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleAuthRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleAuthRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleAuthRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleAuthRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleAuthRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleAuthRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleAuthRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleAuthRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleAuthRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleAuthRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleAuthRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleAuthRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleAuthRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleAuthRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleAuthRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleAuthRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleAuthRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleAuthRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleAuthRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleAuthRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleAuthRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetMessage

`func (o *V2RuleAuthRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleAuthRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleAuthRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleAuthRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleAuthRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleAuthRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetAuthUser

`func (o *V2RuleAuthRequest) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *V2RuleAuthRequest) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *V2RuleAuthRequest) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.


### GetAuthPass

`func (o *V2RuleAuthRequest) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *V2RuleAuthRequest) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *V2RuleAuthRequest) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


