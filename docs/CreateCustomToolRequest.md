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
**OutputSchema** | Pointer to **map[string]interface{}** | JSON Schema defining the tool&#39;s output structure | [optional] 
**OutputSchemaDescription** | Pointer to **NullableString** | Human-readable description of the tool&#39;s output | [optional] 
**Category** | Pointer to **NullableString** | Category to group related tools | [optional] 
**ResponseMode** | Pointer to **NullableString** | How the tool response is handled: llm (passed back to model) or direct (returned to user) | [optional] 

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

### GetOutputSchema

`func (o *CreateCustomToolRequest) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *CreateCustomToolRequest) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *CreateCustomToolRequest) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *CreateCustomToolRequest) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *CreateCustomToolRequest) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *CreateCustomToolRequest) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil
### GetOutputSchemaDescription

`func (o *CreateCustomToolRequest) GetOutputSchemaDescription() string`

GetOutputSchemaDescription returns the OutputSchemaDescription field if non-nil, zero value otherwise.

### GetOutputSchemaDescriptionOk

`func (o *CreateCustomToolRequest) GetOutputSchemaDescriptionOk() (*string, bool)`

GetOutputSchemaDescriptionOk returns a tuple with the OutputSchemaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchemaDescription

`func (o *CreateCustomToolRequest) SetOutputSchemaDescription(v string)`

SetOutputSchemaDescription sets OutputSchemaDescription field to given value.

### HasOutputSchemaDescription

`func (o *CreateCustomToolRequest) HasOutputSchemaDescription() bool`

HasOutputSchemaDescription returns a boolean if a field has been set.

### SetOutputSchemaDescriptionNil

`func (o *CreateCustomToolRequest) SetOutputSchemaDescriptionNil(b bool)`

 SetOutputSchemaDescriptionNil sets the value for OutputSchemaDescription to be an explicit nil

### UnsetOutputSchemaDescription
`func (o *CreateCustomToolRequest) UnsetOutputSchemaDescription()`

UnsetOutputSchemaDescription ensures that no value is present for OutputSchemaDescription, not even an explicit nil
### GetCategory

`func (o *CreateCustomToolRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateCustomToolRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateCustomToolRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateCustomToolRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CreateCustomToolRequest) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CreateCustomToolRequest) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetResponseMode

`func (o *CreateCustomToolRequest) GetResponseMode() string`

GetResponseMode returns the ResponseMode field if non-nil, zero value otherwise.

### GetResponseModeOk

`func (o *CreateCustomToolRequest) GetResponseModeOk() (*string, bool)`

GetResponseModeOk returns a tuple with the ResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMode

`func (o *CreateCustomToolRequest) SetResponseMode(v string)`

SetResponseMode sets ResponseMode field to given value.

### HasResponseMode

`func (o *CreateCustomToolRequest) HasResponseMode() bool`

HasResponseMode returns a boolean if a field has been set.

### SetResponseModeNil

`func (o *CreateCustomToolRequest) SetResponseModeNil(b bool)`

 SetResponseModeNil sets the value for ResponseMode to be an explicit nil

### UnsetResponseMode
`func (o *CreateCustomToolRequest) UnsetResponseMode()`

UnsetResponseMode ensures that no value is present for ResponseMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


