# ListOrchestrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orchestrations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListOrchestrations200Response

`func NewListOrchestrations200Response() *ListOrchestrations200Response`

NewListOrchestrations200Response instantiates a new ListOrchestrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrchestrations200ResponseWithDefaults

`func NewListOrchestrations200ResponseWithDefaults() *ListOrchestrations200Response`

NewListOrchestrations200ResponseWithDefaults instantiates a new ListOrchestrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrchestrations

`func (o *ListOrchestrations200Response) GetOrchestrations() []map[string]interface{}`

GetOrchestrations returns the Orchestrations field if non-nil, zero value otherwise.

### GetOrchestrationsOk

`func (o *ListOrchestrations200Response) GetOrchestrationsOk() (*[]map[string]interface{}, bool)`

GetOrchestrationsOk returns a tuple with the Orchestrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrations

`func (o *ListOrchestrations200Response) SetOrchestrations(v []map[string]interface{})`

SetOrchestrations sets Orchestrations field to given value.

### HasOrchestrations

`func (o *ListOrchestrations200Response) HasOrchestrations() bool`

HasOrchestrations returns a boolean if a field has been set.

### GetNextCursor

`func (o *ListOrchestrations200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListOrchestrations200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListOrchestrations200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListOrchestrations200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ListOrchestrations200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListOrchestrations200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


