/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quantadmingo

import (
	"encoding/json"
)

// checks if the CrawlerRequestUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrawlerRequestUpdate{}

// CrawlerRequestUpdate struct for CrawlerRequestUpdate
type CrawlerRequestUpdate struct {
	Name *string `json:"name,omitempty"`
	Domain *string `json:"domain,omitempty"`
	BrowserMode *bool `json:"browser_mode,omitempty"`
	Urls []string `json:"urls,omitempty"`
	Headers *map[string]string `json:"headers,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CrawlerRequestUpdate CrawlerRequestUpdate

// NewCrawlerRequestUpdate instantiates a new CrawlerRequestUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrawlerRequestUpdate() *CrawlerRequestUpdate {
	this := CrawlerRequestUpdate{}
	var browserMode bool = false
	this.BrowserMode = &browserMode
	return &this
}

// NewCrawlerRequestUpdateWithDefaults instantiates a new CrawlerRequestUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrawlerRequestUpdateWithDefaults() *CrawlerRequestUpdate {
	this := CrawlerRequestUpdate{}
	var browserMode bool = false
	this.BrowserMode = &browserMode
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CrawlerRequestUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerRequestUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CrawlerRequestUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CrawlerRequestUpdate) SetName(v string) {
	o.Name = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CrawlerRequestUpdate) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerRequestUpdate) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CrawlerRequestUpdate) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CrawlerRequestUpdate) SetDomain(v string) {
	o.Domain = &v
}

// GetBrowserMode returns the BrowserMode field value if set, zero value otherwise.
func (o *CrawlerRequestUpdate) GetBrowserMode() bool {
	if o == nil || IsNil(o.BrowserMode) {
		var ret bool
		return ret
	}
	return *o.BrowserMode
}

// GetBrowserModeOk returns a tuple with the BrowserMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerRequestUpdate) GetBrowserModeOk() (*bool, bool) {
	if o == nil || IsNil(o.BrowserMode) {
		return nil, false
	}
	return o.BrowserMode, true
}

// HasBrowserMode returns a boolean if a field has been set.
func (o *CrawlerRequestUpdate) HasBrowserMode() bool {
	if o != nil && !IsNil(o.BrowserMode) {
		return true
	}

	return false
}

// SetBrowserMode gets a reference to the given bool and assigns it to the BrowserMode field.
func (o *CrawlerRequestUpdate) SetBrowserMode(v bool) {
	o.BrowserMode = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *CrawlerRequestUpdate) GetUrls() []string {
	if o == nil || IsNil(o.Urls) {
		var ret []string
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerRequestUpdate) GetUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *CrawlerRequestUpdate) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *CrawlerRequestUpdate) SetUrls(v []string) {
	o.Urls = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *CrawlerRequestUpdate) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerRequestUpdate) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *CrawlerRequestUpdate) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *CrawlerRequestUpdate) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *CrawlerRequestUpdate) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrawlerRequestUpdate) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *CrawlerRequestUpdate) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *CrawlerRequestUpdate) SetExclude(v []string) {
	o.Exclude = v
}

func (o CrawlerRequestUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrawlerRequestUpdate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CrawlerRequestUpdate) UnmarshalJSON(data []byte) (err error) {
	varCrawlerRequestUpdate := _CrawlerRequestUpdate{}

	err = json.Unmarshal(data, &varCrawlerRequestUpdate)

	if err != nil {
		return err
	}

	*o = CrawlerRequestUpdate(varCrawlerRequestUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "browser_mode")
		delete(additionalProperties, "urls")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCrawlerRequestUpdate struct {
	value *CrawlerRequestUpdate
	isSet bool
}

func (v NullableCrawlerRequestUpdate) Get() *CrawlerRequestUpdate {
	return v.value
}

func (v *NullableCrawlerRequestUpdate) Set(val *CrawlerRequestUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCrawlerRequestUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCrawlerRequestUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrawlerRequestUpdate(val *CrawlerRequestUpdate) *NullableCrawlerRequestUpdate {
	return &NullableCrawlerRequestUpdate{value: val, isSet: true}
}

func (v NullableCrawlerRequestUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrawlerRequestUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


