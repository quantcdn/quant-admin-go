# GetMyUsage200ResponseQuotaMonthlyLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitCents** | Pointer to **int32** | The configured monthly cap in US cents | [optional] 
**UsedPercent** | Pointer to **float32** | Percentage of the cap consumed this month (0–100+) | [optional] 
**RemainingCents** | Pointer to **int32** | Cents remaining before the cap is hit; can be negative if overspent | [optional] 

## Methods

### NewGetMyUsage200ResponseQuotaMonthlyLimit

`func NewGetMyUsage200ResponseQuotaMonthlyLimit() *GetMyUsage200ResponseQuotaMonthlyLimit`

NewGetMyUsage200ResponseQuotaMonthlyLimit instantiates a new GetMyUsage200ResponseQuotaMonthlyLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyUsage200ResponseQuotaMonthlyLimitWithDefaults

`func NewGetMyUsage200ResponseQuotaMonthlyLimitWithDefaults() *GetMyUsage200ResponseQuotaMonthlyLimit`

NewGetMyUsage200ResponseQuotaMonthlyLimitWithDefaults instantiates a new GetMyUsage200ResponseQuotaMonthlyLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitCents

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) GetLimitCents() int32`

GetLimitCents returns the LimitCents field if non-nil, zero value otherwise.

### GetLimitCentsOk

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) GetLimitCentsOk() (*int32, bool)`

GetLimitCentsOk returns a tuple with the LimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCents

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) SetLimitCents(v int32)`

SetLimitCents sets LimitCents field to given value.

### HasLimitCents

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) HasLimitCents() bool`

HasLimitCents returns a boolean if a field has been set.

### GetUsedPercent

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) GetUsedPercent() float32`

GetUsedPercent returns the UsedPercent field if non-nil, zero value otherwise.

### GetUsedPercentOk

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) GetUsedPercentOk() (*float32, bool)`

GetUsedPercentOk returns a tuple with the UsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercent

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) SetUsedPercent(v float32)`

SetUsedPercent sets UsedPercent field to given value.

### HasUsedPercent

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) HasUsedPercent() bool`

HasUsedPercent returns a boolean if a field has been set.

### GetRemainingCents

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) GetRemainingCents() int32`

GetRemainingCents returns the RemainingCents field if non-nil, zero value otherwise.

### GetRemainingCentsOk

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) GetRemainingCentsOk() (*int32, bool)`

GetRemainingCentsOk returns a tuple with the RemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCents

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) SetRemainingCents(v int32)`

SetRemainingCents sets RemainingCents field to given value.

### HasRemainingCents

`func (o *GetMyUsage200ResponseQuotaMonthlyLimit) HasRemainingCents() bool`

HasRemainingCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


