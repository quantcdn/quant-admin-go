# V2Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Domain ID | 
**Domain** | **string** | Domain name | 
**DnsEngaged** | **int32** | DNS engagement status | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


