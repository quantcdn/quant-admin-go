# CreateCustomTool201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Tool** | Pointer to **map[string]interface{}** |  | [optional] 
**EdgeFunctionUrl** | Pointer to **string** | Computed edge function URL (read-only) | [optional] 
**EdgeFunctionCode** | Pointer to **string** | The deployed edge function code | [optional] 
**IsUpdate** | Pointer to **bool** | Whether this was an update to an existing tool | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateCustomTool201Response

`func NewCreateCustomTool201Response() *CreateCustomTool201Response`

NewCreateCustomTool201Response instantiates a new CreateCustomTool201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomTool201ResponseWithDefaults

`func NewCreateCustomTool201ResponseWithDefaults() *CreateCustomTool201Response`

NewCreateCustomTool201ResponseWithDefaults instantiates a new CreateCustomTool201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateCustomTool201Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateCustomTool201Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateCustomTool201Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateCustomTool201Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTool

`func (o *CreateCustomTool201Response) GetTool() map[string]interface{}`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *CreateCustomTool201Response) GetToolOk() (*map[string]interface{}, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *CreateCustomTool201Response) SetTool(v map[string]interface{})`

SetTool sets Tool field to given value.

### HasTool

`func (o *CreateCustomTool201Response) HasTool() bool`

HasTool returns a boolean if a field has been set.

### GetEdgeFunctionUrl

`func (o *CreateCustomTool201Response) GetEdgeFunctionUrl() string`

GetEdgeFunctionUrl returns the EdgeFunctionUrl field if non-nil, zero value otherwise.

### GetEdgeFunctionUrlOk

`func (o *CreateCustomTool201Response) GetEdgeFunctionUrlOk() (*string, bool)`

GetEdgeFunctionUrlOk returns a tuple with the EdgeFunctionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionUrl

`func (o *CreateCustomTool201Response) SetEdgeFunctionUrl(v string)`

SetEdgeFunctionUrl sets EdgeFunctionUrl field to given value.

### HasEdgeFunctionUrl

`func (o *CreateCustomTool201Response) HasEdgeFunctionUrl() bool`

HasEdgeFunctionUrl returns a boolean if a field has been set.

### GetEdgeFunctionCode

`func (o *CreateCustomTool201Response) GetEdgeFunctionCode() string`

GetEdgeFunctionCode returns the EdgeFunctionCode field if non-nil, zero value otherwise.

### GetEdgeFunctionCodeOk

`func (o *CreateCustomTool201Response) GetEdgeFunctionCodeOk() (*string, bool)`

GetEdgeFunctionCodeOk returns a tuple with the EdgeFunctionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionCode

`func (o *CreateCustomTool201Response) SetEdgeFunctionCode(v string)`

SetEdgeFunctionCode sets EdgeFunctionCode field to given value.

### HasEdgeFunctionCode

`func (o *CreateCustomTool201Response) HasEdgeFunctionCode() bool`

HasEdgeFunctionCode returns a boolean if a field has been set.

### GetIsUpdate

`func (o *CreateCustomTool201Response) GetIsUpdate() bool`

GetIsUpdate returns the IsUpdate field if non-nil, zero value otherwise.

### GetIsUpdateOk

`func (o *CreateCustomTool201Response) GetIsUpdateOk() (*bool, bool)`

GetIsUpdateOk returns a tuple with the IsUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpdate

`func (o *CreateCustomTool201Response) SetIsUpdate(v bool)`

SetIsUpdate sets IsUpdate field to given value.

### HasIsUpdate

`func (o *CreateCustomTool201Response) HasIsUpdate() bool`

HasIsUpdate returns a boolean if a field has been set.

### GetMessage

`func (o *CreateCustomTool201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateCustomTool201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateCustomTool201Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateCustomTool201Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


