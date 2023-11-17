# \ProjectsAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationProjectsGet**](ProjectsAPI.md#OrganizationsOrganizationProjectsGet) | **Get** /organizations/{organization}/projects | Get projects for an organization
[**OrganizationsOrganizationProjectsPost**](ProjectsAPI.md#OrganizationsOrganizationProjectsPost) | **Post** /organizations/{organization}/projects | Create a new project for an organization
[**OrganizationsOrganizationProjectsProjectDelete**](ProjectsAPI.md#OrganizationsOrganizationProjectsProjectDelete) | **Delete** /organizations/{organization}/projects/{project} | Delete a project
[**OrganizationsOrganizationProjectsProjectGet**](ProjectsAPI.md#OrganizationsOrganizationProjectsProjectGet) | **Get** /organizations/{organization}/projects/{project} | Get project details
[**OrganizationsOrganizationProjectsProjectPatch**](ProjectsAPI.md#OrganizationsOrganizationProjectsProjectPatch) | **Patch** /organizations/{organization}/projects/{project} | Update project details



## OrganizationsOrganizationProjectsGet

> OrganizationsOrganizationProjectsGet200Response OrganizationsOrganizationProjectsGet(ctx, organization).Execute()

Get projects for an organization

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.OrganizationsOrganizationProjectsGet(context.Background(), organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.OrganizationsOrganizationProjectsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsGet`: OrganizationsOrganizationProjectsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.OrganizationsOrganizationProjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationsOrganizationProjectsGet200Response**](OrganizationsOrganizationProjectsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsPost

> OrganizationsOrganizationProjectsGet200Response OrganizationsOrganizationProjectsPost(ctx, organization).ProjectRequest(projectRequest).Execute()

Create a new project for an organization

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
    projectRequest := *openapiclient.NewProjectRequest() // ProjectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.OrganizationsOrganizationProjectsPost(context.Background(), organization).ProjectRequest(projectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.OrganizationsOrganizationProjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsPost`: OrganizationsOrganizationProjectsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.OrganizationsOrganizationProjectsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectRequest** | [**ProjectRequest**](ProjectRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsGet200Response**](OrganizationsOrganizationProjectsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectDelete

> OrganizationsOrganizationProjectsProjectGet200Response OrganizationsOrganizationProjectsProjectDelete(ctx, organization, project).Execute()

Delete a project

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
    resp, r, err := apiClient.ProjectsAPI.OrganizationsOrganizationProjectsProjectDelete(context.Background(), organization, project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.OrganizationsOrganizationProjectsProjectDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectDelete`: OrganizationsOrganizationProjectsProjectGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.OrganizationsOrganizationProjectsProjectDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationsOrganizationProjectsProjectGet200Response**](OrganizationsOrganizationProjectsProjectGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectGet

> OrganizationsOrganizationProjectsProjectGet200Response OrganizationsOrganizationProjectsProjectGet(ctx, organization, project).Execute()

Get project details

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
    resp, r, err := apiClient.ProjectsAPI.OrganizationsOrganizationProjectsProjectGet(context.Background(), organization, project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.OrganizationsOrganizationProjectsProjectGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectGet`: OrganizationsOrganizationProjectsProjectGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.OrganizationsOrganizationProjectsProjectGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationsOrganizationProjectsProjectGet200Response**](OrganizationsOrganizationProjectsProjectGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectPatch

> OrganizationsOrganizationProjectsProjectGet200Response OrganizationsOrganizationProjectsProjectPatch(ctx, organization, project).ProjectRequest(projectRequest).Execute()

Update project details

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
    projectRequest := *openapiclient.NewProjectRequest() // ProjectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsAPI.OrganizationsOrganizationProjectsProjectPatch(context.Background(), organization, project).ProjectRequest(projectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.OrganizationsOrganizationProjectsProjectPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectPatch`: OrganizationsOrganizationProjectsProjectGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.OrganizationsOrganizationProjectsProjectPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectRequest** | [**ProjectRequest**](ProjectRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectGet200Response**](OrganizationsOrganizationProjectsProjectGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

