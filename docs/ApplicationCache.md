# ApplicationCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheEndpoint** | Pointer to **NullableString** | Cache cluster endpoint | [optional] 
**CacheIdentifier** | Pointer to **NullableString** | Cache cluster identifier | [optional] 
**DataStorageMaxGb** | Pointer to **NullableInt32** | Maximum cache storage in GB | [optional] 

## Methods

### NewApplicationCache

`func NewApplicationCache() *ApplicationCache`

NewApplicationCache instantiates a new ApplicationCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCacheWithDefaults

`func NewApplicationCacheWithDefaults() *ApplicationCache`

NewApplicationCacheWithDefaults instantiates a new ApplicationCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheEndpoint

`func (o *ApplicationCache) GetCacheEndpoint() string`

GetCacheEndpoint returns the CacheEndpoint field if non-nil, zero value otherwise.

### GetCacheEndpointOk

`func (o *ApplicationCache) GetCacheEndpointOk() (*string, bool)`

GetCacheEndpointOk returns a tuple with the CacheEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheEndpoint

`func (o *ApplicationCache) SetCacheEndpoint(v string)`

SetCacheEndpoint sets CacheEndpoint field to given value.

### HasCacheEndpoint

`func (o *ApplicationCache) HasCacheEndpoint() bool`

HasCacheEndpoint returns a boolean if a field has been set.

### SetCacheEndpointNil

`func (o *ApplicationCache) SetCacheEndpointNil(b bool)`

 SetCacheEndpointNil sets the value for CacheEndpoint to be an explicit nil

### UnsetCacheEndpoint
`func (o *ApplicationCache) UnsetCacheEndpoint()`

UnsetCacheEndpoint ensures that no value is present for CacheEndpoint, not even an explicit nil
### GetCacheIdentifier

`func (o *ApplicationCache) GetCacheIdentifier() string`

GetCacheIdentifier returns the CacheIdentifier field if non-nil, zero value otherwise.

### GetCacheIdentifierOk

`func (o *ApplicationCache) GetCacheIdentifierOk() (*string, bool)`

GetCacheIdentifierOk returns a tuple with the CacheIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheIdentifier

`func (o *ApplicationCache) SetCacheIdentifier(v string)`

SetCacheIdentifier sets CacheIdentifier field to given value.

### HasCacheIdentifier

`func (o *ApplicationCache) HasCacheIdentifier() bool`

HasCacheIdentifier returns a boolean if a field has been set.

### SetCacheIdentifierNil

`func (o *ApplicationCache) SetCacheIdentifierNil(b bool)`

 SetCacheIdentifierNil sets the value for CacheIdentifier to be an explicit nil

### UnsetCacheIdentifier
`func (o *ApplicationCache) UnsetCacheIdentifier()`

UnsetCacheIdentifier ensures that no value is present for CacheIdentifier, not even an explicit nil
### GetDataStorageMaxGb

`func (o *ApplicationCache) GetDataStorageMaxGb() int32`

GetDataStorageMaxGb returns the DataStorageMaxGb field if non-nil, zero value otherwise.

### GetDataStorageMaxGbOk

`func (o *ApplicationCache) GetDataStorageMaxGbOk() (*int32, bool)`

GetDataStorageMaxGbOk returns a tuple with the DataStorageMaxGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStorageMaxGb

`func (o *ApplicationCache) SetDataStorageMaxGb(v int32)`

SetDataStorageMaxGb sets DataStorageMaxGb field to given value.

### HasDataStorageMaxGb

`func (o *ApplicationCache) HasDataStorageMaxGb() bool`

HasDataStorageMaxGb returns a boolean if a field has been set.

### SetDataStorageMaxGbNil

`func (o *ApplicationCache) SetDataStorageMaxGbNil(b bool)`

 SetDataStorageMaxGbNil sets the value for DataStorageMaxGb to be an explicit nil

### UnsetDataStorageMaxGb
`func (o *ApplicationCache) UnsetDataStorageMaxGb()`

UnsetDataStorageMaxGb ensures that no value is present for DataStorageMaxGb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


