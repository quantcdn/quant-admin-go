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

// checks if the Organisation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organisation{}

// Organisation struct for Organisation
type Organisation struct {
	Name *string `json:"name,omitempty"`
	MachineName *string `json:"machine_name,omitempty"`
}

// NewOrganisation instantiates a new Organisation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisation() *Organisation {
	this := Organisation{}
	return &this
}

// NewOrganisationWithDefaults instantiates a new Organisation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationWithDefaults() *Organisation {
	this := Organisation{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Organisation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organisation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Organisation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Organisation) SetName(v string) {
	o.Name = &v
}

// GetMachineName returns the MachineName field value if set, zero value otherwise.
func (o *Organisation) GetMachineName() string {
	if o == nil || IsNil(o.MachineName) {
		var ret string
		return ret
	}
	return *o.MachineName
}

// GetMachineNameOk returns a tuple with the MachineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organisation) GetMachineNameOk() (*string, bool) {
	if o == nil || IsNil(o.MachineName) {
		return nil, false
	}
	return o.MachineName, true
}

// HasMachineName returns a boolean if a field has been set.
func (o *Organisation) HasMachineName() bool {
	if o != nil && !IsNil(o.MachineName) {
		return true
	}

	return false
}

// SetMachineName gets a reference to the given string and assigns it to the MachineName field.
func (o *Organisation) SetMachineName(v string) {
	o.MachineName = &v
}

func (o Organisation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organisation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MachineName) {
		toSerialize["machine_name"] = o.MachineName
	}
	return toSerialize, nil
}

type NullableOrganisation struct {
	value *Organisation
	isSet bool
}

func (v NullableOrganisation) Get() *Organisation {
	return v.value
}

func (v *NullableOrganisation) Set(val *Organisation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisation(val *Organisation) *NullableOrganisation {
	return &NullableOrganisation{value: val, isSet: true}
}

func (v NullableOrganisation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


