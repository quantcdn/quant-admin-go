# V1SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **map[string]interface{}** | Search configuration for the index | [optional] 
**Index** | Pointer to **map[string]interface{}** | Detail related to index size and status | [optional] 

## Methods

### NewV1SearchResponse

`func NewV1SearchResponse() *V1SearchResponse`

NewV1SearchResponse instantiates a new V1SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SearchResponseWithDefaults

`func NewV1SearchResponseWithDefaults() *V1SearchResponse`

NewV1SearchResponseWithDefaults instantiates a new V1SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *V1SearchResponse) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *V1SearchResponse) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *V1SearchResponse) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *V1SearchResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetIndex

`func (o *V1SearchResponse) GetIndex() map[string]interface{}`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *V1SearchResponse) GetIndexOk() (*map[string]interface{}, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *V1SearchResponse) SetIndex(v map[string]interface{})`

SetIndex sets Index field to given value.

### HasIndex

`func (o *V1SearchResponse) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


