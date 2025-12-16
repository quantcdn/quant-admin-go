# \AIMonitoringAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAIUsageStats**](AIMonitoringAPI.md#GetAIUsageStats) | **Get** /api/v3/organizations/{organisation}/ai/usage | Get AI usage statistics



## GetAIUsageStats

> GetAIUsageStats200Response GetAIUsageStats(ctx, organisation).Month(month).Execute()

Get AI usage statistics

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
	month := "2025-10" // string | Month to retrieve statistics for (YYYY-MM format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMonitoringAPI.GetAIUsageStats(context.Background(), organisation).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMonitoringAPI.GetAIUsageStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIUsageStats`: GetAIUsageStats200Response
	fmt.Fprintf(os.Stdout, "Response from `AIMonitoringAPI.GetAIUsageStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIUsageStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **month** | **string** | Month to retrieve statistics for (YYYY-MM format) | 

### Return type

[**GetAIUsageStats200Response**](GetAIUsageStats200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

