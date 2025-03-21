# \DefaultAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesServeStaticCreate**](DefaultAPI.md#RulesServeStaticCreate) | **Post** /organizations/{organization}/projects/{project}/rules/serve-static | 
[**RulesServeStaticDelete**](DefaultAPI.md#RulesServeStaticDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/serve-static/{rule} | 
[**RulesServeStaticList**](DefaultAPI.md#RulesServeStaticList) | **Get** /organizations/{organization}/projects/{project}/rules/serve-static | 
[**RulesServeStaticRead**](DefaultAPI.md#RulesServeStaticRead) | **Get** /organizations/{organization}/projects/{project}/rules/serve-static/{rule} | 
[**RulesServeStaticUpdate**](DefaultAPI.md#RulesServeStaticUpdate) | **Patch** /organizations/{organization}/projects/{project}/rules/serve-static/{rule} | 



## RulesServeStaticCreate

> RuleServeStatic RulesServeStaticCreate(ctx, organization, project).RuleServeStaticRequest(ruleServeStaticRequest).Execute()



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
	ruleServeStaticRequest := *openapiclient.NewRuleServeStaticRequest([]string{"Domain_example"}, false, []string{"Url_example"}, "StaticFilePath_example") // RuleServeStaticRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RulesServeStaticCreate(context.Background(), organization, project).RuleServeStaticRequest(ruleServeStaticRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RulesServeStaticCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticCreate`: RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RulesServeStaticCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesServeStaticCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleServeStaticRequest** | [**RuleServeStaticRequest**](RuleServeStaticRequest.md) |  | 

### Return type

[**RuleServeStatic**](RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticDelete

> RuleServeStatic RulesServeStaticDelete(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.DefaultAPI.RulesServeStaticDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RulesServeStaticDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticDelete`: RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RulesServeStaticDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesServeStaticDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleServeStatic**](RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticList

> []RuleServeStatic RulesServeStaticList(ctx, organization, project).Execute()



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
	resp, r, err := apiClient.DefaultAPI.RulesServeStaticList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RulesServeStaticList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticList`: []RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RulesServeStaticList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesServeStaticListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RuleServeStatic**](RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticRead

> RuleServeStatic RulesServeStaticRead(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.DefaultAPI.RulesServeStaticRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RulesServeStaticRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticRead`: RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RulesServeStaticRead`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesServeStaticReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleServeStatic**](RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesServeStaticUpdate

> RuleServeStatic RulesServeStaticUpdate(ctx, organization, project, rule).RuleServeStaticRequestUpdate(ruleServeStaticRequestUpdate).Execute()



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
	ruleServeStaticRequestUpdate := *openapiclient.NewRuleServeStaticRequestUpdate() // RuleServeStaticRequestUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RulesServeStaticUpdate(context.Background(), organization, project, rule).RuleServeStaticRequestUpdate(ruleServeStaticRequestUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RulesServeStaticUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesServeStaticUpdate`: RuleServeStatic
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RulesServeStaticUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesServeStaticUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleServeStaticRequestUpdate** | [**RuleServeStaticRequestUpdate**](RuleServeStaticRequestUpdate.md) |  | 

### Return type

[**RuleServeStatic**](RuleServeStatic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

