# V2RuleCustomResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Rule name | [optional] 
**Uuid** | **string** | Rule UUID | 
**RuleId** | Pointer to **string** | Rule ID | [optional] 
**Weight** | Pointer to **int32** | Rule weight | [optional] [default to 0]
**Url** | Pointer to **[]string** | URL patterns | [optional] 
**Domain** | Pointer to **[]string** | Domain patterns | [optional] 
**Disabled** | **bool** | Whether rule is disabled | [default to false]
**OnlyWithCookie** | Pointer to **string** | Only apply with cookie | [optional] 
**Method** | Pointer to **string** | HTTP method | [optional] 
**MethodIs** | Pointer to **[]string** | Allowed HTTP methods | [optional] 
**MethodIsNot** | Pointer to **[]string** | Excluded HTTP methods | [optional] 
**Ip** | Pointer to **string** | IP address | [optional] 
**IpIs** | Pointer to **[]string** | Allowed IP addresses | [optional] 
**IpIsNot** | Pointer to **[]string** | Excluded IP addresses | [optional] 
**Country** | Pointer to **string** | Country code | [optional] 
**CountryIs** | Pointer to **[]string** | Allowed countries | [optional] 
**CountryIsNot** | Pointer to **[]string** | Excluded countries | [optional] 
**Action** | **string** | Rule action | 
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**ActionConfig** | Pointer to [**V2RuleCustomResponseAction**](V2RuleCustomResponseAction.md) |  | [optional] 

## Methods

### NewV2RuleCustomResponse

`func NewV2RuleCustomResponse(uuid string, disabled bool, action string, message string, error_ bool, ) *V2RuleCustomResponse`

NewV2RuleCustomResponse instantiates a new V2RuleCustomResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleCustomResponseWithDefaults

`func NewV2RuleCustomResponseWithDefaults() *V2RuleCustomResponse`

NewV2RuleCustomResponseWithDefaults instantiates a new V2RuleCustomResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2RuleCustomResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleCustomResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleCustomResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleCustomResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleCustomResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleCustomResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleCustomResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetRuleId

`func (o *V2RuleCustomResponse) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *V2RuleCustomResponse) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *V2RuleCustomResponse) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *V2RuleCustomResponse) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleCustomResponse) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleCustomResponse) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleCustomResponse) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleCustomResponse) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleCustomResponse) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleCustomResponse) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleCustomResponse) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V2RuleCustomResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *V2RuleCustomResponse) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleCustomResponse) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleCustomResponse) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V2RuleCustomResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleCustomResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleCustomResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleCustomResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetOnlyWithCookie

`func (o *V2RuleCustomResponse) GetOnlyWithCookie() string`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *V2RuleCustomResponse) GetOnlyWithCookieOk() (*string, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *V2RuleCustomResponse) SetOnlyWithCookie(v string)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *V2RuleCustomResponse) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleCustomResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleCustomResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleCustomResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleCustomResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleCustomResponse) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleCustomResponse) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleCustomResponse) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleCustomResponse) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleCustomResponse) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleCustomResponse) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleCustomResponse) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleCustomResponse) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleCustomResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleCustomResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleCustomResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleCustomResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleCustomResponse) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleCustomResponse) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleCustomResponse) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleCustomResponse) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleCustomResponse) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleCustomResponse) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleCustomResponse) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleCustomResponse) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetCountry

`func (o *V2RuleCustomResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleCustomResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleCustomResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleCustomResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleCustomResponse) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleCustomResponse) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleCustomResponse) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleCustomResponse) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleCustomResponse) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleCustomResponse) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleCustomResponse) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleCustomResponse) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetAction

`func (o *V2RuleCustomResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V2RuleCustomResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V2RuleCustomResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetMessage

`func (o *V2RuleCustomResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleCustomResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleCustomResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleCustomResponse) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleCustomResponse) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleCustomResponse) SetError(v bool)`

SetError sets Error field to given value.


### GetActionConfig

`func (o *V2RuleCustomResponse) GetActionConfig() V2RuleCustomResponseAction`

GetActionConfig returns the ActionConfig field if non-nil, zero value otherwise.

### GetActionConfigOk

`func (o *V2RuleCustomResponse) GetActionConfigOk() (*V2RuleCustomResponseAction, bool)`

GetActionConfigOk returns a tuple with the ActionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConfig

`func (o *V2RuleCustomResponse) SetActionConfig(v V2RuleCustomResponseAction)`

SetActionConfig sets ActionConfig field to given value.

### HasActionConfig

`func (o *V2RuleCustomResponse) HasActionConfig() bool`

HasActionConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


