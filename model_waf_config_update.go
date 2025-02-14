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

// checks if the WAFConfigUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WAFConfigUpdate{}

// WAFConfigUpdate struct for WAFConfigUpdate
type WAFConfigUpdate struct {
	Mode *string `json:"mode,omitempty"`
	ParanoiaLevel *int32 `json:"paranoia_level,omitempty"`
	AllowRules []string `json:"allow_rules,omitempty"`
	AllowIp []string `json:"allow_ip,omitempty"`
	BlockIp []string `json:"block_ip,omitempty"`
	BlockUa []string `json:"block_ua,omitempty"`
	BlockReferer []string `json:"block_referer,omitempty"`
	Thresholds []Threshold `json:"thresholds,omitempty"`
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
	AdditionalProperties map[string]interface{}
}

type _WAFConfigUpdate WAFConfigUpdate

// NewWAFConfigUpdate instantiates a new WAFConfigUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWAFConfigUpdate() *WAFConfigUpdate {
	this := WAFConfigUpdate{}
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

// NewWAFConfigUpdateWithDefaults instantiates a new WAFConfigUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWAFConfigUpdateWithDefaults() *WAFConfigUpdate {
	this := WAFConfigUpdate{}
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

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *WAFConfigUpdate) SetMode(v string) {
	o.Mode = &v
}

// GetParanoiaLevel returns the ParanoiaLevel field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetParanoiaLevel() int32 {
	if o == nil || IsNil(o.ParanoiaLevel) {
		var ret int32
		return ret
	}
	return *o.ParanoiaLevel
}

// GetParanoiaLevelOk returns a tuple with the ParanoiaLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetParanoiaLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ParanoiaLevel) {
		return nil, false
	}
	return o.ParanoiaLevel, true
}

// HasParanoiaLevel returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasParanoiaLevel() bool {
	if o != nil && !IsNil(o.ParanoiaLevel) {
		return true
	}

	return false
}

// SetParanoiaLevel gets a reference to the given int32 and assigns it to the ParanoiaLevel field.
func (o *WAFConfigUpdate) SetParanoiaLevel(v int32) {
	o.ParanoiaLevel = &v
}

// GetAllowRules returns the AllowRules field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetAllowRules() []string {
	if o == nil || IsNil(o.AllowRules) {
		var ret []string
		return ret
	}
	return o.AllowRules
}

// GetAllowRulesOk returns a tuple with the AllowRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetAllowRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowRules) {
		return nil, false
	}
	return o.AllowRules, true
}

// HasAllowRules returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasAllowRules() bool {
	if o != nil && !IsNil(o.AllowRules) {
		return true
	}

	return false
}

// SetAllowRules gets a reference to the given []string and assigns it to the AllowRules field.
func (o *WAFConfigUpdate) SetAllowRules(v []string) {
	o.AllowRules = v
}

// GetAllowIp returns the AllowIp field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetAllowIp() []string {
	if o == nil || IsNil(o.AllowIp) {
		var ret []string
		return ret
	}
	return o.AllowIp
}

// GetAllowIpOk returns a tuple with the AllowIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetAllowIpOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowIp) {
		return nil, false
	}
	return o.AllowIp, true
}

// HasAllowIp returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasAllowIp() bool {
	if o != nil && !IsNil(o.AllowIp) {
		return true
	}

	return false
}

// SetAllowIp gets a reference to the given []string and assigns it to the AllowIp field.
func (o *WAFConfigUpdate) SetAllowIp(v []string) {
	o.AllowIp = v
}

// GetBlockIp returns the BlockIp field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetBlockIp() []string {
	if o == nil || IsNil(o.BlockIp) {
		var ret []string
		return ret
	}
	return o.BlockIp
}

// GetBlockIpOk returns a tuple with the BlockIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetBlockIpOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockIp) {
		return nil, false
	}
	return o.BlockIp, true
}

// HasBlockIp returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasBlockIp() bool {
	if o != nil && !IsNil(o.BlockIp) {
		return true
	}

	return false
}

// SetBlockIp gets a reference to the given []string and assigns it to the BlockIp field.
func (o *WAFConfigUpdate) SetBlockIp(v []string) {
	o.BlockIp = v
}

// GetBlockUa returns the BlockUa field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetBlockUa() []string {
	if o == nil || IsNil(o.BlockUa) {
		var ret []string
		return ret
	}
	return o.BlockUa
}

// GetBlockUaOk returns a tuple with the BlockUa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetBlockUaOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockUa) {
		return nil, false
	}
	return o.BlockUa, true
}

// HasBlockUa returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasBlockUa() bool {
	if o != nil && !IsNil(o.BlockUa) {
		return true
	}

	return false
}

// SetBlockUa gets a reference to the given []string and assigns it to the BlockUa field.
func (o *WAFConfigUpdate) SetBlockUa(v []string) {
	o.BlockUa = v
}

// GetBlockReferer returns the BlockReferer field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetBlockReferer() []string {
	if o == nil || IsNil(o.BlockReferer) {
		var ret []string
		return ret
	}
	return o.BlockReferer
}

// GetBlockRefererOk returns a tuple with the BlockReferer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetBlockRefererOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockReferer) {
		return nil, false
	}
	return o.BlockReferer, true
}

// HasBlockReferer returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasBlockReferer() bool {
	if o != nil && !IsNil(o.BlockReferer) {
		return true
	}

	return false
}

// SetBlockReferer gets a reference to the given []string and assigns it to the BlockReferer field.
func (o *WAFConfigUpdate) SetBlockReferer(v []string) {
	o.BlockReferer = v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetThresholds() []Threshold {
	if o == nil || IsNil(o.Thresholds) {
		var ret []Threshold
		return ret
	}
	return o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetThresholdsOk() ([]Threshold, bool) {
	if o == nil || IsNil(o.Thresholds) {
		return nil, false
	}
	return o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasThresholds() bool {
	if o != nil && !IsNil(o.Thresholds) {
		return true
	}

	return false
}

// SetThresholds gets a reference to the given []Threshold and assigns it to the Thresholds field.
func (o *WAFConfigUpdate) SetThresholds(v []Threshold) {
	o.Thresholds = v
}

// GetHttpblEnabled returns the HttpblEnabled field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetHttpblEnabled() map[string]bool {
	if o == nil || IsNil(o.HttpblEnabled) {
		var ret map[string]bool
		return ret
	}
	return *o.HttpblEnabled
}

// GetHttpblEnabledOk returns a tuple with the HttpblEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetHttpblEnabledOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.HttpblEnabled) {
		return nil, false
	}
	return o.HttpblEnabled, true
}

// HasHttpblEnabled returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasHttpblEnabled() bool {
	if o != nil && !IsNil(o.HttpblEnabled) {
		return true
	}

	return false
}

// SetHttpblEnabled gets a reference to the given map[string]bool and assigns it to the HttpblEnabled field.
func (o *WAFConfigUpdate) SetHttpblEnabled(v map[string]bool) {
	o.HttpblEnabled = &v
}

// GetNotifySlackHitsRpm returns the NotifySlackHitsRpm field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetNotifySlackHitsRpm() int32 {
	if o == nil || IsNil(o.NotifySlackHitsRpm) {
		var ret int32
		return ret
	}
	return *o.NotifySlackHitsRpm
}

// GetNotifySlackHitsRpmOk returns a tuple with the NotifySlackHitsRpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetNotifySlackHitsRpmOk() (*int32, bool) {
	if o == nil || IsNil(o.NotifySlackHitsRpm) {
		return nil, false
	}
	return o.NotifySlackHitsRpm, true
}

// HasNotifySlackHitsRpm returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasNotifySlackHitsRpm() bool {
	if o != nil && !IsNil(o.NotifySlackHitsRpm) {
		return true
	}

	return false
}

// SetNotifySlackHitsRpm gets a reference to the given int32 and assigns it to the NotifySlackHitsRpm field.
func (o *WAFConfigUpdate) SetNotifySlackHitsRpm(v int32) {
	o.NotifySlackHitsRpm = &v
}

// GetIpRatelimitMode returns the IpRatelimitMode field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetIpRatelimitMode() string {
	if o == nil || IsNil(o.IpRatelimitMode) {
		var ret string
		return ret
	}
	return *o.IpRatelimitMode
}

// GetIpRatelimitModeOk returns a tuple with the IpRatelimitMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetIpRatelimitModeOk() (*string, bool) {
	if o == nil || IsNil(o.IpRatelimitMode) {
		return nil, false
	}
	return o.IpRatelimitMode, true
}

// HasIpRatelimitMode returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasIpRatelimitMode() bool {
	if o != nil && !IsNil(o.IpRatelimitMode) {
		return true
	}

	return false
}

// SetIpRatelimitMode gets a reference to the given string and assigns it to the IpRatelimitMode field.
func (o *WAFConfigUpdate) SetIpRatelimitMode(v string) {
	o.IpRatelimitMode = &v
}

// GetIpRatelimitRps returns the IpRatelimitRps field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetIpRatelimitRps() int32 {
	if o == nil || IsNil(o.IpRatelimitRps) {
		var ret int32
		return ret
	}
	return *o.IpRatelimitRps
}

// GetIpRatelimitRpsOk returns a tuple with the IpRatelimitRps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetIpRatelimitRpsOk() (*int32, bool) {
	if o == nil || IsNil(o.IpRatelimitRps) {
		return nil, false
	}
	return o.IpRatelimitRps, true
}

// HasIpRatelimitRps returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasIpRatelimitRps() bool {
	if o != nil && !IsNil(o.IpRatelimitRps) {
		return true
	}

	return false
}

// SetIpRatelimitRps gets a reference to the given int32 and assigns it to the IpRatelimitRps field.
func (o *WAFConfigUpdate) SetIpRatelimitRps(v int32) {
	o.IpRatelimitRps = &v
}

// GetIpRatelimitCooldown returns the IpRatelimitCooldown field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetIpRatelimitCooldown() int32 {
	if o == nil || IsNil(o.IpRatelimitCooldown) {
		var ret int32
		return ret
	}
	return *o.IpRatelimitCooldown
}

// GetIpRatelimitCooldownOk returns a tuple with the IpRatelimitCooldown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetIpRatelimitCooldownOk() (*int32, bool) {
	if o == nil || IsNil(o.IpRatelimitCooldown) {
		return nil, false
	}
	return o.IpRatelimitCooldown, true
}

// HasIpRatelimitCooldown returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasIpRatelimitCooldown() bool {
	if o != nil && !IsNil(o.IpRatelimitCooldown) {
		return true
	}

	return false
}

// SetIpRatelimitCooldown gets a reference to the given int32 and assigns it to the IpRatelimitCooldown field.
func (o *WAFConfigUpdate) SetIpRatelimitCooldown(v int32) {
	o.IpRatelimitCooldown = &v
}

// GetRequestHeaderRatelimitMode returns the RequestHeaderRatelimitMode field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetRequestHeaderRatelimitMode() string {
	if o == nil || IsNil(o.RequestHeaderRatelimitMode) {
		var ret string
		return ret
	}
	return *o.RequestHeaderRatelimitMode
}

// GetRequestHeaderRatelimitModeOk returns a tuple with the RequestHeaderRatelimitMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetRequestHeaderRatelimitModeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestHeaderRatelimitMode) {
		return nil, false
	}
	return o.RequestHeaderRatelimitMode, true
}

// HasRequestHeaderRatelimitMode returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasRequestHeaderRatelimitMode() bool {
	if o != nil && !IsNil(o.RequestHeaderRatelimitMode) {
		return true
	}

	return false
}

// SetRequestHeaderRatelimitMode gets a reference to the given string and assigns it to the RequestHeaderRatelimitMode field.
func (o *WAFConfigUpdate) SetRequestHeaderRatelimitMode(v string) {
	o.RequestHeaderRatelimitMode = &v
}

// GetRequestHeaderName returns the RequestHeaderName field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetRequestHeaderName() string {
	if o == nil || IsNil(o.RequestHeaderName) {
		var ret string
		return ret
	}
	return *o.RequestHeaderName
}

// GetRequestHeaderNameOk returns a tuple with the RequestHeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetRequestHeaderNameOk() (*string, bool) {
	if o == nil || IsNil(o.RequestHeaderName) {
		return nil, false
	}
	return o.RequestHeaderName, true
}

// HasRequestHeaderName returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasRequestHeaderName() bool {
	if o != nil && !IsNil(o.RequestHeaderName) {
		return true
	}

	return false
}

// SetRequestHeaderName gets a reference to the given string and assigns it to the RequestHeaderName field.
func (o *WAFConfigUpdate) SetRequestHeaderName(v string) {
	o.RequestHeaderName = &v
}

// GetRequestHeaderRatelimitRps returns the RequestHeaderRatelimitRps field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetRequestHeaderRatelimitRps() int32 {
	if o == nil || IsNil(o.RequestHeaderRatelimitRps) {
		var ret int32
		return ret
	}
	return *o.RequestHeaderRatelimitRps
}

// GetRequestHeaderRatelimitRpsOk returns a tuple with the RequestHeaderRatelimitRps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetRequestHeaderRatelimitRpsOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestHeaderRatelimitRps) {
		return nil, false
	}
	return o.RequestHeaderRatelimitRps, true
}

// HasRequestHeaderRatelimitRps returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasRequestHeaderRatelimitRps() bool {
	if o != nil && !IsNil(o.RequestHeaderRatelimitRps) {
		return true
	}

	return false
}

// SetRequestHeaderRatelimitRps gets a reference to the given int32 and assigns it to the RequestHeaderRatelimitRps field.
func (o *WAFConfigUpdate) SetRequestHeaderRatelimitRps(v int32) {
	o.RequestHeaderRatelimitRps = &v
}

// GetRequestHeaderRatelimitCooldown returns the RequestHeaderRatelimitCooldown field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetRequestHeaderRatelimitCooldown() int32 {
	if o == nil || IsNil(o.RequestHeaderRatelimitCooldown) {
		var ret int32
		return ret
	}
	return *o.RequestHeaderRatelimitCooldown
}

// GetRequestHeaderRatelimitCooldownOk returns a tuple with the RequestHeaderRatelimitCooldown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetRequestHeaderRatelimitCooldownOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestHeaderRatelimitCooldown) {
		return nil, false
	}
	return o.RequestHeaderRatelimitCooldown, true
}

// HasRequestHeaderRatelimitCooldown returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasRequestHeaderRatelimitCooldown() bool {
	if o != nil && !IsNil(o.RequestHeaderRatelimitCooldown) {
		return true
	}

	return false
}

// SetRequestHeaderRatelimitCooldown gets a reference to the given int32 and assigns it to the RequestHeaderRatelimitCooldown field.
func (o *WAFConfigUpdate) SetRequestHeaderRatelimitCooldown(v int32) {
	o.RequestHeaderRatelimitCooldown = &v
}

// GetWafRatelimitMode returns the WafRatelimitMode field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetWafRatelimitMode() string {
	if o == nil || IsNil(o.WafRatelimitMode) {
		var ret string
		return ret
	}
	return *o.WafRatelimitMode
}

// GetWafRatelimitModeOk returns a tuple with the WafRatelimitMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetWafRatelimitModeOk() (*string, bool) {
	if o == nil || IsNil(o.WafRatelimitMode) {
		return nil, false
	}
	return o.WafRatelimitMode, true
}

// HasWafRatelimitMode returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasWafRatelimitMode() bool {
	if o != nil && !IsNil(o.WafRatelimitMode) {
		return true
	}

	return false
}

// SetWafRatelimitMode gets a reference to the given string and assigns it to the WafRatelimitMode field.
func (o *WAFConfigUpdate) SetWafRatelimitMode(v string) {
	o.WafRatelimitMode = &v
}

// GetWafRatelimitHits returns the WafRatelimitHits field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetWafRatelimitHits() int32 {
	if o == nil || IsNil(o.WafRatelimitHits) {
		var ret int32
		return ret
	}
	return *o.WafRatelimitHits
}

// GetWafRatelimitHitsOk returns a tuple with the WafRatelimitHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetWafRatelimitHitsOk() (*int32, bool) {
	if o == nil || IsNil(o.WafRatelimitHits) {
		return nil, false
	}
	return o.WafRatelimitHits, true
}

// HasWafRatelimitHits returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasWafRatelimitHits() bool {
	if o != nil && !IsNil(o.WafRatelimitHits) {
		return true
	}

	return false
}

// SetWafRatelimitHits gets a reference to the given int32 and assigns it to the WafRatelimitHits field.
func (o *WAFConfigUpdate) SetWafRatelimitHits(v int32) {
	o.WafRatelimitHits = &v
}

// GetWafRatelimitRps returns the WafRatelimitRps field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetWafRatelimitRps() int32 {
	if o == nil || IsNil(o.WafRatelimitRps) {
		var ret int32
		return ret
	}
	return *o.WafRatelimitRps
}

// GetWafRatelimitRpsOk returns a tuple with the WafRatelimitRps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetWafRatelimitRpsOk() (*int32, bool) {
	if o == nil || IsNil(o.WafRatelimitRps) {
		return nil, false
	}
	return o.WafRatelimitRps, true
}

// HasWafRatelimitRps returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasWafRatelimitRps() bool {
	if o != nil && !IsNil(o.WafRatelimitRps) {
		return true
	}

	return false
}

// SetWafRatelimitRps gets a reference to the given int32 and assigns it to the WafRatelimitRps field.
func (o *WAFConfigUpdate) SetWafRatelimitRps(v int32) {
	o.WafRatelimitRps = &v
}

// GetWafRatelimitCooldown returns the WafRatelimitCooldown field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetWafRatelimitCooldown() int32 {
	if o == nil || IsNil(o.WafRatelimitCooldown) {
		var ret int32
		return ret
	}
	return *o.WafRatelimitCooldown
}

// GetWafRatelimitCooldownOk returns a tuple with the WafRatelimitCooldown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetWafRatelimitCooldownOk() (*int32, bool) {
	if o == nil || IsNil(o.WafRatelimitCooldown) {
		return nil, false
	}
	return o.WafRatelimitCooldown, true
}

// HasWafRatelimitCooldown returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasWafRatelimitCooldown() bool {
	if o != nil && !IsNil(o.WafRatelimitCooldown) {
		return true
	}

	return false
}

// SetWafRatelimitCooldown gets a reference to the given int32 and assigns it to the WafRatelimitCooldown field.
func (o *WAFConfigUpdate) SetWafRatelimitCooldown(v int32) {
	o.WafRatelimitCooldown = &v
}

// GetNotifyEmail returns the NotifyEmail field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetNotifyEmail() []string {
	if o == nil || IsNil(o.NotifyEmail) {
		var ret []string
		return ret
	}
	return o.NotifyEmail
}

// GetNotifyEmailOk returns a tuple with the NotifyEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetNotifyEmailOk() ([]string, bool) {
	if o == nil || IsNil(o.NotifyEmail) {
		return nil, false
	}
	return o.NotifyEmail, true
}

// HasNotifyEmail returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasNotifyEmail() bool {
	if o != nil && !IsNil(o.NotifyEmail) {
		return true
	}

	return false
}

// SetNotifyEmail gets a reference to the given []string and assigns it to the NotifyEmail field.
func (o *WAFConfigUpdate) SetNotifyEmail(v []string) {
	o.NotifyEmail = v
}

// GetNotifySlack returns the NotifySlack field value if set, zero value otherwise.
func (o *WAFConfigUpdate) GetNotifySlack() string {
	if o == nil || IsNil(o.NotifySlack) {
		var ret string
		return ret
	}
	return *o.NotifySlack
}

// GetNotifySlackOk returns a tuple with the NotifySlack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFConfigUpdate) GetNotifySlackOk() (*string, bool) {
	if o == nil || IsNil(o.NotifySlack) {
		return nil, false
	}
	return o.NotifySlack, true
}

// HasNotifySlack returns a boolean if a field has been set.
func (o *WAFConfigUpdate) HasNotifySlack() bool {
	if o != nil && !IsNil(o.NotifySlack) {
		return true
	}

	return false
}

// SetNotifySlack gets a reference to the given string and assigns it to the NotifySlack field.
func (o *WAFConfigUpdate) SetNotifySlack(v string) {
	o.NotifySlack = &v
}

func (o WAFConfigUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WAFConfigUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WAFConfigUpdate) UnmarshalJSON(data []byte) (err error) {
	varWAFConfigUpdate := _WAFConfigUpdate{}

	err = json.Unmarshal(data, &varWAFConfigUpdate)

	if err != nil {
		return err
	}

	*o = WAFConfigUpdate(varWAFConfigUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mode")
		delete(additionalProperties, "paranoia_level")
		delete(additionalProperties, "allow_rules")
		delete(additionalProperties, "allow_ip")
		delete(additionalProperties, "block_ip")
		delete(additionalProperties, "block_ua")
		delete(additionalProperties, "block_referer")
		delete(additionalProperties, "thresholds")
		delete(additionalProperties, "httpbl_enabled")
		delete(additionalProperties, "notify_slack_hits_rpm")
		delete(additionalProperties, "ip_ratelimit_mode")
		delete(additionalProperties, "ip_ratelimit_rps")
		delete(additionalProperties, "ip_ratelimit_cooldown")
		delete(additionalProperties, "request_header_ratelimit_mode")
		delete(additionalProperties, "request_header_name")
		delete(additionalProperties, "request_header_ratelimit_rps")
		delete(additionalProperties, "request_header_ratelimit_cooldown")
		delete(additionalProperties, "waf_ratelimit_mode")
		delete(additionalProperties, "waf_ratelimit_hits")
		delete(additionalProperties, "waf_ratelimit_rps")
		delete(additionalProperties, "waf_ratelimit_cooldown")
		delete(additionalProperties, "notify_email")
		delete(additionalProperties, "notify_slack")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWAFConfigUpdate struct {
	value *WAFConfigUpdate
	isSet bool
}

func (v NullableWAFConfigUpdate) Get() *WAFConfigUpdate {
	return v.value
}

func (v *NullableWAFConfigUpdate) Set(val *WAFConfigUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWAFConfigUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWAFConfigUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWAFConfigUpdate(val *WAFConfigUpdate) *NullableWAFConfigUpdate {
	return &NullableWAFConfigUpdate{value: val, isSet: true}
}

func (v NullableWAFConfigUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWAFConfigUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


