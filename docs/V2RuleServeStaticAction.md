# V2RuleServeStaticAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | 
**Error** | **bool** | Error flag | 
**StaticFilePath** | **string** | Path to the static file to serve | 

## Methods

### NewV2RuleServeStaticAction

`func NewV2RuleServeStaticAction(message string, error_ bool, staticFilePath string, ) *V2RuleServeStaticAction`

NewV2RuleServeStaticAction instantiates a new V2RuleServeStaticAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RuleServeStaticActionWithDefaults

`func NewV2RuleServeStaticActionWithDefaults() *V2RuleServeStaticAction`

NewV2RuleServeStaticActionWithDefaults instantiates a new V2RuleServeStaticAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *V2RuleServeStaticAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V2RuleServeStaticAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V2RuleServeStaticAction) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *V2RuleServeStaticAction) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V2RuleServeStaticAction) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V2RuleServeStaticAction) SetError(v bool)`

SetError sets Error field to given value.


### GetStaticFilePath

`func (o *V2RuleServeStaticAction) GetStaticFilePath() string`

GetStaticFilePath returns the StaticFilePath field if non-nil, zero value otherwise.

### GetStaticFilePathOk

`func (o *V2RuleServeStaticAction) GetStaticFilePathOk() (*string, bool)`

GetStaticFilePathOk returns a tuple with the StaticFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticFilePath

`func (o *V2RuleServeStaticAction) SetStaticFilePath(v string)`

SetStaticFilePath sets StaticFilePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


