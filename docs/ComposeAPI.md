# \ComposeAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironmentCompose**](ComposeAPI.md#GetEnvironmentCompose) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/compose | Get the compose file for an environment
[**PatchEnvironmentCompose**](ComposeAPI.md#PatchEnvironmentCompose) | **Patch** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/compose | Partially Update Environment Compose Definition
[**ValidateCompose**](ComposeAPI.md#ValidateCompose) | **Post** /api/v3/organizations/{organisation}/compose/validate | Validate a compose file



## GetEnvironmentCompose

> Compose GetEnvironmentCompose(ctx, organisation, application, environment).Execute()

Get the compose file for an environment

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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeAPI.GetEnvironmentCompose(context.Background(), organisation, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeAPI.GetEnvironmentCompose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentCompose`: Compose
	fmt.Fprintf(os.Stdout, "Response from `ComposeAPI.GetEnvironmentCompose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentComposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Compose**](Compose.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEnvironmentCompose

> PatchEnvironmentCompose202Response PatchEnvironmentCompose(ctx, organisation, application, environment).PatchEnvironmentComposeRequest(patchEnvironmentComposeRequest).Execute()

Partially Update Environment Compose Definition



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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	patchEnvironmentComposeRequest := *openapiclient.NewPatchEnvironmentComposeRequest() // PatchEnvironmentComposeRequest | Partial compose definition updates. All fields are optional.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeAPI.PatchEnvironmentCompose(context.Background(), organisation, application, environment).PatchEnvironmentComposeRequest(patchEnvironmentComposeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeAPI.PatchEnvironmentCompose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEnvironmentCompose`: PatchEnvironmentCompose202Response
	fmt.Fprintf(os.Stdout, "Response from `ComposeAPI.PatchEnvironmentCompose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEnvironmentComposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchEnvironmentComposeRequest** | [**PatchEnvironmentComposeRequest**](PatchEnvironmentComposeRequest.md) | Partial compose definition updates. All fields are optional. | 

### Return type

[**PatchEnvironmentCompose202Response**](PatchEnvironmentCompose202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCompose

> ValidateCompose200Response ValidateCompose(ctx, organisation).ValidateComposeRequest(validateComposeRequest).ImageSuffix(imageSuffix).Execute()

Validate a compose file



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
	validateComposeRequest := *openapiclient.NewValidateComposeRequest("Compose_example") // ValidateComposeRequest | The docker-compose.yml file content. Can be sent as raw YAML string or as a JSON wrapper containing both yamlContent (string) and imageSuffix (string) fields. Query parameter imageSuffix takes precedence if both are provided.
	imageSuffix := "pr-456" // string | Optional. Image tag suffix to apply during translation. Transforms internal image tags to consistent '{containerName}-{suffix}' format (e.g., 'nginx-pr-456'). External images are left unchanged. Useful for feature branch deployments. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeAPI.ValidateCompose(context.Background(), organisation).ValidateComposeRequest(validateComposeRequest).ImageSuffix(imageSuffix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeAPI.ValidateCompose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateCompose`: ValidateCompose200Response
	fmt.Fprintf(os.Stdout, "Response from `ComposeAPI.ValidateCompose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateComposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validateComposeRequest** | [**ValidateComposeRequest**](ValidateComposeRequest.md) | The docker-compose.yml file content. Can be sent as raw YAML string or as a JSON wrapper containing both yamlContent (string) and imageSuffix (string) fields. Query parameter imageSuffix takes precedence if both are provided. | 
 **imageSuffix** | **string** | Optional. Image tag suffix to apply during translation. Transforms internal image tags to consistent &#39;{containerName}-{suffix}&#39; format (e.g., &#39;nginx-pr-456&#39;). External images are left unchanged. Useful for feature branch deployments. | 

### Return type

[**ValidateCompose200Response**](ValidateCompose200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

