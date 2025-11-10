# SpotConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strategy** | **string** | Spot instance strategy. &#39;off&#39; &#x3D; On-Demand only (highest reliability, no savings). &#39;spot-only&#39; &#x3D; 100% Spot instances (~70% savings, default for non-prod). &#39;mixed-safe&#39; &#x3D; 50% Spot instances (~35% savings, requires multiple instances). &#39;mixed-aggressive&#39; &#x3D; 80% Spot instances (~56% savings, requires multiple instances). | [default to "spot-only"]

## Methods

### NewSpotConfiguration

`func NewSpotConfiguration(strategy string, ) *SpotConfiguration`

NewSpotConfiguration instantiates a new SpotConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotConfigurationWithDefaults

`func NewSpotConfigurationWithDefaults() *SpotConfiguration`

NewSpotConfigurationWithDefaults instantiates a new SpotConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategy

`func (o *SpotConfiguration) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *SpotConfiguration) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *SpotConfiguration) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


