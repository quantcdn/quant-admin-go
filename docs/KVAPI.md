# \KVAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KVCreate**](KVAPI.md#KVCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv | Add a kv store
[**KVDelete**](KVAPI.md#KVDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id} | Delete a kv store
[**KVItemsCreate**](KVAPI.md#KVItemsCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items | Add an item to a kv store
[**KVItemsDelete**](KVAPI.md#KVItemsDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key} | Delete an item from a kv store
[**KVItemsList**](KVAPI.md#KVItemsList) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items | List items in a kv store
[**KVItemsShow**](KVAPI.md#KVItemsShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key} | Get an item from a kv store
[**KVItemsUpdate**](KVAPI.md#KVItemsUpdate) | **Put** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key} | Update an item in a kv store
[**KVList**](KVAPI.md#KVList) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv | List key-value stores
[**KVShow**](KVAPI.md#KVShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id} | Get a kv store



## KVCreate

> V2Store KVCreate(ctx, organization, project).V2StoreRequest(v2StoreRequest).Execute()

Add a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	v2StoreRequest := *openapiclient.NewV2StoreRequest("session-data") // V2StoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVCreate(context.Background(), organization, project).V2StoreRequest(v2StoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVCreate`: V2Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2StoreRequest** | [**V2StoreRequest**](V2StoreRequest.md) |  | 

### Return type

[**V2Store**](V2Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVDelete

> KVDelete(ctx, organization, project, storeId).Execute()

Delete a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KVAPI.KVDelete(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsCreate

> KVItemsCreate200Response KVItemsCreate(ctx, organization, project, storeId).V2StoreItemRequest(v2StoreItemRequest).Execute()

Add an item to a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	v2StoreItemRequest := *openapiclient.NewV2StoreItemRequest("user-session-123", "Value_example") // V2StoreItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsCreate(context.Background(), organization, project, storeId).V2StoreItemRequest(v2StoreItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsCreate`: KVItemsCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2StoreItemRequest** | [**V2StoreItemRequest**](V2StoreItemRequest.md) |  | 

### Return type

[**KVItemsCreate200Response**](KVItemsCreate200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsDelete

> KVItemsDelete200Response KVItemsDelete(ctx, organization, project, storeId, key).Execute()

Delete an item from a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsDelete(context.Background(), organization, project, storeId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsDelete`: KVItemsDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**KVItemsDelete200Response**](KVItemsDelete200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsList

> V2StoreItemsListResponse KVItemsList(ctx, organization, project, storeId).Cursor(cursor).Limit(limit).Search(search).IncludeValues(includeValues).Execute()

List items in a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	cursor := "cursor_example" // string | Cursor for pagination (optional)
	limit := int32(56) // int32 | Number of items to return (optional) (default to 10)
	search := "search_example" // string | Search filter for keys (optional)
	includeValues := true // bool | Include values in the response. Secret values will be redacted as '[ENCRYPTED]' for security. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsList(context.Background(), organization, project, storeId).Cursor(cursor).Limit(limit).Search(search).IncludeValues(includeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsList`: V2StoreItemsListResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **string** | Cursor for pagination | 
 **limit** | **int32** | Number of items to return | [default to 10]
 **search** | **string** | Search filter for keys | 
 **includeValues** | **bool** | Include values in the response. Secret values will be redacted as &#39;[ENCRYPTED]&#39; for security. | [default to false]

### Return type

[**V2StoreItemsListResponse**](V2StoreItemsListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsShow

> KVItemsShow200Response KVItemsShow(ctx, organization, project, storeId, key).Execute()

Get an item from a kv store



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsShow(context.Background(), organization, project, storeId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsShow`: KVItemsShow200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**KVItemsShow200Response**](KVItemsShow200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsUpdate

> KVItemsCreate200Response KVItemsUpdate(ctx, organization, project, storeId, key).V2StoreItemUpdateRequest(v2StoreItemUpdateRequest).Execute()

Update an item in a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	key := "key_example" // string | 
	v2StoreItemUpdateRequest := *openapiclient.NewV2StoreItemUpdateRequest("Value_example") // V2StoreItemUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsUpdate(context.Background(), organization, project, storeId, key).V2StoreItemUpdateRequest(v2StoreItemUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsUpdate`: KVItemsCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2StoreItemUpdateRequest** | [**V2StoreItemUpdateRequest**](V2StoreItemUpdateRequest.md) |  | 

### Return type

[**KVItemsCreate200Response**](KVItemsCreate200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVList

> []V2Store KVList(ctx, organization, project).Execute()

List key-value stores

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVList`: []V2Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2Store**](V2Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVShow

> V2Store KVShow(ctx, organization, project, storeId).Execute()

Get a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVShow(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVShow`: V2Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2Store**](V2Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

