# ExtendAISessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMinutes** | Pointer to **int32** | Minutes to add to expiration time (default 60, max 1440) | [optional] [default to 60]

## Methods

### NewExtendAISessionRequest

`func NewExtendAISessionRequest() *ExtendAISessionRequest`

NewExtendAISessionRequest instantiates a new ExtendAISessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendAISessionRequestWithDefaults

`func NewExtendAISessionRequestWithDefaults() *ExtendAISessionRequest`

NewExtendAISessionRequestWithDefaults instantiates a new ExtendAISessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMinutes

`func (o *ExtendAISessionRequest) GetAdditionalMinutes() int32`

GetAdditionalMinutes returns the AdditionalMinutes field if non-nil, zero value otherwise.

### GetAdditionalMinutesOk

`func (o *ExtendAISessionRequest) GetAdditionalMinutesOk() (*int32, bool)`

GetAdditionalMinutesOk returns a tuple with the AdditionalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMinutes

`func (o *ExtendAISessionRequest) SetAdditionalMinutes(v int32)`

SetAdditionalMinutes sets AdditionalMinutes field to given value.

### HasAdditionalMinutes

`func (o *ExtendAISessionRequest) HasAdditionalMinutes() bool`

HasAdditionalMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


