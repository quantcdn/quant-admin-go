# \IdleSleepAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdleSleep**](IdleSleepAPI.md#GetIdleSleep) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/idle-sleep | Get Idle Sleep Setting
[**SetIdleSleep**](IdleSleepAPI.md#SetIdleSleep) | **Put** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/idle-sleep | Set Idle Sleep Setting



## GetIdleSleep

> IdleSleepResponse GetIdleSleep(ctx, organisation, application, environment).Execute()

Get Idle Sleep Setting



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
	organisation := "organisation_example" // string | 
	application := "application_example" // string | 
	environment := "environment_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdleSleepAPI.GetIdleSleep(context.Background(), organisation, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdleSleepAPI.GetIdleSleep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdleSleep`: IdleSleepResponse
	fmt.Fprintf(os.Stdout, "Response from `IdleSleepAPI.GetIdleSleep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**application** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdleSleepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IdleSleepResponse**](IdleSleepResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIdleSleep

> IdleSleepResponse SetIdleSleep(ctx, organisation, application, environment).SetIdleSleepRequest(setIdleSleepRequest).Execute()

Set Idle Sleep Setting



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
	organisation := "organisation_example" // string | 
	application := "application_example" // string | 
	environment := "environment_example" // string | 
	setIdleSleepRequest := *openapiclient.NewSetIdleSleepRequest(false) // SetIdleSleepRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdleSleepAPI.SetIdleSleep(context.Background(), organisation, application, environment).SetIdleSleepRequest(setIdleSleepRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdleSleepAPI.SetIdleSleep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetIdleSleep`: IdleSleepResponse
	fmt.Fprintf(os.Stdout, "Response from `IdleSleepAPI.SetIdleSleep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**application** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetIdleSleepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **setIdleSleepRequest** | [**SetIdleSleepRequest**](SetIdleSleepRequest.md) |  | 

### Return type

[**IdleSleepResponse**](IdleSleepResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

