# SubmitToolCallbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackId** | **string** | The callbackId from the waiting_callback status response | 
**ToolResults** | [**[]SubmitToolCallbackRequestToolResultsInner**](SubmitToolCallbackRequestToolResultsInner.md) | Results of client-executed tools | 

## Methods

### NewSubmitToolCallbackRequest

`func NewSubmitToolCallbackRequest(callbackId string, toolResults []SubmitToolCallbackRequestToolResultsInner, ) *SubmitToolCallbackRequest`

NewSubmitToolCallbackRequest instantiates a new SubmitToolCallbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitToolCallbackRequestWithDefaults

`func NewSubmitToolCallbackRequestWithDefaults() *SubmitToolCallbackRequest`

NewSubmitToolCallbackRequestWithDefaults instantiates a new SubmitToolCallbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackId

`func (o *SubmitToolCallbackRequest) GetCallbackId() string`

GetCallbackId returns the CallbackId field if non-nil, zero value otherwise.

### GetCallbackIdOk

`func (o *SubmitToolCallbackRequest) GetCallbackIdOk() (*string, bool)`

GetCallbackIdOk returns a tuple with the CallbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackId

`func (o *SubmitToolCallbackRequest) SetCallbackId(v string)`

SetCallbackId sets CallbackId field to given value.


### GetToolResults

`func (o *SubmitToolCallbackRequest) GetToolResults() []SubmitToolCallbackRequestToolResultsInner`

GetToolResults returns the ToolResults field if non-nil, zero value otherwise.

### GetToolResultsOk

`func (o *SubmitToolCallbackRequest) GetToolResultsOk() (*[]SubmitToolCallbackRequestToolResultsInner, bool)`

GetToolResultsOk returns a tuple with the ToolResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolResults

`func (o *SubmitToolCallbackRequest) SetToolResults(v []SubmitToolCallbackRequestToolResultsInner)`

SetToolResults sets ToolResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


