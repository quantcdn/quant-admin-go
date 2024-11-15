# \KVItemsAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KVItemsCreate**](KVItemsAPI.md#KVItemsCreate) | **Post** /organizations/{organization}/projects/{project}/kv/{store}/items | 
[**KVItemsDelete**](KVItemsAPI.md#KVItemsDelete) | **Delete** /organizations/{organization}/projects/{project}/kv/{store}/items/{item} | 
[**KVItemsList**](KVItemsAPI.md#KVItemsList) | **Get** /organizations/{organization}/projects/{project}/kv/{store}/items | 
[**KVItemsRead**](KVItemsAPI.md#KVItemsRead) | **Get** /organizations/{organization}/projects/{project}/kv/{store}/items/{item} | 
[**KVItemsUpdate**](KVItemsAPI.md#KVItemsUpdate) | **Patch** /organizations/{organization}/projects/{project}/kv/{store}/items/{item} | 



## KVItemsCreate

> StoreItem KVItemsCreate(ctx, organization, project, store).StoreItem(storeItem).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	store := "store_example" // string | 
	storeItem := *openapiclient.NewStoreItem("Key_example", "Value_example") // StoreItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVItemsAPI.KVItemsCreate(context.Background(), organization, project, store).StoreItem(storeItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVItemsAPI.KVItemsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsCreate`: StoreItem
	fmt.Fprintf(os.Stdout, "Response from `KVItemsAPI.KVItemsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**store** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storeItem** | [**StoreItem**](StoreItem.md) |  | 

### Return type

[**StoreItem**](StoreItem.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsDelete

> StoreItem KVItemsDelete(ctx, organization, project, store, item).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	store := "store_example" // string | 
	item := "item_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVItemsAPI.KVItemsDelete(context.Background(), organization, project, store, item).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVItemsAPI.KVItemsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsDelete`: StoreItem
	fmt.Fprintf(os.Stdout, "Response from `KVItemsAPI.KVItemsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**store** | **string** |  | 
**item** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StoreItem**](StoreItem.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsList

> []StoreItem KVItemsList(ctx, organization, project, store).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	store := "store_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVItemsAPI.KVItemsList(context.Background(), organization, project, store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVItemsAPI.KVItemsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsList`: []StoreItem
	fmt.Fprintf(os.Stdout, "Response from `KVItemsAPI.KVItemsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**store** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]StoreItem**](StoreItem.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsRead

> StoreItem KVItemsRead(ctx, organization, project, store, item).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	store := "store_example" // string | 
	item := "item_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVItemsAPI.KVItemsRead(context.Background(), organization, project, store, item).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVItemsAPI.KVItemsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsRead`: StoreItem
	fmt.Fprintf(os.Stdout, "Response from `KVItemsAPI.KVItemsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**store** | **string** |  | 
**item** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StoreItem**](StoreItem.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsUpdate

> StoreItem KVItemsUpdate(ctx, organization, project, store, item).StoreItemUpdate(storeItemUpdate).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	store := "store_example" // string | 
	item := "item_example" // string | 
	storeItemUpdate := *openapiclient.NewStoreItemUpdate() // StoreItemUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVItemsAPI.KVItemsUpdate(context.Background(), organization, project, store, item).StoreItemUpdate(storeItemUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVItemsAPI.KVItemsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsUpdate`: StoreItem
	fmt.Fprintf(os.Stdout, "Response from `KVItemsAPI.KVItemsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**store** | **string** |  | 
**item** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **storeItemUpdate** | [**StoreItemUpdate**](StoreItemUpdate.md) |  | 

### Return type

[**StoreItem**](StoreItem.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

