# V1Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **bool** |  | 
**Message** | **string** |  | 

## Methods

### NewV1Error

`func NewV1Error(error_ bool, message string, ) *V1Error`

NewV1Error instantiates a new V1Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ErrorWithDefaults

`func NewV1ErrorWithDefaults() *V1Error`

NewV1ErrorWithDefaults instantiates a new V1Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V1Error) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1Error) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1Error) SetError(v bool)`

SetError sets Error field to given value.


### GetMessage

`func (o *V1Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1Error) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


