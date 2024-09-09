# PurgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheKeys** | **[]string** |  | 
**Scope** | **string** |  | [default to "project"]

## Methods

### NewPurgeRequest

`func NewPurgeRequest(cacheKeys []string, scope string, ) *PurgeRequest`

NewPurgeRequest instantiates a new PurgeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeRequestWithDefaults

`func NewPurgeRequestWithDefaults() *PurgeRequest`

NewPurgeRequestWithDefaults instantiates a new PurgeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheKeys

`func (o *PurgeRequest) GetCacheKeys() []string`

GetCacheKeys returns the CacheKeys field if non-nil, zero value otherwise.

### GetCacheKeysOk

`func (o *PurgeRequest) GetCacheKeysOk() (*[]string, bool)`

GetCacheKeysOk returns a tuple with the CacheKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheKeys

`func (o *PurgeRequest) SetCacheKeys(v []string)`

SetCacheKeys sets CacheKeys field to given value.


### GetScope

`func (o *PurgeRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PurgeRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PurgeRequest) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


