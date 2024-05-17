# \RulesCustomResponseAPI

All URIs are relative to *https://portal.stage.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesCustomResponseCreate**](RulesCustomResponseAPI.md#RulesCustomResponseCreate) | **Post** /organizations/{organization}/projects/{project}/rules/custom-response | 
[**RulesCustomResponseDelete**](RulesCustomResponseAPI.md#RulesCustomResponseDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | 
[**RulesCustomResponseList**](RulesCustomResponseAPI.md#RulesCustomResponseList) | **Get** /organizations/{organization}/projects/{project}/rules/custom-response | 
[**RulesCustomResponseRead**](RulesCustomResponseAPI.md#RulesCustomResponseRead) | **Get** /organizations/{organization}/projects/{project}/rules/custom-response/{rule} | 



## RulesCustomResponseCreate

> Rule RulesCustomResponseCreate(ctx, organization, project).RuleCustomResponserequest(ruleCustomResponserequest).Execute()



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
	ruleCustomResponserequest := *openapiclient.NewRuleCustomResponserequest(int32(123), "CustomResponseBody_example", "Domain_example", false) // RuleCustomResponserequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesCustomResponseAPI.RulesCustomResponseCreate(context.Background(), organization, project).RuleCustomResponserequest(ruleCustomResponserequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesCustomResponseAPI.RulesCustomResponseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesCustomResponseCreate`: Rule
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


 **ruleCustomResponserequest** | [**RuleCustomResponserequest**](RuleCustomResponserequest.md) |  | 

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


## RulesCustomResponseDelete

> Rule RulesCustomResponseDelete(ctx, organization, project, rule).Execute()



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
	// response from `RulesCustomResponseDelete`: Rule
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

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseList

> []Rule RulesCustomResponseList(ctx, organization, project).Execute()



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
	// response from `RulesCustomResponseList`: []Rule
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

[**[]Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesCustomResponseRead

> Rule RulesCustomResponseRead(ctx, organization, project, rule).Execute()



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
	// response from `RulesCustomResponseRead`: Rule
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

[**Rule**](Rule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
