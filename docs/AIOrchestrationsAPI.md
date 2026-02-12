# \AIOrchestrationsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrchestration**](AIOrchestrationsAPI.md#CancelOrchestration) | **Post** /api/v3/organizations/{organisation}/ai/orchestrations/{orchestrationId}/cancel | Cancel Durable Orchestration
[**CreateOrchestration**](AIOrchestrationsAPI.md#CreateOrchestration) | **Post** /api/v3/organizations/{organisation}/ai/orchestrations | Create Durable Orchestration
[**DeleteOrchestration**](AIOrchestrationsAPI.md#DeleteOrchestration) | **Delete** /api/v3/organizations/{organisation}/ai/orchestrations/{orchestrationId} | Delete Durable Orchestration
[**GetOrchestration**](AIOrchestrationsAPI.md#GetOrchestration) | **Get** /api/v3/organizations/{organisation}/ai/orchestrations/{orchestrationId} | Get Durable Orchestration
[**ListOrchestrationBatches**](AIOrchestrationsAPI.md#ListOrchestrationBatches) | **Get** /api/v3/organizations/{organisation}/ai/orchestrations/{orchestrationId}/batches | List Orchestration Batches
[**ListOrchestrations**](AIOrchestrationsAPI.md#ListOrchestrations) | **Get** /api/v3/organizations/{organisation}/ai/orchestrations | List Durable Orchestrations
[**PauseOrchestration**](AIOrchestrationsAPI.md#PauseOrchestration) | **Post** /api/v3/organizations/{organisation}/ai/orchestrations/{orchestrationId}/pause | Pause Durable Orchestration
[**ResumeOrchestration**](AIOrchestrationsAPI.md#ResumeOrchestration) | **Post** /api/v3/organizations/{organisation}/ai/orchestrations/{orchestrationId}/resume | Resume Durable Orchestration
[**StartOrchestration**](AIOrchestrationsAPI.md#StartOrchestration) | **Post** /api/v3/organizations/{organisation}/ai/orchestrations/{orchestrationId}/start | Start Durable Orchestration



## CancelOrchestration

> map[string]interface{} CancelOrchestration(ctx, organisation, orchestrationId).Execute()

Cancel Durable Orchestration



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
	organisation := "organisation_example" // string | The organisation machine name
	orchestrationId := "orchestrationId_example" // string | Orchestration identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOrchestrationsAPI.CancelOrchestration(context.Background(), organisation, orchestrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.CancelOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrchestration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIOrchestrationsAPI.CancelOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**orchestrationId** | **string** | Orchestration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrchestration

> map[string]interface{} CreateOrchestration(ctx, organisation).CreateOrchestrationRequest(createOrchestrationRequest).Execute()

Create Durable Orchestration



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
	organisation := "organisation_example" // string | The organisation machine name
	createOrchestrationRequest := *openapiclient.NewCreateOrchestrationRequest("Process batch documents", *openapiclient.NewCreateOrchestrationRequestInputSource("Type_example")) // CreateOrchestrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOrchestrationsAPI.CreateOrchestration(context.Background(), organisation).CreateOrchestrationRequest(createOrchestrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.CreateOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrchestration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIOrchestrationsAPI.CreateOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrchestrationRequest** | [**CreateOrchestrationRequest**](CreateOrchestrationRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrchestration

> DeleteOrchestration(ctx, organisation, orchestrationId).Execute()

Delete Durable Orchestration



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
	organisation := "organisation_example" // string | The organisation machine name
	orchestrationId := "orchestrationId_example" // string | Orchestration identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIOrchestrationsAPI.DeleteOrchestration(context.Background(), organisation, orchestrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.DeleteOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**orchestrationId** | **string** | Orchestration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrchestration

> map[string]interface{} GetOrchestration(ctx, organisation, orchestrationId).Execute()

Get Durable Orchestration



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
	organisation := "organisation_example" // string | The organisation machine name
	orchestrationId := "orch_1704067200_abc123xyz" // string | Orchestration identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOrchestrationsAPI.GetOrchestration(context.Background(), organisation, orchestrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.GetOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrchestration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIOrchestrationsAPI.GetOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**orchestrationId** | **string** | Orchestration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrchestrationBatches

> ListOrchestrationBatches200Response ListOrchestrationBatches(ctx, organisation, orchestrationId).Limit(limit).Cursor(cursor).Execute()

List Orchestration Batches



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
	organisation := "organisation_example" // string | The organisation machine name
	orchestrationId := "orchestrationId_example" // string | Orchestration identifier
	limit := int32(56) // int32 | Maximum number of batches to return (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor from previous response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOrchestrationsAPI.ListOrchestrationBatches(context.Background(), organisation, orchestrationId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.ListOrchestrationBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrchestrationBatches`: ListOrchestrationBatches200Response
	fmt.Fprintf(os.Stdout, "Response from `AIOrchestrationsAPI.ListOrchestrationBatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**orchestrationId** | **string** | Orchestration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrchestrationBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Maximum number of batches to return | [default to 20]
 **cursor** | **string** | Pagination cursor from previous response | 

### Return type

[**ListOrchestrationBatches200Response**](ListOrchestrationBatches200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrchestrations

> ListOrchestrations200Response ListOrchestrations(ctx, organisation).Status(status).Limit(limit).Cursor(cursor).Execute()

List Durable Orchestrations



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
	organisation := "organisation_example" // string | The organisation machine name
	status := "status_example" // string | Filter by orchestration status (optional)
	limit := int32(56) // int32 | Maximum number of results (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor from previous response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOrchestrationsAPI.ListOrchestrations(context.Background(), organisation).Status(status).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.ListOrchestrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrchestrations`: ListOrchestrations200Response
	fmt.Fprintf(os.Stdout, "Response from `AIOrchestrationsAPI.ListOrchestrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrchestrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by orchestration status | 
 **limit** | **int32** | Maximum number of results | [default to 20]
 **cursor** | **string** | Pagination cursor from previous response | 

### Return type

[**ListOrchestrations200Response**](ListOrchestrations200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseOrchestration

> map[string]interface{} PauseOrchestration(ctx, organisation, orchestrationId).Execute()

Pause Durable Orchestration



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
	organisation := "organisation_example" // string | The organisation machine name
	orchestrationId := "orchestrationId_example" // string | Orchestration identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOrchestrationsAPI.PauseOrchestration(context.Background(), organisation, orchestrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.PauseOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseOrchestration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIOrchestrationsAPI.PauseOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**orchestrationId** | **string** | Orchestration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeOrchestration

> map[string]interface{} ResumeOrchestration(ctx, organisation, orchestrationId).Execute()

Resume Durable Orchestration



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
	organisation := "organisation_example" // string | The organisation machine name
	orchestrationId := "orchestrationId_example" // string | Orchestration identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOrchestrationsAPI.ResumeOrchestration(context.Background(), organisation, orchestrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.ResumeOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeOrchestration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIOrchestrationsAPI.ResumeOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**orchestrationId** | **string** | Orchestration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOrchestration

> map[string]interface{} StartOrchestration(ctx, organisation, orchestrationId).Execute()

Start Durable Orchestration



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
	organisation := "organisation_example" // string | The organisation machine name
	orchestrationId := "orchestrationId_example" // string | Orchestration identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOrchestrationsAPI.StartOrchestration(context.Background(), organisation, orchestrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOrchestrationsAPI.StartOrchestration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOrchestration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIOrchestrationsAPI.StartOrchestration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**orchestrationId** | **string** | Orchestration identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

