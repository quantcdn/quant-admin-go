# \CIValuesAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCiValue**](CIValuesAPI.md#DeleteCiValue) | **Delete** /api/v3/organizations/{organisation}/ci/values/{name} | Delete a CI secret or variable
[**ListCiValues**](CIValuesAPI.md#ListCiValues) | **Get** /api/v3/organizations/{organisation}/ci/values | List the CI secrets and variables
[**PutCiValue**](CIValuesAPI.md#PutCiValue) | **Put** /api/v3/organizations/{organisation}/ci/values/{name} | Create or replace a CI secret or variable



## DeleteCiValue

> DeleteCiValue(ctx, organisation, name).Repo(repo).Execute()

Delete a CI secret or variable



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
	organisation := "test-org" // string | The organisation ID
	name := "DEPLOY_KEY" // string | The name, such as DEPLOY_KEY
	repo := "acme/web" // string | Delete the repository's value of this name. Leave out for the organisation's value. Lower-cased by the server. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CIValuesAPI.DeleteCiValue(context.Background(), organisation, name).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIValuesAPI.DeleteCiValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**name** | **string** | The name, such as DEPLOY_KEY | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCiValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **repo** | **string** | Delete the repository&#39;s value of this name. Leave out for the organisation&#39;s value. Lower-cased by the server. | 

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


## ListCiValues

> ListCiValues200Response ListCiValues(ctx, organisation).Repo(repo).Execute()

List the CI secrets and variables



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
	organisation := "test-org" // string | The organisation ID
	repo := "acme/web" // string | Only the organisation's values and this repository's values. Lower-cased by the server. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CIValuesAPI.ListCiValues(context.Background(), organisation).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIValuesAPI.ListCiValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCiValues`: ListCiValues200Response
	fmt.Fprintf(os.Stdout, "Response from `CIValuesAPI.ListCiValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCiValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repo** | **string** | Only the organisation&#39;s values and this repository&#39;s values. Lower-cased by the server. | 

### Return type

[**ListCiValues200Response**](ListCiValues200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCiValue

> CiValueWriteResult PutCiValue(ctx, organisation, name).PutCiValueRequest(putCiValueRequest).Execute()

Create or replace a CI secret or variable



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
	organisation := "test-org" // string | The organisation ID
	name := "DEPLOY_KEY" // string | The name, such as DEPLOY_KEY
	putCiValueRequest := *openapiclient.NewPutCiValueRequest("Kind_example", "Value_example") // PutCiValueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CIValuesAPI.PutCiValue(context.Background(), organisation, name).PutCiValueRequest(putCiValueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIValuesAPI.PutCiValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCiValue`: CiValueWriteResult
	fmt.Fprintf(os.Stdout, "Response from `CIValuesAPI.PutCiValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**name** | **string** | The name, such as DEPLOY_KEY | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCiValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putCiValueRequest** | [**PutCiValueRequest**](PutCiValueRequest.md) |  | 

### Return type

[**CiValueWriteResult**](CiValueWriteResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

