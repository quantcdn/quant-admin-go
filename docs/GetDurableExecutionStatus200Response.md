# GetDurableExecutionStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** |  | [optional] 
**ExecutionArn** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AwsStatus** | Pointer to **string** | Raw AWS Step Functions status | [optional] 
**CallbackId** | Pointer to **string** | Present when status is waiting_callback - use with /chat/callback | [optional] 
**PendingTools** | Pointer to [**[]GetDurableExecutionStatus200ResponsePendingToolsInner**](GetDurableExecutionStatus200ResponsePendingToolsInner.md) | Present when status is waiting_callback - tools waiting for results | [optional] 
**Result** | Pointer to [**GetDurableExecutionStatus200ResponseResult**](GetDurableExecutionStatus200ResponseResult.md) |  | [optional] 
**Error** | Pointer to [**GetDurableExecutionStatus200ResponseError**](GetDurableExecutionStatus200ResponseError.md) |  | [optional] 

## Methods

### NewGetDurableExecutionStatus200Response

`func NewGetDurableExecutionStatus200Response() *GetDurableExecutionStatus200Response`

NewGetDurableExecutionStatus200Response instantiates a new GetDurableExecutionStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDurableExecutionStatus200ResponseWithDefaults

`func NewGetDurableExecutionStatus200ResponseWithDefaults() *GetDurableExecutionStatus200Response`

NewGetDurableExecutionStatus200ResponseWithDefaults instantiates a new GetDurableExecutionStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *GetDurableExecutionStatus200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetDurableExecutionStatus200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetDurableExecutionStatus200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetDurableExecutionStatus200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetExecutionArn

`func (o *GetDurableExecutionStatus200Response) GetExecutionArn() string`

GetExecutionArn returns the ExecutionArn field if non-nil, zero value otherwise.

### GetExecutionArnOk

`func (o *GetDurableExecutionStatus200Response) GetExecutionArnOk() (*string, bool)`

GetExecutionArnOk returns a tuple with the ExecutionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionArn

`func (o *GetDurableExecutionStatus200Response) SetExecutionArn(v string)`

SetExecutionArn sets ExecutionArn field to given value.

### HasExecutionArn

`func (o *GetDurableExecutionStatus200Response) HasExecutionArn() bool`

HasExecutionArn returns a boolean if a field has been set.

### GetStatus

`func (o *GetDurableExecutionStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDurableExecutionStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDurableExecutionStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDurableExecutionStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAwsStatus

`func (o *GetDurableExecutionStatus200Response) GetAwsStatus() string`

GetAwsStatus returns the AwsStatus field if non-nil, zero value otherwise.

### GetAwsStatusOk

`func (o *GetDurableExecutionStatus200Response) GetAwsStatusOk() (*string, bool)`

GetAwsStatusOk returns a tuple with the AwsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStatus

`func (o *GetDurableExecutionStatus200Response) SetAwsStatus(v string)`

SetAwsStatus sets AwsStatus field to given value.

### HasAwsStatus

`func (o *GetDurableExecutionStatus200Response) HasAwsStatus() bool`

HasAwsStatus returns a boolean if a field has been set.

### GetCallbackId

`func (o *GetDurableExecutionStatus200Response) GetCallbackId() string`

GetCallbackId returns the CallbackId field if non-nil, zero value otherwise.

### GetCallbackIdOk

`func (o *GetDurableExecutionStatus200Response) GetCallbackIdOk() (*string, bool)`

GetCallbackIdOk returns a tuple with the CallbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackId

`func (o *GetDurableExecutionStatus200Response) SetCallbackId(v string)`

SetCallbackId sets CallbackId field to given value.

### HasCallbackId

`func (o *GetDurableExecutionStatus200Response) HasCallbackId() bool`

HasCallbackId returns a boolean if a field has been set.

### GetPendingTools

`func (o *GetDurableExecutionStatus200Response) GetPendingTools() []GetDurableExecutionStatus200ResponsePendingToolsInner`

GetPendingTools returns the PendingTools field if non-nil, zero value otherwise.

### GetPendingToolsOk

`func (o *GetDurableExecutionStatus200Response) GetPendingToolsOk() (*[]GetDurableExecutionStatus200ResponsePendingToolsInner, bool)`

GetPendingToolsOk returns a tuple with the PendingTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTools

`func (o *GetDurableExecutionStatus200Response) SetPendingTools(v []GetDurableExecutionStatus200ResponsePendingToolsInner)`

SetPendingTools sets PendingTools field to given value.

### HasPendingTools

`func (o *GetDurableExecutionStatus200Response) HasPendingTools() bool`

HasPendingTools returns a boolean if a field has been set.

### GetResult

`func (o *GetDurableExecutionStatus200Response) GetResult() GetDurableExecutionStatus200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDurableExecutionStatus200Response) GetResultOk() (*GetDurableExecutionStatus200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDurableExecutionStatus200Response) SetResult(v GetDurableExecutionStatus200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDurableExecutionStatus200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *GetDurableExecutionStatus200Response) GetError() GetDurableExecutionStatus200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetDurableExecutionStatus200Response) GetErrorOk() (*GetDurableExecutionStatus200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetDurableExecutionStatus200Response) SetError(v GetDurableExecutionStatus200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *GetDurableExecutionStatus200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


