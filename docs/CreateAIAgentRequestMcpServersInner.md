# CreateAIAgentRequestMcpServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | **string** | The registered server name | 
**Tools** | Pointer to **[]string** | Optional. Only these tools of the server. Omit for every tool the server allows; an empty array allows none. | [optional] 

## Methods

### NewCreateAIAgentRequestMcpServersInner

`func NewCreateAIAgentRequestMcpServersInner(server string, ) *CreateAIAgentRequestMcpServersInner`

NewCreateAIAgentRequestMcpServersInner instantiates a new CreateAIAgentRequestMcpServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAIAgentRequestMcpServersInnerWithDefaults

`func NewCreateAIAgentRequestMcpServersInnerWithDefaults() *CreateAIAgentRequestMcpServersInner`

NewCreateAIAgentRequestMcpServersInnerWithDefaults instantiates a new CreateAIAgentRequestMcpServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *CreateAIAgentRequestMcpServersInner) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateAIAgentRequestMcpServersInner) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateAIAgentRequestMcpServersInner) SetServer(v string)`

SetServer sets Server field to given value.


### GetTools

`func (o *CreateAIAgentRequestMcpServersInner) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CreateAIAgentRequestMcpServersInner) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CreateAIAgentRequestMcpServersInner) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CreateAIAgentRequestMcpServersInner) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


