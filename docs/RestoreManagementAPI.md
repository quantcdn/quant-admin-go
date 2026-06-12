# \RestoreManagementAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRestoreStatus**](RestoreManagementAPI.md#GetRestoreStatus) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/restores/{restoreId} | Get the status of a restore operation
[**RestoreDatabase**](RestoreManagementAPI.md#RestoreDatabase) | **Post** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/backups/{backupId}/restore-database | Restore a database backup to a target environment
[**RestoreFilesystem**](RestoreManagementAPI.md#RestoreFilesystem) | **Post** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/backups/{backupId}/restore-filesystem | Restore a filesystem backup to a target environment



## GetRestoreStatus

> GetRestoreStatus200Response GetRestoreStatus(ctx, organisation, application, environment, restoreId).Execute()

Get the status of a restore operation



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
	environment := "staging" // string | The environment ID
	restoreId := "restore-abc123" // string | The restore operation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestoreManagementAPI.GetRestoreStatus(context.Background(), organisation, application, environment, restoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestoreManagementAPI.GetRestoreStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestoreStatus`: GetRestoreStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `RestoreManagementAPI.GetRestoreStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**restoreId** | **string** | The restore operation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetRestoreStatus200Response**](GetRestoreStatus200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDatabase

> RestoreDatabase202Response RestoreDatabase(ctx, organisation, application, environment, backupId).RestoreDatabaseRequest(restoreDatabaseRequest).Execute()

Restore a database backup to a target environment



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
	environment := "staging" // string | The TARGET environment ID to restore INTO
	backupId := "quant-gov-dashboard-staging-2025-09-19T21-50-27-145Z" // string | The backup ID to restore from
	restoreDatabaseRequest := *openapiclient.NewRestoreDatabaseRequest("BackupId_example", true) // RestoreDatabaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestoreManagementAPI.RestoreDatabase(context.Background(), organisation, application, environment, backupId).RestoreDatabaseRequest(restoreDatabaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestoreManagementAPI.RestoreDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreDatabase`: RestoreDatabase202Response
	fmt.Fprintf(os.Stdout, "Response from `RestoreManagementAPI.RestoreDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The TARGET environment ID to restore INTO | 
**backupId** | **string** | The backup ID to restore from | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restoreDatabaseRequest** | [**RestoreDatabaseRequest**](RestoreDatabaseRequest.md) |  | 

### Return type

[**RestoreDatabase202Response**](RestoreDatabase202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFilesystem

> RestoreFilesystem202Response RestoreFilesystem(ctx, organisation, application, environment, backupId).RestoreFilesystemRequest(restoreFilesystemRequest).Execute()

Restore a filesystem backup to a target environment



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
	environment := "staging" // string | The TARGET environment ID to restore INTO
	backupId := "quant-gov-dashboard-staging-2025-09-19T21-50-27-145Z" // string | The backup ID to restore from
	restoreFilesystemRequest := *openapiclient.NewRestoreFilesystemRequest("BackupId_example", false) // RestoreFilesystemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RestoreManagementAPI.RestoreFilesystem(context.Background(), organisation, application, environment, backupId).RestoreFilesystemRequest(restoreFilesystemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RestoreManagementAPI.RestoreFilesystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreFilesystem`: RestoreFilesystem202Response
	fmt.Fprintf(os.Stdout, "Response from `RestoreManagementAPI.RestoreFilesystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The TARGET environment ID to restore INTO | 
**backupId** | **string** | The backup ID to restore from | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFilesystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restoreFilesystemRequest** | [**RestoreFilesystemRequest**](RestoreFilesystemRequest.md) |  | 

### Return type

[**RestoreFilesystem202Response**](RestoreFilesystem202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

