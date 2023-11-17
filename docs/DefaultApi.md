# \DefaultApi

All URIs are relative to *http://localhost:8001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DefaultApi.md#CreateDomain) | **Post** /domain/create | Create domaim
[**CreateProject**](DefaultApi.md#CreateProject) | **Post** /project/create | Create project
[**DeleteDomain**](DefaultApi.md#DeleteDomain) | **Delete** /domain/delete | Delete domain
[**DeleteProject**](DefaultApi.md#DeleteProject) | **Delete** /project/delete | Delete project
[**EditProject**](DefaultApi.md#EditProject) | **Patch** /project/edit | Edit project
[**GetCrawlRun**](DefaultApi.md#GetCrawlRun) | **Get** /crawl/runs/{runId} | Get crawl run status
[**GetCrawlers**](DefaultApi.md#GetCrawlers) | **Get** /crawl/configs | Get crawl configs
[**GetOrganisations**](DefaultApi.md#GetOrganisations) | **Get** /organisations | Get organisations
[**GetProject**](DefaultApi.md#GetProject) | **Get** /project | Get project
[**GetProjects**](DefaultApi.md#GetProjects) | **Get** /projects | Get projects
[**RunCrawler**](DefaultApi.md#RunCrawler) | **Post** /crawl/run/{uuid} | Run crawl config
[**SslRenewDomain**](DefaultApi.md#SslRenewDomain) | **Post** /domain/ssl-renew | Renew domain SSL certificate



## CreateDomain

> Message CreateDomain(ctx).QuantOrganisation(quantOrganisation).QuantProject(quantProject).DomainCreate(domainCreate).Execute()

Create domaim



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn" // string |
    domainCreate := *openapiclient.NewDomainCreate("example.quantcdn.io") // DomainCreate | Domain schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateDomain(context.Background()).QuantOrganisation(quantOrganisation).QuantProject(quantProject).DomainCreate(domainCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomain`: Message
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |
 **domainCreate** | [**DomainCreate**](DomainCreate.md) | Domain schema |

### Return type

[**Message**](Message.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> ProjectResponse CreateProject(ctx).QuantOrganisation(quantOrganisation).ProjectCreate(projectCreate).Execute()

Create project



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
    quantOrganisation := "quant" // string |
    projectCreate := *openapiclient.NewProjectCreate("quantcdn-project", "au") // ProjectCreate | Project schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateProject(context.Background()).QuantOrganisation(quantOrganisation).ProjectCreate(projectCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **projectCreate** | [**ProjectCreate**](ProjectCreate.md) | Project schema |

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomain

> Message DeleteDomain(ctx).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Domain(domain).Execute()

Delete domain



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn" // string |
    domain := "example.quantcdn.io" // string |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteDomain(context.Background()).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDomain`: Message
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |
 **domain** | **string** |  |

### Return type

[**Message**](Message.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> Message DeleteProject(ctx).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Execute()

Delete project



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn" // string |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteProject(context.Background()).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProject`: Message
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |

### Return type

[**Message**](Message.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProject

> ProjectResponse EditProject(ctx).QuantOrganisation(quantOrganisation).QuantProject(quantProject).ProjectEdit(projectEdit).Execute()

Edit project



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn" // string |
    projectEdit := *openapiclient.NewProjectEdit("quantcdn-project", "au", true) // ProjectEdit | Project schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EditProject(context.Background()).QuantOrganisation(quantOrganisation).QuantProject(quantProject).ProjectEdit(projectEdit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProject`: ProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EditProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |
 **projectEdit** | [**ProjectEdit**](ProjectEdit.md) | Project schema |

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrawlRun

> GetCrawlRun200Response GetCrawlRun(ctx, runId).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Execute()

Get crawl run status



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn" // string |
    runId := "runId_example" // string | Run ID to retrieve status for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCrawlRun(context.Background(), runId).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCrawlRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrawlRun`: GetCrawlRun200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCrawlRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Run ID to retrieve status for |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrawlRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |


### Return type

[**GetCrawlRun200Response**](GetCrawlRun200Response.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrawlers

> CrawlerConfigResponse GetCrawlers(ctx).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Execute()

Get crawl configs



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn" // string |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCrawlers(context.Background()).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCrawlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrawlers`: CrawlerConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCrawlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrawlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |

### Return type

[**CrawlerConfigResponse**](CrawlerConfigResponse.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisations

> Organisations GetOrganisations(ctx).Execute()

Get organisations



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetOrganisations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOrganisations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganisations`: Organisations
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOrganisations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationsRequest struct via the builder pattern


### Return type

[**Organisations**](Organisations.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> InlineResponse200 GetProject(ctx).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Execute()

Get project



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn" // string |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetProject(context.Background()).QuantOrganisation(quantOrganisation).QuantProject(quantProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> Projects GetProjects(ctx).QuantOrganisation(quantOrganisation).Execute()

Get projects



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
    quantOrganisation := "quant" // string |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetProjects(context.Background()).QuantOrganisation(quantOrganisation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjects`: Projects
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |

### Return type

[**Projects**](Projects.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunCrawler

> Message RunCrawler(ctx, uuid).QuantOrganisation(quantOrganisation).QuantProject(quantProject).RunCrawlerRequest(runCrawlerRequest).Execute()

Run crawl config



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn-project" // string |
    uuid := "uuid_example" // string | UUID of the crawler to run
    runCrawlerRequest := *openapiclient.NewRunCrawlerRequest() // RunCrawlerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RunCrawler(context.Background(), uuid).QuantOrganisation(quantOrganisation).QuantProject(quantProject).RunCrawlerRequest(runCrawlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RunCrawler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunCrawler`: Message
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RunCrawler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the crawler to run |

### Other Parameters

Other parameters are passed through a pointer to a apiRunCrawlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |

 **runCrawlerRequest** | [**RunCrawlerRequest**](RunCrawlerRequest.md) |  |

### Return type

[**Message**](Message.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SslRenewDomain

> Message SslRenewDomain(ctx).QuantOrganisation(quantOrganisation).QuantProject(quantProject).DomainSSLRenew(domainSSLRenew).Execute()

Renew domain SSL certificate



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
    quantOrganisation := "quant" // string |
    quantProject := "quantcdn-project" // string |
    domainSSLRenew := *openapiclient.NewDomainSSLRenew("example.quantcdn.io") // DomainSSLRenew |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SslRenewDomain(context.Background()).QuantOrganisation(quantOrganisation).QuantProject(quantProject).DomainSSLRenew(domainSSLRenew).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SslRenewDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslRenewDomain`: Message
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SslRenewDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSslRenewDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantOrganisation** | **string** |  |
 **quantProject** | **string** |  |
 **domainSSLRenew** | [**DomainSSLRenew**](DomainSSLRenew.md) |  |

### Return type

[**Message**](Message.md)

### Authorization

[ApiTokenAuth](../README.md#ApiTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

