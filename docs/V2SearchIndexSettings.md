# V2SearchIndexSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facets** | Pointer to **[]string** | Fields offered as facets | [optional] 
**SearchableAttributes** | Pointer to **[]string** | Fields a query matches against | [optional] 
**SortableAttributes** | Pointer to **[]string** | Fields results may be sorted by | [optional] 
**CustomRanking** | Pointer to **[]string** | Ranking rules, e.g. asc(title) | [optional] 
**Synonyms** | Pointer to **string** | Newline-separated synonym groups, each a comma-separated list of terms | [optional] 

## Methods

### NewV2SearchIndexSettings

`func NewV2SearchIndexSettings() *V2SearchIndexSettings`

NewV2SearchIndexSettings instantiates a new V2SearchIndexSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2SearchIndexSettingsWithDefaults

`func NewV2SearchIndexSettingsWithDefaults() *V2SearchIndexSettings`

NewV2SearchIndexSettingsWithDefaults instantiates a new V2SearchIndexSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacets

`func (o *V2SearchIndexSettings) GetFacets() []string`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *V2SearchIndexSettings) GetFacetsOk() (*[]string, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *V2SearchIndexSettings) SetFacets(v []string)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *V2SearchIndexSettings) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### GetSearchableAttributes

`func (o *V2SearchIndexSettings) GetSearchableAttributes() []string`

GetSearchableAttributes returns the SearchableAttributes field if non-nil, zero value otherwise.

### GetSearchableAttributesOk

`func (o *V2SearchIndexSettings) GetSearchableAttributesOk() (*[]string, bool)`

GetSearchableAttributesOk returns a tuple with the SearchableAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableAttributes

`func (o *V2SearchIndexSettings) SetSearchableAttributes(v []string)`

SetSearchableAttributes sets SearchableAttributes field to given value.

### HasSearchableAttributes

`func (o *V2SearchIndexSettings) HasSearchableAttributes() bool`

HasSearchableAttributes returns a boolean if a field has been set.

### GetSortableAttributes

`func (o *V2SearchIndexSettings) GetSortableAttributes() []string`

GetSortableAttributes returns the SortableAttributes field if non-nil, zero value otherwise.

### GetSortableAttributesOk

`func (o *V2SearchIndexSettings) GetSortableAttributesOk() (*[]string, bool)`

GetSortableAttributesOk returns a tuple with the SortableAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortableAttributes

`func (o *V2SearchIndexSettings) SetSortableAttributes(v []string)`

SetSortableAttributes sets SortableAttributes field to given value.

### HasSortableAttributes

`func (o *V2SearchIndexSettings) HasSortableAttributes() bool`

HasSortableAttributes returns a boolean if a field has been set.

### GetCustomRanking

`func (o *V2SearchIndexSettings) GetCustomRanking() []string`

GetCustomRanking returns the CustomRanking field if non-nil, zero value otherwise.

### GetCustomRankingOk

`func (o *V2SearchIndexSettings) GetCustomRankingOk() (*[]string, bool)`

GetCustomRankingOk returns a tuple with the CustomRanking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRanking

`func (o *V2SearchIndexSettings) SetCustomRanking(v []string)`

SetCustomRanking sets CustomRanking field to given value.

### HasCustomRanking

`func (o *V2SearchIndexSettings) HasCustomRanking() bool`

HasCustomRanking returns a boolean if a field has been set.

### GetSynonyms

`func (o *V2SearchIndexSettings) GetSynonyms() string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *V2SearchIndexSettings) GetSynonymsOk() (*string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *V2SearchIndexSettings) SetSynonyms(v string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *V2SearchIndexSettings) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


