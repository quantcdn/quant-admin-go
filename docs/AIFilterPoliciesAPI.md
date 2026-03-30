# \AIFilterPoliciesAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFilterPolicy**](AIFilterPoliciesAPI.md#CreateFilterPolicy) | **Post** /api/v3/organizations/{organisation}/ai/filter-policies | Create an AI filter policy for an organisation
[**DeleteFilterPolicy**](AIFilterPoliciesAPI.md#DeleteFilterPolicy) | **Delete** /api/v3/organizations/{organisation}/ai/filter-policies/{policyId} | Delete a specific AI filter policy
[**DisableFilterPolicy**](AIFilterPoliciesAPI.md#DisableFilterPolicy) | **Put** /api/v3/organizations/{organisation}/ai/filter-policies/{policyId}/disable | Disable a specific AI filter policy
[**EnableFilterPolicy**](AIFilterPoliciesAPI.md#EnableFilterPolicy) | **Put** /api/v3/organizations/{organisation}/ai/filter-policies/{policyId}/enable | Enable a specific AI filter policy
[**GetFilterPolicy**](AIFilterPoliciesAPI.md#GetFilterPolicy) | **Get** /api/v3/organizations/{organisation}/ai/filter-policies/{policyId} | Get a specific AI filter policy
[**ListFilterPolicies**](AIFilterPoliciesAPI.md#ListFilterPolicies) | **Get** /api/v3/organizations/{organisation}/ai/filter-policies | List AI filter policies for an organisation
[**UpdateFilterPolicy**](AIFilterPoliciesAPI.md#UpdateFilterPolicy) | **Put** /api/v3/organizations/{organisation}/ai/filter-policies/{policyId} | Update a specific AI filter policy



## CreateFilterPolicy

> map[string]interface{} CreateFilterPolicy(ctx, organisation).CreateFilterPolicyRequest(createFilterPolicyRequest).Execute()

Create an AI filter policy for an organisation

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
	createFilterPolicyRequest := *openapiclient.NewCreateFilterPolicyRequest("Name_example", []openapiclient.CreateFilterPolicyRequestRulesInner{*openapiclient.NewCreateFilterPolicyRequestRulesInner()}) // CreateFilterPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFilterPoliciesAPI.CreateFilterPolicy(context.Background(), organisation).CreateFilterPolicyRequest(createFilterPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFilterPoliciesAPI.CreateFilterPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFilterPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIFilterPoliciesAPI.CreateFilterPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilterPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFilterPolicyRequest** | [**CreateFilterPolicyRequest**](CreateFilterPolicyRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFilterPolicy

> map[string]interface{} DeleteFilterPolicy(ctx, organisation, policyId).Execute()

Delete a specific AI filter policy

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
	organisation := "organisation_example" // string | 
	policyId := "policyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFilterPoliciesAPI.DeleteFilterPolicy(context.Background(), organisation, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFilterPoliciesAPI.DeleteFilterPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFilterPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIFilterPoliciesAPI.DeleteFilterPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFilterPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableFilterPolicy

> map[string]interface{} DisableFilterPolicy(ctx, organisation, policyId).Execute()

Disable a specific AI filter policy

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
	organisation := "organisation_example" // string | 
	policyId := "policyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFilterPoliciesAPI.DisableFilterPolicy(context.Background(), organisation, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFilterPoliciesAPI.DisableFilterPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableFilterPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIFilterPoliciesAPI.DisableFilterPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableFilterPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableFilterPolicy

> map[string]interface{} EnableFilterPolicy(ctx, organisation, policyId).Execute()

Enable a specific AI filter policy

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
	organisation := "organisation_example" // string | 
	policyId := "policyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFilterPoliciesAPI.EnableFilterPolicy(context.Background(), organisation, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFilterPoliciesAPI.EnableFilterPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableFilterPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIFilterPoliciesAPI.EnableFilterPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableFilterPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilterPolicy

> map[string]interface{} GetFilterPolicy(ctx, organisation, policyId).Execute()

Get a specific AI filter policy

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
	organisation := "organisation_example" // string | 
	policyId := "policyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFilterPoliciesAPI.GetFilterPolicy(context.Background(), organisation, policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFilterPoliciesAPI.GetFilterPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilterPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIFilterPoliciesAPI.GetFilterPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilterPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFilterPolicies

> map[string]interface{} ListFilterPolicies(ctx, organisation).Execute()

List AI filter policies for an organisation

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
	resp, r, err := apiClient.AIFilterPoliciesAPI.ListFilterPolicies(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFilterPoliciesAPI.ListFilterPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFilterPolicies`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIFilterPoliciesAPI.ListFilterPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFilterPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFilterPolicy

> map[string]interface{} UpdateFilterPolicy(ctx, organisation, policyId).UpdateFilterPolicyRequest(updateFilterPolicyRequest).Execute()

Update a specific AI filter policy

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
	organisation := "organisation_example" // string | 
	policyId := "policyId_example" // string | 
	updateFilterPolicyRequest := *openapiclient.NewUpdateFilterPolicyRequest() // UpdateFilterPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIFilterPoliciesAPI.UpdateFilterPolicy(context.Background(), organisation, policyId).UpdateFilterPolicyRequest(updateFilterPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIFilterPoliciesAPI.UpdateFilterPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFilterPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIFilterPoliciesAPI.UpdateFilterPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFilterPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFilterPolicyRequest** | [**UpdateFilterPolicyRequest**](UpdateFilterPolicyRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

