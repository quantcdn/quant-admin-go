# \CDNMetricsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyMetrics**](CDNMetricsAPI.md#GetDailyMetrics) | **Get** /v2/organizations/{organization}/projects/{project}/metrics/daily | Get daily metrics
[**GetHourlyMetrics**](CDNMetricsAPI.md#GetHourlyMetrics) | **Get** /v2/organizations/{organization}/projects/{project}/metrics/hourly | Get hourly metrics
[**GetMonthlyMetrics**](CDNMetricsAPI.md#GetMonthlyMetrics) | **Get** /v2/organizations/{organization}/projects/{project}/metrics/monthly | Get monthly metrics



## GetDailyMetrics

> V2MetricsResponse GetDailyMetrics(ctx, organization, project).Domain(domain).Metrics(metrics).TimestampFormat(timestampFormat).Execute()

Get daily metrics



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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	domain := "domain_example" // string | Filter by domain ID or domain name (optional)
	metrics := []string{"Inner_example"} // []string | Metrics to return (default: hits, bytes) (optional)
	timestampFormat := "timestampFormat_example" // string | Timestamp format in response (optional) (default to "iso8601")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDNMetricsAPI.GetDailyMetrics(context.Background(), organization, project).Domain(domain).Metrics(metrics).TimestampFormat(timestampFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDNMetricsAPI.GetDailyMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDailyMetrics`: V2MetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `CDNMetricsAPI.GetDailyMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domain** | **string** | Filter by domain ID or domain name | 
 **metrics** | **[]string** | Metrics to return (default: hits, bytes) | 
 **timestampFormat** | **string** | Timestamp format in response | [default to &quot;iso8601&quot;]

### Return type

[**V2MetricsResponse**](V2MetricsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHourlyMetrics

> V2MetricsResponse GetHourlyMetrics(ctx, organization, project).Domain(domain).Metrics(metrics).TimestampFormat(timestampFormat).Execute()

Get hourly metrics



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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	domain := "domain_example" // string | Filter by domain ID or domain name (optional)
	metrics := []string{"Inner_example"} // []string | Metrics to return (default: hits, bytes) (optional)
	timestampFormat := "timestampFormat_example" // string | Timestamp format in response (optional) (default to "iso8601")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDNMetricsAPI.GetHourlyMetrics(context.Background(), organization, project).Domain(domain).Metrics(metrics).TimestampFormat(timestampFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDNMetricsAPI.GetHourlyMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHourlyMetrics`: V2MetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `CDNMetricsAPI.GetHourlyMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHourlyMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domain** | **string** | Filter by domain ID or domain name | 
 **metrics** | **[]string** | Metrics to return (default: hits, bytes) | 
 **timestampFormat** | **string** | Timestamp format in response | [default to &quot;iso8601&quot;]

### Return type

[**V2MetricsResponse**](V2MetricsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonthlyMetrics

> V2MetricsResponse GetMonthlyMetrics(ctx, organization, project).Domain(domain).Metrics(metrics).TimestampFormat(timestampFormat).Execute()

Get monthly metrics



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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	domain := "domain_example" // string | Filter by domain ID or domain name (optional)
	metrics := []string{"Inner_example"} // []string | Metrics to return (default: hits, bytes) (optional)
	timestampFormat := "timestampFormat_example" // string | Timestamp format in response (optional) (default to "iso8601")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CDNMetricsAPI.GetMonthlyMetrics(context.Background(), organization, project).Domain(domain).Metrics(metrics).TimestampFormat(timestampFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CDNMetricsAPI.GetMonthlyMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMonthlyMetrics`: V2MetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `CDNMetricsAPI.GetMonthlyMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonthlyMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domain** | **string** | Filter by domain ID or domain name | 
 **metrics** | **[]string** | Metrics to return (default: hits, bytes) | 
 **timestampFormat** | **string** | Timestamp format in response | [default to &quot;iso8601&quot;]

### Return type

[**V2MetricsResponse**](V2MetricsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

