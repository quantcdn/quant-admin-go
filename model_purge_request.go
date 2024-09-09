/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PurgeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurgeRequest{}

// PurgeRequest struct for PurgeRequest
type PurgeRequest struct {
	CacheKeys []string `json:"cache_keys"`
	Scope string `json:"scope"`
}

type _PurgeRequest PurgeRequest

// NewPurgeRequest instantiates a new PurgeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurgeRequest(cacheKeys []string, scope string) *PurgeRequest {
	this := PurgeRequest{}
	this.CacheKeys = cacheKeys
	this.Scope = scope
	return &this
}

// NewPurgeRequestWithDefaults instantiates a new PurgeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurgeRequestWithDefaults() *PurgeRequest {
	this := PurgeRequest{}
	var scope string = "project"
	this.Scope = scope
	return &this
}

// GetCacheKeys returns the CacheKeys field value
func (o *PurgeRequest) GetCacheKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CacheKeys
}

// GetCacheKeysOk returns a tuple with the CacheKeys field value
// and a boolean to check if the value has been set.
func (o *PurgeRequest) GetCacheKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CacheKeys, true
}

// SetCacheKeys sets field value
func (o *PurgeRequest) SetCacheKeys(v []string) {
	o.CacheKeys = v
}

// GetScope returns the Scope field value
func (o *PurgeRequest) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *PurgeRequest) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *PurgeRequest) SetScope(v string) {
	o.Scope = v
}

func (o PurgeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PurgeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cache_keys"] = o.CacheKeys
	toSerialize["scope"] = o.Scope
	return toSerialize, nil
}

func (o *PurgeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cache_keys",
		"scope",
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

	varPurgeRequest := _PurgeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPurgeRequest)

	if err != nil {
		return err
	}

	*o = PurgeRequest(varPurgeRequest)

	return err
}

type NullablePurgeRequest struct {
	value *PurgeRequest
	isSet bool
}

func (v NullablePurgeRequest) Get() *PurgeRequest {
	return v.value
}

func (v *NullablePurgeRequest) Set(val *PurgeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurgeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurgeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurgeRequest(val *PurgeRequest) *NullablePurgeRequest {
	return &NullablePurgeRequest{value: val, isSet: true}
}

func (v NullablePurgeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurgeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


