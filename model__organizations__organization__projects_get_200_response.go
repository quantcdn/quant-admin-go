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

// checks if the OrganizationsOrganizationProjectsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrganizationProjectsGet200Response{}

// OrganizationsOrganizationProjectsGet200Response struct for OrganizationsOrganizationProjectsGet200Response
type OrganizationsOrganizationProjectsGet200Response struct {
	Data *OrganizationsOrganizationProjectsGet200ResponseData `json:"data,omitempty"`
}

// NewOrganizationsOrganizationProjectsGet200Response instantiates a new OrganizationsOrganizationProjectsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationProjectsGet200Response() *OrganizationsOrganizationProjectsGet200Response {
	this := OrganizationsOrganizationProjectsGet200Response{}
	return &this
}

// NewOrganizationsOrganizationProjectsGet200ResponseWithDefaults instantiates a new OrganizationsOrganizationProjectsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationProjectsGet200ResponseWithDefaults() *OrganizationsOrganizationProjectsGet200Response {
	this := OrganizationsOrganizationProjectsGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrganizationsOrganizationProjectsGet200Response) GetData() OrganizationsOrganizationProjectsGet200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret OrganizationsOrganizationProjectsGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationProjectsGet200Response) GetDataOk() (*OrganizationsOrganizationProjectsGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrganizationsOrganizationProjectsGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OrganizationsOrganizationProjectsGet200ResponseData and assigns it to the Data field.
func (o *OrganizationsOrganizationProjectsGet200Response) SetData(v OrganizationsOrganizationProjectsGet200ResponseData) {
	o.Data = &v
}

func (o OrganizationsOrganizationProjectsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrganizationProjectsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrganizationsOrganizationProjectsGet200Response struct {
	value *OrganizationsOrganizationProjectsGet200Response
	isSet bool
}

func (v NullableOrganizationsOrganizationProjectsGet200Response) Get() *OrganizationsOrganizationProjectsGet200Response {
	return v.value
}

func (v *NullableOrganizationsOrganizationProjectsGet200Response) Set(val *OrganizationsOrganizationProjectsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationProjectsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationProjectsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationProjectsGet200Response(val *OrganizationsOrganizationProjectsGet200Response) *NullableOrganizationsOrganizationProjectsGet200Response {
	return &NullableOrganizationsOrganizationProjectsGet200Response{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationProjectsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationProjectsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

