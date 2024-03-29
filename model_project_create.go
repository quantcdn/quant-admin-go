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

// checks if the ProjectCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCreate{}

// ProjectCreate struct for ProjectCreate
type ProjectCreate struct {
	Name string `json:"name"`
	Region string `json:"region"`
	AllowQueryParams *bool `json:"allow_query_params,omitempty"`
	BasicAuthUsername *string `json:"basic_auth_username,omitempty"`
	BasicAuthPassword *string `json:"basic_auth_password,omitempty"`
	BasicAuthPreviewOnly *bool `json:"basic_auth_preview_only,omitempty"`
}

// NewProjectCreate instantiates a new ProjectCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreate(name string, region string) *ProjectCreate {
	this := ProjectCreate{}
	this.Name = name
	this.Region = region
	return &this
}

// NewProjectCreateWithDefaults instantiates a new ProjectCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreateWithDefaults() *ProjectCreate {
	this := ProjectCreate{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectCreate) SetName(v string) {
	o.Name = v
}

// GetRegion returns the Region field value
func (o *ProjectCreate) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ProjectCreate) SetRegion(v string) {
	o.Region = v
}

// GetAllowQueryParams returns the AllowQueryParams field value if set, zero value otherwise.
func (o *ProjectCreate) GetAllowQueryParams() bool {
	if o == nil || IsNil(o.AllowQueryParams) {
		var ret bool
		return ret
	}
	return *o.AllowQueryParams
}

// GetAllowQueryParamsOk returns a tuple with the AllowQueryParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetAllowQueryParamsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowQueryParams) {
		return nil, false
	}
	return o.AllowQueryParams, true
}

// HasAllowQueryParams returns a boolean if a field has been set.
func (o *ProjectCreate) HasAllowQueryParams() bool {
	if o != nil && !IsNil(o.AllowQueryParams) {
		return true
	}

	return false
}

// SetAllowQueryParams gets a reference to the given bool and assigns it to the AllowQueryParams field.
func (o *ProjectCreate) SetAllowQueryParams(v bool) {
	o.AllowQueryParams = &v
}

// GetBasicAuthUsername returns the BasicAuthUsername field value if set, zero value otherwise.
func (o *ProjectCreate) GetBasicAuthUsername() string {
	if o == nil || IsNil(o.BasicAuthUsername) {
		var ret string
		return ret
	}
	return *o.BasicAuthUsername
}

// GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetBasicAuthUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthUsername) {
		return nil, false
	}
	return o.BasicAuthUsername, true
}

// HasBasicAuthUsername returns a boolean if a field has been set.
func (o *ProjectCreate) HasBasicAuthUsername() bool {
	if o != nil && !IsNil(o.BasicAuthUsername) {
		return true
	}

	return false
}

// SetBasicAuthUsername gets a reference to the given string and assigns it to the BasicAuthUsername field.
func (o *ProjectCreate) SetBasicAuthUsername(v string) {
	o.BasicAuthUsername = &v
}

// GetBasicAuthPassword returns the BasicAuthPassword field value if set, zero value otherwise.
func (o *ProjectCreate) GetBasicAuthPassword() string {
	if o == nil || IsNil(o.BasicAuthPassword) {
		var ret string
		return ret
	}
	return *o.BasicAuthPassword
}

// GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetBasicAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthPassword) {
		return nil, false
	}
	return o.BasicAuthPassword, true
}

// HasBasicAuthPassword returns a boolean if a field has been set.
func (o *ProjectCreate) HasBasicAuthPassword() bool {
	if o != nil && !IsNil(o.BasicAuthPassword) {
		return true
	}

	return false
}

// SetBasicAuthPassword gets a reference to the given string and assigns it to the BasicAuthPassword field.
func (o *ProjectCreate) SetBasicAuthPassword(v string) {
	o.BasicAuthPassword = &v
}

// GetBasicAuthPreviewOnly returns the BasicAuthPreviewOnly field value if set, zero value otherwise.
func (o *ProjectCreate) GetBasicAuthPreviewOnly() bool {
	if o == nil || IsNil(o.BasicAuthPreviewOnly) {
		var ret bool
		return ret
	}
	return *o.BasicAuthPreviewOnly
}

// GetBasicAuthPreviewOnlyOk returns a tuple with the BasicAuthPreviewOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCreate) GetBasicAuthPreviewOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.BasicAuthPreviewOnly) {
		return nil, false
	}
	return o.BasicAuthPreviewOnly, true
}

// HasBasicAuthPreviewOnly returns a boolean if a field has been set.
func (o *ProjectCreate) HasBasicAuthPreviewOnly() bool {
	if o != nil && !IsNil(o.BasicAuthPreviewOnly) {
		return true
	}

	return false
}

// SetBasicAuthPreviewOnly gets a reference to the given bool and assigns it to the BasicAuthPreviewOnly field.
func (o *ProjectCreate) SetBasicAuthPreviewOnly(v bool) {
	o.BasicAuthPreviewOnly = &v
}

func (o ProjectCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["region"] = o.Region
	if !IsNil(o.AllowQueryParams) {
		toSerialize["allow_query_params"] = o.AllowQueryParams
	}
	if !IsNil(o.BasicAuthUsername) {
		toSerialize["basic_auth_username"] = o.BasicAuthUsername
	}
	if !IsNil(o.BasicAuthPassword) {
		toSerialize["basic_auth_password"] = o.BasicAuthPassword
	}
	if !IsNil(o.BasicAuthPreviewOnly) {
		toSerialize["basic_auth_preview_only"] = o.BasicAuthPreviewOnly
	}
	return toSerialize, nil
}

type NullableProjectCreate struct {
	value *ProjectCreate
	isSet bool
}

func (v NullableProjectCreate) Get() *ProjectCreate {
	return v.value
}

func (v *NullableProjectCreate) Set(val *ProjectCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreate(val *ProjectCreate) *NullableProjectCreate {
	return &NullableProjectCreate{value: val, isSet: true}
}

func (v NullableProjectCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


