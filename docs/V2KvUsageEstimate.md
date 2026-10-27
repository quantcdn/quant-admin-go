# V2KvUsageEstimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**ReadCost** | **float32** |  | 
**WriteCost** | **float32** |  | 
**RequestCost30Days** | **float32** | readCost + writeCost. | 
**StorageCostPerMonth** | **float32** |  | 
**PitrCostPerMonth** | **float32** |  | 
**Prices** | [**V2KvUsageEstimatePrices**](V2KvUsageEstimatePrices.md) |  | 
**Note** | **string** |  | 

## Methods

### NewV2KvUsageEstimate

`func NewV2KvUsageEstimate(currency string, readCost float32, writeCost float32, requestCost30Days float32, storageCostPerMonth float32, pitrCostPerMonth float32, prices V2KvUsageEstimatePrices, note string, ) *V2KvUsageEstimate`

NewV2KvUsageEstimate instantiates a new V2KvUsageEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvUsageEstimateWithDefaults

`func NewV2KvUsageEstimateWithDefaults() *V2KvUsageEstimate`

NewV2KvUsageEstimateWithDefaults instantiates a new V2KvUsageEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *V2KvUsageEstimate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *V2KvUsageEstimate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *V2KvUsageEstimate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReadCost

`func (o *V2KvUsageEstimate) GetReadCost() float32`

GetReadCost returns the ReadCost field if non-nil, zero value otherwise.

### GetReadCostOk

`func (o *V2KvUsageEstimate) GetReadCostOk() (*float32, bool)`

GetReadCostOk returns a tuple with the ReadCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCost

`func (o *V2KvUsageEstimate) SetReadCost(v float32)`

SetReadCost sets ReadCost field to given value.


### GetWriteCost

`func (o *V2KvUsageEstimate) GetWriteCost() float32`

GetWriteCost returns the WriteCost field if non-nil, zero value otherwise.

### GetWriteCostOk

`func (o *V2KvUsageEstimate) GetWriteCostOk() (*float32, bool)`

GetWriteCostOk returns a tuple with the WriteCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCost

`func (o *V2KvUsageEstimate) SetWriteCost(v float32)`

SetWriteCost sets WriteCost field to given value.


### GetRequestCost30Days

`func (o *V2KvUsageEstimate) GetRequestCost30Days() float32`

GetRequestCost30Days returns the RequestCost30Days field if non-nil, zero value otherwise.

### GetRequestCost30DaysOk

`func (o *V2KvUsageEstimate) GetRequestCost30DaysOk() (*float32, bool)`

GetRequestCost30DaysOk returns a tuple with the RequestCost30Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCost30Days

`func (o *V2KvUsageEstimate) SetRequestCost30Days(v float32)`

SetRequestCost30Days sets RequestCost30Days field to given value.


### GetStorageCostPerMonth

`func (o *V2KvUsageEstimate) GetStorageCostPerMonth() float32`

GetStorageCostPerMonth returns the StorageCostPerMonth field if non-nil, zero value otherwise.

### GetStorageCostPerMonthOk

`func (o *V2KvUsageEstimate) GetStorageCostPerMonthOk() (*float32, bool)`

GetStorageCostPerMonthOk returns a tuple with the StorageCostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCostPerMonth

`func (o *V2KvUsageEstimate) SetStorageCostPerMonth(v float32)`

SetStorageCostPerMonth sets StorageCostPerMonth field to given value.


### GetPitrCostPerMonth

`func (o *V2KvUsageEstimate) GetPitrCostPerMonth() float32`

GetPitrCostPerMonth returns the PitrCostPerMonth field if non-nil, zero value otherwise.

### GetPitrCostPerMonthOk

`func (o *V2KvUsageEstimate) GetPitrCostPerMonthOk() (*float32, bool)`

GetPitrCostPerMonthOk returns a tuple with the PitrCostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrCostPerMonth

`func (o *V2KvUsageEstimate) SetPitrCostPerMonth(v float32)`

SetPitrCostPerMonth sets PitrCostPerMonth field to given value.


### GetPrices

`func (o *V2KvUsageEstimate) GetPrices() V2KvUsageEstimatePrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *V2KvUsageEstimate) GetPricesOk() (*V2KvUsageEstimatePrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *V2KvUsageEstimate) SetPrices(v V2KvUsageEstimatePrices)`

SetPrices sets Prices field to given value.


### GetNote

`func (o *V2KvUsageEstimate) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *V2KvUsageEstimate) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *V2KvUsageEstimate) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


