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

// checks if the ProjectEdit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEdit{}

// ProjectEdit struct for ProjectEdit
type ProjectEdit struct {
	Name string `json:"name"`
	Region string `json:"region"`
	AllowQueryParams bool `json:"allow_query_params"`
	BasicAuthUsername *string `json:"basic_auth_username,omitempty"`
	BasicAuthPassword *string `json:"basic_auth_password,omitempty"`
	BasicAuthPreviewOnly *bool `json:"basic_auth_preview_only,omitempty"`
	// Optional bucket name for S3 sync function
	CustomS3SyncBucket *string `json:"custom_s3_sync_bucket,omitempty"`
	// Bucket region for S3 sync. Required if s3 sync bucket is provided
	CustomS3SyncRegion *string `json:"custom_s3_sync_region,omitempty"`
	// Access key for S3 sync. Required if s3 sync bucket is provided
	CustomS3SyncAccessKey *string `json:"custom_s3_sync_access_key,omitempty"`
	// Secret key for S3 sync. Required if s3 sync bucket is provided
	CustomS3SyncSecretKey *string `json:"custom_s3_sync_secret_key,omitempty"`
}

// NewProjectEdit instantiates a new ProjectEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEdit(name string, region string, allowQueryParams bool) *ProjectEdit {
	this := ProjectEdit{}
	this.Name = name
	this.Region = region
	this.AllowQueryParams = allowQueryParams
	return &this
}

// NewProjectEditWithDefaults instantiates a new ProjectEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEditWithDefaults() *ProjectEdit {
	this := ProjectEdit{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectEdit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectEdit) SetName(v string) {
	o.Name = v
}

// GetRegion returns the Region field value
func (o *ProjectEdit) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ProjectEdit) SetRegion(v string) {
	o.Region = v
}

// GetAllowQueryParams returns the AllowQueryParams field value
func (o *ProjectEdit) GetAllowQueryParams() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowQueryParams
}

// GetAllowQueryParamsOk returns a tuple with the AllowQueryParams field value
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetAllowQueryParamsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowQueryParams, true
}

// SetAllowQueryParams sets field value
func (o *ProjectEdit) SetAllowQueryParams(v bool) {
	o.AllowQueryParams = v
}

// GetBasicAuthUsername returns the BasicAuthUsername field value if set, zero value otherwise.
func (o *ProjectEdit) GetBasicAuthUsername() string {
	if o == nil || IsNil(o.BasicAuthUsername) {
		var ret string
		return ret
	}
	return *o.BasicAuthUsername
}

// GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetBasicAuthUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthUsername) {
		return nil, false
	}
	return o.BasicAuthUsername, true
}

// HasBasicAuthUsername returns a boolean if a field has been set.
func (o *ProjectEdit) HasBasicAuthUsername() bool {
	if o != nil && !IsNil(o.BasicAuthUsername) {
		return true
	}

	return false
}

// SetBasicAuthUsername gets a reference to the given string and assigns it to the BasicAuthUsername field.
func (o *ProjectEdit) SetBasicAuthUsername(v string) {
	o.BasicAuthUsername = &v
}

// GetBasicAuthPassword returns the BasicAuthPassword field value if set, zero value otherwise.
func (o *ProjectEdit) GetBasicAuthPassword() string {
	if o == nil || IsNil(o.BasicAuthPassword) {
		var ret string
		return ret
	}
	return *o.BasicAuthPassword
}

// GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetBasicAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.BasicAuthPassword) {
		return nil, false
	}
	return o.BasicAuthPassword, true
}

// HasBasicAuthPassword returns a boolean if a field has been set.
func (o *ProjectEdit) HasBasicAuthPassword() bool {
	if o != nil && !IsNil(o.BasicAuthPassword) {
		return true
	}

	return false
}

// SetBasicAuthPassword gets a reference to the given string and assigns it to the BasicAuthPassword field.
func (o *ProjectEdit) SetBasicAuthPassword(v string) {
	o.BasicAuthPassword = &v
}

// GetBasicAuthPreviewOnly returns the BasicAuthPreviewOnly field value if set, zero value otherwise.
func (o *ProjectEdit) GetBasicAuthPreviewOnly() bool {
	if o == nil || IsNil(o.BasicAuthPreviewOnly) {
		var ret bool
		return ret
	}
	return *o.BasicAuthPreviewOnly
}

// GetBasicAuthPreviewOnlyOk returns a tuple with the BasicAuthPreviewOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetBasicAuthPreviewOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.BasicAuthPreviewOnly) {
		return nil, false
	}
	return o.BasicAuthPreviewOnly, true
}

// HasBasicAuthPreviewOnly returns a boolean if a field has been set.
func (o *ProjectEdit) HasBasicAuthPreviewOnly() bool {
	if o != nil && !IsNil(o.BasicAuthPreviewOnly) {
		return true
	}

	return false
}

// SetBasicAuthPreviewOnly gets a reference to the given bool and assigns it to the BasicAuthPreviewOnly field.
func (o *ProjectEdit) SetBasicAuthPreviewOnly(v bool) {
	o.BasicAuthPreviewOnly = &v
}

// GetCustomS3SyncBucket returns the CustomS3SyncBucket field value if set, zero value otherwise.
func (o *ProjectEdit) GetCustomS3SyncBucket() string {
	if o == nil || IsNil(o.CustomS3SyncBucket) {
		var ret string
		return ret
	}
	return *o.CustomS3SyncBucket
}

// GetCustomS3SyncBucketOk returns a tuple with the CustomS3SyncBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetCustomS3SyncBucketOk() (*string, bool) {
	if o == nil || IsNil(o.CustomS3SyncBucket) {
		return nil, false
	}
	return o.CustomS3SyncBucket, true
}

// HasCustomS3SyncBucket returns a boolean if a field has been set.
func (o *ProjectEdit) HasCustomS3SyncBucket() bool {
	if o != nil && !IsNil(o.CustomS3SyncBucket) {
		return true
	}

	return false
}

// SetCustomS3SyncBucket gets a reference to the given string and assigns it to the CustomS3SyncBucket field.
func (o *ProjectEdit) SetCustomS3SyncBucket(v string) {
	o.CustomS3SyncBucket = &v
}

// GetCustomS3SyncRegion returns the CustomS3SyncRegion field value if set, zero value otherwise.
func (o *ProjectEdit) GetCustomS3SyncRegion() string {
	if o == nil || IsNil(o.CustomS3SyncRegion) {
		var ret string
		return ret
	}
	return *o.CustomS3SyncRegion
}

// GetCustomS3SyncRegionOk returns a tuple with the CustomS3SyncRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetCustomS3SyncRegionOk() (*string, bool) {
	if o == nil || IsNil(o.CustomS3SyncRegion) {
		return nil, false
	}
	return o.CustomS3SyncRegion, true
}

// HasCustomS3SyncRegion returns a boolean if a field has been set.
func (o *ProjectEdit) HasCustomS3SyncRegion() bool {
	if o != nil && !IsNil(o.CustomS3SyncRegion) {
		return true
	}

	return false
}

// SetCustomS3SyncRegion gets a reference to the given string and assigns it to the CustomS3SyncRegion field.
func (o *ProjectEdit) SetCustomS3SyncRegion(v string) {
	o.CustomS3SyncRegion = &v
}

// GetCustomS3SyncAccessKey returns the CustomS3SyncAccessKey field value if set, zero value otherwise.
func (o *ProjectEdit) GetCustomS3SyncAccessKey() string {
	if o == nil || IsNil(o.CustomS3SyncAccessKey) {
		var ret string
		return ret
	}
	return *o.CustomS3SyncAccessKey
}

// GetCustomS3SyncAccessKeyOk returns a tuple with the CustomS3SyncAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetCustomS3SyncAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CustomS3SyncAccessKey) {
		return nil, false
	}
	return o.CustomS3SyncAccessKey, true
}

// HasCustomS3SyncAccessKey returns a boolean if a field has been set.
func (o *ProjectEdit) HasCustomS3SyncAccessKey() bool {
	if o != nil && !IsNil(o.CustomS3SyncAccessKey) {
		return true
	}

	return false
}

// SetCustomS3SyncAccessKey gets a reference to the given string and assigns it to the CustomS3SyncAccessKey field.
func (o *ProjectEdit) SetCustomS3SyncAccessKey(v string) {
	o.CustomS3SyncAccessKey = &v
}

// GetCustomS3SyncSecretKey returns the CustomS3SyncSecretKey field value if set, zero value otherwise.
func (o *ProjectEdit) GetCustomS3SyncSecretKey() string {
	if o == nil || IsNil(o.CustomS3SyncSecretKey) {
		var ret string
		return ret
	}
	return *o.CustomS3SyncSecretKey
}

// GetCustomS3SyncSecretKeyOk returns a tuple with the CustomS3SyncSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEdit) GetCustomS3SyncSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CustomS3SyncSecretKey) {
		return nil, false
	}
	return o.CustomS3SyncSecretKey, true
}

// HasCustomS3SyncSecretKey returns a boolean if a field has been set.
func (o *ProjectEdit) HasCustomS3SyncSecretKey() bool {
	if o != nil && !IsNil(o.CustomS3SyncSecretKey) {
		return true
	}

	return false
}

// SetCustomS3SyncSecretKey gets a reference to the given string and assigns it to the CustomS3SyncSecretKey field.
func (o *ProjectEdit) SetCustomS3SyncSecretKey(v string) {
	o.CustomS3SyncSecretKey = &v
}

func (o ProjectEdit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEdit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["region"] = o.Region
	toSerialize["allow_query_params"] = o.AllowQueryParams
	if !IsNil(o.BasicAuthUsername) {
		toSerialize["basic_auth_username"] = o.BasicAuthUsername
	}
	if !IsNil(o.BasicAuthPassword) {
		toSerialize["basic_auth_password"] = o.BasicAuthPassword
	}
	if !IsNil(o.BasicAuthPreviewOnly) {
		toSerialize["basic_auth_preview_only"] = o.BasicAuthPreviewOnly
	}
	if !IsNil(o.CustomS3SyncBucket) {
		toSerialize["custom_s3_sync_bucket"] = o.CustomS3SyncBucket
	}
	if !IsNil(o.CustomS3SyncRegion) {
		toSerialize["custom_s3_sync_region"] = o.CustomS3SyncRegion
	}
	if !IsNil(o.CustomS3SyncAccessKey) {
		toSerialize["custom_s3_sync_access_key"] = o.CustomS3SyncAccessKey
	}
	if !IsNil(o.CustomS3SyncSecretKey) {
		toSerialize["custom_s3_sync_secret_key"] = o.CustomS3SyncSecretKey
	}
	return toSerialize, nil
}

type NullableProjectEdit struct {
	value *ProjectEdit
	isSet bool
}

func (v NullableProjectEdit) Get() *ProjectEdit {
	return v.value
}

func (v *NullableProjectEdit) Set(val *ProjectEdit) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEdit) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEdit(val *ProjectEdit) *NullableProjectEdit {
	return &NullableProjectEdit{value: val, isSet: true}
}

func (v NullableProjectEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

