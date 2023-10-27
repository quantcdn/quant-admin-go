/*
QuantCDN Dashboard API

Provides programmatic interface for projects within QuantCDN

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectResponse{}

// ProjectResponse struct for ProjectResponse
type ProjectResponse struct {
	Data *ProjectResponseData `json:"data,omitempty"`
}

// NewProjectResponse instantiates a new ProjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectResponse() *ProjectResponse {
	this := ProjectResponse{}
	return &this
}

// NewProjectResponseWithDefaults instantiates a new ProjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectResponseWithDefaults() *ProjectResponse {
	this := ProjectResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProjectResponse) GetData() ProjectResponseData {
	if o == nil || IsNil(o.Data) {
		var ret ProjectResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResponse) GetDataOk() (*ProjectResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProjectResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ProjectResponseData and assigns it to the Data field.
func (o *ProjectResponse) SetData(v ProjectResponseData) {
	o.Data = &v
}

func (o ProjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableProjectResponse struct {
	value *ProjectResponse
	isSet bool
}

func (v NullableProjectResponse) Get() *ProjectResponse {
	return v.value
}

func (v *NullableProjectResponse) Set(val *ProjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectResponse(val *ProjectResponse) *NullableProjectResponse {
	return &NullableProjectResponse{value: val, isSet: true}
}

func (v NullableProjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


