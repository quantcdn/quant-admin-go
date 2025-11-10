# V1PingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Configuration associated with a project | [optional] 
**Project** | Pointer to **string** | The project machine name | [optional] 

## Methods

### NewV1PingResponse

`func NewV1PingResponse() *V1PingResponse`

NewV1PingResponse instantiates a new V1PingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PingResponseWithDefaults

`func NewV1PingResponseWithDefaults() *V1PingResponse`

NewV1PingResponseWithDefaults instantiates a new V1PingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *V1PingResponse) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1PingResponse) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1PingResponse) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1PingResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetProject

`func (o *V1PingResponse) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1PingResponse) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1PingResponse) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1PingResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


