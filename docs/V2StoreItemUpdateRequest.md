# V2StoreItemUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Value** | **string** | Item value (can be JSON string) | 

## Methods

### NewV2StoreItemUpdateRequest

`func NewV2StoreItemUpdateRequest(message string, error_ bool, value string, ) *V2StoreItemUpdateRequest`

NewV2StoreItemUpdateRequest instantiates a new V2StoreItemUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2StoreItemUpdateRequestWithDefaults

`func NewV2StoreItemUpdateRequestWithDefaults() *V2StoreItemUpdateRequest`

NewV2StoreItemUpdateRequestWithDefaults instantiates a new V2StoreItemUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2StoreItemUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2StoreItemUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2StoreItemUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2StoreItemUpdateRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2StoreItemUpdateRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2StoreItemUpdateRequest) SetError(v bool)`

SetError sets Error field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


