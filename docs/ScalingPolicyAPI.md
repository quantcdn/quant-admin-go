# \ScalingPolicyAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteScalingPolicy**](ScalingPolicyAPI.md#DeleteScalingPolicy) | **Delete** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/scaling-policies | Delete Scaling Policy
[**ListScalingPolicies**](ScalingPolicyAPI.md#ListScalingPolicies) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/scaling-policies | List Scaling Policies
[**UpsertScalingPolicy**](ScalingPolicyAPI.md#UpsertScalingPolicy) | **Put** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/scaling-policies | Upsert Scaling Policy



## DeleteScalingPolicy

> DeleteScalingPolicy(ctx, organisation, application, environment).Metric(metric).PolicyName(policyName).Execute()

Delete Scaling Policy



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
	application := "application_example" // string | 
	environment := "environment_example" // string | 
	metric := "metric_example" // string | Optional. Delete by metric type. (optional)
	policyName := "policyName_example" // string | Optional. Delete by exact policy name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScalingPolicyAPI.DeleteScalingPolicy(context.Background(), organisation, application, environment).Metric(metric).PolicyName(policyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingPolicyAPI.DeleteScalingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**application** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScalingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metric** | **string** | Optional. Delete by metric type. | 
 **policyName** | **string** | Optional. Delete by exact policy name. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScalingPolicies

> ScalingPolicyListResponse ListScalingPolicies(ctx, organisation, application, environment).Metric(metric).PolicyName(policyName).Execute()

List Scaling Policies



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
	application := "application_example" // string | 
	environment := "environment_example" // string | 
	metric := "metric_example" // string | Optional. Filter policies by metric type. (optional)
	policyName := "policyName_example" // string | Optional. Filter policies by exact policy name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScalingPolicyAPI.ListScalingPolicies(context.Background(), organisation, application, environment).Metric(metric).PolicyName(policyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingPolicyAPI.ListScalingPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListScalingPolicies`: ScalingPolicyListResponse
	fmt.Fprintf(os.Stdout, "Response from `ScalingPolicyAPI.ListScalingPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**application** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScalingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metric** | **string** | Optional. Filter policies by metric type. | 
 **policyName** | **string** | Optional. Filter policies by exact policy name. | 

### Return type

[**ScalingPolicyListResponse**](ScalingPolicyListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertScalingPolicy

> GetScalingPolicyResponse UpsertScalingPolicy(ctx, organisation, application, environment).SetScalingPolicyRequest(setScalingPolicyRequest).PolicyName(policyName).Execute()

Upsert Scaling Policy



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
	application := "application_example" // string | 
	environment := "environment_example" // string | 
	setScalingPolicyRequest := *openapiclient.NewSetScalingPolicyRequest("Metric_example", float64(75)) // SetScalingPolicyRequest | 
	policyName := "policyName_example" // string | Optional. Specify a custom policy name to upsert. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScalingPolicyAPI.UpsertScalingPolicy(context.Background(), organisation, application, environment).SetScalingPolicyRequest(setScalingPolicyRequest).PolicyName(policyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingPolicyAPI.UpsertScalingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertScalingPolicy`: GetScalingPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ScalingPolicyAPI.UpsertScalingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**application** | **string** |  | 
**environment** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertScalingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **setScalingPolicyRequest** | [**SetScalingPolicyRequest**](SetScalingPolicyRequest.md) |  | 
 **policyName** | **string** | Optional. Specify a custom policy name to upsert. | 

### Return type

[**GetScalingPolicyResponse**](GetScalingPolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

