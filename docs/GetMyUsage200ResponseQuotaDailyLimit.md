# GetMyUsage200ResponseQuotaDailyLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitCents** | Pointer to **int32** | The configured daily cap in US cents | [optional] 
**UsedPercent** | Pointer to **float32** | Percentage of the cap consumed today (0–100+) | [optional] 
**RemainingCents** | Pointer to **int32** | Cents remaining before the cap is hit; can be negative if overspent | [optional] 
**ResetsAt** | Pointer to **time.Time** | UTC timestamp when the daily counter resets (always next UTC midnight) | [optional] 

## Methods

### NewGetMyUsage200ResponseQuotaDailyLimit

`func NewGetMyUsage200ResponseQuotaDailyLimit() *GetMyUsage200ResponseQuotaDailyLimit`

NewGetMyUsage200ResponseQuotaDailyLimit instantiates a new GetMyUsage200ResponseQuotaDailyLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyUsage200ResponseQuotaDailyLimitWithDefaults

`func NewGetMyUsage200ResponseQuotaDailyLimitWithDefaults() *GetMyUsage200ResponseQuotaDailyLimit`

NewGetMyUsage200ResponseQuotaDailyLimitWithDefaults instantiates a new GetMyUsage200ResponseQuotaDailyLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitCents

`func (o *GetMyUsage200ResponseQuotaDailyLimit) GetLimitCents() int32`

GetLimitCents returns the LimitCents field if non-nil, zero value otherwise.

### GetLimitCentsOk

`func (o *GetMyUsage200ResponseQuotaDailyLimit) GetLimitCentsOk() (*int32, bool)`

GetLimitCentsOk returns a tuple with the LimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCents

`func (o *GetMyUsage200ResponseQuotaDailyLimit) SetLimitCents(v int32)`

SetLimitCents sets LimitCents field to given value.

### HasLimitCents

`func (o *GetMyUsage200ResponseQuotaDailyLimit) HasLimitCents() bool`

HasLimitCents returns a boolean if a field has been set.

### GetUsedPercent

`func (o *GetMyUsage200ResponseQuotaDailyLimit) GetUsedPercent() float32`

GetUsedPercent returns the UsedPercent field if non-nil, zero value otherwise.

### GetUsedPercentOk

`func (o *GetMyUsage200ResponseQuotaDailyLimit) GetUsedPercentOk() (*float32, bool)`

GetUsedPercentOk returns a tuple with the UsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercent

`func (o *GetMyUsage200ResponseQuotaDailyLimit) SetUsedPercent(v float32)`

SetUsedPercent sets UsedPercent field to given value.

### HasUsedPercent

`func (o *GetMyUsage200ResponseQuotaDailyLimit) HasUsedPercent() bool`

HasUsedPercent returns a boolean if a field has been set.

### GetRemainingCents

`func (o *GetMyUsage200ResponseQuotaDailyLimit) GetRemainingCents() int32`

GetRemainingCents returns the RemainingCents field if non-nil, zero value otherwise.

### GetRemainingCentsOk

`func (o *GetMyUsage200ResponseQuotaDailyLimit) GetRemainingCentsOk() (*int32, bool)`

GetRemainingCentsOk returns a tuple with the RemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCents

`func (o *GetMyUsage200ResponseQuotaDailyLimit) SetRemainingCents(v int32)`

SetRemainingCents sets RemainingCents field to given value.

### HasRemainingCents

`func (o *GetMyUsage200ResponseQuotaDailyLimit) HasRemainingCents() bool`

HasRemainingCents returns a boolean if a field has been set.

### GetResetsAt

`func (o *GetMyUsage200ResponseQuotaDailyLimit) GetResetsAt() time.Time`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *GetMyUsage200ResponseQuotaDailyLimit) GetResetsAtOk() (*time.Time, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *GetMyUsage200ResponseQuotaDailyLimit) SetResetsAt(v time.Time)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *GetMyUsage200ResponseQuotaDailyLimit) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


