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

// checks if the OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig{}

// OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig struct for OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig
type OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig struct {
	Mode *string `json:"mode,omitempty"`
}

// NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig instantiates a new OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig() *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig {
	this := OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig{}
	return &this
}

// NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfigWithDefaults instantiates a new OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfigWithDefaults() *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig {
	this := OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) SetMode(v string) {
	o.Mode = &v
}

func (o OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig struct {
	value *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig
	isSet bool
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) Get() *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig {
	return v.value
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) Set(val *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig(val *OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) *NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig {
	return &NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequestWafConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

