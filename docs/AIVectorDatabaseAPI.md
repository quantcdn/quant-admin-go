# \AIVectorDatabaseAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVectorCollection**](AIVectorDatabaseAPI.md#CreateVectorCollection) | **Post** /api/v3/organizations/{organisation}/ai/vector-db/collections | Create Vector Database Collection
[**DeleteVectorCollection**](AIVectorDatabaseAPI.md#DeleteVectorCollection) | **Delete** /api/v3/organizations/{organisation}/ai/vector-db/collections/{collectionId} | Delete Collection
[**DeleteVectorDocuments**](AIVectorDatabaseAPI.md#DeleteVectorDocuments) | **Delete** /api/v3/organizations/{organisation}/ai/vector-db/collections/{collectionId}/documents | Delete Documents from Collection
[**GetVectorCollection**](AIVectorDatabaseAPI.md#GetVectorCollection) | **Get** /api/v3/organizations/{organisation}/ai/vector-db/collections/{collectionId} | Get Collection Details
[**ListVectorCollections**](AIVectorDatabaseAPI.md#ListVectorCollections) | **Get** /api/v3/organizations/{organisation}/ai/vector-db/collections | List Vector Database Collections
[**QueryVectorCollection**](AIVectorDatabaseAPI.md#QueryVectorCollection) | **Post** /api/v3/organizations/{organisation}/ai/vector-db/collections/{collectionId}/query | Semantic Search Query
[**UploadVectorDocuments**](AIVectorDatabaseAPI.md#UploadVectorDocuments) | **Post** /api/v3/organizations/{organisation}/ai/vector-db/collections/{collectionId}/documents | Upload Documents to Collection



## CreateVectorCollection

> CreateVectorCollection201Response CreateVectorCollection(ctx, organisation).CreateVectorCollectionRequest(createVectorCollectionRequest).Execute()

Create Vector Database Collection



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
	createVectorCollectionRequest := *openapiclient.NewCreateVectorCollectionRequest("product-documentation") // CreateVectorCollectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIVectorDatabaseAPI.CreateVectorCollection(context.Background(), organisation).CreateVectorCollectionRequest(createVectorCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorDatabaseAPI.CreateVectorCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVectorCollection`: CreateVectorCollection201Response
	fmt.Fprintf(os.Stdout, "Response from `AIVectorDatabaseAPI.CreateVectorCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVectorCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVectorCollectionRequest** | [**CreateVectorCollectionRequest**](CreateVectorCollectionRequest.md) |  | 

### Return type

[**CreateVectorCollection201Response**](CreateVectorCollection201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVectorCollection

> DeleteVectorCollection200Response DeleteVectorCollection(ctx, organisation, collectionId).Execute()

Delete Collection



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The collection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIVectorDatabaseAPI.DeleteVectorCollection(context.Background(), organisation, collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorDatabaseAPI.DeleteVectorCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVectorCollection`: DeleteVectorCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `AIVectorDatabaseAPI.DeleteVectorCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**collectionId** | **string** | The collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVectorCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteVectorCollection200Response**](DeleteVectorCollection200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVectorDocuments

> DeleteVectorDocuments200Response DeleteVectorDocuments(ctx, organisation, collectionId).DeleteVectorDocumentsRequest(deleteVectorDocumentsRequest).Execute()

Delete Documents from Collection



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
	organisation := "organisation_example" // string | Organisation machine name
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Collection UUID
	deleteVectorDocumentsRequest := *openapiclient.NewDeleteVectorDocumentsRequest() // DeleteVectorDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIVectorDatabaseAPI.DeleteVectorDocuments(context.Background(), organisation, collectionId).DeleteVectorDocumentsRequest(deleteVectorDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorDatabaseAPI.DeleteVectorDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVectorDocuments`: DeleteVectorDocuments200Response
	fmt.Fprintf(os.Stdout, "Response from `AIVectorDatabaseAPI.DeleteVectorDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | Organisation machine name | 
**collectionId** | **string** | Collection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVectorDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteVectorDocumentsRequest** | [**DeleteVectorDocumentsRequest**](DeleteVectorDocumentsRequest.md) |  | 

### Return type

[**DeleteVectorDocuments200Response**](DeleteVectorDocuments200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVectorCollection

> GetVectorCollection200Response GetVectorCollection(ctx, organisation, collectionId).Execute()

Get Collection Details



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The collection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIVectorDatabaseAPI.GetVectorCollection(context.Background(), organisation, collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorDatabaseAPI.GetVectorCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVectorCollection`: GetVectorCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `AIVectorDatabaseAPI.GetVectorCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**collectionId** | **string** | The collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVectorCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVectorCollection200Response**](GetVectorCollection200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVectorCollections

> ListVectorCollections200Response ListVectorCollections(ctx, organisation).Execute()

List Vector Database Collections



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIVectorDatabaseAPI.ListVectorCollections(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorDatabaseAPI.ListVectorCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVectorCollections`: ListVectorCollections200Response
	fmt.Fprintf(os.Stdout, "Response from `AIVectorDatabaseAPI.ListVectorCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVectorCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListVectorCollections200Response**](ListVectorCollections200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryVectorCollection

> QueryVectorCollection200Response QueryVectorCollection(ctx, organisation, collectionId).QueryVectorCollectionRequest(queryVectorCollectionRequest).Execute()

Semantic Search Query



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The collection ID
	queryVectorCollectionRequest := *openapiclient.NewQueryVectorCollectionRequest() // QueryVectorCollectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIVectorDatabaseAPI.QueryVectorCollection(context.Background(), organisation, collectionId).QueryVectorCollectionRequest(queryVectorCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorDatabaseAPI.QueryVectorCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryVectorCollection`: QueryVectorCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `AIVectorDatabaseAPI.QueryVectorCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**collectionId** | **string** | The collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryVectorCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryVectorCollectionRequest** | [**QueryVectorCollectionRequest**](QueryVectorCollectionRequest.md) |  | 

### Return type

[**QueryVectorCollection200Response**](QueryVectorCollection200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVectorDocuments

> UploadVectorDocuments200Response UploadVectorDocuments(ctx, organisation, collectionId).UploadVectorDocumentsRequest(uploadVectorDocumentsRequest).Execute()

Upload Documents to Collection



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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The collection ID
	uploadVectorDocumentsRequest := *openapiclient.NewUploadVectorDocumentsRequest([]openapiclient.UploadVectorDocumentsRequestDocumentsInner{*openapiclient.NewUploadVectorDocumentsRequestDocumentsInner("Content_example")}) // UploadVectorDocumentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIVectorDatabaseAPI.UploadVectorDocuments(context.Background(), organisation, collectionId).UploadVectorDocumentsRequest(uploadVectorDocumentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorDatabaseAPI.UploadVectorDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadVectorDocuments`: UploadVectorDocuments200Response
	fmt.Fprintf(os.Stdout, "Response from `AIVectorDatabaseAPI.UploadVectorDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**collectionId** | **string** | The collection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadVectorDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadVectorDocumentsRequest** | [**UploadVectorDocumentsRequest**](UploadVectorDocumentsRequest.md) |  | 

### Return type

[**UploadVectorDocuments200Response**](UploadVectorDocuments200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

