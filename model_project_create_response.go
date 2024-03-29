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

// checks if the ProjectCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCreateResponse{}

// ProjectCreateResponse struct for ProjectCreateResponse
type ProjectCreateResponse struct {
	Data *ProjectCreateResponseData `json:"data,omitempty"`
}

// NewProjectCreateResponse instantiates a new ProjectCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreateResponse() *ProjectCreateResponse {
	this := ProjectCreateResponse{}
	return &this
}

// NewProjectCreateResponseWithDefaults instantiates a new ProjectCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreateResponseWithDefaults() *ProjectCreateResponse {
	this := ProjectCreateResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProjectCreateResponse) GetData() ProjectCreateResponseData {
	if o == nil || IsNil(o.Data) {
		var ret ProjectCreateResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreateResponse) GetDataOk() (*ProjectCreateResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProjectCreateResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ProjectCreateResponseData and assigns it to the Data field.
func (o *ProjectCreateResponse) SetData(v ProjectCreateResponseData) {
	o.Data = &v
}

func (o ProjectCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableProjectCreateResponse struct {
	value *ProjectCreateResponse
	isSet bool
}

func (v NullableProjectCreateResponse) Get() *ProjectCreateResponse {
	return v.value
}

func (v *NullableProjectCreateResponse) Set(val *ProjectCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreateResponse(val *ProjectCreateResponse) *NullableProjectCreateResponse {
	return &NullableProjectCreateResponse{value: val, isSet: true}
}

func (v NullableProjectCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


