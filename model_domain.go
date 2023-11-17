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

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain struct for Domain
type Domain struct {
	Id *int32 `json:"id,omitempty"`
	Domain *string `json:"domain,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	ProjectId *int32 `json:"project_id,omitempty"`
	InSection *int32 `json:"in_section,omitempty"`
	SectionMessage *string `json:"section_message,omitempty"`
	DnsEngaged *int32 `json:"dns_engaged,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Domain) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Domain) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Domain) SetId(v int32) {
	o.Id = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Domain) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Domain) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Domain) SetDomain(v string) {
	o.Domain = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Domain) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Domain) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Domain) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Domain) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Domain) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Domain) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Domain) GetDeletedAt() string {
	if o == nil || IsNil(o.DeletedAt) {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDeletedAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Domain) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Domain) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Domain) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Domain) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *Domain) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetInSection returns the InSection field value if set, zero value otherwise.
func (o *Domain) GetInSection() int32 {
	if o == nil || IsNil(o.InSection) {
		var ret int32
		return ret
	}
	return *o.InSection
}

// GetInSectionOk returns a tuple with the InSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetInSectionOk() (*int32, bool) {
	if o == nil || IsNil(o.InSection) {
		return nil, false
	}
	return o.InSection, true
}

// HasInSection returns a boolean if a field has been set.
func (o *Domain) HasInSection() bool {
	if o != nil && !IsNil(o.InSection) {
		return true
	}

	return false
}

// SetInSection gets a reference to the given int32 and assigns it to the InSection field.
func (o *Domain) SetInSection(v int32) {
	o.InSection = &v
}

// GetSectionMessage returns the SectionMessage field value if set, zero value otherwise.
func (o *Domain) GetSectionMessage() string {
	if o == nil || IsNil(o.SectionMessage) {
		var ret string
		return ret
	}
	return *o.SectionMessage
}

// GetSectionMessageOk returns a tuple with the SectionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetSectionMessageOk() (*string, bool) {
	if o == nil || IsNil(o.SectionMessage) {
		return nil, false
	}
	return o.SectionMessage, true
}

// HasSectionMessage returns a boolean if a field has been set.
func (o *Domain) HasSectionMessage() bool {
	if o != nil && !IsNil(o.SectionMessage) {
		return true
	}

	return false
}

// SetSectionMessage gets a reference to the given string and assigns it to the SectionMessage field.
func (o *Domain) SetSectionMessage(v string) {
	o.SectionMessage = &v
}

// GetDnsEngaged returns the DnsEngaged field value if set, zero value otherwise.
func (o *Domain) GetDnsEngaged() int32 {
	if o == nil || IsNil(o.DnsEngaged) {
		var ret int32
		return ret
	}
	return *o.DnsEngaged
}

// GetDnsEngagedOk returns a tuple with the DnsEngaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDnsEngagedOk() (*int32, bool) {
	if o == nil || IsNil(o.DnsEngaged) {
		return nil, false
	}
	return o.DnsEngaged, true
}

// HasDnsEngaged returns a boolean if a field has been set.
func (o *Domain) HasDnsEngaged() bool {
	if o != nil && !IsNil(o.DnsEngaged) {
		return true
	}

	return false
}

// SetDnsEngaged gets a reference to the given int32 and assigns it to the DnsEngaged field.
func (o *Domain) SetDnsEngaged(v int32) {
	o.DnsEngaged = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
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
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.InSection) {
		toSerialize["in_section"] = o.InSection
	}
	if !IsNil(o.SectionMessage) {
		toSerialize["section_message"] = o.SectionMessage
	}
	if !IsNil(o.DnsEngaged) {
		toSerialize["dns_engaged"] = o.DnsEngaged
	}
	return toSerialize, nil
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


