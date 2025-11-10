# \PurgeAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PurgeCreate**](PurgeAPI.md#PurgeCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/purge | Purge cache via URL or cache keys



## PurgeCreate

> string PurgeCreate(ctx, organization, project).PurgeCreateRequest(purgeCreateRequest).Execute()

Purge cache via URL or cache keys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	purgeCreateRequest := *openapiclient.NewPurgeCreateRequest([]string{"CacheKeys_example"}, "Scope_example") // PurgeCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurgeAPI.PurgeCreate(context.Background(), organization, project).PurgeCreateRequest(purgeCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurgeAPI.PurgeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurgeCreate`: string
	fmt.Fprintf(os.Stdout, "Response from `PurgeAPI.PurgeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **purgeCreateRequest** | [**PurgeCreateRequest**](PurgeCreateRequest.md) |  | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

