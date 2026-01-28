# \AIToolsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAIOrchestrationStatus**](AIToolsAPI.md#GetAIOrchestrationStatus) | **Get** /api/v3/organizations/{organisation}/ai/tools/orchestrations/{orchestrationId} | Get Orchestration Status
[**GetAIToolExecutionStatus**](AIToolsAPI.md#GetAIToolExecutionStatus) | **Get** /api/v3/organizations/{organisation}/ai/tools/executions/{executionId} | Get async tool execution status and result
[**ListAIToolExecutions**](AIToolsAPI.md#ListAIToolExecutions) | **Get** /api/v3/organizations/{organisation}/ai/tools/executions | List tool executions for monitoring and debugging
[**ListAIToolNames**](AIToolsAPI.md#ListAIToolNames) | **Get** /api/v3/organizations/{organisation}/ai/tools/names | List tool names only (lightweight response)
[**ListAITools**](AIToolsAPI.md#ListAITools) | **Get** /api/v3/organizations/{organisation}/ai/tools | List available built-in tools for function calling



## GetAIOrchestrationStatus

> GetAIOrchestrationStatus200Response GetAIOrchestrationStatus(ctx, organisation, orchestrationId).Execute()

Get Orchestration Status



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
	orchestrationId := "orch_abc123def456789012345678901234" // string | Orchestration identifier for aggregated async tool executions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.GetAIOrchestrationStatus(context.Background(), organisation, orchestrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.GetAIOrchestrationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIOrchestrationStatus`: GetAIOrchestrationStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.GetAIOrchestrationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**orchestrationId** | **string** | Orchestration identifier for aggregated async tool executions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIOrchestrationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAIOrchestrationStatus200Response**](GetAIOrchestrationStatus200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIToolExecutionStatus

> GetAIToolExecutionStatus200Response GetAIToolExecutionStatus(ctx, organisation, executionId).Execute()

Get async tool execution status and result



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
	executionId := "exec_0123456789abcdef0123456789abcdef" // string | Tool execution identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.GetAIToolExecutionStatus(context.Background(), organisation, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.GetAIToolExecutionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIToolExecutionStatus`: GetAIToolExecutionStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.GetAIToolExecutionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**executionId** | **string** | Tool execution identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIToolExecutionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAIToolExecutionStatus200Response**](GetAIToolExecutionStatus200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIToolExecutions

> ListAIToolExecutions200Response ListAIToolExecutions(ctx, organisation).Status(status).Limit(limit).Execute()

List tool executions for monitoring and debugging



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
	status := "status_example" // string | Filter by execution status (optional)
	limit := int32(56) // int32 | Maximum number of executions to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.ListAIToolExecutions(context.Background(), organisation).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.ListAIToolExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIToolExecutions`: ListAIToolExecutions200Response
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.ListAIToolExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIToolExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by execution status | 
 **limit** | **int32** | Maximum number of executions to return | [default to 50]

### Return type

[**ListAIToolExecutions200Response**](ListAIToolExecutions200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIToolNames

> ListAIToolNames200Response ListAIToolNames(ctx, organisation).Execute()

List tool names only (lightweight response)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.ListAIToolNames(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.ListAIToolNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIToolNames`: ListAIToolNames200Response
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.ListAIToolNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIToolNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAIToolNames200Response**](ListAIToolNames200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAITools

> ListAITools200Response ListAITools(ctx, organisation).Execute()

List available built-in tools for function calling



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.ListAITools(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.ListAITools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAITools`: ListAITools200Response
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.ListAITools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAITools200Response**](ListAITools200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

