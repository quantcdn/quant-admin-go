# \SubscriptionAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriptionCloudUsage**](SubscriptionAPI.md#GetSubscriptionCloudUsage) | **Get** /api/v3/organizations/{organisation}/subscription/cloud-usage | Cloud-app cost breakdown for the subscription page



## GetSubscriptionCloudUsage

> GetSubscriptionCloudUsage(ctx, organisation).Month(month).Execute()

Cloud-app cost breakdown for the subscription page



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
	organisation := "organisation_example" // string | 
	month := "month_example" // string | YYYY-MM, defaults to current month (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionAPI.GetSubscriptionCloudUsage(context.Background(), organisation).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.GetSubscriptionCloudUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionCloudUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **month** | **string** | YYYY-MM, defaults to current month | 

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

