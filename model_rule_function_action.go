/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quantadmingo

import (
	"encoding/json"
	"fmt"
)

// checks if the RuleFunctionAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleFunctionAction{}

// RuleFunctionAction struct for RuleFunctionAction
type RuleFunctionAction struct {
	FnUuid string `json:"fn_uuid"`
	AdditionalProperties map[string]interface{}
}

type _RuleFunctionAction RuleFunctionAction

// NewRuleFunctionAction instantiates a new RuleFunctionAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleFunctionAction(fnUuid string) *RuleFunctionAction {
	this := RuleFunctionAction{}
	this.FnUuid = fnUuid
	return &this
}

// NewRuleFunctionActionWithDefaults instantiates a new RuleFunctionAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleFunctionActionWithDefaults() *RuleFunctionAction {
	this := RuleFunctionAction{}
	return &this
}

// GetFnUuid returns the FnUuid field value
func (o *RuleFunctionAction) GetFnUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FnUuid
}

// GetFnUuidOk returns a tuple with the FnUuid field value
// and a boolean to check if the value has been set.
func (o *RuleFunctionAction) GetFnUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FnUuid, true
}

// SetFnUuid sets field value
func (o *RuleFunctionAction) SetFnUuid(v string) {
	o.FnUuid = v
}

func (o RuleFunctionAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleFunctionAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fn_uuid"] = o.FnUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleFunctionAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fn_uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRuleFunctionAction := _RuleFunctionAction{}

	err = json.Unmarshal(data, &varRuleFunctionAction)

	if err != nil {
		return err
	}

	*o = RuleFunctionAction(varRuleFunctionAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fn_uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleFunctionAction struct {
	value *RuleFunctionAction
	isSet bool
}

func (v NullableRuleFunctionAction) Get() *RuleFunctionAction {
	return v.value
}

func (v *NullableRuleFunctionAction) Set(val *RuleFunctionAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleFunctionAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleFunctionAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleFunctionAction(val *RuleFunctionAction) *NullableRuleFunctionAction {
	return &NullableRuleFunctionAction{value: val, isSet: true}
}

func (v NullableRuleFunctionAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleFunctionAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


