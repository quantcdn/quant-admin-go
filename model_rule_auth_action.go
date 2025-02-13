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

// checks if the RuleAuthAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleAuthAction{}

// RuleAuthAction struct for RuleAuthAction
type RuleAuthAction struct {
	AuthUser string `json:"auth_user"`
	AuthPass string `json:"auth_pass"`
	AdditionalProperties map[string]interface{}
}

type _RuleAuthAction RuleAuthAction

// NewRuleAuthAction instantiates a new RuleAuthAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleAuthAction(authUser string, authPass string) *RuleAuthAction {
	this := RuleAuthAction{}
	this.AuthUser = authUser
	this.AuthPass = authPass
	return &this
}

// NewRuleAuthActionWithDefaults instantiates a new RuleAuthAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleAuthActionWithDefaults() *RuleAuthAction {
	this := RuleAuthAction{}
	return &this
}

// GetAuthUser returns the AuthUser field value
func (o *RuleAuthAction) GetAuthUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthUser
}

// GetAuthUserOk returns a tuple with the AuthUser field value
// and a boolean to check if the value has been set.
func (o *RuleAuthAction) GetAuthUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthUser, true
}

// SetAuthUser sets field value
func (o *RuleAuthAction) SetAuthUser(v string) {
	o.AuthUser = v
}

// GetAuthPass returns the AuthPass field value
func (o *RuleAuthAction) GetAuthPass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthPass
}

// GetAuthPassOk returns a tuple with the AuthPass field value
// and a boolean to check if the value has been set.
func (o *RuleAuthAction) GetAuthPassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthPass, true
}

// SetAuthPass sets field value
func (o *RuleAuthAction) SetAuthPass(v string) {
	o.AuthPass = v
}

func (o RuleAuthAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleAuthAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth_user"] = o.AuthUser
	toSerialize["auth_pass"] = o.AuthPass

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleAuthAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auth_user",
		"auth_pass",
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

	varRuleAuthAction := _RuleAuthAction{}

	err = json.Unmarshal(data, &varRuleAuthAction)

	if err != nil {
		return err
	}

	*o = RuleAuthAction(varRuleAuthAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auth_user")
		delete(additionalProperties, "auth_pass")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleAuthAction struct {
	value *RuleAuthAction
	isSet bool
}

func (v NullableRuleAuthAction) Get() *RuleAuthAction {
	return v.value
}

func (v *NullableRuleAuthAction) Set(val *RuleAuthAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleAuthAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleAuthAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleAuthAction(val *RuleAuthAction) *NullableRuleAuthAction {
	return &NullableRuleAuthAction{value: val, isSet: true}
}

func (v NullableRuleAuthAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleAuthAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


