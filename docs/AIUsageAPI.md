# \AIUsageAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyUsage**](AIUsageAPI.md#GetMyUsage) | **Get** /api/v3/organizations/{organisation}/ai/usage/me | Get AI usage summary for the authenticated user



## GetMyUsage

> GetMyUsage200Response GetMyUsage(ctx, organisation).Execute()

Get AI usage summary for the authenticated user

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
	resp, r, err := apiClient.AIUsageAPI.GetMyUsage(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIUsageAPI.GetMyUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyUsage`: GetMyUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `AIUsageAPI.GetMyUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMyUsage200Response**](GetMyUsage200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

