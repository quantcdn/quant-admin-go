# \RulesCustomResponseAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesCustomResponseCreate**](RulesCustomResponseAPI.md#RulesCustomResponseCreate) | **Post** /organizations/{organization}/projects/{project}/rules/custom-response | 
[**RulesCustomResponseDelete**](RulesCustomResponseAPI.md#RulesCustomResponseDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | 
[**RulesCustomResponseList**](RulesCustomResponseAPI.md#RulesCustomResponseList) | **Get** /organizations/{organization}/projects/{project}/rules/custom-response | 
[**RulesCustomResponseRead**](RulesCustomResponseAPI.md#RulesCustomResponseRead) | **Get** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | 
[**RulesCustomResponseUpdate**](RulesCustomResponseAPI.md#RulesCustomResponseUpdate) | **Patch** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | 



## RulesCustomResponseCreate

> RuleCustomResponse RulesCustomResponseCreate(ctx, organization, project).RuleCustomResponseRequest(ruleCustomResponseRequest).Execute()



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
	ruleCustomResponseRequest := *openapiclient.NewRuleCustomResponseRequest([]string{"Domain_example"}, false, []string{"Url_example"}, "CustomResponseBody_example", int32(123)) // RuleCustomResponseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesCustomResponseAPI.RulesCustomResponseCreate(context.Background(), organization, project).RuleCustomResponseRequest(ruleCustomResponseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesCustomResponseAPI.RulesCustomResponseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseCreate`: RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesCustomResponseAPI.RulesCustomResponseCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesCustomResponseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleCustomResponseRequest** | [**RuleCustomResponseRequest**](RuleCustomResponseRequest.md) |  | 

### Return type

[**RuleCustomResponse**](RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseDelete

> RuleCustomResponse RulesCustomResponseDelete(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.RulesCustomResponseAPI.RulesCustomResponseDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesCustomResponseAPI.RulesCustomResponseDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseDelete`: RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesCustomResponseAPI.RulesCustomResponseDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesCustomResponseDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleCustomResponse**](RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseList

> []RuleCustomResponse RulesCustomResponseList(ctx, organization, project).Execute()



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
	resp, r, err := apiClient.RulesCustomResponseAPI.RulesCustomResponseList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesCustomResponseAPI.RulesCustomResponseList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseList`: []RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesCustomResponseAPI.RulesCustomResponseList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesCustomResponseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RuleCustomResponse**](RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseRead

> RuleCustomResponse RulesCustomResponseRead(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.RulesCustomResponseAPI.RulesCustomResponseRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesCustomResponseAPI.RulesCustomResponseRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseRead`: RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesCustomResponseAPI.RulesCustomResponseRead`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesCustomResponseReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleCustomResponse**](RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseUpdate

> RuleCustomResponse RulesCustomResponseUpdate(ctx, organization, project, rule).RuleCustomResponseRequestUpdate(ruleCustomResponseRequestUpdate).Execute()



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
	ruleCustomResponseRequestUpdate := *openapiclient.NewRuleCustomResponseRequestUpdate() // RuleCustomResponseRequestUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesCustomResponseAPI.RulesCustomResponseUpdate(context.Background(), organization, project, rule).RuleCustomResponseRequestUpdate(ruleCustomResponseRequestUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesCustomResponseAPI.RulesCustomResponseUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseUpdate`: RuleCustomResponse
	fmt.Fprintf(os.Stdout, "Response from `RulesCustomResponseAPI.RulesCustomResponseUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesCustomResponseUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleCustomResponseRequestUpdate** | [**RuleCustomResponseRequestUpdate**](RuleCustomResponseRequestUpdate.md) |  | 

### Return type

[**RuleCustomResponse**](RuleCustomResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

