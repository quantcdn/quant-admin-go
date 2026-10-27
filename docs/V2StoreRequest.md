# V2StoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Store name | 
**Format** | Pointer to **int32** | Store format. Format 1 is a plain key-value store. Format 2 enables indexes, ledgers, conditional writes, transactions and the cursor key list. The format is chosen when the store is created and cannot be changed by this request; omit the field for format 1. | [optional] [default to 1]

## Methods

### NewV2StoreRequest

`func NewV2StoreRequest(name string, ) *V2StoreRequest`

NewV2StoreRequest instantiates a new V2StoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2StoreRequestWithDefaults

`func NewV2StoreRequestWithDefaults() *V2StoreRequest`

NewV2StoreRequestWithDefaults instantiates a new V2StoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2StoreRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2StoreRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2StoreRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFormat

`func (o *V2StoreRequest) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2StoreRequest) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2StoreRequest) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *V2StoreRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


