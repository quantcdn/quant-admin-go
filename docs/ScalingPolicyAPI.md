# \ScalingPolicyAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteScalingPolicy**](ScalingPolicyAPI.md#DeleteScalingPolicy) | **Delete** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/scaling-policies/{policyName} | Delete the scaling policy for an environment
[**GetScalingPolicies**](ScalingPolicyAPI.md#GetScalingPolicies) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/scaling-policies | Get the scaling policies for an environment
[**UpdateScalingPolicy**](ScalingPolicyAPI.md#UpdateScalingPolicy) | **Put** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/scaling-policies | Update the scaling policy for an environment



## DeleteScalingPolicy

> DeleteScalingPolicy(ctx, organisation, application, environment, policyName).Execute()

Delete the scaling policy for an environment

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
	policyName := "policyName_example" // string | The policy name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScalingPolicyAPI.DeleteScalingPolicy(context.Background(), organisation, application, environment, policyName).Execute()
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
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**policyName** | **string** | The policy name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScalingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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


## GetScalingPolicies

> GetScalingPolicies(ctx, organisation, application, environment).Execute()

Get the scaling policies for an environment

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScalingPolicyAPI.GetScalingPolicies(context.Background(), organisation, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingPolicyAPI.GetScalingPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScalingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## UpdateScalingPolicy

> UpdateScalingPolicy(ctx, organisation, application, environment).ScalingPolicy(scalingPolicy).Execute()

Update the scaling policy for an environment

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
	scalingPolicy := *openapiclient.NewScalingPolicy() // ScalingPolicy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScalingPolicyAPI.UpdateScalingPolicy(context.Background(), organisation, application, environment).ScalingPolicy(scalingPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScalingPolicyAPI.UpdateScalingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScalingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scalingPolicy** | [**ScalingPolicy**](ScalingPolicy.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

