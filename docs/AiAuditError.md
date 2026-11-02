# AiAuditError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **string** |  | 
**Code** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 

## Methods

### NewAiAuditError

`func NewAiAuditError(message string, error_ string, ) *AiAuditError`

NewAiAuditError instantiates a new AiAuditError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAuditErrorWithDefaults

`func NewAiAuditErrorWithDefaults() *AiAuditError`

NewAiAuditErrorWithDefaults instantiates a new AiAuditError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiAuditError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiAuditError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiAuditError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *AiAuditError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiAuditError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiAuditError) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *AiAuditError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiAuditError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiAuditError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AiAuditError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRetryable

`func (o *AiAuditError) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *AiAuditError) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *AiAuditError) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *AiAuditError) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


