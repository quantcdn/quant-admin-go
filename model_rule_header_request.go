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

// checks if the RuleHeaderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleHeaderRequest{}

// RuleHeaderRequest struct for RuleHeaderRequest
type RuleHeaderRequest struct {
	Headers map[string]string `json:"headers"`
	Domain string `json:"domain"`
	Name *string `json:"name,omitempty"`
	Disabled bool `json:"disabled"`
	Country *string `json:"country,omitempty"`
	CountryIs []string `json:"country_is,omitempty"`
	CountryIsNot []string `json:"country_is_not,omitempty"`
	Method *string `json:"method,omitempty"`
	MethodIs []string `json:"method_is,omitempty"`
	MethodIsNot []string `json:"method_is_not,omitempty"`
	Ip *string `json:"ip,omitempty"`
	IpIs []string `json:"ip_is,omitempty"`
	IpIsNot []string `json:"ip_is_not,omitempty"`
}

type _RuleHeaderRequest RuleHeaderRequest

// NewRuleHeaderRequest instantiates a new RuleHeaderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleHeaderRequest(headers map[string]string, domain string, disabled bool) *RuleHeaderRequest {
	this := RuleHeaderRequest{}
	this.Domain = domain
	this.Disabled = disabled
	return &this
}

// NewRuleHeaderRequestWithDefaults instantiates a new RuleHeaderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleHeaderRequestWithDefaults() *RuleHeaderRequest {
	this := RuleHeaderRequest{}
	var domain string = "any"
	this.Domain = domain
	var disabled bool = false
	this.Disabled = disabled
	return &this
}

// GetHeaders returns the Headers field value
func (o *RuleHeaderRequest) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *RuleHeaderRequest) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetDomain returns the Domain field value
func (o *RuleHeaderRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *RuleHeaderRequest) SetDomain(v string) {
	o.Domain = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleHeaderRequest) SetName(v string) {
	o.Name = &v
}

// GetDisabled returns the Disabled field value
func (o *RuleHeaderRequest) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *RuleHeaderRequest) SetDisabled(v bool) {
	o.Disabled = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *RuleHeaderRequest) SetCountry(v string) {
	o.Country = &v
}

// GetCountryIs returns the CountryIs field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetCountryIs() []string {
	if o == nil || IsNil(o.CountryIs) {
		var ret []string
		return ret
	}
	return o.CountryIs
}

// GetCountryIsOk returns a tuple with the CountryIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetCountryIsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIs) {
		return nil, false
	}
	return o.CountryIs, true
}

// HasCountryIs returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasCountryIs() bool {
	if o != nil && !IsNil(o.CountryIs) {
		return true
	}

	return false
}

// SetCountryIs gets a reference to the given []string and assigns it to the CountryIs field.
func (o *RuleHeaderRequest) SetCountryIs(v []string) {
	o.CountryIs = v
}

// GetCountryIsNot returns the CountryIsNot field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetCountryIsNot() []string {
	if o == nil || IsNil(o.CountryIsNot) {
		var ret []string
		return ret
	}
	return o.CountryIsNot
}

// GetCountryIsNotOk returns a tuple with the CountryIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetCountryIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIsNot) {
		return nil, false
	}
	return o.CountryIsNot, true
}

// HasCountryIsNot returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasCountryIsNot() bool {
	if o != nil && !IsNil(o.CountryIsNot) {
		return true
	}

	return false
}

// SetCountryIsNot gets a reference to the given []string and assigns it to the CountryIsNot field.
func (o *RuleHeaderRequest) SetCountryIsNot(v []string) {
	o.CountryIsNot = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RuleHeaderRequest) SetMethod(v string) {
	o.Method = &v
}

// GetMethodIs returns the MethodIs field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetMethodIs() []string {
	if o == nil || IsNil(o.MethodIs) {
		var ret []string
		return ret
	}
	return o.MethodIs
}

// GetMethodIsOk returns a tuple with the MethodIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetMethodIsOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIs) {
		return nil, false
	}
	return o.MethodIs, true
}

// HasMethodIs returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasMethodIs() bool {
	if o != nil && !IsNil(o.MethodIs) {
		return true
	}

	return false
}

// SetMethodIs gets a reference to the given []string and assigns it to the MethodIs field.
func (o *RuleHeaderRequest) SetMethodIs(v []string) {
	o.MethodIs = v
}

// GetMethodIsNot returns the MethodIsNot field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetMethodIsNot() []string {
	if o == nil || IsNil(o.MethodIsNot) {
		var ret []string
		return ret
	}
	return o.MethodIsNot
}

// GetMethodIsNotOk returns a tuple with the MethodIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetMethodIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIsNot) {
		return nil, false
	}
	return o.MethodIsNot, true
}

// HasMethodIsNot returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasMethodIsNot() bool {
	if o != nil && !IsNil(o.MethodIsNot) {
		return true
	}

	return false
}

// SetMethodIsNot gets a reference to the given []string and assigns it to the MethodIsNot field.
func (o *RuleHeaderRequest) SetMethodIsNot(v []string) {
	o.MethodIsNot = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RuleHeaderRequest) SetIp(v string) {
	o.Ip = &v
}

// GetIpIs returns the IpIs field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetIpIs() []string {
	if o == nil || IsNil(o.IpIs) {
		var ret []string
		return ret
	}
	return o.IpIs
}

// GetIpIsOk returns a tuple with the IpIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetIpIsOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIs) {
		return nil, false
	}
	return o.IpIs, true
}

// HasIpIs returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasIpIs() bool {
	if o != nil && !IsNil(o.IpIs) {
		return true
	}

	return false
}

// SetIpIs gets a reference to the given []string and assigns it to the IpIs field.
func (o *RuleHeaderRequest) SetIpIs(v []string) {
	o.IpIs = v
}

// GetIpIsNot returns the IpIsNot field value if set, zero value otherwise.
func (o *RuleHeaderRequest) GetIpIsNot() []string {
	if o == nil || IsNil(o.IpIsNot) {
		var ret []string
		return ret
	}
	return o.IpIsNot
}

// GetIpIsNotOk returns a tuple with the IpIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleHeaderRequest) GetIpIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIsNot) {
		return nil, false
	}
	return o.IpIsNot, true
}

// HasIpIsNot returns a boolean if a field has been set.
func (o *RuleHeaderRequest) HasIpIsNot() bool {
	if o != nil && !IsNil(o.IpIsNot) {
		return true
	}

	return false
}

// SetIpIsNot gets a reference to the given []string and assigns it to the IpIsNot field.
func (o *RuleHeaderRequest) SetIpIsNot(v []string) {
	o.IpIsNot = v
}

func (o RuleHeaderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleHeaderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["headers"] = o.Headers
	toSerialize["domain"] = o.Domain
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["disabled"] = o.Disabled
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
	return toSerialize, nil
}

func (o *RuleHeaderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"headers",
		"domain",
		"disabled",
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

	varRuleHeaderRequest := _RuleHeaderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRuleHeaderRequest)

	if err != nil {
		return err
	}

	*o = RuleHeaderRequest(varRuleHeaderRequest)

	return err
}

type NullableRuleHeaderRequest struct {
	value *RuleHeaderRequest
	isSet bool
}

func (v NullableRuleHeaderRequest) Get() *RuleHeaderRequest {
	return v.value
}

func (v *NullableRuleHeaderRequest) Set(val *RuleHeaderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleHeaderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleHeaderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleHeaderRequest(val *RuleHeaderRequest) *NullableRuleHeaderRequest {
	return &NullableRuleHeaderRequest{value: val, isSet: true}
}

func (v NullableRuleHeaderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleHeaderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


