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

// checks if the CrawlerScheduleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrawlerScheduleRequest{}

// CrawlerScheduleRequest struct for CrawlerScheduleRequest
type CrawlerScheduleRequest struct {
	Name *string `json:"name,omitempty"`
	ScheduleCronString string `json:"schedule_cron_string"`
	AdditionalProperties map[string]interface{}
}

type _CrawlerScheduleRequest CrawlerScheduleRequest

// NewCrawlerScheduleRequest instantiates a new CrawlerScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrawlerScheduleRequest(scheduleCronString string) *CrawlerScheduleRequest {
	this := CrawlerScheduleRequest{}
	this.ScheduleCronString = scheduleCronString
	return &this
}

// NewCrawlerScheduleRequestWithDefaults instantiates a new CrawlerScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrawlerScheduleRequestWithDefaults() *CrawlerScheduleRequest {
	this := CrawlerScheduleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CrawlerScheduleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerScheduleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CrawlerScheduleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CrawlerScheduleRequest) SetName(v string) {
	o.Name = &v
}

// GetScheduleCronString returns the ScheduleCronString field value
func (o *CrawlerScheduleRequest) GetScheduleCronString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduleCronString
}

// GetScheduleCronStringOk returns a tuple with the ScheduleCronString field value
// and a boolean to check if the value has been set.
func (o *CrawlerScheduleRequest) GetScheduleCronStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleCronString, true
}

// SetScheduleCronString sets field value
func (o *CrawlerScheduleRequest) SetScheduleCronString(v string) {
	o.ScheduleCronString = v
}

func (o CrawlerScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrawlerScheduleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["schedule_cron_string"] = o.ScheduleCronString

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CrawlerScheduleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schedule_cron_string",
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

	varCrawlerScheduleRequest := _CrawlerScheduleRequest{}

	err = json.Unmarshal(data, &varCrawlerScheduleRequest)

	if err != nil {
		return err
	}

	*o = CrawlerScheduleRequest(varCrawlerScheduleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "schedule_cron_string")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCrawlerScheduleRequest struct {
	value *CrawlerScheduleRequest
	isSet bool
}

func (v NullableCrawlerScheduleRequest) Get() *CrawlerScheduleRequest {
	return v.value
}

func (v *NullableCrawlerScheduleRequest) Set(val *CrawlerScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCrawlerScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCrawlerScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrawlerScheduleRequest(val *CrawlerScheduleRequest) *NullableCrawlerScheduleRequest {
	return &NullableCrawlerScheduleRequest{value: val, isSet: true}
}

func (v NullableCrawlerScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrawlerScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


