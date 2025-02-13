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

// checks if the HeadersDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeadersDeleteRequest{}

// HeadersDeleteRequest struct for HeadersDeleteRequest
type HeadersDeleteRequest struct {
	Headers []string `json:"headers"`
	AdditionalProperties map[string]interface{}
}

type _HeadersDeleteRequest HeadersDeleteRequest

// NewHeadersDeleteRequest instantiates a new HeadersDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeadersDeleteRequest(headers []string) *HeadersDeleteRequest {
	this := HeadersDeleteRequest{}
	this.Headers = headers
	return &this
}

// NewHeadersDeleteRequestWithDefaults instantiates a new HeadersDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeadersDeleteRequestWithDefaults() *HeadersDeleteRequest {
	this := HeadersDeleteRequest{}
	return &this
}

// GetHeaders returns the Headers field value
func (o *HeadersDeleteRequest) GetHeaders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *HeadersDeleteRequest) GetHeadersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Headers, true
}

// SetHeaders sets field value
func (o *HeadersDeleteRequest) SetHeaders(v []string) {
	o.Headers = v
}

func (o HeadersDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeadersDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["headers"] = o.Headers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HeadersDeleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"headers",
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

	varHeadersDeleteRequest := _HeadersDeleteRequest{}

	err = json.Unmarshal(data, &varHeadersDeleteRequest)

	if err != nil {
		return err
	}

	*o = HeadersDeleteRequest(varHeadersDeleteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "headers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHeadersDeleteRequest struct {
	value *HeadersDeleteRequest
	isSet bool
}

func (v NullableHeadersDeleteRequest) Get() *HeadersDeleteRequest {
	return v.value
}

func (v *NullableHeadersDeleteRequest) Set(val *HeadersDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHeadersDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHeadersDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeadersDeleteRequest(val *HeadersDeleteRequest) *NullableHeadersDeleteRequest {
	return &NullableHeadersDeleteRequest{value: val, isSet: true}
}

func (v NullableHeadersDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeadersDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


