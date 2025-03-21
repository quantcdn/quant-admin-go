# RuleBotChallengeRequestUpdate

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
**RobotChallengeType** | Pointer to **string** |  | [optional] 
**RobotChallengeVerificationTtl** | Pointer to **int32** |  | [optional] [default to 10800]
**RobotChallengeChallengeTtl** | Pointer to **int32** |  | [optional] [default to 30]

## Methods

### NewRuleBotChallengeRequestUpdate

`func NewRuleBotChallengeRequestUpdate() *RuleBotChallengeRequestUpdate`

NewRuleBotChallengeRequestUpdate instantiates a new RuleBotChallengeRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleBotChallengeRequestUpdateWithDefaults

`func NewRuleBotChallengeRequestUpdateWithDefaults() *RuleBotChallengeRequestUpdate`

NewRuleBotChallengeRequestUpdateWithDefaults instantiates a new RuleBotChallengeRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RuleBotChallengeRequestUpdate) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RuleBotChallengeRequestUpdate) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RuleBotChallengeRequestUpdate) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RuleBotChallengeRequestUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *RuleBotChallengeRequestUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleBotChallengeRequestUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleBotChallengeRequestUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleBotChallengeRequestUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *RuleBotChallengeRequestUpdate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RuleBotChallengeRequestUpdate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RuleBotChallengeRequestUpdate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RuleBotChallengeRequestUpdate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *RuleBotChallengeRequestUpdate) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RuleBotChallengeRequestUpdate) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RuleBotChallengeRequestUpdate) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RuleBotChallengeRequestUpdate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *RuleBotChallengeRequestUpdate) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleBotChallengeRequestUpdate) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleBotChallengeRequestUpdate) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RuleBotChallengeRequestUpdate) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *RuleBotChallengeRequestUpdate) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RuleBotChallengeRequestUpdate) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RuleBotChallengeRequestUpdate) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RuleBotChallengeRequestUpdate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCountry

`func (o *RuleBotChallengeRequestUpdate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RuleBotChallengeRequestUpdate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RuleBotChallengeRequestUpdate) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RuleBotChallengeRequestUpdate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *RuleBotChallengeRequestUpdate) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *RuleBotChallengeRequestUpdate) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *RuleBotChallengeRequestUpdate) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *RuleBotChallengeRequestUpdate) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *RuleBotChallengeRequestUpdate) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *RuleBotChallengeRequestUpdate) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *RuleBotChallengeRequestUpdate) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *RuleBotChallengeRequestUpdate) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *RuleBotChallengeRequestUpdate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RuleBotChallengeRequestUpdate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RuleBotChallengeRequestUpdate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RuleBotChallengeRequestUpdate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *RuleBotChallengeRequestUpdate) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *RuleBotChallengeRequestUpdate) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *RuleBotChallengeRequestUpdate) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *RuleBotChallengeRequestUpdate) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *RuleBotChallengeRequestUpdate) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *RuleBotChallengeRequestUpdate) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *RuleBotChallengeRequestUpdate) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *RuleBotChallengeRequestUpdate) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *RuleBotChallengeRequestUpdate) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RuleBotChallengeRequestUpdate) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RuleBotChallengeRequestUpdate) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RuleBotChallengeRequestUpdate) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *RuleBotChallengeRequestUpdate) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *RuleBotChallengeRequestUpdate) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *RuleBotChallengeRequestUpdate) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *RuleBotChallengeRequestUpdate) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *RuleBotChallengeRequestUpdate) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *RuleBotChallengeRequestUpdate) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *RuleBotChallengeRequestUpdate) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *RuleBotChallengeRequestUpdate) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetRobotChallengeType

`func (o *RuleBotChallengeRequestUpdate) GetRobotChallengeType() string`

GetRobotChallengeType returns the RobotChallengeType field if non-nil, zero value otherwise.

### GetRobotChallengeTypeOk

`func (o *RuleBotChallengeRequestUpdate) GetRobotChallengeTypeOk() (*string, bool)`

GetRobotChallengeTypeOk returns a tuple with the RobotChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeType

`func (o *RuleBotChallengeRequestUpdate) SetRobotChallengeType(v string)`

SetRobotChallengeType sets RobotChallengeType field to given value.

### HasRobotChallengeType

`func (o *RuleBotChallengeRequestUpdate) HasRobotChallengeType() bool`

HasRobotChallengeType returns a boolean if a field has been set.

### GetRobotChallengeVerificationTtl

`func (o *RuleBotChallengeRequestUpdate) GetRobotChallengeVerificationTtl() int32`

GetRobotChallengeVerificationTtl returns the RobotChallengeVerificationTtl field if non-nil, zero value otherwise.

### GetRobotChallengeVerificationTtlOk

`func (o *RuleBotChallengeRequestUpdate) GetRobotChallengeVerificationTtlOk() (*int32, bool)`

GetRobotChallengeVerificationTtlOk returns a tuple with the RobotChallengeVerificationTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeVerificationTtl

`func (o *RuleBotChallengeRequestUpdate) SetRobotChallengeVerificationTtl(v int32)`

SetRobotChallengeVerificationTtl sets RobotChallengeVerificationTtl field to given value.

### HasRobotChallengeVerificationTtl

`func (o *RuleBotChallengeRequestUpdate) HasRobotChallengeVerificationTtl() bool`

HasRobotChallengeVerificationTtl returns a boolean if a field has been set.

### GetRobotChallengeChallengeTtl

`func (o *RuleBotChallengeRequestUpdate) GetRobotChallengeChallengeTtl() int32`

GetRobotChallengeChallengeTtl returns the RobotChallengeChallengeTtl field if non-nil, zero value otherwise.

### GetRobotChallengeChallengeTtlOk

`func (o *RuleBotChallengeRequestUpdate) GetRobotChallengeChallengeTtlOk() (*int32, bool)`

GetRobotChallengeChallengeTtlOk returns a tuple with the RobotChallengeChallengeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRobotChallengeChallengeTtl

`func (o *RuleBotChallengeRequestUpdate) SetRobotChallengeChallengeTtl(v int32)`

SetRobotChallengeChallengeTtl sets RobotChallengeChallengeTtl field to given value.

### HasRobotChallengeChallengeTtl

`func (o *RuleBotChallengeRequestUpdate) HasRobotChallengeChallengeTtl() bool`

HasRobotChallengeChallengeTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


