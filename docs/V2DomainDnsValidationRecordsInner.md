# V2DomainDnsValidationRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | DNS record name (host/subdomain) | [optional] 
**Type** | Pointer to **string** | DNS record type (typically CNAME) | [optional] 
**Value** | Pointer to **string** | DNS record value to point to | [optional] 

## Methods

### NewV2DomainDnsValidationRecordsInner

`func NewV2DomainDnsValidationRecordsInner() *V2DomainDnsValidationRecordsInner`

NewV2DomainDnsValidationRecordsInner instantiates a new V2DomainDnsValidationRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2DomainDnsValidationRecordsInnerWithDefaults

`func NewV2DomainDnsValidationRecordsInnerWithDefaults() *V2DomainDnsValidationRecordsInner`

NewV2DomainDnsValidationRecordsInnerWithDefaults instantiates a new V2DomainDnsValidationRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2DomainDnsValidationRecordsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2DomainDnsValidationRecordsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2DomainDnsValidationRecordsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2DomainDnsValidationRecordsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *V2DomainDnsValidationRecordsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V2DomainDnsValidationRecordsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V2DomainDnsValidationRecordsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V2DomainDnsValidationRecordsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *V2DomainDnsValidationRecordsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2DomainDnsValidationRecordsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2DomainDnsValidationRecordsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V2DomainDnsValidationRecordsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


