# V2DomainDnsGoLiveRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | DNS record type (CNAME, A, or ALIAS) | [optional] 
**Name** | Pointer to **string** | DNS record name/host (@ for apex/root domains, subdomain name for subdomains) | [optional] 
**Value** | Pointer to **NullableString** | DNS record value (IP addresses for A records, domain name for CNAME/ALIAS) | [optional] 
**Description** | Pointer to **string** | Human-readable instructions for configuring this DNS record | [optional] 

## Methods

### NewV2DomainDnsGoLiveRecordsInner

`func NewV2DomainDnsGoLiveRecordsInner() *V2DomainDnsGoLiveRecordsInner`

NewV2DomainDnsGoLiveRecordsInner instantiates a new V2DomainDnsGoLiveRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2DomainDnsGoLiveRecordsInnerWithDefaults

`func NewV2DomainDnsGoLiveRecordsInnerWithDefaults() *V2DomainDnsGoLiveRecordsInner`

NewV2DomainDnsGoLiveRecordsInnerWithDefaults instantiates a new V2DomainDnsGoLiveRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V2DomainDnsGoLiveRecordsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2DomainDnsGoLiveRecordsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2DomainDnsGoLiveRecordsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V2DomainDnsGoLiveRecordsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *V2DomainDnsGoLiveRecordsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2DomainDnsGoLiveRecordsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2DomainDnsGoLiveRecordsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2DomainDnsGoLiveRecordsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *V2DomainDnsGoLiveRecordsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2DomainDnsGoLiveRecordsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2DomainDnsGoLiveRecordsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V2DomainDnsGoLiveRecordsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *V2DomainDnsGoLiveRecordsInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *V2DomainDnsGoLiveRecordsInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDescription

`func (o *V2DomainDnsGoLiveRecordsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2DomainDnsGoLiveRecordsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2DomainDnsGoLiveRecordsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2DomainDnsGoLiveRecordsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


