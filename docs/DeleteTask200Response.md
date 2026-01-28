# DeleteTask200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** |  | [optional] 
**DeletedCount** | Pointer to **int32** | Number of tasks deleted (more than 1 for cascade) | [optional] 
**DeletedTaskIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeleteTask200Response

`func NewDeleteTask200Response() *DeleteTask200Response`

NewDeleteTask200Response instantiates a new DeleteTask200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTask200ResponseWithDefaults

`func NewDeleteTask200ResponseWithDefaults() *DeleteTask200Response`

NewDeleteTask200ResponseWithDefaults instantiates a new DeleteTask200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *DeleteTask200Response) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteTask200Response) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteTask200Response) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteTask200Response) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDeletedCount

`func (o *DeleteTask200Response) GetDeletedCount() int32`

GetDeletedCount returns the DeletedCount field if non-nil, zero value otherwise.

### GetDeletedCountOk

`func (o *DeleteTask200Response) GetDeletedCountOk() (*int32, bool)`

GetDeletedCountOk returns a tuple with the DeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedCount

`func (o *DeleteTask200Response) SetDeletedCount(v int32)`

SetDeletedCount sets DeletedCount field to given value.

### HasDeletedCount

`func (o *DeleteTask200Response) HasDeletedCount() bool`

HasDeletedCount returns a boolean if a field has been set.

### GetDeletedTaskIds

`func (o *DeleteTask200Response) GetDeletedTaskIds() []string`

GetDeletedTaskIds returns the DeletedTaskIds field if non-nil, zero value otherwise.

### GetDeletedTaskIdsOk

`func (o *DeleteTask200Response) GetDeletedTaskIdsOk() (*[]string, bool)`

GetDeletedTaskIdsOk returns a tuple with the DeletedTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedTaskIds

`func (o *DeleteTask200Response) SetDeletedTaskIds(v []string)`

SetDeletedTaskIds sets DeletedTaskIds field to given value.

### HasDeletedTaskIds

`func (o *DeleteTask200Response) HasDeletedTaskIds() bool`

HasDeletedTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


