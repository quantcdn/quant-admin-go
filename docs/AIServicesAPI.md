# \AIServicesAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatInference**](AIServicesAPI.md#ChatInference) | **Post** /api/v3/organizations/{organisation}/ai/chat | Chat inference via API Gateway (buffered responses) with multimodal support
[**ChatInferenceStream**](AIServicesAPI.md#ChatInferenceStream) | **Post** /api/v3/organizations/{organisation}/ai/chat/stream | Chat inference via streaming endpoint (true HTTP streaming) with multimodal support
[**CreateAISession**](AIServicesAPI.md#CreateAISession) | **Post** /api/v3/organizations/{organisation}/ai/sessions | Create a new chat session with multi-tenant isolation
[**DeleteAISession**](AIServicesAPI.md#DeleteAISession) | **Delete** /api/v3/organizations/{organisation}/ai/sessions/{sessionId} | Delete a chat session
[**Embeddings**](AIServicesAPI.md#Embeddings) | **Post** /api/v3/organizations/{organisation}/ai/embeddings | Generate text embeddings for semantic search and RAG applications
[**GetAIConfig**](AIServicesAPI.md#GetAIConfig) | **Get** /api/v3/organizations/{organisation}/ai/config | Get AI configuration for an organization
[**GetAISession**](AIServicesAPI.md#GetAISession) | **Get** /api/v3/organizations/{organisation}/ai/sessions/{sessionId} | Get a specific chat session
[**GetAIUsageStats**](AIServicesAPI.md#GetAIUsageStats) | **Get** /api/v3/organizations/{organisation}/ai/usage | Get AI usage statistics
[**GetToolExecutionStatus**](AIServicesAPI.md#GetToolExecutionStatus) | **Get** /api/v3/organizations/{organisation}/ai/tools/executions/{executionId} | Get async tool execution status and result
[**ImageGeneration**](AIServicesAPI.md#ImageGeneration) | **Post** /api/v3/organizations/{organisation}/ai/image-generation | Generate images with Amazon Nova Canvas
[**ListAIModels**](AIServicesAPI.md#ListAIModels) | **Get** /api/v3/organizations/{organisation}/ai/models | List available AI models for an organization
[**ListAISessions**](AIServicesAPI.md#ListAISessions) | **Get** /api/v3/organizations/{organisation}/ai/sessions | List chat sessions with multi-tenant filtering
[**ListAIToolNames**](AIServicesAPI.md#ListAIToolNames) | **Get** /api/v3/organizations/{organisation}/ai/tools/names | List tool names only (lightweight response)
[**ListAITools**](AIServicesAPI.md#ListAITools) | **Get** /api/v3/organizations/{organisation}/ai/tools | List available built-in tools for function calling
[**ListToolExecutions**](AIServicesAPI.md#ListToolExecutions) | **Get** /api/v3/organizations/{organisation}/ai/tools/executions | List tool executions for monitoring and debugging
[**UpdateAIConfig**](AIServicesAPI.md#UpdateAIConfig) | **Put** /api/v3/organizations/{organisation}/ai/config | Update AI configuration for an organization



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
	resp, r, err := apiClient.AIServicesAPI.ChatInference(context.Background(), organisation).ChatInferenceRequest(chatInferenceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.ChatInference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatInference`: ChatInference200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.ChatInference`: %v\n", resp)
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
	resp, r, err := apiClient.AIServicesAPI.ChatInferenceStream(context.Background(), organisation).ChatInferenceStreamRequest(chatInferenceStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.ChatInferenceStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatInferenceStream`: string
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.ChatInferenceStream`: %v\n", resp)
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


## CreateAISession

> CreateAISession201Response CreateAISession(ctx, organisation).CreateAISessionRequest(createAISessionRequest).Execute()

Create a new chat session with multi-tenant isolation



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
	createAISessionRequest := *openapiclient.NewCreateAISessionRequest("user-12345") // CreateAISessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIServicesAPI.CreateAISession(context.Background(), organisation).CreateAISessionRequest(createAISessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.CreateAISession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAISession`: CreateAISession201Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.CreateAISession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAISessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAISessionRequest** | [**CreateAISessionRequest**](CreateAISessionRequest.md) |  | 

### Return type

[**CreateAISession201Response**](CreateAISession201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAISession

> DeleteAISession200Response DeleteAISession(ctx, organisation, sessionId).Execute()

Delete a chat session

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
	sessionId := "sessionId_example" // string | The session ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIServicesAPI.DeleteAISession(context.Background(), organisation, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.DeleteAISession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAISession`: DeleteAISession200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.DeleteAISession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**sessionId** | **string** | The session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAISessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteAISession200Response**](DeleteAISession200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	resp, r, err := apiClient.AIServicesAPI.Embeddings(context.Background(), organisation).EmbeddingsRequest(embeddingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.Embeddings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Embeddings`: Embeddings200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.Embeddings`: %v\n", resp)
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


## GetAIConfig

> GetAIConfig200Response GetAIConfig(ctx, organisation).Execute()

Get AI configuration for an organization

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
	resp, r, err := apiClient.AIServicesAPI.GetAIConfig(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.GetAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfig`: GetAIConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.GetAIConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAIConfig200Response**](GetAIConfig200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAISession

> GetAISession200Response GetAISession(ctx, organisation, sessionId).Execute()

Get a specific chat session

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
	sessionId := "sessionId_example" // string | The session ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIServicesAPI.GetAISession(context.Background(), organisation, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.GetAISession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAISession`: GetAISession200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.GetAISession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**sessionId** | **string** | The session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAISessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAISession200Response**](GetAISession200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIUsageStats

> GetAIUsageStats200Response GetAIUsageStats(ctx, organisation).Month(month).Execute()

Get AI usage statistics

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
	month := "2025-10" // string | Month to retrieve statistics for (YYYY-MM format) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIServicesAPI.GetAIUsageStats(context.Background(), organisation).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.GetAIUsageStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIUsageStats`: GetAIUsageStats200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.GetAIUsageStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIUsageStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **month** | **string** | Month to retrieve statistics for (YYYY-MM format) | 

### Return type

[**GetAIUsageStats200Response**](GetAIUsageStats200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolExecutionStatus

> GetToolExecutionStatus200Response GetToolExecutionStatus(ctx, organisation, executionId).Execute()

Get async tool execution status and result



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
	executionId := "exec_0123456789abcdef0123456789abcdef" // string | Tool execution identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIServicesAPI.GetToolExecutionStatus(context.Background(), organisation, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.GetToolExecutionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToolExecutionStatus`: GetToolExecutionStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.GetToolExecutionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**executionId** | **string** | Tool execution identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolExecutionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetToolExecutionStatus200Response**](GetToolExecutionStatus200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.AIServicesAPI.ImageGeneration(context.Background(), organisation).ImageGenerationRequest(imageGenerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.ImageGeneration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImageGeneration`: ImageGeneration200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.ImageGeneration`: %v\n", resp)
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
	resp, r, err := apiClient.AIServicesAPI.ListAIModels(context.Background(), organisation).Feature(feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.ListAIModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIModels`: ListAIModels200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.ListAIModels`: %v\n", resp)
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


## ListAISessions

> []ListAISessions200ResponseInner ListAISessions(ctx, organisation).UserId(userId).SessionGroup(sessionGroup).Limit(limit).Offset(offset).Model(model).Execute()

List chat sessions with multi-tenant filtering



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
	userId := "user-12345" // string | Filter sessions by user ID (optional)
	sessionGroup := "drupal-production" // string | Filter by session group. Returns only sessions matching the specified group. (optional)
	limit := int32(56) // int32 | Maximum number of sessions to return (default 50, max 100) (optional) (default to 50)
	offset := int32(56) // int32 | Offset for pagination (optional)
	model := "model_example" // string | Filter by model ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIServicesAPI.ListAISessions(context.Background(), organisation).UserId(userId).SessionGroup(sessionGroup).Limit(limit).Offset(offset).Model(model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.ListAISessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAISessions`: []ListAISessions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.ListAISessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAISessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Filter sessions by user ID | 
 **sessionGroup** | **string** | Filter by session group. Returns only sessions matching the specified group. | 
 **limit** | **int32** | Maximum number of sessions to return (default 50, max 100) | [default to 50]
 **offset** | **int32** | Offset for pagination | 
 **model** | **string** | Filter by model ID | 

### Return type

[**[]ListAISessions200ResponseInner**](ListAISessions200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIToolNames

> ListAIToolNames200Response ListAIToolNames(ctx, organisation).Execute()

List tool names only (lightweight response)



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
	resp, r, err := apiClient.AIServicesAPI.ListAIToolNames(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.ListAIToolNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIToolNames`: ListAIToolNames200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.ListAIToolNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIToolNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAIToolNames200Response**](ListAIToolNames200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAITools

> ListAITools200Response ListAITools(ctx, organisation).Execute()

List available built-in tools for function calling



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
	resp, r, err := apiClient.AIServicesAPI.ListAITools(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.ListAITools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAITools`: ListAITools200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.ListAITools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListAITools200Response**](ListAITools200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListToolExecutions

> ListToolExecutions200Response ListToolExecutions(ctx, organisation).Status(status).Limit(limit).Execute()

List tool executions for monitoring and debugging



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
	status := "status_example" // string | Filter by execution status (optional)
	limit := int32(56) // int32 | Maximum number of executions to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIServicesAPI.ListToolExecutions(context.Background(), organisation).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.ListToolExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListToolExecutions`: ListToolExecutions200Response
	fmt.Fprintf(os.Stdout, "Response from `AIServicesAPI.ListToolExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListToolExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by execution status | 
 **limit** | **int32** | Maximum number of executions to return | [default to 50]

### Return type

[**ListToolExecutions200Response**](ListToolExecutions200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAIConfig

> UpdateAIConfig(ctx, organisation).UpdateAIConfigRequest(updateAIConfigRequest).Execute()

Update AI configuration for an organization

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
	updateAIConfigRequest := *openapiclient.NewUpdateAIConfigRequest() // UpdateAIConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIServicesAPI.UpdateAIConfig(context.Background(), organisation).UpdateAIConfigRequest(updateAIConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIServicesAPI.UpdateAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAIConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAIConfigRequest** | [**UpdateAIConfigRequest**](UpdateAIConfigRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

