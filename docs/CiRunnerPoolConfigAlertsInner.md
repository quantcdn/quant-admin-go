# CiRunnerPoolConfigAlertsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | pool_error: runner VMs failed to start, VMs reported themselves unhealthy, or a job never reached this pool&#39;s runners (detail says which). labels_unfit: a job asked for labels or an architecture this pool&#39;s runners do not have, so the pool did not take it. | [optional] 
**Since** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewCiRunnerPoolConfigAlertsInner

`func NewCiRunnerPoolConfigAlertsInner() *CiRunnerPoolConfigAlertsInner`

NewCiRunnerPoolConfigAlertsInner instantiates a new CiRunnerPoolConfigAlertsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiRunnerPoolConfigAlertsInnerWithDefaults

`func NewCiRunnerPoolConfigAlertsInnerWithDefaults() *CiRunnerPoolConfigAlertsInner`

NewCiRunnerPoolConfigAlertsInnerWithDefaults instantiates a new CiRunnerPoolConfigAlertsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CiRunnerPoolConfigAlertsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CiRunnerPoolConfigAlertsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CiRunnerPoolConfigAlertsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CiRunnerPoolConfigAlertsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSince

`func (o *CiRunnerPoolConfigAlertsInner) GetSince() string`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *CiRunnerPoolConfigAlertsInner) GetSinceOk() (*string, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *CiRunnerPoolConfigAlertsInner) SetSince(v string)`

SetSince sets Since field to given value.

### HasSince

`func (o *CiRunnerPoolConfigAlertsInner) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetDetail

`func (o *CiRunnerPoolConfigAlertsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CiRunnerPoolConfigAlertsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CiRunnerPoolConfigAlertsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CiRunnerPoolConfigAlertsInner) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


