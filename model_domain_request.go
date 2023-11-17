/*
Quant administration API

The Quant administration API provides programmatic access to manage projects within your available organizations. 

API version: 2.0.0
Contact: apiteam@quantcdn.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DomainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainRequest{}

// DomainRequest struct for DomainRequest
type DomainRequest struct {
	Name *string `json:"name,omitempty"`
	Domain *string `json:"domain,omitempty"`
	BrowserMode *bool `json:"browser_mode,omitempty"`
	UrlList []string `json:"url_list,omitempty"`
	Headers map[string]interface{} `json:"headers,omitempty"`
}

// NewDomainRequest instantiates a new DomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainRequest() *DomainRequest {
	this := DomainRequest{}
	return &this
}

// NewDomainRequestWithDefaults instantiates a new DomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainRequestWithDefaults() *DomainRequest {
	this := DomainRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DomainRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DomainRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DomainRequest) SetName(v string) {
	o.Name = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DomainRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DomainRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DomainRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetBrowserMode returns the BrowserMode field value if set, zero value otherwise.
func (o *DomainRequest) GetBrowserMode() bool {
	if o == nil || IsNil(o.BrowserMode) {
		var ret bool
		return ret
	}
	return *o.BrowserMode
}

// GetBrowserModeOk returns a tuple with the BrowserMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetBrowserModeOk() (*bool, bool) {
	if o == nil || IsNil(o.BrowserMode) {
		return nil, false
	}
	return o.BrowserMode, true
}

// HasBrowserMode returns a boolean if a field has been set.
func (o *DomainRequest) HasBrowserMode() bool {
	if o != nil && !IsNil(o.BrowserMode) {
		return true
	}

	return false
}

// SetBrowserMode gets a reference to the given bool and assigns it to the BrowserMode field.
func (o *DomainRequest) SetBrowserMode(v bool) {
	o.BrowserMode = &v
}

// GetUrlList returns the UrlList field value if set, zero value otherwise.
func (o *DomainRequest) GetUrlList() []string {
	if o == nil || IsNil(o.UrlList) {
		var ret []string
		return ret
	}
	return o.UrlList
}

// GetUrlListOk returns a tuple with the UrlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetUrlListOk() ([]string, bool) {
	if o == nil || IsNil(o.UrlList) {
		return nil, false
	}
	return o.UrlList, true
}

// HasUrlList returns a boolean if a field has been set.
func (o *DomainRequest) HasUrlList() bool {
	if o != nil && !IsNil(o.UrlList) {
		return true
	}

	return false
}

// SetUrlList gets a reference to the given []string and assigns it to the UrlList field.
func (o *DomainRequest) SetUrlList(v []string) {
	o.UrlList = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *DomainRequest) GetHeaders() map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainRequest) GetHeadersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]interface{}{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *DomainRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *DomainRequest) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

func (o DomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.BrowserMode) {
		toSerialize["browser_mode"] = o.BrowserMode
	}
	if !IsNil(o.UrlList) {
		toSerialize["url_list"] = o.UrlList
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

type NullableDomainRequest struct {
	value *DomainRequest
	isSet bool
}

func (v NullableDomainRequest) Get() *DomainRequest {
	return v.value
}

func (v *NullableDomainRequest) Set(val *DomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainRequest(val *DomainRequest) *NullableDomainRequest {
	return &NullableDomainRequest{value: val, isSet: true}
}

func (v NullableDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


