# \OpenAICompatibilityAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OaiChatCompletions**](OpenAICompatibilityAPI.md#OaiChatCompletions) | **Post** /oai/v1/chat/completions | Create a chat completion (OpenAI-compatible)
[**OaiEmbeddings**](OpenAICompatibilityAPI.md#OaiEmbeddings) | **Post** /oai/v1/embeddings | Create embeddings (OpenAI-compatible)
[**OaiGetModel**](OpenAICompatibilityAPI.md#OaiGetModel) | **Get** /oai/v1/models/{model} | Retrieve a model (OpenAI-compatible)
[**OaiListModels**](OpenAICompatibilityAPI.md#OaiListModels) | **Get** /oai/v1/models | List available models (OpenAI-compatible)



## OaiChatCompletions

> OaiChatCompletions200Response OaiChatCompletions(ctx).OaiChatCompletionsRequest(oaiChatCompletionsRequest).Execute()

Create a chat completion (OpenAI-compatible)



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
	oaiChatCompletionsRequest := *openapiclient.NewOaiChatCompletionsRequest("anthropic.claude-sonnet-4-6", []openapiclient.OaiChatCompletionsRequestMessagesInner{*openapiclient.NewOaiChatCompletionsRequestMessagesInner()}) // OaiChatCompletionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAICompatibilityAPI.OaiChatCompletions(context.Background()).OaiChatCompletionsRequest(oaiChatCompletionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAICompatibilityAPI.OaiChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OaiChatCompletions`: OaiChatCompletions200Response
	fmt.Fprintf(os.Stdout, "Response from `OpenAICompatibilityAPI.OaiChatCompletions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOaiChatCompletionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oaiChatCompletionsRequest** | [**OaiChatCompletionsRequest**](OaiChatCompletionsRequest.md) |  | 

### Return type

[**OaiChatCompletions200Response**](OaiChatCompletions200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OaiEmbeddings

> OaiEmbeddings200Response OaiEmbeddings(ctx).OaiEmbeddingsRequest(oaiEmbeddingsRequest).Execute()

Create embeddings (OpenAI-compatible)



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
	oaiEmbeddingsRequest := *openapiclient.NewOaiEmbeddingsRequest("amazon.titan-embed-text-v2:0", interface{}(123)) // OaiEmbeddingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAICompatibilityAPI.OaiEmbeddings(context.Background()).OaiEmbeddingsRequest(oaiEmbeddingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAICompatibilityAPI.OaiEmbeddings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OaiEmbeddings`: OaiEmbeddings200Response
	fmt.Fprintf(os.Stdout, "Response from `OpenAICompatibilityAPI.OaiEmbeddings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOaiEmbeddingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oaiEmbeddingsRequest** | [**OaiEmbeddingsRequest**](OaiEmbeddingsRequest.md) |  | 

### Return type

[**OaiEmbeddings200Response**](OaiEmbeddings200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OaiGetModel

> OaiGetModel200Response OaiGetModel(ctx, model).Execute()

Retrieve a model (OpenAI-compatible)



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
	model := "amazon.nova-micro-v1:0" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAICompatibilityAPI.OaiGetModel(context.Background(), model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAICompatibilityAPI.OaiGetModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OaiGetModel`: OaiGetModel200Response
	fmt.Fprintf(os.Stdout, "Response from `OpenAICompatibilityAPI.OaiGetModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOaiGetModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OaiGetModel200Response**](OaiGetModel200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OaiListModels

> OaiListModels200Response OaiListModels(ctx).Execute()

List available models (OpenAI-compatible)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAICompatibilityAPI.OaiListModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAICompatibilityAPI.OaiListModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OaiListModels`: OaiListModels200Response
	fmt.Fprintf(os.Stdout, "Response from `OpenAICompatibilityAPI.OaiListModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOaiListModelsRequest struct via the builder pattern


### Return type

[**OaiListModels200Response**](OaiListModels200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

