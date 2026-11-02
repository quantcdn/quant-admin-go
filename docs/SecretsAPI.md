# \SecretsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEnvironmentSecret**](SecretsAPI.md#DeleteEnvironmentSecret) | **Delete** /api/v3/organizations/{api_organisation}/applications/{api_application}/environments/{api_environment}/secrets/{api_secret} | Delete a secret
[**ListEnvironmentSecrets**](SecretsAPI.md#ListEnvironmentSecrets) | **Get** /api/v3/organizations/{api_organisation}/applications/{api_application}/environments/{api_environment}/secrets | List the secret names for an environment
[**UpdateEnvironmentSecret**](SecretsAPI.md#UpdateEnvironmentSecret) | **Put** /api/v3/organizations/{api_organisation}/applications/{api_application}/environments/{api_environment}/secrets/{api_secret} | Create or replace a secret



## DeleteEnvironmentSecret

> DeleteEnvironmentSecret200Response DeleteEnvironmentSecret(ctx, apiOrganisation, apiApplication, apiEnvironment, apiSecret).Execute()

Delete a secret



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
	apiSecret := "apiSecret_example" // string | The secret key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.DeleteEnvironmentSecret(context.Background(), apiOrganisation, apiApplication, apiEnvironment, apiSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.DeleteEnvironmentSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEnvironmentSecret`: DeleteEnvironmentSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.DeleteEnvironmentSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiOrganisation** | **string** | The organisation ID | 
**apiApplication** | **string** | The application ID | 
**apiEnvironment** | **string** | The environment ID | 
**apiSecret** | **string** | The secret key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DeleteEnvironmentSecret200Response**](DeleteEnvironmentSecret200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentSecrets

> []SecretKey ListEnvironmentSecrets(ctx, apiOrganisation, apiApplication, apiEnvironment).Execute()

List the secret names for an environment



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
	resp, r, err := apiClient.SecretsAPI.ListEnvironmentSecrets(context.Background(), apiOrganisation, apiApplication, apiEnvironment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.ListEnvironmentSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironmentSecrets`: []SecretKey
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.ListEnvironmentSecrets`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListEnvironmentSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]SecretKey**](SecretKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironmentSecret

> SecretWriteResult UpdateEnvironmentSecret(ctx, apiOrganisation, apiApplication, apiEnvironment, apiSecret).UpdateEnvironmentSecretRequest(updateEnvironmentSecretRequest).Execute()

Create or replace a secret



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
	apiSecret := "DB_PASSWORD" // string | The secret key
	updateEnvironmentSecretRequest := *openapiclient.NewUpdateEnvironmentSecretRequest("Value_example") // UpdateEnvironmentSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.UpdateEnvironmentSecret(context.Background(), apiOrganisation, apiApplication, apiEnvironment, apiSecret).UpdateEnvironmentSecretRequest(updateEnvironmentSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.UpdateEnvironmentSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnvironmentSecret`: SecretWriteResult
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.UpdateEnvironmentSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiOrganisation** | **string** | The organisation ID | 
**apiApplication** | **string** | The application ID | 
**apiEnvironment** | **string** | The environment ID | 
**apiSecret** | **string** | The secret key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateEnvironmentSecretRequest** | [**UpdateEnvironmentSecretRequest**](UpdateEnvironmentSecretRequest.md) |  | 

### Return type

[**SecretWriteResult**](SecretWriteResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

