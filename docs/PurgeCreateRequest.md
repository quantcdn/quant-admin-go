# PurgeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheKeys** | **[]string** | Cache keys to purge | 
**Urls** | Pointer to **[]string** | URLs to purge | [optional] 
**Scope** | **string** | Purge scope (org or project) | [default to "project"]
**Soft** | Pointer to **bool** | Soft purge | [optional] [default to true]

## Methods

### NewPurgeCreateRequest

`func NewPurgeCreateRequest(cacheKeys []string, scope string, ) *PurgeCreateRequest`

NewPurgeCreateRequest instantiates a new PurgeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeCreateRequestWithDefaults

`func NewPurgeCreateRequestWithDefaults() *PurgeCreateRequest`

NewPurgeCreateRequestWithDefaults instantiates a new PurgeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheKeys

`func (o *PurgeCreateRequest) GetCacheKeys() []string`

GetCacheKeys returns the CacheKeys field if non-nil, zero value otherwise.

### GetCacheKeysOk

`func (o *PurgeCreateRequest) GetCacheKeysOk() (*[]string, bool)`

GetCacheKeysOk returns a tuple with the CacheKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheKeys

`func (o *PurgeCreateRequest) SetCacheKeys(v []string)`

SetCacheKeys sets CacheKeys field to given value.


### GetUrls

`func (o *PurgeCreateRequest) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *PurgeCreateRequest) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *PurgeCreateRequest) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *PurgeCreateRequest) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetScope

`func (o *PurgeCreateRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PurgeCreateRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PurgeCreateRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSoft

`func (o *PurgeCreateRequest) GetSoft() bool`

GetSoft returns the Soft field if non-nil, zero value otherwise.

### GetSoftOk

`func (o *PurgeCreateRequest) GetSoftOk() (*bool, bool)`

GetSoftOk returns a tuple with the Soft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoft

`func (o *PurgeCreateRequest) SetSoft(v bool)`

SetSoft sets Soft field to given value.

### HasSoft

`func (o *PurgeCreateRequest) HasSoft() bool`

HasSoft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


