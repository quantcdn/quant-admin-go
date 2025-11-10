# \EnvironmentsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](EnvironmentsAPI.md#CreateEnvironment) | **Post** /api/v3/organizations/{organisation}/applications/{application}/environments | Create a new environment
[**DeleteEnvironment**](EnvironmentsAPI.md#DeleteEnvironment) | **Delete** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment} | Delete an environment
[**GetEnvironment**](EnvironmentsAPI.md#GetEnvironment) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment} | Get a single environment
[**GetEnvironmentLogs**](EnvironmentsAPI.md#GetEnvironmentLogs) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/logs | Get the logs for an environment
[**GetEnvironmentMetrics**](EnvironmentsAPI.md#GetEnvironmentMetrics) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/metrics | Get the metrics for an environment
[**ListEnvironments**](EnvironmentsAPI.md#ListEnvironments) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments | Get all environments for an application
[**ListSyncOperations**](EnvironmentsAPI.md#ListSyncOperations) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/sync/{type} | List the sync operations for an environment
[**SyncToEnvironment**](EnvironmentsAPI.md#SyncToEnvironment) | **Post** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/sync/{type} | Perform a sync operation from a source environment to the current environment
[**UpdateEnvironment**](EnvironmentsAPI.md#UpdateEnvironment) | **Put** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment} | Update Environment Compose Definition
[**UpdateEnvironmentState**](EnvironmentsAPI.md#UpdateEnvironmentState) | **Put** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/state | Update the state of an environment



## CreateEnvironment

> Environment CreateEnvironment(ctx, organisation, application).CreateEnvironmentRequest(createEnvironmentRequest).Execute()

Create a new environment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	createEnvironmentRequest := *openapiclient.NewCreateEnvironmentRequest("EnvName_example") // CreateEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.CreateEnvironment(context.Background(), organisation, application).CreateEnvironmentRequest(createEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.CreateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironment`: Environment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createEnvironmentRequest** | [**CreateEnvironmentRequest**](CreateEnvironmentRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, organisation, application, environment).Execute()

Delete an environment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentsAPI.DeleteEnvironment(context.Background(), organisation, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.DeleteEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironment

> Environment GetEnvironment(ctx, organisation, application, environment).Execute()

Get a single environment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.GetEnvironment(context.Background(), organisation, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.GetEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironment`: Environment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Environment**](Environment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentLogs

> GetEnvironmentLogs200Response GetEnvironmentLogs(ctx, organisation, application, environment).StartTime(startTime).EndTime(endTime).ContainerName(containerName).FilterPattern(filterPattern).Limit(limit).NextToken(nextToken).Execute()

Get the logs for an environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	startTime := "startTime_example" // string | Start time for log retrieval (ISO 8601 format or Unix timestamp) (optional)
	endTime := "endTime_example" // string | End time for log retrieval (ISO 8601 format or Unix timestamp) (optional)
	containerName := "containerName_example" // string | Filter logs by specific container name (optional)
	filterPattern := "filterPattern_example" // string | CloudWatch Logs filter pattern for searching log content (optional)
	limit := int32(56) // int32 | Maximum number of log entries to return per page (optional)
	nextToken := "nextToken_example" // string | Pagination token from previous response for retrieving next page of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.GetEnvironmentLogs(context.Background(), organisation, application, environment).StartTime(startTime).EndTime(endTime).ContainerName(containerName).FilterPattern(filterPattern).Limit(limit).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.GetEnvironmentLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentLogs`: GetEnvironmentLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.GetEnvironmentLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **string** | Start time for log retrieval (ISO 8601 format or Unix timestamp) | 
 **endTime** | **string** | End time for log retrieval (ISO 8601 format or Unix timestamp) | 
 **containerName** | **string** | Filter logs by specific container name | 
 **filterPattern** | **string** | CloudWatch Logs filter pattern for searching log content | 
 **limit** | **int32** | Maximum number of log entries to return per page | 
 **nextToken** | **string** | Pagination token from previous response for retrieving next page of results | 

### Return type

[**GetEnvironmentLogs200Response**](GetEnvironmentLogs200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentMetrics

> map[string]interface{} GetEnvironmentMetrics(ctx, organisation, application, environment).StartTime(startTime).EndTime(endTime).Period(period).Statistics(statistics).ContainerName(containerName).Execute()

Get the metrics for an environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	startTime := int32(56) // int32 | Start time for metrics retrieval (Unix timestamp in milliseconds) (optional)
	endTime := int32(56) // int32 | End time for metrics retrieval (Unix timestamp in milliseconds) (optional)
	period := int32(56) // int32 | Period in seconds for metric aggregation (e.g., 60 for 1 minute, 300 for 5 minutes) (optional)
	statistics := "statistics_example" // string | Comma-separated list of CloudWatch statistics (e.g., Average, Maximum, Minimum, Sum, SampleCount) (optional)
	containerName := "containerName_example" // string | Filter metrics by specific container name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.GetEnvironmentMetrics(context.Background(), organisation, application, environment).StartTime(startTime).EndTime(endTime).Period(period).Statistics(statistics).ContainerName(containerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.GetEnvironmentMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.GetEnvironmentMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **int32** | Start time for metrics retrieval (Unix timestamp in milliseconds) | 
 **endTime** | **int32** | End time for metrics retrieval (Unix timestamp in milliseconds) | 
 **period** | **int32** | Period in seconds for metric aggregation (e.g., 60 for 1 minute, 300 for 5 minutes) | 
 **statistics** | **string** | Comma-separated list of CloudWatch statistics (e.g., Average, Maximum, Minimum, Sum, SampleCount) | 
 **containerName** | **string** | Filter metrics by specific container name | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironments

> []Environment ListEnvironments(ctx, organisation, application).Execute()

Get all environments for an application

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.ListEnvironments(context.Background(), organisation, application).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.ListEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironments`: []Environment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.ListEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Environment**](Environment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncOperations

> []SyncOperation ListSyncOperations(ctx, organisation, application, environment, type_).Execute()

List the sync operations for an environment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	type_ := "type__example" // string | The sync type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.ListSyncOperations(context.Background(), organisation, application, environment, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.ListSyncOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSyncOperations`: []SyncOperation
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.ListSyncOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**type_** | **string** | The sync type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSyncOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]SyncOperation**](SyncOperation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncToEnvironment

> SyncOperation SyncToEnvironment(ctx, organisation, application, environment, type_).SyncToEnvironmentRequest(syncToEnvironmentRequest).Execute()

Perform a sync operation from a source environment to the current environment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	type_ := "type__example" // string | The sync type
	syncToEnvironmentRequest := *openapiclient.NewSyncToEnvironmentRequest() // SyncToEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.SyncToEnvironment(context.Background(), organisation, application, environment, type_).SyncToEnvironmentRequest(syncToEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.SyncToEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncToEnvironment`: SyncOperation
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.SyncToEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**type_** | **string** | The sync type | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncToEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **syncToEnvironmentRequest** | [**SyncToEnvironmentRequest**](SyncToEnvironmentRequest.md) |  | 

### Return type

[**SyncOperation**](SyncOperation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironment

> UpdateEnvironment(ctx, organisation, application, environment).UpdateEnvironmentRequest(updateEnvironmentRequest).Execute()

Update Environment Compose Definition



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	updateEnvironmentRequest := *openapiclient.NewUpdateEnvironmentRequest(*openapiclient.NewCompose()) // UpdateEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentsAPI.UpdateEnvironment(context.Background(), organisation, application, environment).UpdateEnvironmentRequest(updateEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.UpdateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateEnvironmentRequest** | [**UpdateEnvironmentRequest**](UpdateEnvironmentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironmentState

> UpdateEnvironmentState(ctx, organisation, application, environment).UpdateEnvironmentStateRequest(updateEnvironmentStateRequest).Execute()

Update the state of an environment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	updateEnvironmentStateRequest := *openapiclient.NewUpdateEnvironmentStateRequest() // UpdateEnvironmentStateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentsAPI.UpdateEnvironmentState(context.Background(), organisation, application, environment).UpdateEnvironmentStateRequest(updateEnvironmentStateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.UpdateEnvironmentState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateEnvironmentStateRequest** | [**UpdateEnvironmentStateRequest**](UpdateEnvironmentStateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

