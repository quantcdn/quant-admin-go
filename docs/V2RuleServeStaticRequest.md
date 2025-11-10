# V2RuleServeStaticRequest

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
**StaticFilePath** | **string** | Path to the static file to serve | 

## Methods

### NewV2RuleServeStaticRequest

`func NewV2RuleServeStaticRequest(domain []string, url []string, message string, error_ bool, staticFilePath string, ) *V2RuleServeStaticRequest`

NewV2RuleServeStaticRequest instantiates a new V2RuleServeStaticRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleServeStaticRequestWithDefaults

`func NewV2RuleServeStaticRequestWithDefaults() *V2RuleServeStaticRequest`

NewV2RuleServeStaticRequestWithDefaults instantiates a new V2RuleServeStaticRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *V2RuleServeStaticRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleServeStaticRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleServeStaticRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *V2RuleServeStaticRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleServeStaticRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleServeStaticRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleServeStaticRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleServeStaticRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleServeStaticRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleServeStaticRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2RuleServeStaticRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleServeStaticRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleServeStaticRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleServeStaticRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleServeStaticRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleServeStaticRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleServeStaticRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleServeStaticRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V2RuleServeStaticRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleServeStaticRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleServeStaticRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleServeStaticRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *V2RuleServeStaticRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleServeStaticRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleServeStaticRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleServeStaticRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleServeStaticRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleServeStaticRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleServeStaticRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleServeStaticRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleServeStaticRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleServeStaticRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleServeStaticRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleServeStaticRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleServeStaticRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleServeStaticRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleServeStaticRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleServeStaticRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleServeStaticRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleServeStaticRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleServeStaticRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleServeStaticRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleServeStaticRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleServeStaticRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleServeStaticRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleServeStaticRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleServeStaticRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleServeStaticRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleServeStaticRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleServeStaticRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleServeStaticRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleServeStaticRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleServeStaticRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleServeStaticRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleServeStaticRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleServeStaticRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleServeStaticRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleServeStaticRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetMessage

`func (o *V2RuleServeStaticRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleServeStaticRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleServeStaticRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleServeStaticRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleServeStaticRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleServeStaticRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetStaticFilePath

`func (o *V2RuleServeStaticRequest) GetStaticFilePath() string`

GetStaticFilePath returns the StaticFilePath field if non-nil, zero value otherwise.

### GetStaticFilePathOk

`func (o *V2RuleServeStaticRequest) GetStaticFilePathOk() (*string, bool)`

GetStaticFilePathOk returns a tuple with the StaticFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticFilePath

`func (o *V2RuleServeStaticRequest) SetStaticFilePath(v string)`

SetStaticFilePath sets StaticFilePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


