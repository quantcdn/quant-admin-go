# CreateCustomToolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique tool name (alphanumeric and underscores only) | 
**Description** | **string** | Human-readable description of what the tool does | 
**EdgeFunctionUrl** | **string** | HTTPS URL of the edge function | 
**InputSchema** | **map[string]interface{}** | JSON Schema defining the tool&#39;s input parameters | 
**IsAsync** | Pointer to **bool** | Whether this tool runs asynchronously (&gt;5 seconds) | [optional] [default to false]
**TimeoutSeconds** | Pointer to **int32** | Tool execution timeout | [optional] [default to 30]

## Methods

### NewCreateCustomToolRequest

`func NewCreateCustomToolRequest(name string, description string, edgeFunctionUrl string, inputSchema map[string]interface{}, ) *CreateCustomToolRequest`

NewCreateCustomToolRequest instantiates a new CreateCustomToolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomToolRequestWithDefaults

`func NewCreateCustomToolRequestWithDefaults() *CreateCustomToolRequest`

NewCreateCustomToolRequestWithDefaults instantiates a new CreateCustomToolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCustomToolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomToolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomToolRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCustomToolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomToolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomToolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEdgeFunctionUrl

`func (o *CreateCustomToolRequest) GetEdgeFunctionUrl() string`

GetEdgeFunctionUrl returns the EdgeFunctionUrl field if non-nil, zero value otherwise.

### GetEdgeFunctionUrlOk

`func (o *CreateCustomToolRequest) GetEdgeFunctionUrlOk() (*string, bool)`

GetEdgeFunctionUrlOk returns a tuple with the EdgeFunctionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionUrl

`func (o *CreateCustomToolRequest) SetEdgeFunctionUrl(v string)`

SetEdgeFunctionUrl sets EdgeFunctionUrl field to given value.


### GetInputSchema

`func (o *CreateCustomToolRequest) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *CreateCustomToolRequest) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *CreateCustomToolRequest) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetIsAsync

`func (o *CreateCustomToolRequest) GetIsAsync() bool`

GetIsAsync returns the IsAsync field if non-nil, zero value otherwise.

### GetIsAsyncOk

`func (o *CreateCustomToolRequest) GetIsAsyncOk() (*bool, bool)`

GetIsAsyncOk returns a tuple with the IsAsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsync

`func (o *CreateCustomToolRequest) SetIsAsync(v bool)`

SetIsAsync sets IsAsync field to given value.

### HasIsAsync

`func (o *CreateCustomToolRequest) HasIsAsync() bool`

HasIsAsync returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *CreateCustomToolRequest) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *CreateCustomToolRequest) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *CreateCustomToolRequest) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *CreateCustomToolRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


