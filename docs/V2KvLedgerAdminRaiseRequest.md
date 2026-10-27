# V2KvLedgerAdminRaiseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chains** | Pointer to **int32** | Not lower than the current value. | [optional] 
**RetentionDays** | Pointer to **int32** | Not lower than the current value. The archive worker extends the retention of the existing archive objects. | [optional] 
**LockMode** | Pointer to **string** | governance can become compliance, never back. | [optional] 

## Methods

### NewV2KvLedgerAdminRaiseRequest

`func NewV2KvLedgerAdminRaiseRequest() *V2KvLedgerAdminRaiseRequest`

NewV2KvLedgerAdminRaiseRequest instantiates a new V2KvLedgerAdminRaiseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerAdminRaiseRequestWithDefaults

`func NewV2KvLedgerAdminRaiseRequestWithDefaults() *V2KvLedgerAdminRaiseRequest`

NewV2KvLedgerAdminRaiseRequestWithDefaults instantiates a new V2KvLedgerAdminRaiseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChains

`func (o *V2KvLedgerAdminRaiseRequest) GetChains() int32`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *V2KvLedgerAdminRaiseRequest) GetChainsOk() (*int32, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *V2KvLedgerAdminRaiseRequest) SetChains(v int32)`

SetChains sets Chains field to given value.

### HasChains

`func (o *V2KvLedgerAdminRaiseRequest) HasChains() bool`

HasChains returns a boolean if a field has been set.

### GetRetentionDays

`func (o *V2KvLedgerAdminRaiseRequest) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *V2KvLedgerAdminRaiseRequest) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *V2KvLedgerAdminRaiseRequest) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *V2KvLedgerAdminRaiseRequest) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetLockMode

`func (o *V2KvLedgerAdminRaiseRequest) GetLockMode() string`

GetLockMode returns the LockMode field if non-nil, zero value otherwise.

### GetLockModeOk

`func (o *V2KvLedgerAdminRaiseRequest) GetLockModeOk() (*string, bool)`

GetLockModeOk returns a tuple with the LockMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockMode

`func (o *V2KvLedgerAdminRaiseRequest) SetLockMode(v string)`

SetLockMode sets LockMode field to given value.

### HasLockMode

`func (o *V2KvLedgerAdminRaiseRequest) HasLockMode() bool`

HasLockMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


