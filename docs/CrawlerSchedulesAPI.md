# \CrawlerSchedulesAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrawlerSchedulesAdd**](CrawlerSchedulesAPI.md#CrawlerSchedulesAdd) | **Post** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules | Add a new schedule
[**CrawlerSchedulesDelete**](CrawlerSchedulesAPI.md#CrawlerSchedulesDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule} | Delete a schedule
[**CrawlerSchedulesEdit**](CrawlerSchedulesAPI.md#CrawlerSchedulesEdit) | **Patch** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule} | Edit a schedule
[**CrawlerSchedulesList**](CrawlerSchedulesAPI.md#CrawlerSchedulesList) | **Get** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules | List schedules for a crawler
[**CrawlerSchedulesShow**](CrawlerSchedulesAPI.md#CrawlerSchedulesShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler}/schedules/{crawler_schedule} | Show a specific schedule



## CrawlerSchedulesAdd

> V2CrawlerSchedule CrawlerSchedulesAdd(ctx, organization, project, crawler).V2CrawlerScheduleRequest(v2CrawlerScheduleRequest).Execute()

Add a new schedule

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
	crawler := "crawler_example" // string | Crawler identifier
	v2CrawlerScheduleRequest := *openapiclient.NewV2CrawlerScheduleRequest("Test schedule", "0 2 * * *") // V2CrawlerScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesAdd(context.Background(), organization, project, crawler).V2CrawlerScheduleRequest(v2CrawlerScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesAdd`: V2CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerSchedulesAPI.CrawlerSchedulesAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**crawler** | **string** | Crawler identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2CrawlerScheduleRequest** | [**V2CrawlerScheduleRequest**](V2CrawlerScheduleRequest.md) |  | 

### Return type

[**V2CrawlerSchedule**](V2CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlerSchedulesDelete

> CrawlerSchedulesDelete(ctx, organization, project, crawler, crawlerSchedule).Execute()

Delete a schedule

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
	crawler := "crawler_example" // string | Crawler identifier
	crawlerSchedule := "crawlerSchedule_example" // string | Crawler schedule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesDelete(context.Background(), organization, project, crawler, crawlerSchedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesDelete``: %v\n", err)
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
**crawler** | **string** | Crawler identifier | 
**crawlerSchedule** | **string** | Crawler schedule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesDeleteRequest struct via the builder pattern


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


## CrawlerSchedulesEdit

> V2CrawlerSchedule CrawlerSchedulesEdit(ctx, organization, project, crawler, crawlerSchedule).V2CrawlerScheduleRequest(v2CrawlerScheduleRequest).Execute()

Edit a schedule

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
	crawler := "crawler_example" // string | Crawler identifier
	crawlerSchedule := "crawlerSchedule_example" // string | Crawler schedule identifier
	v2CrawlerScheduleRequest := *openapiclient.NewV2CrawlerScheduleRequest("Test schedule", "0 2 * * *") // V2CrawlerScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesEdit(context.Background(), organization, project, crawler, crawlerSchedule).V2CrawlerScheduleRequest(v2CrawlerScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesEdit`: V2CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerSchedulesAPI.CrawlerSchedulesEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**crawler** | **string** | Crawler identifier | 
**crawlerSchedule** | **string** | Crawler schedule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2CrawlerScheduleRequest** | [**V2CrawlerScheduleRequest**](V2CrawlerScheduleRequest.md) |  | 

### Return type

[**V2CrawlerSchedule**](V2CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlerSchedulesList

> []V2CrawlerSchedule CrawlerSchedulesList(ctx, organization, project, crawler).Execute()

List schedules for a crawler

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
	// response from `CrawlerSchedulesList`: []V2CrawlerSchedule
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

[**[]V2CrawlerSchedule**](V2CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlerSchedulesShow

> V2CrawlerSchedule CrawlerSchedulesShow(ctx, organization, project, crawler, crawlerSchedule).Execute()

Show a specific schedule

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
	crawler := "crawler_example" // string | Crawler identifier
	crawlerSchedule := "crawlerSchedule_example" // string | Crawler schedule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerSchedulesAPI.CrawlerSchedulesShow(context.Background(), organization, project, crawler, crawlerSchedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerSchedulesAPI.CrawlerSchedulesShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesShow`: V2CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerSchedulesAPI.CrawlerSchedulesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**crawler** | **string** | Crawler identifier | 
**crawlerSchedule** | **string** | Crawler schedule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlerSchedulesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**V2CrawlerSchedule**](V2CrawlerSchedule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

