# \AIModelsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAIModel**](AIModelsAPI.md#GetAIModel) | **Get** /api/v3/organizations/{organisation}/ai/models/{modelId} | Get AI Model Details
[**ListAIModels**](AIModelsAPI.md#ListAIModels) | **Get** /api/v3/organizations/{organisation}/ai/models | List available AI models for an organization



## GetAIModel

> GetAIModel200Response GetAIModel(ctx, organisation, modelId).Execute()

Get AI Model Details



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
	modelId := "amazon.nova-lite-v1:0" // string | The model identifier (e.g., amazon.nova-lite-v1:0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelsAPI.GetAIModel(context.Background(), organisation, modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.GetAIModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIModel`: GetAIModel200Response
	fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.GetAIModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**modelId** | **string** | The model identifier (e.g., amazon.nova-lite-v1:0) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAIModel200Response**](GetAIModel200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIModels

> ListAIModels200Response ListAIModels(ctx, organisation).Feature(feature).Execute()

List available AI models for an organization

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
	feature := "embeddings" // string | Filter models by supported feature (optional) (default to "all")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIModelsAPI.ListAIModels(context.Background(), organisation).Feature(feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIModelsAPI.ListAIModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModels`: ListAIModels200Response
	fmt.Fprintf(os.Stdout, "Response from `AIModelsAPI.ListAIModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feature** | **string** | Filter models by supported feature | [default to &quot;all&quot;]

### Return type

[**ListAIModels200Response**](ListAIModels200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

