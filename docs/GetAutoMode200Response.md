# GetAutoMode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoMode** | [**NullableAIAutoModeConfig**](AIAutoModeConfig.md) |  | 
**Eligible** | [**NullableAIAutoModeEligibility**](AIAutoModeEligibility.md) |  | 

## Methods

### NewGetAutoMode200Response

`func NewGetAutoMode200Response(autoMode NullableAIAutoModeConfig, eligible NullableAIAutoModeEligibility, ) *GetAutoMode200Response`

NewGetAutoMode200Response instantiates a new GetAutoMode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAutoMode200ResponseWithDefaults

`func NewGetAutoMode200ResponseWithDefaults() *GetAutoMode200Response`

NewGetAutoMode200ResponseWithDefaults instantiates a new GetAutoMode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoMode

`func (o *GetAutoMode200Response) GetAutoMode() AIAutoModeConfig`

GetAutoMode returns the AutoMode field if non-nil, zero value otherwise.

### GetAutoModeOk

`func (o *GetAutoMode200Response) GetAutoModeOk() (*AIAutoModeConfig, bool)`

GetAutoModeOk returns a tuple with the AutoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMode

`func (o *GetAutoMode200Response) SetAutoMode(v AIAutoModeConfig)`

SetAutoMode sets AutoMode field to given value.


### SetAutoModeNil

`func (o *GetAutoMode200Response) SetAutoModeNil(b bool)`

 SetAutoModeNil sets the value for AutoMode to be an explicit nil

### UnsetAutoMode
`func (o *GetAutoMode200Response) UnsetAutoMode()`

UnsetAutoMode ensures that no value is present for AutoMode, not even an explicit nil
### GetEligible

`func (o *GetAutoMode200Response) GetEligible() AIAutoModeEligibility`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *GetAutoMode200Response) GetEligibleOk() (*AIAutoModeEligibility, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *GetAutoMode200Response) SetEligible(v AIAutoModeEligibility)`

SetEligible sets Eligible field to given value.


### SetEligibleNil

`func (o *GetAutoMode200Response) SetEligibleNil(b bool)`

 SetEligibleNil sets the value for Eligible to be an explicit nil

### UnsetEligible
`func (o *GetAutoMode200Response) UnsetEligible()`

UnsetEligible ensures that no value is present for Eligible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


