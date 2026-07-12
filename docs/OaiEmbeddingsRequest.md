# OaiEmbeddingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Input** | **interface{}** | A string or array of strings to embed | 

## Methods

### NewOaiEmbeddingsRequest

`func NewOaiEmbeddingsRequest(model string, input interface{}, ) *OaiEmbeddingsRequest`

NewOaiEmbeddingsRequest instantiates a new OaiEmbeddingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOaiEmbeddingsRequestWithDefaults

`func NewOaiEmbeddingsRequestWithDefaults() *OaiEmbeddingsRequest`

NewOaiEmbeddingsRequestWithDefaults instantiates a new OaiEmbeddingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *OaiEmbeddingsRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OaiEmbeddingsRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OaiEmbeddingsRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *OaiEmbeddingsRequest) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *OaiEmbeddingsRequest) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *OaiEmbeddingsRequest) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *OaiEmbeddingsRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *OaiEmbeddingsRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


