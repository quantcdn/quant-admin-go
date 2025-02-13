# \RulesContentFilterAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesContentFilterList**](RulesContentFilterAPI.md#RulesContentFilterList) | **Get** /organizations/{organization}/projects/{project}/rules/content-filter | 



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

