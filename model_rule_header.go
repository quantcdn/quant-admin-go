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

// checks if the RuleHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleHeader{}

// RuleHeader struct for RuleHeader
type RuleHeader struct {
	Kind *string `json:"kind,omitempty"`
	ActionConfig RuleHeaderAction `json:"action_config"`
	Name *string `json:"name,omitempty"`
	Uuid string `json:"uuid"`
	RuleId *string `json:"rule_id,omitempty"`
	Url []string `json:"url,omitempty"`
	Domain []string `json:"domain,omitempty"`
	Disabled bool `json:"disabled"`
	OnlyWithCookie *string `json:"only_with_cookie,omitempty"`
	Method *string `json:"method,omitempty"`
	MethodIs []string `json:"method_is,omitempty"`
	MethodIsNot []string `json:"method_is_not,omitempty"`
	Ip *string `json:"ip,omitempty"`
	IpIs []string `json:"ip_is,omitempty"`
	IpIsNot []string `json:"ip_is_not,omitempty"`
	Country *string `json:"country,omitempty"`
	CountryIs []string `json:"country_is,omitempty"`
	CountryIsNot []string `json:"country_is_not,omitempty"`
	Action string `json:"action"`
	AdditionalProperties map[string]interface{}
}

type _RuleHeader RuleHeader

// NewRuleHeader instantiates a new RuleHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleHeader(actionConfig RuleHeaderAction, uuid string, disabled bool, action string) *RuleHeader {
	this := RuleHeader{}
	this.Uuid = uuid
	this.Disabled = disabled
	this.Action = action
	return &this
}

// NewRuleHeaderWithDefaults instantiates a new RuleHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleHeaderWithDefaults() *RuleHeader {
	this := RuleHeader{}
	var kind string = "rule_header"
	this.Kind = &kind
	var disabled bool = false
	this.Disabled = disabled
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RuleHeader) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RuleHeader) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RuleHeader) SetKind(v string) {
	o.Kind = &v
}

// GetActionConfig returns the ActionConfig field value
func (o *RuleHeader) GetActionConfig() RuleHeaderAction {
	if o == nil {
		var ret RuleHeaderAction
		return ret
	}

	return o.ActionConfig
}

// GetActionConfigOk returns a tuple with the ActionConfig field value
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetActionConfigOk() (*RuleHeaderAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionConfig, true
}

// SetActionConfig sets field value
func (o *RuleHeader) SetActionConfig(v RuleHeaderAction) {
	o.ActionConfig = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleHeader) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleHeader) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleHeader) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value
func (o *RuleHeader) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *RuleHeader) SetUuid(v string) {
	o.Uuid = v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *RuleHeader) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *RuleHeader) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *RuleHeader) SetRuleId(v string) {
	o.RuleId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RuleHeader) GetUrl() []string {
	if o == nil || IsNil(o.Url) {
		var ret []string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetUrlOk() ([]string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RuleHeader) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given []string and assigns it to the Url field.
func (o *RuleHeader) SetUrl(v []string) {
	o.Url = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *RuleHeader) GetDomain() []string {
	if o == nil || IsNil(o.Domain) {
		var ret []string
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetDomainOk() ([]string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *RuleHeader) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given []string and assigns it to the Domain field.
func (o *RuleHeader) SetDomain(v []string) {
	o.Domain = v
}

// GetDisabled returns the Disabled field value
func (o *RuleHeader) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *RuleHeader) SetDisabled(v bool) {
	o.Disabled = v
}

// GetOnlyWithCookie returns the OnlyWithCookie field value if set, zero value otherwise.
func (o *RuleHeader) GetOnlyWithCookie() string {
	if o == nil || IsNil(o.OnlyWithCookie) {
		var ret string
		return ret
	}
	return *o.OnlyWithCookie
}

// GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetOnlyWithCookieOk() (*string, bool) {
	if o == nil || IsNil(o.OnlyWithCookie) {
		return nil, false
	}
	return o.OnlyWithCookie, true
}

// HasOnlyWithCookie returns a boolean if a field has been set.
func (o *RuleHeader) HasOnlyWithCookie() bool {
	if o != nil && !IsNil(o.OnlyWithCookie) {
		return true
	}

	return false
}

// SetOnlyWithCookie gets a reference to the given string and assigns it to the OnlyWithCookie field.
func (o *RuleHeader) SetOnlyWithCookie(v string) {
	o.OnlyWithCookie = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RuleHeader) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RuleHeader) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RuleHeader) SetMethod(v string) {
	o.Method = &v
}

// GetMethodIs returns the MethodIs field value if set, zero value otherwise.
func (o *RuleHeader) GetMethodIs() []string {
	if o == nil || IsNil(o.MethodIs) {
		var ret []string
		return ret
	}
	return o.MethodIs
}

// GetMethodIsOk returns a tuple with the MethodIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetMethodIsOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIs) {
		return nil, false
	}
	return o.MethodIs, true
}

// HasMethodIs returns a boolean if a field has been set.
func (o *RuleHeader) HasMethodIs() bool {
	if o != nil && !IsNil(o.MethodIs) {
		return true
	}

	return false
}

// SetMethodIs gets a reference to the given []string and assigns it to the MethodIs field.
func (o *RuleHeader) SetMethodIs(v []string) {
	o.MethodIs = v
}

// GetMethodIsNot returns the MethodIsNot field value if set, zero value otherwise.
func (o *RuleHeader) GetMethodIsNot() []string {
	if o == nil || IsNil(o.MethodIsNot) {
		var ret []string
		return ret
	}
	return o.MethodIsNot
}

// GetMethodIsNotOk returns a tuple with the MethodIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetMethodIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIsNot) {
		return nil, false
	}
	return o.MethodIsNot, true
}

// HasMethodIsNot returns a boolean if a field has been set.
func (o *RuleHeader) HasMethodIsNot() bool {
	if o != nil && !IsNil(o.MethodIsNot) {
		return true
	}

	return false
}

// SetMethodIsNot gets a reference to the given []string and assigns it to the MethodIsNot field.
func (o *RuleHeader) SetMethodIsNot(v []string) {
	o.MethodIsNot = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RuleHeader) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RuleHeader) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RuleHeader) SetIp(v string) {
	o.Ip = &v
}

// GetIpIs returns the IpIs field value if set, zero value otherwise.
func (o *RuleHeader) GetIpIs() []string {
	if o == nil || IsNil(o.IpIs) {
		var ret []string
		return ret
	}
	return o.IpIs
}

// GetIpIsOk returns a tuple with the IpIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetIpIsOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIs) {
		return nil, false
	}
	return o.IpIs, true
}

// HasIpIs returns a boolean if a field has been set.
func (o *RuleHeader) HasIpIs() bool {
	if o != nil && !IsNil(o.IpIs) {
		return true
	}

	return false
}

// SetIpIs gets a reference to the given []string and assigns it to the IpIs field.
func (o *RuleHeader) SetIpIs(v []string) {
	o.IpIs = v
}

// GetIpIsNot returns the IpIsNot field value if set, zero value otherwise.
func (o *RuleHeader) GetIpIsNot() []string {
	if o == nil || IsNil(o.IpIsNot) {
		var ret []string
		return ret
	}
	return o.IpIsNot
}

// GetIpIsNotOk returns a tuple with the IpIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetIpIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIsNot) {
		return nil, false
	}
	return o.IpIsNot, true
}

// HasIpIsNot returns a boolean if a field has been set.
func (o *RuleHeader) HasIpIsNot() bool {
	if o != nil && !IsNil(o.IpIsNot) {
		return true
	}

	return false
}

// SetIpIsNot gets a reference to the given []string and assigns it to the IpIsNot field.
func (o *RuleHeader) SetIpIsNot(v []string) {
	o.IpIsNot = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *RuleHeader) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *RuleHeader) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *RuleHeader) SetCountry(v string) {
	o.Country = &v
}

// GetCountryIs returns the CountryIs field value if set, zero value otherwise.
func (o *RuleHeader) GetCountryIs() []string {
	if o == nil || IsNil(o.CountryIs) {
		var ret []string
		return ret
	}
	return o.CountryIs
}

// GetCountryIsOk returns a tuple with the CountryIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetCountryIsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIs) {
		return nil, false
	}
	return o.CountryIs, true
}

// HasCountryIs returns a boolean if a field has been set.
func (o *RuleHeader) HasCountryIs() bool {
	if o != nil && !IsNil(o.CountryIs) {
		return true
	}

	return false
}

// SetCountryIs gets a reference to the given []string and assigns it to the CountryIs field.
func (o *RuleHeader) SetCountryIs(v []string) {
	o.CountryIs = v
}

// GetCountryIsNot returns the CountryIsNot field value if set, zero value otherwise.
func (o *RuleHeader) GetCountryIsNot() []string {
	if o == nil || IsNil(o.CountryIsNot) {
		var ret []string
		return ret
	}
	return o.CountryIsNot
}

// GetCountryIsNotOk returns a tuple with the CountryIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetCountryIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIsNot) {
		return nil, false
	}
	return o.CountryIsNot, true
}

// HasCountryIsNot returns a boolean if a field has been set.
func (o *RuleHeader) HasCountryIsNot() bool {
	if o != nil && !IsNil(o.CountryIsNot) {
		return true
	}

	return false
}

// SetCountryIsNot gets a reference to the given []string and assigns it to the CountryIsNot field.
func (o *RuleHeader) SetCountryIsNot(v []string) {
	o.CountryIsNot = v
}

// GetAction returns the Action field value
func (o *RuleHeader) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RuleHeader) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *RuleHeader) SetAction(v string) {
	o.Action = v
}

func (o RuleHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["action_config"] = o.ActionConfig
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	toSerialize["disabled"] = o.Disabled
	if !IsNil(o.OnlyWithCookie) {
		toSerialize["only_with_cookie"] = o.OnlyWithCookie
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
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CountryIs) {
		toSerialize["country_is"] = o.CountryIs
	}
	if !IsNil(o.CountryIsNot) {
		toSerialize["country_is_not"] = o.CountryIsNot
	}
	toSerialize["action"] = o.Action

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleHeader) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action_config",
		"uuid",
		"disabled",
		"action",
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

	varRuleHeader := _RuleHeader{}

	err = json.Unmarshal(data, &varRuleHeader)

	if err != nil {
		return err
	}

	*o = RuleHeader(varRuleHeader)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "kind")
		delete(additionalProperties, "action_config")
		delete(additionalProperties, "name")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "rule_id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "only_with_cookie")
		delete(additionalProperties, "method")
		delete(additionalProperties, "method_is")
		delete(additionalProperties, "method_is_not")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "ip_is")
		delete(additionalProperties, "ip_is_not")
		delete(additionalProperties, "country")
		delete(additionalProperties, "country_is")
		delete(additionalProperties, "country_is_not")
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleHeader struct {
	value *RuleHeader
	isSet bool
}

func (v NullableRuleHeader) Get() *RuleHeader {
	return v.value
}

func (v *NullableRuleHeader) Set(val *RuleHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleHeader(val *RuleHeader) *NullableRuleHeader {
	return &NullableRuleHeader{value: val, isSet: true}
}

func (v NullableRuleHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


