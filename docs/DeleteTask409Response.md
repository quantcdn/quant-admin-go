# DeleteTask409Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**DependentTaskIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeleteTask409Response

`func NewDeleteTask409Response() *DeleteTask409Response`

NewDeleteTask409Response instantiates a new DeleteTask409Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTask409ResponseWithDefaults

`func NewDeleteTask409ResponseWithDefaults() *DeleteTask409Response`

NewDeleteTask409ResponseWithDefaults instantiates a new DeleteTask409Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DeleteTask409Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteTask409Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteTask409Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DeleteTask409Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *DeleteTask409Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteTask409Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteTask409Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeleteTask409Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDependentTaskIds

`func (o *DeleteTask409Response) GetDependentTaskIds() []string`

GetDependentTaskIds returns the DependentTaskIds field if non-nil, zero value otherwise.

### GetDependentTaskIdsOk

`func (o *DeleteTask409Response) GetDependentTaskIdsOk() (*[]string, bool)`

GetDependentTaskIdsOk returns a tuple with the DependentTaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentTaskIds

`func (o *DeleteTask409Response) SetDependentTaskIds(v []string)`

SetDependentTaskIds sets DependentTaskIds field to given value.

### HasDependentTaskIds

`func (o *DeleteTask409Response) HasDependentTaskIds() bool`

HasDependentTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


