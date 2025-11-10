# ImageGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | Pointer to **string** | Model to use for image generation | [optional] [default to "amazon.nova-canvas-v1:0"]
**TaskType** | **string** | Type of image generation task | 
**TextToImageParams** | Pointer to [**ImageGenerationRequestTextToImageParams**](ImageGenerationRequestTextToImageParams.md) |  | [optional] 
**ColorGuidedGenerationParams** | Pointer to [**ImageGenerationRequestColorGuidedGenerationParams**](ImageGenerationRequestColorGuidedGenerationParams.md) |  | [optional] 
**ImageVariationParams** | Pointer to [**ImageGenerationRequestImageVariationParams**](ImageGenerationRequestImageVariationParams.md) |  | [optional] 
**InPaintingParams** | Pointer to [**ImageGenerationRequestInPaintingParams**](ImageGenerationRequestInPaintingParams.md) |  | [optional] 
**OutPaintingParams** | Pointer to [**ImageGenerationRequestOutPaintingParams**](ImageGenerationRequestOutPaintingParams.md) |  | [optional] 
**BackgroundRemovalParams** | Pointer to [**ImageGenerationRequestBackgroundRemovalParams**](ImageGenerationRequestBackgroundRemovalParams.md) |  | [optional] 
**VirtualTryOnParams** | Pointer to **map[string]interface{}** | Parameters for VIRTUAL_TRY_ON task | [optional] 
**ImageGenerationConfig** | Pointer to [**ImageGenerationRequestImageGenerationConfig**](ImageGenerationRequestImageGenerationConfig.md) |  | [optional] 
**Region** | Pointer to **string** | AWS region for Nova Canvas | [optional] [default to "us-east-1"]

## Methods

### NewImageGenerationRequest

`func NewImageGenerationRequest(taskType string, ) *ImageGenerationRequest`

NewImageGenerationRequest instantiates a new ImageGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageGenerationRequestWithDefaults

`func NewImageGenerationRequestWithDefaults() *ImageGenerationRequest`

NewImageGenerationRequestWithDefaults instantiates a new ImageGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *ImageGenerationRequest) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ImageGenerationRequest) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ImageGenerationRequest) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ImageGenerationRequest) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetTaskType

`func (o *ImageGenerationRequest) GetTaskType() string`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ImageGenerationRequest) GetTaskTypeOk() (*string, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ImageGenerationRequest) SetTaskType(v string)`

SetTaskType sets TaskType field to given value.


### GetTextToImageParams

`func (o *ImageGenerationRequest) GetTextToImageParams() ImageGenerationRequestTextToImageParams`

GetTextToImageParams returns the TextToImageParams field if non-nil, zero value otherwise.

### GetTextToImageParamsOk

`func (o *ImageGenerationRequest) GetTextToImageParamsOk() (*ImageGenerationRequestTextToImageParams, bool)`

GetTextToImageParamsOk returns a tuple with the TextToImageParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextToImageParams

`func (o *ImageGenerationRequest) SetTextToImageParams(v ImageGenerationRequestTextToImageParams)`

SetTextToImageParams sets TextToImageParams field to given value.

### HasTextToImageParams

`func (o *ImageGenerationRequest) HasTextToImageParams() bool`

HasTextToImageParams returns a boolean if a field has been set.

### GetColorGuidedGenerationParams

`func (o *ImageGenerationRequest) GetColorGuidedGenerationParams() ImageGenerationRequestColorGuidedGenerationParams`

GetColorGuidedGenerationParams returns the ColorGuidedGenerationParams field if non-nil, zero value otherwise.

### GetColorGuidedGenerationParamsOk

`func (o *ImageGenerationRequest) GetColorGuidedGenerationParamsOk() (*ImageGenerationRequestColorGuidedGenerationParams, bool)`

GetColorGuidedGenerationParamsOk returns a tuple with the ColorGuidedGenerationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorGuidedGenerationParams

`func (o *ImageGenerationRequest) SetColorGuidedGenerationParams(v ImageGenerationRequestColorGuidedGenerationParams)`

SetColorGuidedGenerationParams sets ColorGuidedGenerationParams field to given value.

### HasColorGuidedGenerationParams

`func (o *ImageGenerationRequest) HasColorGuidedGenerationParams() bool`

HasColorGuidedGenerationParams returns a boolean if a field has been set.

### GetImageVariationParams

`func (o *ImageGenerationRequest) GetImageVariationParams() ImageGenerationRequestImageVariationParams`

GetImageVariationParams returns the ImageVariationParams field if non-nil, zero value otherwise.

### GetImageVariationParamsOk

`func (o *ImageGenerationRequest) GetImageVariationParamsOk() (*ImageGenerationRequestImageVariationParams, bool)`

GetImageVariationParamsOk returns a tuple with the ImageVariationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVariationParams

`func (o *ImageGenerationRequest) SetImageVariationParams(v ImageGenerationRequestImageVariationParams)`

SetImageVariationParams sets ImageVariationParams field to given value.

### HasImageVariationParams

`func (o *ImageGenerationRequest) HasImageVariationParams() bool`

HasImageVariationParams returns a boolean if a field has been set.

### GetInPaintingParams

`func (o *ImageGenerationRequest) GetInPaintingParams() ImageGenerationRequestInPaintingParams`

GetInPaintingParams returns the InPaintingParams field if non-nil, zero value otherwise.

### GetInPaintingParamsOk

`func (o *ImageGenerationRequest) GetInPaintingParamsOk() (*ImageGenerationRequestInPaintingParams, bool)`

GetInPaintingParamsOk returns a tuple with the InPaintingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInPaintingParams

`func (o *ImageGenerationRequest) SetInPaintingParams(v ImageGenerationRequestInPaintingParams)`

SetInPaintingParams sets InPaintingParams field to given value.

### HasInPaintingParams

`func (o *ImageGenerationRequest) HasInPaintingParams() bool`

HasInPaintingParams returns a boolean if a field has been set.

### GetOutPaintingParams

`func (o *ImageGenerationRequest) GetOutPaintingParams() ImageGenerationRequestOutPaintingParams`

GetOutPaintingParams returns the OutPaintingParams field if non-nil, zero value otherwise.

### GetOutPaintingParamsOk

`func (o *ImageGenerationRequest) GetOutPaintingParamsOk() (*ImageGenerationRequestOutPaintingParams, bool)`

GetOutPaintingParamsOk returns a tuple with the OutPaintingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutPaintingParams

`func (o *ImageGenerationRequest) SetOutPaintingParams(v ImageGenerationRequestOutPaintingParams)`

SetOutPaintingParams sets OutPaintingParams field to given value.

### HasOutPaintingParams

`func (o *ImageGenerationRequest) HasOutPaintingParams() bool`

HasOutPaintingParams returns a boolean if a field has been set.

### GetBackgroundRemovalParams

`func (o *ImageGenerationRequest) GetBackgroundRemovalParams() ImageGenerationRequestBackgroundRemovalParams`

GetBackgroundRemovalParams returns the BackgroundRemovalParams field if non-nil, zero value otherwise.

### GetBackgroundRemovalParamsOk

`func (o *ImageGenerationRequest) GetBackgroundRemovalParamsOk() (*ImageGenerationRequestBackgroundRemovalParams, bool)`

GetBackgroundRemovalParamsOk returns a tuple with the BackgroundRemovalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundRemovalParams

`func (o *ImageGenerationRequest) SetBackgroundRemovalParams(v ImageGenerationRequestBackgroundRemovalParams)`

SetBackgroundRemovalParams sets BackgroundRemovalParams field to given value.

### HasBackgroundRemovalParams

`func (o *ImageGenerationRequest) HasBackgroundRemovalParams() bool`

HasBackgroundRemovalParams returns a boolean if a field has been set.

### GetVirtualTryOnParams

`func (o *ImageGenerationRequest) GetVirtualTryOnParams() map[string]interface{}`

GetVirtualTryOnParams returns the VirtualTryOnParams field if non-nil, zero value otherwise.

### GetVirtualTryOnParamsOk

`func (o *ImageGenerationRequest) GetVirtualTryOnParamsOk() (*map[string]interface{}, bool)`

GetVirtualTryOnParamsOk returns a tuple with the VirtualTryOnParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualTryOnParams

`func (o *ImageGenerationRequest) SetVirtualTryOnParams(v map[string]interface{})`

SetVirtualTryOnParams sets VirtualTryOnParams field to given value.

### HasVirtualTryOnParams

`func (o *ImageGenerationRequest) HasVirtualTryOnParams() bool`

HasVirtualTryOnParams returns a boolean if a field has been set.

### GetImageGenerationConfig

`func (o *ImageGenerationRequest) GetImageGenerationConfig() ImageGenerationRequestImageGenerationConfig`

GetImageGenerationConfig returns the ImageGenerationConfig field if non-nil, zero value otherwise.

### GetImageGenerationConfigOk

`func (o *ImageGenerationRequest) GetImageGenerationConfigOk() (*ImageGenerationRequestImageGenerationConfig, bool)`

GetImageGenerationConfigOk returns a tuple with the ImageGenerationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageGenerationConfig

`func (o *ImageGenerationRequest) SetImageGenerationConfig(v ImageGenerationRequestImageGenerationConfig)`

SetImageGenerationConfig sets ImageGenerationConfig field to given value.

### HasImageGenerationConfig

`func (o *ImageGenerationRequest) HasImageGenerationConfig() bool`

HasImageGenerationConfig returns a boolean if a field has been set.

### GetRegion

`func (o *ImageGenerationRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ImageGenerationRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ImageGenerationRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ImageGenerationRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


