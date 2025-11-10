# V2SecretStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Name** | **string** | Secret store name | 

## Methods

### NewV2SecretStoreRequest

`func NewV2SecretStoreRequest(message string, error_ bool, name string, ) *V2SecretStoreRequest`

NewV2SecretStoreRequest instantiates a new V2SecretStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SecretStoreRequestWithDefaults

`func NewV2SecretStoreRequestWithDefaults() *V2SecretStoreRequest`

NewV2SecretStoreRequestWithDefaults instantiates a new V2SecretStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2SecretStoreRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2SecretStoreRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2SecretStoreRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2SecretStoreRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2SecretStoreRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2SecretStoreRequest) SetError(v bool)`

SetError sets Error field to given value.


### GetName

`func (o *V2SecretStoreRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2SecretStoreRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2SecretStoreRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


