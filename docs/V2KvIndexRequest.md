# V2KvIndexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique in the store. | 
**Collection** | **string** | A collection, or a pattern where * as one whole segment matches exactly one segment (customer:*:world). At most 512 UTF-8 bytes. | 
**Fields** | Pointer to **[]string** | Top-level fields of a plaintext JSON object value (0-3). A unique index needs at least one. A field that a ledger on the collection marks as personal cannot be indexed (INDEX_FIELD_ENCRYPTED). | [optional] 
**SortBy** | Pointer to **string** | Optional single sort field. | [optional] 
**Unique** | Pointer to **bool** | One key per value. Enables KV_indexes_lookup. Not allowed on a ledger collection (LEDGER_INDEX_KIND). | [optional] [default to false]
**Count** | Pointer to **bool** | Keeps a live count per value. Enables KV_indexes_count. Not allowed on a ledger collection (LEDGER_INDEX_KIND) or on items with a ttl (COUNT_WITH_TTL). | [optional] [default to false]

## Methods

### NewV2KvIndexRequest

`func NewV2KvIndexRequest(name string, collection string, ) *V2KvIndexRequest`

NewV2KvIndexRequest instantiates a new V2KvIndexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvIndexRequestWithDefaults

`func NewV2KvIndexRequestWithDefaults() *V2KvIndexRequest`

NewV2KvIndexRequestWithDefaults instantiates a new V2KvIndexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V2KvIndexRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2KvIndexRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2KvIndexRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCollection

`func (o *V2KvIndexRequest) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *V2KvIndexRequest) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *V2KvIndexRequest) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetFields

`func (o *V2KvIndexRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *V2KvIndexRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *V2KvIndexRequest) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *V2KvIndexRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSortBy

`func (o *V2KvIndexRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *V2KvIndexRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *V2KvIndexRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *V2KvIndexRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetUnique

`func (o *V2KvIndexRequest) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *V2KvIndexRequest) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *V2KvIndexRequest) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *V2KvIndexRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetCount

`func (o *V2KvIndexRequest) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *V2KvIndexRequest) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *V2KvIndexRequest) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *V2KvIndexRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


