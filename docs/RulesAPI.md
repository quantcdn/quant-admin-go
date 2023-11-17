# \RulesAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsOrganizationProjectsProjectRulesAuthPost**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesAuthPost) | **Post** /organizations/{organization}/projects/{project}/rules/auth | Create authentication rules
[**OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/auth/{rule} | Delete an authentication rule
[**OrganizationsOrganizationProjectsProjectRulesAuthRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesAuthRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/auth/{rule} | Get authentication rule details
[**OrganizationsOrganizationProjectsProjectRulesAuthRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesAuthRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/auth/{rule} | Update authentication rule details
[**OrganizationsOrganizationProjectsProjectRulesCustomResponsePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesCustomResponsePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/custom-response | Update custom response rule details
[**OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Delete a custom response rule
[**OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Get custom repsonse rule details
[**OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Update custom response rule details
[**OrganizationsOrganizationProjectsProjectRulesGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesGet) | **Get** /organizations/{organization}/projects/{project}/rules | Get rules for a project
[**OrganizationsOrganizationProjectsProjectRulesHeaderPatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesHeaderPatch) | **Patch** /organizations/{organization}/projects/{project}/rules/header | Create header rules
[**OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/header/{rule} | Delete a header rule
[**OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/header/{rule} | Get header rule details
[**OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/header/{rule} | Update header rule details
[**OrganizationsOrganizationProjectsProjectRulesProxyPost**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesProxyPost) | **Post** /organizations/{organization}/projects/{project}/rules/proxy | Update proxy rule details
[**OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | Delete a proxy rule
[**OrganizationsOrganizationProjectsProjectRulesProxyRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesProxyRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | Get proxy rule details
[**OrganizationsOrganizationProjectsProjectRulesProxyRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesProxyRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | Update proxy rule details
[**OrganizationsOrganizationProjectsProjectRulesRedirectPost**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesRedirectPost) | **Post** /organizations/{organization}/projects/{project}/rules/redirect | Create a new redirect rule
[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | Delete a redirect rule
[**OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet) | **Get** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | Get redirect rule details
[**OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch**](RulesAPI.md#OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch) | **Patch** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | Update redirect rule details



## OrganizationsOrganizationProjectsProjectRulesAuthPost

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesAuthPost(ctx, organization, project, rule).Body(body).Execute()

Create authentication rules

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    body := RuleAuthRequest(987) // RuleAuthRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthPost(context.Background(), organization, project, rule).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesAuthPost`: OrganizationsOrganizationProjectsProjectRulesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthPost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesAuthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **RuleAuthRequest** |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete(ctx, organization, project, rule).Execute()

Delete an authentication rule

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
    // response from `OrganizationsOrganizationProjectsProjectRulesAuthRuleDelete`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesAuthRuleGet

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesAuthRuleGet(ctx, organization, project, rule).Execute()

Get authentication rule details

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
    // response from `OrganizationsOrganizationProjectsProjectRulesAuthRuleGet`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesAuthRulePatch

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesAuthRulePatch(ctx, organization, project, rule).Body(body).Execute()

Update authentication rule details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    body := RuleAuthRequest(987) // RuleAuthRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRulePatch(context.Background(), organization, project, rule).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesAuthRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesAuthRulePatch`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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



 **body** | **RuleAuthRequest** |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesCustomResponsePatch

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesCustomResponsePatch(ctx, organization, project, rule).Body(body).Execute()

Update custom response rule details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    body := RuleCustomResponseRequest(987) // RuleCustomResponseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponsePatch(context.Background(), organization, project, rule).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponsePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesCustomResponsePatch`: OrganizationsOrganizationProjectsProjectRulesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponsePatch`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesCustomResponsePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **RuleCustomResponseRequest** |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete(ctx, organization, project, rule).Execute()

Delete a custom response rule

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
    // response from `OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleDelete`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet(ctx, organization, project, rule).Execute()

Get custom repsonse rule details

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
    // response from `OrganizationsOrganizationProjectsProjectRulesCustomResponseRuleGet`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch(ctx, organization, project, rule).Body(body).Execute()

Update custom response rule details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    body := RuleCustomResponseRequest(987) // RuleCustomResponseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch(context.Background(), organization, project, rule).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesCustomResponseRulePatch`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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



 **body** | **RuleCustomResponseRequest** |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesHeaderPatch

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesHeaderPatch(ctx, organization, project, rule).Body(body).Execute()

Create header rules

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    body := RuleHeaderRequest(987) // RuleHeaderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderPatch(context.Background(), organization, project, rule).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesHeaderPatch`: OrganizationsOrganizationProjectsProjectRulesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderPatch`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesHeaderPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **RuleHeaderRequest** |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete(ctx, organization, project, rule).Execute()

Delete a header rule

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
    // response from `OrganizationsOrganizationProjectsProjectRulesHeaderRuleDelete`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet(ctx, organization, project, rule).Execute()

Get header rule details

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
    // response from `OrganizationsOrganizationProjectsProjectRulesHeaderRuleGet`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch(ctx, organization, project, rule).Body(body).Execute()

Update header rule details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    body := RuleHeaderRequest(987) // RuleHeaderRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch(context.Background(), organization, project, rule).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesHeaderRulePatch`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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



 **body** | **RuleHeaderRequest** |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesProxyPost

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesProxyPost(ctx, organization, project, rule).Body(body).Execute()

Update proxy rule details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    body := RuleProxyRequest(987) // RuleProxyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyPost(context.Background(), organization, project, rule).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesProxyPost`: OrganizationsOrganizationProjectsProjectRulesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyPost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesProxyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **RuleProxyRequest** |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete(ctx, organization, project, rule).Execute()

Delete a proxy rule

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
    // response from `OrganizationsOrganizationProjectsProjectRulesProxyRuleDelete`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesProxyRuleGet

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesProxyRuleGet(ctx, organization, project, rule).Execute()

Get proxy rule details

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
    // response from `OrganizationsOrganizationProjectsProjectRulesProxyRuleGet`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesProxyRulePatch

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesProxyRulePatch(ctx, organization, project, rule).Body(body).Execute()

Update proxy rule details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    body := RuleProxyRequest(987) // RuleProxyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRulePatch(context.Background(), organization, project, rule).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesProxyRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesProxyRulePatch`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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



 **body** | **RuleProxyRequest** |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesRedirectPost

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesRedirectPost(ctx, organization, project, rule).RuleRedirectRequest(ruleRedirectRequest).Execute()

Create a new redirect rule

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    ruleRedirectRequest := *openapiclient.NewRuleRedirectRequest() // RuleRedirectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectPost(context.Background(), organization, project, rule).RuleRedirectRequest(ruleRedirectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesRedirectPost`: OrganizationsOrganizationProjectsProjectRulesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectPost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrganizationsOrganizationProjectsProjectRulesRedirectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleRedirectRequest** | [**RuleRedirectRequest**](RuleRedirectRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete(ctx, organization, project, rule).Execute()

Delete a redirect rule

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
    // response from `OrganizationsOrganizationProjectsProjectRulesRedirectRuleDelete`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet(ctx, organization, project, rule).Execute()

Get redirect rule details

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
    // response from `OrganizationsOrganizationProjectsProjectRulesRedirectRuleGet`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch

> OrganizationsOrganizationProjectsProjectRulesGet200Response OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch(ctx, organization, project, rule).RuleRedirectRequest(ruleRedirectRequest).Execute()

Update redirect rule details

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
    organization := TODO // interface{} | Organization machine name
    project := TODO // interface{} | Project machine name
    rule := TODO // interface{} |
    ruleRedirectRequest := *openapiclient.NewRuleRedirectRequest() // RuleRedirectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch(context.Background(), organization, project, rule).RuleRedirectRequest(ruleRedirectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationsOrganizationProjectsProjectRulesRedirectRulePatch`: OrganizationsOrganizationProjectsProjectRulesGet200Response
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



 **ruleRedirectRequest** | [**RuleRedirectRequest**](RuleRedirectRequest.md) |  |

### Return type

[**OrganizationsOrganizationProjectsProjectRulesGet200Response**](OrganizationsOrganizationProjectsProjectRulesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

