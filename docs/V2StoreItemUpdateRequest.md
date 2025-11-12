# V2StoreItemUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Item value (can be JSON string) | 
**Secret** | Pointer to **bool** | Store as secret with KMS encryption. Note: Encryption status cannot be changed after initial creation - this value is preserved from the original item. | [optional] 

## Methods

### NewV2StoreItemUpdateRequest

`func NewV2StoreItemUpdateRequest(value string, ) *V2StoreItemUpdateRequest`

NewV2StoreItemUpdateRequest instantiates a new V2StoreItemUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2StoreItemUpdateRequestWithDefaults

`func NewV2StoreItemUpdateRequestWithDefaults() *V2StoreItemUpdateRequest`

NewV2StoreItemUpdateRequestWithDefaults instantiates a new V2StoreItemUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *V2StoreItemUpdateRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2StoreItemUpdateRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2StoreItemUpdateRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecret

`func (o *V2StoreItemUpdateRequest) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V2StoreItemUpdateRequest) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V2StoreItemUpdateRequest) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *V2StoreItemUpdateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


