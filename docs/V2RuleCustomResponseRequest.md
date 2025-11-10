# V2RuleCustomResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
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
**CustomResponseBody** | **string** | Custom response body content | 
**CustomResponseStatusCode** | Pointer to **int32** | HTTP status code for custom response | [optional] [default to 200]
**StatusCode** | Pointer to **int32** | Legacy field for status code (deprecated) | [optional] 
**Body** | Pointer to **string** | Legacy field for response body (deprecated) | [optional] 

## Methods

### NewV2RuleCustomResponseRequest

`func NewV2RuleCustomResponseRequest(message string, error_ bool, domain []string, url []string, customResponseBody string, ) *V2RuleCustomResponseRequest`

NewV2RuleCustomResponseRequest instantiates a new V2RuleCustomResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleCustomResponseRequestWithDefaults

`func NewV2RuleCustomResponseRequestWithDefaults() *V2RuleCustomResponseRequest`

NewV2RuleCustomResponseRequestWithDefaults instantiates a new V2RuleCustomResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleCustomResponseRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleCustomResponseRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleCustomResponseRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleCustomResponseRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleCustomResponseRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleCustomResponseRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetDomain

`func (o *V2RuleCustomResponseRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleCustomResponseRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleCustomResponseRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *V2RuleCustomResponseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleCustomResponseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleCustomResponseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleCustomResponseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleCustomResponseRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleCustomResponseRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleCustomResponseRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2RuleCustomResponseRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleCustomResponseRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleCustomResponseRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleCustomResponseRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleCustomResponseRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleCustomResponseRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleCustomResponseRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleCustomResponseRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V2RuleCustomResponseRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleCustomResponseRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleCustomResponseRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleCustomResponseRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *V2RuleCustomResponseRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleCustomResponseRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleCustomResponseRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleCustomResponseRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleCustomResponseRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleCustomResponseRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleCustomResponseRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleCustomResponseRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleCustomResponseRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleCustomResponseRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleCustomResponseRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleCustomResponseRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleCustomResponseRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleCustomResponseRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleCustomResponseRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleCustomResponseRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleCustomResponseRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleCustomResponseRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleCustomResponseRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleCustomResponseRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleCustomResponseRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleCustomResponseRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleCustomResponseRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleCustomResponseRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleCustomResponseRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleCustomResponseRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleCustomResponseRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleCustomResponseRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleCustomResponseRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleCustomResponseRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleCustomResponseRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleCustomResponseRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleCustomResponseRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleCustomResponseRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleCustomResponseRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleCustomResponseRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetCustomResponseBody

`func (o *V2RuleCustomResponseRequest) GetCustomResponseBody() string`

GetCustomResponseBody returns the CustomResponseBody field if non-nil, zero value otherwise.

### GetCustomResponseBodyOk

`func (o *V2RuleCustomResponseRequest) GetCustomResponseBodyOk() (*string, bool)`

GetCustomResponseBodyOk returns a tuple with the CustomResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseBody

`func (o *V2RuleCustomResponseRequest) SetCustomResponseBody(v string)`

SetCustomResponseBody sets CustomResponseBody field to given value.


### GetCustomResponseStatusCode

`func (o *V2RuleCustomResponseRequest) GetCustomResponseStatusCode() int32`

GetCustomResponseStatusCode returns the CustomResponseStatusCode field if non-nil, zero value otherwise.

### GetCustomResponseStatusCodeOk

`func (o *V2RuleCustomResponseRequest) GetCustomResponseStatusCodeOk() (*int32, bool)`

GetCustomResponseStatusCodeOk returns a tuple with the CustomResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseStatusCode

`func (o *V2RuleCustomResponseRequest) SetCustomResponseStatusCode(v int32)`

SetCustomResponseStatusCode sets CustomResponseStatusCode field to given value.

### HasCustomResponseStatusCode

`func (o *V2RuleCustomResponseRequest) HasCustomResponseStatusCode() bool`

HasCustomResponseStatusCode returns a boolean if a field has been set.

### GetStatusCode

`func (o *V2RuleCustomResponseRequest) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *V2RuleCustomResponseRequest) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *V2RuleCustomResponseRequest) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *V2RuleCustomResponseRequest) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetBody

`func (o *V2RuleCustomResponseRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *V2RuleCustomResponseRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *V2RuleCustomResponseRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *V2RuleCustomResponseRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


