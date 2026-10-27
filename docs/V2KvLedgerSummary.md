# V2KvLedgerSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveLagSeconds** | **int32** |  | 
**Broken** | **bool** |  | 
**RetainUntil** | Pointer to **time.Time** |  | [optional] 
**AnchorFailingSince** | Pointer to **time.Time** |  | [optional] 
**LastVerify** | Pointer to [**V2KvLedgerLastVerify**](V2KvLedgerLastVerify.md) |  | [optional] 
**At** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewV2KvLedgerSummary

`func NewV2KvLedgerSummary(archiveLagSeconds int32, broken bool, ) *V2KvLedgerSummary`

NewV2KvLedgerSummary instantiates a new V2KvLedgerSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2KvLedgerSummaryWithDefaults

`func NewV2KvLedgerSummaryWithDefaults() *V2KvLedgerSummary`

NewV2KvLedgerSummaryWithDefaults instantiates a new V2KvLedgerSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveLagSeconds

`func (o *V2KvLedgerSummary) GetArchiveLagSeconds() int32`

GetArchiveLagSeconds returns the ArchiveLagSeconds field if non-nil, zero value otherwise.

### GetArchiveLagSecondsOk

`func (o *V2KvLedgerSummary) GetArchiveLagSecondsOk() (*int32, bool)`

GetArchiveLagSecondsOk returns a tuple with the ArchiveLagSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLagSeconds

`func (o *V2KvLedgerSummary) SetArchiveLagSeconds(v int32)`

SetArchiveLagSeconds sets ArchiveLagSeconds field to given value.


### GetBroken

`func (o *V2KvLedgerSummary) GetBroken() bool`

GetBroken returns the Broken field if non-nil, zero value otherwise.

### GetBrokenOk

`func (o *V2KvLedgerSummary) GetBrokenOk() (*bool, bool)`

GetBrokenOk returns a tuple with the Broken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroken

`func (o *V2KvLedgerSummary) SetBroken(v bool)`

SetBroken sets Broken field to given value.


### GetRetainUntil

`func (o *V2KvLedgerSummary) GetRetainUntil() time.Time`

GetRetainUntil returns the RetainUntil field if non-nil, zero value otherwise.

### GetRetainUntilOk

`func (o *V2KvLedgerSummary) GetRetainUntilOk() (*time.Time, bool)`

GetRetainUntilOk returns a tuple with the RetainUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntil

`func (o *V2KvLedgerSummary) SetRetainUntil(v time.Time)`

SetRetainUntil sets RetainUntil field to given value.

### HasRetainUntil

`func (o *V2KvLedgerSummary) HasRetainUntil() bool`

HasRetainUntil returns a boolean if a field has been set.

### GetAnchorFailingSince

`func (o *V2KvLedgerSummary) GetAnchorFailingSince() time.Time`

GetAnchorFailingSince returns the AnchorFailingSince field if non-nil, zero value otherwise.

### GetAnchorFailingSinceOk

`func (o *V2KvLedgerSummary) GetAnchorFailingSinceOk() (*time.Time, bool)`

GetAnchorFailingSinceOk returns a tuple with the AnchorFailingSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorFailingSince

`func (o *V2KvLedgerSummary) SetAnchorFailingSince(v time.Time)`

SetAnchorFailingSince sets AnchorFailingSince field to given value.

### HasAnchorFailingSince

`func (o *V2KvLedgerSummary) HasAnchorFailingSince() bool`

HasAnchorFailingSince returns a boolean if a field has been set.

### GetLastVerify

`func (o *V2KvLedgerSummary) GetLastVerify() V2KvLedgerLastVerify`

GetLastVerify returns the LastVerify field if non-nil, zero value otherwise.

### GetLastVerifyOk

`func (o *V2KvLedgerSummary) GetLastVerifyOk() (*V2KvLedgerLastVerify, bool)`

GetLastVerifyOk returns a tuple with the LastVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerify

`func (o *V2KvLedgerSummary) SetLastVerify(v V2KvLedgerLastVerify)`

SetLastVerify sets LastVerify field to given value.

### HasLastVerify

`func (o *V2KvLedgerSummary) HasLastVerify() bool`

HasLastVerify returns a boolean if a field has been set.

### GetAt

`func (o *V2KvLedgerSummary) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *V2KvLedgerSummary) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *V2KvLedgerSummary) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *V2KvLedgerSummary) HasAt() bool`

HasAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


