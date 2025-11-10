# ImageGeneration200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | **[]string** | Array of base64-encoded generated images | 
**MaskImage** | Pointer to **string** | Base64-encoded mask image (for virtual try-on) | [optional] 
**Error** | Pointer to **string** | Error message if any images were blocked by content moderation | [optional] 

## Methods

### NewImageGeneration200Response

`func NewImageGeneration200Response(images []string, ) *ImageGeneration200Response`

NewImageGeneration200Response instantiates a new ImageGeneration200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageGeneration200ResponseWithDefaults

`func NewImageGeneration200ResponseWithDefaults() *ImageGeneration200Response`

NewImageGeneration200ResponseWithDefaults instantiates a new ImageGeneration200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *ImageGeneration200Response) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ImageGeneration200Response) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ImageGeneration200Response) SetImages(v []string)`

SetImages sets Images field to given value.


### GetMaskImage

`func (o *ImageGeneration200Response) GetMaskImage() string`

GetMaskImage returns the MaskImage field if non-nil, zero value otherwise.

### GetMaskImageOk

`func (o *ImageGeneration200Response) GetMaskImageOk() (*string, bool)`

GetMaskImageOk returns a tuple with the MaskImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskImage

`func (o *ImageGeneration200Response) SetMaskImage(v string)`

SetMaskImage sets MaskImage field to given value.

### HasMaskImage

`func (o *ImageGeneration200Response) HasMaskImage() bool`

HasMaskImage returns a boolean if a field has been set.

### GetError

`func (o *ImageGeneration200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ImageGeneration200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ImageGeneration200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ImageGeneration200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


