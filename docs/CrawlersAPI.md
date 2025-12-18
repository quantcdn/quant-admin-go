# \CrawlersAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrawlersCreate**](CrawlersAPI.md#CrawlersCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/crawlers | Create a new crawler
[**CrawlersDelete**](CrawlersAPI.md#CrawlersDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler} | Delete a crawler
[**CrawlersGetRunById**](CrawlersAPI.md#CrawlersGetRunById) | **Get** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler}/runs/{run_id} | Get a run by ID
[**CrawlersGetRuns**](CrawlersAPI.md#CrawlersGetRuns) | **Get** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler}/runs | Get all runs for a crawler
[**CrawlersList**](CrawlersAPI.md#CrawlersList) | **Get** /api/v2/organizations/{organization}/projects/{project}/crawlers | List crawlers for the project
[**CrawlersRead**](CrawlersAPI.md#CrawlersRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler} | Get details of a single crawler
[**CrawlersRun**](CrawlersAPI.md#CrawlersRun) | **Post** /api/v2/organizations/{organization}/projects/{project}/crawlers/{crawler}/run | Run a crawler
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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	crawler := "00000000-0000-0000-0000-000000000000" // string | The UUID of the crawler

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
**crawler** | **string** | The UUID of the crawler | 

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


## CrawlersGetRunById

> V2CrawlerRun CrawlersGetRunById(ctx, organization, project, crawler, runId).Execute()

Get a run by ID

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	crawler := "00000000-0000-0000-0000-000000000000" // string | Crawler identifier
	runId := int32(1) // int32 | Run identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.CrawlersGetRunById(context.Background(), organization, project, crawler, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersGetRunById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlersGetRunById`: V2CrawlerRun
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.CrawlersGetRunById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**crawler** | **string** | Crawler identifier | 
**runId** | **int32** | Run identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlersGetRunByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**V2CrawlerRun**](V2CrawlerRun.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlersGetRuns

> []V2CrawlerRun CrawlersGetRuns(ctx, organization, project, crawler).Execute()

Get all runs for a crawler

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	crawler := "00000000-0000-0000-0000-000000000000" // string | Crawler identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.CrawlersGetRuns(context.Background(), organization, project, crawler).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersGetRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlersGetRuns`: []V2CrawlerRun
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.CrawlersGetRuns`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCrawlersGetRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]V2CrawlerRun**](V2CrawlerRun.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrawlersList

> CrawlersList(ctx, organization, project).Execute()

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CrawlersAPI.CrawlersList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersList``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiCrawlersListRequest struct via the builder pattern


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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	crawler := "00000000-0000-0000-0000-000000000000" // string | The UUID of the crawler

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
**crawler** | **string** | The UUID of the crawler | 

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


## CrawlersRun

> CrawlersRun200Response CrawlersRun(ctx, organization, project, crawler).CrawlersRunRequest(crawlersRunRequest).Execute()

Run a crawler

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	crawler := "00000000-0000-0000-0000-000000000000" // string | Crawler identifier
	crawlersRunRequest := *openapiclient.NewCrawlersRunRequest() // CrawlersRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.CrawlersRun(context.Background(), organization, project, crawler).CrawlersRunRequest(crawlersRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CrawlersRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrawlersRun`: CrawlersRun200Response
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.CrawlersRun`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCrawlersRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **crawlersRunRequest** | [**CrawlersRunRequest**](CrawlersRunRequest.md) |  | 

### Return type

[**CrawlersRun200Response**](CrawlersRun200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	crawler := "00000000-0000-0000-0000-000000000000" // string | The UUID of the crawler
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
**crawler** | **string** | The UUID of the crawler | 

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

