# GetGovernanceConfig200ResponseSpendLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonthlyBudgetCents** | Pointer to **NullableInt32** |  | [optional] 
**DailyBudgetCents** | Pointer to **NullableInt32** |  | [optional] 
**PerUserMonthlyBudgetCents** | Pointer to **NullableInt32** |  | [optional] 
**PerUserDailyBudgetCents** | Pointer to **NullableInt32** |  | [optional] 
**WarningThresholdPercent** | Pointer to **NullableInt32** |  | [optional] 
**InterfaceLimits** | Pointer to [**map[string]GetGovernanceConfig200ResponseSpendLimitsInterfaceLimitsValue**](GetGovernanceConfig200ResponseSpendLimitsInterfaceLimitsValue.md) | Aggregate spend caps per interface label (slack, autonomous, api-gateway, streaming, websocket). Keys are interface labels. | [optional] 
**UserOverrides** | Pointer to [**map[string]GetGovernanceConfig200ResponseSpendLimitsUserOverridesValue**](GetGovernanceConfig200ResponseSpendLimitsUserOverridesValue.md) | Per-user budget overrides keyed by userId (numeric portal id, slack-&lt;id&gt;, or system:code-agent). Replaces the flat per-user budget for that user; unlimited&#x3D;true exempts them. | [optional] 

## Methods

### NewGetGovernanceConfig200ResponseSpendLimits

`func NewGetGovernanceConfig200ResponseSpendLimits() *GetGovernanceConfig200ResponseSpendLimits`

NewGetGovernanceConfig200ResponseSpendLimits instantiates a new GetGovernanceConfig200ResponseSpendLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGovernanceConfig200ResponseSpendLimitsWithDefaults

`func NewGetGovernanceConfig200ResponseSpendLimitsWithDefaults() *GetGovernanceConfig200ResponseSpendLimits`

NewGetGovernanceConfig200ResponseSpendLimitsWithDefaults instantiates a new GetGovernanceConfig200ResponseSpendLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonthlyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetMonthlyBudgetCents() int32`

GetMonthlyBudgetCents returns the MonthlyBudgetCents field if non-nil, zero value otherwise.

### GetMonthlyBudgetCentsOk

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetMonthlyBudgetCentsOk() (*int32, bool)`

GetMonthlyBudgetCentsOk returns a tuple with the MonthlyBudgetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetMonthlyBudgetCents(v int32)`

SetMonthlyBudgetCents sets MonthlyBudgetCents field to given value.

### HasMonthlyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) HasMonthlyBudgetCents() bool`

HasMonthlyBudgetCents returns a boolean if a field has been set.

### SetMonthlyBudgetCentsNil

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetMonthlyBudgetCentsNil(b bool)`

 SetMonthlyBudgetCentsNil sets the value for MonthlyBudgetCents to be an explicit nil

### UnsetMonthlyBudgetCents
`func (o *GetGovernanceConfig200ResponseSpendLimits) UnsetMonthlyBudgetCents()`

UnsetMonthlyBudgetCents ensures that no value is present for MonthlyBudgetCents, not even an explicit nil
### GetDailyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetDailyBudgetCents() int32`

GetDailyBudgetCents returns the DailyBudgetCents field if non-nil, zero value otherwise.

### GetDailyBudgetCentsOk

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetDailyBudgetCentsOk() (*int32, bool)`

GetDailyBudgetCentsOk returns a tuple with the DailyBudgetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetDailyBudgetCents(v int32)`

SetDailyBudgetCents sets DailyBudgetCents field to given value.

### HasDailyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) HasDailyBudgetCents() bool`

HasDailyBudgetCents returns a boolean if a field has been set.

### SetDailyBudgetCentsNil

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetDailyBudgetCentsNil(b bool)`

 SetDailyBudgetCentsNil sets the value for DailyBudgetCents to be an explicit nil

### UnsetDailyBudgetCents
`func (o *GetGovernanceConfig200ResponseSpendLimits) UnsetDailyBudgetCents()`

UnsetDailyBudgetCents ensures that no value is present for DailyBudgetCents, not even an explicit nil
### GetPerUserMonthlyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetPerUserMonthlyBudgetCents() int32`

GetPerUserMonthlyBudgetCents returns the PerUserMonthlyBudgetCents field if non-nil, zero value otherwise.

### GetPerUserMonthlyBudgetCentsOk

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetPerUserMonthlyBudgetCentsOk() (*int32, bool)`

GetPerUserMonthlyBudgetCentsOk returns a tuple with the PerUserMonthlyBudgetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUserMonthlyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetPerUserMonthlyBudgetCents(v int32)`

SetPerUserMonthlyBudgetCents sets PerUserMonthlyBudgetCents field to given value.

### HasPerUserMonthlyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) HasPerUserMonthlyBudgetCents() bool`

HasPerUserMonthlyBudgetCents returns a boolean if a field has been set.

### SetPerUserMonthlyBudgetCentsNil

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetPerUserMonthlyBudgetCentsNil(b bool)`

 SetPerUserMonthlyBudgetCentsNil sets the value for PerUserMonthlyBudgetCents to be an explicit nil

### UnsetPerUserMonthlyBudgetCents
`func (o *GetGovernanceConfig200ResponseSpendLimits) UnsetPerUserMonthlyBudgetCents()`

UnsetPerUserMonthlyBudgetCents ensures that no value is present for PerUserMonthlyBudgetCents, not even an explicit nil
### GetPerUserDailyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetPerUserDailyBudgetCents() int32`

GetPerUserDailyBudgetCents returns the PerUserDailyBudgetCents field if non-nil, zero value otherwise.

### GetPerUserDailyBudgetCentsOk

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetPerUserDailyBudgetCentsOk() (*int32, bool)`

GetPerUserDailyBudgetCentsOk returns a tuple with the PerUserDailyBudgetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUserDailyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetPerUserDailyBudgetCents(v int32)`

SetPerUserDailyBudgetCents sets PerUserDailyBudgetCents field to given value.

### HasPerUserDailyBudgetCents

`func (o *GetGovernanceConfig200ResponseSpendLimits) HasPerUserDailyBudgetCents() bool`

HasPerUserDailyBudgetCents returns a boolean if a field has been set.

### SetPerUserDailyBudgetCentsNil

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetPerUserDailyBudgetCentsNil(b bool)`

 SetPerUserDailyBudgetCentsNil sets the value for PerUserDailyBudgetCents to be an explicit nil

### UnsetPerUserDailyBudgetCents
`func (o *GetGovernanceConfig200ResponseSpendLimits) UnsetPerUserDailyBudgetCents()`

UnsetPerUserDailyBudgetCents ensures that no value is present for PerUserDailyBudgetCents, not even an explicit nil
### GetWarningThresholdPercent

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetWarningThresholdPercent() int32`

GetWarningThresholdPercent returns the WarningThresholdPercent field if non-nil, zero value otherwise.

### GetWarningThresholdPercentOk

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetWarningThresholdPercentOk() (*int32, bool)`

GetWarningThresholdPercentOk returns a tuple with the WarningThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThresholdPercent

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetWarningThresholdPercent(v int32)`

SetWarningThresholdPercent sets WarningThresholdPercent field to given value.

### HasWarningThresholdPercent

`func (o *GetGovernanceConfig200ResponseSpendLimits) HasWarningThresholdPercent() bool`

HasWarningThresholdPercent returns a boolean if a field has been set.

### SetWarningThresholdPercentNil

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetWarningThresholdPercentNil(b bool)`

 SetWarningThresholdPercentNil sets the value for WarningThresholdPercent to be an explicit nil

### UnsetWarningThresholdPercent
`func (o *GetGovernanceConfig200ResponseSpendLimits) UnsetWarningThresholdPercent()`

UnsetWarningThresholdPercent ensures that no value is present for WarningThresholdPercent, not even an explicit nil
### GetInterfaceLimits

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetInterfaceLimits() map[string]GetGovernanceConfig200ResponseSpendLimitsInterfaceLimitsValue`

GetInterfaceLimits returns the InterfaceLimits field if non-nil, zero value otherwise.

### GetInterfaceLimitsOk

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetInterfaceLimitsOk() (*map[string]GetGovernanceConfig200ResponseSpendLimitsInterfaceLimitsValue, bool)`

GetInterfaceLimitsOk returns a tuple with the InterfaceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceLimits

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetInterfaceLimits(v map[string]GetGovernanceConfig200ResponseSpendLimitsInterfaceLimitsValue)`

SetInterfaceLimits sets InterfaceLimits field to given value.

### HasInterfaceLimits

`func (o *GetGovernanceConfig200ResponseSpendLimits) HasInterfaceLimits() bool`

HasInterfaceLimits returns a boolean if a field has been set.

### GetUserOverrides

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetUserOverrides() map[string]GetGovernanceConfig200ResponseSpendLimitsUserOverridesValue`

GetUserOverrides returns the UserOverrides field if non-nil, zero value otherwise.

### GetUserOverridesOk

`func (o *GetGovernanceConfig200ResponseSpendLimits) GetUserOverridesOk() (*map[string]GetGovernanceConfig200ResponseSpendLimitsUserOverridesValue, bool)`

GetUserOverridesOk returns a tuple with the UserOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOverrides

`func (o *GetGovernanceConfig200ResponseSpendLimits) SetUserOverrides(v map[string]GetGovernanceConfig200ResponseSpendLimitsUserOverridesValue)`

SetUserOverrides sets UserOverrides field to given value.

### HasUserOverrides

`func (o *GetGovernanceConfig200ResponseSpendLimits) HasUserOverrides() bool`

HasUserOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


