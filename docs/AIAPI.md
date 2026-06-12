# \AIAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAIUsageStats**](AIAPI.md#GetAIUsageStats) | **Get** /api/v3/organizations/{organisation}/ai/usage | Organisation AI usage breakdown (subscription page parity)



## GetAIUsageStats

> GetAIUsageStats(ctx, organisation).Month(month).GroupBy(groupBy).Include(include).UserId(userId).TokenId(tokenId).Execute()

Organisation AI usage breakdown (subscription page parity)



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
	groupBy := "groupBy_example" // string |  (optional) (default to "model")
	include := "include_example" // string | Set to 'daily' to append a 30-day series (optional)
	userId := "userId_example" // string | Scope the daily series to a user (optional)
	tokenId := "tokenId_example" // string | Scope the daily series to a token (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIAPI.GetAIUsageStats(context.Background(), organisation).Month(month).GroupBy(groupBy).Include(include).UserId(userId).TokenId(tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAPI.GetAIUsageStats``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAIUsageStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **month** | **string** | YYYY-MM, defaults to current month | 
 **groupBy** | **string** |  | [default to &quot;model&quot;]
 **include** | **string** | Set to &#39;daily&#39; to append a 30-day series | 
 **userId** | **string** | Scope the daily series to a user | 
 **tokenId** | **string** | Scope the daily series to a token | 

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

