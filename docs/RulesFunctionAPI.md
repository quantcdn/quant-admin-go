# \RulesFunctionAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesFunctionCreate**](RulesFunctionAPI.md#RulesFunctionCreate) | **Post** /organizations/{organization}/projects/{project}/rules/function | 
[**RulesFunctionDelete**](RulesFunctionAPI.md#RulesFunctionDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/function/{rule} | 
[**RulesFunctionList**](RulesFunctionAPI.md#RulesFunctionList) | **Get** /organizations/{organization}/projects/{project}/rules/function | 
[**RulesFunctionRead**](RulesFunctionAPI.md#RulesFunctionRead) | **Get** /organizations/{organization}/projects/{project}/rules/function/{rule} | 
[**RulesFunctionUpdate**](RulesFunctionAPI.md#RulesFunctionUpdate) | **Patch** /organizations/{organization}/projects/{project}/rules/function/{rule} | 



## RulesFunctionCreate

> RuleFunction RulesFunctionCreate(ctx, organization, project).RuleFunctionRequest(ruleFunctionRequest).Execute()



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
	ruleFunctionRequest := *openapiclient.NewRuleFunctionRequest([]string{"Domain_example"}, false, []string{"Url_example"}, "FnUuid_example") // RuleFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesFunctionAPI.RulesFunctionCreate(context.Background(), organization, project).RuleFunctionRequest(ruleFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesFunctionAPI.RulesFunctionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionCreate`: RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesFunctionAPI.RulesFunctionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesFunctionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleFunctionRequest** | [**RuleFunctionRequest**](RuleFunctionRequest.md) |  | 

### Return type

[**RuleFunction**](RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionDelete

> RuleFunction RulesFunctionDelete(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.RulesFunctionAPI.RulesFunctionDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesFunctionAPI.RulesFunctionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionDelete`: RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesFunctionAPI.RulesFunctionDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesFunctionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleFunction**](RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionList

> []RuleFunction RulesFunctionList(ctx, organization, project).Execute()



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
	resp, r, err := apiClient.RulesFunctionAPI.RulesFunctionList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesFunctionAPI.RulesFunctionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionList`: []RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesFunctionAPI.RulesFunctionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesFunctionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RuleFunction**](RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionRead

> RuleFunction RulesFunctionRead(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.RulesFunctionAPI.RulesFunctionRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesFunctionAPI.RulesFunctionRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionRead`: RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesFunctionAPI.RulesFunctionRead`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesFunctionReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleFunction**](RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesFunctionUpdate

> RuleFunction RulesFunctionUpdate(ctx, organization, project, rule).RuleFunctionRequestUpdate(ruleFunctionRequestUpdate).Execute()



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
	ruleFunctionRequestUpdate := *openapiclient.NewRuleFunctionRequestUpdate() // RuleFunctionRequestUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesFunctionAPI.RulesFunctionUpdate(context.Background(), organization, project, rule).RuleFunctionRequestUpdate(ruleFunctionRequestUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesFunctionAPI.RulesFunctionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesFunctionUpdate`: RuleFunction
	fmt.Fprintf(os.Stdout, "Response from `RulesFunctionAPI.RulesFunctionUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesFunctionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleFunctionRequestUpdate** | [**RuleFunctionRequestUpdate**](RuleFunctionRequestUpdate.md) |  | 

### Return type

[**RuleFunction**](RuleFunction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

