# ListTasks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | Pointer to [**[]ListTasks200ResponseTasksInner**](ListTasks200ResponseTasksInner.md) |  | [optional] 
**TaskIds** | Pointer to **[]string** | Task IDs (only with dependsOn filter) | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**DependsOn** | Pointer to **string** | The queried task ID (only with dependsOn filter) | [optional] 

## Methods

### NewListTasks200Response

`func NewListTasks200Response() *ListTasks200Response`

NewListTasks200Response instantiates a new ListTasks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasks200ResponseWithDefaults

`func NewListTasks200ResponseWithDefaults() *ListTasks200Response`

NewListTasks200ResponseWithDefaults instantiates a new ListTasks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *ListTasks200Response) GetTasks() []ListTasks200ResponseTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ListTasks200Response) GetTasksOk() (*[]ListTasks200ResponseTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ListTasks200Response) SetTasks(v []ListTasks200ResponseTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ListTasks200Response) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTaskIds

`func (o *ListTasks200Response) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *ListTasks200Response) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *ListTasks200Response) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *ListTasks200Response) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetCount

`func (o *ListTasks200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListTasks200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListTasks200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListTasks200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDependsOn

`func (o *ListTasks200Response) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ListTasks200Response) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ListTasks200Response) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *ListTasks200Response) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


