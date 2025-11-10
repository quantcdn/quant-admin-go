# V2StoreItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Key** | **string** | Item key | 
**Value** | **string** | Item value (can be JSON string) | 

## Methods

### NewV2StoreItem

`func NewV2StoreItem(message string, error_ bool, key string, value string, ) *V2StoreItem`

NewV2StoreItem instantiates a new V2StoreItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2StoreItemWithDefaults

`func NewV2StoreItemWithDefaults() *V2StoreItem`

NewV2StoreItemWithDefaults instantiates a new V2StoreItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2StoreItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2StoreItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2StoreItem) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2StoreItem) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2StoreItem) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2StoreItem) SetError(v bool)`

SetError sets Error field to given value.


### GetKey

`func (o *V2StoreItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V2StoreItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V2StoreItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *V2StoreItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V2StoreItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V2StoreItem) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


