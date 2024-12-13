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

// checks if the ProxyConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyConfig{}

// ProxyConfig struct for ProxyConfig
type ProxyConfig struct {
	To string `json:"to"`
	Host *string `json:"host,omitempty"`
	AuthUser *string `json:"auth_user,omitempty"`
	AuthPass *string `json:"auth_pass,omitempty"`
	DisableSslVerify *bool `json:"disable_ssl_verify,omitempty"`
	CacheLifetime *int32 `json:"cache_lifetime,omitempty"`
	OnlyProxy404 *bool `json:"only_proxy_404,omitempty"`
	InjectHeaders *map[string]string `json:"inject_headers,omitempty"`
	ProxyStripHeaders []string `json:"proxy_strip_headers,omitempty"`
	ProxyStripRequestHeaders []string `json:"proxy_strip_request_headers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProxyConfig ProxyConfig

// NewProxyConfig instantiates a new ProxyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyConfig(to string) *ProxyConfig {
	this := ProxyConfig{}
	this.To = to
	return &this
}

// NewProxyConfigWithDefaults instantiates a new ProxyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyConfigWithDefaults() *ProxyConfig {
	this := ProxyConfig{}
	return &this
}

// GetTo returns the To field value
func (o *ProxyConfig) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *ProxyConfig) SetTo(v string) {
	o.To = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ProxyConfig) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ProxyConfig) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ProxyConfig) SetHost(v string) {
	o.Host = &v
}

// GetAuthUser returns the AuthUser field value if set, zero value otherwise.
func (o *ProxyConfig) GetAuthUser() string {
	if o == nil || IsNil(o.AuthUser) {
		var ret string
		return ret
	}
	return *o.AuthUser
}

// GetAuthUserOk returns a tuple with the AuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetAuthUserOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUser) {
		return nil, false
	}
	return o.AuthUser, true
}

// HasAuthUser returns a boolean if a field has been set.
func (o *ProxyConfig) HasAuthUser() bool {
	if o != nil && !IsNil(o.AuthUser) {
		return true
	}

	return false
}

// SetAuthUser gets a reference to the given string and assigns it to the AuthUser field.
func (o *ProxyConfig) SetAuthUser(v string) {
	o.AuthUser = &v
}

// GetAuthPass returns the AuthPass field value if set, zero value otherwise.
func (o *ProxyConfig) GetAuthPass() string {
	if o == nil || IsNil(o.AuthPass) {
		var ret string
		return ret
	}
	return *o.AuthPass
}

// GetAuthPassOk returns a tuple with the AuthPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetAuthPassOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPass) {
		return nil, false
	}
	return o.AuthPass, true
}

// HasAuthPass returns a boolean if a field has been set.
func (o *ProxyConfig) HasAuthPass() bool {
	if o != nil && !IsNil(o.AuthPass) {
		return true
	}

	return false
}

// SetAuthPass gets a reference to the given string and assigns it to the AuthPass field.
func (o *ProxyConfig) SetAuthPass(v string) {
	o.AuthPass = &v
}

// GetDisableSslVerify returns the DisableSslVerify field value if set, zero value otherwise.
func (o *ProxyConfig) GetDisableSslVerify() bool {
	if o == nil || IsNil(o.DisableSslVerify) {
		var ret bool
		return ret
	}
	return *o.DisableSslVerify
}

// GetDisableSslVerifyOk returns a tuple with the DisableSslVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetDisableSslVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableSslVerify) {
		return nil, false
	}
	return o.DisableSslVerify, true
}

// HasDisableSslVerify returns a boolean if a field has been set.
func (o *ProxyConfig) HasDisableSslVerify() bool {
	if o != nil && !IsNil(o.DisableSslVerify) {
		return true
	}

	return false
}

// SetDisableSslVerify gets a reference to the given bool and assigns it to the DisableSslVerify field.
func (o *ProxyConfig) SetDisableSslVerify(v bool) {
	o.DisableSslVerify = &v
}

// GetCacheLifetime returns the CacheLifetime field value if set, zero value otherwise.
func (o *ProxyConfig) GetCacheLifetime() int32 {
	if o == nil || IsNil(o.CacheLifetime) {
		var ret int32
		return ret
	}
	return *o.CacheLifetime
}

// GetCacheLifetimeOk returns a tuple with the CacheLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetCacheLifetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheLifetime) {
		return nil, false
	}
	return o.CacheLifetime, true
}

// HasCacheLifetime returns a boolean if a field has been set.
func (o *ProxyConfig) HasCacheLifetime() bool {
	if o != nil && !IsNil(o.CacheLifetime) {
		return true
	}

	return false
}

// SetCacheLifetime gets a reference to the given int32 and assigns it to the CacheLifetime field.
func (o *ProxyConfig) SetCacheLifetime(v int32) {
	o.CacheLifetime = &v
}

// GetOnlyProxy404 returns the OnlyProxy404 field value if set, zero value otherwise.
func (o *ProxyConfig) GetOnlyProxy404() bool {
	if o == nil || IsNil(o.OnlyProxy404) {
		var ret bool
		return ret
	}
	return *o.OnlyProxy404
}

// GetOnlyProxy404Ok returns a tuple with the OnlyProxy404 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetOnlyProxy404Ok() (*bool, bool) {
	if o == nil || IsNil(o.OnlyProxy404) {
		return nil, false
	}
	return o.OnlyProxy404, true
}

// HasOnlyProxy404 returns a boolean if a field has been set.
func (o *ProxyConfig) HasOnlyProxy404() bool {
	if o != nil && !IsNil(o.OnlyProxy404) {
		return true
	}

	return false
}

// SetOnlyProxy404 gets a reference to the given bool and assigns it to the OnlyProxy404 field.
func (o *ProxyConfig) SetOnlyProxy404(v bool) {
	o.OnlyProxy404 = &v
}

// GetInjectHeaders returns the InjectHeaders field value if set, zero value otherwise.
func (o *ProxyConfig) GetInjectHeaders() map[string]string {
	if o == nil || IsNil(o.InjectHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.InjectHeaders
}

// GetInjectHeadersOk returns a tuple with the InjectHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetInjectHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.InjectHeaders) {
		return nil, false
	}
	return o.InjectHeaders, true
}

// HasInjectHeaders returns a boolean if a field has been set.
func (o *ProxyConfig) HasInjectHeaders() bool {
	if o != nil && !IsNil(o.InjectHeaders) {
		return true
	}

	return false
}

// SetInjectHeaders gets a reference to the given map[string]string and assigns it to the InjectHeaders field.
func (o *ProxyConfig) SetInjectHeaders(v map[string]string) {
	o.InjectHeaders = &v
}

// GetProxyStripHeaders returns the ProxyStripHeaders field value if set, zero value otherwise.
func (o *ProxyConfig) GetProxyStripHeaders() []string {
	if o == nil || IsNil(o.ProxyStripHeaders) {
		var ret []string
		return ret
	}
	return o.ProxyStripHeaders
}

// GetProxyStripHeadersOk returns a tuple with the ProxyStripHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetProxyStripHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.ProxyStripHeaders) {
		return nil, false
	}
	return o.ProxyStripHeaders, true
}

// HasProxyStripHeaders returns a boolean if a field has been set.
func (o *ProxyConfig) HasProxyStripHeaders() bool {
	if o != nil && !IsNil(o.ProxyStripHeaders) {
		return true
	}

	return false
}

// SetProxyStripHeaders gets a reference to the given []string and assigns it to the ProxyStripHeaders field.
func (o *ProxyConfig) SetProxyStripHeaders(v []string) {
	o.ProxyStripHeaders = v
}

// GetProxyStripRequestHeaders returns the ProxyStripRequestHeaders field value if set, zero value otherwise.
func (o *ProxyConfig) GetProxyStripRequestHeaders() []string {
	if o == nil || IsNil(o.ProxyStripRequestHeaders) {
		var ret []string
		return ret
	}
	return o.ProxyStripRequestHeaders
}

// GetProxyStripRequestHeadersOk returns a tuple with the ProxyStripRequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyConfig) GetProxyStripRequestHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.ProxyStripRequestHeaders) {
		return nil, false
	}
	return o.ProxyStripRequestHeaders, true
}

// HasProxyStripRequestHeaders returns a boolean if a field has been set.
func (o *ProxyConfig) HasProxyStripRequestHeaders() bool {
	if o != nil && !IsNil(o.ProxyStripRequestHeaders) {
		return true
	}

	return false
}

// SetProxyStripRequestHeaders gets a reference to the given []string and assigns it to the ProxyStripRequestHeaders field.
func (o *ProxyConfig) SetProxyStripRequestHeaders(v []string) {
	o.ProxyStripRequestHeaders = v
}

func (o ProxyConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["to"] = o.To
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProxyConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varProxyConfig := _ProxyConfig{}

	err = json.Unmarshal(data, &varProxyConfig)

	if err != nil {
		return err
	}

	*o = ProxyConfig(varProxyConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProxyConfig struct {
	value *ProxyConfig
	isSet bool
}

func (v NullableProxyConfig) Get() *ProxyConfig {
	return v.value
}

func (v *NullableProxyConfig) Set(val *ProxyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyConfig(val *ProxyConfig) *NullableProxyConfig {
	return &NullableProxyConfig{value: val, isSet: true}
}

func (v NullableProxyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


