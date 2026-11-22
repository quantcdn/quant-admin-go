# CiRunnerPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | **string** | Hard cap on concurrent jobs: small 2, medium 4, large 8. When the pool is full, jobs wait. | 
**Architectures** | Pointer to **[]string** |  | [optional] [default to ["x64"]]
**NetworkMode** | Pointer to **string** |  | [optional] [default to "public"]
**EgressAllowlist** | Pointer to **[]string** | Domains allowed when networkMode is restricted. Wildcard only as a leading *. label. | [optional] 
**IdleTimeoutMinutes** | Pointer to **int32** |  | [optional] [default to 15]
**BusinessHoursWarm** | Pointer to [**CiRunnerPoolConfigBusinessHoursWarm**](CiRunnerPoolConfigBusinessHoursWarm.md) |  | [optional] 
**Connection** | Pointer to [**CiRunnerPoolConfigConnection**](CiRunnerPoolConfigConnection.md) |  | [optional] 
**CacheResidency** | Pointer to **string** | Where the pool&#39;s GitHub Actions build cache may live. au-only keeps cache data in Australia; if the pool&#39;s cache is unavailable, jobs run without a cache. allow-github lets jobs use GitHub&#39;s cache when the pool&#39;s cache is unavailable. When not set, a pool that accepts public repositories uses allow-github and every other pool uses au-only. A change applies to the next job each runner machine starts. | [optional] 
**Alerts** | Pointer to [**[]CiRunnerPoolConfigAlertsInner**](CiRunnerPoolConfigAlertsInner.md) |  | [optional] [readonly] 

## Methods

### NewCiRunnerPoolConfig

`func NewCiRunnerPoolConfig(tier string, ) *CiRunnerPoolConfig`

NewCiRunnerPoolConfig instantiates a new CiRunnerPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiRunnerPoolConfigWithDefaults

`func NewCiRunnerPoolConfigWithDefaults() *CiRunnerPoolConfig`

NewCiRunnerPoolConfigWithDefaults instantiates a new CiRunnerPoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *CiRunnerPoolConfig) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CiRunnerPoolConfig) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CiRunnerPoolConfig) SetTier(v string)`

SetTier sets Tier field to given value.


### GetArchitectures

`func (o *CiRunnerPoolConfig) GetArchitectures() []string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *CiRunnerPoolConfig) GetArchitecturesOk() (*[]string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *CiRunnerPoolConfig) SetArchitectures(v []string)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *CiRunnerPoolConfig) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### GetNetworkMode

`func (o *CiRunnerPoolConfig) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *CiRunnerPoolConfig) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *CiRunnerPoolConfig) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *CiRunnerPoolConfig) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetEgressAllowlist

`func (o *CiRunnerPoolConfig) GetEgressAllowlist() []string`

GetEgressAllowlist returns the EgressAllowlist field if non-nil, zero value otherwise.

### GetEgressAllowlistOk

`func (o *CiRunnerPoolConfig) GetEgressAllowlistOk() (*[]string, bool)`

GetEgressAllowlistOk returns a tuple with the EgressAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressAllowlist

`func (o *CiRunnerPoolConfig) SetEgressAllowlist(v []string)`

SetEgressAllowlist sets EgressAllowlist field to given value.

### HasEgressAllowlist

`func (o *CiRunnerPoolConfig) HasEgressAllowlist() bool`

HasEgressAllowlist returns a boolean if a field has been set.

### GetIdleTimeoutMinutes

`func (o *CiRunnerPoolConfig) GetIdleTimeoutMinutes() int32`

GetIdleTimeoutMinutes returns the IdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutMinutesOk

`func (o *CiRunnerPoolConfig) GetIdleTimeoutMinutesOk() (*int32, bool)`

GetIdleTimeoutMinutesOk returns a tuple with the IdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMinutes

`func (o *CiRunnerPoolConfig) SetIdleTimeoutMinutes(v int32)`

SetIdleTimeoutMinutes sets IdleTimeoutMinutes field to given value.

### HasIdleTimeoutMinutes

`func (o *CiRunnerPoolConfig) HasIdleTimeoutMinutes() bool`

HasIdleTimeoutMinutes returns a boolean if a field has been set.

### GetBusinessHoursWarm

`func (o *CiRunnerPoolConfig) GetBusinessHoursWarm() CiRunnerPoolConfigBusinessHoursWarm`

GetBusinessHoursWarm returns the BusinessHoursWarm field if non-nil, zero value otherwise.

### GetBusinessHoursWarmOk

`func (o *CiRunnerPoolConfig) GetBusinessHoursWarmOk() (*CiRunnerPoolConfigBusinessHoursWarm, bool)`

GetBusinessHoursWarmOk returns a tuple with the BusinessHoursWarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessHoursWarm

`func (o *CiRunnerPoolConfig) SetBusinessHoursWarm(v CiRunnerPoolConfigBusinessHoursWarm)`

SetBusinessHoursWarm sets BusinessHoursWarm field to given value.

### HasBusinessHoursWarm

`func (o *CiRunnerPoolConfig) HasBusinessHoursWarm() bool`

HasBusinessHoursWarm returns a boolean if a field has been set.

### GetConnection

`func (o *CiRunnerPoolConfig) GetConnection() CiRunnerPoolConfigConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *CiRunnerPoolConfig) GetConnectionOk() (*CiRunnerPoolConfigConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *CiRunnerPoolConfig) SetConnection(v CiRunnerPoolConfigConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *CiRunnerPoolConfig) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetCacheResidency

`func (o *CiRunnerPoolConfig) GetCacheResidency() string`

GetCacheResidency returns the CacheResidency field if non-nil, zero value otherwise.

### GetCacheResidencyOk

`func (o *CiRunnerPoolConfig) GetCacheResidencyOk() (*string, bool)`

GetCacheResidencyOk returns a tuple with the CacheResidency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheResidency

`func (o *CiRunnerPoolConfig) SetCacheResidency(v string)`

SetCacheResidency sets CacheResidency field to given value.

### HasCacheResidency

`func (o *CiRunnerPoolConfig) HasCacheResidency() bool`

HasCacheResidency returns a boolean if a field has been set.

### GetAlerts

`func (o *CiRunnerPoolConfig) GetAlerts() []CiRunnerPoolConfigAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *CiRunnerPoolConfig) GetAlertsOk() (*[]CiRunnerPoolConfigAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *CiRunnerPoolConfig) SetAlerts(v []CiRunnerPoolConfigAlertsInner)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *CiRunnerPoolConfig) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


