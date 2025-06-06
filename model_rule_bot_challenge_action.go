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

// checks if the RuleBotChallengeAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleBotChallengeAction{}

// RuleBotChallengeAction struct for RuleBotChallengeAction
type RuleBotChallengeAction struct {
	RobotChallengeType string `json:"robot_challenge_type"`
	RobotChallengeVerificationTtl int32 `json:"robot_challenge_verification_ttl"`
	RobotChallengeChallengeTtl int32 `json:"robot_challenge_challenge_ttl"`
	AdditionalProperties map[string]interface{}
}

type _RuleBotChallengeAction RuleBotChallengeAction

// NewRuleBotChallengeAction instantiates a new RuleBotChallengeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleBotChallengeAction(robotChallengeType string, robotChallengeVerificationTtl int32, robotChallengeChallengeTtl int32) *RuleBotChallengeAction {
	this := RuleBotChallengeAction{}
	this.RobotChallengeType = robotChallengeType
	this.RobotChallengeVerificationTtl = robotChallengeVerificationTtl
	this.RobotChallengeChallengeTtl = robotChallengeChallengeTtl
	return &this
}

// NewRuleBotChallengeActionWithDefaults instantiates a new RuleBotChallengeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleBotChallengeActionWithDefaults() *RuleBotChallengeAction {
	this := RuleBotChallengeAction{}
	var robotChallengeVerificationTtl int32 = 10800
	this.RobotChallengeVerificationTtl = robotChallengeVerificationTtl
	var robotChallengeChallengeTtl int32 = 30
	this.RobotChallengeChallengeTtl = robotChallengeChallengeTtl
	return &this
}

// GetRobotChallengeType returns the RobotChallengeType field value
func (o *RuleBotChallengeAction) GetRobotChallengeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RobotChallengeType
}

// GetRobotChallengeTypeOk returns a tuple with the RobotChallengeType field value
// and a boolean to check if the value has been set.
func (o *RuleBotChallengeAction) GetRobotChallengeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RobotChallengeType, true
}

// SetRobotChallengeType sets field value
func (o *RuleBotChallengeAction) SetRobotChallengeType(v string) {
	o.RobotChallengeType = v
}

// GetRobotChallengeVerificationTtl returns the RobotChallengeVerificationTtl field value
func (o *RuleBotChallengeAction) GetRobotChallengeVerificationTtl() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RobotChallengeVerificationTtl
}

// GetRobotChallengeVerificationTtlOk returns a tuple with the RobotChallengeVerificationTtl field value
// and a boolean to check if the value has been set.
func (o *RuleBotChallengeAction) GetRobotChallengeVerificationTtlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RobotChallengeVerificationTtl, true
}

// SetRobotChallengeVerificationTtl sets field value
func (o *RuleBotChallengeAction) SetRobotChallengeVerificationTtl(v int32) {
	o.RobotChallengeVerificationTtl = v
}

// GetRobotChallengeChallengeTtl returns the RobotChallengeChallengeTtl field value
func (o *RuleBotChallengeAction) GetRobotChallengeChallengeTtl() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RobotChallengeChallengeTtl
}

// GetRobotChallengeChallengeTtlOk returns a tuple with the RobotChallengeChallengeTtl field value
// and a boolean to check if the value has been set.
func (o *RuleBotChallengeAction) GetRobotChallengeChallengeTtlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RobotChallengeChallengeTtl, true
}

// SetRobotChallengeChallengeTtl sets field value
func (o *RuleBotChallengeAction) SetRobotChallengeChallengeTtl(v int32) {
	o.RobotChallengeChallengeTtl = v
}

func (o RuleBotChallengeAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleBotChallengeAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["robot_challenge_type"] = o.RobotChallengeType
	toSerialize["robot_challenge_verification_ttl"] = o.RobotChallengeVerificationTtl
	toSerialize["robot_challenge_challenge_ttl"] = o.RobotChallengeChallengeTtl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleBotChallengeAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"robot_challenge_type",
		"robot_challenge_verification_ttl",
		"robot_challenge_challenge_ttl",
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

	varRuleBotChallengeAction := _RuleBotChallengeAction{}

	err = json.Unmarshal(data, &varRuleBotChallengeAction)

	if err != nil {
		return err
	}

	*o = RuleBotChallengeAction(varRuleBotChallengeAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "robot_challenge_type")
		delete(additionalProperties, "robot_challenge_verification_ttl")
		delete(additionalProperties, "robot_challenge_challenge_ttl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleBotChallengeAction struct {
	value *RuleBotChallengeAction
	isSet bool
}

func (v NullableRuleBotChallengeAction) Get() *RuleBotChallengeAction {
	return v.value
}

func (v *NullableRuleBotChallengeAction) Set(val *RuleBotChallengeAction) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleBotChallengeAction) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleBotChallengeAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleBotChallengeAction(val *RuleBotChallengeAction) *NullableRuleBotChallengeAction {
	return &NullableRuleBotChallengeAction{value: val, isSet: true}
}

func (v NullableRuleBotChallengeAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleBotChallengeAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


