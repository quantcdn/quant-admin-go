# \AISkillsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSkill**](AISkillsAPI.md#CreateSkill) | **Post** /api/v3/organizations/{organisation}/ai/skills | Create Inline Skill
[**DeleteSkill**](AISkillsAPI.md#DeleteSkill) | **Delete** /api/v3/organizations/{organisation}/ai/skills/{skillId} | Delete Skill
[**DeleteSkillCollection**](AISkillsAPI.md#DeleteSkillCollection) | **Delete** /api/v3/organizations/{organisation}/ai/skills/collections/{namespace} | Delete Skill Collection
[**GetSkill**](AISkillsAPI.md#GetSkill) | **Get** /api/v3/organizations/{organisation}/ai/skills/{skillId} | Get Skill Details
[**ImportSkill**](AISkillsAPI.md#ImportSkill) | **Post** /api/v3/organizations/{organisation}/ai/skills/import | Import Skill from External Source
[**ImportSkillCollection**](AISkillsAPI.md#ImportSkillCollection) | **Post** /api/v3/organizations/{organisation}/ai/skills/import-collection | Import Skill Collection from GitHub
[**ListSkillCollections**](AISkillsAPI.md#ListSkillCollections) | **Get** /api/v3/organizations/{organisation}/ai/skills/collections | List Skill Collections
[**ListSkills**](AISkillsAPI.md#ListSkills) | **Get** /api/v3/organizations/{organisation}/ai/skills | List Organization&#39;s Skills
[**SyncSkill**](AISkillsAPI.md#SyncSkill) | **Post** /api/v3/organizations/{organisation}/ai/skills/{skillId}/sync | Sync Skill from Source
[**SyncSkillCollection**](AISkillsAPI.md#SyncSkillCollection) | **Post** /api/v3/organizations/{organisation}/ai/skills/collections/{namespace}/sync | Sync Skill Collection
[**UpdateSkill**](AISkillsAPI.md#UpdateSkill) | **Put** /api/v3/organizations/{organisation}/ai/skills/{skillId} | Update Skill



## CreateSkill

> CreateSkill201Response CreateSkill(ctx, organisation).CreateSkillRequest(createSkillRequest).Execute()

Create Inline Skill



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
	createSkillRequest := *openapiclient.NewCreateSkillRequest("code-review", "You are an expert code reviewer...", "When the user asks for code review") // CreateSkillRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.CreateSkill(context.Background(), organisation).CreateSkillRequest(createSkillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.CreateSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSkill`: CreateSkill201Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.CreateSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSkillRequest** | [**CreateSkillRequest**](CreateSkillRequest.md) |  | 

### Return type

[**CreateSkill201Response**](CreateSkill201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSkill

> DeleteSkill200Response DeleteSkill(ctx, organisation, skillId).Execute()

Delete Skill



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
	skillId := "skillId_example" // string | The skill ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.DeleteSkill(context.Background(), organisation, skillId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.DeleteSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSkill`: DeleteSkill200Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.DeleteSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**skillId** | **string** | The skill ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteSkill200Response**](DeleteSkill200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSkillCollection

> DeleteSkillCollection200Response DeleteSkillCollection(ctx, organisation, namespace).Execute()

Delete Skill Collection



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
	namespace := "namespace_example" // string | Collection namespace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.DeleteSkillCollection(context.Background(), organisation, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.DeleteSkillCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSkillCollection`: DeleteSkillCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.DeleteSkillCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**namespace** | **string** | Collection namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteSkillCollection200Response**](DeleteSkillCollection200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkill

> GetSkill200Response GetSkill(ctx, organisation, skillId).Execute()

Get Skill Details



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
	skillId := "skillId_example" // string | The skill ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.GetSkill(context.Background(), organisation, skillId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.GetSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSkill`: GetSkill200Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.GetSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**skillId** | **string** | The skill ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSkill200Response**](GetSkill200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSkill

> ImportSkill201Response ImportSkill(ctx, organisation).ImportSkillRequest(importSkillRequest).Execute()

Import Skill from External Source



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
	importSkillRequest := *openapiclient.NewImportSkillRequest(*openapiclient.NewImportSkillRequestSource("Type_example")) // ImportSkillRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.ImportSkill(context.Background(), organisation).ImportSkillRequest(importSkillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.ImportSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportSkill`: ImportSkill201Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.ImportSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importSkillRequest** | [**ImportSkillRequest**](ImportSkillRequest.md) |  | 

### Return type

[**ImportSkill201Response**](ImportSkill201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSkillCollection

> ImportSkillCollection201Response ImportSkillCollection(ctx, organisation).ImportSkillCollectionRequest(importSkillCollectionRequest).Execute()

Import Skill Collection from GitHub



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
	importSkillCollectionRequest := *openapiclient.NewImportSkillCollectionRequest("superpowers", *openapiclient.NewImportSkillCollectionRequestSource("Type_example", "obra/superpowers")) // ImportSkillCollectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.ImportSkillCollection(context.Background(), organisation).ImportSkillCollectionRequest(importSkillCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.ImportSkillCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportSkillCollection`: ImportSkillCollection201Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.ImportSkillCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSkillCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importSkillCollectionRequest** | [**ImportSkillCollectionRequest**](ImportSkillCollectionRequest.md) |  | 

### Return type

[**ImportSkillCollection201Response**](ImportSkillCollection201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSkillCollections

> ListSkillCollections200Response ListSkillCollections(ctx, organisation).Execute()

List Skill Collections



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
	resp, r, err := apiClient.AISkillsAPI.ListSkillCollections(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.ListSkillCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSkillCollections`: ListSkillCollections200Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.ListSkillCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSkillCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSkillCollections200Response**](ListSkillCollections200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSkills

> ListSkills200Response ListSkills(ctx, organisation).Tag(tag).Namespace(namespace).Limit(limit).Execute()

List Organization's Skills



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
	tag := "tag_example" // string | Filter skills by tag (optional)
	namespace := "namespace_example" // string | Filter skills by collection namespace (e.g. 'superpowers') (optional)
	limit := int32(56) // int32 | Maximum number of skills to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.ListSkills(context.Background(), organisation).Tag(tag).Namespace(namespace).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.ListSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSkills`: ListSkills200Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.ListSkills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **string** | Filter skills by tag | 
 **namespace** | **string** | Filter skills by collection namespace (e.g. &#39;superpowers&#39;) | 
 **limit** | **int32** | Maximum number of skills to return | [default to 50]

### Return type

[**ListSkills200Response**](ListSkills200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncSkill

> ImportSkill201Response SyncSkill(ctx, organisation, skillId).Execute()

Sync Skill from Source



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
	skillId := "skillId_example" // string | The skill ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.SyncSkill(context.Background(), organisation, skillId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.SyncSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncSkill`: ImportSkill201Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.SyncSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**skillId** | **string** | The skill ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ImportSkill201Response**](ImportSkill201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncSkillCollection

> SyncSkillCollection200Response SyncSkillCollection(ctx, organisation, namespace).Execute()

Sync Skill Collection



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
	namespace := "namespace_example" // string | Collection namespace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.SyncSkillCollection(context.Background(), organisation, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.SyncSkillCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncSkillCollection`: SyncSkillCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.SyncSkillCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**namespace** | **string** | Collection namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncSkillCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SyncSkillCollection200Response**](SyncSkillCollection200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSkill

> UpdateSkill200Response UpdateSkill(ctx, organisation, skillId).UpdateSkillRequest(updateSkillRequest).Execute()

Update Skill



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
	skillId := "skillId_example" // string | The skill ID
	updateSkillRequest := *openapiclient.NewUpdateSkillRequest() // UpdateSkillRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISkillsAPI.UpdateSkill(context.Background(), organisation, skillId).UpdateSkillRequest(updateSkillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISkillsAPI.UpdateSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSkill`: UpdateSkill200Response
	fmt.Fprintf(os.Stdout, "Response from `AISkillsAPI.UpdateSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**skillId** | **string** | The skill ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSkillRequest** | [**UpdateSkillRequest**](UpdateSkillRequest.md) |  | 

### Return type

[**UpdateSkill200Response**](UpdateSkill200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

