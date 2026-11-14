# \AlertChannelsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertChannelsList**](AlertChannelsAPI.md#AlertChannelsList) | **Get** /api/v3/organizations/{organisation}/projects/{project}/alert-channels | List the organisation&#39;s alert channels



## AlertChannelsList

> AlertChannelList AlertChannelsList(ctx, organisation, project).Execute()

List the organisation's alert channels



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
	organisation := "test-org" // string | The organisation machine name
	project := "test-project" // string | The project machine name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertChannelsAPI.AlertChannelsList(context.Background(), organisation, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertChannelsAPI.AlertChannelsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertChannelsList`: AlertChannelList
	fmt.Fprintf(os.Stdout, "Response from `AlertChannelsAPI.AlertChannelsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**project** | **string** | The project machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertChannelsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertChannelList**](AlertChannelList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

