# CiCacheDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | true clears the whole cache of the pool. Must be the only field. | [optional] 
**RepoId** | Pointer to **int64** | GitHub repository id. A digit string is also accepted. | [optional] 
**Ref** | Pointer to **string** | Git ref; starts with refs/, a valid git ref name without %. | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The exact cache key. Deletes only that entry, not keys that start with it. | [optional] 

## Methods

### NewCiCacheDeleteRequest

`func NewCiCacheDeleteRequest() *CiCacheDeleteRequest`

NewCiCacheDeleteRequest instantiates a new CiCacheDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiCacheDeleteRequestWithDefaults

`func NewCiCacheDeleteRequestWithDefaults() *CiCacheDeleteRequest`

NewCiCacheDeleteRequestWithDefaults instantiates a new CiCacheDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *CiCacheDeleteRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CiCacheDeleteRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CiCacheDeleteRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *CiCacheDeleteRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetRepoId

`func (o *CiCacheDeleteRequest) GetRepoId() int64`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *CiCacheDeleteRequest) GetRepoIdOk() (*int64, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *CiCacheDeleteRequest) SetRepoId(v int64)`

SetRepoId sets RepoId field to given value.

### HasRepoId

`func (o *CiCacheDeleteRequest) HasRepoId() bool`

HasRepoId returns a boolean if a field has been set.

### GetRef

`func (o *CiCacheDeleteRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CiCacheDeleteRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CiCacheDeleteRequest) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CiCacheDeleteRequest) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetVersion

`func (o *CiCacheDeleteRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CiCacheDeleteRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CiCacheDeleteRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CiCacheDeleteRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetKey

`func (o *CiCacheDeleteRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CiCacheDeleteRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CiCacheDeleteRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CiCacheDeleteRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


