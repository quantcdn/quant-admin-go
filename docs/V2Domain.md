# V2Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Domain ID | 
**Domain** | **string** | Domain name | 
**DnsEngaged** | **int32** | DNS engagement status. 1 indicates DNS is properly configured and engaged, 0 indicates DNS configuration is pending or incomplete. | 
**DnsValidationRecords** | Pointer to [**[]V2DomainDnsValidationRecordsInner**](V2DomainDnsValidationRecordsInner.md) | DNS validation records required for SSL certificate validation. Present for domains pending certificate validation. Each record contains the CNAME information needed to validate domain ownership. | [optional] 
**DnsGoLiveRecords** | Pointer to [**[]V2DomainDnsGoLiveRecordsInner**](V2DomainDnsGoLiveRecordsInner.md) | DNS records required to route traffic to the CDN. These records differ based on domain type (apex vs subdomain). Present when the CDN is configured and ready to receive traffic. | [optional] 

## Methods

### NewV2Domain

`func NewV2Domain(id int32, domain string, dnsEngaged int32, ) *V2Domain`

NewV2Domain instantiates a new V2Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2DomainWithDefaults

`func NewV2DomainWithDefaults() *V2Domain`

NewV2DomainWithDefaults instantiates a new V2Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2Domain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2Domain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2Domain) SetId(v int32)`

SetId sets Id field to given value.


### GetDomain

`func (o *V2Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V2Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V2Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDnsEngaged

`func (o *V2Domain) GetDnsEngaged() int32`

GetDnsEngaged returns the DnsEngaged field if non-nil, zero value otherwise.

### GetDnsEngagedOk

`func (o *V2Domain) GetDnsEngagedOk() (*int32, bool)`

GetDnsEngagedOk returns a tuple with the DnsEngaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEngaged

`func (o *V2Domain) SetDnsEngaged(v int32)`

SetDnsEngaged sets DnsEngaged field to given value.


### GetDnsValidationRecords

`func (o *V2Domain) GetDnsValidationRecords() []V2DomainDnsValidationRecordsInner`

GetDnsValidationRecords returns the DnsValidationRecords field if non-nil, zero value otherwise.

### GetDnsValidationRecordsOk

`func (o *V2Domain) GetDnsValidationRecordsOk() (*[]V2DomainDnsValidationRecordsInner, bool)`

GetDnsValidationRecordsOk returns a tuple with the DnsValidationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsValidationRecords

`func (o *V2Domain) SetDnsValidationRecords(v []V2DomainDnsValidationRecordsInner)`

SetDnsValidationRecords sets DnsValidationRecords field to given value.

### HasDnsValidationRecords

`func (o *V2Domain) HasDnsValidationRecords() bool`

HasDnsValidationRecords returns a boolean if a field has been set.

### SetDnsValidationRecordsNil

`func (o *V2Domain) SetDnsValidationRecordsNil(b bool)`

 SetDnsValidationRecordsNil sets the value for DnsValidationRecords to be an explicit nil

### UnsetDnsValidationRecords
`func (o *V2Domain) UnsetDnsValidationRecords()`

UnsetDnsValidationRecords ensures that no value is present for DnsValidationRecords, not even an explicit nil
### GetDnsGoLiveRecords

`func (o *V2Domain) GetDnsGoLiveRecords() []V2DomainDnsGoLiveRecordsInner`

GetDnsGoLiveRecords returns the DnsGoLiveRecords field if non-nil, zero value otherwise.

### GetDnsGoLiveRecordsOk

`func (o *V2Domain) GetDnsGoLiveRecordsOk() (*[]V2DomainDnsGoLiveRecordsInner, bool)`

GetDnsGoLiveRecordsOk returns a tuple with the DnsGoLiveRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsGoLiveRecords

`func (o *V2Domain) SetDnsGoLiveRecords(v []V2DomainDnsGoLiveRecordsInner)`

SetDnsGoLiveRecords sets DnsGoLiveRecords field to given value.

### HasDnsGoLiveRecords

`func (o *V2Domain) HasDnsGoLiveRecords() bool`

HasDnsGoLiveRecords returns a boolean if a field has been set.

### SetDnsGoLiveRecordsNil

`func (o *V2Domain) SetDnsGoLiveRecordsNil(b bool)`

 SetDnsGoLiveRecordsNil sets the value for DnsGoLiveRecords to be an explicit nil

### UnsetDnsGoLiveRecords
`func (o *V2Domain) UnsetDnsGoLiveRecords()`

UnsetDnsGoLiveRecords ensures that no value is present for DnsGoLiveRecords, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


