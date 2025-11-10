# \DomainsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsCreate**](DomainsAPI.md#DomainsCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/domains | Add a new domain
[**DomainsDelete**](DomainsAPI.md#DomainsDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/domains/{domain} | Delete a domain
[**DomainsList**](DomainsAPI.md#DomainsList) | **Get** /api/v2/organizations/{organization}/projects/{project}/domains | List all domains for a project
[**DomainsRead**](DomainsAPI.md#DomainsRead) | **Get** /api/v2/organizations/{organization}/projects/{project}/domains/{domain} | Get details of a single domain
[**DomainsRenew**](DomainsAPI.md#DomainsRenew) | **Post** /api/v2/organizations/{organization}/projects/{project}/domains/{domain}/renew | Renew the SSL certificate for a domain



## DomainsCreate

> V2Domain DomainsCreate(ctx, organization, project).V2DomainRequest(v2DomainRequest).Execute()

Add a new domain

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	v2DomainRequest := *openapiclient.NewV2DomainRequest("The requested resource was not found", true, "test-domain.com") // V2DomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsCreate(context.Background(), organization, project).V2DomainRequest(v2DomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsCreate`: V2Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2DomainRequest** | [**V2DomainRequest**](V2DomainRequest.md) |  | 

### Return type

[**V2Domain**](V2Domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDelete

> DomainsDelete(ctx, organization, project, domain).Execute()

Delete a domain

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.DomainsDelete(context.Background(), organization, project, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsDelete``: %v\n", err)
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
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDeleteRequest struct via the builder pattern


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


## DomainsList

> []V2Domain DomainsList(ctx, organization, project).Execute()

List all domains for a project

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsList`: []V2Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2Domain**](V2Domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsRead

> V2Domain DomainsRead(ctx, organization, project, domain).Execute()

Get details of a single domain

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DomainsRead(context.Background(), organization, project, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsRead`: V2Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DomainsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2Domain**](V2Domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsRenew

> DomainsRenew(ctx, organization, project, domain).Execute()

Renew the SSL certificate for a domain

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
	organization := "organization_example" // string | Organization identifier
	project := "project_example" // string | Project identifier
	domain := "domain_example" // string | Domain identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DomainsAPI.DomainsRenew(context.Background(), organization, project, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DomainsRenew``: %v\n", err)
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
**domain** | **string** | Domain identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsRenewRequest struct via the builder pattern


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

