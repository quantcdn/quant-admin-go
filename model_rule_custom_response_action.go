/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the RuleCustomResponseAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleCustomResponseAction{}

// RuleCustomResponseAction struct for RuleCustomResponseAction
type RuleCustomResponseAction struct {
	CustomResponseBody string `json:"custom_response_body"`
	CustomResponseStatusCode int32 `json:"custom_response_status_code"`
	AdditionalProperties map[string]interface{}
}

type _RuleCustomResponseAction RuleCustomResponseAction

// NewRuleCustomResponseAction instantiates a new RuleCustomResponseAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleCustomResponseAction(customResponseBody string, customResponseStatusCode int32) *RuleCustomResponseAction {
	this := RuleCustomResponseAction{}
	this.CustomResponseBody = customResponseBody
	this.CustomResponseStatusCode = customResponseStatusCode
	return &this
}

// NewRuleCustomResponseActionWithDefaults instantiates a new RuleCustomResponseAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleCustomResponseActionWithDefaults() *RuleCustomResponseAction {
	this := RuleCustomResponseAction{}
	var customResponseStatusCode int32 = 200
	this.CustomResponseStatusCode = customResponseStatusCode
	return &this
}

// GetCustomResponseBody returns the CustomResponseBody field value
func (o *RuleCustomResponseAction) GetCustomResponseBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomResponseBody
}

// GetCustomResponseBodyOk returns a tuple with the CustomResponseBody field value
// and a boolean to check if the value has been set.
func (o *RuleCustomResponseAction) GetCustomResponseBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomResponseBody, true
}

// SetCustomResponseBody sets field value
func (o *RuleCustomResponseAction) SetCustomResponseBody(v string) {
	o.CustomResponseBody = v
}

// GetCustomResponseStatusCode returns the CustomResponseStatusCode field value
func (o *RuleCustomResponseAction) GetCustomResponseStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomResponseStatusCode
}

// GetCustomResponseStatusCodeOk returns a tuple with the CustomResponseStatusCode field value
// and a boolean to check if the value has been set.
func (o *RuleCustomResponseAction) GetCustomResponseStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomResponseStatusCode, true
}

// SetCustomResponseStatusCode sets field value
func (o *RuleCustomResponseAction) SetCustomResponseStatusCode(v int32) {
	o.CustomResponseStatusCode = v
}

func (o RuleCustomResponseAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleCustomResponseAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["custom_response_body"] = o.CustomResponseBody
	toSerialize["custom_response_status_code"] = o.CustomResponseStatusCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleCustomResponseAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"custom_response_body",
		"custom_response_status_code",
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

	varRuleCustomResponseAction := _RuleCustomResponseAction{}

	err = json.Unmarshal(data, &varRuleCustomResponseAction)

	if err != nil {
		return err
	}

	*o = RuleCustomResponseAction(varRuleCustomResponseAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "custom_response_body")
		delete(additionalProperties, "custom_response_status_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleCustomResponseAction struct {
	value *RuleCustomResponseAction
	isSet bool
}

func (v NullableRuleCustomResponseAction) Get() *RuleCustomResponseAction {
	return v.value
}

func (v *NullableRuleCustomResponseAction) Set(val *RuleCustomResponseAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleCustomResponseAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleCustomResponseAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleCustomResponseAction(val *RuleCustomResponseAction) *NullableRuleCustomResponseAction {
	return &NullableRuleCustomResponseAction{value: val, isSet: true}
}

func (v NullableRuleCustomResponseAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleCustomResponseAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

