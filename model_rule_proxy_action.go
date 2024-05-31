/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RuleProxyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleProxyAction{}

// RuleProxyAction struct for RuleProxyAction
type RuleProxyAction struct {
	Host *string `json:"host,omitempty"`
	WafEnabled bool `json:"waf_enabled"`
	ProxyStripHeaders []string `json:"proxy_strip_headers,omitempty"`
	AuthPass *string `json:"auth_pass,omitempty"`
	OriginTimeout *int32 `json:"origin_timeout,omitempty"`
	ProxyAlertEnabled *bool `json:"proxy_alert_enabled,omitempty"`
	Notify *string `json:"notify,omitempty"`
	FailoverOriginStatusCode []string `json:"failover_origin_status_code,omitempty"`
	FailoverOriginTtfb *int32 `json:"failover_origin_ttfb,omitempty"`
	FailoverLifetime *int32 `json:"failover_lifetime,omitempty"`
	OnlyProxy404 *bool `json:"only_proxy_404,omitempty"`
	InjectHeaders *map[string]string `json:"inject_headers,omitempty"`
	WafConfig *RuleWAFConfig `json:"waf_config,omitempty"`
	To *string `json:"to,omitempty"`
	CacheLifetime *int32 `json:"cache_lifetime,omitempty"`
	FailoverMode *bool `json:"failover_mode,omitempty"`
	DisableSslVerify *bool `json:"disable_ssl_verify,omitempty"`
	NotifyConfig ProxyNotifyConfig `json:"notify_config"`
}

type _RuleProxyAction RuleProxyAction

// NewRuleProxyAction instantiates a new RuleProxyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleProxyAction(wafEnabled bool, notifyConfig ProxyNotifyConfig) *RuleProxyAction {
	this := RuleProxyAction{}
	this.WafEnabled = wafEnabled
	this.NotifyConfig = notifyConfig
	return &this
}

// NewRuleProxyActionWithDefaults instantiates a new RuleProxyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleProxyActionWithDefaults() *RuleProxyAction {
	this := RuleProxyAction{}
	var wafEnabled bool = false
	this.WafEnabled = wafEnabled
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
func (o *RuleProxyAction) GetOriginTimeout() int32 {
	if o == nil || IsNil(o.OriginTimeout) {
		var ret int32
		return ret
	}
	return *o.OriginTimeout
}

// GetOriginTimeoutOk returns a tuple with the OriginTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetOriginTimeoutOk() (*int32, bool) {
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

// SetOriginTimeout gets a reference to the given int32 and assigns it to the OriginTimeout field.
func (o *RuleProxyAction) SetOriginTimeout(v int32) {
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

// GetFailoverOriginStatusCode returns the FailoverOriginStatusCode field value if set, zero value otherwise.
func (o *RuleProxyAction) GetFailoverOriginStatusCode() []string {
	if o == nil || IsNil(o.FailoverOriginStatusCode) {
		var ret []string
		return ret
	}
	return o.FailoverOriginStatusCode
}

// GetFailoverOriginStatusCodeOk returns a tuple with the FailoverOriginStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetFailoverOriginStatusCodeOk() ([]string, bool) {
	if o == nil || IsNil(o.FailoverOriginStatusCode) {
		return nil, false
	}
	return o.FailoverOriginStatusCode, true
}

// HasFailoverOriginStatusCode returns a boolean if a field has been set.
func (o *RuleProxyAction) HasFailoverOriginStatusCode() bool {
	if o != nil && !IsNil(o.FailoverOriginStatusCode) {
		return true
	}

	return false
}

// SetFailoverOriginStatusCode gets a reference to the given []string and assigns it to the FailoverOriginStatusCode field.
func (o *RuleProxyAction) SetFailoverOriginStatusCode(v []string) {
	o.FailoverOriginStatusCode = v
}

// GetFailoverOriginTtfb returns the FailoverOriginTtfb field value if set, zero value otherwise.
func (o *RuleProxyAction) GetFailoverOriginTtfb() int32 {
	if o == nil || IsNil(o.FailoverOriginTtfb) {
		var ret int32
		return ret
	}
	return *o.FailoverOriginTtfb
}

// GetFailoverOriginTtfbOk returns a tuple with the FailoverOriginTtfb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetFailoverOriginTtfbOk() (*int32, bool) {
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

// SetFailoverOriginTtfb gets a reference to the given int32 and assigns it to the FailoverOriginTtfb field.
func (o *RuleProxyAction) SetFailoverOriginTtfb(v int32) {
	o.FailoverOriginTtfb = &v
}

// GetFailoverLifetime returns the FailoverLifetime field value if set, zero value otherwise.
func (o *RuleProxyAction) GetFailoverLifetime() int32 {
	if o == nil || IsNil(o.FailoverLifetime) {
		var ret int32
		return ret
	}
	return *o.FailoverLifetime
}

// GetFailoverLifetimeOk returns a tuple with the FailoverLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetFailoverLifetimeOk() (*int32, bool) {
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

// SetFailoverLifetime gets a reference to the given int32 and assigns it to the FailoverLifetime field.
func (o *RuleProxyAction) SetFailoverLifetime(v int32) {
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

// GetWafConfig returns the WafConfig field value if set, zero value otherwise.
func (o *RuleProxyAction) GetWafConfig() RuleWAFConfig {
	if o == nil || IsNil(o.WafConfig) {
		var ret RuleWAFConfig
		return ret
	}
	return *o.WafConfig
}

// GetWafConfigOk returns a tuple with the WafConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetWafConfigOk() (*RuleWAFConfig, bool) {
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

// SetWafConfig gets a reference to the given RuleWAFConfig and assigns it to the WafConfig field.
func (o *RuleProxyAction) SetWafConfig(v RuleWAFConfig) {
	o.WafConfig = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RuleProxyAction) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RuleProxyAction) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *RuleProxyAction) SetTo(v string) {
	o.To = &v
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

// GetNotifyConfig returns the NotifyConfig field value
func (o *RuleProxyAction) GetNotifyConfig() ProxyNotifyConfig {
	if o == nil {
		var ret ProxyNotifyConfig
		return ret
	}

	return o.NotifyConfig
}

// GetNotifyConfigOk returns a tuple with the NotifyConfig field value
// and a boolean to check if the value has been set.
func (o *RuleProxyAction) GetNotifyConfigOk() (*ProxyNotifyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyConfig, true
}

// SetNotifyConfig sets field value
func (o *RuleProxyAction) SetNotifyConfig(v ProxyNotifyConfig) {
	o.NotifyConfig = v
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
	if !IsNil(o.FailoverOriginStatusCode) {
		toSerialize["failover_origin_status_code"] = o.FailoverOriginStatusCode
	}
	if !IsNil(o.FailoverOriginTtfb) {
		toSerialize["failover_origin_ttfb"] = o.FailoverOriginTtfb
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
	if !IsNil(o.WafConfig) {
		toSerialize["waf_config"] = o.WafConfig
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.CacheLifetime) {
		toSerialize["cache_lifetime"] = o.CacheLifetime
	}
	if !IsNil(o.FailoverMode) {
		toSerialize["failover_mode"] = o.FailoverMode
	}
	if !IsNil(o.DisableSslVerify) {
		toSerialize["disable_ssl_verify"] = o.DisableSslVerify
	}
	toSerialize["notify_config"] = o.NotifyConfig
	return toSerialize, nil
}

func (o *RuleProxyAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"waf_enabled",
		"notify_config",
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

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRuleProxyAction)

	if err != nil {
		return err
	}

	*o = RuleProxyAction(varRuleProxyAction)

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


