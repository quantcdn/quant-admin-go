# V2RuleProxyRequest

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
**To** | **string** | Target URL to proxy to | 
**Host** | Pointer to **string** | Host header override | [optional] 
**AuthUser** | Pointer to **string** | Basic auth username | [optional] [default to ""]
**AuthPass** | Pointer to **string** | Basic auth password | [optional] [default to ""]
**DisableSslVerify** | Pointer to **bool** | Disable SSL verification | [optional] [default to false]
**CacheLifetime** | Pointer to **NullableString** | Cache lifetime | [optional] 
**OnlyProxy404** | Pointer to **bool** | Only proxy 404 responses | [optional] [default to false]
**InjectHeaders** | Pointer to **map[string]string** | Headers to inject | [optional] 
**ProxyStripHeaders** | Pointer to **[]string** | Headers to strip from response | [optional] 
**ProxyStripRequestHeaders** | Pointer to **[]string** | Headers to strip from request | [optional] 
**OriginTimeout** | Pointer to **string** | Origin timeout | [optional] 
**FailoverMode** | Pointer to **bool** | Enable failover mode | [optional] [default to false]
**FailoverOriginTtfb** | Pointer to **string** | Failover TTFB threshold | [optional] [default to "2000"]
**FailoverOriginStatusCodes** | Pointer to **[]string** | Status codes for failover (default: 200,404,301,302,304) | [optional] 
**FailoverLifetime** | Pointer to **string** | Failover cache lifetime | [optional] [default to "300"]
**Notify** | Pointer to **string** | Notification type (none, slack) | [optional] [default to "none"]
**NotifyConfig** | Pointer to [**NullableV2RuleProxyActionAllOfNotifyConfig**](V2RuleProxyActionAllOfNotifyConfig.md) |  | [optional] 
**WafEnabled** | Pointer to **bool** | WAF enabled | [optional] [default to false]
**WafConfig** | Pointer to [**NullableWafConfig**](WafConfig.md) |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** | Proxy alert enabled | [optional] 
**ProxyInlineFnEnabled** | Pointer to **bool** | Proxy inline function enabled | [optional] [default to false]
**QuantCloudSelection** | Pointer to [**NullableV2RuleProxyActionAllOfQuantCloudSelection**](V2RuleProxyActionAllOfQuantCloudSelection.md) |  | [optional] 
**StaticErrorPage** | Pointer to **string** | Static error page | [optional] 
**StaticErrorPageStatusCodes** | Pointer to **[]string** | Status codes for static error page | [optional] 
**ApplicationProxy** | Pointer to **bool** | Application proxy enabled | [optional] [default to false]
**ApplicationName** | Pointer to **string** | Application name | [optional] 
**ApplicationEnvironment** | Pointer to **string** | Application environment | [optional] 
**ApplicationContainer** | Pointer to **string** | Application container | [optional] 
**ApplicationPort** | Pointer to **int32** | Application port | [optional] 

## Methods

### NewV2RuleProxyRequest

`func NewV2RuleProxyRequest(message string, error_ bool, domain []string, url []string, to string, ) *V2RuleProxyRequest`

NewV2RuleProxyRequest instantiates a new V2RuleProxyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleProxyRequestWithDefaults

`func NewV2RuleProxyRequestWithDefaults() *V2RuleProxyRequest`

NewV2RuleProxyRequestWithDefaults instantiates a new V2RuleProxyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleProxyRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleProxyRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleProxyRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleProxyRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleProxyRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleProxyRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetDomain

`func (o *V2RuleProxyRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleProxyRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleProxyRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetName

`func (o *V2RuleProxyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleProxyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleProxyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleProxyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleProxyRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleProxyRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleProxyRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *V2RuleProxyRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleProxyRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleProxyRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleProxyRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleProxyRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleProxyRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleProxyRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleProxyRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *V2RuleProxyRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleProxyRequest) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleProxyRequest) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleProxyRequest) SetUrl(v []string)`

SetUrl sets Url field to given value.


### GetCountry

`func (o *V2RuleProxyRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleProxyRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleProxyRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleProxyRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleProxyRequest) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleProxyRequest) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleProxyRequest) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleProxyRequest) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleProxyRequest) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleProxyRequest) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleProxyRequest) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleProxyRequest) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleProxyRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleProxyRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleProxyRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleProxyRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleProxyRequest) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleProxyRequest) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleProxyRequest) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleProxyRequest) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleProxyRequest) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleProxyRequest) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleProxyRequest) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleProxyRequest) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleProxyRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleProxyRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleProxyRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleProxyRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleProxyRequest) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleProxyRequest) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleProxyRequest) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleProxyRequest) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleProxyRequest) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleProxyRequest) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleProxyRequest) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleProxyRequest) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetTo

`func (o *V2RuleProxyRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V2RuleProxyRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V2RuleProxyRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetHost

`func (o *V2RuleProxyRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2RuleProxyRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2RuleProxyRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2RuleProxyRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAuthUser

`func (o *V2RuleProxyRequest) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *V2RuleProxyRequest) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *V2RuleProxyRequest) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *V2RuleProxyRequest) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetAuthPass

`func (o *V2RuleProxyRequest) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *V2RuleProxyRequest) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *V2RuleProxyRequest) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *V2RuleProxyRequest) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *V2RuleProxyRequest) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *V2RuleProxyRequest) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *V2RuleProxyRequest) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *V2RuleProxyRequest) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *V2RuleProxyRequest) GetCacheLifetime() string`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *V2RuleProxyRequest) GetCacheLifetimeOk() (*string, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *V2RuleProxyRequest) SetCacheLifetime(v string)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *V2RuleProxyRequest) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### SetCacheLifetimeNil

`func (o *V2RuleProxyRequest) SetCacheLifetimeNil(b bool)`

 SetCacheLifetimeNil sets the value for CacheLifetime to be an explicit nil

### UnsetCacheLifetime
`func (o *V2RuleProxyRequest) UnsetCacheLifetime()`

UnsetCacheLifetime ensures that no value is present for CacheLifetime, not even an explicit nil
### GetOnlyProxy404

`func (o *V2RuleProxyRequest) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *V2RuleProxyRequest) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *V2RuleProxyRequest) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *V2RuleProxyRequest) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *V2RuleProxyRequest) GetInjectHeaders() map[string]string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *V2RuleProxyRequest) GetInjectHeadersOk() (*map[string]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *V2RuleProxyRequest) SetInjectHeaders(v map[string]string)`

SetInjectHeaders sets InjectHeaders field to given value.

### HasInjectHeaders

`func (o *V2RuleProxyRequest) HasInjectHeaders() bool`

HasInjectHeaders returns a boolean if a field has been set.

### SetInjectHeadersNil

`func (o *V2RuleProxyRequest) SetInjectHeadersNil(b bool)`

 SetInjectHeadersNil sets the value for InjectHeaders to be an explicit nil

### UnsetInjectHeaders
`func (o *V2RuleProxyRequest) UnsetInjectHeaders()`

UnsetInjectHeaders ensures that no value is present for InjectHeaders, not even an explicit nil
### GetProxyStripHeaders

`func (o *V2RuleProxyRequest) GetProxyStripHeaders() []string`

GetProxyStripHeaders returns the ProxyStripHeaders field if non-nil, zero value otherwise.

### GetProxyStripHeadersOk

`func (o *V2RuleProxyRequest) GetProxyStripHeadersOk() (*[]string, bool)`

GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripHeaders

`func (o *V2RuleProxyRequest) SetProxyStripHeaders(v []string)`

SetProxyStripHeaders sets ProxyStripHeaders field to given value.

### HasProxyStripHeaders

`func (o *V2RuleProxyRequest) HasProxyStripHeaders() bool`

HasProxyStripHeaders returns a boolean if a field has been set.

### GetProxyStripRequestHeaders

`func (o *V2RuleProxyRequest) GetProxyStripRequestHeaders() []string`

GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field if non-nil, zero value otherwise.

### GetProxyStripRequestHeadersOk

`func (o *V2RuleProxyRequest) GetProxyStripRequestHeadersOk() (*[]string, bool)`

GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripRequestHeaders

`func (o *V2RuleProxyRequest) SetProxyStripRequestHeaders(v []string)`

SetProxyStripRequestHeaders sets ProxyStripRequestHeaders field to given value.

### HasProxyStripRequestHeaders

`func (o *V2RuleProxyRequest) HasProxyStripRequestHeaders() bool`

HasProxyStripRequestHeaders returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *V2RuleProxyRequest) GetOriginTimeout() string`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *V2RuleProxyRequest) GetOriginTimeoutOk() (*string, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *V2RuleProxyRequest) SetOriginTimeout(v string)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *V2RuleProxyRequest) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetFailoverMode

`func (o *V2RuleProxyRequest) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *V2RuleProxyRequest) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *V2RuleProxyRequest) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *V2RuleProxyRequest) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *V2RuleProxyRequest) GetFailoverOriginTtfb() string`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *V2RuleProxyRequest) GetFailoverOriginTtfbOk() (*string, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *V2RuleProxyRequest) SetFailoverOriginTtfb(v string)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *V2RuleProxyRequest) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverOriginStatusCodes

`func (o *V2RuleProxyRequest) GetFailoverOriginStatusCodes() []string`

GetFailoverOriginStatusCodes returns the FailoverOriginStatusCodes field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodesOk

`func (o *V2RuleProxyRequest) GetFailoverOriginStatusCodesOk() (*[]string, bool)`

GetFailoverOriginStatusCodesOk returns a tuple with the FailoverOriginStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCodes

`func (o *V2RuleProxyRequest) SetFailoverOriginStatusCodes(v []string)`

SetFailoverOriginStatusCodes sets FailoverOriginStatusCodes field to given value.

### HasFailoverOriginStatusCodes

`func (o *V2RuleProxyRequest) HasFailoverOriginStatusCodes() bool`

HasFailoverOriginStatusCodes returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *V2RuleProxyRequest) GetFailoverLifetime() string`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *V2RuleProxyRequest) GetFailoverLifetimeOk() (*string, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *V2RuleProxyRequest) SetFailoverLifetime(v string)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *V2RuleProxyRequest) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetNotify

`func (o *V2RuleProxyRequest) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *V2RuleProxyRequest) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *V2RuleProxyRequest) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *V2RuleProxyRequest) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *V2RuleProxyRequest) GetNotifyConfig() V2RuleProxyActionAllOfNotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *V2RuleProxyRequest) GetNotifyConfigOk() (*V2RuleProxyActionAllOfNotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *V2RuleProxyRequest) SetNotifyConfig(v V2RuleProxyActionAllOfNotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *V2RuleProxyRequest) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

### SetNotifyConfigNil

`func (o *V2RuleProxyRequest) SetNotifyConfigNil(b bool)`

 SetNotifyConfigNil sets the value for NotifyConfig to be an explicit nil

### UnsetNotifyConfig
`func (o *V2RuleProxyRequest) UnsetNotifyConfig()`

UnsetNotifyConfig ensures that no value is present for NotifyConfig, not even an explicit nil
### GetWafEnabled

`func (o *V2RuleProxyRequest) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *V2RuleProxyRequest) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *V2RuleProxyRequest) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *V2RuleProxyRequest) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.

### GetWafConfig

`func (o *V2RuleProxyRequest) GetWafConfig() WafConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *V2RuleProxyRequest) GetWafConfigOk() (*WafConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *V2RuleProxyRequest) SetWafConfig(v WafConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *V2RuleProxyRequest) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.

### SetWafConfigNil

`func (o *V2RuleProxyRequest) SetWafConfigNil(b bool)`

 SetWafConfigNil sets the value for WafConfig to be an explicit nil

### UnsetWafConfig
`func (o *V2RuleProxyRequest) UnsetWafConfig()`

UnsetWafConfig ensures that no value is present for WafConfig, not even an explicit nil
### GetProxyAlertEnabled

`func (o *V2RuleProxyRequest) GetProxyAlertEnabled() bool`

GetProxyAlertEnabled returns the ProxyAlertEnabled field if non-nil, zero value otherwise.

### GetProxyAlertEnabledOk

`func (o *V2RuleProxyRequest) GetProxyAlertEnabledOk() (*bool, bool)`

GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAlertEnabled

`func (o *V2RuleProxyRequest) SetProxyAlertEnabled(v bool)`

SetProxyAlertEnabled sets ProxyAlertEnabled field to given value.

### HasProxyAlertEnabled

`func (o *V2RuleProxyRequest) HasProxyAlertEnabled() bool`

HasProxyAlertEnabled returns a boolean if a field has been set.

### GetProxyInlineFnEnabled

`func (o *V2RuleProxyRequest) GetProxyInlineFnEnabled() bool`

GetProxyInlineFnEnabled returns the ProxyInlineFnEnabled field if non-nil, zero value otherwise.

### GetProxyInlineFnEnabledOk

`func (o *V2RuleProxyRequest) GetProxyInlineFnEnabledOk() (*bool, bool)`

GetProxyInlineFnEnabledOk returns a tuple with the ProxyInlineFnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyInlineFnEnabled

`func (o *V2RuleProxyRequest) SetProxyInlineFnEnabled(v bool)`

SetProxyInlineFnEnabled sets ProxyInlineFnEnabled field to given value.

### HasProxyInlineFnEnabled

`func (o *V2RuleProxyRequest) HasProxyInlineFnEnabled() bool`

HasProxyInlineFnEnabled returns a boolean if a field has been set.

### GetQuantCloudSelection

`func (o *V2RuleProxyRequest) GetQuantCloudSelection() V2RuleProxyActionAllOfQuantCloudSelection`

GetQuantCloudSelection returns the QuantCloudSelection field if non-nil, zero value otherwise.

### GetQuantCloudSelectionOk

`func (o *V2RuleProxyRequest) GetQuantCloudSelectionOk() (*V2RuleProxyActionAllOfQuantCloudSelection, bool)`

GetQuantCloudSelectionOk returns a tuple with the QuantCloudSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantCloudSelection

`func (o *V2RuleProxyRequest) SetQuantCloudSelection(v V2RuleProxyActionAllOfQuantCloudSelection)`

SetQuantCloudSelection sets QuantCloudSelection field to given value.

### HasQuantCloudSelection

`func (o *V2RuleProxyRequest) HasQuantCloudSelection() bool`

HasQuantCloudSelection returns a boolean if a field has been set.

### SetQuantCloudSelectionNil

`func (o *V2RuleProxyRequest) SetQuantCloudSelectionNil(b bool)`

 SetQuantCloudSelectionNil sets the value for QuantCloudSelection to be an explicit nil

### UnsetQuantCloudSelection
`func (o *V2RuleProxyRequest) UnsetQuantCloudSelection()`

UnsetQuantCloudSelection ensures that no value is present for QuantCloudSelection, not even an explicit nil
### GetStaticErrorPage

`func (o *V2RuleProxyRequest) GetStaticErrorPage() string`

GetStaticErrorPage returns the StaticErrorPage field if non-nil, zero value otherwise.

### GetStaticErrorPageOk

`func (o *V2RuleProxyRequest) GetStaticErrorPageOk() (*string, bool)`

GetStaticErrorPageOk returns a tuple with the StaticErrorPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticErrorPage

`func (o *V2RuleProxyRequest) SetStaticErrorPage(v string)`

SetStaticErrorPage sets StaticErrorPage field to given value.

### HasStaticErrorPage

`func (o *V2RuleProxyRequest) HasStaticErrorPage() bool`

HasStaticErrorPage returns a boolean if a field has been set.

### GetStaticErrorPageStatusCodes

`func (o *V2RuleProxyRequest) GetStaticErrorPageStatusCodes() []string`

GetStaticErrorPageStatusCodes returns the StaticErrorPageStatusCodes field if non-nil, zero value otherwise.

### GetStaticErrorPageStatusCodesOk

`func (o *V2RuleProxyRequest) GetStaticErrorPageStatusCodesOk() (*[]string, bool)`

GetStaticErrorPageStatusCodesOk returns a tuple with the StaticErrorPageStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticErrorPageStatusCodes

`func (o *V2RuleProxyRequest) SetStaticErrorPageStatusCodes(v []string)`

SetStaticErrorPageStatusCodes sets StaticErrorPageStatusCodes field to given value.

### HasStaticErrorPageStatusCodes

`func (o *V2RuleProxyRequest) HasStaticErrorPageStatusCodes() bool`

HasStaticErrorPageStatusCodes returns a boolean if a field has been set.

### GetApplicationProxy

`func (o *V2RuleProxyRequest) GetApplicationProxy() bool`

GetApplicationProxy returns the ApplicationProxy field if non-nil, zero value otherwise.

### GetApplicationProxyOk

`func (o *V2RuleProxyRequest) GetApplicationProxyOk() (*bool, bool)`

GetApplicationProxyOk returns a tuple with the ApplicationProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationProxy

`func (o *V2RuleProxyRequest) SetApplicationProxy(v bool)`

SetApplicationProxy sets ApplicationProxy field to given value.

### HasApplicationProxy

`func (o *V2RuleProxyRequest) HasApplicationProxy() bool`

HasApplicationProxy returns a boolean if a field has been set.

### GetApplicationName

`func (o *V2RuleProxyRequest) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *V2RuleProxyRequest) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *V2RuleProxyRequest) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *V2RuleProxyRequest) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationEnvironment

`func (o *V2RuleProxyRequest) GetApplicationEnvironment() string`

GetApplicationEnvironment returns the ApplicationEnvironment field if non-nil, zero value otherwise.

### GetApplicationEnvironmentOk

`func (o *V2RuleProxyRequest) GetApplicationEnvironmentOk() (*string, bool)`

GetApplicationEnvironmentOk returns a tuple with the ApplicationEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEnvironment

`func (o *V2RuleProxyRequest) SetApplicationEnvironment(v string)`

SetApplicationEnvironment sets ApplicationEnvironment field to given value.

### HasApplicationEnvironment

`func (o *V2RuleProxyRequest) HasApplicationEnvironment() bool`

HasApplicationEnvironment returns a boolean if a field has been set.

### GetApplicationContainer

`func (o *V2RuleProxyRequest) GetApplicationContainer() string`

GetApplicationContainer returns the ApplicationContainer field if non-nil, zero value otherwise.

### GetApplicationContainerOk

`func (o *V2RuleProxyRequest) GetApplicationContainerOk() (*string, bool)`

GetApplicationContainerOk returns a tuple with the ApplicationContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationContainer

`func (o *V2RuleProxyRequest) SetApplicationContainer(v string)`

SetApplicationContainer sets ApplicationContainer field to given value.

### HasApplicationContainer

`func (o *V2RuleProxyRequest) HasApplicationContainer() bool`

HasApplicationContainer returns a boolean if a field has been set.

### GetApplicationPort

`func (o *V2RuleProxyRequest) GetApplicationPort() int32`

GetApplicationPort returns the ApplicationPort field if non-nil, zero value otherwise.

### GetApplicationPortOk

`func (o *V2RuleProxyRequest) GetApplicationPortOk() (*int32, bool)`

GetApplicationPortOk returns a tuple with the ApplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPort

`func (o *V2RuleProxyRequest) SetApplicationPort(v int32)`

SetApplicationPort sets ApplicationPort field to given value.

### HasApplicationPort

`func (o *V2RuleProxyRequest) HasApplicationPort() bool`

HasApplicationPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


