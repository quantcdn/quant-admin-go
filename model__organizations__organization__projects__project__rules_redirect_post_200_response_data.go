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

// checks if the OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData{}

// OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData struct for OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData
type OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData struct {
	Rules []Rule `json:"rules,omitempty"`
}

// NewOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData instantiates a new OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData() *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData {
	this := OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData{}
	return &this
}

// NewOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseDataWithDefaults instantiates a new OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseDataWithDefaults() *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData {
	this := OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) GetRules() []Rule {
	if o == nil || IsNil(o.Rules) {
		var ret []Rule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) GetRulesOk() ([]Rule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []Rule and assigns it to the Rules field.
func (o *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) SetRules(v []Rule) {
	o.Rules = v
}

func (o OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData struct {
	value *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData
	isSet bool
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) Get() *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData {
	return v.value
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) Set(val *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData(val *OrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) *NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData {
	return &NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesRedirectPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


