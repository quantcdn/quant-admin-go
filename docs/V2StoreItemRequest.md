# V2StoreItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Key** | **string** | Item key | 
**Value** | **string** | Item value (can be JSON string) | 

## Methods

### NewV2StoreItemRequest

`func NewV2StoreItemRequest(message string, error_ bool, key string, value string, ) *V2StoreItemRequest`

NewV2StoreItemRequest instantiates a new V2StoreItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2StoreItemRequestWithDefaults

`func NewV2StoreItemRequestWithDefaults() *V2StoreItemRequest`

NewV2StoreItemRequestWithDefaults instantiates a new V2StoreItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2StoreItemRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2StoreItemRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2StoreItemRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2StoreItemRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2StoreItemRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2StoreItemRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetKey

`func (o *V2StoreItemRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2StoreItemRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2StoreItemRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *V2StoreItemRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2StoreItemRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2StoreItemRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


