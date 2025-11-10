# \CrawlersAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrawlersCreate**](CrawlersAPI.md#CrawlersCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/crawlers | Create a new crawler
[**CrawlersDelete**](CrawlersAPI.md#CrawlersDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler} | Delete a crawler
[**CrawlersList**](CrawlersAPI.md#CrawlersList) | **Get** /api/v2/organizations/{organization}/projects/{project}/crawlers | List crawlers for the project
[**CrawlersRead**](CrawlersAPI.md#CrawlersRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler} | Get details of a single crawler
[**CrawlersUpdate**](CrawlersAPI.md#CrawlersUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler} | Update a crawler



## CrawlersCreate

> V2Crawler CrawlersCreate(ctx, organization, project).V2CrawlerRequest(v2CrawlerRequest).Execute()

Create a new crawler

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
	v2CrawlerRequest := *openapiclient.NewV2CrawlerRequest("test-domain.com") // V2CrawlerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.CrawlersCreate(context.Background(), organization, project).V2CrawlerRequest(v2CrawlerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlersCreate`: V2Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.CrawlersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2CrawlerRequest** | [**V2CrawlerRequest**](V2CrawlerRequest.md) |  | 

### Return type

[**V2Crawler**](V2Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlersDelete

> CrawlersDelete(ctx, organization, project, crawler).Execute()

Delete a crawler

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
	crawler := "crawler_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CrawlersAPI.CrawlersDelete(context.Background(), organization, project, crawler).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersDelete``: %v\n", err)
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
**crawler** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlersDeleteRequest struct via the builder pattern


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


## CrawlersList

> []V2Crawler CrawlersList(ctx, organization, project).Execute()

List crawlers for the project

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.CrawlersList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlersList`: []V2Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.CrawlersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2Crawler**](V2Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlersRead

> V2Crawler CrawlersRead(ctx, organization, project, crawler).Execute()

Get details of a single crawler

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
	crawler := "crawler_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.CrawlersRead(context.Background(), organization, project, crawler).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlersRead`: V2Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.CrawlersRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**crawler** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlersReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2Crawler**](V2Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlersUpdate

> V2Crawler CrawlersUpdate(ctx, organization, project, crawler).V2CrawlerRequest(v2CrawlerRequest).Execute()

Update a crawler

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
	crawler := "crawler_example" // string | 
	v2CrawlerRequest := *openapiclient.NewV2CrawlerRequest("test-domain.com") // V2CrawlerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.CrawlersUpdate(context.Background(), organization, project, crawler).V2CrawlerRequest(v2CrawlerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlersUpdate`: V2Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.CrawlersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**crawler** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2CrawlerRequest** | [**V2CrawlerRequest**](V2CrawlerRequest.md) |  | 

### Return type

[**V2Crawler**](V2Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

