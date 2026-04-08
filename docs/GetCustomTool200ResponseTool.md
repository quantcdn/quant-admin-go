# GetCustomTool200ResponseTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EdgeFunctionUrl** | Pointer to **string** |  | [optional] 
**EdgeFunctionCode** | Pointer to **string** | The deployed edge function source code | [optional] 
**IsAsync** | Pointer to **bool** |  | [optional] 
**InputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**OutputSchemaDescription** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**ResponseMode** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetCustomTool200ResponseTool

`func NewGetCustomTool200ResponseTool() *GetCustomTool200ResponseTool`

NewGetCustomTool200ResponseTool instantiates a new GetCustomTool200ResponseTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomTool200ResponseToolWithDefaults

`func NewGetCustomTool200ResponseToolWithDefaults() *GetCustomTool200ResponseTool`

NewGetCustomTool200ResponseToolWithDefaults instantiates a new GetCustomTool200ResponseTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetCustomTool200ResponseTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCustomTool200ResponseTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCustomTool200ResponseTool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCustomTool200ResponseTool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetCustomTool200ResponseTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCustomTool200ResponseTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCustomTool200ResponseTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCustomTool200ResponseTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEdgeFunctionUrl

`func (o *GetCustomTool200ResponseTool) GetEdgeFunctionUrl() string`

GetEdgeFunctionUrl returns the EdgeFunctionUrl field if non-nil, zero value otherwise.

### GetEdgeFunctionUrlOk

`func (o *GetCustomTool200ResponseTool) GetEdgeFunctionUrlOk() (*string, bool)`

GetEdgeFunctionUrlOk returns a tuple with the EdgeFunctionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionUrl

`func (o *GetCustomTool200ResponseTool) SetEdgeFunctionUrl(v string)`

SetEdgeFunctionUrl sets EdgeFunctionUrl field to given value.

### HasEdgeFunctionUrl

`func (o *GetCustomTool200ResponseTool) HasEdgeFunctionUrl() bool`

HasEdgeFunctionUrl returns a boolean if a field has been set.

### GetEdgeFunctionCode

`func (o *GetCustomTool200ResponseTool) GetEdgeFunctionCode() string`

GetEdgeFunctionCode returns the EdgeFunctionCode field if non-nil, zero value otherwise.

### GetEdgeFunctionCodeOk

`func (o *GetCustomTool200ResponseTool) GetEdgeFunctionCodeOk() (*string, bool)`

GetEdgeFunctionCodeOk returns a tuple with the EdgeFunctionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionCode

`func (o *GetCustomTool200ResponseTool) SetEdgeFunctionCode(v string)`

SetEdgeFunctionCode sets EdgeFunctionCode field to given value.

### HasEdgeFunctionCode

`func (o *GetCustomTool200ResponseTool) HasEdgeFunctionCode() bool`

HasEdgeFunctionCode returns a boolean if a field has been set.

### GetIsAsync

`func (o *GetCustomTool200ResponseTool) GetIsAsync() bool`

GetIsAsync returns the IsAsync field if non-nil, zero value otherwise.

### GetIsAsyncOk

`func (o *GetCustomTool200ResponseTool) GetIsAsyncOk() (*bool, bool)`

GetIsAsyncOk returns a tuple with the IsAsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsync

`func (o *GetCustomTool200ResponseTool) SetIsAsync(v bool)`

SetIsAsync sets IsAsync field to given value.

### HasIsAsync

`func (o *GetCustomTool200ResponseTool) HasIsAsync() bool`

HasIsAsync returns a boolean if a field has been set.

### GetInputSchema

`func (o *GetCustomTool200ResponseTool) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *GetCustomTool200ResponseTool) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *GetCustomTool200ResponseTool) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *GetCustomTool200ResponseTool) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *GetCustomTool200ResponseTool) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *GetCustomTool200ResponseTool) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *GetCustomTool200ResponseTool) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *GetCustomTool200ResponseTool) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *GetCustomTool200ResponseTool) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *GetCustomTool200ResponseTool) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil
### GetOutputSchemaDescription

`func (o *GetCustomTool200ResponseTool) GetOutputSchemaDescription() string`

GetOutputSchemaDescription returns the OutputSchemaDescription field if non-nil, zero value otherwise.

### GetOutputSchemaDescriptionOk

`func (o *GetCustomTool200ResponseTool) GetOutputSchemaDescriptionOk() (*string, bool)`

GetOutputSchemaDescriptionOk returns a tuple with the OutputSchemaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchemaDescription

`func (o *GetCustomTool200ResponseTool) SetOutputSchemaDescription(v string)`

SetOutputSchemaDescription sets OutputSchemaDescription field to given value.

### HasOutputSchemaDescription

`func (o *GetCustomTool200ResponseTool) HasOutputSchemaDescription() bool`

HasOutputSchemaDescription returns a boolean if a field has been set.

### SetOutputSchemaDescriptionNil

`func (o *GetCustomTool200ResponseTool) SetOutputSchemaDescriptionNil(b bool)`

 SetOutputSchemaDescriptionNil sets the value for OutputSchemaDescription to be an explicit nil

### UnsetOutputSchemaDescription
`func (o *GetCustomTool200ResponseTool) UnsetOutputSchemaDescription()`

UnsetOutputSchemaDescription ensures that no value is present for OutputSchemaDescription, not even an explicit nil
### GetCategory

`func (o *GetCustomTool200ResponseTool) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetCustomTool200ResponseTool) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetCustomTool200ResponseTool) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetCustomTool200ResponseTool) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetCustomTool200ResponseTool) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetCustomTool200ResponseTool) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetResponseMode

`func (o *GetCustomTool200ResponseTool) GetResponseMode() string`

GetResponseMode returns the ResponseMode field if non-nil, zero value otherwise.

### GetResponseModeOk

`func (o *GetCustomTool200ResponseTool) GetResponseModeOk() (*string, bool)`

GetResponseModeOk returns a tuple with the ResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMode

`func (o *GetCustomTool200ResponseTool) SetResponseMode(v string)`

SetResponseMode sets ResponseMode field to given value.

### HasResponseMode

`func (o *GetCustomTool200ResponseTool) HasResponseMode() bool`

HasResponseMode returns a boolean if a field has been set.

### SetResponseModeNil

`func (o *GetCustomTool200ResponseTool) SetResponseModeNil(b bool)`

 SetResponseModeNil sets the value for ResponseMode to be an explicit nil

### UnsetResponseMode
`func (o *GetCustomTool200ResponseTool) UnsetResponseMode()`

UnsetResponseMode ensures that no value is present for ResponseMode, not even an explicit nil
### GetCreatedAt

`func (o *GetCustomTool200ResponseTool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetCustomTool200ResponseTool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetCustomTool200ResponseTool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetCustomTool200ResponseTool) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


