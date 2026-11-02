# \AISlackBotsAPI

All URIs are relative to *https://dashboard.quantcdn.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSlackInteraction**](AISlackBotsAPI.md#CancelSlackInteraction) | **Post** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/interactions/{interactionId}/cancel | Cancel an unanswered Slack interaction
[**CreateSlackBot**](AISlackBotsAPI.md#CreateSlackBot) | **Post** /api/v3/organizations/{organisation}/ai/slack-bots | Create Slack Bot
[**CreateSlackInteraction**](AISlackBotsAPI.md#CreateSlackInteraction) | **Post** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/interactions | Ask for a single Slack button response
[**DeleteSlackBot**](AISlackBotsAPI.md#DeleteSlackBot) | **Delete** /api/v3/organizations/{organisation}/ai/slack-bots/{botId} | Delete Slack Bot
[**GetSlackBot**](AISlackBotsAPI.md#GetSlackBot) | **Get** /api/v3/organizations/{organisation}/ai/slack-bots/{botId} | Get Slack Bot
[**GetSlackBotThread**](AISlackBotsAPI.md#GetSlackBotThread) | **Get** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/threads/{channel}/{threadTs} | Read a thread
[**GetSlackInteraction**](AISlackBotsAPI.md#GetSlackInteraction) | **Get** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/interactions/{interactionId} | Read an external Slack interaction
[**ListSlackBotChannelMessages**](AISlackBotsAPI.md#ListSlackBotChannelMessages) | **Get** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/channels/{channel}/messages | Read recent channel messages
[**ListSlackBots**](AISlackBotsAPI.md#ListSlackBots) | **Get** /api/v3/organizations/{organisation}/ai/slack-bots | List Slack Bots
[**ListSlackBotsApiEnabled**](AISlackBotsAPI.md#ListSlackBotsApiEnabled) | **Get** /api/v3/organizations/{organisation}/ai/slack-bots/api-enabled | List API-enabled Slack bots
[**PostSlackBotFile**](AISlackBotsAPI.md#PostSlackBotFile) | **Post** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/files | Post an image as the bot
[**PostSlackBotMessage**](AISlackBotsAPI.md#PostSlackBotMessage) | **Post** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/messages | Post a message as the bot
[**SearchSlackWorkspaceChannels**](AISlackBotsAPI.md#SearchSlackWorkspaceChannels) | **Get** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/workspace/channels | Search Slack Workspace Channels
[**SearchSlackWorkspaceUsers**](AISlackBotsAPI.md#SearchSlackWorkspaceUsers) | **Get** /api/v3/organizations/{organisation}/ai/slack-bots/{botId}/workspace/users | Search Slack Workspace Users
[**UpdateSlackBot**](AISlackBotsAPI.md#UpdateSlackBot) | **Put** /api/v3/organizations/{organisation}/ai/slack-bots/{botId} | Update Slack Bot



## CancelSlackInteraction

> CancelSlackInteraction(ctx, organisation, botId, interactionId).Execute()

Cancel an unanswered Slack interaction

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
	botId := "botId_example" // string | 
	interactionId := "interactionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISlackBotsAPI.CancelSlackInteraction(context.Background(), organisation, botId, interactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.CancelSlackInteraction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**botId** | **string** |  | 
**interactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSlackInteractionRequest struct via the builder pattern


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


## CreateSlackBot

> CreateSlackBot201Response CreateSlackBot(ctx, organisation).CreateSlackBotRequest(createSlackBotRequest).Execute()

Create Slack Bot



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
	createSlackBotRequest := *openapiclient.NewCreateSlackBotRequest("Name_example", "SetupType_example", "SystemPrompt_example", "ModelId_example") // CreateSlackBotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.CreateSlackBot(context.Background(), organisation).CreateSlackBotRequest(createSlackBotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.CreateSlackBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSlackBot`: CreateSlackBot201Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.CreateSlackBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlackBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSlackBotRequest** | [**CreateSlackBotRequest**](CreateSlackBotRequest.md) |  | 

### Return type

[**CreateSlackBot201Response**](CreateSlackBot201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSlackInteraction

> CreateSlackInteraction(ctx, organisation, botId).IdempotencyKey(idempotencyKey).CreateSlackInteractionRequest(createSlackInteractionRequest).Execute()

Ask for a single Slack button response



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
	botId := "botId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | 
	createSlackInteractionRequest := *openapiclient.NewCreateSlackInteractionRequest("Channel_example", "Text_example", []openapiclient.CreateSlackInteractionRequestActionsInner{*openapiclient.NewCreateSlackInteractionRequestActionsInner("Id_example", "Label_example")}) // CreateSlackInteractionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISlackBotsAPI.CreateSlackInteraction(context.Background(), organisation, botId).IdempotencyKey(idempotencyKey).CreateSlackInteractionRequest(createSlackInteractionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.CreateSlackInteraction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**botId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlackInteractionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** |  | 
 **createSlackInteractionRequest** | [**CreateSlackInteractionRequest**](CreateSlackInteractionRequest.md) |  | 

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


## DeleteSlackBot

> DeleteSlackBot200Response DeleteSlackBot(ctx, organisation, botId).Execute()

Delete Slack Bot



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
	botId := "botId_example" // string | The Slack bot ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.DeleteSlackBot(context.Background(), organisation, botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.DeleteSlackBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSlackBot`: DeleteSlackBot200Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.DeleteSlackBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**botId** | **string** | The Slack bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSlackBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeleteSlackBot200Response**](DeleteSlackBot200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackBot

> GetSlackBot200Response GetSlackBot(ctx, organisation, botId).Execute()

Get Slack Bot



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
	botId := "botId_example" // string | The Slack bot ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.GetSlackBot(context.Background(), organisation, botId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.GetSlackBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlackBot`: GetSlackBot200Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.GetSlackBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**botId** | **string** | The Slack bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSlackBot200Response**](GetSlackBot200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackBotThread

> GetSlackBotThread200Response GetSlackBotThread(ctx, organisation, botId, channel, threadTs).Since(since).Limit(limit).Cursor(cursor).Execute()

Read a thread



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
	botId := "botId_example" // string | 
	channel := "channel_example" // string | 
	threadTs := "threadTs_example" // string | 
	since := "since_example" // string | Slack ts, exclusive. Only replies after it; the root is always included. (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.GetSlackBotThread(context.Background(), organisation, botId, channel, threadTs).Since(since).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.GetSlackBotThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSlackBotThread`: GetSlackBotThread200Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.GetSlackBotThread`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**botId** | **string** |  | 
**channel** | **string** |  | 
**threadTs** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackBotThreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **since** | **string** | Slack ts, exclusive. Only replies after it; the root is always included. | 
 **limit** | **int32** |  | [default to 100]
 **cursor** | **string** |  | 

### Return type

[**GetSlackBotThread200Response**](GetSlackBotThread200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSlackInteraction

> GetSlackInteraction(ctx, organisation, botId, interactionId).Execute()

Read an external Slack interaction



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
	botId := "botId_example" // string | 
	interactionId := "interactionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AISlackBotsAPI.GetSlackInteraction(context.Background(), organisation, botId, interactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.GetSlackInteraction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** |  | 
**botId** | **string** |  | 
**interactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlackInteractionRequest struct via the builder pattern


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


## ListSlackBotChannelMessages

> ListSlackBotChannelMessages200Response ListSlackBotChannelMessages(ctx, organisation, botId, channel).Since(since).Limit(limit).Cursor(cursor).Execute()

Read recent channel messages



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
	botId := "botId_example" // string | 
	channel := "channel_example" // string | 
	since := "since_example" // string | Slack ts, exclusive (optional)
	limit := int32(56) // int32 |  (optional) (default to 20)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.ListSlackBotChannelMessages(context.Background(), organisation, botId, channel).Since(since).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.ListSlackBotChannelMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlackBotChannelMessages`: ListSlackBotChannelMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.ListSlackBotChannelMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**botId** | **string** |  | 
**channel** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlackBotChannelMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **since** | **string** | Slack ts, exclusive | 
 **limit** | **int32** |  | [default to 20]
 **cursor** | **string** |  | 

### Return type

[**ListSlackBotChannelMessages200Response**](ListSlackBotChannelMessages200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlackBots

> ListSlackBots200Response ListSlackBots(ctx, organisation).Execute()

List Slack Bots



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
	resp, r, err := apiClient.AISlackBotsAPI.ListSlackBots(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.ListSlackBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlackBots`: ListSlackBots200Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.ListSlackBots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlackBotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSlackBots200Response**](ListSlackBots200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSlackBotsApiEnabled

> ListSlackBotsApiEnabled200Response ListSlackBotsApiEnabled(ctx, organisation).Execute()

List API-enabled Slack bots



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
	resp, r, err := apiClient.AISlackBotsAPI.ListSlackBotsApiEnabled(context.Background(), organisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.ListSlackBotsApiEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSlackBotsApiEnabled`: ListSlackBotsApiEnabled200Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.ListSlackBotsApiEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSlackBotsApiEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSlackBotsApiEnabled200Response**](ListSlackBotsApiEnabled200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSlackBotFile

> PostSlackBotFile201Response PostSlackBotFile(ctx, organisation, botId).PostSlackBotFileRequest(postSlackBotFileRequest).IdempotencyKey(idempotencyKey).Execute()

Post an image as the bot



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
	botId := "botId_example" // string | 
	postSlackBotFileRequest := *openapiclient.NewPostSlackBotFileRequest("Channel_example", "Content_example") // PostSlackBotFileRequest | 
	idempotencyKey := "idempotencyKey_example" // string | 8-128 visible ASCII characters. The same key and body within 10 minutes returns the first result and does not upload again. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.PostSlackBotFile(context.Background(), organisation, botId).PostSlackBotFileRequest(postSlackBotFileRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.PostSlackBotFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSlackBotFile`: PostSlackBotFile201Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.PostSlackBotFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**botId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSlackBotFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postSlackBotFileRequest** | [**PostSlackBotFileRequest**](PostSlackBotFileRequest.md) |  | 
 **idempotencyKey** | **string** | 8-128 visible ASCII characters. The same key and body within 10 minutes returns the first result and does not upload again. | 

### Return type

[**PostSlackBotFile201Response**](PostSlackBotFile201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSlackBotMessage

> PostSlackBotMessage201Response PostSlackBotMessage(ctx, organisation, botId).PostSlackBotMessageRequest(postSlackBotMessageRequest).IdempotencyKey(idempotencyKey).Execute()

Post a message as the bot



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
	botId := "botId_example" // string | 
	postSlackBotMessageRequest := *openapiclient.NewPostSlackBotMessageRequest("Channel_example", "Text_example") // PostSlackBotMessageRequest | 
	idempotencyKey := "idempotencyKey_example" // string | 8-128 visible ASCII characters. An empty Idempotency-Key header is refused with 400 INVALID_REQUEST. The same key and body within 10 minutes returns the first result and does not post again. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.PostSlackBotMessage(context.Background(), organisation, botId).PostSlackBotMessageRequest(postSlackBotMessageRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.PostSlackBotMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSlackBotMessage`: PostSlackBotMessage201Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.PostSlackBotMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation machine name | 
**botId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSlackBotMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postSlackBotMessageRequest** | [**PostSlackBotMessageRequest**](PostSlackBotMessageRequest.md) |  | 
 **idempotencyKey** | **string** | 8-128 visible ASCII characters. An empty Idempotency-Key header is refused with 400 INVALID_REQUEST. The same key and body within 10 minutes returns the first result and does not post again. | 

### Return type

[**PostSlackBotMessage201Response**](PostSlackBotMessage201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSlackWorkspaceChannels

> SearchSlackWorkspaceChannels200Response SearchSlackWorkspaceChannels(ctx, organisation, botId).Q(q).Execute()

Search Slack Workspace Channels



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
	botId := "botId_example" // string | The Slack bot ID
	q := "q_example" // string | Search query to filter channels by name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.SearchSlackWorkspaceChannels(context.Background(), organisation, botId).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.SearchSlackWorkspaceChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSlackWorkspaceChannels`: SearchSlackWorkspaceChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.SearchSlackWorkspaceChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**botId** | **string** | The Slack bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSlackWorkspaceChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Search query to filter channels by name | 

### Return type

[**SearchSlackWorkspaceChannels200Response**](SearchSlackWorkspaceChannels200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSlackWorkspaceUsers

> SearchSlackWorkspaceUsers200Response SearchSlackWorkspaceUsers(ctx, organisation, botId).Q(q).Execute()

Search Slack Workspace Users



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
	botId := "botId_example" // string | The Slack bot ID
	q := "q_example" // string | Search query to filter users by name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.SearchSlackWorkspaceUsers(context.Background(), organisation, botId).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.SearchSlackWorkspaceUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSlackWorkspaceUsers`: SearchSlackWorkspaceUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.SearchSlackWorkspaceUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**botId** | **string** | The Slack bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSlackWorkspaceUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Search query to filter users by name | 

### Return type

[**SearchSlackWorkspaceUsers200Response**](SearchSlackWorkspaceUsers200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSlackBot

> CreateSlackBot201Response UpdateSlackBot(ctx, organisation, botId).UpdateSlackBotRequest(updateSlackBotRequest).Execute()

Update Slack Bot



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
	botId := "botId_example" // string | The Slack bot ID
	updateSlackBotRequest := *openapiclient.NewUpdateSlackBotRequest() // UpdateSlackBotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISlackBotsAPI.UpdateSlackBot(context.Background(), organisation, botId).UpdateSlackBotRequest(updateSlackBotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISlackBotsAPI.UpdateSlackBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSlackBot`: CreateSlackBot201Response
	fmt.Fprintf(os.Stdout, "Response from `AISlackBotsAPI.UpdateSlackBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisation** | **string** | The organisation ID | 
**botId** | **string** | The Slack bot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSlackBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSlackBotRequest** | [**UpdateSlackBotRequest**](UpdateSlackBotRequest.md) |  | 

### Return type

[**CreateSlackBot201Response**](CreateSlackBot201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

