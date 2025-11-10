# V2RuleFunctionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**FnUuid** | **string** | Function UUID | 

## Methods

### NewV2RuleFunctionAction

`func NewV2RuleFunctionAction(message string, error_ bool, fnUuid string, ) *V2RuleFunctionAction`

NewV2RuleFunctionAction instantiates a new V2RuleFunctionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleFunctionActionWithDefaults

`func NewV2RuleFunctionActionWithDefaults() *V2RuleFunctionAction`

NewV2RuleFunctionActionWithDefaults instantiates a new V2RuleFunctionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleFunctionAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleFunctionAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleFunctionAction) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleFunctionAction) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleFunctionAction) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleFunctionAction) SetError(v bool)`

SetError sets Error field to given value.


### GetFnUuid

`func (o *V2RuleFunctionAction) GetFnUuid() string`

GetFnUuid returns the FnUuid field if non-nil, zero value otherwise.

### GetFnUuidOk

`func (o *V2RuleFunctionAction) GetFnUuidOk() (*string, bool)`

GetFnUuidOk returns a tuple with the FnUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFnUuid

`func (o *V2RuleFunctionAction) SetFnUuid(v string)`

SetFnUuid sets FnUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


