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

// checks if the ProjectResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectResponseData{}

// ProjectResponseData struct for ProjectResponseData
type ProjectResponseData struct {
	Project *Project `json:"project,omitempty"`
}

// NewProjectResponseData instantiates a new ProjectResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectResponseData() *ProjectResponseData {
	this := ProjectResponseData{}
	return &this
}

// NewProjectResponseDataWithDefaults instantiates a new ProjectResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectResponseDataWithDefaults() *ProjectResponseData {
	this := ProjectResponseData{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectResponseData) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResponseData) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectResponseData) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ProjectResponseData) SetProject(v Project) {
	o.Project = &v
}

func (o ProjectResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	return toSerialize, nil
}

type NullableProjectResponseData struct {
	value *ProjectResponseData
	isSet bool
}

func (v NullableProjectResponseData) Get() *ProjectResponseData {
	return v.value
}

func (v *NullableProjectResponseData) Set(val *ProjectResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectResponseData(val *ProjectResponseData) *NullableProjectResponseData {
	return &NullableProjectResponseData{value: val, isSet: true}
}

func (v NullableProjectResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


