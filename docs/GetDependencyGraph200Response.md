# GetDependencyGraph200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskListId** | Pointer to **string** |  | [optional] 
**TaskCount** | Pointer to **int32** |  | [optional] 
**Roots** | Pointer to **[]string** | Task IDs with no dependencies | [optional] 
**Leaves** | Pointer to **[]string** | Task IDs with no dependents | [optional] 
**Graph** | Pointer to **map[string]interface{}** | Adjacency list with task summaries, dependsOn, and dependents arrays | [optional] 

## Methods

### NewGetDependencyGraph200Response

`func NewGetDependencyGraph200Response() *GetDependencyGraph200Response`

NewGetDependencyGraph200Response instantiates a new GetDependencyGraph200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDependencyGraph200ResponseWithDefaults

`func NewGetDependencyGraph200ResponseWithDefaults() *GetDependencyGraph200Response`

NewGetDependencyGraph200ResponseWithDefaults instantiates a new GetDependencyGraph200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskListId

`func (o *GetDependencyGraph200Response) GetTaskListId() string`

GetTaskListId returns the TaskListId field if non-nil, zero value otherwise.

### GetTaskListIdOk

`func (o *GetDependencyGraph200Response) GetTaskListIdOk() (*string, bool)`

GetTaskListIdOk returns a tuple with the TaskListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskListId

`func (o *GetDependencyGraph200Response) SetTaskListId(v string)`

SetTaskListId sets TaskListId field to given value.

### HasTaskListId

`func (o *GetDependencyGraph200Response) HasTaskListId() bool`

HasTaskListId returns a boolean if a field has been set.

### GetTaskCount

`func (o *GetDependencyGraph200Response) GetTaskCount() int32`

GetTaskCount returns the TaskCount field if non-nil, zero value otherwise.

### GetTaskCountOk

`func (o *GetDependencyGraph200Response) GetTaskCountOk() (*int32, bool)`

GetTaskCountOk returns a tuple with the TaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCount

`func (o *GetDependencyGraph200Response) SetTaskCount(v int32)`

SetTaskCount sets TaskCount field to given value.

### HasTaskCount

`func (o *GetDependencyGraph200Response) HasTaskCount() bool`

HasTaskCount returns a boolean if a field has been set.

### GetRoots

`func (o *GetDependencyGraph200Response) GetRoots() []string`

GetRoots returns the Roots field if non-nil, zero value otherwise.

### GetRootsOk

`func (o *GetDependencyGraph200Response) GetRootsOk() (*[]string, bool)`

GetRootsOk returns a tuple with the Roots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoots

`func (o *GetDependencyGraph200Response) SetRoots(v []string)`

SetRoots sets Roots field to given value.

### HasRoots

`func (o *GetDependencyGraph200Response) HasRoots() bool`

HasRoots returns a boolean if a field has been set.

### GetLeaves

`func (o *GetDependencyGraph200Response) GetLeaves() []string`

GetLeaves returns the Leaves field if non-nil, zero value otherwise.

### GetLeavesOk

`func (o *GetDependencyGraph200Response) GetLeavesOk() (*[]string, bool)`

GetLeavesOk returns a tuple with the Leaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaves

`func (o *GetDependencyGraph200Response) SetLeaves(v []string)`

SetLeaves sets Leaves field to given value.

### HasLeaves

`func (o *GetDependencyGraph200Response) HasLeaves() bool`

HasLeaves returns a boolean if a field has been set.

### GetGraph

`func (o *GetDependencyGraph200Response) GetGraph() map[string]interface{}`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *GetDependencyGraph200Response) GetGraphOk() (*map[string]interface{}, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *GetDependencyGraph200Response) SetGraph(v map[string]interface{})`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *GetDependencyGraph200Response) HasGraph() bool`

HasGraph returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


