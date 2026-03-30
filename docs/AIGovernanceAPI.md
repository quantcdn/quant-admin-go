# \AIGovernanceAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGovernanceConfig**](AIGovernanceAPI.md#GetGovernanceConfig) | **Get** /api/v3/organizations/{organisation}/ai/governance | Get AI governance configuration for an organisation
[**GetGovernanceSpend**](AIGovernanceAPI.md#GetGovernanceSpend) | **Get** /api/v3/organizations/{organisation}/ai/governance/spend | Get AI spend summary for an organisation
[**UpdateGovernanceConfig**](AIGovernanceAPI.md#UpdateGovernanceConfig) | **Put** /api/v3/organizations/{organisation}/ai/governance | Update AI governance configuration for an organisation



## GetGovernanceConfig

> GetGovernanceConfig200Response GetGovernanceConfig(ctx, organisation).Execute()

Get AI governance configuration for an organisation

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
	resp, r, err := apiClient.AIGovernanceAPI.GetGovernanceConfig(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGovernanceAPI.GetGovernanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGovernanceConfig`: GetGovernanceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `AIGovernanceAPI.GetGovernanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGovernanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetGovernanceConfig200Response**](GetGovernanceConfig200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGovernanceSpend

> map[string]interface{} GetGovernanceSpend(ctx, organisation).Execute()

Get AI spend summary for an organisation

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
	resp, r, err := apiClient.AIGovernanceAPI.GetGovernanceSpend(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGovernanceAPI.GetGovernanceSpend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGovernanceSpend`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIGovernanceAPI.GetGovernanceSpend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGovernanceSpendRequest struct via the builder pattern


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


## UpdateGovernanceConfig

> UpdateGovernanceConfig200Response UpdateGovernanceConfig(ctx, organisation).UpdateGovernanceConfigRequest(updateGovernanceConfigRequest).Execute()

Update AI governance configuration for an organisation

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
	updateGovernanceConfigRequest := *openapiclient.NewUpdateGovernanceConfigRequest(false, "ModelPolicy_example") // UpdateGovernanceConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIGovernanceAPI.UpdateGovernanceConfig(context.Background(), organisation).UpdateGovernanceConfigRequest(updateGovernanceConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIGovernanceAPI.UpdateGovernanceConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGovernanceConfig`: UpdateGovernanceConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `AIGovernanceAPI.UpdateGovernanceConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGovernanceConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGovernanceConfigRequest** | [**UpdateGovernanceConfigRequest**](UpdateGovernanceConfigRequest.md) |  | 

### Return type

[**UpdateGovernanceConfig200Response**](UpdateGovernanceConfig200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

