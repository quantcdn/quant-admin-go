# ImageGenerationRequestImageGenerationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] [default to "standard"]
**CfgScale** | Pointer to **float32** |  | [optional] [default to 6.5]
**Seed** | Pointer to **int32** |  | [optional] 
**NumberOfImages** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewImageGenerationRequestImageGenerationConfig

`func NewImageGenerationRequestImageGenerationConfig() *ImageGenerationRequestImageGenerationConfig`

NewImageGenerationRequestImageGenerationConfig instantiates a new ImageGenerationRequestImageGenerationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageGenerationRequestImageGenerationConfigWithDefaults

`func NewImageGenerationRequestImageGenerationConfigWithDefaults() *ImageGenerationRequestImageGenerationConfig`

NewImageGenerationRequestImageGenerationConfigWithDefaults instantiates a new ImageGenerationRequestImageGenerationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *ImageGenerationRequestImageGenerationConfig) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ImageGenerationRequestImageGenerationConfig) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ImageGenerationRequestImageGenerationConfig) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ImageGenerationRequestImageGenerationConfig) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ImageGenerationRequestImageGenerationConfig) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ImageGenerationRequestImageGenerationConfig) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ImageGenerationRequestImageGenerationConfig) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ImageGenerationRequestImageGenerationConfig) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetQuality

`func (o *ImageGenerationRequestImageGenerationConfig) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ImageGenerationRequestImageGenerationConfig) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ImageGenerationRequestImageGenerationConfig) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ImageGenerationRequestImageGenerationConfig) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetCfgScale

`func (o *ImageGenerationRequestImageGenerationConfig) GetCfgScale() float32`

GetCfgScale returns the CfgScale field if non-nil, zero value otherwise.

### GetCfgScaleOk

`func (o *ImageGenerationRequestImageGenerationConfig) GetCfgScaleOk() (*float32, bool)`

GetCfgScaleOk returns a tuple with the CfgScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfgScale

`func (o *ImageGenerationRequestImageGenerationConfig) SetCfgScale(v float32)`

SetCfgScale sets CfgScale field to given value.

### HasCfgScale

`func (o *ImageGenerationRequestImageGenerationConfig) HasCfgScale() bool`

HasCfgScale returns a boolean if a field has been set.

### GetSeed

`func (o *ImageGenerationRequestImageGenerationConfig) GetSeed() int32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *ImageGenerationRequestImageGenerationConfig) GetSeedOk() (*int32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *ImageGenerationRequestImageGenerationConfig) SetSeed(v int32)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *ImageGenerationRequestImageGenerationConfig) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetNumberOfImages

`func (o *ImageGenerationRequestImageGenerationConfig) GetNumberOfImages() int32`

GetNumberOfImages returns the NumberOfImages field if non-nil, zero value otherwise.

### GetNumberOfImagesOk

`func (o *ImageGenerationRequestImageGenerationConfig) GetNumberOfImagesOk() (*int32, bool)`

GetNumberOfImagesOk returns a tuple with the NumberOfImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfImages

`func (o *ImageGenerationRequestImageGenerationConfig) SetNumberOfImages(v int32)`

SetNumberOfImages sets NumberOfImages field to given value.

### HasNumberOfImages

`func (o *ImageGenerationRequestImageGenerationConfig) HasNumberOfImages() bool`

HasNumberOfImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


