# \AISessionsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAISession**](AISessionsAPI.md#CreateAISession) | **Post** /api/v3/organizations/{organisation}/ai/sessions | Create a new chat session with multi-tenant isolation
[**DeleteAISession**](AISessionsAPI.md#DeleteAISession) | **Delete** /api/v3/organizations/{organisation}/ai/sessions/{sessionId} | Delete a chat session
[**ExtendAISession**](AISessionsAPI.md#ExtendAISession) | **Put** /api/v3/organizations/{organisation}/ai/sessions/{sessionId}/extend | Extend Session Expiration
[**GetAISession**](AISessionsAPI.md#GetAISession) | **Get** /api/v3/organizations/{organisation}/ai/sessions/{sessionId} | Get a specific chat session
[**ListAISessions**](AISessionsAPI.md#ListAISessions) | **Get** /api/v3/organizations/{organisation}/ai/sessions | List chat sessions with multi-tenant filtering
[**UpdateAISession**](AISessionsAPI.md#UpdateAISession) | **Put** /api/v3/organizations/{organisation}/ai/sessions/{sessionId} | Update Session



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
	resp, r, err := apiClient.AISessionsAPI.CreateAISession(context.Background(), organisation).CreateAISessionRequest(createAISessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISessionsAPI.CreateAISession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAISession`: CreateAISession201Response
	fmt.Fprintf(os.Stdout, "Response from `AISessionsAPI.CreateAISession`: %v\n", resp)
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
	resp, r, err := apiClient.AISessionsAPI.DeleteAISession(context.Background(), organisation, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISessionsAPI.DeleteAISession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAISession`: DeleteAISession200Response
	fmt.Fprintf(os.Stdout, "Response from `AISessionsAPI.DeleteAISession`: %v\n", resp)
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


## ExtendAISession

> ExtendAISession200Response ExtendAISession(ctx, organisation, sessionId).ExtendAISessionRequest(extendAISessionRequest).Execute()

Extend Session Expiration



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The session ID
	extendAISessionRequest := *openapiclient.NewExtendAISessionRequest() // ExtendAISessionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISessionsAPI.ExtendAISession(context.Background(), organisation, sessionId).ExtendAISessionRequest(extendAISessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISessionsAPI.ExtendAISession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtendAISession`: ExtendAISession200Response
	fmt.Fprintf(os.Stdout, "Response from `AISessionsAPI.ExtendAISession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**sessionId** | **string** | The session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendAISessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extendAISessionRequest** | [**ExtendAISessionRequest**](ExtendAISessionRequest.md) |  | 

### Return type

[**ExtendAISession200Response**](ExtendAISession200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.AISessionsAPI.GetAISession(context.Background(), organisation, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISessionsAPI.GetAISession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAISession`: GetAISession200Response
	fmt.Fprintf(os.Stdout, "Response from `AISessionsAPI.GetAISession`: %v\n", resp)
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
	resp, r, err := apiClient.AISessionsAPI.ListAISessions(context.Background(), organisation).UserId(userId).SessionGroup(sessionGroup).Limit(limit).Offset(offset).Model(model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISessionsAPI.ListAISessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAISessions`: []ListAISessions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AISessionsAPI.ListAISessions`: %v\n", resp)
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


## UpdateAISession

> UpdateAISession200Response UpdateAISession(ctx, organisation, sessionId).UpdateAISessionRequest(updateAISessionRequest).Execute()

Update Session



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The session ID
	updateAISessionRequest := *openapiclient.NewUpdateAISessionRequest() // UpdateAISessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISessionsAPI.UpdateAISession(context.Background(), organisation, sessionId).UpdateAISessionRequest(updateAISessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISessionsAPI.UpdateAISession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAISession`: UpdateAISession200Response
	fmt.Fprintf(os.Stdout, "Response from `AISessionsAPI.UpdateAISession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**sessionId** | **string** | The session ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAISessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAISessionRequest** | [**UpdateAISessionRequest**](UpdateAISessionRequest.md) |  | 

### Return type

[**UpdateAISession200Response**](UpdateAISession200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

