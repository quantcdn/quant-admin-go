# V2RuleAuthAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**AuthUser** | **string** | Authentication username | 
**AuthPass** | **string** | Authentication password | 

## Methods

### NewV2RuleAuthAction

`func NewV2RuleAuthAction(message string, error_ bool, authUser string, authPass string, ) *V2RuleAuthAction`

NewV2RuleAuthAction instantiates a new V2RuleAuthAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleAuthActionWithDefaults

`func NewV2RuleAuthActionWithDefaults() *V2RuleAuthAction`

NewV2RuleAuthActionWithDefaults instantiates a new V2RuleAuthAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleAuthAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleAuthAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleAuthAction) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleAuthAction) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleAuthAction) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleAuthAction) SetError(v bool)`

SetError sets Error field to given value.


### GetAuthUser

`func (o *V2RuleAuthAction) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *V2RuleAuthAction) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *V2RuleAuthAction) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.


### GetAuthPass

`func (o *V2RuleAuthAction) GetAuthPass() string`

GetAuthPass returns the AuthPass field if non-nil, zero value otherwise.

### GetAuthPassOk

`func (o *V2RuleAuthAction) GetAuthPassOk() (*string, bool)`

GetAuthPassOk returns a tuple with the AuthPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPass

`func (o *V2RuleAuthAction) SetAuthPass(v string)`

SetAuthPass sets AuthPass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


