# AiSearchSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Limit** | Pointer to **int32** |  | [optional] 
**MinScore** | Pointer to **float32** |  | [optional] 

## Methods

### NewAiSearchSearchRequest

`func NewAiSearchSearchRequest(query string, ) *AiSearchSearchRequest`

NewAiSearchSearchRequest instantiates a new AiSearchSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSearchSearchRequestWithDefaults

`func NewAiSearchSearchRequestWithDefaults() *AiSearchSearchRequest`

NewAiSearchSearchRequestWithDefaults instantiates a new AiSearchSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *AiSearchSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AiSearchSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AiSearchSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetLimit

`func (o *AiSearchSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AiSearchSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AiSearchSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AiSearchSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMinScore

`func (o *AiSearchSearchRequest) GetMinScore() float32`

GetMinScore returns the MinScore field if non-nil, zero value otherwise.

### GetMinScoreOk

`func (o *AiSearchSearchRequest) GetMinScoreOk() (*float32, bool)`

GetMinScoreOk returns a tuple with the MinScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScore

`func (o *AiSearchSearchRequest) SetMinScore(v float32)`

SetMinScore sets MinScore field to given value.

### HasMinScore

`func (o *AiSearchSearchRequest) HasMinScore() bool`

HasMinScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


