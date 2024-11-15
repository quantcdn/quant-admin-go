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

// checks if the RuleRedirectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleRedirectRequest{}

// RuleRedirectRequest struct for RuleRedirectRequest
type RuleRedirectRequest struct {
	Domain []string `json:"domain"`
	Name *string `json:"name,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Disabled bool `json:"disabled"`
	Url []string `json:"url"`
	Country *string `json:"country,omitempty"`
	CountryIs []string `json:"country_is,omitempty"`
	CountryIsNot []string `json:"country_is_not,omitempty"`
	Method *string `json:"method,omitempty"`
	MethodIs []string `json:"method_is,omitempty"`
	MethodIsNot []string `json:"method_is_not,omitempty"`
	Ip *string `json:"ip,omitempty"`
	IpIs []string `json:"ip_is,omitempty"`
	IpIsNot []string `json:"ip_is_not,omitempty"`
	OnlyWithCookie *bool `json:"only_with_cookie,omitempty"`
	CookieName *string `json:"cookie_name,omitempty"`
	RedirectTo string `json:"redirect_to"`
	RedirectCode string `json:"redirect_code"`
	AdditionalProperties map[string]interface{}
}

type _RuleRedirectRequest RuleRedirectRequest

// NewRuleRedirectRequest instantiates a new RuleRedirectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleRedirectRequest(domain []string, disabled bool, url []string, redirectTo string, redirectCode string) *RuleRedirectRequest {
	this := RuleRedirectRequest{}
	this.Domain = domain
	this.Disabled = disabled
	this.Url = url
	var onlyWithCookie bool = false
	this.OnlyWithCookie = &onlyWithCookie
	this.RedirectTo = redirectTo
	this.RedirectCode = redirectCode
	return &this
}

// NewRuleRedirectRequestWithDefaults instantiates a new RuleRedirectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleRedirectRequestWithDefaults() *RuleRedirectRequest {
	this := RuleRedirectRequest{}
	var disabled bool = false
	this.Disabled = disabled
	var onlyWithCookie bool = false
	this.OnlyWithCookie = &onlyWithCookie
	var redirectCode string = "301"
	this.RedirectCode = redirectCode
	return &this
}

// GetDomain returns the Domain field value
func (o *RuleRedirectRequest) GetDomain() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetDomainOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain, true
}

// SetDomain sets field value
func (o *RuleRedirectRequest) SetDomain(v []string) {
	o.Domain = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleRedirectRequest) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RuleRedirectRequest) SetUuid(v string) {
	o.Uuid = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *RuleRedirectRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetDisabled returns the Disabled field value
func (o *RuleRedirectRequest) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *RuleRedirectRequest) SetDisabled(v bool) {
	o.Disabled = v
}

// GetUrl returns the Url field value
func (o *RuleRedirectRequest) GetUrl() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetUrlOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url, true
}

// SetUrl sets field value
func (o *RuleRedirectRequest) SetUrl(v []string) {
	o.Url = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *RuleRedirectRequest) SetCountry(v string) {
	o.Country = &v
}

// GetCountryIs returns the CountryIs field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetCountryIs() []string {
	if o == nil || IsNil(o.CountryIs) {
		var ret []string
		return ret
	}
	return o.CountryIs
}

// GetCountryIsOk returns a tuple with the CountryIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetCountryIsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIs) {
		return nil, false
	}
	return o.CountryIs, true
}

// HasCountryIs returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasCountryIs() bool {
	if o != nil && !IsNil(o.CountryIs) {
		return true
	}

	return false
}

// SetCountryIs gets a reference to the given []string and assigns it to the CountryIs field.
func (o *RuleRedirectRequest) SetCountryIs(v []string) {
	o.CountryIs = v
}

// GetCountryIsNot returns the CountryIsNot field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetCountryIsNot() []string {
	if o == nil || IsNil(o.CountryIsNot) {
		var ret []string
		return ret
	}
	return o.CountryIsNot
}

// GetCountryIsNotOk returns a tuple with the CountryIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetCountryIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIsNot) {
		return nil, false
	}
	return o.CountryIsNot, true
}

// HasCountryIsNot returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasCountryIsNot() bool {
	if o != nil && !IsNil(o.CountryIsNot) {
		return true
	}

	return false
}

// SetCountryIsNot gets a reference to the given []string and assigns it to the CountryIsNot field.
func (o *RuleRedirectRequest) SetCountryIsNot(v []string) {
	o.CountryIsNot = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *RuleRedirectRequest) SetMethod(v string) {
	o.Method = &v
}

// GetMethodIs returns the MethodIs field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetMethodIs() []string {
	if o == nil || IsNil(o.MethodIs) {
		var ret []string
		return ret
	}
	return o.MethodIs
}

// GetMethodIsOk returns a tuple with the MethodIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetMethodIsOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIs) {
		return nil, false
	}
	return o.MethodIs, true
}

// HasMethodIs returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasMethodIs() bool {
	if o != nil && !IsNil(o.MethodIs) {
		return true
	}

	return false
}

// SetMethodIs gets a reference to the given []string and assigns it to the MethodIs field.
func (o *RuleRedirectRequest) SetMethodIs(v []string) {
	o.MethodIs = v
}

// GetMethodIsNot returns the MethodIsNot field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetMethodIsNot() []string {
	if o == nil || IsNil(o.MethodIsNot) {
		var ret []string
		return ret
	}
	return o.MethodIsNot
}

// GetMethodIsNotOk returns a tuple with the MethodIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetMethodIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.MethodIsNot) {
		return nil, false
	}
	return o.MethodIsNot, true
}

// HasMethodIsNot returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasMethodIsNot() bool {
	if o != nil && !IsNil(o.MethodIsNot) {
		return true
	}

	return false
}

// SetMethodIsNot gets a reference to the given []string and assigns it to the MethodIsNot field.
func (o *RuleRedirectRequest) SetMethodIsNot(v []string) {
	o.MethodIsNot = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RuleRedirectRequest) SetIp(v string) {
	o.Ip = &v
}

// GetIpIs returns the IpIs field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetIpIs() []string {
	if o == nil || IsNil(o.IpIs) {
		var ret []string
		return ret
	}
	return o.IpIs
}

// GetIpIsOk returns a tuple with the IpIs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetIpIsOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIs) {
		return nil, false
	}
	return o.IpIs, true
}

// HasIpIs returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasIpIs() bool {
	if o != nil && !IsNil(o.IpIs) {
		return true
	}

	return false
}

// SetIpIs gets a reference to the given []string and assigns it to the IpIs field.
func (o *RuleRedirectRequest) SetIpIs(v []string) {
	o.IpIs = v
}

// GetIpIsNot returns the IpIsNot field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetIpIsNot() []string {
	if o == nil || IsNil(o.IpIsNot) {
		var ret []string
		return ret
	}
	return o.IpIsNot
}

// GetIpIsNotOk returns a tuple with the IpIsNot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetIpIsNotOk() ([]string, bool) {
	if o == nil || IsNil(o.IpIsNot) {
		return nil, false
	}
	return o.IpIsNot, true
}

// HasIpIsNot returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasIpIsNot() bool {
	if o != nil && !IsNil(o.IpIsNot) {
		return true
	}

	return false
}

// SetIpIsNot gets a reference to the given []string and assigns it to the IpIsNot field.
func (o *RuleRedirectRequest) SetIpIsNot(v []string) {
	o.IpIsNot = v
}

// GetOnlyWithCookie returns the OnlyWithCookie field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetOnlyWithCookie() bool {
	if o == nil || IsNil(o.OnlyWithCookie) {
		var ret bool
		return ret
	}
	return *o.OnlyWithCookie
}

// GetOnlyWithCookieOk returns a tuple with the OnlyWithCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetOnlyWithCookieOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyWithCookie) {
		return nil, false
	}
	return o.OnlyWithCookie, true
}

// HasOnlyWithCookie returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasOnlyWithCookie() bool {
	if o != nil && !IsNil(o.OnlyWithCookie) {
		return true
	}

	return false
}

// SetOnlyWithCookie gets a reference to the given bool and assigns it to the OnlyWithCookie field.
func (o *RuleRedirectRequest) SetOnlyWithCookie(v bool) {
	o.OnlyWithCookie = &v
}

// GetCookieName returns the CookieName field value if set, zero value otherwise.
func (o *RuleRedirectRequest) GetCookieName() string {
	if o == nil || IsNil(o.CookieName) {
		var ret string
		return ret
	}
	return *o.CookieName
}

// GetCookieNameOk returns a tuple with the CookieName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetCookieNameOk() (*string, bool) {
	if o == nil || IsNil(o.CookieName) {
		return nil, false
	}
	return o.CookieName, true
}

// HasCookieName returns a boolean if a field has been set.
func (o *RuleRedirectRequest) HasCookieName() bool {
	if o != nil && !IsNil(o.CookieName) {
		return true
	}

	return false
}

// SetCookieName gets a reference to the given string and assigns it to the CookieName field.
func (o *RuleRedirectRequest) SetCookieName(v string) {
	o.CookieName = &v
}

// GetRedirectTo returns the RedirectTo field value
func (o *RuleRedirectRequest) GetRedirectTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectTo
}

// GetRedirectToOk returns a tuple with the RedirectTo field value
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetRedirectToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectTo, true
}

// SetRedirectTo sets field value
func (o *RuleRedirectRequest) SetRedirectTo(v string) {
	o.RedirectTo = v
}

// GetRedirectCode returns the RedirectCode field value
func (o *RuleRedirectRequest) GetRedirectCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectCode
}

// GetRedirectCodeOk returns a tuple with the RedirectCode field value
// and a boolean to check if the value has been set.
func (o *RuleRedirectRequest) GetRedirectCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectCode, true
}

// SetRedirectCode sets field value
func (o *RuleRedirectRequest) SetRedirectCode(v string) {
	o.RedirectCode = v
}

func (o RuleRedirectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleRedirectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	toSerialize["disabled"] = o.Disabled
	toSerialize["url"] = o.Url
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CountryIs) {
		toSerialize["country_is"] = o.CountryIs
	}
	if !IsNil(o.CountryIsNot) {
		toSerialize["country_is_not"] = o.CountryIsNot
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.MethodIs) {
		toSerialize["method_is"] = o.MethodIs
	}
	if !IsNil(o.MethodIsNot) {
		toSerialize["method_is_not"] = o.MethodIsNot
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.IpIs) {
		toSerialize["ip_is"] = o.IpIs
	}
	if !IsNil(o.IpIsNot) {
		toSerialize["ip_is_not"] = o.IpIsNot
	}
	if !IsNil(o.OnlyWithCookie) {
		toSerialize["only_with_cookie"] = o.OnlyWithCookie
	}
	if !IsNil(o.CookieName) {
		toSerialize["cookie_name"] = o.CookieName
	}
	toSerialize["redirect_to"] = o.RedirectTo
	toSerialize["redirect_code"] = o.RedirectCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RuleRedirectRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"disabled",
		"url",
		"redirect_to",
		"redirect_code",
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

	varRuleRedirectRequest := _RuleRedirectRequest{}

	err = json.Unmarshal(data, &varRuleRedirectRequest)

	if err != nil {
		return err
	}

	*o = RuleRedirectRequest(varRuleRedirectRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "name")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "url")
		delete(additionalProperties, "country")
		delete(additionalProperties, "country_is")
		delete(additionalProperties, "country_is_not")
		delete(additionalProperties, "method")
		delete(additionalProperties, "method_is")
		delete(additionalProperties, "method_is_not")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "ip_is")
		delete(additionalProperties, "ip_is_not")
		delete(additionalProperties, "only_with_cookie")
		delete(additionalProperties, "cookie_name")
		delete(additionalProperties, "redirect_to")
		delete(additionalProperties, "redirect_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRuleRedirectRequest struct {
	value *RuleRedirectRequest
	isSet bool
}

func (v NullableRuleRedirectRequest) Get() *RuleRedirectRequest {
	return v.value
}

func (v *NullableRuleRedirectRequest) Set(val *RuleRedirectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleRedirectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleRedirectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleRedirectRequest(val *RuleRedirectRequest) *NullableRuleRedirectRequest {
	return &NullableRuleRedirectRequest{value: val, isSet: true}
}

func (v NullableRuleRedirectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleRedirectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


