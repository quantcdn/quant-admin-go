# \RulesBotChallengeAPI

All URIs are relative to *https://dashboard.quantcdn.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RulesBotChallengeCreate**](RulesBotChallengeAPI.md#RulesBotChallengeCreate) | **Post** /organizations/{organization}/projects/{project}/rules/bot-challenge | 
[**RulesBotChallengeDelete**](RulesBotChallengeAPI.md#RulesBotChallengeDelete) | **Delete** /organizations/{organization}/projects/{project}/rules/bot-challenge/{rule} | 
[**RulesBotChallengeList**](RulesBotChallengeAPI.md#RulesBotChallengeList) | **Get** /organizations/{organization}/projects/{project}/rules/bot-challenge | 
[**RulesBotChallengeRead**](RulesBotChallengeAPI.md#RulesBotChallengeRead) | **Get** /organizations/{organization}/projects/{project}/rules/bot-challenge/{rule} | 
[**RulesBotChallengeUpdate**](RulesBotChallengeAPI.md#RulesBotChallengeUpdate) | **Patch** /organizations/{organization}/projects/{project}/rules/bot-challenge/{rule} | 



## RulesBotChallengeCreate

> RuleBotChallenge RulesBotChallengeCreate(ctx, organization, project).RuleBotChallengeRequest(ruleBotChallengeRequest).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	ruleBotChallengeRequest := *openapiclient.NewRuleBotChallengeRequest([]string{"Domain_example"}, false, []string{"Url_example"}, "RobotChallengeType_example", int32(123), int32(123)) // RuleBotChallengeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesBotChallengeAPI.RulesBotChallengeCreate(context.Background(), organization, project).RuleBotChallengeRequest(ruleBotChallengeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesBotChallengeAPI.RulesBotChallengeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeCreate`: RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesBotChallengeAPI.RulesBotChallengeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ruleBotChallengeRequest** | [**RuleBotChallengeRequest**](RuleBotChallengeRequest.md) |  | 

### Return type

[**RuleBotChallenge**](RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeDelete

> RuleBotChallenge RulesBotChallengeDelete(ctx, organization, project, rule).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesBotChallengeAPI.RulesBotChallengeDelete(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesBotChallengeAPI.RulesBotChallengeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeDelete`: RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesBotChallengeAPI.RulesBotChallengeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleBotChallenge**](RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeList

> []RuleBotChallenge RulesBotChallengeList(ctx, organization, project).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesBotChallengeAPI.RulesBotChallengeList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesBotChallengeAPI.RulesBotChallengeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeList`: []RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesBotChallengeAPI.RulesBotChallengeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RuleBotChallenge**](RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeRead

> RuleBotChallenge RulesBotChallengeRead(ctx, organization, project, rule).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesBotChallengeAPI.RulesBotChallengeRead(context.Background(), organization, project, rule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesBotChallengeAPI.RulesBotChallengeRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeRead`: RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesBotChallengeAPI.RulesBotChallengeRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RuleBotChallenge**](RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RulesBotChallengeUpdate

> RuleBotChallenge RulesBotChallengeUpdate(ctx, organization, project, rule).RuleBotChallengeRequestUpdate(ruleBotChallengeRequestUpdate).Execute()



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
	organization := "organization_example" // string | 
	project := "project_example" // string | 
	rule := "rule_example" // string | 
	ruleBotChallengeRequestUpdate := *openapiclient.NewRuleBotChallengeRequestUpdate() // RuleBotChallengeRequestUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesBotChallengeAPI.RulesBotChallengeUpdate(context.Background(), organization, project, rule).RuleBotChallengeRequestUpdate(ruleBotChallengeRequestUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesBotChallengeAPI.RulesBotChallengeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RulesBotChallengeUpdate`: RuleBotChallenge
	fmt.Fprintf(os.Stdout, "Response from `RulesBotChallengeAPI.RulesBotChallengeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** |  | 
**project** | **string** |  | 
**rule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRulesBotChallengeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ruleBotChallengeRequestUpdate** | [**RuleBotChallengeRequestUpdate**](RuleBotChallengeRequestUpdate.md) |  | 

### Return type

[**RuleBotChallenge**](RuleBotChallenge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

