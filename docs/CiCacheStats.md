# CiCacheStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** |  | 
**Hits** | **float64** | Cache lookups that found an entry. | 
**Misses** | **float64** | Cache lookups that found no entry. | 
**Saves** | **float64** | Entries saved. | 
**BytesIn** | **float64** | Bytes saved to the cache. | 
**BytesOut** | **float64** | Bytes restored from the cache. | 
**HitRate** | **NullableFloat64** | hits / (hits + misses); null when both are 0. | 

## Methods

### NewCiCacheStats

`func NewCiCacheStats(days int32, hits float64, misses float64, saves float64, bytesIn float64, bytesOut float64, hitRate NullableFloat64, ) *CiCacheStats`

NewCiCacheStats instantiates a new CiCacheStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiCacheStatsWithDefaults

`func NewCiCacheStatsWithDefaults() *CiCacheStats`

NewCiCacheStatsWithDefaults instantiates a new CiCacheStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *CiCacheStats) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *CiCacheStats) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *CiCacheStats) SetDays(v int32)`

SetDays sets Days field to given value.


### GetHits

`func (o *CiCacheStats) GetHits() float64`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *CiCacheStats) GetHitsOk() (*float64, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *CiCacheStats) SetHits(v float64)`

SetHits sets Hits field to given value.


### GetMisses

`func (o *CiCacheStats) GetMisses() float64`

GetMisses returns the Misses field if non-nil, zero value otherwise.

### GetMissesOk

`func (o *CiCacheStats) GetMissesOk() (*float64, bool)`

GetMissesOk returns a tuple with the Misses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMisses

`func (o *CiCacheStats) SetMisses(v float64)`

SetMisses sets Misses field to given value.


### GetSaves

`func (o *CiCacheStats) GetSaves() float64`

GetSaves returns the Saves field if non-nil, zero value otherwise.

### GetSavesOk

`func (o *CiCacheStats) GetSavesOk() (*float64, bool)`

GetSavesOk returns a tuple with the Saves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaves

`func (o *CiCacheStats) SetSaves(v float64)`

SetSaves sets Saves field to given value.


### GetBytesIn

`func (o *CiCacheStats) GetBytesIn() float64`

GetBytesIn returns the BytesIn field if non-nil, zero value otherwise.

### GetBytesInOk

`func (o *CiCacheStats) GetBytesInOk() (*float64, bool)`

GetBytesInOk returns a tuple with the BytesIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesIn

`func (o *CiCacheStats) SetBytesIn(v float64)`

SetBytesIn sets BytesIn field to given value.


### GetBytesOut

`func (o *CiCacheStats) GetBytesOut() float64`

GetBytesOut returns the BytesOut field if non-nil, zero value otherwise.

### GetBytesOutOk

`func (o *CiCacheStats) GetBytesOutOk() (*float64, bool)`

GetBytesOutOk returns a tuple with the BytesOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOut

`func (o *CiCacheStats) SetBytesOut(v float64)`

SetBytesOut sets BytesOut field to given value.


### GetHitRate

`func (o *CiCacheStats) GetHitRate() float64`

GetHitRate returns the HitRate field if non-nil, zero value otherwise.

### GetHitRateOk

`func (o *CiCacheStats) GetHitRateOk() (*float64, bool)`

GetHitRateOk returns a tuple with the HitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitRate

`func (o *CiCacheStats) SetHitRate(v float64)`

SetHitRate sets HitRate field to given value.


### SetHitRateNil

`func (o *CiCacheStats) SetHitRateNil(b bool)`

 SetHitRateNil sets the value for HitRate to be an explicit nil

### UnsetHitRate
`func (o *CiCacheStats) UnsetHitRate()`

UnsetHitRate ensures that no value is present for HitRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


