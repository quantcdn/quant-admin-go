# \CronAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCronJob**](CronAPI.md#CreateCronJob) | **Post** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/cron | Create a new cron job
[**DeleteCronJob**](CronAPI.md#DeleteCronJob) | **Delete** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/cron/{cron} | Delete a cron job
[**GetCronJob**](CronAPI.md#GetCronJob) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/cron/{cron} | Get a cron job
[**GetCronRun**](CronAPI.md#GetCronRun) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/cron/{cron}/runs/{run} | Get a cron run
[**ListCronJobRuns**](CronAPI.md#ListCronJobRuns) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/cron/{cron}/runs | Get all runs for a cron job
[**ListCronJobs**](CronAPI.md#ListCronJobs) | **Get** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/cron | Get all cron jobs for an environment
[**UpdateCronJob**](CronAPI.md#UpdateCronJob) | **Patch** /api/v3/organizations/{organisation}/applications/{application}/environments/{environment}/cron/{cron} | Update a cron job



## CreateCronJob

> Cron CreateCronJob(ctx, organisation, application, environment).CreateCronJobRequest(createCronJobRequest).Execute()

Create a new cron job

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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	createCronJobRequest := *openapiclient.NewCreateCronJobRequest("Name_example", "ScheduleExpression_example", []string{"Command_example"}) // CreateCronJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronAPI.CreateCronJob(context.Background(), organisation, application, environment).CreateCronJobRequest(createCronJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.CreateCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCronJob`: Cron
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.CreateCronJob`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createCronJobRequest** | [**CreateCronJobRequest**](CreateCronJobRequest.md) |  | 

### Return type

[**Cron**](Cron.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCronJob

> DeleteCronJob(ctx, organisation, application, environment, cron).Execute()

Delete a cron job

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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	cron := "cron_example" // string | The cron job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronAPI.DeleteCronJob(context.Background(), organisation, application, environment, cron).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.DeleteCronJob``: %v\n", err)
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
**cron** | **string** | The cron job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCronJobRequest struct via the builder pattern


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


## GetCronJob

> Cron GetCronJob(ctx, organisation, application, environment, cron).Execute()

Get a cron job

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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	cron := "cron_example" // string | The cron job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronAPI.GetCronJob(context.Background(), organisation, application, environment, cron).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.GetCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCronJob`: Cron
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.GetCronJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**cron** | **string** | The cron job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Cron**](Cron.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCronRun

> CronRun GetCronRun(ctx, organisation, application, environment, cron, run).Execute()

Get a cron run

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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	cron := "cron_example" // string | The cron job ID
	run := "run_example" // string | The cron run ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronAPI.GetCronRun(context.Background(), organisation, application, environment, cron, run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.GetCronRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCronRun`: CronRun
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.GetCronRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**cron** | **string** | The cron job ID | 
**run** | **string** | The cron run ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCronRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**CronRun**](CronRun.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCronJobRuns

> []CronRun ListCronJobRuns(ctx, organisation, application, environment, cron).Execute()

Get all runs for a cron job

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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	cron := "cron_example" // string | The cron job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronAPI.ListCronJobRuns(context.Background(), organisation, application, environment, cron).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.ListCronJobRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCronJobRuns`: []CronRun
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.ListCronJobRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**cron** | **string** | The cron job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCronJobRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]CronRun**](CronRun.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCronJobs

> Cron ListCronJobs(ctx, organisation, application, environment).Execute()

Get all cron jobs for an environment

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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronAPI.ListCronJobs(context.Background(), organisation, application, environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.ListCronJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCronJobs`: Cron
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.ListCronJobs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListCronJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Cron**](Cron.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCronJob

> Cron UpdateCronJob(ctx, organisation, application, environment, cron).UpdateCronJobRequest(updateCronJobRequest).Execute()

Update a cron job

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
	application := "application_example" // string | The application ID
	environment := "environment_example" // string | The environment ID
	cron := "cron_example" // string | The cron job ID
	updateCronJobRequest := *openapiclient.NewUpdateCronJobRequest() // UpdateCronJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronAPI.UpdateCronJob(context.Background(), organisation, application, environment, cron).UpdateCronJobRequest(updateCronJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.UpdateCronJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCronJob`: Cron
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.UpdateCronJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**application** | **string** | The application ID | 
**environment** | **string** | The environment ID | 
**cron** | **string** | The cron job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCronJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateCronJobRequest** | [**UpdateCronJobRequest**](UpdateCronJobRequest.md) |  | 

### Return type

[**Cron**](Cron.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

