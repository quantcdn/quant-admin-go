# \KVAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KVArchivesList**](KVAPI.md#KVArchivesList) | **Get** /api/v2/organizations/{organization}/kv/archives | List retained ledger archives
[**KVArchivesPurge**](KVAPI.md#KVArchivesPurge) | **Post** /api/v2/organizations/{organization}/kv/archives/{store_name}/purge | Dry run or start a purge of a retained archive
[**KVArchivesPurgeManifest**](KVAPI.md#KVArchivesPurgeManifest) | **Get** /api/v2/organizations/{organization}/kv/archives/{store_name}/purges/{purge_id}/manifest | Get the manifest of a purge
[**KVArchivesPurgeReceipt**](KVAPI.md#KVArchivesPurgeReceipt) | **Get** /api/v2/organizations/{organization}/kv/archives/{store_name}/purges/{purge_id}/receipt | Get the signed receipt of a purge
[**KVArchivesPurgeShow**](KVAPI.md#KVArchivesPurgeShow) | **Get** /api/v2/organizations/{organization}/kv/archives/{store_name}/purges/{purge_id} | Show one purge of an archive
[**KVArchivesShow**](KVAPI.md#KVArchivesShow) | **Get** /api/v2/organizations/{organization}/kv/archives/{store_name} | Show the archive of a deleted store
[**KVCreate**](KVAPI.md#KVCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv | Add a kv store
[**KVDelete**](KVAPI.md#KVDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id} | Delete a kv store
[**KVEntriesDelete**](KVAPI.md#KVEntriesDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key}/entry | Delete a key, optionally only at a version
[**KVEntriesPut**](KVAPI.md#KVEntriesPut) | **Put** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key}/entry | Set a key, optionally only if absent or only at a version
[**KVEntriesShow**](KVAPI.md#KVEntriesShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key}/entry | Get a key with its version and metadata
[**KVExportCreate**](KVAPI.md#KVExportCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/export | Start an export of the whole store to gzip JSON Lines files
[**KVExportResume**](KVAPI.md#KVExportResume) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/export/resume | Queue a stalled export step again
[**KVExportShow**](KVAPI.md#KVExportShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/export | Get the latest export and, when done, its download links
[**KVIndexesCount**](KVAPI.md#KVIndexesCount) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/indexes/{index_name}/count | Count the keys that hold a value of a count index
[**KVIndexesCreate**](KVAPI.md#KVIndexesCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/indexes | Declare an index
[**KVIndexesDelete**](KVAPI.md#KVIndexesDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/indexes/{index_name} | Drop an index
[**KVIndexesList**](KVAPI.md#KVIndexesList) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/indexes | List the declared indexes of a store
[**KVIndexesLookup**](KVAPI.md#KVIndexesLookup) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/indexes/{index_name}/lookup | Get the key that holds a value of a unique index
[**KVIndexesResume**](KVAPI.md#KVIndexesResume) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/indexes/{index_name}/resume | Queue a stalled index job step again
[**KVIndexesShow**](KVAPI.md#KVIndexesShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/indexes/{index_name} | Get one index with its state and job
[**KVItemsBatchGet**](KVAPI.md#KVItemsBatchGet) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/batch-get | Get up to 100 keys in one request
[**KVItemsCreate**](KVAPI.md#KVItemsCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items | Add an item to a kv store
[**KVItemsDelete**](KVAPI.md#KVItemsDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key} | Delete an item from a kv store
[**KVItemsIncrement**](KVAPI.md#KVItemsIncrement) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key}/increment | Atomically add to a numeric value
[**KVItemsList**](KVAPI.md#KVItemsList) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items | List items in a kv store
[**KVItemsPatch**](KVAPI.md#KVItemsPatch) | **Patch** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key} | Update an item in a kv store (PATCH alias of PUT)
[**KVItemsPurge**](KVAPI.md#KVItemsPurge) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items | Delete items in bulk by prefix and/or age
[**KVItemsShow**](KVAPI.md#KVItemsShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key} | Get an item from a kv store
[**KVItemsTransact**](KVAPI.md#KVItemsTransact) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/transact | Run 1 to 25 operations as one transaction
[**KVItemsUpdate**](KVAPI.md#KVItemsUpdate) | **Put** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/items/{key} | Update an item in a kv store
[**KVKeysList**](KVAPI.md#KVKeysList) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/keys | List keys with a cursor: by collection, prefix, filter, index or ledger time window
[**KVLedgersAppend**](KVAPI.md#KVLedgersAppend) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/append | Append an entry to a ledger
[**KVLedgersArchive**](KVAPI.md#KVLedgersArchive) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/archive | Get download links for the ledger archive, digests and anchor files
[**KVLedgersCreate**](KVAPI.md#KVLedgersCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers | Declare a ledger
[**KVLedgersDelete**](KVAPI.md#KVLedgersDelete) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name} | Drop a ledger
[**KVLedgersList**](KVAPI.md#KVLedgersList) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers | List the ledgers of a store
[**KVLedgersProof**](KVAPI.md#KVLedgersProof) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/proof | Get the inclusion proof of one ledger entry
[**KVLedgersPublicKey**](KVAPI.md#KVLedgersPublicKey) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/public-key | Get the public key that verifies ledger signatures
[**KVLedgersShow**](KVAPI.md#KVLedgersShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name} | Get a ledger definition with its retention end and last verify
[**KVLedgersShred**](KVAPI.md#KVLedgersShred) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/shred | Shred a data subject (cannot be undone)
[**KVLedgersShredsEntries**](KVAPI.md#KVLedgersShredsEntries) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/shreds/{shred_id}/entries | List the shred entries of an erasure record
[**KVLedgersShredsList**](KVAPI.md#KVLedgersShredsList) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/shreds | List the erasure records of a ledger
[**KVLedgersShredsReceipt**](KVAPI.md#KVLedgersShredsReceipt) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/shreds/{shred_id}/receipt | Get the signed receipt of a shred
[**KVLedgersStatus**](KVAPI.md#KVLedgersStatus) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/status | Get the chains, archive lag, digest, anchor and verify state of a ledger
[**KVLedgersSubjectEntries**](KVAPI.md#KVLedgersSubjectEntries) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/subject-entries | List the ledger entries of one data subject
[**KVLedgersSubjectPreview**](KVAPI.md#KVLedgersSubjectPreview) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/subject-preview | Preview a shred of a data subject
[**KVLedgersUpdate**](KVAPI.md#KVLedgersUpdate) | **Patch** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name} | Raise a ledger
[**KVLedgersVerificationsReport**](KVAPI.md#KVLedgersVerificationsReport) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/verifications/{verification_id}/report | Get the signed report of one verification run
[**KVLedgersVerify**](KVAPI.md#KVLedgersVerify) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/ledgers/{ledger_name}/verify | Verify a ledger against its archive and signed digests
[**KVLinkToProject**](KVAPI.md#KVLinkToProject) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/link | Link a KV store to another project
[**KVList**](KVAPI.md#KVList) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv | List key-value stores
[**KVShow**](KVAPI.md#KVShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id} | Get a kv store
[**KVUnlinkFromProject**](KVAPI.md#KVUnlinkFromProject) | **Delete** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/link | Unlink a KV store from this project
[**KVUpgradeCreate**](KVAPI.md#KVUpgradeCreate) | **Post** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/upgrade | Start, roll back, finalize or resume a format 1 to format 2 upgrade
[**KVUpgradeShow**](KVAPI.md#KVUpgradeShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/upgrade | Get the store format and the latest upgrade
[**KVUsageShow**](KVAPI.md#KVUsageShow) | **Get** /api/v2/organizations/{organization}/projects/{project}/kv/{store_id}/usage | Get 30-day usage and a cost estimate



## KVArchivesList

> V2KvArchiveList KVArchivesList(ctx, organization).Limit(limit).Cursor(cursor).Execute()

List retained ledger archives



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
	organization := "test-org" // string | Organization identifier
	limit := int32(56) // int32 | The maximum number of archives in the page (1-100, default 50). (optional) (default to 50)
	cursor := "cursor_example" // string | The next_cursor of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVArchivesList(context.Background(), organization).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVArchivesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVArchivesList`: V2KvArchiveList
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVArchivesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVArchivesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of archives in the page (1-100, default 50). | [default to 50]
 **cursor** | **string** | The next_cursor of the previous page. | 

### Return type

[**V2KvArchiveList**](V2KvArchiveList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVArchivesPurge

> V2KvArchivePurgeAnswer KVArchivesPurge(ctx, organization, storeName).V2KvArchivePurgeRequest(v2KvArchivePurgeRequest).IdempotencyKey(idempotencyKey).Execute()

Dry run or start a purge of a retained archive



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
	organization := "test-org" // string | Organization identifier
	storeName := "orders" // string | The id of the deleted store, as KV_list showed it before the delete
	v2KvArchivePurgeRequest := *openapiclient.NewV2KvArchivePurgeRequest("Confirm_example") // V2KvArchivePurgeRequest | 
	idempotencyKey := "idempotencyKey_example" // string | 1 to 128 characters of letters, digits and . _ : -. A repeat with the same key answers 202 with the existing purge. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVArchivesPurge(context.Background(), organization, storeName).V2KvArchivePurgeRequest(v2KvArchivePurgeRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVArchivesPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVArchivesPurge`: V2KvArchivePurgeAnswer
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVArchivesPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**storeName** | **string** | The id of the deleted store, as KV_list showed it before the delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVArchivesPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2KvArchivePurgeRequest** | [**V2KvArchivePurgeRequest**](V2KvArchivePurgeRequest.md) |  | 
 **idempotencyKey** | **string** | 1 to 128 characters of letters, digits and . _ : -. A repeat with the same key answers 202 with the existing purge. | 

### Return type

[**V2KvArchivePurgeAnswer**](V2KvArchivePurgeAnswer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVArchivesPurgeManifest

> V2KvArchiveManifest KVArchivesPurgeManifest(ctx, organization, storeName, purgeId).Execute()

Get the manifest of a purge



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
	organization := "test-org" // string | Organization identifier
	storeName := "orders" // string | The id of the deleted store, as KV_list showed it before the delete
	purgeId := "01JBX5Q8M7R2T4V6W8YZ0A1B2C" // string | The purge id (a 26-character ULID), as the purge start answered it

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVArchivesPurgeManifest(context.Background(), organization, storeName, purgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVArchivesPurgeManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVArchivesPurgeManifest`: V2KvArchiveManifest
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVArchivesPurgeManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**storeName** | **string** | The id of the deleted store, as KV_list showed it before the delete | 
**purgeId** | **string** | The purge id (a 26-character ULID), as the purge start answered it | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVArchivesPurgeManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2KvArchiveManifest**](V2KvArchiveManifest.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVArchivesPurgeReceipt

> V2KvArchivePurgeReceipt KVArchivesPurgeReceipt(ctx, organization, storeName, purgeId).Execute()

Get the signed receipt of a purge



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
	organization := "test-org" // string | Organization identifier
	storeName := "orders" // string | The id of the deleted store, as KV_list showed it before the delete
	purgeId := "01JBX5Q8M7R2T4V6W8YZ0A1B2C" // string | The purge id (a 26-character ULID), as the purge start answered it

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVArchivesPurgeReceipt(context.Background(), organization, storeName, purgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVArchivesPurgeReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVArchivesPurgeReceipt`: V2KvArchivePurgeReceipt
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVArchivesPurgeReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**storeName** | **string** | The id of the deleted store, as KV_list showed it before the delete | 
**purgeId** | **string** | The purge id (a 26-character ULID), as the purge start answered it | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVArchivesPurgeReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2KvArchivePurgeReceipt**](V2KvArchivePurgeReceipt.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVArchivesPurgeShow

> V2KvArchivePurge KVArchivesPurgeShow(ctx, organization, storeName, purgeId).Execute()

Show one purge of an archive



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
	organization := "test-org" // string | Organization identifier
	storeName := "orders" // string | The id of the deleted store, as KV_list showed it before the delete
	purgeId := "01JBX5Q8M7R2T4V6W8YZ0A1B2C" // string | The purge id (a 26-character ULID), as the purge start answered it

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVArchivesPurgeShow(context.Background(), organization, storeName, purgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVArchivesPurgeShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVArchivesPurgeShow`: V2KvArchivePurge
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVArchivesPurgeShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**storeName** | **string** | The id of the deleted store, as KV_list showed it before the delete | 
**purgeId** | **string** | The purge id (a 26-character ULID), as the purge start answered it | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVArchivesPurgeShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2KvArchivePurge**](V2KvArchivePurge.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVArchivesShow

> V2KvArchiveStatus KVArchivesShow(ctx, organization, storeName).Execute()

Show the archive of a deleted store



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
	organization := "test-org" // string | Organization identifier
	storeName := "orders" // string | The id of the deleted store, as KV_list showed it before the delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVArchivesShow(context.Background(), organization, storeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVArchivesShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVArchivesShow`: V2KvArchiveStatus
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVArchivesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**storeName** | **string** | The id of the deleted store, as KV_list showed it before the delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVArchivesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**V2KvArchiveStatus**](V2KvArchiveStatus.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVCreate

> V2Store KVCreate(ctx, organization, project).V2StoreRequest(v2StoreRequest).Execute()

Add a kv store



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	v2StoreRequest := *openapiclient.NewV2StoreRequest("session-data") // V2StoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVCreate(context.Background(), organization, project).V2StoreRequest(v2StoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVCreate`: V2Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **v2StoreRequest** | [**V2StoreRequest**](V2StoreRequest.md) |  | 

### Return type

[**V2Store**](V2Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVDelete

> KVDelete(ctx, organization, project, storeId).Force(force).Execute()

Delete a kv store



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	force := true // bool | Delete the store even if it still holds keys. Without it a non-empty store returns 409. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KVAPI.KVDelete(context.Background(), organization, project, storeId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVDelete``: %v\n", err)
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
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | Delete the store even if it still holds keys. Without it a non-empty store returns 409. | [default to false]

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


## KVEntriesDelete

> V2KvEntryDeleteResponse KVEntriesDelete(ctx, organization, project, storeId, key).IfVersion(ifVersion).Execute()

Delete a key, optionally only at a version



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	key := "order:1001" // string | The key. It can contain '/'.
	ifVersion := "ifVersion_example" // string | Format 2 stores only. Delete only at this version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVEntriesDelete(context.Background(), organization, project, storeId, key).IfVersion(ifVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVEntriesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVEntriesDelete`: V2KvEntryDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVEntriesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**key** | **string** | The key. It can contain &#39;/&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVEntriesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ifVersion** | **string** | Format 2 stores only. Delete only at this version. | 

### Return type

[**V2KvEntryDeleteResponse**](V2KvEntryDeleteResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVEntriesPut

> V2KvEntryPutResponse KVEntriesPut(ctx, organization, project, storeId, key).V2KvEntryPutRequest(v2KvEntryPutRequest).Execute()

Set a key, optionally only if absent or only at a version



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	key := "order:1001" // string | The key. It can contain '/'.
	v2KvEntryPutRequest := *openapiclient.NewV2KvEntryPutRequest(interface{}(123)) // V2KvEntryPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVEntriesPut(context.Background(), organization, project, storeId, key).V2KvEntryPutRequest(v2KvEntryPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVEntriesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVEntriesPut`: V2KvEntryPutResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVEntriesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**key** | **string** | The key. It can contain &#39;/&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVEntriesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2KvEntryPutRequest** | [**V2KvEntryPutRequest**](V2KvEntryPutRequest.md) |  | 

### Return type

[**V2KvEntryPutResponse**](V2KvEntryPutResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVEntriesShow

> V2KvEntry KVEntriesShow(ctx, organization, project, storeId, key).Consistent(consistent).Reveal(reveal).Execute()

Get a key with its version and metadata



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	key := "order:1001" // string | The key. It can contain '/'.
	consistent := true // bool | false for an eventually consistent read. (optional) (default to true)
	reveal := true // bool | Ledger rows only: return personal fields in clear. Needs the ledger_reveal permission; each reveal is logged. Without it, personal fields are redacted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVEntriesShow(context.Background(), organization, project, storeId, key).Consistent(consistent).Reveal(reveal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVEntriesShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVEntriesShow`: V2KvEntry
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVEntriesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**key** | **string** | The key. It can contain &#39;/&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVEntriesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **consistent** | **bool** | false for an eventually consistent read. | [default to true]
 **reveal** | **bool** | Ledger rows only: return personal fields in clear. Needs the ledger_reveal permission; each reveal is logged. Without it, personal fields are redacted. | [default to false]

### Return type

[**V2KvEntry**](V2KvEntry.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVExportCreate

> V2KvExportResponse KVExportCreate(ctx, organization, project, storeId).Body(body).Execute()

Start an export of the whole store to gzip JSON Lines files



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVExportCreate(context.Background(), organization, project, storeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVExportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVExportCreate`: V2KvExportResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVExportCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVExportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **map[string]interface{}** |  | 

### Return type

[**V2KvExportResponse**](V2KvExportResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVExportResume

> V2KvExportResponse KVExportResume(ctx, organization, project, storeId).Execute()

Queue a stalled export step again



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVExportResume(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVExportResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVExportResume`: V2KvExportResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVExportResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVExportResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2KvExportResponse**](V2KvExportResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVExportShow

> V2KvExportResponse KVExportShow(ctx, organization, project, storeId).Execute()

Get the latest export and, when done, its download links



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVExportShow(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVExportShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVExportShow`: V2KvExportResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVExportShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVExportShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2KvExportResponse**](V2KvExportResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVIndexesCount

> V2KvCount KVIndexesCount(ctx, organization, project, storeId, indexName).Equals(equals).Execute()

Count the keys that hold a value of a count index



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	indexName := "byCustomer" // string | 
	equals := "{"plan":"pro"}" // string | A JSON object that gives a JSON scalar for exactly the index fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVIndexesCount(context.Background(), organization, project, storeId, indexName).Equals(equals).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVIndexesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVIndexesCount`: V2KvCount
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVIndexesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVIndexesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **equals** | **string** | A JSON object that gives a JSON scalar for exactly the index fields. | 

### Return type

[**V2KvCount**](V2KvCount.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVIndexesCreate

> V2KvIndex KVIndexesCreate(ctx, organization, project, storeId).V2KvIndexRequest(v2KvIndexRequest).Execute()

Declare an index



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	v2KvIndexRequest := *openapiclient.NewV2KvIndexRequest("byCustomer", "order") // V2KvIndexRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVIndexesCreate(context.Background(), organization, project, storeId).V2KvIndexRequest(v2KvIndexRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVIndexesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVIndexesCreate`: V2KvIndex
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVIndexesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVIndexesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2KvIndexRequest** | [**V2KvIndexRequest**](V2KvIndexRequest.md) |  | 

### Return type

[**V2KvIndex**](V2KvIndex.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVIndexesDelete

> V2KvIndex KVIndexesDelete(ctx, organization, project, storeId, indexName).Execute()

Drop an index



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	indexName := "byCustomer" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVIndexesDelete(context.Background(), organization, project, storeId, indexName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVIndexesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVIndexesDelete`: V2KvIndex
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVIndexesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVIndexesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**V2KvIndex**](V2KvIndex.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVIndexesList

> V2KvIndexList KVIndexesList(ctx, organization, project, storeId).Limit(limit).Execute()

List the declared indexes of a store



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	limit := int32(56) // int32 | The maximum number of indexes to return (1-1000). Omit it for all indexes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVIndexesList(context.Background(), organization, project, storeId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVIndexesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVIndexesList`: V2KvIndexList
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVIndexesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVIndexesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The maximum number of indexes to return (1-1000). Omit it for all indexes. | 

### Return type

[**V2KvIndexList**](V2KvIndexList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVIndexesLookup

> V2KvIndexLookup KVIndexesLookup(ctx, organization, project, storeId, indexName).Equals(equals).Reveal(reveal).Execute()

Get the key that holds a value of a unique index



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	indexName := "byCustomer" // string | 
	equals := "{"email":"ann@example.com"}" // string | A JSON object that gives a JSON scalar for exactly the index fields. Values are type-tagged (1 and \"1\" differ).
	reveal := true // bool | Ledger rows only: return personal fields in clear. Needs the ledger_reveal permission; each reveal is logged. Without it, personal fields are redacted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVIndexesLookup(context.Background(), organization, project, storeId, indexName).Equals(equals).Reveal(reveal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVIndexesLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVIndexesLookup`: V2KvIndexLookup
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVIndexesLookup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVIndexesLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **equals** | **string** | A JSON object that gives a JSON scalar for exactly the index fields. Values are type-tagged (1 and \&quot;1\&quot; differ). | 
 **reveal** | **bool** | Ledger rows only: return personal fields in clear. Needs the ledger_reveal permission; each reveal is logged. Without it, personal fields are redacted. | [default to false]

### Return type

[**V2KvIndexLookup**](V2KvIndexLookup.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVIndexesResume

> V2KvIndex KVIndexesResume(ctx, organization, project, storeId, indexName).Execute()

Queue a stalled index job step again



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	indexName := "byCustomer" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVIndexesResume(context.Background(), organization, project, storeId, indexName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVIndexesResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVIndexesResume`: V2KvIndex
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVIndexesResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVIndexesResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**V2KvIndex**](V2KvIndex.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVIndexesShow

> V2KvIndex KVIndexesShow(ctx, organization, project, storeId, indexName).Execute()

Get one index with its state and job



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	indexName := "byCustomer" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVIndexesShow(context.Background(), organization, project, storeId, indexName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVIndexesShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVIndexesShow`: V2KvIndex
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVIndexesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**indexName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVIndexesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**V2KvIndex**](V2KvIndex.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsBatchGet

> V2KvBatchGetResponse KVItemsBatchGet(ctx, organization, project, storeId).V2KvBatchGetRequest(v2KvBatchGetRequest).Reveal(reveal).Execute()

Get up to 100 keys in one request



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	v2KvBatchGetRequest := *openapiclient.NewV2KvBatchGetRequest([]string{"Keys_example"}) // V2KvBatchGetRequest | 
	reveal := true // bool | Ledger rows only: return personal fields in clear. Needs the ledger_reveal permission; each reveal is logged. Without it, personal fields are redacted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsBatchGet(context.Background(), organization, project, storeId).V2KvBatchGetRequest(v2KvBatchGetRequest).Reveal(reveal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsBatchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsBatchGet`: V2KvBatchGetResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsBatchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsBatchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2KvBatchGetRequest** | [**V2KvBatchGetRequest**](V2KvBatchGetRequest.md) |  | 
 **reveal** | **bool** | Ledger rows only: return personal fields in clear. Needs the ledger_reveal permission; each reveal is logged. Without it, personal fields are redacted. | [default to false]

### Return type

[**V2KvBatchGetResponse**](V2KvBatchGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsCreate

> KVItemsCreate200Response KVItemsCreate(ctx, organization, project, storeId).V2StoreItemRequest(v2StoreItemRequest).Execute()

Add an item to a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	v2StoreItemRequest := *openapiclient.NewV2StoreItemRequest("user-session-123", "Value_example") // V2StoreItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsCreate(context.Background(), organization, project, storeId).V2StoreItemRequest(v2StoreItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsCreate`: KVItemsCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2StoreItemRequest** | [**V2StoreItemRequest**](V2StoreItemRequest.md) |  | 

### Return type

[**KVItemsCreate200Response**](KVItemsCreate200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsDelete

> KVItemsDelete200Response KVItemsDelete(ctx, organization, project, storeId, key).Execute()

Delete an item from a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsDelete(context.Background(), organization, project, storeId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsDelete`: KVItemsDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**KVItemsDelete200Response**](KVItemsDelete200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsIncrement

> V2KvIncrementResponse KVItemsIncrement(ctx, organization, project, storeId, key).V2KvIncrementRequest(v2KvIncrementRequest).Execute()

Atomically add to a numeric value



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	key := "order:1001" // string | The key. It can contain '/'.
	v2KvIncrementRequest := *openapiclient.NewV2KvIncrementRequest() // V2KvIncrementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsIncrement(context.Background(), organization, project, storeId, key).V2KvIncrementRequest(v2KvIncrementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsIncrement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsIncrement`: V2KvIncrementResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsIncrement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**key** | **string** | The key. It can contain &#39;/&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsIncrementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2KvIncrementRequest** | [**V2KvIncrementRequest**](V2KvIncrementRequest.md) |  | 

### Return type

[**V2KvIncrementResponse**](V2KvIncrementResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsList

> V2StoreItemsListResponse KVItemsList(ctx, organization, project, storeId).Cursor(cursor).Limit(limit).Search(search).IncludeValues(includeValues).Execute()

List items in a kv store



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	cursor := "cursor_example" // string | Cursor for pagination (optional)
	limit := int32(56) // int32 | Number of items to return (1 to 1000, default 10) (optional) (default to 10)
	search := "search_example" // string | Search filter for keys (optional)
	includeValues := true // bool | Include values in the response. Secret values will be redacted as '[ENCRYPTED]' for security. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsList(context.Background(), organization, project, storeId).Cursor(cursor).Limit(limit).Search(search).IncludeValues(includeValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsList`: V2StoreItemsListResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **string** | Cursor for pagination | 
 **limit** | **int32** | Number of items to return (1 to 1000, default 10) | [default to 10]
 **search** | **string** | Search filter for keys | 
 **includeValues** | **bool** | Include values in the response. Secret values will be redacted as &#39;[ENCRYPTED]&#39; for security. | [default to false]

### Return type

[**V2StoreItemsListResponse**](V2StoreItemsListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsPatch

> KVItemsCreate200Response KVItemsPatch(ctx, organization, project, storeId, key).V2StoreItemUpdateRequest(v2StoreItemUpdateRequest).Execute()

Update an item in a kv store (PATCH alias of PUT)

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	key := "key_example" // string | 
	v2StoreItemUpdateRequest := *openapiclient.NewV2StoreItemUpdateRequest("Value_example") // V2StoreItemUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsPatch(context.Background(), organization, project, storeId, key).V2StoreItemUpdateRequest(v2StoreItemUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsPatch`: KVItemsCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2StoreItemUpdateRequest** | [**V2StoreItemUpdateRequest**](V2StoreItemUpdateRequest.md) |  | 

### Return type

[**KVItemsCreate200Response**](KVItemsCreate200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsPurge

> KVItemsPurge200Response KVItemsPurge(ctx, organization, project, storeId).Prefix(prefix).OlderThan(olderThan).Execute()

Delete items in bulk by prefix and/or age



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	prefix := "oauth_state:" // string | Only delete keys that start with this string. (optional)
	olderThan := "24h" // string | Only delete keys last updated before this instant. ISO 8601, or a duration with unit s, m, h or d. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsPurge(context.Background(), organization, project, storeId).Prefix(prefix).OlderThan(olderThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsPurge`: KVItemsPurge200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **prefix** | **string** | Only delete keys that start with this string. | 
 **olderThan** | **string** | Only delete keys last updated before this instant. ISO 8601, or a duration with unit s, m, h or d. | 

### Return type

[**KVItemsPurge200Response**](KVItemsPurge200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsShow

> KVItemsShow200Response KVItemsShow(ctx, organization, project, storeId, key).Execute()

Get an item from a kv store



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsShow(context.Background(), organization, project, storeId, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsShow`: KVItemsShow200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**KVItemsShow200Response**](KVItemsShow200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsTransact

> V2KvTransactResponse KVItemsTransact(ctx, organization, project, storeId).V2KvTransactRequest(v2KvTransactRequest).Execute()

Run 1 to 25 operations as one transaction



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	v2KvTransactRequest := *openapiclient.NewV2KvTransactRequest([]openapiclient.V2KvTransactOperation{*openapiclient.NewV2KvTransactOperation("Type_example", "Key_example")}) // V2KvTransactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsTransact(context.Background(), organization, project, storeId).V2KvTransactRequest(v2KvTransactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsTransact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsTransact`: V2KvTransactResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsTransact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsTransactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2KvTransactRequest** | [**V2KvTransactRequest**](V2KvTransactRequest.md) |  | 

### Return type

[**V2KvTransactResponse**](V2KvTransactResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVItemsUpdate

> KVItemsCreate200Response KVItemsUpdate(ctx, organization, project, storeId, key).V2StoreItemUpdateRequest(v2StoreItemUpdateRequest).Execute()

Update an item in a kv store

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 
	key := "key_example" // string | 
	v2StoreItemUpdateRequest := *openapiclient.NewV2StoreItemUpdateRequest("Value_example") // V2StoreItemUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVItemsUpdate(context.Background(), organization, project, storeId, key).V2StoreItemUpdateRequest(v2StoreItemUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVItemsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVItemsUpdate`: KVItemsCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVItemsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVItemsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2StoreItemUpdateRequest** | [**V2StoreItemUpdateRequest**](V2StoreItemUpdateRequest.md) |  | 

### Return type

[**KVItemsCreate200Response**](KVItemsCreate200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVKeysList

> V2KvKeysPage KVKeysList(ctx, organization, project, storeId).Collection(collection).Prefix(prefix).Pattern(pattern).StartsWith(startsWith).After(after).Before(before).Reverse(reverse).Since(since).Until(until).Index(index).Equals(equals).Where(where).Type_(type_).IncludeValues(includeValues).SubjectTag(subjectTag).Limit(limit).Cursor(cursor).Reveal(reveal).Execute()

List keys with a cursor: by collection, prefix, filter, index or ledger time window



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	collection := "collection_example" // string | One collection (the key text before the last ':'). Empty: the root collection. Not with prefix, pattern or index. (optional)
	prefix := "prefix_example" // string | Literal key prefix (a scan). Not with pattern or collection. (optional)
	pattern := "pattern_example" // string | Wildcard pattern such as user_*, *_temp or *session* (a scan). (optional)
	startsWith := "startsWith_example" // string | With collection: only ids that start with this text (no ':'). Not with after or before. (optional)
	after := "after_example" // string | Exclusive lower bound: an id with collection, a JSON scalar with index. (optional)
	before := "before_example" // string | Exclusive upper bound: an id with collection, a JSON scalar with index. (optional)
	reverse := true // bool | With collection or index: descending order. Format 2 only (else 400 REQUIRES_FORMAT_2). (optional) (default to false)
	since := "2026-10-16T00:00:00.000Z" // string | Ledger collection only. Epoch milliseconds or ISO 8601 with a zone. Inclusive. (optional)
	until := "1792195200000" // string | Ledger collection only. Epoch milliseconds or ISO 8601 with a zone. Exclusive; after since. (optional)
	index := "index_example" // string | List the members of this declared index (format 2). Not with collection, prefix, pattern or starts_with. (optional)
	equals := "{"customerId":"customer-123"}" // string | With index: a JSON object with exactly the index fields. A filter that names the subjectField of a ledger of the store is refused (400 VALIDATION): use KV_ledgers_subject_entries. (optional)
	where := "{"status":"ready"}" // string | A JSON object of 1 to 8 top-level equalities on plaintext object values (a filter, not an index). A filter that names the subjectField of a ledger of the store is refused (400 VALIDATION): use KV_ledgers_subject_entries. (optional)
	type_ := "type__example" // string | Filter by encryption. (optional) (default to "all")
	includeValues := true // bool | Return the values. Secret values read [ENCRYPTED] (write-only through the API). (optional) (default to false)
	subjectTag := "subjectTag_example" // string | Ledger collections only: the entries of one subject token, as shown in personal=sealed reads. (optional)
	limit := int32(56) // int32 | The maximum number of items in the page (1-1000, default 25). A page can be shorter while has_more is true (byte or read budget; see page_end). (optional) (default to 25)
	cursor := "cursor_example" // string | The next_cursor of the previous page. (optional)
	reveal := true // bool | Ledger rows only: return personal fields in clear. Needs the ledger_reveal permission; each reveal is logged. Without it, personal fields are redacted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVKeysList(context.Background(), organization, project, storeId).Collection(collection).Prefix(prefix).Pattern(pattern).StartsWith(startsWith).After(after).Before(before).Reverse(reverse).Since(since).Until(until).Index(index).Equals(equals).Where(where).Type_(type_).IncludeValues(includeValues).SubjectTag(subjectTag).Limit(limit).Cursor(cursor).Reveal(reveal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVKeysList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVKeysList`: V2KvKeysPage
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVKeysList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVKeysListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **collection** | **string** | One collection (the key text before the last &#39;:&#39;). Empty: the root collection. Not with prefix, pattern or index. | 
 **prefix** | **string** | Literal key prefix (a scan). Not with pattern or collection. | 
 **pattern** | **string** | Wildcard pattern such as user_*, *_temp or *session* (a scan). | 
 **startsWith** | **string** | With collection: only ids that start with this text (no &#39;:&#39;). Not with after or before. | 
 **after** | **string** | Exclusive lower bound: an id with collection, a JSON scalar with index. | 
 **before** | **string** | Exclusive upper bound: an id with collection, a JSON scalar with index. | 
 **reverse** | **bool** | With collection or index: descending order. Format 2 only (else 400 REQUIRES_FORMAT_2). | [default to false]
 **since** | **string** | Ledger collection only. Epoch milliseconds or ISO 8601 with a zone. Inclusive. | 
 **until** | **string** | Ledger collection only. Epoch milliseconds or ISO 8601 with a zone. Exclusive; after since. | 
 **index** | **string** | List the members of this declared index (format 2). Not with collection, prefix, pattern or starts_with. | 
 **equals** | **string** | With index: a JSON object with exactly the index fields. A filter that names the subjectField of a ledger of the store is refused (400 VALIDATION): use KV_ledgers_subject_entries. | 
 **where** | **string** | A JSON object of 1 to 8 top-level equalities on plaintext object values (a filter, not an index). A filter that names the subjectField of a ledger of the store is refused (400 VALIDATION): use KV_ledgers_subject_entries. | 
 **type_** | **string** | Filter by encryption. | [default to &quot;all&quot;]
 **includeValues** | **bool** | Return the values. Secret values read [ENCRYPTED] (write-only through the API). | [default to false]
 **subjectTag** | **string** | Ledger collections only: the entries of one subject token, as shown in personal&#x3D;sealed reads. | 
 **limit** | **int32** | The maximum number of items in the page (1-1000, default 25). A page can be shorter while has_more is true (byte or read budget; see page_end). | [default to 25]
 **cursor** | **string** | The next_cursor of the previous page. | 
 **reveal** | **bool** | Ledger rows only: return personal fields in clear. Needs the ledger_reveal permission; each reveal is logged. Without it, personal fields are redacted. | [default to false]

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


## KVLedgersAppend

> V2KvLedgerReceipt KVLedgersAppend(ctx, organization, project, storeId).V2KvLedgerAppendRequest(v2KvLedgerAppendRequest).Execute()

Append an entry to a ledger



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	v2KvLedgerAppendRequest := *openapiclient.NewV2KvLedgerAppendRequest("Collection_example", map[string]interface{}{"key": interface{}(123)}) // V2KvLedgerAppendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersAppend(context.Background(), organization, project, storeId).V2KvLedgerAppendRequest(v2KvLedgerAppendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersAppend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersAppend`: V2KvLedgerReceipt
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersAppend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersAppendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2KvLedgerAppendRequest** | [**V2KvLedgerAppendRequest**](V2KvLedgerAppendRequest.md) |  | 

### Return type

[**V2KvLedgerReceipt**](V2KvLedgerReceipt.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersArchive

> V2KvLedgerArchive KVLedgersArchive(ctx, organization, project, storeId, ledgerName).Cursor(cursor).DigestCursor(digestCursor).Execute()

Get download links for the ledger archive, digests and anchor files



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	cursor := "cursor_example" // string | The next_cursor of the previous page of objects. (optional)
	digestCursor := "digestCursor_example" // string | The next_digest_cursor of the previous digest page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersArchive(context.Background(), organization, project, storeId, ledgerName).Cursor(cursor).DigestCursor(digestCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersArchive`: V2KvLedgerArchive
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **cursor** | **string** | The next_cursor of the previous page of objects. | 
 **digestCursor** | **string** | The next_digest_cursor of the previous digest page. | 

### Return type

[**V2KvLedgerArchive**](V2KvLedgerArchive.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersCreate

> V2KvLedger KVLedgersCreate(ctx, organization, project, storeId).V2KvLedgerAdminDeclareRequest(v2KvLedgerAdminDeclareRequest).Execute()

Declare a ledger



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	v2KvLedgerAdminDeclareRequest := *openapiclient.NewV2KvLedgerAdminDeclareRequest("Name_example", "Collection_example", int32(123), int32(123), "LockMode_example") // V2KvLedgerAdminDeclareRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersCreate(context.Background(), organization, project, storeId).V2KvLedgerAdminDeclareRequest(v2KvLedgerAdminDeclareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersCreate`: V2KvLedger
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2KvLedgerAdminDeclareRequest** | [**V2KvLedgerAdminDeclareRequest**](V2KvLedgerAdminDeclareRequest.md) |  | 

### Return type

[**V2KvLedger**](V2KvLedger.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersDelete

> V2KvLedger KVLedgersDelete(ctx, organization, project, storeId, ledgerName).Execute()

Drop a ledger



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersDelete(context.Background(), organization, project, storeId, ledgerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersDelete`: V2KvLedger
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**V2KvLedger**](V2KvLedger.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersList

> V2KvLedgerList KVLedgersList(ctx, organization, project, storeId).Execute()

List the ledgers of a store



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersList(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersList`: V2KvLedgerList
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2KvLedgerList**](V2KvLedgerList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersProof

> V2KvLedgerProof KVLedgersProof(ctx, organization, project, storeId, ledgerName).Key(key).Execute()

Get the inclusion proof of one ledger entry



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	key := "ai-use:agency-1:01M5104A000000000000000000-00" // string | The key of a ledger entry (collection:id), as in the append receipt.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersProof(context.Background(), organization, project, storeId, ledgerName).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersProof`: V2KvLedgerProof
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersProof`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **key** | **string** | The key of a ledger entry (collection:id), as in the append receipt. | 

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


## KVLedgersPublicKey

> V2KvLedgerPublicKey KVLedgersPublicKey(ctx, organization, project, storeId).Execute()

Get the public key that verifies ledger signatures



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersPublicKey(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersPublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersPublicKey`: V2KvLedgerPublicKey
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersPublicKeyRequest struct via the builder pattern


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


## KVLedgersShow

> V2KvLedgerDetail KVLedgersShow(ctx, organization, project, storeId, ledgerName).Execute()

Get a ledger definition with its retention end and last verify



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersShow(context.Background(), organization, project, storeId, ledgerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersShow`: V2KvLedgerDetail
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**V2KvLedgerDetail**](V2KvLedgerDetail.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersShred

> V2KvLedgerErasureShredResult KVLedgersShred(ctx, organization, project, storeId, ledgerName).V2KvLedgerErasureShredRequest(v2KvLedgerErasureShredRequest).Execute()

Shred a data subject (cannot be undone)



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	v2KvLedgerErasureShredRequest := *openapiclient.NewV2KvLedgerErasureShredRequest(openapiclient.V2KvLedgerSubjectId{Int64: new(int64)}, "Reason_example", "Authority_example", "RequestedBy_example") // V2KvLedgerErasureShredRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersShred(context.Background(), organization, project, storeId, ledgerName).V2KvLedgerErasureShredRequest(v2KvLedgerErasureShredRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersShred``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersShred`: V2KvLedgerErasureShredResult
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersShred`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersShredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2KvLedgerErasureShredRequest** | [**V2KvLedgerErasureShredRequest**](V2KvLedgerErasureShredRequest.md) |  | 

### Return type

[**V2KvLedgerErasureShredResult**](V2KvLedgerErasureShredResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersShredsEntries

> V2KvLedgerErasureShredEntriesPage KVLedgersShredsEntries(ctx, organization, project, storeId, ledgerName, shredId).Cursor(cursor).Execute()

List the shred entries of an erasure record



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	shredId := "01JB2Z6V8Q3M4N5P6R7S8T9V0W" // string | The shredId of a shred answer, or the id of an erasure record (KV_ledgers_shreds_list). A ULID.
	cursor := "cursor_example" // string | The next_cursor of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersShredsEntries(context.Background(), organization, project, storeId, ledgerName, shredId).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersShredsEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersShredsEntries`: V2KvLedgerErasureShredEntriesPage
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersShredsEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 
**shredId** | **string** | The shredId of a shred answer, or the id of an erasure record (KV_ledgers_shreds_list). A ULID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersShredsEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **cursor** | **string** | The next_cursor of the previous page. | 

### Return type

[**V2KvLedgerErasureShredEntriesPage**](V2KvLedgerErasureShredEntriesPage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersShredsList

> V2KvLedgerErasureShredPage KVLedgersShredsList(ctx, organization, project, storeId, ledgerName).Cursor(cursor).Execute()

List the erasure records of a ledger



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersShredsList(context.Background(), organization, project, storeId, ledgerName).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersShredsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersShredsList`: V2KvLedgerErasureShredPage
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersShredsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersShredsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **cursor** | **string** | The next_cursor of the previous page. | 

### Return type

[**V2KvLedgerErasureShredPage**](V2KvLedgerErasureShredPage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersShredsReceipt

> V2KvLedgerErasureShredReceiptDocument KVLedgersShredsReceipt(ctx, organization, project, storeId, ledgerName, shredId).Execute()

Get the signed receipt of a shred



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	shredId := "01JB2Z6V8Q3M4N5P6R7S8T9V0W" // string | The shredId of a shred answer, or the id of an erasure record (KV_ledgers_shreds_list). A ULID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersShredsReceipt(context.Background(), organization, project, storeId, ledgerName, shredId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersShredsReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersShredsReceipt`: V2KvLedgerErasureShredReceiptDocument
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersShredsReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 
**shredId** | **string** | The shredId of a shred answer, or the id of an erasure record (KV_ledgers_shreds_list). A ULID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersShredsReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**V2KvLedgerErasureShredReceiptDocument**](V2KvLedgerErasureShredReceiptDocument.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersStatus

> V2KvLedgerStatus KVLedgersStatus(ctx, organization, project, storeId, ledgerName).Cursor(cursor).Execute()

Get the chains, archive lag, digest, anchor and verify state of a ledger



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	cursor := "cursor_example" // string | The next_cursor of the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersStatus(context.Background(), organization, project, storeId, ledgerName).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersStatus`: V2KvLedgerStatus
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersStatusRequest struct via the builder pattern


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


## KVLedgersSubjectEntries

> V2KvLedgerSubjectEntries KVLedgersSubjectEntries(ctx, organization, project, storeId, ledgerName).V2KvLedgerSubjectEntriesRequest(v2KvLedgerSubjectEntriesRequest).Execute()

List the ledger entries of one data subject



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	v2KvLedgerSubjectEntriesRequest := *openapiclient.NewV2KvLedgerSubjectEntriesRequest(openapiclient.V2KvLedgerSubjectId{Int64: new(int64)}) // V2KvLedgerSubjectEntriesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersSubjectEntries(context.Background(), organization, project, storeId, ledgerName).V2KvLedgerSubjectEntriesRequest(v2KvLedgerSubjectEntriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersSubjectEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersSubjectEntries`: V2KvLedgerSubjectEntries
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersSubjectEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersSubjectEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2KvLedgerSubjectEntriesRequest** | [**V2KvLedgerSubjectEntriesRequest**](V2KvLedgerSubjectEntriesRequest.md) |  | 

### Return type

[**V2KvLedgerSubjectEntries**](V2KvLedgerSubjectEntries.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersSubjectPreview

> V2KvLedgerErasureSubjectPreview KVLedgersSubjectPreview(ctx, organization, project, storeId, ledgerName).V2KvLedgerErasureSubjectPreviewRequest(v2KvLedgerErasureSubjectPreviewRequest).Execute()

Preview a shred of a data subject



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	v2KvLedgerErasureSubjectPreviewRequest := *openapiclient.NewV2KvLedgerErasureSubjectPreviewRequest(openapiclient.V2KvLedgerSubjectId{Int64: new(int64)}) // V2KvLedgerErasureSubjectPreviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersSubjectPreview(context.Background(), organization, project, storeId, ledgerName).V2KvLedgerErasureSubjectPreviewRequest(v2KvLedgerErasureSubjectPreviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersSubjectPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersSubjectPreview`: V2KvLedgerErasureSubjectPreview
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersSubjectPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersSubjectPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2KvLedgerErasureSubjectPreviewRequest** | [**V2KvLedgerErasureSubjectPreviewRequest**](V2KvLedgerErasureSubjectPreviewRequest.md) |  | 

### Return type

[**V2KvLedgerErasureSubjectPreview**](V2KvLedgerErasureSubjectPreview.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersUpdate

> V2KvLedger KVLedgersUpdate(ctx, organization, project, storeId, ledgerName).V2KvLedgerAdminRaiseRequest(v2KvLedgerAdminRaiseRequest).Execute()

Raise a ledger



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	v2KvLedgerAdminRaiseRequest := *openapiclient.NewV2KvLedgerAdminRaiseRequest() // V2KvLedgerAdminRaiseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersUpdate(context.Background(), organization, project, storeId, ledgerName).V2KvLedgerAdminRaiseRequest(v2KvLedgerAdminRaiseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersUpdate`: V2KvLedger
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2KvLedgerAdminRaiseRequest** | [**V2KvLedgerAdminRaiseRequest**](V2KvLedgerAdminRaiseRequest.md) |  | 

### Return type

[**V2KvLedger**](V2KvLedger.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersVerificationsReport

> V2KvLedgerVerificationReport KVLedgersVerificationsReport(ctx, organization, project, storeId, ledgerName, verificationId).Execute()

Get the signed report of one verification run



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	verificationId := "01M52R6QW07CCAW8KHYZM9210J" // string | The verificationId of a verify answer, or lastVerify.id of the ledger.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersVerificationsReport(context.Background(), organization, project, storeId, ledgerName, verificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersVerificationsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersVerificationsReport`: V2KvLedgerVerificationReport
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersVerificationsReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 
**verificationId** | **string** | The verificationId of a verify answer, or lastVerify.id of the ledger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersVerificationsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**V2KvLedgerVerificationReport**](V2KvLedgerVerificationReport.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLedgersVerify

> V2KvLedgerVerifyResult KVLedgersVerify(ctx, organization, project, storeId, ledgerName).V2KvLedgerVerifyRequest(v2KvLedgerVerifyRequest).Execute()

Verify a ledger against its archive and signed digests



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	ledgerName := "aiUse" // string | 
	v2KvLedgerVerifyRequest := *openapiclient.NewV2KvLedgerVerifyRequest() // V2KvLedgerVerifyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLedgersVerify(context.Background(), organization, project, storeId, ledgerName).V2KvLedgerVerifyRequest(v2KvLedgerVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLedgersVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLedgersVerify`: V2KvLedgerVerifyResult
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLedgersVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 
**ledgerName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLedgersVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **v2KvLedgerVerifyRequest** | [**V2KvLedgerVerifyRequest**](V2KvLedgerVerifyRequest.md) |  | 

### Return type

[**V2KvLedgerVerifyResult**](V2KvLedgerVerifyResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVLinkToProject

> KVLinkToProject200Response KVLinkToProject(ctx, organization, project, storeId).KVLinkToProjectRequest(kVLinkToProjectRequest).Execute()

Link a KV store to another project



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Source project identifier
	storeId := "store-123" // string | KV store identifier
	kVLinkToProjectRequest := *openapiclient.NewKVLinkToProjectRequest(int32(123)) // KVLinkToProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVLinkToProject(context.Background(), organization, project, storeId).KVLinkToProjectRequest(kVLinkToProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVLinkToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVLinkToProject`: KVLinkToProject200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVLinkToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Source project identifier | 
**storeId** | **string** | KV store identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVLinkToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **kVLinkToProjectRequest** | [**KVLinkToProjectRequest**](KVLinkToProjectRequest.md) |  | 

### Return type

[**KVLinkToProject200Response**](KVLinkToProject200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVList

> []V2Store KVList(ctx, organization, project).Execute()

List key-value stores

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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVList(context.Background(), organization, project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVList`: []V2Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V2Store**](V2Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVShow

> V2Store KVShow(ctx, organization, project, storeId).Execute()

Get a kv store



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "0000" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVShow(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVShow`: V2Store
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2Store**](V2Store.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVUnlinkFromProject

> KVLinkToProject200Response KVUnlinkFromProject(ctx, organization, project, storeId).Execute()

Unlink a KV store from this project



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "store-123" // string | KV store identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVUnlinkFromProject(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVUnlinkFromProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVUnlinkFromProject`: KVLinkToProject200Response
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVUnlinkFromProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVUnlinkFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**KVLinkToProject200Response**](KVLinkToProject200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVUpgradeCreate

> V2KvUpgradeResponse KVUpgradeCreate(ctx, organization, project, storeId).V2KvUpgradeRequest(v2KvUpgradeRequest).Execute()

Start, roll back, finalize or resume a format 1 to format 2 upgrade



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list
	v2KvUpgradeRequest := *openapiclient.NewV2KvUpgradeRequest("Action_example") // V2KvUpgradeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVUpgradeCreate(context.Background(), organization, project, storeId).V2KvUpgradeRequest(v2KvUpgradeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVUpgradeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVUpgradeCreate`: V2KvUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVUpgradeCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVUpgradeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **v2KvUpgradeRequest** | [**V2KvUpgradeRequest**](V2KvUpgradeRequest.md) |  | 

### Return type

[**V2KvUpgradeResponse**](V2KvUpgradeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVUpgradeShow

> V2KvUpgradeState KVUpgradeShow(ctx, organization, project, storeId).Execute()

Get the store format and the latest upgrade



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVUpgradeShow(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVUpgradeShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVUpgradeShow`: V2KvUpgradeState
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVUpgradeShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVUpgradeShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2KvUpgradeState**](V2KvUpgradeState.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KVUsageShow

> V2KvUsage KVUsageShow(ctx, organization, project, storeId).Execute()

Get 30-day usage and a cost estimate



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
	organization := "test-org" // string | Organization identifier
	project := "test-project" // string | Project identifier
	storeId := "test-org-test-project-orders" // string | KV store id, as returned by KV_list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KVAPI.KVUsageShow(context.Background(), organization, project, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KVAPI.KVUsageShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KVUsageShow`: V2KvUsage
	fmt.Fprintf(os.Stdout, "Response from `KVAPI.KVUsageShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organization** | **string** | Organization identifier | 
**project** | **string** | Project identifier | 
**storeId** | **string** | KV store id, as returned by KV_list | 

### Other Parameters

Other parameters are passed through a pointer to a apiKVUsageShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**V2KvUsage**](V2KvUsage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

