# \KVAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KVCreate**](KVAPI.md#KVCreate) | **Post** /organizations/{organization}/projects/{project}/kv | 
[**KVDelete**](KVAPI.md#KVDelete) | **Delete** /organizations/{organization}/projects/{project}/kv/{store} | 
[**KVList**](KVAPI.md#KVList) | **Get** /organizations/{organization}/projects/{project}/kv | 
[**KVRead**](KVAPI.md#KVRead) | **Get** /organizations/{organization}/projects/{project}/kv/{store} | 
[**KVUpdate**](KVAPI.md#KVUpdate) | **Patch** /organizations/{organization}/projects/{project}/kv/{store} | 



## KVCreate

> Store KVCreate(ctx, organization, project).Store(store).Execute()



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
	store := *openapiclient.NewStore("Name_example") // Store | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVCreate(context.Background(), organization, project).Store(store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVCreate`: Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **store** | [**Store**](Store.md) |  | 

### Return type

[**Store**](Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVDelete

> Store KVDelete(ctx, organization, project, store).Execute()



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
	resp, r, err := apiClient.KVAPI.KVDelete(context.Background(), organization, project, store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVDelete`: Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiKVDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Store**](Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVList

> []Store KVList(ctx, organization, project).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVList`: []Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Store**](Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVRead

> Store KVRead(ctx, organization, project, store).Execute()



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
	resp, r, err := apiClient.KVAPI.KVRead(context.Background(), organization, project, store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVRead`: Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVRead`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiKVReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Store**](Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVUpdate

> Store KVUpdate(ctx, organization, project, store).StoreUpdate(storeUpdate).Execute()



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
	storeUpdate := *openapiclient.NewStoreUpdate() // StoreUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVUpdate(context.Background(), organization, project, store).StoreUpdate(storeUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVUpdate`: Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiKVUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **storeUpdate** | [**StoreUpdate**](StoreUpdate.md) |  | 

### Return type

[**Store**](Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

