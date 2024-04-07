# \RulesAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRuleAuth**](RulesAPI.md#CreateRuleAuth) | **Post** /organizations/{organization}/projects/{project}/rules/auth | 
[**CreateRuleCustomResponse**](RulesAPI.md#CreateRuleCustomResponse) | **Post** /organizations/{organization}/projects/{project}/rules/custom-response | 
[**CreateRuleHeaders**](RulesAPI.md#CreateRuleHeaders) | **Post** /organizations/{organization}/projects/{project}/rules/headers | 
[**CreateRuleProxy**](RulesAPI.md#CreateRuleProxy) | **Post** /organizations/{organization}/projects/{project}/rules/proxy | 
[**CreateRuleRedirect**](RulesAPI.md#CreateRuleRedirect) | **Post** /organizations/{organization}/projects/{project}/rules/redirect | 
[**DeleteRuleAuth**](RulesAPI.md#DeleteRuleAuth) | **Delete** /organizations/{organization}/projects/{project}/rules/auth/{rule} | 
[**DeleteRuleCustomResponse**](RulesAPI.md#DeleteRuleCustomResponse) | **Delete** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | 
[**DeleteRuleHeaders**](RulesAPI.md#DeleteRuleHeaders) | **Delete** /organizations/{organization}/projects/{project}/rules/headers/{rule} | 
[**DeleteRuleProxy**](RulesAPI.md#DeleteRuleProxy) | **Delete** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | 
[**DeleteRuleRedirect**](RulesAPI.md#DeleteRuleRedirect) | **Delete** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | 
[**GetRuleAuth**](RulesAPI.md#GetRuleAuth) | **Get** /organizations/{organization}/projects/{project}/rules/auth/{rule} | 
[**GetRuleCustomResponse**](RulesAPI.md#GetRuleCustomResponse) | **Get** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | 
[**GetRuleHeaders**](RulesAPI.md#GetRuleHeaders) | **Get** /organizations/{organization}/projects/{project}/rules/headers/{rule} | 
[**GetRuleProxy**](RulesAPI.md#GetRuleProxy) | **Get** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | 
[**GetRuleRedirect**](RulesAPI.md#GetRuleRedirect) | **Get** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | 
[**ListRules**](RulesAPI.md#ListRules) | **Get** /organizations/{organization}/projects/{project}/rules | 
[**UpdateRuleAuth**](RulesAPI.md#UpdateRuleAuth) | **Patch** /organizations/{organization}/projects/{project}/rules/auth/{rule} | 
[**UpdateRuleCustomResponse**](RulesAPI.md#UpdateRuleCustomResponse) | **Patch** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | 
[**UpdateRuleHeaders**](RulesAPI.md#UpdateRuleHeaders) | **Patch** /organizations/{organization}/projects/{project}/rules/headers/{rule} | 
[**UpdateRuleProxy**](RulesAPI.md#UpdateRuleProxy) | **Patch** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | 
[**UpdateRuleRedirect**](RulesAPI.md#UpdateRuleRedirect) | **Patch** /organizations/{organization}/projects/{project}/rules/redirect/{rule} | 



## CreateRuleAuth

> Rule CreateRuleAuth(ctx, organization, project).RuleAuthRequest(ruleAuthRequest).Execute()





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
	ruleAuthRequest := *openapiclient.NewRuleAuthRequest() // RuleAuthRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.CreateRuleAuth(context.Background(), organization, project).RuleAuthRequest(ruleAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateRuleAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleAuth`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateRuleAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleAuthRequest** | [**RuleAuthRequest**](RuleAuthRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleCustomResponse

> Rule CreateRuleCustomResponse(ctx, organization, project).RuleCustomResponseRequest(ruleCustomResponseRequest).Execute()





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
	ruleCustomResponseRequest := *openapiclient.NewRuleCustomResponseRequest() // RuleCustomResponseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.CreateRuleCustomResponse(context.Background(), organization, project).RuleCustomResponseRequest(ruleCustomResponseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateRuleCustomResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleCustomResponse`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateRuleCustomResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleCustomResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleCustomResponseRequest** | [**RuleCustomResponseRequest**](RuleCustomResponseRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleHeaders

> Rule CreateRuleHeaders(ctx, organization, project).RuleHeaderRequest(ruleHeaderRequest).Execute()





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
	ruleHeaderRequest := *openapiclient.NewRuleHeaderRequest() // RuleHeaderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.CreateRuleHeaders(context.Background(), organization, project).RuleHeaderRequest(ruleHeaderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateRuleHeaders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleHeaders`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateRuleHeaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleHeaderRequest** | [**RuleHeaderRequest**](RuleHeaderRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleProxy

> Rule CreateRuleProxy(ctx, organization, project).RuleProxyRequest(ruleProxyRequest).Execute()





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
	ruleProxyRequest := *openapiclient.NewRuleProxyRequest() // RuleProxyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.CreateRuleProxy(context.Background(), organization, project).RuleProxyRequest(ruleProxyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateRuleProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleProxy`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateRuleProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleProxyRequest** | [**RuleProxyRequest**](RuleProxyRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleRedirect

> Rule CreateRuleRedirect(ctx, organization, project).RuleRedirectRequest(ruleRedirectRequest).Execute()





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
	ruleRedirectRequest := *openapiclient.NewRuleRedirectRequest() // RuleRedirectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.CreateRuleRedirect(context.Background(), organization, project).RuleRedirectRequest(ruleRedirectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.CreateRuleRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRuleRedirect`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.CreateRuleRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleRedirectRequest** | [**RuleRedirectRequest**](RuleRedirectRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleAuth

> Rule DeleteRuleAuth(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.DeleteRuleAuth(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRuleAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleAuth`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.DeleteRuleAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleCustomResponse

> Rule DeleteRuleCustomResponse(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.DeleteRuleCustomResponse(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRuleCustomResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleCustomResponse`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.DeleteRuleCustomResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleCustomResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleHeaders

> Rule DeleteRuleHeaders(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.DeleteRuleHeaders(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRuleHeaders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleHeaders`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.DeleteRuleHeaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleProxy

> Rule DeleteRuleProxy(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.DeleteRuleProxy(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRuleProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleProxy`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.DeleteRuleProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleRedirect

> Rule DeleteRuleRedirect(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.DeleteRuleRedirect(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.DeleteRuleRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRuleRedirect`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.DeleteRuleRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleAuth

> Rule GetRuleAuth(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleAuth(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleAuth`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleCustomResponse

> Rule GetRuleCustomResponse(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleCustomResponse(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleCustomResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleCustomResponse`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleCustomResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleCustomResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleHeaders

> Rule GetRuleHeaders(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleHeaders(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleHeaders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleHeaders`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleHeaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleProxy

> Rule GetRuleProxy(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleProxy(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleProxy`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleRedirect

> Rule GetRuleRedirect(ctx, organization, project, rule).Execute()





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
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.GetRuleRedirect(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.GetRuleRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleRedirect`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.GetRuleRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> []Rule ListRules(ctx, organization, project).Execute()





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
	resp, r, err := apiClient.RulesAPI.ListRules(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ListRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRules`: []Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ListRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleAuth

> Rule UpdateRuleAuth(ctx, organization, project, rule).RuleAuthRequest(ruleAuthRequest).Execute()





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
	rule := "rule_example" // string | 
	ruleAuthRequest := *openapiclient.NewRuleAuthRequest() // RuleAuthRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateRuleAuth(context.Background(), organization, project, rule).RuleAuthRequest(ruleAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateRuleAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleAuth`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateRuleAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleAuthRequest** | [**RuleAuthRequest**](RuleAuthRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleCustomResponse

> Rule UpdateRuleCustomResponse(ctx, organization, project, rule).RuleCustomResponseRequest(ruleCustomResponseRequest).Execute()





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
	rule := "rule_example" // string | 
	ruleCustomResponseRequest := *openapiclient.NewRuleCustomResponseRequest() // RuleCustomResponseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateRuleCustomResponse(context.Background(), organization, project, rule).RuleCustomResponseRequest(ruleCustomResponseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateRuleCustomResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleCustomResponse`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateRuleCustomResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleCustomResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleCustomResponseRequest** | [**RuleCustomResponseRequest**](RuleCustomResponseRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleHeaders

> Rule UpdateRuleHeaders(ctx, organization, project, rule).RuleHeaderRequest(ruleHeaderRequest).Execute()





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
	rule := "rule_example" // string | 
	ruleHeaderRequest := *openapiclient.NewRuleHeaderRequest() // RuleHeaderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateRuleHeaders(context.Background(), organization, project, rule).RuleHeaderRequest(ruleHeaderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateRuleHeaders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleHeaders`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateRuleHeaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleHeaderRequest** | [**RuleHeaderRequest**](RuleHeaderRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleProxy

> Rule UpdateRuleProxy(ctx, organization, project, rule).RuleProxyRequest(ruleProxyRequest).Execute()





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
	rule := "rule_example" // string | 
	ruleProxyRequest := *openapiclient.NewRuleProxyRequest() // RuleProxyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateRuleProxy(context.Background(), organization, project, rule).RuleProxyRequest(ruleProxyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateRuleProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleProxy`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateRuleProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleProxyRequest** | [**RuleProxyRequest**](RuleProxyRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleRedirect

> Rule UpdateRuleRedirect(ctx, organization, project, rule).RuleRedirectRequest(ruleRedirectRequest).Execute()





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
	rule := "rule_example" // string | 
	ruleRedirectRequest := *openapiclient.NewRuleRedirectRequest() // RuleRedirectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.UpdateRuleRedirect(context.Background(), organization, project, rule).RuleRedirectRequest(ruleRedirectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.UpdateRuleRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRuleRedirect`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.UpdateRuleRedirect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleRedirectRequest** | [**RuleRedirectRequest**](RuleRedirectRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

