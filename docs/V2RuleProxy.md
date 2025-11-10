# V2RuleProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
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
**To** | **string** | Target URL to proxy to | 
**Host** | Pointer to **string** | Host header override | [optional] 
**AuthUser** | Pointer to **string** | Basic auth username | [optional] 
**AuthPass** | Pointer to **string** | Basic auth password | [optional] 
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
**WafConfig** | Pointer to [**WafConfig**](WafConfig.md) |  | [optional] 
**ProxyAlertEnabled** | Pointer to **bool** | Proxy alert enabled | [optional] 
**ProxyInlineFnEnabled** | Pointer to **bool** | Proxy inline function enabled | [optional] [default to false]
**QuantCloudSelection** | Pointer to [**NullableV2RuleProxyActionAllOfQuantCloudSelection**](V2RuleProxyActionAllOfQuantCloudSelection.md) |  | [optional] 
**ActionConfig** | [**V2RuleProxyAction**](V2RuleProxyAction.md) |  | 

## Methods

### NewV2RuleProxy

`func NewV2RuleProxy(message string, error_ bool, uuid string, disabled bool, action string, to string, actionConfig V2RuleProxyAction, ) *V2RuleProxy`

NewV2RuleProxy instantiates a new V2RuleProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleProxyWithDefaults

`func NewV2RuleProxyWithDefaults() *V2RuleProxy`

NewV2RuleProxyWithDefaults instantiates a new V2RuleProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleProxy) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleProxy) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleProxy) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleProxy) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleProxy) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleProxy) SetError(v bool)`

SetError sets Error field to given value.


### GetName

`func (o *V2RuleProxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2RuleProxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2RuleProxy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2RuleProxy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *V2RuleProxy) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *V2RuleProxy) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *V2RuleProxy) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetRuleId

`func (o *V2RuleProxy) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *V2RuleProxy) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *V2RuleProxy) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *V2RuleProxy) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetWeight

`func (o *V2RuleProxy) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V2RuleProxy) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V2RuleProxy) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V2RuleProxy) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetUrl

`func (o *V2RuleProxy) GetUrl() []string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V2RuleProxy) GetUrlOk() (*[]string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V2RuleProxy) SetUrl(v []string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V2RuleProxy) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *V2RuleProxy) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2RuleProxy) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2RuleProxy) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V2RuleProxy) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDisabled

`func (o *V2RuleProxy) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *V2RuleProxy) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *V2RuleProxy) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetOnlyWithCookie

`func (o *V2RuleProxy) GetOnlyWithCookie() string`

GetOnlyWithCookie returns the OnlyWithCookie field if non-nil, zero value otherwise.

### GetOnlyWithCookieOk

`func (o *V2RuleProxy) GetOnlyWithCookieOk() (*string, bool)`

GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyWithCookie

`func (o *V2RuleProxy) SetOnlyWithCookie(v string)`

SetOnlyWithCookie sets OnlyWithCookie field to given value.

### HasOnlyWithCookie

`func (o *V2RuleProxy) HasOnlyWithCookie() bool`

HasOnlyWithCookie returns a boolean if a field has been set.

### GetMethod

`func (o *V2RuleProxy) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V2RuleProxy) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V2RuleProxy) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V2RuleProxy) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethodIs

`func (o *V2RuleProxy) GetMethodIs() []string`

GetMethodIs returns the MethodIs field if non-nil, zero value otherwise.

### GetMethodIsOk

`func (o *V2RuleProxy) GetMethodIsOk() (*[]string, bool)`

GetMethodIsOk returns a tuple with the MethodIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIs

`func (o *V2RuleProxy) SetMethodIs(v []string)`

SetMethodIs sets MethodIs field to given value.

### HasMethodIs

`func (o *V2RuleProxy) HasMethodIs() bool`

HasMethodIs returns a boolean if a field has been set.

### GetMethodIsNot

`func (o *V2RuleProxy) GetMethodIsNot() []string`

GetMethodIsNot returns the MethodIsNot field if non-nil, zero value otherwise.

### GetMethodIsNotOk

`func (o *V2RuleProxy) GetMethodIsNotOk() (*[]string, bool)`

GetMethodIsNotOk returns a tuple with the MethodIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodIsNot

`func (o *V2RuleProxy) SetMethodIsNot(v []string)`

SetMethodIsNot sets MethodIsNot field to given value.

### HasMethodIsNot

`func (o *V2RuleProxy) HasMethodIsNot() bool`

HasMethodIsNot returns a boolean if a field has been set.

### GetIp

`func (o *V2RuleProxy) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *V2RuleProxy) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *V2RuleProxy) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *V2RuleProxy) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpIs

`func (o *V2RuleProxy) GetIpIs() []string`

GetIpIs returns the IpIs field if non-nil, zero value otherwise.

### GetIpIsOk

`func (o *V2RuleProxy) GetIpIsOk() (*[]string, bool)`

GetIpIsOk returns a tuple with the IpIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIs

`func (o *V2RuleProxy) SetIpIs(v []string)`

SetIpIs sets IpIs field to given value.

### HasIpIs

`func (o *V2RuleProxy) HasIpIs() bool`

HasIpIs returns a boolean if a field has been set.

### GetIpIsNot

`func (o *V2RuleProxy) GetIpIsNot() []string`

GetIpIsNot returns the IpIsNot field if non-nil, zero value otherwise.

### GetIpIsNotOk

`func (o *V2RuleProxy) GetIpIsNotOk() (*[]string, bool)`

GetIpIsNotOk returns a tuple with the IpIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpIsNot

`func (o *V2RuleProxy) SetIpIsNot(v []string)`

SetIpIsNot sets IpIsNot field to given value.

### HasIpIsNot

`func (o *V2RuleProxy) HasIpIsNot() bool`

HasIpIsNot returns a boolean if a field has been set.

### GetCountry

`func (o *V2RuleProxy) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *V2RuleProxy) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *V2RuleProxy) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *V2RuleProxy) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryIs

`func (o *V2RuleProxy) GetCountryIs() []string`

GetCountryIs returns the CountryIs field if non-nil, zero value otherwise.

### GetCountryIsOk

`func (o *V2RuleProxy) GetCountryIsOk() (*[]string, bool)`

GetCountryIsOk returns a tuple with the CountryIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIs

`func (o *V2RuleProxy) SetCountryIs(v []string)`

SetCountryIs sets CountryIs field to given value.

### HasCountryIs

`func (o *V2RuleProxy) HasCountryIs() bool`

HasCountryIs returns a boolean if a field has been set.

### GetCountryIsNot

`func (o *V2RuleProxy) GetCountryIsNot() []string`

GetCountryIsNot returns the CountryIsNot field if non-nil, zero value otherwise.

### GetCountryIsNotOk

`func (o *V2RuleProxy) GetCountryIsNotOk() (*[]string, bool)`

GetCountryIsNotOk returns a tuple with the CountryIsNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsNot

`func (o *V2RuleProxy) SetCountryIsNot(v []string)`

SetCountryIsNot sets CountryIsNot field to given value.

### HasCountryIsNot

`func (o *V2RuleProxy) HasCountryIsNot() bool`

HasCountryIsNot returns a boolean if a field has been set.

### GetAction

`func (o *V2RuleProxy) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V2RuleProxy) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V2RuleProxy) SetAction(v string)`

SetAction sets Action field to given value.


### GetTo

`func (o *V2RuleProxy) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *V2RuleProxy) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *V2RuleProxy) SetTo(v string)`

SetTo sets To field to given value.


### GetHost

`func (o *V2RuleProxy) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V2RuleProxy) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V2RuleProxy) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V2RuleProxy) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAuthUser

`func (o *V2RuleProxy) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *V2RuleProxy) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *V2RuleProxy) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *V2RuleProxy) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetAuthPass

`func (o *V2RuleProxy) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *V2RuleProxy) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *V2RuleProxy) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.

### HasAuthPass

`func (o *V2RuleProxy) HasAuthPass() bool`

HasAuthPass returns a boolean if a field has been set.

### GetDisableSslVerify

`func (o *V2RuleProxy) GetDisableSslVerify() bool`

GetDisableSslVerify returns the DisableSslVerify field if non-nil, zero value otherwise.

### GetDisableSslVerifyOk

`func (o *V2RuleProxy) GetDisableSslVerifyOk() (*bool, bool)`

GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSslVerify

`func (o *V2RuleProxy) SetDisableSslVerify(v bool)`

SetDisableSslVerify sets DisableSslVerify field to given value.

### HasDisableSslVerify

`func (o *V2RuleProxy) HasDisableSslVerify() bool`

HasDisableSslVerify returns a boolean if a field has been set.

### GetCacheLifetime

`func (o *V2RuleProxy) GetCacheLifetime() string`

GetCacheLifetime returns the CacheLifetime field if non-nil, zero value otherwise.

### GetCacheLifetimeOk

`func (o *V2RuleProxy) GetCacheLifetimeOk() (*string, bool)`

GetCacheLifetimeOk returns a tuple with the CacheLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetime

`func (o *V2RuleProxy) SetCacheLifetime(v string)`

SetCacheLifetime sets CacheLifetime field to given value.

### HasCacheLifetime

`func (o *V2RuleProxy) HasCacheLifetime() bool`

HasCacheLifetime returns a boolean if a field has been set.

### SetCacheLifetimeNil

`func (o *V2RuleProxy) SetCacheLifetimeNil(b bool)`

 SetCacheLifetimeNil sets the value for CacheLifetime to be an explicit nil

### UnsetCacheLifetime
`func (o *V2RuleProxy) UnsetCacheLifetime()`

UnsetCacheLifetime ensures that no value is present for CacheLifetime, not even an explicit nil
### GetOnlyProxy404

`func (o *V2RuleProxy) GetOnlyProxy404() bool`

GetOnlyProxy404 returns the OnlyProxy404 field if non-nil, zero value otherwise.

### GetOnlyProxy404Ok

`func (o *V2RuleProxy) GetOnlyProxy404Ok() (*bool, bool)`

GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyProxy404

`func (o *V2RuleProxy) SetOnlyProxy404(v bool)`

SetOnlyProxy404 sets OnlyProxy404 field to given value.

### HasOnlyProxy404

`func (o *V2RuleProxy) HasOnlyProxy404() bool`

HasOnlyProxy404 returns a boolean if a field has been set.

### GetInjectHeaders

`func (o *V2RuleProxy) GetInjectHeaders() map[string]string`

GetInjectHeaders returns the InjectHeaders field if non-nil, zero value otherwise.

### GetInjectHeadersOk

`func (o *V2RuleProxy) GetInjectHeadersOk() (*map[string]string, bool)`

GetInjectHeadersOk returns a tuple with the InjectHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectHeaders

`func (o *V2RuleProxy) SetInjectHeaders(v map[string]string)`

SetInjectHeaders sets InjectHeaders field to given value.

### HasInjectHeaders

`func (o *V2RuleProxy) HasInjectHeaders() bool`

HasInjectHeaders returns a boolean if a field has been set.

### SetInjectHeadersNil

`func (o *V2RuleProxy) SetInjectHeadersNil(b bool)`

 SetInjectHeadersNil sets the value for InjectHeaders to be an explicit nil

### UnsetInjectHeaders
`func (o *V2RuleProxy) UnsetInjectHeaders()`

UnsetInjectHeaders ensures that no value is present for InjectHeaders, not even an explicit nil
### GetProxyStripHeaders

`func (o *V2RuleProxy) GetProxyStripHeaders() []string`

GetProxyStripHeaders returns the ProxyStripHeaders field if non-nil, zero value otherwise.

### GetProxyStripHeadersOk

`func (o *V2RuleProxy) GetProxyStripHeadersOk() (*[]string, bool)`

GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripHeaders

`func (o *V2RuleProxy) SetProxyStripHeaders(v []string)`

SetProxyStripHeaders sets ProxyStripHeaders field to given value.

### HasProxyStripHeaders

`func (o *V2RuleProxy) HasProxyStripHeaders() bool`

HasProxyStripHeaders returns a boolean if a field has been set.

### GetProxyStripRequestHeaders

`func (o *V2RuleProxy) GetProxyStripRequestHeaders() []string`

GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field if non-nil, zero value otherwise.

### GetProxyStripRequestHeadersOk

`func (o *V2RuleProxy) GetProxyStripRequestHeadersOk() (*[]string, bool)`

GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyStripRequestHeaders

`func (o *V2RuleProxy) SetProxyStripRequestHeaders(v []string)`

SetProxyStripRequestHeaders sets ProxyStripRequestHeaders field to given value.

### HasProxyStripRequestHeaders

`func (o *V2RuleProxy) HasProxyStripRequestHeaders() bool`

HasProxyStripRequestHeaders returns a boolean if a field has been set.

### GetOriginTimeout

`func (o *V2RuleProxy) GetOriginTimeout() string`

GetOriginTimeout returns the OriginTimeout field if non-nil, zero value otherwise.

### GetOriginTimeoutOk

`func (o *V2RuleProxy) GetOriginTimeoutOk() (*string, bool)`

GetOriginTimeoutOk returns a tuple with the OriginTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTimeout

`func (o *V2RuleProxy) SetOriginTimeout(v string)`

SetOriginTimeout sets OriginTimeout field to given value.

### HasOriginTimeout

`func (o *V2RuleProxy) HasOriginTimeout() bool`

HasOriginTimeout returns a boolean if a field has been set.

### GetFailoverMode

`func (o *V2RuleProxy) GetFailoverMode() bool`

GetFailoverMode returns the FailoverMode field if non-nil, zero value otherwise.

### GetFailoverModeOk

`func (o *V2RuleProxy) GetFailoverModeOk() (*bool, bool)`

GetFailoverModeOk returns a tuple with the FailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMode

`func (o *V2RuleProxy) SetFailoverMode(v bool)`

SetFailoverMode sets FailoverMode field to given value.

### HasFailoverMode

`func (o *V2RuleProxy) HasFailoverMode() bool`

HasFailoverMode returns a boolean if a field has been set.

### GetFailoverOriginTtfb

`func (o *V2RuleProxy) GetFailoverOriginTtfb() string`

GetFailoverOriginTtfb returns the FailoverOriginTtfb field if non-nil, zero value otherwise.

### GetFailoverOriginTtfbOk

`func (o *V2RuleProxy) GetFailoverOriginTtfbOk() (*string, bool)`

GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginTtfb

`func (o *V2RuleProxy) SetFailoverOriginTtfb(v string)`

SetFailoverOriginTtfb sets FailoverOriginTtfb field to given value.

### HasFailoverOriginTtfb

`func (o *V2RuleProxy) HasFailoverOriginTtfb() bool`

HasFailoverOriginTtfb returns a boolean if a field has been set.

### GetFailoverOriginStatusCodes

`func (o *V2RuleProxy) GetFailoverOriginStatusCodes() []string`

GetFailoverOriginStatusCodes returns the FailoverOriginStatusCodes field if non-nil, zero value otherwise.

### GetFailoverOriginStatusCodesOk

`func (o *V2RuleProxy) GetFailoverOriginStatusCodesOk() (*[]string, bool)`

GetFailoverOriginStatusCodesOk returns a tuple with the FailoverOriginStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverOriginStatusCodes

`func (o *V2RuleProxy) SetFailoverOriginStatusCodes(v []string)`

SetFailoverOriginStatusCodes sets FailoverOriginStatusCodes field to given value.

### HasFailoverOriginStatusCodes

`func (o *V2RuleProxy) HasFailoverOriginStatusCodes() bool`

HasFailoverOriginStatusCodes returns a boolean if a field has been set.

### GetFailoverLifetime

`func (o *V2RuleProxy) GetFailoverLifetime() string`

GetFailoverLifetime returns the FailoverLifetime field if non-nil, zero value otherwise.

### GetFailoverLifetimeOk

`func (o *V2RuleProxy) GetFailoverLifetimeOk() (*string, bool)`

GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverLifetime

`func (o *V2RuleProxy) SetFailoverLifetime(v string)`

SetFailoverLifetime sets FailoverLifetime field to given value.

### HasFailoverLifetime

`func (o *V2RuleProxy) HasFailoverLifetime() bool`

HasFailoverLifetime returns a boolean if a field has been set.

### GetNotify

`func (o *V2RuleProxy) GetNotify() string`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *V2RuleProxy) GetNotifyOk() (*string, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *V2RuleProxy) SetNotify(v string)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *V2RuleProxy) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNotifyConfig

`func (o *V2RuleProxy) GetNotifyConfig() V2RuleProxyActionAllOfNotifyConfig`

GetNotifyConfig returns the NotifyConfig field if non-nil, zero value otherwise.

### GetNotifyConfigOk

`func (o *V2RuleProxy) GetNotifyConfigOk() (*V2RuleProxyActionAllOfNotifyConfig, bool)`

GetNotifyConfigOk returns a tuple with the NotifyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyConfig

`func (o *V2RuleProxy) SetNotifyConfig(v V2RuleProxyActionAllOfNotifyConfig)`

SetNotifyConfig sets NotifyConfig field to given value.

### HasNotifyConfig

`func (o *V2RuleProxy) HasNotifyConfig() bool`

HasNotifyConfig returns a boolean if a field has been set.

### SetNotifyConfigNil

`func (o *V2RuleProxy) SetNotifyConfigNil(b bool)`

 SetNotifyConfigNil sets the value for NotifyConfig to be an explicit nil

### UnsetNotifyConfig
`func (o *V2RuleProxy) UnsetNotifyConfig()`

UnsetNotifyConfig ensures that no value is present for NotifyConfig, not even an explicit nil
### GetWafEnabled

`func (o *V2RuleProxy) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *V2RuleProxy) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *V2RuleProxy) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *V2RuleProxy) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.

### GetWafConfig

`func (o *V2RuleProxy) GetWafConfig() WafConfig`

GetWafConfig returns the WafConfig field if non-nil, zero value otherwise.

### GetWafConfigOk

`func (o *V2RuleProxy) GetWafConfigOk() (*WafConfig, bool)`

GetWafConfigOk returns a tuple with the WafConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafConfig

`func (o *V2RuleProxy) SetWafConfig(v WafConfig)`

SetWafConfig sets WafConfig field to given value.

### HasWafConfig

`func (o *V2RuleProxy) HasWafConfig() bool`

HasWafConfig returns a boolean if a field has been set.

### GetProxyAlertEnabled

`func (o *V2RuleProxy) GetProxyAlertEnabled() bool`

GetProxyAlertEnabled returns the ProxyAlertEnabled field if non-nil, zero value otherwise.

### GetProxyAlertEnabledOk

`func (o *V2RuleProxy) GetProxyAlertEnabledOk() (*bool, bool)`

GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAlertEnabled

`func (o *V2RuleProxy) SetProxyAlertEnabled(v bool)`

SetProxyAlertEnabled sets ProxyAlertEnabled field to given value.

### HasProxyAlertEnabled

`func (o *V2RuleProxy) HasProxyAlertEnabled() bool`

HasProxyAlertEnabled returns a boolean if a field has been set.

### GetProxyInlineFnEnabled

`func (o *V2RuleProxy) GetProxyInlineFnEnabled() bool`

GetProxyInlineFnEnabled returns the ProxyInlineFnEnabled field if non-nil, zero value otherwise.

### GetProxyInlineFnEnabledOk

`func (o *V2RuleProxy) GetProxyInlineFnEnabledOk() (*bool, bool)`

GetProxyInlineFnEnabledOk returns a tuple with the ProxyInlineFnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyInlineFnEnabled

`func (o *V2RuleProxy) SetProxyInlineFnEnabled(v bool)`

SetProxyInlineFnEnabled sets ProxyInlineFnEnabled field to given value.

### HasProxyInlineFnEnabled

`func (o *V2RuleProxy) HasProxyInlineFnEnabled() bool`

HasProxyInlineFnEnabled returns a boolean if a field has been set.

### GetQuantCloudSelection

`func (o *V2RuleProxy) GetQuantCloudSelection() V2RuleProxyActionAllOfQuantCloudSelection`

GetQuantCloudSelection returns the QuantCloudSelection field if non-nil, zero value otherwise.

### GetQuantCloudSelectionOk

`func (o *V2RuleProxy) GetQuantCloudSelectionOk() (*V2RuleProxyActionAllOfQuantCloudSelection, bool)`

GetQuantCloudSelectionOk returns a tuple with the QuantCloudSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantCloudSelection

`func (o *V2RuleProxy) SetQuantCloudSelection(v V2RuleProxyActionAllOfQuantCloudSelection)`

SetQuantCloudSelection sets QuantCloudSelection field to given value.

### HasQuantCloudSelection

`func (o *V2RuleProxy) HasQuantCloudSelection() bool`

HasQuantCloudSelection returns a boolean if a field has been set.

### SetQuantCloudSelectionNil

`func (o *V2RuleProxy) SetQuantCloudSelectionNil(b bool)`

 SetQuantCloudSelectionNil sets the value for QuantCloudSelection to be an explicit nil

### UnsetQuantCloudSelection
`func (o *V2RuleProxy) UnsetQuantCloudSelection()`

UnsetQuantCloudSelection ensures that no value is present for QuantCloudSelection, not even an explicit nil
### GetActionConfig

`func (o *V2RuleProxy) GetActionConfig() V2RuleProxyAction`

GetActionConfig returns the ActionConfig field if non-nil, zero value otherwise.

### GetActionConfigOk

`func (o *V2RuleProxy) GetActionConfigOk() (*V2RuleProxyAction, bool)`

GetActionConfigOk returns a tuple with the ActionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConfig

`func (o *V2RuleProxy) SetActionConfig(v V2RuleProxyAction)`

SetActionConfig sets ActionConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


