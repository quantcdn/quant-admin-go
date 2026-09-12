# ListMcpServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayBaseUrl** | Pointer to **string** | Origin of the MCP gateway host (scheme://host[:port]) — clients treat gateway URLs on this origin as bearer-authenticated with the end user&#39;s Quant token | [optional] 
**Servers** | Pointer to [**[]ListMcpServers200ResponseServersInner**](ListMcpServers200ResponseServersInner.md) |  | [optional] 

## Methods

### NewListMcpServers200Response

`func NewListMcpServers200Response() *ListMcpServers200Response`

NewListMcpServers200Response instantiates a new ListMcpServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMcpServers200ResponseWithDefaults

`func NewListMcpServers200ResponseWithDefaults() *ListMcpServers200Response`

NewListMcpServers200ResponseWithDefaults instantiates a new ListMcpServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayBaseUrl

`func (o *ListMcpServers200Response) GetGatewayBaseUrl() string`

GetGatewayBaseUrl returns the GatewayBaseUrl field if non-nil, zero value otherwise.

### GetGatewayBaseUrlOk

`func (o *ListMcpServers200Response) GetGatewayBaseUrlOk() (*string, bool)`

GetGatewayBaseUrlOk returns a tuple with the GatewayBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayBaseUrl

`func (o *ListMcpServers200Response) SetGatewayBaseUrl(v string)`

SetGatewayBaseUrl sets GatewayBaseUrl field to given value.

### HasGatewayBaseUrl

`func (o *ListMcpServers200Response) HasGatewayBaseUrl() bool`

HasGatewayBaseUrl returns a boolean if a field has been set.

### GetServers

`func (o *ListMcpServers200Response) GetServers() []ListMcpServers200ResponseServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ListMcpServers200Response) GetServersOk() (*[]ListMcpServers200ResponseServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ListMcpServers200Response) SetServers(v []ListMcpServers200ResponseServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ListMcpServers200Response) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


