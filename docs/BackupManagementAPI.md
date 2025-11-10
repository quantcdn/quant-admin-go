# \BackupManagementAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](BackupManagementAPI.md#CreateBackup) | **Post** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/backups/{type} | Create a backup for an environment
[**DeleteBackup**](BackupManagementAPI.md#DeleteBackup) | **Delete** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/backups/{type}/{backupId} | Delete a backup
[**DownloadBackup**](BackupManagementAPI.md#DownloadBackup) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/backups/{type}/{backupId}/download | Generate a download URL for a backup
[**ListBackups**](BackupManagementAPI.md#ListBackups) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/backups/{type} | List backups for an environment



## CreateBackup

> CreateBackup202Response CreateBackup(ctx, organisation, application, environment, type_).CreateBackupRequest(createBackupRequest).Execute()

Create a backup for an environment

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
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	type_ := "type__example" // string | The backup type
	createBackupRequest := *openapiclient.NewCreateBackupRequest() // CreateBackupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupManagementAPI.CreateBackup(context.Background(), organisation, application, environment, type_).CreateBackupRequest(createBackupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupManagementAPI.CreateBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackup`: CreateBackup202Response
	fmt.Fprintf(os.Stdout, "Response from `BackupManagementAPI.CreateBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**type_** | **string** | The backup type | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **createBackupRequest** | [**CreateBackupRequest**](CreateBackupRequest.md) |  | 

### Return type

[**CreateBackup202Response**](CreateBackup202Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackup

> DeleteBackup200Response DeleteBackup(ctx, organisation, application, environment, type_, backupId).Execute()

Delete a backup

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
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	type_ := "type__example" // string | The backup type
	backupId := "backupId_example" // string | The backup ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupManagementAPI.DeleteBackup(context.Background(), organisation, application, environment, type_, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupManagementAPI.DeleteBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackup`: DeleteBackup200Response
	fmt.Fprintf(os.Stdout, "Response from `BackupManagementAPI.DeleteBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**type_** | **string** | The backup type | 
**backupId** | **string** | The backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**DeleteBackup200Response**](DeleteBackup200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadBackup

> DownloadBackup200Response DownloadBackup(ctx, organisation, application, environment, type_, backupId).Execute()

Generate a download URL for a backup

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
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	type_ := "type__example" // string | The backup type
	backupId := "backupId_example" // string | The backup ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupManagementAPI.DownloadBackup(context.Background(), organisation, application, environment, type_, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupManagementAPI.DownloadBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadBackup`: DownloadBackup200Response
	fmt.Fprintf(os.Stdout, "Response from `BackupManagementAPI.DownloadBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**type_** | **string** | The backup type | 
**backupId** | **string** | The backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**DownloadBackup200Response**](DownloadBackup200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackups

> ListBackups200Response ListBackups(ctx, organisation, application, environment, type_).Order(order).Limit(limit).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Status(status).NextToken(nextToken).Execute()

List backups for an environment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organisation := "organisation_example" // string | The organisation ID
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	type_ := "type__example" // string | The backup type
	order := "order_example" // string | Sort order for backups by creation date (asc = oldest first, desc = newest first) (optional) (default to "desc")
	limit := int32(56) // int32 | Maximum number of backups to return (max 100) (optional) (default to 50)
	createdBefore := time.Now() // time.Time | Only return backups created before this ISO 8601 timestamp (e.g., 2025-01-01T00:00:00Z) (optional)
	createdAfter := time.Now() // time.Time | Only return backups created after this ISO 8601 timestamp (e.g., 2024-12-01T00:00:00Z) (optional)
	status := "status_example" // string | Filter backups by status (optional)
	nextToken := "nextToken_example" // string | Token for retrieving the next page of results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupManagementAPI.ListBackups(context.Background(), organisation, application, environment, type_).Order(order).Limit(limit).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Status(status).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupManagementAPI.ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackups`: ListBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `BackupManagementAPI.ListBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**type_** | **string** | The backup type | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **order** | **string** | Sort order for backups by creation date (asc &#x3D; oldest first, desc &#x3D; newest first) | [default to &quot;desc&quot;]
 **limit** | **int32** | Maximum number of backups to return (max 100) | [default to 50]
 **createdBefore** | **time.Time** | Only return backups created before this ISO 8601 timestamp (e.g., 2025-01-01T00:00:00Z) | 
 **createdAfter** | **time.Time** | Only return backups created after this ISO 8601 timestamp (e.g., 2024-12-01T00:00:00Z) | 
 **status** | **string** | Filter backups by status | 
 **nextToken** | **string** | Token for retrieving the next page of results | 

### Return type

[**ListBackups200Response**](ListBackups200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

