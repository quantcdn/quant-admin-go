/*
QuantCDN Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain struct for Domain
type Domain struct {
	Id int32 `json:"id"`
	Domain string `json:"domain"`
	ProjectId int32 `json:"project_id"`
	InSection int32 `json:"in_section"`
	DnsEngaged int32 `json:"dns_engaged"`
	SectionMessage *string `json:"section_message,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type _Domain Domain

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain(id int32, domain string, projectId int32, inSection int32, dnsEngaged int32) *Domain {
	this := Domain{}
	this.Id = id
	this.Domain = domain
	this.ProjectId = projectId
	this.InSection = inSection
	this.DnsEngaged = dnsEngaged
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetId returns the Id field value
func (o *Domain) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Domain) SetId(v int32) {
	o.Id = v
}

// GetDomain returns the Domain field value
func (o *Domain) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *Domain) SetDomain(v string) {
	o.Domain = v
}

// GetProjectId returns the ProjectId field value
func (o *Domain) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Domain) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Domain) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetInSection returns the InSection field value
func (o *Domain) GetInSection() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InSection
}

// GetInSectionOk returns a tuple with the InSection field value
// and a boolean to check if the value has been set.
func (o *Domain) GetInSectionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InSection, true
}

// SetInSection sets field value
func (o *Domain) SetInSection(v int32) {
	o.InSection = v
}

// GetDnsEngaged returns the DnsEngaged field value
func (o *Domain) GetDnsEngaged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DnsEngaged
}

// GetDnsEngagedOk returns a tuple with the DnsEngaged field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDnsEngagedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsEngaged, true
}

// SetDnsEngaged sets field value
func (o *Domain) SetDnsEngaged(v int32) {
	o.DnsEngaged = v
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

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["domain"] = o.Domain
	toSerialize["project_id"] = o.ProjectId
	toSerialize["in_section"] = o.InSection
	toSerialize["dns_engaged"] = o.DnsEngaged
	if !IsNil(o.SectionMessage) {
		toSerialize["section_message"] = o.SectionMessage
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
	return toSerialize, nil
}

func (o *Domain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"domain",
		"project_id",
		"in_section",
		"dns_engaged",
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

	varDomain := _Domain{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDomain)

	if err != nil {
		return err
	}

	*o = Domain(varDomain)

	return err
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


