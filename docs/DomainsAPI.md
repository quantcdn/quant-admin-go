# \DomainsAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationProjectsProjectDomainsDomainDelete**](DomainsAPI.md#OrganizationsOrganizationProjectsProjectDomainsDomainDelete) | **Delete** /organizations/{organization}/projects/{project}/domains/{domain} | Delete a domain
[**OrganizationsOrganizationProjectsProjectDomainsDomainGet**](DomainsAPI.md#OrganizationsOrganizationProjectsProjectDomainsDomainGet) | **Get** /organizations/{organization}/projects/{project}/domains/{domain} | Get domain details
[**OrganizationsOrganizationProjectsProjectDomainsDomainPatch**](DomainsAPI.md#OrganizationsOrganizationProjectsProjectDomainsDomainPatch) | **Patch** /organizations/{organization}/projects/{project}/domains/{domain} | Update domain details
[**OrganizationsOrganizationProjectsProjectDomainsGet**](DomainsAPI.md#OrganizationsOrganizationProjectsProjectDomainsGet) | **Get** /organizations/{organization}/projects/{project}/domains | List domains for a project
[**OrganizationsOrganizationProjectsProjectDomainsPost**](DomainsAPI.md#OrganizationsOrganizationProjectsProjectDomainsPost) | **Post** /organizations/{organization}/projects/{project}/domains | Create a new domain for a project



## OrganizationsOrganizationProjectsProjectDomainsDomainDelete

> OrganizationsOrganizationProjectsProjectDomainsGet200Response OrganizationsOrganizationProjectsProjectDomainsDomainDelete(ctx, organization, project, domain).Execute()

Delete a domain

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
    domain := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainDelete(context.Background(), organization, project, domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectDomainsDomainDelete`: OrganizationsOrganizationProjectsProjectDomainsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**domain** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectDomainsDomainDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectDomainsGet200Response**](OrganizationsOrganizationProjectsProjectDomainsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectDomainsDomainGet

> OrganizationsOrganizationProjectsProjectDomainsGet200Response OrganizationsOrganizationProjectsProjectDomainsDomainGet(ctx, organization, project, domain).Execute()

Get domain details

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
    domain := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainGet(context.Background(), organization, project, domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectDomainsDomainGet`: OrganizationsOrganizationProjectsProjectDomainsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**domain** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectDomainsDomainGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectDomainsGet200Response**](OrganizationsOrganizationProjectsProjectDomainsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectDomainsDomainPatch

> OrganizationsOrganizationProjectsProjectDomainsGet200Response OrganizationsOrganizationProjectsProjectDomainsDomainPatch(ctx, organization, project, domain).DomainRequest(domainRequest).Execute()

Update domain details

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
    domain := TODO // interface{} |
    domainRequest := *openapiclient.NewDomainRequest() // DomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainPatch(context.Background(), organization, project, domain).DomainRequest(domainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectDomainsDomainPatch`: OrganizationsOrganizationProjectsProjectDomainsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsDomainPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**domain** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectDomainsDomainPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **domainRequest** | [**DomainRequest**](DomainRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectDomainsGet200Response**](OrganizationsOrganizationProjectsProjectDomainsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectDomainsGet

> OrganizationsOrganizationProjectsProjectDomainsGet200Response OrganizationsOrganizationProjectsProjectDomainsGet(ctx, organization, project).Execute()

List domains for a project

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
    resp, r, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsGet(context.Background(), organization, project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectDomainsGet`: OrganizationsOrganizationProjectsProjectDomainsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectDomainsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationsOrganizationProjectsProjectDomainsGet200Response**](OrganizationsOrganizationProjectsProjectDomainsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectDomainsPost

> OrganizationsOrganizationProjectsProjectDomainsGet200Response OrganizationsOrganizationProjectsProjectDomainsPost(ctx, organization, project).DomainRequest(domainRequest).Execute()

Create a new domain for a project

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
    domainRequest := *openapiclient.NewDomainRequest() // DomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsPost(context.Background(), organization, project).DomainRequest(domainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectDomainsPost`: OrganizationsOrganizationProjectsProjectDomainsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.OrganizationsOrganizationProjectsProjectDomainsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectDomainsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domainRequest** | [**DomainRequest**](DomainRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectDomainsGet200Response**](OrganizationsOrganizationProjectsProjectDomainsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

