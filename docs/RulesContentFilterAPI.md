# \RulesContentFilterAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesContentFilterCreate**](RulesContentFilterAPI.md#RulesContentFilterCreate) | **Post** /organizations/{organization}/projects/{project}/rules/content-filter | 
[**RulesContentFilterDelete**](RulesContentFilterAPI.md#RulesContentFilterDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/content-filter/{rule} | 
[**RulesContentFilterList**](RulesContentFilterAPI.md#RulesContentFilterList) | **Get** /organizations/{organization}/projects/{project}/rules/content-filter | 
[**RulesContentFilterRead**](RulesContentFilterAPI.md#RulesContentFilterRead) | **Get** /organizations/{organization}/projects/{project}/rules/content-filter/{rule} | 
[**RulesContentFilterUpdate**](RulesContentFilterAPI.md#RulesContentFilterUpdate) | **Patch** /organizations/{organization}/projects/{project}/rules/content-filter/{rule} | 



## RulesContentFilterCreate

> RuleContentFilter RulesContentFilterCreate(ctx, organization, project).RuleContentFilterRequest(ruleContentFilterRequest).Execute()



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
	ruleContentFilterRequest := *openapiclient.NewRuleContentFilterRequest([]string{"Domain_example"}, false, []string{"Url_example"}, "FnUuid_example") // RuleContentFilterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesContentFilterAPI.RulesContentFilterCreate(context.Background(), organization, project).RuleContentFilterRequest(ruleContentFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesContentFilterAPI.RulesContentFilterCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterCreate`: RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesContentFilterAPI.RulesContentFilterCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesContentFilterCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleContentFilterRequest** | [**RuleContentFilterRequest**](RuleContentFilterRequest.md) |  | 

### Return type

[**RuleContentFilter**](RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterDelete

> RuleContentFilter RulesContentFilterDelete(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.RulesContentFilterAPI.RulesContentFilterDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesContentFilterAPI.RulesContentFilterDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterDelete`: RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesContentFilterAPI.RulesContentFilterDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesContentFilterDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleContentFilter**](RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterList

> []RuleContentFilter RulesContentFilterList(ctx, organization, project).Execute()



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
	resp, r, err := apiClient.RulesContentFilterAPI.RulesContentFilterList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesContentFilterAPI.RulesContentFilterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterList`: []RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesContentFilterAPI.RulesContentFilterList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesContentFilterListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RuleContentFilter**](RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterRead

> RuleContentFilter RulesContentFilterRead(ctx, organization, project, rule).Execute()



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
	resp, r, err := apiClient.RulesContentFilterAPI.RulesContentFilterRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesContentFilterAPI.RulesContentFilterRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterRead`: RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesContentFilterAPI.RulesContentFilterRead`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesContentFilterReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleContentFilter**](RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesContentFilterUpdate

> RuleContentFilter RulesContentFilterUpdate(ctx, organization, project, rule).RuleContentFilterRequestUpdate(ruleContentFilterRequestUpdate).Execute()



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
	ruleContentFilterRequestUpdate := *openapiclient.NewRuleContentFilterRequestUpdate() // RuleContentFilterRequestUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesContentFilterAPI.RulesContentFilterUpdate(context.Background(), organization, project, rule).RuleContentFilterRequestUpdate(ruleContentFilterRequestUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesContentFilterAPI.RulesContentFilterUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesContentFilterUpdate`: RuleContentFilter
	fmt.Fprintf(os.Stdout, "Response from `RulesContentFilterAPI.RulesContentFilterUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRulesContentFilterUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleContentFilterRequestUpdate** | [**RuleContentFilterRequestUpdate**](RuleContentFilterRequestUpdate.md) |  | 

### Return type

[**RuleContentFilter**](RuleContentFilter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

