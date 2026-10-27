# V2KvUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreName** | **string** |  | 
**Format** | **int32** |  | 
**State** | **string** |  | 
**Table** | [**V2KvUsageTable**](V2KvUsageTable.md) |  | 
**RetainedTable** | Pointer to [**V2KvUsageRetainedTable**](V2KvUsageRetainedTable.md) |  | [optional] 
**Period** | [**V2KvUsagePeriod**](V2KvUsagePeriod.md) |  | 
**Daily** | [**[]V2KvUsageDailyInner**](V2KvUsageDailyInner.md) | Exactly period.days rows, oldest first. Units are rounded to 1 decimal. | 
**Totals** | [**V2KvUsageTotals**](V2KvUsageTotals.md) |  | 
**Estimate** | [**V2KvUsageEstimate**](V2KvUsageEstimate.md) |  | 

## Methods

### NewV2KvUsage

`func NewV2KvUsage(storeName string, format int32, state string, table V2KvUsageTable, period V2KvUsagePeriod, daily []V2KvUsageDailyInner, totals V2KvUsageTotals, estimate V2KvUsageEstimate, ) *V2KvUsage`

NewV2KvUsage instantiates a new V2KvUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvUsageWithDefaults

`func NewV2KvUsageWithDefaults() *V2KvUsage`

NewV2KvUsageWithDefaults instantiates a new V2KvUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreName

`func (o *V2KvUsage) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *V2KvUsage) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *V2KvUsage) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.


### GetFormat

`func (o *V2KvUsage) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V2KvUsage) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V2KvUsage) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetState

`func (o *V2KvUsage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V2KvUsage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V2KvUsage) SetState(v string)`

SetState sets State field to given value.


### GetTable

`func (o *V2KvUsage) GetTable() V2KvUsageTable`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *V2KvUsage) GetTableOk() (*V2KvUsageTable, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *V2KvUsage) SetTable(v V2KvUsageTable)`

SetTable sets Table field to given value.


### GetRetainedTable

`func (o *V2KvUsage) GetRetainedTable() V2KvUsageRetainedTable`

GetRetainedTable returns the RetainedTable field if non-nil, zero value otherwise.

### GetRetainedTableOk

`func (o *V2KvUsage) GetRetainedTableOk() (*V2KvUsageRetainedTable, bool)`

GetRetainedTableOk returns a tuple with the RetainedTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedTable

`func (o *V2KvUsage) SetRetainedTable(v V2KvUsageRetainedTable)`

SetRetainedTable sets RetainedTable field to given value.

### HasRetainedTable

`func (o *V2KvUsage) HasRetainedTable() bool`

HasRetainedTable returns a boolean if a field has been set.

### GetPeriod

`func (o *V2KvUsage) GetPeriod() V2KvUsagePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *V2KvUsage) GetPeriodOk() (*V2KvUsagePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *V2KvUsage) SetPeriod(v V2KvUsagePeriod)`

SetPeriod sets Period field to given value.


### GetDaily

`func (o *V2KvUsage) GetDaily() []V2KvUsageDailyInner`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *V2KvUsage) GetDailyOk() (*[]V2KvUsageDailyInner, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *V2KvUsage) SetDaily(v []V2KvUsageDailyInner)`

SetDaily sets Daily field to given value.


### GetTotals

`func (o *V2KvUsage) GetTotals() V2KvUsageTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *V2KvUsage) GetTotalsOk() (*V2KvUsageTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *V2KvUsage) SetTotals(v V2KvUsageTotals)`

SetTotals sets Totals field to given value.


### GetEstimate

`func (o *V2KvUsage) GetEstimate() V2KvUsageEstimate`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *V2KvUsage) GetEstimateOk() (*V2KvUsageEstimate, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *V2KvUsage) SetEstimate(v V2KvUsageEstimate)`

SetEstimate sets Estimate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


