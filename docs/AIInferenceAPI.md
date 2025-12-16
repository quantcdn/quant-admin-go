# \AIInferenceAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatInference**](AIInferenceAPI.md#ChatInference) | **Post** /api/v3/organizations/{organisation}/ai/chat | Chat inference via API Gateway (buffered responses) with multimodal support
[**ChatInferenceStream**](AIInferenceAPI.md#ChatInferenceStream) | **Post** /api/v3/organizations/{organisation}/ai/chat/stream | Chat inference via streaming endpoint (true HTTP streaming) with multimodal support
[**Embeddings**](AIInferenceAPI.md#Embeddings) | **Post** /api/v3/organizations/{organisation}/ai/embeddings | Generate text embeddings for semantic search and RAG applications
[**ImageGeneration**](AIInferenceAPI.md#ImageGeneration) | **Post** /api/v3/organizations/{organisation}/ai/image-generation | Generate images with Amazon Nova Canvas



## ChatInference

> ChatInference200Response ChatInference(ctx, organisation).ChatInferenceRequest(chatInferenceRequest).Execute()

Chat inference via API Gateway (buffered responses) with multimodal support



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
	chatInferenceRequest := *openapiclient.NewChatInferenceRequest([]openapiclient.ChatInferenceRequestMessagesInner{*openapiclient.NewChatInferenceRequestMessagesInner("Role_example", openapiclient.chatInference_request_messages_inner_content{ArrayOfChatInferenceRequestMessagesInnerContentOneOfInner: new([]ChatInferenceRequestMessagesInnerContentOneOfInner)})}, "amazon.nova-lite-v1:0") // ChatInferenceRequest | Chat request with optional multimodal content blocks

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIInferenceAPI.ChatInference(context.Background(), organisation).ChatInferenceRequest(chatInferenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIInferenceAPI.ChatInference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatInference`: ChatInference200Response
	fmt.Fprintf(os.Stdout, "Response from `AIInferenceAPI.ChatInference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatInferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatInferenceRequest** | [**ChatInferenceRequest**](ChatInferenceRequest.md) | Chat request with optional multimodal content blocks | 

### Return type

[**ChatInference200Response**](ChatInference200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatInferenceStream

> string ChatInferenceStream(ctx, organisation).ChatInferenceStreamRequest(chatInferenceStreamRequest).Execute()

Chat inference via streaming endpoint (true HTTP streaming) with multimodal support



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
	chatInferenceStreamRequest := *openapiclient.NewChatInferenceStreamRequest([]openapiclient.ChatInferenceStreamRequestMessagesInner{*openapiclient.NewChatInferenceStreamRequestMessagesInner("Role_example", openapiclient.chatInferenceStream_request_messages_inner_content{ArrayOfMapmapOfStringAny: new([]map[string]interface{})})}, "amazon.nova-lite-v1:0") // ChatInferenceStreamRequest | Chat request with optional multimodal content blocks

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIInferenceAPI.ChatInferenceStream(context.Background(), organisation).ChatInferenceStreamRequest(chatInferenceStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIInferenceAPI.ChatInferenceStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatInferenceStream`: string
	fmt.Fprintf(os.Stdout, "Response from `AIInferenceAPI.ChatInferenceStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatInferenceStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatInferenceStreamRequest** | [**ChatInferenceStreamRequest**](ChatInferenceStreamRequest.md) | Chat request with optional multimodal content blocks | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Embeddings

> Embeddings200Response Embeddings(ctx, organisation).EmbeddingsRequest(embeddingsRequest).Execute()

Generate text embeddings for semantic search and RAG applications



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
	embeddingsRequest := *openapiclient.NewEmbeddingsRequest(openapiclient.embeddings_request_input{ArrayOfString: new([]string)}) // EmbeddingsRequest | Embedding request with single or multiple texts

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIInferenceAPI.Embeddings(context.Background(), organisation).EmbeddingsRequest(embeddingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIInferenceAPI.Embeddings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Embeddings`: Embeddings200Response
	fmt.Fprintf(os.Stdout, "Response from `AIInferenceAPI.Embeddings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmbeddingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **embeddingsRequest** | [**EmbeddingsRequest**](EmbeddingsRequest.md) | Embedding request with single or multiple texts | 

### Return type

[**Embeddings200Response**](Embeddings200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageGeneration

> ImageGeneration200Response ImageGeneration(ctx, organisation).ImageGenerationRequest(imageGenerationRequest).Execute()

Generate images with Amazon Nova Canvas



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
	imageGenerationRequest := *openapiclient.NewImageGenerationRequest("TaskType_example") // ImageGenerationRequest | Image generation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIInferenceAPI.ImageGeneration(context.Background(), organisation).ImageGenerationRequest(imageGenerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIInferenceAPI.ImageGeneration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageGeneration`: ImageGeneration200Response
	fmt.Fprintf(os.Stdout, "Response from `AIInferenceAPI.ImageGeneration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageGenerationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageGenerationRequest** | [**ImageGenerationRequest**](ImageGenerationRequest.md) | Image generation request | 

### Return type

[**ImageGeneration200Response**](ImageGeneration200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

