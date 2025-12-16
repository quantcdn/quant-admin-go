# GetAIModel200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ContextWindow** | Pointer to **int32** |  | [optional] 
**MaxOutputTokens** | Pointer to **int32** |  | [optional] 
**Capabilities** | Pointer to [**GetAIModel200ResponseCapabilities**](GetAIModel200ResponseCapabilities.md) |  | [optional] 
**Pricing** | Pointer to [**GetAIModel200ResponsePricing**](GetAIModel200ResponsePricing.md) |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**Available** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetAIModel200Response

`func NewGetAIModel200Response() *GetAIModel200Response`

NewGetAIModel200Response instantiates a new GetAIModel200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAIModel200ResponseWithDefaults

`func NewGetAIModel200ResponseWithDefaults() *GetAIModel200Response`

NewGetAIModel200ResponseWithDefaults instantiates a new GetAIModel200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAIModel200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAIModel200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAIModel200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetAIModel200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetAIModel200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAIModel200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAIModel200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAIModel200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *GetAIModel200Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetAIModel200Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetAIModel200Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetAIModel200Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetDescription

`func (o *GetAIModel200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAIModel200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAIModel200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAIModel200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContextWindow

`func (o *GetAIModel200Response) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *GetAIModel200Response) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *GetAIModel200Response) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *GetAIModel200Response) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *GetAIModel200Response) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *GetAIModel200Response) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *GetAIModel200Response) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *GetAIModel200Response) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### GetCapabilities

`func (o *GetAIModel200Response) GetCapabilities() GetAIModel200ResponseCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetAIModel200Response) GetCapabilitiesOk() (*GetAIModel200ResponseCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetAIModel200Response) SetCapabilities(v GetAIModel200ResponseCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *GetAIModel200Response) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetPricing

`func (o *GetAIModel200Response) GetPricing() GetAIModel200ResponsePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *GetAIModel200Response) GetPricingOk() (*GetAIModel200ResponsePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *GetAIModel200Response) SetPricing(v GetAIModel200ResponsePricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *GetAIModel200Response) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetReleaseDate

`func (o *GetAIModel200Response) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *GetAIModel200Response) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *GetAIModel200Response) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *GetAIModel200Response) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetDeprecated

`func (o *GetAIModel200Response) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *GetAIModel200Response) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *GetAIModel200Response) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *GetAIModel200Response) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetAvailable

`func (o *GetAIModel200Response) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetAIModel200Response) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetAIModel200Response) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *GetAIModel200Response) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


