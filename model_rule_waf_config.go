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

// checks if the RuleWAFConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleWAFConfig{}

// RuleWAFConfig struct for RuleWAFConfig
type RuleWAFConfig struct {
	Mode string `json:"mode"`
	ParanoiaLevel *int32 `json:"paranoia_level,omitempty"`
	AllowRules []string `json:"allow_rules,omitempty"`
	AllowIp []string `json:"allow_ip,omitempty"`
	BlockIp []string `json:"block_ip,omitempty"`
	BlockUa []string `json:"block_ua,omitempty"`
	BlockReferer []string `json:"block_referer,omitempty"`
	BlockLists *WafConfigBlockLists `json:"block_lists,omitempty"`
	Httpbl *ProxyConfigHttpbl `json:"httpbl,omitempty"`
	Thresholds []map[string]interface{} `json:"thresholds,omitempty"`
	HttpblEnabled *map[string]bool `json:"httpbl_enabled,omitempty"`
	NotifySlackHitsRpm *int32 `json:"notify_slack_hits_rpm,omitempty"`
	IpRatelimitMode *string `json:"ip_ratelimit_mode,omitempty"`
	IpRatelimitRps *int32 `json:"ip_ratelimit_rps,omitempty"`
	IpRatelimitCooldown *int32 `json:"ip_ratelimit_cooldown,omitempty"`
	RequestHeaderRatelimitMode *string `json:"request_header_ratelimit_mode,omitempty"`
	RequestHeaderName *string `json:"request_header_name,omitempty"`
	RequestHeaderRatelimitRps *int32 `json:"request_header_ratelimit_rps,omitempty"`
	RequestHeaderRatelimitCooldown *int32 `json:"request_header_ratelimit_cooldown,omitempty"`
	WafRatelimitMode *string `json:"waf_ratelimit_mode,omitempty"`
	WafRatelimitHits *int32 `json:"waf_ratelimit_hits,omitempty"`
	WafRatelimitRps *int32 `json:"waf_ratelimit_rps,omitempty"`
	WafRatelimitCooldown *int32 `json:"waf_ratelimit_cooldown,omitempty"`
	NotifyEmail []string `json:"notify_email,omitempty"`
	NotifySlack *string `json:"notify_slack,omitempty"`
	NotifySlackRpm *int32 `json:"notify_slack_rpm,omitempty"`
}

type _RuleWAFConfig RuleWAFConfig

// NewRuleWAFConfig instantiates a new RuleWAFConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleWAFConfig(mode string) *RuleWAFConfig {
	this := RuleWAFConfig{}
	this.Mode = mode
	var paranoiaLevel int32 = 1
	this.ParanoiaLevel = &paranoiaLevel
	var ipRatelimitMode string = "disabled"
	this.IpRatelimitMode = &ipRatelimitMode
	var ipRatelimitRps int32 = 5
	this.IpRatelimitRps = &ipRatelimitRps
	var ipRatelimitCooldown int32 = 30
	this.IpRatelimitCooldown = &ipRatelimitCooldown
	var requestHeaderRatelimitMode string = "disabled"
	this.RequestHeaderRatelimitMode = &requestHeaderRatelimitMode
	var requestHeaderRatelimitRps int32 = 5
	this.RequestHeaderRatelimitRps = &requestHeaderRatelimitRps
	var requestHeaderRatelimitCooldown int32 = 30
	this.RequestHeaderRatelimitCooldown = &requestHeaderRatelimitCooldown
	var wafRatelimitMode string = "disabled"
	this.WafRatelimitMode = &wafRatelimitMode
	var wafRatelimitHits int32 = 10
	this.WafRatelimitHits = &wafRatelimitHits
	var wafRatelimitRps int32 = 5
	this.WafRatelimitRps = &wafRatelimitRps
	var wafRatelimitCooldown int32 = 300
	this.WafRatelimitCooldown = &wafRatelimitCooldown
	return &this
}

// NewRuleWAFConfigWithDefaults instantiates a new RuleWAFConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleWAFConfigWithDefaults() *RuleWAFConfig {
	this := RuleWAFConfig{}
	var paranoiaLevel int32 = 1
	this.ParanoiaLevel = &paranoiaLevel
	var ipRatelimitMode string = "disabled"
	this.IpRatelimitMode = &ipRatelimitMode
	var ipRatelimitRps int32 = 5
	this.IpRatelimitRps = &ipRatelimitRps
	var ipRatelimitCooldown int32 = 30
	this.IpRatelimitCooldown = &ipRatelimitCooldown
	var requestHeaderRatelimitMode string = "disabled"
	this.RequestHeaderRatelimitMode = &requestHeaderRatelimitMode
	var requestHeaderRatelimitRps int32 = 5
	this.RequestHeaderRatelimitRps = &requestHeaderRatelimitRps
	var requestHeaderRatelimitCooldown int32 = 30
	this.RequestHeaderRatelimitCooldown = &requestHeaderRatelimitCooldown
	var wafRatelimitMode string = "disabled"
	this.WafRatelimitMode = &wafRatelimitMode
	var wafRatelimitHits int32 = 10
	this.WafRatelimitHits = &wafRatelimitHits
	var wafRatelimitRps int32 = 5
	this.WafRatelimitRps = &wafRatelimitRps
	var wafRatelimitCooldown int32 = 300
	this.WafRatelimitCooldown = &wafRatelimitCooldown
	return &this
}

// GetMode returns the Mode field value
func (o *RuleWAFConfig) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *RuleWAFConfig) SetMode(v string) {
	o.Mode = v
}

// GetParanoiaLevel returns the ParanoiaLevel field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetParanoiaLevel() int32 {
	if o == nil || IsNil(o.ParanoiaLevel) {
		var ret int32
		return ret
	}
	return *o.ParanoiaLevel
}

// GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetParanoiaLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ParanoiaLevel) {
		return nil, false
	}
	return o.ParanoiaLevel, true
}

// HasParanoiaLevel returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasParanoiaLevel() bool {
	if o != nil && !IsNil(o.ParanoiaLevel) {
		return true
	}

	return false
}

// SetParanoiaLevel gets a reference to the given int32 and assigns it to the ParanoiaLevel field.
func (o *RuleWAFConfig) SetParanoiaLevel(v int32) {
	o.ParanoiaLevel = &v
}

// GetAllowRules returns the AllowRules field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetAllowRules() []string {
	if o == nil || IsNil(o.AllowRules) {
		var ret []string
		return ret
	}
	return o.AllowRules
}

// GetAllowRulesOk returns a tuple with the AllowRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetAllowRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowRules) {
		return nil, false
	}
	return o.AllowRules, true
}

// HasAllowRules returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasAllowRules() bool {
	if o != nil && !IsNil(o.AllowRules) {
		return true
	}

	return false
}

// SetAllowRules gets a reference to the given []string and assigns it to the AllowRules field.
func (o *RuleWAFConfig) SetAllowRules(v []string) {
	o.AllowRules = v
}

// GetAllowIp returns the AllowIp field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetAllowIp() []string {
	if o == nil || IsNil(o.AllowIp) {
		var ret []string
		return ret
	}
	return o.AllowIp
}

// GetAllowIpOk returns a tuple with the AllowIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetAllowIpOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowIp) {
		return nil, false
	}
	return o.AllowIp, true
}

// HasAllowIp returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasAllowIp() bool {
	if o != nil && !IsNil(o.AllowIp) {
		return true
	}

	return false
}

// SetAllowIp gets a reference to the given []string and assigns it to the AllowIp field.
func (o *RuleWAFConfig) SetAllowIp(v []string) {
	o.AllowIp = v
}

// GetBlockIp returns the BlockIp field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetBlockIp() []string {
	if o == nil || IsNil(o.BlockIp) {
		var ret []string
		return ret
	}
	return o.BlockIp
}

// GetBlockIpOk returns a tuple with the BlockIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetBlockIpOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockIp) {
		return nil, false
	}
	return o.BlockIp, true
}

// HasBlockIp returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasBlockIp() bool {
	if o != nil && !IsNil(o.BlockIp) {
		return true
	}

	return false
}

// SetBlockIp gets a reference to the given []string and assigns it to the BlockIp field.
func (o *RuleWAFConfig) SetBlockIp(v []string) {
	o.BlockIp = v
}

// GetBlockUa returns the BlockUa field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetBlockUa() []string {
	if o == nil || IsNil(o.BlockUa) {
		var ret []string
		return ret
	}
	return o.BlockUa
}

// GetBlockUaOk returns a tuple with the BlockUa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetBlockUaOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockUa) {
		return nil, false
	}
	return o.BlockUa, true
}

// HasBlockUa returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasBlockUa() bool {
	if o != nil && !IsNil(o.BlockUa) {
		return true
	}

	return false
}

// SetBlockUa gets a reference to the given []string and assigns it to the BlockUa field.
func (o *RuleWAFConfig) SetBlockUa(v []string) {
	o.BlockUa = v
}

// GetBlockReferer returns the BlockReferer field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetBlockReferer() []string {
	if o == nil || IsNil(o.BlockReferer) {
		var ret []string
		return ret
	}
	return o.BlockReferer
}

// GetBlockRefererOk returns a tuple with the BlockReferer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetBlockRefererOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockReferer) {
		return nil, false
	}
	return o.BlockReferer, true
}

// HasBlockReferer returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasBlockReferer() bool {
	if o != nil && !IsNil(o.BlockReferer) {
		return true
	}

	return false
}

// SetBlockReferer gets a reference to the given []string and assigns it to the BlockReferer field.
func (o *RuleWAFConfig) SetBlockReferer(v []string) {
	o.BlockReferer = v
}

// GetBlockLists returns the BlockLists field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetBlockLists() WafConfigBlockLists {
	if o == nil || IsNil(o.BlockLists) {
		var ret WafConfigBlockLists
		return ret
	}
	return *o.BlockLists
}

// GetBlockListsOk returns a tuple with the BlockLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetBlockListsOk() (*WafConfigBlockLists, bool) {
	if o == nil || IsNil(o.BlockLists) {
		return nil, false
	}
	return o.BlockLists, true
}

// HasBlockLists returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasBlockLists() bool {
	if o != nil && !IsNil(o.BlockLists) {
		return true
	}

	return false
}

// SetBlockLists gets a reference to the given WafConfigBlockLists and assigns it to the BlockLists field.
func (o *RuleWAFConfig) SetBlockLists(v WafConfigBlockLists) {
	o.BlockLists = &v
}

// GetHttpbl returns the Httpbl field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetHttpbl() ProxyConfigHttpbl {
	if o == nil || IsNil(o.Httpbl) {
		var ret ProxyConfigHttpbl
		return ret
	}
	return *o.Httpbl
}

// GetHttpblOk returns a tuple with the Httpbl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetHttpblOk() (*ProxyConfigHttpbl, bool) {
	if o == nil || IsNil(o.Httpbl) {
		return nil, false
	}
	return o.Httpbl, true
}

// HasHttpbl returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasHttpbl() bool {
	if o != nil && !IsNil(o.Httpbl) {
		return true
	}

	return false
}

// SetHttpbl gets a reference to the given ProxyConfigHttpbl and assigns it to the Httpbl field.
func (o *RuleWAFConfig) SetHttpbl(v ProxyConfigHttpbl) {
	o.Httpbl = &v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetThresholds() []map[string]interface{} {
	if o == nil || IsNil(o.Thresholds) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetThresholdsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Thresholds) {
		return nil, false
	}
	return o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasThresholds() bool {
	if o != nil && !IsNil(o.Thresholds) {
		return true
	}

	return false
}

// SetThresholds gets a reference to the given []map[string]interface{} and assigns it to the Thresholds field.
func (o *RuleWAFConfig) SetThresholds(v []map[string]interface{}) {
	o.Thresholds = v
}

// GetHttpblEnabled returns the HttpblEnabled field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetHttpblEnabled() map[string]bool {
	if o == nil || IsNil(o.HttpblEnabled) {
		var ret map[string]bool
		return ret
	}
	return *o.HttpblEnabled
}

// GetHttpblEnabledOk returns a tuple with the HttpblEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetHttpblEnabledOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.HttpblEnabled) {
		return nil, false
	}
	return o.HttpblEnabled, true
}

// HasHttpblEnabled returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasHttpblEnabled() bool {
	if o != nil && !IsNil(o.HttpblEnabled) {
		return true
	}

	return false
}

// SetHttpblEnabled gets a reference to the given map[string]bool and assigns it to the HttpblEnabled field.
func (o *RuleWAFConfig) SetHttpblEnabled(v map[string]bool) {
	o.HttpblEnabled = &v
}

// GetNotifySlackHitsRpm returns the NotifySlackHitsRpm field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetNotifySlackHitsRpm() int32 {
	if o == nil || IsNil(o.NotifySlackHitsRpm) {
		var ret int32
		return ret
	}
	return *o.NotifySlackHitsRpm
}

// GetNotifySlackHitsRpmOk returns a tuple with the NotifySlackHitsRpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetNotifySlackHitsRpmOk() (*int32, bool) {
	if o == nil || IsNil(o.NotifySlackHitsRpm) {
		return nil, false
	}
	return o.NotifySlackHitsRpm, true
}

// HasNotifySlackHitsRpm returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasNotifySlackHitsRpm() bool {
	if o != nil && !IsNil(o.NotifySlackHitsRpm) {
		return true
	}

	return false
}

// SetNotifySlackHitsRpm gets a reference to the given int32 and assigns it to the NotifySlackHitsRpm field.
func (o *RuleWAFConfig) SetNotifySlackHitsRpm(v int32) {
	o.NotifySlackHitsRpm = &v
}

// GetIpRatelimitMode returns the IpRatelimitMode field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetIpRatelimitMode() string {
	if o == nil || IsNil(o.IpRatelimitMode) {
		var ret string
		return ret
	}
	return *o.IpRatelimitMode
}

// GetIpRatelimitModeOk returns a tuple with the IpRatelimitMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetIpRatelimitModeOk() (*string, bool) {
	if o == nil || IsNil(o.IpRatelimitMode) {
		return nil, false
	}
	return o.IpRatelimitMode, true
}

// HasIpRatelimitMode returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasIpRatelimitMode() bool {
	if o != nil && !IsNil(o.IpRatelimitMode) {
		return true
	}

	return false
}

// SetIpRatelimitMode gets a reference to the given string and assigns it to the IpRatelimitMode field.
func (o *RuleWAFConfig) SetIpRatelimitMode(v string) {
	o.IpRatelimitMode = &v
}

// GetIpRatelimitRps returns the IpRatelimitRps field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetIpRatelimitRps() int32 {
	if o == nil || IsNil(o.IpRatelimitRps) {
		var ret int32
		return ret
	}
	return *o.IpRatelimitRps
}

// GetIpRatelimitRpsOk returns a tuple with the IpRatelimitRps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetIpRatelimitRpsOk() (*int32, bool) {
	if o == nil || IsNil(o.IpRatelimitRps) {
		return nil, false
	}
	return o.IpRatelimitRps, true
}

// HasIpRatelimitRps returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasIpRatelimitRps() bool {
	if o != nil && !IsNil(o.IpRatelimitRps) {
		return true
	}

	return false
}

// SetIpRatelimitRps gets a reference to the given int32 and assigns it to the IpRatelimitRps field.
func (o *RuleWAFConfig) SetIpRatelimitRps(v int32) {
	o.IpRatelimitRps = &v
}

// GetIpRatelimitCooldown returns the IpRatelimitCooldown field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetIpRatelimitCooldown() int32 {
	if o == nil || IsNil(o.IpRatelimitCooldown) {
		var ret int32
		return ret
	}
	return *o.IpRatelimitCooldown
}

// GetIpRatelimitCooldownOk returns a tuple with the IpRatelimitCooldown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetIpRatelimitCooldownOk() (*int32, bool) {
	if o == nil || IsNil(o.IpRatelimitCooldown) {
		return nil, false
	}
	return o.IpRatelimitCooldown, true
}

// HasIpRatelimitCooldown returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasIpRatelimitCooldown() bool {
	if o != nil && !IsNil(o.IpRatelimitCooldown) {
		return true
	}

	return false
}

// SetIpRatelimitCooldown gets a reference to the given int32 and assigns it to the IpRatelimitCooldown field.
func (o *RuleWAFConfig) SetIpRatelimitCooldown(v int32) {
	o.IpRatelimitCooldown = &v
}

// GetRequestHeaderRatelimitMode returns the RequestHeaderRatelimitMode field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetRequestHeaderRatelimitMode() string {
	if o == nil || IsNil(o.RequestHeaderRatelimitMode) {
		var ret string
		return ret
	}
	return *o.RequestHeaderRatelimitMode
}

// GetRequestHeaderRatelimitModeOk returns a tuple with the RequestHeaderRatelimitMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetRequestHeaderRatelimitModeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestHeaderRatelimitMode) {
		return nil, false
	}
	return o.RequestHeaderRatelimitMode, true
}

// HasRequestHeaderRatelimitMode returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasRequestHeaderRatelimitMode() bool {
	if o != nil && !IsNil(o.RequestHeaderRatelimitMode) {
		return true
	}

	return false
}

// SetRequestHeaderRatelimitMode gets a reference to the given string and assigns it to the RequestHeaderRatelimitMode field.
func (o *RuleWAFConfig) SetRequestHeaderRatelimitMode(v string) {
	o.RequestHeaderRatelimitMode = &v
}

// GetRequestHeaderName returns the RequestHeaderName field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetRequestHeaderName() string {
	if o == nil || IsNil(o.RequestHeaderName) {
		var ret string
		return ret
	}
	return *o.RequestHeaderName
}

// GetRequestHeaderNameOk returns a tuple with the RequestHeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetRequestHeaderNameOk() (*string, bool) {
	if o == nil || IsNil(o.RequestHeaderName) {
		return nil, false
	}
	return o.RequestHeaderName, true
}

// HasRequestHeaderName returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasRequestHeaderName() bool {
	if o != nil && !IsNil(o.RequestHeaderName) {
		return true
	}

	return false
}

// SetRequestHeaderName gets a reference to the given string and assigns it to the RequestHeaderName field.
func (o *RuleWAFConfig) SetRequestHeaderName(v string) {
	o.RequestHeaderName = &v
}

// GetRequestHeaderRatelimitRps returns the RequestHeaderRatelimitRps field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetRequestHeaderRatelimitRps() int32 {
	if o == nil || IsNil(o.RequestHeaderRatelimitRps) {
		var ret int32
		return ret
	}
	return *o.RequestHeaderRatelimitRps
}

// GetRequestHeaderRatelimitRpsOk returns a tuple with the RequestHeaderRatelimitRps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetRequestHeaderRatelimitRpsOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestHeaderRatelimitRps) {
		return nil, false
	}
	return o.RequestHeaderRatelimitRps, true
}

// HasRequestHeaderRatelimitRps returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasRequestHeaderRatelimitRps() bool {
	if o != nil && !IsNil(o.RequestHeaderRatelimitRps) {
		return true
	}

	return false
}

// SetRequestHeaderRatelimitRps gets a reference to the given int32 and assigns it to the RequestHeaderRatelimitRps field.
func (o *RuleWAFConfig) SetRequestHeaderRatelimitRps(v int32) {
	o.RequestHeaderRatelimitRps = &v
}

// GetRequestHeaderRatelimitCooldown returns the RequestHeaderRatelimitCooldown field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetRequestHeaderRatelimitCooldown() int32 {
	if o == nil || IsNil(o.RequestHeaderRatelimitCooldown) {
		var ret int32
		return ret
	}
	return *o.RequestHeaderRatelimitCooldown
}

// GetRequestHeaderRatelimitCooldownOk returns a tuple with the RequestHeaderRatelimitCooldown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetRequestHeaderRatelimitCooldownOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestHeaderRatelimitCooldown) {
		return nil, false
	}
	return o.RequestHeaderRatelimitCooldown, true
}

// HasRequestHeaderRatelimitCooldown returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasRequestHeaderRatelimitCooldown() bool {
	if o != nil && !IsNil(o.RequestHeaderRatelimitCooldown) {
		return true
	}

	return false
}

// SetRequestHeaderRatelimitCooldown gets a reference to the given int32 and assigns it to the RequestHeaderRatelimitCooldown field.
func (o *RuleWAFConfig) SetRequestHeaderRatelimitCooldown(v int32) {
	o.RequestHeaderRatelimitCooldown = &v
}

// GetWafRatelimitMode returns the WafRatelimitMode field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetWafRatelimitMode() string {
	if o == nil || IsNil(o.WafRatelimitMode) {
		var ret string
		return ret
	}
	return *o.WafRatelimitMode
}

// GetWafRatelimitModeOk returns a tuple with the WafRatelimitMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetWafRatelimitModeOk() (*string, bool) {
	if o == nil || IsNil(o.WafRatelimitMode) {
		return nil, false
	}
	return o.WafRatelimitMode, true
}

// HasWafRatelimitMode returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasWafRatelimitMode() bool {
	if o != nil && !IsNil(o.WafRatelimitMode) {
		return true
	}

	return false
}

// SetWafRatelimitMode gets a reference to the given string and assigns it to the WafRatelimitMode field.
func (o *RuleWAFConfig) SetWafRatelimitMode(v string) {
	o.WafRatelimitMode = &v
}

// GetWafRatelimitHits returns the WafRatelimitHits field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetWafRatelimitHits() int32 {
	if o == nil || IsNil(o.WafRatelimitHits) {
		var ret int32
		return ret
	}
	return *o.WafRatelimitHits
}

// GetWafRatelimitHitsOk returns a tuple with the WafRatelimitHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetWafRatelimitHitsOk() (*int32, bool) {
	if o == nil || IsNil(o.WafRatelimitHits) {
		return nil, false
	}
	return o.WafRatelimitHits, true
}

// HasWafRatelimitHits returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasWafRatelimitHits() bool {
	if o != nil && !IsNil(o.WafRatelimitHits) {
		return true
	}

	return false
}

// SetWafRatelimitHits gets a reference to the given int32 and assigns it to the WafRatelimitHits field.
func (o *RuleWAFConfig) SetWafRatelimitHits(v int32) {
	o.WafRatelimitHits = &v
}

// GetWafRatelimitRps returns the WafRatelimitRps field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetWafRatelimitRps() int32 {
	if o == nil || IsNil(o.WafRatelimitRps) {
		var ret int32
		return ret
	}
	return *o.WafRatelimitRps
}

// GetWafRatelimitRpsOk returns a tuple with the WafRatelimitRps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetWafRatelimitRpsOk() (*int32, bool) {
	if o == nil || IsNil(o.WafRatelimitRps) {
		return nil, false
	}
	return o.WafRatelimitRps, true
}

// HasWafRatelimitRps returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasWafRatelimitRps() bool {
	if o != nil && !IsNil(o.WafRatelimitRps) {
		return true
	}

	return false
}

// SetWafRatelimitRps gets a reference to the given int32 and assigns it to the WafRatelimitRps field.
func (o *RuleWAFConfig) SetWafRatelimitRps(v int32) {
	o.WafRatelimitRps = &v
}

// GetWafRatelimitCooldown returns the WafRatelimitCooldown field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetWafRatelimitCooldown() int32 {
	if o == nil || IsNil(o.WafRatelimitCooldown) {
		var ret int32
		return ret
	}
	return *o.WafRatelimitCooldown
}

// GetWafRatelimitCooldownOk returns a tuple with the WafRatelimitCooldown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetWafRatelimitCooldownOk() (*int32, bool) {
	if o == nil || IsNil(o.WafRatelimitCooldown) {
		return nil, false
	}
	return o.WafRatelimitCooldown, true
}

// HasWafRatelimitCooldown returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasWafRatelimitCooldown() bool {
	if o != nil && !IsNil(o.WafRatelimitCooldown) {
		return true
	}

	return false
}

// SetWafRatelimitCooldown gets a reference to the given int32 and assigns it to the WafRatelimitCooldown field.
func (o *RuleWAFConfig) SetWafRatelimitCooldown(v int32) {
	o.WafRatelimitCooldown = &v
}

// GetNotifyEmail returns the NotifyEmail field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetNotifyEmail() []string {
	if o == nil || IsNil(o.NotifyEmail) {
		var ret []string
		return ret
	}
	return o.NotifyEmail
}

// GetNotifyEmailOk returns a tuple with the NotifyEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetNotifyEmailOk() ([]string, bool) {
	if o == nil || IsNil(o.NotifyEmail) {
		return nil, false
	}
	return o.NotifyEmail, true
}

// HasNotifyEmail returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasNotifyEmail() bool {
	if o != nil && !IsNil(o.NotifyEmail) {
		return true
	}

	return false
}

// SetNotifyEmail gets a reference to the given []string and assigns it to the NotifyEmail field.
func (o *RuleWAFConfig) SetNotifyEmail(v []string) {
	o.NotifyEmail = v
}

// GetNotifySlack returns the NotifySlack field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetNotifySlack() string {
	if o == nil || IsNil(o.NotifySlack) {
		var ret string
		return ret
	}
	return *o.NotifySlack
}

// GetNotifySlackOk returns a tuple with the NotifySlack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetNotifySlackOk() (*string, bool) {
	if o == nil || IsNil(o.NotifySlack) {
		return nil, false
	}
	return o.NotifySlack, true
}

// HasNotifySlack returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasNotifySlack() bool {
	if o != nil && !IsNil(o.NotifySlack) {
		return true
	}

	return false
}

// SetNotifySlack gets a reference to the given string and assigns it to the NotifySlack field.
func (o *RuleWAFConfig) SetNotifySlack(v string) {
	o.NotifySlack = &v
}

// GetNotifySlackRpm returns the NotifySlackRpm field value if set, zero value otherwise.
func (o *RuleWAFConfig) GetNotifySlackRpm() int32 {
	if o == nil || IsNil(o.NotifySlackRpm) {
		var ret int32
		return ret
	}
	return *o.NotifySlackRpm
}

// GetNotifySlackRpmOk returns a tuple with the NotifySlackRpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleWAFConfig) GetNotifySlackRpmOk() (*int32, bool) {
	if o == nil || IsNil(o.NotifySlackRpm) {
		return nil, false
	}
	return o.NotifySlackRpm, true
}

// HasNotifySlackRpm returns a boolean if a field has been set.
func (o *RuleWAFConfig) HasNotifySlackRpm() bool {
	if o != nil && !IsNil(o.NotifySlackRpm) {
		return true
	}

	return false
}

// SetNotifySlackRpm gets a reference to the given int32 and assigns it to the NotifySlackRpm field.
func (o *RuleWAFConfig) SetNotifySlackRpm(v int32) {
	o.NotifySlackRpm = &v
}

func (o RuleWAFConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleWAFConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mode"] = o.Mode
	if !IsNil(o.ParanoiaLevel) {
		toSerialize["paranoia_level"] = o.ParanoiaLevel
	}
	if !IsNil(o.AllowRules) {
		toSerialize["allow_rules"] = o.AllowRules
	}
	if !IsNil(o.AllowIp) {
		toSerialize["allow_ip"] = o.AllowIp
	}
	if !IsNil(o.BlockIp) {
		toSerialize["block_ip"] = o.BlockIp
	}
	if !IsNil(o.BlockUa) {
		toSerialize["block_ua"] = o.BlockUa
	}
	if !IsNil(o.BlockReferer) {
		toSerialize["block_referer"] = o.BlockReferer
	}
	if !IsNil(o.BlockLists) {
		toSerialize["block_lists"] = o.BlockLists
	}
	if !IsNil(o.Httpbl) {
		toSerialize["httpbl"] = o.Httpbl
	}
	if !IsNil(o.Thresholds) {
		toSerialize["thresholds"] = o.Thresholds
	}
	if !IsNil(o.HttpblEnabled) {
		toSerialize["httpbl_enabled"] = o.HttpblEnabled
	}
	if !IsNil(o.NotifySlackHitsRpm) {
		toSerialize["notify_slack_hits_rpm"] = o.NotifySlackHitsRpm
	}
	if !IsNil(o.IpRatelimitMode) {
		toSerialize["ip_ratelimit_mode"] = o.IpRatelimitMode
	}
	if !IsNil(o.IpRatelimitRps) {
		toSerialize["ip_ratelimit_rps"] = o.IpRatelimitRps
	}
	if !IsNil(o.IpRatelimitCooldown) {
		toSerialize["ip_ratelimit_cooldown"] = o.IpRatelimitCooldown
	}
	if !IsNil(o.RequestHeaderRatelimitMode) {
		toSerialize["request_header_ratelimit_mode"] = o.RequestHeaderRatelimitMode
	}
	if !IsNil(o.RequestHeaderName) {
		toSerialize["request_header_name"] = o.RequestHeaderName
	}
	if !IsNil(o.RequestHeaderRatelimitRps) {
		toSerialize["request_header_ratelimit_rps"] = o.RequestHeaderRatelimitRps
	}
	if !IsNil(o.RequestHeaderRatelimitCooldown) {
		toSerialize["request_header_ratelimit_cooldown"] = o.RequestHeaderRatelimitCooldown
	}
	if !IsNil(o.WafRatelimitMode) {
		toSerialize["waf_ratelimit_mode"] = o.WafRatelimitMode
	}
	if !IsNil(o.WafRatelimitHits) {
		toSerialize["waf_ratelimit_hits"] = o.WafRatelimitHits
	}
	if !IsNil(o.WafRatelimitRps) {
		toSerialize["waf_ratelimit_rps"] = o.WafRatelimitRps
	}
	if !IsNil(o.WafRatelimitCooldown) {
		toSerialize["waf_ratelimit_cooldown"] = o.WafRatelimitCooldown
	}
	if !IsNil(o.NotifyEmail) {
		toSerialize["notify_email"] = o.NotifyEmail
	}
	if !IsNil(o.NotifySlack) {
		toSerialize["notify_slack"] = o.NotifySlack
	}
	if !IsNil(o.NotifySlackRpm) {
		toSerialize["notify_slack_rpm"] = o.NotifySlackRpm
	}
	return toSerialize, nil
}

func (o *RuleWAFConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
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

	varRuleWAFConfig := _RuleWAFConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRuleWAFConfig)

	if err != nil {
		return err
	}

	*o = RuleWAFConfig(varRuleWAFConfig)

	return err
}

type NullableRuleWAFConfig struct {
	value *RuleWAFConfig
	isSet bool
}

func (v NullableRuleWAFConfig) Get() *RuleWAFConfig {
	return v.value
}

func (v *NullableRuleWAFConfig) Set(val *RuleWAFConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleWAFConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleWAFConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleWAFConfig(val *RuleWAFConfig) *NullableRuleWAFConfig {
	return &NullableRuleWAFConfig{value: val, isSet: true}
}

func (v NullableRuleWAFConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleWAFConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


