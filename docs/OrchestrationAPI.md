# \OrchestrationAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAIOrchestrationStatus**](OrchestrationAPI.md#GetAIOrchestrationStatus) | **Get** /api/v3/organizations/{organisation}/ai/tools/orchestrations/{orchestrationId} | Get Orchestration Status



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
	resp, r, err := apiClient.OrchestrationAPI.GetAIOrchestrationStatus(context.Background(), organisation, orchestrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrchestrationAPI.GetAIOrchestrationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIOrchestrationStatus`: GetAIOrchestrationStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `OrchestrationAPI.GetAIOrchestrationStatus`: %v\n", resp)
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

