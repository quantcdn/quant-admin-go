# \SSHAccessAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSshAccessCredentials**](SSHAccessAPI.md#GetSshAccessCredentials) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/ssh-access | Get SSH access credentials for an environment



## GetSshAccessCredentials

> GetSshAccessCredentials200Response GetSshAccessCredentials(ctx, organisation, application, environment).Execute()

Get SSH access credentials for an environment

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
	application := "application_example" // string | The application name
	environment := "environment_example" // string | The environment name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHAccessAPI.GetSshAccessCredentials(context.Background(), organisation, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHAccessAPI.GetSshAccessCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshAccessCredentials`: GetSshAccessCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `SSHAccessAPI.GetSshAccessCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**application** | **string** | The application name | 
**environment** | **string** | The environment name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshAccessCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetSshAccessCredentials200Response**](GetSshAccessCredentials200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

