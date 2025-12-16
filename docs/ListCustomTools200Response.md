# ListCustomTools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to [**[]ListCustomTools200ResponseToolsInner**](ListCustomTools200ResponseToolsInner.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewListCustomTools200Response

`func NewListCustomTools200Response() *ListCustomTools200Response`

NewListCustomTools200Response instantiates a new ListCustomTools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomTools200ResponseWithDefaults

`func NewListCustomTools200ResponseWithDefaults() *ListCustomTools200Response`

NewListCustomTools200ResponseWithDefaults instantiates a new ListCustomTools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *ListCustomTools200Response) GetTools() []ListCustomTools200ResponseToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ListCustomTools200Response) GetToolsOk() (*[]ListCustomTools200ResponseToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ListCustomTools200Response) SetTools(v []ListCustomTools200ResponseToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ListCustomTools200Response) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCount

`func (o *ListCustomTools200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCustomTools200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCustomTools200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCustomTools200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


