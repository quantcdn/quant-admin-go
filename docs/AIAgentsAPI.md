# \AIAgentsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatWithAIAgent**](AIAgentsAPI.md#ChatWithAIAgent) | **Post** /api/v3/organizations/{organisation}/ai/agents/{agentId}/chat | Chat with AI Agent
[**CreateAIAgent**](AIAgentsAPI.md#CreateAIAgent) | **Post** /api/v3/organizations/{organisation}/ai/agents | Create AI Agent
[**DeleteAIAgent**](AIAgentsAPI.md#DeleteAIAgent) | **Delete** /api/v3/organizations/{organisation}/ai/agents/{agentId} | Delete Agent
[**GetAIAgent**](AIAgentsAPI.md#GetAIAgent) | **Get** /api/v3/organizations/{organisation}/ai/agents/{agentId} | Get Agent Details
[**ListAIAgents**](AIAgentsAPI.md#ListAIAgents) | **Get** /api/v3/organizations/{organisation}/ai/agents | List AI Agents
[**UpdateAIAgent**](AIAgentsAPI.md#UpdateAIAgent) | **Put** /api/v3/organizations/{organisation}/ai/agents/{agentId} | Update Agent



## ChatWithAIAgent

> ChatWithAIAgent200Response ChatWithAIAgent(ctx, organisation, agentId).ChatWithAIAgentRequest(chatWithAIAgentRequest).Execute()

Chat with AI Agent



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The agent ID
	chatWithAIAgentRequest := *openapiclient.NewChatWithAIAgentRequest("Message_example") // ChatWithAIAgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.ChatWithAIAgent(context.Background(), organisation, agentId).ChatWithAIAgentRequest(chatWithAIAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.ChatWithAIAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatWithAIAgent`: ChatWithAIAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.ChatWithAIAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**agentId** | **string** | The agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatWithAIAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chatWithAIAgentRequest** | [**ChatWithAIAgentRequest**](ChatWithAIAgentRequest.md) |  | 

### Return type

[**ChatWithAIAgent200Response**](ChatWithAIAgent200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAIAgent

> CreateAIAgent201Response CreateAIAgent(ctx, organisation).CreateAIAgentRequest(createAIAgentRequest).Execute()

Create AI Agent



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
	createAIAgentRequest := *openapiclient.NewCreateAIAgentRequest("Code Review Assistant", "Reviews code for security vulnerabilities and best practices", "You are a senior software engineer specializing in secure code review.", "anthropic.claude-3-5-sonnet-20241022-v2:0") // CreateAIAgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.CreateAIAgent(context.Background(), organisation).CreateAIAgentRequest(createAIAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.CreateAIAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAIAgent`: CreateAIAgent201Response
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.CreateAIAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAIAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAIAgentRequest** | [**CreateAIAgentRequest**](CreateAIAgentRequest.md) |  | 

### Return type

[**CreateAIAgent201Response**](CreateAIAgent201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAIAgent

> DeleteAIAgent200Response DeleteAIAgent(ctx, organisation, agentId).Execute()

Delete Agent



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.DeleteAIAgent(context.Background(), organisation, agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.DeleteAIAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAIAgent`: DeleteAIAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.DeleteAIAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**agentId** | **string** | The agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAIAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteAIAgent200Response**](DeleteAIAgent200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIAgent

> GetAIAgent200Response GetAIAgent(ctx, organisation, agentId).Execute()

Get Agent Details



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.GetAIAgent(context.Background(), organisation, agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAIAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIAgent`: GetAIAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAIAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**agentId** | **string** | The agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAIAgent200Response**](GetAIAgent200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIAgents

> ListAIAgents200Response ListAIAgents(ctx, organisation).Group(group).Execute()

List AI Agents



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
	group := "group_example" // string | Optional group filter (e.g., 'development', 'compliance') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.ListAIAgents(context.Background(), organisation).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.ListAIAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIAgents`: ListAIAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.ListAIAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | **string** | Optional group filter (e.g., &#39;development&#39;, &#39;compliance&#39;) | 

### Return type

[**ListAIAgents200Response**](ListAIAgents200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAIAgent

> UpdateAIAgent200Response UpdateAIAgent(ctx, organisation, agentId).UpdateAIAgentRequest(updateAIAgentRequest).Execute()

Update Agent



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The agent ID
	updateAIAgentRequest := *openapiclient.NewUpdateAIAgentRequest() // UpdateAIAgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.UpdateAIAgent(context.Background(), organisation, agentId).UpdateAIAgentRequest(updateAIAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.UpdateAIAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAIAgent`: UpdateAIAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.UpdateAIAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**agentId** | **string** | The agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAIAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAIAgentRequest** | [**UpdateAIAgentRequest**](UpdateAIAgentRequest.md) |  | 

### Return type

[**UpdateAIAgent200Response**](UpdateAIAgent200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

