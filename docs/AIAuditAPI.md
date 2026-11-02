# \AIAuditAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiAuditExport**](AIAuditAPI.md#AiAuditExport) | **Get** /api/v3/organizations/{organisation}/ai/audit/export | Export AI audit records as JSON Lines
[**AiAuditLedgerEntries**](AIAuditAPI.md#AiAuditLedgerEntries) | **Get** /api/v3/organizations/{organisation}/ai/audit/ledger/entries | List AI governance ledger entries
[**AiAuditLedgerProof**](AIAuditAPI.md#AiAuditLedgerProof) | **Get** /api/v3/organizations/{organisation}/ai/audit/ledger/proof | Get the inclusion proof of a governance ledger entry
[**AiAuditLedgerPublicKey**](AIAuditAPI.md#AiAuditLedgerPublicKey) | **Get** /api/v3/organizations/{organisation}/ai/audit/ledger/public-key | Get the governance ledger public key
[**AiAuditLedgerStatus**](AIAuditAPI.md#AiAuditLedgerStatus) | **Get** /api/v3/organizations/{organisation}/ai/audit/ledger | Get the AI governance ledger status
[**AiAuditRecordGet**](AIAuditAPI.md#AiAuditRecordGet) | **Get** /api/v3/organizations/{organisation}/ai/audit/record | Get one AI audit record
[**AiAuditRecordsList**](AIAuditAPI.md#AiAuditRecordsList) | **Get** /api/v3/organizations/{organisation}/ai/audit/records | List AI audit records



## AiAuditExport

> string AiAuditExport(ctx, organisation).StartDate(startDate).EndDate(endDate).UserId(userId).ModelId(modelId).RecordType(recordType).Decision(decision).Tool(tool).Outcome(outcome).Execute()

Export AI audit records as JSON Lines



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func main() {
	organisation := "organisation_example" // string | The organisation machine name
	startDate := time.Now() // string | 
	endDate := time.Now() // string | 
	userId := "userId_example" // string |  (optional)
	modelId := "modelId_example" // string |  (optional)
	recordType := "recordType_example" // string |  (optional)
	decision := "decision_example" // string |  (optional)
	tool := "tool_example" // string |  (optional)
	outcome := "outcome_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAuditAPI.AiAuditExport(context.Background(), organisation).StartDate(startDate).EndDate(endDate).UserId(userId).ModelId(modelId).RecordType(recordType).Decision(decision).Tool(tool).Outcome(outcome).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAuditAPI.AiAuditExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAuditExport`: string
	fmt.Fprintf(os.Stdout, "Response from `AIAuditAPI.AiAuditExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAuditExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **userId** | **string** |  | 
 **modelId** | **string** |  | 
 **recordType** | **string** |  | 
 **decision** | **string** |  | 
 **tool** | **string** |  | 
 **outcome** | **string** |  | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAuditLedgerEntries

> V2KvKeysPage AiAuditLedgerEntries(ctx, organisation).Since(since).Until(until).Limit(limit).Cursor(cursor).Execute()

List AI governance ledger entries



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
	organisation := "organisation_example" // string | The organisation machine name
	since := "since_example" // string | Epoch milliseconds or ISO 8601 with a zone. Inclusive. (optional)
	until := "until_example" // string | Epoch milliseconds or ISO 8601 with a zone. Exclusive. (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	cursor := "cursor_example" // string | The next_cursor of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAuditAPI.AiAuditLedgerEntries(context.Background(), organisation).Since(since).Until(until).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAuditAPI.AiAuditLedgerEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAuditLedgerEntries`: V2KvKeysPage
	fmt.Fprintf(os.Stdout, "Response from `AIAuditAPI.AiAuditLedgerEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAuditLedgerEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **string** | Epoch milliseconds or ISO 8601 with a zone. Inclusive. | 
 **until** | **string** | Epoch milliseconds or ISO 8601 with a zone. Exclusive. | 
 **limit** | **int32** |  | [default to 100]
 **cursor** | **string** | The next_cursor of the previous page. | 

### Return type

[**V2KvKeysPage**](V2KvKeysPage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAuditLedgerProof

> V2KvLedgerProof AiAuditLedgerProof(ctx, organisation).Key(key).Execute()

Get the inclusion proof of a governance ledger entry



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
	organisation := "organisation_example" // string | The organisation machine name
	key := "events:governance:01JB0000000000000000000000" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAuditAPI.AiAuditLedgerProof(context.Background(), organisation).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAuditAPI.AiAuditLedgerProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAuditLedgerProof`: V2KvLedgerProof
	fmt.Fprintf(os.Stdout, "Response from `AIAuditAPI.AiAuditLedgerProof`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAuditLedgerProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 

### Return type

[**V2KvLedgerProof**](V2KvLedgerProof.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAuditLedgerPublicKey

> V2KvLedgerPublicKey AiAuditLedgerPublicKey(ctx, organisation).Execute()

Get the governance ledger public key



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
	organisation := "organisation_example" // string | The organisation machine name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAuditAPI.AiAuditLedgerPublicKey(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAuditAPI.AiAuditLedgerPublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAuditLedgerPublicKey`: V2KvLedgerPublicKey
	fmt.Fprintf(os.Stdout, "Response from `AIAuditAPI.AiAuditLedgerPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAuditLedgerPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2KvLedgerPublicKey**](V2KvLedgerPublicKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAuditLedgerStatus

> V2KvLedgerStatus AiAuditLedgerStatus(ctx, organisation).Cursor(cursor).Execute()

Get the AI governance ledger status



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
	organisation := "organisation_example" // string | The organisation machine name
	cursor := "cursor_example" // string | The next_cursor of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAuditAPI.AiAuditLedgerStatus(context.Background(), organisation).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAuditAPI.AiAuditLedgerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAuditLedgerStatus`: V2KvLedgerStatus
	fmt.Fprintf(os.Stdout, "Response from `AIAuditAPI.AiAuditLedgerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAuditLedgerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The next_cursor of the previous page. | 

### Return type

[**V2KvLedgerStatus**](V2KvLedgerStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAuditRecordGet

> AiAuditRecordGet200Response AiAuditRecordGet(ctx, organisation).Key(key).Execute()

Get one AI audit record



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
	organisation := "organisation_example" // string | The organisation machine name
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAuditAPI.AiAuditRecordGet(context.Background(), organisation).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAuditAPI.AiAuditRecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAuditRecordGet`: AiAuditRecordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAuditAPI.AiAuditRecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAuditRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 

### Return type

[**AiAuditRecordGet200Response**](AiAuditRecordGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAuditRecordsList

> AiAuditRecordsPage AiAuditRecordsList(ctx, organisation).StartDate(startDate).EndDate(endDate).Since(since).Until(until).UserId(userId).TokenId(tokenId).ModelId(modelId).RecordType(recordType).Decision(decision).Tool(tool).Outcome(outcome).Interface_(interface_).Limit(limit).Cursor(cursor).Execute()

List AI audit records



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/quantcdn/quant-admin-go"
)

func main() {
	organisation := "organisation_example" // string | The organisation machine name
	startDate := time.Now() // string | First UTC day (inclusive). At most 31 days, not in the future, not older than 365 days.
	endDate := time.Now() // string | Last UTC day (inclusive).
	since := time.Now() // time.Time | Record time from (inclusive). (optional)
	until := time.Now() // time.Time | Record time to (exclusive). (optional)
	userId := "userId_example" // string | Portal user id (also matches the requester, approver and reader of tool and access records). (optional)
	tokenId := "tokenId_example" // string |  (optional)
	modelId := "modelId_example" // string |  (optional)
	recordType := "inference" // string |  (optional)
	decision := "approved" // string | approved, rejected, expired, or denied (tool_denied). (optional)
	tool := "tool_example" // string |  (optional)
	outcome := "failed" // string | outcome or status: succeeded, failed, blocked, ok, refused, error. (optional)
	interface_ := "api-gateway" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := "cursor_example" // string | nextCursor of the previous page (same organisation and range). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAuditAPI.AiAuditRecordsList(context.Background(), organisation).StartDate(startDate).EndDate(endDate).Since(since).Until(until).UserId(userId).TokenId(tokenId).ModelId(modelId).RecordType(recordType).Decision(decision).Tool(tool).Outcome(outcome).Interface_(interface_).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAuditAPI.AiAuditRecordsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAuditRecordsList`: AiAuditRecordsPage
	fmt.Fprintf(os.Stdout, "Response from `AIAuditAPI.AiAuditRecordsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAuditRecordsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | First UTC day (inclusive). At most 31 days, not in the future, not older than 365 days. | 
 **endDate** | **string** | Last UTC day (inclusive). | 
 **since** | **time.Time** | Record time from (inclusive). | 
 **until** | **time.Time** | Record time to (exclusive). | 
 **userId** | **string** | Portal user id (also matches the requester, approver and reader of tool and access records). | 
 **tokenId** | **string** |  | 
 **modelId** | **string** |  | 
 **recordType** | **string** |  | 
 **decision** | **string** | approved, rejected, expired, or denied (tool_denied). | 
 **tool** | **string** |  | 
 **outcome** | **string** | outcome or status: succeeded, failed, blocked, ok, refused, error. | 
 **interface_** | **string** |  | 
 **limit** | **int32** |  | [default to 50]
 **cursor** | **string** | nextCursor of the previous page (same organisation and range). | 

### Return type

[**AiAuditRecordsPage**](AiAuditRecordsPage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

