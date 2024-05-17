# \RulesProxyAPI

All URIs are relative to *https://portal.stage.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesProxyCreate**](RulesProxyAPI.md#RulesProxyCreate) | **Post** /organizations/{organization}/projects/{project}/rules/proxy | 
[**RulesProxyDelete**](RulesProxyAPI.md#RulesProxyDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | 
[**RulesProxyList**](RulesProxyAPI.md#RulesProxyList) | **Get** /organizations/{organization}/projects/{project}/rules/proxy | 
[**RulesProxyRead**](RulesProxyAPI.md#RulesProxyRead) | **Get** /organizations/{organization}/projects/{project}/rules/proxy/{rule} | 



## RulesProxyCreate

> Rule RulesProxyCreate(ctx, organization, project).RuleProxyRequest(ruleProxyRequest).Execute()



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
	ruleProxyRequest := *openapiclient.NewRuleProxyRequest("To_example", "Domain_example", false) // RuleProxyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesProxyAPI.RulesProxyCreate(context.Background(), organization, project).RuleProxyRequest(ruleProxyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesProxyAPI.RulesProxyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesProxyCreate`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesProxyAPI.RulesProxyCreate`: %v\n", resp)
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


 **ruleProxyRequest** | [**RuleProxyRequest**](RuleProxyRequest.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesProxyDelete

> Rule RulesProxyDelete(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.RulesProxyAPI.RulesProxyDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesProxyAPI.RulesProxyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesProxyDelete`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesProxyAPI.RulesProxyDelete`: %v\n", resp)
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

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesProxyList

> []Rule RulesProxyList(ctx, organization, project).Execute()



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
	resp, r, err := apiClient.RulesProxyAPI.RulesProxyList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesProxyAPI.RulesProxyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesProxyList`: []Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesProxyAPI.RulesProxyList`: %v\n", resp)
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

[**[]Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesProxyRead

> Rule RulesProxyRead(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.RulesProxyAPI.RulesProxyRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesProxyAPI.RulesProxyRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesProxyRead`: Rule
	fmt.Fprintf(os.Stdout, "Response from `RulesProxyAPI.RulesProxyRead`: %v\n", resp)
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

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

