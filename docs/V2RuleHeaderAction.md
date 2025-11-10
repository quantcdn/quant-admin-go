# V2RuleHeaderAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**Headers** | **map[string]string** | Headers to set | 

## Methods

### NewV2RuleHeaderAction

`func NewV2RuleHeaderAction(message string, error_ bool, headers map[string]string, ) *V2RuleHeaderAction`

NewV2RuleHeaderAction instantiates a new V2RuleHeaderAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleHeaderActionWithDefaults

`func NewV2RuleHeaderActionWithDefaults() *V2RuleHeaderAction`

NewV2RuleHeaderActionWithDefaults instantiates a new V2RuleHeaderAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleHeaderAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleHeaderAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleHeaderAction) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleHeaderAction) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleHeaderAction) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleHeaderAction) SetError(v bool)`

SetError sets Error field to given value.


### GetHeaders

`func (o *V2RuleHeaderAction) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *V2RuleHeaderAction) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *V2RuleHeaderAction) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


