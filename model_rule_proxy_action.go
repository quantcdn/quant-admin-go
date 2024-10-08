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

// checks if the RuleProxyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleProxyAction{}

// RuleProxyAction struct for RuleProxyAction
type RuleProxyAction struct {
	Host *string `json:"host,omitempty"`
	WafEnabled bool `json:"waf_enabled"`
	ProxyStripHeaders []string `json:"proxy_strip_headers,omitempty"`
	AuthUser *string `json:"auth_user,omitempty"`
	AuthPass *string `json:"auth_pass,omitempty"`
	OriginTimeout *string `json:"origin_timeout,omitempty"`
	ProxyAlertEnabled *bool `json:"proxy_alert_enabled,omitempty"`
	Notify *string `json:"notify,omitempty"`
	NotifyConfig *RuleProxyActionNotifyConfig `json:"notify_config,omitempty"`
	ProxyStripRequestHeaders []string `json:"proxy_strip_request_headers,omitempty"`
	FailoverOriginStatusCodes []string `json:"failover_origin_status_codes,omitempty"`
	FailoverOriginTtfb *string `json:"failover_origin_ttfb,omitempty"`
	FailoverMode *bool `json:"failover_mode,omitempty"`
	FailoverLifetime *string `json:"failover_lifetime,omitempty"`
	OnlyProxy404 *bool `json:"only_proxy_404,omitempty"`
	InjectHeaders *map[string]string `json:"inject_headers,omitempty"`
	To string `json:"to"`
	CacheLifetime *int32 `json:"cache_lifetime,omitempty"`
	DisableSslVerify *bool `json:"disable_ssl_verify,omitempty"`
	NotifyEmail *string `json:"notify_email,omitempty"`
	WafConfig *WAFConfig `json:"waf_config,omitempty"`
	ProxyInlineFnEnabled *bool `json:"proxy_inline_fn_enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RuleProxyAction RuleProxyAction

// NewRuleProxyAction instantiates a new RuleProxyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleProxyAction(wafEnabled bool, to string) *RuleProxyAction {
	this := RuleProxyAction{}
	this.WafEnabled = wafEnabled
	this.To = to
	var cacheLifetime int32 = 0
	this.CacheLifetime = &cacheLifetime
	var disableSslVerify bool = true
	this.DisableSslVerify = &disableSslVerify
	var proxyInlineFnEnabled bool = false
	this.ProxyInlineFnEnabled = &proxyInlineFnEnabled
	return &this
}

// NewRuleProxyActionWithDefaults instantiates a new RuleProxyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleProxyActionWithDefaults() *RuleProxyAction {
	this := RuleProxyAction{}
	var wafEnabled bool = false
	this.WafEnabled = wafEnabled
	var cacheLifetime int32 = 0
	this.CacheLifetime = &cacheLifetime
	var disableSslVerify bool = true
	this.DisableSslVerify = &disableSslVerify
	var proxyInlineFnEnabled bool = false
	this.ProxyInlineFnEnabled = &proxyInlineFnEnabled
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *RuleProxyAction) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *RuleProxyAction) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *RuleProxyAction) SetHost(v string) {
	o.Host = &v
}

// GetWafEnabled returns the WafEnabled field value
func (o *RuleProxyAction) GetWafEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WafEnabled
}

// GetWafEnabledOk returns a tuple with the WafEnabled field value
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetWafEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WafEnabled, true
}

// SetWafEnabled sets field value
func (o *RuleProxyAction) SetWafEnabled(v bool) {
	o.WafEnabled = v
}

// GetProxyStripHeaders returns the ProxyStripHeaders field value if set, zero value otherwise.
func (o *RuleProxyAction) GetProxyStripHeaders() []string {
	if o == nil || IsNil(o.ProxyStripHeaders) {
		var ret []string
		return ret
	}
	return o.ProxyStripHeaders
}

// GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetProxyStripHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.ProxyStripHeaders) {
		return nil, false
	}
	return o.ProxyStripHeaders, true
}

// HasProxyStripHeaders returns a boolean if a field has been set.
func (o *RuleProxyAction) HasProxyStripHeaders() bool {
	if o != nil && !IsNil(o.ProxyStripHeaders) {
		return true
	}

	return false
}

// SetProxyStripHeaders gets a reference to the given []string and assigns it to the ProxyStripHeaders field.
func (o *RuleProxyAction) SetProxyStripHeaders(v []string) {
	o.ProxyStripHeaders = v
}

// GetAuthUser returns the AuthUser field value if set, zero value otherwise.
func (o *RuleProxyAction) GetAuthUser() string {
	if o == nil || IsNil(o.AuthUser) {
		var ret string
		return ret
	}
	return *o.AuthUser
}

// GetAuthUserOk returns a tuple with the AuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetAuthUserOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUser) {
		return nil, false
	}
	return o.AuthUser, true
}

// HasAuthUser returns a boolean if a field has been set.
func (o *RuleProxyAction) HasAuthUser() bool {
	if o != nil && !IsNil(o.AuthUser) {
		return true
	}

	return false
}

// SetAuthUser gets a reference to the given string and assigns it to the AuthUser field.
func (o *RuleProxyAction) SetAuthUser(v string) {
	o.AuthUser = &v
}

// GetAuthPass returns the AuthPass field value if set, zero value otherwise.
func (o *RuleProxyAction) GetAuthPass() string {
	if o == nil || IsNil(o.AuthPass) {
		var ret string
		return ret
	}
	return *o.AuthPass
}

// GetAuthPassOk returns a tuple with the AuthPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetAuthPassOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPass) {
		return nil, false
	}
	return o.AuthPass, true
}

// HasAuthPass returns a boolean if a field has been set.
func (o *RuleProxyAction) HasAuthPass() bool {
	if o != nil && !IsNil(o.AuthPass) {
		return true
	}

	return false
}

// SetAuthPass gets a reference to the given string and assigns it to the AuthPass field.
func (o *RuleProxyAction) SetAuthPass(v string) {
	o.AuthPass = &v
}

// GetOriginTimeout returns the OriginTimeout field value if set, zero value otherwise.
func (o *RuleProxyAction) GetOriginTimeout() string {
	if o == nil || IsNil(o.OriginTimeout) {
		var ret string
		return ret
	}
	return *o.OriginTimeout
}

// GetOriginTimeoutOk returns a tuple with the OriginTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetOriginTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.OriginTimeout) {
		return nil, false
	}
	return o.OriginTimeout, true
}

// HasOriginTimeout returns a boolean if a field has been set.
func (o *RuleProxyAction) HasOriginTimeout() bool {
	if o != nil && !IsNil(o.OriginTimeout) {
		return true
	}

	return false
}

// SetOriginTimeout gets a reference to the given string and assigns it to the OriginTimeout field.
func (o *RuleProxyAction) SetOriginTimeout(v string) {
	o.OriginTimeout = &v
}

// GetProxyAlertEnabled returns the ProxyAlertEnabled field value if set, zero value otherwise.
func (o *RuleProxyAction) GetProxyAlertEnabled() bool {
	if o == nil || IsNil(o.ProxyAlertEnabled) {
		var ret bool
		return ret
	}
	return *o.ProxyAlertEnabled
}

// GetProxyAlertEnabledOk returns a tuple with the ProxyAlertEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetProxyAlertEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProxyAlertEnabled) {
		return nil, false
	}
	return o.ProxyAlertEnabled, true
}

// HasProxyAlertEnabled returns a boolean if a field has been set.
func (o *RuleProxyAction) HasProxyAlertEnabled() bool {
	if o != nil && !IsNil(o.ProxyAlertEnabled) {
		return true
	}

	return false
}

// SetProxyAlertEnabled gets a reference to the given bool and assigns it to the ProxyAlertEnabled field.
func (o *RuleProxyAction) SetProxyAlertEnabled(v bool) {
	o.ProxyAlertEnabled = &v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *RuleProxyAction) GetNotify() string {
	if o == nil || IsNil(o.Notify) {
		var ret string
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetNotifyOk() (*string, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *RuleProxyAction) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given string and assigns it to the Notify field.
func (o *RuleProxyAction) SetNotify(v string) {
	o.Notify = &v
}

// GetNotifyConfig returns the NotifyConfig field value if set, zero value otherwise.
func (o *RuleProxyAction) GetNotifyConfig() RuleProxyActionNotifyConfig {
	if o == nil || IsNil(o.NotifyConfig) {
		var ret RuleProxyActionNotifyConfig
		return ret
	}
	return *o.NotifyConfig
}

// GetNotifyConfigOk returns a tuple with the NotifyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetNotifyConfigOk() (*RuleProxyActionNotifyConfig, bool) {
	if o == nil || IsNil(o.NotifyConfig) {
		return nil, false
	}
	return o.NotifyConfig, true
}

// HasNotifyConfig returns a boolean if a field has been set.
func (o *RuleProxyAction) HasNotifyConfig() bool {
	if o != nil && !IsNil(o.NotifyConfig) {
		return true
	}

	return false
}

// SetNotifyConfig gets a reference to the given RuleProxyActionNotifyConfig and assigns it to the NotifyConfig field.
func (o *RuleProxyAction) SetNotifyConfig(v RuleProxyActionNotifyConfig) {
	o.NotifyConfig = &v
}

// GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field value if set, zero value otherwise.
func (o *RuleProxyAction) GetProxyStripRequestHeaders() []string {
	if o == nil || IsNil(o.ProxyStripRequestHeaders) {
		var ret []string
		return ret
	}
	return o.ProxyStripRequestHeaders
}

// GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetProxyStripRequestHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.ProxyStripRequestHeaders) {
		return nil, false
	}
	return o.ProxyStripRequestHeaders, true
}

// HasProxyStripRequestHeaders returns a boolean if a field has been set.
func (o *RuleProxyAction) HasProxyStripRequestHeaders() bool {
	if o != nil && !IsNil(o.ProxyStripRequestHeaders) {
		return true
	}

	return false
}

// SetProxyStripRequestHeaders gets a reference to the given []string and assigns it to the ProxyStripRequestHeaders field.
func (o *RuleProxyAction) SetProxyStripRequestHeaders(v []string) {
	o.ProxyStripRequestHeaders = v
}

// GetFailoverOriginStatusCodes returns the FailoverOriginStatusCodes field value if set, zero value otherwise.
func (o *RuleProxyAction) GetFailoverOriginStatusCodes() []string {
	if o == nil || IsNil(o.FailoverOriginStatusCodes) {
		var ret []string
		return ret
	}
	return o.FailoverOriginStatusCodes
}

// GetFailoverOriginStatusCodesOk returns a tuple with the FailoverOriginStatusCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetFailoverOriginStatusCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.FailoverOriginStatusCodes) {
		return nil, false
	}
	return o.FailoverOriginStatusCodes, true
}

// HasFailoverOriginStatusCodes returns a boolean if a field has been set.
func (o *RuleProxyAction) HasFailoverOriginStatusCodes() bool {
	if o != nil && !IsNil(o.FailoverOriginStatusCodes) {
		return true
	}

	return false
}

// SetFailoverOriginStatusCodes gets a reference to the given []string and assigns it to the FailoverOriginStatusCodes field.
func (o *RuleProxyAction) SetFailoverOriginStatusCodes(v []string) {
	o.FailoverOriginStatusCodes = v
}

// GetFailoverOriginTtfb returns the FailoverOriginTtfb field value if set, zero value otherwise.
func (o *RuleProxyAction) GetFailoverOriginTtfb() string {
	if o == nil || IsNil(o.FailoverOriginTtfb) {
		var ret string
		return ret
	}
	return *o.FailoverOriginTtfb
}

// GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetFailoverOriginTtfbOk() (*string, bool) {
	if o == nil || IsNil(o.FailoverOriginTtfb) {
		return nil, false
	}
	return o.FailoverOriginTtfb, true
}

// HasFailoverOriginTtfb returns a boolean if a field has been set.
func (o *RuleProxyAction) HasFailoverOriginTtfb() bool {
	if o != nil && !IsNil(o.FailoverOriginTtfb) {
		return true
	}

	return false
}

// SetFailoverOriginTtfb gets a reference to the given string and assigns it to the FailoverOriginTtfb field.
func (o *RuleProxyAction) SetFailoverOriginTtfb(v string) {
	o.FailoverOriginTtfb = &v
}

// GetFailoverMode returns the FailoverMode field value if set, zero value otherwise.
func (o *RuleProxyAction) GetFailoverMode() bool {
	if o == nil || IsNil(o.FailoverMode) {
		var ret bool
		return ret
	}
	return *o.FailoverMode
}

// GetFailoverModeOk returns a tuple with the FailoverMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetFailoverModeOk() (*bool, bool) {
	if o == nil || IsNil(o.FailoverMode) {
		return nil, false
	}
	return o.FailoverMode, true
}

// HasFailoverMode returns a boolean if a field has been set.
func (o *RuleProxyAction) HasFailoverMode() bool {
	if o != nil && !IsNil(o.FailoverMode) {
		return true
	}

	return false
}

// SetFailoverMode gets a reference to the given bool and assigns it to the FailoverMode field.
func (o *RuleProxyAction) SetFailoverMode(v bool) {
	o.FailoverMode = &v
}

// GetFailoverLifetime returns the FailoverLifetime field value if set, zero value otherwise.
func (o *RuleProxyAction) GetFailoverLifetime() string {
	if o == nil || IsNil(o.FailoverLifetime) {
		var ret string
		return ret
	}
	return *o.FailoverLifetime
}

// GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetFailoverLifetimeOk() (*string, bool) {
	if o == nil || IsNil(o.FailoverLifetime) {
		return nil, false
	}
	return o.FailoverLifetime, true
}

// HasFailoverLifetime returns a boolean if a field has been set.
func (o *RuleProxyAction) HasFailoverLifetime() bool {
	if o != nil && !IsNil(o.FailoverLifetime) {
		return true
	}

	return false
}

// SetFailoverLifetime gets a reference to the given string and assigns it to the FailoverLifetime field.
func (o *RuleProxyAction) SetFailoverLifetime(v string) {
	o.FailoverLifetime = &v
}

// GetOnlyProxy404 returns the OnlyProxy404 field value if set, zero value otherwise.
func (o *RuleProxyAction) GetOnlyProxy404() bool {
	if o == nil || IsNil(o.OnlyProxy404) {
		var ret bool
		return ret
	}
	return *o.OnlyProxy404
}

// GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetOnlyProxy404Ok() (*bool, bool) {
	if o == nil || IsNil(o.OnlyProxy404) {
		return nil, false
	}
	return o.OnlyProxy404, true
}

// HasOnlyProxy404 returns a boolean if a field has been set.
func (o *RuleProxyAction) HasOnlyProxy404() bool {
	if o != nil && !IsNil(o.OnlyProxy404) {
		return true
	}

	return false
}

// SetOnlyProxy404 gets a reference to the given bool and assigns it to the OnlyProxy404 field.
func (o *RuleProxyAction) SetOnlyProxy404(v bool) {
	o.OnlyProxy404 = &v
}

// GetInjectHeaders returns the InjectHeaders field value if set, zero value otherwise.
func (o *RuleProxyAction) GetInjectHeaders() map[string]string {
	if o == nil || IsNil(o.InjectHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.InjectHeaders
}

// GetInjectHeadersOk returns a tuple with the InjectHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetInjectHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.InjectHeaders) {
		return nil, false
	}
	return o.InjectHeaders, true
}

// HasInjectHeaders returns a boolean if a field has been set.
func (o *RuleProxyAction) HasInjectHeaders() bool {
	if o != nil && !IsNil(o.InjectHeaders) {
		return true
	}

	return false
}

// SetInjectHeaders gets a reference to the given map[string]string and assigns it to the InjectHeaders field.
func (o *RuleProxyAction) SetInjectHeaders(v map[string]string) {
	o.InjectHeaders = &v
}

// GetTo returns the To field value
func (o *RuleProxyAction) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *RuleProxyAction) SetTo(v string) {
	o.To = v
}

// GetCacheLifetime returns the CacheLifetime field value if set, zero value otherwise.
func (o *RuleProxyAction) GetCacheLifetime() int32 {
	if o == nil || IsNil(o.CacheLifetime) {
		var ret int32
		return ret
	}
	return *o.CacheLifetime
}

// GetCacheLifetimeOk returns a tuple with the CacheLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetCacheLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheLifetime) {
		return nil, false
	}
	return o.CacheLifetime, true
}

// HasCacheLifetime returns a boolean if a field has been set.
func (o *RuleProxyAction) HasCacheLifetime() bool {
	if o != nil && !IsNil(o.CacheLifetime) {
		return true
	}

	return false
}

// SetCacheLifetime gets a reference to the given int32 and assigns it to the CacheLifetime field.
func (o *RuleProxyAction) SetCacheLifetime(v int32) {
	o.CacheLifetime = &v
}

// GetDisableSslVerify returns the DisableSslVerify field value if set, zero value otherwise.
func (o *RuleProxyAction) GetDisableSslVerify() bool {
	if o == nil || IsNil(o.DisableSslVerify) {
		var ret bool
		return ret
	}
	return *o.DisableSslVerify
}

// GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetDisableSslVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableSslVerify) {
		return nil, false
	}
	return o.DisableSslVerify, true
}

// HasDisableSslVerify returns a boolean if a field has been set.
func (o *RuleProxyAction) HasDisableSslVerify() bool {
	if o != nil && !IsNil(o.DisableSslVerify) {
		return true
	}

	return false
}

// SetDisableSslVerify gets a reference to the given bool and assigns it to the DisableSslVerify field.
func (o *RuleProxyAction) SetDisableSslVerify(v bool) {
	o.DisableSslVerify = &v
}

// GetNotifyEmail returns the NotifyEmail field value if set, zero value otherwise.
func (o *RuleProxyAction) GetNotifyEmail() string {
	if o == nil || IsNil(o.NotifyEmail) {
		var ret string
		return ret
	}
	return *o.NotifyEmail
}

// GetNotifyEmailOk returns a tuple with the NotifyEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetNotifyEmailOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyEmail) {
		return nil, false
	}
	return o.NotifyEmail, true
}

// HasNotifyEmail returns a boolean if a field has been set.
func (o *RuleProxyAction) HasNotifyEmail() bool {
	if o != nil && !IsNil(o.NotifyEmail) {
		return true
	}

	return false
}

// SetNotifyEmail gets a reference to the given string and assigns it to the NotifyEmail field.
func (o *RuleProxyAction) SetNotifyEmail(v string) {
	o.NotifyEmail = &v
}

// GetWafConfig returns the WafConfig field value if set, zero value otherwise.
func (o *RuleProxyAction) GetWafConfig() WAFConfig {
	if o == nil || IsNil(o.WafConfig) {
		var ret WAFConfig
		return ret
	}
	return *o.WafConfig
}

// GetWafConfigOk returns a tuple with the WafConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetWafConfigOk() (*WAFConfig, bool) {
	if o == nil || IsNil(o.WafConfig) {
		return nil, false
	}
	return o.WafConfig, true
}

// HasWafConfig returns a boolean if a field has been set.
func (o *RuleProxyAction) HasWafConfig() bool {
	if o != nil && !IsNil(o.WafConfig) {
		return true
	}

	return false
}

// SetWafConfig gets a reference to the given WAFConfig and assigns it to the WafConfig field.
func (o *RuleProxyAction) SetWafConfig(v WAFConfig) {
	o.WafConfig = &v
}

// GetProxyInlineFnEnabled returns the ProxyInlineFnEnabled field value if set, zero value otherwise.
func (o *RuleProxyAction) GetProxyInlineFnEnabled() bool {
	if o == nil || IsNil(o.ProxyInlineFnEnabled) {
		var ret bool
		return ret
	}
	return *o.ProxyInlineFnEnabled
}

// GetProxyInlineFnEnabledOk returns a tuple with the ProxyInlineFnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetProxyInlineFnEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProxyInlineFnEnabled) {
		return nil, false
	}
	return o.ProxyInlineFnEnabled, true
}

// HasProxyInlineFnEnabled returns a boolean if a field has been set.
func (o *RuleProxyAction) HasProxyInlineFnEnabled() bool {
	if o != nil && !IsNil(o.ProxyInlineFnEnabled) {
		return true
	}

	return false
}

// SetProxyInlineFnEnabled gets a reference to the given bool and assigns it to the ProxyInlineFnEnabled field.
func (o *RuleProxyAction) SetProxyInlineFnEnabled(v bool) {
	o.ProxyInlineFnEnabled = &v
}

func (o RuleProxyAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleProxyAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	toSerialize["waf_enabled"] = o.WafEnabled
	if !IsNil(o.ProxyStripHeaders) {
		toSerialize["proxy_strip_headers"] = o.ProxyStripHeaders
	}
	if !IsNil(o.AuthUser) {
		toSerialize["auth_user"] = o.AuthUser
	}
	if !IsNil(o.AuthPass) {
		toSerialize["auth_pass"] = o.AuthPass
	}
	if !IsNil(o.OriginTimeout) {
		toSerialize["origin_timeout"] = o.OriginTimeout
	}
	if !IsNil(o.ProxyAlertEnabled) {
		toSerialize["proxy_alert_enabled"] = o.ProxyAlertEnabled
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.NotifyConfig) {
		toSerialize["notify_config"] = o.NotifyConfig
	}
	if !IsNil(o.ProxyStripRequestHeaders) {
		toSerialize["proxy_strip_request_headers"] = o.ProxyStripRequestHeaders
	}
	if !IsNil(o.FailoverOriginStatusCodes) {
		toSerialize["failover_origin_status_codes"] = o.FailoverOriginStatusCodes
	}
	if !IsNil(o.FailoverOriginTtfb) {
		toSerialize["failover_origin_ttfb"] = o.FailoverOriginTtfb
	}
	if !IsNil(o.FailoverMode) {
		toSerialize["failover_mode"] = o.FailoverMode
	}
	if !IsNil(o.FailoverLifetime) {
		toSerialize["failover_lifetime"] = o.FailoverLifetime
	}
	if !IsNil(o.OnlyProxy404) {
		toSerialize["only_proxy_404"] = o.OnlyProxy404
	}
	if !IsNil(o.InjectHeaders) {
		toSerialize["inject_headers"] = o.InjectHeaders
	}
	toSerialize["to"] = o.To
	if !IsNil(o.CacheLifetime) {
		toSerialize["cache_lifetime"] = o.CacheLifetime
	}
	if !IsNil(o.DisableSslVerify) {
		toSerialize["disable_ssl_verify"] = o.DisableSslVerify
	}
	if !IsNil(o.NotifyEmail) {
		toSerialize["notify_email"] = o.NotifyEmail
	}
	if !IsNil(o.WafConfig) {
		toSerialize["waf_config"] = o.WafConfig
	}
	if !IsNil(o.ProxyInlineFnEnabled) {
		toSerialize["proxy_inline_fn_enabled"] = o.ProxyInlineFnEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleProxyAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"waf_enabled",
		"to",
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

	varRuleProxyAction := _RuleProxyAction{}

	err = json.Unmarshal(data, &varRuleProxyAction)

	if err != nil {
		return err
	}

	*o = RuleProxyAction(varRuleProxyAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "host")
		delete(additionalProperties, "waf_enabled")
		delete(additionalProperties, "proxy_strip_headers")
		delete(additionalProperties, "auth_user")
		delete(additionalProperties, "auth_pass")
		delete(additionalProperties, "origin_timeout")
		delete(additionalProperties, "proxy_alert_enabled")
		delete(additionalProperties, "notify")
		delete(additionalProperties, "notify_config")
		delete(additionalProperties, "proxy_strip_request_headers")
		delete(additionalProperties, "failover_origin_status_codes")
		delete(additionalProperties, "failover_origin_ttfb")
		delete(additionalProperties, "failover_mode")
		delete(additionalProperties, "failover_lifetime")
		delete(additionalProperties, "only_proxy_404")
		delete(additionalProperties, "inject_headers")
		delete(additionalProperties, "to")
		delete(additionalProperties, "cache_lifetime")
		delete(additionalProperties, "disable_ssl_verify")
		delete(additionalProperties, "notify_email")
		delete(additionalProperties, "waf_config")
		delete(additionalProperties, "proxy_inline_fn_enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleProxyAction struct {
	value *RuleProxyAction
	isSet bool
}

func (v NullableRuleProxyAction) Get() *RuleProxyAction {
	return v.value
}

func (v *NullableRuleProxyAction) Set(val *RuleProxyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleProxyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleProxyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleProxyAction(val *RuleProxyAction) *NullableRuleProxyAction {
	return &NullableRuleProxyAction{value: val, isSet: true}
}

func (v NullableRuleProxyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleProxyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


