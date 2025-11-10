# ApplicationEnvironmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvName** | Pointer to **string** | Environment name | [optional] 
**Status** | Pointer to **string** | Environment status | [optional] 
**RunningCount** | Pointer to **int32** | Running task count | [optional] 
**DesiredCount** | Pointer to **int32** | Desired task count | [optional] 

## Methods

### NewApplicationEnvironmentsInner

`func NewApplicationEnvironmentsInner() *ApplicationEnvironmentsInner`

NewApplicationEnvironmentsInner instantiates a new ApplicationEnvironmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEnvironmentsInnerWithDefaults

`func NewApplicationEnvironmentsInnerWithDefaults() *ApplicationEnvironmentsInner`

NewApplicationEnvironmentsInnerWithDefaults instantiates a new ApplicationEnvironmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvName

`func (o *ApplicationEnvironmentsInner) GetEnvName() string`

GetEnvName returns the EnvName field if non-nil, zero value otherwise.

### GetEnvNameOk

`func (o *ApplicationEnvironmentsInner) GetEnvNameOk() (*string, bool)`

GetEnvNameOk returns a tuple with the EnvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvName

`func (o *ApplicationEnvironmentsInner) SetEnvName(v string)`

SetEnvName sets EnvName field to given value.

### HasEnvName

`func (o *ApplicationEnvironmentsInner) HasEnvName() bool`

HasEnvName returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationEnvironmentsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationEnvironmentsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationEnvironmentsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationEnvironmentsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRunningCount

`func (o *ApplicationEnvironmentsInner) GetRunningCount() int32`

GetRunningCount returns the RunningCount field if non-nil, zero value otherwise.

### GetRunningCountOk

`func (o *ApplicationEnvironmentsInner) GetRunningCountOk() (*int32, bool)`

GetRunningCountOk returns a tuple with the RunningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningCount

`func (o *ApplicationEnvironmentsInner) SetRunningCount(v int32)`

SetRunningCount sets RunningCount field to given value.

### HasRunningCount

`func (o *ApplicationEnvironmentsInner) HasRunningCount() bool`

HasRunningCount returns a boolean if a field has been set.

### GetDesiredCount

`func (o *ApplicationEnvironmentsInner) GetDesiredCount() int32`

GetDesiredCount returns the DesiredCount field if non-nil, zero value otherwise.

### GetDesiredCountOk

`func (o *ApplicationEnvironmentsInner) GetDesiredCountOk() (*int32, bool)`

GetDesiredCountOk returns a tuple with the DesiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredCount

`func (o *ApplicationEnvironmentsInner) SetDesiredCount(v int32)`

SetDesiredCount sets DesiredCount field to given value.

### HasDesiredCount

`func (o *ApplicationEnvironmentsInner) HasDesiredCount() bool`

HasDesiredCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


