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

// checks if the Crawler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Crawler{}

// Crawler struct for Crawler
type Crawler struct {
	Id *int32 `json:"id,omitempty"`
	ProjectId *int32 `json:"project_id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Config *string `json:"config,omitempty"`
	UrlsList []string `json:"urls_list,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Domain *string `json:"domain,omitempty"`
	DomainVerified *int32 `json:"domain_verified,omitempty"`
}

// NewCrawler instantiates a new Crawler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrawler() *Crawler {
	this := Crawler{}
	return &this
}

// NewCrawlerWithDefaults instantiates a new Crawler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrawlerWithDefaults() *Crawler {
	this := Crawler{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Crawler) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Crawler) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Crawler) SetId(v int32) {
	o.Id = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Crawler) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Crawler) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *Crawler) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Crawler) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Crawler) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Crawler) SetUuid(v string) {
	o.Uuid = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Crawler) GetConfig() string {
	if o == nil || IsNil(o.Config) {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetConfigOk() (*string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Crawler) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *Crawler) SetConfig(v string) {
	o.Config = &v
}

// GetUrlsList returns the UrlsList field value if set, zero value otherwise.
func (o *Crawler) GetUrlsList() []string {
	if o == nil || IsNil(o.UrlsList) {
		var ret []string
		return ret
	}
	return o.UrlsList
}

// GetUrlsListOk returns a tuple with the UrlsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetUrlsListOk() ([]string, bool) {
	if o == nil || IsNil(o.UrlsList) {
		return nil, false
	}
	return o.UrlsList, true
}

// HasUrlsList returns a boolean if a field has been set.
func (o *Crawler) HasUrlsList() bool {
	if o != nil && !IsNil(o.UrlsList) {
		return true
	}

	return false
}

// SetUrlsList gets a reference to the given []string and assigns it to the UrlsList field.
func (o *Crawler) SetUrlsList(v []string) {
	o.UrlsList = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Crawler) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Crawler) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Crawler) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Crawler) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Crawler) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Crawler) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Crawler) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Crawler) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Crawler) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Crawler) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Crawler) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Crawler) SetDomain(v string) {
	o.Domain = &v
}

// GetDomainVerified returns the DomainVerified field value if set, zero value otherwise.
func (o *Crawler) GetDomainVerified() int32 {
	if o == nil || IsNil(o.DomainVerified) {
		var ret int32
		return ret
	}
	return *o.DomainVerified
}

// GetDomainVerifiedOk returns a tuple with the DomainVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crawler) GetDomainVerifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.DomainVerified) {
		return nil, false
	}
	return o.DomainVerified, true
}

// HasDomainVerified returns a boolean if a field has been set.
func (o *Crawler) HasDomainVerified() bool {
	if o != nil && !IsNil(o.DomainVerified) {
		return true
	}

	return false
}

// SetDomainVerified gets a reference to the given int32 and assigns it to the DomainVerified field.
func (o *Crawler) SetDomainVerified(v int32) {
	o.DomainVerified = &v
}

func (o Crawler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Crawler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.UrlsList) {
		toSerialize["urls_list"] = o.UrlsList
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.DomainVerified) {
		toSerialize["domain_verified"] = o.DomainVerified
	}
	return toSerialize, nil
}

type NullableCrawler struct {
	value *Crawler
	isSet bool
}

func (v NullableCrawler) Get() *Crawler {
	return v.value
}

func (v *NullableCrawler) Set(val *Crawler) {
	v.value = val
	v.isSet = true
}

func (v NullableCrawler) IsSet() bool {
	return v.isSet
}

func (v *NullableCrawler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrawler(val *Crawler) *NullableCrawler {
	return &NullableCrawler{value: val, isSet: true}
}

func (v NullableCrawler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrawler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

