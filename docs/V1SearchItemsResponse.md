# V1SearchItemsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]V1SearchHit**](V1SearchHit.md) |  | [optional] 
**NbPages** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1SearchItemsResponse

`func NewV1SearchItemsResponse() *V1SearchItemsResponse`

NewV1SearchItemsResponse instantiates a new V1SearchItemsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SearchItemsResponseWithDefaults

`func NewV1SearchItemsResponseWithDefaults() *V1SearchItemsResponse`

NewV1SearchItemsResponseWithDefaults instantiates a new V1SearchItemsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *V1SearchItemsResponse) GetHits() []V1SearchHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *V1SearchItemsResponse) GetHitsOk() (*[]V1SearchHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *V1SearchItemsResponse) SetHits(v []V1SearchHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *V1SearchItemsResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetNbPages

`func (o *V1SearchItemsResponse) GetNbPages() int32`

GetNbPages returns the NbPages field if non-nil, zero value otherwise.

### GetNbPagesOk

`func (o *V1SearchItemsResponse) GetNbPagesOk() (*int32, bool)`

GetNbPagesOk returns a tuple with the NbPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPages

`func (o *V1SearchItemsResponse) SetNbPages(v int32)`

SetNbPages sets NbPages field to given value.

### HasNbPages

`func (o *V1SearchItemsResponse) HasNbPages() bool`

HasNbPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


