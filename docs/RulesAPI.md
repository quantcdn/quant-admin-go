# \RulesAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesAuthCreate**](RulesAPI.md#RulesAuthCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/auth | Create an authentication rule
[**RulesAuthDelete**](RulesAPI.md#RulesAuthDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/auth/{rule} | Delete an authentication rule
[**RulesAuthList**](RulesAPI.md#RulesAuthList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/auth | List authentication rules
[**RulesAuthRead**](RulesAPI.md#RulesAuthRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/auth/{rule} | Get details of an authentication rule
[**RulesAuthUpdate**](RulesAPI.md#RulesAuthUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/auth/{rule} | Update an authentication rule
[**RulesBotChallengeCreate**](RulesAPI.md#RulesBotChallengeCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/bot-challenge | Create a bot challenge rule
[**RulesBotChallengeDelete**](RulesAPI.md#RulesBotChallengeDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/bot-challenge/{rule} | Delete a bot challenge rule
[**RulesBotChallengeList**](RulesAPI.md#RulesBotChallengeList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/bot-challenge | List bot challenge rules
[**RulesBotChallengeRead**](RulesAPI.md#RulesBotChallengeRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/bot-challenge/{rule} | Get details of a bot challenge rule
[**RulesBotChallengeUpdate**](RulesAPI.md#RulesBotChallengeUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/bot-challenge/{rule} | Update a bot challenge rule
[**RulesContentFilterCreate**](RulesAPI.md#RulesContentFilterCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/content-filter | Create a content filter rule
[**RulesContentFilterDelete**](RulesAPI.md#RulesContentFilterDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/content-filter/{rule} | Delete a content filter rule
[**RulesContentFilterList**](RulesAPI.md#RulesContentFilterList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/content-filter | List content filter rules
[**RulesContentFilterRead**](RulesAPI.md#RulesContentFilterRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/content-filter/{rule} | Get details of a content filter rule
[**RulesContentFilterUpdate**](RulesAPI.md#RulesContentFilterUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/content-filter/{rule} | Update a content filter rule
[**RulesCustomResponseCreate**](RulesAPI.md#RulesCustomResponseCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/custom-response | Create a custom response rule
[**RulesCustomResponseDelete**](RulesAPI.md#RulesCustomResponseDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Delete a custom response rule
[**RulesCustomResponseList**](RulesAPI.md#RulesCustomResponseList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/custom-response | List custom response rules
[**RulesCustomResponseRead**](RulesAPI.md#RulesCustomResponseRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Get details of a custom response rule
[**RulesCustomResponseUpdate**](RulesAPI.md#RulesCustomResponseUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/custom-response/{rule} | Update a custom response rule
[**RulesFunctionCreate**](RulesAPI.md#RulesFunctionCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/function | Create an edge function rule
[**RulesFunctionDelete**](RulesAPI.md#RulesFunctionDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/function/{rule} | Delete an edge function rule
[**RulesFunctionList**](RulesAPI.md#RulesFunctionList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/function | List edge function rules
[**RulesFunctionRead**](RulesAPI.md#RulesFunctionRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/function/{rule} | Get details of an edge function rule
[**RulesFunctionUpdate**](RulesAPI.md#RulesFunctionUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/function/{rule} | Update an edge function rule
[**RulesHeadersCreate**](RulesAPI.md#RulesHeadersCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/headers | Create a header rule
[**RulesHeadersDelete**](RulesAPI.md#RulesHeadersDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/headers/{rule} | Delete a header rule
[**RulesHeadersList**](RulesAPI.md#RulesHeadersList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/headers | List header rules
[**RulesHeadersRead**](RulesAPI.md#RulesHeadersRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/headers/{rule} | Get details of a header rule
[**RulesHeadersUpdate**](RulesAPI.md#RulesHeadersUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/headers/{rule} | Update a header rule
[**RulesProxyCreate**](RulesAPI.md#RulesProxyCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/proxy | Create a proxy rule
[**RulesProxyDelete**](RulesAPI.md#RulesProxyDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/proxy/{rule} | Delete a proxy rule
[**RulesProxyList**](RulesAPI.md#RulesProxyList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/proxy | List proxy rules
[**RulesProxyRead**](RulesAPI.md#RulesProxyRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/proxy/{rule} | Get details of a proxy rule
[**RulesProxyUpdate**](RulesAPI.md#RulesProxyUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/proxy/{rule} | Update a proxy rule
[**RulesRedirectCreate**](RulesAPI.md#RulesRedirectCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/redirect | Create a redirect rule
[**RulesRedirectDelete**](RulesAPI.md#RulesRedirectDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/redirect/{rule} | Delete a redirect rule
[**RulesRedirectList**](RulesAPI.md#RulesRedirectList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/redirect | List redirect rules
[**RulesRedirectRead**](RulesAPI.md#RulesRedirectRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/redirect/{rule} | Get details of a redirect rule
[**RulesRedirectUpdate**](RulesAPI.md#RulesRedirectUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/redirect/{rule} | Update a redirect rule
[**RulesServeStaticCreate**](RulesAPI.md#RulesServeStaticCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/rules/serve-static | Create a serve static rule
[**RulesServeStaticDelete**](RulesAPI.md#RulesServeStaticDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/rules/serve-static/{rule} | Delete a serve static rule
[**RulesServeStaticList**](RulesAPI.md#RulesServeStaticList) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/serve-static | List serve static rules
[**RulesServeStaticRead**](RulesAPI.md#RulesServeStaticRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/rules/serve-static/{rule} | Get details of a serve static rule
[**RulesServeStaticUpdate**](RulesAPI.md#RulesServeStaticUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/rules/serve-static/{rule} | Update a serve static rule



## RulesAuthCreate

> V2RuleAuth RulesAuthCreate(ctx, organization, project).V2RuleAuthRequest(v2RuleAuthRequest).Execute()

Create an authentication rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	v2RuleAuthRequest := *openapiclient.NewV2RuleAuthRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", "admin", "secure_password123") // V2RuleAuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesAuthCreate(context.Background(), organization, project).V2RuleAuthRequest(v2RuleAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesAuthCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesAuthCreate`: V2RuleAuth
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesAuthCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesAuthCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleAuthRequest** | [**V2RuleAuthRequest**](V2RuleAuthRequest.md) |  | 

### Return type

[**V2RuleAuth**](V2RuleAuth.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesAuthDelete

> RulesAuthDelete(ctx, organization, project, rule).Execute()

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesAuthDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesAuthDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesAuthDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesAuthList

> []V2RuleAuth RulesAuthList(ctx, organization, project).Execute()

List authentication rules

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesAuthList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesAuthList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesAuthList`: []V2RuleAuth
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesAuthList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesAuthListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleAuth**](V2RuleAuth.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesAuthRead

> V2RuleAuth RulesAuthRead(ctx, organization, project, rule).Execute()

Get details of an authentication rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesAuthRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesAuthRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesAuthRead`: V2RuleAuth
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesAuthRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesAuthReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleAuth**](V2RuleAuth.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesAuthUpdate

> V2RuleAuth RulesAuthUpdate(ctx, organization, project, rule).V2RuleAuthRequest(v2RuleAuthRequest).Execute()

Update an authentication rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier
	v2RuleAuthRequest := *openapiclient.NewV2RuleAuthRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", "admin", "secure_password123") // V2RuleAuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesAuthUpdate(context.Background(), organization, project, rule).V2RuleAuthRequest(v2RuleAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesAuthUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesAuthUpdate`: V2RuleAuth
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesAuthUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesAuthUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleAuthRequest** | [**V2RuleAuthRequest**](V2RuleAuthRequest.md) |  | 

### Return type

[**V2RuleAuth**](V2RuleAuth.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeCreate

> V2RuleBotChallenge RulesBotChallengeCreate(ctx, organization, project).V2RuleBotChallengeRequest(v2RuleBotChallengeRequest).Execute()

Create a bot challenge rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	v2RuleBotChallengeRequest := *openapiclient.NewV2RuleBotChallengeRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", "invisible") // V2RuleBotChallengeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesBotChallengeCreate(context.Background(), organization, project).V2RuleBotChallengeRequest(v2RuleBotChallengeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesBotChallengeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeCreate`: V2RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesBotChallengeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleBotChallengeRequest** | [**V2RuleBotChallengeRequest**](V2RuleBotChallengeRequest.md) |  | 

### Return type

[**V2RuleBotChallenge**](V2RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeDelete

> RulesBotChallengeDelete(ctx, organization, project, rule).Execute()

Delete a bot challenge rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesBotChallengeDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesBotChallengeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeList

> []V2RuleBotChallenge RulesBotChallengeList(ctx, organization, project).Execute()

List bot challenge rules

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesBotChallengeList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesBotChallengeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeList`: []V2RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesBotChallengeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleBotChallenge**](V2RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeRead

> V2RuleBotChallenge RulesBotChallengeRead(ctx, organization, project, rule).Execute()

Get details of a bot challenge rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesBotChallengeRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesBotChallengeRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeRead`: V2RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesBotChallengeRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleBotChallenge**](V2RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeUpdate

> V2RuleBotChallenge RulesBotChallengeUpdate(ctx, organization, project, rule).V2RuleBotChallengeRequest(v2RuleBotChallengeRequest).Execute()

Update a bot challenge rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier
	v2RuleBotChallengeRequest := *openapiclient.NewV2RuleBotChallengeRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", "invisible") // V2RuleBotChallengeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesBotChallengeUpdate(context.Background(), organization, project, rule).V2RuleBotChallengeRequest(v2RuleBotChallengeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesBotChallengeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeUpdate`: V2RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesBotChallengeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleBotChallengeRequest** | [**V2RuleBotChallengeRequest**](V2RuleBotChallengeRequest.md) |  | 

### Return type

[**V2RuleBotChallenge**](V2RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterCreate

> V2RuleContentFilter RulesContentFilterCreate(ctx, organization, project).V2RuleContentFilterRequest(v2RuleContentFilterRequest).Execute()

Create a content filter rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	v2RuleContentFilterRequest := *openapiclient.NewV2RuleContentFilterRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "8d3f4820-8536-51ef-b827-f18gc2g01bf8") // V2RuleContentFilterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesContentFilterCreate(context.Background(), organization, project).V2RuleContentFilterRequest(v2RuleContentFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesContentFilterCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterCreate`: V2RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesContentFilterCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesContentFilterCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleContentFilterRequest** | [**V2RuleContentFilterRequest**](V2RuleContentFilterRequest.md) |  | 

### Return type

[**V2RuleContentFilter**](V2RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterDelete

> RulesContentFilterDelete(ctx, organization, project, rule).Execute()

Delete a content filter rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesContentFilterDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesContentFilterDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesContentFilterDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterList

> []V2RuleContentFilter RulesContentFilterList(ctx, organization, project).Execute()

List content filter rules

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesContentFilterList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesContentFilterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterList`: []V2RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesContentFilterList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesContentFilterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleContentFilter**](V2RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterRead

> V2RuleContentFilter RulesContentFilterRead(ctx, organization, project, rule).Execute()

Get details of a content filter rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesContentFilterRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesContentFilterRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterRead`: V2RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesContentFilterRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesContentFilterReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleContentFilter**](V2RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterUpdate

> V2RuleContentFilter RulesContentFilterUpdate(ctx, organization, project, rule).V2RuleContentFilterRequest(v2RuleContentFilterRequest).Execute()

Update a content filter rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier
	v2RuleContentFilterRequest := *openapiclient.NewV2RuleContentFilterRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "8d3f4820-8536-51ef-b827-f18gc2g01bf8") // V2RuleContentFilterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesContentFilterUpdate(context.Background(), organization, project, rule).V2RuleContentFilterRequest(v2RuleContentFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesContentFilterUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterUpdate`: V2RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesContentFilterUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesContentFilterUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleContentFilterRequest** | [**V2RuleContentFilterRequest**](V2RuleContentFilterRequest.md) |  | 

### Return type

[**V2RuleContentFilter**](V2RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseCreate

> V2RuleCustomResponse RulesCustomResponseCreate(ctx, organization, project).V2RuleCustomResponseRequest(v2RuleCustomResponseRequest).Execute()

Create a custom response rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	v2RuleCustomResponseRequest := *openapiclient.NewV2RuleCustomResponseRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "<html><body>Custom maintenance page</body></html>") // V2RuleCustomResponseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesCustomResponseCreate(context.Background(), organization, project).V2RuleCustomResponseRequest(v2RuleCustomResponseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesCustomResponseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseCreate`: V2RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesCustomResponseCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesCustomResponseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleCustomResponseRequest** | [**V2RuleCustomResponseRequest**](V2RuleCustomResponseRequest.md) |  | 

### Return type

[**V2RuleCustomResponse**](V2RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseDelete

> RulesCustomResponseDelete(ctx, organization, project, rule).Execute()

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesCustomResponseDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesCustomResponseDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesCustomResponseDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseList

> []V2RuleCustomResponse RulesCustomResponseList(ctx, organization, project).Execute()

List custom response rules

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesCustomResponseList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesCustomResponseList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseList`: []V2RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesCustomResponseList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesCustomResponseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleCustomResponse**](V2RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseRead

> V2RuleCustomResponse RulesCustomResponseRead(ctx, organization, project, rule).Execute()

Get details of a custom response rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesCustomResponseRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesCustomResponseRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseRead`: V2RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesCustomResponseRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesCustomResponseReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleCustomResponse**](V2RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseUpdate

> V2RuleCustomResponse RulesCustomResponseUpdate(ctx, organization, project, rule).V2RuleCustomResponseRequest(v2RuleCustomResponseRequest).Execute()

Update a custom response rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier
	v2RuleCustomResponseRequest := *openapiclient.NewV2RuleCustomResponseRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "<html><body>Custom maintenance page</body></html>") // V2RuleCustomResponseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesCustomResponseUpdate(context.Background(), organization, project, rule).V2RuleCustomResponseRequest(v2RuleCustomResponseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesCustomResponseUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseUpdate`: V2RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesCustomResponseUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesCustomResponseUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleCustomResponseRequest** | [**V2RuleCustomResponseRequest**](V2RuleCustomResponseRequest.md) |  | 

### Return type

[**V2RuleCustomResponse**](V2RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionCreate

> V2RuleFunction RulesFunctionCreate(ctx, organization, project).V2RuleFunctionRequest(v2RuleFunctionRequest).Execute()

Create an edge function rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	v2RuleFunctionRequest := *openapiclient.NewV2RuleFunctionRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", "7c9e6679-7425-40de-944b-e07fc1f90ae7") // V2RuleFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesFunctionCreate(context.Background(), organization, project).V2RuleFunctionRequest(v2RuleFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesFunctionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionCreate`: V2RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesFunctionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesFunctionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleFunctionRequest** | [**V2RuleFunctionRequest**](V2RuleFunctionRequest.md) |  | 

### Return type

[**V2RuleFunction**](V2RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionDelete

> RulesFunctionDelete(ctx, organization, project, rule).Execute()

Delete an edge function rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesFunctionDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesFunctionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesFunctionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionList

> []V2RuleFunction RulesFunctionList(ctx, organization, project).Execute()

List edge function rules

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesFunctionList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesFunctionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionList`: []V2RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesFunctionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesFunctionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleFunction**](V2RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionRead

> V2RuleFunction RulesFunctionRead(ctx, organization, project, rule).Execute()

Get details of an edge function rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesFunctionRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesFunctionRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionRead`: V2RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesFunctionRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesFunctionReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleFunction**](V2RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionUpdate

> V2RuleFunction RulesFunctionUpdate(ctx, organization, project, rule).V2RuleFunctionRequest(v2RuleFunctionRequest).Execute()

Update an edge function rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier
	v2RuleFunctionRequest := *openapiclient.NewV2RuleFunctionRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", "7c9e6679-7425-40de-944b-e07fc1f90ae7") // V2RuleFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesFunctionUpdate(context.Background(), organization, project, rule).V2RuleFunctionRequest(v2RuleFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesFunctionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionUpdate`: V2RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesFunctionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesFunctionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleFunctionRequest** | [**V2RuleFunctionRequest**](V2RuleFunctionRequest.md) |  | 

### Return type

[**V2RuleFunction**](V2RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesHeadersCreate

> V2RuleHeader RulesHeadersCreate(ctx, organization, project).V2RuleHeaderRequest(v2RuleHeaderRequest).Execute()

Create a header rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	v2RuleHeaderRequest := *openapiclient.NewV2RuleHeaderRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", map[string]string{"key": "Inner_example"}) // V2RuleHeaderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesHeadersCreate(context.Background(), organization, project).V2RuleHeaderRequest(v2RuleHeaderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesHeadersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesHeadersCreate`: V2RuleHeader
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesHeadersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesHeadersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleHeaderRequest** | [**V2RuleHeaderRequest**](V2RuleHeaderRequest.md) |  | 

### Return type

[**V2RuleHeader**](V2RuleHeader.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesHeadersDelete

> RulesHeadersDelete(ctx, organization, project, rule).Execute()

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesHeadersDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesHeadersDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesHeadersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesHeadersList

> []V2RuleHeader RulesHeadersList(ctx, organization, project).Execute()

List header rules

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesHeadersList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesHeadersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesHeadersList`: []V2RuleHeader
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesHeadersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesHeadersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleHeader**](V2RuleHeader.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesHeadersRead

> V2RuleHeader RulesHeadersRead(ctx, organization, project, rule).Execute()

Get details of a header rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesHeadersRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesHeadersRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesHeadersRead`: V2RuleHeader
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesHeadersRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesHeadersReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleHeader**](V2RuleHeader.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesHeadersUpdate

> V2RuleHeader RulesHeadersUpdate(ctx, organization, project, rule).V2RuleHeaderRequest(v2RuleHeaderRequest).Execute()

Update a header rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier
	v2RuleHeaderRequest := *openapiclient.NewV2RuleHeaderRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", map[string]string{"key": "Inner_example"}) // V2RuleHeaderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesHeadersUpdate(context.Background(), organization, project, rule).V2RuleHeaderRequest(v2RuleHeaderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesHeadersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesHeadersUpdate`: V2RuleHeader
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesHeadersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesHeadersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleHeaderRequest** | [**V2RuleHeaderRequest**](V2RuleHeaderRequest.md) |  | 

### Return type

[**V2RuleHeader**](V2RuleHeader.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesProxyCreate

> V2RuleProxy RulesProxyCreate(ctx, organization, project).V2RuleProxyRequest(v2RuleProxyRequest).Execute()

Create a proxy rule

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	v2RuleProxyRequest := *openapiclient.NewV2RuleProxyRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com") // V2RuleProxyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesProxyCreate(context.Background(), organization, project).V2RuleProxyRequest(v2RuleProxyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesProxyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesProxyCreate`: V2RuleProxy
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesProxyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesProxyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleProxyRequest** | [**V2RuleProxyRequest**](V2RuleProxyRequest.md) |  | 

### Return type

[**V2RuleProxy**](V2RuleProxy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesProxyDelete

> RulesProxyDelete(ctx, organization, project, rule).Execute()

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesProxyDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesProxyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesProxyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesProxyList

> []V2RuleProxy RulesProxyList(ctx, organization, project).Execute()

List proxy rules

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesProxyList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesProxyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesProxyList`: []V2RuleProxy
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesProxyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesProxyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleProxy**](V2RuleProxy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesProxyRead

> V2RuleProxy RulesProxyRead(ctx, organization, project, rule).Execute()

Get details of a proxy rule

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesProxyRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesProxyRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesProxyRead`: V2RuleProxy
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesProxyRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesProxyReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleProxy**](V2RuleProxy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesProxyUpdate

> V2RuleProxy RulesProxyUpdate(ctx, organization, project, rule).V2RuleProxyRequest(v2RuleProxyRequest).Execute()

Update a proxy rule

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 
	v2RuleProxyRequest := *openapiclient.NewV2RuleProxyRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com") // V2RuleProxyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesProxyUpdate(context.Background(), organization, project, rule).V2RuleProxyRequest(v2RuleProxyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesProxyUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesProxyUpdate`: V2RuleProxy
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesProxyUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesProxyUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleProxyRequest** | [**V2RuleProxyRequest**](V2RuleProxyRequest.md) |  | 

### Return type

[**V2RuleProxy**](V2RuleProxy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesRedirectCreate

> V2RuleRedirect RulesRedirectCreate(ctx, organization, project).V2RuleRedirectRequest(v2RuleRedirectRequest).Execute()

Create a redirect rule

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	v2RuleRedirectRequest := *openapiclient.NewV2RuleRedirectRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://test-domain.com/new-path") // V2RuleRedirectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesRedirectCreate(context.Background(), organization, project).V2RuleRedirectRequest(v2RuleRedirectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesRedirectCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesRedirectCreate`: V2RuleRedirect
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesRedirectCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesRedirectCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleRedirectRequest** | [**V2RuleRedirectRequest**](V2RuleRedirectRequest.md) |  | 

### Return type

[**V2RuleRedirect**](V2RuleRedirect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesRedirectDelete

> RulesRedirectDelete(ctx, organization, project, rule).Execute()

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesRedirectDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesRedirectDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesRedirectDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesRedirectList

> []V2RuleRedirect RulesRedirectList(ctx, organization, project).Execute()

List redirect rules

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesRedirectList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesRedirectList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesRedirectList`: []V2RuleRedirect
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesRedirectList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesRedirectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleRedirect**](V2RuleRedirect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesRedirectRead

> V2RuleRedirect RulesRedirectRead(ctx, organization, project, rule).Execute()

Get details of a redirect rule

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesRedirectRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesRedirectRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesRedirectRead`: V2RuleRedirect
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesRedirectRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesRedirectReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleRedirect**](V2RuleRedirect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesRedirectUpdate

> V2RuleRedirect RulesRedirectUpdate(ctx, organization, project, rule).V2RuleRedirectRequest(v2RuleRedirectRequest).Execute()

Update a redirect rule

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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 
	v2RuleRedirectRequest := *openapiclient.NewV2RuleRedirectRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://test-domain.com/new-path") // V2RuleRedirectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesRedirectUpdate(context.Background(), organization, project, rule).V2RuleRedirectRequest(v2RuleRedirectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesRedirectUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesRedirectUpdate`: V2RuleRedirect
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesRedirectUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesRedirectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleRedirectRequest** | [**V2RuleRedirectRequest**](V2RuleRedirectRequest.md) |  | 

### Return type

[**V2RuleRedirect**](V2RuleRedirect.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticCreate

> V2RuleServeStatic RulesServeStaticCreate(ctx, organization, project).V2RuleServeStaticRequest(v2RuleServeStaticRequest).Execute()

Create a serve static rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	v2RuleServeStaticRequest := *openapiclient.NewV2RuleServeStaticRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", "/index.html") // V2RuleServeStaticRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesServeStaticCreate(context.Background(), organization, project).V2RuleServeStaticRequest(v2RuleServeStaticRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesServeStaticCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticCreate`: V2RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesServeStaticCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesServeStaticCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2RuleServeStaticRequest** | [**V2RuleServeStaticRequest**](V2RuleServeStaticRequest.md) |  | 

### Return type

[**V2RuleServeStatic**](V2RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticDelete

> RulesServeStaticDelete(ctx, organization, project, rule).Execute()

Delete a serve static rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RulesAPI.RulesServeStaticDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesServeStaticDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesServeStaticDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticList

> []V2RuleServeStatic RulesServeStaticList(ctx, organization, project).Execute()

List serve static rules

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesServeStaticList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesServeStaticList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticList`: []V2RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesServeStaticList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesServeStaticListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2RuleServeStatic**](V2RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticRead

> V2RuleServeStatic RulesServeStaticRead(ctx, organization, project, rule).Execute()

Get details of a serve static rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesServeStaticRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesServeStaticRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticRead`: V2RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesServeStaticRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesServeStaticReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2RuleServeStatic**](V2RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticUpdate

> V2RuleServeStatic RulesServeStaticUpdate(ctx, organization, project, rule).V2RuleServeStaticRequest(v2RuleServeStaticRequest).Execute()

Update a serve static rule

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	rule := "rule_example" // string | Rule identifier
	v2RuleServeStaticRequest := *openapiclient.NewV2RuleServeStaticRequest("The requested resource was not found", true, []string{"Domain_example"}, []string{"Url_example"}, "https://backend.test-domain.com", "/index.html") // V2RuleServeStaticRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.RulesServeStaticUpdate(context.Background(), organization, project, rule).V2RuleServeStaticRequest(v2RuleServeStaticRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.RulesServeStaticUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticUpdate`: V2RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.RulesServeStaticUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**rule** | **string** | Rule identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesServeStaticUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2RuleServeStaticRequest** | [**V2RuleServeStaticRequest**](V2RuleServeStaticRequest.md) |  | 

### Return type

[**V2RuleServeStatic**](V2RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

