# \AISearchAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiSearchChat**](AISearchAPI.md#AiSearchChat) | **Post** /api/v3/organisations/{organisation}/projects/{project}/ai-search/chat | RAG chat with AI Search content
[**AiSearchDeletePages**](AISearchAPI.md#AiSearchDeletePages) | **Delete** /api/v3/organisations/{organisation}/projects/{project}/ai-search/pages | Delete pages by URLs or patterns
[**AiSearchDisable**](AISearchAPI.md#AiSearchDisable) | **Post** /api/v3/organisations/{organisation}/projects/{project}/ai-search/disable | Disable AI Search for a project
[**AiSearchEnable**](AISearchAPI.md#AiSearchEnable) | **Post** /api/v3/organisations/{organisation}/projects/{project}/ai-search/enable | Enable AI Search for a project
[**AiSearchGetCrawl**](AISearchAPI.md#AiSearchGetCrawl) | **Get** /api/v3/organisations/{organisation}/projects/{project}/ai-search/crawls/{jobId} | Get AI Search ingest job status
[**AiSearchGetCrawlPages**](AISearchAPI.md#AiSearchGetCrawlPages) | **Get** /api/v3/organisations/{organisation}/projects/{project}/ai-search/crawls/{jobId}/pages | Get per-page ingest results for a crawl job
[**AiSearchGetSettings**](AISearchAPI.md#AiSearchGetSettings) | **Get** /api/v3/organisations/{organisation}/projects/{project}/ai-search/settings | Get AI Search public access and rate limit settings
[**AiSearchIngestPages**](AISearchAPI.md#AiSearchIngestPages) | **Post** /api/v3/organisations/{organisation}/projects/{project}/ai-search/pages | Ingest pages into the AI Search index
[**AiSearchListCrawls**](AISearchAPI.md#AiSearchListCrawls) | **Get** /api/v3/organisations/{organisation}/projects/{project}/ai-search/crawls | List AI Search ingest jobs
[**AiSearchListPages**](AISearchAPI.md#AiSearchListPages) | **Get** /api/v3/organisations/{organisation}/projects/{project}/ai-search/pages | List indexed pages with cursor pagination
[**AiSearchPurgeIndex**](AISearchAPI.md#AiSearchPurgeIndex) | **Delete** /api/v3/organisations/{organisation}/projects/{project}/ai-search/index | Purge the entire AI Search index
[**AiSearchSearch**](AISearchAPI.md#AiSearchSearch) | **Post** /api/v3/organisations/{organisation}/projects/{project}/ai-search/search | Semantic search across the AI Search index
[**AiSearchStatus**](AISearchAPI.md#AiSearchStatus) | **Get** /api/v3/organisations/{organisation}/projects/{project}/ai-search | Get AI Search status for a project
[**AiSearchTopQueries**](AISearchAPI.md#AiSearchTopQueries) | **Get** /api/v3/organisations/{organisation}/projects/{project}/ai-search/top-queries | Get the most popular AI Search queries
[**AiSearchTriggerCrawl**](AISearchAPI.md#AiSearchTriggerCrawl) | **Post** /api/v3/organisations/{organisation}/projects/{project}/ai-search/crawls | Trigger a crawler run that ingests into AI Search
[**AiSearchUpdateSettings**](AISearchAPI.md#AiSearchUpdateSettings) | **Put** /api/v3/organisations/{organisation}/projects/{project}/ai-search/settings | Update AI Search public access and rate limit settings
[**AiSearchUsage**](AISearchAPI.md#AiSearchUsage) | **Get** /api/v3/organisations/{organisation}/projects/{project}/ai-search/usage | Get usage statistics for the AI Search site



## AiSearchChat

> AiSearchChat(ctx, organisation, project).AiSearchChatRequest(aiSearchChatRequest).Execute()

RAG chat with AI Search content

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
	project := "project_example" // string | 
	aiSearchChatRequest := *openapiclient.NewAiSearchChatRequest("Message_example") // AiSearchChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchChat(context.Background(), organisation, project).AiSearchChatRequest(aiSearchChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aiSearchChatRequest** | [**AiSearchChatRequest**](AiSearchChatRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSearchDeletePages

> AiSearchDeletePages(ctx, organisation, project).AiSearchDeletePagesRequest(aiSearchDeletePagesRequest).Execute()

Delete pages by URLs or patterns

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
	project := "project_example" // string | 
	aiSearchDeletePagesRequest := *openapiclient.NewAiSearchDeletePagesRequest() // AiSearchDeletePagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchDeletePages(context.Background(), organisation, project).AiSearchDeletePagesRequest(aiSearchDeletePagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchDeletePages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchDeletePagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aiSearchDeletePagesRequest** | [**AiSearchDeletePagesRequest**](AiSearchDeletePagesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSearchDisable

> AiSearchDisable(ctx, organisation, project).Execute()

Disable AI Search for a project

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
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchDisable(context.Background(), organisation, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AiSearchEnable

> AiSearchEnable(ctx, organisation, project).AiSearchEnableRequest(aiSearchEnableRequest).Execute()

Enable AI Search for a project

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
	project := "project_example" // string | 
	aiSearchEnableRequest := *openapiclient.NewAiSearchEnableRequest() // AiSearchEnableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchEnable(context.Background(), organisation, project).AiSearchEnableRequest(aiSearchEnableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aiSearchEnableRequest** | [**AiSearchEnableRequest**](AiSearchEnableRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSearchGetCrawl

> AiSearchGetCrawl(ctx, organisation, project, jobId).Execute()

Get AI Search ingest job status

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
	project := "project_example" // string | 
	jobId := "jobId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchGetCrawl(context.Background(), organisation, project, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchGetCrawl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchGetCrawlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## AiSearchGetCrawlPages

> AiSearchGetCrawlPages(ctx, organisation, project, jobId).Limit(limit).StatusCode(statusCode).ProcessingStatus(processingStatus).Execute()

Get per-page ingest results for a crawl job

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
	project := "project_example" // string | 
	jobId := "jobId_example" // string | 
	limit := int32(56) // int32 |  (optional)
	statusCode := int32(56) // int32 |  (optional)
	processingStatus := "processingStatus_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchGetCrawlPages(context.Background(), organisation, project, jobId).Limit(limit).StatusCode(statusCode).ProcessingStatus(processingStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchGetCrawlPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchGetCrawlPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** |  | 
 **statusCode** | **int32** |  | 
 **processingStatus** | **string** |  | 

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


## AiSearchGetSettings

> AiSearchGetSettings(ctx, organisation, project).Execute()

Get AI Search public access and rate limit settings

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
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchGetSettings(context.Background(), organisation, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchGetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AiSearchIngestPages

> AiSearchIngestPages(ctx, organisation, project).AiSearchIngestPagesRequest(aiSearchIngestPagesRequest).Execute()

Ingest pages into the AI Search index

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
	project := "project_example" // string | 
	aiSearchIngestPagesRequest := *openapiclient.NewAiSearchIngestPagesRequest([]openapiclient.AiSearchIngestPagesRequestPagesInner{*openapiclient.NewAiSearchIngestPagesRequestPagesInner("Url_example", "Title_example", "Content_example")}) // AiSearchIngestPagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchIngestPages(context.Background(), organisation, project).AiSearchIngestPagesRequest(aiSearchIngestPagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchIngestPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchIngestPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aiSearchIngestPagesRequest** | [**AiSearchIngestPagesRequest**](AiSearchIngestPagesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSearchListCrawls

> AiSearchListCrawls(ctx, organisation, project).Limit(limit).Execute()

List AI Search ingest jobs

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
	project := "project_example" // string | 
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchListCrawls(context.Background(), organisation, project).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchListCrawls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchListCrawlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | 

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


## AiSearchListPages

> AiSearchListPages(ctx, organisation, project).Limit(limit).Cursor(cursor).Search(search).Execute()

List indexed pages with cursor pagination

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
	project := "project_example" // string | 
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchListPages(context.Background(), organisation, project).Limit(limit).Cursor(cursor).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchListPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchListPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | 
 **cursor** | **string** |  | 
 **search** | **string** |  | 

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


## AiSearchPurgeIndex

> AiSearchPurgeIndex(ctx, organisation, project).Execute()

Purge the entire AI Search index

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
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchPurgeIndex(context.Background(), organisation, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchPurgeIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchPurgeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AiSearchSearch

> AiSearchSearch(ctx, organisation, project).AiSearchSearchRequest(aiSearchSearchRequest).Execute()

Semantic search across the AI Search index

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
	project := "project_example" // string | 
	aiSearchSearchRequest := *openapiclient.NewAiSearchSearchRequest("Query_example") // AiSearchSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchSearch(context.Background(), organisation, project).AiSearchSearchRequest(aiSearchSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aiSearchSearchRequest** | [**AiSearchSearchRequest**](AiSearchSearchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSearchStatus

> AiSearchStatus(ctx, organisation, project).Execute()

Get AI Search status for a project

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
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchStatus(context.Background(), organisation, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AiSearchTopQueries

> AiSearchTopQueries(ctx, organisation, project).Range_(range_).Limit(limit).Execute()

Get the most popular AI Search queries

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
	project := "project_example" // string | 
	range_ := "range__example" // string |  (optional) (default to "30d")
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchTopQueries(context.Background(), organisation, project).Range_(range_).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchTopQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchTopQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **range_** | **string** |  | [default to &quot;30d&quot;]
 **limit** | **int32** |  | 

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


## AiSearchTriggerCrawl

> AiSearchTriggerCrawl(ctx, organisation, project).AiSearchTriggerCrawlRequest(aiSearchTriggerCrawlRequest).Execute()

Trigger a crawler run that ingests into AI Search

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
	project := "project_example" // string | 
	aiSearchTriggerCrawlRequest := *openapiclient.NewAiSearchTriggerCrawlRequest("CrawlerUuid_example") // AiSearchTriggerCrawlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchTriggerCrawl(context.Background(), organisation, project).AiSearchTriggerCrawlRequest(aiSearchTriggerCrawlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchTriggerCrawl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchTriggerCrawlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aiSearchTriggerCrawlRequest** | [**AiSearchTriggerCrawlRequest**](AiSearchTriggerCrawlRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSearchUpdateSettings

> AiSearchUpdateSettings(ctx, organisation, project).AiSearchUpdateSettingsRequest(aiSearchUpdateSettingsRequest).Execute()

Update AI Search public access and rate limit settings

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
	project := "project_example" // string | 
	aiSearchUpdateSettingsRequest := *openapiclient.NewAiSearchUpdateSettingsRequest() // AiSearchUpdateSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchUpdateSettings(context.Background(), organisation, project).AiSearchUpdateSettingsRequest(aiSearchUpdateSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchUpdateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aiSearchUpdateSettingsRequest** | [**AiSearchUpdateSettingsRequest**](AiSearchUpdateSettingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSearchUsage

> AiSearchUsage(ctx, organisation, project).Range_(range_).Execute()

Get usage statistics for the AI Search site

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
	project := "project_example" // string | 
	range_ := "range__example" // string |  (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISearchAPI.AiSearchUsage(context.Background(), organisation, project).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISearchAPI.AiSearchUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiSearchUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **range_** | **string** |  | [default to &quot;30d&quot;]

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

