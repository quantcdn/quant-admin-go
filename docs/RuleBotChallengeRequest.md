# RuleBotChallengeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **[]string** |  | [default to ["any"]]
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] [default to 0]
**Disabled** | **bool** |  | [default to false]
**Url** | **[]string** |  | 
**Country** | Pointer to **string** |  | [optional] 
**CountryIs** | Pointer to **[]string** |  | [optional] 
**CountryIsNot** | Pointer to **[]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**MethodIs** | Pointer to **[]string** |  | [optional] 
**MethodIsNot** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**IpIs** | Pointer to **[]string** |  | [optional] 
**IpIsNot** | Pointer to **[]string** |  | [optional] 
**RobotChallengeType** | **string** |  | 
**RobotChallengeVerificationTtl** | **int32** |  | [default to 10800]
**RobotChallengeChallengeTtl** | **int32** |  | [default to 30]

## Methods

### NewRuleBotChallengeRequest

`func NewRuleBotChallengeRequest(domain []string, disabled bool, url []string, robotChallengeType string, robotChallengeVerificationTtl int32, robotChallengeChallengeTtl int32, ) *RuleBotChallengeRequest`

NewRuleBotChallengeRequest instantiates a new RuleBotChallengeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleBotChallengeRequestWithDefaults

`func NewRuleBotChallengeRequestWithDefaults() *RuleBotChallengeRequest`

NewRuleBotChallengeRequestWithDefaults instantiates a new RuleBotChallengeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleBotChallengeRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleBotChallengeRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleBotChallengeRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *RuleBotChallengeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleBotChallengeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleBotChallengeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleBotChallengeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleBotChallengeRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleBotChallengeRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleBotChallengeRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleBotChallengeRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *RuleBotChallengeRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RuleBotChallengeRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RuleBotChallengeRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RuleBotChallengeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleBotChallengeRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleBotChallengeRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleBotChallengeRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetUrl

`func (o *RuleBotChallengeRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleBotChallengeRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleBotChallengeRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *RuleBotChallengeRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleBotChallengeRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleBotChallengeRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleBotChallengeRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleBotChallengeRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleBotChallengeRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleBotChallengeRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleBotChallengeRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleBotChallengeRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleBotChallengeRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleBotChallengeRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleBotChallengeRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleBotChallengeRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleBotChallengeRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleBotChallengeRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleBotChallengeRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleBotChallengeRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleBotChallengeRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleBotChallengeRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleBotChallengeRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleBotChallengeRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleBotChallengeRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleBotChallengeRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleBotChallengeRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleBotChallengeRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleBotChallengeRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleBotChallengeRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleBotChallengeRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleBotChallengeRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleBotChallengeRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleBotChallengeRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleBotChallengeRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleBotChallengeRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleBotChallengeRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleBotChallengeRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleBotChallengeRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetRobotChallengeType

`func (o *RuleBotChallengeRequest) GetRobotChallengeType() string`

GetRobotChallengeType returns the RobotChallengeType field if non-nil, zero value otherwise.

### GetRobotChallengeTypeOk

`func (o *RuleBotChallengeRequest) GetRobotChallengeTypeOk() (*string, bool)`

GetRobotChallengeTypeOk returns a tuple with the RobotChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeType

`func (o *RuleBotChallengeRequest) SetRobotChallengeType(v string)`

SetRobotChallengeType sets RobotChallengeType field to given value.


### GetRobotChallengeVerificationTtl

`func (o *RuleBotChallengeRequest) GetRobotChallengeVerificationTtl() int32`

GetRobotChallengeVerificationTtl returns the RobotChallengeVerificationTtl field if non-nil, zero value otherwise.

### GetRobotChallengeVerificationTtlOk

`func (o *RuleBotChallengeRequest) GetRobotChallengeVerificationTtlOk() (*int32, bool)`

GetRobotChallengeVerificationTtlOk returns a tuple with the RobotChallengeVerificationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeVerificationTtl

`func (o *RuleBotChallengeRequest) SetRobotChallengeVerificationTtl(v int32)`

SetRobotChallengeVerificationTtl sets RobotChallengeVerificationTtl field to given value.


### GetRobotChallengeChallengeTtl

`func (o *RuleBotChallengeRequest) GetRobotChallengeChallengeTtl() int32`

GetRobotChallengeChallengeTtl returns the RobotChallengeChallengeTtl field if non-nil, zero value otherwise.

### GetRobotChallengeChallengeTtlOk

`func (o *RuleBotChallengeRequest) GetRobotChallengeChallengeTtlOk() (*int32, bool)`

GetRobotChallengeChallengeTtlOk returns a tuple with the RobotChallengeChallengeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeChallengeTtl

`func (o *RuleBotChallengeRequest) SetRobotChallengeChallengeTtl(v int32)`

SetRobotChallengeChallengeTtl sets RobotChallengeChallengeTtl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


