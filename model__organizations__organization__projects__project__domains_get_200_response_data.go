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

// checks if the OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData{}

// OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData struct for OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData
type OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData struct {
	Organizations []Domain `json:"organizations,omitempty"`
}

// NewOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData instantiates a new OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData() *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData {
	this := OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData{}
	return &this
}

// NewOrganizationsOrganizationProjectsProjectDomainsGet200ResponseDataWithDefaults instantiates a new OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationProjectsProjectDomainsGet200ResponseDataWithDefaults() *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData {
	this := OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) GetOrganizations() []Domain {
	if o == nil || IsNil(o.Organizations) {
		var ret []Domain
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) GetOrganizationsOk() ([]Domain, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []Domain and assigns it to the Organizations field.
func (o *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) SetOrganizations(v []Domain) {
	o.Organizations = v
}

func (o OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	return toSerialize, nil
}

type NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData struct {
	value *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData
	isSet bool
}

func (v NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) Get() *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData {
	return v.value
}

func (v *NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) Set(val *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData(val *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) *NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData {
	return &NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


