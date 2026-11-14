# UpdateAutoMode400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Eligible** | Pointer to [**NullableAIAutoModeEligibility**](AIAutoModeEligibility.md) |  | [optional] 

## Methods

### NewUpdateAutoMode400Response

`func NewUpdateAutoMode400Response(error_ string, ) *UpdateAutoMode400Response`

NewUpdateAutoMode400Response instantiates a new UpdateAutoMode400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAutoMode400ResponseWithDefaults

`func NewUpdateAutoMode400ResponseWithDefaults() *UpdateAutoMode400Response`

NewUpdateAutoMode400ResponseWithDefaults instantiates a new UpdateAutoMode400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UpdateAutoMode400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpdateAutoMode400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpdateAutoMode400Response) SetError(v string)`

SetError sets Error field to given value.


### GetEligible

`func (o *UpdateAutoMode400Response) GetEligible() AIAutoModeEligibility`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *UpdateAutoMode400Response) GetEligibleOk() (*AIAutoModeEligibility, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *UpdateAutoMode400Response) SetEligible(v AIAutoModeEligibility)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *UpdateAutoMode400Response) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### SetEligibleNil

`func (o *UpdateAutoMode400Response) SetEligibleNil(b bool)`

 SetEligibleNil sets the value for Eligible to be an explicit nil

### UnsetEligible
`func (o *UpdateAutoMode400Response) UnsetEligible()`

UnsetEligible ensures that no value is present for Eligible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


