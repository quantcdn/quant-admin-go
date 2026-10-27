# \AIToolApprovalAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecideToolProposal**](AIToolApprovalAPI.md#DecideToolProposal) | **Post** /api/v3/organizations/{organisation}/ai/tool-proposals/{proposalId}/decision | Decide a tool approval proposal
[**GetToolProposal**](AIToolApprovalAPI.md#GetToolProposal) | **Get** /api/v3/organizations/{organisation}/ai/tool-proposals/{proposalId} | Get a tool approval proposal



## DecideToolProposal

> DecideToolProposal(ctx, organisation, proposalId).DecideToolProposalRequest(decideToolProposalRequest).XQuantApprovalNonce(xQuantApprovalNonce).Execute()

Decide a tool approval proposal



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
	proposalId := "proposalId_example" // string | The proposal ID
	decideToolProposalRequest := *openapiclient.NewDecideToolProposalRequest("Decision_example") // DecideToolProposalRequest | 
	xQuantApprovalNonce := "xQuantApprovalNonce_example" // string | One-time approval nonce from the approval prompt. It is forwarded to the platform and never logged. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIToolApprovalAPI.DecideToolProposal(context.Background(), organisation, proposalId).DecideToolProposalRequest(decideToolProposalRequest).XQuantApprovalNonce(xQuantApprovalNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolApprovalAPI.DecideToolProposal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**proposalId** | **string** | The proposal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecideToolProposalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **decideToolProposalRequest** | [**DecideToolProposalRequest**](DecideToolProposalRequest.md) |  | 
 **xQuantApprovalNonce** | **string** | One-time approval nonce from the approval prompt. It is forwarded to the platform and never logged. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToolProposal

> GetToolProposal(ctx, organisation, proposalId).Execute()

Get a tool approval proposal



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
	proposalId := "proposalId_example" // string | The proposal ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIToolApprovalAPI.GetToolProposal(context.Background(), organisation, proposalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolApprovalAPI.GetToolProposal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**proposalId** | **string** | The proposal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolProposalRequest struct via the builder pattern


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

