# \CommandsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommand**](CommandsAPI.md#CreateCommand) | **Post** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/commands | Create a command for an environment
[**GetCommand**](CommandsAPI.md#GetCommand) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/commands/{command} | Get a command
[**ListCommands**](CommandsAPI.md#ListCommands) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/commands | Get all commands for an environment



## CreateCommand

> Command CreateCommand(ctx, organisation, application, environment).CreateCommandRequest(createCommandRequest).Execute()

Create a command for an environment

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
	application := "test-app" // string | The application ID
	environment := "test-env" // string | The environment ID
	createCommandRequest := *openapiclient.NewCreateCommandRequest("Command_example") // CreateCommandRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommandsAPI.CreateCommand(context.Background(), organisation, application, environment).CreateCommandRequest(createCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandsAPI.CreateCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommand`: Command
	fmt.Fprintf(os.Stdout, "Response from `CommandsAPI.CreateCommand`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createCommandRequest** | [**CreateCommandRequest**](CreateCommandRequest.md) |  | 

### Return type

[**Command**](Command.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommand

> Command GetCommand(ctx, organisation, application, environment, command).Execute()

Get a command

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
	application := "test-app" // string | The application ID
	environment := "test-env" // string | The environment ID
	command := "test-cmd" // string | The command run ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommandsAPI.GetCommand(context.Background(), organisation, application, environment, command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandsAPI.GetCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommand`: Command
	fmt.Fprintf(os.Stdout, "Response from `CommandsAPI.GetCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**command** | **string** | The command run ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Command**](Command.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommands

> ListCommands(ctx, organisation, application, environment).Execute()

Get all commands for an environment

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
	application := "test-app" // string | The application ID
	environment := "test-env" // string | The environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommandsAPI.ListCommands(context.Background(), organisation, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandsAPI.ListCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

Other parameters are passed through a pointer to a apiListCommandsRequest struct via the builder pattern


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

