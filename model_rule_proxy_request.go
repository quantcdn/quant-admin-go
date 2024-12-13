/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the RuleProxyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleProxyRequest{}

// RuleProxyRequest struct for RuleProxyRequest
type RuleProxyRequest struct {
	Domain []string `json:"domain"`
	Name *string `json:"name,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Disabled bool `json:"disabled"`
	Url []string `json:"url"`
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
	Proxy ProxyConfig `json:"proxy"`
	Failover *FailoverConfig `json:"failover,omitempty"`
	Notify *string `json:"notify,omitempty"`
	NotifyConfig *NotifyConfig `json:"notify_config,omitempty"`
	WafEnabled *bool `json:"waf_enabled,omitempty"`
	WafConfig *WAFConfig `json:"waf_config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RuleProxyRequest RuleProxyRequest

// NewRuleProxyRequest instantiates a new RuleProxyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleProxyRequest(domain []string, disabled bool, url []string, proxy ProxyConfig) *RuleProxyRequest {
	this := RuleProxyRequest{}
	this.Domain = domain
	this.Disabled = disabled
	this.Url = url
	var onlyWithCookie bool = false
	this.OnlyWithCookie = &onlyWithCookie
	this.Proxy = proxy
	var notify string = "none"
	this.Notify = &notify
	var wafEnabled bool = false
	this.WafEnabled = &wafEnabled
	return &this
}

// NewRuleProxyRequestWithDefaults instantiates a new RuleProxyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleProxyRequestWithDefaults() *RuleProxyRequest {
	this := RuleProxyRequest{}
	var disabled bool = false
	this.Disabled = disabled
	var onlyWithCookie bool = false
	this.OnlyWithCookie = &onlyWithCookie
	var notify string = "none"
	this.Notify = &notify
	var wafEnabled bool = false
	this.WafEnabled = &wafEnabled
	return &this
}

// GetDomain returns the Domain field value
func (o *RuleProxyRequest) GetDomain() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetDomainOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain, true
}

// SetDomain sets field value
func (o *RuleProxyRequest) SetDomain(v []string) {
	o.Domain = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleProxyRequest) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RuleProxyRequest) SetUuid(v string) {
	o.Uuid = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *RuleProxyRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetDisabled returns the Disabled field value
func (o *RuleProxyRequest) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *RuleProxyRequest) SetDisabled(v bool) {
	o.Disabled = v
}

// GetUrl returns the Url field value
func (o *RuleProxyRequest) GetUrl() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetUrlOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url, true
}

// SetUrl sets field value
func (o *RuleProxyRequest) SetUrl(v []string) {
	o.Url = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *RuleProxyRequest) SetCountry(v string) {
	o.Country = &v
}

// GetCountryIs returns the CountryIs field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetCountryIs() []string {
	if o == nil || IsNil(o.CountryIs) {
		var ret []string
		return ret
	}
	return o.CountryIs
}

// GetCountryIsOk returns a tuple with the CountryIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetCountryIsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIs) {
		return nil, false
	}
	return o.CountryIs, true
}

// HasCountryIs returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasCountryIs() bool {
	if o != nil && !IsNil(o.CountryIs) {
		return true
	}

	return false
}

// SetCountryIs gets a reference to the given []string and assigns it to the CountryIs field.
func (o *RuleProxyRequest) SetCountryIs(v []string) {
	o.CountryIs = v
}

// GetCountryIsNot returns the CountryIsNot field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetCountryIsNot() []string {
	if o == nil || IsNil(o.CountryIsNot) {
		var ret []string
		return ret
	}
	return o.CountryIsNot
}

// GetCountryIsNotOk returns a tuple with the CountryIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetCountryIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIsNot) {
		return nil, false
	}
	return o.CountryIsNot, true
}

// HasCountryIsNot returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasCountryIsNot() bool {
	if o != nil && !IsNil(o.CountryIsNot) {
		return true
	}

	return false
}

// SetCountryIsNot gets a reference to the given []string and assigns it to the CountryIsNot field.
func (o *RuleProxyRequest) SetCountryIsNot(v []string) {
	o.CountryIsNot = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RuleProxyRequest) SetMethod(v string) {
	o.Method = &v
}

// GetMethodIs returns the MethodIs field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetMethodIs() []string {
	if o == nil || IsNil(o.MethodIs) {
		var ret []string
		return ret
	}
	return o.MethodIs
}

// GetMethodIsOk returns a tuple with the MethodIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetMethodIsOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIs) {
		return nil, false
	}
	return o.MethodIs, true
}

// HasMethodIs returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasMethodIs() bool {
	if o != nil && !IsNil(o.MethodIs) {
		return true
	}

	return false
}

// SetMethodIs gets a reference to the given []string and assigns it to the MethodIs field.
func (o *RuleProxyRequest) SetMethodIs(v []string) {
	o.MethodIs = v
}

// GetMethodIsNot returns the MethodIsNot field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetMethodIsNot() []string {
	if o == nil || IsNil(o.MethodIsNot) {
		var ret []string
		return ret
	}
	return o.MethodIsNot
}

// GetMethodIsNotOk returns a tuple with the MethodIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetMethodIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIsNot) {
		return nil, false
	}
	return o.MethodIsNot, true
}

// HasMethodIsNot returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasMethodIsNot() bool {
	if o != nil && !IsNil(o.MethodIsNot) {
		return true
	}

	return false
}

// SetMethodIsNot gets a reference to the given []string and assigns it to the MethodIsNot field.
func (o *RuleProxyRequest) SetMethodIsNot(v []string) {
	o.MethodIsNot = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RuleProxyRequest) SetIp(v string) {
	o.Ip = &v
}

// GetIpIs returns the IpIs field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetIpIs() []string {
	if o == nil || IsNil(o.IpIs) {
		var ret []string
		return ret
	}
	return o.IpIs
}

// GetIpIsOk returns a tuple with the IpIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetIpIsOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIs) {
		return nil, false
	}
	return o.IpIs, true
}

// HasIpIs returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasIpIs() bool {
	if o != nil && !IsNil(o.IpIs) {
		return true
	}

	return false
}

// SetIpIs gets a reference to the given []string and assigns it to the IpIs field.
func (o *RuleProxyRequest) SetIpIs(v []string) {
	o.IpIs = v
}

// GetIpIsNot returns the IpIsNot field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetIpIsNot() []string {
	if o == nil || IsNil(o.IpIsNot) {
		var ret []string
		return ret
	}
	return o.IpIsNot
}

// GetIpIsNotOk returns a tuple with the IpIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetIpIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIsNot) {
		return nil, false
	}
	return o.IpIsNot, true
}

// HasIpIsNot returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasIpIsNot() bool {
	if o != nil && !IsNil(o.IpIsNot) {
		return true
	}

	return false
}

// SetIpIsNot gets a reference to the given []string and assigns it to the IpIsNot field.
func (o *RuleProxyRequest) SetIpIsNot(v []string) {
	o.IpIsNot = v
}

// GetOnlyWithCookie returns the OnlyWithCookie field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetOnlyWithCookie() bool {
	if o == nil || IsNil(o.OnlyWithCookie) {
		var ret bool
		return ret
	}
	return *o.OnlyWithCookie
}

// GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetOnlyWithCookieOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyWithCookie) {
		return nil, false
	}
	return o.OnlyWithCookie, true
}

// HasOnlyWithCookie returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasOnlyWithCookie() bool {
	if o != nil && !IsNil(o.OnlyWithCookie) {
		return true
	}

	return false
}

// SetOnlyWithCookie gets a reference to the given bool and assigns it to the OnlyWithCookie field.
func (o *RuleProxyRequest) SetOnlyWithCookie(v bool) {
	o.OnlyWithCookie = &v
}

// GetCookieName returns the CookieName field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetCookieName() string {
	if o == nil || IsNil(o.CookieName) {
		var ret string
		return ret
	}
	return *o.CookieName
}

// GetCookieNameOk returns a tuple with the CookieName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetCookieNameOk() (*string, bool) {
	if o == nil || IsNil(o.CookieName) {
		return nil, false
	}
	return o.CookieName, true
}

// HasCookieName returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasCookieName() bool {
	if o != nil && !IsNil(o.CookieName) {
		return true
	}

	return false
}

// SetCookieName gets a reference to the given string and assigns it to the CookieName field.
func (o *RuleProxyRequest) SetCookieName(v string) {
	o.CookieName = &v
}

// GetProxy returns the Proxy field value
func (o *RuleProxyRequest) GetProxy() ProxyConfig {
	if o == nil {
		var ret ProxyConfig
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetProxyOk() (*ProxyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *RuleProxyRequest) SetProxy(v ProxyConfig) {
	o.Proxy = v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetFailover() FailoverConfig {
	if o == nil || IsNil(o.Failover) {
		var ret FailoverConfig
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetFailoverOk() (*FailoverConfig, bool) {
	if o == nil || IsNil(o.Failover) {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasFailover() bool {
	if o != nil && !IsNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given FailoverConfig and assigns it to the Failover field.
func (o *RuleProxyRequest) SetFailover(v FailoverConfig) {
	o.Failover = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetNotify() string {
	if o == nil || IsNil(o.Notify) {
		var ret string
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetNotifyOk() (*string, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given string and assigns it to the Notify field.
func (o *RuleProxyRequest) SetNotify(v string) {
	o.Notify = &v
}

// GetNotifyConfig returns the NotifyConfig field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetNotifyConfig() NotifyConfig {
	if o == nil || IsNil(o.NotifyConfig) {
		var ret NotifyConfig
		return ret
	}
	return *o.NotifyConfig
}

// GetNotifyConfigOk returns a tuple with the NotifyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetNotifyConfigOk() (*NotifyConfig, bool) {
	if o == nil || IsNil(o.NotifyConfig) {
		return nil, false
	}
	return o.NotifyConfig, true
}

// HasNotifyConfig returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasNotifyConfig() bool {
	if o != nil && !IsNil(o.NotifyConfig) {
		return true
	}

	return false
}

// SetNotifyConfig gets a reference to the given NotifyConfig and assigns it to the NotifyConfig field.
func (o *RuleProxyRequest) SetNotifyConfig(v NotifyConfig) {
	o.NotifyConfig = &v
}

// GetWafEnabled returns the WafEnabled field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetWafEnabled() bool {
	if o == nil || IsNil(o.WafEnabled) {
		var ret bool
		return ret
	}
	return *o.WafEnabled
}

// GetWafEnabledOk returns a tuple with the WafEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetWafEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WafEnabled) {
		return nil, false
	}
	return o.WafEnabled, true
}

// HasWafEnabled returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasWafEnabled() bool {
	if o != nil && !IsNil(o.WafEnabled) {
		return true
	}

	return false
}

// SetWafEnabled gets a reference to the given bool and assigns it to the WafEnabled field.
func (o *RuleProxyRequest) SetWafEnabled(v bool) {
	o.WafEnabled = &v
}

// GetWafConfig returns the WafConfig field value if set, zero value otherwise.
func (o *RuleProxyRequest) GetWafConfig() WAFConfig {
	if o == nil || IsNil(o.WafConfig) {
		var ret WAFConfig
		return ret
	}
	return *o.WafConfig
}

// GetWafConfigOk returns a tuple with the WafConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyRequest) GetWafConfigOk() (*WAFConfig, bool) {
	if o == nil || IsNil(o.WafConfig) {
		return nil, false
	}
	return o.WafConfig, true
}

// HasWafConfig returns a boolean if a field has been set.
func (o *RuleProxyRequest) HasWafConfig() bool {
	if o != nil && !IsNil(o.WafConfig) {
		return true
	}

	return false
}

// SetWafConfig gets a reference to the given WAFConfig and assigns it to the WafConfig field.
func (o *RuleProxyRequest) SetWafConfig(v WAFConfig) {
	o.WafConfig = &v
}

func (o RuleProxyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleProxyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	toSerialize["disabled"] = o.Disabled
	toSerialize["url"] = o.Url
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
	toSerialize["proxy"] = o.Proxy
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

func (o *RuleProxyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"disabled",
		"url",
		"proxy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRuleProxyRequest := _RuleProxyRequest{}

	err = json.Unmarshal(data, &varRuleProxyRequest)

	if err != nil {
		return err
	}

	*o = RuleProxyRequest(varRuleProxyRequest)

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
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "failover")
		delete(additionalProperties, "notify")
		delete(additionalProperties, "notify_config")
		delete(additionalProperties, "waf_enabled")
		delete(additionalProperties, "waf_config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleProxyRequest struct {
	value *RuleProxyRequest
	isSet bool
}

func (v NullableRuleProxyRequest) Get() *RuleProxyRequest {
	return v.value
}

func (v *NullableRuleProxyRequest) Set(val *RuleProxyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleProxyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleProxyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleProxyRequest(val *RuleProxyRequest) *NullableRuleProxyRequest {
	return &NullableRuleProxyRequest{value: val, isSet: true}
}

func (v NullableRuleProxyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleProxyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


