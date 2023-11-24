/*
Quant administration API

The Quant administration API provides programmatic access to manage projects within your available organizations. 

API version: 2.0.0
Contact: apiteam@quantcdn.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RuleAuthRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleAuthRequest{}

// RuleAuthRequest struct for RuleAuthRequest
type RuleAuthRequest struct {
	Domain *string `json:"domain,omitempty"`
	Country *string `json:"country,omitempty"`
	CountryIs []string `json:"country_is,omitempty"`
	CountryIsNot []string `json:"country_is_not,omitempty"`
	Method *string `json:"method,omitempty"`
	MethodIs []string `json:"method_is,omitempty"`
	MethodIsNot []string `json:"method_is_not,omitempty"`
	Ip *string `json:"ip,omitempty"`
	IpIs []string `json:"ip_is,omitempty"`
	IpIsNot []string `json:"ip_is_not,omitempty"`
	OnlyWithCookie *string `json:"only_with_cookie,omitempty"`
	Url *string `json:"url,omitempty"`
	Name *string `json:"name,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	To *string `json:"to,omitempty"`
	Host *string `json:"host,omitempty"`
	AuthUser *string `json:"auth_user,omitempty"`
	AuthPass *string `json:"auth_pass,omitempty"`
	DisableSslVerify *bool `json:"disable_ssl_verify,omitempty"`
	CacheLifetime *int32 `json:"cache_lifetime,omitempty"`
	OnlyProxy404 *bool `json:"only_proxy_404,omitempty"`
	StripHeaders []string `json:"strip_headers,omitempty"`
}

// NewRuleAuthRequest instantiates a new RuleAuthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleAuthRequest() *RuleAuthRequest {
	this := RuleAuthRequest{}
	return &this
}

// NewRuleAuthRequestWithDefaults instantiates a new RuleAuthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleAuthRequestWithDefaults() *RuleAuthRequest {
	this := RuleAuthRequest{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *RuleAuthRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *RuleAuthRequest) SetCountry(v string) {
	o.Country = &v
}

// GetCountryIs returns the CountryIs field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetCountryIs() []string {
	if o == nil || IsNil(o.CountryIs) {
		var ret []string
		return ret
	}
	return o.CountryIs
}

// GetCountryIsOk returns a tuple with the CountryIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetCountryIsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIs) {
		return nil, false
	}
	return o.CountryIs, true
}

// HasCountryIs returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasCountryIs() bool {
	if o != nil && !IsNil(o.CountryIs) {
		return true
	}

	return false
}

// SetCountryIs gets a reference to the given []string and assigns it to the CountryIs field.
func (o *RuleAuthRequest) SetCountryIs(v []string) {
	o.CountryIs = v
}

// GetCountryIsNot returns the CountryIsNot field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetCountryIsNot() []string {
	if o == nil || IsNil(o.CountryIsNot) {
		var ret []string
		return ret
	}
	return o.CountryIsNot
}

// GetCountryIsNotOk returns a tuple with the CountryIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetCountryIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIsNot) {
		return nil, false
	}
	return o.CountryIsNot, true
}

// HasCountryIsNot returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasCountryIsNot() bool {
	if o != nil && !IsNil(o.CountryIsNot) {
		return true
	}

	return false
}

// SetCountryIsNot gets a reference to the given []string and assigns it to the CountryIsNot field.
func (o *RuleAuthRequest) SetCountryIsNot(v []string) {
	o.CountryIsNot = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RuleAuthRequest) SetMethod(v string) {
	o.Method = &v
}

// GetMethodIs returns the MethodIs field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetMethodIs() []string {
	if o == nil || IsNil(o.MethodIs) {
		var ret []string
		return ret
	}
	return o.MethodIs
}

// GetMethodIsOk returns a tuple with the MethodIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetMethodIsOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIs) {
		return nil, false
	}
	return o.MethodIs, true
}

// HasMethodIs returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasMethodIs() bool {
	if o != nil && !IsNil(o.MethodIs) {
		return true
	}

	return false
}

// SetMethodIs gets a reference to the given []string and assigns it to the MethodIs field.
func (o *RuleAuthRequest) SetMethodIs(v []string) {
	o.MethodIs = v
}

// GetMethodIsNot returns the MethodIsNot field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetMethodIsNot() []string {
	if o == nil || IsNil(o.MethodIsNot) {
		var ret []string
		return ret
	}
	return o.MethodIsNot
}

// GetMethodIsNotOk returns a tuple with the MethodIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetMethodIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIsNot) {
		return nil, false
	}
	return o.MethodIsNot, true
}

// HasMethodIsNot returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasMethodIsNot() bool {
	if o != nil && !IsNil(o.MethodIsNot) {
		return true
	}

	return false
}

// SetMethodIsNot gets a reference to the given []string and assigns it to the MethodIsNot field.
func (o *RuleAuthRequest) SetMethodIsNot(v []string) {
	o.MethodIsNot = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RuleAuthRequest) SetIp(v string) {
	o.Ip = &v
}

// GetIpIs returns the IpIs field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetIpIs() []string {
	if o == nil || IsNil(o.IpIs) {
		var ret []string
		return ret
	}
	return o.IpIs
}

// GetIpIsOk returns a tuple with the IpIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetIpIsOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIs) {
		return nil, false
	}
	return o.IpIs, true
}

// HasIpIs returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasIpIs() bool {
	if o != nil && !IsNil(o.IpIs) {
		return true
	}

	return false
}

// SetIpIs gets a reference to the given []string and assigns it to the IpIs field.
func (o *RuleAuthRequest) SetIpIs(v []string) {
	o.IpIs = v
}

// GetIpIsNot returns the IpIsNot field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetIpIsNot() []string {
	if o == nil || IsNil(o.IpIsNot) {
		var ret []string
		return ret
	}
	return o.IpIsNot
}

// GetIpIsNotOk returns a tuple with the IpIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetIpIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIsNot) {
		return nil, false
	}
	return o.IpIsNot, true
}

// HasIpIsNot returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasIpIsNot() bool {
	if o != nil && !IsNil(o.IpIsNot) {
		return true
	}

	return false
}

// SetIpIsNot gets a reference to the given []string and assigns it to the IpIsNot field.
func (o *RuleAuthRequest) SetIpIsNot(v []string) {
	o.IpIsNot = v
}

// GetOnlyWithCookie returns the OnlyWithCookie field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetOnlyWithCookie() string {
	if o == nil || IsNil(o.OnlyWithCookie) {
		var ret string
		return ret
	}
	return *o.OnlyWithCookie
}

// GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetOnlyWithCookieOk() (*string, bool) {
	if o == nil || IsNil(o.OnlyWithCookie) {
		return nil, false
	}
	return o.OnlyWithCookie, true
}

// HasOnlyWithCookie returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasOnlyWithCookie() bool {
	if o != nil && !IsNil(o.OnlyWithCookie) {
		return true
	}

	return false
}

// SetOnlyWithCookie gets a reference to the given string and assigns it to the OnlyWithCookie field.
func (o *RuleAuthRequest) SetOnlyWithCookie(v string) {
	o.OnlyWithCookie = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RuleAuthRequest) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleAuthRequest) SetName(v string) {
	o.Name = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *RuleAuthRequest) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *RuleAuthRequest) SetTo(v string) {
	o.To = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *RuleAuthRequest) SetHost(v string) {
	o.Host = &v
}

// GetAuthUser returns the AuthUser field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetAuthUser() string {
	if o == nil || IsNil(o.AuthUser) {
		var ret string
		return ret
	}
	return *o.AuthUser
}

// GetAuthUserOk returns a tuple with the AuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetAuthUserOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUser) {
		return nil, false
	}
	return o.AuthUser, true
}

// HasAuthUser returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasAuthUser() bool {
	if o != nil && !IsNil(o.AuthUser) {
		return true
	}

	return false
}

// SetAuthUser gets a reference to the given string and assigns it to the AuthUser field.
func (o *RuleAuthRequest) SetAuthUser(v string) {
	o.AuthUser = &v
}

// GetAuthPass returns the AuthPass field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetAuthPass() string {
	if o == nil || IsNil(o.AuthPass) {
		var ret string
		return ret
	}
	return *o.AuthPass
}

// GetAuthPassOk returns a tuple with the AuthPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetAuthPassOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPass) {
		return nil, false
	}
	return o.AuthPass, true
}

// HasAuthPass returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasAuthPass() bool {
	if o != nil && !IsNil(o.AuthPass) {
		return true
	}

	return false
}

// SetAuthPass gets a reference to the given string and assigns it to the AuthPass field.
func (o *RuleAuthRequest) SetAuthPass(v string) {
	o.AuthPass = &v
}

// GetDisableSslVerify returns the DisableSslVerify field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetDisableSslVerify() bool {
	if o == nil || IsNil(o.DisableSslVerify) {
		var ret bool
		return ret
	}
	return *o.DisableSslVerify
}

// GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetDisableSslVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableSslVerify) {
		return nil, false
	}
	return o.DisableSslVerify, true
}

// HasDisableSslVerify returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasDisableSslVerify() bool {
	if o != nil && !IsNil(o.DisableSslVerify) {
		return true
	}

	return false
}

// SetDisableSslVerify gets a reference to the given bool and assigns it to the DisableSslVerify field.
func (o *RuleAuthRequest) SetDisableSslVerify(v bool) {
	o.DisableSslVerify = &v
}

// GetCacheLifetime returns the CacheLifetime field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetCacheLifetime() int32 {
	if o == nil || IsNil(o.CacheLifetime) {
		var ret int32
		return ret
	}
	return *o.CacheLifetime
}

// GetCacheLifetimeOk returns a tuple with the CacheLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetCacheLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheLifetime) {
		return nil, false
	}
	return o.CacheLifetime, true
}

// HasCacheLifetime returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasCacheLifetime() bool {
	if o != nil && !IsNil(o.CacheLifetime) {
		return true
	}

	return false
}

// SetCacheLifetime gets a reference to the given int32 and assigns it to the CacheLifetime field.
func (o *RuleAuthRequest) SetCacheLifetime(v int32) {
	o.CacheLifetime = &v
}

// GetOnlyProxy404 returns the OnlyProxy404 field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetOnlyProxy404() bool {
	if o == nil || IsNil(o.OnlyProxy404) {
		var ret bool
		return ret
	}
	return *o.OnlyProxy404
}

// GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetOnlyProxy404Ok() (*bool, bool) {
	if o == nil || IsNil(o.OnlyProxy404) {
		return nil, false
	}
	return o.OnlyProxy404, true
}

// HasOnlyProxy404 returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasOnlyProxy404() bool {
	if o != nil && !IsNil(o.OnlyProxy404) {
		return true
	}

	return false
}

// SetOnlyProxy404 gets a reference to the given bool and assigns it to the OnlyProxy404 field.
func (o *RuleAuthRequest) SetOnlyProxy404(v bool) {
	o.OnlyProxy404 = &v
}

// GetStripHeaders returns the StripHeaders field value if set, zero value otherwise.
func (o *RuleAuthRequest) GetStripHeaders() []string {
	if o == nil || IsNil(o.StripHeaders) {
		var ret []string
		return ret
	}
	return o.StripHeaders
}

// GetStripHeadersOk returns a tuple with the StripHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAuthRequest) GetStripHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.StripHeaders) {
		return nil, false
	}
	return o.StripHeaders, true
}

// HasStripHeaders returns a boolean if a field has been set.
func (o *RuleAuthRequest) HasStripHeaders() bool {
	if o != nil && !IsNil(o.StripHeaders) {
		return true
	}

	return false
}

// SetStripHeaders gets a reference to the given []string and assigns it to the StripHeaders field.
func (o *RuleAuthRequest) SetStripHeaders(v []string) {
	o.StripHeaders = v
}

func (o RuleAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleAuthRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
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
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
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
	if !IsNil(o.StripHeaders) {
		toSerialize["strip_headers"] = o.StripHeaders
	}
	return toSerialize, nil
}

type NullableRuleAuthRequest struct {
	value *RuleAuthRequest
	isSet bool
}

func (v NullableRuleAuthRequest) Get() *RuleAuthRequest {
	return v.value
}

func (v *NullableRuleAuthRequest) Set(val *RuleAuthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleAuthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleAuthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleAuthRequest(val *RuleAuthRequest) *NullableRuleAuthRequest {
	return &NullableRuleAuthRequest{value: val, isSet: true}
}

func (v NullableRuleAuthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleAuthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


