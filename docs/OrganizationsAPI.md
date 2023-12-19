# \OrganizationsAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsGet**](OrganizationsAPI.md#OrganizationsGet) | **Get** /organizations | List organizations
[**OrganizationsOrganizationGet**](OrganizationsAPI.md#OrganizationsOrganizationGet) | **Get** /organizations/{organization} | Get organization by machine name



## OrganizationsGet

> OrganizationsGet200Response OrganizationsGet(ctx).Execute()

List organizations

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
    resp, r, err := apiClient.OrganizationsAPI.OrganizationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsGet`: OrganizationsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsGetRequest struct via the builder pattern


### Return type

[**OrganizationsGet200Response**](OrganizationsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationGet

> OrganizationsGet200Response OrganizationsOrganizationGet(ctx, organization).Execute()

Get organization by machine name

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
    resp, r, err := apiClient.OrganizationsAPI.OrganizationsOrganizationGet(context.Background(), organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.OrganizationsOrganizationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationGet`: OrganizationsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.OrganizationsOrganizationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationsGet200Response**](OrganizationsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

