# ListVectorCollections200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]ListVectorCollections200ResponseCollectionsInner**](ListVectorCollections200ResponseCollectionsInner.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewListVectorCollections200Response

`func NewListVectorCollections200Response() *ListVectorCollections200Response`

NewListVectorCollections200Response instantiates a new ListVectorCollections200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVectorCollections200ResponseWithDefaults

`func NewListVectorCollections200ResponseWithDefaults() *ListVectorCollections200Response`

NewListVectorCollections200ResponseWithDefaults instantiates a new ListVectorCollections200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *ListVectorCollections200Response) GetCollections() []ListVectorCollections200ResponseCollectionsInner`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ListVectorCollections200Response) GetCollectionsOk() (*[]ListVectorCollections200ResponseCollectionsInner, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ListVectorCollections200Response) SetCollections(v []ListVectorCollections200ResponseCollectionsInner)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *ListVectorCollections200Response) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetCount

`func (o *ListVectorCollections200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListVectorCollections200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListVectorCollections200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListVectorCollections200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


