# CiCacheEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoId** | **int64** | GitHub repository id. | 
**Repo** | **NullableString** | owner/name recorded for repoId, or null when none is recorded. | 
**Ref** | **string** | The git ref that saved the entry. A pull request saves to refs/pull/&lt;number&gt;/merge. | 
**Key** | **string** | The actions/cache key. | 
**Version** | **string** | The cache version (a hash of the paths and compression the job used). | 
**SizeBytes** | **int64** |  | 
**CreatedAt** | **NullableTime** | When the entry was saved. Entries expire 14 days after they are saved. | 

## Methods

### NewCiCacheEntry

`func NewCiCacheEntry(repoId int64, repo NullableString, ref string, key string, version string, sizeBytes int64, createdAt NullableTime, ) *CiCacheEntry`

NewCiCacheEntry instantiates a new CiCacheEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiCacheEntryWithDefaults

`func NewCiCacheEntryWithDefaults() *CiCacheEntry`

NewCiCacheEntryWithDefaults instantiates a new CiCacheEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoId

`func (o *CiCacheEntry) GetRepoId() int64`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *CiCacheEntry) GetRepoIdOk() (*int64, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *CiCacheEntry) SetRepoId(v int64)`

SetRepoId sets RepoId field to given value.


### GetRepo

`func (o *CiCacheEntry) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CiCacheEntry) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CiCacheEntry) SetRepo(v string)`

SetRepo sets Repo field to given value.


### SetRepoNil

`func (o *CiCacheEntry) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *CiCacheEntry) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetRef

`func (o *CiCacheEntry) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CiCacheEntry) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CiCacheEntry) SetRef(v string)`

SetRef sets Ref field to given value.


### GetKey

`func (o *CiCacheEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CiCacheEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CiCacheEntry) SetKey(v string)`

SetKey sets Key field to given value.


### GetVersion

`func (o *CiCacheEntry) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CiCacheEntry) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CiCacheEntry) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSizeBytes

`func (o *CiCacheEntry) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *CiCacheEntry) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *CiCacheEntry) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.


### GetCreatedAt

`func (o *CiCacheEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CiCacheEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CiCacheEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CiCacheEntry) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CiCacheEntry) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


