# V2RuleCustomResponseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**CustomResponseBody** | **string** | Custom response body content | 
**CustomResponseStatusCode** | Pointer to **int32** | HTTP status code for custom response | [optional] [default to 200]

## Methods

### NewV2RuleCustomResponseAction

`func NewV2RuleCustomResponseAction(message string, error_ bool, customResponseBody string, ) *V2RuleCustomResponseAction`

NewV2RuleCustomResponseAction instantiates a new V2RuleCustomResponseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleCustomResponseActionWithDefaults

`func NewV2RuleCustomResponseActionWithDefaults() *V2RuleCustomResponseAction`

NewV2RuleCustomResponseActionWithDefaults instantiates a new V2RuleCustomResponseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleCustomResponseAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleCustomResponseAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleCustomResponseAction) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleCustomResponseAction) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleCustomResponseAction) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleCustomResponseAction) SetError(v bool)`

SetError sets Error field to given value.


### GetCustomResponseBody

`func (o *V2RuleCustomResponseAction) GetCustomResponseBody() string`

GetCustomResponseBody returns the CustomResponseBody field if non-nil, zero value otherwise.

### GetCustomResponseBodyOk

`func (o *V2RuleCustomResponseAction) GetCustomResponseBodyOk() (*string, bool)`

GetCustomResponseBodyOk returns a tuple with the CustomResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseBody

`func (o *V2RuleCustomResponseAction) SetCustomResponseBody(v string)`

SetCustomResponseBody sets CustomResponseBody field to given value.


### GetCustomResponseStatusCode

`func (o *V2RuleCustomResponseAction) GetCustomResponseStatusCode() int32`

GetCustomResponseStatusCode returns the CustomResponseStatusCode field if non-nil, zero value otherwise.

### GetCustomResponseStatusCodeOk

`func (o *V2RuleCustomResponseAction) GetCustomResponseStatusCodeOk() (*int32, bool)`

GetCustomResponseStatusCodeOk returns a tuple with the CustomResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomResponseStatusCode

`func (o *V2RuleCustomResponseAction) SetCustomResponseStatusCode(v int32)`

SetCustomResponseStatusCode sets CustomResponseStatusCode field to given value.

### HasCustomResponseStatusCode

`func (o *V2RuleCustomResponseAction) HasCustomResponseStatusCode() bool`

HasCustomResponseStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


