# \AITaskManagementAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTask**](AITaskManagementAPI.md#CreateTask) | **Post** /api/v3/organizations/{organisation}/ai/tasks | Create a new task
[**DeleteTask**](AITaskManagementAPI.md#DeleteTask) | **Delete** /api/v3/organizations/{organisation}/ai/tasks/{taskId} | Delete a task
[**GetDependencyGraph**](AITaskManagementAPI.md#GetDependencyGraph) | **Get** /api/v3/organizations/{organisation}/ai/tasks/{taskListId}/dependency-graph | Get dependency graph for a task list
[**GetTask**](AITaskManagementAPI.md#GetTask) | **Get** /api/v3/organizations/{organisation}/ai/tasks/{taskId} | Get task details
[**ListTasks**](AITaskManagementAPI.md#ListTasks) | **Get** /api/v3/organizations/{organisation}/ai/tasks | List tasks with optional filtering
[**UpdateTask**](AITaskManagementAPI.md#UpdateTask) | **Put** /api/v3/organizations/{organisation}/ai/tasks/{taskId} | Update a task



## CreateTask

> CreateTask201Response CreateTask(ctx, organisation).CreateTaskRequest(createTaskRequest).Execute()

Create a new task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	createTaskRequest := *openapiclient.NewCreateTaskRequest("Process document and create summary") // CreateTaskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AITaskManagementAPI.CreateTask(context.Background(), organisation).CreateTaskRequest(createTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AITaskManagementAPI.CreateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTask`: CreateTask201Response
	fmt.Fprintf(os.Stdout, "Response from `AITaskManagementAPI.CreateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTaskRequest** | [**CreateTaskRequest**](CreateTaskRequest.md) |  | 

### Return type

[**CreateTask201Response**](CreateTask201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTask

> DeleteTask200Response DeleteTask(ctx, organisation, taskId).Cascade(cascade).Execute()

Delete a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The task UUID
	cascade := true // bool | If true, delete task and all dependent tasks recursively (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AITaskManagementAPI.DeleteTask(context.Background(), organisation, taskId).Cascade(cascade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AITaskManagementAPI.DeleteTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTask`: DeleteTask200Response
	fmt.Fprintf(os.Stdout, "Response from `AITaskManagementAPI.DeleteTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**taskId** | **string** | The task UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cascade** | **bool** | If true, delete task and all dependent tasks recursively | [default to false]

### Return type

[**DeleteTask200Response**](DeleteTask200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDependencyGraph

> GetDependencyGraph200Response GetDependencyGraph(ctx, organisation, taskListId).Execute()

Get dependency graph for a task list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	taskListId := "world-1" // string | The task list ID to get the dependency graph for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AITaskManagementAPI.GetDependencyGraph(context.Background(), organisation, taskListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AITaskManagementAPI.GetDependencyGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDependencyGraph`: GetDependencyGraph200Response
	fmt.Fprintf(os.Stdout, "Response from `AITaskManagementAPI.GetDependencyGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**taskListId** | **string** | The task list ID to get the dependency graph for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDependencyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDependencyGraph200Response**](GetDependencyGraph200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> GetTask200Response GetTask(ctx, organisation, taskId).Execute()

Get task details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The task UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AITaskManagementAPI.GetTask(context.Background(), organisation, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AITaskManagementAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: GetTask200Response
	fmt.Fprintf(os.Stdout, "Response from `AITaskManagementAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**taskId** | **string** | The task UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTask200Response**](GetTask200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> ListTasks200Response ListTasks(ctx, organisation).TaskListId(taskListId).Status(status).AssignedAgentId(assignedAgentId).Limit(limit).DependsOn(dependsOn).IncludeDetails(includeDetails).Execute()

List tasks with optional filtering



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	taskListId := "world-1" // string | Filter tasks by task list ID. Task lists are implicit groupings - any string can be used. (optional)
	status := "pending" // string | Filter tasks by status (optional)
	assignedAgentId := "agent-code-reviewer" // string | Filter tasks by assigned agent ID (optional)
	limit := int32(56) // int32 | Maximum number of tasks to return (default 50, max 100) (optional) (default to 50)
	dependsOn := "550e8400-e29b-41d4-a716-446655440000" // string | Reverse lookup: find tasks that depend on this task ID. Returns tasks waiting for the specified task to complete. (optional)
	includeDetails := true // bool | When using dependsOn, return full task objects in addition to IDs. Default false (IDs only for lightweight responses). (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AITaskManagementAPI.ListTasks(context.Background(), organisation).TaskListId(taskListId).Status(status).AssignedAgentId(assignedAgentId).Limit(limit).DependsOn(dependsOn).IncludeDetails(includeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AITaskManagementAPI.ListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasks`: ListTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `AITaskManagementAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskListId** | **string** | Filter tasks by task list ID. Task lists are implicit groupings - any string can be used. | 
 **status** | **string** | Filter tasks by status | 
 **assignedAgentId** | **string** | Filter tasks by assigned agent ID | 
 **limit** | **int32** | Maximum number of tasks to return (default 50, max 100) | [default to 50]
 **dependsOn** | **string** | Reverse lookup: find tasks that depend on this task ID. Returns tasks waiting for the specified task to complete. | 
 **includeDetails** | **bool** | When using dependsOn, return full task objects in addition to IDs. Default false (IDs only for lightweight responses). | [default to false]

### Return type

[**ListTasks200Response**](ListTasks200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> UpdateTask200Response UpdateTask(ctx, organisation, taskId).UpdateTaskRequest(updateTaskRequest).Execute()

Update a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	taskId := "550e8400-e29b-41d4-a716-446655440000" // string | The task UUID
	updateTaskRequest := *openapiclient.NewUpdateTaskRequest() // UpdateTaskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AITaskManagementAPI.UpdateTask(context.Background(), organisation, taskId).UpdateTaskRequest(updateTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AITaskManagementAPI.UpdateTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTask`: UpdateTask200Response
	fmt.Fprintf(os.Stdout, "Response from `AITaskManagementAPI.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**taskId** | **string** | The task UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateTaskRequest** | [**UpdateTaskRequest**](UpdateTaskRequest.md) |  | 

### Return type

[**UpdateTask200Response**](UpdateTask200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

