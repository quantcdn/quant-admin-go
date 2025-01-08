# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Domain** | **string** |  | 
**ProjectId** | Pointer to **int32** |  | [optional] 
**InSection** | Pointer to **int32** |  | [optional] 
**DnsEngaged** | **int32** |  | 
**SectionMessage** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDomain

`func NewDomain(id int32, domain string, dnsEngaged int32, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v int32)`

SetId sets Id field to given value.


### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetProjectId

`func (o *Domain) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Domain) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Domain) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Domain) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetInSection

`func (o *Domain) GetInSection() int32`

GetInSection returns the InSection field if non-nil, zero value otherwise.

### GetInSectionOk

`func (o *Domain) GetInSectionOk() (*int32, bool)`

GetInSectionOk returns a tuple with the InSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSection

`func (o *Domain) SetInSection(v int32)`

SetInSection sets InSection field to given value.

### HasInSection

`func (o *Domain) HasInSection() bool`

HasInSection returns a boolean if a field has been set.

### GetDnsEngaged

`func (o *Domain) GetDnsEngaged() int32`

GetDnsEngaged returns the DnsEngaged field if non-nil, zero value otherwise.

### GetDnsEngagedOk

`func (o *Domain) GetDnsEngagedOk() (*int32, bool)`

GetDnsEngagedOk returns a tuple with the DnsEngaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEngaged

`func (o *Domain) SetDnsEngaged(v int32)`

SetDnsEngaged sets DnsEngaged field to given value.


### GetSectionMessage

`func (o *Domain) GetSectionMessage() string`

GetSectionMessage returns the SectionMessage field if non-nil, zero value otherwise.

### GetSectionMessageOk

`func (o *Domain) GetSectionMessageOk() (*string, bool)`

GetSectionMessageOk returns a tuple with the SectionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionMessage

`func (o *Domain) SetSectionMessage(v string)`

SetSectionMessage sets SectionMessage field to given value.

### HasSectionMessage

`func (o *Domain) HasSectionMessage() bool`

HasSectionMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Domain) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Domain) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Domain) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Domain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Domain) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Domain) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Domain) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Domain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Domain) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Domain) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Domain) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Domain) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


