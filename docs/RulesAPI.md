# \RulesAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/auth/{rule} | Delete an authentication rule
[**OrganizationsOrganizationProjectsProjectRulesAuthRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesAuthRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/auth/{rule} | Get authentication rule details
[**OrganizationsOrganizationProjectsProjectRulesAuthRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesAuthRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/auth/{rule} | Update authentication rule details
[**OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Delete a custom response rule
[**OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Get custom repsonse rule details
[**OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Update custom response rule details
[**OrganizationsOrganizationProjectsProjectRulesGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesGet) | **Get** /organizations/{organization}/projects/{project}/rules | Get rules for a project
[**OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/header/{rule} | Delete a header rule
[**OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/header/{rule} | Get header rule details
[**OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/header/{rule} | Update header rule details
[**OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | Delete a proxy rule
[**OrganizationsOrganizationProjectsProjectRulesProxyRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesProxyRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | Get proxy rule details
[**OrganizationsOrganizationProjectsProjectRulesProxyRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesProxyRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | Update proxy rule details
[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | Delete a redirect rule
[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | Get redirect rule details
[**OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | Update redirect rule details



## OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete(ctx, organization, project, rule).Execute()

Delete an authentication rule

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesAuthRuleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesAuthRuleGet

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesAuthRuleGet(ctx, organization, project, rule).Execute()

Get authentication rule details

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRuleGet(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesAuthRuleGet`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRuleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesAuthRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesAuthRulePatch

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesAuthRulePatch(ctx, organization, project, rule).OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest(organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest).Execute()

Update authentication rule details

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
    rule := TODO // interface{} |
    organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest := *openapiclient.NewOrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest() // OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRulePatch(context.Background(), organization, project, rule).OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest(organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesAuthRulePatch`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRulePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest** | [**OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest**](OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete(ctx, organization, project, rule).Execute()

Delete a custom response rule

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet(ctx, organization, project, rule).Execute()

Get custom repsonse rule details

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch(ctx, organization, project, rule).OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest(organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest).Execute()

Update custom response rule details

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
    rule := TODO // interface{} |
    organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest := *openapiclient.NewOrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest() // OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch(context.Background(), organization, project, rule).OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest(organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest** | [**OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest**](OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesGet

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesGet(ctx, organization, project).Execute()

Get rules for a project

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
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesGet(context.Background(), organization, project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesGet`: OrganizationsOrganizationProjectsProjectRulesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete(ctx, organization, project, rule).Execute()

Delete a header rule

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesHeaderRuleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet(ctx, organization, project, rule).Execute()

Get header rule details

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesHeaderRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch(ctx, organization, project, rule).OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest(organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest).Execute()

Update header rule details

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
    rule := TODO // interface{} |
    organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest := *openapiclient.NewOrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest() // OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch(context.Background(), organization, project, rule).OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest(organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesHeaderRulePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **organizationsOrganizationProjectsProjectRulesAuthRulePatchRequest** | [**OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest**](OrganizationsOrganizationProjectsProjectRulesAuthRulePatchRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete(ctx, organization, project, rule).Execute()

Delete a proxy rule

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesProxyRuleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesProxyRuleGet

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesProxyRuleGet(ctx, organization, project, rule).Execute()

Get proxy rule details

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRuleGet(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesProxyRuleGet`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRuleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesProxyRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesProxyRulePatch

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesProxyRulePatch(ctx, organization, project, rule).OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest(organizationsOrganizationProjectsProjectRulesProxyRulePatchRequest).Execute()

Update proxy rule details

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
    rule := TODO // interface{} |
    organizationsOrganizationProjectsProjectRulesProxyRulePatchRequest := *openapiclient.NewOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest() // OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRulePatch(context.Background(), organization, project, rule).OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest(organizationsOrganizationProjectsProjectRulesProxyRulePatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesProxyRulePatch`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRulePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **organizationsOrganizationProjectsProjectRulesProxyRulePatchRequest** | [**OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest**](OrganizationsOrganizationProjectsProjectRulesProxyRulePatchRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete(ctx, organization, project, rule).Execute()

Delete a redirect rule

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesRedirectRuleDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet(ctx, organization, project, rule).Execute()

Get redirect rule details

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
    rule := TODO // interface{} |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet(context.Background(), organization, project, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesRedirectRuleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch

> OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch(ctx, organization, project, rule).OrganizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest(organizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest).Execute()

Update redirect rule details

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
    rule := TODO // interface{} |
    organizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest := *openapiclient.NewOrganizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest() // OrganizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch(context.Background(), organization, project, rule).OrganizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest(organizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch`: OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | [**interface{}**](.md) | Organization machine name |
**project** | [**interface{}**](.md) | Project machine name |
**rule** | [**interface{}**](.md) |  |

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **organizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest** | [**OrganizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest**](OrganizationsOrganizationProjectsProjectRulesRedirectRulePatchRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response**](OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

