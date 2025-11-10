# ListAITools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | [**[]ListAITools200ResponseToolsInner**](ListAITools200ResponseToolsInner.md) | Array of available tool specifications | 
**Count** | **int32** | Number of available tools | 

## Methods

### NewListAITools200Response

`func NewListAITools200Response(tools []ListAITools200ResponseToolsInner, count int32, ) *ListAITools200Response`

NewListAITools200Response instantiates a new ListAITools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAITools200ResponseWithDefaults

`func NewListAITools200ResponseWithDefaults() *ListAITools200Response`

NewListAITools200ResponseWithDefaults instantiates a new ListAITools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *ListAITools200Response) GetTools() []ListAITools200ResponseToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ListAITools200Response) GetToolsOk() (*[]ListAITools200ResponseToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ListAITools200Response) SetTools(v []ListAITools200ResponseToolsInner)`

SetTools sets Tools field to given value.


### GetCount

`func (o *ListAITools200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListAITools200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListAITools200Response) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


