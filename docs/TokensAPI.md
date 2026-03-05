# \TokensAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokensCreate**](TokensAPI.md#TokensCreate) | **Post** /api/v2/organizations/{organization}/tokens | Create a new API token scoped to this organization
[**TokensDelete**](TokensAPI.md#TokensDelete) | **Delete** /api/v2/organizations/{organization}/tokens/{token_id} | Revoke an API token
[**TokensList**](TokensAPI.md#TokensList) | **Get** /api/v2/organizations/{organization}/tokens | List API tokens scoped to this organization



## TokensCreate

> TokensCreate201Response TokensCreate(ctx, organization).TokensCreateRequest(tokensCreateRequest).Execute()

Create a new API token scoped to this organization

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
	tokensCreateRequest := *openapiclient.NewTokensCreateRequest("My deploy token") // TokensCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.TokensCreate(context.Background(), organization).TokensCreateRequest(tokensCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.TokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensCreate`: TokensCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.TokensCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokensCreateRequest** | [**TokensCreateRequest**](TokensCreateRequest.md) |  | 

### Return type

[**TokensCreate201Response**](TokensCreate201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensDelete

> TokensDelete200Response TokensDelete(ctx, organization, tokenId).Execute()

Revoke an API token

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
	tokenId := int32(42) // int32 | Token ID to revoke

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.TokensDelete(context.Background(), organization, tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.TokensDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensDelete`: TokensDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.TokensDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**tokenId** | **int32** | Token ID to revoke | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokensDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TokensDelete200Response**](TokensDelete200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokensList

> []TokensList200ResponseInner TokensList(ctx, organization).Execute()

List API tokens scoped to this organization

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.TokensList(context.Background(), organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.TokensList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokensList`: []TokensList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.TokensList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TokensList200ResponseInner**](TokensList200ResponseInner.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

