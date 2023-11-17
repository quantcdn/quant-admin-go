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

// checks if the OrganizationsOrganizationProjectsProjectRulesGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrganizationProjectsProjectRulesGet200ResponseData{}

// OrganizationsOrganizationProjectsProjectRulesGet200ResponseData struct for OrganizationsOrganizationProjectsProjectRulesGet200ResponseData
type OrganizationsOrganizationProjectsProjectRulesGet200ResponseData struct {
	Rules []Rule `json:"rules,omitempty"`
}

// NewOrganizationsOrganizationProjectsProjectRulesGet200ResponseData instantiates a new OrganizationsOrganizationProjectsProjectRulesGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationProjectsProjectRulesGet200ResponseData() *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData {
	this := OrganizationsOrganizationProjectsProjectRulesGet200ResponseData{}
	return &this
}

// NewOrganizationsOrganizationProjectsProjectRulesGet200ResponseDataWithDefaults instantiates a new OrganizationsOrganizationProjectsProjectRulesGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationProjectsProjectRulesGet200ResponseDataWithDefaults() *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData {
	this := OrganizationsOrganizationProjectsProjectRulesGet200ResponseData{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData) GetRules() []Rule {
	if o == nil || IsNil(o.Rules) {
		var ret []Rule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData) GetRulesOk() ([]Rule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []Rule and assigns it to the Rules field.
func (o *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData) SetRules(v []Rule) {
	o.Rules = v
}

func (o OrganizationsOrganizationProjectsProjectRulesGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrganizationProjectsProjectRulesGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData struct {
	value *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData
	isSet bool
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData) Get() *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData {
	return v.value
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData) Set(val *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData(val *OrganizationsOrganizationProjectsProjectRulesGet200ResponseData) *NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData {
	return &NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationProjectsProjectRulesGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


