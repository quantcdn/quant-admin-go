# \CrawlerAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrawlerSchedulesCreate**](CrawlerAPI.md#CrawlerSchedulesCreate) | **Post** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedule | 
[**CrawlerSchedulesDelete**](CrawlerAPI.md#CrawlerSchedulesDelete) | **Delete** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedule/{schedule} | 
[**CrawlerSchedulesList**](CrawlerAPI.md#CrawlerSchedulesList) | **Get** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedule | 
[**CrawlerSchedulesRead**](CrawlerAPI.md#CrawlerSchedulesRead) | **Get** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedule/{schedule} | 
[**CrawlerSchedulesUpdate**](CrawlerAPI.md#CrawlerSchedulesUpdate) | **Patch** /organizations/{organization}/projects/{project}/crawlers/{crawler}/schedule/{schedule} | 



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
	resp, r, err := apiClient.CrawlerAPI.CrawlerSchedulesCreate(context.Background(), organization, project, crawler).CrawlerScheduleRequest(crawlerScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerAPI.CrawlerSchedulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesCreate`: CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerAPI.CrawlerSchedulesCreate`: %v\n", resp)
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

> CrawlerSchedule CrawlerSchedulesDelete(ctx, organization, project, crawler, schedule).Execute()



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
	schedule := "schedule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerAPI.CrawlerSchedulesDelete(context.Background(), organization, project, crawler, schedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerAPI.CrawlerSchedulesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesDelete`: CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerAPI.CrawlerSchedulesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**crawler** | **string** |  | 
**schedule** | **string** |  | 

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
	resp, r, err := apiClient.CrawlerAPI.CrawlerSchedulesList(context.Background(), organization, project, crawler).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerAPI.CrawlerSchedulesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesList`: []CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerAPI.CrawlerSchedulesList`: %v\n", resp)
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

> CrawlerSchedule CrawlerSchedulesRead(ctx, organization, project, crawler, schedule).Execute()



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
	schedule := "schedule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerAPI.CrawlerSchedulesRead(context.Background(), organization, project, crawler, schedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerAPI.CrawlerSchedulesRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesRead`: CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerAPI.CrawlerSchedulesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**crawler** | **string** |  | 
**schedule** | **string** |  | 

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

> CrawlerSchedule CrawlerSchedulesUpdate(ctx, organization, project, crawler, schedule).CrawlerScheduleRequestUpdate(crawlerScheduleRequestUpdate).Execute()



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
	schedule := "schedule_example" // string | 
	crawlerScheduleRequestUpdate := *openapiclient.NewCrawlerScheduleRequestUpdate() // CrawlerScheduleRequestUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlerAPI.CrawlerSchedulesUpdate(context.Background(), organization, project, crawler, schedule).CrawlerScheduleRequestUpdate(crawlerScheduleRequestUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlerAPI.CrawlerSchedulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlerSchedulesUpdate`: CrawlerSchedule
	fmt.Fprintf(os.Stdout, "Response from `CrawlerAPI.CrawlerSchedulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**crawler** | **string** |  | 
**schedule** | **string** |  | 

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

