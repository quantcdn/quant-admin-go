# \AICustomToolsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomTool**](AICustomToolsAPI.md#CreateCustomTool) | **Post** /api/v3/organizations/{organisation}/ai/custom-tools | Register Custom Edge Function Tool
[**DeleteCustomTool**](AICustomToolsAPI.md#DeleteCustomTool) | **Delete** /api/v3/organizations/{organisation}/ai/custom-tools/{toolName} | Delete Custom Tool
[**ListCustomTools**](AICustomToolsAPI.md#ListCustomTools) | **Get** /api/v3/organizations/{organisation}/ai/custom-tools | List Custom Tools



## CreateCustomTool

> CreateCustomTool201Response CreateCustomTool(ctx, organisation).CreateCustomToolRequest(createCustomToolRequest).Execute()

Register Custom Edge Function Tool



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
	organisation := "organisation_example" // string | The organisation ID
	createCustomToolRequest := *openapiclient.NewCreateCustomToolRequest("check_inventory", "Check product inventory levels in warehouse", "https://inventory-api.my-project.quant.cloud", map[string]interface{}(123)) // CreateCustomToolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AICustomToolsAPI.CreateCustomTool(context.Background(), organisation).CreateCustomToolRequest(createCustomToolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AICustomToolsAPI.CreateCustomTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomTool`: CreateCustomTool201Response
	fmt.Fprintf(os.Stdout, "Response from `AICustomToolsAPI.CreateCustomTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCustomToolRequest** | [**CreateCustomToolRequest**](CreateCustomToolRequest.md) |  | 

### Return type

[**CreateCustomTool201Response**](CreateCustomTool201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomTool

> DeleteCustomTool200Response DeleteCustomTool(ctx, organisation, toolName).Execute()

Delete Custom Tool



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
	organisation := "organisation_example" // string | The organisation ID
	toolName := "toolName_example" // string | The tool name to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AICustomToolsAPI.DeleteCustomTool(context.Background(), organisation, toolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AICustomToolsAPI.DeleteCustomTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomTool`: DeleteCustomTool200Response
	fmt.Fprintf(os.Stdout, "Response from `AICustomToolsAPI.DeleteCustomTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**toolName** | **string** | The tool name to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteCustomTool200Response**](DeleteCustomTool200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomTools

> ListCustomTools200Response ListCustomTools(ctx, organisation).Execute()

List Custom Tools



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
	organisation := "organisation_example" // string | The organisation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AICustomToolsAPI.ListCustomTools(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AICustomToolsAPI.ListCustomTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomTools`: ListCustomTools200Response
	fmt.Fprintf(os.Stdout, "Response from `AICustomToolsAPI.ListCustomTools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListCustomTools200Response**](ListCustomTools200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

