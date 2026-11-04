# \ProjectsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectEdgeEvents**](ProjectsAPI.md#GetProjectEdgeEvents) | **Get** /api/v2/organizations/{organization}/projects/{project}/edge-events | Get edge events for a project
[**GetProjectLogs**](ProjectsAPI.md#GetProjectLogs) | **Get** /api/v2/organizations/{organization}/projects/{project}/logs | Get CDN access logs for a project
[**ProjectsCreate**](ProjectsAPI.md#ProjectsCreate) | **Post** /api/v2/organizations/{organization}/projects | Create a new project
[**ProjectsDelete**](ProjectsAPI.md#ProjectsDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project} | Delete a project
[**ProjectsList**](ProjectsAPI.md#ProjectsList) | **Get** /api/v2/organizations/{organization}/projects | Retrieve all projects for an organization
[**ProjectsRead**](ProjectsAPI.md#ProjectsRead) | **Get** /api/v2/organizations/{organization}/projects/{project} | Get details of a single project
[**ProjectsUpdate**](ProjectsAPI.md#ProjectsUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project} | Update a project



## GetProjectEdgeEvents

> GetProjectEdgeEvents200Response GetProjectEdgeEvents(ctx, organization, project).Type_(type_).Function(function).Level(level).Search(search).Limit(limit).StartTime(startTime).EndTime(endTime).NextToken(nextToken).Execute()

Get edge events for a project



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
	organization := "organization_example" // string | The organization machine name
	project := "project_example" // string | The project machine name
	type_ := "console,function" // string | Comma-separated event types: console, function, waf-hit, static-failover (optional)
	function := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only events from this edge function (UUID) (optional)
	level := "level_example" // string | Only console events at this level (optional)
	search := "search_example" // string | Case-insensitive substring of the message (optional)
	limit := int32(56) // int32 | Maximum number of events per page (default 100) (optional) (default to 100)
	startTime := "startTime_example" // string | Start of the time range. ISO 8601 or Unix epoch milliseconds. (optional)
	endTime := "endTime_example" // string | End of the time range. ISO 8601 or Unix epoch milliseconds. (optional)
	nextToken := "nextToken_example" // string | Opaque pagination token from the previous response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectEdgeEvents(context.Background(), organization, project).Type_(type_).Function(function).Level(level).Search(search).Limit(limit).StartTime(startTime).EndTime(endTime).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectEdgeEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectEdgeEvents`: GetProjectEdgeEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectEdgeEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | The organization machine name | 
**project** | **string** | The project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEdgeEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | Comma-separated event types: console, function, waf-hit, static-failover | 
 **function** | **string** | Only events from this edge function (UUID) | 
 **level** | **string** | Only console events at this level | 
 **search** | **string** | Case-insensitive substring of the message | 
 **limit** | **int32** | Maximum number of events per page (default 100) | [default to 100]
 **startTime** | **string** | Start of the time range. ISO 8601 or Unix epoch milliseconds. | 
 **endTime** | **string** | End of the time range. ISO 8601 or Unix epoch milliseconds. | 
 **nextToken** | **string** | Opaque pagination token from the previous response | 

### Return type

[**GetProjectEdgeEvents200Response**](GetProjectEdgeEvents200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectLogs

> GetProjectLogs200Response GetProjectLogs(ctx, organization, project).Limit(limit).StartTime(startTime).EndTime(endTime).Filter(filter).Domain(domain).NextToken(nextToken).Execute()

Get CDN access logs for a project



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
	organization := "test-org" // string | The organization machine name
	project := "test-project" // string | The project machine name
	limit := int32(56) // int32 | Maximum number of log entries to return per page (default 100) (optional) (default to 100)
	startTime := "2024-11-16T00:00:00Z" // string | Start of the time range. ISO 8601 or Unix epoch milliseconds. (optional)
	endTime := "2024-11-16T23:59:59Z" // string | End of the time range. ISO 8601 or Unix epoch milliseconds. (optional)
	filter := "$.status_code = 404" // string | JSON filter expression AND-ed with the project constraint, e.g. $.status_code = 404. Outer braces are optional; nested braces are rejected. (optional)
	domain := "www.example.com" // string | Only return entries for this domain (optional)
	nextToken := "nextToken_example" // string | Opaque pagination token from the previous response. Pass back unchanged to fetch the next page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectLogs(context.Background(), organization, project).Limit(limit).StartTime(startTime).EndTime(endTime).Filter(filter).Domain(domain).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectLogs`: GetProjectLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | The organization machine name | 
**project** | **string** | The project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Maximum number of log entries to return per page (default 100) | [default to 100]
 **startTime** | **string** | Start of the time range. ISO 8601 or Unix epoch milliseconds. | 
 **endTime** | **string** | End of the time range. ISO 8601 or Unix epoch milliseconds. | 
 **filter** | **string** | JSON filter expression AND-ed with the project constraint, e.g. $.status_code &#x3D; 404. Outer braces are optional; nested braces are rejected. | 
 **domain** | **string** | Only return entries for this domain | 
 **nextToken** | **string** | Opaque pagination token from the previous response. Pass back unchanged to fetch the next page. | 

### Return type

[**GetProjectLogs200Response**](GetProjectLogs200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreate

> V2Project ProjectsCreate(ctx, organization).V2ProjectRequest(v2ProjectRequest).Execute()

Create a new project

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
	organization := "test-org" // string | Organization identifier
	v2ProjectRequest := *openapiclient.NewV2ProjectRequest() // V2ProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsCreate(context.Background(), organization).V2ProjectRequest(v2ProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCreate`: V2Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ProjectRequest** | [**V2ProjectRequest**](V2ProjectRequest.md) |  | 

### Return type

[**V2Project**](V2Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDelete

> ProjectsDelete(ctx, organization, project).Execute()

Delete a project

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsDelete(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> []V2Project ProjectsList(ctx, organization).Execute()

Retrieve all projects for an organization

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
	organization := "test-org" // string | Organization identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsList(context.Background(), organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsList`: []V2Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]V2Project**](V2Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsRead

> V2Project ProjectsRead(ctx, organization, project).WithToken(withToken).Execute()

Get details of a single project

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	withToken := true // bool |  (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsRead(context.Background(), organization, project).WithToken(withToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsRead`: V2Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withToken** | **bool** |  | [default to false]

### Return type

[**V2Project**](V2Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsUpdate

> V2Project ProjectsUpdate(ctx, organization, project).V2ProjectRequest(v2ProjectRequest).Execute()

Update a project

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	v2ProjectRequest := *openapiclient.NewV2ProjectRequest() // V2ProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsUpdate(context.Background(), organization, project).V2ProjectRequest(v2ProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsUpdate`: V2Project
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2ProjectRequest** | [**V2ProjectRequest**](V2ProjectRequest.md) |  | 

### Return type

[**V2Project**](V2Project.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

