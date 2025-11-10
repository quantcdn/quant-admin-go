# ImageGenerationRequestTextToImageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Text prompt | [optional] 
**NegativeText** | Pointer to **string** | What NOT to include | [optional] 
**Style** | Pointer to **string** |  | [optional] 
**ConditionImage** | Pointer to **string** | Base64-encoded conditioning image | [optional] 
**ControlMode** | Pointer to **string** |  | [optional] [default to "CANNY_EDGE"]
**ControlStrength** | Pointer to **float32** |  | [optional] [default to 0.7]

## Methods

### NewImageGenerationRequestTextToImageParams

`func NewImageGenerationRequestTextToImageParams() *ImageGenerationRequestTextToImageParams`

NewImageGenerationRequestTextToImageParams instantiates a new ImageGenerationRequestTextToImageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageGenerationRequestTextToImageParamsWithDefaults

`func NewImageGenerationRequestTextToImageParamsWithDefaults() *ImageGenerationRequestTextToImageParams`

NewImageGenerationRequestTextToImageParamsWithDefaults instantiates a new ImageGenerationRequestTextToImageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ImageGenerationRequestTextToImageParams) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ImageGenerationRequestTextToImageParams) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ImageGenerationRequestTextToImageParams) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ImageGenerationRequestTextToImageParams) HasText() bool`

HasText returns a boolean if a field has been set.

### GetNegativeText

`func (o *ImageGenerationRequestTextToImageParams) GetNegativeText() string`

GetNegativeText returns the NegativeText field if non-nil, zero value otherwise.

### GetNegativeTextOk

`func (o *ImageGenerationRequestTextToImageParams) GetNegativeTextOk() (*string, bool)`

GetNegativeTextOk returns a tuple with the NegativeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeText

`func (o *ImageGenerationRequestTextToImageParams) SetNegativeText(v string)`

SetNegativeText sets NegativeText field to given value.

### HasNegativeText

`func (o *ImageGenerationRequestTextToImageParams) HasNegativeText() bool`

HasNegativeText returns a boolean if a field has been set.

### GetStyle

`func (o *ImageGenerationRequestTextToImageParams) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ImageGenerationRequestTextToImageParams) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ImageGenerationRequestTextToImageParams) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ImageGenerationRequestTextToImageParams) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetConditionImage

`func (o *ImageGenerationRequestTextToImageParams) GetConditionImage() string`

GetConditionImage returns the ConditionImage field if non-nil, zero value otherwise.

### GetConditionImageOk

`func (o *ImageGenerationRequestTextToImageParams) GetConditionImageOk() (*string, bool)`

GetConditionImageOk returns a tuple with the ConditionImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionImage

`func (o *ImageGenerationRequestTextToImageParams) SetConditionImage(v string)`

SetConditionImage sets ConditionImage field to given value.

### HasConditionImage

`func (o *ImageGenerationRequestTextToImageParams) HasConditionImage() bool`

HasConditionImage returns a boolean if a field has been set.

### GetControlMode

`func (o *ImageGenerationRequestTextToImageParams) GetControlMode() string`

GetControlMode returns the ControlMode field if non-nil, zero value otherwise.

### GetControlModeOk

`func (o *ImageGenerationRequestTextToImageParams) GetControlModeOk() (*string, bool)`

GetControlModeOk returns a tuple with the ControlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlMode

`func (o *ImageGenerationRequestTextToImageParams) SetControlMode(v string)`

SetControlMode sets ControlMode field to given value.

### HasControlMode

`func (o *ImageGenerationRequestTextToImageParams) HasControlMode() bool`

HasControlMode returns a boolean if a field has been set.

### GetControlStrength

`func (o *ImageGenerationRequestTextToImageParams) GetControlStrength() float32`

GetControlStrength returns the ControlStrength field if non-nil, zero value otherwise.

### GetControlStrengthOk

`func (o *ImageGenerationRequestTextToImageParams) GetControlStrengthOk() (*float32, bool)`

GetControlStrengthOk returns a tuple with the ControlStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlStrength

`func (o *ImageGenerationRequestTextToImageParams) SetControlStrength(v float32)`

SetControlStrength sets ControlStrength field to given value.

### HasControlStrength

`func (o *ImageGenerationRequestTextToImageParams) HasControlStrength() bool`

HasControlStrength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


