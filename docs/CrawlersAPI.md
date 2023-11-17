# \CrawlersAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationProjectsProjectCrawlersCrawlerDelete**](CrawlersAPI.md#OrganizationsOrganizationProjectsProjectCrawlersCrawlerDelete) | **Delete** /organizations/{organization}/projects/{project}/crawlers/{crawler} | Delete a crawler
[**OrganizationsOrganizationProjectsProjectCrawlersCrawlerGet**](CrawlersAPI.md#OrganizationsOrganizationProjectsProjectCrawlersCrawlerGet) | **Get** /organizations/{organization}/projects/{project}/crawlers/{crawler} | Get crawler details
[**OrganizationsOrganizationProjectsProjectCrawlersCrawlerPatch**](CrawlersAPI.md#OrganizationsOrganizationProjectsProjectCrawlersCrawlerPatch) | **Patch** /organizations/{organization}/projects/{project}/crawlers/{crawler} | Update crawler details
[**OrganizationsOrganizationProjectsProjectCrawlersGet**](CrawlersAPI.md#OrganizationsOrganizationProjectsProjectCrawlersGet) | **Get** /organizations/{organization}/projects/{project}/crawlers | Get crawlers for a project
[**OrganizationsOrganizationProjectsProjectCrawlersPost**](CrawlersAPI.md#OrganizationsOrganizationProjectsProjectCrawlersPost) | **Post** /organizations/{organization}/projects/{project}/crawlers | Create a new crawler for a project



## OrganizationsOrganizationProjectsProjectCrawlersCrawlerDelete

> OrganizationsOrganizationProjectsProjectCrawlersGet200Response OrganizationsOrganizationProjectsProjectCrawlersCrawlerDelete(ctx, organization, project, crawler).Execute()

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    crawler := TODO // interface{} | Crawler uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerDelete(context.Background(), organization, project, crawler).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectCrawlersCrawlerDelete`: OrganizationsOrganizationProjectsProjectCrawlersGet200Response
    fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**crawler** | [**interface{}**](.md) | Crawler uuid |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectCrawlersCrawlerDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectCrawlersGet200Response**](OrganizationsOrganizationProjectsProjectCrawlersGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectCrawlersCrawlerGet

> OrganizationsOrganizationProjectsProjectCrawlersGet200Response OrganizationsOrganizationProjectsProjectCrawlersCrawlerGet(ctx, organization, project, crawler).Execute()

Get crawler details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    crawler := TODO // interface{} | Crawler uuid

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerGet(context.Background(), organization, project, crawler).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectCrawlersCrawlerGet`: OrganizationsOrganizationProjectsProjectCrawlersGet200Response
    fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**crawler** | [**interface{}**](.md) | Crawler uuid |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectCrawlersCrawlerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectCrawlersGet200Response**](OrganizationsOrganizationProjectsProjectCrawlersGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectCrawlersCrawlerPatch

> OrganizationsOrganizationProjectsProjectCrawlersGet200Response OrganizationsOrganizationProjectsProjectCrawlersCrawlerPatch(ctx, organization, project, crawler).OrganizationsOrganizationProjectsProjectCrawlersPostRequest(organizationsOrganizationProjectsProjectCrawlersPostRequest).Execute()

Update crawler details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    crawler := TODO // interface{} | Crawler uuid
    organizationsOrganizationProjectsProjectCrawlersPostRequest := *openapiclient.NewOrganizationsOrganizationProjectsProjectCrawlersPostRequest() // OrganizationsOrganizationProjectsProjectCrawlersPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerPatch(context.Background(), organization, project, crawler).OrganizationsOrganizationProjectsProjectCrawlersPostRequest(organizationsOrganizationProjectsProjectCrawlersPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectCrawlersCrawlerPatch`: OrganizationsOrganizationProjectsProjectCrawlersGet200Response
    fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersCrawlerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**crawler** | [**interface{}**](.md) | Crawler uuid |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectCrawlersCrawlerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **organizationsOrganizationProjectsProjectCrawlersPostRequest** | [**OrganizationsOrganizationProjectsProjectCrawlersPostRequest**](OrganizationsOrganizationProjectsProjectCrawlersPostRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectCrawlersGet200Response**](OrganizationsOrganizationProjectsProjectCrawlersGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectCrawlersGet

> OrganizationsOrganizationProjectsProjectCrawlersGet200Response OrganizationsOrganizationProjectsProjectCrawlersGet(ctx, organization, project).Execute()

Get crawlers for a project

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersGet(context.Background(), organization, project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectCrawlersGet`: OrganizationsOrganizationProjectsProjectCrawlersGet200Response
    fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectCrawlersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationsOrganizationProjectsProjectCrawlersGet200Response**](OrganizationsOrganizationProjectsProjectCrawlersGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectCrawlersPost

> OrganizationsOrganizationProjectsProjectCrawlersGet200Response OrganizationsOrganizationProjectsProjectCrawlersPost(ctx, organization, project).OrganizationsOrganizationProjectsProjectCrawlersPostRequest(organizationsOrganizationProjectsProjectCrawlersPostRequest).Execute()

Create a new crawler for a project

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    organizationsOrganizationProjectsProjectCrawlersPostRequest := *openapiclient.NewOrganizationsOrganizationProjectsProjectCrawlersPostRequest() // OrganizationsOrganizationProjectsProjectCrawlersPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersPost(context.Background(), organization, project).OrganizationsOrganizationProjectsProjectCrawlersPostRequest(organizationsOrganizationProjectsProjectCrawlersPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectCrawlersPost`: OrganizationsOrganizationProjectsProjectCrawlersGet200Response
    fmt.Fprintf(os.Stdout, "Response from `CrawlersAPI.OrganizationsOrganizationProjectsProjectCrawlersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectCrawlersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationsOrganizationProjectsProjectCrawlersPostRequest** | [**OrganizationsOrganizationProjectsProjectCrawlersPostRequest**](OrganizationsOrganizationProjectsProjectCrawlersPostRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectCrawlersGet200Response**](OrganizationsOrganizationProjectsProjectCrawlersGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

