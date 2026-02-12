# CreateOrchestrationRequestInputSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Input source type (api type not yet supported) | 
**Items** | Pointer to **[]map[string]interface{}** | Static items (for type&#x3D;static) | [optional] 
**TaskQuery** | Pointer to **map[string]interface{}** | Task query filters (for type&#x3D;task_query) | [optional] 
**GeneratorPrompt** | Pointer to **string** | AI prompt (for type&#x3D;generator) | [optional] 

## Methods

### NewCreateOrchestrationRequestInputSource

`func NewCreateOrchestrationRequestInputSource(type_ string, ) *CreateOrchestrationRequestInputSource`

NewCreateOrchestrationRequestInputSource instantiates a new CreateOrchestrationRequestInputSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrchestrationRequestInputSourceWithDefaults

`func NewCreateOrchestrationRequestInputSourceWithDefaults() *CreateOrchestrationRequestInputSource`

NewCreateOrchestrationRequestInputSourceWithDefaults instantiates a new CreateOrchestrationRequestInputSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOrchestrationRequestInputSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrchestrationRequestInputSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrchestrationRequestInputSource) SetType(v string)`

SetType sets Type field to given value.


### GetItems

`func (o *CreateOrchestrationRequestInputSource) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateOrchestrationRequestInputSource) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateOrchestrationRequestInputSource) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *CreateOrchestrationRequestInputSource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTaskQuery

`func (o *CreateOrchestrationRequestInputSource) GetTaskQuery() map[string]interface{}`

GetTaskQuery returns the TaskQuery field if non-nil, zero value otherwise.

### GetTaskQueryOk

`func (o *CreateOrchestrationRequestInputSource) GetTaskQueryOk() (*map[string]interface{}, bool)`

GetTaskQueryOk returns a tuple with the TaskQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQuery

`func (o *CreateOrchestrationRequestInputSource) SetTaskQuery(v map[string]interface{})`

SetTaskQuery sets TaskQuery field to given value.

### HasTaskQuery

`func (o *CreateOrchestrationRequestInputSource) HasTaskQuery() bool`

HasTaskQuery returns a boolean if a field has been set.

### GetGeneratorPrompt

`func (o *CreateOrchestrationRequestInputSource) GetGeneratorPrompt() string`

GetGeneratorPrompt returns the GeneratorPrompt field if non-nil, zero value otherwise.

### GetGeneratorPromptOk

`func (o *CreateOrchestrationRequestInputSource) GetGeneratorPromptOk() (*string, bool)`

GetGeneratorPromptOk returns a tuple with the GeneratorPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorPrompt

`func (o *CreateOrchestrationRequestInputSource) SetGeneratorPrompt(v string)`

SetGeneratorPrompt sets GeneratorPrompt field to given value.

### HasGeneratorPrompt

`func (o *CreateOrchestrationRequestInputSource) HasGeneratorPrompt() bool`

HasGeneratorPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


