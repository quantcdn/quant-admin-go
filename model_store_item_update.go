/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quantadmingo

import (
	"encoding/json"
)

// checks if the StoreItemUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreItemUpdate{}

// StoreItemUpdate struct for StoreItemUpdate
type StoreItemUpdate struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoreItemUpdate StoreItemUpdate

// NewStoreItemUpdate instantiates a new StoreItemUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreItemUpdate() *StoreItemUpdate {
	this := StoreItemUpdate{}
	return &this
}

// NewStoreItemUpdateWithDefaults instantiates a new StoreItemUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreItemUpdateWithDefaults() *StoreItemUpdate {
	this := StoreItemUpdate{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StoreItemUpdate) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreItemUpdate) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StoreItemUpdate) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StoreItemUpdate) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StoreItemUpdate) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreItemUpdate) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StoreItemUpdate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *StoreItemUpdate) SetValue(v string) {
	o.Value = &v
}

func (o StoreItemUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreItemUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StoreItemUpdate) UnmarshalJSON(data []byte) (err error) {
	varStoreItemUpdate := _StoreItemUpdate{}

	err = json.Unmarshal(data, &varStoreItemUpdate)

	if err != nil {
		return err
	}

	*o = StoreItemUpdate(varStoreItemUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoreItemUpdate struct {
	value *StoreItemUpdate
	isSet bool
}

func (v NullableStoreItemUpdate) Get() *StoreItemUpdate {
	return v.value
}

func (v *NullableStoreItemUpdate) Set(val *StoreItemUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreItemUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreItemUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreItemUpdate(val *StoreItemUpdate) *NullableStoreItemUpdate {
	return &NullableStoreItemUpdate{value: val, isSet: true}
}

func (v NullableStoreItemUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreItemUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


