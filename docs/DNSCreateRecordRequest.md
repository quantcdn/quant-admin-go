# DNSCreateRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name relative to the zone; @ denotes the apex | 
**Type** | **string** |  | 
**Value** | **string** |  | 
**Ttl** | Pointer to **int32** |  | [optional] [default to 300]

## Methods

### NewDNSCreateRecordRequest

`func NewDNSCreateRecordRequest(name string, type_ string, value string, ) *DNSCreateRecordRequest`

NewDNSCreateRecordRequest instantiates a new DNSCreateRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSCreateRecordRequestWithDefaults

`func NewDNSCreateRecordRequestWithDefaults() *DNSCreateRecordRequest`

NewDNSCreateRecordRequestWithDefaults instantiates a new DNSCreateRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DNSCreateRecordRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSCreateRecordRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSCreateRecordRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DNSCreateRecordRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSCreateRecordRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSCreateRecordRequest) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *DNSCreateRecordRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DNSCreateRecordRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DNSCreateRecordRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetTtl

`func (o *DNSCreateRecordRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSCreateRecordRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSCreateRecordRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSCreateRecordRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


