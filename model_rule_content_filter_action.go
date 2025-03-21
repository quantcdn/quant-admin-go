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

// checks if the RuleContentFilterAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleContentFilterAction{}

// RuleContentFilterAction struct for RuleContentFilterAction
type RuleContentFilterAction struct {
	FnUuid string `json:"fn_uuid"`
	AdditionalProperties map[string]interface{}
}

type _RuleContentFilterAction RuleContentFilterAction

// NewRuleContentFilterAction instantiates a new RuleContentFilterAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleContentFilterAction(fnUuid string) *RuleContentFilterAction {
	this := RuleContentFilterAction{}
	this.FnUuid = fnUuid
	return &this
}

// NewRuleContentFilterActionWithDefaults instantiates a new RuleContentFilterAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleContentFilterActionWithDefaults() *RuleContentFilterAction {
	this := RuleContentFilterAction{}
	return &this
}

// GetFnUuid returns the FnUuid field value
func (o *RuleContentFilterAction) GetFnUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FnUuid
}

// GetFnUuidOk returns a tuple with the FnUuid field value
// and a boolean to check if the value has been set.
func (o *RuleContentFilterAction) GetFnUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FnUuid, true
}

// SetFnUuid sets field value
func (o *RuleContentFilterAction) SetFnUuid(v string) {
	o.FnUuid = v
}

func (o RuleContentFilterAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleContentFilterAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fn_uuid"] = o.FnUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleContentFilterAction) UnmarshalJSON(data []byte) (err error) {
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

	varRuleContentFilterAction := _RuleContentFilterAction{}

	err = json.Unmarshal(data, &varRuleContentFilterAction)

	if err != nil {
		return err
	}

	*o = RuleContentFilterAction(varRuleContentFilterAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fn_uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleContentFilterAction struct {
	value *RuleContentFilterAction
	isSet bool
}

func (v NullableRuleContentFilterAction) Get() *RuleContentFilterAction {
	return v.value
}

func (v *NullableRuleContentFilterAction) Set(val *RuleContentFilterAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleContentFilterAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleContentFilterAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleContentFilterAction(val *RuleContentFilterAction) *NullableRuleContentFilterAction {
	return &NullableRuleContentFilterAction{value: val, isSet: true}
}

func (v NullableRuleContentFilterAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleContentFilterAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


