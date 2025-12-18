# \HeadersAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HeadersCreate**](HeadersAPI.md#HeadersCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/custom-headers | Create or update custom headers
[**HeadersDelete**](HeadersAPI.md#HeadersDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/custom-headers | Delete custom headers
[**HeadersList**](HeadersAPI.md#HeadersList) | **Get** /api/v2/organizations/{organization}/projects/{project}/custom-headers | List custom headers for a project



## HeadersCreate

> map[string]string HeadersCreate(ctx, organization, project).V2CustomHeaderRequest(v2CustomHeaderRequest).Execute()

Create or update custom headers

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	v2CustomHeaderRequest := *openapiclient.NewV2CustomHeaderRequest(map[string]string{"key": "Inner_example"}) // V2CustomHeaderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeadersAPI.HeadersCreate(context.Background(), organization, project).V2CustomHeaderRequest(v2CustomHeaderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeadersAPI.HeadersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadersCreate`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HeadersAPI.HeadersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2CustomHeaderRequest** | [**V2CustomHeaderRequest**](V2CustomHeaderRequest.md) |  | 

### Return type

**map[string]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadersDelete

> HeadersDelete(ctx, organization, project).V2CustomHeaderRequest(v2CustomHeaderRequest).Execute()

Delete custom headers

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	v2CustomHeaderRequest := *openapiclient.NewV2CustomHeaderRequest(map[string]string{"key": "Inner_example"}) // V2CustomHeaderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HeadersAPI.HeadersDelete(context.Background(), organization, project).V2CustomHeaderRequest(v2CustomHeaderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeadersAPI.HeadersDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2CustomHeaderRequest** | [**V2CustomHeaderRequest**](V2CustomHeaderRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadersList

> map[string]string HeadersList(ctx, organization, project).Execute()

List custom headers for a project

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HeadersAPI.HeadersList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HeadersAPI.HeadersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadersList`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HeadersAPI.HeadersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

