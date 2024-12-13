# \CrawlerSchedulesAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrawlerSchedulesCreate**](CrawlerSchedulesAPI.md#CrawlerSchedulesCreate) | **Post** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules | 
[**CrawlerSchedulesDelete**](CrawlerSchedulesAPI.md#CrawlerSchedulesDelete) | **Delete** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule} | 
[**CrawlerSchedulesList**](CrawlerSchedulesAPI.md#CrawlerSchedulesList) | **Get** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules | 
[**CrawlerSchedulesRead**](CrawlerSchedulesAPI.md#CrawlerSchedulesRead) | **Get** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule} | 
[**CrawlerSchedulesUpdate**](CrawlerSchedulesAPI.md#CrawlerSchedulesUpdate) | **Patch** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule} | 



## CrawlerSchedulesCreate

> CrawlerSchedule CrawlerSchedulesCreate(ctx, organization, project, crawler).CrawlerScheduleRequest(crawlerScheduleRequest).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	crawler := "crawler_example" // string | 
	crawlerScheduleRequest := *openapiclient.NewCrawlerScheduleRequest("ScheduleCronString_example") // CrawlerScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesCreate(context.Background(), organization, project, crawler).CrawlerScheduleRequest(crawlerScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesCreate`: CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerSchedulesAPI.CrawlerSchedulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**crawler** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **crawlerScheduleRequest** | [**CrawlerScheduleRequest**](CrawlerScheduleRequest.md) |  | 

### Return type

[**CrawlerSchedule**](CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlerSchedulesDelete

> CrawlerSchedule CrawlerSchedulesDelete(ctx, organization, project, crawler, crawlerSchedule).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	crawler := "crawler_example" // string | 
	crawlerSchedule := "crawlerSchedule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesDelete(context.Background(), organization, project, crawler, crawlerSchedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesDelete`: CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerSchedulesAPI.CrawlerSchedulesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**crawler** | **string** |  | 
**crawlerSchedule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**CrawlerSchedule**](CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlerSchedulesList

> []CrawlerSchedule CrawlerSchedulesList(ctx, organization, project, crawler).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	crawler := "crawler_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesList(context.Background(), organization, project, crawler).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesList`: []CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerSchedulesAPI.CrawlerSchedulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**crawler** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]CrawlerSchedule**](CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlerSchedulesRead

> CrawlerSchedule CrawlerSchedulesRead(ctx, organization, project, crawler, crawlerSchedule).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	crawler := "crawler_example" // string | 
	crawlerSchedule := "crawlerSchedule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesRead(context.Background(), organization, project, crawler, crawlerSchedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesRead`: CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerSchedulesAPI.CrawlerSchedulesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**crawler** | **string** |  | 
**crawlerSchedule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**CrawlerSchedule**](CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlerSchedulesUpdate

> CrawlerSchedule CrawlerSchedulesUpdate(ctx, organization, project, crawler, crawlerSchedule).CrawlerScheduleRequestUpdate(crawlerScheduleRequestUpdate).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	crawler := "crawler_example" // string | 
	crawlerSchedule := "crawlerSchedule_example" // string | 
	crawlerScheduleRequestUpdate := *openapiclient.NewCrawlerScheduleRequestUpdate() // CrawlerScheduleRequestUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesUpdate(context.Background(), organization, project, crawler, crawlerSchedule).CrawlerScheduleRequestUpdate(crawlerScheduleRequestUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesUpdate`: CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerSchedulesAPI.CrawlerSchedulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**crawler** | **string** |  | 
**crawlerSchedule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **crawlerScheduleRequestUpdate** | [**CrawlerScheduleRequestUpdate**](CrawlerScheduleRequestUpdate.md) |  | 

### Return type

[**CrawlerSchedule**](CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

