# V1Transition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**DateTimestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1Transition

`func NewV1Transition() *V1Transition`

NewV1Transition instantiates a new V1Transition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TransitionWithDefaults

`func NewV1TransitionWithDefaults() *V1Transition`

NewV1TransitionWithDefaults instantiates a new V1Transition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *V1Transition) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V1Transition) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V1Transition) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *V1Transition) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDateTimestamp

`func (o *V1Transition) GetDateTimestamp() int32`

GetDateTimestamp returns the DateTimestamp field if non-nil, zero value otherwise.

### GetDateTimestampOk

`func (o *V1Transition) GetDateTimestampOk() (*int32, bool)`

GetDateTimestampOk returns a tuple with the DateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimestamp

`func (o *V1Transition) SetDateTimestamp(v int32)`

SetDateTimestamp sets DateTimestamp field to given value.

### HasDateTimestamp

`func (o *V1Transition) HasDateTimestamp() bool`

HasDateTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


