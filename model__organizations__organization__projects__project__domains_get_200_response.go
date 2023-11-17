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

// checks if the OrganizationsOrganizationProjectsProjectDomainsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrganizationProjectsProjectDomainsGet200Response{}

// OrganizationsOrganizationProjectsProjectDomainsGet200Response struct for OrganizationsOrganizationProjectsProjectDomainsGet200Response
type OrganizationsOrganizationProjectsProjectDomainsGet200Response struct {
	Data *OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData `json:"data,omitempty"`
}

// NewOrganizationsOrganizationProjectsProjectDomainsGet200Response instantiates a new OrganizationsOrganizationProjectsProjectDomainsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationProjectsProjectDomainsGet200Response() *OrganizationsOrganizationProjectsProjectDomainsGet200Response {
	this := OrganizationsOrganizationProjectsProjectDomainsGet200Response{}
	return &this
}

// NewOrganizationsOrganizationProjectsProjectDomainsGet200ResponseWithDefaults instantiates a new OrganizationsOrganizationProjectsProjectDomainsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationProjectsProjectDomainsGet200ResponseWithDefaults() *OrganizationsOrganizationProjectsProjectDomainsGet200Response {
	this := OrganizationsOrganizationProjectsProjectDomainsGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrganizationsOrganizationProjectsProjectDomainsGet200Response) GetData() OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationProjectsProjectDomainsGet200Response) GetDataOk() (*OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrganizationsOrganizationProjectsProjectDomainsGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData and assigns it to the Data field.
func (o *OrganizationsOrganizationProjectsProjectDomainsGet200Response) SetData(v OrganizationsOrganizationProjectsProjectDomainsGet200ResponseData) {
	o.Data = &v
}

func (o OrganizationsOrganizationProjectsProjectDomainsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrganizationProjectsProjectDomainsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response struct {
	value *OrganizationsOrganizationProjectsProjectDomainsGet200Response
	isSet bool
}

func (v NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response) Get() *OrganizationsOrganizationProjectsProjectDomainsGet200Response {
	return v.value
}

func (v *NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response) Set(val *OrganizationsOrganizationProjectsProjectDomainsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationProjectsProjectDomainsGet200Response(val *OrganizationsOrganizationProjectsProjectDomainsGet200Response) *NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response {
	return &NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationProjectsProjectDomainsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


