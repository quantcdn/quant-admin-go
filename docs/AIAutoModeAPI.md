# \AIAutoModeAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutoMode**](AIAutoModeAPI.md#DeleteAutoMode) | **Delete** /api/v3/organizations/{organisation}/ai/auto-mode | Remove the auto mode setting for an organisation
[**GetAutoMode**](AIAutoModeAPI.md#GetAutoMode) | **Get** /api/v3/organizations/{organisation}/ai/auto-mode | Get the auto mode setting for an organisation
[**UpdateAutoMode**](AIAutoModeAPI.md#UpdateAutoMode) | **Put** /api/v3/organizations/{organisation}/ai/auto-mode | Save the auto mode setting for an organisation



## DeleteAutoMode

> DeleteAutoMode200Response DeleteAutoMode(ctx, organisation).Execute()

Remove the auto mode setting for an organisation

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
	resp, r, err := apiClient.AIAutoModeAPI.DeleteAutoMode(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAutoModeAPI.DeleteAutoMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAutoMode`: DeleteAutoMode200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAutoModeAPI.DeleteAutoMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAutoMode200Response**](DeleteAutoMode200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoMode

> GetAutoMode200Response GetAutoMode(ctx, organisation).Execute()

Get the auto mode setting for an organisation

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
	resp, r, err := apiClient.AIAutoModeAPI.GetAutoMode(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAutoModeAPI.GetAutoMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoMode`: GetAutoMode200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAutoModeAPI.GetAutoMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAutoMode200Response**](GetAutoMode200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoMode

> UpdateAutoMode200Response UpdateAutoMode(ctx, organisation).AIAutoModeInput(aIAutoModeInput).Execute()

Save the auto mode setting for an organisation

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
	aIAutoModeInput := *openapiclient.NewAIAutoModeInput(false, *openapiclient.NewAIAutoModeInputGroup(*openapiclient.NewAIAutoModeConfigGroupRoles([]string{"Fast_example"}, []string{"Standard_example"}, []string{"Deep_example"})), int32(123), "Enforcement_example", false) // AIAutoModeInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAutoModeAPI.UpdateAutoMode(context.Background(), organisation).AIAutoModeInput(aIAutoModeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAutoModeAPI.UpdateAutoMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoMode`: UpdateAutoMode200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAutoModeAPI.UpdateAutoMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aIAutoModeInput** | [**AIAutoModeInput**](AIAutoModeInput.md) |  | 

### Return type

[**UpdateAutoMode200Response**](UpdateAutoMode200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

