# \VariablesAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkSetEnvironmentVariables**](VariablesAPI.md#BulkSetEnvironmentVariables) | **Put** /api/v3/organizations/{api_organisation}/applications/{api_application}/environments/{api_environment}/variables | Bulk set/replace environment variables
[**DeleteEnvironmentVariable**](VariablesAPI.md#DeleteEnvironmentVariable) | **Delete** /api/v3/organizations/{api_organisation}/applications/{api_application}/environments/{api_environment}/variables/{api_variable} | Delete a variable
[**ListEnvironmentVariables**](VariablesAPI.md#ListEnvironmentVariables) | **Get** /api/v3/organizations/{api_organisation}/applications/{api_application}/environments/{api_environment}/variables | Get all variables for an environment
[**UpdateEnvironmentVariable**](VariablesAPI.md#UpdateEnvironmentVariable) | **Put** /api/v3/organizations/{api_organisation}/applications/{api_application}/environments/{api_environment}/variables/{api_variable} | Update a variable



## BulkSetEnvironmentVariables

> BulkSetEnvironmentVariables(ctx, apiOrganisation, apiApplication, apiEnvironment).BulkSetEnvironmentVariablesRequest(bulkSetEnvironmentVariablesRequest).Execute()

Bulk set/replace environment variables

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
	apiOrganisation := "test-org" // string | The organisation ID
	apiApplication := "test-app" // string | The application ID
	apiEnvironment := "test-env" // string | The environment ID
	bulkSetEnvironmentVariablesRequest := *openapiclient.NewBulkSetEnvironmentVariablesRequest([]openapiclient.BulkSetEnvironmentVariablesRequestEnvironmentInner{*openapiclient.NewBulkSetEnvironmentVariablesRequestEnvironmentInner("Name_example", "Value_example")}) // BulkSetEnvironmentVariablesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VariablesAPI.BulkSetEnvironmentVariables(context.Background(), apiOrganisation, apiApplication, apiEnvironment).BulkSetEnvironmentVariablesRequest(bulkSetEnvironmentVariablesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.BulkSetEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiOrganisation** | **string** | The organisation ID | 
**apiApplication** | **string** | The application ID | 
**apiEnvironment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkSetEnvironmentVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bulkSetEnvironmentVariablesRequest** | [**BulkSetEnvironmentVariablesRequest**](BulkSetEnvironmentVariablesRequest.md) |  | 

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


## DeleteEnvironmentVariable

> DeleteEnvironmentVariable(ctx, apiOrganisation, apiApplication, apiEnvironment, apiVariable).Execute()

Delete a variable

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
	apiOrganisation := "apiOrganisation_example" // string | The organisation ID
	apiApplication := "apiApplication_example" // string | The application ID
	apiEnvironment := "apiEnvironment_example" // string | The environment ID
	apiVariable := "apiVariable_example" // string | The variable key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VariablesAPI.DeleteEnvironmentVariable(context.Background(), apiOrganisation, apiApplication, apiEnvironment, apiVariable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.DeleteEnvironmentVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiOrganisation** | **string** | The organisation ID | 
**apiApplication** | **string** | The application ID | 
**apiEnvironment** | **string** | The environment ID | 
**apiVariable** | **string** | The variable key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentVariableRequest struct via the builder pattern


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


## ListEnvironmentVariables

> ListEnvironmentVariables(ctx, apiOrganisation, apiApplication, apiEnvironment).Execute()

Get all variables for an environment

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
	apiOrganisation := "test-org" // string | The organisation ID
	apiApplication := "test-app" // string | The application ID
	apiEnvironment := "test-env" // string | The environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VariablesAPI.ListEnvironmentVariables(context.Background(), apiOrganisation, apiApplication, apiEnvironment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.ListEnvironmentVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiOrganisation** | **string** | The organisation ID | 
**apiApplication** | **string** | The application ID | 
**apiEnvironment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentVariablesRequest struct via the builder pattern


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


## UpdateEnvironmentVariable

> UpdateEnvironmentVariable(ctx, apiOrganisation, apiApplication, apiEnvironment, apiVariable).UpdateEnvironmentVariableRequest(updateEnvironmentVariableRequest).Execute()

Update a variable

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
	apiOrganisation := "test-org" // string | The organisation ID
	apiApplication := "test-app" // string | The application ID
	apiEnvironment := "test-env" // string | The environment ID
	apiVariable := "apiVariable_example" // string | The variable key
	updateEnvironmentVariableRequest := *openapiclient.NewUpdateEnvironmentVariableRequest() // UpdateEnvironmentVariableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VariablesAPI.UpdateEnvironmentVariable(context.Background(), apiOrganisation, apiApplication, apiEnvironment, apiVariable).UpdateEnvironmentVariableRequest(updateEnvironmentVariableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.UpdateEnvironmentVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiOrganisation** | **string** | The organisation ID | 
**apiApplication** | **string** | The application ID | 
**apiEnvironment** | **string** | The environment ID | 
**apiVariable** | **string** | The variable key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateEnvironmentVariableRequest** | [**UpdateEnvironmentVariableRequest**](UpdateEnvironmentVariableRequest.md) |  | 

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

