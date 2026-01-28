# \AIFileStorageAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFile**](AIFileStorageAPI.md#DeleteFile) | **Delete** /api/v3/organizations/{organisation}/ai/files/{fileId} | Delete File
[**GetFile**](AIFileStorageAPI.md#GetFile) | **Get** /api/v3/organizations/{organisation}/ai/files/{fileId} | Get File
[**ListFiles**](AIFileStorageAPI.md#ListFiles) | **Get** /api/v3/organizations/{organisation}/ai/files | List Files
[**UploadFile**](AIFileStorageAPI.md#UploadFile) | **Post** /api/v3/organizations/{organisation}/ai/files | Upload File to S3



## DeleteFile

> DeleteFile200Response DeleteFile(ctx, organisation, fileId).Execute()

Delete File



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The file ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFileStorageAPI.DeleteFile(context.Background(), organisation, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFileStorageAPI.DeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFile`: DeleteFile200Response
	fmt.Fprintf(os.Stdout, "Response from `AIFileStorageAPI.DeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**fileId** | **string** | The file ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteFile200Response**](DeleteFile200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> GetFile200Response GetFile(ctx, organisation, fileId).Execute()

Get File



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The file ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFileStorageAPI.GetFile(context.Background(), organisation, fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFileStorageAPI.GetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFile`: GetFile200Response
	fmt.Fprintf(os.Stdout, "Response from `AIFileStorageAPI.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**fileId** | **string** | The file ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFile200Response**](GetFile200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFiles

> ListFiles200Response ListFiles(ctx, organisation).Filter(filter).Limit(limit).Cursor(cursor).Execute()

List Files



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
	filter := "{}" // string | JSON-encoded metadata filter. Supports exact match and array contains filters. (optional)
	limit := int32(56) // int32 | Maximum files to return (optional) (default to 50)
	cursor := "cursor_example" // string | Pagination cursor from previous response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFileStorageAPI.ListFiles(context.Background(), organisation).Filter(filter).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFileStorageAPI.ListFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFiles`: ListFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `AIFileStorageAPI.ListFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | JSON-encoded metadata filter. Supports exact match and array contains filters. | 
 **limit** | **int32** | Maximum files to return | [default to 50]
 **cursor** | **string** | Pagination cursor from previous response | 

### Return type

[**ListFiles200Response**](ListFiles200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> UploadFile201Response UploadFile(ctx, organisation).UploadFileRequest(uploadFileRequest).Execute()

Upload File to S3



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
	uploadFileRequest := *openapiclient.NewUploadFileRequest("image/png") // UploadFileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFileStorageAPI.UploadFile(context.Background(), organisation).UploadFileRequest(uploadFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFileStorageAPI.UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFile`: UploadFile201Response
	fmt.Fprintf(os.Stdout, "Response from `AIFileStorageAPI.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadFileRequest** | [**UploadFileRequest**](UploadFileRequest.md) |  | 

### Return type

[**UploadFile201Response**](UploadFile201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

