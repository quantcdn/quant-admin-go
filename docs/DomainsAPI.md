# \DomainsAPI

All URIs are relative to *http://localhost:8001/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DomainsAPI.md#CreateDomain) | **Post** /organizations/{organization}/projects/{project}/domains | 
[**DeleteDomain**](DomainsAPI.md#DeleteDomain) | **Delete** /organizations/{organization}/projects/{project}/domains/{domain} | 
[**GetDomain**](DomainsAPI.md#GetDomain) | **Get** /organizations/{organization}/projects/{project}/domains/{domain} | 
[**ListDomains**](DomainsAPI.md#ListDomains) | **Get** /organizations/{organization}/projects/{project}/domains | 
[**UpdateDomain**](DomainsAPI.md#UpdateDomain) | **Patch** /organizations/{organization}/projects/{project}/domains/{domain} | 



## CreateDomain

> Domain CreateDomain(ctx, organization, project).DomainRequest(domainRequest).Execute()





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
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name
	domainRequest := *openapiclient.NewDomainRequest() // DomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.CreateDomain(context.Background(), organization, project).DomainRequest(domainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.CreateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomain`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.CreateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domainRequest** | [**DomainRequest**](DomainRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDomain

> Domain DeleteDomain(ctx, organization, project, domain).Execute()





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
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DeleteDomain(context.Background(), organization, project, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDomain`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DeleteDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Domain**](Domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomain

> Domain GetDomain(ctx, organization, project, domain).Execute()





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
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name
	domain := "domain_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetDomain(context.Background(), organization, project, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomain`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Domain**](Domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomains

> []Domain ListDomains(ctx, organization, project).Execute()





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
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.ListDomains(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.ListDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDomains`: []Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.ListDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Domain**](Domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomain

> Domain UpdateDomain(ctx, organization, project, domain).DomainRequest(domainRequest).Execute()





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
	organization := "organization_example" // string | Organization machine name
	project := "project_example" // string | Project machine name
	domain := "domain_example" // string | 
	domainRequest := *openapiclient.NewDomainRequest() // DomainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.UpdateDomain(context.Background(), organization, project, domain).DomainRequest(domainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.UpdateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDomain`: Domain
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.UpdateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization machine name | 
**project** | **string** | Project machine name | 
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **domainRequest** | [**DomainRequest**](DomainRequest.md) |  | 

### Return type

[**Domain**](Domain.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

