/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quantadmingo

import (
	"encoding/json"
)

// checks if the RuleProxyRequestUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleProxyRequestUpdate{}

// RuleProxyRequestUpdate struct for RuleProxyRequestUpdate
type RuleProxyRequestUpdate struct {
	Domain []string `json:"domain,omitempty"`
	Name *string `json:"name,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Url []string `json:"url,omitempty"`
	Country *string `json:"country,omitempty"`
	CountryIs []string `json:"country_is,omitempty"`
	CountryIsNot []string `json:"country_is_not,omitempty"`
	Method *string `json:"method,omitempty"`
	MethodIs []string `json:"method_is,omitempty"`
	MethodIsNot []string `json:"method_is_not,omitempty"`
	Ip *string `json:"ip,omitempty"`
	IpIs []string `json:"ip_is,omitempty"`
	IpIsNot []string `json:"ip_is_not,omitempty"`
	OnlyWithCookie *bool `json:"only_with_cookie,omitempty"`
	CookieName *string `json:"cookie_name,omitempty"`
	To *string `json:"to,omitempty"`
	Host *string `json:"host,omitempty"`
	AuthUser *string `json:"auth_user,omitempty"`
	AuthPass *string `json:"auth_pass,omitempty"`
	DisableSslVerify *bool `json:"disable_ssl_verify,omitempty"`
	CacheLifetime *int32 `json:"cache_lifetime,omitempty"`
	OnlyProxy404 *bool `json:"only_proxy_404,omitempty"`
	InjectHeaders *map[string]string `json:"inject_headers,omitempty"`
	ProxyStripHeaders []string `json:"proxy_strip_headers,omitempty"`
	ProxyStripRequestHeaders []string `json:"proxy_strip_request_headers,omitempty"`
	Failover *FailoverConfig `json:"failover,omitempty"`
	Notify *string `json:"notify,omitempty"`
	NotifyConfig *NotifyConfig `json:"notify_config,omitempty"`
	WafEnabled *bool `json:"waf_enabled,omitempty"`
	WafConfig *WAFConfigUpdate `json:"waf_config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RuleProxyRequestUpdate RuleProxyRequestUpdate

// NewRuleProxyRequestUpdate instantiates a new RuleProxyRequestUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleProxyRequestUpdate() *RuleProxyRequestUpdate {
	this := RuleProxyRequestUpdate{}
	var disabled bool = false
	this.Disabled = &disabled
	var onlyWithCookie bool = false
	this.OnlyWithCookie = &onlyWithCookie
	var notify string = "none"
	this.Notify = &notify
	var wafEnabled bool = false
	this.WafEnabled = &wafEnabled
	return &this
}

// NewRuleProxyRequestUpdateWithDefaults instantiates a new RuleProxyRequestUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleProxyRequestUpdateWithDefaults() *RuleProxyRequestUpdate {
	this := RuleProxyRequestUpdate{}
	var disabled bool = false
	this.Disabled = &disabled
	var onlyWithCookie bool = false
	this.OnlyWithCookie = &onlyWithCookie
	var notify string = "none"
	this.Notify = &notify
	var wafEnabled bool = false
	this.WafEnabled = &wafEnabled
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetDomain() []string {
	if o == nil || IsNil(o.Domain) {
		var ret []string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetDomainOk() ([]string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given []string and assigns it to the Domain field.
func (o *RuleProxyRequestUpdate) SetDomain(v []string) {
	o.Domain = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleProxyRequestUpdate) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RuleProxyRequestUpdate) SetUuid(v string) {
	o.Uuid = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *RuleProxyRequestUpdate) SetWeight(v int32) {
	o.Weight = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *RuleProxyRequestUpdate) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetUrl() []string {
	if o == nil || IsNil(o.Url) {
		var ret []string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetUrlOk() ([]string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given []string and assigns it to the Url field.
func (o *RuleProxyRequestUpdate) SetUrl(v []string) {
	o.Url = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *RuleProxyRequestUpdate) SetCountry(v string) {
	o.Country = &v
}

// GetCountryIs returns the CountryIs field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetCountryIs() []string {
	if o == nil || IsNil(o.CountryIs) {
		var ret []string
		return ret
	}
	return o.CountryIs
}

// GetCountryIsOk returns a tuple with the CountryIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetCountryIsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIs) {
		return nil, false
	}
	return o.CountryIs, true
}

// HasCountryIs returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasCountryIs() bool {
	if o != nil && !IsNil(o.CountryIs) {
		return true
	}

	return false
}

// SetCountryIs gets a reference to the given []string and assigns it to the CountryIs field.
func (o *RuleProxyRequestUpdate) SetCountryIs(v []string) {
	o.CountryIs = v
}

// GetCountryIsNot returns the CountryIsNot field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetCountryIsNot() []string {
	if o == nil || IsNil(o.CountryIsNot) {
		var ret []string
		return ret
	}
	return o.CountryIsNot
}

// GetCountryIsNotOk returns a tuple with the CountryIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetCountryIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIsNot) {
		return nil, false
	}
	return o.CountryIsNot, true
}

// HasCountryIsNot returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasCountryIsNot() bool {
	if o != nil && !IsNil(o.CountryIsNot) {
		return true
	}

	return false
}

// SetCountryIsNot gets a reference to the given []string and assigns it to the CountryIsNot field.
func (o *RuleProxyRequestUpdate) SetCountryIsNot(v []string) {
	o.CountryIsNot = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RuleProxyRequestUpdate) SetMethod(v string) {
	o.Method = &v
}

// GetMethodIs returns the MethodIs field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetMethodIs() []string {
	if o == nil || IsNil(o.MethodIs) {
		var ret []string
		return ret
	}
	return o.MethodIs
}

// GetMethodIsOk returns a tuple with the MethodIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetMethodIsOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIs) {
		return nil, false
	}
	return o.MethodIs, true
}

// HasMethodIs returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasMethodIs() bool {
	if o != nil && !IsNil(o.MethodIs) {
		return true
	}

	return false
}

// SetMethodIs gets a reference to the given []string and assigns it to the MethodIs field.
func (o *RuleProxyRequestUpdate) SetMethodIs(v []string) {
	o.MethodIs = v
}

// GetMethodIsNot returns the MethodIsNot field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetMethodIsNot() []string {
	if o == nil || IsNil(o.MethodIsNot) {
		var ret []string
		return ret
	}
	return o.MethodIsNot
}

// GetMethodIsNotOk returns a tuple with the MethodIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetMethodIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIsNot) {
		return nil, false
	}
	return o.MethodIsNot, true
}

// HasMethodIsNot returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasMethodIsNot() bool {
	if o != nil && !IsNil(o.MethodIsNot) {
		return true
	}

	return false
}

// SetMethodIsNot gets a reference to the given []string and assigns it to the MethodIsNot field.
func (o *RuleProxyRequestUpdate) SetMethodIsNot(v []string) {
	o.MethodIsNot = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RuleProxyRequestUpdate) SetIp(v string) {
	o.Ip = &v
}

// GetIpIs returns the IpIs field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetIpIs() []string {
	if o == nil || IsNil(o.IpIs) {
		var ret []string
		return ret
	}
	return o.IpIs
}

// GetIpIsOk returns a tuple with the IpIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetIpIsOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIs) {
		return nil, false
	}
	return o.IpIs, true
}

// HasIpIs returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasIpIs() bool {
	if o != nil && !IsNil(o.IpIs) {
		return true
	}

	return false
}

// SetIpIs gets a reference to the given []string and assigns it to the IpIs field.
func (o *RuleProxyRequestUpdate) SetIpIs(v []string) {
	o.IpIs = v
}

// GetIpIsNot returns the IpIsNot field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetIpIsNot() []string {
	if o == nil || IsNil(o.IpIsNot) {
		var ret []string
		return ret
	}
	return o.IpIsNot
}

// GetIpIsNotOk returns a tuple with the IpIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetIpIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIsNot) {
		return nil, false
	}
	return o.IpIsNot, true
}

// HasIpIsNot returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasIpIsNot() bool {
	if o != nil && !IsNil(o.IpIsNot) {
		return true
	}

	return false
}

// SetIpIsNot gets a reference to the given []string and assigns it to the IpIsNot field.
func (o *RuleProxyRequestUpdate) SetIpIsNot(v []string) {
	o.IpIsNot = v
}

// GetOnlyWithCookie returns the OnlyWithCookie field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetOnlyWithCookie() bool {
	if o == nil || IsNil(o.OnlyWithCookie) {
		var ret bool
		return ret
	}
	return *o.OnlyWithCookie
}

// GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetOnlyWithCookieOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyWithCookie) {
		return nil, false
	}
	return o.OnlyWithCookie, true
}

// HasOnlyWithCookie returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasOnlyWithCookie() bool {
	if o != nil && !IsNil(o.OnlyWithCookie) {
		return true
	}

	return false
}

// SetOnlyWithCookie gets a reference to the given bool and assigns it to the OnlyWithCookie field.
func (o *RuleProxyRequestUpdate) SetOnlyWithCookie(v bool) {
	o.OnlyWithCookie = &v
}

// GetCookieName returns the CookieName field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetCookieName() string {
	if o == nil || IsNil(o.CookieName) {
		var ret string
		return ret
	}
	return *o.CookieName
}

// GetCookieNameOk returns a tuple with the CookieName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetCookieNameOk() (*string, bool) {
	if o == nil || IsNil(o.CookieName) {
		return nil, false
	}
	return o.CookieName, true
}

// HasCookieName returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasCookieName() bool {
	if o != nil && !IsNil(o.CookieName) {
		return true
	}

	return false
}

// SetCookieName gets a reference to the given string and assigns it to the CookieName field.
func (o *RuleProxyRequestUpdate) SetCookieName(v string) {
	o.CookieName = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *RuleProxyRequestUpdate) SetTo(v string) {
	o.To = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *RuleProxyRequestUpdate) SetHost(v string) {
	o.Host = &v
}

// GetAuthUser returns the AuthUser field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetAuthUser() string {
	if o == nil || IsNil(o.AuthUser) {
		var ret string
		return ret
	}
	return *o.AuthUser
}

// GetAuthUserOk returns a tuple with the AuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetAuthUserOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUser) {
		return nil, false
	}
	return o.AuthUser, true
}

// HasAuthUser returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasAuthUser() bool {
	if o != nil && !IsNil(o.AuthUser) {
		return true
	}

	return false
}

// SetAuthUser gets a reference to the given string and assigns it to the AuthUser field.
func (o *RuleProxyRequestUpdate) SetAuthUser(v string) {
	o.AuthUser = &v
}

// GetAuthPass returns the AuthPass field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetAuthPass() string {
	if o == nil || IsNil(o.AuthPass) {
		var ret string
		return ret
	}
	return *o.AuthPass
}

// GetAuthPassOk returns a tuple with the AuthPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetAuthPassOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPass) {
		return nil, false
	}
	return o.AuthPass, true
}

// HasAuthPass returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasAuthPass() bool {
	if o != nil && !IsNil(o.AuthPass) {
		return true
	}

	return false
}

// SetAuthPass gets a reference to the given string and assigns it to the AuthPass field.
func (o *RuleProxyRequestUpdate) SetAuthPass(v string) {
	o.AuthPass = &v
}

// GetDisableSslVerify returns the DisableSslVerify field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetDisableSslVerify() bool {
	if o == nil || IsNil(o.DisableSslVerify) {
		var ret bool
		return ret
	}
	return *o.DisableSslVerify
}

// GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetDisableSslVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableSslVerify) {
		return nil, false
	}
	return o.DisableSslVerify, true
}

// HasDisableSslVerify returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasDisableSslVerify() bool {
	if o != nil && !IsNil(o.DisableSslVerify) {
		return true
	}

	return false
}

// SetDisableSslVerify gets a reference to the given bool and assigns it to the DisableSslVerify field.
func (o *RuleProxyRequestUpdate) SetDisableSslVerify(v bool) {
	o.DisableSslVerify = &v
}

// GetCacheLifetime returns the CacheLifetime field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetCacheLifetime() int32 {
	if o == nil || IsNil(o.CacheLifetime) {
		var ret int32
		return ret
	}
	return *o.CacheLifetime
}

// GetCacheLifetimeOk returns a tuple with the CacheLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetCacheLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheLifetime) {
		return nil, false
	}
	return o.CacheLifetime, true
}

// HasCacheLifetime returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasCacheLifetime() bool {
	if o != nil && !IsNil(o.CacheLifetime) {
		return true
	}

	return false
}

// SetCacheLifetime gets a reference to the given int32 and assigns it to the CacheLifetime field.
func (o *RuleProxyRequestUpdate) SetCacheLifetime(v int32) {
	o.CacheLifetime = &v
}

// GetOnlyProxy404 returns the OnlyProxy404 field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetOnlyProxy404() bool {
	if o == nil || IsNil(o.OnlyProxy404) {
		var ret bool
		return ret
	}
	return *o.OnlyProxy404
}

// GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetOnlyProxy404Ok() (*bool, bool) {
	if o == nil || IsNil(o.OnlyProxy404) {
		return nil, false
	}
	return o.OnlyProxy404, true
}

// HasOnlyProxy404 returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasOnlyProxy404() bool {
	if o != nil && !IsNil(o.OnlyProxy404) {
		return true
	}

	return false
}

// SetOnlyProxy404 gets a reference to the given bool and assigns it to the OnlyProxy404 field.
func (o *RuleProxyRequestUpdate) SetOnlyProxy404(v bool) {
	o.OnlyProxy404 = &v
}

// GetInjectHeaders returns the InjectHeaders field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetInjectHeaders() map[string]string {
	if o == nil || IsNil(o.InjectHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.InjectHeaders
}

// GetInjectHeadersOk returns a tuple with the InjectHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetInjectHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.InjectHeaders) {
		return nil, false
	}
	return o.InjectHeaders, true
}

// HasInjectHeaders returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasInjectHeaders() bool {
	if o != nil && !IsNil(o.InjectHeaders) {
		return true
	}

	return false
}

// SetInjectHeaders gets a reference to the given map[string]string and assigns it to the InjectHeaders field.
func (o *RuleProxyRequestUpdate) SetInjectHeaders(v map[string]string) {
	o.InjectHeaders = &v
}

// GetProxyStripHeaders returns the ProxyStripHeaders field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetProxyStripHeaders() []string {
	if o == nil || IsNil(o.ProxyStripHeaders) {
		var ret []string
		return ret
	}
	return o.ProxyStripHeaders
}

// GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetProxyStripHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.ProxyStripHeaders) {
		return nil, false
	}
	return o.ProxyStripHeaders, true
}

// HasProxyStripHeaders returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasProxyStripHeaders() bool {
	if o != nil && !IsNil(o.ProxyStripHeaders) {
		return true
	}

	return false
}

// SetProxyStripHeaders gets a reference to the given []string and assigns it to the ProxyStripHeaders field.
func (o *RuleProxyRequestUpdate) SetProxyStripHeaders(v []string) {
	o.ProxyStripHeaders = v
}

// GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetProxyStripRequestHeaders() []string {
	if o == nil || IsNil(o.ProxyStripRequestHeaders) {
		var ret []string
		return ret
	}
	return o.ProxyStripRequestHeaders
}

// GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetProxyStripRequestHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.ProxyStripRequestHeaders) {
		return nil, false
	}
	return o.ProxyStripRequestHeaders, true
}

// HasProxyStripRequestHeaders returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasProxyStripRequestHeaders() bool {
	if o != nil && !IsNil(o.ProxyStripRequestHeaders) {
		return true
	}

	return false
}

// SetProxyStripRequestHeaders gets a reference to the given []string and assigns it to the ProxyStripRequestHeaders field.
func (o *RuleProxyRequestUpdate) SetProxyStripRequestHeaders(v []string) {
	o.ProxyStripRequestHeaders = v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetFailover() FailoverConfig {
	if o == nil || IsNil(o.Failover) {
		var ret FailoverConfig
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetFailoverOk() (*FailoverConfig, bool) {
	if o == nil || IsNil(o.Failover) {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasFailover() bool {
	if o != nil && !IsNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given FailoverConfig and assigns it to the Failover field.
func (o *RuleProxyRequestUpdate) SetFailover(v FailoverConfig) {
	o.Failover = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetNotify() string {
	if o == nil || IsNil(o.Notify) {
		var ret string
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetNotifyOk() (*string, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given string and assigns it to the Notify field.
func (o *RuleProxyRequestUpdate) SetNotify(v string) {
	o.Notify = &v
}

// GetNotifyConfig returns the NotifyConfig field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetNotifyConfig() NotifyConfig {
	if o == nil || IsNil(o.NotifyConfig) {
		var ret NotifyConfig
		return ret
	}
	return *o.NotifyConfig
}

// GetNotifyConfigOk returns a tuple with the NotifyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetNotifyConfigOk() (*NotifyConfig, bool) {
	if o == nil || IsNil(o.NotifyConfig) {
		return nil, false
	}
	return o.NotifyConfig, true
}

// HasNotifyConfig returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasNotifyConfig() bool {
	if o != nil && !IsNil(o.NotifyConfig) {
		return true
	}

	return false
}

// SetNotifyConfig gets a reference to the given NotifyConfig and assigns it to the NotifyConfig field.
func (o *RuleProxyRequestUpdate) SetNotifyConfig(v NotifyConfig) {
	o.NotifyConfig = &v
}

// GetWafEnabled returns the WafEnabled field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetWafEnabled() bool {
	if o == nil || IsNil(o.WafEnabled) {
		var ret bool
		return ret
	}
	return *o.WafEnabled
}

// GetWafEnabledOk returns a tuple with the WafEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetWafEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WafEnabled) {
		return nil, false
	}
	return o.WafEnabled, true
}

// HasWafEnabled returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasWafEnabled() bool {
	if o != nil && !IsNil(o.WafEnabled) {
		return true
	}

	return false
}

// SetWafEnabled gets a reference to the given bool and assigns it to the WafEnabled field.
func (o *RuleProxyRequestUpdate) SetWafEnabled(v bool) {
	o.WafEnabled = &v
}

// GetWafConfig returns the WafConfig field value if set, zero value otherwise.
func (o *RuleProxyRequestUpdate) GetWafConfig() WAFConfigUpdate {
	if o == nil || IsNil(o.WafConfig) {
		var ret WAFConfigUpdate
		return ret
	}
	return *o.WafConfig
}

// GetWafConfigOk returns a tuple with the WafConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequestUpdate) GetWafConfigOk() (*WAFConfigUpdate, bool) {
	if o == nil || IsNil(o.WafConfig) {
		return nil, false
	}
	return o.WafConfig, true
}

// HasWafConfig returns a boolean if a field has been set.
func (o *RuleProxyRequestUpdate) HasWafConfig() bool {
	if o != nil && !IsNil(o.WafConfig) {
		return true
	}

	return false
}

// SetWafConfig gets a reference to the given WAFConfigUpdate and assigns it to the WafConfig field.
func (o *RuleProxyRequestUpdate) SetWafConfig(v WAFConfigUpdate) {
	o.WafConfig = &v
}

func (o RuleProxyRequestUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleProxyRequestUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CountryIs) {
		toSerialize["country_is"] = o.CountryIs
	}
	if !IsNil(o.CountryIsNot) {
		toSerialize["country_is_not"] = o.CountryIsNot
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.MethodIs) {
		toSerialize["method_is"] = o.MethodIs
	}
	if !IsNil(o.MethodIsNot) {
		toSerialize["method_is_not"] = o.MethodIsNot
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.IpIs) {
		toSerialize["ip_is"] = o.IpIs
	}
	if !IsNil(o.IpIsNot) {
		toSerialize["ip_is_not"] = o.IpIsNot
	}
	if !IsNil(o.OnlyWithCookie) {
		toSerialize["only_with_cookie"] = o.OnlyWithCookie
	}
	if !IsNil(o.CookieName) {
		toSerialize["cookie_name"] = o.CookieName
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.AuthUser) {
		toSerialize["auth_user"] = o.AuthUser
	}
	if !IsNil(o.AuthPass) {
		toSerialize["auth_pass"] = o.AuthPass
	}
	if !IsNil(o.DisableSslVerify) {
		toSerialize["disable_ssl_verify"] = o.DisableSslVerify
	}
	if !IsNil(o.CacheLifetime) {
		toSerialize["cache_lifetime"] = o.CacheLifetime
	}
	if !IsNil(o.OnlyProxy404) {
		toSerialize["only_proxy_404"] = o.OnlyProxy404
	}
	if !IsNil(o.InjectHeaders) {
		toSerialize["inject_headers"] = o.InjectHeaders
	}
	if !IsNil(o.ProxyStripHeaders) {
		toSerialize["proxy_strip_headers"] = o.ProxyStripHeaders
	}
	if !IsNil(o.ProxyStripRequestHeaders) {
		toSerialize["proxy_strip_request_headers"] = o.ProxyStripRequestHeaders
	}
	if !IsNil(o.Failover) {
		toSerialize["failover"] = o.Failover
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.NotifyConfig) {
		toSerialize["notify_config"] = o.NotifyConfig
	}
	if !IsNil(o.WafEnabled) {
		toSerialize["waf_enabled"] = o.WafEnabled
	}
	if !IsNil(o.WafConfig) {
		toSerialize["waf_config"] = o.WafConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleProxyRequestUpdate) UnmarshalJSON(data []byte) (err error) {
	varRuleProxyRequestUpdate := _RuleProxyRequestUpdate{}

	err = json.Unmarshal(data, &varRuleProxyRequestUpdate)

	if err != nil {
		return err
	}

	*o = RuleProxyRequestUpdate(varRuleProxyRequestUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "name")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "url")
		delete(additionalProperties, "country")
		delete(additionalProperties, "country_is")
		delete(additionalProperties, "country_is_not")
		delete(additionalProperties, "method")
		delete(additionalProperties, "method_is")
		delete(additionalProperties, "method_is_not")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "ip_is")
		delete(additionalProperties, "ip_is_not")
		delete(additionalProperties, "only_with_cookie")
		delete(additionalProperties, "cookie_name")
		delete(additionalProperties, "to")
		delete(additionalProperties, "host")
		delete(additionalProperties, "auth_user")
		delete(additionalProperties, "auth_pass")
		delete(additionalProperties, "disable_ssl_verify")
		delete(additionalProperties, "cache_lifetime")
		delete(additionalProperties, "only_proxy_404")
		delete(additionalProperties, "inject_headers")
		delete(additionalProperties, "proxy_strip_headers")
		delete(additionalProperties, "proxy_strip_request_headers")
		delete(additionalProperties, "failover")
		delete(additionalProperties, "notify")
		delete(additionalProperties, "notify_config")
		delete(additionalProperties, "waf_enabled")
		delete(additionalProperties, "waf_config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleProxyRequestUpdate struct {
	value *RuleProxyRequestUpdate
	isSet bool
}

func (v NullableRuleProxyRequestUpdate) Get() *RuleProxyRequestUpdate {
	return v.value
}

func (v *NullableRuleProxyRequestUpdate) Set(val *RuleProxyRequestUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleProxyRequestUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleProxyRequestUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleProxyRequestUpdate(val *RuleProxyRequestUpdate) *NullableRuleProxyRequestUpdate {
	return &NullableRuleProxyRequestUpdate{value: val, isSet: true}
}

func (v NullableRuleProxyRequestUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleProxyRequestUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


