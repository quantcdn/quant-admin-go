# \ResourcesAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachOrgResource**](ResourcesAPI.md#AttachOrgResource) | **Post** /api/v3/organizations/{organisation}/resources/{resource}/attachments | Attach a resource to an application environment
[**CreateOrgResource**](ResourcesAPI.md#CreateOrgResource) | **Post** /api/v3/organizations/{organisation}/resources | Create a shared resource
[**DeleteOrgResource**](ResourcesAPI.md#DeleteOrgResource) | **Delete** /api/v3/organizations/{organisation}/resources/{resource} | Delete a shared resource
[**DetachOrgResource**](ResourcesAPI.md#DetachOrgResource) | **Delete** /api/v3/organizations/{organisation}/resources/{resource}/attachments/{application}/{environment} | Detach a resource from an application environment
[**GetOrgResource**](ResourcesAPI.md#GetOrgResource) | **Get** /api/v3/organizations/{organisation}/resources/{resource} | Get a shared resource and its attachments
[**GetOrgResourceCredentials**](ResourcesAPI.md#GetOrgResourceCredentials) | **Get** /api/v3/organizations/{organisation}/resources/{resource}/credentials | Get a cache&#39;s administrative credential
[**ListOrgResources**](ResourcesAPI.md#ListOrgResources) | **Get** /api/v3/organizations/{organisation}/resources | List an organisation&#39;s shared resources
[**PurgeOrgResource**](ResourcesAPI.md#PurgeOrgResource) | **Post** /api/v3/organizations/{organisation}/resources/{resource}/purge | Purge keys from a cache



## AttachOrgResource

> ResourceAttachment AttachOrgResource(ctx, organisation, resource).AttachOrgResourceRequest(attachOrgResourceRequest).Execute()

Attach a resource to an application environment



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
	resource := "res-abc123" // string | The resource ID
	attachOrgResourceRequest := *openapiclient.NewAttachOrgResourceRequest("test-app", "production") // AttachOrgResourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.AttachOrgResource(context.Background(), organisation, resource).AttachOrgResourceRequest(attachOrgResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.AttachOrgResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachOrgResource`: ResourceAttachment
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.AttachOrgResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**resource** | **string** | The resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachOrgResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachOrgResourceRequest** | [**AttachOrgResourceRequest**](AttachOrgResourceRequest.md) |  | 

### Return type

[**ResourceAttachment**](ResourceAttachment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgResource

> OrgResource CreateOrgResource(ctx, organisation).CreateOrgResourceRequest(createOrgResourceRequest).Execute()

Create a shared resource



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
	createOrgResourceRequest := *openapiclient.NewCreateOrgResourceRequest("object-storage", "buzz-media") // CreateOrgResourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.CreateOrgResource(context.Background(), organisation).CreateOrgResourceRequest(createOrgResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.CreateOrgResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgResource`: OrgResource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.CreateOrgResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrgResourceRequest** | [**CreateOrgResourceRequest**](CreateOrgResourceRequest.md) |  | 

### Return type

[**OrgResource**](OrgResource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgResource

> DeleteOrgResource(ctx, organisation, resource).Force(force).Execute()

Delete a shared resource



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
	resource := "res-abc123" // string | The resource ID
	force := true // bool | Delete even if the resource is attached or in an error state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.DeleteOrgResource(context.Background(), organisation, resource).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.DeleteOrgResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**resource** | **string** | The resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Delete even if the resource is attached or in an error state | 

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


## DetachOrgResource

> DetachOrgResource(ctx, organisation, resource, application, environment).Execute()

Detach a resource from an application environment



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
	resource := "res-abc123" // string | The resource ID
	application := "test-app" // string | The application ID
	environment := "production" // string | The environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.DetachOrgResource(context.Background(), organisation, resource, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.DetachOrgResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**resource** | **string** | The resource ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachOrgResourceRequest struct via the builder pattern


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


## GetOrgResource

> OrgResource GetOrgResource(ctx, organisation, resource).Execute()

Get a shared resource and its attachments

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
	resource := "res-abc123" // string | The resource ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetOrgResource(context.Background(), organisation, resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetOrgResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgResource`: OrgResource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetOrgResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**resource** | **string** | The resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgResource**](OrgResource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgResourceCredentials

> GetOrgResourceCredentials200Response GetOrgResourceCredentials(ctx, organisation, resource).Execute()

Get a cache's administrative credential



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
	resource := "vk-sessions" // string | The resource ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.GetOrgResourceCredentials(context.Background(), organisation, resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.GetOrgResourceCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgResourceCredentials`: GetOrgResourceCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.GetOrgResourceCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**resource** | **string** | The resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgResourceCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrgResourceCredentials200Response**](GetOrgResourceCredentials200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgResources

> []OrgResource ListOrgResources(ctx, organisation).Execute()

List an organisation's shared resources

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ListOrgResources(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ListOrgResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgResources`: []OrgResource
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ListOrgResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrgResource**](OrgResource.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeOrgResource

> PurgeOrgResource200Response PurgeOrgResource(ctx, organisation, resource).PurgeOrgResourceRequest(purgeOrgResourceRequest).Execute()

Purge keys from a cache



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
	resource := "vk-sessions" // string | The resource ID
	purgeOrgResourceRequest := *openapiclient.NewPurgeOrgResourceRequest("Scope_example") // PurgeOrgResourceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.PurgeOrgResource(context.Background(), organisation, resource).PurgeOrgResourceRequest(purgeOrgResourceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.PurgeOrgResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurgeOrgResource`: PurgeOrgResource200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.PurgeOrgResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**resource** | **string** | The resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeOrgResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **purgeOrgResourceRequest** | [**PurgeOrgResourceRequest**](PurgeOrgResourceRequest.md) |  | 

### Return type

[**PurgeOrgResource200Response**](PurgeOrgResource200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

