# V2SearchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extractors** | Pointer to [**V2SearchExtractors**](V2SearchExtractors.md) |  | [optional] 
**Algolia** | Pointer to [**V2SearchIndexSettings**](V2SearchIndexSettings.md) |  | [optional] 

## Methods

### NewV2SearchConfig

`func NewV2SearchConfig() *V2SearchConfig`

NewV2SearchConfig instantiates a new V2SearchConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SearchConfigWithDefaults

`func NewV2SearchConfigWithDefaults() *V2SearchConfig`

NewV2SearchConfigWithDefaults instantiates a new V2SearchConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtractors

`func (o *V2SearchConfig) GetExtractors() V2SearchExtractors`

GetExtractors returns the Extractors field if non-nil, zero value otherwise.

### GetExtractorsOk

`func (o *V2SearchConfig) GetExtractorsOk() (*V2SearchExtractors, bool)`

GetExtractorsOk returns a tuple with the Extractors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractors

`func (o *V2SearchConfig) SetExtractors(v V2SearchExtractors)`

SetExtractors sets Extractors field to given value.

### HasExtractors

`func (o *V2SearchConfig) HasExtractors() bool`

HasExtractors returns a boolean if a field has been set.

### GetAlgolia

`func (o *V2SearchConfig) GetAlgolia() V2SearchIndexSettings`

GetAlgolia returns the Algolia field if non-nil, zero value otherwise.

### GetAlgoliaOk

`func (o *V2SearchConfig) GetAlgoliaOk() (*V2SearchIndexSettings, bool)`

GetAlgoliaOk returns a tuple with the Algolia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgolia

`func (o *V2SearchConfig) SetAlgolia(v V2SearchIndexSettings)`

SetAlgolia sets Algolia field to given value.

### HasAlgolia

`func (o *V2SearchConfig) HasAlgolia() bool`

HasAlgolia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


