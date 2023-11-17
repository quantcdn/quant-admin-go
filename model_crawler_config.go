/*
QuantCDN Dashboard API

Provides programmatic interface for projects within QuantCDN

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CrawlerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrawlerConfig{}

// CrawlerConfig struct for CrawlerConfig
type CrawlerConfig struct {
	Name *string `json:"name,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Domain *string `json:"domain,omitempty"`
	UrlCount *int32 `json:"url_count,omitempty"`
}

// NewCrawlerConfig instantiates a new CrawlerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrawlerConfig() *CrawlerConfig {
	this := CrawlerConfig{}
	return &this
}

// NewCrawlerConfigWithDefaults instantiates a new CrawlerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrawlerConfigWithDefaults() *CrawlerConfig {
	this := CrawlerConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CrawlerConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CrawlerConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CrawlerConfig) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CrawlerConfig) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerConfig) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CrawlerConfig) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CrawlerConfig) SetUuid(v string) {
	o.Uuid = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CrawlerConfig) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerConfig) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CrawlerConfig) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CrawlerConfig) SetDomain(v string) {
	o.Domain = &v
}

// GetUrlCount returns the UrlCount field value if set, zero value otherwise.
func (o *CrawlerConfig) GetUrlCount() int32 {
	if o == nil || IsNil(o.UrlCount) {
		var ret int32
		return ret
	}
	return *o.UrlCount
}

// GetUrlCountOk returns a tuple with the UrlCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerConfig) GetUrlCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UrlCount) {
		return nil, false
	}
	return o.UrlCount, true
}

// HasUrlCount returns a boolean if a field has been set.
func (o *CrawlerConfig) HasUrlCount() bool {
	if o != nil && !IsNil(o.UrlCount) {
		return true
	}

	return false
}

// SetUrlCount gets a reference to the given int32 and assigns it to the UrlCount field.
func (o *CrawlerConfig) SetUrlCount(v int32) {
	o.UrlCount = &v
}

func (o CrawlerConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrawlerConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.UrlCount) {
		toSerialize["url_count"] = o.UrlCount
	}
	return toSerialize, nil
}

type NullableCrawlerConfig struct {
	value *CrawlerConfig
	isSet bool
}

func (v NullableCrawlerConfig) Get() *CrawlerConfig {
	return v.value
}

func (v *NullableCrawlerConfig) Set(val *CrawlerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCrawlerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCrawlerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrawlerConfig(val *CrawlerConfig) *NullableCrawlerConfig {
	return &NullableCrawlerConfig{value: val, isSet: true}
}

func (v NullableCrawlerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrawlerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

