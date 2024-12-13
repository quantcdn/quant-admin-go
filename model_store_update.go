/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StoreUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreUpdate{}

// StoreUpdate struct for StoreUpdate
type StoreUpdate struct {
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreUpdate StoreUpdate

// NewStoreUpdate instantiates a new StoreUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreUpdate() *StoreUpdate {
	this := StoreUpdate{}
	return &this
}

// NewStoreUpdateWithDefaults instantiates a new StoreUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreUpdateWithDefaults() *StoreUpdate {
	this := StoreUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StoreUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StoreUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StoreUpdate) SetName(v string) {
	o.Name = &v
}

func (o StoreUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoreUpdate) UnmarshalJSON(data []byte) (err error) {
	varStoreUpdate := _StoreUpdate{}

	err = json.Unmarshal(data, &varStoreUpdate)

	if err != nil {
		return err
	}

	*o = StoreUpdate(varStoreUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreUpdate struct {
	value *StoreUpdate
	isSet bool
}

func (v NullableStoreUpdate) Get() *StoreUpdate {
	return v.value
}

func (v *NullableStoreUpdate) Set(val *StoreUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreUpdate(val *StoreUpdate) *NullableStoreUpdate {
	return &NullableStoreUpdate{value: val, isSet: true}
}

func (v NullableStoreUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

