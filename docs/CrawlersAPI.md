# \CrawlersAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCrawlers**](CrawlersAPI.md#CreateCrawlers) | **Post** /organizations/{organization}/projects/{project}/crawlers | 
[**DeleteCrawler**](CrawlersAPI.md#DeleteCrawler) | **Delete** /organizations/{organization}/projects/{project}/crawlers/{crawler} | 
[**GetCrawler**](CrawlersAPI.md#GetCrawler) | **Get** /organizations/{organization}/projects/{project}/crawlers/{crawler} | 
[**ListCrawlers**](CrawlersAPI.md#ListCrawlers) | **Get** /organizations/{organization}/projects/{project}/crawlers | 
[**UpdateCrawler**](CrawlersAPI.md#UpdateCrawler) | **Patch** /organizations/{organization}/projects/{project}/crawlers/{crawler} | 



## CreateCrawlers

> Crawler CreateCrawlers(ctx, organization, project).CrawlerRequest(crawlerRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name
	crawlerRequest := *openapiclient.NewCrawlerRequest() // CrawlerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.CreateCrawlers(context.Background(), organization, project).CrawlerRequest(crawlerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.CreateCrawlers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCrawlers`: Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.CreateCrawlers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCrawlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **crawlerRequest** | [**CrawlerRequest**](CrawlerRequest.md) |  | 

### Return type

[**Crawler**](Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCrawler

> Crawler DeleteCrawler(ctx, organization, project, crawler).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name
	crawler := "crawler_example" // string | Crawler uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.DeleteCrawler(context.Background(), organization, project, crawler).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.DeleteCrawler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCrawler`: Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.DeleteCrawler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**crawler** | **string** | Crawler uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrawlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Crawler**](Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrawler

> Crawler GetCrawler(ctx, organization, project, crawler).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name
	crawler := "crawler_example" // string | Crawler uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.GetCrawler(context.Background(), organization, project, crawler).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.GetCrawler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrawler`: Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.GetCrawler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**crawler** | **string** | Crawler uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrawlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Crawler**](Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCrawlers

> []Crawler ListCrawlers(ctx, organization, project).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.ListCrawlers(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.ListCrawlers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCrawlers`: []Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.ListCrawlers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCrawlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Crawler**](Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCrawler

> Crawler UpdateCrawler(ctx, organization, project, crawler).CrawlerRequest(crawlerRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name
	crawler := "crawler_example" // string | Crawler uuid
	crawlerRequest := *openapiclient.NewCrawlerRequest() // CrawlerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrawlersAPI.UpdateCrawler(context.Background(), organization, project, crawler).CrawlerRequest(crawlerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.UpdateCrawler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCrawler`: Crawler
	fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.UpdateCrawler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**crawler** | **string** | Crawler uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCrawlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **crawlerRequest** | [**CrawlerRequest**](CrawlerRequest.md) |  | 

### Return type

[**Crawler**](Crawler.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

